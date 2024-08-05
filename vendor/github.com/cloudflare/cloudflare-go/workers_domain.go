package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

var (
	ErrMissingHostname    = errors.New("required hostname missing")
	ErrMissingService     = errors.New("required service missing")
	ErrMissingEnvironment = errors.New("required environment missing")
)

type AttachWorkersDomainParams struct {
	ID          string `json:"id,omitempty"`
	ZoneID      string `json:"zone_id,omitempty"`
	ZoneName    string `json:"zone_name,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Service     string `json:"service,omitempty"`
	Environment string `json:"environment,omitempty"`
}

type WorkersDomain struct {
	ID          string `json:"id,omitempty"`
	ZoneID      string `json:"zone_id,omitempty"`
	ZoneName    string `json:"zone_name,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Service     string `json:"service,omitempty"`
	Environment string `json:"environment,omitempty"`
}

type WorkersDomainResponse struct {
	Response
	Result WorkersDomain `json:"result"`
}

type ListWorkersDomainParams struct {
	ZoneID      string `url:"zone_id,omitempty"`
	ZoneName    string `url:"zone_name,omitempty"`
	Hostname    string `url:"hostname,omitempty"`
	Service     string `url:"service,omitempty"`
	Environment string `url:"environment,omitempty"`
}

type WorkersDomainListResponse struct {
	Response
	Result []WorkersDomain `json:"result"`
}

// ListWorkersDomains lists all Worker Domains.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-domain-list-domains
func (api *API) ListWorkersDomains(ctx context.Context, rc *ResourceContainer, params ListWorkersDomainParams) ([]WorkersDomain, error) {
	if rc.Identifier == "" {
		return []WorkersDomain{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/workers/domains", rc.Identifier), params)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []WorkersDomain{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r WorkersDomainListResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return []WorkersDomain{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// AttachWorkersDomain attaches a worker to a zone and hostname.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-domain-attach-to-domain
func (api *API) AttachWorkersDomain(ctx context.Context, rc *ResourceContainer, domain AttachWorkersDomainParams) (WorkersDomain, error) {
	if rc.Identifier == "" {
		return WorkersDomain{}, ErrMissingAccountID
	}

	if domain.ZoneID == "" {
		return WorkersDomain{}, ErrMissingZoneID
	}

	if domain.Hostname == "" {
		return WorkersDomain{}, ErrMissingHostname
	}

	if domain.Service == "" {
		return WorkersDomain{}, ErrMissingService
	}

	if domain.Environment == "" {
		return WorkersDomain{}, ErrMissingEnvironment
	}

	uri := fmt.Sprintf("/accounts/%s/workers/domains", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, domain)
	if err != nil {
		return WorkersDomain{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r WorkersDomainResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return WorkersDomain{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// GetWorkersDomain gets a single Worker Domain.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-domain-get-a-domain
func (api *API) GetWorkersDomain(ctx context.Context, rc *ResourceContainer, domainID string) (WorkersDomain, error) {
	if rc.Identifier == "" {
		return WorkersDomain{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/domains/%s", rc.Identifier, domainID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkersDomain{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var r WorkersDomainResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return WorkersDomain{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r.Result, nil
}

// DetachWorkersDomain detaches a worker from a zone and hostname.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-domain-detach-from-domain
func (api *API) DetachWorkersDomain(ctx context.Context, rc *ResourceContainer, domainID string) error {
	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/domains/%s", rc.Identifier, domainID)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	return nil
}
