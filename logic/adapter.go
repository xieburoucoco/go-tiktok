package logic

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"go-tiktok/consts"
	"go-tiktok/consts/endpoint"
)

type TypeName string

const (
	SearchTypeName        TypeName = endpoint.SEARCH_ENDPOINT_NAME
	ItemDetailTypeName    TypeName = endpoint.ITEM_DETAIL_ENDPOINT_NAME
	MusicItemListTypeName TypeName = endpoint.MUSIC_ITEM_LIST_ENDPOINT_NAME
	UserDetailTypeName    TypeName = endpoint.USER_DETAIL_ENDPOINT_NAME
	UserListTypeName      TypeName = endpoint.USER_LIST_ENDPOINT_NAME
)

type ParamAdapter struct {
}

func NewParamAdapter() *ParamAdapter {
	return &ParamAdapter{}
}

func (s *ParamAdapter) GetEndpointRequestData(ctx context.Context, typeName TypeName, args ...interface{}) (httpMethod consts.HTTPMethodType, route string, params map[string]interface{}, cookies map[string]interface{}, model interface{}, err error) {
	switch typeName {
	case SearchTypeName:
		httpMethod, route, params, cookies, model, err = endpoint.BuildSearchEndpoint(args[0].(endpoint.TypeSearchEndpointName), args[1].(string), args[2].(string), args[3].(string))
	case ItemDetailTypeName:
		httpMethod, route, params, cookies, model, err = endpoint.BuildItemDetailEndpoint(args[0].(string))
	case UserDetailTypeName:
		httpMethod, route, params, cookies, model, err = endpoint.BuildUserDetailEndpoint(args[0].(string))
	case UserListTypeName:
		httpMethod, route, params, cookies, model, err = endpoint.BuildUserListEndpoint(args[0].(endpoint.SceneType), args[1].(string), args[2].(string), args[3].(string))
	case MusicItemListTypeName:
		httpMethod, route, params, cookies, model, err = endpoint.BuildMusicItemListlEndpoint(args[0].(string), args[1].(string))
	default:
		return httpMethod, route, params, cookies, model, fmt.Errorf("unsupported type name: %s", typeName)
	}
	return httpMethod, route, params, cookies, model, nil
}

func (s *ParamAdapter) FetchEndpoint(ctx context.Context, proxyURL string, msToken string, typeName TypeName, args ...interface{}) (response *resty.Response, responseBody []byte, err error) {
	httpMethod, route, params, cookies, model, err := NewParamAdapter().GetEndpointRequestData(ctx, typeName, args...)
	if err != nil {
		return response, responseBody, err
	}
	cli := NewSTikTokCli(TikTokOption{
		Method:  httpMethod,
		URL:     route,
		Params:  params,
		Cookies: cookies,
		Model:   model,
		MsToken: msToken,
		Proxy:   proxyURL,
	})
	response, err = cli.Do(ctx)
	if err != nil {
		return response, responseBody, err
	}
	responseBody = response.Body()
	return response, responseBody, err
}
