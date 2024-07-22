package urlsubmitter

import (
	"os"
	"testing"
)

func TestBingSubmitter(t *testing.T) {
	host := os.Getenv("BING_SUBMIT_HOST")
	if host == "" {
		t.Fatal("BING_SUBMIT_HOST environment variable is not set")
	}
	key := os.Getenv("BING_KEY")
	if key == "" {
		t.Fatal("BING_KEY environment variable is not set")
	}
	keyLocation := os.Getenv("BING_KEY_LOCATION")
	if keyLocation == "" {
		t.Fatal("BING_KEY_LOCATION environment variable is not set")
	}
	submitURL := os.Getenv("SUBMIT_URL")
	if submitURL == "" {
		t.Fatal("SUBMIT_URL environment variable is not set")
	}

	submitter := NewBingSubmitter(key, keyLocation, host)

	urls := []string{
		submitURL,
	}

	result, err := submitter.SubmitURLs(urls)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	t.Logf("bing result %v", result)

	expected := "200 OK"
	if result != expected {
		t.Errorf("Expected result %v, got %v", expected, result)
	}
}
