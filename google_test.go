package urlsubmitter

import (
	"os"
	"testing"
)

func TestGoogleSubmitter(t *testing.T) {
	credentialsFile := os.Getenv("GOOGLE_CREDENTIALS_FILE")
	if credentialsFile == "" {
		t.Fatal("GOOGLE_CREDENTIALS_FILE environment variable is not set")
	}
	submitURL := os.Getenv("SUBMIT_URL")
	if submitURL == "" {
		t.Fatal("SUBMIT_URL environment variable is not set")
	}

	submitter := NewGoogleSubmitter(credentialsFile)

	urls := []string{
		submitURL,
	}

	result, err := submitter.SubmitURLs(urls)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	t.Logf("google result %v", result)

	if result == "" {
		t.Errorf("Expected a non-empty result")
	}
}
