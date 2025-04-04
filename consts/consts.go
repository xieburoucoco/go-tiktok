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
	HOME                        = "https://www.tiktok.com"
	BROWSER_VERSION             = "5.0 (Windows NT 10.0; Win64; x64)"
	USER_AGENT_KEY              = "user-agent"
	USER_AGENT_DEFAULT_VALUE    = "Mozilla/" + BROWSER_VERSION + " AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36"
	TIKTOK_HOME_EXTRACT_PATTERN = `<script id="__UNIVERSAL_DATA_FOR_REHYDRATION__" type="application/json">(.*?)</script>`
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
		"browser_language": "en-US",
		"browser_name":     "Mozilla",
		"browser_online":   "true",
		"browser_platform": "Win32",
		"browser_version":  BROWSER_VERSION,
		"channel":          "tiktok_web",
		"cookie_enabled":   "true",
		"device_id":        "7106080012156032519",
		"device_platform":  "web_pc",
		"focus_state":      "true",
		"from_page":        "user",
		"history_len":      "1",
		"is_fullscreen":    "false",
		"is_page_visible":  "true",
		"os":               "windows",
		"priority_region":  "",
		"referer":          HOME,
		"region":           "US",
		"screen_height":    "1080",
		"screen_width":     "1920",
		"tz_name":          "Asia/Shanghai",
		"webcast_language": "en",
	}

}
