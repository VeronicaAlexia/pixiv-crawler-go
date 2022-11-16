package app

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/config"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/file"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/input"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/request"
	"github.com/VeronicaAlexia/pixiv-crawler-go/src/download"
	"github.com/VeronicaAlexia/pixiv-crawler-go/src/pixiv"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils"
)

var App = pixiv.NewApp()

func DownloaderSingly(illust_id string) {
	var urls []string
	if utils.ListFind(file.ShowFileList("./imageFile"), illust_id) {
		fmt.Println(illust_id, "is exist, skip")
	} else {
		illust, err := App.IllustDetail(illust_id)
		if err != nil {
			fmt.Println("download fail", err)
			return
		}
		if illust == nil || illust.MetaSinglePage == nil {
			fmt.Println("download fail, illust is nil")
			return
		}
		if illust.MetaSinglePage.OriginalImageURL == "" {
			for _, img := range illust.MetaPages {
				urls = append(urls, img.Images.Original)
			}
		} else {
			urls = append(urls, illust.MetaSinglePage.OriginalImageURL)
		}
		for index, url := range urls {
			fmt.Println("download", illust.Title, "\timage", index+1, "of", len(urls))
			download.ImagesSingly(url, nil)
		}
		fmt.Println("\033[2J")
	}
}

func ShellUserFollowing(UserID int) {
	if UserID == 0 {
		UserID = config.Vars.UserID
	}
	following, err := App.UserFollowing(UserID, "public", 0)
	if err != nil {
		fmt.Println("Request user following fail,please check network", err)
	}
	for index, user := range following.UserPreviews {
		fmt.Println("index:", index, "\tuser_id:", user.User.ID, "\tuser_name:", user.User.Name)
	}
	fmt.Println("一共", len(following.UserPreviews), "个关注的用户")
	for _, user := range following.UserPreviews {
		ShellAuthor("", user.User.ID)
	}
	// 刷新屏幕
}

func ShellStarsImages(user_id int, next_url string) {
	bookmarks, err := App.UserBookmarksIllust(user_id, next_url)
	if err != nil {
		fmt.Println("Request user bookmarks illust fail,please check network", err)
	} else {
		download.DownloadTask(bookmarks.Illusts, true)
		if bookmarks.NextURL != "" {
			ShellStarsImages(user_id, bookmarks.NextURL)
		}
	}
}

func ShellRanking() {
	RankingMode := []string{"day", "week", "month", "day_male", "day_female", "week_original", "week_rookie", "day_manga"}
	for index, mode := range RankingMode {
		fmt.Println("index:", index, "\tmodel:", mode)
	}
	illusts, err := App.IllustRanking(RankingMode[input.OutputInt(">", ">", len(RankingMode))])
	if err != nil {
		fmt.Println("Ranking request fail,please check network", err)
	} else {
		download.DownloadTask(illusts.Illusts, true)
	}
}

func ShellRecommend(next_url string, auth bool) {
	if recommended, err := App.Recommended(next_url, auth); err != nil {
		fmt.Println("Recommended request fail,please check network", err)
	} else {
		download.DownloadTask(recommended.Illusts, true)
		if recommended.NextURL != "" {
			ShellRecommend(recommended.NextURL, auth)
		}
	}
}

func ShellAuthor(next_url string, author_id int) {
	if illusts, err := App.UserIllusts(author_id, next_url); err == nil {
		download.DownloadTask(illusts.Illusts, true)
		if illusts.NextURL != "" { // If there is a next page, continue to request
			ShellAuthor(illusts.NextURL, author_id)
		}
	} else {
		fmt.Println("Request author info fail,please check network", err)
	}
}

func ShellLoginPixiv() {
	if accessToken, err := request.ChromeDriverLogin(); err != nil {
		fmt.Println("Login fail,please check network", err)
	} else {
		config.VarsFile.Vipers.Set("pixiv_refresh_token", accessToken.RefreshToken)
		config.VarsFile.Vipers.Set("pixiv_token", accessToken.AccessToken)
		config.VarsFile.Vipers.Set("PIXIV_USER_ID", accessToken.User.ID)
		config.VarsFile.SaveConfig()
	}
}

func ShellNovel(novel_id string) {
	file.NewFile("novel")
	response, err := App.NovelDetail(novel_id)
	if err != nil {
		fmt.Println("Request novel fail,please check network", err)
	}
	if chapter_content, ok := App.NovelContent(novel_id); ok == nil {
		book_info := "title: " + response.Novel.Title + "\n"
		book_info += "author: " + chapter_content["author_namer"] + "\n"
		book_info += "novel_id: " + chapter_content["novel_id"] + "\n"
		book_info += "update: " + chapter_content["update_date"] + "\n"
		book_info += "intro: " + chapter_content["description"] + "\n"
		book_info += "tags: "
		for index, tag := range response.Novel.Tags {
			book_info += tag.Name
			if index != 0 {
				book_info += ", "
			}
		}
		file_path := "novel/" + response.Novel.Title + ".txt"
		file.Open(file_path, "w", book_info+"\n\n"+chapter_content["content_text"])
	} else {
		fmt.Println("Request novel content fail,please check network", err)
	}
}
