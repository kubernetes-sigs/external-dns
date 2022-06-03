package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AuthIdCharacteristics a single option from
// configuration?properties=auth_id_characteristics.
type AuthIdCharacteristics struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

// APIShield is all the available options under
// configuration?properties=auth_id_characteristics.
type APIShield struct {
	AuthIdCharacteristics []AuthIdCharacteristics `json:"auth_id_characteristics"`
}

// APIShieldResponse represents the response from the api_gateway/configuration endpoint.
type APIShieldResponse struct {
	Result APIShield `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type UpdateAPIShieldParams struct {
	AuthIdCharacteristics []AuthIdCharacteristics `json:"auth_id_characteristics"`
}

// GetAPIShieldConfiguration gets a zone API shield configuration.
//
// API documentation: https://api.cloudflare.com/#api-shield-settings-retrieve-information-about-specific-configuration-properties
func (api *API) GetAPIShieldConfiguration(ctx context.Context, rc *ResourceContainer) (APIShield, ResultInfo, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/configuration?properties=auth_id_characteristics", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return APIShield{}, ResultInfo{}, err
	}
	var asResponse APIShieldResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return APIShield{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, asResponse.ResultInfo, nil
}

// UpdateAPIShieldConfiguration sets a zone API shield configuration.
//
// API documentation: https://api.cloudflare.com/#api-shield-settings-set-configuration-properties
func (api *API) UpdateAPIShieldConfiguration(ctx context.Context, rc *ResourceContainer, params UpdateAPIShieldParams) (Response, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/configuration", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return Response{}, err
	}
	var asResponse Response
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return Response{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse, nil
}
