package civogo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/civo/civogo/utils"
)

// Client is the means of connecting to the Civo API service
type Client struct {
	BaseURL          *url.URL
	UserAgent        string
	APIKey           string
	Region           string
	LastJSONResponse string

	httpClient *http.Client
}

// Component is a struct to define a User-Agent from a client
type Component struct {
	ID, Name, Version string
}

// HTTPError is the error returned when the API fails with an HTTP error
type HTTPError struct {
	Code   int
	Status string
	Reason string
}

// Result is the result of a SimpleResponse
type Result string

// SimpleResponse is a structure that returns success and/or any error
type SimpleResponse struct {
	ID           string `json:"id"`
	Result       Result `json:"result"`
	ErrorCode    string `json:"code"`
	ErrorReason  string `json:"reason"`
	ErrorDetails string `json:"details"`
}

// ConfigAdvanceClientForTesting initializes a Client connecting to a local test server and allows for specifying methods
type ConfigAdvanceClientForTesting struct {
	Method string
	Value  []ValueAdvanceClientForTesting
}

// ValueAdvanceClientForTesting is a struct that holds the URL and the request body
type ValueAdvanceClientForTesting struct {
	RequestBody  string
	URL          string
	ResponseBody string
}

// ResultSuccess represents a successful SimpleResponse
const ResultSuccess = "success"

func (e HTTPError) Error() string {
	return fmt.Sprintf("%d: %s, %s", e.Code, e.Status, e.Reason)
}

// NewClientWithURL initializes a Client with a specific API URL
func NewClientWithURL(apiKey, civoAPIURL, region string) (*Client, error) {
	if apiKey == "" {
		err := errors.New("no API Key supplied, this is required")
		return nil, NoAPIKeySuppliedError.wrap(err)
	}
	parsedURL, err := url.Parse(civoAPIURL)
	if err != nil {
		return nil, err
	}

	var httpTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
	}

	client := &Client{
		BaseURL:   parsedURL,
		UserAgent: "civogo/" + utils.GetVersion(),
		APIKey:    apiKey,
		Region:    region,
		httpClient: &http.Client{
			Transport: httpTransport,
		},
	}
	return client, nil
}

// NewClient initializes a Client connecting to the production API
func NewClient(apiKey, region string) (*Client, error) {
	return NewClientWithURL(apiKey, "https://api.civo.com", region)
}

// NewAdvancedClientForTesting initializes a Client connecting to a local test server and allows for specifying methods
func NewAdvancedClientForTesting(responses []ConfigAdvanceClientForTesting) (*Client, *httptest.Server, error) {
	var responseSent bool

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			return
		}

		req.Body = io.NopCloser(bytes.NewBuffer(body))

		for _, criteria := range responses {
			// we check the method first
			if criteria.Method == "PUT" || criteria.Method == "POST" || criteria.Method == "PATCH" {
				for _, criteria := range criteria.Value {
					if req.URL.Path == criteria.URL {
						if strings.TrimSpace(string(body)) == strings.TrimSpace(criteria.RequestBody) {
							responseSent = true
							rw.Write([]byte(criteria.ResponseBody))
						}
					}
				}
			} else {
				for _, criteria := range criteria.Value {
					if req.URL.Path == criteria.URL {
						responseSent = true
						rw.Write([]byte(criteria.ResponseBody))
					}
				}
			}
		}

		if !responseSent {
			fmt.Println("Failed to find a matching request!")
			fmt.Println("Request body:", string(body))
			fmt.Println("Method:", req.Method)
			fmt.Println("URL:", req.URL.String())
			rw.Write([]byte(`{"result": "failed to find a matching request"}`))
		}
	}))

	client, err := NewClientForTestingWithServer(server)

	return client, server, err
}

// NewClientForTesting initializes a Client connecting to a local test server
func NewClientForTesting(responses map[string]string) (*Client, *httptest.Server, error) {
	var responseSent bool

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		for url, response := range responses {
			if strings.Contains(req.URL.String(), url) {
				responseSent = true
				rw.Write([]byte(response))
			}
		}

		if !responseSent {
			fmt.Println("Failed to find a matching request!")
			fmt.Println("URL:", req.URL.String())

			rw.Write([]byte(`{"result": "failed to find a matching request"}`))
		}
	}))

	client, err := NewClientForTestingWithServer(server)

	return client, server, err
}

// NewClientForTestingWithServer initializes a Client connecting to a passed-in local test server
func NewClientForTestingWithServer(server *httptest.Server) (*Client, error) {
	client, err := NewClientWithURL("TEST-API-KEY", server.URL, "TEST")
	if err != nil {
		return nil, err
	}
	client.httpClient = server.Client()
	return client, err
}

func (c *Client) prepareClientURL(requestURL string) *url.URL {
	u, _ := url.Parse(c.BaseURL.String() + requestURL)
	return u
}

func (c *Client) sendRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("Authorization", fmt.Sprintf("bearer %s", c.APIKey))

	c.httpClient.Transport = &http.Transport{
		DisableCompression: false,
	}

	if req.Method == "GET" || req.Method == "DELETE" {
		// add the region param
		param := req.URL.Query()
		param.Add("region", c.Region)
		req.URL.RawQuery = param.Encode()
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	c.LastJSONResponse = string(body)

	if resp.StatusCode >= 300 {
		return nil, HTTPError{Code: resp.StatusCode, Status: resp.Status, Reason: string(body)}
	}

	return body, err
}

// SendGetRequest sends a correctly authenticated get request to the API server
func (c *Client) SendGetRequest(requestURL string) ([]byte, error) {
	u := c.prepareClientURL(requestURL)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)
}

// SendPostRequest sends a correctly authenticated post request to the API server
func (c *Client) SendPostRequest(requestURL string, params interface{}) ([]byte, error) {
	u := c.prepareClientURL(requestURL)

	// we create a new buffer and encode everything to json to send it in the request
	jsonValue, _ := json.Marshal(params)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req)
}

// SendPutRequest sends a correctly authenticated put request to the API server
func (c *Client) SendPutRequest(requestURL string, params interface{}) ([]byte, error) {
	u := c.prepareClientURL(requestURL)

	// we create a new buffer and encode everything to json to send it in the request
	jsonValue, _ := json.Marshal(params)

	req, err := http.NewRequest("PUT", u.String(), bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	return c.sendRequest(req)
}

// SendDeleteRequest sends a correctly authenticated delete request to the API server
func (c *Client) SendDeleteRequest(requestURL string) ([]byte, error) {
	u := c.prepareClientURL(requestURL)
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, err
	}

	return c.sendRequest(req)
}

// DecodeSimpleResponse parses a response body in to a SimpleResponse object
func (c *Client) DecodeSimpleResponse(resp []byte) (*SimpleResponse, error) {
	response := SimpleResponse{}
	err := json.NewDecoder(bytes.NewReader(resp)).Decode(&response)
	return &response, err
}

// SetUserAgent sets the user agent for the client
func (c *Client) SetUserAgent(component *Component) {
	if component.ID == "" {
		c.UserAgent = fmt.Sprintf("%s/%s %s", component.Name, component.Version, c.UserAgent)
	} else {
		c.UserAgent = fmt.Sprintf("%s/%s-%s %s", component.Name, component.Version, component.ID, c.UserAgent)
	}
}
