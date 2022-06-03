package cloudflare

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

// ErrMissingDomain is for when domain is needed but not given.
var ErrMissingDomain = errors.New("required domain missing")

// DomainDetails represents details for a domain.
type DomainDetails struct {
	Domain                string                `json:"domain"`
	ResolvesToRefs        []ResolvesToRefs      `json:"resolves_to_refs"`
	PopularityRank        int                   `json:"popularity_rank"`
	Application           Application           `json:"application"`
	RiskTypes             []interface{}         `json:"risk_types"`
	ContentCategories     []ContentCategories   `json:"content_categories"`
	AdditionalInformation AdditionalInformation `json:"additional_information"`
}

// ResolvesToRefs what a domain resolves to.
type ResolvesToRefs struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

type Application struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ContentCategories represents the categories for a domain.
type ContentCategories struct {
	ID              int    `json:"id"`
	SuperCategoryID int    `json:"super_category_id"`
	Name            string `json:"name"`
}

// AdditionalInformation represents any additional information for a domain.
type AdditionalInformation struct {
	SuspectedMalwareFamily string `json:"suspected_malware_family"`
}

// DomainHistory represents the history for a domain.
type DomainHistory struct {
	Domain          string            `json:"domain"`
	Categorizations []Categorizations `json:"categorizations"`
}

// Categories represents categories for a domain.
type Categories struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Categorizations represents the categories and when those categories were set.
type Categorizations struct {
	Categories []Categories `json:"categories"`
	Start      string       `json:"start"`
	End        string       `json:"end"`
}

// GetDomainDetailsParameters represent the parameters for a domain details request.
type GetDomainDetailsParameters struct {
	AccountID string `url:"-"`
	Domain    string `url:"domain,omitempty"`
}

// DomainDetailsResponse represents an API response for domain details.
type DomainDetailsResponse struct {
	Response
	Result DomainDetails `json:"result,omitempty"`
}

// GetBulkDomainDetailsParameters represents the parameters for bulk domain details request.
type GetBulkDomainDetailsParameters struct {
	AccountID string   `url:"-"`
	Domains   []string `url:"domain"`
}

// GetBulkDomainDetailsResponse represents an API response for bulk domain details.
type GetBulkDomainDetailsResponse struct {
	Response
	Result []DomainDetails `json:"result,omitempty"`
}

// GetDomainHistoryParameters represents the parameters for domain history request.
type GetDomainHistoryParameters struct {
	AccountID string `url:"-"`
	Domain    string `url:"domain,omitempty"`
}

// GetDomainHistoryResponse represents an API response for domain history.
type GetDomainHistoryResponse struct {
	Response
	Result []DomainHistory `json:"result,omitempty"`
}

// IntelligenceDomainDetails gets domain information.
//
// API Reference: https://api.cloudflare.com/#domain-intelligence-get-domain-details
func (api *API) IntelligenceDomainDetails(ctx context.Context, params GetDomainDetailsParameters) (DomainDetails, error) {
	if params.AccountID == "" {
		return DomainDetails{}, ErrMissingAccountID
	}

	if params.Domain == "" {
		return DomainDetails{}, ErrMissingDomain
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/domain", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return DomainDetails{}, err
	}

	var domainDetails DomainDetailsResponse
	if err := json.Unmarshal(res, &domainDetails); err != nil {
		return DomainDetails{}, err
	}
	return domainDetails.Result, nil
}

// IntelligenceBulkDomainDetails gets domain information for a list of domains.
//
// API Reference: https://api.cloudflare.com/#domain-intelligence-get-multiple-domain-details
func (api *API) IntelligenceBulkDomainDetails(ctx context.Context, params GetBulkDomainDetailsParameters) ([]DomainDetails, error) {
	if params.AccountID == "" {
		return []DomainDetails{}, ErrMissingAccountID
	}

	if len(params.Domains) == 0 {
		return []DomainDetails{}, ErrMissingDomain
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/domain/bulk", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DomainDetails{}, err
	}

	var domainDetails GetBulkDomainDetailsResponse
	if err := json.Unmarshal(res, &domainDetails); err != nil {
		return []DomainDetails{}, err
	}
	return domainDetails.Result, nil
}

// IntelligenceDomainHistory get domain history for given domain
//
// API Reference: https://api.cloudflare.com/#domain-history-get-domain-history
func (api *API) IntelligenceDomainHistory(ctx context.Context, params GetDomainHistoryParameters) ([]DomainHistory, error) {
	if params.AccountID == "" {
		return []DomainHistory{}, ErrMissingAccountID
	}

	if params.Domain == "" {
		return []DomainHistory{}, ErrMissingDomain
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/intel/domain-history", params.AccountID), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []DomainHistory{}, err
	}

	var domainDetails GetDomainHistoryResponse
	if err := json.Unmarshal(res, &domainDetails); err != nil {
		return []DomainHistory{}, err
	}
	return domainDetails.Result, nil
}
