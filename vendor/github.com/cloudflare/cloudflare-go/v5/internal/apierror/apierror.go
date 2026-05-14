// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	Errors     []shared.ErrorData `json:"errors,required"`
	JSON       errorJSON          `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// errorJSON contains the JSON metadata for the struct [Error]
type errorJSON struct {
	Errors      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorJSON) RawJSON() string {
	return r.raw
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s \"%s\": %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.RawJSON())
}

func (r *Error) DumpRequest(body bool) []byte {
	if r.Request.GetBody != nil {
		r.Request.Body, _ = r.Request.GetBody()
	}
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}
