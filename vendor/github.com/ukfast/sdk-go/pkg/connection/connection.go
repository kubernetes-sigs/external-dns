package connection

//go:generate mockgen -package mocks -destination ../../test/mocks/mock_connection.go github.com/ukfast/sdk-go/pkg/connection Connection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/ukfast/sdk-go/pkg/logging"
)

const apiURI = "api.ukfast.io"
const apiScheme = "https"
const userAgent = "ukfast-sdk-go"

type Connection interface {
	Get(resource string, parameters APIRequestParameters) (*APIResponse, error)
	Post(resource string, body interface{}) (*APIResponse, error)
	Put(resource string, body interface{}) (*APIResponse, error)
	Patch(resource string, body interface{}) (*APIResponse, error)
	Delete(resource string, body interface{}) (*APIResponse, error)
	Invoke(request APIRequest) (*APIResponse, error)
}

type APIConnection struct {
	HTTPClient  *http.Client
	Credentials Credentials
	APIURI      string
	APIScheme   string
	Headers     http.Header
	UserAgent   string
}

// NewAPIKeyCredentialsAPIConnection creates a new client
func NewAPIKeyCredentialsAPIConnection(apiKey string) *APIConnection {
	return NewAPIConnection(&APIKeyCredentials{APIKey: apiKey})
}

func NewAPIConnection(credentials Credentials) *APIConnection {
	return &APIConnection{
		Credentials: credentials,
		HTTPClient:  &http.Client{},
		APIURI:      apiURI,
		APIScheme:   apiScheme,
		UserAgent:   userAgent,
	}
}

// composeURI returns a composed URI for given resource and request modifiers
func (c *APIConnection) composeURI(resource string, pagination APIRequestPagination, sorting APIRequestSorting, filtering []APIRequestFiltering) string {
	data := url.Values{}
	c.hydratePaginationQuery(&data, pagination)
	c.hydrateSortingQuery(&data, sorting)
	c.hydrateFilteringQuery(&data, filtering)

	q := data.Encode()
	// Add query parameter start
	if q != "" {
		q = "?" + q
	}

	return fmt.Sprintf("%s://%s/%s%s", c.APIScheme, c.APIURI, strings.Trim(resource, "/"), q)
}

// hydratePaginationQuery populates query parameters with pagination query parameters, if any
func (c *APIConnection) hydratePaginationQuery(q *url.Values, pagination APIRequestPagination) {
	if pagination.Page != 0 {
		q.Add("page", strconv.Itoa(pagination.Page))
	}
	if pagination.PerPage != 0 {
		q.Add("per_page", strconv.Itoa(pagination.PerPage))
	}
}

// hydrateFilteringQuery populates query parameters with filtering query parameters, if any
func (c *APIConnection) hydrateFilteringQuery(q *url.Values, filtering []APIRequestFiltering) {
	for _, filter := range filtering {
		q.Add(fmt.Sprintf("%s:%s", filter.Property, filter.Operator.String()), strings.Join(filter.Value, ","))
	}
}

// hydrateSortingQuery populates query parameters with sorting query parameters, if any
func (c *APIConnection) hydrateSortingQuery(q *url.Values, sorting APIRequestSorting) {
	if sorting.Property != "" {
		direction := "asc"
		if sorting.Descending {
			direction = "desc"
		}

		q.Add("sort", fmt.Sprintf("%s:%s", sorting.Property, direction))
	}
}

// Get invokes a GET request, returning an APIResponse
func (c *APIConnection) Get(resource string, parameters APIRequestParameters) (*APIResponse, error) {
	return c.Invoke(APIRequest{
		Method:     "GET",
		Resource:   resource,
		Parameters: parameters,
	})
}

// Post invokes a POST request, returning an APIResponse
func (c *APIConnection) Post(resource string, body interface{}) (*APIResponse, error) {
	return c.Invoke(APIRequest{
		Method:   "POST",
		Resource: resource,
		Body:     body,
	})
}

// Put invokes a PUT request, returning an APIResponse
func (c *APIConnection) Put(resource string, body interface{}) (*APIResponse, error) {
	return c.Invoke(APIRequest{
		Method:   "PUT",
		Resource: resource,
		Body:     body,
	})
}

// Patch invokes a PATCH request, returning an APIResponse
func (c *APIConnection) Patch(resource string, body interface{}) (*APIResponse, error) {
	return c.Invoke(APIRequest{
		Method:   "PATCH",
		Resource: resource,
		Body:     body,
	})
}

// Delete invokes a DELETE request, returning an APIResponse
func (c *APIConnection) Delete(resource string, body interface{}) (*APIResponse, error) {
	return c.Invoke(APIRequest{
		Method:   "DELETE",
		Resource: resource,
		Body:     body,
	})
}

// Invoke invokes a request, returning an APIResponse
func (c *APIConnection) Invoke(request APIRequest) (*APIResponse, error) {
	req, err := c.NewRequest(request)
	if err != nil {
		return nil, err
	}

	return c.InvokeRequest(req)
}

// NewRequest generates a new Request from given parameters
func (c *APIConnection) getBody(request APIRequest) (io.Reader, error) {
	buf := new(bytes.Buffer)
	if request.Body != nil {
		if reader, ok := request.Body.(io.Reader); ok {
			return reader, nil
		}

		if v, ok := request.Body.(Validatable); ok {
			valErr := v.Validate()
			if valErr != nil {
				return nil, valErr
			}
		}

		err := json.NewEncoder(buf).Encode(request.Body)
		if err != nil {
			return nil, err
		}
		logging.Tracef("Encoded body: %s", buf)
	}

	return buf, nil
}

// NewRequest generates a new Request from given parameters
func (c *APIConnection) NewRequest(request APIRequest) (*http.Request, error) {
	uri := c.composeURI(request.Resource, request.Parameters.Pagination, request.Parameters.Sorting, request.Parameters.Filtering)

	logging.Debugf("Generated URI: %s", uri)

	body, err := c.getBody(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(request.Method, uri, body)
	if err != nil {
		return nil, err
	}

	// Add standard headers
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)
	for headerKey, headerValue := range c.Credentials.GetAuthHeaders() {
		req.Header.Add(headerKey, headerValue)
	}

	// Append additional headers if defined
	if c.Headers != nil {
		for headerKey, headerValues := range c.Headers {
			for _, headerValue := range headerValues {
				req.Header.Add(headerKey, headerValue)
			}
		}
	}

	return req, nil
}

// InvokeRequest invokes a request, returning an APIResponse
func (c *APIConnection) InvokeRequest(req *http.Request) (*APIResponse, error) {
	resp := &APIResponse{}

	logging.Debugf("Executing request: %s %s", req.Method, req.URL.String())

	r, err := c.HTTPClient.Do(req)
	if err != nil {
		return resp, fmt.Errorf("api request failed: %w", err)
	}

	logging.Debugf("Got response: StatusCode=[%d]", r.StatusCode)

	resp.Response = r

	return resp, nil
}
