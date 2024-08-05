package cloudflare

import (
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
)

// SplitTunnelResponse represents the response from the get split
// tunnel endpoints.
type SplitTunnelResponse struct {
	Response
	Result []SplitTunnel `json:"result"`
}

// SplitTunnel represents the individual tunnel struct.
type SplitTunnel struct {
	Address     string `json:"address,omitempty"`
	Host        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
}

// ListSplitTunnel returns all include or exclude split tunnel  within an account.
//
// API reference for include: https://api.cloudflare.com/#device-policy-get-split-tunnel-include-list
// API reference for exclude: https://api.cloudflare.com/#device-policy-get-split-tunnel-exclude-list
func (api *API) ListSplitTunnels(ctx context.Context, accountID string, mode string) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/%s", AccountRouteRoot, accountID, mode)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelResponse SplitTunnelResponse
	err = json.Unmarshal(res, &splitTunnelResponse)
	if err != nil {
		return []SplitTunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return splitTunnelResponse.Result, nil
}

// UpdateSplitTunnel updates the existing split tunnel policy.
//
// API reference for include: https://api.cloudflare.com/#device-policy-set-split-tunnel-include-list
// API reference for exclude: https://api.cloudflare.com/#device-policy-set-split-tunnel-exclude-list
func (api *API) UpdateSplitTunnel(ctx context.Context, accountID string, mode string, tunnels []SplitTunnel) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/%s", AccountRouteRoot, accountID, mode)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// SplitTunnelResponse represents the response from the get split
// tunnel endpoints.
type SplitTunnelResponse struct {
	Response
	Result []SplitTunnel `json:"result"`
}

// SplitTunnel represents the individual tunnel struct.
type SplitTunnel struct {
	Address     string `json:"address,omitempty"`
	Host        string `json:"host,omitempty"`
	Description string `json:"description,omitempty"`
}

// ListSplitTunnel returns all include or exclude split tunnel  within an account.
//
// API reference for include: https://api.cloudflare.com/#device-policy-get-split-tunnel-include-list
// API reference for exclude: https://api.cloudflare.com/#device-policy-get-split-tunnel-exclude-list
func (api *API) ListSplitTunnels(ctx context.Context, accountID string, mode string) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/%s", AccountRouteRoot, accountID, mode)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelResponse SplitTunnelResponse
	err = json.Unmarshal(res, &splitTunnelResponse)
	if err != nil {
		return []SplitTunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return splitTunnelResponse.Result, nil
}

// UpdateSplitTunnel updates the existing split tunnel policy.
//
// API reference for include: https://api.cloudflare.com/#device-policy-set-split-tunnel-include-list
// API reference for exclude: https://api.cloudflare.com/#device-policy-set-split-tunnel-exclude-list
func (api *API) UpdateSplitTunnel(ctx context.Context, accountID string, mode string, tunnels []SplitTunnel) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/%s", AccountRouteRoot, accountID, mode)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, tunnels)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelResponse SplitTunnelResponse
	err = json.Unmarshal(res, &splitTunnelResponse)
	if err != nil {
		return []SplitTunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return splitTunnelResponse.Result, nil
}

// ListSplitTunnelDeviceSettingsPolicy returns all include or exclude split tunnel within a device settings policy
//
// API reference for include: https://api.cloudflare.com/#device-policy-get-split-tunnel-include-list
// API reference for exclude: https://api.cloudflare.com/#device-policy-get-split-tunnel-exclude-list
func (api *API) ListSplitTunnelsDeviceSettingsPolicy(ctx context.Context, accountID, policyID string, mode string) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/%s/%s", AccountRouteRoot, accountID, policyID, mode)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelResponse SplitTunnelResponse
	err = json.Unmarshal(res, &splitTunnelResponse)
	if err != nil {
		return []SplitTunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return splitTunnelResponse.Result, nil
}

// UpdateSplitTunnelDeviceSettingsPolicy updates the existing split tunnel policy within a device settings policy
//
// API reference for include: https://api.cloudflare.com/#device-policy-set-split-tunnel-include-list
// API reference for exclude: https://api.cloudflare.com/#device-policy-set-split-tunnel-exclude-list
func (api *API) UpdateSplitTunnelDeviceSettingsPolicy(ctx context.Context, accountID, policyID string, mode string, tunnels []SplitTunnel) ([]SplitTunnel, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/%s/%s", AccountRouteRoot, accountID, policyID, mode)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, tunnels)
	if err != nil {
		return []SplitTunnel{}, err
	}

	var splitTunnelResponse SplitTunnelResponse
	err = json.Unmarshal(res, &splitTunnelResponse)
	if err != nil {
		return []SplitTunnel{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return splitTunnelResponse.Result, nil
}
