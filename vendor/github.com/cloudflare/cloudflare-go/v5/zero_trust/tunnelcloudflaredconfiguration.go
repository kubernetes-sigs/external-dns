// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// TunnelCloudflaredConfigurationService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTunnelCloudflaredConfigurationService] method instead.
type TunnelCloudflaredConfigurationService struct {
	Options []option.RequestOption
}

// NewTunnelCloudflaredConfigurationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTunnelCloudflaredConfigurationService(opts ...option.RequestOption) (r *TunnelCloudflaredConfigurationService) {
	r = &TunnelCloudflaredConfigurationService{}
	r.Options = opts
	return
}

// Adds or updates the configuration for a remotely-managed tunnel.
func (r *TunnelCloudflaredConfigurationService) Update(ctx context.Context, tunnelID string, params TunnelCloudflaredConfigurationUpdateParams, opts ...option.RequestOption) (res *TunnelCloudflaredConfigurationUpdateResponse, err error) {
	var env TunnelCloudflaredConfigurationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", params.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the configuration for a remotely-managed tunnel
func (r *TunnelCloudflaredConfigurationService) Get(ctx context.Context, tunnelID string, query TunnelCloudflaredConfigurationGetParams, opts ...option.RequestOption) (res *TunnelCloudflaredConfigurationGetResponse, err error) {
	var env TunnelCloudflaredConfigurationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tunnelID == "" {
		err = errors.New("missing required tunnel_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cfd_tunnel/%s/configurations", query.AccountID, tunnelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Cloudflare Tunnel configuration
type TunnelCloudflaredConfigurationUpdateResponse struct {
	// Identifier.
	AccountID string `json:"account_id"`
	// The tunnel configuration and ingress rules.
	Config    TunnelCloudflaredConfigurationUpdateResponseConfig `json:"config"`
	CreatedAt time.Time                                          `json:"created_at" format:"date-time"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel's configuration on the Zero Trust dashboard.
	Source TunnelCloudflaredConfigurationUpdateResponseSource `json:"source"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// The version of the Tunnel Configuration.
	Version int64                                            `json:"version"`
	JSON    tunnelCloudflaredConfigurationUpdateResponseJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseJSON contains the JSON metadata for
// the struct [TunnelCloudflaredConfigurationUpdateResponse]
type tunnelCloudflaredConfigurationUpdateResponseJSON struct {
	AccountID   apijson.Field
	Config      apijson.Field
	CreatedAt   apijson.Field
	Source      apijson.Field
	TunnelID    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The tunnel configuration and ingress rules.
type TunnelCloudflaredConfigurationUpdateResponseConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress []TunnelCloudflaredConfigurationUpdateResponseConfigIngress `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequest `json:"originRequest"`
	// Enable private network access from WARP users to private network routes. This is
	// enabled if the tunnel has an assigned route.
	WARPRouting TunnelCloudflaredConfigurationUpdateResponseConfigWARPRouting `json:"warp-routing"`
	JSON        tunnelCloudflaredConfigurationUpdateResponseConfigJSON        `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseConfigJSON contains the JSON
// metadata for the struct [TunnelCloudflaredConfigurationUpdateResponseConfig]
type tunnelCloudflaredConfigurationUpdateResponseConfigJSON struct {
	Ingress       apijson.Field
	OriginRequest apijson.Field
	WARPRouting   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseConfigJSON) RawJSON() string {
	return r.raw
}

// Public hostname
type TunnelCloudflaredConfigurationUpdateResponseConfigIngress struct {
	// Public hostname for this service.
	Hostname string `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service string `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequest `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path string                                                        `json:"path"`
	JSON tunnelCloudflaredConfigurationUpdateResponseConfigIngressJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseConfigIngressJSON contains the JSON
// metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseConfigIngress]
type tunnelCloudflaredConfigurationUpdateResponseConfigIngressJSON struct {
	Hostname      apijson.Field
	Service       apijson.Field
	OriginRequest apijson.Field
	Path          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseConfigIngress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseConfigIngressJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                                      `json:"tlsTimeout"`
	JSON       tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestJSON
// contains the JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequest]
type tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                                             `json:"required"`
	JSON     tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON
// contains the JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccess]
type tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseConfigIngressOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                               `json:"tlsTimeout"`
	JSON       tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequest]
type tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                                      `json:"required"`
	JSON     tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccessJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccessJSON
// contains the JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccess]
type tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseConfigOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelCloudflaredConfigurationUpdateResponseConfigWARPRouting struct {
	Enabled bool                                                              `json:"enabled"`
	JSON    tunnelCloudflaredConfigurationUpdateResponseConfigWARPRoutingJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseConfigWARPRoutingJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseConfigWARPRouting]
type tunnelCloudflaredConfigurationUpdateResponseConfigWARPRoutingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseConfigWARPRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseConfigWARPRoutingJSON) RawJSON() string {
	return r.raw
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel's configuration on the Zero Trust dashboard.
type TunnelCloudflaredConfigurationUpdateResponseSource string

const (
	TunnelCloudflaredConfigurationUpdateResponseSourceLocal      TunnelCloudflaredConfigurationUpdateResponseSource = "local"
	TunnelCloudflaredConfigurationUpdateResponseSourceCloudflare TunnelCloudflaredConfigurationUpdateResponseSource = "cloudflare"
)

func (r TunnelCloudflaredConfigurationUpdateResponseSource) IsKnown() bool {
	switch r {
	case TunnelCloudflaredConfigurationUpdateResponseSourceLocal, TunnelCloudflaredConfigurationUpdateResponseSourceCloudflare:
		return true
	}
	return false
}

// Cloudflare Tunnel configuration
type TunnelCloudflaredConfigurationGetResponse struct {
	// Identifier.
	AccountID string `json:"account_id"`
	// The tunnel configuration and ingress rules.
	Config    TunnelCloudflaredConfigurationGetResponseConfig `json:"config"`
	CreatedAt time.Time                                       `json:"created_at" format:"date-time"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel's configuration on the Zero Trust dashboard.
	Source TunnelCloudflaredConfigurationGetResponseSource `json:"source"`
	// UUID of the tunnel.
	TunnelID string `json:"tunnel_id" format:"uuid"`
	// The version of the Tunnel Configuration.
	Version int64                                         `json:"version"`
	JSON    tunnelCloudflaredConfigurationGetResponseJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseJSON contains the JSON metadata for the
// struct [TunnelCloudflaredConfigurationGetResponse]
type tunnelCloudflaredConfigurationGetResponseJSON struct {
	AccountID   apijson.Field
	Config      apijson.Field
	CreatedAt   apijson.Field
	Source      apijson.Field
	TunnelID    apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseJSON) RawJSON() string {
	return r.raw
}

// The tunnel configuration and ingress rules.
type TunnelCloudflaredConfigurationGetResponseConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress []TunnelCloudflaredConfigurationGetResponseConfigIngress `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelCloudflaredConfigurationGetResponseConfigOriginRequest `json:"originRequest"`
	// Enable private network access from WARP users to private network routes. This is
	// enabled if the tunnel has an assigned route.
	WARPRouting TunnelCloudflaredConfigurationGetResponseConfigWARPRouting `json:"warp-routing"`
	JSON        tunnelCloudflaredConfigurationGetResponseConfigJSON        `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseConfigJSON contains the JSON metadata
// for the struct [TunnelCloudflaredConfigurationGetResponseConfig]
type tunnelCloudflaredConfigurationGetResponseConfigJSON struct {
	Ingress       apijson.Field
	OriginRequest apijson.Field
	WARPRouting   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseConfig) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseConfigJSON) RawJSON() string {
	return r.raw
}

// Public hostname
type TunnelCloudflaredConfigurationGetResponseConfigIngress struct {
	// Public hostname for this service.
	Hostname string `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service string `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequest `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path string                                                     `json:"path"`
	JSON tunnelCloudflaredConfigurationGetResponseConfigIngressJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseConfigIngressJSON contains the JSON
// metadata for the struct [TunnelCloudflaredConfigurationGetResponseConfigIngress]
type tunnelCloudflaredConfigurationGetResponseConfigIngressJSON struct {
	Hostname      apijson.Field
	Service       apijson.Field
	OriginRequest apijson.Field
	Path          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseConfigIngress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseConfigIngressJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                                   `json:"tlsTimeout"`
	JSON       tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestJSON contains
// the JSON metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequest]
type tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                                          `json:"required"`
	JSON     tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccessJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccessJSON
// contains the JSON metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccess]
type tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseConfigIngressOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelCloudflaredConfigurationGetResponseConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access TunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccess `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool string `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout int64 `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding bool `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin bool `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader string `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections int64 `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout int64 `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs bool `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify bool `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName string `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType string `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive int64 `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout int64                                                            `json:"tlsTimeout"`
	JSON       tunnelCloudflaredConfigurationGetResponseConfigOriginRequestJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseConfigOriginRequestJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseConfigOriginRequest]
type tunnelCloudflaredConfigurationGetResponseConfigOriginRequestJSON struct {
	Access                 apijson.Field
	CAPool                 apijson.Field
	ConnectTimeout         apijson.Field
	DisableChunkedEncoding apijson.Field
	HTTP2Origin            apijson.Field
	HTTPHostHeader         apijson.Field
	KeepAliveConnections   apijson.Field
	KeepAliveTimeout       apijson.Field
	NoHappyEyeballs        apijson.Field
	NoTLSVerify            apijson.Field
	OriginServerName       apijson.Field
	ProxyType              apijson.Field
	TCPKeepAlive           apijson.Field
	TLSTimeout             apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseConfigOriginRequest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseConfigOriginRequestJSON) RawJSON() string {
	return r.raw
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   []string `json:"audTag,required"`
	TeamName string   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required bool                                                                   `json:"required"`
	JSON     tunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccessJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccessJSON contains
// the JSON metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccess]
type tunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccessJSON struct {
	AUDTag      apijson.Field
	TeamName    apijson.Field
	Required    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccess) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseConfigOriginRequestAccessJSON) RawJSON() string {
	return r.raw
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelCloudflaredConfigurationGetResponseConfigWARPRouting struct {
	Enabled bool                                                           `json:"enabled"`
	JSON    tunnelCloudflaredConfigurationGetResponseConfigWARPRoutingJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseConfigWARPRoutingJSON contains the JSON
// metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseConfigWARPRouting]
type tunnelCloudflaredConfigurationGetResponseConfigWARPRoutingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseConfigWARPRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseConfigWARPRoutingJSON) RawJSON() string {
	return r.raw
}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel's configuration on the Zero Trust dashboard.
type TunnelCloudflaredConfigurationGetResponseSource string

const (
	TunnelCloudflaredConfigurationGetResponseSourceLocal      TunnelCloudflaredConfigurationGetResponseSource = "local"
	TunnelCloudflaredConfigurationGetResponseSourceCloudflare TunnelCloudflaredConfigurationGetResponseSource = "cloudflare"
)

func (r TunnelCloudflaredConfigurationGetResponseSource) IsKnown() bool {
	switch r {
	case TunnelCloudflaredConfigurationGetResponseSourceLocal, TunnelCloudflaredConfigurationGetResponseSourceCloudflare:
		return true
	}
	return false
}

type TunnelCloudflaredConfigurationUpdateParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The tunnel configuration and ingress rules.
	Config param.Field[TunnelCloudflaredConfigurationUpdateParamsConfig] `json:"config"`
}

func (r TunnelCloudflaredConfigurationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The tunnel configuration and ingress rules.
type TunnelCloudflaredConfigurationUpdateParamsConfig struct {
	// List of public hostname definitions. At least one ingress rule needs to be
	// defined for the tunnel.
	Ingress param.Field[[]TunnelCloudflaredConfigurationUpdateParamsConfigIngress] `json:"ingress"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequest] `json:"originRequest"`
}

func (r TunnelCloudflaredConfigurationUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public hostname
type TunnelCloudflaredConfigurationUpdateParamsConfigIngress struct {
	// Public hostname for this service.
	Hostname param.Field[string] `json:"hostname,required"`
	// Protocol and address of destination server. Supported protocols: http://,
	// https://, unix://, tcp://, ssh://, rdp://, unix+tls://, smb://. Alternatively
	// can return a HTTP status code http_status:[code] e.g. 'http_status:404'.
	Service param.Field[string] `json:"service,required"`
	// Configuration parameters for the public hostname specific connection settings
	// between cloudflared and origin server.
	OriginRequest param.Field[TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequest] `json:"originRequest"`
	// Requests with this path route to this public hostname.
	Path param.Field[string] `json:"path"`
}

func (r TunnelCloudflaredConfigurationUpdateParamsConfigIngress) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin param.Field[bool] `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader param.Field[string] `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections param.Field[int64] `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout param.Field[int64] `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs param.Field[bool] `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelCloudflaredConfigurationUpdateParamsConfigIngressOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration parameters for the public hostname specific connection settings
// between cloudflared and origin server.
type TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequest struct {
	// For all L7 requests to this hostname, cloudflared will validate each request's
	// Cf-Access-Jwt-Assertion request header.
	Access param.Field[TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequestAccess] `json:"access"`
	// Path to the certificate authority (CA) for the certificate of your origin. This
	// option should be used only if your certificate is not signed by Cloudflare.
	CAPool param.Field[string] `json:"caPool"`
	// Timeout for establishing a new TCP connection to your origin server. This
	// excludes the time taken to establish TLS, which is controlled by tlsTimeout.
	ConnectTimeout param.Field[int64] `json:"connectTimeout"`
	// Disables chunked transfer encoding. Useful if you are running a WSGI server.
	DisableChunkedEncoding param.Field[bool] `json:"disableChunkedEncoding"`
	// Attempt to connect to origin using HTTP2. Origin must be configured as https.
	HTTP2Origin param.Field[bool] `json:"http2Origin"`
	// Sets the HTTP Host header on requests sent to the local service.
	HTTPHostHeader param.Field[string] `json:"httpHostHeader"`
	// Maximum number of idle keepalive connections between Tunnel and your origin.
	// This does not restrict the total number of concurrent connections.
	KeepAliveConnections param.Field[int64] `json:"keepAliveConnections"`
	// Timeout after which an idle keepalive connection can be discarded.
	KeepAliveTimeout param.Field[int64] `json:"keepAliveTimeout"`
	// Disable the “happy eyeballs” algorithm for IPv4/IPv6 fallback if your local
	// network has misconfigured one of the protocols.
	NoHappyEyeballs param.Field[bool] `json:"noHappyEyeballs"`
	// Disables TLS verification of the certificate presented by your origin. Will
	// allow any certificate from the origin to be accepted.
	NoTLSVerify param.Field[bool] `json:"noTLSVerify"`
	// Hostname that cloudflared should expect from your origin server certificate.
	OriginServerName param.Field[string] `json:"originServerName"`
	// cloudflared starts a proxy server to translate HTTP traffic into TCP when
	// proxying, for example, SSH or RDP. This configures what type of proxy will be
	// started. Valid options are: "" for the regular proxy and "socks" for a SOCKS5
	// proxy.
	ProxyType param.Field[string] `json:"proxyType"`
	// The timeout after which a TCP keepalive packet is sent on a connection between
	// Tunnel and the origin server.
	TCPKeepAlive param.Field[int64] `json:"tcpKeepAlive"`
	// Timeout for completing a TLS handshake to your origin server, if you have chosen
	// to connect Tunnel to an HTTPS server.
	TLSTimeout param.Field[int64] `json:"tlsTimeout"`
}

func (r TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequest) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// For all L7 requests to this hostname, cloudflared will validate each request's
// Cf-Access-Jwt-Assertion request header.
type TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequestAccess struct {
	// Access applications that are allowed to reach this hostname for this Tunnel.
	// Audience tags can be identified in the dashboard or via the List Access policies
	// API.
	AUDTag   param.Field[[]string] `json:"audTag,required"`
	TeamName param.Field[string]   `json:"teamName,required"`
	// Deny traffic that has not fulfilled Access authorization.
	Required param.Field[bool] `json:"required"`
}

func (r TunnelCloudflaredConfigurationUpdateParamsConfigOriginRequestAccess) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable private network access from WARP users to private network routes. This is
// enabled if the tunnel has an assigned route.
type TunnelCloudflaredConfigurationUpdateParamsConfigWARPRouting struct {
	Enabled param.Field[bool] `json:"enabled"`
}

func (r TunnelCloudflaredConfigurationUpdateParamsConfigWARPRouting) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TunnelCloudflaredConfigurationUpdateResponseEnvelope struct {
	Errors   []TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TunnelCloudflaredConfigurationUpdateResponseEnvelopeSuccess `json:"success,required"`
	// Cloudflare Tunnel configuration
	Result TunnelCloudflaredConfigurationUpdateResponse             `json:"result"`
	JSON   tunnelCloudflaredConfigurationUpdateResponseEnvelopeJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [TunnelCloudflaredConfigurationUpdateResponseEnvelope]
type tunnelCloudflaredConfigurationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrors struct {
	Code             int64                                                            `json:"code,required"`
	Message          string                                                           `json:"message,required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrors]
type tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSource]
type tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessages struct {
	Code             int64                                                              `json:"code,required"`
	Message          string                                                             `json:"message,required"`
	DocumentationURL string                                                             `json:"documentation_url"`
	Source           TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessages]
type tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                                 `json:"pointer"`
	JSON    tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSource]
type tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TunnelCloudflaredConfigurationUpdateResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredConfigurationUpdateResponseEnvelopeSuccessTrue TunnelCloudflaredConfigurationUpdateResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredConfigurationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredConfigurationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TunnelCloudflaredConfigurationGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type TunnelCloudflaredConfigurationGetResponseEnvelope struct {
	Errors   []TunnelCloudflaredConfigurationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TunnelCloudflaredConfigurationGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success TunnelCloudflaredConfigurationGetResponseEnvelopeSuccess `json:"success,required"`
	// Cloudflare Tunnel configuration
	Result TunnelCloudflaredConfigurationGetResponse             `json:"result"`
	JSON   tunnelCloudflaredConfigurationGetResponseEnvelopeJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [TunnelCloudflaredConfigurationGetResponseEnvelope]
type tunnelCloudflaredConfigurationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationGetResponseEnvelopeErrors struct {
	Code             int64                                                         `json:"code,required"`
	Message          string                                                        `json:"message,required"`
	DocumentationURL string                                                        `json:"documentation_url"`
	Source           TunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseEnvelopeErrors]
type tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                            `json:"pointer"`
	JSON    tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSource]
type tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationGetResponseEnvelopeMessages struct {
	Code             int64                                                           `json:"code,required"`
	Message          string                                                          `json:"message,required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           TunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseEnvelopeMessages]
type tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [TunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSource]
type tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tunnelCloudflaredConfigurationGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TunnelCloudflaredConfigurationGetResponseEnvelopeSuccess bool

const (
	TunnelCloudflaredConfigurationGetResponseEnvelopeSuccessTrue TunnelCloudflaredConfigurationGetResponseEnvelopeSuccess = true
)

func (r TunnelCloudflaredConfigurationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TunnelCloudflaredConfigurationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
