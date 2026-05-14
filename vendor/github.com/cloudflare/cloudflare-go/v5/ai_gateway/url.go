// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// URLService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewURLService] method instead.
type URLService struct {
	Options []option.RequestOption
}

// NewURLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewURLService(opts ...option.RequestOption) (r *URLService) {
	r = &URLService{}
	r.Options = opts
	return
}

// Get Gateway URL
func (r *URLService) Get(ctx context.Context, gatewayID string, provider string, query URLGetParams, opts ...option.RequestOption) (res *string, err error) {
	var env URLGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if provider == "" {
		err = errors.New("missing required provider parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/url/%s", query.AccountID, gatewayID, provider)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type URLGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type URLGetResponseEnvelope struct {
	Result  string                     `json:"result,required"`
	Success bool                       `json:"success,required"`
	JSON    urlGetResponseEnvelopeJSON `json:"-"`
}

// urlGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [URLGetResponseEnvelope]
type urlGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
