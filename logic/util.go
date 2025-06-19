package logic

import (
	"github.com/xieburoucoco/go-tiktok/consts"
	"github.com/xieburoucoco/go-tiktok/consts/endpoint"
)

type ITiktokUtil interface {
	UnmarshalUserList(body endpoint.UserListRes) (profiles []TiktokProfile, nextCursor int, err error)
}

type sTiktokUtil struct {
}

func NewSTiktokUtil() ITiktokUtil {
	return &sTiktokUtil{}
}

func (s *sTiktokUtil) UnmarshalUserList(body endpoint.UserListRes) (profiles []TiktokProfile, nextCursor int, err error) {
	profiles = make([]TiktokProfile, 0)
	for _, elem := range body.UserList {
		temp := TiktokProfile{
			UniqueId:       elem.User.UniqueId,
			SecUid:         elem.User.SecUid,
			OnlineUrl:      consts.HOME + "/@" + elem.User.SecUid,
			FollowerCount:  elem.Stats.FollowerCount,
			FollowingCount: elem.Stats.FollowingCount,
			Heart:          elem.Stats.Heart,
			HeartCount:     elem.Stats.HeartCount,
			VideoCount:     elem.Stats.VideoCount,
			DiggCount:      elem.Stats.DiggCount,
			FriendCount:    elem.Stats.FriendCount,
		}
		profiles = append(profiles, temp)
	}
	nextCursor = body.MinCursor
	return profiles, nextCursor, err
}

type TiktokProfile struct {
	UniqueId       string `json:"uniqueId"          orm:"unique_id"           description:"用户名唯一标志"`
	SecUid         string `json:"secUid"            orm:"sec_uid"             description:"用户主页id"`
	OnlineUrl      string `json:"onlineUrl"         orm:"online_url"          description:"在线链接"`
	FollowerCount  int    `json:"followerCount"     orm:"follower_count"      description:"粉丝数"`
	FollowingCount int    `json:"followingCount"    orm:"following_count"     description:"关注数"`
	Heart          int    `json:"heart"             orm:"heart"               description:"点赞"`
	HeartCount     int    `json:"heartCount"        orm:"heart_count"         description:"点赞"`
	VideoCount     int    `json:"videoCount"        orm:"video_count"         description:"视频个数"`
	DiggCount      int    `json:"diggCount"         orm:"digg_count"          description:"点赞"`
	FriendCount    int    `json:"friendCount"       orm:"friend_count"        description:"朋友数"`
}
