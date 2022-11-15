package pixiv

import (
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/request"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils/pixivstruct"
	"github.com/pkg/errors"
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
