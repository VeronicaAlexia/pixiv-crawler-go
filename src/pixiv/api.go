package pixiv

import (
	"github.com/VeronicaAlexia/pixiv-crawler-go/pkg/request"
	"github.com/VeronicaAlexia/pixiv-crawler-go/utils/pixivstruct"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

// AppPixivAPI -- App-API (6.x - app-api.pixiv.net)
type AppPixivAPI struct{}

func NewApp() *AppPixivAPI {
	return &AppPixivAPI{}
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
func (a *AppPixivAPI) UserIllusts(uid string, next_url string) (*pixivstruct.IllustsResponse, error) {
	params := map[string]string{"user_id": uid, "filter": "for_ios", "type": "illust", "offset": "0"}
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

// IllustComments Comments posted in a pixiv artwork
func (a *AppPixivAPI) IllustComments(illustID int, next_url string) (*pixivstruct.IllustComments, error) {
	params := map[string]string{"illust_id": strconv.Itoa(illustID), "offset": "0", "include_total_comments": "true"}
	response := request.Get(NextUrl(next_url, COMMENTS, params)).Json(&pixivstruct.IllustComments{}).(*pixivstruct.IllustComments)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// IllustCommentAdd adds a comment to given illustID
func (a *AppPixivAPI) IllustCommentAdd(illustID int, comment string, parentCommentID int) (*pixivstruct.IllustCommentAddResult, error) {
	params := map[string]string{"illust_id": strconv.Itoa(illustID), "comment": comment, "parent_comment_id": strconv.Itoa(parentCommentID)}
	response := request.Post(API_BASE+ADD, params).Json(&pixivstruct.IllustCommentAddResult{}).(*pixivstruct.IllustCommentAddResult)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// IllustRelated returns Related works
func (a *AppPixivAPI) IllustRelated(illustID int, filter string, seedIllustIDs []string) (*pixivstruct.IllustsResponse, error) {
	params := map[string]string{"illust_id": strconv.Itoa(illustID)}
	if filter == "" {
		params["filter"] = "for_ios"
	}
	if seedIllustIDs != nil {
		params["seed_illust_ids"] = "[" + strings.Join(seedIllustIDs, ",") + "]"
	}
	response := request.Post(API_BASE+RELATED, params).Json(&pixivstruct.IllustsResponse{}).(*pixivstruct.IllustsResponse)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// IllustRecommended Home Recommendation
// contentType: [illust, manga]

func (a *AppPixivAPI) Recommended(next_url string, requireAuth bool) (*pixivstruct.IllustRecommended, error) {
	params := map[string]string{"include_privacy_policy": "true", "include_ranking_illusts": "true"}
	var path = RECOMMENDED_NO_LOGIN
	if requireAuth {
		path = RECOMMENDED // auth required
	}
	response := request.Get(NextUrl(next_url, path, params)).Json(&pixivstruct.IllustRecommended{}).(*pixivstruct.IllustRecommended)
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

// TrendingTagsIllust Trend label
func (a *AppPixivAPI) TrendingTagsIllust(filter string) (*pixivstruct.TrendingTagsIllust, error) {
	params := map[string]string{"filter": filter}
	response := request.Get(API_BASE+TRENDING_TAGS, params).Json(&pixivstruct.TrendingTagsIllust{}).(*pixivstruct.TrendingTagsIllust)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// SearchIllust search for
// searchTarget - Search type
//
//	"partial_match_for_tags"  - The label part is consistent
//	"exact_match_for_tags"    - The labels are exactly the same
//	"title_and_caption"       - Title description
//
// sort: [date_desc, date_asc]
//
// duration: [within_last_day, within_last_week, within_last_month]
func (a *AppPixivAPI) SearchIllust(word string, next_url string) (*pixivstruct.SearchIllustResult, error) {
	params := map[string]string{
		"word":          word,
		"search_target": "partial_match_for_tags",
		"sort":          "date_desc",
		"filter":        "for_ios",
		"duration":      "within_last_day",
		"offset":        "0",
	}
	response := request.Get(NextUrl(next_url, SEARCH, params)).Json(&pixivstruct.SearchIllustResult{}).(*pixivstruct.SearchIllustResult)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// IllustBookmarkDetail Bookmark details
func (a *AppPixivAPI) IllustBookmarkDetail(illustID int) (*pixivstruct.IllustBookmarkDetail, error) {
	params := map[string]string{"illust_id": strconv.Itoa(illustID)}
	response := request.Get(API_BASE+BOOKMARK_DETAIL, params).Json(&pixivstruct.IllustBookmarkDetail{}).(*pixivstruct.IllustBookmarkDetail)
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

// UserBookmarkTagsIllust User favorite tag list
func (a *AppPixivAPI) UserBookmarkTagsIllust(restrict string, next_url string) (*pixivstruct.UserBookmarkTags, error) {
	params := map[string]string{"restrict": restrict, "offset": "0"}
	response := request.Get(NextUrl(next_url, BOOKMARK_TAG, params)).Json(&pixivstruct.UserBookmarkTags{}).(*pixivstruct.UserBookmarkTags)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
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

// UserMyPixiv Users in MyPixiv
func (a *AppPixivAPI) UserMyPixiv(userID int, next_url string) (*pixivstruct.UserFollowList, error) {
	params := map[string]string{"user_id": strconv.Itoa(userID), "offset": "0"}
	response := request.Get(NextUrl(next_url, USER_MYPIXIV, params)).Json(&pixivstruct.UserFollowList{}).(*pixivstruct.UserFollowList)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
}

// UserList Blacklisted users
func (a *AppPixivAPI) UserList(userID int, next_url string) (*pixivstruct.UserList, error) {
	params := map[string]string{"user_id": strconv.Itoa(userID), "offset": "0", "filter": "for_ios"}
	response := request.Get(NextUrl(next_url, USER_LIST, params)).Json(&pixivstruct.UserList{}).(*pixivstruct.UserList)
	if response.Error.Message != "" {
		return nil, errors.New(response.Error.Message)
	}
	return response, nil
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
