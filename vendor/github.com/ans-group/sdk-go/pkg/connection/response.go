package connection

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ans-group/sdk-go/pkg/logging"
)

type ResponseBody interface {
	Error() error
	Pagination() APIResponseMetadataPagination
}

// APIResponse represents the base API response
type APIResponse struct {
	*http.Response
}

// APIResponseBody represents the base API response body
type APIResponseBody struct {
	APIResponseBodyError

	Metadata APIResponseMetadata `json:"meta"`
}

type APIResponseBodyError struct {
	Errors  []APIResponseBodyErrorItem `json:"errors"`
	Message string                     `json:"message"`
}

func (e *APIResponseBodyError) Error() string {
	var errArr []string

	// Message will be populated in certain circumstances, populate error array if exists
	if e.Message != "" {
		errArr = append(errArr, fmt.Sprintf("message=\"%s\"", e.Message))
	}

	for _, err := range e.Errors {
		errArr = append(errArr, err.Error())
	}

	return strings.Join(errArr, "; ")
}

// APIResponseBodyErrorItem represents an API response error
type APIResponseBodyErrorItem struct {
	Title  string `json:"title"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
	Source string `json:"source"`
}

func (a *APIResponseBodyErrorItem) Error() string {
	return fmt.Sprintf("title=\"%s\", detail=\"%s\", status=\"%d\", source=\"%s\"", a.Title, a.Detail, a.Status, a.Source)
}

// APIResponseMetadata represents the API response metadata
type APIResponseMetadata struct {
	Pagination APIResponseMetadataPagination `json:"pagination"`
}

// APIResponseMetadataPagination represents the API response pagination data
type APIResponseMetadataPagination struct {
	Total      int                                `json:"total"`
	Count      int                                `json:"count"`
	PerPage    int                                `json:"per_page"`
	TotalPages int                                `json:"total_pages"`
	Links      APIResponseMetadataPaginationLinks `json:"links"`
}

// APIResponseMetadataPaginationLinks represents the links returned within the API response pagination data
type APIResponseMetadataPaginationLinks struct {
	Next     string `json:"next"`
	Previous string `json:"previous"`
	First    string `json:"first"`
	Last     string `json:"last"`
}

// DeserializeResponseBody deserializes the API response body and stores the result
// in parameter out
func (r *APIResponse) DeserializeResponseBody(out interface{}) error {
	defer r.Response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(r.Response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body with response status code %d: %s", r.StatusCode, err)
	}

	logging.Debugf("Response body: %s", string(bodyBytes))

	if len(bodyBytes) > 0 {
		return json.Unmarshal(bodyBytes, out)
	}

	return nil
}

// ValidateStatusCode validates the API response
func (r *APIResponse) ValidateStatusCode(codes []int, respBody ResponseBody) error {
	if len(codes) > 0 {
		for _, code := range codes {
			if r.StatusCode == code {
				return nil
			}
		}
	} else {
		if r.StatusCode >= 200 && r.StatusCode <= 299 {
			return nil
		}
	}

	return fmt.Errorf("unexpected status code (%d): %w", r.StatusCode, respBody.Error())
}

type ResponseHandler func(resp *APIResponse) error

func StatusCodeResponseHandler(code int, err error) ResponseHandler {
	return func(resp *APIResponse) error {
		if resp.StatusCode == code {
			return err
		}

		return nil
	}
}

func NotFoundResponseHandler(err error) ResponseHandler {
	return StatusCodeResponseHandler(404, err)
}

// HandleResponse deserializes the response body into provided respBody, and validates
// the response using the optionally provided ResponseHandler handler
func (r *APIResponse) HandleResponse(respBody ResponseBody, handlers ...ResponseHandler) error {
	err := r.DeserializeResponseBody(respBody)
	if err != nil {
		return err
	}

	for _, handler := range handlers {
		if handler != nil {
			err = handler(r)
			if err != nil {
				return err
			}
		}
	}

	return r.ValidateStatusCode([]int{}, respBody)
}

// Error returns an error struct with embedded errors from body
func (a *APIResponseBody) Error() error {
	return &APIResponseBodyError{
		Message: a.Message,
		Errors:  a.Errors,
	}
}

// TotalPages returns amount of pages for API response
func (a *APIResponseBody) Pagination() APIResponseMetadataPagination {
	return a.Metadata.Pagination
}

// APIResponseBodyStringData represents the API response body containing generic data
type APIResponseBodyData[T any] struct {
	APIResponseBody

	Data T `json:"data"`
}

// APIResponseBodyStringData represents the API response body containing string data
type APIResponseBodyStringData struct {
	APIResponseBody

	Data string `json:"data"`
}
