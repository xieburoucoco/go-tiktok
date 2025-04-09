package endpoint

import (
	"fmt"
	"go-tiktok/consts"
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
	params["count"] = 30
	params["maxCursor"] = maxCursor
	params["minCursor"] = minCursor
	return params, nil
}

func BuildUserListEndpoint(scene SceneType, secUid, maxCursor, minCursor string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, UserListRes, error) {
	res := UserListRes{}
	cookies := make(map[string]interface{})
	params, err := GetUserListParams(scene, secUid, maxCursor, minCursor)
	if err != nil {
		return USER_LIST_METHOD, GetUserListRoute(), params, cookies, res, err
	}
	return USER_LIST_METHOD, GetUserListRoute(), params, cookies, res, nil
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
