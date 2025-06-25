package endpoint

import (
	"encoding/json"
	"fmt"
	"github.com/xieburoucoco/go-tiktok/consts"
	"net/url"
)

const (
	PROFILE_ENDPOINT_NAME = "Profile"
	PROFILE_METHOD        = consts.GET
)

func GetProfileRoute() string {
	return consts.HOME + "/tiktok/v1/im/user/profile/"
}

func BuildProfileRoute(userIds []string) (string, error) {
	if len(userIds) == 0 {
		return "", fmt.Errorf("uid is empty")
	}

	// 基础 URL
	baseURL := GetProfileRoute()

	// 将 uid 切片编码为 JSON 数组字符串
	userIDsJSON, err := json.Marshal(userIds)
	if err != nil {
		return "", fmt.Errorf("failed to marshal user_ids: %v", err)
	}

	// 使用 url.Values 构建查询参数
	params := url.Values{}
	params.Add("aid", "1988")
	params.Add("user_ids", string(userIDsJSON))
	fullUrl := baseURL + "?" + params.Encode()
	// 构造完整 URL
	return fullUrl, nil
}

func GetProfileParams(scene SceneType, secUid, maxCursor, minCursor string) (map[string]interface{}, error) {
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

func BuildProfileEndpoint(sessionId, ttTargetIdc string, userIds []string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, ProfileRes, error) {
	res := ProfileRes{}
	cookies := map[string]interface{}{
		"sessionid":     sessionId,
		"tt-target-idc": ttTargetIdc,
	}
	route, err := BuildProfileRoute(userIds)
	return PROFILE_METHOD, route, make(map[string]interface{}), cookies, res, err
}

type ProfileRes struct {
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
	Profile     []struct {
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
	} `json:"Profile"`
}
