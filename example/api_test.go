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
	ttwid := "t6NL9TUgftL2TJwac-tvewiC9UN5FZamAl6_189t6C-9lSC7sXvc7vMSgdF9KDBsHvczb0YmNarPzODKES89ofY7KYzb4S0LBj0a2DN506OpkQuj1JQ5gwTc5o3brI1D6v-vK6nwzEXahzHiRqwviMDYMw==" // See readme.md file for how to obtain the ttwid
	_, body, res, err := api.Search(ctx, endpoint.LIVE, "Trump", "", ttwid, msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}
