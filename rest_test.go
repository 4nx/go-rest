package rest

import (
	"fmt"
	"net/http"
	"testing"
)

var (
	method      = Get
	baseURL     = "https://api.example.net"
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

func TestAddQueryParams(t *testing.T) {
	t.Parallel()
	queryKey1 := "testQueryKey1"
	queryKey2 := "testQueryKey2"
	queryValue1 := "testQueryValue1"
	queryValue2 := "testQueryValue2"
	request.QueryParams = make(map[string]string)
	request.QueryParams[queryKey1] = queryValue1
	request.QueryParams[queryKey2] = queryValue2

	testQueryURL := AddQueryParams(request.BaseURL, request.QueryParams)
	if testQueryURL != request.BaseURL+"?"+queryKey1+"="+queryValue1+"&"+queryKey2+"="+queryValue2 {
		t.Error("Bad AddQueryParams result: query url not correct")
	}
}
