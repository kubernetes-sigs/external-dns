package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EmailRoutingSettings struct {
	Tag        string     `json:"tag,omitempty"`
	Name       string     `json:"name,omitempty"`
	Enabled    bool       `json:"enabled,omitempty"`
	Created    *time.Time `json:"created,omitempty"`
	Modified   *time.Time `json:"modified,omitempty"`
	SkipWizard *bool      `json:"skip_wizard,omitempty"`
	Status     string     `json:"status,omitempty"`
}

type EmailRoutingSettingsResponse struct {
	Result EmailRoutingSettings `json:"result,omitempty"`
	Response
}

type EmailRoutingDNSSettingsResponse struct {
	Result []DNSRecord `json:"result,omitempty"`
	Response
}

// GetEmailRoutingSettings Get information about the settings for your Email Routing zone.
//
// API reference: https://api.cloudflare.com/#email-routing-settings-get-email-routing-settings
func (api *API) GetEmailRoutingSettings(ctx context.Context, rc *ResourceContainer) (EmailRoutingSettings, error) {
	if rc.Identifier == "" {
		return EmailRoutingSettings{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/email/routing", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return EmailRoutingSettings{}, err
	}

	var r EmailRoutingSettingsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// EnableEmailRouting Enable you Email Routing zone. Add and lock the necessary MX and SPF records.
//
// API reference: https://api.cloudflare.com/#email-routing-settings-enable-email-routing
func (api *API) EnableEmailRouting(ctx context.Context, rc *ResourceContainer) (EmailRoutingSettings, error) {
	if rc.Identifier == "" {
		return EmailRoutingSettings{}, ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/email/routing/enable", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return EmailRoutingSettings{}, err
	}

	var r EmailRoutingSettingsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DisableEmailRouting Disable your Email Routing zone. Also removes additional MX records previously required for Email Routing to work.
//
// API reference: https://api.cloudflare.com/#email-routing-settings-disable-email-routing
func (api *API) DisableEmailRouting(ctx context.Context, rc *ResourceContainer) (EmailRoutingSettings, error) {
	if rc.Identifier == "" {
		return EmailRoutingSettings{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/email/routing/disable", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return EmailRoutingSettings{}, err
	}

	var r EmailRoutingSettingsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return EmailRoutingSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// GetEmailRoutingDNSSettings Show the DNS records needed to configure your Email Routing zone.
//
// API reference: https://api.cloudflare.com/#email-routing-settings-email-routing---dns-settings
func (api *API) GetEmailRoutingDNSSettings(ctx context.Context, rc *ResourceContainer) ([]DNSRecord, error) {
	if rc.Identifier == "" {
		return []DNSRecord{}, ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/email/routing/dns", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DNSRecord{}, err
	}

	var r EmailRoutingDNSSettingsResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return []DNSRecord{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
