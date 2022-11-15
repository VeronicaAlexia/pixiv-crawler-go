package main

import (
	"fmt"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/arguments"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/cli"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/config"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/file"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/request"
	"github.com/VeronicaAlexia/pixiv-crawler-go/src/app"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils"
	"log"
	"os"
)

func init() {
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
		res, err := app.App.UserDetail(config.Vars.UserID)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("account: ", res.User.Name, "\tid: ", res.User.ID)
		}
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
		config.Vars.ThreadMax = arguments.CommandLines.Max
		if !command_line_shell(arguments.CommandLines) {
			if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
				_ = cli.ShowAppHelp(c)
			}
		}
	}
	if err := cli_app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
func command_line_shell(args arguments.Command) bool {
	if args.IllustID != "" {
		app.DownloaderSingly(args.IllustID)

	} else if args.AuthorID != 0 {
		app.ShellAuthor("", args.AuthorID)

	} else if args.URL != "" {
		app.DownloaderSingly(utils.GetInt(args.URL))

	} else if args.Following {
		app.GET_USER_FOLLOWING(args.UserID)

	} else if args.Recommend {
		app.ShellRecommend("", true)

	} else if args.Ranking {
		app.ShellRanking()

	} else if args.Stars {
		app.ShellStars(config.Vars.UserID, "")

	} else {
		return false

	}
	return true
}
