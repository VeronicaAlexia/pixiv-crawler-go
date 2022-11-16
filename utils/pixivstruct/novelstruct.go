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
type NovelContent struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		BookmarkCount     int       `json:"bookmarkCount"`
		CommentCount      int       `json:"commentCount"`
		MarkerCount       int       `json:"markerCount"`
		CreateDate        time.Time `json:"createDate"`
		UploadDate        time.Time `json:"uploadDate"`
		Description       string    `json:"description"`
		ID                string    `json:"id"`
		Title             string    `json:"title"`
		LikeCount         int       `json:"likeCount"`
		PageCount         string    `json:"pageCount"`
		UserID            string    `json:"userId"`
		UserName          string    `json:"userName"`
		ViewCount         int       `json:"viewCount"`
		IsOriginal        bool      `json:"isOriginal"`
		IsBungei          bool      `json:"isBungei"`
		XRestrict         int       `json:"xRestrict"`
		Restrict          int       `json:"restrict"`
		Content           string    `json:"content"`
		CoverURL          string    `json:"coverUrl"`
		SuggestedSettings struct {
			ViewMode        int         `json:"viewMode"`
			ThemeBackground int         `json:"themeBackground"`
			ThemeSize       interface{} `json:"themeSize"`
			ThemeSpacing    interface{} `json:"themeSpacing"`
		} `json:"suggestedSettings"`
		IsBookmarkable bool        `json:"isBookmarkable"`
		BookmarkData   interface{} `json:"bookmarkData"`
		LikeData       bool        `json:"likeData"`
		PollData       interface{} `json:"pollData"`
		Marker         interface{} `json:"marker"`
		Tags           struct {
			AuthorID string `json:"authorId"`
			IsLocked bool   `json:"isLocked"`
			Tags     []struct {
				Tag       string `json:"tag"`
				Locked    bool   `json:"locked"`
				Deletable bool   `json:"deletable"`
				UserID    string `json:"userId"`
				UserName  string `json:"userName"`
			} `json:"tags"`
			Writable bool `json:"writable"`
		} `json:"tags"`
		SeriesNavData        interface{}   `json:"seriesNavData"`
		DescriptionBoothID   interface{}   `json:"descriptionBoothId"`
		DescriptionYoutubeID interface{}   `json:"descriptionYoutubeId"`
		ComicPromotion       interface{}   `json:"comicPromotion"`
		FanboxPromotion      interface{}   `json:"fanboxPromotion"`
		ContestBanners       []interface{} `json:"contestBanners"`
		ContestData          interface{}   `json:"contestData"`
		Request              interface{}   `json:"request"`
		ImageResponseOutData []interface{} `json:"imageResponseOutData"`
		ImageResponseData    []interface{} `json:"imageResponseData"`
		ImageResponseCount   int           `json:"imageResponseCount"`
		HasGlossary          bool          `json:"hasGlossary"`
		ZoneConfig           struct {
			Responsive struct {
				URL string `json:"url"`
			} `json:"responsive"`
			Rectangle struct {
				URL string `json:"url"`
			} `json:"rectangle"`
			Five00X500 struct {
				URL string `json:"url"`
			} `json:"500x500"`
			Header struct {
				URL string `json:"url"`
			} `json:"header"`
			Footer struct {
				URL string `json:"url"`
			} `json:"footer"`
			ExpandedFooter struct {
				URL string `json:"url"`
			} `json:"expandedFooter"`
			Logo struct {
				URL string `json:"url"`
			} `json:"logo"`
			Relatedworks struct {
				URL string `json:"url"`
			} `json:"relatedworks"`
		} `json:"zoneConfig"`
		ExtraData struct {
			Meta struct {
				Title             string `json:"title"`
				Description       string `json:"description"`
				Canonical         string `json:"canonical"`
				DescriptionHeader string `json:"descriptionHeader"`
				Ogp               struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Type        string `json:"type"`
				} `json:"ogp"`
				Twitter struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Card        string `json:"card"`
				} `json:"twitter"`
			} `json:"meta"`
		} `json:"extraData"`
		TitleCaptionTranslation struct {
			WorkTitle   interface{} `json:"workTitle"`
			WorkCaption interface{} `json:"workCaption"`
		} `json:"titleCaptionTranslation"`
		IsUnlisted         bool        `json:"isUnlisted"`
		Language           string      `json:"language"`
		TextEmbeddedImages interface{} `json:"textEmbeddedImages"`
		CommentOff         int         `json:"commentOff"`
		CharacterCount     int         `json:"characterCount"`
		WordCount          int         `json:"wordCount"`
		UseWordCount       bool        `json:"useWordCount"`
		ReadingTime        int         `json:"readingTime"`
		AiType             int         `json:"aiType"`
	} `json:"body"`
}
type SeriesContent struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		SeriesContents []struct {
			Id     string `json:"id"`
			UserId string `json:"userId"`
			Series struct {
				Id           int `json:"id"`
				ViewableType int `json:"viewableType"`
				ContentOrder int `json:"contentOrder"`
			} `json:"series"`
			Title             string      `json:"title"`
			CommentHtml       string      `json:"commentHtml"`
			Tags              []string    `json:"tags"`
			Restrict          int         `json:"restrict"`
			XRestrict         int         `json:"xRestrict"`
			IsOriginal        bool        `json:"isOriginal"`
			TextLength        int         `json:"textLength"`
			CharacterCount    int         `json:"characterCount"`
			WordCount         int         `json:"wordCount"`
			UseWordCount      bool        `json:"useWordCount"`
			ReadingTime       int         `json:"readingTime"`
			BookmarkCount     int         `json:"bookmarkCount"`
			Url               string      `json:"url"`
			UploadTimestamp   int         `json:"uploadTimestamp"`
			ReuploadTimestamp int         `json:"reuploadTimestamp"`
			IsBookmarkable    bool        `json:"isBookmarkable"`
			BookmarkData      interface{} `json:"bookmarkData"`
			AiType            int         `json:"aiType"`
		} `json:"seriesContents"`
	} `json:"body"`
}
