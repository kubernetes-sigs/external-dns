package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type TieredCacheType int

const (
	TieredCacheOff     TieredCacheType = 0
	TieredCacheGeneric TieredCacheType = 1
	TieredCacheSmart   TieredCacheType = 2
)

func (e TieredCacheType) String() string {
	switch e {
	case TieredCacheGeneric:
		return "generic"
	case TieredCacheSmart:
		return "smart"
	case TieredCacheOff:
		return "off"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type TieredCache struct {
	Type         TieredCacheType
	LastModified time.Time
}

// GetTieredCache allows you to retrieve the current Tiered Cache Settings for a Zone.
// This function does not support custom topologies, only Generic and Smart Tiered Caching.
//
// API Reference: https://api.cloudflare.com/#smart-tiered-cache-get-smart-tiered-cache-setting
// API Reference: https://api.cloudflare.com/#tiered-cache-get-tiered-cache-setting
func (api *API) GetTieredCache(ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	var lastModified time.Time

	generic, err := getGenericTieredCache(api, ctx, rc)
	if err != nil {
		return TieredCache{}, err
	}
	lastModified = generic.LastModified

	smart, err := getSmartTieredCache(api, ctx, rc)
	if err != nil {
		return TieredCache{}, err
	}

	if smart.LastModified.After(lastModified) {
		lastModified = smart.LastModified
	}

	if generic.Type == TieredCacheOff {
		return TieredCache{Type: TieredCacheOff, LastModified: lastModified}, nil
	}

	if smart.Type == TieredCacheOff {
		return TieredCache{Type: TieredCacheGeneric, LastModified: lastModified}, nil
	}

	return TieredCache{Type: TieredCacheSmart, LastModified: lastModified}, nil
}

// SetTieredCache allows you to set a zone's tiered cache topology between the available types.
// Using the value of TieredCacheOff will disable Tiered Cache entirely.
//
// API Reference: https://api.cloudflare.com/#smart-tiered-cache-patch-smart-tiered-cache-setting
// API Reference: https://api.cloudflare.com/#tiered-cache-patch-tiered-cache-setting
func (api *API) SetTieredCache(ctx context.Context, rc *ResourceContainer, value TieredCacheType) (TieredCache, error) {
	if value == TieredCacheOff {
		return api.DeleteTieredCache(ctx, rc)
	}

	var lastModified time.Time

	if value == TieredCacheGeneric {
		result, err := deleteSmartTieredCache(api, ctx, rc)
		if err != nil {
			return TieredCache{}, err
		}
		lastModified = result.LastModified

		result, err = enableGenericTieredCache(api, ctx, rc)
		if err != nil {
			return TieredCache{}, err
		}

		if result.LastModified.After(lastModified) {
			lastModified = result.LastModified
		}
		return TieredCache{Type: TieredCacheGeneric, LastModified: lastModified}, nil
	}

	result, err := enableGenericTieredCache(api, ctx, rc)
	if err != nil {
		return TieredCache{}, err
	}
	lastModified = result.LastModified

	result, err = enableSmartTieredCache(api, ctx, rc)
	if err != nil {
		return TieredCache{}, err
	}

	if result.LastModified.After(lastModified) {
		lastModified = result.LastModified
	}
	return TieredCache{Type: TieredCacheSmart, LastModified: lastModified}, nil
}

// DeleteTieredCache allows you to delete the tiered cache settings for a zone.
// This is equivalent to using SetTieredCache with the value of TieredCacheOff.
//
// API Reference: https://api.cloudflare.com/#smart-tiered-cache-delete-smart-tiered-cache-setting
// API Reference: https://api.cloudflare.com/#tiered-cache-patch-tiered-cache-setting
func (api *API) DeleteTieredCache(ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	var lastModified time.Time

	result, err := deleteSmartTieredCache(api, ctx, rc)
	if err != nil {
		return TieredCache{}, err
	}
	lastModified = result.LastModified

	result, err = disableGenericTieredCache(api, ctx, rc)
	if err != nil {
		return TieredCache{}, err
	}

	if result.LastModified.After(lastModified) {
		lastModified = result.LastModified
	}
	return TieredCache{Type: TieredCacheOff, LastModified: lastModified}, nil
}

type tieredCacheResult struct {
	ID           string    `json:"id"`
	Value        string    `json:"value,omitempty"`
	LastModified time.Time `json:"modified_on"`
}

type tieredCacheResponse struct {
	Result tieredCacheResult `json:"result"`
	Response
}

type tieredCacheSetting struct {
	Value string `json:"value"`
}

func getGenericTieredCache(api *API, ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	uri := fmt.Sprintf("/zones/%s/argo/tiered_caching", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	var response tieredCacheResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	if !response.Success {
		return TieredCache{Type: TieredCacheOff}, errors.New("request to retrieve generic tiered cache failed")
	}

	if response.Result.Value == "off" {
		return TieredCache{Type: TieredCacheOff, LastModified: response.Result.LastModified}, nil
	}

	return TieredCache{Type: TieredCacheGeneric, LastModified: response.Result.LastModified}, nil
}

func getSmartTieredCache(api *API, ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	uri := fmt.Sprintf("/zones/%s/cache/tiered_cache_smart_topology_enable", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		var notFoundError *NotFoundError
		if errors.As(err, &notFoundError) {
			return TieredCache{Type: TieredCacheOff}, nil
		}
		return TieredCache{Type: TieredCacheOff}, err
	}

	var response tieredCacheResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	if !response.Success {
		return TieredCache{Type: TieredCacheOff}, errors.New("request to retrieve smart tiered cache failed")
	}

	if response.Result.Value == "off" {
		return TieredCache{Type: TieredCacheOff, LastModified: response.Result.LastModified}, nil
	}
	return TieredCache{Type: TieredCacheSmart, LastModified: response.Result.LastModified}, nil
}

func enableGenericTieredCache(api *API, ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	uri := fmt.Sprintf("/zones/%s/argo/tiered_caching", rc.Identifier)
	setting := tieredCacheSetting{
		Value: "on",
	}
	body, err := json.Marshal(setting)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, body)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	var response tieredCacheResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	if !response.Success {
		return TieredCache{Type: TieredCacheOff}, errors.New("request to enable generic tiered cache failed")
	}

	return TieredCache{Type: TieredCacheGeneric, LastModified: response.Result.LastModified}, nil
}

func enableSmartTieredCache(api *API, ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	uri := fmt.Sprintf("/zones/%s/cache/tiered_cache_smart_topology_enable", rc.Identifier)
	setting := tieredCacheSetting{
		Value: "on",
	}
	body, err := json.Marshal(setting)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, body)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	var response tieredCacheResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	if !response.Success {
		return TieredCache{Type: TieredCacheOff}, errors.New("request to enable smart tiered cache failed")
	}

	return TieredCache{Type: TieredCacheSmart, LastModified: response.Result.LastModified}, nil
}

func disableGenericTieredCache(api *API, ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	uri := fmt.Sprintf("/zones/%s/argo/tiered_caching", rc.Identifier)
	setting := tieredCacheSetting{
		Value: "off",
	}
	body, err := json.Marshal(setting)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, body)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	var response tieredCacheResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	if !response.Success {
		return TieredCache{Type: TieredCacheOff}, errors.New("request to disable generic tiered cache failed")
	}

	return TieredCache{Type: TieredCacheOff, LastModified: response.Result.LastModified}, nil
}

func deleteSmartTieredCache(api *API, ctx context.Context, rc *ResourceContainer) (TieredCache, error) {
	uri := fmt.Sprintf("/zones/%s/cache/tiered_cache_smart_topology_enable", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		var notFoundError *NotFoundError
		if errors.As(err, &notFoundError) {
			return TieredCache{Type: TieredCacheOff}, nil
		}
		return TieredCache{Type: TieredCacheOff}, err
	}

	var response tieredCacheResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return TieredCache{Type: TieredCacheOff}, err
	}

	if !response.Success {
		return TieredCache{Type: TieredCacheOff}, errors.New("request to disable smart tiered cache failed")
	}

	return TieredCache{Type: TieredCacheOff, LastModified: response.Result.LastModified}, nil
}
