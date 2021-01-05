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

func TestAddJSONContentType(t *testing.T) {
	t.Parallel()

	emptyBodyRequest, _ := http.NewRequest(string(request.Method), request.BaseURL, nil)
	testEmptyBodyRequest, err := AddJSONContentType([]byte(""), emptyBodyRequest)
	if err != nil {
		t.Error("Bad AddJSONContentType result: empty body request thrown error - should be nil")
	}
	contentType := testEmptyBodyRequest.Header.Get("Content-Type")
	if contentType == "application/json" || len(contentType) != 0 {
		t.Error("Bad AddJSONContentType result: empty body request created Content-Type - should not exist")
	}

	nonEmptyBodyRequest, _ := http.NewRequest(string(request.Method), request.BaseURL, nil)
	testNonEmptyBodyRequest, err := AddJSONContentType([]byte("ABCDEFG"), nonEmptyBodyRequest)
	if err != nil {
		t.Error("Bad AddJSONContentType result: Header set - should be nil")
	}
	if _, nonEmptyBodyExists := testNonEmptyBodyRequest.Header["Content-Type"]; !nonEmptyBodyExists {
		t.Error("Bad AddJSONContentType result: non empty body request created no Content-Type - should be application/json")
	}

	existentContentTypeRequest, _ := http.NewRequest(string(request.Method), request.BaseURL, nil)
	existentContentTypeRequest.Header.Set("Content-Type", "application/json")
	existentContenTypeRequest, err := AddJSONContentType([]byte(""), existentContentTypeRequest)
	if err == nil {
		t.Error("Bad AddJSONContentType result: no error was thrown - should not be nil")
	}
	if existentContenTypeRequest != nil {
		t.Error("Bad AddJSONContentType result: request was given back - should be nil")
	}
}

func TestAddHeaders(t *testing.T) {
	t.Parallel()
	testKey1 := "testKey1"
	testValue1 := "testValue1"
	request.Headers = make(map[string]string)
	request.Headers[testKey1] = testValue1

	validHeaderRequest, _ := http.NewRequest(string(request.Method), request.BaseURL, nil)
	testValidHeaderRequest, err := AddHeaders(request.Headers, validHeaderRequest)
	if err != nil {
		t.Error("Bad AddHeaders result: error was thrown - should be nil")
	}
	testValidHeader := testValidHeaderRequest.Header.Get(testKey1)
	if testValidHeader != testValue1 {
		t.Error(fmt.Sprintf("Bad AddHeaders result: header is not set correctly - should be %s", testValue1))
	}
}
