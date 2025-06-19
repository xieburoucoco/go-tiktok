package work

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/xieburoucoco/go-tiktok/consts/endpoint"
	"github.com/xieburoucoco/go-tiktok/logic"
	"os"
	"strconv"
	"testing"
)

// Get the user's attention as well as a list of followers. example: https://www.tiktok.com/api/user/list
func TestUserList(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "" // See readme.md file for how to obtain the msToken // See readme.md file for how to obtain the msToken
	secUid := ""  // See readme.md file for how to obtain the secUid
	minCursor := ""
	currentCursor := "0"
	total := make([]logic.TiktokProfile, 0)
	for minCursor != currentCursor {
		const maxRetries = 3
		var res endpoint.UserListRes
		var err error
		for i := 0; i < maxRetries; i++ {
			_, _, res, err = api.UserList(ctx, endpoint.Following, secUid, "0", currentCursor, msToken, "")
			if err == nil {
				break // 没有错误，跳出循环
			}
			if i == maxRetries-1 {
				t.Errorf("UserList failed after %d retries: %v", maxRetries, err)
				return
			}
			fmt.Printf("UserList failed, retrying (%d/%d): %v\n", i+1, maxRetries, err)
		}
		profiles, cursor, err := logic.NewSTiktokUtil().UnmarshalUserList(res)
		if err != nil {
			t.Error(err)
			return
		}
		total = append(total, profiles...)

		minCursor = currentCursor
		currentCursor = strconv.Itoa(cursor)
		fmt.Println("爬取数量", len(profiles), "记录总数", len(total), "当前游标", currentCursor, "最小游标", minCursor, "总数", res.Total)
		t.Log(cursor)
		if cursor == -1 {
			break
		}
	}

	// 将 total 数据写入 CSV 文件
	file, err := os.Create("./demo.csv")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 写入 CSV 文件的表头
	headers := []string{
		"用户名唯一标志", "用户主页id", "在线链接", "粉丝数", "关注数",
		"点赞", "点赞", "视频个数", "点赞", "朋友数",
	}
	if err := writer.Write(headers); err != nil {
		t.Error(err)
		return
	}

	// 写入数据行
	for _, profile := range total {
		record := []string{
			profile.UniqueId, profile.SecUid, profile.OnlineUrl,
			strconv.Itoa(profile.FollowerCount), strconv.Itoa(profile.FollowingCount),
			strconv.Itoa(profile.Heart), strconv.Itoa(profile.HeartCount),
			strconv.Itoa(profile.VideoCount), strconv.Itoa(profile.DiggCount),
			strconv.Itoa(profile.FriendCount),
		}
		if err := writer.Write(record); err != nil {
			t.Error(err)
			return
		}
	}
}
