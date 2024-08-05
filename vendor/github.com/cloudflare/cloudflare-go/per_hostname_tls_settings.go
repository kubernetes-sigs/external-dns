package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// HostnameTLSSetting represents the metadata for a user-created tls setting.
type HostnameTLSSetting struct {
	Hostname  string     `json:"hostname"`
	Value     string     `json:"value"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// HostnameTLSSettingResponse represents the response from the PUT and DELETE endpoints for per-hostname tls settings.
type HostnameTLSSettingResponse struct {
	Response
	Result HostnameTLSSetting `json:"result"`
}

// HostnameTLSSettingsResponse represents the response from the retrieval endpoint for per-hostname tls settings.
type HostnameTLSSettingsResponse struct {
	Response
	Result     []HostnameTLSSetting `json:"result"`
	ResultInfo `json:"result_info"`
}

// ListHostnameTLSSettingsParams represents the data related to per-hostname tls settings being retrieved.
type ListHostnameTLSSettingsParams struct {
	Setting           string `json:"-" url:"setting,omitempty"`
	PaginationOptions `json:"-"`
	Limit             int      `json:"-" url:"limit,omitempty"`
	Offset            int      `json:"-" url:"offset,omitempty"`
	Hostname          []string `json:"-" url:"hostname,omitempty"`
}

// UpdateHostnameTLSSettingParams represents the data related to the per-hostname tls setting being updated.
type UpdateHostnameTLSSettingParams struct {
	Setting  string
	Hostname string
	Value    string `json:"value"`
}

// DeleteHostnameTLSSettingParams represents the data related to the per-hostname tls setting being deleted.
type DeleteHostnameTLSSettingParams struct {
	Setting  string
	Hostname string
}

var (
	ErrMissingHostnameTLSSettingName = errors.New("tls setting name required but missing")
)

// ListHostnameTLSSettings returns a list of all user-created tls setting values for the specified setting and hostnames.
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-list
func (api *API) ListHostnameTLSSettings(ctx context.Context, rc *ResourceContainer, params ListHostnameTLSSettingsParams) ([]HostnameTLSSetting, ResultInfo, error) {
	if rc.Identifier == "" {
		return []HostnameTLSSetting{}, ResultInfo{}, ErrMissingZoneID
	}
	if params.Setting == "" {
		return []HostnameTLSSetting{}, ResultInfo{}, ErrMissingHostnameTLSSettingName
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/hostnames/settings/%s", rc.Identifier, params.Setting), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []HostnameTLSSetting{}, ResultInfo{}, err
	}
	var r HostnameTLSSettingsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []HostnameTLSSetting{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, r.ResultInfo, err
}

// UpdateHostnameTLSSetting will update the per-hostname tls setting for the specified hostname.
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-put
func (api *API) UpdateHostnameTLSSetting(ctx context.Context, rc *ResourceContainer, params UpdateHostnameTLSSettingParams) (HostnameTLSSetting, error) {
	if rc.Identifier == "" {
		return HostnameTLSSetting{}, ErrMissingZoneID
	}
	if params.Setting == "" {
		return HostnameTLSSetting{}, ErrMissingHostnameTLSSettingName
	}
	if params.Hostname == "" {
		return HostnameTLSSetting{}, ErrMissingHostname
	}

	uri := fmt.Sprintf("/zones/%s/hostnames/settings/%s/%s", rc.Identifier, params.Setting, params.Hostname)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return HostnameTLSSetting{}, err
	}
	var r HostnameTLSSettingResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameTLSSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteHostnameTLSSetting will delete the specified per-hostname tls setting.
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-delete
func (api *API) DeleteHostnameTLSSetting(ctx context.Context, rc *ResourceContainer, params DeleteHostnameTLSSettingParams) (HostnameTLSSetting, error) {
	if rc.Identifier == "" {
		return HostnameTLSSetting{}, ErrMissingZoneID
	}
	if params.Setting == "" {
		return HostnameTLSSetting{}, ErrMissingHostnameTLSSettingName
	}
	if params.Hostname == "" {
		return HostnameTLSSetting{}, ErrMissingHostname
	}

	uri := fmt.Sprintf("/zones/%s/hostnames/settings/%s/%s", rc.Identifier, params.Setting, params.Hostname)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return HostnameTLSSetting{}, err
	}
	var r HostnameTLSSettingResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameTLSSetting{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// HostnameTLSSettingCiphers represents the metadata for a user-created ciphers tls setting.
type HostnameTLSSettingCiphers struct {
	Hostname  string     `json:"hostname"`
	Value     []string   `json:"value"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// HostnameTLSSettingCiphersResponse represents the response from the PUT and DELETE endpoints for per-hostname ciphers tls settings.
type HostnameTLSSettingCiphersResponse struct {
	Response
	Result HostnameTLSSettingCiphers `json:"result"`
}

// HostnameTLSSettingsCiphersResponse represents the response from the retrieval endpoint for per-hostname ciphers tls settings.
type HostnameTLSSettingsCiphersResponse struct {
	Response
	Result     []HostnameTLSSettingCiphers `json:"result"`
	ResultInfo `json:"result_info"`
}

// ListHostnameTLSSettingsCiphersParams represents the data related to per-hostname ciphers tls settings being retrieved.
type ListHostnameTLSSettingsCiphersParams struct {
	PaginationOptions
	Limit    int      `json:"-" url:"limit,omitempty"`
	Offset   int      `json:"-" url:"offset,omitempty"`
	Hostname []string `json:"-" url:"hostname,omitempty"`
}

// UpdateHostnameTLSSettingCiphersParams represents the data related to the per-hostname ciphers tls setting being updated.
type UpdateHostnameTLSSettingCiphersParams struct {
	Hostname string
	Value    []string `json:"value"`
}

// DeleteHostnameTLSSettingCiphersParams represents the data related to the per-hostname ciphers tls setting being deleted.
type DeleteHostnameTLSSettingCiphersParams struct {
	Hostname string
}

// ListHostnameTLSSettingsCiphers returns a list of all user-created tls setting ciphers values for the specified setting and hostnames.
// Ciphers functions are separate due to the API returning a list of strings as the value, rather than a string (as is the case for the other tls settings).
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-list
func (api *API) ListHostnameTLSSettingsCiphers(ctx context.Context, rc *ResourceContainer, params ListHostnameTLSSettingsCiphersParams) ([]HostnameTLSSettingCiphers, ResultInfo, error) {
	if rc.Identifier == "" {
		return []HostnameTLSSettingCiphers{}, ResultInfo{}, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/hostnames/settings/ciphers", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []HostnameTLSSettingCiphers{}, ResultInfo{}, err
	}
	var r HostnameTLSSettingsCiphersResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []HostnameTLSSettingCiphers{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, r.ResultInfo, err
}

// UpdateHostnameTLSSettingCiphers will update the per-hostname ciphers tls setting for the specified hostname.
// Ciphers functions are separate due to the API returning a list of strings as the value, rather than a string (as is the case for the other tls settings).
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-put
func (api *API) UpdateHostnameTLSSettingCiphers(ctx context.Context, rc *ResourceContainer, params UpdateHostnameTLSSettingCiphersParams) (HostnameTLSSettingCiphers, error) {
	if rc.Identifier == "" {
		return HostnameTLSSettingCiphers{}, ErrMissingZoneID
	}
	if params.Hostname == "" {
		return HostnameTLSSettingCiphers{}, ErrMissingHostname
	}

	uri := fmt.Sprintf("/zones/%s/hostnames/settings/ciphers/%s", rc.Identifier, params.Hostname)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return HostnameTLSSettingCiphers{}, err
	}
	var r HostnameTLSSettingCiphersResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameTLSSettingCiphers{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteHostnameTLSSettingCiphers will delete the specified per-hostname ciphers tls setting value.
// Ciphers functions are separate due to the API returning a list of strings as the value, rather than a string (as is the case for the other tls settings).
//
// API reference: https://developers.cloudflare.com/api/operations/per-hostname-tls-settings-delete
func (api *API) DeleteHostnameTLSSettingCiphers(ctx context.Context, rc *ResourceContainer, params DeleteHostnameTLSSettingCiphersParams) (HostnameTLSSettingCiphers, error) {
	if rc.Identifier == "" {
		return HostnameTLSSettingCiphers{}, ErrMissingZoneID
	}
	if params.Hostname == "" {
		return HostnameTLSSettingCiphers{}, ErrMissingHostname
	}

	uri := fmt.Sprintf("/zones/%s/hostnames/settings/ciphers/%s", rc.Identifier, params.Hostname)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return HostnameTLSSettingCiphers{}, err
	}
	// Unmarshal into HostnameTLSSettingResponse first because the API returns an empty string
	var r HostnameTLSSettingResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return HostnameTLSSettingCiphers{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return HostnameTLSSettingCiphers{
		Hostname:  r.Result.Hostname,
		Value:     []string{},
		Status:    r.Result.Status,
		CreatedAt: r.Result.CreatedAt,
		UpdatedAt: r.Result.UpdatedAt,
	}, nil
}
