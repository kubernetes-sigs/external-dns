// Package dnsimple provides a client for the DNSimple API.
// In order to use this package you will need a DNSimple account.
package dnsimple

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	// Version identifies the current library version.
	// This is a pro-forma convention given that Go dependencies
	// tends to be fetched directly from the repo.
	// It is also used in the user-agent identify the client.
	Version = "0.60.0"

	// defaultBaseURL to the DNSimple production API.
	defaultBaseURL = "https://api.dnsimple.com"

	// userAgent represents the default user agent used
	// when no other user agent is set.
	defaultUserAgent = "dnsimple-go/" + Version

	apiVersion = "v2"
)

// Client represents a client to the DNSimple API.
type Client struct {
	// httpClient is the underlying HTTP client
	// used to communicate with the API.
	httpClient *http.Client

	// BaseURL for API requests.
	// Defaults to the public DNSimple API, but can be set to a different endpoint (e.g. the sandbox).
	BaseURL string

	// UserAgent used when communicating with the DNSimple API.
	UserAgent string

	// Services used for talking to different parts of the DNSimple API.
	Identity          *IdentityService
	Accounts          *AccountsService
	Certificates      *CertificatesService
	Contacts          *ContactsService
	Domains           *DomainsService
	Oauth             *OauthService
	Registrar         *RegistrarService
	Services          *ServicesService
	Templates         *TemplatesService
	Tlds              *TldsService
	VanityNameServers *VanityNameServersService
	Webhooks          *WebhooksService
	Zones             *ZonesService

	// Set to true to output debugging logs during API calls
	Debug bool
}

// ListOptions contains the common options you can pass to a List method
// in order to control parameters such as pagination and page number.
type ListOptions struct {
	// The page to return
	Page *int `url:"page,omitempty"`

	// The number of entries to return per page
	PerPage *int `url:"per_page,omitempty"`

	// The order criteria to sort the results.
	// The value is a comma-separated list of field[:direction],
	// eg. name | name:desc | name:desc,expiration:desc
	Sort *string `url:"sort,omitempty"`
}

// NewClient returns a new DNSimple API client.
//
// To authenticate you must provide an http.Client that will perform authentication
// for you with one of the currently supported mechanisms: OAuth or HTTP Basic.
func NewClient(httpClient *http.Client) *Client {
	c := &Client{httpClient: httpClient, BaseURL: defaultBaseURL}
	c.Identity = &IdentityService{client: c}
	c.Accounts = &AccountsService{client: c}
	c.Certificates = &CertificatesService{client: c}
	c.Contacts = &ContactsService{client: c}
	c.Domains = &DomainsService{client: c}
	c.Oauth = &OauthService{client: c}
	c.Registrar = &RegistrarService{client: c}
	c.Services = &ServicesService{client: c}
	c.Templates = &TemplatesService{client: c}
	c.Tlds = &TldsService{client: c}
	c.VanityNameServers = &VanityNameServersService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Zones = &ZonesService{client: c}
	return c
}

// SetUserAgent overrides the default UserAgent.
//
// When a custom user agent is provided, the final user agent is the combination of the custom user agent
// prepended by the default user agent.
//
//     customAgentFlag dnsimple-go/1.0
//
func (c *Client) SetUserAgent(ua string) {
	c.UserAgent = ua
}

// formatUserAgent builds the final user agent to use for HTTP requests.
//
// If no custom user agent is provided, the default user agent is used.
//
//     dnsimple-go/1.0
//
// If a custom user agent is provided, the final user agent is the combination of the custom user agent
// prepended by the default user agent.
//
//     customAgentFlag dnsimple-go/1.0
//
func formatUserAgent(customUserAgent string) string {
	if customUserAgent == "" {
		return defaultUserAgent
	}

	return fmt.Sprintf("%s %s", customUserAgent, defaultUserAgent)
}

func versioned(path string) string {
	return fmt.Sprintf("/%s/%s", apiVersion, strings.Trim(path, "/"))
}

func (c *Client) get(ctx context.Context, path string, obj interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodGet, path, nil, obj, nil)
}

func (c *Client) post(ctx context.Context, path string, payload, obj interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodPost, path, payload, obj, nil)
}

func (c *Client) put(ctx context.Context, path string, payload, obj interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodPut, path, payload, obj, nil)
}

func (c *Client) patch(ctx context.Context, path string, payload, obj interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodPatch, path, payload, obj, nil)
}

func (c *Client) delete(ctx context.Context, path string, payload, obj interface{}) (*http.Response, error) {
	return c.makeRequest(ctx, http.MethodDelete, path, payload, obj, nil)
}

// Request executes an API request with the current client scope, and returns the response.
func (c *Client) Request(ctx context.Context, method, path string, payload, obj interface{}, headers http.Header) (*http.Response, error) {
	return c.makeRequest(ctx, method, path, payload, obj, headers)
}

// makeRequest executes an API request and returns the HTTP response.
//
// The content pointed by payload is serialized and used as body of the request.
// The HTTP response is JSON decoded and stored in the value pointed by obj.
func (c *Client) makeRequest(ctx context.Context, method, path string, payload, obj interface{}, headers http.Header) (*http.Response, error) {
	req, err := c.newRequestWithHeaders(method, path, payload, headers)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		log.Printf("Request (%v): %#v", req.URL, req)
	}

	resp, err := c.request(ctx, req, obj)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		log.Printf("Response: %#v", resp)
	}

	return resp, nil
}

// newRequest creates an API request.
//
// The path is expected to be a relative path and will be resolved according to the BaseURL of the Client.
// Paths should always be specified without a preceding slash.
func (c *Client) newRequest(method, path string, payload interface{}) (*http.Request, error) {
	return c.newRequestWithHeaders(method, path, payload, nil)
}

// newRequestWithHeaders creates an API request, with custom headers.
func (c *Client) newRequestWithHeaders(method, path string, payload interface{}, headers http.Header) (*http.Request, error) {
	url := c.BaseURL + path

	body := new(bytes.Buffer)
	if payload != nil {
		err := json.NewEncoder(body).Encode(payload)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", formatUserAgent(c.UserAgent))

	return req, err
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}

// request sends the HTTP request and returns the HTTP response.
//
// The HTTP response is JSON decoded and stored in the value pointed by obj,
// or returned as an error if an API error has occurred.
// If obj implements the io.Writer interface, the raw response body will be written to obj,
// without attempting to decode it.
func (c *Client) request(ctx context.Context, req *http.Request, obj interface{}) (*http.Response, error) {
	if ctx == nil {
		return nil, errors.New("context must be non-nil")
	}
	req = req.WithContext(ctx)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return resp, err
	}

	// If obj implements the io.Writer,
	// the response body is decoded into v.
	if obj != nil {
		if w, ok := obj.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
		} else {
			err = json.NewDecoder(resp.Body).Decode(obj)
		}
	}

	return resp, err
}

// A Response represents an API response.
type Response struct {
	// HTTP response
	HTTPResponse *http.Response

	// If the response is paginated, the Pagination will store them.
	Pagination *Pagination `json:"pagination"`
}

// RateLimit returns the maximum amount of requests this account can send in an hour.
func (r *Response) RateLimit() int {
	value, _ := strconv.Atoi(r.HTTPResponse.Header.Get("X-RateLimit-Limit"))
	return value
}

// RateLimitRemaining returns the remaining amount of requests this account can send within this hour window.
func (r *Response) RateLimitRemaining() int {
	value, _ := strconv.Atoi(r.HTTPResponse.Header.Get("X-RateLimit-Remaining"))
	return value
}

// RateLimitReset returns when the throttling window will be reset for this account.
func (r *Response) RateLimitReset() time.Time {
	value, _ := strconv.ParseInt(r.HTTPResponse.Header.Get("X-RateLimit-Reset"), 10, 64)
	return time.Unix(value, 0)
}

// Pagination represents the pagination information, if the response is paginated.
type Pagination struct {
	CurrentPage  int `json:"current_page"`
	PerPage      int `json:"per_page"`
	TotalPages   int `json:"total_pages"`
	TotalEntries int `json:"total_entries"`
}

// An ErrorResponse represents an API response that generated an error.
type ErrorResponse struct {
	Response

	// human-readable message
	Message string `json:"message"`
}

// Error implements the error interface.
func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %v %v",
		r.HTTPResponse.Request.Method, r.HTTPResponse.Request.URL,
		r.HTTPResponse.StatusCode, r.Message)
}

// CheckResponse checks the API response for errors, and returns them if present.
// A response is considered an error if the status code is different than 2xx. Specific requests
// may have additional requirements, but this is sufficient in most of the cases.
func CheckResponse(resp *http.Response) error {
	if code := resp.StatusCode; 200 <= code && code <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{}
	errorResponse.HTTPResponse = resp

	err := json.NewDecoder(resp.Body).Decode(errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

// addOptions adds the parameters in opt as URL query parameters to s.  opt
// must be a struct whose fields may contain "url" tags.
func addURLQueryOptions(path string, options interface{}) (string, error) {
	opt := reflect.ValueOf(options)

	// options is a pointer
	// return if the value of the pointer is nil,
	if opt.Kind() == reflect.Ptr && opt.IsNil() {
		return path, nil
	}

	// append the options to the URL
	u, err := url.Parse(path)
	if err != nil {
		return path, err
	}

	qs, err := query.Values(options)
	if err != nil {
		return path, err
	}

	uqs := u.Query()
	for k := range qs {
		uqs.Set(k, qs.Get(k))
	}
	u.RawQuery = uqs.Encode()

	return u.String(), nil
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }
