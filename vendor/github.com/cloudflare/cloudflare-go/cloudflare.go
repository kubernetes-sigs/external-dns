// Package cloudflare implements the Cloudflare v4 API.
package cloudflare

import (
	"bytes"
	"context"
<<<<<<< HEAD
	"encoding/json"
<<<<<<< HEAD
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"errors"

	"golang.org/x/time/rate"
)

var (
	Version string = "v4"

<<<<<<< HEAD
const (
<<<<<<< HEAD
<<<<<<< HEAD
	originCARootCertEccURL = "https://developers.cloudflare.com/ssl/0d2cd0f374da0fb6dbf53128b60bbbf7/origin_ca_ecc_root.pem"
	originCARootCertRsaURL = "https://developers.cloudflare.com/ssl/e2b9968022bf23b071d95229b5678452/origin_ca_rsa_root.pem"
||||||| parent of e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
const (
	originCARootCertEccURL = "https://developers.cloudflare.com/ssl/0d2cd0f374da0fb6dbf53128b60bbbf7/origin_ca_ecc_root.pem"
	originCARootCertRsaURL = "https://developers.cloudflare.com/ssl/e2b9968022bf23b071d95229b5678452/origin_ca_rsa_root.pem"
=======
	// Deprecated: Use `client.New` configuration instead.
	apiURL = fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath)
>>>>>>> e1cd8261c (UPSTREAM: <carry>: update vendored files v0.13.1)
)

const (
	// AuthKeyEmail specifies that we should authenticate with API key and email address.
	AuthKeyEmail = 1 << iota
	// AuthUserService specifies that we should authenticate with a User-Service key.
	AuthUserService
	// AuthToken specifies that we should authenticate with an API Token.
	AuthToken
)

// API holds the configuration for the current API client. A client should not
// be modified concurrently.
type API struct {
	APIKey            string
	APIEmail          string
	APIUserServiceKey string
	APIToken          string
	BaseURL           string
	AccountID         string
	UserAgent         string
	headers           http.Header
	httpClient        *http.Client
	authType          int
	rateLimiter       *rate.Limiter
	retryPolicy       RetryPolicy
	logger            Logger
	Debug             bool
}

// newClient provides shared logic for New and NewWithUserServiceKey.
func newClient(opts ...Option) (*API, error) {
	silentLogger := log.New(ioutil.Discard, "", log.LstdFlags)

	api := &API{
		BaseURL:     fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath),
		UserAgent:   userAgent + "/" + Version,
		headers:     make(http.Header),
		rateLimiter: rate.NewLimiter(rate.Limit(4), 1), // 4rps equates to default api limit (1200 req/5 min)
		retryPolicy: RetryPolicy{
			MaxRetries:    3,
			MinRetryDelay: time.Duration(1) * time.Second,
			MaxRetryDelay: time.Duration(30) * time.Second,
		},
		logger: silentLogger,
	}

	err := api.parseOptions(opts...)
	if err != nil {
		return nil, fmt.Errorf("options parsing failed: %w", err)
	}

	// Fall back to http.DefaultClient if the package user does not provide
	// their own.
	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	}

	return api, nil
}

// New creates a new Cloudflare v4 API client.
func New(key, email string, opts ...Option) (*API, error) {
	if key == "" || email == "" {
		return nil, errors.New(errEmptyCredentials)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIKey = key
	api.APIEmail = email
	api.authType = AuthKeyEmail

	return api, nil
}

// NewWithAPIToken creates a new Cloudflare v4 API client using API Tokens.
func NewWithAPIToken(token string, opts ...Option) (*API, error) {
	if token == "" {
		return nil, errors.New(errEmptyAPIToken)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIToken = token
	api.authType = AuthToken

	return api, nil
}

// NewWithUserServiceKey creates a new Cloudflare v4 API client using service key authentication.
func NewWithUserServiceKey(key string, opts ...Option) (*API, error) {
	if key == "" {
		return nil, errors.New(errEmptyCredentials)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIUserServiceKey = key
	api.authType = AuthUserService

	return api, nil
}

// SetAuthType sets the authentication method (AuthKeyEmail, AuthToken, or AuthUserService).
func (api *API) SetAuthType(authType int) {
	api.authType = authType
}

// ZoneIDByName retrieves a zone's ID from the name.
func (api *API) ZoneIDByName(zoneName string) (string, error) {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	zoneName = normalizeZoneName(zoneName)
	res, err := api.ListZonesContext(context.TODO(), WithZoneFilters(zoneName, "", ""))
	if err != nil {
		return "", errors.Wrap(err, "ListZonesContext command failed")
	}

	if len(res.Result) > 1 && api.AccountID == "" {
		return "", errors.New("ambiguous zone name used without an account ID")
	}

	for _, zone := range res.Result {
		if api.AccountID != "" {
			if zone.Name == zoneName && api.AccountID == zone.Account.ID {
				return zone.ID, nil
			}
		} else {
			if zone.Name == zoneName {
				return zone.ID, nil
			}
		}
	}

	return "", errors.New("Zone could not be found")
}

// makeRequest makes a HTTP request and returns the body as a byte slice,
// closing it before returning. params will be serialized to JSON.
func (api *API) makeRequest(method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(context.TODO(), method, uri, params, api.authType)
}

func (api *API) makeRequestContext(ctx context.Context, method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(ctx, method, uri, params, api.authType)
}

func (api *API) makeRequestWithHeaders(method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(context.TODO(), method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithAuthType(ctx context.Context, method, uri string, params interface{}, authType int) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, authType, nil)
}

func (api *API) makeRequestWithAuthTypeAndHeaders(ctx context.Context, method, uri string, params interface{}, authType int, headers http.Header) ([]byte, error) {
	// Replace nil with a JSON object if needed
	var jsonBody []byte
	var err error

	if params != nil {
		if paramBytes, ok := params.([]byte); ok {
			jsonBody = paramBytes
		} else {
			jsonBody, err = json.Marshal(params)
			if err != nil {
				return nil, errors.Wrap(err, "error marshalling params to JSON")
			}
		}
	} else {
		jsonBody = nil
	}

	var resp *http.Response
	var respErr error
	var reqBody io.Reader
	var respBody []byte
	for i := 0; i <= api.retryPolicy.MaxRetries; i++ {
		if jsonBody != nil {
			reqBody = bytes.NewReader(jsonBody)
		}
		if i > 0 {
			// expect the backoff introduced here on errored requests to dominate the effect of rate limiting
			// don't need a random component here as the rate limiter should do something similar
			// nb time duration could truncate an arbitrary float. Since our inputs are all ints, we should be ok
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(api.retryPolicy.MinRetryDelay))

			if sleepDuration > api.retryPolicy.MaxRetryDelay {
				sleepDuration = api.retryPolicy.MaxRetryDelay
			}
			// useful to do some simple logging here, maybe introduce levels later
			api.logger.Printf("Sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)
			time.Sleep(sleepDuration)

		}
		err = api.rateLimiter.Wait(context.TODO())
		if err != nil {
			return nil, errors.Wrap(err, "Error caused by request rate limiting")
		}
		resp, respErr = api.request(ctx, method, uri, reqBody, authType, headers)

		// retry if the server is rate limiting us or if it failed
		// assumes server operations are rolled back on failure
		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			// if we got a valid http response, try to read body so we can reuse the connection
			// see https://golang.org/pkg/net/http/#Client.Do
			if respErr == nil {
				respBody, err = ioutil.ReadAll(resp.Body)
				resp.Body.Close()

				respErr = errors.Wrap(err, "could not read response body")

				api.logger.Printf("Request: %s %s got an error response %d: %s\n", method, uri, resp.StatusCode,
					strings.Replace(strings.Replace(string(respBody), "\n", "", -1), "\t", "", -1))
			} else {
				api.logger.Printf("Error performing request: %s %s : %s \n", method, uri, respErr.Error())
			}
			continue
		} else {
			respBody, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, errors.Wrap(err, "could not read response body")
			}
			break
		}
	}
	if respErr != nil {
		return nil, respErr
	}

	switch {
	case resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices:
	case resp.StatusCode == http.StatusUnauthorized:
		return nil, errorFromResponse(resp.StatusCode, respBody)
	case resp.StatusCode == http.StatusForbidden:
		return nil, errorFromResponse(resp.StatusCode, respBody)
	case resp.StatusCode == http.StatusServiceUnavailable,
		resp.StatusCode == http.StatusBadGateway,
		resp.StatusCode == http.StatusGatewayTimeout,
		resp.StatusCode == 522,
		resp.StatusCode == 523,
		resp.StatusCode == 524:
		return nil, errors.Errorf("HTTP status %d: service failure", resp.StatusCode)
	// This isn't a great solution due to the way the `default` case is
	// a catch all and that the `filters/validate-expr` returns a HTTP 400
	// yet the clients need to use the HTTP body as a JSON string.
	case resp.StatusCode == 400 && strings.HasSuffix(resp.Request.URL.Path, "/filters/validate-expr"):
		return nil, errors.Errorf("%s", respBody)
	default:
		var s string
		if respBody != nil {
			s = string(respBody)
		}
		return nil, errors.Errorf("HTTP status %d: content %q", resp.StatusCode, s)
	}

	return respBody, nil
}

// request makes a HTTP request to the given API endpoint, returning the raw
// *http.Response, or an error if one occurred. The caller is responsible for
// closing the response body.
func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader, authType int, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP request creation failed")
	}
	req.WithContext(ctx)

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if authType&AuthKeyEmail != 0 {
		req.Header.Set("X-Auth-Key", api.APIKey)
		req.Header.Set("X-Auth-Email", api.APIEmail)
	}
	if authType&AuthUserService != 0 {
		req.Header.Set("X-Auth-User-Service-Key", api.APIUserServiceKey)
	}
	if authType&AuthToken != 0 {
		req.Header.Set("Authorization", "Bearer "+api.APIToken)
	}

	if api.UserAgent != "" {
		req.Header.Set("User-Agent", api.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP request failed")
	}

	return resp, nil
}

// Returns the base URL to use for API endpoints that exist for accounts.
// If an account option was used when creating the API instance, returns
// the account URL.
//
// accountBase is the base URL for endpoints referring to the current user.
// It exists as a parameter because it is not consistent across APIs.
func (api *API) userBaseURL(accountBase string) string {
	if api.AccountID != "" {
		return "/accounts/" + api.AccountID
	}
	return accountBase
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}

// ResponseInfo contains a code and message returned by the API as errors or
// informational messages inside the response.
type ResponseInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response is a template.  There will also be a result struct.  There will be a
// unique response type for each response, which will include this type.
type Response struct {
	Success  bool           `json:"success"`
	Errors   []ResponseInfo `json:"errors"`
	Messages []ResponseInfo `json:"messages"`
}

// ResultInfoCursors contains information about cursors.
type ResultInfoCursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// ResultInfo contains metadata about the Response.
type ResultInfo struct {
	Page       int               `json:"page"`
	PerPage    int               `json:"per_page"`
	TotalPages int               `json:"total_pages"`
	Count      int               `json:"count"`
	Total      int               `json:"total_count"`
	Cursor     string            `json:"cursor"`
	Cursors    ResultInfoCursors `json:"cursors"`
}

// RawResponse keeps the result as JSON form
type RawResponse struct {
	Response
	Result json.RawMessage `json:"result"`
}

// Raw makes a HTTP request with user provided params and returns the
// result as untouched JSON.
func (api *API) Raw(method, endpoint string, data interface{}) (json.RawMessage, error) {
	res, err := api.makeRequest(method, endpoint, data)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var r RawResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// PaginationOptions can be passed to a list request to configure paging
// These values will be defaulted if omitted, and PerPage has min/max limits set by resource
type PaginationOptions struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// RetryPolicy specifies number of retries and min/max retry delays
// This config is used when the client exponentially backs off after errored requests
type RetryPolicy struct {
	MaxRetries    int
	MinRetryDelay time.Duration
	MaxRetryDelay time.Duration
}

// Logger defines the interface this library needs to use logging
// This is a subset of the methods implemented in the log package
type Logger interface {
	Printf(format string, v ...interface{})
}

// ReqOption is a functional option for configuring API requests
type ReqOption func(opt *reqOption)
type reqOption struct {
	params url.Values
}

// WithZoneFilters applies a filter based on zone properties.
func WithZoneFilters(zoneName, accountID, status string) ReqOption {
	return func(opt *reqOption) {
		if zoneName != "" {
			opt.params.Set("name", normalizeZoneName(zoneName))
		}

		if accountID != "" {
			opt.params.Set("account.id", accountID)
		}

		if status != "" {
			opt.params.Set("status", status)
		}
	}
}

// WithPagination configures the pagination for a response.
func WithPagination(opts PaginationOptions) ReqOption {
	return func(opt *reqOption) {
		opt.params.Set("page", strconv.Itoa(opts.Page))
		opt.params.Set("per_page", strconv.Itoa(opts.PerPage))
	}
}

// errorFromResponse returns a formatted error from the status code and error messages
// from the response body.
func errorFromResponse(statusCode int, respBody []byte) error {
	var r Response
	err := json.Unmarshal(respBody, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}

	errMsgs := []string{}
	for _, v := range r.Errors {
		errMsgs = append(errMsgs, v.Message)
	}
	return errors.Errorf("HTTP status %d: %s", statusCode, strings.Join(errMsgs, " "))
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	res, err := api.ListZonesContext(context.TODO(), WithZoneFilter(zoneName))
||||||| parent of 5ce8c7613 (update vendored files)
	res, err := api.ListZonesContext(context.TODO(), WithZoneFilter(zoneName))
=======
	zoneName = normalizeZoneName(zoneName)
	res, err := api.ListZonesContext(context.TODO(), WithZoneFilters(zoneName, "", ""))
>>>>>>> 5ce8c7613 (update vendored files)
	if err != nil {
		return "", errors.Wrap(err, "ListZonesContext command failed")
	}

	if len(res.Result) > 1 && api.AccountID == "" {
		return "", errors.New("ambiguous zone name used without an account ID")
	}

	for _, zone := range res.Result {
		if api.AccountID != "" {
			if zone.Name == zoneName && api.AccountID == zone.Account.ID {
				return zone.ID, nil
			}
		} else {
			if zone.Name == zoneName {
				return zone.ID, nil
			}
		}
	}

	return "", errors.New("Zone could not be found")
}

// makeRequest makes a HTTP request and returns the body as a byte slice,
// closing it before returning. params will be serialized to JSON.
func (api *API) makeRequest(method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(context.TODO(), method, uri, params, api.authType)
}

func (api *API) makeRequestContext(ctx context.Context, method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(ctx, method, uri, params, api.authType)
}

func (api *API) makeRequestWithHeaders(method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(context.TODO(), method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithAuthType(ctx context.Context, method, uri string, params interface{}, authType int) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, authType, nil)
}

func (api *API) makeRequestWithAuthTypeAndHeaders(ctx context.Context, method, uri string, params interface{}, authType int, headers http.Header) ([]byte, error) {
	// Replace nil with a JSON object if needed
	var jsonBody []byte
	var err error

	if params != nil {
		if paramBytes, ok := params.([]byte); ok {
			jsonBody = paramBytes
		} else {
			jsonBody, err = json.Marshal(params)
			if err != nil {
				return nil, errors.Wrap(err, "error marshalling params to JSON")
			}
		}
	} else {
		jsonBody = nil
	}

	var resp *http.Response
	var respErr error
	var reqBody io.Reader
	var respBody []byte
	for i := 0; i <= api.retryPolicy.MaxRetries; i++ {
		if jsonBody != nil {
			reqBody = bytes.NewReader(jsonBody)
		}
		if i > 0 {
			// expect the backoff introduced here on errored requests to dominate the effect of rate limiting
			// don't need a random component here as the rate limiter should do something similar
			// nb time duration could truncate an arbitrary float. Since our inputs are all ints, we should be ok
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(api.retryPolicy.MinRetryDelay))

			if sleepDuration > api.retryPolicy.MaxRetryDelay {
				sleepDuration = api.retryPolicy.MaxRetryDelay
			}
			// useful to do some simple logging here, maybe introduce levels later
			api.logger.Printf("Sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)
			time.Sleep(sleepDuration)

		}
		err = api.rateLimiter.Wait(context.TODO())
		if err != nil {
			return nil, errors.Wrap(err, "Error caused by request rate limiting")
		}
		resp, respErr = api.request(ctx, method, uri, reqBody, authType, headers)

		// retry if the server is rate limiting us or if it failed
		// assumes server operations are rolled back on failure
		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			// if we got a valid http response, try to read body so we can reuse the connection
			// see https://golang.org/pkg/net/http/#Client.Do
			if respErr == nil {
				respBody, err = ioutil.ReadAll(resp.Body)
				resp.Body.Close()

				respErr = errors.Wrap(err, "could not read response body")

				api.logger.Printf("Request: %s %s got an error response %d: %s\n", method, uri, resp.StatusCode,
					strings.Replace(strings.Replace(string(respBody), "\n", "", -1), "\t", "", -1))
			} else {
				api.logger.Printf("Error performing request: %s %s : %s \n", method, uri, respErr.Error())
			}
			continue
		} else {
			respBody, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, errors.Wrap(err, "could not read response body")
			}
			break
		}
	}
	if respErr != nil {
		return nil, respErr
	}

	switch {
	case resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices:
	case resp.StatusCode == http.StatusUnauthorized:
		return nil, errorFromResponse(resp.StatusCode, respBody)
	case resp.StatusCode == http.StatusForbidden:
		return nil, errorFromResponse(resp.StatusCode, respBody)
	case resp.StatusCode == http.StatusServiceUnavailable,
		resp.StatusCode == http.StatusBadGateway,
		resp.StatusCode == http.StatusGatewayTimeout,
		resp.StatusCode == 522,
		resp.StatusCode == 523,
		resp.StatusCode == 524:
		return nil, errors.Errorf("HTTP status %d: service failure", resp.StatusCode)
	// This isn't a great solution due to the way the `default` case is
	// a catch all and that the `filters/validate-expr` returns a HTTP 400
	// yet the clients need to use the HTTP body as a JSON string.
	case resp.StatusCode == 400 && strings.HasSuffix(resp.Request.URL.Path, "/filters/validate-expr"):
		return nil, errors.Errorf("%s", respBody)
	default:
		var s string
		if respBody != nil {
			s = string(respBody)
		}
		return nil, errors.Errorf("HTTP status %d: content %q", resp.StatusCode, s)
	}

	return respBody, nil
}

// request makes a HTTP request to the given API endpoint, returning the raw
// *http.Response, or an error if one occurred. The caller is responsible for
// closing the response body.
func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader, authType int, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequest(method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP request creation failed")
	}
	req.WithContext(ctx)

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if authType&AuthKeyEmail != 0 {
		req.Header.Set("X-Auth-Key", api.APIKey)
		req.Header.Set("X-Auth-Email", api.APIEmail)
	}
	if authType&AuthUserService != 0 {
		req.Header.Set("X-Auth-User-Service-Key", api.APIUserServiceKey)
	}
	if authType&AuthToken != 0 {
		req.Header.Set("Authorization", "Bearer "+api.APIToken)
	}

	if api.UserAgent != "" {
		req.Header.Set("User-Agent", api.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP request failed")
	}

	return resp, nil
}

// Returns the base URL to use for API endpoints that exist for accounts.
// If an account option was used when creating the API instance, returns
// the account URL.
//
// accountBase is the base URL for endpoints referring to the current user.
// It exists as a parameter because it is not consistent across APIs.
func (api *API) userBaseURL(accountBase string) string {
	if api.AccountID != "" {
		return "/accounts/" + api.AccountID
	}
	return accountBase
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}

// ResponseInfo contains a code and message returned by the API as errors or
// informational messages inside the response.
type ResponseInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response is a template.  There will also be a result struct.  There will be a
// unique response type for each response, which will include this type.
type Response struct {
	Success  bool           `json:"success"`
	Errors   []ResponseInfo `json:"errors"`
	Messages []ResponseInfo `json:"messages"`
}

// ResultInfoCursors contains information about cursors.
type ResultInfoCursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// ResultInfo contains metadata about the Response.
type ResultInfo struct {
	Page       int               `json:"page"`
	PerPage    int               `json:"per_page"`
	TotalPages int               `json:"total_pages"`
	Count      int               `json:"count"`
	Total      int               `json:"total_count"`
	Cursor     string            `json:"cursor"`
	Cursors    ResultInfoCursors `json:"cursors"`
}

// RawResponse keeps the result as JSON form
type RawResponse struct {
	Response
	Result json.RawMessage `json:"result"`
}

// Raw makes a HTTP request with user provided params and returns the
// result as untouched JSON.
func (api *API) Raw(method, endpoint string, data interface{}) (json.RawMessage, error) {
	res, err := api.makeRequest(method, endpoint, data)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}

	var r RawResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// PaginationOptions can be passed to a list request to configure paging
// These values will be defaulted if omitted, and PerPage has min/max limits set by resource
type PaginationOptions struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// RetryPolicy specifies number of retries and min/max retry delays
// This config is used when the client exponentially backs off after errored requests
type RetryPolicy struct {
	MaxRetries    int
	MinRetryDelay time.Duration
	MaxRetryDelay time.Duration
}

// Logger defines the interface this library needs to use logging
// This is a subset of the methods implemented in the log package
type Logger interface {
	Printf(format string, v ...interface{})
}

// ReqOption is a functional option for configuring API requests
type ReqOption func(opt *reqOption)
type reqOption struct {
	params url.Values
}

// WithZoneFilters applies a filter based on zone properties.
func WithZoneFilters(zoneName, accountID, status string) ReqOption {
	return func(opt *reqOption) {
		if zoneName != "" {
			opt.params.Set("name", normalizeZoneName(zoneName))
		}

		if accountID != "" {
			opt.params.Set("account.id", accountID)
		}

		if status != "" {
			opt.params.Set("status", status)
		}
	}
}

// WithPagination configures the pagination for a response.
func WithPagination(opts PaginationOptions) ReqOption {
	return func(opt *reqOption) {
		opt.params.Set("page", strconv.Itoa(opts.Page))
		opt.params.Set("per_page", strconv.Itoa(opts.PerPage))
	}
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// errorFromResponse returns a formatted error from the status code and error messages
// from the response body.
func errorFromResponse(statusCode int, respBody []byte) error {
	var r Response
	err := json.Unmarshal(respBody, &r)
	if err != nil {
		return errors.Wrap(err, errUnmarshalError)
	}

	errMsgs := []string{}
	for _, v := range r.Errors {
		errMsgs = append(errMsgs, v.Message)
	}
	return errors.Errorf("HTTP status %d: %s", statusCode, strings.Join(errMsgs, " "))
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	res, err := api.ListZonesContext(context.TODO(), WithZoneFilter(zoneName))
||||||| parent of 6b7ce455e (update vendored files)
	res, err := api.ListZonesContext(context.TODO(), WithZoneFilter(zoneName))
=======
	zoneName = normalizeZoneName(zoneName)
	res, err := api.ListZonesContext(context.Background(), WithZoneFilters(zoneName, api.AccountID, ""))
>>>>>>> 6b7ce455e (update vendored files)
	if err != nil {
		return "", errors.Wrap(err, "ListZonesContext command failed")
	}

	switch len(res.Result) {
	case 0:
		return "", errors.New("zone could not be found")
	case 1:
		return res.Result[0].ID, nil
	default:
		return "", errors.New("ambiguous zone name; an account ID might help")
	}
}

// makeRequest makes a HTTP request and returns the body as a byte slice,
// closing it before returning. params will be serialized to JSON.
func (api *API) makeRequest(method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(context.Background(), method, uri, params, api.authType)
}

func (api *API) makeRequestContext(ctx context.Context, method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(ctx, method, uri, params, api.authType)
}

func (api *API) makeRequestContextWithHeaders(ctx context.Context, method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithHeaders(method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(context.Background(), method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithAuthType(ctx context.Context, method, uri string, params interface{}, authType int) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, authType, nil)
}

func (api *API) makeRequestWithAuthTypeAndHeaders(ctx context.Context, method, uri string, params interface{}, authType int, headers http.Header) ([]byte, error) {
	// Replace nil with a JSON object if needed
	var jsonBody []byte
	var err error

	if params != nil {
		if paramBytes, ok := params.([]byte); ok {
			jsonBody = paramBytes
		} else {
			jsonBody, err = json.Marshal(params)
			if err != nil {
				return nil, errors.Wrap(err, "error marshalling params to JSON")
			}
		}
	} else {
		jsonBody = nil
	}

	var resp *http.Response
	var respErr error
	var reqBody io.Reader
	var respBody []byte
	for i := 0; i <= api.retryPolicy.MaxRetries; i++ {
		if jsonBody != nil {
			reqBody = bytes.NewReader(jsonBody)
		}
		if i > 0 {
			// expect the backoff introduced here on errored requests to dominate the effect of rate limiting
			// don't need a random component here as the rate limiter should do something similar
			// nb time duration could truncate an arbitrary float. Since our inputs are all ints, we should be ok
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(api.retryPolicy.MinRetryDelay))

			if sleepDuration > api.retryPolicy.MaxRetryDelay {
				sleepDuration = api.retryPolicy.MaxRetryDelay
			}
			// useful to do some simple logging here, maybe introduce levels later
			api.logger.Printf("Sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)

			select {
			case <-time.After(sleepDuration):
			case <-ctx.Done():
				return nil, errors.Wrap(ctx.Err(), "operation aborted during backoff")
			}

		}
		err = api.rateLimiter.Wait(context.Background())
		if err != nil {
			return nil, errors.Wrap(err, "Error caused by request rate limiting")
		}
		resp, respErr = api.request(ctx, method, uri, reqBody, authType, headers)

		// retry if the server is rate limiting us or if it failed
		// assumes server operations are rolled back on failure
		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			// if we got a valid http response, try to read body so we can reuse the connection
			// see https://golang.org/pkg/net/http/#Client.Do
			if respErr == nil {
				respBody, err = ioutil.ReadAll(resp.Body)
				resp.Body.Close()

				respErr = errors.Wrap(err, "could not read response body")

				api.logger.Printf("Request: %s %s got an error response %d: %s\n", method, uri, resp.StatusCode,
					strings.Replace(strings.Replace(string(respBody), "\n", "", -1), "\t", "", -1))
			} else {
				api.logger.Printf("Error performing request: %s %s : %s \n", method, uri, respErr.Error())
			}
			continue
		} else {
			respBody, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, errors.Wrap(err, "could not read response body")
			}
			break
		}
	}
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode >= http.StatusBadRequest {
		if strings.HasSuffix(resp.Request.URL.Path, "/filters/validate-expr") {
			return nil, errors.Errorf("%s", respBody)
		}

		if resp.StatusCode > http.StatusInternalServerError {
			return nil, errors.Errorf("HTTP status %d: service failure", resp.StatusCode)
		}

		errBody := &Response{}
		err = json.Unmarshal(respBody, &errBody)
		if err != nil {
			return nil, errors.Wrap(err, errUnmarshalErrorBody)
		}

		return nil, &APIRequestError{
			StatusCode: resp.StatusCode,
			Errors:     errBody.Errors,
		}
	}

	return respBody, nil
}

// request makes a HTTP request to the given API endpoint, returning the raw
// *http.Response, or an error if one occurred. The caller is responsible for
// closing the response body.
func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader, authType int, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP request creation failed")
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if authType&AuthKeyEmail != 0 {
		req.Header.Set("X-Auth-Key", api.APIKey)
		req.Header.Set("X-Auth-Email", api.APIEmail)
	}
	if authType&AuthUserService != 0 {
		req.Header.Set("X-Auth-User-Service-Key", api.APIUserServiceKey)
	}
	if authType&AuthToken != 0 {
		req.Header.Set("Authorization", "Bearer "+api.APIToken)
	}

	if api.UserAgent != "" {
		req.Header.Set("User-Agent", api.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "HTTP request failed")
	}

	return resp, nil
}

// Returns the base URL to use for API endpoints that exist for accounts.
// If an account option was used when creating the API instance, returns
// the account URL.
//
// accountBase is the base URL for endpoints referring to the current user.
// It exists as a parameter because it is not consistent across APIs.
func (api *API) userBaseURL(accountBase string) string {
	if api.AccountID != "" {
		return "/accounts/" + api.AccountID
	}
	return accountBase
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}

// ResponseInfo contains a code and message returned by the API as errors or
// informational messages inside the response.
type ResponseInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response is a template.  There will also be a result struct.  There will be a
// unique response type for each response, which will include this type.
type Response struct {
	Success  bool           `json:"success"`
	Errors   []ResponseInfo `json:"errors"`
	Messages []ResponseInfo `json:"messages"`
}

// ResultInfoCursors contains information about cursors.
type ResultInfoCursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

// ResultInfo contains metadata about the Response.
type ResultInfo struct {
	Page       int               `json:"page"`
	PerPage    int               `json:"per_page"`
	TotalPages int               `json:"total_pages"`
	Count      int               `json:"count"`
	Total      int               `json:"total_count"`
	Cursor     string            `json:"cursor"`
	Cursors    ResultInfoCursors `json:"cursors"`
}

// RawResponse keeps the result as JSON form
type RawResponse struct {
	Response
	Result json.RawMessage `json:"result"`
}

// Raw makes a HTTP request with user provided params and returns the
// result as untouched JSON.
func (api *API) Raw(method, endpoint string, data interface{}) (json.RawMessage, error) {
	res, err := api.makeRequest(method, endpoint, data)
	if err != nil {
		return nil, err
	}

	var r RawResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// PaginationOptions can be passed to a list request to configure paging
// These values will be defaulted if omitted, and PerPage has min/max limits set by resource
type PaginationOptions struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

// RetryPolicy specifies number of retries and min/max retry delays
// This config is used when the client exponentially backs off after errored requests
type RetryPolicy struct {
	MaxRetries    int
	MinRetryDelay time.Duration
	MaxRetryDelay time.Duration
}

// Logger defines the interface this library needs to use logging
// This is a subset of the methods implemented in the log package
type Logger interface {
	Printf(format string, v ...interface{})
}

// ReqOption is a functional option for configuring API requests
type ReqOption func(opt *reqOption)
type reqOption struct {
	params url.Values
}

// WithZoneFilters applies a filter based on zone properties.
func WithZoneFilters(zoneName, accountID, status string) ReqOption {
	return func(opt *reqOption) {
		if zoneName != "" {
			opt.params.Set("name", normalizeZoneName(zoneName))
		}

		if accountID != "" {
			opt.params.Set("account.id", accountID)
		}

		if status != "" {
			opt.params.Set("status", status)
		}
	}
}

// WithPagination configures the pagination for a response.
func WithPagination(opts PaginationOptions) ReqOption {
	return func(opt *reqOption) {
		if opts.Page > 0 {
			opt.params.Set("page", strconv.Itoa(opts.Page))
		}

		if opts.PerPage > 0 {
			opt.params.Set("per_page", strconv.Itoa(opts.PerPage))
		}
	}
}

// checkResultInfo checks whether ResultInfo is reasonable except that it currently
// ignores the cursor information. perPage, page, and count are the requested #items
// per page, the requested page number, and the actual length of the Result array.
//
// Responses from the actual Cloudflare servers should pass all these checks (or we
// discover a serious bug in the Cloudflare servers). However, the unit tests can
// easily violate these constraints and this utility function can help debugging.
// Correct pagination information is crucial for more advanced List* functions that
// handle pagination automatically and fetch different pages in parallel.
//
// TODO: check cursors as well.
func checkResultInfo(perPage, page, count int, info *ResultInfo) bool {
	if info.Cursor != "" || info.Cursors.Before != "" || info.Cursors.After != "" {
		panic("checkResultInfo could not handle cursors yet.")
	}

	switch {
	case info.PerPage != perPage || info.Page != page || info.Count != count:
		return false

	case info.PerPage <= 0:
		return false

	case info.Total == 0 && info.TotalPages == 0 && info.Page == 1 && info.Count == 0:
		return true

	case info.Total <= 0 || info.TotalPages <= 0:
		return false

	case info.Total > info.PerPage*info.TotalPages || info.Total <= info.PerPage*(info.TotalPages-1):
		return false
	}

	switch {
	case info.Page > info.TotalPages || info.Page <= 0:
		return false

	case info.Page < info.TotalPages:
		return info.Count == info.PerPage

	case info.Page == info.TotalPages:
		return info.Count == info.Total-info.PerPage*(info.TotalPages-1)

	default:
		// This is actually impossible, but Go compiler does not know trichotomy
		panic("checkResultInfo: impossible")
	}
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	originCARootCertEccURL = "https://developers.cloudflare.com/ssl/0d2cd0f374da0fb6dbf53128b60bbbf7/origin_ca_ecc_root.pem"
	originCARootCertRsaURL = "https://developers.cloudflare.com/ssl/e2b9968022bf23b071d95229b5678452/origin_ca_rsa_root.pem"
)

const (
>>>>>>> 4d7e5ad26 (update vendored files)
	// AuthKeyEmail specifies that we should authenticate with API key and email address
	AuthKeyEmail = 1 << iota
	// AuthUserService specifies that we should authenticate with a User-Service key
	AuthUserService
	// AuthToken specifies that we should authenticate with an API Token
	AuthToken
)

// API holds the configuration for the current API client. A client should not
// be modified concurrently.
type API struct {
	APIKey            string
	APIEmail          string
	APIUserServiceKey string
	APIToken          string
	BaseURL           string
	AccountID         string
	UserAgent         string
	headers           http.Header
	httpClient        *http.Client
	authType          int
	rateLimiter       *rate.Limiter
	retryPolicy       RetryPolicy
	logger            Logger
}

// newClient provides shared logic for New and NewWithUserServiceKey
func newClient(opts ...Option) (*API, error) {
	silentLogger := log.New(ioutil.Discard, "", log.LstdFlags)

	api := &API{
		BaseURL:     apiURL,
		headers:     make(http.Header),
		rateLimiter: rate.NewLimiter(rate.Limit(4), 1), // 4rps equates to default api limit (1200 req/5 min)
		retryPolicy: RetryPolicy{
			MaxRetries:    3,
			MinRetryDelay: time.Duration(1) * time.Second,
			MaxRetryDelay: time.Duration(30) * time.Second,
		},
		logger: silentLogger,
	}

	err := api.parseOptions(opts...)
	if err != nil {
		return nil, errors.Wrap(err, "options parsing failed")
	}

	// Fall back to http.DefaultClient if the package user does not provide
	// their own.
	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	}

	return api, nil
}

// New creates a new Cloudflare v4 API client.
func New(key, email string, opts ...Option) (*API, error) {
	if key == "" || email == "" {
		return nil, errors.New(errEmptyCredentials)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIKey = key
	api.APIEmail = email
	api.authType = AuthKeyEmail

	return api, nil
}

// NewWithAPIToken creates a new Cloudflare v4 API client using API Tokens
func NewWithAPIToken(token string, opts ...Option) (*API, error) {
	if token == "" {
		return nil, errors.New(errEmptyAPIToken)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIToken = token
	api.authType = AuthToken

	return api, nil
}

// NewWithUserServiceKey creates a new Cloudflare v4 API client using service key authentication.
func NewWithUserServiceKey(key string, opts ...Option) (*API, error) {
	if key == "" {
		return nil, errors.New(errEmptyCredentials)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIUserServiceKey = key
	api.authType = AuthUserService

	return api, nil
}

// SetAuthType sets the authentication method (AuthKeyEmail, AuthToken, or AuthUserService).
func (api *API) SetAuthType(authType int) {
	api.authType = authType
}

// ZoneIDByName retrieves a zone's ID from the name.
func (api *API) ZoneIDByName(zoneName string) (string, error) {
	zoneName = normalizeZoneName(zoneName)
	res, err := api.ListZonesContext(context.Background(), WithZoneFilters(zoneName, api.AccountID, ""))
	if err != nil {
		return "", fmt.Errorf("ListZonesContext command failed: %w", err)
	}

	switch len(res.Result) {
	case 0:
		return "", errors.New("zone could not be found")
	case 1:
		return res.Result[0].ID, nil
	default:
		return "", errors.New("ambiguous zone name; an account ID might help")
	}
}

// makeRequest makes a HTTP request and returns the body as a byte slice,
// closing it before returning. params will be serialized to JSON.
func (api *API) makeRequest(method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(context.Background(), method, uri, params, api.authType)
}

func (api *API) makeRequestContext(ctx context.Context, method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(ctx, method, uri, params, api.authType)
}

func (api *API) makeRequestContextWithHeaders(ctx context.Context, method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithAuthType(ctx context.Context, method, uri string, params interface{}, authType int) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, authType, nil)
}

// APIResponse holds the structure for a response from the API. It looks alot
// like `http.Response` however, uses a `[]byte` for the `Body` instead of a
// `io.ReadCloser`.
//
// This may go away in the experimental client in favour of `http.Response`.
type APIResponse struct {
	Body       []byte
	Status     string
	StatusCode int
	Headers    http.Header
}

func (api *API) makeRequestWithAuthTypeAndHeaders(ctx context.Context, method, uri string, params interface{}, authType int, headers http.Header) ([]byte, error) {
	res, err := api.makeRequestWithAuthTypeAndHeadersComplete(ctx, method, uri, params, authType, headers)
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

// Use this method if an API response can have different Content-Type headers and different body formats.
func (api *API) makeRequestContextWithHeadersComplete(ctx context.Context, method, uri string, params interface{}, headers http.Header) (*APIResponse, error) {
	return api.makeRequestWithAuthTypeAndHeadersComplete(ctx, method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithAuthTypeAndHeadersComplete(ctx context.Context, method, uri string, params interface{}, authType int, headers http.Header) (*APIResponse, error) {
	var err error
	var resp *http.Response
	var respErr error
	var respBody []byte

	for i := 0; i <= api.retryPolicy.MaxRetries; i++ {
		var reqBody io.Reader
		if params != nil {
			if r, ok := params.(io.Reader); ok {
				reqBody = r
			} else if paramBytes, ok := params.([]byte); ok {
				reqBody = bytes.NewReader(paramBytes)
			} else {
				var jsonBody []byte
				jsonBody, err = json.Marshal(params)
				if err != nil {
					return nil, fmt.Errorf("error marshalling params to JSON: %w", err)
				}
				reqBody = bytes.NewReader(jsonBody)
			}
		}

		if i > 0 {
			// expect the backoff introduced here on errored requests to dominate the effect of rate limiting
			// don't need a random component here as the rate limiter should do something similar
			// nb time duration could truncate an arbitrary float. Since our inputs are all ints, we should be ok
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(api.retryPolicy.MinRetryDelay))

			if sleepDuration > api.retryPolicy.MaxRetryDelay {
				sleepDuration = api.retryPolicy.MaxRetryDelay
			}
			// useful to do some simple logging here, maybe introduce levels later
			api.logger.Printf("Sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)

			select {
			case <-time.After(sleepDuration):
			case <-ctx.Done():
				return nil, fmt.Errorf("operation aborted during backoff: %w", ctx.Err())
			}
		}

		err = api.rateLimiter.Wait(ctx)
		if err != nil {
			return nil, fmt.Errorf("error caused by request rate limiting: %w", err)
		}

		if api.Debug {
			if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
				buf := &bytes.Buffer{}
				tee := io.TeeReader(reqBody, buf)
				debugBody, _ := ioutil.ReadAll(tee)
				payloadBody, _ := ioutil.ReadAll(buf)
				fmt.Printf("cloudflare-go [DEBUG] REQUEST Method:%v URI:%s Headers:%#v Body:%v\n", method, api.BaseURL+uri, headers, string(debugBody))
				// ensure we recreate the io.Reader for use
				reqBody = bytes.NewReader(payloadBody)
			} else {
				fmt.Printf("cloudflare-go [DEBUG] REQUEST Method:%v URI:%s Headers:%#v Body:%v\n", method, api.BaseURL+uri, headers, nil)
			}
		}

		resp, respErr = api.request(ctx, method, uri, reqBody, authType, headers)

		// short circuit processing on context timeouts
		if respErr != nil && errors.Is(respErr, context.DeadlineExceeded) {
			return nil, respErr
		}

		// retry if the server is rate limiting us or if it failed
		// assumes server operations are rolled back on failure
		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
				respErr = errors.New("exceeded available rate limit retries")
			}

			// if we got a valid http response, try to read body so we can reuse the connection
			// see https://golang.org/pkg/net/http/#Client.Do
			if respErr == nil {
				respBody, err = ioutil.ReadAll(resp.Body)
				resp.Body.Close()

				respErr = fmt.Errorf("could not read response body: %w", err)

				api.logger.Printf("Request: %s %s got an error response %d: %s\n", method, uri, resp.StatusCode,
					strings.Replace(strings.Replace(string(respBody), "\n", "", -1), "\t", "", -1))
			} else {
				api.logger.Printf("Error performing request: %s %s : %s \n", method, uri, respErr.Error())
			}
			continue
		} else {
			respBody, err = ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, fmt.Errorf("could not read response body: %w", err)
			}
			break
		}
	}

	// still had an error after all retries
	if respErr != nil {
		return nil, respErr
	}

	if api.Debug {
		fmt.Printf("cloudflare-go [DEBUG] RESPONSE StatusCode:%d RayID:%s ContentType:%s Body:%#v\n", resp.StatusCode, resp.Header.Get("cf-ray"), resp.Header.Get("content-type"), string(respBody))
	}

	if resp.StatusCode >= http.StatusBadRequest {
		if strings.HasSuffix(resp.Request.URL.Path, "/filters/validate-expr") {
			return nil, fmt.Errorf("%s", respBody)
		}

		if resp.StatusCode >= http.StatusInternalServerError {
			return nil, &ServiceError{cloudflareError: &Error{
				StatusCode: resp.StatusCode,
				RayID:      resp.Header.Get("cf-ray"),
				Errors: []ResponseInfo{{
					Message: errInternalServiceError,
				}},
			}}
		}

		errBody := &Response{}
		err = json.Unmarshal(respBody, &errBody)
		if err != nil {
			return nil, fmt.Errorf(errUnmarshalErrorBody+": %w", err)
		}

		errCodes := make([]int, 0, len(errBody.Errors))
		errMsgs := make([]string, 0, len(errBody.Errors))
		for _, e := range errBody.Errors {
			errCodes = append(errCodes, e.Code)
			errMsgs = append(errMsgs, e.Message)
		}

		err := &Error{
			StatusCode:    resp.StatusCode,
			RayID:         resp.Header.Get("cf-ray"),
			Errors:        errBody.Errors,
			ErrorCodes:    errCodes,
			ErrorMessages: errMsgs,
		}

		switch resp.StatusCode {
		case http.StatusUnauthorized:
			err.Type = ErrorTypeAuthorization
			return nil, &AuthorizationError{cloudflareError: err}
		case http.StatusForbidden:
			err.Type = ErrorTypeAuthentication
			return nil, &AuthenticationError{cloudflareError: err}
		case http.StatusNotFound:
			err.Type = ErrorTypeNotFound
			return nil, &NotFoundError{cloudflareError: err}
		case http.StatusTooManyRequests:
			err.Type = ErrorTypeRateLimit
			return nil, &RatelimitError{cloudflareError: err}
		default:
			err.Type = ErrorTypeRequest
			return nil, &RequestError{cloudflareError: err}
		}
	}

	return &APIResponse{
		Body:       respBody,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header,
	}, nil
}

// request makes a HTTP request to the given API endpoint, returning the raw
// *http.Response, or an error if one occurred. The caller is responsible for
// closing the response body.
func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader, authType int, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if authType&AuthKeyEmail != 0 {
		req.Header.Set("X-Auth-Key", api.APIKey)
		req.Header.Set("X-Auth-Email", api.APIEmail)
	}
	if authType&AuthUserService != 0 {
		req.Header.Set("X-Auth-User-Service-Key", api.APIUserServiceKey)
	}
	if authType&AuthToken != 0 {
		req.Header.Set("Authorization", "Bearer "+api.APIToken)
	}

	if api.UserAgent != "" {
		req.Header.Set("User-Agent", api.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	return resp, nil
}

// Returns the base URL to use for API endpoints that exist for accounts.
// If an account option was used when creating the API instance, returns
// the account URL.
//
// accountBase is the base URL for endpoints referring to the current user.
// It exists as a parameter because it is not consistent across APIs.
func (api *API) userBaseURL(accountBase string) string {
	if api.AccountID != "" {
		return "/accounts/" + api.AccountID
	}
	return accountBase
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}

// ResponseInfo contains a code and message returned by the API as errors or
// informational messages inside the response.
type ResponseInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response is a template.  There will also be a result struct.  There will be a
// unique response type for each response, which will include this type.
type Response struct {
	Success  bool           `json:"success"`
	Errors   []ResponseInfo `json:"errors"`
	Messages []ResponseInfo `json:"messages"`
}

// ResultInfoCursors contains information about cursors.
type ResultInfoCursors struct {
	Before string `json:"before" url:"before,omitempty"`
	After  string `json:"after" url:"after,omitempty"`
}

// ResultInfo contains metadata about the Response.
type ResultInfo struct {
	Page       int               `json:"page" url:"page,omitempty"`
	PerPage    int               `json:"per_page" url:"per_page,omitempty"`
	TotalPages int               `json:"total_pages" url:"-"`
	Count      int               `json:"count" url:"-"`
	Total      int               `json:"total_count" url:"-"`
	Cursor     string            `json:"cursor" url:"cursor,omitempty"`
	Cursors    ResultInfoCursors `json:"cursors" url:"cursors,omitempty"`
}

// RawResponse keeps the result as JSON form.
type RawResponse struct {
	Response
	Result json.RawMessage `json:"result"`
}

// Raw makes a HTTP request with user provided params and returns the
// result as untouched JSON.
func (api *API) Raw(ctx context.Context, method, endpoint string, data interface{}, headers http.Header) (json.RawMessage, error) {
	res, err := api.makeRequestContextWithHeaders(ctx, method, endpoint, data, headers)
	if err != nil {
		return nil, err
	}

	var r RawResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// PaginationOptions can be passed to a list request to configure paging
// These values will be defaulted if omitted, and PerPage has min/max limits set by resource.
type PaginationOptions struct {
	Page    int `json:"page,omitempty" url:"page,omitempty"`
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"`
}

// RetryPolicy specifies number of retries and min/max retry delays
// This config is used when the client exponentially backs off after errored requests.
type RetryPolicy struct {
	MaxRetries    int
	MinRetryDelay time.Duration
	MaxRetryDelay time.Duration
}

// Logger defines the interface this library needs to use logging
// This is a subset of the methods implemented in the log package.
type Logger interface {
	Printf(format string, v ...interface{})
}

// ReqOption is a functional option for configuring API requests.
type ReqOption func(opt *reqOption)

type reqOption struct {
	params url.Values
}

// WithZoneFilters applies a filter based on zone properties.
func WithZoneFilters(zoneName, accountID, status string) ReqOption {
	return func(opt *reqOption) {
		if zoneName != "" {
			opt.params.Set("name", normalizeZoneName(zoneName))
		}

		if accountID != "" {
			opt.params.Set("account.id", accountID)
		}

		if status != "" {
			opt.params.Set("status", status)
		}
	}
}

// WithPagination configures the pagination for a response.
func WithPagination(opts PaginationOptions) ReqOption {
	return func(opt *reqOption) {
		if opts.Page > 0 {
			opt.params.Set("page", strconv.Itoa(opts.Page))
		}

		if opts.PerPage > 0 {
			opt.params.Set("per_page", strconv.Itoa(opts.PerPage))
		}
	}
}

// checkResultInfo checks whether ResultInfo is reasonable except that it currently
// ignores the cursor information. perPage, page, and count are the requested #items
// per page, the requested page number, and the actual length of the Result array.
//
// Responses from the actual Cloudflare servers should pass all these checks (or we
// discover a serious bug in the Cloudflare servers). However, the unit tests can
// easily violate these constraints and this utility function can help debugging.
// Correct pagination information is crucial for more advanced List* functions that
// handle pagination automatically and fetch different pages in parallel.
//
// TODO: check cursors as well.
func checkResultInfo(perPage, page, count int, info *ResultInfo) bool {
	if info.Cursor != "" || info.Cursors.Before != "" || info.Cursors.After != "" {
		panic("checkResultInfo could not handle cursors yet.")
	}

	switch {
	case info.PerPage != perPage || info.Page != page || info.Count != count:
		return false

	case info.PerPage <= 0:
		return false

	case info.Total == 0 && info.TotalPages == 0 && info.Page == 1 && info.Count == 0:
		return true

	case info.Total <= 0 || info.TotalPages <= 0:
		return false

	case info.Total > info.PerPage*info.TotalPages || info.Total <= info.PerPage*(info.TotalPages-1):
		return false
	}

	switch {
	case info.Page > info.TotalPages || info.Page <= 0:
		return false

	case info.Page < info.TotalPages:
		return info.Count == info.PerPage

	case info.Page == info.TotalPages:
		return info.Count == info.Total-info.PerPage*(info.TotalPages-1)

	default:
		// This is actually impossible, but Go compiler does not know trichotomy
		panic("checkResultInfo: impossible")
	}
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"encoding/json"
=======
	"errors"
	"fmt"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"io"
	"log"
	"math"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/goccy/go-json"

	"golang.org/x/time/rate"
)

var (
	Version string = "v4"

	// Deprecated: Use `client.New` configuration instead.
	apiURL = fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath)
)

const (
	// AuthKeyEmail specifies that we should authenticate with API key and email address.
	AuthKeyEmail = 1 << iota
	// AuthUserService specifies that we should authenticate with a User-Service key.
	AuthUserService
	// AuthToken specifies that we should authenticate with an API Token.
	AuthToken
)

// API holds the configuration for the current API client. A client should not
// be modified concurrently.
type API struct {
	APIKey            string
	APIEmail          string
	APIUserServiceKey string
	APIToken          string
	BaseURL           string
	UserAgent         string
	headers           http.Header
	httpClient        *http.Client
	authType          int
	rateLimiter       *rate.Limiter
	retryPolicy       RetryPolicy
	logger            Logger
	Debug             bool
}

// newClient provides shared logic for New and NewWithUserServiceKey.
func newClient(opts ...Option) (*API, error) {
	silentLogger := log.New(io.Discard, "", log.LstdFlags)

	api := &API{
		BaseURL:     fmt.Sprintf("%s://%s%s", defaultScheme, defaultHostname, defaultBasePath),
		UserAgent:   userAgent + "/" + Version,
		headers:     make(http.Header),
		rateLimiter: rate.NewLimiter(rate.Limit(4), 1), // 4rps equates to default api limit (1200 req/5 min)
		retryPolicy: RetryPolicy{
			MaxRetries:    3,
			MinRetryDelay: 1 * time.Second,
			MaxRetryDelay: 30 * time.Second,
		},
		logger: silentLogger,
	}

	err := api.parseOptions(opts...)
	if err != nil {
		return nil, fmt.Errorf("options parsing failed: %w", err)
	}

	// Fall back to http.DefaultClient if the package user does not provide
	// their own.
	if api.httpClient == nil {
		api.httpClient = http.DefaultClient
	}

	return api, nil
}

// New creates a new Cloudflare v4 API client.
func New(key, email string, opts ...Option) (*API, error) {
	if key == "" || email == "" {
		return nil, errors.New(errEmptyCredentials)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIKey = key
	api.APIEmail = email
	api.authType = AuthKeyEmail

	return api, nil
}

// NewWithAPIToken creates a new Cloudflare v4 API client using API Tokens.
func NewWithAPIToken(token string, opts ...Option) (*API, error) {
	if token == "" {
		return nil, errors.New(errEmptyAPIToken)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIToken = token
	api.authType = AuthToken

	return api, nil
}

// NewWithUserServiceKey creates a new Cloudflare v4 API client using service key authentication.
func NewWithUserServiceKey(key string, opts ...Option) (*API, error) {
	if key == "" {
		return nil, errors.New(errEmptyCredentials)
	}

	api, err := newClient(opts...)
	if err != nil {
		return nil, err
	}

	api.APIUserServiceKey = key
	api.authType = AuthUserService

	return api, nil
}

// SetAuthType sets the authentication method (AuthKeyEmail, AuthToken, or AuthUserService).
func (api *API) SetAuthType(authType int) {
	api.authType = authType
}

// ZoneIDByName retrieves a zone's ID from the name.
func (api *API) ZoneIDByName(zoneName string) (string, error) {
	zoneName = normalizeZoneName(zoneName)
	res, err := api.ListZonesContext(context.Background(), WithZoneFilters(zoneName, "", ""))
	if err != nil {
		return "", fmt.Errorf("ListZonesContext command failed: %w", err)
	}

	switch len(res.Result) {
	case 0:
		return "", errors.New("zone could not be found")
	case 1:
		return res.Result[0].ID, nil
	default:
		return "", errors.New("ambiguous zone name; an account ID might help")
	}
}

// makeRequest makes a HTTP request and returns the body as a byte slice,
// closing it before returning. params will be serialized to JSON.
func (api *API) makeRequest(method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(context.Background(), method, uri, params, api.authType)
}

func (api *API) makeRequestContext(ctx context.Context, method, uri string, params interface{}) ([]byte, error) {
	return api.makeRequestWithAuthType(ctx, method, uri, params, api.authType)
}

func (api *API) makeRequestContextWithHeaders(ctx context.Context, method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithAuthType(ctx context.Context, method, uri string, params interface{}, authType int) ([]byte, error) {
	return api.makeRequestWithAuthTypeAndHeaders(ctx, method, uri, params, authType, nil)
}

// APIResponse holds the structure for a response from the API. It looks alot
// like `http.Response` however, uses a `[]byte` for the `Body` instead of a
// `io.ReadCloser`.
//
// This may go away in the experimental client in favour of `http.Response`.
type APIResponse struct {
	Body       []byte
	Status     string
	StatusCode int
	Headers    http.Header
}

func (api *API) makeRequestWithAuthTypeAndHeaders(ctx context.Context, method, uri string, params interface{}, authType int, headers http.Header) ([]byte, error) {
	res, err := api.makeRequestWithAuthTypeAndHeadersComplete(ctx, method, uri, params, authType, headers)
	if err != nil {
		return nil, err
	}
	return res.Body, err
}

// Use this method if an API response can have different Content-Type headers and different body formats.
func (api *API) makeRequestContextWithHeadersComplete(ctx context.Context, method, uri string, params interface{}, headers http.Header) (*APIResponse, error) {
	return api.makeRequestWithAuthTypeAndHeadersComplete(ctx, method, uri, params, api.authType, headers)
}

func (api *API) makeRequestWithAuthTypeAndHeadersComplete(ctx context.Context, method, uri string, params interface{}, authType int, headers http.Header) (*APIResponse, error) {
	var err error
	var resp *http.Response
	var respErr error
	var respBody []byte

	for i := 0; i <= api.retryPolicy.MaxRetries; i++ {
		var reqBody io.Reader
		if params != nil {
			if r, ok := params.(io.Reader); ok {
				reqBody = r
			} else if paramBytes, ok := params.([]byte); ok {
				reqBody = bytes.NewReader(paramBytes)
			} else {
				var jsonBody []byte
				jsonBody, err = json.Marshal(params)
				if err != nil {
					return nil, fmt.Errorf("error marshalling params to JSON: %w", err)
				}
				reqBody = bytes.NewReader(jsonBody)
			}
		}

		if i > 0 {
			// expect the backoff introduced here on errored requests to dominate the effect of rate limiting
			// don't need a random component here as the rate limiter should do something similar
			// nb time duration could truncate an arbitrary float. Since our inputs are all ints, we should be ok
			sleepDuration := time.Duration(math.Pow(2, float64(i-1)) * float64(api.retryPolicy.MinRetryDelay))

			if sleepDuration > api.retryPolicy.MaxRetryDelay {
				sleepDuration = api.retryPolicy.MaxRetryDelay
			}
			// useful to do some simple logging here, maybe introduce levels later
			api.logger.Printf("Sleeping %s before retry attempt number %d for request %s %s", sleepDuration.String(), i, method, uri)

			select {
			case <-time.After(sleepDuration):
			case <-ctx.Done():
				return nil, fmt.Errorf("operation aborted during backoff: %w", ctx.Err())
			}
		}

		err = api.rateLimiter.Wait(ctx)
		if err != nil {
			return nil, fmt.Errorf("error caused by request rate limiting: %w", err)
		}

		resp, respErr = api.request(ctx, method, uri, reqBody, authType, headers)

		// short circuit processing on context timeouts
		if respErr != nil && errors.Is(respErr, context.DeadlineExceeded) {
			return nil, respErr
		}

		// retry if the server is rate limiting us or if it failed
		// assumes server operations are rolled back on failure
		if respErr != nil || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			if resp != nil && resp.StatusCode == http.StatusTooManyRequests {
				respErr = errors.New("exceeded available rate limit retries")
			}

			if respErr == nil {
				respErr = fmt.Errorf("received %s response (HTTP %d), please try again later", strings.ToLower(http.StatusText(resp.StatusCode)), resp.StatusCode)
			}
			continue
		} else {
			respBody, err = io.ReadAll(resp.Body)
			defer resp.Body.Close()
			if err != nil {
				return nil, fmt.Errorf("could not read response body: %w", err)
			}

			break
		}
	}

	// still had an error after all retries
	if respErr != nil {
		return nil, respErr
	}

	if resp.StatusCode >= http.StatusBadRequest {
		if strings.HasSuffix(resp.Request.URL.Path, "/filters/validate-expr") {
			return nil, fmt.Errorf("%s", respBody)
		}

		if resp.StatusCode >= http.StatusInternalServerError {
			return nil, &ServiceError{cloudflareError: &Error{
				StatusCode: resp.StatusCode,
				RayID:      resp.Header.Get("cf-ray"),
				Errors: []ResponseInfo{{
					Message: errInternalServiceError,
				}},
			}}
		}

		errBody := &Response{}
		err = json.Unmarshal(respBody, &errBody)
		if err != nil {
			return nil, fmt.Errorf(errUnmarshalErrorBody+": %w", err)
		}

		errCodes := make([]int, 0, len(errBody.Errors))
		errMsgs := make([]string, 0, len(errBody.Errors))
		for _, e := range errBody.Errors {
			errCodes = append(errCodes, e.Code)
			errMsgs = append(errMsgs, e.Message)
		}

		err := &Error{
			StatusCode:    resp.StatusCode,
			RayID:         resp.Header.Get("cf-ray"),
			Errors:        errBody.Errors,
			ErrorCodes:    errCodes,
			ErrorMessages: errMsgs,
			Messages:      errBody.Messages,
		}

		switch resp.StatusCode {
		case http.StatusUnauthorized:
			err.Type = ErrorTypeAuthorization
			return nil, &AuthorizationError{cloudflareError: err}
		case http.StatusForbidden:
			err.Type = ErrorTypeAuthentication
			return nil, &AuthenticationError{cloudflareError: err}
		case http.StatusNotFound:
			err.Type = ErrorTypeNotFound
			return nil, &NotFoundError{cloudflareError: err}
		case http.StatusTooManyRequests:
			err.Type = ErrorTypeRateLimit
			return nil, &RatelimitError{cloudflareError: err}
		default:
			err.Type = ErrorTypeRequest
			return nil, &RequestError{cloudflareError: err}
		}
	}

	return &APIResponse{
		Body:       respBody,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Headers:    resp.Header,
	}, nil
}

// request makes a HTTP request to the given API endpoint, returning the raw
// *http.Response, or an error if one occurred. The caller is responsible for
// closing the response body.
func (api *API) request(ctx context.Context, method, uri string, reqBody io.Reader, authType int, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, api.BaseURL+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, api.headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if authType&AuthKeyEmail != 0 {
		req.Header.Set("X-Auth-Key", api.APIKey)
		req.Header.Set("X-Auth-Email", api.APIEmail)
	}
	if authType&AuthUserService != 0 {
		req.Header.Set("X-Auth-User-Service-Key", api.APIUserServiceKey)
	}
	if authType&AuthToken != 0 {
		req.Header.Set("Authorization", "Bearer "+api.APIToken)
	}

	if api.UserAgent != "" {
		req.Header.Set("User-Agent", api.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	if api.Debug {
		dump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}

		// Strip out any sensitive information from the request payload.
		sensitiveKeys := []string{api.APIKey, api.APIEmail, api.APIToken, api.APIUserServiceKey}
		for _, key := range sensitiveKeys {
			if key != "" {
				valueRegex := regexp.MustCompile(fmt.Sprintf("(?m)%s", key))
				dump = valueRegex.ReplaceAll(dump, []byte("[redacted]"))
			}
		}
		log.Printf("\n%s", string(dump))
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	if api.Debug {
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return resp, err
		}
		log.Printf("\n%s", string(dump))
	}

	return resp, nil
}

// copyHeader copies all headers for `source` and sets them on `target`.
// based on https://godoc.org/github.com/golang/gddo/httputil/header#Copy
func copyHeader(target, source http.Header) {
	for k, vs := range source {
		target[k] = vs
	}
}

// ResponseInfo contains a code and message returned by the API as errors or
// informational messages inside the response.
type ResponseInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Response is a template.  There will also be a result struct.  There will be a
// unique response type for each response, which will include this type.
type Response struct {
	Success  bool           `json:"success"`
	Errors   []ResponseInfo `json:"errors"`
	Messages []ResponseInfo `json:"messages"`
}

// ResultInfoCursors contains information about cursors.
type ResultInfoCursors struct {
	Before string `json:"before" url:"before,omitempty"`
	After  string `json:"after" url:"after,omitempty"`
}

// ResultInfo contains metadata about the Response.
type ResultInfo struct {
	Page       int               `json:"page" url:"page,omitempty"`
	PerPage    int               `json:"per_page" url:"per_page,omitempty"`
	TotalPages int               `json:"total_pages" url:"-"`
	Count      int               `json:"count" url:"-"`
	Total      int               `json:"total_count" url:"-"`
	Cursor     string            `json:"cursor" url:"cursor,omitempty"`
	Cursors    ResultInfoCursors `json:"cursors" url:"cursors,omitempty"`
}

// RawResponse keeps the result as JSON form.
type RawResponse struct {
	Response
	Result     json.RawMessage `json:"result"`
	ResultInfo *ResultInfo     `json:"result_info,omitempty"`
}

// Raw makes a HTTP request with user provided params and returns the
// result as a RawResponse, which contains the untouched JSON result.
func (api *API) Raw(ctx context.Context, method, endpoint string, data interface{}, headers http.Header) (RawResponse, error) {
	var r RawResponse
	res, err := api.makeRequestContextWithHeaders(ctx, method, endpoint, data, headers)
	if err != nil {
		return r, err
	}

	if err := json.Unmarshal(res, &r); err != nil {
		return r, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// PaginationOptions can be passed to a list request to configure paging
// These values will be defaulted if omitted, and PerPage has min/max limits set by resource.
type PaginationOptions struct {
	Page    int `json:"page,omitempty" url:"page,omitempty"`
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"`
}

// RetryPolicy specifies number of retries and min/max retry delays
// This config is used when the client exponentially backs off after errored requests.
type RetryPolicy struct {
	MaxRetries    int
	MinRetryDelay time.Duration
	MaxRetryDelay time.Duration
}

// Logger defines the interface this library needs to use logging
// This is a subset of the methods implemented in the log package.
type Logger interface {
	Printf(format string, v ...interface{})
}

// ReqOption is a functional option for configuring API requests.
type ReqOption func(opt *reqOption)

type reqOption struct {
	params url.Values
}

// WithZoneFilters applies a filter based on zone properties.
func WithZoneFilters(zoneName, accountID, status string) ReqOption {
	return func(opt *reqOption) {
		if zoneName != "" {
			opt.params.Set("name", normalizeZoneName(zoneName))
		}

		if accountID != "" {
			opt.params.Set("account.id", accountID)
		}

		if status != "" {
			opt.params.Set("status", status)
		}
	}
}

// WithPagination configures the pagination for a response.
func WithPagination(opts PaginationOptions) ReqOption {
	return func(opt *reqOption) {
		if opts.Page > 0 {
			opt.params.Set("page", strconv.Itoa(opts.Page))
		}

		if opts.PerPage > 0 {
			opt.params.Set("per_page", strconv.Itoa(opts.PerPage))
		}
	}
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// checkResultInfo checks whether ResultInfo is reasonable except that it currently
// ignores the cursor information. perPage, page, and count are the requested #items
// per page, the requested page number, and the actual length of the Result array.
//
// Responses from the actual Cloudflare servers should pass all these checks (or we
// discover a serious bug in the Cloudflare servers). However, the unit tests can
// easily violate these constraints and this utility function can help debugging.
// Correct pagination information is crucial for more advanced List* functions that
// handle pagination automatically and fetch different pages in parallel.
//
// TODO: check cursors as well.
func checkResultInfo(perPage, page, count int, info *ResultInfo) bool {
	if info.Cursor != "" || info.Cursors.Before != "" || info.Cursors.After != "" {
		panic("checkResultInfo could not handle cursors yet.")
	}

	switch {
	case info.PerPage != perPage || info.Page != page || info.Count != count:
		return false

	case info.PerPage <= 0:
		return false

	case info.Total == 0 && info.TotalPages == 0 && info.Page == 1 && info.Count == 0:
		return true

	case info.Total <= 0 || info.TotalPages <= 0:
		return false

	case info.Total > info.PerPage*info.TotalPages || info.Total <= info.PerPage*(info.TotalPages-1):
		return false
	}

	switch {
	case info.Page > info.TotalPages || info.Page <= 0:
		return false

	case info.Page < info.TotalPages:
		return info.Count == info.PerPage

	case info.Page == info.TotalPages:
		return info.Count == info.Total-info.PerPage*(info.TotalPages-1)

	default:
		// This is actually impossible, but Go compiler does not know trichotomy
		panic("checkResultInfo: impossible")
	}
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "asc"
	OrderDirectionDesc OrderDirection = "desc"
)
