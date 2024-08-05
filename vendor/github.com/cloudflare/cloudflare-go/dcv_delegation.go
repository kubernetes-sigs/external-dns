package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type DCVDelegation struct {
	UUID string `json:"uuid"`
}

// DCVDelegationResponse represents the response from the dcv_delegation/uuid endpoint.
type DCVDelegationResponse struct {
	Result DCVDelegation `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type GetDCVDelegationParams struct{}

// GetDCVDelegation gets a zone DCV Delegation UUID.
//
// API documentation: https://developers.cloudflare.com/api/operations/dcv-delegation-uuid-get
func (api *API) GetDCVDelegation(ctx context.Context, rc *ResourceContainer, params GetDCVDelegationParams) (DCVDelegation, ResultInfo, error) {
	uri := fmt.Sprintf("/zones/%s/dcv_delegation/uuid", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DCVDelegation{}, ResultInfo{}, err
	}
	var dcvResponse DCVDelegationResponse
	err = json.Unmarshal(res, &dcvResponse)
	if err != nil {
		return DCVDelegation{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return dcvResponse.Result, dcvResponse.ResultInfo, nil
}
