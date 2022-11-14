package pixiv

import (
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/config"
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/request"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils/pixivstruct"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/dghubble/sling"
	"github.com/pkg/errors"
)

// AppPixivAPI -- App-API (6.x - app-api.pixiv.net)
type AppPixivAPI struct {
	sling   *sling.Sling
	timeout time.Duration
	proxy   *url.URL
}

func NewApp() *AppPixivAPI {
	s := sling.New().Base(API_BASE).Set("User-Agent", "PixivIOSApp/7.6.2 (iOS 12.2; iPhone9,1)").Set("App-Version", "7.6.2").Set("App-OS-VERSION", "12.2").Set("App-OS", "ios")
	return &AppPixivAPI{sling: s}
}

func (a *AppPixivAPI) request(path string, params, data interface{}, auth bool) (err error) {
	var res *http.Response
	if auth {
		res, err = a.sling.New().Get(path).Set("Authorization", "Bearer "+config.Vars.PixivToken).QueryStruct(params).ReceiveSuccess(data)
		if res.StatusCode == 400 {
			if !request.RefreshAuth() {
				return errors.New("refresh token failed")
			} else {
				return a.request(path, params, data, auth)
			}
		}
	} else {
		res, err = a.sling.New().Get(path).QueryStruct(params).ReceiveSuccess(data)
	}
	return err
}

func (a *AppPixivAPI) WithDownloadTimeout(timeout time.Duration) *AppPixivAPI {
	a.timeout = timeout
	return a
}

func (a *AppPixivAPI) WithDownloadProxy(proxy *url.URL) *AppPixivAPI {
	a.proxy = proxy
	return a
}

func (a *AppPixivAPI) post(path string, params, data interface{}, auth bool) (err error) {
	if auth {
		_, err = a.sling.New().Post(path).Set("Authorization", "Bearer "+config.Vars.PixivToken).BodyForm(params).ReceiveSuccess(data)
	} else {
		_, err = a.sling.New().Post(path).BodyForm(params).ReceiveSuccess(data)
	}
	return err
}

func (a *AppPixivAPI) UserDetail(uid int) (*pixivstruct.UserDetail, error) {
	params := map[string]string{"user_id": strconv.Itoa(uid), "filter": "for_ios"}
	response := request.Get(API_BASE+USER_DETAIL, params).Json(&pixivstruct.UserDetail{}).(*pixivstruct.UserDetail)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	} else {
		return response, nil
	}
}

// UserIllusts type: [illust, manga]
func (a *AppPixivAPI) UserIllusts(uid int, next_url string) (*pixivstruct.IllustsResponse, error) {
	params := map[string]string{"user_id": strconv.Itoa(uid), "filter": "for_ios", "type": "illust", "offset": "0"}
	response := request.Get(NextUrl(next_url, USER_AUTHOR, params)).Json(&pixivstruct.IllustsResponse{}).(*pixivstruct.IllustsResponse)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// UserBookmarksIllust restrict: [public, private]
func (a *AppPixivAPI) UserBookmarksIllust(uid int, next_url string) (*pixivstruct.IllustsResponse, error) {
	params := map[string]string{"user_id": strconv.Itoa(uid), "restrict": "public"}
	response := request.Get(NextUrl(next_url, BOOKMARKS, params)).Json(&pixivstruct.IllustsResponse{}).(*pixivstruct.IllustsResponse)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	} else {
		return response, nil
	}
}

// IllustFollow restrict: [public, private]
func (a *AppPixivAPI) IllustFollow(restrict string, offset int) ([]pixivstruct.Illust, error) {
	params := map[string]string{"restrict": restrict, "offset": strconv.Itoa(offset)}
	response := request.Get(API_BASE+FOLLOW, params).Json(&pixivstruct.IllustsResponse{}).(*pixivstruct.IllustsResponse)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response.Illusts, nil
}

func (a *AppPixivAPI) IllustDetail(id string) (*pixivstruct.Illust, error) {
	params := map[string]string{"illust_id": id}
	response := request.Get(API_BASE+DETAIL, params).Json(&pixivstruct.IllustResponse{}).(*pixivstruct.IllustResponse)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return &response.Illust, nil
}

type illustCommentsParams struct {
	IllustID             int  `url:"illust_id,omitemtpy"`
	Offset               int  `url:"offset,omitempty"`
	IncludeTotalComments bool `url:"include_total_comments,omitempty"`
}

// IllustComments Comments posted in a pixiv artwork
func (a *AppPixivAPI) IllustComments(illustID int, offset int, includeTotalComments bool) (*pixivstruct.IllustComments, error) {
	data := &pixivstruct.IllustComments{}
	params := &illustCommentsParams{
		IllustID:             illustID,
		IncludeTotalComments: includeTotalComments,
		Offset:               offset,
	}

	if err := a.request(COMMENTS, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

type illustCommentAddParams struct {
	IllustID        uint64 `url:"illust_id,omitempty"`
	Comment         string `url:"comment,omitempty"`
	ParentCommentID int    `url:"parent_comment_id,omitempty"`
}

// IllustCommentAdd adds a comment to given illustID
func (a *AppPixivAPI) IllustCommentAdd(illustID uint64, comment string, parentCommentID int) (*pixivstruct.IllustCommentAddResult, error) {
	data := &pixivstruct.IllustCommentAddResult{}
	params := &illustCommentAddParams{
		IllustID:        illustID,
		Comment:         comment,
		ParentCommentID: parentCommentID,
	}
	if err := a.post(ADD, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

type illustRelatedParams struct {
	IllustID      uint64   `url:"illust_id,omitempty"`
	Filter        string   `url:"filter,omitempty"`
	SeedIllustIDs []string `url:"seed_illust_ids[],omitempty,omitempty"`
}

// IllustRelated returns Related works
func (a *AppPixivAPI) IllustRelated(illustID uint64, filter string, seedIllustIDs []string) (*pixivstruct.IllustsResponse, error) {
	data := &pixivstruct.IllustsResponse{}
	if filter == "" {
		filter = "for_ios"
	}
	params := &illustRelatedParams{
		IllustID: illustID,
		Filter:   filter,
	}
	if seedIllustIDs != nil {
		params.SeedIllustIDs = seedIllustIDs
	}

	if err := a.request(RELATED, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

// IllustRecommended Home Recommendation
// contentType: [illust, manga]

func (a *AppPixivAPI) Recommended(url string, requireAuth bool) (*pixivstruct.IllustRecommended, error) {
	params := map[string]string{"include_privacy_policy": "true", "include_ranking_illusts": "true"}
	if url == "" {
		if requireAuth {
			url = RECOMMENDED // auth required
		} else {
			url = RECOMMENDED_NO_LOGIN // no auth required
		}
	} else {
		url = strings.ReplaceAll(url, API_BASE, "")
		params = nil
	}
	response := request.Get(API_BASE+url, params).Json(&pixivstruct.IllustRecommended{}).(*pixivstruct.IllustRecommended)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// IllustRanking mode: [day, week, month, day_male, day_female, week_original, week_rookie, day_manga]  date: yyyy-mm-dd
func (a *AppPixivAPI) IllustRanking(mode string) (*pixivstruct.IllustsResponse, error) {
	params := map[string]string{"mode": mode}
	response := request.Get(API_BASE+RANKING, params).Json(&pixivstruct.IllustsResponse{}).(*pixivstruct.IllustsResponse)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

type trendingTagsIllustParams struct {
	Filter string `url:"filter,omitempty"`
}

// TrendingTagsIllust Trend label
func (a *AppPixivAPI) TrendingTagsIllust(filter string) (*pixivstruct.TrendingTagsIllust, error) {
	data := &pixivstruct.TrendingTagsIllust{}
	params := &trendingTagsIllustParams{
		Filter: filter,
	}
	if err := a.request(TRENDING_TAGS, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

type searchIllustParams struct {
	Word         string `url:"word,omitempty"`
	SearchTarget string `url:"search_target,omitempty"`
	Sort         string `url:"sort,omitempty"`
	Filter       string `url:"filter,omitempty"`
	Duration     string `url:"duration,omitempty"`
	Offset       int    `url:"offset,omitempty"`
}

// SearchIllust search for
//
// searchTarget - Search type
//
//	"partial_match_for_tags"  - The label part is consistent
//	"exact_match_for_tags"    - The labels are exactly the same
//	"title_and_caption"       - Title description
//
// sort: [date_desc, date_asc]
//
// duration: [within_last_day, within_last_week, within_last_month]
func (a *AppPixivAPI) SearchIllust(word string, searchTarget string, sort string, duration string, filter string, offset int) (*pixivstruct.SearchIllustResult, error) {
	data := &pixivstruct.SearchIllustResult{}
	params := &searchIllustParams{
		Word:         word,
		SearchTarget: searchTarget,
		Sort:         sort,
		Filter:       filter,
		Duration:     duration,
		Offset:       offset,
	}
	if err := a.request(SEARCH, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

// IllustBookmarkDetail Bookmark details
func (a *AppPixivAPI) IllustBookmarkDetail(illustID int) (*pixivstruct.IllustBookmarkDetail, error) {
	response := request.Get(
		API_BASE+BOOKMARK_DETAIL, map[string]string{"illust_id": strconv.Itoa(illustID)},
	).Json(&pixivstruct.IllustBookmarkDetail{}).(*pixivstruct.IllustBookmarkDetail)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// IllustBookmarkAdd Add bookmark
func (a *AppPixivAPI) IllustBookmarkAdd(illustID int, restrict string, tags []string) error {
	params := map[string]string{"illust_id": strconv.Itoa(illustID), "restrict": restrict}
	if tags != nil {
		params["tags"] = "[" + strings.Join(tags, ",") + "]"
	}
	if response := request.Post(API_BASE+BOOKMARK_ADD, params).Content(); response == nil {
		return errors.New("IllustBookmarkAdd failed")
	}
	return nil
}

// IllustBookmarkDelete Remove bookmark
func (a *AppPixivAPI) IllustBookmarkDelete(illustID int) error {
	params := map[string]string{"illust_id": strconv.Itoa(illustID)}
	if response := request.Post(API_BASE+BOOKMARK_DELETE, params).Content(); response == nil {
		return errors.New("delete bookmark failed")
	}
	return nil
}

type userBookmarkTagsIllustParams struct {
	Restrict string
	Offset   int
}

// UserBookmarkTagsIllust User favorite tag list
func (a *AppPixivAPI) UserBookmarkTagsIllust(restrict string, offset int) (*pixivstruct.UserBookmarkTags, error) {
	data := &pixivstruct.UserBookmarkTags{}
	params := &userBookmarkTagsIllustParams{
		Restrict: restrict,
		Offset:   offset,
	}
	if err := a.request(BOOKMARK_TAG, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

func userFollowStats(urlEnd string, userID int, restrict string, offset int) (*pixivstruct.UserFollowList, error) {
	params := map[string]string{"user_id": strconv.Itoa(userID), "restrict": restrict, "offset": strconv.Itoa(offset)}
	response := request.Get(API_BASE+USER+urlEnd, params).Json(&pixivstruct.UserFollowList{}).(*pixivstruct.UserFollowList)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// UserFollowing Following user list
func (a *AppPixivAPI) UserFollowing(userID int, restrict string, offset int) (*pixivstruct.UserFollowList, error) {
	return userFollowStats("following", userID, restrict, offset)
}

// UserFollower Follower user list
func (a *AppPixivAPI) UserFollower(userID int, restrict string, offset int) (*pixivstruct.UserFollowList, error) {
	return userFollowStats("follower", userID, restrict, offset)
}

func userFollowPost(urlEnd string, userID int, restrict string) error {
	if response := request.Post(
		API_BASE+USER_FOLLOW+urlEnd, map[string]string{"user_id": strconv.Itoa(userID), "restrict": restrict},
	).Content(); response == nil {
		return errors.New("the user is already followed")
	}
	return nil
}

// UserFollowAdd Follow users
func (a *AppPixivAPI) UserFollowAdd(userID int, restrict string) error {
	return userFollowPost("add", userID, restrict)
}

// UserFollowDelete Unfollow users
func (a *AppPixivAPI) UserFollowDelete(userID int, restrict string) error {
	return userFollowPost("delete", userID, restrict)
}

type userMyPixivParams struct {
	UserID uint64 `url:"user_id,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// UserMyPixiv Users in MyPixiv
func (a *AppPixivAPI) UserMyPixiv(userID uint64, offset int) (*pixivstruct.UserFollowList, error) {
	data := &pixivstruct.UserFollowList{}
	params := &userMyPixivParams{
		UserID: userID,
		Offset: offset,
	}
	if err := a.request(USER_MYPIXIV, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

type userListParams struct {
	UserID uint64 `url:"user_id,omitempty"`
	Filter string `url:"filter,omitempty"`
	Offset int    `url:"offset,omitempty"`
}

// UserList Blacklisted users
func (a *AppPixivAPI) UserList(userID uint64, filter string, offset int) (*pixivstruct.UserList, error) {
	data := &pixivstruct.UserList{}
	params := &userListParams{
		UserID: userID,
		Filter: filter,
		Offset: offset,
	}
	if err := a.request(USER_LIST, params, data, true); err != nil {
		return nil, err
	}
	return data, nil
}

// UgoiraMetadata Ugoira Info
func (a *AppPixivAPI) UgoiraMetadata(illustID int) (*pixivstruct.UgoiraMetadata, error) {
	params := map[string]string{"illust_id": strconv.Itoa(illustID)}
	response := request.Get(API_BASE+METADATA, params).Json(&pixivstruct.UgoiraMetadata{}).(*pixivstruct.UgoiraMetadata)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// ShowcaseArticle Special feature details (disguised as Chrome)
func (a *AppPixivAPI) ShowcaseArticle(showcaseID string) (*pixivstruct.ShowcaseArticle, error) {
	params := map[string]string{"article_id": showcaseID}
	response := request.Get(
		WEB_BASE+"/"+WEB_ARTICLE,
		params, map[string]string{
			"Referer":    WEB_BASE,
			"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
		},
	).Json(&pixivstruct.ShowcaseArticle{}).(*pixivstruct.ShowcaseArticle)
	if response.Message != "" {
		return nil, errors.New(response.Message)
	}
	return response, nil
}
