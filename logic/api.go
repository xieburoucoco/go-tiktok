package logic

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"go-tiktok/consts"
	"go-tiktok/consts/endpoint"
)

type ITikTokAPI interface {
	Search(ctx context.Context, typeRoute endpoint.TypeSearchEndpointName, keyword string, offset string, ttwid string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.SearchRes, err error)
	ItemDetail(ctx context.Context, itemId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error)
	UserDetail(ctx context.Context, uniqueId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserDetailRes, err error)
	MusicItemList(ctx context.Context, musicId string, cursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error)
	Home(ctx context.Context, uniqueId string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.HomeRes, err error)
	UserList(ctx context.Context, scene endpoint.SceneType, secUid, maxCursor, minCursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserListRes, err error)
}

type STikTokAPI struct {
	Adapter ParamAdapter
}

func NewITikTokAPI(adapter ParamAdapter) ITikTokAPI {
	return &STikTokAPI{
		Adapter: adapter,
	}
}

// todo need login
func (s *STikTokAPI) Search(ctx context.Context, typeRoute endpoint.TypeSearchEndpointName, keyword string, offset string, ttwid string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.SearchRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, SearchTypeName, typeRoute, keyword, offset, ttwid)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

func (s *STikTokAPI) ItemDetail(ctx context.Context, itemId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, ItemDetailTypeName, itemId)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

func (s *STikTokAPI) UserDetail(ctx context.Context, uniqueId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserDetailRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, UserDetailTypeName, uniqueId)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}
func (s *STikTokAPI) UserList(ctx context.Context, scene endpoint.SceneType, secUid, maxCursor, minCursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserListRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, UserListTypeName, scene, secUid, maxCursor, minCursor)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

func (s *STikTokAPI) MusicItemList(ctx context.Context, musicId string, cursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, MusicItemListTypeName, musicId, cursor)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

func (s *STikTokAPI) Home(ctx context.Context, uniqueId string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.HomeRes, err error) {
	_, route, res, err := endpoint.BuildHomeEndpoint(uniqueId)
	if err != nil {
		return response, responseBody, res, err
	}
	cli := resty.New()
	if len(proxyURL) != 0 {
		cli.SetProxy(proxyURL)
	}
	response, err = cli.R().SetHeaders(consts.GetBaseHeaders()).Get(route)
	if err != nil {
		return response, responseBody, res, err
	}
	responseBody = response.Body()
	res, err = endpoint.UnmarshalHomeResponse(responseBody)
	return response, responseBody, res, err
}
