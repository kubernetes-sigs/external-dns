// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package stream

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// EmbedService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmbedService] method instead.
type EmbedService struct {
	Options []option.RequestOption
}

// NewEmbedService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEmbedService(opts ...option.RequestOption) (r *EmbedService) {
	r = &EmbedService{}
	r.Options = opts
	return
}

// Fetches an HTML code snippet to embed a video in a web page delivered through
// Cloudflare. On success, returns an HTML fragment for use on web pages to display
// a video. On failure, returns a JSON response body.
func (r *EmbedService) Get(ctx context.Context, identifier string, query EmbedGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/embed", query.AccountID, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EmbedGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}
