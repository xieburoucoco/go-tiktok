package endpoint

import (
	"context"
	"github.com/xieburoucoco/go-tiktok/consts"
)

const (
	COMMENT_ENDPOINT_NAME = "Comment"
	COMMENT_METHOD        = consts.GET
)

func GetCommentRoute() string {
	return consts.API_ENDPOINT + "comment/list/"
}

func BuildCommentRoute(ctx context.Context, msToken, awemeId, cursor string) (string, error) {
	buildUrl := GetCommentRoute() + "?"
	fullRoute, err := consts.NewSRoute(buildUrl).
		BuildParam("WebIdLastTime", "1738896054").
		BuildParam("aid", "1988").
		BuildParam("app_language", "ja-JP").
		BuildParam("app_name", "tiktok_web").
		BuildParam("aweme_id", awemeId).
		BuildParam("browser_language", "zh-CN").
		BuildParam("browser_name", "Mozilla").
		BuildParam("browser_online", "true").
		BuildParam("browser_platform", "Win32").
		BuildParamReplace("browser_version", consts.BROWSER_VERSION, "+", "%20").
		//BuildParam("browser_version", consts.BROWSER_VERSION).
		BuildParam("channel", "tiktok_web").
		BuildParam("cookie_enabled", "true").
		BuildParam("count", "20").
		BuildParam("current_region", "JP").
		BuildParam("cursor", cursor).
		BuildParam("data_collection_enabled", "true").
		BuildParam("device_id", "7468501464287299079").
		BuildParam("device_platform", "web_pc").
		BuildParam("enter_from", "tiktok_web").
		BuildParam("focus_state", "false").
		BuildParam("fromWeb", "1").
		BuildParam("from_page", "video").
		BuildParam("history_len", "3").
		BuildParam("is_fullscreen", "false").
		BuildParam("is_non_personalized", "false").
		BuildParam("is_page_visible", "true").
		BuildParam("os", "windows").
		BuildParam("priority_region", "").
		BuildParam("referer", "").
		BuildParam("region", "US").
		BuildParam("screen_height", "1152").
		BuildParam("screen_width", "2048").
		BuildParam("tz_name", "Asia/Shanghai").
		BuildParam("webcast_language", "zh-Hans").
		BuildMsTokenRoute(msToken).
		BuildXBogusRoute(ctx, consts.USER_AGENT_DEFAULT_VALUE)
	return string(fullRoute), err
}

func BuildCommentEndpoint(ctx context.Context, msToken, awemeId, cursor string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, CommentRes, error) {
	res := CommentRes{}
	route, err := BuildCommentRoute(ctx, msToken, awemeId, cursor)
	return COMMENT_METHOD, route, make(map[string]interface{}), make(map[string]interface{}), res, err
}

type CommentRes struct {
	AliasCommentDeleted bool `json:"alias_comment_deleted"`
	Comments            []struct {
		AuthorPin             bool        `json:"author_pin"`
		AwemeId               string      `json:"aweme_id"`
		Cid                   string      `json:"cid"`
		CollectStat           int         `json:"collect_stat"`
		CommentLanguage       string      `json:"comment_language"`
		CommentPostItemIds    interface{} `json:"comment_post_item_ids"`
		CreateTime            int         `json:"create_time"`
		DiggCount             int         `json:"digg_count"`
		ForbidReplyWithVideo  bool        `json:"forbid_reply_with_video"`
		ImageList             interface{} `json:"image_list"`
		IsAuthorDigged        bool        `json:"is_author_digged"`
		IsCommentTranslatable bool        `json:"is_comment_translatable"`
		LabelList             interface{} `json:"label_list"`
		NoShow                bool        `json:"no_show"`
		ReplyComment          interface{} `json:"reply_comment"`
		ReplyCommentTotal     int         `json:"reply_comment_total"`
		ReplyId               string      `json:"reply_id"`
		ReplyToReplyId        string      `json:"reply_to_reply_id"`
		ShareInfo             struct {
			Acl struct {
				Code  int    `json:"code"`
				Extra string `json:"extra"`
			} `json:"acl"`
			Desc  string `json:"desc"`
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"share_info"`
		SortExtraScore struct {
			ReplyScore    float64 `json:"reply_score"`
			ShowMoreScore float64 `json:"show_more_score"`
		} `json:"sort_extra_score"`
		SortTags      string        `json:"sort_tags"`
		Status        int           `json:"status"`
		StickPosition int           `json:"stick_position"`
		Text          string        `json:"text"`
		TextExtra     []interface{} `json:"text_extra"`
		TransBtnStyle int           `json:"trans_btn_style"`
		User          struct {
			AccountLabels           interface{} `json:"account_labels"`
			AdCoverUrl              interface{} `json:"ad_cover_url"`
			AdvanceFeatureItemOrder interface{} `json:"advance_feature_item_order"`
			AdvancedFeatureInfo     interface{} `json:"advanced_feature_info"`
			AvatarThumb             struct {
				Uri       string      `json:"uri"`
				UrlList   []string    `json:"url_list"`
				UrlPrefix interface{} `json:"url_prefix"`
			} `json:"avatar_thumb"`
			BoldFields                 interface{} `json:"bold_fields"`
			CanMessageFollowStatusList interface{} `json:"can_message_follow_status_list"`
			CanSetGeofencing           interface{} `json:"can_set_geofencing"`
			ChaList                    interface{} `json:"cha_list"`
			CoverUrl                   interface{} `json:"cover_url"`
			CustomVerify               string      `json:"custom_verify"`
			EnterpriseVerifyReason     string      `json:"enterprise_verify_reason"`
			Events                     interface{} `json:"events"`
			FollowersDetail            interface{} `json:"followers_detail"`
			Geofencing                 interface{} `json:"geofencing"`
			HomepageBottomToast        interface{} `json:"homepage_bottom_toast"`
			ItemList                   interface{} `json:"item_list"`
			MutualRelationAvatars      interface{} `json:"mutual_relation_avatars"`
			NeedPoints                 interface{} `json:"need_points"`
			Nickname                   string      `json:"nickname"`
			PlatformSyncInfo           interface{} `json:"platform_sync_info"`
			PredictedAgeGroup          string      `json:"predicted_age_group"`
			RelativeUsers              interface{} `json:"relative_users"`
			SearchHighlight            interface{} `json:"search_highlight"`
			SecUid                     string      `json:"sec_uid"`
			ShieldEditFieldInfo        interface{} `json:"shield_edit_field_info"`
			TypeLabel                  interface{} `json:"type_label"`
			Uid                        string      `json:"uid"`
			UniqueId                   string      `json:"unique_id"`
			UserProfileGuide           interface{} `json:"user_profile_guide"`
			UserTags                   interface{} `json:"user_tags"`
			WhiteCoverUrl              interface{} `json:"white_cover_url"`
		} `json:"user"`
		UserBuried bool `json:"user_buried"`
		UserDigged int  `json:"user_digged"`
	} `json:"comments"`
	Cursor int `json:"cursor"`
	Extra  struct {
		ApiDebugInfo interface{} `json:"api_debug_info"`
		FatalItemIds interface{} `json:"fatal_item_ids"`
		Now          int64       `json:"now"`
	} `json:"extra"`
	HasFilteredComments int `json:"has_filtered_comments"`
	HasMore             int `json:"has_more"`
	LogPb               struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	ReplyStyle int         `json:"reply_style"`
	StatusCode int         `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	TopGifts   interface{} `json:"top_gifts"`
	Total      int         `json:"total"`
}
