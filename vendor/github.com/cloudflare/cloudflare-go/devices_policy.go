package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Enabled struct {
	Enabled bool `json:"enabled"`
}

// DeviceClientCertificatesZone identifies if the zero trust zone is configured for an account.
type DeviceClientCertificatesZone struct {
	Response
	Result Enabled
}

// UpdateDeviceClientCertificates controls the zero trust zone used to provision client certificates.
//
// API reference: https://api.cloudflare.com/#device-client-certificates
func (api *API) UpdateDeviceClientCertificatesZone(ctx context.Context, zoneID string, enable bool) (DeviceClientCertificatesZone, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/certificates", ZoneRouteRoot, zoneID)

	result := DeviceClientCertificatesZone{}
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, Enabled{enable})
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}

// GetDeviceClientCertificatesZone controls the zero trust zone used to provision client certificates.
//
// API reference: https://api.cloudflare.com/#device-client-certificates
func (api *API) GetDeviceClientCertificatesZone(ctx context.Context, zoneID string) (DeviceClientCertificatesZone, error) {
	uri := fmt.Sprintf("/%s/%s/devices/policy/certificates", ZoneRouteRoot, zoneID)

	result := DeviceClientCertificatesZone{}
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(res, &result); err != nil {
		return result, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result, err
}
