package linodego

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

// paginatedResponse represents a single response from a paginated
// endpoint.
type paginatedResponse[T any] struct {
	Page    int `json:"page"    url:"page,omitempty"`
	Pages   int `json:"pages"   url:"pages,omitempty"`
	Results int `json:"results" url:"results,omitempty"`
	Data    []T `json:"data"`
}

// handlePaginatedResults aggregates results from the given
// paginated endpoint using the provided ListOptions and HTTP method.
// nolint:funlen
func handlePaginatedResults[T any, O any](
	ctx context.Context,
	client *Client,
	endpoint string,
	opts *ListOptions,
	method string,
	options ...O,
) ([]T, error) {
	result := make([]T, 0)

	if opts == nil {
		opts = &ListOptions{PageOptions: &PageOptions{Page: 0}}
	}

	if opts.PageOptions == nil {
		opts.PageOptions = &PageOptions{Page: 0}
	}

	// Validate options
	numOpts := len(options)
	if numOpts > 1 {
		return nil, fmt.Errorf("invalid number of options: expected 0 or 1, got %d", numOpts)
	}

	// Prepare request body if options are provided
	var reqBody string

	if numOpts > 0 && !isNil(options[0]) {
		body, err := json.Marshal(options[0])
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}

		reqBody = string(body)
	}

	// Makes a request to a particular page and appends the response to the result
	handlePage := func(page int) error {
		var resultType paginatedResponse[T]

		// Override the page to be applied in applyListOptionsToRequest(...)
		opts.Page = page

		// This request object cannot be reused for each page request
		// because it can lead to possible data corruption
		req := client.R(ctx).SetResult(&resultType)

		// Apply all user-provided list options to the request
		if err := applyListOptionsToRequest(opts, req); err != nil {
			return err
		}

		// Set request body if provided
		if reqBody != "" {
			req.SetBody(reqBody)
		}

		var response *paginatedResponse[T]
		// Execute the appropriate HTTP method
		switch method {
		case "GET":
			res, err := coupleAPIErrors(req.Get(endpoint))
			if err != nil {
				return err
			}

			response = res.Result().(*paginatedResponse[T])
		case "PUT":
			res, err := coupleAPIErrors(req.Put(endpoint))
			if err != nil {
				return err
			}

			response = res.Result().(*paginatedResponse[T])
		case "POST":
			res, err := coupleAPIErrors(req.Post(endpoint))
			if err != nil {
				return err
			}

			response = res.Result().(*paginatedResponse[T])
		default:
			return fmt.Errorf("unsupported HTTP method: %s", method)
		}

		// Update pagination metadata
		opts.Page = page
		opts.Pages = response.Pages
		opts.Results = response.Results
		result = append(result, response.Data...)

		return nil
	}

	// Determine starting page
	startingPage := 1
	pageDefined := opts.Page > 0

	if pageDefined {
		startingPage = opts.Page
	}

	// Get the first page
	if err := handlePage(startingPage); err != nil {
		return nil, err
	}

	// If the user has explicitly specified a page, we don't
	// need to get any other pages.
	if pageDefined {
		return result, nil
	}

	// Get the rest of the pages
	for page := 2; page <= opts.Pages; page++ {
		if err := handlePage(page); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// getPaginatedResults aggregates results from the given
// paginated endpoint using the provided ListOptions.
func getPaginatedResults[T any](
	ctx context.Context,
	client *Client,
	endpoint string,
	opts *ListOptions,
) ([]T, error) {
	return handlePaginatedResults[T, any](ctx, client, endpoint, opts, "GET")
}

// putPaginatedResults sends a PUT request and aggregates the results from the given
// paginated endpoint using the provided ListOptions.
func putPaginatedResults[T, O any](
	ctx context.Context,
	client *Client,
	endpoint string,
	opts *ListOptions,
	options ...O,
) ([]T, error) {
	return handlePaginatedResults[T, O](ctx, client, endpoint, opts, "PUT", options...)
}

// postPaginatedResults sends a POST request and aggregates the results from the given
// paginated endpoint using the provided ListOptions.
func postPaginatedResults[T, O any](
	ctx context.Context,
	client *Client,
	endpoint string,
	opts *ListOptions,
	options ...O,
) ([]T, error) {
	return handlePaginatedResults[T, O](ctx, client, endpoint, opts, "POST", options...)
}

// doGETRequest runs a GET request using the given client and API endpoint,
// and returns the result
func doGETRequest[T any](
	ctx context.Context,
	client *Client,
	endpoint string,
) (*T, error) {
	var resultType T

	req := client.R(ctx).SetResult(&resultType)

	r, err := coupleAPIErrors(req.Get(endpoint))
	if err != nil {
		return nil, err
	}

	return r.Result().(*T), nil
}

// doPOSTRequest runs a PUT request using the given client, API endpoint,
// and options/body.
func doPOSTRequest[T, O any](
	ctx context.Context,
	client *Client,
	endpoint string,
	options ...O,
) (*T, error) {
	var resultType T

	numOpts := len(options)

	if numOpts > 1 {
		return nil, fmt.Errorf("invalid number of options: %d", len(options))
	}

	req := client.R(ctx).SetResult(&resultType)

	if numOpts > 0 && !isNil(options[0]) {
		body, err := json.Marshal(options[0])
		if err != nil {
			return nil, err
		}

		req.SetBody(string(body))
	}

	r, err := coupleAPIErrors(req.Post(endpoint))
	if err != nil {
		return nil, err
	}

	return r.Result().(*T), nil
}

// doPOSTRequestNoResponseBody runs a POST request using the given client, API endpoint,
// and options/body. It expects only empty response from the endpoint.
func doPOSTRequestNoResponseBody[T any](
	ctx context.Context,
	client *Client,
	endpoint string,
	options ...T,
) error {
	_, err := doPOSTRequest[any, T](ctx, client, endpoint, options...)
	return err
}

// doPOSTRequestNoRequestResponseBody runs a POST request where no request body is needed and no response body
// is expected from the endpoints.
func doPOSTRequestNoRequestResponseBody(
	ctx context.Context,
	client *Client,
	endpoint string,
) error {
	return doPOSTRequestNoResponseBody(ctx, client, endpoint, struct{}{})
}

// doPUTRequest runs a PUT request using the given client, API endpoint,
// and options/body.
func doPUTRequest[T, O any](
	ctx context.Context,
	client *Client,
	endpoint string,
	options ...O,
) (*T, error) {
	var resultType T

	numOpts := len(options)

	if numOpts > 1 {
		return nil, fmt.Errorf("invalid number of options: %d", len(options))
	}

	req := client.R(ctx).SetResult(&resultType)

	if numOpts > 0 && !isNil(options[0]) {
		body, err := json.Marshal(options[0])
		if err != nil {
			return nil, err
		}

		req.SetBody(string(body))
	}

	r, err := coupleAPIErrors(req.Put(endpoint))
	if err != nil {
		return nil, err
	}

	return r.Result().(*T), nil
}

// doDELETERequest runs a DELETE request using the given client
// and API endpoint.
func doDELETERequest(
	ctx context.Context,
	client *Client,
	endpoint string,
) error {
	req := client.R(ctx)
	_, err := coupleAPIErrors(req.Delete(endpoint))

	return err
}

// formatAPIPath allows us to safely build an API request with path escaping
func formatAPIPath(format string, args ...any) string {
	escapedArgs := make([]any, len(args))
	for i, arg := range args {
		if typeStr, ok := arg.(string); ok {
			arg = url.PathEscape(typeStr)
		}

		escapedArgs[i] = arg
	}

	return fmt.Sprintf(format, escapedArgs...)
}

func isNil(i any) bool {
	if i == nil {
		return true
	}

	// Check for nil pointers
	v := reflect.ValueOf(i)

	return v.Kind() == reflect.Pointer && v.IsNil()
}
