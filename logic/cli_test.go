package logic

import (
	"context"
	"go-tiktok/consts"
	"go-tiktok/consts/endpoint"
	"testing"
)

func TestCli(t *testing.T) {
	op := TikTokOption{
		Method: consts.GET,
		URL:    "",
		Params: map[string]interface{}{
			"1": "123",
		},
		Body:    nil,
		Model:   nil,
		MsToken: "123",
	}
	cli := NewSTikTokCli(op)
	cli.client.SetProxy("http://localhost:7890")
	response, err := cli.req.Get("https://www.tiktok.com/@gemdzq")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(response)
}

func TestCliOp(t *testing.T) {
	ctx := context.Background()
	op := TikTokOption{
		Method: consts.GET,
		URL:    "111",
		Params: map[string]interface{}{
			"1": "123",
		},
		Body:    nil,
		Model:   nil,
		MsToken: "123",
	}
	cli := NewSTikTokCli(op)
	res, err := NewITikTokCli(cli).Do(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

// Get the details of the home page user. example: https://www.tiktok.com/@gemdzq
func TestHome(t *testing.T) {
	ctx := context.Background()
	api := NewITikTokAPI(*NewParamAdapter())
	_, body, res, err := api.Home(ctx, "gemdzq", "http://localhost:7897")
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
	api := NewITikTokAPI(*NewParamAdapter())
	msToken := "input your msToken for Browser Cookie" // See readme.md file for how to obtain the msToken
	_, body, res, err := api.ItemDetail(ctx, "7273529185589562625", msToken, "http://localhost:7890")
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
	api := NewITikTokAPI(*NewParamAdapter())
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
	api := NewITikTokAPI(*NewParamAdapter())
	msToken := "input your msToken for Browser Cookie" // See readme.md file for how to obtain the msToken
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
	api := NewITikTokAPI(*NewParamAdapter())
	msToken := "input your msToken for Browser Cookie"                                       // See readme.md file for how to obtain the msToken
	secUid := "MS4wLjABAAAADWVixuGqt-G8FDQ9yx9TLQD-4fFpwQtBhXe6EDCJ32wiprPkgzEzdGCjCR1PEwmf" // See readme.md file for how to obtain the secUid
	_, body, res, err := api.UserList(ctx, endpoint.Following, secUid, "0", "1744095875", msToken, "http://localhost:7897")
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
	api := NewITikTokAPI(*NewParamAdapter())
	msToken := "input your msToken for Browser Cookie" // See readme.md file for how to obtain the msToken
	ttwid := "input your ttwid for Browser Cookie"     // See readme.md file for how to obtain the ttwid
	_, body, res, err := api.Search(ctx, endpoint.LIVE, "Trump", "", ttwid, msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}
