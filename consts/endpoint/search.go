package endpoint

import (
	"fmt"
	"github.com/xieburoucoco/go-tiktok/consts"
)

type TypeSearchEndpointName string

const (
	SEARCH_ENDPOINT_NAME                        = "search"
	SEARCH_METHOD                               = consts.GET
	GENERAL              TypeSearchEndpointName = "general"
	USER                 TypeSearchEndpointName = "user"
	ITEM                 TypeSearchEndpointName = "item"
	LIVE                 TypeSearchEndpointName = "live"
)

func GetSearchRoute(endpointName TypeSearchEndpointName) string {
	return consts.API_ENDPOINT + "search/" + string(endpointName) + "/full/"
}

func GetSearchParams(keyword string, offset string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if len(keyword) == 0 {
		return params, fmt.Errorf("keyword is empty")
	}
	if len(offset) == 0 {
		offset = "0"
	}
	params["keyword"] = keyword
	params["offset"] = offset
	return params, nil
}

func BuildSearchEndpoint(endpointName TypeSearchEndpointName, keyword string, offset string, ttwid string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, SearchRes, error) {
	res := SearchRes{}
	params, err := GetSearchParams(keyword, offset)
	if err != nil {
		return SEARCH_METHOD, GetSearchRoute(endpointName), params, nil, res, err
	}
	cookies := make(map[string]interface{})
	cookies["ttwid"] = ttwid
	return SEARCH_METHOD, GetSearchRoute(endpointName), params, cookies, res, nil
}

type SearchRes struct {
	StatusCode int `json:"status_code"`
	Data       []struct {
		LiveInfo struct {
			RawData  string `json:"raw_data"`
			RoomInfo struct {
				HasCommerceGoods bool `json:"has_commerce_goods"`
				IsBattle         bool `json:"is_battle"`
			} `json:"room_info"`
		} `json:"live_info"`
		Type int `json:"type"`
		Item struct {
			Id         string `json:"id"`
			Desc       string `json:"desc"`
			CreateTime int    `json:"createTime"`
			Video      struct {
				Id            string   `json:"id"`
				Height        int      `json:"height"`
				Width         int      `json:"width"`
				Duration      int      `json:"duration"`
				Ratio         string   `json:"ratio"`
				Cover         string   `json:"cover"`
				OriginCover   string   `json:"originCover"`
				DynamicCover  string   `json:"dynamicCover"`
				PlayAddr      string   `json:"playAddr"`
				DownloadAddr  string   `json:"downloadAddr,omitempty"`
				ShareCover    []string `json:"shareCover"`
				ReflowCover   string   `json:"reflowCover"`
				Bitrate       int      `json:"bitrate"`
				EncodedType   string   `json:"encodedType"`
				Format        string   `json:"format"`
				VideoQuality  string   `json:"videoQuality"`
				EncodeUserTag string   `json:"encodeUserTag"`
			} `json:"video"`
			Author struct {
				Id              string `json:"id"`
				UniqueId        string `json:"uniqueId"`
				Nickname        string `json:"nickname"`
				AvatarThumb     string `json:"avatarThumb"`
				AvatarMedium    string `json:"avatarMedium"`
				AvatarLarger    string `json:"avatarLarger"`
				Signature       string `json:"signature"`
				Verified        bool   `json:"verified"`
				SecUid          string `json:"secUid"`
				Secret          bool   `json:"secret"`
				Ftc             bool   `json:"ftc"`
				Relation        int    `json:"relation"`
				OpenFavorite    bool   `json:"openFavorite"`
				CommentSetting  int    `json:"commentSetting"`
				DuetSetting     int    `json:"duetSetting"`
				StitchSetting   int    `json:"stitchSetting"`
				PrivateAccount  bool   `json:"privateAccount"`
				DownloadSetting int    `json:"downloadSetting"`
			} `json:"author"`
			Music struct {
				Id          string `json:"id"`
				Title       string `json:"title"`
				PlayUrl     string `json:"playUrl"`
				CoverThumb  string `json:"coverThumb"`
				CoverMedium string `json:"coverMedium"`
				CoverLarge  string `json:"coverLarge"`
				AuthorName  string `json:"authorName"`
				Original    bool   `json:"original"`
				Duration    int    `json:"duration"`
				Album       string `json:"album"`
			} `json:"music"`
			Challenges []struct {
				Id            string `json:"id"`
				Title         string `json:"title"`
				Desc          string `json:"desc"`
				ProfileThumb  string `json:"profileThumb"`
				ProfileMedium string `json:"profileMedium"`
				ProfileLarger string `json:"profileLarger"`
				CoverThumb    string `json:"coverThumb"`
				CoverMedium   string `json:"coverMedium"`
				CoverLarger   string `json:"coverLarger"`
				IsCommerce    bool   `json:"isCommerce"`
			} `json:"challenges"`
			Stats struct {
				DiggCount    int `json:"diggCount"`
				ShareCount   int `json:"shareCount"`
				CommentCount int `json:"commentCount"`
				PlayCount    int `json:"playCount"`
				CollectCount int `json:"collectCount"`
			} `json:"stats"`
			DuetInfo struct {
				DuetFromId string `json:"duetFromId"`
			} `json:"duetInfo"`
			OriginalItem bool `json:"originalItem"`
			OfficalItem  bool `json:"officalItem"`
			TextExtra    []struct {
				AwemeId      string `json:"awemeId"`
				Start        int    `json:"start"`
				End          int    `json:"end"`
				HashtagName  string `json:"hashtagName"`
				HashtagId    string `json:"hashtagId"`
				Type         int    `json:"type"`
				UserId       string `json:"userId"`
				IsCommerce   bool   `json:"isCommerce"`
				UserUniqueId string `json:"userUniqueId"`
				SecUid       string `json:"secUid"`
				SubType      int    `json:"subType"`
			} `json:"textExtra"`
			Secret            bool `json:"secret"`
			ForFriend         bool `json:"forFriend"`
			Digged            bool `json:"digged"`
			ItemCommentStatus int  `json:"itemCommentStatus"`
			ShowNotPass       bool `json:"showNotPass"`
			Vl1               bool `json:"vl1"`
			ItemMute          bool `json:"itemMute"`
			AuthorStats       struct {
				FollowingCount int `json:"followingCount"`
				FollowerCount  int `json:"followerCount"`
				HeartCount     int `json:"heartCount"`
				VideoCount     int `json:"videoCount"`
				DiggCount      int `json:"diggCount"`
				Heart          int `json:"heart"`
			} `json:"authorStats"`
			PrivateItem    bool `json:"privateItem"`
			DuetEnabled    bool `json:"duetEnabled"`
			StitchEnabled  bool `json:"stitchEnabled"`
			ShareEnabled   bool `json:"shareEnabled"`
			IsAd           bool `json:"isAd"`
			Collected      bool `json:"collected"`
			EffectStickers []struct {
				Name string `json:"name"`
				ID   string `json:"ID"`
			} `json:"effectStickers,omitempty"`
			StickersOnItem []struct {
				StickerType int      `json:"stickerType"`
				StickerText []string `json:"stickerText"`
			} `json:"stickersOnItem,omitempty"`
			Anchors []struct {
				Id      string `json:"id"`
				Type    int    `json:"type"`
				Keyword string `json:"keyword"`
				Icon    struct {
					UrlList []string `json:"urlList"`
				} `json:"icon"`
				Description string `json:"description"`
				Thumbnail   struct {
					UrlList []string `json:"urlList"`
					Width   int      `json:"width"`
					Height  int      `json:"height"`
				} `json:"thumbnail"`
				ExtraInfo struct {
					Subtype string `json:"subtype"`
				} `json:"extraInfo"`
				Schema   string `json:"schema,omitempty"`
				LogExtra string `json:"logExtra,omitempty"`
			} `json:"anchors,omitempty"`
		} `json:"item"`
		Common struct {
			DocIdStr string `json:"doc_id_str"`
		} `json:"common"`
	} `json:"data"` // /search/general/full/ or /search/live/full/
	UserList []struct {
		UserInfo struct {
			Uid         string `json:"uid"`
			Nickname    string `json:"nickname"`
			Signature   string `json:"signature"`
			AvatarThumb struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				Width     int         `json:"width"`
				Height    int         `json:"height"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb"`
			FollowStatus               int         `json:"follow_status"`
			FollowerCount              int         `json:"follower_count"`
			CustomVerify               string      `json:"custom_verify"`
			UniqueId                   string      `json:"unique_id"`
			RoomId                     int         `json:"room_id"`
			EnterpriseVerifyReason     string      `json:"enterprise_verify_reason"`
			FollowersDetail            interface{} `json:"followers_detail"`
			PlatformSyncInfo           interface{} `json:"platform_sync_info"`
			Geofencing                 interface{} `json:"geofencing"`
			CoverUrl                   interface{} `json:"cover_url"`
			ItemList                   interface{} `json:"item_list"`
			TypeLabel                  interface{} `json:"type_label"`
			AdCoverUrl                 interface{} `json:"ad_cover_url"`
			RelativeUsers              interface{} `json:"relative_users"`
			ChaList                    interface{} `json:"cha_list"`
			SecUid                     string      `json:"sec_uid"`
			NeedPoints                 interface{} `json:"need_points"`
			HomepageBottomToast        interface{} `json:"homepage_bottom_toast"`
			CanSetGeofencing           interface{} `json:"can_set_geofencing"`
			WhiteCoverUrl              interface{} `json:"white_cover_url"`
			UserTags                   interface{} `json:"user_tags"`
			BoldFields                 interface{} `json:"bold_fields"`
			SearchHighlight            interface{} `json:"search_highlight"`
			MutualRelationAvatars      interface{} `json:"mutual_relation_avatars"`
			RoomIdStr                  string      `json:"room_id_str"`
			Events                     interface{} `json:"events"`
			AdvanceFeatureItemOrder    interface{} `json:"advance_feature_item_order"`
			AdvancedFeatureInfo        interface{} `json:"advanced_feature_info"`
			UserProfileGuide           interface{} `json:"user_profile_guide"`
			ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
			CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
			AccountLabels              interface{} `json:"account_labels"`
		} `json:"user_info"`
		Position       interface{} `json:"position"`
		UniqidPosition interface{} `json:"uniqid_position"`
		Effects        interface{} `json:"effects"`
		Musics         interface{} `json:"musics"`
		Items          interface{} `json:"items"`
		MixList        interface{} `json:"mix_list"`
		Challenges     interface{} `json:"challenges"`
	} `json:"user_list"` // /search/user/full/
	Cursor  int `json:"cursor"`
	HasMore int `json:"has_more"`
	AdInfo  struct {
	} `json:"ad_info"`
	Extra struct {
		Now             int64         `json:"now"`
		Logid           string        `json:"logid"`
		FatalItemIds    []interface{} `json:"fatal_item_ids"`
		SearchRequestId string        `json:"search_request_id"`
		ApiDebugInfo    interface{}   `json:"api_debug_info"`
	} `json:"extra"`
	LogPb struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	GlobalDoodleConfig struct {
		FeedbackSurvey interface{} `json:"feedback_survey"`
	} `json:"global_doodle_config"`
	Backtrace     string      `json:"backtrace"`
	Type          int         `json:"type"`
	ChallengeList interface{} `json:"challenge_list"`
	MusicList     interface{} `json:"music_list"`
	Qc            string      `json:"qc"`
	Rid           string      `json:"rid"`
	InputKeyword  string      `json:"input_keyword"`
	FeedbackType  string      `json:"feedback_type"`
	ItemList      []struct {
		Id         string `json:"id"`
		Desc       string `json:"desc"`
		CreateTime int    `json:"createTime"`
		Video      struct {
			Id            string   `json:"id"`
			Height        int      `json:"height"`
			Width         int      `json:"width"`
			Duration      int      `json:"duration"`
			Ratio         string   `json:"ratio"`
			Cover         string   `json:"cover"`
			OriginCover   string   `json:"originCover"`
			DynamicCover  string   `json:"dynamicCover"`
			PlayAddr      string   `json:"playAddr"`
			DownloadAddr  string   `json:"downloadAddr,omitempty"`
			ShareCover    []string `json:"shareCover"`
			ReflowCover   string   `json:"reflowCover"`
			Bitrate       int      `json:"bitrate"`
			EncodedType   string   `json:"encodedType"`
			Format        string   `json:"format"`
			VideoQuality  string   `json:"videoQuality"`
			EncodeUserTag string   `json:"encodeUserTag"`
		} `json:"video"`
		Author struct {
			Id              string `json:"id"`
			UniqueId        string `json:"uniqueId"`
			Nickname        string `json:"nickname"`
			AvatarThumb     string `json:"avatarThumb"`
			AvatarMedium    string `json:"avatarMedium"`
			AvatarLarger    string `json:"avatarLarger"`
			Signature       string `json:"signature"`
			Verified        bool   `json:"verified"`
			SecUid          string `json:"secUid"`
			Secret          bool   `json:"secret"`
			Ftc             bool   `json:"ftc"`
			Relation        int    `json:"relation"`
			OpenFavorite    bool   `json:"openFavorite"`
			CommentSetting  int    `json:"commentSetting"`
			DuetSetting     int    `json:"duetSetting"`
			StitchSetting   int    `json:"stitchSetting"`
			PrivateAccount  bool   `json:"privateAccount"`
			DownloadSetting int    `json:"downloadSetting"`
		} `json:"author"`
		Music struct {
			Id          string `json:"id"`
			Title       string `json:"title"`
			PlayUrl     string `json:"playUrl"`
			CoverThumb  string `json:"coverThumb"`
			CoverMedium string `json:"coverMedium"`
			CoverLarge  string `json:"coverLarge"`
			AuthorName  string `json:"authorName"`
			Original    bool   `json:"original"`
			Duration    int    `json:"duration"`
			Album       string `json:"album"`
		} `json:"music"`
		Challenges []struct {
			Id            string `json:"id"`
			Title         string `json:"title"`
			Desc          string `json:"desc"`
			ProfileThumb  string `json:"profileThumb"`
			ProfileMedium string `json:"profileMedium"`
			ProfileLarger string `json:"profileLarger"`
			CoverThumb    string `json:"coverThumb"`
			CoverMedium   string `json:"coverMedium"`
			CoverLarger   string `json:"coverLarger"`
			IsCommerce    bool   `json:"isCommerce"`
		} `json:"challenges,omitempty"`
		Stats struct {
			DiggCount    int `json:"diggCount"`
			ShareCount   int `json:"shareCount"`
			CommentCount int `json:"commentCount"`
			PlayCount    int `json:"playCount"`
			CollectCount int `json:"collectCount"`
		} `json:"stats"`
		DuetInfo struct {
			DuetFromId string `json:"duetFromId"`
		} `json:"duetInfo"`
		OriginalItem bool `json:"originalItem"`
		OfficalItem  bool `json:"officalItem"`
		TextExtra    []struct {
			AwemeId      string `json:"awemeId"`
			Start        int    `json:"start"`
			End          int    `json:"end"`
			HashtagName  string `json:"hashtagName"`
			HashtagId    string `json:"hashtagId"`
			Type         int    `json:"type"`
			UserId       string `json:"userId"`
			IsCommerce   bool   `json:"isCommerce"`
			UserUniqueId string `json:"userUniqueId"`
			SecUid       string `json:"secUid"`
			SubType      int    `json:"subType"`
		} `json:"textExtra,omitempty"`
		Secret            bool `json:"secret"`
		ForFriend         bool `json:"forFriend"`
		Digged            bool `json:"digged"`
		ItemCommentStatus int  `json:"itemCommentStatus"`
		ShowNotPass       bool `json:"showNotPass"`
		Vl1               bool `json:"vl1"`
		ItemMute          bool `json:"itemMute"`
		AuthorStats       struct {
			FollowingCount int `json:"followingCount"`
			FollowerCount  int `json:"followerCount"`
			HeartCount     int `json:"heartCount"`
			VideoCount     int `json:"videoCount"`
			DiggCount      int `json:"diggCount"`
			Heart          int `json:"heart"`
		} `json:"authorStats"`
		PrivateItem    bool `json:"privateItem"`
		DuetEnabled    bool `json:"duetEnabled"`
		StitchEnabled  bool `json:"stitchEnabled"`
		ShareEnabled   bool `json:"shareEnabled"`
		StickersOnItem []struct {
			StickerType int      `json:"stickerType"`
			StickerText []string `json:"stickerText"`
		} `json:"stickersOnItem,omitempty"`
		IsAd           bool `json:"isAd"`
		Collected      bool `json:"collected"`
		EffectStickers []struct {
			Name string `json:"name"`
			ID   string `json:"ID"`
		} `json:"effectStickers,omitempty"`
		Anchors []struct {
			Id      string `json:"id"`
			Type    int    `json:"type"`
			Keyword string `json:"keyword"`
			Icon    struct {
				UrlList []string `json:"urlList"`
			} `json:"icon"`
			Description string `json:"description"`
			Thumbnail   struct {
				UrlList []string `json:"urlList"`
				Width   int      `json:"width"`
				Height  int      `json:"height"`
			} `json:"thumbnail"`
			ExtraInfo struct {
				Subtype string `json:"subtype"`
			} `json:"extraInfo"`
			Schema   string `json:"schema,omitempty"`
			LogExtra string `json:"logExtra,omitempty"`
		} `json:"anchors,omitempty"`
		VideoSuggestWordsList struct {
			VideoSuggestWordsStruct []struct {
				Words []struct {
					Word   string `json:"word"`
					WordId string `json:"word_id"`
				} `json:"words"`
				Scene    string `json:"scene"`
				HintText string `json:"hint_text"`
			} `json:"video_suggest_words_struct"`
		} `json:"videoSuggestWordsList,omitempty"`
	} `json:"item_list"` // /search/item/full/
}
