package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// APIShieldOperation represents an operation stored in API Shield Endpoint Management.
type APIShieldOperation struct {
	APIShieldBasicOperation
	ID          string         `json:"operation_id"`
	LastUpdated *time.Time     `json:"last_updated"`
	Features    map[string]any `json:"features,omitempty"`
}

// GetAPIShieldOperationParams represents the parameters to pass when retrieving an operation.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-an-operation
type GetAPIShieldOperationParams struct {
	// The Operation ID to retrieve
	OperationID string `url:"-"`
	// Features represents a set of features to return in `features` object when
	// performing making read requests against an Operation or listing operations.
	Features []string `url:"feature,omitempty"`
}

// CreateAPIShieldOperationsParams represents the parameters to pass when adding one or more operations.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-add-operations-to-a-zone
type CreateAPIShieldOperationsParams struct {
	// Operations are a slice of operations to be created in API Shield Endpoint Management
	Operations []APIShieldBasicOperation `url:"-"`
}

// APIShieldBasicOperation should be used when creating an operation in API Shield Endpoint Management.
type APIShieldBasicOperation struct {
	Method   string `json:"method"`
	Host     string `json:"host"`
	Endpoint string `json:"endpoint"`
}

// DeleteAPIShieldOperationParams represents the parameters to pass to delete an operation.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-delete-an-operation
type DeleteAPIShieldOperationParams struct {
	// OperationID is the operation to be deleted
	OperationID string `url:"-"`
}

// ListAPIShieldOperationsParams represents the parameters to pass when retrieving operations
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone
type ListAPIShieldOperationsParams struct {
	// Features represents a set of features to return in `features` object when
	// performing making read requests against an Operation or listing operations.
	Features []string `url:"feature,omitempty"`
	// Direction to order results.
	Direction string `url:"direction,omitempty"`
	// OrderBy when requesting a feature, the feature keys are available for ordering as well, e.g., thresholds.suggested_threshold.
	OrderBy string `url:"order,omitempty"`
	// Filters to only return operations that match filtering criteria, see APIShieldGetOperationsFilters
	APIShieldListOperationsFilters
	// Pagination options to apply to the request.
	PaginationOptions
}

// APIShieldListOperationsFilters represents the filtering query parameters to set when retrieving operations
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone
type APIShieldListOperationsFilters struct {
	// Hosts filters results to only include the specified hosts.
	Hosts []string `url:"host,omitempty"`
	// Methods filters results to only include the specified methods.
	Methods []string `url:"method,omitempty"`
	// Endpoint filter results to only include endpoints containing this pattern.
	Endpoint string `url:"endpoint,omitempty"`
}

// APIShieldGetOperationResponse represents the response from the api_gateway/operations/{id} endpoint.
type APIShieldGetOperationResponse struct {
	Result APIShieldOperation `json:"result"`
	Response
}

// APIShieldGetOperationsResponse represents the response from the api_gateway/operations endpoint.
type APIShieldGetOperationsResponse struct {
	Result     []APIShieldOperation `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// APIShieldDeleteOperationResponse represents the response from the api_gateway/operations/{id} endpoint (DELETE).
type APIShieldDeleteOperationResponse struct {
	Result interface{} `json:"result"`
	Response
}

// GetAPIShieldOperation returns information about an operation
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-an-operation
func (api *API) GetAPIShieldOperation(ctx context.Context, rc *ResourceContainer, params GetAPIShieldOperationParams) (*APIShieldOperation, error) {
	path := fmt.Sprintf("/zones/%s/api_gateway/operations/%s", rc.Identifier, params.OperationID)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var asResponse APIShieldGetOperationResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// ListAPIShieldOperations retrieve information about all operations on a zone
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-retrieve-information-about-all-operations-on-a-zone
func (api *API) ListAPIShieldOperations(ctx context.Context, rc *ResourceContainer, params ListAPIShieldOperationsParams) ([]APIShieldOperation, ResultInfo, error) {
	path := fmt.Sprintf("/zones/%s/api_gateway/operations", rc.Identifier)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var asResponse APIShieldGetOperationsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, asResponse.ResultInfo, nil
}

// CreateAPIShieldOperations add one or more operations to a zone.
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-add-operations-to-a-zone
func (api *API) CreateAPIShieldOperations(ctx context.Context, rc *ResourceContainer, params CreateAPIShieldOperationsParams) ([]APIShieldOperation, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/operations", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params.Operations)
	if err != nil {
		return nil, err
	}

	// Result should be all the operations added to the zone, similar to doing GetAPIShieldOperations
	var asResponse APIShieldGetOperationsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, nil
}

// DeleteAPIShieldOperation deletes a single operation
//
// API documentation https://developers.cloudflare.com/api/operations/api-shield-endpoint-management-delete-an-operation
func (api *API) DeleteAPIShieldOperation(ctx context.Context, rc *ResourceContainer, params DeleteAPIShieldOperationParams) error {
	uri := fmt.Sprintf("/zones/%s/api_gateway/operations/%s", rc.Identifier, params.OperationID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var asResponse APIShieldDeleteOperationResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}
