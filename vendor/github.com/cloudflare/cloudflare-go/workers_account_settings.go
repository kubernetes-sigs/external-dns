package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type WorkersAccountSettings struct {
	DefaultUsageModel string `json:"default_usage_model,omitempty"`
	GreenCompute      bool   `json:"green_compute,omitempty"`
}

type CreateWorkersAccountSettingsParameters struct {
	DefaultUsageModel string `json:"default_usage_model,omitempty"`
	GreenCompute      bool   `json:"green_compute,omitempty"`
}

type CreateWorkersAccountSettingsResponse struct {
	Response
	Result WorkersAccountSettings
}

type WorkersAccountSettingsParameters struct{}

type WorkersAccountSettingsResponse struct {
	Response
	Result WorkersAccountSettings
}

// CreateWorkersAccountSettings sets the account settings for Workers.
//
// API reference: https://api.cloudflare.com/#worker-account-settings-create-worker-account-settings
func (api *API) CreateWorkersAccountSettings(ctx context.Context, rc *ResourceContainer, params CreateWorkersAccountSettingsParameters) (WorkersAccountSettings, error) {
	if rc.Identifier == "" {
		return WorkersAccountSettings{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/account-settings", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return WorkersAccountSettings{}, err
	}

	var workersAccountSettingsResponse CreateWorkersAccountSettingsResponse
	if err := json.Unmarshal(res, &workersAccountSettingsResponse); err != nil {
		return WorkersAccountSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return workersAccountSettingsResponse.Result, nil
}

// WorkersAccountSettings returns the current account settings for Workers.
//
// API reference: https://api.cloudflare.com/#worker-account-settings-fetch-worker-account-settings
func (api *API) WorkersAccountSettings(ctx context.Context, rc *ResourceContainer, params WorkersAccountSettingsParameters) (WorkersAccountSettings, error) {
	if rc.Identifier == "" {
		return WorkersAccountSettings{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/account-settings", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return WorkersAccountSettings{}, err
	}

	var workersAccountSettingsResponse CreateWorkersAccountSettingsResponse
	if err := json.Unmarshal(res, &workersAccountSettingsResponse); err != nil {
		return WorkersAccountSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return workersAccountSettingsResponse.Result, nil
}
