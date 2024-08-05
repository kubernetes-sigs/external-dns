package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type Config struct {
	TlsSockAddr string `json:"tls_sockaddr,omitempty"`
	Sha256      string `json:"sha256,omitempty"`
}

type DeviceManagedNetwork struct {
	NetworkID string  `json:"network_id,omitempty"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Config    *Config `json:"config"`
}

type DeviceManagedNetworkResponse struct {
	Response
	Result DeviceManagedNetwork `json:"result"`
}

type DeviceManagedNetworkListResponse struct {
	Response
	Result []DeviceManagedNetwork `json:"result"`
}

type ListDeviceManagedNetworksParams struct{}

type CreateDeviceManagedNetworkParams struct {
	NetworkID string  `json:"network_id,omitempty"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Config    *Config `json:"config"`
}

type UpdateDeviceManagedNetworkParams struct {
	NetworkID string  `json:"network_id,omitempty"`
	Type      string  `json:"type"`
	Name      string  `json:"name"`
	Config    *Config `json:"config"`
}

// ListDeviceManagedNetwork returns all Device Managed Networks for a given
// account.
//
// API reference : https://api.cloudflare.com/#device-managed-networks-list-device-managed-networks
func (api *API) ListDeviceManagedNetworks(ctx context.Context, rc *ResourceContainer, params ListDeviceManagedNetworksParams) ([]DeviceManagedNetwork, error) {
	if rc.Level != AccountRouteLevel {
		return []DeviceManagedNetwork{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/networks", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DeviceManagedNetwork{}, err
	}

	var response DeviceManagedNetworkListResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return []DeviceManagedNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// CreateDeviceManagedNetwork creates a new Device Managed Network.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-create-device-managed-network
func (api *API) CreateDeviceManagedNetwork(ctx context.Context, rc *ResourceContainer, params CreateDeviceManagedNetworkParams) (DeviceManagedNetwork, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceManagedNetwork{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/networks", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return DeviceManagedNetwork{}, err
	}

	var deviceManagedNetworksResponse DeviceManagedNetworkResponse
	if err := json.Unmarshal(res, &deviceManagedNetworksResponse); err != nil {
		return DeviceManagedNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceManagedNetworksResponse.Result, err
}

// UpdateDeviceManagedNetwork Update a Device Managed Network.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-update-device-managed-network
func (api *API) UpdateDeviceManagedNetwork(ctx context.Context, rc *ResourceContainer, params UpdateDeviceManagedNetworkParams) (DeviceManagedNetwork, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceManagedNetwork{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/networks/%s", rc.Level, rc.Identifier, params.NetworkID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return DeviceManagedNetwork{}, err
	}

	var deviceManagedNetworksResponse DeviceManagedNetworkResponse

	if err := json.Unmarshal(res, &deviceManagedNetworksResponse); err != nil {
		return DeviceManagedNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceManagedNetworksResponse.Result, err
}

// GetDeviceManagedNetwork gets a single Device Managed Network.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-device-managed-network-details
func (api *API) GetDeviceManagedNetwork(ctx context.Context, rc *ResourceContainer, networkID string) (DeviceManagedNetwork, error) {
	if rc.Level != AccountRouteLevel {
		return DeviceManagedNetwork{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/networks/%s", rc.Level, rc.Identifier, networkID)

	deviceManagedNetworksResponse := DeviceManagedNetworkResponse{}
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DeviceManagedNetwork{}, err
	}

	if err := json.Unmarshal(res, &deviceManagedNetworksResponse); err != nil {
		return DeviceManagedNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return deviceManagedNetworksResponse.Result, err
}

// DeleteManagedNetworks deletes a Device Managed Network.
//
// API reference: https://api.cloudflare.com/#device-managed-networks-delete-device-managed-network
func (api *API) DeleteManagedNetworks(ctx context.Context, rc *ResourceContainer, networkID string) ([]DeviceManagedNetwork, error) {
	if rc.Level != AccountRouteLevel {
		return []DeviceManagedNetwork{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/devices/networks/%s", rc.Level, rc.Identifier, networkID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return []DeviceManagedNetwork{}, err
	}

	var response DeviceManagedNetworkListResponse
	if err := json.Unmarshal(res, &response); err != nil {
		return []DeviceManagedNetwork{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, err
}
