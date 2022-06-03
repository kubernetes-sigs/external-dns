package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"errors"
)

var (
	ErrMissingNetwork      = errors.New("missing required network parameter")
	ErrInvalidNetworkValue = errors.New("invalid IP parameter. Cannot use CIDR ranges for this endpoint.")
)

// TunnelRoute is the full record for a route.
type TunnelRoute struct {
	Network          string     `json:"network"`
	TunnelID         string     `json:"tunnel_id"`
	TunnelName       string     `json:"tunnel_name"`
	Comment          string     `json:"comment"`
	CreatedAt        *time.Time `json:"created_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
	VirtualNetworkID string     `json:"virtual_network_id"`
}

type TunnelRoutesListParams struct {
	TunnelID         string     `url:"tunnel_id,omitempty"`
	Comment          string     `url:"comment,omitempty"`
	IsDeleted        *bool      `url:"is_deleted,omitempty"`
	NetworkSubset    string     `url:"network_subset,omitempty"`
	NetworkSuperset  string     `url:"network_superset,omitempty"`
	ExistedAt        *time.Time `url:"existed_at,omitempty"`
	VirtualNetworkID string     `url:"virtual_network_id,omitempty"`
	PaginationOptions
}

type TunnelRoutesCreateParams struct {
	Network          string `json:"-"`
	TunnelID         string `json:"tunnel_id"`
	Comment          string `json:"comment,omitempty"`
	VirtualNetworkID string `json:"virtual_network_id,omitempty"`
}

type TunnelRoutesUpdateParams struct {
	Network          string `json:"network"`
	TunnelID         string `json:"tunnel_id"`
	Comment          string `json:"comment,omitempty"`
	VirtualNetworkID string `json:"virtual_network_id,omitempty"`
}

type TunnelRoutesForIPParams struct {
	Network          string `url:"-"`
	VirtualNetworkID string `url:"virtual_network_id,omitempty"`
}

type TunnelRoutesDeleteParams struct {
	Network          string `url:"-"`
	VirtualNetworkID string `url:"virtual_network_id,omitempty"`
}

// tunnelRouteListResponse is the API response for listing tunnel routes.
type tunnelRouteListResponse struct {
	Response
	Result []TunnelRoute `json:"result"`
}

type tunnelRouteResponse struct {
	Response
	Result TunnelRoute `json:"result"`
}

// ListTunnelRoutes lists all defined routes for tunnels in the account.
//
// See: https://api.cloudflare.com/#tunnel-route-list-tunnel-routes
func (api *API) ListTunnelRoutes(ctx context.Context, rc *ResourceContainer, params TunnelRoutesListParams) ([]TunnelRoute, error) {
	if rc.Identifier == "" {
		return []TunnelRoute{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/teamnet/routes", AccountRouteRoot, rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TunnelRoute{}, err
	}

	var resp tunnelRouteListResponse
	err = json.Unmarshal(res, &resp)
	if err != nil {
		return []TunnelRoute{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return resp.Result, nil
}

// GetTunnelRouteForIP finds the Tunnel Route that encompasses the given IP.
//
// See: https://api.cloudflare.com/#tunnel-route-get-tunnel-route-by-ip
func (api *API) GetTunnelRouteForIP(ctx context.Context, rc *ResourceContainer, params TunnelRoutesForIPParams) (TunnelRoute, error) {
	if rc.Identifier == "" {
		return TunnelRoute{}, ErrMissingAccountID
	}

	if params.Network == "" {
		return TunnelRoute{}, ErrMissingNetwork
	}

	if strings.Contains(params.Network, "/") {
		return TunnelRoute{}, ErrInvalidNetworkValue
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/teamnet/routes/ip/%s", AccountRouteRoot, rc.Identifier, params.Network), params)

	responseBody, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return TunnelRoute{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return routeResponse.Result, nil
}

// CreateTunnelRoute add a new route to the account routing table for the given
// tunnel.
//
// See: https://api.cloudflare.com/#tunnel-route-create-route
func (api *API) CreateTunnelRoute(ctx context.Context, rc *ResourceContainer, params TunnelRoutesCreateParams) (TunnelRoute, error) {
	if rc.Identifier == "" {
		return TunnelRoute{}, ErrMissingAccountID
	}

	if params.Network == "" {
		return TunnelRoute{}, ErrMissingNetwork
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, rc.Identifier, url.PathEscape(params.Network))

	responseBody, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return TunnelRoute{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return routeResponse.Result, nil
}

// DeleteTunnelRoute delete an existing route from the account routing table.
//
// See: https://api.cloudflare.com/#tunnel-route-delete-route
func (api *API) DeleteTunnelRoute(ctx context.Context, rc *ResourceContainer, params TunnelRoutesDeleteParams) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	if params.Network == "" {
		return ErrMissingNetwork
	}

	// Cannot fully utilize buildURI here because it tries to escape "%" sign
	// from the already escaped "/" sign from Network field.
	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s%s", AccountRouteRoot, rc.Identifier, url.PathEscape(params.Network), buildURI("", params))

	responseBody, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// UpdateTunnelRoute updates an existing route in the account routing table for
// the given tunnel.
//
// See: https://api.cloudflare.com/#tunnel-route-update-route
func (api *API) UpdateTunnelRoute(ctx context.Context, rc *ResourceContainer, params TunnelRoutesUpdateParams) (TunnelRoute, error) {
	if rc.Identifier == "" {
		return TunnelRoute{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/%s/%s/teamnet/routes/network/%s", AccountRouteRoot, rc.Identifier, url.PathEscape(params.Network))

	responseBody, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return TunnelRoute{}, err
	}

	var routeResponse tunnelRouteResponse
	err = json.Unmarshal(responseBody, &routeResponse)
	if err != nil {
		return TunnelRoute{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return routeResponse.Result, nil
}
