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
	novel_id := "18729784"
	init_test()
	file.NewFile("novel")
	if response, err := app.App.NovelDetail(novel_id); err != nil {
		t.Error(err)
	} else {
		if chapter_content, ok := app.App.NovelContent(novel_id); ok == nil {
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
			t.Error(ok)
		}
	}
}
