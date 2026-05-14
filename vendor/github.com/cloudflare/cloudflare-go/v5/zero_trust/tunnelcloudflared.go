// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// TunnelCloudflaredService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelCloudflaredService] method instead.
type TunnelCloudflaredService struct {
	Options        []option.RequestOption
	Configurations *TunnelCloudflaredConfigurationService
	Connections    *TunnelCloudflaredConnectionService
	Token          *TunnelCloudflaredTokenService
	Connectors     *TunnelCloudflaredConnectorService
	Management     *TunnelCloudflaredManagementService
}

// NewTunnelCloudflaredService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTunnelCloudflaredService(opts ...option.RequestOption) (r *TunnelCloudflaredService) {
	r = &TunnelCloudflaredService{}
	r.Options = opts
	r.Configurations = NewTunnelCloudflaredConfigurationService(opts...)
	r.Connections = NewTunnelCloudflaredConnectionService(opts...)
	r.Token = NewTunnelCloudflaredTokenService(opts...)
	r.Connectors = NewTunnelCloudflaredConnectorService(opts...)
	r.Management = NewTunnelCloudflaredManagementService(opts...)
	return
}

// Creates a new Cloudflare Tunnel in an account.
func (r *TunnelCloudflaredService) New(ctx context.Context, params TunnelCloudflaredNewParams, opts ...option.RequestOption) (res *TunnelCloudflaredNewResponse, err error) {
	var env TunnelCloudflaredNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists and filters Cloudflare Tunnels in an account.
func (r *TunnelCloudflaredService) List(ctx context.Context, params TunnelCloudflaredListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[TunnelCloudflaredListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists and filters Cloudflare Tunnels in an account.
func (r *TunnelCloudflaredService) ListAutoPaging(ctx context.Context, params TunnelCloudflaredListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[TunnelCloudflaredListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a Cloudflare Tunnel from an account.
func (r *TunnelCloudflaredService) Delete(ctx context.Context, tunnelID string, body TunnelCloudflaredDeleteParams, opts ...option.RequestOption) (res *TunnelCloudflaredDeleteResponse, err error) {
	var env TunnelCloudflaredDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", body.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing Cloudflare Tunnel.
func (r *TunnelCloudflaredService) Edit(ctx context.Context, tunnelID string, params TunnelCloudflaredEditParams, opts ...option.RequestOption) (res *TunnelCloudflaredEditResponse, err error) {
	var env TunnelCloudflaredEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Cloudflare Tunnel.
func (r *TunnelCloudflaredService) Get(ctx context.Context, tunnelID string, query TunnelCloudflaredGetParams, opts ...option.RequestOption) (res *TunnelCloudflaredGetResponse, err error) {
	var env TunnelCloudflaredGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredNewResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc TunnelCloudflaredNewResponseConfigSrc `json:"config_src"`
	// This field can have the runtime type of [[]shared.CloudflareTunnelConnection],
	// [[]TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnection].
	Connections interface{} `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	//
	// Deprecated: Use the config_src field instead.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredNewResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredNewResponseTunType `json:"tun_type"`
	JSON    tunnelCloudflaredNewResponseJSON    `json:"-"`
	union   TunnelCloudflaredNewResponseUnion
}

// tunnelCloudflaredNewResponseJSON contains the JSON metadata for the struct
// [TunnelCloudflaredNewResponse]
type tunnelCloudflaredNewResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	ConfigSrc       apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelCloudflaredNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelCloudflaredNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelCloudflaredNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelCloudflaredNewResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.CloudflareTunnel],
// [TunnelCloudflaredNewResponseTunnelWARPConnectorTunnel].
func (r TunnelCloudflaredNewResponse) AsUnion() TunnelCloudflaredNewResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [TunnelCloudflaredNewResponseTunnelWARPConnectorTunnel].
type TunnelCloudflaredNewResponseUnion interface {
	ImplementsTunnelCloudflaredNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelCloudflaredNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelCloudflaredNewResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredNewResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [TunnelCloudflaredNewResponseTunnelWARPConnectorTunnel]
type tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TunnelCloudflaredNewResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelCloudflaredNewResponseTunnelWARPConnectorTunnel) ImplementsTunnelCloudflaredNewResponse() {
}

type TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                              `json:"uuid" format:"uuid"`
	JSON tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnection]
type tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredNewResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatus string

const (
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusInactive TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatus = "inactive"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusDegraded TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatus = "degraded"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusHealthy  TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatus = "healthy"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusDown     TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusInactive, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusDegraded, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusHealthy, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeWARP          TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType = "warp"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeMagic         TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType = "magic"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeCNI           TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeWARP, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeMagic, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelCloudflaredNewResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelCloudflaredNewResponseConfigSrc string

const (
	TunnelCloudflaredNewResponseConfigSrcLocal      TunnelCloudflaredNewResponseConfigSrc = "local"
	TunnelCloudflaredNewResponseConfigSrcCloudflare TunnelCloudflaredNewResponseConfigSrc = "cloudflare"
)

func (r TunnelCloudflaredNewResponseConfigSrc) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewResponseConfigSrcLocal, TunnelCloudflaredNewResponseConfigSrcCloudflare:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredNewResponseStatus string

const (
	TunnelCloudflaredNewResponseStatusInactive TunnelCloudflaredNewResponseStatus = "inactive"
	TunnelCloudflaredNewResponseStatusDegraded TunnelCloudflaredNewResponseStatus = "degraded"
	TunnelCloudflaredNewResponseStatusHealthy  TunnelCloudflaredNewResponseStatus = "healthy"
	TunnelCloudflaredNewResponseStatusDown     TunnelCloudflaredNewResponseStatus = "down"
)

func (r TunnelCloudflaredNewResponseStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewResponseStatusInactive, TunnelCloudflaredNewResponseStatusDegraded, TunnelCloudflaredNewResponseStatusHealthy, TunnelCloudflaredNewResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredNewResponseTunType string

const (
	TunnelCloudflaredNewResponseTunTypeCfdTunnel     TunnelCloudflaredNewResponseTunType = "cfd_tunnel"
	TunnelCloudflaredNewResponseTunTypeWARPConnector TunnelCloudflaredNewResponseTunType = "warp_connector"
	TunnelCloudflaredNewResponseTunTypeWARP          TunnelCloudflaredNewResponseTunType = "warp"
	TunnelCloudflaredNewResponseTunTypeMagic         TunnelCloudflaredNewResponseTunType = "magic"
	TunnelCloudflaredNewResponseTunTypeIPSec         TunnelCloudflaredNewResponseTunType = "ip_sec"
	TunnelCloudflaredNewResponseTunTypeGRE           TunnelCloudflaredNewResponseTunType = "gre"
	TunnelCloudflaredNewResponseTunTypeCNI           TunnelCloudflaredNewResponseTunType = "cni"
)

func (r TunnelCloudflaredNewResponseTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewResponseTunTypeCfdTunnel, TunnelCloudflaredNewResponseTunTypeWARPConnector, TunnelCloudflaredNewResponseTunTypeWARP, TunnelCloudflaredNewResponseTunTypeMagic, TunnelCloudflaredNewResponseTunTypeIPSec, TunnelCloudflaredNewResponseTunTypeGRE, TunnelCloudflaredNewResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredListResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc TunnelCloudflaredListResponseConfigSrc `json:"config_src"`
	// This field can have the runtime type of [[]shared.CloudflareTunnelConnection],
	// [[]TunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnection].
	Connections interface{} `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	//
	// Deprecated: Use the config_src field instead.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredListResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredListResponseTunType `json:"tun_type"`
	JSON    tunnelCloudflaredListResponseJSON    `json:"-"`
	union   TunnelCloudflaredListResponseUnion
}

// tunnelCloudflaredListResponseJSON contains the JSON metadata for the struct
// [TunnelCloudflaredListResponse]
type tunnelCloudflaredListResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	ConfigSrc       apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelCloudflaredListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelCloudflaredListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelCloudflaredListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelCloudflaredListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.CloudflareTunnel],
// [TunnelCloudflaredListResponseTunnelWARPConnectorTunnel].
func (r TunnelCloudflaredListResponse) AsUnion() TunnelCloudflaredListResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [TunnelCloudflaredListResponseTunnelWARPConnectorTunnel].
type TunnelCloudflaredListResponseUnion interface {
	ImplementsTunnelCloudflaredListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelCloudflaredListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelCloudflaredListResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredListResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelCloudflaredListResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelCloudflaredListResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [TunnelCloudflaredListResponseTunnelWARPConnectorTunnel]
type tunnelCloudflaredListResponseTunnelWARPConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TunnelCloudflaredListResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredListResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelCloudflaredListResponseTunnelWARPConnectorTunnel) ImplementsTunnelCloudflaredListResponse() {
}

type TunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                               `json:"uuid" format:"uuid"`
	JSON tunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnectionJSON contains
// the JSON metadata for the struct
// [TunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnection]
type tunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *TunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredListResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatus string

const (
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusInactive TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatus = "inactive"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusDegraded TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatus = "degraded"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusHealthy  TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatus = "healthy"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusDown     TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusInactive, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusDegraded, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusHealthy, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeWARP          TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType = "warp"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeMagic         TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType = "magic"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeCNI           TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeWARP, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeMagic, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelCloudflaredListResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelCloudflaredListResponseConfigSrc string

const (
	TunnelCloudflaredListResponseConfigSrcLocal      TunnelCloudflaredListResponseConfigSrc = "local"
	TunnelCloudflaredListResponseConfigSrcCloudflare TunnelCloudflaredListResponseConfigSrc = "cloudflare"
)

func (r TunnelCloudflaredListResponseConfigSrc) IsKnown() bool {
	switch r {
	case TunnelCloudflaredListResponseConfigSrcLocal, TunnelCloudflaredListResponseConfigSrcCloudflare:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredListResponseStatus string

const (
	TunnelCloudflaredListResponseStatusInactive TunnelCloudflaredListResponseStatus = "inactive"
	TunnelCloudflaredListResponseStatusDegraded TunnelCloudflaredListResponseStatus = "degraded"
	TunnelCloudflaredListResponseStatusHealthy  TunnelCloudflaredListResponseStatus = "healthy"
	TunnelCloudflaredListResponseStatusDown     TunnelCloudflaredListResponseStatus = "down"
)

func (r TunnelCloudflaredListResponseStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredListResponseStatusInactive, TunnelCloudflaredListResponseStatusDegraded, TunnelCloudflaredListResponseStatusHealthy, TunnelCloudflaredListResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredListResponseTunType string

const (
	TunnelCloudflaredListResponseTunTypeCfdTunnel     TunnelCloudflaredListResponseTunType = "cfd_tunnel"
	TunnelCloudflaredListResponseTunTypeWARPConnector TunnelCloudflaredListResponseTunType = "warp_connector"
	TunnelCloudflaredListResponseTunTypeWARP          TunnelCloudflaredListResponseTunType = "warp"
	TunnelCloudflaredListResponseTunTypeMagic         TunnelCloudflaredListResponseTunType = "magic"
	TunnelCloudflaredListResponseTunTypeIPSec         TunnelCloudflaredListResponseTunType = "ip_sec"
	TunnelCloudflaredListResponseTunTypeGRE           TunnelCloudflaredListResponseTunType = "gre"
	TunnelCloudflaredListResponseTunTypeCNI           TunnelCloudflaredListResponseTunType = "cni"
)

func (r TunnelCloudflaredListResponseTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredListResponseTunTypeCfdTunnel, TunnelCloudflaredListResponseTunTypeWARPConnector, TunnelCloudflaredListResponseTunTypeWARP, TunnelCloudflaredListResponseTunTypeMagic, TunnelCloudflaredListResponseTunTypeIPSec, TunnelCloudflaredListResponseTunTypeGRE, TunnelCloudflaredListResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredDeleteResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc TunnelCloudflaredDeleteResponseConfigSrc `json:"config_src"`
	// This field can have the runtime type of [[]shared.CloudflareTunnelConnection],
	// [[]TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnection].
	Connections interface{} `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	//
	// Deprecated: Use the config_src field instead.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredDeleteResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredDeleteResponseTunType `json:"tun_type"`
	JSON    tunnelCloudflaredDeleteResponseJSON    `json:"-"`
	union   TunnelCloudflaredDeleteResponseUnion
}

// tunnelCloudflaredDeleteResponseJSON contains the JSON metadata for the struct
// [TunnelCloudflaredDeleteResponse]
type tunnelCloudflaredDeleteResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	ConfigSrc       apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelCloudflaredDeleteResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelCloudflaredDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelCloudflaredDeleteResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelCloudflaredDeleteResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.CloudflareTunnel],
// [TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnel].
func (r TunnelCloudflaredDeleteResponse) AsUnion() TunnelCloudflaredDeleteResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnel].
type TunnelCloudflaredDeleteResponseUnion interface {
	ImplementsTunnelCloudflaredDeleteResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelCloudflaredDeleteResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct
// [TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnel]
type tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnel) ImplementsTunnelCloudflaredDeleteResponse() {
}

type TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                                 `json:"uuid" format:"uuid"`
	JSON tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnectionJSON contains
// the JSON metadata for the struct
// [TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnection]
type tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatus string

const (
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusInactive TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatus = "inactive"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusDegraded TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatus = "degraded"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusHealthy  TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatus = "healthy"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusDown     TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusInactive, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusDegraded, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusHealthy, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeWARP          TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType = "warp"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeMagic         TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType = "magic"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeCNI           TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeWARP, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeMagic, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelCloudflaredDeleteResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelCloudflaredDeleteResponseConfigSrc string

const (
	TunnelCloudflaredDeleteResponseConfigSrcLocal      TunnelCloudflaredDeleteResponseConfigSrc = "local"
	TunnelCloudflaredDeleteResponseConfigSrcCloudflare TunnelCloudflaredDeleteResponseConfigSrc = "cloudflare"
)

func (r TunnelCloudflaredDeleteResponseConfigSrc) IsKnown() bool {
	switch r {
	case TunnelCloudflaredDeleteResponseConfigSrcLocal, TunnelCloudflaredDeleteResponseConfigSrcCloudflare:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredDeleteResponseStatus string

const (
	TunnelCloudflaredDeleteResponseStatusInactive TunnelCloudflaredDeleteResponseStatus = "inactive"
	TunnelCloudflaredDeleteResponseStatusDegraded TunnelCloudflaredDeleteResponseStatus = "degraded"
	TunnelCloudflaredDeleteResponseStatusHealthy  TunnelCloudflaredDeleteResponseStatus = "healthy"
	TunnelCloudflaredDeleteResponseStatusDown     TunnelCloudflaredDeleteResponseStatus = "down"
)

func (r TunnelCloudflaredDeleteResponseStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredDeleteResponseStatusInactive, TunnelCloudflaredDeleteResponseStatusDegraded, TunnelCloudflaredDeleteResponseStatusHealthy, TunnelCloudflaredDeleteResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredDeleteResponseTunType string

const (
	TunnelCloudflaredDeleteResponseTunTypeCfdTunnel     TunnelCloudflaredDeleteResponseTunType = "cfd_tunnel"
	TunnelCloudflaredDeleteResponseTunTypeWARPConnector TunnelCloudflaredDeleteResponseTunType = "warp_connector"
	TunnelCloudflaredDeleteResponseTunTypeWARP          TunnelCloudflaredDeleteResponseTunType = "warp"
	TunnelCloudflaredDeleteResponseTunTypeMagic         TunnelCloudflaredDeleteResponseTunType = "magic"
	TunnelCloudflaredDeleteResponseTunTypeIPSec         TunnelCloudflaredDeleteResponseTunType = "ip_sec"
	TunnelCloudflaredDeleteResponseTunTypeGRE           TunnelCloudflaredDeleteResponseTunType = "gre"
	TunnelCloudflaredDeleteResponseTunTypeCNI           TunnelCloudflaredDeleteResponseTunType = "cni"
)

func (r TunnelCloudflaredDeleteResponseTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredDeleteResponseTunTypeCfdTunnel, TunnelCloudflaredDeleteResponseTunTypeWARPConnector, TunnelCloudflaredDeleteResponseTunTypeWARP, TunnelCloudflaredDeleteResponseTunTypeMagic, TunnelCloudflaredDeleteResponseTunTypeIPSec, TunnelCloudflaredDeleteResponseTunTypeGRE, TunnelCloudflaredDeleteResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredEditResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc TunnelCloudflaredEditResponseConfigSrc `json:"config_src"`
	// This field can have the runtime type of [[]shared.CloudflareTunnelConnection],
	// [[]TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnection].
	Connections interface{} `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	//
	// Deprecated: Use the config_src field instead.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredEditResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredEditResponseTunType `json:"tun_type"`
	JSON    tunnelCloudflaredEditResponseJSON    `json:"-"`
	union   TunnelCloudflaredEditResponseUnion
}

// tunnelCloudflaredEditResponseJSON contains the JSON metadata for the struct
// [TunnelCloudflaredEditResponse]
type tunnelCloudflaredEditResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	ConfigSrc       apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelCloudflaredEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelCloudflaredEditResponse) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelCloudflaredEditResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelCloudflaredEditResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.CloudflareTunnel],
// [TunnelCloudflaredEditResponseTunnelWARPConnectorTunnel].
func (r TunnelCloudflaredEditResponse) AsUnion() TunnelCloudflaredEditResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [TunnelCloudflaredEditResponseTunnelWARPConnectorTunnel].
type TunnelCloudflaredEditResponseUnion interface {
	ImplementsTunnelCloudflaredEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelCloudflaredEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelCloudflaredEditResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredEditResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [TunnelCloudflaredEditResponseTunnelWARPConnectorTunnel]
type tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TunnelCloudflaredEditResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelCloudflaredEditResponseTunnelWARPConnectorTunnel) ImplementsTunnelCloudflaredEditResponse() {
}

type TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                               `json:"uuid" format:"uuid"`
	JSON tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnectionJSON contains
// the JSON metadata for the struct
// [TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnection]
type tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredEditResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatus string

const (
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusInactive TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatus = "inactive"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusDegraded TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatus = "degraded"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusHealthy  TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatus = "healthy"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusDown     TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusInactive, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusDegraded, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusHealthy, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeWARP          TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType = "warp"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeMagic         TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType = "magic"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeCNI           TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeWARP, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeMagic, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelCloudflaredEditResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelCloudflaredEditResponseConfigSrc string

const (
	TunnelCloudflaredEditResponseConfigSrcLocal      TunnelCloudflaredEditResponseConfigSrc = "local"
	TunnelCloudflaredEditResponseConfigSrcCloudflare TunnelCloudflaredEditResponseConfigSrc = "cloudflare"
)

func (r TunnelCloudflaredEditResponseConfigSrc) IsKnown() bool {
	switch r {
	case TunnelCloudflaredEditResponseConfigSrcLocal, TunnelCloudflaredEditResponseConfigSrcCloudflare:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredEditResponseStatus string

const (
	TunnelCloudflaredEditResponseStatusInactive TunnelCloudflaredEditResponseStatus = "inactive"
	TunnelCloudflaredEditResponseStatusDegraded TunnelCloudflaredEditResponseStatus = "degraded"
	TunnelCloudflaredEditResponseStatusHealthy  TunnelCloudflaredEditResponseStatus = "healthy"
	TunnelCloudflaredEditResponseStatusDown     TunnelCloudflaredEditResponseStatus = "down"
)

func (r TunnelCloudflaredEditResponseStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredEditResponseStatusInactive, TunnelCloudflaredEditResponseStatusDegraded, TunnelCloudflaredEditResponseStatusHealthy, TunnelCloudflaredEditResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredEditResponseTunType string

const (
	TunnelCloudflaredEditResponseTunTypeCfdTunnel     TunnelCloudflaredEditResponseTunType = "cfd_tunnel"
	TunnelCloudflaredEditResponseTunTypeWARPConnector TunnelCloudflaredEditResponseTunType = "warp_connector"
	TunnelCloudflaredEditResponseTunTypeWARP          TunnelCloudflaredEditResponseTunType = "warp"
	TunnelCloudflaredEditResponseTunTypeMagic         TunnelCloudflaredEditResponseTunType = "magic"
	TunnelCloudflaredEditResponseTunTypeIPSec         TunnelCloudflaredEditResponseTunType = "ip_sec"
	TunnelCloudflaredEditResponseTunTypeGRE           TunnelCloudflaredEditResponseTunType = "gre"
	TunnelCloudflaredEditResponseTunTypeCNI           TunnelCloudflaredEditResponseTunType = "cni"
)

func (r TunnelCloudflaredEditResponseTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredEditResponseTunTypeCfdTunnel, TunnelCloudflaredEditResponseTunTypeWARPConnector, TunnelCloudflaredEditResponseTunTypeWARP, TunnelCloudflaredEditResponseTunTypeMagic, TunnelCloudflaredEditResponseTunTypeIPSec, TunnelCloudflaredEditResponseTunTypeGRE, TunnelCloudflaredEditResponseTunTypeCNI:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredGetResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc TunnelCloudflaredGetResponseConfigSrc `json:"config_src"`
	// This field can have the runtime type of [[]shared.CloudflareTunnelConnection],
	// [[]TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnection].
	Connections interface{} `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// This field can have the runtime type of [interface{}].
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	//
	// Deprecated: Use the config_src field instead.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredGetResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredGetResponseTunType `json:"tun_type"`
	JSON    tunnelCloudflaredGetResponseJSON    `json:"-"`
	union   TunnelCloudflaredGetResponseUnion
}

// tunnelCloudflaredGetResponseJSON contains the JSON metadata for the struct
// [TunnelCloudflaredGetResponse]
type tunnelCloudflaredGetResponseJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	ConfigSrc       apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r tunnelCloudflaredGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelCloudflaredGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelCloudflaredGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelCloudflaredGetResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [shared.CloudflareTunnel],
// [TunnelCloudflaredGetResponseTunnelWARPConnectorTunnel].
func (r TunnelCloudflaredGetResponse) AsUnion() TunnelCloudflaredGetResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [TunnelCloudflaredGetResponseTunnelWARPConnectorTunnel].
type TunnelCloudflaredGetResponseUnion interface {
	ImplementsTunnelCloudflaredGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelCloudflaredGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelCloudflaredGetResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelCloudflaredGetResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelJSON contains the JSON
// metadata for the struct [TunnelCloudflaredGetResponseTunnelWARPConnectorTunnel]
type tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *TunnelCloudflaredGetResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelCloudflaredGetResponseTunnelWARPConnectorTunnel) ImplementsTunnelCloudflaredGetResponse() {
}

type TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                              `json:"uuid" format:"uuid"`
	JSON tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnectionJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnection]
type tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredGetResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatus string

const (
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusInactive TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatus = "inactive"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusDegraded TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatus = "degraded"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusHealthy  TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatus = "healthy"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusDown     TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusInactive, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusDegraded, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusHealthy, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeWARP          TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType = "warp"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeMagic         TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType = "magic"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeCNI           TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeWARP, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeMagic, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelCloudflaredGetResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelCloudflaredGetResponseConfigSrc string

const (
	TunnelCloudflaredGetResponseConfigSrcLocal      TunnelCloudflaredGetResponseConfigSrc = "local"
	TunnelCloudflaredGetResponseConfigSrcCloudflare TunnelCloudflaredGetResponseConfigSrc = "cloudflare"
)

func (r TunnelCloudflaredGetResponseConfigSrc) IsKnown() bool {
	switch r {
	case TunnelCloudflaredGetResponseConfigSrcLocal, TunnelCloudflaredGetResponseConfigSrcCloudflare:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredGetResponseStatus string

const (
	TunnelCloudflaredGetResponseStatusInactive TunnelCloudflaredGetResponseStatus = "inactive"
	TunnelCloudflaredGetResponseStatusDegraded TunnelCloudflaredGetResponseStatus = "degraded"
	TunnelCloudflaredGetResponseStatusHealthy  TunnelCloudflaredGetResponseStatus = "healthy"
	TunnelCloudflaredGetResponseStatusDown     TunnelCloudflaredGetResponseStatus = "down"
)

func (r TunnelCloudflaredGetResponseStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredGetResponseStatusInactive, TunnelCloudflaredGetResponseStatusDegraded, TunnelCloudflaredGetResponseStatusHealthy, TunnelCloudflaredGetResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelCloudflaredGetResponseTunType string

const (
	TunnelCloudflaredGetResponseTunTypeCfdTunnel     TunnelCloudflaredGetResponseTunType = "cfd_tunnel"
	TunnelCloudflaredGetResponseTunTypeWARPConnector TunnelCloudflaredGetResponseTunType = "warp_connector"
	TunnelCloudflaredGetResponseTunTypeWARP          TunnelCloudflaredGetResponseTunType = "warp"
	TunnelCloudflaredGetResponseTunTypeMagic         TunnelCloudflaredGetResponseTunType = "magic"
	TunnelCloudflaredGetResponseTunTypeIPSec         TunnelCloudflaredGetResponseTunType = "ip_sec"
	TunnelCloudflaredGetResponseTunTypeGRE           TunnelCloudflaredGetResponseTunType = "gre"
	TunnelCloudflaredGetResponseTunTypeCNI           TunnelCloudflaredGetResponseTunType = "cni"
)

func (r TunnelCloudflaredGetResponseTunType) IsKnown() bool {
	switch r {
	case TunnelCloudflaredGetResponseTunTypeCfdTunnel, TunnelCloudflaredGetResponseTunTypeWARPConnector, TunnelCloudflaredGetResponseTunTypeWARP, TunnelCloudflaredGetResponseTunTypeMagic, TunnelCloudflaredGetResponseTunTypeIPSec, TunnelCloudflaredGetResponseTunTypeGRE, TunnelCloudflaredGetResponseTunTypeCNI:
		return true
	}
	return false
}

type TunnelCloudflaredNewParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name,required"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc param.Field[TunnelCloudflaredNewParamsConfigSrc] `json:"config_src"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r TunnelCloudflaredNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelCloudflaredNewParamsConfigSrc string

const (
	TunnelCloudflaredNewParamsConfigSrcLocal      TunnelCloudflaredNewParamsConfigSrc = "local"
	TunnelCloudflaredNewParamsConfigSrcCloudflare TunnelCloudflaredNewParamsConfigSrc = "cloudflare"
)

func (r TunnelCloudflaredNewParamsConfigSrc) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewParamsConfigSrcLocal, TunnelCloudflaredNewParamsConfigSrcCloudflare:
		return true
	}
	return false
}

type TunnelCloudflaredNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelCloudflaredNewResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredNewResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredNewResponseEnvelope]
type tunnelCloudflaredNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredNewResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredNewResponseEnvelopeSuccessTrue TunnelCloudflaredNewResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredListParams struct {
	// Cloudflare account ID
	AccountID     param.Field[string] `path:"account_id,required"`
	ExcludePrefix param.Field[string] `query:"exclude_prefix"`
	// If provided, include only resources that were created (and not deleted) before
	// this time. URL encoded.
	ExistedAt     param.Field[string] `query:"existed_at" format:"url-encoded-date-time"`
	IncludePrefix param.Field[string] `query:"include_prefix"`
	// If `true`, only include deleted tunnels. If `false`, exclude deleted tunnels. If
	// empty, all tunnels will be included.
	IsDeleted param.Field[bool] `query:"is_deleted"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status param.Field[TunnelCloudflaredListParamsStatus] `query:"status"`
	// UUID of the tunnel.
	UUID          param.Field[string]    `query:"uuid" format:"uuid"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [TunnelCloudflaredListParams]'s query parameters as
// `url.Values`.
func (r TunnelCloudflaredListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelCloudflaredListParamsStatus string

const (
	TunnelCloudflaredListParamsStatusInactive TunnelCloudflaredListParamsStatus = "inactive"
	TunnelCloudflaredListParamsStatusDegraded TunnelCloudflaredListParamsStatus = "degraded"
	TunnelCloudflaredListParamsStatusHealthy  TunnelCloudflaredListParamsStatus = "healthy"
	TunnelCloudflaredListParamsStatusDown     TunnelCloudflaredListParamsStatus = "down"
)

func (r TunnelCloudflaredListParamsStatus) IsKnown() bool {
	switch r {
	case TunnelCloudflaredListParamsStatusInactive, TunnelCloudflaredListParamsStatusDegraded, TunnelCloudflaredListParamsStatusHealthy, TunnelCloudflaredListParamsStatusDown:
		return true
	}
	return false
}

type TunnelCloudflaredDeleteParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelCloudflaredDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelCloudflaredDeleteResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredDeleteResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredDeleteResponseEnvelope]
type tunnelCloudflaredDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredDeleteResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredDeleteResponseEnvelopeSuccessTrue TunnelCloudflaredDeleteResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredEditParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// A user-friendly name for a tunnel.
	Name param.Field[string] `json:"name"`
	// Sets the password required to run a locally-managed tunnel. Must be at least 32
	// bytes and encoded as a base64 string.
	TunnelSecret param.Field[string] `json:"tunnel_secret"`
}

func (r TunnelCloudflaredEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelCloudflaredEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelCloudflaredEditResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredEditResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredEditResponseEnvelope]
type tunnelCloudflaredEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredEditResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredEditResponseEnvelopeSuccessTrue TunnelCloudflaredEditResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredGetParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelCloudflaredGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
	Result TunnelCloudflaredGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success TunnelCloudflaredGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    tunnelCloudflaredGetResponseEnvelopeJSON    `json:"-"`
}

// tunnelCloudflaredGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TunnelCloudflaredGetResponseEnvelope]
type tunnelCloudflaredGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TunnelCloudflaredGetResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredGetResponseEnvelopeSuccessTrue TunnelCloudflaredGetResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
