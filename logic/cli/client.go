package cli

import (
	"context"
	"github.com/go-resty/resty/v2"
	"go-tiktok/consts"
)

type ITikTokCli interface {
	UnmarshalHomeResponse(ctx context.Context, body []byte) (HomeJson, error)
}

type STikTokCli struct {
	client *resty.Client
	req    *resty.Request
}

func NewITikTokCli(s *STikTokCli) ITikTokCli {
	return s
}

func New2(s *STikTokCli) ITikTokCli {
	return s
}

func NewSTikTokCli() *STikTokCli {
	cli := resty.New()
	cli.SetHeaders(consts.GetBaseHeaders()).
		SetQueryParams(consts.GetBaseParams())
	return &STikTokCli{
		client: cli,
		req:    cli.R(),
	}
}
