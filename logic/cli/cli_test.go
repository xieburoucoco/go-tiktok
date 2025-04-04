package cli

import (
	"context"
	"testing"
)

func TestHome(t *testing.T) {
	cli := NewSTikTokCli()
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
