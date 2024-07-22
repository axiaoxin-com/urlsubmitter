package urlsubmitter

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// BingSubmitter is a URL submitter for Microsoft's IndexNow API.
type BingSubmitter struct {
	API         string // API endpoint for submitting URLs to IndexNow.
	Key         string // Key for the IndexNow API.
	KeyLocation string // Location of the key file.
	Host        string // Host name of the site.
}

// NewBingSubmitter creates a new BingSubmitter with the given parameters.
func NewBingSubmitter(key, keyLocation, host string) *BingSubmitter {
	api := "https://api.indexnow.org/IndexNow"
	return &BingSubmitter{
		API:         api,
		Key:         key,
		KeyLocation: keyLocation,
		Host:        host,
	}
}

// SubmitURLs submits the given URLs to Bing's IndexNow API.
// docs: https://www.bing.com/indexnow/getstarted
func (m *BingSubmitter) SubmitURLs(urls []string) (string, error) {
	client := &http.Client{}
	data := map[string]interface{}{
		"host":        m.Host,
		"key":         m.Key,
		"keyLocation": m.KeyLocation,
		"urlList":     urls,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", m.API, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	return resp.Status, nil
}
