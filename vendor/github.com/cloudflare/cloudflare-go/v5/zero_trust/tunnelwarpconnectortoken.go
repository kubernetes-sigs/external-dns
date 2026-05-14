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

// TunnelWARPConnectorTokenService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelWARPConnectorTokenService] method instead.
type TunnelWARPConnectorTokenService struct {
	Options []option.RequestOption
}

// NewTunnelWARPConnectorTokenService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelWARPConnectorTokenService(opts ...option.RequestOption) (r *TunnelWARPConnectorTokenService) {
	r = &TunnelWARPConnectorTokenService{}
	r.Options = opts
	return
}

// Gets the token used to associate warp device with a specific Warp Connector
// tunnel.
func (r *TunnelWARPConnectorTokenService) Get(ctx context.Context, tunnelID string, query TunnelWARPConnectorTokenGetParams, opts ...option.RequestOption) (res *string, err error) {
	var env TunnelWARPConnectorTokenGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/warp_connector/%s/token", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelWARPConnectorTokenGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelWARPConnectorTokenGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The Tunnel Token is used as a mechanism to authenticate the operation of a
	// tunnel.
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success TunnelWARPConnectorTokenGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelWARPConnectorTokenGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelWARPConnectorTokenGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [TunnelWARPConnectorTokenGetResponseEnvelope]
type tunnelWARPConnectorTokenGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelWARPConnectorTokenGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelWARPConnectorTokenGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelWARPConnectorTokenGetResponseEnvelopeSuccess bool

const (
	TunnelWARPConnectorTokenGetResponseEnvelopeSuccessTrue TunnelWARPConnectorTokenGetResponseEnvelopeSuccess = true
)

func (r TunnelWARPConnectorTokenGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelWARPConnectorTokenGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
