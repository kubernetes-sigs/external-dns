package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// Retrieve whether the zone is subject to a zone hold, and metadata about the
// hold.
type ZoneHold struct {
	Hold              *bool      `json:"hold,omitempty"`
	IncludeSubdomains *bool      `json:"include_subdomains,omitempty"`
	HoldAfter         *time.Time `json:"hold_after,omitempty"`
}

// ZoneHoldResponse represents a response from the Zone Hold endpoint.
type ZoneHoldResponse struct {
	Result ZoneHold `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// CreateZoneHoldParams represents params for the Create Zone Hold
// endpoint.
type CreateZoneHoldParams struct {
	IncludeSubdomains *bool `url:"include_subdomains,omitempty"`
}

// DeleteZoneHoldParams represents params for the Delete Zone Hold
// endpoint.
type DeleteZoneHoldParams struct {
	HoldAfter *time.Time `url:"hold_after,omitempty"`
}

type GetZoneHoldParams struct{}

// CreateZoneHold enforces a zone hold on the zone, blocking the creation and
// activation of zone.
//
// API reference: https://developers.cloudflare.com/api/operations/zones-0-hold-post
func (api *API) CreateZoneHold(ctx context.Context, rc *ResourceContainer, params CreateZoneHoldParams) (ZoneHold, error) {
	if rc.Level != ZoneRouteLevel {
		return ZoneHold{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/hold", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return ZoneHold{}, err
	}

	response := &ZoneHoldResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneHold{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// DeleteZoneHold removes enforcement of a zone hold on the zone, permanently or
// temporarily, allowing the creation and activation of zones with this hostname.
//
// API reference:https://developers.cloudflare.com/api/operations/zones-0-hold-delete
func (api *API) DeleteZoneHold(ctx context.Context, rc *ResourceContainer, params DeleteZoneHoldParams) (ZoneHold, error) {
	if rc.Level != ZoneRouteLevel {
		return ZoneHold{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/hold", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return ZoneHold{}, err
	}

	response := &ZoneHoldResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneHold{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}

// GetZoneHold retrieves whether the zone is subject to a zone hold, and the
// metadata about the hold.
//
// API reference: https://developers.cloudflare.com/api/operations/zones-0-hold-get
func (api *API) GetZoneHold(ctx context.Context, rc *ResourceContainer, params GetZoneHoldParams) (ZoneHold, error) {
	if rc.Level != ZoneRouteLevel {
		return ZoneHold{}, ErrRequiredZoneLevelResourceContainer
	}

	uri := fmt.Sprintf("/zones/%s/hold", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneHold{}, err
	}

	response := &ZoneHoldResponse{}
	err = json.Unmarshal(res, &response)
	if err != nil {
		return ZoneHold{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return response.Result, nil
}
