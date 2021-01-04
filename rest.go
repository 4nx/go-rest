package rest

import (
	"bytes"
	"net/http"
	"net/url"
)

// Version represents the current version of this library
const Version = "0.0.1"

// Method holds the support HTTP methods
type Method string

// Supported HTTP methods
const (
	// Get GET HTTP method
	Get Method = "GET"

	// Post POST HTTP method
	Post Method = "POST"
)

// Request holds the request which should be made
type Request struct {
	Method      Method
	BaseURL     string
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
}

// Response holds the response from the API call
type Response struct {
	StatusCode int
	Body       string
	Headers    map[string][]string
}

// Error is a struct for the error handling
type Error struct {
	Response *Response
}

// DefaultHTTPClient will be used if no custom HTTP client is defined
var DefaultHTTPClient = &Client{HTTPClient: &http.Client{}}

// Client allows modification of client headers, redirect policy
// and other settings
// See https://golang.org/pkg/net/http
type Client struct {
	HTTPClient *http.Client
}

// AddContentType adds the json content type to the headers
func AddContentType(request *http.Request, body []byte) *http.Request {
	_, exists := request.Header["Content-Type"]
	if len(body) > 0 && !exists {
		request.Header.Set("Content-Type", "application/json")
	}
	return request
}

// AddQueryParams adds the query parameters to the provided base URL
func AddQueryParams(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

// CreateRequest creates the API request object
func CreateRequest(request Request) (*http.Request, error) {
	if len(request.QueryParams) != 0 {
		request.BaseURL = AddQueryParams(request.BaseURL, request.QueryParams)
	}
	req, err := http.NewRequest(string(request.Method), request.BaseURL, bytes.NewBuffer(request.Body))
	if err != nil {
		return nil, err
	}
	for key, value := range request.Headers {
		req.Header.Set(key, value)
	}
	req = AddContentType(req, request.Body)
	return req, err
}
