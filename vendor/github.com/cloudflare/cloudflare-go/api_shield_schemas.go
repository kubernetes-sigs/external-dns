package cloudflare

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/goccy/go-json"
)

// APIShieldSchema represents a schema stored in API Shield Schema Validation 2.0.
type APIShieldSchema struct {
	// ID represents the ID of the schema
	ID string `json:"schema_id"`
	// Name represents the name of the schema
	Name string `json:"name"`
	// Kind of the schema
	Kind string `json:"kind"`
	// Source is the contents of the schema
	Source string `json:"source,omitempty"`
	// CreatedAt is the time the schema was created
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ValidationEnabled controls if schema is used for validation
	ValidationEnabled bool `json:"validation_enabled,omitempty"`
}

// CreateAPIShieldSchemaParams represents the parameters to pass when creating a schema in Schema Validation 2.0.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-post-schema
type CreateAPIShieldSchemaParams struct {
	// Source is a io.Reader containing the contents of the schema
	Source io.Reader
	// Name represents the name of the schema.
	Name string
	// Kind of the schema. This is always set to openapi_v3.
	Kind string
	// ValidationEnabled controls if schema is used for validation
	ValidationEnabled *bool
}

// GetAPIShieldSchemaParams represents the parameters to pass when retrieving a schema with a given schema ID.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-retrieve-information-about-specific-schema
type GetAPIShieldSchemaParams struct {
	// SchemaID is the ID of the schema to retrieve
	SchemaID string `url:"-"`

	// OmitSource specifies whether the contents of the schema should be returned in the "Source" field.
	OmitSource *bool `url:"omit_source,omitempty"`
}

// ListAPIShieldSchemasParams represents the parameters to pass when retrieving all schemas.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-retrieve-information-about-all-schemas
type ListAPIShieldSchemasParams struct {
	// OmitSource specifies whether the contents of the schema should be returned in the "Source" field.
	OmitSource *bool `url:"omit_source,omitempty"`

	// ValidationEnabled specifies whether to return only schemas that have validation enabled.
	ValidationEnabled *bool `url:"validation_enabled,omitempty"`

	// PaginationOptions to apply to the request.
	PaginationOptions
}

// DeleteAPIShieldSchemaParams represents the parameters to pass to delete a schema.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-delete-a-schema
type DeleteAPIShieldSchemaParams struct {
	// SchemaID is the schema to be deleted
	SchemaID string `url:"-"`
}

// UpdateAPIShieldSchemaParams represents the parameters to pass to patch certain fields on an existing schema
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-enable-validation-for-a-schema
type UpdateAPIShieldSchemaParams struct {
	// SchemaID is the schema to be patched
	SchemaID string `json:"-" url:"-"`

	// ValidationEnabled controls if schema is used for validation
	ValidationEnabled *bool `json:"validation_enabled" url:"-"`
}

// APIShieldGetSchemaResponse represents the response from the GET api_gateway/user_schemas/{id} endpoint.
type APIShieldGetSchemaResponse struct {
	Result APIShieldSchema `json:"result"`
	Response
}

// APIShieldListSchemasResponse represents the response from the GET api_gateway/user_schemas endpoint.
type APIShieldListSchemasResponse struct {
	Result     []APIShieldSchema `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

// APIShieldCreateSchemaResponse represents the response from the POST api_gateway/user_schemas endpoint.
type APIShieldCreateSchemaResponse struct {
	Result APIShieldCreateSchemaResult `json:"result"`
	Response
}

// APIShieldDeleteSchemaResponse represents the response from the DELETE api_gateway/user_schemas/{id} endpoint.
type APIShieldDeleteSchemaResponse struct {
	Result interface{} `json:"result"`
	Response
}

// APIShieldPatchSchemaResponse represents the response from the PATCH api_gateway/user_schemas/{id} endpoint.
type APIShieldPatchSchemaResponse struct {
	Result APIShieldSchema `json:"result"`
	Response
}

// APIShieldCreateSchemaResult represents the successful result of creating a schema in Schema Validation 2.0.
type APIShieldCreateSchemaResult struct {
	// APIShieldSchema is the schema that was created
	Schema APIShieldSchema `json:"schema"`
	// APIShieldCreateSchemaEvents are non-critical event logs that occurred during processing.
	Events APIShieldCreateSchemaEvents `json:"upload_details"`
}

// APIShieldCreateSchemaEvents are event logs that occurred during processing.
//
// The logs are separated into levels of severity.
type APIShieldCreateSchemaEvents struct {
	Critical *APIShieldCreateSchemaEventWithLocation   `json:"critical,omitempty"`
	Errors   []APIShieldCreateSchemaEventWithLocations `json:"errors,omitempty"`
	Warnings []APIShieldCreateSchemaEventWithLocations `json:"warnings,omitempty"`
}

// APIShieldCreateSchemaEvent is an event log that occurred during processing.
type APIShieldCreateSchemaEvent struct {
	// Code identifies the event that occurred
	Code uint `json:"code"`
	// Message describes the event that occurred
	Message string `json:"message"`
}

// APIShieldCreateSchemaEventWithLocation is an event log that occurred during processing, with the location
// in the schema where the event occurred.
type APIShieldCreateSchemaEventWithLocation struct {
	APIShieldCreateSchemaEvent

	// Location is where the event occurred
	// See https://goessner.net/articles/JsonPath/ for JSONPath specification.
	Location string `json:"location,omitempty"`
}

// APIShieldCreateSchemaEventWithLocations is an event log that occurred during processing, with the location(s)
// in the schema where the event occurred.
type APIShieldCreateSchemaEventWithLocations struct {
	APIShieldCreateSchemaEvent

	// Locations lists JSONPath locations where the event occurred
	// See https://goessner.net/articles/JsonPath/ for JSONPath specification
	Locations []string `json:"locations"`
}

func (cse APIShieldCreateSchemaEventWithLocations) String() string {
	var s string
	s += cse.Message

	if len(cse.Locations) == 0 || (len(cse.Locations) == 1 && cse.Locations[0] == "") {
		// append nothing
	} else if len(cse.Locations) == 1 {
		s += fmt.Sprintf(" (%s)", cse.Locations[0])
	} else {
		s += " (multiple locations)"
	}

	return s
}

// GetAPIShieldSchema retrieves information about a specific schema on a zone
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-retrieve-information-about-specific-schema
func (api *API) GetAPIShieldSchema(ctx context.Context, rc *ResourceContainer, params GetAPIShieldSchemaParams) (*APIShieldSchema, error) {
	if params.SchemaID == "" {
		return nil, fmt.Errorf("schema ID must be provided")
	}

	path := fmt.Sprintf("/zones/%s/api_gateway/user_schemas/%s", rc.Identifier, params.SchemaID)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var asResponse APIShieldGetSchemaResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// ListAPIShieldSchemas retrieves all schemas for a zone
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-retrieve-information-about-all-schemas
func (api *API) ListAPIShieldSchemas(ctx context.Context, rc *ResourceContainer, params ListAPIShieldSchemasParams) ([]APIShieldSchema, ResultInfo, error) {
	path := fmt.Sprintf("/zones/%s/api_gateway/user_schemas", rc.Identifier)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, ResultInfo{}, err
	}

	var asResponse APIShieldListSchemasResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return asResponse.Result, asResponse.ResultInfo, nil
}

// CreateAPIShieldSchema uploads a schema to a zone
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-post-schema
func (api *API) CreateAPIShieldSchema(ctx context.Context, rc *ResourceContainer, params CreateAPIShieldSchemaParams) (*APIShieldCreateSchemaResult, error) {
	uri := fmt.Sprintf("/zones/%s/api_gateway/user_schemas", rc.Identifier)

	if params.Name == "" {
		return nil, fmt.Errorf("name must not be empty")
	}

	if params.Source == nil {
		return nil, fmt.Errorf("source must not be nil")
	}

	// Prepare the form to be submitted
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	// write fields
	if err := w.WriteField("name", params.Name); err != nil {
		return nil, fmt.Errorf("error during multi-part form construction: %w", err)
	}
	if err := w.WriteField("kind", params.Kind); err != nil {
		return nil, fmt.Errorf("error during multi-part form construction: %w", err)
	}

	if params.ValidationEnabled != nil {
		if err := w.WriteField("validation_enabled", strconv.FormatBool(*params.ValidationEnabled)); err != nil {
			return nil, fmt.Errorf("error during multi-part form construction: %w", err)
		}
	}

	// write schema contents
	part, err := w.CreateFormFile("file", params.Name)
	if err != nil {
		return nil, fmt.Errorf("error during multi-part form construction: %w", err)
	}
	if _, err := io.Copy(part, params.Source); err != nil {
		return nil, fmt.Errorf("error during multi-part form construction: %w", err)
	}
	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("error during multi-part form construction: %w", err)
	}

	res, err := api.makeRequestContextWithHeaders(ctx, http.MethodPost, uri, &b, http.Header{
		"Content-Type": []string{w.FormDataContentType()},
	})
	if err != nil {
		return nil, err
	}

	var asResponse APIShieldCreateSchemaResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// DeleteAPIShieldSchema deletes a single schema
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-delete-a-schema
func (api *API) DeleteAPIShieldSchema(ctx context.Context, rc *ResourceContainer, params DeleteAPIShieldSchemaParams) error {
	if params.SchemaID == "" {
		return fmt.Errorf("schema ID must be provided")
	}

	uri := fmt.Sprintf("/zones/%s/api_gateway/user_schemas/%s", rc.Identifier, params.SchemaID)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var asResponse APIShieldDeleteSchemaResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// UpdateAPIShieldSchema updates certain fields on an existing schema.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-enable-validation-for-a-schema
func (api *API) UpdateAPIShieldSchema(ctx context.Context, rc *ResourceContainer, params UpdateAPIShieldSchemaParams) (*APIShieldSchema, error) {
	if params.SchemaID == "" {
		return nil, fmt.Errorf("schema ID must be provided")
	}

	uri := fmt.Sprintf("/zones/%s/api_gateway/user_schemas/%s", rc.Identifier, params.SchemaID)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return nil, err
	}

	// Result should be the updated schema that was patched
	var asResponse APIShieldPatchSchemaResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// Schema Validation Settings

// APIShieldSchemaValidationSettings represents zone level schema validation settings for
// API Shield Schema Validation 2.0.
type APIShieldSchemaValidationSettings struct {
	// DefaultMitigationAction is the mitigation to apply when there is no operation-level
	// mitigation action defined
	DefaultMitigationAction string `json:"validation_default_mitigation_action" url:"-"`
	// OverrideMitigationAction when set, will apply to all requests regardless of
	// zone-level/operation-level setting
	OverrideMitigationAction *string `json:"validation_override_mitigation_action" url:"-"`
}

// UpdateAPIShieldSchemaValidationSettingsParams represents the parameters to pass to update certain fields
// on schema validation settings on the zone
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-patch-zone-level-settings
type UpdateAPIShieldSchemaValidationSettingsParams struct {
	// DefaultMitigationAction is the mitigation to apply when there is no operation-level
	// mitigation action defined
	//
	// passing a `nil` value will have no effect on this setting
	DefaultMitigationAction *string `json:"validation_default_mitigation_action" url:"-"`

	// OverrideMitigationAction when set, will apply to all requests regardless of
	// zone-level/operation-level setting
	//
	// passing a `nil` value will have no effect on this setting
	OverrideMitigationAction *string `json:"validation_override_mitigation_action" url:"-"`
}

// APIShieldSchemaValidationSettingsResponse represents the response from the GET api_gateway/settings/schema_validation endpoint.
type APIShieldSchemaValidationSettingsResponse struct {
	Result APIShieldSchemaValidationSettings `json:"result"`
	Response
}

// GetAPIShieldSchemaValidationSettings retrieves zone level schema validation settings
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-retrieve-zone-level-settings
func (api *API) GetAPIShieldSchemaValidationSettings(ctx context.Context, rc *ResourceContainer) (*APIShieldSchemaValidationSettings, error) {
	path := fmt.Sprintf("/zones/%s/api_gateway/settings/schema_validation", rc.Identifier)

	uri := buildURI(path, nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	var asResponse APIShieldSchemaValidationSettingsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// UpdateAPIShieldSchemaValidationSettings updates certain fields on zone level schema validation settings
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-patch-zone-level-settings
func (api *API) UpdateAPIShieldSchemaValidationSettings(ctx context.Context, rc *ResourceContainer, params UpdateAPIShieldSchemaValidationSettingsParams) (*APIShieldSchemaValidationSettings, error) {
	path := fmt.Sprintf("/zones/%s/api_gateway/settings/schema_validation", rc.Identifier)

	uri := buildURI(path, params)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return nil, err
	}

	var asResponse APIShieldSchemaValidationSettingsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// APIShieldOperationSchemaValidationSettings represents operation level schema validation settings for
// API Shield Schema Validation 2.0.
type APIShieldOperationSchemaValidationSettings struct {
	// MitigationAction is the mitigation to apply to the operation
	MitigationAction *string `json:"mitigation_action" url:"-"`
}

// GetAPIShieldOperationSchemaValidationSettingsParams represents the parameters to pass to retrieve
// the schema validation settings set on the operation.
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-retrieve-operation-level-settings
type GetAPIShieldOperationSchemaValidationSettingsParams struct {
	// The Operation ID to apply the mitigation action to
	OperationID string `url:"-"`
}

// UpdateAPIShieldOperationSchemaValidationSettings maps operation IDs to APIShieldOperationSchemaValidationSettings
//
// # This can be used to bulk update operations in one call
//
// Example:
//
//	UpdateAPIShieldOperationSchemaValidationSettings{
//			"99522293-a505-45e5-bbad-bbc339f5dc40": APIShieldOperationSchemaValidationSettings{ MitigationAction: nil },
//	}
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-update-multiple-operation-level-settings
type UpdateAPIShieldOperationSchemaValidationSettings map[string]APIShieldOperationSchemaValidationSettings

// APIShieldOperationSchemaValidationSettingsResponse represents the response from the GET api_gateway/operation/{operationID}/schema_validation endpoint.
type APIShieldOperationSchemaValidationSettingsResponse struct {
	Result APIShieldOperationSchemaValidationSettings `json:"result"`
	Response
}

// UpdateAPIShieldOperationSchemaValidationSettingsResponse represents the response from the PATCH api_gateway/operations/schema_validation endpoint.
type UpdateAPIShieldOperationSchemaValidationSettingsResponse struct {
	Result UpdateAPIShieldOperationSchemaValidationSettings `json:"result"`
	Response
}

// GetAPIShieldOperationSchemaValidationSettings retrieves operation level schema validation settings
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-retrieve-operation-level-settings
func (api *API) GetAPIShieldOperationSchemaValidationSettings(ctx context.Context, rc *ResourceContainer, params GetAPIShieldOperationSchemaValidationSettingsParams) (*APIShieldOperationSchemaValidationSettings, error) {
	if params.OperationID == "" {
		return nil, fmt.Errorf("operation ID must be provided")
	}

	path := fmt.Sprintf("/zones/%s/api_gateway/operations/%s/schema_validation", rc.Identifier, params.OperationID)

	uri := buildURI(path, nil)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, params)
	if err != nil {
		return nil, err
	}

	var asResponse APIShieldOperationSchemaValidationSettingsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}

// UpdateAPIShieldOperationSchemaValidationSettings update multiple operation level schema validation settings
//
// API documentation: https://developers.cloudflare.com/api/operations/api-shield-schema-validation-update-multiple-operation-level-settings
func (api *API) UpdateAPIShieldOperationSchemaValidationSettings(ctx context.Context, rc *ResourceContainer, params UpdateAPIShieldOperationSchemaValidationSettings) (*UpdateAPIShieldOperationSchemaValidationSettings, error) {
	path := fmt.Sprintf("/zones/%s/api_gateway/operations/schema_validation", rc.Identifier)

	uri := buildURI(path, nil)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return nil, err
	}

	var asResponse UpdateAPIShieldOperationSchemaValidationSettingsResponse
	err = json.Unmarshal(res, &asResponse)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return &asResponse.Result, nil
}
