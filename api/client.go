package api

import "net/http"

const (
	// BaseURL is YouTube endpoint base URL.
	BaseURL = "https://www.youtube.com"

	// ClientVersion for request.
	ClientVersion = "2.20200925.01.00"
)

// Client is common HTTP client.
var Client = http.Client{}
