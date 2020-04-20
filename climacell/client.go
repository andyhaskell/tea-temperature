package climacell

import (
	"net/http"
	"net/url"
	"time"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "api.climacell.co",
	Path:   "/v3/",
}

// Client is a client for sending requests to the ClimaCell API.
type Client struct {
	c      *http.Client
	apiKey string
}

// New creates a new Go client for the ClimaCell API.
func New(apiKey string) *Client {
	c := &http.Client{Timeout: 30 * time.Second}

	return &Client{
		c:      c,
		apiKey: apiKey,
	}
}
