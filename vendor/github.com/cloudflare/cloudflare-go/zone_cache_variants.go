package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ZoneCacheVariantsValues struct {
	Avif []string `json:"avif,omitempty"`
	Bmp  []string `json:"bmp,omitempty"`
	Gif  []string `json:"gif,omitempty"`
	Jpeg []string `json:"jpeg,omitempty"`
	Jpg  []string `json:"jpg,omitempty"`
	Jpg2 []string `json:"jpg2,omitempty"`
	Jp2  []string `json:"jp2,omitempty"`
	Png  []string `json:"png,omitempty"`
	Tiff []string `json:"tiff,omitempty"`
	Tif  []string `json:"tif,omitempty"`
	Webp []string `json:"webp,omitempty"`
}

type ZoneCacheVariants struct {
	ModifiedOn time.Time               `json:"modified_on"`
	Value      ZoneCacheVariantsValues `json:"value"`
}

type updateZoneCacheVariantsRequest struct {
	Value ZoneCacheVariantsValues `json:"value"`
}

type zoneCacheVariantsSingleResponse struct {
	Response
	Result ZoneCacheVariants `json:"result"`
}

// ZoneCacheVariants returns information about the current cache variants
//
// API reference: https://api.cloudflare.com/#zone-cache-settings-get-variants-setting
func (api *API) ZoneCacheVariants(ctx context.Context, zoneID string) (ZoneCacheVariants, error) {
	uri := fmt.Sprintf("/zones/%s/cache/variants", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneCacheVariants{}, err
	}
	var r zoneCacheVariantsSingleResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return ZoneCacheVariants{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateZoneCacheVariants updates the cache variants for a given zone.
//
// API reference: https://api.cloudflare.com/#zone-cache-settings-change-variants-setting
func (api *API) UpdateZoneCacheVariants(ctx context.Context, zoneID string, variants ZoneCacheVariantsValues) (ZoneCacheVariants, error) {
	uri := fmt.Sprintf("/zones/%s/cache/variants", zoneID)

	updateReq := updateZoneCacheVariantsRequest{Value: variants}
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, updateReq)
	if err != nil {
		return ZoneCacheVariants{}, err
	}

	response := &zoneCacheVariantsSingleResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneCacheVariants{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// DeleteZoneCacheVariants deletes cache variants for a given zone.
//
// API reference: https://api.cloudflare.com/#zone-cache-settings-delete-variants-setting
func (api *API) DeleteZoneCacheVariants(ctx context.Context, zoneID string) error {
	uri := fmt.Sprintf("/zones/%s/cache/variants", zoneID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
