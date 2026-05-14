// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// TunnelCloudflaredTokenService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelCloudflaredTokenService] method instead.
type TunnelCloudflaredTokenService struct {
	Options []option.RequestOption
}

// NewTunnelCloudflaredTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelCloudflaredTokenService(opts ...option.RequestOption) (r *TunnelCloudflaredTokenService) {
	r = &TunnelCloudflaredTokenService{}
	r.Options = opts
	return
}

// Gets the token used to associate cloudflared with a specific tunnel.
func (r *TunnelCloudflaredTokenService) Get(ctx context.Context, tunnelID string, query TunnelCloudflaredTokenGetParams, opts ...option.RequestOption) (res *string, err error) {
	var env TunnelCloudflaredTokenGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/token", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelCloudflaredTokenGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelCloudflaredTokenGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The Tunnel Token is used as a mechanism to authenticate the operation of a
	// tunnel.
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredTokenGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredTokenGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredTokenGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredTokenGetResponseEnvelope]
type tunnelCloudflaredTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredTokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredTokenGetResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredTokenGetResponseEnvelopeSuccessTrue TunnelCloudflaredTokenGetResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredTokenGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredTokenGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
