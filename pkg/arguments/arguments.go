package arguments

import (
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/cli"
)

type Command struct {
	IllustID  string
	Name      string
	URL       string
	Following bool
	Recommend bool
	Ranking   bool
	Stars     bool
	Login     bool
	IsNovel   bool
	AuthorID  int
	ThreadMax int
	UserID    int
}

var CommandLines = Command{}

var CommandLineFlag = []cli.Flag{
	cli.StringFlag{
		Name:        "d, download",
		Value:       "",
		Usage:       "input IllustID to download",
		Destination: &CommandLines.IllustID,
	},
	cli.StringFlag{
		Name:        "u, url",
		Value:       "",
		Usage:       "input pixiv url to download",
		Destination: &CommandLines.URL,
	},
	cli.IntFlag{
		Name:        "a, author",
		Value:       0,
		Usage:       "author id",
		Destination: &CommandLines.AuthorID,
	},
	cli.IntFlag{
		Name:        "user, userid",
		Value:       0,
		Usage:       "input user id",
		Destination: &CommandLines.UserID,
	},
	cli.IntFlag{
		Name:        "m, max",
		Value:       16,
		Usage:       "input max thread number",
		Destination: &CommandLines.ThreadMax,
	},
	cli.BoolFlag{
		Name:        "f, following",
		Usage:       "following",
		Destination: &CommandLines.Following,
	},
	cli.BoolFlag{
		Name:        "r, recommend",
		Usage:       "recommend illust",
		Destination: &CommandLines.Recommend,
	},
	cli.BoolFlag{
		Name:        "s, stars",
		Usage:       "download stars",
		Destination: &CommandLines.Stars,
	},
	cli.BoolFlag{
		Name:        "rk, ranking",
		Usage:       "ranking illust",
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
