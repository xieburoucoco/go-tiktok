package logic

import (
	"context"
	"go-tiktok/consts"
	"go-tiktok/consts/endpoint"
	"testing"
)

func TestHome(t *testing.T) {
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
	h, err := NewITikTokCli(cli).UnmarshalHomeResponse(context.Background(), response.Body())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(h)
}

func TestCli(t *testing.T) {
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

func TestApi(t *testing.T) {
	ctx := context.Background()
	api := NewITikTokAPI(*NewParamAdapter())
	msToken := "Dux1qw1mvMzqFddt152mCDIvRwQdSq0isVyS-cVLZP7b4vTe0wQCUqyInOOjVTHKCmXIY0MTTioVelBoWlFX0Yf-hQIN49WUIne1Sv9zAllBIRPX171WrhRwqqYbhPOODhCjIXdZuQDzrZG_d2pv6D_VXV0="
	body, res, err := api.Search(ctx, endpoint.GENERAL, "123", msToken, "http://localhost:7890")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

func TestItemDetail(t *testing.T) {
	ctx := context.Background()
	api := NewITikTokAPI(*NewParamAdapter())
	msToken := ""
	body, res, err := api.ItemDetail(ctx, "7273529185589562625", msToken, "http://localhost:7890")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

func TestMusicItemList(t *testing.T) {
	ctx := context.Background()
	api := NewITikTokAPI(*NewParamAdapter())
	msToken := ""
	body, res, err := api.MusicItemList(ctx, "7389877976516627216", "0", msToken, "http://localhost:7890")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}
