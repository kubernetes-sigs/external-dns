package cloudflare

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	// ErrMissingIdentifier is for when identifier is required but missing.
	ErrMissingIdentifier = errors.New("identifier required but missing")
	// ErrMissingName is for when name is required but missing.
	ErrMissingName = errors.New("name required but missing")
	// ErrMissingTarget is for when target is required but missing.
	ErrMissingTarget = errors.New("target required but missing")
)

// Web3Hostname represents a web3 hostname.
type Web3Hostname struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status,omitempty"`
	Target      string     `json:"target,omitempty"`
	Dnslink     string     `json:"dnslink,omitempty"`
	CreatedOn   *time.Time `json:"created_on,omitempty"`
	ModifiedOn  *time.Time `json:"modified_on,omitempty"`
}

// Web3HostnameListParameters represents the parameters for listing web3 hostnames.
type Web3HostnameListParameters struct {
	ZoneID string
}

// Web3HostnameListResponse represents the API response body for listing web3 hostnames.
type Web3HostnameListResponse struct {
	Response
	Result []Web3Hostname `json:"result"`
}

// Web3HostnameCreateParameters represents the parameters for creating a web3 hostname.
type Web3HostnameCreateParameters struct {
	ZoneID      string
	Name        string `json:"name,omitempty"`
	Target      string `json:"target,omitempty"`
	Description string `json:"description,omitempty"`
	DNSLink     string `json:"dnslink,omitempty"`
}

// Web3HostnameResponse represents an API response body for a web3 hostname.
type Web3HostnameResponse struct {
	Response
	Result Web3Hostname `json:"result,omitempty"`
}

// Web3HostnameDetailsParameters represents the parameters for getting a single web3 hostname.
type Web3HostnameDetailsParameters struct {
	ZoneID     string
	Identifier string
}

// Web3HostnameUpdateParameters represents the parameters for editing a web3 hostname.
type Web3HostnameUpdateParameters struct {
	ZoneID      string
	Identifier  string
	Description string `json:"description,omitempty"`
	DNSLink     string `json:"dnslink,omitempty"`
}

// Web3HostnameDeleteResult represents the result of deleting a web3 hostname.
type Web3HostnameDeleteResult struct {
	ID string `json:"id,omitempty"`
}

// Web3HostnameDeleteResponse represents the API response body for deleting a web3 hostname.
type Web3HostnameDeleteResponse struct {
	Response
	Result Web3HostnameDeleteResult `json:"result,omitempty"`
}

// ListWeb3Hostnames lists all web3 hostnames.
//
// API Reference: https://api.cloudflare.com/#web3-hostname-list-web3-hostnames
func (api *API) ListWeb3Hostnames(ctx context.Context, params Web3HostnameListParameters) ([]Web3Hostname, error) {
	if params.ZoneID == "" {
		return []Web3Hostname{}, ErrMissingZoneID
	}

	uri := fmt.Sprintf("/zones/%s/web3/hostnames", params.ZoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []Web3Hostname{}, err
	}
	var web3ListResponse Web3HostnameListResponse
	if err := json.Unmarshal(res, &web3ListResponse); err != nil {
		return []Web3Hostname{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return web3ListResponse.Result, nil
}

// CreateWeb3Hostname creates a web3 hostname.
//
// API Reference: https://api.cloudflare.com/#web3-hostname-create-web3-hostname
func (api *API) CreateWeb3Hostname(ctx context.Context, params Web3HostnameCreateParameters) (Web3Hostname, error) {
	if params.ZoneID == "" {
		return Web3Hostname{}, ErrMissingZoneID
	}
	if params.Name == "" {
		return Web3Hostname{}, ErrMissingName
	}
	if params.Target == "" {
		return Web3Hostname{}, ErrMissingTarget
	}

	uri := fmt.Sprintf("/zones/%s/web3/hostnames", params.ZoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return Web3Hostname{}, err
	}
	var web3Response Web3HostnameResponse
	if err := json.Unmarshal(res, &web3Response); err != nil {
		return Web3Hostname{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return web3Response.Result, nil
}

// GetWeb3Hostname gets a single web3 hostname by identifier.
//
// API Reference: https://api.cloudflare.com/#web3-hostname-web3-hostname-details
func (api *API) GetWeb3Hostname(ctx context.Context, params Web3HostnameDetailsParameters) (Web3Hostname, error) {
	if params.ZoneID == "" {
		return Web3Hostname{}, ErrMissingZoneID
	}
	if params.Identifier == "" {
		return Web3Hostname{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/zones/%s/web3/hostnames/%s", params.ZoneID, params.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return Web3Hostname{}, err
	}
	var web3Response Web3HostnameResponse
	if err := json.Unmarshal(res, &web3Response); err != nil {
		return Web3Hostname{}, err
	}
	return web3Response.Result, nil
}

// UpdateWeb3Hostname edits a web3 hostname.
//
// API Reference: https://api.cloudflare.com/#web3-hostname-edit-web3-hostname
func (api *API) UpdateWeb3Hostname(ctx context.Context, params Web3HostnameUpdateParameters) (Web3Hostname, error) {
	if params.ZoneID == "" {
		return Web3Hostname{}, ErrMissingZoneID
	}
	if params.Identifier == "" {
		return Web3Hostname{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/zones/%s/web3/hostnames/%s", params.ZoneID, params.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return Web3Hostname{}, err
	}
	var web3Response Web3HostnameResponse
	if err := json.Unmarshal(res, &web3Response); err != nil {
		return Web3Hostname{}, err
	}
	return web3Response.Result, nil
}

// DeleteWeb3Hostname deletes a web3 hostname.
//
// API Reference: https://api.cloudflare.com/#web3-hostname-delete-web3-hostname
func (api *API) DeleteWeb3Hostname(ctx context.Context, params Web3HostnameDetailsParameters) (Web3HostnameDeleteResult, error) {
	if params.ZoneID == "" {
		return Web3HostnameDeleteResult{}, ErrMissingZoneID
	}
	if params.Identifier == "" {
		return Web3HostnameDeleteResult{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/zones/%s/web3/hostnames/%s", params.ZoneID, params.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return Web3HostnameDeleteResult{}, err
	}
	var web3Response Web3HostnameDeleteResponse
	if err := json.Unmarshal(res, &web3Response); err != nil {
		return Web3HostnameDeleteResult{}, err
	}
	return web3Response.Result, nil
}
