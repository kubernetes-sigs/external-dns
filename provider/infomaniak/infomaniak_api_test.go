package infomaniak

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

var GetDoFunc func(req *http.Request) (*http.Response, error)

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

func TestNewInfomaniakAPI(t *testing.T) {
	api := NewInfomaniakAPI("TestToken")
	assert.Equal(t, api.apiToken, "TestToken")
}

func TestRequestInvalidMethod(t *testing.T) {
	api := NewInfomaniakAPI("OkToken")
	_, err := api.request("INVALID METHOD", "/", nil)
	assert.NotEqual(t, err, nil)
}

func TestRequest(t *testing.T) {
	api := NewInfomaniakAPI("OkToken")

	Client = &MockClient{}

	// Test that the token is used
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, req.Header.Get("Authorization"), "Bearer OkToken")
		// on get request, this method must match
		assert.Equal(t, req.Method, "GET")
		// Path always starts with a /
		assert.Equal(t, strings.HasPrefix(req.URL.Path, "/"), true)
		// simulates server unreachable or something
		return nil, errors.New(
			"Error from web server",
		)
	}

	// Test that the request produces an error on server error
	_, err := api.request("GET", "/url", nil)
	assert.NotEqual(t, err, nil)
	api.request("GET", "url", nil)

	// simulates server wrong response format
	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		var json string
		if req.Header.Get("Authorization") != "Bearer OkToken" {
			json = `{"result":"error","error":{"code":"not_authorized","description":"Not authorized."}}`
		} else {
			switch req.URL.Path {
			case "/format_error":
				json = `"oh my not json"`
			case "/not_success":
				json = `{"result": "super result", "data": "oh my data"}`
			default:
				json = `{"result": "success", "data": "oh yeah"}`
			}
		}
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	_, err = api.request("GET", "/format_error", nil)
	assert.NotEqual(t, err, nil)

	_, err = api.request("GET", "/not_success", nil)
	assert.NotEqual(t, err, nil)

	_, err = api.request("GET", "/no_error", nil)
	assert.Equal(t, err, nil)

	api.apiToken = "NotOkToken"
	_, err = api.request("GET", "/no_error", nil)
	assert.NotEqual(t, err, nil)
}

func TestMethods(t *testing.T) {
	Client = &MockClient{}
	api := NewInfomaniakAPI("OkToken")
	var Method string

	GetDoFunc = func(req *http.Request) (*http.Response, error) {
		assert.Equal(t, req.Method, Method)
		json := `{"result": "success", "data": "oh yeah"}`
		r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}
	Method = "GET"
	api.get("abc", nil)

	// fail with control chars
	_, err := api.get("abc\ndef", nil)
	assert.NotEqual(t, err, nil, "should fail with control chars")

	_, err = api.get("abc", url.Values{"should":{"def"}})
	assert.Equal(t, err, nil)

	Method = "POST"
	api.post("abc", bytes.NewBufferString(""))
	assert.Equal(t, err, nil)

	Method = "PUT"
	api.put("abc", bytes.NewBufferString(""))
	assert.Equal(t, err, nil)

	Method = "DELETE"
	api.delete("abc")
	assert.Equal(t, err, nil)
}

func MockServer(req *http.Request) (*http.Response, error) {
	var json string
	switch req.Method {
	case "GET":
		switch req.URL.Path {
		case "/1/product" :
			customer_name := req.URL.Query()["customer_name"]
			if customer_name != nil {
				json = `{"result":"ok", "data":[
					{
						"id": 1234567,
						"service_name": "domain",
						"customer_name": "def"
					}]}`
			} else {
				json = `{"result":"ok", "data":[
					{
						"id": 1234567,
						"service_name": "domain",
						"customer_name": "domain1.com"
					},
					{
						"id": 1234568,
						"service_name": "domain",
						"customer_name": "domain2.com"
					}]}`
			}
		}

		fmt.Println(req.URL.RawQuery)
		json = `"oh my not json"`
	}
	switch req.URL.Path {
	case "/format_error":
		json = `"oh my not json"`
	case "/not_success":
		json = `{"result": "super result", "data": "oh my data"}`
	default:
		json = `{"result": "success", "data": "oh yeah"}`
	}
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))
	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestGetDomainByName(t *testing.T) {
	Client = &MockClient{}
	api := NewInfomaniakAPI("OkToken")
	GetDoFunc = MockServer

	api.GetDomainByName("abc.com")

}

