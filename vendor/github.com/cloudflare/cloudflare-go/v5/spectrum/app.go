// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

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
	"github.com/tidwall/gjson"
)

// AppService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppService] method instead.
type AppService struct {
	Options []option.RequestOption
}

// NewAppService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAppService(opts ...option.RequestOption) (r *AppService) {
	r = &AppService{}
	r.Options = opts
	return
}

// Creates a new Spectrum application from a configuration using a name for the
// origin.
func (r *AppService) New(ctx context.Context, params AppNewParams, opts ...option.RequestOption) (res *AppNewResponse, err error) {
	var env AppNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a previously existing application's configuration that uses a name for
// the origin.
func (r *AppService) Update(ctx context.Context, appID string, params AppUpdateParams, opts ...option.RequestOption) (res *AppUpdateResponse, err error) {
	var env AppUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", params.ZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *AppService) List(ctx context.Context, params AppListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AppListResponseUnion], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps", params.ZoneID)
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

// Retrieves a list of currently existing Spectrum applications inside a zone.
func (r *AppService) ListAutoPaging(ctx context.Context, params AppListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AppListResponseUnion] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a previously existing application.
func (r *AppService) Delete(ctx context.Context, appID string, body AppDeleteParams, opts ...option.RequestOption) (res *AppDeleteResponse, err error) {
	var env AppDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", body.ZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the application configuration of a specific application inside a zone.
func (r *AppService) Get(ctx context.Context, appID string, query AppGetParams, opts ...option.RequestOption) (res *AppGetResponse, err error) {
	var env AppGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/apps/%s", query.ZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AppNewResponse struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs EdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// This field can have the runtime type of [[]string].
	OriginDirect interface{} `json:"origin_direct"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS OriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort OriginPortUnion `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppNewResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS AppNewResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppNewResponseTrafficType `json:"traffic_type"`
	JSON        appNewResponseJSON        `json:"-"`
	union       AppNewResponseUnion
}

// appNewResponseJSON contains the JSON metadata for the struct [AppNewResponse]
type appNewResponseJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r appNewResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AppNewResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AppNewResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppNewResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AppNewResponseSpectrumConfigAppConfig],
// [AppNewResponseSpectrumConfigPaygoAppConfig].
func (r AppNewResponse) AsUnion() AppNewResponseUnion {
	return r.union
}

// Union satisfied by [AppNewResponseSpectrumConfigAppConfig] or
// [AppNewResponseSpectrumConfigPaygoAppConfig].
type AppNewResponseUnion interface {
	implementsAppNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppNewResponseSpectrumConfigAppConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppNewResponseSpectrumConfigPaygoAppConfig{}),
		},
	)
}

type AppNewResponseSpectrumConfigAppConfig struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppNewResponseSpectrumConfigAppConfigTrafficType `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs EdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS OriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort OriginPortUnion `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppNewResponseSpectrumConfigAppConfigProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS  AppNewResponseSpectrumConfigAppConfigTLS  `json:"tls"`
	JSON appNewResponseSpectrumConfigAppConfigJSON `json:"-"`
}

// appNewResponseSpectrumConfigAppConfigJSON contains the JSON metadata for the
// struct [AppNewResponseSpectrumConfigAppConfig]
type appNewResponseSpectrumConfigAppConfigJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	TrafficType      apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppNewResponseSpectrumConfigAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseSpectrumConfigAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r AppNewResponseSpectrumConfigAppConfig) implementsAppNewResponse() {}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewResponseSpectrumConfigAppConfigTrafficType string

const (
	AppNewResponseSpectrumConfigAppConfigTrafficTypeDirect AppNewResponseSpectrumConfigAppConfigTrafficType = "direct"
	AppNewResponseSpectrumConfigAppConfigTrafficTypeHTTP   AppNewResponseSpectrumConfigAppConfigTrafficType = "http"
	AppNewResponseSpectrumConfigAppConfigTrafficTypeHTTPS  AppNewResponseSpectrumConfigAppConfigTrafficType = "https"
)

func (r AppNewResponseSpectrumConfigAppConfigTrafficType) IsKnown() bool {
	switch r {
	case AppNewResponseSpectrumConfigAppConfigTrafficTypeDirect, AppNewResponseSpectrumConfigAppConfigTrafficTypeHTTP, AppNewResponseSpectrumConfigAppConfigTrafficTypeHTTPS:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppNewResponseSpectrumConfigAppConfigProxyProtocol string

const (
	AppNewResponseSpectrumConfigAppConfigProxyProtocolOff    AppNewResponseSpectrumConfigAppConfigProxyProtocol = "off"
	AppNewResponseSpectrumConfigAppConfigProxyProtocolV1     AppNewResponseSpectrumConfigAppConfigProxyProtocol = "v1"
	AppNewResponseSpectrumConfigAppConfigProxyProtocolV2     AppNewResponseSpectrumConfigAppConfigProxyProtocol = "v2"
	AppNewResponseSpectrumConfigAppConfigProxyProtocolSimple AppNewResponseSpectrumConfigAppConfigProxyProtocol = "simple"
)

func (r AppNewResponseSpectrumConfigAppConfigProxyProtocol) IsKnown() bool {
	switch r {
	case AppNewResponseSpectrumConfigAppConfigProxyProtocolOff, AppNewResponseSpectrumConfigAppConfigProxyProtocolV1, AppNewResponseSpectrumConfigAppConfigProxyProtocolV2, AppNewResponseSpectrumConfigAppConfigProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppNewResponseSpectrumConfigAppConfigTLS string

const (
	AppNewResponseSpectrumConfigAppConfigTLSOff      AppNewResponseSpectrumConfigAppConfigTLS = "off"
	AppNewResponseSpectrumConfigAppConfigTLSFlexible AppNewResponseSpectrumConfigAppConfigTLS = "flexible"
	AppNewResponseSpectrumConfigAppConfigTLSFull     AppNewResponseSpectrumConfigAppConfigTLS = "full"
	AppNewResponseSpectrumConfigAppConfigTLSStrict   AppNewResponseSpectrumConfigAppConfigTLS = "strict"
)

func (r AppNewResponseSpectrumConfigAppConfigTLS) IsKnown() bool {
	switch r {
	case AppNewResponseSpectrumConfigAppConfigTLSOff, AppNewResponseSpectrumConfigAppConfigTLSFlexible, AppNewResponseSpectrumConfigAppConfigTLSFull, AppNewResponseSpectrumConfigAppConfigTLSStrict:
		return true
	}
	return false
}

type AppNewResponseSpectrumConfigPaygoAppConfig struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string                                       `json:"origin_direct" format:"URI"`
	JSON         appNewResponseSpectrumConfigPaygoAppConfigJSON `json:"-"`
}

// appNewResponseSpectrumConfigPaygoAppConfigJSON contains the JSON metadata for
// the struct [AppNewResponseSpectrumConfigPaygoAppConfig]
type appNewResponseSpectrumConfigPaygoAppConfigJSON struct {
	ID           apijson.Field
	CreatedOn    apijson.Field
	DNS          apijson.Field
	ModifiedOn   apijson.Field
	Protocol     apijson.Field
	OriginDirect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppNewResponseSpectrumConfigPaygoAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseSpectrumConfigPaygoAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r AppNewResponseSpectrumConfigPaygoAppConfig) implementsAppNewResponse() {}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppNewResponseProxyProtocol string

const (
	AppNewResponseProxyProtocolOff    AppNewResponseProxyProtocol = "off"
	AppNewResponseProxyProtocolV1     AppNewResponseProxyProtocol = "v1"
	AppNewResponseProxyProtocolV2     AppNewResponseProxyProtocol = "v2"
	AppNewResponseProxyProtocolSimple AppNewResponseProxyProtocol = "simple"
)

func (r AppNewResponseProxyProtocol) IsKnown() bool {
	switch r {
	case AppNewResponseProxyProtocolOff, AppNewResponseProxyProtocolV1, AppNewResponseProxyProtocolV2, AppNewResponseProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppNewResponseTLS string

const (
	AppNewResponseTLSOff      AppNewResponseTLS = "off"
	AppNewResponseTLSFlexible AppNewResponseTLS = "flexible"
	AppNewResponseTLSFull     AppNewResponseTLS = "full"
	AppNewResponseTLSStrict   AppNewResponseTLS = "strict"
)

func (r AppNewResponseTLS) IsKnown() bool {
	switch r {
	case AppNewResponseTLSOff, AppNewResponseTLSFlexible, AppNewResponseTLSFull, AppNewResponseTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewResponseTrafficType string

const (
	AppNewResponseTrafficTypeDirect AppNewResponseTrafficType = "direct"
	AppNewResponseTrafficTypeHTTP   AppNewResponseTrafficType = "http"
	AppNewResponseTrafficTypeHTTPS  AppNewResponseTrafficType = "https"
)

func (r AppNewResponseTrafficType) IsKnown() bool {
	switch r {
	case AppNewResponseTrafficTypeDirect, AppNewResponseTrafficTypeHTTP, AppNewResponseTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppUpdateResponse struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs EdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// This field can have the runtime type of [[]string].
	OriginDirect interface{} `json:"origin_direct"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS OriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort OriginPortUnion `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppUpdateResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS AppUpdateResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppUpdateResponseTrafficType `json:"traffic_type"`
	JSON        appUpdateResponseJSON        `json:"-"`
	union       AppUpdateResponseUnion
}

// appUpdateResponseJSON contains the JSON metadata for the struct
// [AppUpdateResponse]
type appUpdateResponseJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r appUpdateResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AppUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AppUpdateResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppUpdateResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are
// [AppUpdateResponseSpectrumConfigAppConfig],
// [AppUpdateResponseSpectrumConfigPaygoAppConfig].
func (r AppUpdateResponse) AsUnion() AppUpdateResponseUnion {
	return r.union
}

// Union satisfied by [AppUpdateResponseSpectrumConfigAppConfig] or
// [AppUpdateResponseSpectrumConfigPaygoAppConfig].
type AppUpdateResponseUnion interface {
	implementsAppUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppUpdateResponseSpectrumConfigAppConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppUpdateResponseSpectrumConfigPaygoAppConfig{}),
		},
	)
}

type AppUpdateResponseSpectrumConfigAppConfig struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppUpdateResponseSpectrumConfigAppConfigTrafficType `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs EdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS OriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort OriginPortUnion `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppUpdateResponseSpectrumConfigAppConfigProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS  AppUpdateResponseSpectrumConfigAppConfigTLS  `json:"tls"`
	JSON appUpdateResponseSpectrumConfigAppConfigJSON `json:"-"`
}

// appUpdateResponseSpectrumConfigAppConfigJSON contains the JSON metadata for the
// struct [AppUpdateResponseSpectrumConfigAppConfig]
type appUpdateResponseSpectrumConfigAppConfigJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	TrafficType      apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppUpdateResponseSpectrumConfigAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseSpectrumConfigAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r AppUpdateResponseSpectrumConfigAppConfig) implementsAppUpdateResponse() {}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateResponseSpectrumConfigAppConfigTrafficType string

const (
	AppUpdateResponseSpectrumConfigAppConfigTrafficTypeDirect AppUpdateResponseSpectrumConfigAppConfigTrafficType = "direct"
	AppUpdateResponseSpectrumConfigAppConfigTrafficTypeHTTP   AppUpdateResponseSpectrumConfigAppConfigTrafficType = "http"
	AppUpdateResponseSpectrumConfigAppConfigTrafficTypeHTTPS  AppUpdateResponseSpectrumConfigAppConfigTrafficType = "https"
)

func (r AppUpdateResponseSpectrumConfigAppConfigTrafficType) IsKnown() bool {
	switch r {
	case AppUpdateResponseSpectrumConfigAppConfigTrafficTypeDirect, AppUpdateResponseSpectrumConfigAppConfigTrafficTypeHTTP, AppUpdateResponseSpectrumConfigAppConfigTrafficTypeHTTPS:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppUpdateResponseSpectrumConfigAppConfigProxyProtocol string

const (
	AppUpdateResponseSpectrumConfigAppConfigProxyProtocolOff    AppUpdateResponseSpectrumConfigAppConfigProxyProtocol = "off"
	AppUpdateResponseSpectrumConfigAppConfigProxyProtocolV1     AppUpdateResponseSpectrumConfigAppConfigProxyProtocol = "v1"
	AppUpdateResponseSpectrumConfigAppConfigProxyProtocolV2     AppUpdateResponseSpectrumConfigAppConfigProxyProtocol = "v2"
	AppUpdateResponseSpectrumConfigAppConfigProxyProtocolSimple AppUpdateResponseSpectrumConfigAppConfigProxyProtocol = "simple"
)

func (r AppUpdateResponseSpectrumConfigAppConfigProxyProtocol) IsKnown() bool {
	switch r {
	case AppUpdateResponseSpectrumConfigAppConfigProxyProtocolOff, AppUpdateResponseSpectrumConfigAppConfigProxyProtocolV1, AppUpdateResponseSpectrumConfigAppConfigProxyProtocolV2, AppUpdateResponseSpectrumConfigAppConfigProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppUpdateResponseSpectrumConfigAppConfigTLS string

const (
	AppUpdateResponseSpectrumConfigAppConfigTLSOff      AppUpdateResponseSpectrumConfigAppConfigTLS = "off"
	AppUpdateResponseSpectrumConfigAppConfigTLSFlexible AppUpdateResponseSpectrumConfigAppConfigTLS = "flexible"
	AppUpdateResponseSpectrumConfigAppConfigTLSFull     AppUpdateResponseSpectrumConfigAppConfigTLS = "full"
	AppUpdateResponseSpectrumConfigAppConfigTLSStrict   AppUpdateResponseSpectrumConfigAppConfigTLS = "strict"
)

func (r AppUpdateResponseSpectrumConfigAppConfigTLS) IsKnown() bool {
	switch r {
	case AppUpdateResponseSpectrumConfigAppConfigTLSOff, AppUpdateResponseSpectrumConfigAppConfigTLSFlexible, AppUpdateResponseSpectrumConfigAppConfigTLSFull, AppUpdateResponseSpectrumConfigAppConfigTLSStrict:
		return true
	}
	return false
}

type AppUpdateResponseSpectrumConfigPaygoAppConfig struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string                                          `json:"origin_direct" format:"URI"`
	JSON         appUpdateResponseSpectrumConfigPaygoAppConfigJSON `json:"-"`
}

// appUpdateResponseSpectrumConfigPaygoAppConfigJSON contains the JSON metadata for
// the struct [AppUpdateResponseSpectrumConfigPaygoAppConfig]
type appUpdateResponseSpectrumConfigPaygoAppConfigJSON struct {
	ID           apijson.Field
	CreatedOn    apijson.Field
	DNS          apijson.Field
	ModifiedOn   apijson.Field
	Protocol     apijson.Field
	OriginDirect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppUpdateResponseSpectrumConfigPaygoAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseSpectrumConfigPaygoAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r AppUpdateResponseSpectrumConfigPaygoAppConfig) implementsAppUpdateResponse() {}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppUpdateResponseProxyProtocol string

const (
	AppUpdateResponseProxyProtocolOff    AppUpdateResponseProxyProtocol = "off"
	AppUpdateResponseProxyProtocolV1     AppUpdateResponseProxyProtocol = "v1"
	AppUpdateResponseProxyProtocolV2     AppUpdateResponseProxyProtocol = "v2"
	AppUpdateResponseProxyProtocolSimple AppUpdateResponseProxyProtocol = "simple"
)

func (r AppUpdateResponseProxyProtocol) IsKnown() bool {
	switch r {
	case AppUpdateResponseProxyProtocolOff, AppUpdateResponseProxyProtocolV1, AppUpdateResponseProxyProtocolV2, AppUpdateResponseProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppUpdateResponseTLS string

const (
	AppUpdateResponseTLSOff      AppUpdateResponseTLS = "off"
	AppUpdateResponseTLSFlexible AppUpdateResponseTLS = "flexible"
	AppUpdateResponseTLSFull     AppUpdateResponseTLS = "full"
	AppUpdateResponseTLSStrict   AppUpdateResponseTLS = "strict"
)

func (r AppUpdateResponseTLS) IsKnown() bool {
	switch r {
	case AppUpdateResponseTLSOff, AppUpdateResponseTLSFlexible, AppUpdateResponseTLSFull, AppUpdateResponseTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateResponseTrafficType string

const (
	AppUpdateResponseTrafficTypeDirect AppUpdateResponseTrafficType = "direct"
	AppUpdateResponseTrafficTypeHTTP   AppUpdateResponseTrafficType = "http"
	AppUpdateResponseTrafficTypeHTTPS  AppUpdateResponseTrafficType = "https"
)

func (r AppUpdateResponseTrafficType) IsKnown() bool {
	switch r {
	case AppUpdateResponseTrafficTypeDirect, AppUpdateResponseTrafficTypeHTTP, AppUpdateResponseTrafficTypeHTTPS:
		return true
	}
	return false
}

// Union satisfied by [AppListResponseArray] or [AppListResponseArray].
type AppListResponseUnion interface {
	implementsAppListResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppListResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppListResponseArray{}),
		},
	)
}

type AppListResponseArray []AppListResponseArrayItem

func (r AppListResponseArray) implementsAppListResponseUnion() {}

type AppListResponseArrayItem struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppListResponseArrayTrafficType `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs EdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS OriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort OriginPortUnion `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppListResponseArrayProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS  AppListResponseArrayTLS      `json:"tls"`
	JSON appListResponseArrayItemJSON `json:"-"`
}

// appListResponseArrayItemJSON contains the JSON metadata for the struct
// [AppListResponseArrayItem]
type appListResponseArrayItemJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	TrafficType      apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppListResponseArrayItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appListResponseArrayItemJSON) RawJSON() string {
	return r.raw
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppListResponseArrayTrafficType string

const (
	AppListResponseArrayTrafficTypeDirect AppListResponseArrayTrafficType = "direct"
	AppListResponseArrayTrafficTypeHTTP   AppListResponseArrayTrafficType = "http"
	AppListResponseArrayTrafficTypeHTTPS  AppListResponseArrayTrafficType = "https"
)

func (r AppListResponseArrayTrafficType) IsKnown() bool {
	switch r {
	case AppListResponseArrayTrafficTypeDirect, AppListResponseArrayTrafficTypeHTTP, AppListResponseArrayTrafficTypeHTTPS:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppListResponseArrayProxyProtocol string

const (
	AppListResponseArrayProxyProtocolOff    AppListResponseArrayProxyProtocol = "off"
	AppListResponseArrayProxyProtocolV1     AppListResponseArrayProxyProtocol = "v1"
	AppListResponseArrayProxyProtocolV2     AppListResponseArrayProxyProtocol = "v2"
	AppListResponseArrayProxyProtocolSimple AppListResponseArrayProxyProtocol = "simple"
)

func (r AppListResponseArrayProxyProtocol) IsKnown() bool {
	switch r {
	case AppListResponseArrayProxyProtocolOff, AppListResponseArrayProxyProtocolV1, AppListResponseArrayProxyProtocolV2, AppListResponseArrayProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppListResponseArrayTLS string

const (
	AppListResponseArrayTLSOff      AppListResponseArrayTLS = "off"
	AppListResponseArrayTLSFlexible AppListResponseArrayTLS = "flexible"
	AppListResponseArrayTLSFull     AppListResponseArrayTLS = "full"
	AppListResponseArrayTLSStrict   AppListResponseArrayTLS = "strict"
)

func (r AppListResponseArrayTLS) IsKnown() bool {
	switch r {
	case AppListResponseArrayTLSOff, AppListResponseArrayTLSFlexible, AppListResponseArrayTLSFull, AppListResponseArrayTLSStrict:
		return true
	}
	return false
}

type AppDeleteResponse struct {
	// Identifier.
	ID   string                `json:"id,required"`
	JSON appDeleteResponseJSON `json:"-"`
}

// appDeleteResponseJSON contains the JSON metadata for the struct
// [AppDeleteResponse]
type appDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AppGetResponse struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs EdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// This field can have the runtime type of [[]string].
	OriginDirect interface{} `json:"origin_direct"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS OriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort OriginPortUnion `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppGetResponseProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS AppGetResponseTLS `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppGetResponseTrafficType `json:"traffic_type"`
	JSON        appGetResponseJSON        `json:"-"`
	union       AppGetResponseUnion
}

// appGetResponseJSON contains the JSON metadata for the struct [AppGetResponse]
type appGetResponseJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	TrafficType      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r appGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *AppGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = AppGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AppGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [AppGetResponseSpectrumConfigAppConfig],
// [AppGetResponseSpectrumConfigPaygoAppConfig].
func (r AppGetResponse) AsUnion() AppGetResponseUnion {
	return r.union
}

// Union satisfied by [AppGetResponseSpectrumConfigAppConfig] or
// [AppGetResponseSpectrumConfigPaygoAppConfig].
type AppGetResponseUnion interface {
	implementsAppGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppGetResponseSpectrumConfigAppConfig{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppGetResponseSpectrumConfigPaygoAppConfig{}),
		},
	)
}

type AppGetResponseSpectrumConfigAppConfig struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType AppGetResponseSpectrumConfigAppConfigTrafficType `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting bool `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs EdgeIPs `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall bool `json:"ip_firewall"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS OriginDNS `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort OriginPortUnion `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol AppGetResponseSpectrumConfigAppConfigProxyProtocol `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS  AppGetResponseSpectrumConfigAppConfigTLS  `json:"tls"`
	JSON appGetResponseSpectrumConfigAppConfigJSON `json:"-"`
}

// appGetResponseSpectrumConfigAppConfigJSON contains the JSON metadata for the
// struct [AppGetResponseSpectrumConfigAppConfig]
type appGetResponseSpectrumConfigAppConfigJSON struct {
	ID               apijson.Field
	CreatedOn        apijson.Field
	DNS              apijson.Field
	ModifiedOn       apijson.Field
	Protocol         apijson.Field
	TrafficType      apijson.Field
	ArgoSmartRouting apijson.Field
	EdgeIPs          apijson.Field
	IPFirewall       apijson.Field
	OriginDirect     apijson.Field
	OriginDNS        apijson.Field
	OriginPort       apijson.Field
	ProxyProtocol    apijson.Field
	TLS              apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppGetResponseSpectrumConfigAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseSpectrumConfigAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r AppGetResponseSpectrumConfigAppConfig) implementsAppGetResponse() {}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppGetResponseSpectrumConfigAppConfigTrafficType string

const (
	AppGetResponseSpectrumConfigAppConfigTrafficTypeDirect AppGetResponseSpectrumConfigAppConfigTrafficType = "direct"
	AppGetResponseSpectrumConfigAppConfigTrafficTypeHTTP   AppGetResponseSpectrumConfigAppConfigTrafficType = "http"
	AppGetResponseSpectrumConfigAppConfigTrafficTypeHTTPS  AppGetResponseSpectrumConfigAppConfigTrafficType = "https"
)

func (r AppGetResponseSpectrumConfigAppConfigTrafficType) IsKnown() bool {
	switch r {
	case AppGetResponseSpectrumConfigAppConfigTrafficTypeDirect, AppGetResponseSpectrumConfigAppConfigTrafficTypeHTTP, AppGetResponseSpectrumConfigAppConfigTrafficTypeHTTPS:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppGetResponseSpectrumConfigAppConfigProxyProtocol string

const (
	AppGetResponseSpectrumConfigAppConfigProxyProtocolOff    AppGetResponseSpectrumConfigAppConfigProxyProtocol = "off"
	AppGetResponseSpectrumConfigAppConfigProxyProtocolV1     AppGetResponseSpectrumConfigAppConfigProxyProtocol = "v1"
	AppGetResponseSpectrumConfigAppConfigProxyProtocolV2     AppGetResponseSpectrumConfigAppConfigProxyProtocol = "v2"
	AppGetResponseSpectrumConfigAppConfigProxyProtocolSimple AppGetResponseSpectrumConfigAppConfigProxyProtocol = "simple"
)

func (r AppGetResponseSpectrumConfigAppConfigProxyProtocol) IsKnown() bool {
	switch r {
	case AppGetResponseSpectrumConfigAppConfigProxyProtocolOff, AppGetResponseSpectrumConfigAppConfigProxyProtocolV1, AppGetResponseSpectrumConfigAppConfigProxyProtocolV2, AppGetResponseSpectrumConfigAppConfigProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppGetResponseSpectrumConfigAppConfigTLS string

const (
	AppGetResponseSpectrumConfigAppConfigTLSOff      AppGetResponseSpectrumConfigAppConfigTLS = "off"
	AppGetResponseSpectrumConfigAppConfigTLSFlexible AppGetResponseSpectrumConfigAppConfigTLS = "flexible"
	AppGetResponseSpectrumConfigAppConfigTLSFull     AppGetResponseSpectrumConfigAppConfigTLS = "full"
	AppGetResponseSpectrumConfigAppConfigTLSStrict   AppGetResponseSpectrumConfigAppConfigTLS = "strict"
)

func (r AppGetResponseSpectrumConfigAppConfigTLS) IsKnown() bool {
	switch r {
	case AppGetResponseSpectrumConfigAppConfigTLSOff, AppGetResponseSpectrumConfigAppConfigTLSFlexible, AppGetResponseSpectrumConfigAppConfigTLSFull, AppGetResponseSpectrumConfigAppConfigTLSStrict:
		return true
	}
	return false
}

type AppGetResponseSpectrumConfigPaygoAppConfig struct {
	// App identifier.
	ID string `json:"id,required"`
	// When the Application was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The name and type of DNS record for the Spectrum application.
	DNS DNS `json:"dns,required"`
	// When the Application was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol string `json:"protocol,required"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect []string                                       `json:"origin_direct" format:"URI"`
	JSON         appGetResponseSpectrumConfigPaygoAppConfigJSON `json:"-"`
}

// appGetResponseSpectrumConfigPaygoAppConfigJSON contains the JSON metadata for
// the struct [AppGetResponseSpectrumConfigPaygoAppConfig]
type appGetResponseSpectrumConfigPaygoAppConfigJSON struct {
	ID           apijson.Field
	CreatedOn    apijson.Field
	DNS          apijson.Field
	ModifiedOn   apijson.Field
	Protocol     apijson.Field
	OriginDirect apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AppGetResponseSpectrumConfigPaygoAppConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseSpectrumConfigPaygoAppConfigJSON) RawJSON() string {
	return r.raw
}

func (r AppGetResponseSpectrumConfigPaygoAppConfig) implementsAppGetResponse() {}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppGetResponseProxyProtocol string

const (
	AppGetResponseProxyProtocolOff    AppGetResponseProxyProtocol = "off"
	AppGetResponseProxyProtocolV1     AppGetResponseProxyProtocol = "v1"
	AppGetResponseProxyProtocolV2     AppGetResponseProxyProtocol = "v2"
	AppGetResponseProxyProtocolSimple AppGetResponseProxyProtocol = "simple"
)

func (r AppGetResponseProxyProtocol) IsKnown() bool {
	switch r {
	case AppGetResponseProxyProtocolOff, AppGetResponseProxyProtocolV1, AppGetResponseProxyProtocolV2, AppGetResponseProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppGetResponseTLS string

const (
	AppGetResponseTLSOff      AppGetResponseTLS = "off"
	AppGetResponseTLSFlexible AppGetResponseTLS = "flexible"
	AppGetResponseTLSFull     AppGetResponseTLS = "full"
	AppGetResponseTLSStrict   AppGetResponseTLS = "strict"
)

func (r AppGetResponseTLS) IsKnown() bool {
	switch r {
	case AppGetResponseTLSOff, AppGetResponseTLSFlexible, AppGetResponseTLSFull, AppGetResponseTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppGetResponseTrafficType string

const (
	AppGetResponseTrafficTypeDirect AppGetResponseTrafficType = "direct"
	AppGetResponseTrafficTypeHTTP   AppGetResponseTrafficType = "http"
	AppGetResponseTrafficTypeHTTPS  AppGetResponseTrafficType = "https"
)

func (r AppGetResponseTrafficType) IsKnown() bool {
	switch r {
	case AppGetResponseTrafficTypeDirect, AppGetResponseTrafficTypeHTTP, AppGetResponseTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppNewParams struct {
	// Zone identifier.
	ZoneID param.Field[string]   `path:"zone_id,required"`
	Body   AppNewParamsBodyUnion `json:"body,required"`
}

func (r AppNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AppNewParamsBody struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[EdgeIPsUnionParam] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall   param.Field[bool]        `json:"ip_firewall"`
	OriginDirect param.Field[interface{}] `json:"origin_direct"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[OriginDNSParam] `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[OriginPortUnionParam] `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppNewParamsBodyProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppNewParamsBodyTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppNewParamsBodyTrafficType] `json:"traffic_type"`
}

func (r AppNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppNewParamsBody) implementsAppNewParamsBodyUnion() {}

// Satisfied by [spectrum.AppNewParamsBodySpectrumConfigAppConfig],
// [spectrum.AppNewParamsBodySpectrumConfigPaygoAppConfig], [AppNewParamsBody].
type AppNewParamsBodyUnion interface {
	implementsAppNewParamsBodyUnion()
}

type AppNewParamsBodySpectrumConfigAppConfig struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppNewParamsBodySpectrumConfigAppConfigTrafficType] `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[EdgeIPsUnionParam] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect param.Field[[]string] `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[OriginDNSParam] `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[OriginPortUnionParam] `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppNewParamsBodySpectrumConfigAppConfigProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppNewParamsBodySpectrumConfigAppConfigTLS] `json:"tls"`
}

func (r AppNewParamsBodySpectrumConfigAppConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppNewParamsBodySpectrumConfigAppConfig) implementsAppNewParamsBodyUnion() {}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewParamsBodySpectrumConfigAppConfigTrafficType string

const (
	AppNewParamsBodySpectrumConfigAppConfigTrafficTypeDirect AppNewParamsBodySpectrumConfigAppConfigTrafficType = "direct"
	AppNewParamsBodySpectrumConfigAppConfigTrafficTypeHTTP   AppNewParamsBodySpectrumConfigAppConfigTrafficType = "http"
	AppNewParamsBodySpectrumConfigAppConfigTrafficTypeHTTPS  AppNewParamsBodySpectrumConfigAppConfigTrafficType = "https"
)

func (r AppNewParamsBodySpectrumConfigAppConfigTrafficType) IsKnown() bool {
	switch r {
	case AppNewParamsBodySpectrumConfigAppConfigTrafficTypeDirect, AppNewParamsBodySpectrumConfigAppConfigTrafficTypeHTTP, AppNewParamsBodySpectrumConfigAppConfigTrafficTypeHTTPS:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppNewParamsBodySpectrumConfigAppConfigProxyProtocol string

const (
	AppNewParamsBodySpectrumConfigAppConfigProxyProtocolOff    AppNewParamsBodySpectrumConfigAppConfigProxyProtocol = "off"
	AppNewParamsBodySpectrumConfigAppConfigProxyProtocolV1     AppNewParamsBodySpectrumConfigAppConfigProxyProtocol = "v1"
	AppNewParamsBodySpectrumConfigAppConfigProxyProtocolV2     AppNewParamsBodySpectrumConfigAppConfigProxyProtocol = "v2"
	AppNewParamsBodySpectrumConfigAppConfigProxyProtocolSimple AppNewParamsBodySpectrumConfigAppConfigProxyProtocol = "simple"
)

func (r AppNewParamsBodySpectrumConfigAppConfigProxyProtocol) IsKnown() bool {
	switch r {
	case AppNewParamsBodySpectrumConfigAppConfigProxyProtocolOff, AppNewParamsBodySpectrumConfigAppConfigProxyProtocolV1, AppNewParamsBodySpectrumConfigAppConfigProxyProtocolV2, AppNewParamsBodySpectrumConfigAppConfigProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppNewParamsBodySpectrumConfigAppConfigTLS string

const (
	AppNewParamsBodySpectrumConfigAppConfigTLSOff      AppNewParamsBodySpectrumConfigAppConfigTLS = "off"
	AppNewParamsBodySpectrumConfigAppConfigTLSFlexible AppNewParamsBodySpectrumConfigAppConfigTLS = "flexible"
	AppNewParamsBodySpectrumConfigAppConfigTLSFull     AppNewParamsBodySpectrumConfigAppConfigTLS = "full"
	AppNewParamsBodySpectrumConfigAppConfigTLSStrict   AppNewParamsBodySpectrumConfigAppConfigTLS = "strict"
)

func (r AppNewParamsBodySpectrumConfigAppConfigTLS) IsKnown() bool {
	switch r {
	case AppNewParamsBodySpectrumConfigAppConfigTLSOff, AppNewParamsBodySpectrumConfigAppConfigTLSFlexible, AppNewParamsBodySpectrumConfigAppConfigTLSFull, AppNewParamsBodySpectrumConfigAppConfigTLSStrict:
		return true
	}
	return false
}

type AppNewParamsBodySpectrumConfigPaygoAppConfig struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect param.Field[[]string] `json:"origin_direct" format:"URI"`
}

func (r AppNewParamsBodySpectrumConfigPaygoAppConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppNewParamsBodySpectrumConfigPaygoAppConfig) implementsAppNewParamsBodyUnion() {}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppNewParamsBodyProxyProtocol string

const (
	AppNewParamsBodyProxyProtocolOff    AppNewParamsBodyProxyProtocol = "off"
	AppNewParamsBodyProxyProtocolV1     AppNewParamsBodyProxyProtocol = "v1"
	AppNewParamsBodyProxyProtocolV2     AppNewParamsBodyProxyProtocol = "v2"
	AppNewParamsBodyProxyProtocolSimple AppNewParamsBodyProxyProtocol = "simple"
)

func (r AppNewParamsBodyProxyProtocol) IsKnown() bool {
	switch r {
	case AppNewParamsBodyProxyProtocolOff, AppNewParamsBodyProxyProtocolV1, AppNewParamsBodyProxyProtocolV2, AppNewParamsBodyProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppNewParamsBodyTLS string

const (
	AppNewParamsBodyTLSOff      AppNewParamsBodyTLS = "off"
	AppNewParamsBodyTLSFlexible AppNewParamsBodyTLS = "flexible"
	AppNewParamsBodyTLSFull     AppNewParamsBodyTLS = "full"
	AppNewParamsBodyTLSStrict   AppNewParamsBodyTLS = "strict"
)

func (r AppNewParamsBodyTLS) IsKnown() bool {
	switch r {
	case AppNewParamsBodyTLSOff, AppNewParamsBodyTLSFlexible, AppNewParamsBodyTLSFull, AppNewParamsBodyTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppNewParamsBodyTrafficType string

const (
	AppNewParamsBodyTrafficTypeDirect AppNewParamsBodyTrafficType = "direct"
	AppNewParamsBodyTrafficTypeHTTP   AppNewParamsBodyTrafficType = "http"
	AppNewParamsBodyTrafficTypeHTTPS  AppNewParamsBodyTrafficType = "https"
)

func (r AppNewParamsBodyTrafficType) IsKnown() bool {
	switch r {
	case AppNewParamsBodyTrafficTypeDirect, AppNewParamsBodyTrafficTypeHTTP, AppNewParamsBodyTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppNewResponseEnvelope struct {
	Errors   []AppNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AppNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AppNewResponse                `json:"result"`
	JSON    appNewResponseEnvelopeJSON    `json:"-"`
}

// appNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelope]
type appNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           AppNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             appNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// appNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelopeErrors]
type appNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    appNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// appNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelopeErrorsSource]
type appNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AppNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             appNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// appNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppNewResponseEnvelopeMessages]
type appNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppNewResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    appNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// appNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AppNewResponseEnvelopeMessagesSource]
type appNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AppNewResponseEnvelopeSuccess bool

const (
	AppNewResponseEnvelopeSuccessTrue AppNewResponseEnvelopeSuccess = true
)

func (r AppNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppUpdateParams struct {
	// Zone identifier.
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   AppUpdateParamsBodyUnion `json:"body,required"`
}

func (r AppUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type AppUpdateParamsBody struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[EdgeIPsUnionParam] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall   param.Field[bool]        `json:"ip_firewall"`
	OriginDirect param.Field[interface{}] `json:"origin_direct"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[OriginDNSParam] `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[OriginPortUnionParam] `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppUpdateParamsBodyProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppUpdateParamsBodyTLS] `json:"tls"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppUpdateParamsBodyTrafficType] `json:"traffic_type"`
}

func (r AppUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppUpdateParamsBody) implementsAppUpdateParamsBodyUnion() {}

// Satisfied by [spectrum.AppUpdateParamsBodySpectrumConfigAppConfig],
// [spectrum.AppUpdateParamsBodySpectrumConfigPaygoAppConfig],
// [AppUpdateParamsBody].
type AppUpdateParamsBodyUnion interface {
	implementsAppUpdateParamsBodyUnion()
}

type AppUpdateParamsBodySpectrumConfigAppConfig struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// Determines how data travels from the edge to your origin. When set to "direct",
	// Spectrum will send traffic directly to your origin, and the application's type
	// is derived from the `protocol`. When set to "http" or "https", Spectrum will
	// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
	// the application type matches this property exactly.
	TrafficType param.Field[AppUpdateParamsBodySpectrumConfigAppConfigTrafficType] `json:"traffic_type,required"`
	// Enables Argo Smart Routing for this application. Notes: Only available for TCP
	// applications with traffic_type set to "direct".
	ArgoSmartRouting param.Field[bool] `json:"argo_smart_routing"`
	// The anycast edge IP configuration for the hostname of this application.
	EdgeIPs param.Field[EdgeIPsUnionParam] `json:"edge_ips"`
	// Enables IP Access Rules for this application. Notes: Only available for TCP
	// applications.
	IPFirewall param.Field[bool] `json:"ip_firewall"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect param.Field[[]string] `json:"origin_direct" format:"URI"`
	// The name and type of DNS record for the Spectrum application.
	OriginDNS param.Field[OriginDNSParam] `json:"origin_dns"`
	// The destination port at the origin. Only specified in conjunction with
	// origin_dns. May use an integer to specify a single origin port, for example
	// `1000`, or a string to specify a range of origin ports, for example
	// `"1000-2000"`. Notes: If specifying a port range, the number of ports in the
	// range must match the number of ports specified in the "protocol" field.
	OriginPort param.Field[OriginPortUnionParam] `json:"origin_port"`
	// Enables Proxy Protocol to the origin. Refer to
	// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
	// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
	// Proxy Protocol.
	ProxyProtocol param.Field[AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocol] `json:"proxy_protocol"`
	// The type of TLS termination associated with the application.
	TLS param.Field[AppUpdateParamsBodySpectrumConfigAppConfigTLS] `json:"tls"`
}

func (r AppUpdateParamsBodySpectrumConfigAppConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppUpdateParamsBodySpectrumConfigAppConfig) implementsAppUpdateParamsBodyUnion() {}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateParamsBodySpectrumConfigAppConfigTrafficType string

const (
	AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeDirect AppUpdateParamsBodySpectrumConfigAppConfigTrafficType = "direct"
	AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeHTTP   AppUpdateParamsBodySpectrumConfigAppConfigTrafficType = "http"
	AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeHTTPS  AppUpdateParamsBodySpectrumConfigAppConfigTrafficType = "https"
)

func (r AppUpdateParamsBodySpectrumConfigAppConfigTrafficType) IsKnown() bool {
	switch r {
	case AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeDirect, AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeHTTP, AppUpdateParamsBodySpectrumConfigAppConfigTrafficTypeHTTPS:
		return true
	}
	return false
}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocol string

const (
	AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolOff    AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocol = "off"
	AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolV1     AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocol = "v1"
	AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolV2     AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocol = "v2"
	AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolSimple AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocol = "simple"
)

func (r AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocol) IsKnown() bool {
	switch r {
	case AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolOff, AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolV1, AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolV2, AppUpdateParamsBodySpectrumConfigAppConfigProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppUpdateParamsBodySpectrumConfigAppConfigTLS string

const (
	AppUpdateParamsBodySpectrumConfigAppConfigTLSOff      AppUpdateParamsBodySpectrumConfigAppConfigTLS = "off"
	AppUpdateParamsBodySpectrumConfigAppConfigTLSFlexible AppUpdateParamsBodySpectrumConfigAppConfigTLS = "flexible"
	AppUpdateParamsBodySpectrumConfigAppConfigTLSFull     AppUpdateParamsBodySpectrumConfigAppConfigTLS = "full"
	AppUpdateParamsBodySpectrumConfigAppConfigTLSStrict   AppUpdateParamsBodySpectrumConfigAppConfigTLS = "strict"
)

func (r AppUpdateParamsBodySpectrumConfigAppConfigTLS) IsKnown() bool {
	switch r {
	case AppUpdateParamsBodySpectrumConfigAppConfigTLSOff, AppUpdateParamsBodySpectrumConfigAppConfigTLSFlexible, AppUpdateParamsBodySpectrumConfigAppConfigTLSFull, AppUpdateParamsBodySpectrumConfigAppConfigTLSStrict:
		return true
	}
	return false
}

type AppUpdateParamsBodySpectrumConfigPaygoAppConfig struct {
	// The name and type of DNS record for the Spectrum application.
	DNS param.Field[DNSParam] `json:"dns,required"`
	// The port configuration at Cloudflare's edge. May specify a single port, for
	// example `"tcp/1000"`, or a range of ports, for example `"tcp/1000-2000"`.
	Protocol param.Field[string] `json:"protocol,required"`
	// List of origin IP addresses. Array may contain multiple IP addresses for load
	// balancing.
	OriginDirect param.Field[[]string] `json:"origin_direct" format:"URI"`
}

func (r AppUpdateParamsBodySpectrumConfigPaygoAppConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AppUpdateParamsBodySpectrumConfigPaygoAppConfig) implementsAppUpdateParamsBodyUnion() {}

// Enables Proxy Protocol to the origin. Refer to
// [Enable Proxy protocol](https://developers.cloudflare.com/spectrum/getting-started/proxy-protocol/)
// for implementation details on PROXY Protocol V1, PROXY Protocol V2, and Simple
// Proxy Protocol.
type AppUpdateParamsBodyProxyProtocol string

const (
	AppUpdateParamsBodyProxyProtocolOff    AppUpdateParamsBodyProxyProtocol = "off"
	AppUpdateParamsBodyProxyProtocolV1     AppUpdateParamsBodyProxyProtocol = "v1"
	AppUpdateParamsBodyProxyProtocolV2     AppUpdateParamsBodyProxyProtocol = "v2"
	AppUpdateParamsBodyProxyProtocolSimple AppUpdateParamsBodyProxyProtocol = "simple"
)

func (r AppUpdateParamsBodyProxyProtocol) IsKnown() bool {
	switch r {
	case AppUpdateParamsBodyProxyProtocolOff, AppUpdateParamsBodyProxyProtocolV1, AppUpdateParamsBodyProxyProtocolV2, AppUpdateParamsBodyProxyProtocolSimple:
		return true
	}
	return false
}

// The type of TLS termination associated with the application.
type AppUpdateParamsBodyTLS string

const (
	AppUpdateParamsBodyTLSOff      AppUpdateParamsBodyTLS = "off"
	AppUpdateParamsBodyTLSFlexible AppUpdateParamsBodyTLS = "flexible"
	AppUpdateParamsBodyTLSFull     AppUpdateParamsBodyTLS = "full"
	AppUpdateParamsBodyTLSStrict   AppUpdateParamsBodyTLS = "strict"
)

func (r AppUpdateParamsBodyTLS) IsKnown() bool {
	switch r {
	case AppUpdateParamsBodyTLSOff, AppUpdateParamsBodyTLSFlexible, AppUpdateParamsBodyTLSFull, AppUpdateParamsBodyTLSStrict:
		return true
	}
	return false
}

// Determines how data travels from the edge to your origin. When set to "direct",
// Spectrum will send traffic directly to your origin, and the application's type
// is derived from the `protocol`. When set to "http" or "https", Spectrum will
// apply Cloudflare's HTTP/HTTPS features as it sends traffic to your origin, and
// the application type matches this property exactly.
type AppUpdateParamsBodyTrafficType string

const (
	AppUpdateParamsBodyTrafficTypeDirect AppUpdateParamsBodyTrafficType = "direct"
	AppUpdateParamsBodyTrafficTypeHTTP   AppUpdateParamsBodyTrafficType = "http"
	AppUpdateParamsBodyTrafficTypeHTTPS  AppUpdateParamsBodyTrafficType = "https"
)

func (r AppUpdateParamsBodyTrafficType) IsKnown() bool {
	switch r {
	case AppUpdateParamsBodyTrafficTypeDirect, AppUpdateParamsBodyTrafficTypeHTTP, AppUpdateParamsBodyTrafficTypeHTTPS:
		return true
	}
	return false
}

type AppUpdateResponseEnvelope struct {
	Errors   []AppUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AppUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  AppUpdateResponse                `json:"result"`
	JSON    appUpdateResponseEnvelopeJSON    `json:"-"`
}

// appUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelope]
type appUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           AppUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             appUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// appUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelopeErrors]
type appUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    appUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// appUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AppUpdateResponseEnvelopeErrorsSource]
type appUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AppUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             appUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// appUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppUpdateResponseEnvelopeMessages]
type appUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    appUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// appUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AppUpdateResponseEnvelopeMessagesSource]
type appUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AppUpdateResponseEnvelopeSuccess bool

const (
	AppUpdateResponseEnvelopeSuccessTrue AppUpdateResponseEnvelopeSuccess = true
)

func (r AppUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppListParams struct {
	// Zone identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Sets the direction by which results are ordered.
	Direction param.Field[AppListParamsDirection] `query:"direction"`
	// Application field by which results are ordered.
	Order param.Field[AppListParamsOrder] `query:"order"`
	// Page number of paginated results. This parameter is required in order to use
	// other pagination parameters. If included in the query, `result_info` will be
	// present in the response.
	Page param.Field[float64] `query:"page"`
	// Sets the maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AppListParams]'s query parameters as `url.Values`.
func (r AppListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sets the direction by which results are ordered.
type AppListParamsDirection string

const (
	AppListParamsDirectionAsc  AppListParamsDirection = "asc"
	AppListParamsDirectionDesc AppListParamsDirection = "desc"
)

func (r AppListParamsDirection) IsKnown() bool {
	switch r {
	case AppListParamsDirectionAsc, AppListParamsDirectionDesc:
		return true
	}
	return false
}

// Application field by which results are ordered.
type AppListParamsOrder string

const (
	AppListParamsOrderProtocol   AppListParamsOrder = "protocol"
	AppListParamsOrderAppID      AppListParamsOrder = "app_id"
	AppListParamsOrderCreatedOn  AppListParamsOrder = "created_on"
	AppListParamsOrderModifiedOn AppListParamsOrder = "modified_on"
	AppListParamsOrderDNS        AppListParamsOrder = "dns"
)

func (r AppListParamsOrder) IsKnown() bool {
	switch r {
	case AppListParamsOrderProtocol, AppListParamsOrderAppID, AppListParamsOrderCreatedOn, AppListParamsOrderModifiedOn, AppListParamsOrderDNS:
		return true
	}
	return false
}

type AppDeleteParams struct {
	// Zone identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AppDeleteResponseEnvelope struct {
	Errors   []AppDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AppDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  AppDeleteResponse                `json:"result,nullable"`
	JSON    appDeleteResponseEnvelopeJSON    `json:"-"`
}

// appDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelope]
type appDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           AppDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             appDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// appDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelopeErrors]
type appDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    appDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// appDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AppDeleteResponseEnvelopeErrorsSource]
type appDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           AppDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             appDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// appDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppDeleteResponseEnvelopeMessages]
type appDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    appDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// appDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AppDeleteResponseEnvelopeMessagesSource]
type appDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AppDeleteResponseEnvelopeSuccess bool

const (
	AppDeleteResponseEnvelopeSuccessTrue AppDeleteResponseEnvelopeSuccess = true
)

func (r AppDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AppGetParams struct {
	// Zone identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AppGetResponseEnvelope struct {
	Errors   []AppGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AppGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AppGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AppGetResponse                `json:"result"`
	JSON    appGetResponseEnvelopeJSON    `json:"-"`
}

// appGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelope]
type appGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           AppGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             appGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// appGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelopeErrors]
type appGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    appGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// appGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelopeErrorsSource]
type appGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AppGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             appGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// appGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [AppGetResponseEnvelopeMessages]
type appGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AppGetResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    appGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// appGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [AppGetResponseEnvelopeMessagesSource]
type appGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AppGetResponseEnvelopeSuccess bool

const (
	AppGetResponseEnvelopeSuccessTrue AppGetResponseEnvelopeSuccess = true
)

func (r AppGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AppGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
