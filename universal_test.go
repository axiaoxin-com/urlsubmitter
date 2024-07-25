package urlsubmitter

import (
	"os"
	"testing"
)

func TestUniversalSubmitter(t *testing.T) {
	baiduAPI := os.Getenv("BAIDU_API")
	if baiduAPI == "" {
		t.Fatal("BAIDU_API environment variable is not set")
	}
	bingHost := os.Getenv("BING_SUBMIT_HOST")
	if bingHost == "" {
		t.Fatal("BING_SUBMIT_HOST environment variable is not set")
	}
	bingKey := os.Getenv("BING_KEY")
	if bingKey == "" {
		t.Fatal("BING_KEY environment variable is not set")
	}
	bingKeyLocation := os.Getenv("BING_KEY_LOCATION")
	if bingKeyLocation == "" {
		t.Fatal("BING_KEY_LOCATION environment variable is not set")
	}
	googleCredentialsFile := os.Getenv("GOOGLE_CREDENTIALS_FILE")
	if googleCredentialsFile == "" {
		t.Fatal("GOOGLE_CREDENTIALS_FILE environment variable is not set")
	}

	submitURL := os.Getenv("SUBMIT_URL")
	if submitURL == "" {
		t.Fatal("SUBMIT_URL environment variable is not set")
	}

	universal := NewUniversalSubmitter(&UniversalOptions{
		BaiduAPI:              baiduAPI,
		BingKey:               bingKey,
		BingKeyLocation:       bingKeyLocation,
		BingHost:              bingHost,
		GoogleCredentialsFile: googleCredentialsFile,
	})
	if universal == nil {
		t.Fatal("UniversalSubmitter is nil")
	}
	urls := []string{
		submitURL,
	}
	result, err := universal.SubmitURLs(urls)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	t.Logf("universal result %v", result)
}
