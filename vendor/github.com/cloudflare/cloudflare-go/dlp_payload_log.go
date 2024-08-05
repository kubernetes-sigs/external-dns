package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type DLPPayloadLogSettings struct {
	PublicKey string `json:"public_key,omitempty"`

	// Only present in responses
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type GetDLPPayloadLogSettingsParams struct{}

type DLPPayloadLogSettingsResponse struct {
	Response
	Result DLPPayloadLogSettings `json:"result"`
}

// GetDLPPayloadLogSettings gets the current DLP payload logging settings.
//
// API reference: https://api.cloudflare.com/#dlp-payload-log-settings-get-settings
func (api *API) GetDLPPayloadLogSettings(ctx context.Context, rc *ResourceContainer, params GetDLPPayloadLogSettingsParams) (DLPPayloadLogSettings, error) {
	if rc.Identifier == "" {
		return DLPPayloadLogSettings{}, ErrMissingResourceIdentifier
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/payload_log", rc.Level, rc.Identifier), nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DLPPayloadLogSettings{}, err
	}

	var dlpPayloadLogSettingsResponse DLPPayloadLogSettingsResponse
	err = json.Unmarshal(res, &dlpPayloadLogSettingsResponse)
	if err != nil {
		return DLPPayloadLogSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpPayloadLogSettingsResponse.Result, nil
}

// UpdateDLPPayloadLogSettings sets the current DLP payload logging settings to new values.
//
// API reference: https://api.cloudflare.com/#dlp-payload-log-settings-update-settings
func (api *API) UpdateDLPPayloadLogSettings(ctx context.Context, rc *ResourceContainer, settings DLPPayloadLogSettings) (DLPPayloadLogSettings, error) {
	if rc.Identifier == "" {
		return DLPPayloadLogSettings{}, ErrMissingResourceIdentifier
	}

	uri := buildURI(fmt.Sprintf("/%s/%s/dlp/payload_log", rc.Level, rc.Identifier), nil)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, settings)
	if err != nil {
		return DLPPayloadLogSettings{}, err
	}

	var dlpPayloadLogSettingsResponse DLPPayloadLogSettingsResponse
	err = json.Unmarshal(res, &dlpPayloadLogSettingsResponse)
	if err != nil {
		return DLPPayloadLogSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dlpPayloadLogSettingsResponse.Result, nil
}
