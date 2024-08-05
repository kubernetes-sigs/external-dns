package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// PageShield represents the page shield object minus any timestamps.
type PageShield struct {
	Enabled                        *bool `json:"enabled,omitempty"`
	UseCloudflareReportingEndpoint *bool `json:"use_cloudflare_reporting_endpoint,omitempty"`
	UseConnectionURLPath           *bool `json:"use_connection_url_path,omitempty"`
}

type UpdatePageShieldSettingsParams struct {
	Enabled                        *bool `json:"enabled,omitempty"`
	UseCloudflareReportingEndpoint *bool `json:"use_cloudflare_reporting_endpoint,omitempty"`
	UseConnectionURLPath           *bool `json:"use_connection_url_path,omitempty"`
}

// PageShieldSettings represents the page shield settings for a zone.
type PageShieldSettings struct {
	PageShield
	UpdatedAt string `json:"updated_at"`
}

// PageShieldSettingsResponse represents the response from the page shield settings endpoint.
type PageShieldSettingsResponse struct {
	PageShield PageShieldSettings `json:"result"`
	Response
}

type GetPageShieldSettingsParams struct{}

// GetPageShieldSettings returns the page shield settings for a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-get-page-shield-settings
func (api *API) GetPageShieldSettings(ctx context.Context, rc *ResourceContainer, params GetPageShieldSettingsParams) (*PageShieldSettingsResponse, error) {
	uri := fmt.Sprintf("/zones/%s/page_shield", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldSettingsResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}

// UpdatePageShieldSettings updates the page shield settings for a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-update-page-shield-settings
func (api *API) UpdatePageShieldSettings(ctx context.Context, rc *ResourceContainer, params UpdatePageShieldSettingsParams) (*PageShieldSettingsResponse, error) {
	uri := fmt.Sprintf("/zones/%s/page_shield", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldSettingsResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}
