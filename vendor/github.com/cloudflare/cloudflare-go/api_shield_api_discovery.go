package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// APIShieldDiscoveryOrigin is an enumeration on what discovery engine an operation was discovered by.
type APIShieldDiscoveryOrigin string

const (

	// APIShieldDiscoveryOriginML discovered operations that were sourced using ML API Discovery.
	APIShieldDiscoveryOriginML APIShieldDiscoveryOrigin = "ML"
	// APIShieldDiscoveryOriginSessionIdentifier discovered operations that were sourced using Session Identifier
	// API Discovery.
	APIShieldDiscoveryOriginSessionIdentifier APIShieldDiscoveryOrigin = "SessionIdentifier"
)

// APIShieldDiscoveryState is an enumeration on states a discovery operation can be in.
type APIShieldDiscoveryState string

const (
	// APIShieldDiscoveryStateReview discovered operations that are not saved into API Shield Endpoint Management.
	APIShieldDiscoveryStateReview APIShieldDiscoveryState = "review"
	// APIShieldDiscoveryStateSaved discovered operations that are already saved into API Shield Endpoint Management.
	APIShieldDiscoveryStateSaved APIShieldDiscoveryState = "saved"
	// APIShieldDiscoveryStateIgnored discovered operations that have been marked as ignored.
	APIShieldDiscoveryStateIgnored APIShieldDiscoveryState = "ignored"
)

// APIShieldDiscoveryOperation is an operation that was discovered by API Discovery.
type APIShieldDiscoveryOperation struct {
	// ID represents the ID of the operation, formatted as UUID
	ID string `json:"id"`
	// Origin represents the API discovery engine(s) that discovered this operation
	Origin []APIShieldDiscoveryOrigin `json:"origin"`
	// State represents the state of operation in API Discovery
	State APIShieldDiscoveryState `json:"state"`
	// LastUpdated timestamp of when this operation was last updated
	LastUpdated *time.Time `json:"last_updated"`
	// Features are additional data about the operation
	Features map[string]any `json:"features,omitempty"`

	Method   string `json:"method"`
	Host     string `json:"host"`
	Endpoint string `json:"endpoint"`
}

// ListAPIShieldDiscoveryOperationsParams represents the parameters to pass when retrieving discovered operations.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-api-discovery-retrieve-discovered-operations-on-a-zone
type ListAPIShieldDiscoveryOperationsParams struct {
	// Direction to order results.
	Direction string `url:"direction,omitempty"`
	// OrderBy when requesting a feature, the feature keys are available for ordering as well, e.g., thresholds.suggested_threshold.
	OrderBy string `url:"order,omitempty"`
	// Hosts filters results to only include the specified hosts.
	Hosts []string `url:"host,omitempty"`
	// Methods filters results to only include the specified methods.
	Methods []string `url:"method,omitempty"`
	// Endpoint filters results to only include endpoints containing this pattern.
	Endpoint string `url:"endpoint,omitempty"`
	// Diff when true, only return API Discovery results that are not saved into API Shield Endpoint Management
	Diff bool `url:"diff,omitempty"`
	// Origin filters results to only include discovery results sourced from a particular discovery engine
	// See APIShieldDiscoveryOrigin for valid values.
	Origin APIShieldDiscoveryOrigin `url:"origin,omitempty"`
	// State filters results to only include discovery results in a particular state
	// See APIShieldDiscoveryState for valid values.
	State APIShieldDiscoveryState `url:"state,omitempty"`

	// Pagination options to apply to the request.
	PaginationOptions
}

// UpdateAPIShieldDiscoveryOperationParams represents the parameters to pass to patch a discovery operation
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-api-patch-discovered-operation
type UpdateAPIShieldDiscoveryOperationParams struct {
	// OperationID is the ID, formatted as UUID, of the operation to be updated
	OperationID string                  `json:"-" url:"-"`
	State       APIShieldDiscoveryState `json:"state" url:"-"`
}

// UpdateAPIShieldDiscoveryOperationsParams maps discovery operation IDs to PatchAPIShieldDiscoveryOperation structs
//
// Example:
//
//	UpdateAPIShieldDiscoveryOperations{
//			"99522293-a505-45e5-bbad-bbc339f5dc40": PatchAPIShieldDiscoveryOperation{ State: "review" },
//	}
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-api-patch-discovered-operations
type UpdateAPIShieldDiscoveryOperationsParams map[string]UpdateAPIShieldDiscoveryOperation

// UpdateAPIShieldDiscoveryOperation represents the state to set on a discovery operation.
type UpdateAPIShieldDiscoveryOperation struct {
	// State is the state to set on the operation
	State APIShieldDiscoveryState `json:"state" url:"-"`
}

// APIShieldListDiscoveryOperationsResponse represents the response from the api_gateway/discovery/operations endpoint.
type APIShieldListDiscoveryOperationsResponse struct {
	Result     []APIShieldDiscoveryOperation `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// APIShieldPatchDiscoveryOperationResponse represents the response from the PATCH api_gateway/discovery/operations/{id} endpoint.
type APIShieldPatchDiscoveryOperationResponse struct {
	Result UpdateAPIShieldDiscoveryOperation `json:"result"`
	Response
}

// APIShieldPatchDiscoveryOperationsResponse represents the response from the PATCH api_gateway/discovery/operations endpoint.
type APIShieldPatchDiscoveryOperationsResponse struct {
	Result UpdateAPIShieldDiscoveryOperationsParams `json:"result"`
	Response
}

// ListAPIShieldDiscoveryOperations retrieve the most up to date view of discovered operations.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-api-discovery-retrieve-discovered-operations-on-a-zone
func (api *API) ListAPIShieldDiscoveryOperations(ctx context.Context, rc *ResourceContainer, params ListAPIShieldDiscoveryOperationsParams) ([]APIShieldDiscoveryOperation, ResultInfo, error) {
	uri := buildURI(fmt.Sprintf("/zones/%s/api_gateway/discovery/operations", rc.Identifier), params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var asResponse APIShieldListDiscoveryOperationsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, asResponse.ResultInfo, nil
}

// UpdateAPIShieldDiscoveryOperation updates certain fields on a discovered operation.
//
// API Documentation: https://developers.cloudflare.com/api/operations/api-shield-api-patch-discovered-operation
func (api *API) UpdateAPIShieldDiscoveryOperation(ctx context.Context, rc *ResourceContainer, params UpdateAPIShieldDiscoveryOperationParams) (*UpdateAPIShieldDiscoveryOperation, error) {
	if params.OperationID == "" {
		return nil, fmt.Errorf("operation ID must be provided")
	}

	uri := fmt.Sprintf("/zones/%s/api_gateway/discovery/operations/%s", rc.Identifier, params.OperationID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return nil, err
	}

	// Result should be the updated schema that was patched
	var asResponse APIShieldPatchDiscoveryOperationResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// UpdateAPIShieldDiscoveryOperations bulk updates certain fields on multiple discovered operations
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-api-patch-discovered-operations
func (api *API) UpdateAPIShieldDiscoveryOperations(ctx context.Context, rc *ResourceContainer, params UpdateAPIShieldDiscoveryOperationsParams) (*UpdateAPIShieldDiscoveryOperationsParams, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/discovery/operations", rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return nil, err
	}

	// Result should be the updated schema that was patched
	var asResponse APIShieldPatchDiscoveryOperationsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}
