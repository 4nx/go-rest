package rest

import (
	"fmt"
	"net/http"
	"net/url"
)

// Method holds the support HTTP methods
type Method string

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

// Supported HTTP methods
const (
	// Get GET HTTP method
	Get Method = "GET"

	// Post POST HTTP method
	Post Method = "POST"
)

// RequestData holds the API relevant informations
type RequestData struct {
	Method      Method
	BaseURL     string
	Headers     map[string]string
	QueryParams map[string]string
	Body        []byte
}

// AddHeaders adds the provided headers to the request
func AddHeaders(headers map[string]string, body []byte, request *http.Request) {
	if len(body) > 0 {
		request.Header.Set("Content-Type", "application/json")
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
}

// AddQueryParams adds the query params to base URL
func AddQueryParams(baseURL string, queryParams map[string]string) string {
	baseURL += "?"
	params := url.Values{}
	for key, value := range queryParams {
		params.Add(key, value)
	}
	return baseURL + params.Encode()
}

// BuildRequest creates the request
func BuildRequest(requestData RequestData, request *http.Request) error {
	req, err := http.NewRequest(string(requestData.Method), requestData.BaseURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create new request: %v", err)
	}
	AddHeaders(requestData.Headers, requestData.Body, req)
	AddQueryParams(requestData.BaseURL, requestData.QueryParams)
	return nil
}
