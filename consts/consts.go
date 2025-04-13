package consts

import "net/http"

const (
	XB_SCRIPT_PATH = "js/temp_xb.js"
)

type HTTPMethodType string

const (
	GET  HTTPMethodType = http.MethodGet
	POST HTTPMethodType = http.MethodPost
)

const (
	HOME                     = "https://www.tiktok.com"
	API_ENDPOINT             = HOME + "/api/"
	BROWSER_VERSION          = "5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"
	USER_AGENT_KEY           = "user-agent"
	USER_AGENT_DEFAULT_VALUE = "Mozilla/" + BROWSER_VERSION
	XBOGUS                   = "X-Bogus"
	MSTOKEN                  = "msToken"
)

func GetBaseHeaders() map[string]string {
	return map[string]string{
		"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"accept-language":           "zh-CN,zh;q=0.9",
		"cache-control":             "no-cache",
		"pragma":                    "no-cache",
		"priority":                  "u=0, i",
		"sec-ch-ua":                 "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\"",
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "\"Windows\"",
		"sec-fetch-dest":            "document",
		"sec-fetch-mode":            "navigate",
		"sec-fetch-site":            "cross-site",
		"sec-fetch-user":            "?1",
		"upgrade-insecure-requests": "1",
		USER_AGENT_KEY:              USER_AGENT_DEFAULT_VALUE,
	}
}

func GetBaseParams() map[string]string {
	return map[string]string{
		"aid":              "1988",
		"app_name":         "tiktok_web",
		"browser_language": "zh-CN",
		"browser_name":     "Mozilla",
		"browser_online":   "true",
		"browser_platform": "Win32",
		"browser_version":  BROWSER_VERSION,
		"cookie_enabled":   "true",
		"device_id":        "7106080012156032519",
		"device_platform":  "web_pc",
		"os":               "windows",
		"region":           "US",
		"screen_height":    "1080",
		"screen_width":     "1920",
		"tz_name":          "Asia/Shanghai",
	}

}
