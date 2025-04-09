package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-tiktok/consts"
	"go-tiktok/script"
	"net/http"
	"net/url"
)

type TikTokOption struct {
	Method  consts.HTTPMethodType
	URL     string
	Params  map[string]interface{}
	Cookies map[string]interface{}
	Body    interface{}
	Model   interface{}
	MsToken string
	Proxy   string
}

type ITikTokCli interface {
	Do(ctx context.Context) (*resty.Response, error)
	BuildURL(updateParams map[string]interface{}) error
	Unmarshal(body []byte, m interface{}) error
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
	return &STikTokCli{
		client:  cli,
		req:     cli.R(),
		options: options,
	}
}

func (s *STikTokCli) Do(ctx context.Context) (*resty.Response, error) {
	var resp *resty.Response
	var err error
	s.client.SetHeaders(consts.GetBaseHeaders())
	queryParams := consts.GetBaseParams()
	for k, v := range queryParams {
		s.options.Params[k] = v
	}
	ua := s.client.Header.Get(consts.USER_AGENT_KEY)
	for k, v := range s.options.Cookies {
		s.client.SetCookie(&http.Cookie{
			Name:  k,
			Value: fmt.Sprintf("%v", v),
		})
	}
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
	s.options.Params[consts.MSTOKEN] = s.options.MsToken
	err = s.BuildURL(nil)
	xbValue, err := script.NewSNodeScriptUtil().GetXbValueByNodeCmd(ctx, s.options.URL, ua)
	if err != nil {
		return resp, err
	}
	s.options.URL += "&" + consts.XBOGUS + "=" + xbValue
	switch s.options.Method {
	case consts.GET:
		resp, err = s.client.R().Get(s.options.URL)
	case consts.POST:
		resp, err = s.req.Post(s.options.URL)
	default:
		return resp, fmt.Errorf("invalid HTTP method: %s", s.options.Method)
	}
	return resp, err
}

func (s *STikTokCli) Unmarshal(body []byte, m interface{}) error {
	err := json.Unmarshal(body, m)
	return err
}

func (s *STikTokCli) BuildURL(updateParams map[string]interface{}) error {
	baseURL := s.options.URL

	// 解析基础URL
	u, err := url.Parse(baseURL)
	if err != nil {
		return fmt.Errorf("解析URL失败: %v", err)
	}

	// 获取现有查询参数（如果有）
	existingParams := u.Query()

	// 如果没有初始Params，初始化
	if s.options.Params == nil {
		s.options.Params = make(map[string]interface{})
	}

	// 如果有updateParams，更新到s.options.Params中
	if updateParams != nil {
		for k, v := range updateParams {
			s.options.Params[k] = v
		}
	}

	// 将所有参数合并到查询参数中
	for k, v := range s.options.Params {
		existingParams.Set(k, fmt.Sprintf("%v", v))
	}

	// 设置编码后的查询字符串
	u.RawQuery = existingParams.Encode()

	// 更新URL
	s.options.URL = u.String()

	// 调试输出
	//fmt.Println("更新后的URL:", s.options.URL)
	return nil
}
