package rest

import (
	"net/http"
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

// AddJSONContentType adds JSON content type if body data will be send
func AddJSONContentType(body []byte, request *http.Request) {
	if len(body) > 0 {
		request.Header.Set("Content-Type", "application/json")
	}
}

// AddHeaders adds the provided headers to the request
func AddHeaders(headers map[string]string, request *http.Request) {
	for key, value := range headers {
		request.Header.Set(key, value)
	}
}
