// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// TunnelCloudflaredConnectionService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelCloudflaredConnectionService] method instead.
type TunnelCloudflaredConnectionService struct {
	Options []option.RequestOption
}

// NewTunnelCloudflaredConnectionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelCloudflaredConnectionService(opts ...option.RequestOption) (r *TunnelCloudflaredConnectionService) {
	r = &TunnelCloudflaredConnectionService{}
	r.Options = opts
	return
}

// Removes a connection (aka Cloudflare Tunnel Connector) from a Cloudflare Tunnel
// independently of its current state. If no connector id (client_id) is provided
// all connectors will be removed. We recommend running this command after rotating
// tokens.
func (r *TunnelCloudflaredConnectionService) Delete(ctx context.Context, tunnelID string, params TunnelCloudflaredConnectionDeleteParams, opts ...option.RequestOption) (res *TunnelCloudflaredConnectionDeleteResponse, err error) {
	var env TunnelCloudflaredConnectionDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *TunnelCloudflaredConnectionService) Get(ctx context.Context, tunnelID string, query TunnelCloudflaredConnectionGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[Client], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/connections", query.AccountID, tunnelID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches connection details for a Cloudflare Tunnel.
func (r *TunnelCloudflaredConnectionService) GetAutoPaging(ctx context.Context, tunnelID string, query TunnelCloudflaredConnectionGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Client] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, tunnelID, query, opts...))
}

// A client (typically cloudflared) that maintains connections to a Cloudflare data
// center.
type Client struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// The cloudflared OS architecture used to establish this connection.
	Arch string `json:"arch"`
	// The version of the remote tunnel configuration. Used internally to sync
	// cloudflared with the Zero Trust dashboard.
	ConfigVersion int64 `json:"config_version"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	Conns []ClientConn `json:"conns"`
	// Features enabled for the Cloudflare Tunnel.
	Features []string `json:"features"`
	// Timestamp of when the tunnel connection was started.
	RunAt time.Time `json:"run_at" format:"date-time"`
	// The cloudflared version used to establish this connection.
	Version string     `json:"version"`
	JSON    clientJSON `json:"-"`
}

// clientJSON contains the JSON metadata for the struct [Client]
type clientJSON struct {
	ID            apijson.Field
	Arch          apijson.Field
	ConfigVersion apijson.Field
	Conns         apijson.Field
	Features      apijson.Field
	RunAt         apijson.Field
	Version       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Client) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientJSON) RawJSON() string {
	return r.raw
}

type ClientConn struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	UUID string         `json:"uuid" format:"uuid"`
	JSON clientConnJSON `json:"-"`
}

// clientConnJSON contains the JSON metadata for the struct [ClientConn]
type clientConnJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ClientConn) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r clientConnJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConnectionDeleteResponse = interface{}

type TunnelCloudflaredConnectionDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID param.Field[string] `query:"client_id" format:"uuid"`
}

// URLQuery serializes [TunnelCloudflaredConnectionDeleteParams]'s query parameters
// as `url.Values`.
func (r TunnelCloudflaredConnectionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TunnelCloudflaredConnectionDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo                     `json:"errors,required"`
	Messages []shared.ResponseInfo                     `json:"messages,required"`
	Result   TunnelCloudflaredConnectionDeleteResponse `json:"result,required,nullable"`
	// Whether the API call was successful
	Success TunnelCloudflaredConnectionDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredConnectionDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredConnectionDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [TunnelCloudflaredConnectionDeleteResponseEnvelope]
type tunnelCloudflaredConnectionDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConnectionDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConnectionDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredConnectionDeleteResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredConnectionDeleteResponseEnvelopeSuccessTrue TunnelCloudflaredConnectionDeleteResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredConnectionDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredConnectionDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredConnectionGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}
