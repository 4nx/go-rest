package rest

import (
	"bytes"
	"net/http"
	"testing"
)

func TestAddQueryParams(t *testing.T) {
	t.Parallel()
	apiURL := "https://api.example.net"
	queryParams := make(map[string]string)
	queryParams["testKey1"] = "testValue1"
	queryParams["testKey2"] = "testValue2"
	testURL := AddQueryParams(apiURL, queryParams)
	if testURL != apiURL+"?testKey1=testValue1&testKey2=testValue2" {
		t.Error("Bad AddQueryParams result")
	}
}

func TestAddContentType(t *testing.T) {
	t.Parallel()
	method := Get
	baseURL := "http://api.example.net"
	body := []byte("")
	request := Request{
		Method:  method,
		BaseURL: baseURL,
		Body:    body,
	}
	req, _ := http.NewRequest(string(request.Method), request.BaseURL, bytes.NewBuffer(request.Body))

	testRequestEmpty := AddContentType(req, request.Body)
	testHeaderEmpty := testRequestEmpty.Header.Get("Content-Type")
	if testHeaderEmpty != "" {
		t.Error("Bad AddContentType result")
	}

	request.Body = []byte("A")
	testRequestBody := AddContentType(req, request.Body)
	testHeaderBody := testRequestBody.Header.Get("Content-Type")
	if testHeaderBody != "application/json" {
		t.Error("Bad AddContentType result")
	}
}
