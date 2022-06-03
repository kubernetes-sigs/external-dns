package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TeamsProxyEndpointListResponse struct {
	Response
	ResultInfo `json:"result_info"`
	Result     []TeamsProxyEndpoint `json:"result"`
}
type TeamsProxyEndpointDetailResponse struct {
	Response
	Result TeamsProxyEndpoint `json:"result"`
}

type TeamsProxyEndpoint struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	IPs       []string   `json:"ips"`
	Subdomain string     `json:"subdomain"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// TeamsProxyEndpoint returns a single proxy endpoints within an account.
//
// API reference: https://api.cloudflare.com/#zero-trust-gateway-proxy-endpoints-proxy-endpoint-details
func (api *API) TeamsProxyEndpoint(ctx context.Context, accountID, proxyEndpointID string) (TeamsProxyEndpoint, error) {
	uri := fmt.Sprintf("/%s/%s/gateway/proxy_endpoints/%s", AccountRouteRoot, accountID, proxyEndpointID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsProxyEndpoint{}, err
	}

	var teamsProxyEndpointDetailResponse TeamsProxyEndpointDetailResponse
	err = json.Unmarshal(res, &teamsProxyEndpointDetailResponse)
	if err != nil {
		return TeamsProxyEndpoint{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsProxyEndpointDetailResponse.Result, nil
}

// TeamsProxyEndpoints returns all proxy endpoints within an account.
//
// API reference: https://api.cloudflare.com/#zero-trust-gateway-proxy-endpoints-list-proxy-endpoints
func (api *API) TeamsProxyEndpoints(ctx context.Context, accountID string) ([]TeamsProxyEndpoint, ResultInfo, error) {
	uri := fmt.Sprintf("/%s/%s/gateway/proxy_endpoints", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsProxyEndpoint{}, ResultInfo{}, err
	}

	var teamsProxyEndpointListResponse TeamsProxyEndpointListResponse
	err = json.Unmarshal(res, &teamsProxyEndpointListResponse)
	if err != nil {
		return []TeamsProxyEndpoint{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsProxyEndpointListResponse.Result, teamsProxyEndpointListResponse.ResultInfo, nil
}

// CreateTeamsProxyEndpoint creates a new proxy endpoint.
//
// API reference: https://api.cloudflare.com/#zero-trust-gateway-proxy-endpoints-create-proxy-endpoint
func (api *API) CreateTeamsProxyEndpoint(ctx context.Context, accountID string, proxyEndpoint TeamsProxyEndpoint) (TeamsProxyEndpoint, error) {
	uri := fmt.Sprintf("/%s/%s/gateway/proxy_endpoints", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, proxyEndpoint)
	if err != nil {
		return TeamsProxyEndpoint{}, err
	}

	var teamsProxyEndpointDetailResponse TeamsProxyEndpointDetailResponse
	err = json.Unmarshal(res, &teamsProxyEndpointDetailResponse)
	if err != nil {
		return TeamsProxyEndpoint{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsProxyEndpointDetailResponse.Result, nil
}

// UpdateTeamsProxyEndpoint updates an existing teams Proxy Endpoint.
//
// API reference: https://api.cloudflare.com/#zero-trust-gateway-proxy-endpoints-update-proxy-endpoint
func (api *API) UpdateTeamsProxyEndpoint(ctx context.Context, accountID string, proxyEndpoint TeamsProxyEndpoint) (TeamsProxyEndpoint, error) {
	if proxyEndpoint.ID == "" {
		return TeamsProxyEndpoint{}, fmt.Errorf("Proxy Endpoint ID cannot be empty")
	}

	uri := fmt.Sprintf(
		"/%s/%s/gateway/proxy_endpoints/%s",
		AccountRouteRoot,
		accountID,
		proxyEndpoint.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, proxyEndpoint)
	if err != nil {
		return TeamsProxyEndpoint{}, err
	}

	var teamsProxyEndpointDetailResponse TeamsProxyEndpointDetailResponse
	err = json.Unmarshal(res, &teamsProxyEndpointDetailResponse)
	if err != nil {
		return TeamsProxyEndpoint{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsProxyEndpointDetailResponse.Result, nil
}

// DeleteTeamsProxyEndpoint deletes a teams Proxy Endpoint.
//
// API reference: https://api.cloudflare.com/#zero-trust-gateway-proxy-endpoints-delete-proxy-endpoint
func (api *API) DeleteTeamsProxyEndpoint(ctx context.Context, accountID, proxyEndpointID string) error {
	uri := fmt.Sprintf(
		"/%s/%s/gateway/proxy_endpoints/%s",
		AccountRouteRoot,
		accountID,
		proxyEndpointID,
	)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
