# URL Submitter

`urlsubmitter` 是一个用 Go 语言编写的 SEO 工具包，用于将 URL 提交到不同的搜索引擎和服务，包括百度、微软 IndexNow 和 Google Indexing API。

[English Version](./README_en.md)

## 功能

- 将 URL 提交到百度
- 将 URL 提交到微软 IndexNow
- 将 URL 提交到 Google Indexing API

## 安装

首先，确保你已经安装了 Go 语言环境。然后，你可以通过以下命令安装 `urlsubmitter` 包：

```sh
go get github.com/axiaoxin-com/urlsubmitter
```

## 使用前必读

调用 Google Indexing API 和 Bing IndexNow 前期需要有一点准备工作，可以参考以下文章：

- [SEO ｜如何通过 Golang 快速通知 Google 抓取更新网页](https://blog.axiaoxin.com/post/how-to-use-golang-call-google-indexing-api/)
- [SEO ｜让 Bing 更快抓取你的网页：Golang 与 Bing IndexNow 的应用](https://blog.axiaoxin.com/post/how-to-use-golang-call-bing-indexnow/)

## 使用示例

以下是如何使用 `urlsubmitter` 包的示例代码：

```go
package main

import (
    "fmt"
    "log"

    "github.com/axiaoxin-com/urlsubmitter"
)

func main() {
    // 初始化 Baidu 提交器
    baiduSubmitter := urlsubmitter.NewBaiduSubmitter("http://data.zz.baidu.com/urls?site=https://www.example.org&token=baidutoken")

    // 初始化 Bing 提交器
    bingSubmitter := urlsubmitter.NewBingSubmitter(
        "www.example.org",
        "bingkey",
        "https://www.example.org/keylocation.txt",
    )

    // 初始化 Google 提交器
    googleSubmitter := urlsubmitter.NewGoogleSubmitter("/path/to/your-svc-account-keys.json")

    urls := []string{
        "http://www.example.org/1.html",
        "http://www.example.org/2.html",
    }

    // 提交 URL 到 Baidu
    baiduResult, err := baiduSubmitter.SubmitURLs(urls)
    if err != nil {
        log.Fatalf("Error submitting to Baidu: %v", err)
    }
    fmt.Println("Baidu Result:", baiduResult)

    // 提交 URL 到 Bing
    bingResult, err := bingSubmitter.SubmitURLs(urls)
    if err != nil {
        log.Fatalf("Error submitting to Bing: %v", err)
    }
    fmt.Println("Bing Result:", bingResult)

    // 提交 URL 到 Google
    googleResult, err := googleSubmitter.SubmitURLs(urls)
    if err != nil {
        log.Fatalf("Error submitting to Google: %v", err)
    }
    fmt.Println("Google Result:", googleResult)
}
```

`UniversalSubmitter` 用法示例:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/axiaoxin-com/urlsubmitter"
)

func main() {
	options := &urlsubmitter.UniversalOptions{
		BaiduAPI:            os.Getenv("BAIDU_API"),
		BingKey:             os.Getenv("BING_KEY"),
		BingKeyLocation:     os.Getenv("BING_KEY_LOCATION"),
		BingHost:            os.Getenv("BING_SUBMIT_HOST"),
		GoogleCredentialsFile: os.Getenv("GOOGLE_CREDENTIALS_FILE"),
	}

	universal := urlsubmitter.NewUniversalSubmitter(options)
	if universal == nil {
		log.Fatal("No valid submitter configurations provided.")
	}

	urls := []string{
		"http://www.example.com/1.html",
		"http://www.example.com/2.html",
		"https://www.example.org/url1",
		"https://www.example.org/folder/url2",
		"https://example.com/new-url/",
	}

	results, err := universal.SubmitURLs(urls)
	if err != nil {
		log.Fatalf("Failed to submit URLs: %v", err)
	}

	for platform, result := range results {
		fmt.Printf("%s: %s\n", platform, result)
	}
}
```

## 运行单元测试

你可以使用以下命令运行单元测试，通过从环境变量读取相关参数：

```sh
# 环境变量
export SUBMIT_URL="要提交的url"
export BAIDU_API="百度api"
export BING_SUBMIT_HOST="bing host参数"
export BING_KEY="bing key参数"
export BING_KEY_LOCATION="bing keyLocation参数"
export GOOGLE_CREDENTIALS_FILE="google密钥文件参数"
```

运行测试：

```sh
go test ./...
```

## 贡献

欢迎提交问题（Issues）和贡献代码（Pull Requests）。在提交 Pull Request 之前，请确保所有测试通过，并遵循项目的代码规范。

感谢你使用 `urlsubmitter`！希望它能帮助你简化 URL 提交的工作流程。
