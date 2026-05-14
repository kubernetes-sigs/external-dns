// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// NetworkRouteNetworkService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkRouteNetworkService] method instead.
type NetworkRouteNetworkService struct {
	Options []option.RequestOption
}

// NewNetworkRouteNetworkService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkRouteNetworkService(opts ...option.RequestOption) (r *NetworkRouteNetworkService) {
	r = &NetworkRouteNetworkService{}
	r.Options = opts
	return
}

// Routes a private network through a Cloudflare Tunnel. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
//
// Deprecated: This endpoint and its related APIs are deprecated in favor of the
// equivalent Tunnel Route (without CIDR) APIs.
func (r *NetworkRouteNetworkService) New(ctx context.Context, ipNetworkEncoded string, params NetworkRouteNetworkNewParams, opts ...option.RequestOption) (res *Route, err error) {
	var env NetworkRouteNetworkNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipNetworkEncoded == "" {
		err = errors.New("missing required ip_network_encoded parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", params.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a private network route from an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format. If no
// virtual_network_id is provided it will delete the route from the default vnet.
// If no tun_type is provided it will fetch the type from the tunnel_id or if that
// is missing it will assume Cloudflare Tunnel as default. If tunnel_id is provided
// it will delete the route from that tunnel, otherwise it will delete the route
// based on the vnet and tun_type.
//
// Deprecated: This endpoint and its related APIs are deprecated in favor of the
// equivalent Tunnel Route (without CIDR) APIs.
func (r *NetworkRouteNetworkService) Delete(ctx context.Context, ipNetworkEncoded string, params NetworkRouteNetworkDeleteParams, opts ...option.RequestOption) (res *Route, err error) {
	var env NetworkRouteNetworkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipNetworkEncoded == "" {
		err = errors.New("missing required ip_network_encoded parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", params.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing private network route in an account. The CIDR in
// `ip_network_encoded` must be written in URL-encoded format.
//
// Deprecated: This endpoint and its related APIs are deprecated in favor of the
// equivalent Tunnel Route (without CIDR) APIs.
func (r *NetworkRouteNetworkService) Edit(ctx context.Context, ipNetworkEncoded string, body NetworkRouteNetworkEditParams, opts ...option.RequestOption) (res *Route, err error) {
	var env NetworkRouteNetworkEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ipNetworkEncoded == "" {
		err = errors.New("missing required ip_network_encoded parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/teamnet/routes/network/%s", body.AccountID, ipNetworkEncoded)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetworkRouteNetworkNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// UUID of the tunnel.
	TunnelID param.Field[string] `json:"tunnel_id,required" format:"uuid"`
	// Optional remark describing the route.
	Comment param.Field[string] `json:"comment"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `json:"virtual_network_id" format:"uuid"`
}

func (r NetworkRouteNetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkRouteNetworkNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Route                 `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteNetworkNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteNetworkNewResponseEnvelopeJSON    `json:"-"`
}

// networkRouteNetworkNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkRouteNetworkNewResponseEnvelope]
type networkRouteNetworkNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteNetworkNewResponseEnvelopeSuccess bool

const (
	NetworkRouteNetworkNewResponseEnvelopeSuccessTrue NetworkRouteNetworkNewResponseEnvelopeSuccess = true
)

func (r NetworkRouteNetworkNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkRouteNetworkNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkRouteNetworkDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// The type of tunnel.
	TunType param.Field[NetworkRouteNetworkDeleteParamsTunType] `query:"tun_type"`
	// UUID of the tunnel.
	TunnelID param.Field[string] `query:"tunnel_id" format:"uuid"`
	// UUID of the virtual network.
	VirtualNetworkID param.Field[string] `query:"virtual_network_id" format:"uuid"`
}

// URLQuery serializes [NetworkRouteNetworkDeleteParams]'s query parameters as
// `url.Values`.
func (r NetworkRouteNetworkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of tunnel.
type NetworkRouteNetworkDeleteParamsTunType string

const (
	NetworkRouteNetworkDeleteParamsTunTypeCfdTunnel     NetworkRouteNetworkDeleteParamsTunType = "cfd_tunnel"
	NetworkRouteNetworkDeleteParamsTunTypeWARPConnector NetworkRouteNetworkDeleteParamsTunType = "warp_connector"
	NetworkRouteNetworkDeleteParamsTunTypeWARP          NetworkRouteNetworkDeleteParamsTunType = "warp"
	NetworkRouteNetworkDeleteParamsTunTypeMagic         NetworkRouteNetworkDeleteParamsTunType = "magic"
	NetworkRouteNetworkDeleteParamsTunTypeIPSec         NetworkRouteNetworkDeleteParamsTunType = "ip_sec"
	NetworkRouteNetworkDeleteParamsTunTypeGRE           NetworkRouteNetworkDeleteParamsTunType = "gre"
	NetworkRouteNetworkDeleteParamsTunTypeCNI           NetworkRouteNetworkDeleteParamsTunType = "cni"
)

func (r NetworkRouteNetworkDeleteParamsTunType) IsKnown() bool {
	switch r {
	case NetworkRouteNetworkDeleteParamsTunTypeCfdTunnel, NetworkRouteNetworkDeleteParamsTunTypeWARPConnector, NetworkRouteNetworkDeleteParamsTunTypeWARP, NetworkRouteNetworkDeleteParamsTunTypeMagic, NetworkRouteNetworkDeleteParamsTunTypeIPSec, NetworkRouteNetworkDeleteParamsTunTypeGRE, NetworkRouteNetworkDeleteParamsTunTypeCNI:
		return true
	}
	return false
}

type NetworkRouteNetworkDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Route                 `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteNetworkDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteNetworkDeleteResponseEnvelopeJSON    `json:"-"`
}

// networkRouteNetworkDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkRouteNetworkDeleteResponseEnvelope]
type networkRouteNetworkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteNetworkDeleteResponseEnvelopeSuccess bool

const (
	NetworkRouteNetworkDeleteResponseEnvelopeSuccessTrue NetworkRouteNetworkDeleteResponseEnvelopeSuccess = true
)

func (r NetworkRouteNetworkDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkRouteNetworkDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NetworkRouteNetworkEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type NetworkRouteNetworkEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Route                 `json:"result,required"`
	// Whether the API call was successful
	Success NetworkRouteNetworkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkRouteNetworkEditResponseEnvelopeJSON    `json:"-"`
}

// networkRouteNetworkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetworkRouteNetworkEditResponseEnvelope]
type networkRouteNetworkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkRouteNetworkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkRouteNetworkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkRouteNetworkEditResponseEnvelopeSuccess bool

const (
	NetworkRouteNetworkEditResponseEnvelopeSuccessTrue NetworkRouteNetworkEditResponseEnvelopeSuccess = true
)

func (r NetworkRouteNetworkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkRouteNetworkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
