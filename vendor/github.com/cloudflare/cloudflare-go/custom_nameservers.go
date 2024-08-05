package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type CustomNameserverRecord struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type CustomNameserver struct {
	NSName string `json:"ns_name"`
	NSSet  int    `json:"ns_set"`
}

type CustomNameserverResult struct {
	DNSRecords []CustomNameserverRecord `json:"dns_records"`
	NSName     string                   `json:"ns_name"`
	NSSet      int                      `json:"ns_set"`
	Status     string                   `json:"status"`
	ZoneTag    string                   `json:"zone_tag"`
}

type CustomNameserverZoneMetadata struct {
	NSSet   int  `json:"ns_set"`
	Enabled bool `json:"enabled"`
}

type customNameserverListResponse struct {
	Response
	Result []CustomNameserverResult `json:"result"`
}

type customNameserverCreateResponse struct {
	Response
	Result CustomNameserverResult `json:"result"`
}

type getEligibleZonesAccountCustomNameserversResponse struct {
	Result []string `json:"result"`
}

type customNameserverZoneMetadata struct {
	Response
	Result CustomNameserverZoneMetadata
}

type GetCustomNameserversParams struct{}

type CreateCustomNameserversParams struct {
	NSName string `json:"ns_name"`
	NSSet  int    `json:"ns_set"`
}

type DeleteCustomNameserversParams struct {
	NSName string
}

type GetEligibleZonesAccountCustomNameserversParams struct{}

type GetCustomNameserverZoneMetadataParams struct{}

type UpdateCustomNameserverZoneMetadataParams struct {
	NSSet   int  `json:"ns_set"`
	Enabled bool `json:"enabled"`
}

// GetCustomNameservers lists custom nameservers.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-list-account-custom-nameservers
func (api *API) GetCustomNameservers(ctx context.Context, rc *ResourceContainer, params GetCustomNameserversParams) ([]CustomNameserverResult, error) {
	if rc.Level != AccountRouteLevel {
		return []CustomNameserverResult{}, ErrRequiredAccountLevelResourceContainer
	}
	uri := fmt.Sprintf("/%s/%s/custom_ns", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var response customNameserverListResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// CreateCustomNameservers adds a custom nameserver.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-add-account-custom-nameserver
func (api *API) CreateCustomNameservers(ctx context.Context, rc *ResourceContainer, params CreateCustomNameserversParams) (CustomNameserverResult, error) {
	if rc.Level != AccountRouteLevel {
		return CustomNameserverResult{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/custom_ns", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return CustomNameserverResult{}, err
	}

	response := &customNameserverCreateResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CustomNameserverResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// DeleteCustomNameservers removes a custom nameserver.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-delete-account-custom-nameserver
func (api *API) DeleteCustomNameservers(ctx context.Context, rc *ResourceContainer, params DeleteCustomNameserversParams) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	if params.NSName == "" {
		return errors.New("missing required NSName parameter")
	}

	uri := fmt.Sprintf("/%s/%s/custom_ns/%s", rc.Level, rc.Identifier, params.NSName)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetEligibleZonesAccountCustomNameservers lists zones eligible for custom nameservers.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-get-eligible-zones-for-account-custom-nameservers
func (api *API) GetEligibleZonesAccountCustomNameservers(ctx context.Context, rc *ResourceContainer, params GetEligibleZonesAccountCustomNameserversParams) ([]string, error) {
	if rc.Level != AccountRouteLevel {
		return []string{}, ErrRequiredAccountLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/custom_ns/availability", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var response getEligibleZonesAccountCustomNameserversResponse
	err = json.Unmarshal(res, &response)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// GetCustomNameserverZoneMetadata get metadata for custom nameservers on a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-usage-for-a-zone-get-account-custom-nameserver-related-zone-metadata
func (api *API) GetCustomNameserverZoneMetadata(ctx context.Context, rc *ResourceContainer, params GetCustomNameserverZoneMetadataParams) (CustomNameserverZoneMetadata, error) {
	if rc.Level != ZoneRouteLevel {
		return CustomNameserverZoneMetadata{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/custom_ns", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return CustomNameserverZoneMetadata{}, err
	}

	var response customNameserverZoneMetadata
	err = json.Unmarshal(res, &response)
	if err != nil {
		return CustomNameserverZoneMetadata{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// UpdateCustomNameserverZoneMetadata set metadata for custom nameservers on a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/account-level-custom-nameservers-usage-for-a-zone-set-account-custom-nameserver-related-zone-metadata
func (api *API) UpdateCustomNameserverZoneMetadata(ctx context.Context, rc *ResourceContainer, params UpdateCustomNameserverZoneMetadataParams) error {
	if rc.Level != ZoneRouteLevel {
		return ErrRequiredZoneLevelResourceContainer
	}

	uri := fmt.Sprintf("/%s/%s/custom_ns", rc.Level, rc.Identifier)

	_, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return err
	}

	return nil
}
