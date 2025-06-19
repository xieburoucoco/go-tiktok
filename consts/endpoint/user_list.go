package endpoint

import (
	"fmt"
	"github.com/xieburoucoco/go-tiktok/consts"
	"net/url"
	"strings"
)

const (
	USER_LIST_ENDPOINT_NAME = "UserList"
	USER_LIST_METHOD        = consts.GET
)

type SceneType string

const (
	FollowerSceneType           = "21"
	Following         SceneType = "67"
)

func GetUserListRoute() string {
	return consts.API_ENDPOINT + "user/list/"
}

func BuildUserListRoute(scene SceneType, secUid, maxCursor, minCursor string) (string, error) {
	if len(secUid) == 0 {
		return "", fmt.Errorf("secUid is empty")
	}
	if len(maxCursor) == 0 {
		maxCursor = "0"
	}
	if len(minCursor) == 0 {
		return "", fmt.Errorf("minCursor is empty")
	}

	buildUrl := GetUserListRoute() + "?"
	startUrl := buildUrl
	addAndConcatParam := func(key string, value string) {
		if buildUrl != startUrl {
			buildUrl += "&"
		}
		escapedValue := strings.ReplaceAll(url.QueryEscape(value), "+", "%20")
		//escapedValue := url.QueryEscape(value)
		buildUrl += fmt.Sprintf("%s=%s", key, escapedValue)
	}
	addAndConcatParam("WebIdLastTime", "1738896054")
	addAndConcatParam("aid", "1988")
	addAndConcatParam("app_language", "zh-Hans")
	addAndConcatParam("app_name", "tiktok_web")
	addAndConcatParam("browser_language", "zh-CN")
	addAndConcatParam("browser_name", "Mozilla")
	addAndConcatParam("browser_online", "true")
	addAndConcatParam("browser_platform", "Win32")
	addAndConcatParam("browser_version", consts.BROWSER_VERSION)
	addAndConcatParam("channel", "tiktok_web")
	addAndConcatParam("cookie_enabled", "true")

	addAndConcatParam("count", "30")
	addAndConcatParam("data_collection_enabled", "true")

	addAndConcatParam("device_id", "7468501464287299079")
	addAndConcatParam("device_platform", "web_pc")
	addAndConcatParam("focus_state", "false")
	addAndConcatParam("from_page", "user")
	addAndConcatParam("history_len", "4")
	addAndConcatParam("is_fullscreen", "false")
	addAndConcatParam("is_page_visible", "true")

	addAndConcatParam("maxCursor", maxCursor)
	addAndConcatParam("minCursor", minCursor)

	addAndConcatParam("os", "windows")
	addAndConcatParam("priority_region", "")
	addAndConcatParam("referer", "")
	addAndConcatParam("region", "US")
	addAndConcatParam("scene", string(scene))
	addAndConcatParam("screen_height", "1152")
	addAndConcatParam("screen_width", "2048")
	addAndConcatParam("secUid", secUid)
	addAndConcatParam("tz_name", "Asia/Shanghai")
	addAndConcatParam("webcast_language", "zh-Hans")
	return buildUrl, nil
}

func GetUserListParams(scene SceneType, secUid, maxCursor, minCursor string) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if len(secUid) == 0 {
		return nil, fmt.Errorf("secUid is empty")
	}
	if len(maxCursor) == 0 {
		maxCursor = "0"
	}
	if len(minCursor) == 0 {
		minCursor = "0"
	}
	params["scene"] = scene
	params["secUid"] = secUid
	//params["count"] = "30"
	params["maxCursor"] = maxCursor
	params["minCursor"] = minCursor
	return params, nil
}

func BuildUserListEndpoint(scene SceneType, secUid, maxCursor, minCursor string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, UserListRes, error) {
	res := UserListRes{}
	cookies := make(map[string]interface{})
	route, err := BuildUserListRoute(scene, secUid, maxCursor, minCursor)
	return USER_LIST_METHOD, route, make(map[string]interface{}), cookies, res, err
}

type UserListRes struct {
	Extra struct {
		FatalItemIds []interface{} `json:"fatal_item_ids"`
		Logid        string        `json:"logid"`
		Now          int64         `json:"now"`
	} `json:"extra"`
	HasMore bool `json:"hasMore"`
	LogPb   struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	MinCursor   int    `json:"minCursor"`
	StatusCode  int    `json:"statusCode"`
	StatusCode1 int    `json:"status_code"`
	StatusMsg   string `json:"status_msg"`
	Total       int    `json:"total"`
	UserList    []struct {
		Stats struct {
			DiggCount      int `json:"diggCount"`
			FollowerCount  int `json:"followerCount"`
			FollowingCount int `json:"followingCount"`
			FriendCount    int `json:"friendCount"`
			Heart          int `json:"heart"`
			HeartCount     int `json:"heartCount"`
			VideoCount     int `json:"videoCount"`
		} `json:"stats"`
		User struct {
			AvatarLarger    string `json:"avatarLarger"`
			AvatarMedium    string `json:"avatarMedium"`
			AvatarThumb     string `json:"avatarThumb"`
			DownloadSetting int    `json:"downloadSetting"`
			Ftc             bool   `json:"ftc"`
			Id              string `json:"id"`
			IsADVirtual     bool   `json:"isADVirtual"`
			Nickname        string `json:"nickname"`
			OpenFavorite    bool   `json:"openFavorite"`
			PrivateAccount  bool   `json:"privateAccount"`
			Relation        int    `json:"relation"`
			SecUid          string `json:"secUid"`
			Secret          bool   `json:"secret"`
			Signature       string `json:"signature"`
			TtSeller        bool   `json:"ttSeller"`
			UniqueId        string `json:"uniqueId"`
			Verified        bool   `json:"verified"`
			CommentSetting  int    `json:"commentSetting,omitempty"`
			DuetSetting     int    `json:"duetSetting,omitempty"`
			StitchSetting   int    `json:"stitchSetting,omitempty"`
		} `json:"user"`
	} `json:"userList"`
}
