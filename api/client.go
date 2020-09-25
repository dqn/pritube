package api

import "net/http"

// BaseURL is YouTube endpoint base URL.
const BaseURL = "https://www.youtube.com"

// Client is common HTTP client.
var Client = http.Client{}
