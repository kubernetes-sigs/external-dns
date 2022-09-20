/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package godaddy

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

// DefaultTimeout api requests after 180s
const DefaultTimeout = 180 * time.Second

// Errors
var (
	ErrAPIDown = errors.New("godaddy: the GoDaddy API is down")
)

// APIError error
type APIError struct {
	Code    string
	Message string
}

func (err *APIError) Error() string {
	return fmt.Sprintf("Error %s: %q", err.Code, err.Message)
}

// Logger is the interface that should be implemented for loggers that wish to
// log HTTP requests and HTTP responses.
type Logger interface {
	// LogRequest logs an HTTP request.
	LogRequest(*http.Request)

	// LogResponse logs an HTTP response.
	LogResponse(*http.Response)
}

// Client represents a client to call the GoDaddy API
type Client struct {
	// APIKey holds the Application key
	APIKey string

	// APISecret holds the Application secret key
	APISecret string

	// API endpoint
	APIEndPoint string

	// Client is the underlying HTTP client used to run the requests. It may be overloaded but a default one is instanciated in ``NewClient`` by default.
	Client *http.Client

	// Logger is used to log HTTP requests and responses.
	Logger Logger

	Timeout time.Duration
}

// GDErrorField describe the error reason
type GDErrorField struct {
	Code        string `json:"code,omitempty"`
	Message     string `json:"message,omitempty"`
	Path        string `json:"path,omitempty"`
	PathRelated string `json:"pathRelated,omitempty"`
}

// GDErrorResponse is the body response when an API call fails
type GDErrorResponse struct {
	Code    string         `json:"code"`
	Fields  []GDErrorField `json:"fields,omitempty"`
	Message string         `json:"message,omitempty"`
}

func (r GDErrorResponse) String() string {
	if b, err := json.Marshal(r); err == nil {
		return string(b)
	}

	return "<error>"
}

// NewClient represents a new client to call the API
func NewClient(useOTE bool, apiKey, apiSecret string) (*Client, error) {
	var endpoint string

	if useOTE {
		endpoint = " https://api.ote-godaddy.com"
	} else {
		endpoint = "https://api.godaddy.com"
	}

	client := Client{
		APIKey:      apiKey,
		APISecret:   apiSecret,
		APIEndPoint: endpoint,
		Client:      &http.Client{},
		Timeout:     DefaultTimeout,
	}

	// Get and check the configuration
	if err := client.validate(); err != nil {
		return nil, err
	}
	return &client, nil
}

//
// Common request wrappers
//

// Get is a wrapper for the GET method
func (c *Client) Get(url string, resType interface{}) error {
	return c.CallAPI("GET", url, nil, resType, true)
}

// Patch is a wrapper for the POST method
func (c *Client) Patch(url string, reqBody, resType interface{}) error {
	return c.CallAPI("PATCH", url, reqBody, resType, true)
}

// Post is a wrapper for the POST method
func (c *Client) Post(url string, reqBody, resType interface{}) error {
	return c.CallAPI("POST", url, reqBody, resType, true)
}

// Put is a wrapper for the PUT method
func (c *Client) Put(url string, reqBody, resType interface{}) error {
	return c.CallAPI("PUT", url, reqBody, resType, true)
}

// Delete is a wrapper for the DELETE method
func (c *Client) Delete(url string, resType interface{}) error {
	return c.CallAPI("DELETE", url, nil, resType, true)
}

// GetWithContext is a wrapper for the GET method
func (c *Client) GetWithContext(ctx context.Context, url string, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "GET", url, nil, resType, true)
}

// PatchWithContext is a wrapper for the POST method
func (c *Client) PatchWithContext(ctx context.Context, url string, reqBody, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "PATCH", url, reqBody, resType, true)
}

// PostWithContext is a wrapper for the POST method
func (c *Client) PostWithContext(ctx context.Context, url string, reqBody, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "POST", url, reqBody, resType, true)
}

// PutWithContext is a wrapper for the PUT method
func (c *Client) PutWithContext(ctx context.Context, url string, reqBody, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "PUT", url, reqBody, resType, true)
}

// DeleteWithContext is a wrapper for the DELETE method
func (c *Client) DeleteWithContext(ctx context.Context, url string, resType interface{}) error {
	return c.CallAPIWithContext(ctx, "DELETE", url, nil, resType, true)
}

// NewRequest returns a new HTTP request
func (c *Client) NewRequest(method, path string, reqBody interface{}, needAuth bool) (*http.Request, error) {
	var body []byte
	var err error

	if reqBody != nil {
		body, err = json.Marshal(reqBody)
		if err != nil {
			return nil, err
		}
	}

	target := fmt.Sprintf("%s%s", c.APIEndPoint, path)
	req, err := http.NewRequest(method, target, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	// Inject headers
	if body != nil {
		req.Header.Set("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Set("Authorization", fmt.Sprintf("sso-key %s:%s", c.APIKey, c.APISecret))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "ExternalDNS/"+externaldns.Version)

	// Send the request with requested timeout
	c.Client.Timeout = c.Timeout

	return req, nil
}

// Do sends an HTTP request and returns an HTTP response
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	if c.Logger != nil {
		c.Logger.LogRequest(req)
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if c.Logger != nil {
		c.Logger.LogResponse(resp)
	}
	return resp, nil
}

// CallAPI is the lowest level call helper. If needAuth is true,
// inject authentication headers and sign the request.
//
// Request signature is a sha1 hash on following fields, joined by '+':
// - applicationSecret (from Client instance)
// - consumerKey (from Client instance)
// - capitalized method (from arguments)
// - full request url, including any query string argument
// - full serialized request body
// - server current time (takes time delta into account)
//
// Call will automatically assemble the target url from the endpoint
// configured in the client instance and the path argument. If the reqBody
// argument is not nil, it will also serialize it as json and inject
// the required Content-Type header.
//
// If everything went fine, unmarshall response into resType and return nil
// otherwise, return the error
func (c *Client) CallAPI(method, path string, reqBody, resType interface{}, needAuth bool) error {
	return c.CallAPIWithContext(context.Background(), method, path, reqBody, resType, needAuth)
}

// CallAPIWithContext is the lowest level call helper. If needAuth is true,
// inject authentication headers and sign the request.
//
// Request signature is a sha1 hash on following fields, joined by '+':
// - applicationSecret (from Client instance)
// - consumerKey (from Client instance)
// - capitalized method (from arguments)
// - full request url, including any query string argument
// - full serialized request body
// - server current time (takes time delta into account)
//
// # Context is used by http.Client to handle context cancelation
//
// Call will automatically assemble the target url from the endpoint
// configured in the client instance and the path argument. If the reqBody
// argument is not nil, it will also serialize it as json and inject
// the required Content-Type header.
//
// If everything went fine, unmarshall response into resType and return nil
// otherwise, return the error
func (c *Client) CallAPIWithContext(ctx context.Context, method, path string, reqBody, resType interface{}, needAuth bool) error {
	req, err := c.NewRequest(method, path, reqBody, needAuth)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	response, err := c.Do(req)
	if err != nil {
		return err
	}
	return c.UnmarshalResponse(response, resType)
}

// UnmarshalResponse checks the response and unmarshals it into the response
// type if needed Helper function, called from CallAPI
func (c *Client) UnmarshalResponse(response *http.Response, resType interface{}) error {
	// Read all the response body
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// < 200 && >= 300 : API error
	if response.StatusCode < http.StatusOK || response.StatusCode >= http.StatusMultipleChoices {
		apiError := &APIError{
			Code: fmt.Sprintf("HTTPStatus: %d", response.StatusCode),
		}

		if err = json.Unmarshal(body, apiError); err != nil {
			return err
		}

		return apiError
	}

	// Nothing to unmarshal
	if len(body) == 0 || resType == nil {
		return nil
	}

	return json.Unmarshal(body, &resType)
}

func (c *Client) validate() error {
	var response interface{}

	if err := c.Get("/v1/domains?statuses=ACTIVE", response); err != nil {
		return err
	}

	return nil
}
