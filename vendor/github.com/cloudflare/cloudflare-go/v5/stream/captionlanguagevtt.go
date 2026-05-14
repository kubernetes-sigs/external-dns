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

// CaptionLanguageVttService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCaptionLanguageVttService] method instead.
type CaptionLanguageVttService struct {
	Options []option.RequestOption
}

// NewCaptionLanguageVttService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCaptionLanguageVttService(opts ...option.RequestOption) (r *CaptionLanguageVttService) {
	r = &CaptionLanguageVttService{}
	r.Options = opts
	return
}

// Return WebVTT captions for a provided language.
func (r *CaptionLanguageVttService) Get(ctx context.Context, identifier string, language string, query CaptionLanguageVttGetParams, opts ...option.RequestOption) (res *string, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/vtt")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/stream/%s/captions/%s/vtt", query.AccountID, identifier, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CaptionLanguageVttGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
