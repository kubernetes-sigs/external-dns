package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// PagesDomain represents a pages domain.
type PagesDomain struct {
	ID               string           `json:"id"`
	Name             string           `json:"name"`
	Status           string           `json:"status"`
	VerificationData VerificationData `json:"verification_data"`
	ValidationData   ValidationData   `json:"validation_data"`
	ZoneTag          string           `json:"zone_tag"`
	CreatedOn        *time.Time       `json:"created_on"`
}

// VerificationData represents verification data for a domain.
type VerificationData struct {
	Status string `json:"status"`
}

// ValidationData represents validation data for a domain.
type ValidationData struct {
	Status string `json:"status"`
	Method string `json:"method"`
}

// PagesDomainsParameters represents parameters for a pages domains request.
type PagesDomainsParameters struct {
	AccountID   string
	ProjectName string
}

// PagesDomainsResponse represents an API response for a pages domains request.
type PagesDomainsResponse struct {
	Response
	Result []PagesDomain `json:"result,omitempty"`
}

// PagesDomainParameters represents parameters for a pages domain request.
type PagesDomainParameters struct {
	AccountID   string
	ProjectName string
	DomainName  string `json:"name,omitempty"`
}

// PagesDomainResponse represents an API response for a pages domain request.
type PagesDomainResponse struct {
	Response
	Result PagesDomain `json:"result,omitempty"`
}

// GetPagesDomains gets all domains for a pages project.
//
// API Reference: https://api.cloudflare.com/#pages-domains-get-domains
func (api *API) GetPagesDomains(ctx context.Context, params PagesDomainsParameters) ([]PagesDomain, error) {
	if params.AccountID == "" {
		return []PagesDomain{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return []PagesDomain{}, ErrMissingProjectName
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/domains", params.AccountID, params.ProjectName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []PagesDomain{}, err
	}

	var pageDomainResponse PagesDomainsResponse
	if err := json.Unmarshal(res, &pageDomainResponse); err != nil {
		return []PagesDomain{}, err
	}
	return pageDomainResponse.Result, nil
}

// GetPagesDomain gets a single domain.
//
// API Reference: https://api.cloudflare.com/#pages-domains-get-domains
func (api *API) GetPagesDomain(ctx context.Context, params PagesDomainParameters) (PagesDomain, error) {
	if params.AccountID == "" {
		return PagesDomain{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesDomain{}, ErrMissingProjectName
	}

	if params.DomainName == "" {
		return PagesDomain{}, ErrMissingDomain
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/domains/%s", params.AccountID, params.ProjectName, params.DomainName)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return PagesDomain{}, err
	}

	var pagesDomainResponse PagesDomainResponse
	if err := json.Unmarshal(res, &pagesDomainResponse); err != nil {
		return PagesDomain{}, err
	}
	return pagesDomainResponse.Result, nil
}

// PagesPatchDomain retries the validation status of a single domain.
//
// API Reference: https://api.cloudflare.com/#pages-domains-patch-domain
func (api *API) PagesPatchDomain(ctx context.Context, params PagesDomainParameters) (PagesDomain, error) {
	if params.AccountID == "" {
		return PagesDomain{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesDomain{}, ErrMissingProjectName
	}

	if params.DomainName == "" {
		return PagesDomain{}, ErrMissingDomain
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/domains/%s", params.AccountID, params.ProjectName, params.DomainName)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, nil)
	if err != nil {
		return PagesDomain{}, err
	}

	var pagesDomainResponse PagesDomainResponse
	if err := json.Unmarshal(res, &pagesDomainResponse); err != nil {
		return PagesDomain{}, err
	}
	return pagesDomainResponse.Result, nil
}

// PagesAddDomain adds a domain to a pages project.
//
// API Reference: https://api.cloudflare.com/#pages-domains-add-domain
func (api *API) PagesAddDomain(ctx context.Context, params PagesDomainParameters) (PagesDomain, error) {
	if params.AccountID == "" {
		return PagesDomain{}, ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return PagesDomain{}, ErrMissingProjectName
	}

	if params.DomainName == "" {
		return PagesDomain{}, ErrMissingDomain
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/domains", params.AccountID, params.ProjectName)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return PagesDomain{}, err
	}

	var pagesDomainResponse PagesDomainResponse
	if err := json.Unmarshal(res, &pagesDomainResponse); err != nil {
		return PagesDomain{}, err
	}
	return pagesDomainResponse.Result, nil
}

// PagesDeleteDomain removes a domain from a pages project.
//
// API Reference: https://api.cloudflare.com/#pages-domains-delete-domain
func (api *API) PagesDeleteDomain(ctx context.Context, params PagesDomainParameters) error {
	if params.AccountID == "" {
		return ErrMissingAccountID
	}

	if params.ProjectName == "" {
		return ErrMissingProjectName
	}

	if params.DomainName == "" {
		return ErrMissingDomain
	}

	uri := fmt.Sprintf("/accounts/%s/pages/projects/%s/domains/%s", params.AccountID, params.ProjectName, params.DomainName)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, params)
	if err != nil {
		return err
	}
	return nil
}
