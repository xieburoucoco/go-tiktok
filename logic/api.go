package logic

import (
	"context"
	"encoding/json"
	"go-tiktok/consts/endpoint"
)

type ITikTokAPI interface {
	Search(ctx context.Context, typeRoute endpoint.TypeSearchEndpointName, keyword string, msToken string, proxyURL string) ([]byte, interface{}, error)
	ItemDetail(ctx context.Context, itemId string, msToken string, proxyURL string) (responseBody []byte, res endpoint.ItemDetailRes, err error)
	MusicItemList(ctx context.Context, musicId string, cursor string, msToken string, proxyURL string) (responseBody []byte, res endpoint.ItemDetailRes, err error)
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
func (s *STikTokAPI) Search(ctx context.Context, typeRoute endpoint.TypeSearchEndpointName, keyword string, msToken string, proxyURL string) ([]byte, interface{}, error) {
	var temp interface{}
	_, responseBody, err := s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, SearchTypeName, typeRoute, keyword)
	return responseBody, temp, err
}

func (s *STikTokAPI) ItemDetail(ctx context.Context, itemId string, msToken string, proxyURL string) (responseBody []byte, res endpoint.ItemDetailRes, err error) {
	_, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, ItemDetailTypeName, itemId)
	if err != nil {
		return responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return responseBody, res, err
}

func (s *STikTokAPI) MusicItemList(ctx context.Context, musicId string, cursor string, msToken string, proxyURL string) (responseBody []byte, res endpoint.ItemDetailRes, err error) {
	_, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, MusicItemListTypeName, musicId, cursor)
	if err != nil {
		return responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return responseBody, res, err
}
