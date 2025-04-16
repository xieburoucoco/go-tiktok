// Package logic provides a TikTok API client for interacting with various TikTok endpoints,
// such as searching, fetching user details, video metadata, comments, and more.
// It uses the resty HTTP client and supports proxy and token-based authentication.
package logic

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/xieburoucoco/go-tiktok/consts"
	"github.com/xieburoucoco/go-tiktok/consts/endpoint"
)

// ITikTokAPI defines methods for interacting with TikTok API endpoints.
// It supports operations like searching, fetching item details, user details, comments, and more.
// All methods return the raw HTTP response, response body, a typed result, and an error.
type ITikTokAPI interface {
	// Search performs a search on TikTok for the given keyword and type (e.g., LIVE, VIDEO).
	// It requires a ttwid cookie and optionally an msToken for authentication.
	// The offset parameter controls pagination.
	Search(ctx context.Context, typeRoute endpoint.TypeSearchEndpointName, keyword string, offset string, ttwid string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.SearchRes, err error)

	// ItemDetail fetches metadata for a specific TikTok item (e.g., video) by its itemId.
	// It requires an msToken for authentication.
	ItemDetail(ctx context.Context, itemId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error)

	// ItemList retrieves a list of items (e.g., videos) for a user identified by secUid.
	// The cursor parameter controls pagination, and msToken is required for authentication.
	ItemList(ctx context.Context, secUid string, cursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemListRes, err error)

	// UserDetail fetches profile details for a TikTok user identified by uniqueId.
	// It requires an msToken for authentication.
	UserDetail(ctx context.Context, uniqueId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserDetailRes, err error)

	// MusicItemList retrieves a list of items (e.g., videos) associated with a musicId.
	// The cursor parameter controls pagination, and msToken is required for authentication.
	MusicItemList(ctx context.Context, musicId string, cursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error)

	// CommentList fetches comments for a TikTok item identified by awemeId.
	// The cursor parameter controls pagination, and msToken is required for authentication.
	CommentList(ctx context.Context, awemeId, cursor, msToken, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.CommentRes, err error)

	// Home fetches the TikTok home feed for a user identified by uniqueId.
	// It does not require an msToken but supports proxy usage.
	Home(ctx context.Context, uniqueId string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.HomeRes, err error)

	// UserList retrieves a list of users (e.g., followers or following) for a user identified by secUid.
	// The scene parameter specifies the type of list (e.g., followers, following).
	// maxCursor and minCursor control pagination, and msToken is required for authentication.
	UserList(ctx context.Context, scene endpoint.SceneType, secUid, maxCursor, minCursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserListRes, err error)
}

// STikTokAPI is an implementation of ITikTokAPI that uses a ParamAdapter
// to handle endpoint requests and response parsing.
type STikTokAPI struct {
	Adapter ParamAdapter
}

// NewITikTokAPI creates a new ITikTokAPI instance with the provided ParamAdapter.
func NewITikTokAPI(adapter ParamAdapter) ITikTokAPI {
	return &STikTokAPI{
		Adapter: adapter,
	}
}

// Search implements the Search method of ITikTokAPI.
// It fetches search results for the given keyword and type, using the provided tokens and proxy.
// The response is unmarshaled into an endpoint.SearchRes struct.
func (s *STikTokAPI) Search(ctx context.Context, typeRoute endpoint.TypeSearchEndpointName, keyword string, offset string, ttwid string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.SearchRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, SearchTypeName, typeRoute, keyword, offset, ttwid)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

// ItemDetail implements the ItemDetail method of ITikTokAPI.
// It fetches metadata for a specific item by itemId, using the provided msToken and proxy.
// The response is unmarshaled into an endpoint.ItemDetailRes struct.
func (s *STikTokAPI) ItemDetail(ctx context.Context, itemId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, ItemDetailTypeName, itemId)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

// ItemList implements the ItemList method of ITikTokAPI.
// It retrieves a list of items for a user by secUid, using the provided cursor, msToken, and proxy.
// The response is unmarshaled into an endpoint.ItemListRes struct.
func (s *STikTokAPI) ItemList(ctx context.Context, secUid string, cursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemListRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, ItemListTypeName, secUid, cursor)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

// UserDetail implements the UserDetail method of ITikTokAPI.
// It fetches user profile details by uniqueId, using the provided msToken and proxy.
// The response is unmarshaled into an endpoint.UserDetailRes struct.
func (s *STikTokAPI) UserDetail(ctx context.Context, uniqueId string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserDetailRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, UserDetailTypeName, uniqueId)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

// UserList implements the UserList method of ITikTokAPI.
// It retrieves a list of users (e.g., followers or following) by secUid and scene, using the provided cursors, msToken, and proxy.
// The response is unmarshaled into an endpoint.UserListRes struct.
func (s *STikTokAPI) UserList(ctx context.Context, scene endpoint.SceneType, secUid, maxCursor, minCursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.UserListRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, UserListTypeName, scene, secUid, maxCursor, minCursor)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

// MusicItemList implements the MusicItemList method of ITikTokAPI.
// It retrieves a list of items associated with a musicId, using the provided cursor, msToken, and proxy.
// The response is unmarshaled into an endpoint.ItemDetailRes struct.
func (s *STikTokAPI) MusicItemList(ctx context.Context, musicId string, cursor string, msToken string, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.ItemDetailRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, MusicItemListTypeName, musicId, cursor)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

// CommentList implements the CommentList method of ITikTokAPI.
// It fetches comments for an item by awemeId, using the provided cursor, msToken, and proxy.
// The response is unmarshaled into an endpoint.CommentRes struct.
func (s *STikTokAPI) CommentList(ctx context.Context, awemeId, cursor, msToken, proxyURL string) (response *resty.Response, responseBody []byte, res endpoint.CommentRes, err error) {
	response, responseBody, err = s.Adapter.FetchEndpoint(ctx, proxyURL, msToken, CommentTypeName, awemeId, cursor)
	if err != nil {
		return response, responseBody, res, err
	}
	err = json.Unmarshal(responseBody, &res)
	return response, responseBody, res, err
}

// Home implements the Home method of ITikTokAPI.
// It fetches the home feed for a user by uniqueId, using the provided proxy.
// The response is unmarshaled into an endpoint.HomeRes struct.
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
	body, res, err := endpoint.UnmarshalHomeResponse(responseBody)
	return response, body, res, err
}
