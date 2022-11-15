package pixiv

import (
	"bytes"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/request"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils/pixivstruct"
	"github.com/antchfx/htmlquery"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

func (a *AppPixivAPI) NovelDetail(novel_id string) (*pixivstruct.NovelDetail, error) {
	params := map[string]string{"novel_id": novel_id}
	response := request.Get(API_BASE+BOOK_DETAIL, params).Json(&pixivstruct.NovelDetail{}).(*pixivstruct.NovelDetail)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	} else {
		return response, nil
	}
}

func (a *AppPixivAPI) NovelContent(novel_id string) string {
	response := request.Get(API_BASE+BOOK_CONTENT, map[string]string{"id": novel_id}).Content()
	if response != nil {
		xpath_root, _ := htmlquery.Parse(bytes.NewReader(response))
		content := regexp.MustCompile(`"text":"(.*?)"`).
			FindAllStringSubmatch(
				strings.ReplaceAll(htmlquery.FindOne(xpath_root, `/html/head/script[1]/text()`).Data,
					" ", ""), -1)
		content_text, err := utils.UnicodeToZh([]byte(content[0][1]))
		if err == nil && len(content_text) > 0 {
			return strings.ReplaceAll(string(content_text), `\n`, "\n")
		}
	}
	return ""
}
