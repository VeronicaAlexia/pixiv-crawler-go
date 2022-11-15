package pixivstruct

import "time"

type NovelDetail struct {
	Novel NovelInfo `json:"novel"`
	Error Error     `json:"error"`
}
type ImageUrls struct {
	SquareMedium string `json:"square_medium"`
	Medium       string `json:"medium"`
	Large        string `json:"large"`
}
type Tags struct {
	Name                string      `json:"name"`
	TranslatedName      interface{} `json:"translated_name"`
	AddedByUploadedUser bool        `json:"added_by_uploaded_user"`
}
type ProfileImageUrls struct {
	Medium string `json:"medium"`
}
type NovelInfo struct {
	ID                   int       `json:"id"`
	Title                string    `json:"title"`
	Caption              string    `json:"caption"`
	Restrict             int       `json:"restrict"`
	XRestrict            int       `json:"x_restrict"`
	IsOriginal           bool      `json:"is_original"`
	ImageUrls            ImageUrls `json:"image_urls"`
	CreateDate           time.Time `json:"create_date"`
	Tags                 []Tags    `json:"tags"`
	PageCount            int       `json:"page_count"`
	TextLength           int       `json:"text_length"`
	User                 User      `json:"user"`
	Series               Series    `json:"series"`
	IsBookmarked         bool      `json:"is_bookmarked"`
	TotalBookmarks       int       `json:"total_bookmarks"`
	TotalView            int       `json:"total_view"`
	Visible              bool      `json:"visible"`
	TotalComments        int       `json:"total_comments"`
	IsMuted              bool      `json:"is_muted"`
	IsMypixivOnly        bool      `json:"is_mypixiv_only"`
	IsXRestricted        bool      `json:"is_x_restricted"`
	NovelAiType          int       `json:"novel_ai_type"`
	CommentAccessControl int       `json:"comment_access_control"`
}
