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

// TunnelCloudflaredManagementService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelCloudflaredManagementService] method instead.
type TunnelCloudflaredManagementService struct {
	Options []option.RequestOption
}

// NewTunnelCloudflaredManagementService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelCloudflaredManagementService(opts ...option.RequestOption) (r *TunnelCloudflaredManagementService) {
	r = &TunnelCloudflaredManagementService{}
	r.Options = opts
	return
}

// Gets a management token used to access the management resources (i.e. Streaming
// Logs) of a tunnel.
func (r *TunnelCloudflaredManagementService) New(ctx context.Context, tunnelID string, params TunnelCloudflaredManagementNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env TunnelCloudflaredManagementNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/management", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TunnelCloudflaredManagementNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string]                                         `path:"account_id,required"`
	Resources param.Field[[]TunnelCloudflaredManagementNewParamsResource] `json:"resources,required"`
}

func (r TunnelCloudflaredManagementNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Management resources the token will have access to.
type TunnelCloudflaredManagementNewParamsResource string

const (
	TunnelCloudflaredManagementNewParamsResourceLogs TunnelCloudflaredManagementNewParamsResource = "logs"
)

func (r TunnelCloudflaredManagementNewParamsResource) IsKnown() bool {
	switch r {
	case TunnelCloudflaredManagementNewParamsResourceLogs:
		return true
	}
	return false
}

type TunnelCloudflaredManagementNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// The Tunnel Token is used as a mechanism to authenticate the operation of a
	// tunnel.
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredManagementNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredManagementNewResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredManagementNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [TunnelCloudflaredManagementNewResponseEnvelope]
type tunnelCloudflaredManagementNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredManagementNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredManagementNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredManagementNewResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredManagementNewResponseEnvelopeSuccessTrue TunnelCloudflaredManagementNewResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredManagementNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredManagementNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
