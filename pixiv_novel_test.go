package main

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/config"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/file"
	"github.com/VeronicaAlexia/pixiv-crawler-go/src/app"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils"
	"path"
	"testing"
)

func init() {
	config.VarsConfigInit()
	if config.Vars.PixivRefreshToken == "" {
		app.ShellLoginPixiv()
	}
	if res, err := app.App.UserDetail(config.Vars.UserID); err != nil {
		panic("UserDetail error: " + err.Error())
	} else {
		fmt.Println("account: ", res.User.Name, "\tid: ", res.User.ID)
	}
}

func TestPixivNovel(t *testing.T) {
	novel_id := "18729784"
	file.NewFile("novel")
	response, err := app.App.NovelDetail(novel_id)
	if err != nil {
		t.Error(err)

	}
	if chapter_content, ok := app.App.NovelContent(novel_id); ok == nil {
		book_info := utils.MakeBookInfo(response)
		file.Open(path.Join("novel", response.Novel.Title+".txt"), "w", book_info+"\n\n"+chapter_content)

	} else {
		t.Error(ok)
	}
}
