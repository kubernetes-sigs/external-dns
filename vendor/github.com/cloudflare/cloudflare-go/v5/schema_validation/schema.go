// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/api_gateway"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// SchemaService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSchemaService] method instead.
type SchemaService struct {
	Options []option.RequestOption
}

// NewSchemaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSchemaService(opts ...option.RequestOption) (r *SchemaService) {
	r = &SchemaService{}
	r.Options = opts
	return
}

// Upload a schema
func (r *SchemaService) New(ctx context.Context, params SchemaNewParams, opts ...option.RequestOption) (res *SchemaNewResponse, err error) {
	var env SchemaNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/schemas", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all uploaded schemas
func (r *SchemaService) List(ctx context.Context, params SchemaListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SchemaListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/schemas", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all uploaded schemas
func (r *SchemaService) ListAutoPaging(ctx context.Context, params SchemaListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SchemaListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a schema
func (r *SchemaService) Delete(ctx context.Context, schemaID string, body SchemaDeleteParams, opts ...option.RequestOption) (res *SchemaDeleteResponse, err error) {
	var env SchemaDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/schemas/%s", body.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit details of a schema to enable validation
func (r *SchemaService) Edit(ctx context.Context, schemaID string, params SchemaEditParams, opts ...option.RequestOption) (res *SchemaEditResponse, err error) {
	var env SchemaEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/schemas/%s", params.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details of a schema
func (r *SchemaService) Get(ctx context.Context, schemaID string, params SchemaGetParams, opts ...option.RequestOption) (res *SchemaGetResponse, err error) {
	var env SchemaGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/schemas/%s", params.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A schema used in schema validation
type SchemaNewResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The kind of the schema
	Kind SchemaNewResponseKind `json:"kind,required"`
	// A human-readable name for the schema
	Name string `json:"name,required"`
	// A unique identifier of this schema
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// The raw schema, e.g., the OpenAPI schema, either as JSON or YAML
	Source string `json:"source,required"`
	// An indicator if this schema is enabled
	ValidationEnabled bool                  `json:"validation_enabled"`
	JSON              schemaNewResponseJSON `json:"-"`
}

// schemaNewResponseJSON contains the JSON metadata for the struct
// [SchemaNewResponse]
type schemaNewResponseJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SchemaNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaNewResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the schema
type SchemaNewResponseKind string

const (
	SchemaNewResponseKindOpenAPIV3 SchemaNewResponseKind = "openapi_v3"
)

func (r SchemaNewResponseKind) IsKnown() bool {
	switch r {
	case SchemaNewResponseKindOpenAPIV3:
		return true
	}
	return false
}

// A schema used in schema validation
type SchemaListResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The kind of the schema
	Kind SchemaListResponseKind `json:"kind,required"`
	// A human-readable name for the schema
	Name string `json:"name,required"`
	// A unique identifier of this schema
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// The raw schema, e.g., the OpenAPI schema, either as JSON or YAML
	Source string `json:"source,required"`
	// An indicator if this schema is enabled
	ValidationEnabled bool                   `json:"validation_enabled"`
	JSON              schemaListResponseJSON `json:"-"`
}

// schemaListResponseJSON contains the JSON metadata for the struct
// [SchemaListResponse]
type schemaListResponseJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SchemaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaListResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the schema
type SchemaListResponseKind string

const (
	SchemaListResponseKindOpenAPIV3 SchemaListResponseKind = "openapi_v3"
)

func (r SchemaListResponseKind) IsKnown() bool {
	switch r {
	case SchemaListResponseKindOpenAPIV3:
		return true
	}
	return false
}

type SchemaDeleteResponse struct {
	// The ID of the schema that was just deleted
	SchemaID string                   `json:"schema_id,required" format:"uuid"`
	JSON     schemaDeleteResponseJSON `json:"-"`
}

// schemaDeleteResponseJSON contains the JSON metadata for the struct
// [SchemaDeleteResponse]
type schemaDeleteResponseJSON struct {
	SchemaID    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A schema used in schema validation
type SchemaEditResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The kind of the schema
	Kind SchemaEditResponseKind `json:"kind,required"`
	// A human-readable name for the schema
	Name string `json:"name,required"`
	// A unique identifier of this schema
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// The raw schema, e.g., the OpenAPI schema, either as JSON or YAML
	Source string `json:"source,required"`
	// An indicator if this schema is enabled
	ValidationEnabled bool                   `json:"validation_enabled"`
	JSON              schemaEditResponseJSON `json:"-"`
}

// schemaEditResponseJSON contains the JSON metadata for the struct
// [SchemaEditResponse]
type schemaEditResponseJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SchemaEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaEditResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the schema
type SchemaEditResponseKind string

const (
	SchemaEditResponseKindOpenAPIV3 SchemaEditResponseKind = "openapi_v3"
)

func (r SchemaEditResponseKind) IsKnown() bool {
	switch r {
	case SchemaEditResponseKindOpenAPIV3:
		return true
	}
	return false
}

// A schema used in schema validation
type SchemaGetResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The kind of the schema
	Kind SchemaGetResponseKind `json:"kind,required"`
	// A human-readable name for the schema
	Name string `json:"name,required"`
	// A unique identifier of this schema
	SchemaID string `json:"schema_id,required" format:"uuid"`
	// The raw schema, e.g., the OpenAPI schema, either as JSON or YAML
	Source string `json:"source,required"`
	// An indicator if this schema is enabled
	ValidationEnabled bool                  `json:"validation_enabled"`
	JSON              schemaGetResponseJSON `json:"-"`
}

// schemaGetResponseJSON contains the JSON metadata for the struct
// [SchemaGetResponse]
type schemaGetResponseJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *SchemaGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaGetResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the schema
type SchemaGetResponseKind string

const (
	SchemaGetResponseKindOpenAPIV3 SchemaGetResponseKind = "openapi_v3"
)

func (r SchemaGetResponseKind) IsKnown() bool {
	switch r {
	case SchemaGetResponseKindOpenAPIV3:
		return true
	}
	return false
}

type SchemaNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The kind of the schema
	Kind param.Field[SchemaNewParamsKind] `json:"kind,required"`
	// A human-readable name for the schema
	Name param.Field[string] `json:"name,required"`
	// The raw schema, e.g., the OpenAPI schema, either as JSON or YAML
	Source param.Field[string] `json:"source,required"`
	// An indicator if this schema is enabled
	ValidationEnabled param.Field[bool] `json:"validation_enabled,required"`
}

func (r SchemaNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the schema
type SchemaNewParamsKind string

const (
	SchemaNewParamsKindOpenAPIV3 SchemaNewParamsKind = "openapi_v3"
)

func (r SchemaNewParamsKind) IsKnown() bool {
	switch r {
	case SchemaNewParamsKindOpenAPIV3:
		return true
	}
	return false
}

type SchemaNewResponseEnvelope struct {
	Errors   []SchemaNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SchemaNewResponseEnvelopeMessages `json:"messages,required"`
	// A schema used in schema validation
	Result SchemaNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SchemaNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    schemaNewResponseEnvelopeJSON    `json:"-"`
}

// schemaNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SchemaNewResponseEnvelope]
type schemaNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SchemaNewResponseEnvelopeErrors struct {
	// A unique error code that describes the kind of issue with the schema
	Code int64 `json:"code,required"`
	// A short text explaining the issue with the schema
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           SchemaNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             schemaNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// schemaNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SchemaNewResponseEnvelopeErrors]
type schemaNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SchemaNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SchemaNewResponseEnvelopeErrorsSource struct {
	// A list of JSON path expression(s) that describe the location(s) of the issue
	// within the provided resource. See
	// [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/)
	// for JSONPath specification.
	Locations []string                                  `json:"locations"`
	Pointer   string                                    `json:"pointer"`
	JSON      schemaNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// schemaNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [SchemaNewResponseEnvelopeErrorsSource]
type schemaNewResponseEnvelopeErrorsSourceJSON struct {
	Locations   apijson.Field
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SchemaNewResponseEnvelopeMessages struct {
	// A unique error code that describes the kind of issue with the schema
	Code int64 `json:"code,required"`
	// A short text explaining the issue with the schema
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           SchemaNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             schemaNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// schemaNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SchemaNewResponseEnvelopeMessages]
type schemaNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SchemaNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SchemaNewResponseEnvelopeMessagesSource struct {
	// A list of JSON path expression(s) that describe the location(s) of the issue
	// within the provided resource. See
	// [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/)
	// for JSONPath specification.
	Locations []string                                    `json:"locations"`
	Pointer   string                                      `json:"pointer"`
	JSON      schemaNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// schemaNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [SchemaNewResponseEnvelopeMessagesSource]
type schemaNewResponseEnvelopeMessagesSourceJSON struct {
	Locations   apijson.Field
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SchemaNewResponseEnvelopeSuccess bool

const (
	SchemaNewResponseEnvelopeSuccessTrue SchemaNewResponseEnvelopeSuccess = true
)

func (r SchemaNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SchemaNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SchemaListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter for enabled schemas
	ValidationEnabled param.Field[bool] `query:"validation_enabled"`
}

// URLQuery serializes [SchemaListParams]'s query parameters as `url.Values`.
func (r SchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SchemaDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SchemaDeleteResponseEnvelope struct {
	Errors   api_gateway.Message  `json:"errors,required"`
	Messages api_gateway.Message  `json:"messages,required"`
	Result   SchemaDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SchemaDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    schemaDeleteResponseEnvelopeJSON    `json:"-"`
}

// schemaDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SchemaDeleteResponseEnvelope]
type schemaDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SchemaDeleteResponseEnvelopeSuccess bool

const (
	SchemaDeleteResponseEnvelopeSuccessTrue SchemaDeleteResponseEnvelopeSuccess = true
)

func (r SchemaDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SchemaDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SchemaEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[bool] `json:"validation_enabled"`
}

func (r SchemaEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SchemaEditResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// A schema used in schema validation
	Result SchemaEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SchemaEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    schemaEditResponseEnvelopeJSON    `json:"-"`
}

// schemaEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SchemaEditResponseEnvelope]
type schemaEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SchemaEditResponseEnvelopeSuccess bool

const (
	SchemaEditResponseEnvelopeSuccessTrue SchemaEditResponseEnvelopeSuccess = true
)

func (r SchemaEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SchemaEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SchemaGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
}

// URLQuery serializes [SchemaGetParams]'s query parameters as `url.Values`.
func (r SchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SchemaGetResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// A schema used in schema validation
	Result SchemaGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SchemaGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    schemaGetResponseEnvelopeJSON    `json:"-"`
}

// schemaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SchemaGetResponseEnvelope]
type schemaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SchemaGetResponseEnvelopeSuccess bool

const (
	SchemaGetResponseEnvelopeSuccessTrue SchemaGetResponseEnvelopeSuccess = true
)

func (r SchemaGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SchemaGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
