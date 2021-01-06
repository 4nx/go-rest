package rest

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	method      Method = Get
	baseURL     string = "https://api.example.net"
	headers     map[string]string
	queryParams map[string]string
	body        []byte
)

var request = RequestData{
	Method:      method,
	BaseURL:     baseURL,
	Headers:     headers,
	QueryParams: queryParams,
	Body:        body,
}

func TestAddHeaders(t *testing.T) {
	t.Parallel()
	testKey1 := "testKey1"
	testValue1 := "testValue1"
	request.Headers = make(map[string]string)
	request.Headers[testKey1] = testValue1

	emptyBodyRequest, _ := http.NewRequest(string(request.Method), request.BaseURL, nil)
	AddHeaders(request.Headers, []byte(""), emptyBodyRequest)
	contentType := emptyBodyRequest.Header.Get("Content-Type")
	if contentType == "application/json" || len(contentType) != 0 {
		t.Error("Bad AddHeaders result: empty body request created Content-Type - should not exist")
	}
	testValidHeader := emptyBodyRequest.Header.Get(testKey1)
	if testValidHeader != testValue1 {
		t.Error(fmt.Sprintf("Bad AddHeaders result: header is not set correctly - should be %s", testValue1))
	}

	nonEmptyBodyRequest, _ := http.NewRequest(string(request.Method), request.BaseURL, nil)
	AddHeaders(request.Headers, []byte("ABCDEFG"), nonEmptyBodyRequest)
	testNonEmptyBodyRequest := nonEmptyBodyRequest.Header.Get("Content-Type")
	if testNonEmptyBodyRequest != "application/json" {
		t.Error("Bad AddHeaders result: non empty body request created no Content-Type - should be application/json")
	}
}
