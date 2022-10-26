package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"pixiv-cil/pkg/arguments"
	"pixiv-cil/pkg/config"
	"pixiv-cil/pkg/file"
	"pixiv-cil/src"
	"pixiv-cil/src/pixiv"
	"pixiv-cil/utils"
)

func init() {
	config.VarsConfigInit()
	if config.Vars.PixivRefreshToken == "" {
		if accessToken, err := pixiv.ChromeDriverLogin(); err != nil {
			panic(err)
		} else {
			config.VarsFile.Vipers.Set("PIXIV_REFRESH_TOKEN", accessToken.RefreshToken)
			config.VarsFile.Vipers.Set("PIXIV_TOKEN", accessToken.AccessToken)
			config.VarsFile.Vipers.Set("PIXIV_USER_ID", accessToken.User.ID)
			config.VarsFile.SaveConfig()
		}
	}
	pixiv.TokenVariable = config.Vars.PixivToken
	pixiv.RefreshTokenVariable = config.Vars.PixivRefreshToken
}

func main() {
	file.NewFile("imageFile")
	cli_app := cli.NewApp()
	cli_app.Name = "image downloader"
	cli_app.Version = config.Vars.VersionName
	cli_app.Usage = "download image from pixiv "
	cli_app.Flags = arguments.CommandLineFlag
	cli_app.Action = func(c *cli.Context) error {
		if arguments.CommandLines.IllustID != 0 {
			src.CurrentDownloader(arguments.CommandLines.IllustID)
		} else if arguments.CommandLines.AuthorID != 0 {
			src.ThreadDownloadImages(src.GET_AUTHOR_INFO(arguments.CommandLines.AuthorID, 0))
		} else if arguments.CommandLines.URL != "" {
			src.CurrentDownloader(utils.GetInt(arguments.CommandLines.URL))
		} else if arguments.CommandLines.Following {
			src.GET_USER_FOLLOWING(arguments.CommandLines.UserID)
		} else if arguments.CommandLines.Recommend {
			src.GET_RECOMMEND("")
		} else {
			if len(os.Args) == 1 {
				_ = cli.ShowAppHelp(c)
			} else {
				if os.Args[1] == "-h" || os.Args[1] == "--help" {
					_ = cli.ShowAppHelp(c)
				}
			}
		}
		return nil
	}
	if err := cli_app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
