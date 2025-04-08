package endpoint

import "go-tiktok/consts"

type TypeSearchEndpointName string

const (
	SEARCH_ENDPOINT_NAME                        = "search"
	SEARCH_METHOD                               = consts.GET
	GENERAL              TypeSearchEndpointName = "general"
	USER                 TypeSearchEndpointName = "user"
	ITEM                 TypeSearchEndpointName = "item"
	LIVE                 TypeSearchEndpointName = "live"
)

func GetSearchRoute(endpointName TypeSearchEndpointName) string {
	return consts.API_ENDPOINT + "search/" + string(endpointName) + "/full/"
}

func GetSearchParams(keyword string) map[string]interface{} {
	//params := consts.GetBaseParams()
	params := make(map[string]interface{})
	params["keyword"] = keyword
	return params
}

func BuildSearchEndpoint(endpointName TypeSearchEndpointName, keyword string) (consts.HTTPMethodType, string, map[string]interface{}, interface{}, error) {
	res := struct {
	}{}
	return SEARCH_METHOD, GetSearchRoute(endpointName), GetSearchParams(keyword), res, nil
}
