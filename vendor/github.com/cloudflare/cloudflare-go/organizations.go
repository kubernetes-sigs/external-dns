package cloudflare

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Organization represents a multi-user organization.
type Organization struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Status      string   `json:"status,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Roles       []string `json:"roles,omitempty"`
}

// organizationResponse represents the response from the Organization endpoint.
type organizationResponse struct {
	Response
	Result     []Organization `json:"result"`
	ResultInfo `json:"result_info"`
}

// ListOrganizations lists organizations of the logged-in user.
// API reference:
// 	https://api.cloudflare.com/#user-s-organizations-list-organizations
//	GET /user/organizations
func (api *API) ListOrganizations() ([]Organization, ResultInfo, error) {
	var r organizationResponse
	res, err := api.makeRequest("GET", "/user/organizations", nil)
	if err != nil {
		return []Organization{}, ResultInfo{}, errors.Wrap(err, errMakeRequestError)
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return []Organization{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}

	return r.Result, r.ResultInfo, nil
}
