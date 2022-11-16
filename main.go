package main

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/arguments"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/cli"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/config"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/file"
	"github.com/VeronicaAlexia/pixiv-crawler-go/src/app"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils"
	"log"
	"os"
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

func main() {
	file.NewFile("imageFile")
	cli_app := cli.NewApp()
	cli_app.Name = "image downloader"
	cli_app.Version = config.Vars.VersionName
	cli_app.Usage = "download image from pixiv "
	cli_app.Flags = arguments.CommandLineFlag
	cli_app.Action = func(c *cli.Context) {
		config.Vars.ThreadMax = arguments.CommandLines.ThreadMax
		shell(arguments.CommandLines, c)
	}
	if err := cli_app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func shell(args arguments.Command, c *cli.Context) {
	if args.IllustID != "" || args.URL != "" {
		var IllustID string
		if args.IllustID != "" {
			IllustID = args.IllustID
		} else {
			IllustID = utils.GetInt(args.URL)
		}
		if args.IsNovel {
			app.ShellNovel(IllustID)
		} else {
			app.DownloaderSingly(IllustID)
		}

	} else if args.AuthorID != "" {
		app.ShellAuthor("", args.AuthorID)

	} else if args.Following {
		app.ShellUserFollowing(args.UserID)

	} else if args.Recommend {
		app.ShellRecommend("", true)

	} else if args.Stars {
		app.ShellStarsImages(config.Vars.UserID, "")

	} else if args.Ranking {
		app.ShellRanking()

	} else if args.Login {
		app.ShellLoginPixiv()

	} else {
		if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
			_ = cli.ShowAppHelp(c)
		}

	}
}
