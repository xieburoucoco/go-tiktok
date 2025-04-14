package endpoint

import (
	"context"
	"github.com/xieburoucoco/go-tiktok/consts"
)

const (
	ITEM_LIST_ENDPOINT_NAME = "ItemList"
	ITEM_LIST_METHOD        = consts.GET
)

func GetItemListRoute() string {
	return consts.API_ENDPOINT + "post/item_list/"
}

func BuildItemListRoute(ctx context.Context, msToken, secUid, cursor string) (string, error) {
	fullRoute, err := consts.NewSRoute(GetItemListRoute()+"?").
		BuildParam("aid", "1988").
		BuildParam("app_language", "zh-Hans").
		BuildParam("app_name", "tiktok_web").
		BuildParam("browser_language", "zh-CN").
		BuildParam("browser_name", "Mozilla").
		BuildParam("browser_online", "true").
		BuildParam("browser_platform", "Win32").
		BuildParam("browser_version", consts.BROWSER_VERSION).
		BuildParam("channel", "tiktok_web").
		BuildParam("cookie_enabled", "true").
		BuildParam("coverFormat", "2").
		BuildParam("data_collection_enabled", "false").
		BuildParam("device_id", "7419900769212581419").
		BuildParam("device_platform", "web_pc").
		BuildParam("focus_state", "true").
		BuildParam("from_page", "user").
		BuildParam("is_fullscreen", "false").
		BuildParam("is_page_visible", "true").
		BuildParam("language", "zh-Hans").
		BuildParam("needPinnedItemIds", "true").
		BuildParam("odinId", "7419900785709319210").
		BuildParam("os", "windows").
		BuildParam("region", "US").
		BuildParam("screen_height", "1152").
		BuildParam("screen_width", "2048").
		BuildParam("tz_name", "Asia/Shanghai").
		BuildParam("secUid", secUid).
		BuildParam("count", "35").
		BuildParam("cursor", cursor).
		BuildMsTokenRoute(msToken).
		BuildXBogusRoute(ctx, consts.USER_AGENT_DEFAULT_VALUE)
	return string(fullRoute), err
}

func BuildItemListEndpoint(ctx context.Context, msToken, secUid, cursor string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, ItemListRes, error) {
	res := ItemListRes{}
	route, err := BuildItemListRoute(ctx, msToken, secUid, cursor)
	return ITEM_LIST_METHOD, route, make(map[string]interface{}), make(map[string]interface{}), res, err
}

type ItemListRes struct {
	Cursor string `json:"cursor"`
	Extra  struct {
		FatalItemIds []interface{} `json:"fatal_item_ids"`
		Logid        string        `json:"logid"`
		Now          int64         `json:"now"`
	} `json:"extra"`
	HasMore  bool `json:"hasMore"`
	ItemList []struct {
		AIGCDescription string `json:"AIGCDescription"`
		CategoryType    int    `json:"CategoryType"`
		HasPromoteEntry int    `json:"HasPromoteEntry"`
		Author          struct {
			AvatarLarger    string `json:"avatarLarger"`
			AvatarMedium    string `json:"avatarMedium"`
			AvatarThumb     string `json:"avatarThumb"`
			CommentSetting  int    `json:"commentSetting"`
			DownloadSetting int    `json:"downloadSetting"`
			DuetSetting     int    `json:"duetSetting"`
			Ftc             bool   `json:"ftc"`
			Id              string `json:"id"`
			IsADVirtual     bool   `json:"isADVirtual"`
			IsEmbedBanned   bool   `json:"isEmbedBanned"`
			Nickname        string `json:"nickname"`
			OpenFavorite    bool   `json:"openFavorite"`
			PrivateAccount  bool   `json:"privateAccount"`
			Relation        int    `json:"relation"`
			SecUid          string `json:"secUid"`
			Secret          bool   `json:"secret"`
			Signature       string `json:"signature"`
			StitchSetting   int    `json:"stitchSetting"`
			UniqueId        string `json:"uniqueId"`
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
			Id            string `json:"id"`
			ProfileLarger string `json:"profileLarger"`
			ProfileMedium string `json:"profileMedium"`
			ProfileThumb  string `json:"profileThumb"`
			Title         string `json:"title"`
		} `json:"challenges"`
		Collected bool `json:"collected"`
		Contents  []struct {
			Desc      string `json:"desc"`
			TextExtra []struct {
				AwemeId     string `json:"awemeId"`
				End         int    `json:"end"`
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
		DiversificationId int    `json:"diversificationId"`
		DuetDisplay       int    `json:"duetDisplay"`
		DuetEnabled       bool   `json:"duetEnabled"`
		ForFriend         bool   `json:"forFriend"`
		Id                string `json:"id"`
		IsAd              bool   `json:"isAd"`
		ItemCommentStatus int    `json:"itemCommentStatus"`
		ItemControl       struct {
			CanRepost          bool `json:"can_repost"`
			CanComment         bool `json:"can_comment,omitempty"`
			CanCreatorRedirect bool `json:"can_creator_redirect,omitempty"`
			CanMusicRedirect   bool `json:"can_music_redirect,omitempty"`
			CanShare           bool `json:"can_share,omitempty"`
		} `json:"item_control"`
		Music struct {
			AuthorName    string `json:"authorName"`
			CoverLarge    string `json:"coverLarge"`
			CoverMedium   string `json:"coverMedium"`
			CoverThumb    string `json:"coverThumb"`
			Duration      int    `json:"duration"`
			Id            string `json:"id"`
			IsCopyrighted bool   `json:"isCopyrighted"`
			Original      bool   `json:"original"`
			PlayUrl       string `json:"playUrl"`
			Private       bool   `json:"private"`
			Title         string `json:"title"`
			Tt2Dsp        struct {
			} `json:"tt2dsp"`
			Album string `json:"album,omitempty"`
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
		StitchDisplay int  `json:"stitchDisplay"`
		StitchEnabled bool `json:"stitchEnabled"`
		TextExtra     []struct {
			AwemeId     string `json:"awemeId"`
			End         int    `json:"end"`
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
					Uri      string   `json:"Uri"`
					UrlKey   string   `json:"UrlKey"`
					UrlList  []string `json:"UrlList"`
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
			Id            string `json:"id"`
			OriginCover   string `json:"originCover"`
			PlayAddr      string `json:"playAddr"`
			Ratio         string `json:"ratio"`
			Size          int    `json:"size"`
			VideoID       string `json:"videoID"`
			VideoQuality  string `json:"videoQuality"`
			VolumeInfo    struct {
				Loudness float64 `json:"Loudness"`
				Peak     float64 `json:"Peak"`
			} `json:"volumeInfo"`
			Width     int `json:"width"`
			ZoomCover struct {
				Field1 string `json:"240"`
				Field2 string `json:"480"`
				Field3 string `json:"720"`
				Field4 string `json:"960"`
			} `json:"zoomCover"`
		} `json:"video"`
		StickersOnItem []struct {
			StickerText []string `json:"stickerText"`
			StickerType int      `json:"stickerType"`
		} `json:"stickersOnItem,omitempty"`
		AdAuthorization bool `json:"adAuthorization,omitempty"`
	} `json:"itemList"`
	LogPb struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	StatusCode  int    `json:"statusCode"`
	StatusCode1 int    `json:"status_code"`
	StatusMsg   string `json:"status_msg"`
}
