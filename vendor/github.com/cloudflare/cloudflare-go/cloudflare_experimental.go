package cloudflare

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type service struct {
	client *Client
}

type ClientParams struct {
	Key            string
	Email          string
	UserServiceKey string
	Token          string
	STS            *SecurityTokenConfiguration
	BaseURL        *url.URL
	UserAgent      string
	Headers        http.Header
	HTTPClient     *http.Client
	RetryPolicy    RetryPolicy
	Logger         LeveledLoggerInterface
	Debug          bool
}

// A Client manages communication with the Cloudflare API.
type Client struct {
	clientMu sync.Mutex

	*ClientParams

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Zones *ZonesService
}

// Client returns the http.Client used by this Cloudflare client.
func (c *Client) Client() *http.Client {
	c.clientMu.Lock()
	defer c.clientMu.Unlock()
	clientCopy := *c.HTTPClient
	return &clientCopy
}

// Call is the entrypoint to making API calls with the correct request setup.
func (c *Client) Call(ctx context.Context, method, path string, payload interface{}) ([]byte, error) {
	return c.makeRequest(ctx, method, path, payload, nil)
}

// CallWithHeaders is the entrypoint to making API calls with the correct
// request setup (like `Call`) but allows passing in additional HTTP headers
// with the request.
func (c *Client) CallWithHeaders(ctx context.Context, method, path string, payload interface{}, headers http.Header) ([]byte, error) {
	return c.makeRequest(ctx, method, path, payload, headers)
}

// New creates a new instance of the API client by merging ClientParams with the
// default values.
func NewExperimental(config *ClientParams) (*Client, error) {
	c := &Client{ClientParams: &ClientParams{}}
	c.common.client = c

	defaultURL, _ := url.Parse(defaultScheme + "://" + defaultHostname + defaultBasePath)
	if config.BaseURL != nil {
		c.ClientParams.BaseURL = config.BaseURL
	} else {
		c.ClientParams.BaseURL = defaultURL
	}

	if config.UserAgent != "" {
		c.ClientParams.UserAgent = config.UserAgent
	} else {
		c.ClientParams.UserAgent = userAgent + "/" + Version
	}

	if config.HTTPClient != nil {
		c.ClientParams.HTTPClient = config.HTTPClient
	} else {
		retryClient := retryablehttp.NewClient()

		if c.RetryPolicy.MaxRetries > 0 {
			retryClient.RetryMax = c.RetryPolicy.MaxRetries
		} else {
			retryClient.RetryMax = 4
		}

		if c.RetryPolicy.MinRetryDelay > 0 {
			retryClient.RetryWaitMin = c.RetryPolicy.MinRetryDelay
		} else {
			retryClient.RetryWaitMin = time.Duration(1) * time.Second
		}

		if c.RetryPolicy.MaxRetryDelay > 0 {
			retryClient.RetryWaitMax = c.RetryPolicy.MaxRetryDelay
		} else {
			retryClient.RetryWaitMax = time.Duration(30) * time.Second
		}

		retryClient.Logger = silentRetryLogger
		c.ClientParams.HTTPClient = retryClient.StandardClient()
	}

	if config.Headers != nil {
		c.ClientParams.Headers = config.Headers
	} else {
		c.ClientParams.Headers = make(http.Header)
	}

	if config.Key != "" && config.Token != "" {
		return nil, ErrAPIKeysAndTokensAreMutuallyExclusive
	}

	if config.Key != "" {
		c.ClientParams.Key = config.Key
		c.ClientParams.Email = config.Email
	}

	if config.Token != "" {
		c.ClientParams.Token = config.Token
	}

	if config.UserServiceKey != "" {
		c.ClientParams.UserServiceKey = config.UserServiceKey
	}

	c.ClientParams.Debug = config.Debug
	if c.ClientParams.Debug {
		c.ClientParams.Logger = &LeveledLogger{Level: 4}
	} else {
		c.ClientParams.Logger = SilentLeveledLogger
	}

	if config.STS != nil {
		stsToken, err := fetchSTSCredentials(config.STS)
		if err != nil {
			return nil, ErrSTSFailure
		}
		c.ClientParams.Token = stsToken
	}

	c.Zones = (*ZonesService)(&c.common)

	return c, nil
}

// request makes a HTTP request to the given API endpoint, returning the raw
// *http.Response, or an error if one occurred. The caller is responsible for
// closing the response body.
func (c *Client) request(ctx context.Context, method, uri string, reqBody io.Reader, headers http.Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL.String()+uri, reqBody)
	if err != nil {
		return nil, fmt.Errorf("HTTP request creation failed: %w", err)
	}

	combinedHeaders := make(http.Header)
	copyHeader(combinedHeaders, c.Headers)
	copyHeader(combinedHeaders, headers)
	req.Header = combinedHeaders

	if c.Key == "" && c.Email == "" && c.Token == "" && c.UserServiceKey == "" {
		return nil, ErrMissingCredentials
	}

	if c.Key != "" {
		req.Header.Set("X-Auth-Key", c.ClientParams.Key)
		req.Header.Set("X-Auth-Email", c.ClientParams.Email)
	}

	if c.UserServiceKey != "" {
		req.Header.Set("X-Auth-User-Service-Key", c.ClientParams.UserServiceKey)
	}

	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.ClientParams.Token)
	}

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.ClientParams.UserAgent)
	}

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}

	return resp, nil
}

func (c *Client) makeRequest(ctx context.Context, method, uri string, params interface{}, headers http.Header) ([]byte, error) {
	var reqBody io.Reader
	var err error

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

	var resp *http.Response
	var respErr error
	var respBody []byte

	if method == http.MethodPost || method == http.MethodPut || method == http.MethodPatch {
		buf := &bytes.Buffer{}
		tee := io.TeeReader(reqBody, buf)
		debugBody, _ := ioutil.ReadAll(tee)
		payloadBody, _ := ioutil.ReadAll(buf)
		c.Logger.Debugf("REQUEST Method:%v URI:%s Headers:%#v Body:%v\n", method, c.BaseURL.String()+uri, headers, string(debugBody))
		// ensure we recreate the io.Reader for use
		reqBody = bytes.NewReader(payloadBody)
	} else {
		c.Logger.Debugf("REQUEST Method:%v URI:%s Headers:%#v Body:%v\n", method, c.BaseURL.String()+uri, headers, nil) //)
	}

	resp, respErr = c.request(ctx, method, uri, reqBody, headers)
	if respErr != nil {
		return nil, respErr
	}

	c.Logger.Debugf("RESPONSE URI:%s StatusCode:%d Body:%#v RayID:%s\n", c.BaseURL.String()+uri, resp.StatusCode, string(respBody), resp.Header.Get("cf-ray"))

	respBody, err = ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("could not read response body: %w", err)
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
			return nil, fmt.Errorf("failed to unmarshal response body: %w", err)
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

	return respBody, nil
}

func (c *Client) get(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	return c.makeRequest(ctx, http.MethodGet, path, payload, nil)
}

func (c *Client) post(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	return c.makeRequest(ctx, http.MethodPost, path, payload, nil)
}

func (c *Client) patch(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	return c.makeRequest(ctx, http.MethodPatch, path, payload, nil)
}

func (c *Client) put(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	return c.makeRequest(ctx, http.MethodPut, path, payload, nil)
}

func (c *Client) delete(ctx context.Context, path string, payload interface{}) ([]byte, error) {
	return c.makeRequest(ctx, http.MethodDelete, path, payload, nil)
}
