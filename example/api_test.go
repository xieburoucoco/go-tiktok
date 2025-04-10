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
	msToken := "rwyKHF2ULPSVb4vYYJK_ilJ3LZWbYXNJ2M7y57-hMrHm0NzB_gTwVWX74Pd8PpXasOutNgaq-W6KXK9ckibc8Orv-l-W1_i6rfh5Pw7vnf-9Dd6NwENdN8vJ-IImlkVOkAlZEvUSRtwAMT7d88ma4PmKx5k=" // See readme.md file for how to obtain the msToken
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
	msToken := "input your msToken for Browser Cookie" // See readme.md file for how to obtain the msToken
	_, body, res, err := api.MusicItemList(ctx, "7389877976516627216", "0", msToken, "http://localhost:7890")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

// Get the user's details . example:  https://www.tiktok.com/api/user/detail/
func TestUserDetail(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "MUyWP1yEeMKXbZFS9aaw8FXb2VQ60BZOnqQKu-Snmzfeh3DrXNIWcotjHNlNw4dT7hlwhx_XR5aHbsrXGWFoheYc6_HliiD4eNM-sVhe4-kV1lRadi4dFFVdh7HJN54inq-d1LbJWUoe2-0Bu3vBiaXnHnw=" // See readme.md file for how to obtain the msToken
	_, body, res, err := api.UserDetail(ctx, "gemdzq", msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))

	t.Log(res)
}

// Get the user's attention as well as a list of followers. example: https://www.tiktok.com/api/user/list
func TestUserList(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "MUyWP1yEeMKXbZFS9aaw8FXb2VQ60BZOnqQKu-Snmzfeh3DrXNIWcotjHNlNw4dT7hlwhx_XR5aHbsrXGWFoheYc6_HliiD4eNM-sVhe4-kV1lRadi4dFFVdh7HJN54inq-d1LbJWUoe2-0Bu3vBiaXnHnw=" // See readme.md file for how to obtain the msToken
	secUid := ""                                                                                                                                                              // See readme.md file for how to obtain the secUid
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
	ttwid := "1%7CLX32v3zr6FNzupFObOTVmXKBOWtCKeK81YZJHtL9EuI%7C1744293769%7C3835791b940f7b3e5eedb778c92d72290b2b1a5acbf7695c6572ec508f0acb7e" // See readme.md file for how to obtain the ttwid
	_, body, res, err := api.Search(ctx, endpoint.LIVE, "Trump", "", ttwid, msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}
