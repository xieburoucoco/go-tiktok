package cli

import (
	"context"
	"go-tiktok/consts"
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
	do, bytes, err := NewITikTokCli(cli).Do(ctx)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(do)
	t.Log(bytes)
}
