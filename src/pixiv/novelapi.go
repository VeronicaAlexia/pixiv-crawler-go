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

func (a *AppPixivAPI) NovelContent(novel_id string) (string, error) {
	response := request.Get(WEB_BASE+WEB_BOOK_CONTENT+novel_id, nil).Json(&pixivstruct.NovelContent{}).(*pixivstruct.NovelContent)
	if response.Error {
		return "", errors.New(response.Message)
	}
	//return map[string]string{
	//	"novel_id":     response.Body.ID,
	//	"author_id":    response.Body.UserID,
	//	"book_name":    response.Body.Title,
	//	"cover_url":    response.Body.CoverURL,
	//	"description":  response.Body.Description,
	//	"content_text": response.Body.Content,
	//	"author_namer": response.Body.UserName,
	//	"create_date":  response.Body.CreateDate.Format("2006-01-02 15:04:05"),
	//	"update_date":  response.Body.UploadDate.Format("2006-01-02 15:04:05"),
	//}, nil
	return response.Body.Content, nil
}
func (a *AppPixivAPI) AppNovelContent(novel_id string) string {
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
