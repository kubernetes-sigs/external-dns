package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// PageShieldPolicy represents a page shield policy.
type PageShieldPolicy struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	Enabled     *bool  `json:"enabled,omitempty"`
	Expression  string `json:"expression"`
	ID          string `json:"id"`
	Value       string `json:"value"`
}

type CreatePageShieldPolicyParams struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	Enabled     *bool  `json:"enabled,omitempty"`
	Expression  string `json:"expression"`
	ID          string `json:"id"`
	Value       string `json:"value"`
}

type UpdatePageShieldPolicyParams struct {
	Action      string `json:"action"`
	Description string `json:"description"`
	Enabled     *bool  `json:"enabled,omitempty"`
	Expression  string `json:"expression"`
	ID          string `json:"id"`
	Value       string `json:"value"`
}

// ListPageShieldPoliciesResponse represents the response from the list page shield policies endpoint.
type ListPageShieldPoliciesResponse struct {
	Result []PageShieldPolicy `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

type ListPageShieldPoliciesParams struct{}

// ListPageShieldPolicies lists all page shield policies for a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-list-page-shield-policies
func (api *API) ListPageShieldPolicies(ctx context.Context, rc *ResourceContainer, params ListPageShieldPoliciesParams) ([]PageShieldPolicy, ResultInfo, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/policies", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var psResponse ListPageShieldPoliciesResponse
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return psResponse.Result, psResponse.ResultInfo, nil
}

// CreatePageShieldPolicy creates a page shield policy for a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-create-page-shield-policy
func (api *API) CreatePageShieldPolicy(ctx context.Context, rc *ResourceContainer, params CreatePageShieldPolicyParams) (*PageShieldPolicy, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/policies", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, path, params)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldPolicy
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}

// DeletePageShieldPolicy deletes a page shield policy for a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-delete-page-shield-policy
func (api *API) DeletePageShieldPolicy(ctx context.Context, rc *ResourceContainer, policyID string) error {
	path := fmt.Sprintf("/zones/%s/page_shield/policies/%s", rc.Identifier, policyID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetPageShieldPolicy gets a page shield policy for a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-get-page-shield-policy
func (api *API) GetPageShieldPolicy(ctx context.Context, rc *ResourceContainer, policyID string) (*PageShieldPolicy, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/policies/%s", rc.Identifier, policyID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldPolicy
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}

// UpdatePageShieldPolicy updates a page shield policy for a zone.
//
// API documentation: https://developers.cloudflare.com/api/operations/page-shield-update-page-shield-policy
func (api *API) UpdatePageShieldPolicy(ctx context.Context, rc *ResourceContainer, params UpdatePageShieldPolicyParams) (*PageShieldPolicy, error) {
	path := fmt.Sprintf("/zones/%s/page_shield/policies/%s", rc.Identifier, params.ID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, path, params)
	if err != nil {
		return nil, err
	}

	var psResponse PageShieldPolicy
	err = json.Unmarshal(res, &psResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &psResponse, nil
}
