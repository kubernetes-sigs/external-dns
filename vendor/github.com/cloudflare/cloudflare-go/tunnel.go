package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"errors"
)

// ErrMissingTunnelID is for when a required tunnel ID is missing from the
// parameters.
var ErrMissingTunnelID = errors.New("required missing tunnel ID")

// Tunnel is the struct definition of a tunnel.
type Tunnel struct {
	ID             string             `json:"id,omitempty"`
	Name           string             `json:"name,omitempty"`
	Secret         string             `json:"tunnel_secret,omitempty"`
	CreatedAt      *time.Time         `json:"created_at,omitempty"`
	DeletedAt      *time.Time         `json:"deleted_at,omitempty"`
	Connections    []TunnelConnection `json:"connections,omitempty"`
	ConnsActiveAt  *time.Time         `json:"conns_active_at,omitempty"`
	ConnInactiveAt *time.Time         `json:"conns_inactive_at,omitempty"`
}

// Connection is the struct definition of a connection.
type Connection struct {
	ID          string             `json:"id,omitempty"`
	Features    []string           `json:"features,omitempty"`
	Version     string             `json:"version,omitempty"`
	Arch        string             `json:"arch,omitempty"`
	Connections []TunnelConnection `json:"conns,omitempty"`
	RunAt       *time.Time         `json:"run_at,omitempty"`
}

// TunnelConnection represents the connections associated with a tunnel.
type TunnelConnection struct {
	ColoName           string `json:"colo_name"`
	ID                 string `json:"id"`
	IsPendingReconnect bool   `json:"is_pending_reconnect"`
	ClientID           string `json:"client_id"`
	ClientVersion      string `json:"client_version"`
	OpenedAt           string `json:"opened_at"`
	OriginIP           string `json:"origin_ip"`
}

// TunnelsDetailResponse is used for representing the API response payload for
// multiple tunnels.
type TunnelsDetailResponse struct {
	Result []Tunnel `json:"result"`
	Response
}

// TunnelDetailResponse is used for representing the API response payload for
// a single tunnel.
type TunnelDetailResponse struct {
	Result Tunnel `json:"result"`
	Response
}

// TunnelConnectionResponse is used for representing the API response payload for
// connections of a single tunnel.
type TunnelConnectionResponse struct {
	Result []Connection `json:"result"`
	Response
}

type TunnelConfigurationResult struct {
	TunnelID string              `json:"tunnel_id,omitempty"`
	Config   TunnelConfiguration `json:"config,omitempty"`
	Version  int                 `json:"version,omitempty"`
}

// TunnelConfigurationResponse is used for representing the API response payload
// for a single tunnel.
type TunnelConfigurationResponse struct {
	Result TunnelConfigurationResult `json:"result"`
	Response
}

// TunnelTokenResponse is  the API response for a tunnel token.
type TunnelTokenResponse struct {
	Result string `json:"result"`
	Response
}

type TunnelCreateParams struct {
	Name   string `json:"name,omitempty"`
	Secret string `json:"tunnel_secret,omitempty"`
}

type TunnelUpdateParams struct {
	Name   string `json:"name,omitempty"`
	Secret string `json:"tunnel_secret,omitempty"`
}

type UnvalidatedIngressRule struct {
	Hostname string `json:"hostname,omitempty"`
	Path     string `json:"path,omitempty"`
	Service  string `json:"service,omitempty"`
}

// OriginRequestConfig is a set of optional fields that users may set to
// customize how cloudflared sends requests to origin services. It is used to set
// up general config that apply to all rules, and also, specific per-rule
// config.
type OriginRequestConfig struct {
	// HTTP proxy timeout for establishing a new connection
	ConnectTimeout *time.Duration `json:"connectTimeout,omitempty"`
	// HTTP proxy timeout for completing a TLS handshake
	TLSTimeout *time.Duration `json:"tlsTimeout,omitempty"`
	// HTTP proxy TCP keepalive duration
	TCPKeepAlive *time.Duration `json:"tcpKeepAlive,omitempty"`
	// HTTP proxy should disable "happy eyeballs" for IPv4/v6 fallback
	NoHappyEyeballs *bool `json:"noHappyEyeballs,omitempty"`
	// HTTP proxy maximum keepalive connection pool size
	KeepAliveConnections *int `json:"keepAliveConnections,omitempty"`
	// HTTP proxy timeout for closing an idle connection
	KeepAliveTimeout *time.Duration `json:"keepAliveTimeout,omitempty"`
	// Sets the HTTP Host header for the local webserver.
	HTTPHostHeader *string `json:"httpHostHeader,omitempty"`
	// Hostname on the origin server certificate.
	OriginServerName *string `json:"originServerName,omitempty"`
	// Path to the CA for the certificate of your origin.
	// This option should be used only if your certificate is not signed by Cloudflare.
	CAPool *string `json:"caPool,omitempty"`
	// Disables TLS verification of the certificate presented by your origin.
	// Will allow any certificate from the origin to be accepted.
	// Note: The connection from your machine to Cloudflare's Edge is still encrypted.
	NoTLSVerify *bool `json:"noTLSVerify,omitempty"`
	// Disables chunked transfer encoding.
	// Useful if you are running a WSGI server.
	DisableChunkedEncoding *bool `json:"disableChunkedEncoding,omitempty"`
	// Runs as jump host
	BastionMode *bool `json:"bastionMode,omitempty"`
	// Listen address for the proxy.
	ProxyAddress *string `json:"proxyAddress,omitempty"`
	// Listen port for the proxy.
	ProxyPort *uint `json:"proxyPort,omitempty"`
	// Valid options are 'socks' or empty.
	ProxyType *string `json:"proxyType,omitempty"`
	// IP rules for the proxy service
	IPRules []IngressIPRule `json:"ipRules,omitempty"`
}

type IngressIPRule struct {
	Prefix *string `json:"prefix,omitempty"`
	Ports  []int   `json:"ports,omitempty"`
	Allow  bool    `json:"allow,omitempty"`
}

type TunnelConfiguration struct {
	Ingress       []UnvalidatedIngressRule `json:"ingress,omitempty"`
	WarpRouting   *WarpRoutingConfig       `json:"warp-routing,omitempty"`
	OriginRequest OriginRequestConfig      `json:"originRequest,omitempty"`
}

type WarpRoutingConfig struct {
	Enabled bool `json:"enabled,omitempty"`
}

type TunnelConfigurationParams struct {
	TunnelID string              `json:"-"`
	Config   TunnelConfiguration `json:"config,omitempty"`
}

type TunnelListParams struct {
	Name      string     `url:"name,omitempty"`
	UUID      string     `url:"uuid,omitempty"` // the tunnel ID
	IsDeleted *bool      `url:"is_deleted,omitempty"`
	ExistedAt *time.Time `url:"existed_at,omitempty"`
}

// Tunnels lists all tunnels.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-list-cloudflare-tunnels
func (api *API) Tunnels(ctx context.Context, rc *ResourceContainer, params TunnelListParams) ([]Tunnel, error) {
	if rc.Identifier == "" {
		return []Tunnel{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/cfd_tunnel", rc.Identifier), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Tunnel{}, err
	}

	var argoDetailsResponse TunnelsDetailResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return []Tunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return argoDetailsResponse.Result, nil
}

// Tunnel returns a single Argo tunnel.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-get-cloudflare-tunnel
func (api *API) Tunnel(ctx context.Context, rc *ResourceContainer, tunnelID string) (Tunnel, error) {
	if rc.Identifier == "" {
		return Tunnel{}, ErrMissingAccountID
	}

	if tunnelID == "" {
		return Tunnel{}, errors.New("missing tunnel ID")
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel/%s", rc.Identifier, tunnelID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Tunnel{}, err
	}

	var argoDetailsResponse TunnelDetailResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return Tunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return argoDetailsResponse.Result, nil
}

// CreateTunnel creates a new tunnel for the account.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-create-cloudflare-tunnel
func (api *API) CreateTunnel(ctx context.Context, rc *ResourceContainer, params TunnelCreateParams) (Tunnel, error) {
	if rc.Identifier == "" {
		return Tunnel{}, ErrMissingAccountID
	}

	if params.Name == "" {
		return Tunnel{}, errors.New("missing tunnel name")
	}

	if params.Secret == "" {
		return Tunnel{}, errors.New("missing tunnel secret")
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel", rc.Identifier)

	tunnel := Tunnel{Name: params.Name, Secret: params.Secret}

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, tunnel)
	if err != nil {
		return Tunnel{}, err
	}

	var argoDetailsResponse TunnelDetailResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return Tunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return argoDetailsResponse.Result, nil
}

// UpdateTunnel updates an existing tunnel for the account.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-update-cloudflare-tunnel
func (api *API) UpdateTunnel(ctx context.Context, rc *ResourceContainer, params TunnelUpdateParams) (Tunnel, error) {
	if rc.Identifier == "" {
		return Tunnel{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel", rc.Identifier)

	var tunnel Tunnel

	if params.Name != "" {
		tunnel.Name = params.Name
	}

	if params.Secret != "" {
		tunnel.Secret = params.Secret
	}

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, tunnel)
	if err != nil {
		return Tunnel{}, err
	}

	var argoDetailsResponse TunnelDetailResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return Tunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return argoDetailsResponse.Result, nil
}

// UpdateTunnelConfiguration updates an existing tunnel for the account.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-configuration-properties
func (api *API) UpdateTunnelConfiguration(ctx context.Context, rc *ResourceContainer, params TunnelConfigurationParams) (TunnelConfigurationResult, error) {
	if rc.Identifier == "" {
		return TunnelConfigurationResult{}, ErrMissingAccountID
	}

	if params.TunnelID == "" {
		return TunnelConfigurationResult{}, ErrMissingTunnelID
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel/%s/configurations", rc.Identifier, params.TunnelID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return TunnelConfigurationResult{}, err
	}

	var tunnelDetailsResponse TunnelConfigurationResponse
	err = json.Unmarshal(res, &tunnelDetailsResponse)
	if err != nil {
		return TunnelConfigurationResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	var tunnelDetails TunnelConfigurationResult

	tunnelDetails.Config = tunnelDetailsResponse.Result.Config
	tunnelDetails.TunnelID = tunnelDetailsResponse.Result.TunnelID
	tunnelDetails.Version = tunnelDetailsResponse.Result.Version

	return tunnelDetails, nil
}

// GetTunnelConfiguration updates an existing tunnel for the account.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-configuration-properties
func (api *API) GetTunnelConfiguration(ctx context.Context, rc *ResourceContainer, tunnelID string) (TunnelConfigurationResult, error) {
	if rc.Identifier == "" {
		return TunnelConfigurationResult{}, ErrMissingAccountID
	}

	if tunnelID == "" {
		return TunnelConfigurationResult{}, ErrMissingTunnelID
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel/%s/configurations", rc.Identifier, tunnelID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TunnelConfigurationResult{}, err
	}

	var tunnelDetailsResponse TunnelConfigurationResponse
	err = json.Unmarshal(res, &tunnelDetailsResponse)
	if err != nil {
		return TunnelConfigurationResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	var tunnelDetails TunnelConfigurationResult

	tunnelDetails.Config = tunnelDetailsResponse.Result.Config
	tunnelDetails.TunnelID = tunnelDetailsResponse.Result.TunnelID
	tunnelDetails.Version = tunnelDetailsResponse.Result.Version

	return tunnelDetails, nil
}

// TunnelConnections gets all connections on a tunnel.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-list-cloudflare-tunnel-connections
func (api *API) TunnelConnections(ctx context.Context, rc *ResourceContainer, tunnelID string) ([]Connection, error) {
	if rc.Identifier == "" {
		return []Connection{}, ErrMissingAccountID
	}

	if tunnelID == "" {
		return []Connection{}, ErrMissingTunnelID
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel/%s/connections", rc.Identifier, tunnelID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Connection{}, err
	}

	var argoDetailsResponse TunnelConnectionResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return []Connection{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return argoDetailsResponse.Result, nil
}

// DeleteTunnel removes a single Argo tunnel.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-delete-cloudflare-tunnel
func (api *API) DeleteTunnel(ctx context.Context, rc *ResourceContainer, tunnelID string) error {
	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel/%s", rc.Identifier, tunnelID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var argoDetailsResponse TunnelDetailResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// CleanupTunnelConnections deletes any inactive connections on a tunnel.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-clean-up-cloudflare-tunnel-connections
func (api *API) CleanupTunnelConnections(ctx context.Context, rc *ResourceContainer, tunnelID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if tunnelID == "" {
		return errors.New("missing tunnel ID")
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel/%s/connections", rc.Identifier, tunnelID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var argoDetailsResponse TunnelDetailResponse
	err = json.Unmarshal(res, &argoDetailsResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// TunnelToken that allows to run a tunnel.
//
// API reference: https://api.cloudflare.com/#cloudflare-tunnel-get-cloudflare-tunnel-token
func (api *API) TunnelToken(ctx context.Context, rc *ResourceContainer, tunnelID string) (string, error) {
	if rc.Identifier == "" {
		return "", ErrMissingAccountID
	}

	if tunnelID == "" {
		return "", errors.New("missing tunnel ID")
	}

	uri := fmt.Sprintf("/accounts/%s/cfd_tunnel/%s/token", rc.Identifier, tunnelID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return "", err
	}

	var tunnelTokenResponse TunnelTokenResponse
	err = json.Unmarshal(res, &tunnelTokenResponse)
	if err != nil {
		return "", fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return tunnelTokenResponse.Result, nil
}
