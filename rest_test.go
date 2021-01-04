package rest

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httputil"
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

func TestAddJSONContentType(t *testing.T) {
	t.Parallel()
	method := "GET"
	baseURL := "https://api.example.net"
	body1 := []byte("")
	body2 := []byte("A")
	req, _ := http.NewRequest(method, baseURL, bytes.NewBuffer(body1))

	testRequestEmpty := AddJSONContentType(req, body1)
	testHeaderEmpty := testRequestEmpty.Header.Get("Content-Type")
	if testHeaderEmpty != "" {
		t.Error("Bad AddContentType result: Content-Type header not empty")
	}

	testRequestBody := AddJSONContentType(req, body2)
	testHeaderBody := testRequestBody.Header.Get("Content-Type")
	if testHeaderBody != "application/json" {
		t.Error("Bad AddContentType result: Content-Type header not application/json")
	}
}

func TestCreateRequest(t *testing.T) {
	t.Parallel()
	method := Post
	baseURL := "https://api.example.net"
	body := []byte("")
	token := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Headers := make(map[string]string)
	Headers["Content-Type"] = "application/json"
	Headers["Authorization"] = "Bearer " + token
	queryParams := make(map[string]string)
	queryParams["test1"] = "1"
	queryParams["test2"] = "2"
	request := Request{
		Method:      method,
		BaseURL:     baseURL,
		Headers:     Headers,
		QueryParams: queryParams,
		Body:        body,
	}

	req, err := CreateRequest(request)
	if err != nil {
		t.Errorf("Bad CreateRequest result - Returned error: %v", err)
	}
	if req == nil {
		t.Errorf("Bad CreateRequest result: Return is nil")
	}

	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		t.Errorf("Error : %v", err)
	}
	fmt.Println("Request : ", string(requestDump))
}
