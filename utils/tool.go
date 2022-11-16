package utils

import (
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils/pixivstruct"
	"strconv"
	"strings"
)

func UnicodeToZh(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func MakeBookInfo(response *pixivstruct.NovelDetail) string {

	book_info := "title: " + response.Novel.Title + "\n"
	book_info += "author: " + response.Novel.User.Name + "\n"
	book_info += "novel_id: " + strconv.Itoa(response.Novel.ID) + "\n"
	book_info += "create_date: " + response.Novel.CreateDate.Format("2006-01-02 15:04:05") + "\n"
	book_info += "intro: " + response.Novel.Caption + "\n"
	book_info += "tags: "
	for index, tag := range response.Novel.Tags {
		book_info += tag.Name
		if index != 0 {
			book_info += ", "
		}
	}
	return book_info
}
