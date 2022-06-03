package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ListManagedHeadersResponse struct {
	Response
	Result ManagedHeaders `json:"result"`
}

type UpdateManagedHeadersParams struct {
	ManagedHeaders
}

type ManagedHeaders struct {
	ManagedRequestHeaders  []ManagedHeader `json:"managed_request_headers"`
	ManagedResponseHeaders []ManagedHeader `json:"managed_response_headers"`
}

type ManagedHeader struct {
	ID            string   `json:"id"`
	Enabled       bool     `json:"enabled"`
	HasCoflict    bool     `json:"has_conflict,omitempty"`
	ConflictsWith []string `json:"conflicts_with,omitempty"`
}

type ListManagedHeadersParams struct {
	Status string `url:"status,omitempty"`
}

func (api *API) ListZoneManagedHeaders(ctx context.Context, rc *ResourceContainer, params ListManagedHeadersParams) (ManagedHeaders, error) {
	if rc.Identifier == "" {
		return ManagedHeaders{}, ErrMissingZoneID
	}

	uri := buildURI(fmt.Sprintf("/zones/%s/managed_headers", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ManagedHeaders{}, err
	}

	result := ListManagedHeadersResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ManagedHeaders{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}

func (api *API) UpdateZoneManagedHeaders(ctx context.Context, rc *ResourceContainer, params UpdateManagedHeadersParams) (ManagedHeaders, error) {
	if rc.Identifier == "" {
		return ManagedHeaders{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/managed_headers", rc.Identifier)

	payload, err := json.Marshal(params.ManagedHeaders)
	if err != nil {
		return ManagedHeaders{}, err
	}

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, payload)
	if err != nil {
		return ManagedHeaders{}, err
	}

	result := ListManagedHeadersResponse{}
	if err := json.Unmarshal(res, &result); err != nil {
		return ManagedHeaders{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return result.Result, nil
}
