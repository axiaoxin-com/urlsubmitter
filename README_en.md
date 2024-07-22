# URL Submitter

`urlsubmitter` is a Go package designed for SEO purposes, enabling you to submit URLs to various search engines and services, including Baidu, Microsoft IndexNow, and Google Indexing API.

[中文版](./README.md)

## Features

- Submit URLs to Baidu
- Submit URLs to Microsoft IndexNow
- Submit URLs to Google Indexing API

## Installation

First, make sure you have Go installed. Then, you can install the `urlsubmitter` package with the following command:

```sh
go get github.com/axiaoxin-com/urlsubmitter
```

Sure, here's the translated content:

## Important Notice Before Use

Before using the Google Indexing API and Bing IndexNow, some preliminary preparation is required. You can refer to the following articles:

- [SEO | How to Quickly Notify Google to Crawl Updated Pages Using Golang](https://blog.axiaoxin.com/post/how-to-use-golang-call-google-indexing-api/)
- [SEO | Accelerate Bing's Crawling of Your Pages: Application of Golang with Bing IndexNow](https://blog.axiaoxin.com/post/how-to-use-golang-call-bing-indexnow/)

## Usage Example

Here is an example of how to use the `urlsubmitter` package:

```go
package main

import (
    "fmt"
    "log"

    "github.com/axiaoxin-com/urlsubmitter"
)

func main() {
    // Initialize Baidu submitter
    baiduSubmitter := urlsubmitter.NewBaiduSubmitter("http://data.zz.baidu.com/urls?site=https://www.example.org&token=baidutoken")

    // Initialize Bing submitter
    bingSubmitter := urlsubmitter.NewBingSubmitter(
        "https://api.indexnow.org/IndexNow",
        "www.example.org",
        "bingkey",
        "https://www.example.org/keylocation.txt",
    )

    // Initialize Google submitter
    googleSubmitter := urlsubmitter.NewGoogleSubmitter("/path/to/your-svc-account-keys.json")

    urls := []string{
        "http://www.example.org/1.html",
        "http://www.example.org/2.html",
    }

    // Submit URLs to Baidu
    baiduResult, err := baiduSubmitter.SubmitURLs(urls)
    if err != nil {
        log.Fatalf("Error submitting to Baidu: %v", err)
    }
    fmt.Println("Baidu Result:", baiduResult)

    // Submit URLs to Bing
    bingResult, err := bingSubmitter.SubmitURLs(urls)
    if err != nil {
        log.Fatalf("Error submitting to Bing: %v", err)
    }
    fmt.Println("Bing Result:", bingResult)

    // Submit URLs to Google
    googleResult, err := googleSubmitter.SubmitURLs(urls)
    if err != nil {
        log.Fatalf("Error submitting to Google: %v", err)
    }
    fmt.Println("Google Result:", googleResult)
}
```

## Running Unit Tests

You can run unit tests using the following commands, ensuring you set the necessary environment variables:

```sh
# Environment variables
export SUBMIT_URL="URL to submit"
export BAIDU_API="Baidu API"
export BING_SUBMIT_HOST="Bing host parameter"
export BING_KEY="Bing key parameter"
export BING_KEY_LOCATION="Bing keyLocation parameter"
export GOOGLE_CREDENTIALS_FILE="Google credentials file parameter"
```

Run the tests:

```sh
go test ./...
```

## Contributing

Issues and Pull Requests are welcome. Please ensure all tests pass and adhere to the project's coding standards before submitting a Pull Request.

Thank you for using `urlsubmitter`! We hope it helps streamline your URL submission process.
