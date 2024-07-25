package urlsubmitter

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"google.golang.org/api/indexing/v3"
	"google.golang.org/api/option"
)

// GoogleSubmitter is a URL submitter for Google's Indexing API.
type GoogleSubmitter struct {
	CredentialsFile string // 谷歌 indexing-api 服务账号密钥文件路径
}

// NewGoogleSubmitter creates a new GoogleSubmitter with the given credentials file.
func NewGoogleSubmitter(credentialsFile string) *GoogleSubmitter {
	return &GoogleSubmitter{CredentialsFile: credentialsFile}
}

// SubmitURLs submits the given URLs to Google's Indexing API.
// docs: https://developers.google.com/search/apis/indexing-api/v3/quickstart
func (g *GoogleSubmitter) SubmitURLs(urls []string) (string, error) {
	ctx := context.Background()
	client, err := indexing.NewService(ctx, option.WithCredentialsFile(g.CredentialsFile))
	if err != nil {
		return "", err
	}

	results := make([]string, 0, len(urls))
	for _, url := range urls {
		data := &indexing.UrlNotification{
			Type: "URL_UPDATED",
			Url:  url,
		}
		rsp, err := client.UrlNotifications.Publish(data).Do()
		if err != nil {
			results = append(results, fmt.Sprintf("Error: %v", err))
		} else {
			rspjson, _ := json.MarshalIndent(rsp, "", " ")
			results = append(results, string(rspjson))
		}
	}

	return strings.Join(results, "\n"), nil
}
