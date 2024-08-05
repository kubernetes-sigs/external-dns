package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// RegionalTieredCache is the structure of the API object for the regional tiered cache
// setting.
type RegionalTieredCache struct {
	ID         string    `json:"id,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
	Value      string    `json:"value"`
}

// RegionalTieredCacheDetailsResponse is the API response for the regional tiered cache
// setting.
type RegionalTieredCacheDetailsResponse struct {
	Result RegionalTieredCache `json:"result"`
	Response
}

type zoneRegionalTieredCacheSingleResponse struct {
	Response
	Result RegionalTieredCache `json:"result"`
}

type GetRegionalTieredCacheParams struct{}

type UpdateRegionalTieredCacheParams struct {
	Value string `json:"value"`
}

// GetRegionalTieredCache returns information about the current regional tiered
// cache settings.
//
// API reference: https://developers.cloudflare.com/api/operations/zone-cache-settings-get-regional-tiered-cache-setting
func (api *API) GetRegionalTieredCache(ctx context.Context, rc *ResourceContainer, params GetRegionalTieredCacheParams) (RegionalTieredCache, error) {
	if rc.Level != ZoneRouteLevel {
		return RegionalTieredCache{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/cache/regional_tiered_cache", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return RegionalTieredCache{}, err
	}

	var RegionalTieredCacheDetailsResponse RegionalTieredCacheDetailsResponse
	err = json.Unmarshal(res, &RegionalTieredCacheDetailsResponse)
	if err != nil {
		return RegionalTieredCache{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return RegionalTieredCacheDetailsResponse.Result, nil
}

// UpdateRegionalTieredCache updates the regional tiered cache setting for a
// zone.
//
// API reference: https://developers.cloudflare.com/api/operations/zone-cache-settings-change-regional-tiered-cache-setting
func (api *API) UpdateRegionalTieredCache(ctx context.Context, rc *ResourceContainer, params UpdateRegionalTieredCacheParams) (RegionalTieredCache, error) {
	if rc.Level != ZoneRouteLevel {
		return RegionalTieredCache{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/cache/regional_tiered_cache", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return RegionalTieredCache{}, err
	}

	response := &zoneRegionalTieredCacheSingleResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return RegionalTieredCache{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}
