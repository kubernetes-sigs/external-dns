package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

var ErrMissingWorkerRouteID = errors.New("missing required route ID")

type ListWorkerRoutes struct{}

type CreateWorkerRouteParams struct {
	Pattern string `json:"pattern"`
	Script  string `json:"script,omitempty"`
}

type ListWorkerRoutesParams struct{}

type UpdateWorkerRouteParams struct {
	ID      string `json:"id,omitempty"`
	Pattern string `json:"pattern"`
	Script  string `json:"script,omitempty"`
}

// CreateWorkerRoute creates worker route for a script.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-routes-create-route
func (api *API) CreateWorkerRoute(ctx context.Context, rc *ResourceContainer, params CreateWorkerRouteParams) (WorkerRouteResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return WorkerRouteResponse{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if rc.Identifier == "" {
		return WorkerRouteResponse{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/zones/%s/workers/routes", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return WorkerRouteResponse{}, err
	}

	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// DeleteWorkerRoute deletes worker route for a script.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-routes-delete-route
func (api *API) DeleteWorkerRoute(ctx context.Context, rc *ResourceContainer, routeID string) (WorkerRouteResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return WorkerRouteResponse{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if rc.Identifier == "" {
		return WorkerRouteResponse{}, ErrMissingIdentifier
	}

	if routeID == "" {
		return WorkerRouteResponse{}, errors.New("missing required route ID")
	}

	uri := fmt.Sprintf("/zones/%s/workers/routes/%s", rc.Identifier, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// ListWorkerRoutes returns list of Worker routes.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-routes-list-routes
func (api *API) ListWorkerRoutes(ctx context.Context, rc *ResourceContainer, params ListWorkerRoutesParams) (WorkerRoutesResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return WorkerRoutesResponse{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if rc.Identifier == "" {
		return WorkerRoutesResponse{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/zones/%s/workers/routes", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerRoutesResponse{}, err
	}
	var r WorkerRoutesResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRoutesResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return r, nil
}

// GetWorkerRoute returns a Workers route.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-routes-get-route
func (api *API) GetWorkerRoute(ctx context.Context, rc *ResourceContainer, routeID string) (WorkerRouteResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return WorkerRouteResponse{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if rc.Identifier == "" {
		return WorkerRouteResponse{}, ErrMissingIdentifier
	}

	uri := fmt.Sprintf("/zones/%s/workers/routes/%s", rc.Identifier, routeID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}

// UpdateWorkerRoute updates worker route for a script.
//
// API reference: https://developers.cloudflare.com/api/operations/worker-routes-update-route
func (api *API) UpdateWorkerRoute(ctx context.Context, rc *ResourceContainer, params UpdateWorkerRouteParams) (WorkerRouteResponse, error) {
	if rc.Level != ZoneRouteLevel {
		return WorkerRouteResponse{}, fmt.Errorf(errInvalidResourceContainerAccess, ZoneRouteLevel)
	}

	if rc.Identifier == "" {
		return WorkerRouteResponse{}, ErrMissingIdentifier
	}

	if params.ID == "" {
		return WorkerRouteResponse{}, ErrMissingWorkerRouteID
	}

	uri := fmt.Sprintf("/zones/%s/workers/routes/%s", rc.Identifier, params.ID)
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return WorkerRouteResponse{}, err
	}
	var r WorkerRouteResponse
	err = json.Unmarshal(res, &r)
	if err != nil {
		return WorkerRouteResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r, nil
}
