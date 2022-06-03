package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TeamsAccount struct {
	GatewayTag   string `json:"gateway_tag"`   // Internal teams ID
	ProviderName string `json:"provider_name"` // Auth provider
	ID           string `json:"id"`            // cloudflare account ID
}

// TeamsAccountResponse is the API response, containing information on teams
// account.
type TeamsAccountResponse struct {
	Response
	Result TeamsAccount `json:"result"`
}

// TeamsConfigResponse is the API response, containing information on teams
// account config.
type TeamsConfigResponse struct {
	Response
	Result TeamsConfiguration `json:"result"`
}

// TeamsConfiguration data model.
type TeamsConfiguration struct {
	Settings  TeamsAccountSettings `json:"settings"`
	CreatedAt time.Time            `json:"created_at,omitempty"`
	UpdatedAt time.Time            `json:"updated_at,omitempty"`
}

type TeamsAccountSettings struct {
	Antivirus        *TeamsAntivirus   `json:"antivirus,omitempty"`
	TLSDecrypt       *TeamsTLSDecrypt  `json:"tls_decrypt,omitempty"`
	ActivityLog      *TeamsActivityLog `json:"activity_log,omitempty"`
	BlockPage        *TeamsBlockPage   `json:"block_page,omitempty"`
	BrowserIsolation *BrowserIsolation `json:"browser_isolation,omitempty"`
	FIPS             *TeamsFIPS        `json:"fips,omitempty"`
}

type BrowserIsolation struct {
	UrlBrowserIsolationEnabled bool `json:"url_browser_isolation_enabled"`
}

type TeamsAntivirus struct {
	EnabledDownloadPhase bool `json:"enabled_download_phase"`
	EnabledUploadPhase   bool `json:"enabled_upload_phase"`
	FailClosed           bool `json:"fail_closed"`
}

type TeamsFIPS struct {
	TLS bool `json:"tls"`
}

type TeamsTLSDecrypt struct {
	Enabled bool `json:"enabled"`
}

type TeamsActivityLog struct {
	Enabled bool `json:"enabled"`
}

type TeamsBlockPage struct {
	Enabled         *bool  `json:"enabled,omitempty"`
	FooterText      string `json:"footer_text,omitempty"`
	HeaderText      string `json:"header_text,omitempty"`
	LogoPath        string `json:"logo_path,omitempty"`
	BackgroundColor string `json:"background_color,omitempty"`
	Name            string `json:"name,omitempty"`
	MailtoAddress   string `json:"mailto_address,omitempty"`
	MailtoSubject   string `json:"mailto_subject,omitempty"`
	SuppressFooter  *bool  `json:"suppress_footer,omitempty"`
}

type TeamsRuleType = string

const (
	TeamsHttpRuleType TeamsRuleType = "http"
	TeamsDnsRuleType  TeamsRuleType = "dns"
	TeamsL4RuleType   TeamsRuleType = "l4"
)

type TeamsAccountLoggingConfiguration struct {
	LogAll    bool `json:"log_all"`
	LogBlocks bool `json:"log_blocks"`
}

type TeamsLoggingSettings struct {
	LoggingSettingsByRuleType map[TeamsRuleType]TeamsAccountLoggingConfiguration `json:"settings_by_rule_type"`
	RedactPii                 bool                                               `json:"redact_pii,omitempty"`
}

type TeamsDeviceSettings struct {
	GatewayProxyEnabled    bool `json:"gateway_proxy_enabled"`
	GatewayProxyUDPEnabled bool `json:"gateway_udp_proxy_enabled"`
}

type TeamsDeviceSettingsResponse struct {
	Response
	Result TeamsDeviceSettings `json:"result"`
}

type TeamsLoggingSettingsResponse struct {
	Response
	Result TeamsLoggingSettings `json:"result"`
}

// TeamsAccount returns teams account information with internal and external ID.
//
// API reference: TBA.
func (api *API) TeamsAccount(ctx context.Context, accountID string) (TeamsAccount, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsAccount{}, err
	}

	var teamsAccountResponse TeamsAccountResponse
	err = json.Unmarshal(res, &teamsAccountResponse)
	if err != nil {
		return TeamsAccount{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsAccountResponse.Result, nil
}

// TeamsAccountConfiguration returns teams account configuration.
//
// API reference: TBA.
func (api *API) TeamsAccountConfiguration(ctx context.Context, accountID string) (TeamsConfiguration, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/configuration", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsConfiguration{}, err
	}

	var teamsConfigResponse TeamsConfigResponse
	err = json.Unmarshal(res, &teamsConfigResponse)
	if err != nil {
		return TeamsConfiguration{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsConfigResponse.Result, nil
}

// TeamsAccountDeviceConfiguration returns teams account device configuration with udp status.
//
// API reference: TBA.
func (api *API) TeamsAccountDeviceConfiguration(ctx context.Context, accountID string) (TeamsDeviceSettings, error) {
	uri := fmt.Sprintf("/accounts/%s/devices/settings", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsDeviceSettings{}, err
	}

	var teamsDeviceResponse TeamsDeviceSettingsResponse
	err = json.Unmarshal(res, &teamsDeviceResponse)
	if err != nil {
		return TeamsDeviceSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsDeviceResponse.Result, nil
}

// TeamsAccountLoggingConfiguration returns teams account logging configuration.
//
// API reference: TBA.
func (api *API) TeamsAccountLoggingConfiguration(ctx context.Context, accountID string) (TeamsLoggingSettings, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/logging", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsLoggingSettings{}, err
	}

	var teamsConfigResponse TeamsLoggingSettingsResponse
	err = json.Unmarshal(res, &teamsConfigResponse)
	if err != nil {
		return TeamsLoggingSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsConfigResponse.Result, nil
}

// TeamsAccountUpdateConfiguration updates a teams account configuration.
//
// API reference: TBA.
func (api *API) TeamsAccountUpdateConfiguration(ctx context.Context, accountID string, config TeamsConfiguration) (TeamsConfiguration, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/configuration", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, config)
	if err != nil {
		return TeamsConfiguration{}, err
	}

	var teamsConfigResponse TeamsConfigResponse
	err = json.Unmarshal(res, &teamsConfigResponse)
	if err != nil {
		return TeamsConfiguration{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsConfigResponse.Result, nil
}

// TeamsAccountUpdateLoggingConfiguration updates the log settings and returns new teams account logging configuration.
//
// API reference: TBA.
func (api *API) TeamsAccountUpdateLoggingConfiguration(ctx context.Context, accountID string, config TeamsLoggingSettings) (TeamsLoggingSettings, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/logging", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, config)
	if err != nil {
		return TeamsLoggingSettings{}, err
	}

	var teamsConfigResponse TeamsLoggingSettingsResponse
	err = json.Unmarshal(res, &teamsConfigResponse)
	if err != nil {
		return TeamsLoggingSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsConfigResponse.Result, nil
}

// TeamsAccountDeviceUpdateConfiguration updates teams account device configuration including udp filtering status.
//
// API reference: TBA.
func (api *API) TeamsAccountDeviceUpdateConfiguration(ctx context.Context, accountID string, settings TeamsDeviceSettings) (TeamsDeviceSettings, error) {
	uri := fmt.Sprintf("/accounts/%s/devices/settings", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, settings)
	if err != nil {
		return TeamsDeviceSettings{}, err
	}

	var teamsDeviceResponse TeamsDeviceSettingsResponse
	err = json.Unmarshal(res, &teamsDeviceResponse)
	if err != nil {
		return TeamsDeviceSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsDeviceResponse.Result, nil
}
