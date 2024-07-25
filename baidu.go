package urlsubmitter

import (
	"bytes"
	"io"
	"net/http"
	"strings"
)

// BaiduSubmitter is a URL submitter for Baidu.
type BaiduSubmitter struct {
	API string // API endpoint for submitting URLs to Baidu. Baidu搜索资源平台-资源提交-普通收录-API提交-推送接口
}

// NewBaiduSubmitter creates a new BaiduSubmitter with the given API endpoint.
func NewBaiduSubmitter(api string) *BaiduSubmitter {
	return &BaiduSubmitter{API: api}
}

// SubmitURLs submits the given URLs to Baidu's API.
// docs: https://ziyuan.baidu.com/linksubmit/index
func (b *BaiduSubmitter) SubmitURLs(urls []string) (string, error) {
	client := &http.Client{}
	data := strings.Join(urls, "\n")
	req, err := http.NewRequest("POST", b.API, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "text/plain")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
