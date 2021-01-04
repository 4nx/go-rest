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
	method := "GET"
	baseURL := "http://api.example.net"
	body1 := []byte("")
	body2 := []byte("A")
	req, _ := http.NewRequest(method, baseURL, bytes.NewBuffer(body1))

	testRequestEmpty := AddContentType(req, body1)
	testHeaderEmpty := testRequestEmpty.Header.Get("Content-Type")
	if testHeaderEmpty != "" {
		t.Error("Bad AddContentType result: Content-Type header not empty")
	}

	testRequestBody := AddContentType(req, body2)
	testHeaderBody := testRequestBody.Header.Get("Content-Type")
	if testHeaderBody != "application/json" {
		t.Error("Bad AddContentType result: Content-Type header not application/json")
	}
}
