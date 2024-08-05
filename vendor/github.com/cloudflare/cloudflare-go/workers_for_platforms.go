package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type WorkersForPlatformsDispatchNamespace struct {
	NamespaceId   string     `json:"namespace_id"`
	NamespaceName string     `json:"namespace_name"`
	CreatedOn     *time.Time `json:"created_on,omitempty"`
	CreatedBy     string     `json:"created_by"`
	ModifiedOn    *time.Time `json:"modified_on,omitempty"`
	ModifiedBy    string     `json:"modified_by"`
}

type ListWorkersForPlatformsDispatchNamespaceResponse struct {
	Response
	Result []WorkersForPlatformsDispatchNamespace `json:"result"`
}

type GetWorkersForPlatformsDispatchNamespaceResponse struct {
	Response
	Result WorkersForPlatformsDispatchNamespace `json:"result"`
}

type CreateWorkersForPlatformsDispatchNamespaceParams struct {
	Name string `json:"name"`
}

// ListWorkersForPlatformsDispatchNamespaces lists the dispatch namespaces.
//
// API reference: https://developers.cloudflare.com/api/operations/namespace-worker-list
func (api *API) ListWorkersForPlatformsDispatchNamespaces(ctx context.Context, rc *ResourceContainer) (*ListWorkersForPlatformsDispatchNamespaceResponse, error) {
	if rc.Level != AccountRouteLevel {
		return nil, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return nil, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/dispatch/namespaces", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)

	var r ListWorkersForPlatformsDispatchNamespaceResponse
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &r, nil
}

// GetWorkersForPlatformsDispatchNamespace gets a specific dispatch namespace.
//
// API reference: https://developers.cloudflare.com/api/operations/namespace-worker-get-namespace
func (api *API) GetWorkersForPlatformsDispatchNamespace(ctx context.Context, rc *ResourceContainer, name string) (*GetWorkersForPlatformsDispatchNamespaceResponse, error) {
	if rc.Level != AccountRouteLevel {
		return nil, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return nil, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/dispatch/namespaces/%s", rc.Identifier, name)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)

	var r GetWorkersForPlatformsDispatchNamespaceResponse
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &r, nil
}

// CreateWorkersForPlatformsDispatchNamespace creates a new dispatch namespace.
//
// API reference: https://developers.cloudflare.com/api/operations/namespace-worker-create
func (api *API) CreateWorkersForPlatformsDispatchNamespace(ctx context.Context, rc *ResourceContainer, params CreateWorkersForPlatformsDispatchNamespaceParams) (*GetWorkersForPlatformsDispatchNamespaceResponse, error) {
	if rc.Level != AccountRouteLevel {
		return nil, ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return nil, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/dispatch/namespaces", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)

	var r GetWorkersForPlatformsDispatchNamespaceResponse
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &r, nil
}

// DeleteWorkersForPlatformsDispatchNamespace deletes a dispatch namespace.
//
// API reference: https://developers.cloudflare.com/api/operations/namespace-worker-delete-namespace
func (api *API) DeleteWorkersForPlatformsDispatchNamespace(ctx context.Context, rc *ResourceContainer, name string) error {
	if rc.Level != AccountRouteLevel {
		return ErrRequiredAccountLevelResourceContainer
	}

	if rc.Identifier == "" {
		return ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/workers/dispatch/namespaces/%s", rc.Identifier, name)
	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)

	if err != nil {
		return err
	}

	return nil
}
