package endpoint

import "go-tiktok/consts"

const (
	USER_DETAIL_ENDPOINT_NAME = "UserDetail"
	USER_DETAIL_METHOD        = consts.GET
)

func GetUserDetailRoute() string {
	return consts.API_ENDPOINT + "user/detail/"
}

func GetUserDetailParams(uniqueId string) map[string]interface{} {
	params := make(map[string]interface{})
	params["uniqueId"] = uniqueId
	return params
}

func BuildUserDetailEndpoint(itemId string) (consts.HTTPMethodType, string, map[string]interface{}, map[string]interface{}, UserDetailRes, error) {
	res := UserDetailRes{}
	return USER_DETAIL_METHOD, GetUserDetailRoute(), GetUserDetailParams(itemId), make(map[string]interface{}), res, nil
}

type UserDetailRes struct {
	Extra struct {
		FatalItemIds []interface{} `json:"fatal_item_ids"`
		Logid        string        `json:"logid"`
		Now          int64         `json:"now"`
	} `json:"extra"`
	LogPb struct {
		ImprId string `json:"impr_id"`
	} `json:"log_pb"`
	ShareMeta struct {
		Desc  string `json:"desc"`
		Title string `json:"title"`
	} `json:"shareMeta"`
	StatusCode  int    `json:"statusCode"`
	StatusCode1 int    `json:"status_code"`
	StatusMsg   string `json:"status_msg"`
	UserInfo    struct {
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
			AvatarLarger     string `json:"avatarLarger"`
			AvatarMedium     string `json:"avatarMedium"`
			AvatarThumb      string `json:"avatarThumb"`
			CanExpPlaylist   bool   `json:"canExpPlaylist"`
			CommentSetting   int    `json:"commentSetting"`
			CommerceUserInfo struct {
				CommerceUser bool `json:"commerceUser"`
			} `json:"commerceUserInfo"`
			DownloadSetting        int    `json:"downloadSetting"`
			DuetSetting            int    `json:"duetSetting"`
			FollowingVisibility    int    `json:"followingVisibility"`
			Ftc                    bool   `json:"ftc"`
			Id                     string `json:"id"`
			IsADVirtual            bool   `json:"isADVirtual"`
			IsEmbedBanned          bool   `json:"isEmbedBanned"`
			Nickname               string `json:"nickname"`
			OpenFavorite           bool   `json:"openFavorite"`
			PrivateAccount         bool   `json:"privateAccount"`
			ProfileEmbedPermission int    `json:"profileEmbedPermission"`
			ProfileTab             struct {
				ShowMusicTab    bool `json:"showMusicTab"`
				ShowPlayListTab bool `json:"showPlayListTab"`
			} `json:"profileTab"`
			Relation      int    `json:"relation"`
			SecUid        string `json:"secUid"`
			Secret        bool   `json:"secret"`
			Signature     string `json:"signature"`
			StitchSetting int    `json:"stitchSetting"`
			TtSeller      bool   `json:"ttSeller"`
			UniqueId      string `json:"uniqueId"`
			Verified      bool   `json:"verified"`
		} `json:"user"`
	} `json:"userInfo"`
}
