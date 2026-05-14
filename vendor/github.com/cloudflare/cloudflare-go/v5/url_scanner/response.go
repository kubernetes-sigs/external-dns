// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package url_scanner

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ResponseService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResponseService] method instead.
type ResponseService struct {
	Options []option.RequestOption
}

// NewResponseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResponseService(opts ...option.RequestOption) (r *ResponseService) {
	r = &ResponseService{}
	r.Options = opts
	return
}

// Returns the raw response of the network request. Find the `response_id` in the
// `data.requests.response.hash`.
func (r *ResponseService) Get(ctx context.Context, responseID string, query ResponseGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if responseID == "" {
		err = errors.New("missing required response_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/urlscanner/v2/responses/%s", query.AccountID, responseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ResponseGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id,required"`
}
