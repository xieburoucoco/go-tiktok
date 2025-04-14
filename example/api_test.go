package example

import (
	"context"
	"github.com/xieburoucoco/go-tiktok/consts"
	"github.com/xieburoucoco/go-tiktok/consts/endpoint"
	"github.com/xieburoucoco/go-tiktok/logic"
	"testing"
)

func TestCliOp(t *testing.T) {
	ctx := context.Background()
	op := logic.TikTokOption{
		Method: consts.GET,
		URL:    "111",
		Params: map[string]interface{}{
			"1": "123",
		},
		Body:    nil,
		Model:   nil,
		MsToken: "123",
	}
	cli := logic.NewSTikTokCli(op)
	res, err := logic.NewITikTokCli(cli).Do(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

// Get the details of the home page user. example: https://www.tiktok.com/@gemdzq
func TestHome(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	proxyURL := "http://localhost:7897" // input your proxy or empty string
	uniqueId := "gemdzq"                // input your interested uniqueId
	_, body, res, err := api.Home(ctx, uniqueId, proxyURL)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

// Get a list of video metadata . example:  https://www.tiktok.com/api/item/detail/
func TestItemDetail(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "IAf95B5TvP9sbPnY5pocbaybCXzbk2BBGNGSHs3kRbwvDSrttMwr89HjMNNXpO4qPM_G2et3mNcakpGCZMRFsHyCsyVTV1tFwXAPkO2fXclqvdXzV1567clard7P03STvhKk_AdalvcdwXTonPIOJ197Pi8" // See readme.md file for how to obtain the msToken
	_, body, res, err := api.ItemDetail(ctx, "7273529185589562625", msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

// Get a list of video metadata . example:  https://www.tiktok.com/api/post/item_list/
func TestItemList(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "GP0csW6HHSJqLQQpLgcBn4JO0HXbV0MeZIdUgEYYAdqmnVaQaF5no-i0a7hyKWPoyKSsISz3TVBQbDgss_aRfB5Ft2gXLC1mWNsnZUICj_W8k4lizgdZTCdqDEjZDOhaELPeuh_2_3kgPCFO4-ighnSK" // See readme.md file for how to obtain the msToken
	_, body, res, err := api.ItemList(ctx, "MS4wLjABAAAADWVixuGqt-G8FDQ9yx9TLQD-4fFpwQtBhXe6EDCJ32wiprPkgzEzdGCjCR1PEwmf", "0", msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

// Get the Music info list . example: https://www.tiktok.com/api/music/item_list/
func TestMusicItemList(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "sfJ01Kp0B-b9LLB9QRG6qvyuzFOu7j-C1kW78-ESHg603RLe38iXXxkv5R_V0fer0jpwaOrGFFnAH2Db3yGammAR7U5iL9COHiXaetdURouEdriCDwKv8OEE-wjVl1WtQfqbwv--0zI8W9YuWFcdBOxJf80==" // See readme.md file for how to obtain the msToken
	_, body, res, err := api.MusicItemList(ctx, "7196737013066828546", "0", msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

// Get the Comment list . example: https://www.tiktok.com/api/comment/list/
func TestCommentList(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "DNLk-Hr7680j2-yuUy3T9i-oriZ8MDIHNspfoPABV76u9q4rjeoBv-paClkYUOTH3ZCjS3ChO7fdaEoYvMI9Jzq4dPZMjEzobOu2Et-Y_BB_t7fC6bxCwYhQYWD7B4zrw1umopt3f9IHEhbO7E8gVRuNzA==" // See readme.md file for how to obtain the msToken
	_, body, res, err := api.CommentList(ctx, "7339184461612993794", "0", msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

// Get the user's details . example:  https://www.tiktok.com/api/user/detail/
//func TestUserDetail(t *testing.T) {
//	ctx := context.Background()
//	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
//	msToken := "sfJ01Kp0B-b9LLB9QRG6qvyuzFOu7j-C1kW78-ESHg603RLe38iXXxkv5R_V0fer0jpwaOrGFFnAH2Db3yGammAR7U5iL9COHiXaetdURouEdriCDwKv8OEE-wjVl1WtQfqbwv--0zI8W9YuWFcdBOxJf80==" // See readme.md file for how to obtain the msToken
//	_, body, res, err := api.UserDetail(ctx, "gemdzq", msToken, "http://localhost:7897")
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	t.Log(string(body))
//
//	t.Log(res)
//}

// Get the user's attention as well as a list of followers. example: https://www.tiktok.com/api/user/list
func TestUserList(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "dJ78sFyfkBNjyxMAhXC0GUIZVJhdHDzGdpCJXyNMlgwqJXDtKmhNe3MHasDXpKQmcXV_X9bHMOlRVPS5Bg0PJvy4shGIM1NLk9nm6lz1wsh2CDe4kXpNO0cFIAXR1EJ1_DRb97HKdmYCzdKWbtLgTa16" // See readme.md file for how to obtain the msToken
	secUid := "MS4wLjABAAAADWVixuGqt-G8FDQ9yx9TLQD-4fFpwQtBhXe6EDCJ32wiprPkgzEzdGCjCR1PEwmf"                                                                              // See readme.md file for how to obtain the secUid
	_, body, res, err := api.UserList(ctx, endpoint.Following, secUid, "0", "0", msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

// Search video item . example: https://www.tiktok.com/@gemdzq
func TestSearch(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "" // See readme.md file for how to obtain the msToken
	//msToken := "input your msToken for Browser Cookie"                                                                                                                        // See readme.md file for how to obtain the msToken
	ttwid := "1%7CLX32v3zr6FNzupFObOTVmXKBOWtCKeK81YZJHtL9EuI%7C1744637575%7Ce786ea2e5b4169cfe47f1a318485b5856220c27a9469d910660949ee3b22bb23" // See readme.md file for how to obtain the ttwid
	_, body, res, err := api.Search(ctx, endpoint.LIVE, "Trump", "", ttwid, msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}
