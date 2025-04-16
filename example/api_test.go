package example

import (
	"context"
	"github.com/xieburoucoco/go-tiktok/consts/endpoint"
	"github.com/xieburoucoco/go-tiktok/logic"
	"testing"
)

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
	msToken := "Enter the cookie key mstoken you got from your browser" // See readme.md file for how to obtain the msToken
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
	msToken := "Enter the cookie key mstoken you got from your browser" // See readme.md file for how to obtain the msToken
	secUid := "input your secUid"                                       // See readme.md file for how to obtain the secUid
	_, body, res, err := api.ItemList(ctx, secUid, "0", msToken, "http://localhost:7897")
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
	msToken := "Enter the cookie key mstoken you got from your browser" // See readme.md file for how to obtain the msToken // See readme.md file for how to obtain the msToken
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
	msToken := "Enter the cookie key mstoken you got from your browser" // See readme.md file for how to obtain the msToken // See readme.md file for how to obtain the msToken
	_, body, res, err := api.CommentList(ctx, "7339184461612993794", "0", msToken, "http://localhost:7897")
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
	msToken := "Enter the cookie key mstoken you got from your browser" // See readme.md file for how to obtain the msToken // See readme.md file for how to obtain the msToken
	secUid := "input your secUid"                                       // See readme.md file for how to obtain the secUid
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
	msToken := "input your msToken for Browser Cookie" // See readme.md file for how to obtain the msToken
	ttwid := "enter the ttwid you got from the cookie" // See readme.md file for how to obtain the ttwid
	_, body, res, err := api.Search(ctx, endpoint.LIVE, "Trump", "", ttwid, msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}
