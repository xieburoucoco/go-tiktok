package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-tiktok/consts"
	"go-tiktok/script"
	"strings"
)

type TikTokOption struct {
	Method  consts.HTTPMethodType
	URL     string
	Params  map[string]interface{}
	Body    interface{}
	Model   *interface{}
	MsToken string
	Proxy   string
}

type ITikTokCli interface {
	Do(ctx context.Context) (*resty.Response, []byte, error)
	BuildURL(updateParams map[string]interface{}) error
	Unmarshal(body []byte, m interface{}) error
	UnmarshalHomeResponse(ctx context.Context, body []byte) (HomeJson, error)
}

type STikTokCli struct {
	client  *resty.Client
	req     *resty.Request
	options TikTokOption
}

func NewITikTokCli(s *STikTokCli) ITikTokCli {
	return s
}

func NewSTikTokCli(options TikTokOption) *STikTokCli {
	cli := resty.New()
	cli.SetHeaders(consts.GetBaseHeaders()).
		SetQueryParams(consts.GetBaseParams())
	return &STikTokCli{
		client:  cli,
		req:     cli.R(),
		options: options,
	}
}

func (s *STikTokCli) Do(ctx context.Context) (*resty.Response, []byte, error) {
	var resp *resty.Response
	var err error
	ua := s.client.Header.Get(consts.USER_AGENT_KEY)
	if len(ua) == 0 {
		panic("User-Agent is empty")
	}
	if len(s.options.Method) == 0 {
		panic("Method is empty")
	}
	if len(s.options.URL) == 0 {
		panic("URL is empty")
	}
	if len(s.options.Proxy) != 0 {
		s.client.SetProxy(s.options.Proxy)
	}
	if len(s.options.MsToken) == 0 {
		panic("MsToken is empty")
	}
	s.options.Params["msToken"] = s.options.MsToken
	if err = s.BuildURL(nil); err != nil {
		return resp, nil, err
	}
	xbValue, err := script.NewSNodeScriptUtil().GetXbValueByNodeCmd(ctx, s.options.URL, ua)
	if err != nil {
		return resp, nil, err
	}
	err = s.BuildURL(map[string]interface{}{
		"X-Bogus": xbValue,
	})
	switch s.options.Method {
	case consts.GET:
		resp, err = s.req.Get(s.options.URL)
	case consts.POST:
		resp, err = s.req.Post(s.options.URL)
	default:
		return resp, nil, fmt.Errorf("invalid HTTP method: %s", s.options.Method)
	}
	return resp, nil, err
}

func (s *STikTokCli) Unmarshal(body []byte, m interface{}) error {
	err := json.Unmarshal(body, m)
	return err
}

func (s *STikTokCli) BuildURL(updateParams map[string]interface{}) error {
	if !strings.HasSuffix(s.options.URL, "?") {
		s.options.URL = s.options.URL + "?"
	}
	flg := true
	current := func(p1 map[string]interface{}, p2 map[string]interface{}) map[string]interface{} {
		if len(p1) == 0 {
			return s.options.Params
		} else {
			return p2
		}
	}(updateParams, s.options.Params)
	for k, v := range current { // todo
		if flg {
			flg = false
		} else {
			s.options.URL += "&"
		}
		s.options.URL += fmt.Sprintf("%s=%s", k, v)
	}
	return nil
}
