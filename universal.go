package urlsubmitter

import (
	"errors"
	"fmt"
)

type UniversalOptions struct {
	BaiduAPI string // Baidu搜索资源平台-资源提交-普通收录-API提交-推送接口

	BingKey         string // Key for the Bing IndexNow API.
	BingKeyLocation string // Location of the Bing key file.
	BingHost        string // Host name of the site.

	GoogleCredentialsFile string // 谷歌indexing-api服务账号密钥文件路径
}

// UniversalSubmitter submits URLs to all platforms.
type UniversalSubmitter struct {
	BaiduSubmitter  *BaiduSubmitter
	BingSubmitter   *BingSubmitter
	GoogleSubmitter *GoogleSubmitter
}

// NewUniversalSubmitter creates a new UniversalSubmitter.
func NewUniversalSubmitter(options *UniversalOptions) *UniversalSubmitter {
	var baiduSubmitter *BaiduSubmitter
	var bingSubmitter *BingSubmitter
	var googleSubmitter *GoogleSubmitter

	if options.BaiduAPI != "" {
		baiduSubmitter = NewBaiduSubmitter(options.BaiduAPI)
	}
	if options.BingKey != "" && options.BingKeyLocation != "" && options.BingHost != "" {
		bingSubmitter = NewBingSubmitter(options.BingKey, options.BingKeyLocation, options.BingHost)
	}
	if options.GoogleCredentialsFile != "" {
		googleSubmitter = NewGoogleSubmitter(options.GoogleCredentialsFile)
	}

	if baiduSubmitter == nil && bingSubmitter == nil && googleSubmitter == nil {
		return nil
	}

	return &UniversalSubmitter{
		BaiduSubmitter:  baiduSubmitter,
		BingSubmitter:   bingSubmitter,
		GoogleSubmitter: googleSubmitter,
	}
}

// SubmitURLs submits URLs to all platforms.
func (s *UniversalSubmitter) SubmitURLs(urls []string) (map[string]string, error) {
	results := make(map[string]string)
	var errs error

	baiduResult, err := s.BaiduSubmitter.SubmitURLs(urls)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("BaiduError: %w", err))
	} else {
		results["baidu"] = baiduResult
	}

	bingResult, err := s.BingSubmitter.SubmitURLs(urls)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("BingError: %w", err))
	} else {
		results["bing"] = bingResult
	}

	googleResult, err := s.GoogleSubmitter.SubmitURLs(urls)
	if err != nil {
		errs = errors.Join(errs, fmt.Errorf("GoogleError: %w", err))
	} else {
		results["google"] = googleResult
	}

	return results, errs
}
