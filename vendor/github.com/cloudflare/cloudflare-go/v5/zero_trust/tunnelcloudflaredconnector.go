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

// TunnelCloudflaredConnectorService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelCloudflaredConnectorService] method instead.
type TunnelCloudflaredConnectorService struct {
	Options []option.RequestOption
}

// NewTunnelCloudflaredConnectorService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelCloudflaredConnectorService(opts ...option.RequestOption) (r *TunnelCloudflaredConnectorService) {
	r = &TunnelCloudflaredConnectorService{}
	r.Options = opts
	return
}

// Fetches connector and connection details for a Cloudflare Tunnel.
func (r *TunnelCloudflaredConnectorService) Get(ctx context.Context, tunnelID string, connectorID string, query TunnelCloudflaredConnectorGetParams, opts ...option.RequestOption) (res *Client, err error) {
	var env TunnelCloudflaredConnectorGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connectors/%s", query.AccountID, tunnelID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelCloudflaredConnectorGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelCloudflaredConnectorGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A client (typically cloudflared) that maintains connections to a Cloudflare data
	// center.
	Result Client `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredConnectorGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredConnectorGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredConnectorGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [TunnelCloudflaredConnectorGetResponseEnvelope]
type tunnelCloudflaredConnectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConnectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredConnectorGetResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredConnectorGetResponseEnvelopeSuccessTrue TunnelCloudflaredConnectorGetResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredConnectorGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredConnectorGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
