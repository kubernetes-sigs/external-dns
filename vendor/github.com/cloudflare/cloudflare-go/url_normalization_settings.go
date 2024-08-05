package cloudflare

import (
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
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
