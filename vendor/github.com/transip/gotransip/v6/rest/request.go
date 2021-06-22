package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	contentType = "application/json"
)

// Request will be used by all repositories and the client.
// The Request struct can be transformed to a http request with method, url and optional body.
type Request struct {
	// Endpoint is the api endpoint, without the server, which we will receive the request,
	// like: '/products'
	Endpoint string
	// Parameters is a map of strings (url.Values) that will be used to add
	// http query strings to the request, like '/domains?tags=123'
	Parameters url.Values
	// Body is left as an interface because the Request is not coupled to any specific Request.Body struct
	Body interface{}
	// TestMode is used when users want to tinker with the api without touching their real data
	TestMode bool
}

// GetJSONBody returns the request object as a json byte array
func (r *Request) GetJSONBody() ([]byte, error) {
	return json.Marshal(r.Body)
}

// GetBodyReader returns an io.Reader for the json marshalled body of this request
// this will be used by the writer used in the client.
func (r *Request) GetBodyReader() (io.Reader, error) {
	// try to get the marshalled body
	body, err := r.GetJSONBody()
	if err != nil {
		return nil, fmt.Errorf("error when marshaling request: %w", err)
	}

	return bytes.NewReader(body), nil
}

// GetHTTPRequest generates and returns a http.Request object.
// It does this with the Request struct and the basePath and method,
// that are provided by the client itself.
func (r *Request) GetHTTPRequest(basePath string, method string) (*http.Request, error) {
	requestURL := basePath + r.Endpoint

	var bodyReader io.Reader
	if r.Body != nil {
		reader, err := r.GetBodyReader()
		if err != nil {
			return nil, err
		}

		bodyReader = reader
	}

	request, err := http.NewRequest(method, requestURL, bodyReader)
	if err != nil {
		return nil, err
	}

	// set json headers, because our this library sends and expects that
	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)

	// if TestMode is true we always add a test=1 http query string to the url,
	// this is used when users want to tinker with the api without changing their production data
	if r.TestMode {
		// if Parameters is not set yet, we create a new url.Values and set it
		if r.Parameters == nil {
			r.Parameters = url.Values{}
		}

		r.Parameters.Add("test", "1")
	}

	// set the custom parameters on the rawquery
	request.URL.RawQuery = r.Parameters.Encode()

	return request, nil
}
