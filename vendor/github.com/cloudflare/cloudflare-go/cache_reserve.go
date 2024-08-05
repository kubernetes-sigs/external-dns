package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// CacheReserve is the structure of the API object for the cache reserve
// setting.
type CacheReserve struct {
	ID         string    `json:"id,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
	Value      string    `json:"value"`
}

// CacheReserveDetailsResponse is the API response for the cache reserve
// setting.
type CacheReserveDetailsResponse struct {
	Result CacheReserve `json:"result"`
	Response
}

type zoneCacheReserveSingleResponse struct {
	Response
	Result CacheReserve `json:"result"`
}

type GetCacheReserveParams struct{}

type UpdateCacheReserveParams struct {
	Value string `json:"value"`
}

// GetCacheReserve returns information about the current cache reserve settings.
//
// API reference: https://developers.cloudflare.com/api/operations/zone-cache-settings-get-cache-reserve-setting
func (api *API) GetCacheReserve(ctx context.Context, rc *ResourceContainer, params GetCacheReserveParams) (CacheReserve, error) {
	if rc.Level != ZoneRouteLevel {
		return CacheReserve{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/cache/cache_reserve", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return CacheReserve{}, err
	}

	var cacheReserveDetailsResponse CacheReserveDetailsResponse
	err = json.Unmarshal(res, &cacheReserveDetailsResponse)
	if err != nil {
		return CacheReserve{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return cacheReserveDetailsResponse.Result, nil
}

// UpdateCacheReserve updates the cache reserve setting for a zone
//
// API reference: https://developers.cloudflare.com/api/operations/zone-cache-settings-change-cache-reserve-setting
func (api *API) UpdateCacheReserve(ctx context.Context, rc *ResourceContainer, params UpdateCacheReserveParams) (CacheReserve, error) {
	if rc.Level != ZoneRouteLevel {
		return CacheReserve{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/cache/cache_reserve", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return CacheReserve{}, err
	}

	response := &zoneCacheReserveSingleResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CacheReserve{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}
