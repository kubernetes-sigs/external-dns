package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TeamsDevicesList struct {
	Response
	Result []TeamsDeviceListItem `json:"result"`
}

type TeamsDeviceDetail struct {
	Response
	Result TeamsDeviceListItem `json:"result"`
}

type TeamsDeviceListItem struct {
	User             UserItem `json:"user,omitempty"`
	ID               string   `json:"id,omitempty"`
	Key              string   `json:"key,omitempty"`
	DeviceType       string   `json:"device_type,omitempty"`
	Name             string   `json:"name,omitempty"`
	Model            string   `json:"model,omitempty"`
	Manufacturer     string   `json:"manufacturer,omitempty"`
	Deleted          bool     `json:"deleted,omitempty"`
	Version          string   `json:"version,omitempty"`
	SerialNumber     string   `json:"serial_number,omitempty"`
	OSVersion        string   `json:"os_version,omitempty"`
	OSDistroName     string   `json:"os_distro_name,omitempty"`
	OsDistroRevision string   `json:"os_distro_revision,omitempty"`
	MacAddress       string   `json:"mac_address,omitempty"`
	IP               string   `json:"ip,omitempty"`
	Created          string   `json:"created,omitempty"`
	Updated          string   `json:"updated,omitempty"`
	LastSeen         string   `json:"last_seen,omitempty"`
	RevokedAt        string   `json:"revoked_at,omitempty"`
}

type UserItem struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// ListTeamsDevice returns all devices for a given account.
//
// API reference : https://api.cloudflare.com/#devices-list-devices
func (api *API) ListTeamsDevices(ctx context.Context, accountID string) ([]TeamsDeviceListItem, error) {
	uri := fmt.Sprintf("/%s/%s/devices", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsDeviceListItem{}, err
	}

	var response TeamsDevicesList
	err = json.Unmarshal(res, &response)
	if err != nil {
		return []TeamsDeviceListItem{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// RevokeTeamsDevice revokes device with given identifiers.
//
// API reference : https://api.cloudflare.com/#devices-revoke-devices
func (api *API) RevokeTeamsDevices(ctx context.Context, accountID string, deviceIds []string) (Response, error) {
	uri := fmt.Sprintf("/%s/%s/devices/revoke", AccountRouteRoot, accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, deviceIds)
	if err != nil {
		return Response{}, err
	}

	result := Response{}
	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// GetTeamsDeviceDetails gets device details.
//
// API reference : https://api.cloudflare.com/#devices-device-details
func (api *API) GetTeamsDeviceDetails(ctx context.Context, accountID string, deviceID string) (TeamsDeviceListItem, error) {
	uri := fmt.Sprintf("/%s/%s/devices/%s", AccountRouteRoot, accountID, deviceID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsDeviceListItem{}, err
	}

	var response TeamsDeviceDetail
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TeamsDeviceListItem{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}
