以下是根据你的项目结构和测试用例生成的 `README.md` 文件，结构清晰，内容简洁，符合你的要求。

---

# Go-TikTok SDK

一个开源的 TikTok 平台 SDK，使用 Go 语言开发，支持获取用户详情、视频元数据、音乐信息、搜索等功能。

---

## 使用方法

### 前置条件
在使用本 SDK 之前，请确保已安装以下环境：
1. **Node.js**：本项目依赖 Node.js 运行 JavaScript 脚本（如 `script.js`），请先安装 Node.js（建议版本 v20.x 或更高）。
    - 下载地址：https://nodejs.org/
    - 安装后运行 `node -v` 验证版本。

2. **Go**：确保 Go 环境已配置（建议版本 1.18 或更高）。
    - 下载地址：https://golang.org/

### 安装
1. 克隆项目到本地：
   ```bash
   git clone https://github.com/xieburoucoco/go-tiktok.git
   cd go-tiktok
   ```
2. 下载 Go 依赖：
   ```bash
   go mod tidy
   ```

---

## 使用案例

你可以在 `example/api_test.go` 中找到示例代码。以下是如何运行 `TestSearch` 测试用例的步骤：

1. 打开 `example/api_test.go`，找到 `TestSearch` 函数。
2. 根据下方「从浏览器获取参数」章节，填入 `msToken` 和 `ttwid`。
3. 如果需要代理，修改 `http://localhost:7897` 为你的代理地址；否则留空。
4. 运行测试用例：
   ```bash
   go test -v ./example -run TestSearch
   ```
5. 输出示例：
   ```
   === RUN   TestSearch
   --- PASS: TestSearch (1.23s)
       api_test.go:XX: {"data": {"videos": [{"id": "12345", "title": "Trump Live", ...}]}, "status": "success"}
       api_test.go:XX: <http.Response object>
   PASS
   ok      github.com/xieburoucoco/go-tiktok/example       1.234s
   ```

---
## 功能


1. 获取主页详情：
```go
// Get the details of the home page user. example: https://www.tiktok.com/@gemdzq
func TestHome(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	proxyURL := "http://localhost:7897" // input your proxy or empty string
	uniqueId := "gemdzq"                // input your interested uniqueId
	_, body, res, err := api.Home(ctx, uniqueId, proxyURL)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}

```
![img.png](doc%2Fimg%2Fimg.png)

---
2. 获取用户主页视频详情：
```go
// Get the user's attention as well as a list of followers. example: https://www.tiktok.com/api/user/list
func TestUserList(t *testing.T) {
	ctx := context.Background()
	api := logic.NewITikTokAPI(*logic.NewParamAdapter())
	msToken := "Xae_uiErd4zeHXh0PfKhGbWV5EiFE1WqQiMXwhMOS-T0UzKVR3YcO3n9eANeWYiQkZosWXupEuTLrKIGZXexDRS3wI3dlQE_nbv4RcXTBd3xIllS127pAka__vFJD8BSPJyl8u784Zg4a7vczspUPgHcoHo=" // See readme.md file for how to obtain the msToken
	secUid := "MS4wLjABAAAADWVixuGqt-G8FDQ9yx9TLQD-4fFpwQtBhXe6EDCJ32wiprPkgzEzdGCjCR1PEwmf"                                                                                  // See readme.md file for how to obtain the secUid
	_, body, res, err := api.UserList(ctx, endpoint.Following, secUid, "0", "1744095875", msToken, "http://localhost:7897")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(body))
	t.Log(res)
}
```
![img.png](doc%2Fimg%2Fimg.png)

---
## 从浏览器获取参数

以下参数需要从 TikTok 网站通过浏览器获取：

### 1. `msToken`
- **步骤**：
    1. 打开浏览器，访问 `https://www.tiktok.com`。
    2. 按 F12 打开开发者工具，转到「Network」标签。
    3. 刷新页面，筛选请求，找到任意 TikTok API 请求（如 `/api/search`）。
    4. 在请求的 Cookie 中找到 `msToken=xxx`，复制 `xxx` 部分。

### 2. `ttwid`
- **步骤**：
    1. 同上，打开开发者工具，转到「Network」标签。
    2. 找到 TikTok 的请求，在 Cookie 中查找 `ttwid=xxx`，复制 `xxx` 部分。

### 3. `secUid`
- **步骤**：
    1. 访问目标用户主页（如 `https://www.tiktok.com/@gemdzq`）。
    2. 查看页面源代码（Ctrl+U），搜索 `secUid`。
    3. 找到类似 `"secUid": "MS4wLjABAAAAxxx"` 的字段，复制引号中的值。

### 4. `uniqueId`
- **步骤**：
    1. 在用户主页 URL 中，`@` 后的部分即为 `uniqueId`。
    2. 例如：`https://www.tiktok.com/@gemdzq`，`uniqueId` 为 `gemdzq`。

---

## 联系方式

- **邮箱**：xieburoucoco@example.com（请替换为你的真实邮箱）
- **寄语**：欢迎开发者留下 Issue 或提交 PR 请求！任何意见我都会认真听取，期待与大家一起完善这个项目。

---

### 项目结构（参考）
```
go-tiktok/
├── consts/             # 常量定义
├── example/            # 示例代码
├── logic/              # 核心逻辑
├── script/             # 脚本文件（含 JS）
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

---

这个 `README.md` 文件简洁明了，包含了项目名称、使用方法、使用案例、参数获取步骤以及联系方式，完全基于你的项目结构和测试用例生成。你可以直接将其放入项目根目录，并在邮箱部分替换为你的真实邮箱地址。