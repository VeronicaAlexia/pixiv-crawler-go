package main

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/config"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/file"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/request"
	"github.com/VeronicaAlexia/pixiv-crawler-go/src/app"
	"testing"
)

func init_test() {
	config.VarsConfigInit()
	if config.Vars.PixivRefreshToken == "" {
		if accessToken, err := request.ChromeDriverLogin(); err != nil {
			panic(err)
		} else {
			config.VarsFile.Vipers.Set("pixiv_refresh_token", accessToken.RefreshToken)
			config.VarsFile.Vipers.Set("pixiv_token", accessToken.AccessToken)
			config.VarsFile.Vipers.Set("PIXIV_USER_ID", accessToken.User.ID)
			config.VarsFile.SaveConfig()
		}
	} else {
		if res, err := app.App.UserDetail(config.Vars.UserID); err != nil {
			panic(err)
		} else {
			fmt.Println("account: ", res.User.Name, "\tid: ", res.User.ID)
		}
	}
}

func TestPixivNovel(t *testing.T) {
	init_test()
	file.NewFile("novel")
	if response, err := app.App.NovelDetail("18729784"); err != nil {
		t.Error(err)
	} else {
		book_info := "title: " + response.Novel.Title + "\n"
		book_info += "author: " + response.Novel.User.Name + "\n"
		book_info += "novel_id: " + fmt.Sprint(response.Novel.ID) + "\n"
		book_info += "update: " + response.Novel.CreateDate.Format("2006-01-02 15:04:05") + "\n"
		book_info += "intro: " + response.Novel.Caption + "\n"
		book_info += "tags: "
		for _, tag := range response.Novel.Tags {
			book_info += tag.Name + ", "
		}
		book_info += "\n"
		file.Open("novel/"+response.Novel.Title+".txt", "w", book_info+"\n"+app.App.NovelContent("18729784"))

	}
}
