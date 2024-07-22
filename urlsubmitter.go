package urlsubmitter

// URLSubmitter defines an interface for submitting URLs to various services.
type URLSubmitter interface {
	SubmitURLs(urls []string) (string, error)
}
