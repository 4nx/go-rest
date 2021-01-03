package gorest

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.facest.io/v1"
)

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		apiKey:  apiKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}
