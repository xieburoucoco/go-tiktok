package endpoint

import "go-tiktok/consts"

const (
	MUSIC_ITEM_LIST_ENDPOINT_NAME = "musicItemList"
	MUSIC_ITEM_LIST_METHOD        = consts.GET
)

func GetMusicItemListlRoute() string {
	return consts.API_ENDPOINT + "music/item_list/"
}

func GetMusicItemListlParams(musicId string, cursor string) map[string]interface{} {
	params := make(map[string]interface{})
	params["musicID"] = musicId
	params["cursor"] = cursor
	params["count"] = 30
	params["coverFormat"] = 2
	params["from_page"] = "music"
	return params
}

func BuildMusicItemListlEndpoint(musicId string, cursor string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, MusicItemListlRes, error) {
	res := MusicItemListlRes{}
	return MUSIC_ITEM_LIST_METHOD, GetMusicItemListlRoute(), GetMusicItemListlParams(musicId, cursor), make(map[string]interface{}), res, nil
}

type MusicItemListlRes struct {
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
