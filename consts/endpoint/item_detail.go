package endpoint

import (
	"fmt"
	"github.com/xieburoucoco/go-tiktok/consts"
	"net/url"
)

const (
	ITEM_DETAIL_ENDPOINT_NAME = "ItemDetail"
	ITEM_DETAIL_METHOD        = consts.GET
)

func GetItemDetailRoute() string {
	return consts.API_ENDPOINT + "item/detail/"
}

func BuildItemDetailRoute() string {
	buildUrl := GetItemDetailRoute() + "?"
	startUrl := buildUrl
	addAndConcatParam := func(key string, value string) {
		if buildUrl != startUrl {
			buildUrl += "&"
		}
		//escapedValue := strings.ReplaceAll(url.QueryEscape(value), "+", "%20")
		escapedValue := url.QueryEscape(value)
		buildUrl += fmt.Sprintf("%s=%s", key, escapedValue)
	}
	addAndConcatParam("aid", "1988")
	addAndConcatParam("app_name", "tiktok_web")
	addAndConcatParam("browser_language", "zh-CN")
	addAndConcatParam("browser_name", "Mozilla")
	addAndConcatParam("browser_online", "true")
	addAndConcatParam("browser_platform", "Win32")
	addAndConcatParam("browser_version", "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
	addAndConcatParam("cookie_enabled", "true")
	addAndConcatParam("device_id", "7419900769212581419")
	addAndConcatParam("device_platform", "web_pc")
	addAndConcatParam("os", "windows")
	addAndConcatParam("region", "US")
	addAndConcatParam("screen_height", "1080")
	addAndConcatParam("screen_width", "1920")
	addAndConcatParam("tz_name", "Asia/Shanghai")
	return buildUrl
}

func GetItemDetailParams(itemId string) map[string]interface{} {
	params := make(map[string]interface{})
	params["itemId"] = itemId
	return params
}

func BuildItemDetailEndpoint(itemId string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, ItemDetailRes, error) {
	res := ItemDetailRes{}
	return ITEM_DETAIL_METHOD, BuildItemDetailRoute(), GetItemDetailParams(itemId), make(map[string]interface{}), res, nil
}

//s.options.URL += "?"
//startUrl := s.options.URL
//addAndConcatParam := func(key string, value string) {
//	if s.options.URL != startUrl {
//		s.options.URL += "&"
//	}
//	escapedValue := url.QueryEscape(value)
//	s.options.URL += fmt.Sprintf("%s=%s", key, escapedValue)
//}
//addAndConcatParam("WebIdLastTime", "1738896054")
//addAndConcatParam("aid", "1988")
//addAndConcatParam("app_language", "zh-Hans")
//addAndConcatParam("app_name", "tiktok_web")
//addAndConcatParam("browser_language", "zh-CN")
//addAndConcatParam("browser_name", "Mozilla")
//addAndConcatParam("browser_online", "true")
//addAndConcatParam("browser_platform", "Win32")
//addAndConcatParam("browser_version", ua)
//addAndConcatParam("channel", "tiktok_web")
//addAndConcatParam("cookie_enabled", "true")
//
//addAndConcatParam("count", "30")
//addAndConcatParam("coverFormat", "2")
//addAndConcatParam("cursor", "0")
//
//addAndConcatParam("data_collection_enabled", "true")
//addAndConcatParam("device_id", "7468501464287299079")
//addAndConcatParam("device_platform", "web_pc")
//addAndConcatParam("focus_state", "true")
//addAndConcatParam("from_page", "true")
//addAndConcatParam("history_len", "2")
//addAndConcatParam("is_fullscreen", "false")
//addAndConcatParam("is_page_visible", "true")
////addAndConcatParam("maxCursor", "0")
////addAndConcatParam("minCursor", "1744374151")
//addAndConcatParam("language", "zh-Hans")
//addAndConcatParam("os", "windows")
//addAndConcatParam("priority_region", "")
//addAndConcatParam("referer", "")
//addAndConcatParam("region", "US")
//addAndConcatParam("screen_height", "1152")
//addAndConcatParam("screen_width", "2048")
//addAndConcatParam("tz_name", "Asia/Shanghai")
//addAndConcatParam("webcast_language", "zh-Hans")

type ItemDetailRes struct {
	Extra struct {
		FatalItemIds []interface{} `json:"fatal_item_ids"`
		Logid        string        `json:"logid"`
		Now          int64         `json:"now"`
	} `json:"extra"`
	ItemInfo struct {
		ItemStruct struct {
			AIGCDescription string `json:"AIGCDescription"`
			CategoryType    int    `json:"CategoryType"`
			Author          struct {
				AvatarLarger    string `json:"avatarLarger"`
				AvatarMedium    string `json:"avatarMedium"`
				AvatarThumb     string `json:"avatarThumb"`
				CommentSetting  int    `json:"commentSetting"`
				DownloadSetting int    `json:"downloadSetting"`
				DuetSetting     int    `json:"duetSetting"`
				Ftc             bool   `json:"ftc"`
				ID              string `json:"id"`
				IsADVirtual     bool   `json:"isADVirtual"`
				IsEmbedBanned   bool   `json:"isEmbedBanned"`
				Nickname        string `json:"nickname"`
				OpenFavorite    bool   `json:"openFavorite"`
				PrivateAccount  bool   `json:"privateAccount"`
				Relation        int    `json:"relation"`
				SecUID          string `json:"secUid"`
				Secret          bool   `json:"secret"`
				Signature       string `json:"signature"`
				StitchSetting   int    `json:"stitchSetting"`
				UniqueID        string `json:"uniqueId"`
				Verified        bool   `json:"verified"`
			} `json:"author"`
			AuthorStats struct {
				DiggCount      int `json:"diggCount"`
				FollowerCount  int `json:"followerCount"`
				FollowingCount int `json:"followingCount"`
				FriendCount    int `json:"friendCount"`
				Heart          int `json:"heart"`
				HeartCount     int `json:"heartCount"`
				VideoCount     int `json:"videoCount"`
			} `json:"authorStats"`
			BackendSourceEventTracking string `json:"backendSourceEventTracking"`
			Challenges                 []struct {
				CoverLarger   string `json:"coverLarger"`
				CoverMedium   string `json:"coverMedium"`
				CoverThumb    string `json:"coverThumb"`
				Desc          string `json:"desc"`
				ID            string `json:"id"`
				ProfileLarger string `json:"profileLarger"`
				ProfileMedium string `json:"profileMedium"`
				ProfileThumb  string `json:"profileThumb"`
				Title         string `json:"title"`
			} `json:"challenges"`
			Collected bool `json:"collected"`
			Contents  []struct {
				Desc      string `json:"desc"`
				TextExtra []struct {
					AwemeID     string `json:"awemeId"`
					End         int    `json:"end"`
					HashtagID   string `json:"hashtagId"`
					HashtagName string `json:"hashtagName"`
					IsCommerce  bool   `json:"isCommerce"`
					Start       int    `json:"start"`
					SubType     int    `json:"subType"`
					Type        int    `json:"type"`
				} `json:"textExtra,omitempty"`
			} `json:"contents"`
			CreateTime        int    `json:"createTime"`
			Desc              string `json:"desc"`
			Digged            bool   `json:"digged"`
			DiversificationID int    `json:"diversificationId"`
			DuetDisplay       int    `json:"duetDisplay"`
			DuetEnabled       bool   `json:"duetEnabled"`
			ForFriend         bool   `json:"forFriend"`
			ID                string `json:"id"`
			IsAd              bool   `json:"isAd"`
			ItemCommentStatus int    `json:"itemCommentStatus"`
			ItemControl       struct {
				CanRepost bool `json:"can_repost"`
			} `json:"item_control"`
			KeywordTags []struct {
				Keyword  string `json:"keyword"`
				PageType int    `json:"pageType"`
			} `json:"keywordTags"`
			Music struct {
				Album         string `json:"album"`
				AuthorName    string `json:"authorName"`
				CoverLarge    string `json:"coverLarge"`
				CoverMedium   string `json:"coverMedium"`
				CoverThumb    string `json:"coverThumb"`
				Duration      int    `json:"duration"`
				ID            string `json:"id"`
				IsCopyrighted bool   `json:"isCopyrighted"`
				Original      bool   `json:"original"`
				PlayURL       string `json:"playUrl"`
				Private       bool   `json:"private"`
				Title         string `json:"title"`
				Tt2Dsp        struct {
				} `json:"tt2dsp"`
			} `json:"music"`
			OfficalItem  bool `json:"officalItem"`
			OriginalItem bool `json:"originalItem"`
			PrivateItem  bool `json:"privateItem"`
			Secret       bool `json:"secret"`
			ShareEnabled bool `json:"shareEnabled"`
			Stats        struct {
				CollectCount int `json:"collectCount"`
				CommentCount int `json:"commentCount"`
				DiggCount    int `json:"diggCount"`
				PlayCount    int `json:"playCount"`
				ShareCount   int `json:"shareCount"`
			} `json:"stats"`
			StatsV2 struct {
				CollectCount string `json:"collectCount"`
				CommentCount string `json:"commentCount"`
				DiggCount    string `json:"diggCount"`
				PlayCount    string `json:"playCount"`
				RepostCount  string `json:"repostCount"`
				ShareCount   string `json:"shareCount"`
			} `json:"statsV2"`
			StitchDisplay  int      `json:"stitchDisplay"`
			StitchEnabled  bool     `json:"stitchEnabled"`
			SuggestedWords []string `json:"suggestedWords"`
			TextExtra      []struct {
				AwemeID     string `json:"awemeId"`
				End         int    `json:"end"`
				HashtagID   string `json:"hashtagId"`
				HashtagName string `json:"hashtagName"`
				IsCommerce  bool   `json:"isCommerce"`
				Start       int    `json:"start"`
				SubType     int    `json:"subType"`
				Type        int    `json:"type"`
			} `json:"textExtra"`
			TextLanguage     string `json:"textLanguage"`
			TextTranslatable bool   `json:"textTranslatable"`
			Video            struct {
				VQScore     string `json:"VQScore"`
				Bitrate     int    `json:"bitrate"`
				BitrateInfo []struct {
					Bitrate   int    `json:"Bitrate"`
					CodecType string `json:"CodecType"`
					GearName  string `json:"GearName"`
					MVMAF     string `json:"MVMAF"`
					PlayAddr  struct {
						DataSize int      `json:"DataSize"`
						FileCs   string   `json:"FileCs"`
						FileHash string   `json:"FileHash"`
						Height   int      `json:"Height"`
						URI      string   `json:"Uri"`
						URLKey   string   `json:"UrlKey"`
						URLList  []string `json:"UrlList"`
						Width    int      `json:"Width"`
					} `json:"PlayAddr"`
					QualityType int `json:"QualityType"`
				} `json:"bitrateInfo"`
				ClaInfo struct {
					EnableAutoCaption bool `json:"enableAutoCaption"`
					HasOriginalAudio  bool `json:"hasOriginalAudio"`
					NoCaptionReason   int  `json:"noCaptionReason"`
				} `json:"claInfo"`
				CodecType     string `json:"codecType"`
				Cover         string `json:"cover"`
				Definition    string `json:"definition"`
				DownloadAddr  string `json:"downloadAddr"`
				Duration      int    `json:"duration"`
				DynamicCover  string `json:"dynamicCover"`
				EncodeUserTag string `json:"encodeUserTag"`
				EncodedType   string `json:"encodedType"`
				Format        string `json:"format"`
				Height        int    `json:"height"`
				ID            string `json:"id"`
				OriginCover   string `json:"originCover"`
				PlayAddr      string `json:"playAddr"`
				Ratio         string `json:"ratio"`
				Size          int    `json:"size"`
				VideoID       string `json:"videoID"`
				VideoQuality  string `json:"videoQuality"`
				VolumeInfo    struct {
					Loudness float64 `json:"Loudness"`
					Peak     int     `json:"Peak"`
				} `json:"volumeInfo"`
				Width     int `json:"width"`
				ZoomCover struct {
					Num240 string `json:"240"`
					Num480 string `json:"480"`
					Num720 string `json:"720"`
					Num960 string `json:"960"`
				} `json:"zoomCover"`
			} `json:"video"`
		} `json:"itemStruct"`
	} `json:"itemInfo"`
	LogPb struct {
		ImprID string `json:"impr_id"`
	} `json:"log_pb"`
	ShareMeta struct {
		Desc  string `json:"desc"`
		Title string `json:"title"`
	} `json:"shareMeta"`
	StatusCode    int    `json:"statusCode"`
	StatusAndCode int    `json:"status_code"`
	StatusMsg     string `json:"status_msg"`
}
