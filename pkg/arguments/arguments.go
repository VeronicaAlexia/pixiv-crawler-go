package arguments

import (
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/cli"
)

type Command struct {
	IllustID  string
	Name      string
	URL       string
	AuthorID  string
	Following bool
	Recommend bool
	Ranking   bool
	Stars     bool
	Login     bool
	IsNovel   bool
	ThreadMax int
	UserID    int
}

var CommandLines = Command{}

var CommandLineFlag = []cli.Flag{
	cli.StringFlag{
		Name:        "d, download",
		Value:       "",
		Usage:       "Input IllustID to download illusts.",
		Destination: &CommandLines.IllustID,
	},
	cli.StringFlag{
		Name:        "u, url",
		Value:       "",
		Usage:       "Input pixiv url to download illusts",
		Destination: &CommandLines.URL,
	},
	cli.StringFlag{
		Name:        "a, author",
		Value:       "",
		Usage:       "Input AuthorID to download Author illusts.",
		Destination: &CommandLines.AuthorID,
	},
	cli.IntFlag{
		Name:        "user, userid",
		Value:       0,
		Usage:       "Input user id to change default user id",
		Destination: &CommandLines.UserID,
	},
	cli.IntFlag{
		Name:        "m, max",
		Value:       16,
		Usage:       "Input max thread number",
		Destination: &CommandLines.ThreadMax,
	},
	cli.BoolFlag{
		Name:        "f, following",
		Usage:       "Download illusts from following users",
		Destination: &CommandLines.Following,
	},
	cli.BoolFlag{
		Name:        "r, recommend",
		Usage:       "Download recommend illusts",
		Destination: &CommandLines.Recommend,
	},
	cli.BoolFlag{
		Name:        "s, stars",
		Usage:       "download stars illusts.",
		Destination: &CommandLines.Stars,
	},
	cli.BoolFlag{
		Name:        "rk, ranking",
		Usage:       "Download ranking illusts.",
		Destination: &CommandLines.Ranking,
	},
	cli.BoolFlag{
		Name:        "l, login",
		Usage:       "login pixiv account",
		Destination: &CommandLines.Login,
	},
	cli.BoolFlag{
		Name:        "n, novel",
		Usage:       "download novel",
		Destination: &CommandLines.IsNovel,
	},
}
