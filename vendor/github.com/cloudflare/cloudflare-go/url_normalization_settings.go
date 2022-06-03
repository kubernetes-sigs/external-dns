package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type URLNormalizationSettings struct {
	Type  string `json:"type"`
	Scope string `json:"scope"`
}

type URLNormalizationSettingsResponse struct {
	Result URLNormalizationSettings `json:"result"`
	Response
}

type URLNormalizationSettingsUpdateParams struct {
	Type  string `json:"type"`
	Scope string `json:"scope"`
}

// URLNormalizationSettings API reference: https://api.cloudflare.com/#url-normalization-get-url-normalization-settings
func (api *API) URLNormalizationSettings(ctx context.Context, rc *ResourceContainer) (URLNormalizationSettings, error) {
	uri := fmt.Sprintf("/zones/%s/url_normalization", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return URLNormalizationSettings{}, err
	}

	var urlNormalizationSettingsResponse URLNormalizationSettingsResponse
	err = json.Unmarshal(res, &urlNormalizationSettingsResponse)
	if err != nil {
		return URLNormalizationSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return urlNormalizationSettingsResponse.Result, nil
}

// UpdateURLNormalizationSettings https://api.cloudflare.com/#url-normalization-update-url-normalization-settings
func (api *API) UpdateURLNormalizationSettings(ctx context.Context, rc *ResourceContainer, params URLNormalizationSettingsUpdateParams) (URLNormalizationSettings, error) {
	uri := fmt.Sprintf("/zones/%s/url_normalization", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return URLNormalizationSettings{}, err
	}

	var urlNormalizationSettingsResponse URLNormalizationSettingsResponse
	err = json.Unmarshal(res, &urlNormalizationSettingsResponse)
	if err != nil {
		return URLNormalizationSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return urlNormalizationSettingsResponse.Result, nil
}
