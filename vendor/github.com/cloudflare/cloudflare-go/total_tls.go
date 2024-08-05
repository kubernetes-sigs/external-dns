package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type TotalTLS struct {
	Enabled              *bool  `json:"enabled,omitempty"`
	CertificateAuthority string `json:"certificate_authority,omitempty"`
	ValidityDays         int    `json:"validity_days,omitempty"`
}

type TotalTLSResponse struct {
	Response
	Result TotalTLS `json:"result"`
}

// GetTotalTLS Get Total TLS Settings for a Zone.
//
// API Reference: https://api.cloudflare.com/#total-tls-total-tls-settings-details
func (api *API) GetTotalTLS(ctx context.Context, rc *ResourceContainer) (TotalTLS, error) {
	if rc.Identifier == "" {
		return TotalTLS{}, ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/acm/total_tls", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TotalTLS{}, err
	}

	var r TotalTLSResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return TotalTLS{}, err
	}

	return r.Result, nil
}

// SetTotalTLS Set Total TLS Settings or disable the feature for a Zone.
//
// API Reference: https://api.cloudflare.com/#total-tls-enable-or-disable-total-tls
func (api *API) SetTotalTLS(ctx context.Context, rc *ResourceContainer, params TotalTLS) (TotalTLS, error) {
	if rc.Identifier == "" {
		return TotalTLS{}, ErrMissingZoneID
	}
	uri := fmt.Sprintf("/zones/%s/acm/total_tls", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return TotalTLS{}, err
	}

	var r TotalTLSResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return TotalTLS{}, err
	}

	return r.Result, nil
}
