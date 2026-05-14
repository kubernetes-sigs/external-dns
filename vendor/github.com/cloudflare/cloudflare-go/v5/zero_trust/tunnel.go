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

// TunnelService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelService] method instead.
type TunnelService struct {
	Options       []option.RequestOption
	Cloudflared   *TunnelCloudflaredService
	WARPConnector *TunnelWARPConnectorService
}

// NewTunnelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTunnelService(opts ...option.RequestOption) (r *TunnelService) {
	r = &TunnelService{}
	r.Options = opts
	r.Cloudflared = NewTunnelCloudflaredService(opts...)
	r.WARPConnector = NewTunnelWARPConnectorService(opts...)
	return
}

// Lists and filters all types of Tunnels in an account.
func (r *TunnelService) List(ctx context.Context, params TunnelListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[TunnelListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tunnels", params.AccountID)
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

// Lists and filters all types of Tunnels in an account.
func (r *TunnelService) ListAutoPaging(ctx context.Context, params TunnelListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[TunnelListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type TunnelListResponse struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc TunnelListResponseConfigSrc `json:"config_src"`
	// This field can have the runtime type of [[]shared.CloudflareTunnelConnection],
	// [[]TunnelListResponseTunnelWARPConnectorTunnelConnection].
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
	Status TunnelListResponseStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelListResponseTunType `json:"tun_type"`
	JSON    tunnelListResponseJSON    `json:"-"`
	union   TunnelListResponseUnion
}

// tunnelListResponseJSON contains the JSON metadata for the struct
// [TunnelListResponse]
type tunnelListResponseJSON struct {
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

func (r tunnelListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *TunnelListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = TunnelListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [TunnelListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [shared.CloudflareTunnel],
// [TunnelListResponseTunnelWARPConnectorTunnel].
func (r TunnelListResponse) AsUnion() TunnelListResponseUnion {
	return r.union
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// Union satisfied by [shared.CloudflareTunnel] or
// [TunnelListResponseTunnelWARPConnectorTunnel].
type TunnelListResponseUnion interface {
	ImplementsTunnelListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TunnelListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(shared.CloudflareTunnel{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TunnelListResponseTunnelWARPConnectorTunnel{}),
		},
	)
}

// A Warp Connector Tunnel that connects your origin to Cloudflare's edge.
type TunnelListResponseTunnelWARPConnectorTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []TunnelListResponseTunnelWARPConnectorTunnelConnection `json:"connections"`
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
	Status TunnelListResponseTunnelWARPConnectorTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType TunnelListResponseTunnelWARPConnectorTunnelTunType `json:"tun_type"`
	JSON    tunnelListResponseTunnelWARPConnectorTunnelJSON    `json:"-"`
}

// tunnelListResponseTunnelWARPConnectorTunnelJSON contains the JSON metadata for
// the struct [TunnelListResponseTunnelWARPConnectorTunnel]
type tunnelListResponseTunnelWARPConnectorTunnelJSON struct {
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

func (r *TunnelListResponseTunnelWARPConnectorTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelListResponseTunnelWARPConnectorTunnelJSON) RawJSON() string {
	return r.raw
}

func (r TunnelListResponseTunnelWARPConnectorTunnel) ImplementsTunnelListResponse() {}

type TunnelListResponseTunnelWARPConnectorTunnelConnection struct {
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
	UUID string                                                    `json:"uuid" format:"uuid"`
	JSON tunnelListResponseTunnelWARPConnectorTunnelConnectionJSON `json:"-"`
}

// tunnelListResponseTunnelWARPConnectorTunnelConnectionJSON contains the JSON
// metadata for the struct [TunnelListResponseTunnelWARPConnectorTunnelConnection]
type tunnelListResponseTunnelWARPConnectorTunnelConnectionJSON struct {
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

func (r *TunnelListResponseTunnelWARPConnectorTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelListResponseTunnelWARPConnectorTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelListResponseTunnelWARPConnectorTunnelStatus string

const (
	TunnelListResponseTunnelWARPConnectorTunnelStatusInactive TunnelListResponseTunnelWARPConnectorTunnelStatus = "inactive"
	TunnelListResponseTunnelWARPConnectorTunnelStatusDegraded TunnelListResponseTunnelWARPConnectorTunnelStatus = "degraded"
	TunnelListResponseTunnelWARPConnectorTunnelStatusHealthy  TunnelListResponseTunnelWARPConnectorTunnelStatus = "healthy"
	TunnelListResponseTunnelWARPConnectorTunnelStatusDown     TunnelListResponseTunnelWARPConnectorTunnelStatus = "down"
)

func (r TunnelListResponseTunnelWARPConnectorTunnelStatus) IsKnown() bool {
	switch r {
	case TunnelListResponseTunnelWARPConnectorTunnelStatusInactive, TunnelListResponseTunnelWARPConnectorTunnelStatusDegraded, TunnelListResponseTunnelWARPConnectorTunnelStatusHealthy, TunnelListResponseTunnelWARPConnectorTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelListResponseTunnelWARPConnectorTunnelTunType string

const (
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel     TunnelListResponseTunnelWARPConnectorTunnelTunType = "cfd_tunnel"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector TunnelListResponseTunnelWARPConnectorTunnelTunType = "warp_connector"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeWARP          TunnelListResponseTunnelWARPConnectorTunnelTunType = "warp"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeMagic         TunnelListResponseTunnelWARPConnectorTunnelTunType = "magic"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeIPSec         TunnelListResponseTunnelWARPConnectorTunnelTunType = "ip_sec"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeGRE           TunnelListResponseTunnelWARPConnectorTunnelTunType = "gre"
	TunnelListResponseTunnelWARPConnectorTunnelTunTypeCNI           TunnelListResponseTunnelWARPConnectorTunnelTunType = "cni"
)

func (r TunnelListResponseTunnelWARPConnectorTunnelTunType) IsKnown() bool {
	switch r {
	case TunnelListResponseTunnelWARPConnectorTunnelTunTypeCfdTunnel, TunnelListResponseTunnelWARPConnectorTunnelTunTypeWARPConnector, TunnelListResponseTunnelWARPConnectorTunnelTunTypeWARP, TunnelListResponseTunnelWARPConnectorTunnelTunTypeMagic, TunnelListResponseTunnelWARPConnectorTunnelTunTypeIPSec, TunnelListResponseTunnelWARPConnectorTunnelTunTypeGRE, TunnelListResponseTunnelWARPConnectorTunnelTunTypeCNI:
		return true
	}
	return false
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type TunnelListResponseConfigSrc string

const (
	TunnelListResponseConfigSrcLocal      TunnelListResponseConfigSrc = "local"
	TunnelListResponseConfigSrcCloudflare TunnelListResponseConfigSrc = "cloudflare"
)

func (r TunnelListResponseConfigSrc) IsKnown() bool {
	switch r {
	case TunnelListResponseConfigSrcLocal, TunnelListResponseConfigSrcCloudflare:
		return true
	}
	return false
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelListResponseStatus string

const (
	TunnelListResponseStatusInactive TunnelListResponseStatus = "inactive"
	TunnelListResponseStatusDegraded TunnelListResponseStatus = "degraded"
	TunnelListResponseStatusHealthy  TunnelListResponseStatus = "healthy"
	TunnelListResponseStatusDown     TunnelListResponseStatus = "down"
)

func (r TunnelListResponseStatus) IsKnown() bool {
	switch r {
	case TunnelListResponseStatusInactive, TunnelListResponseStatusDegraded, TunnelListResponseStatusHealthy, TunnelListResponseStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelListResponseTunType string

const (
	TunnelListResponseTunTypeCfdTunnel     TunnelListResponseTunType = "cfd_tunnel"
	TunnelListResponseTunTypeWARPConnector TunnelListResponseTunType = "warp_connector"
	TunnelListResponseTunTypeWARP          TunnelListResponseTunType = "warp"
	TunnelListResponseTunTypeMagic         TunnelListResponseTunType = "magic"
	TunnelListResponseTunTypeIPSec         TunnelListResponseTunType = "ip_sec"
	TunnelListResponseTunTypeGRE           TunnelListResponseTunType = "gre"
	TunnelListResponseTunTypeCNI           TunnelListResponseTunType = "cni"
)

func (r TunnelListResponseTunType) IsKnown() bool {
	switch r {
	case TunnelListResponseTunTypeCfdTunnel, TunnelListResponseTunTypeWARPConnector, TunnelListResponseTunTypeWARP, TunnelListResponseTunTypeMagic, TunnelListResponseTunTypeIPSec, TunnelListResponseTunTypeGRE, TunnelListResponseTunTypeCNI:
		return true
	}
	return false
}

type TunnelListParams struct {
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
	// A user-friendly name for the tunnel.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of results to display.
	PerPage param.Field[float64] `query:"per_page"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status param.Field[TunnelListParamsStatus] `query:"status"`
	// The types of tunnels to filter by, separated by commas.
	TunTypes param.Field[[]TunnelListParamsTunType] `query:"tun_types"`
	// UUID of the tunnel.
	UUID          param.Field[string]    `query:"uuid" format:"uuid"`
	WasActiveAt   param.Field[time.Time] `query:"was_active_at" format:"date-time"`
	WasInactiveAt param.Field[time.Time] `query:"was_inactive_at" format:"date-time"`
}

// URLQuery serializes [TunnelListParams]'s query parameters as `url.Values`.
func (r TunnelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type TunnelListParamsStatus string

const (
	TunnelListParamsStatusInactive TunnelListParamsStatus = "inactive"
	TunnelListParamsStatusDegraded TunnelListParamsStatus = "degraded"
	TunnelListParamsStatusHealthy  TunnelListParamsStatus = "healthy"
	TunnelListParamsStatusDown     TunnelListParamsStatus = "down"
)

func (r TunnelListParamsStatus) IsKnown() bool {
	switch r {
	case TunnelListParamsStatusInactive, TunnelListParamsStatusDegraded, TunnelListParamsStatusHealthy, TunnelListParamsStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type TunnelListParamsTunType string

const (
	TunnelListParamsTunTypeCfdTunnel     TunnelListParamsTunType = "cfd_tunnel"
	TunnelListParamsTunTypeWARPConnector TunnelListParamsTunType = "warp_connector"
	TunnelListParamsTunTypeWARP          TunnelListParamsTunType = "warp"
	TunnelListParamsTunTypeMagic         TunnelListParamsTunType = "magic"
	TunnelListParamsTunTypeIPSec         TunnelListParamsTunType = "ip_sec"
	TunnelListParamsTunTypeGRE           TunnelListParamsTunType = "gre"
	TunnelListParamsTunTypeCNI           TunnelListParamsTunType = "cni"
)

func (r TunnelListParamsTunType) IsKnown() bool {
	switch r {
	case TunnelListParamsTunTypeCfdTunnel, TunnelListParamsTunTypeWARPConnector, TunnelListParamsTunTypeWARP, TunnelListParamsTunTypeMagic, TunnelListParamsTunTypeIPSec, TunnelListParamsTunTypeGRE, TunnelListParamsTunTypeCNI:
		return true
	}
	return false
}
