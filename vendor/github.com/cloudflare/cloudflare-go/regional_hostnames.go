package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type Region struct {
	Key   string `json:"key"`
	Label string `json:"label"`
}

type RegionalHostname struct {
	Hostname  string     `json:"hostname"`
	RegionKey string     `json:"region_key"`
	CreatedOn *time.Time `json:"created_on,omitempty"`
}

// regionalHostnameResponse contains an API Response from a Create, Get, Update, or Delete call.
type regionalHostnameResponse struct {
	Response
	Result RegionalHostname `json:"result"`
}

type ListDataLocalizationRegionsParams struct{}
type ListDataLocalizationRegionalHostnamesParams struct{}

type CreateDataLocalizationRegionalHostnameParams struct {
	Hostname  string `json:"hostname"`
	RegionKey string `json:"region_key"`
}

type UpdateDataLocalizationRegionalHostnameParams struct {
	Hostname  string `json:"-"`
	RegionKey string `json:"region_key"`
}

// ListDataLocalizationRegions lists all available regions.
//
// API reference: https://developers.cloudflare.com/data-localization/regional-services/get-started/#configure-regional-services-via-api
func (api *API) ListDataLocalizationRegions(ctx context.Context, rc *ResourceContainer, params ListDataLocalizationRegionsParams) ([]Region, error) {
	if rc.Level != AccountRouteLevel {
		return []Region{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return []Region{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/addressing/regional_hostnames/regions", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Region{}, err
	}
	result := struct {
		Result []Region `json:"result"`
	}{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []Region{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// ListDataLocalizationRegionalHostnames lists all regional hostnames for a zone.
//
// API reference: https://developers.cloudflare.com/data-localization/regional-services/get-started/#configure-regional-services-via-api
func (api *API) ListDataLocalizationRegionalHostnames(ctx context.Context, rc *ResourceContainer, params ListDataLocalizationRegionalHostnamesParams) ([]RegionalHostname, error) {
	if rc.Level != ZoneRouteLevel {
		return []RegionalHostname{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return []RegionalHostname{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/addressing/regional_hostnames", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []RegionalHostname{}, err
	}
	result := struct {
		Result []RegionalHostname `json:"result"`
	}{}
	if err := json.Unmarshal(res, &result); err != nil {
		return []RegionalHostname{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// CreateDataLocalizationRegionalHostname lists all regional hostnames for a zone.
//
// API reference: https://developers.cloudflare.com/data-localization/regional-services/get-started/#configure-regional-services-via-api
func (api *API) CreateDataLocalizationRegionalHostname(ctx context.Context, rc *ResourceContainer, params CreateDataLocalizationRegionalHostnameParams) (RegionalHostname, error) {
	if rc.Level != ZoneRouteLevel {
		return RegionalHostname{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return RegionalHostname{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/addressing/regional_hostnames", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return RegionalHostname{}, err
	}
	result := regionalHostnameResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return RegionalHostname{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// GetDataLocalizationRegionalHostname returns the details of a specific regional hostname.
//
// API reference: https://developers.cloudflare.com/data-localization/regional-services/get-started/#configure-regional-services-via-api
func (api *API) GetDataLocalizationRegionalHostname(ctx context.Context, rc *ResourceContainer, hostname string) (RegionalHostname, error) {
	if rc.Level != ZoneRouteLevel {
		return RegionalHostname{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return RegionalHostname{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/addressing/regional_hostnames/%s", rc.Identifier, hostname)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return RegionalHostname{}, err
	}

	result := regionalHostnameResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return RegionalHostname{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// UpdateDataLocalizationRegionalHostname returns the details of a specific regional hostname.
//
// API reference: https://developers.cloudflare.com/data-localization/regional-services/get-started/#configure-regional-services-via-api
func (api *API) UpdateDataLocalizationRegionalHostname(ctx context.Context, rc *ResourceContainer, params UpdateDataLocalizationRegionalHostnameParams) (RegionalHostname, error) {
	if rc.Level != ZoneRouteLevel {
		return RegionalHostname{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return RegionalHostname{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/addressing/regional_hostnames/%s", rc.Identifier, params.Hostname)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return RegionalHostname{}, err
	}
	result := regionalHostnameResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return RegionalHostname{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return result.Result, nil
}

// DeleteDataLocalizationRegionalHostname deletes a regional hostname.
//
// API reference: https://developers.cloudflare.com/data-localization/regional-services/get-started/#configure-regional-services-via-api
func (api *API) DeleteDataLocalizationRegionalHostname(ctx context.Context, rc *ResourceContainer, hostname string) error {
	if rc.Level != ZoneRouteLevel {
		return fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/addressing/regional_hostnames/%s", rc.Identifier, hostname)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}
	return nil
}
