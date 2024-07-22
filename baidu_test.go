package urlsubmitter

import (
	"os"
	"testing"
)

func TestBaiduSubmitter(t *testing.T) {
	api := os.Getenv("BAIDU_API")
	if api == "" {
		t.Fatal("BAIDU_API environment variable is not set")
	}
	submitter := NewBaiduSubmitter(api)

	submitURL := os.Getenv("SUBMIT_URL")
	if submitURL == "" {
		t.Fatal("SUBMIT_URL environment variable is not set")
	}
	urls := []string{
		submitURL,
	}

	result, err := submitter.SubmitURLs(urls)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	t.Logf("baidu result %v", result)
	// baidu result {"remain":9,"success":1}
}
