// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiform"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// UserSchemaService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSchemaService] method instead.
type UserSchemaService struct {
	Options    []option.RequestOption
	Operations *UserSchemaOperationService
	Hosts      *UserSchemaHostService
}

// NewUserSchemaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserSchemaService(opts ...option.RequestOption) (r *UserSchemaService) {
	r = &UserSchemaService{}
	r.Options = opts
	r.Operations = NewUserSchemaOperationService(opts...)
	r.Hosts = NewUserSchemaHostService(opts...)
	return
}

// Upload a schema to a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaService) New(ctx context.Context, params UserSchemaNewParams, opts ...option.RequestOption) (res *SchemaUpload, err error) {
	var env UserSchemaNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about all schemas on a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaService) List(ctx context.Context, params UserSchemaListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[PublicSchema], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas", params.ZoneID)
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

// Retrieve information about all schemas on a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaService) ListAutoPaging(ctx context.Context, params UserSchemaListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[PublicSchema] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a schema
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaService) Delete(ctx context.Context, schemaID string, body UserSchemaDeleteParams, opts ...option.RequestOption) (res *UserSchemaDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", body.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Enable validation for a schema
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaService) Edit(ctx context.Context, schemaID string, params UserSchemaEditParams, opts ...option.RequestOption) (res *PublicSchema, err error) {
	var env UserSchemaEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", params.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about a specific schema on a zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaService) Get(ctx context.Context, schemaID string, params UserSchemaGetParams, opts ...option.RequestOption) (res *PublicSchema, err error) {
	var env UserSchemaGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s", params.ZoneID, schemaID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Message []MessageItem

type MessageItem struct {
	Code             int64             `json:"code,required"`
	Message          string            `json:"message,required"`
	DocumentationURL string            `json:"documentation_url"`
	Source           MessageItemSource `json:"source"`
	JSON             messageItemJSON   `json:"-"`
}

// messageItemJSON contains the JSON metadata for the struct [MessageItem]
type messageItemJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MessageItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemJSON) RawJSON() string {
	return r.raw
}

type MessageItemSource struct {
	Pointer string                `json:"pointer"`
	JSON    messageItemSourceJSON `json:"-"`
}

// messageItemSourceJSON contains the JSON metadata for the struct
// [MessageItemSource]
type messageItemSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MessageItemSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r messageItemSourceJSON) RawJSON() string {
	return r.raw
}

type PublicSchema struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Kind of schema
	Kind PublicSchemaKind `json:"kind,required"`
	// Name of the schema
	Name string `json:"name,required"`
	// UUID.
	SchemaID string `json:"schema_id,required"`
	// Source of the schema
	Source string `json:"source"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled bool             `json:"validation_enabled"`
	JSON              publicSchemaJSON `json:"-"`
}

// publicSchemaJSON contains the JSON metadata for the struct [PublicSchema]
type publicSchemaJSON struct {
	CreatedAt         apijson.Field
	Kind              apijson.Field
	Name              apijson.Field
	SchemaID          apijson.Field
	Source            apijson.Field
	ValidationEnabled apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PublicSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r publicSchemaJSON) RawJSON() string {
	return r.raw
}

// Kind of schema
type PublicSchemaKind string

const (
	PublicSchemaKindOpenAPIV3 PublicSchemaKind = "openapi_v3"
)

func (r PublicSchemaKind) IsKnown() bool {
	switch r {
	case PublicSchemaKindOpenAPIV3:
		return true
	}
	return false
}

type SchemaUpload struct {
	Schema        PublicSchema              `json:"schema,required"`
	UploadDetails SchemaUploadUploadDetails `json:"upload_details"`
	JSON          schemaUploadJSON          `json:"-"`
}

// schemaUploadJSON contains the JSON metadata for the struct [SchemaUpload]
type schemaUploadJSON struct {
	Schema        apijson.Field
	UploadDetails apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SchemaUpload) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaUploadJSON) RawJSON() string {
	return r.raw
}

type SchemaUploadUploadDetails struct {
	// Diagnostic warning events that occurred during processing. These events are
	// non-critical errors found within the schema.
	Warnings []SchemaUploadUploadDetailsWarning `json:"warnings"`
	JSON     schemaUploadUploadDetailsJSON      `json:"-"`
}

// schemaUploadUploadDetailsJSON contains the JSON metadata for the struct
// [SchemaUploadUploadDetails]
type schemaUploadUploadDetailsJSON struct {
	Warnings    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaUploadUploadDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaUploadUploadDetailsJSON) RawJSON() string {
	return r.raw
}

type SchemaUploadUploadDetailsWarning struct {
	// Code that identifies the event that occurred.
	Code int64 `json:"code,required"`
	// JSONPath location(s) in the schema where these events were encountered. See
	// [https://goessner.net/articles/JsonPath/](https://goessner.net/articles/JsonPath/)
	// for JSONPath specification.
	Locations []string `json:"locations"`
	// Diagnostic message that describes the event.
	Message string                               `json:"message"`
	JSON    schemaUploadUploadDetailsWarningJSON `json:"-"`
}

// schemaUploadUploadDetailsWarningJSON contains the JSON metadata for the struct
// [SchemaUploadUploadDetailsWarning]
type schemaUploadUploadDetailsWarningJSON struct {
	Code        apijson.Field
	Locations   apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaUploadUploadDetailsWarning) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaUploadUploadDetailsWarningJSON) RawJSON() string {
	return r.raw
}

type UserSchemaDeleteResponse struct {
	Errors   Message `json:"errors,required"`
	Messages Message `json:"messages,required"`
	// Whether the API call was successful.
	Success UserSchemaDeleteResponseSuccess `json:"success,required"`
	JSON    userSchemaDeleteResponseJSON    `json:"-"`
}

// userSchemaDeleteResponseJSON contains the JSON metadata for the struct
// [UserSchemaDeleteResponse]
type userSchemaDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserSchemaDeleteResponseSuccess bool

const (
	UserSchemaDeleteResponseSuccessTrue UserSchemaDeleteResponseSuccess = true
)

func (r UserSchemaDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case UserSchemaDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type UserSchemaNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Schema file bytes
	File param.Field[io.Reader] `json:"file,required" format:"binary"`
	// Kind of schema
	Kind param.Field[UserSchemaNewParamsKind] `json:"kind,required"`
	// Name of the schema
	Name param.Field[string] `json:"name"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[UserSchemaNewParamsValidationEnabled] `json:"validation_enabled"`
}

func (r UserSchemaNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Kind of schema
type UserSchemaNewParamsKind string

const (
	UserSchemaNewParamsKindOpenAPIV3 UserSchemaNewParamsKind = "openapi_v3"
)

func (r UserSchemaNewParamsKind) IsKnown() bool {
	switch r {
	case UserSchemaNewParamsKindOpenAPIV3:
		return true
	}
	return false
}

// Flag whether schema is enabled for validation.
type UserSchemaNewParamsValidationEnabled string

const (
	UserSchemaNewParamsValidationEnabledTrue  UserSchemaNewParamsValidationEnabled = "true"
	UserSchemaNewParamsValidationEnabledFalse UserSchemaNewParamsValidationEnabled = "false"
)

func (r UserSchemaNewParamsValidationEnabled) IsKnown() bool {
	switch r {
	case UserSchemaNewParamsValidationEnabledTrue, UserSchemaNewParamsValidationEnabledFalse:
		return true
	}
	return false
}

type UserSchemaNewResponseEnvelope struct {
	Errors   Message      `json:"errors,required"`
	Messages Message      `json:"messages,required"`
	Result   SchemaUpload `json:"result,required"`
	// Whether the API call was successful.
	Success UserSchemaNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    userSchemaNewResponseEnvelopeJSON    `json:"-"`
}

// userSchemaNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserSchemaNewResponseEnvelope]
type userSchemaNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserSchemaNewResponseEnvelopeSuccess bool

const (
	UserSchemaNewResponseEnvelopeSuccessTrue UserSchemaNewResponseEnvelopeSuccess = true
)

func (r UserSchemaNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserSchemaNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserSchemaListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[bool] `query:"validation_enabled"`
}

// URLQuery serializes [UserSchemaListParams]'s query parameters as `url.Values`.
func (r UserSchemaListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserSchemaDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type UserSchemaEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Flag whether schema is enabled for validation.
	ValidationEnabled param.Field[UserSchemaEditParamsValidationEnabled] `json:"validation_enabled"`
}

func (r UserSchemaEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Flag whether schema is enabled for validation.
type UserSchemaEditParamsValidationEnabled bool

const (
	UserSchemaEditParamsValidationEnabledTrue UserSchemaEditParamsValidationEnabled = true
)

func (r UserSchemaEditParamsValidationEnabled) IsKnown() bool {
	switch r {
	case UserSchemaEditParamsValidationEnabledTrue:
		return true
	}
	return false
}

type UserSchemaEditResponseEnvelope struct {
	Errors   Message      `json:"errors,required"`
	Messages Message      `json:"messages,required"`
	Result   PublicSchema `json:"result,required"`
	// Whether the API call was successful.
	Success UserSchemaEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    userSchemaEditResponseEnvelopeJSON    `json:"-"`
}

// userSchemaEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserSchemaEditResponseEnvelope]
type userSchemaEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserSchemaEditResponseEnvelopeSuccess bool

const (
	UserSchemaEditResponseEnvelopeSuccessTrue UserSchemaEditResponseEnvelopeSuccess = true
)

func (r UserSchemaEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserSchemaEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserSchemaGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Omit the source-files of schemas and only retrieve their meta-data.
	OmitSource param.Field[bool] `query:"omit_source"`
}

// URLQuery serializes [UserSchemaGetParams]'s query parameters as `url.Values`.
func (r UserSchemaGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserSchemaGetResponseEnvelope struct {
	Errors   Message      `json:"errors,required"`
	Messages Message      `json:"messages,required"`
	Result   PublicSchema `json:"result,required"`
	// Whether the API call was successful.
	Success UserSchemaGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    userSchemaGetResponseEnvelopeJSON    `json:"-"`
}

// userSchemaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserSchemaGetResponseEnvelope]
type userSchemaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserSchemaGetResponseEnvelopeSuccess bool

const (
	UserSchemaGetResponseEnvelopeSuccessTrue UserSchemaGetResponseEnvelopeSuccess = true
)

func (r UserSchemaGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserSchemaGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
