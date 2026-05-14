// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/api_gateway"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// SettingOperationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingOperationService] method instead.
type SettingOperationService struct {
	Options []option.RequestOption
}

// NewSettingOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSettingOperationService(opts ...option.RequestOption) (r *SettingOperationService) {
	r = &SettingOperationService{}
	r.Options = opts
	return
}

// Update per-operation schema validation setting
func (r *SettingOperationService) Update(ctx context.Context, operationID string, params SettingOperationUpdateParams, opts ...option.RequestOption) (res *SettingOperationUpdateResponse, err error) {
	var env SettingOperationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings/operations/%s", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List per-operation schema validation settings
func (r *SettingOperationService) List(ctx context.Context, params SettingOperationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[SettingOperationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings/operations", params.ZoneID)
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

// List per-operation schema validation settings
func (r *SettingOperationService) ListAutoPaging(ctx context.Context, params SettingOperationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[SettingOperationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete per-operation schema validation setting
func (r *SettingOperationService) Delete(ctx context.Context, operationID string, body SettingOperationDeleteParams, opts ...option.RequestOption) (res *SettingOperationDeleteResponse, err error) {
	var env SettingOperationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings/operations/%s", body.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Bulk edit per-operation schema validation settings
func (r *SettingOperationService) BulkEdit(ctx context.Context, params SettingOperationBulkEditParams, opts ...option.RequestOption) (res *SettingOperationBulkEditResponse, err error) {
	var env SettingOperationBulkEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings/operations", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get per-operation schema validation setting
func (r *SettingOperationService) Get(ctx context.Context, operationID string, query SettingOperationGetParams, opts ...option.RequestOption) (res *SettingOperationGetResponse, err error) {
	var env SettingOperationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings/operations/%s", query.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingOperationUpdateResponse struct {
	// When set, this applies a mitigation action to this operation which supersedes a
	// global schema validation setting just for this operation
	//
	//   - `"log"` - log request when request does not conform to schema for this
	//     operation
	//   - `"block"` - deny access to the site when request does not conform to schema
	//     for this operation
	//   - `"none"` - will skip mitigation for this operation
	MitigationAction SettingOperationUpdateResponseMitigationAction `json:"mitigation_action,required"`
	// UUID.
	OperationID string                             `json:"operation_id,required"`
	JSON        settingOperationUpdateResponseJSON `json:"-"`
}

// settingOperationUpdateResponseJSON contains the JSON metadata for the struct
// [SettingOperationUpdateResponse]
type settingOperationUpdateResponseJSON struct {
	MitigationAction apijson.Field
	OperationID      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingOperationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation which supersedes a
// global schema validation setting just for this operation
//
//   - `"log"` - log request when request does not conform to schema for this
//     operation
//   - `"block"` - deny access to the site when request does not conform to schema
//     for this operation
//   - `"none"` - will skip mitigation for this operation
type SettingOperationUpdateResponseMitigationAction string

const (
	SettingOperationUpdateResponseMitigationActionLog   SettingOperationUpdateResponseMitigationAction = "log"
	SettingOperationUpdateResponseMitigationActionBlock SettingOperationUpdateResponseMitigationAction = "block"
	SettingOperationUpdateResponseMitigationActionNone  SettingOperationUpdateResponseMitigationAction = "none"
)

func (r SettingOperationUpdateResponseMitigationAction) IsKnown() bool {
	switch r {
	case SettingOperationUpdateResponseMitigationActionLog, SettingOperationUpdateResponseMitigationActionBlock, SettingOperationUpdateResponseMitigationActionNone:
		return true
	}
	return false
}

type SettingOperationListResponse struct {
	// When set, this applies a mitigation action to this operation which supersedes a
	// global schema validation setting just for this operation
	//
	//   - `"log"` - log request when request does not conform to schema for this
	//     operation
	//   - `"block"` - deny access to the site when request does not conform to schema
	//     for this operation
	//   - `"none"` - will skip mitigation for this operation
	MitigationAction SettingOperationListResponseMitigationAction `json:"mitigation_action,required"`
	// UUID.
	OperationID string                           `json:"operation_id,required"`
	JSON        settingOperationListResponseJSON `json:"-"`
}

// settingOperationListResponseJSON contains the JSON metadata for the struct
// [SettingOperationListResponse]
type settingOperationListResponseJSON struct {
	MitigationAction apijson.Field
	OperationID      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationListResponseJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation which supersedes a
// global schema validation setting just for this operation
//
//   - `"log"` - log request when request does not conform to schema for this
//     operation
//   - `"block"` - deny access to the site when request does not conform to schema
//     for this operation
//   - `"none"` - will skip mitigation for this operation
type SettingOperationListResponseMitigationAction string

const (
	SettingOperationListResponseMitigationActionLog   SettingOperationListResponseMitigationAction = "log"
	SettingOperationListResponseMitigationActionBlock SettingOperationListResponseMitigationAction = "block"
	SettingOperationListResponseMitigationActionNone  SettingOperationListResponseMitigationAction = "none"
)

func (r SettingOperationListResponseMitigationAction) IsKnown() bool {
	switch r {
	case SettingOperationListResponseMitigationActionLog, SettingOperationListResponseMitigationActionBlock, SettingOperationListResponseMitigationActionNone:
		return true
	}
	return false
}

type SettingOperationDeleteResponse struct {
	// UUID.
	OperationID string                             `json:"operation_id"`
	JSON        settingOperationDeleteResponseJSON `json:"-"`
}

// settingOperationDeleteResponseJSON contains the JSON metadata for the struct
// [SettingOperationDeleteResponse]
type settingOperationDeleteResponseJSON struct {
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOperationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SettingOperationBulkEditResponse map[string]SettingOperationBulkEditResponseItem

type SettingOperationBulkEditResponseItem struct {
	// When set, this applies a mitigation action to this operation which supersedes a
	// global schema validation setting just for this operation
	//
	//   - `"log"` - log request when request does not conform to schema for this
	//     operation
	//   - `"block"` - deny access to the site when request does not conform to schema
	//     for this operation
	//   - `"none"` - will skip mitigation for this operation
	MitigationAction SettingOperationBulkEditResponseItemMitigationAction `json:"mitigation_action,required"`
	// UUID.
	OperationID string                                   `json:"operation_id,required"`
	JSON        settingOperationBulkEditResponseItemJSON `json:"-"`
}

// settingOperationBulkEditResponseItemJSON contains the JSON metadata for the
// struct [SettingOperationBulkEditResponseItem]
type settingOperationBulkEditResponseItemJSON struct {
	MitigationAction apijson.Field
	OperationID      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingOperationBulkEditResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationBulkEditResponseItemJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation which supersedes a
// global schema validation setting just for this operation
//
//   - `"log"` - log request when request does not conform to schema for this
//     operation
//   - `"block"` - deny access to the site when request does not conform to schema
//     for this operation
//   - `"none"` - will skip mitigation for this operation
type SettingOperationBulkEditResponseItemMitigationAction string

const (
	SettingOperationBulkEditResponseItemMitigationActionLog   SettingOperationBulkEditResponseItemMitigationAction = "log"
	SettingOperationBulkEditResponseItemMitigationActionBlock SettingOperationBulkEditResponseItemMitigationAction = "block"
	SettingOperationBulkEditResponseItemMitigationActionNone  SettingOperationBulkEditResponseItemMitigationAction = "none"
)

func (r SettingOperationBulkEditResponseItemMitigationAction) IsKnown() bool {
	switch r {
	case SettingOperationBulkEditResponseItemMitigationActionLog, SettingOperationBulkEditResponseItemMitigationActionBlock, SettingOperationBulkEditResponseItemMitigationActionNone:
		return true
	}
	return false
}

type SettingOperationGetResponse struct {
	// When set, this applies a mitigation action to this operation which supersedes a
	// global schema validation setting just for this operation
	//
	//   - `"log"` - log request when request does not conform to schema for this
	//     operation
	//   - `"block"` - deny access to the site when request does not conform to schema
	//     for this operation
	//   - `"none"` - will skip mitigation for this operation
	MitigationAction SettingOperationGetResponseMitigationAction `json:"mitigation_action,required"`
	// UUID.
	OperationID string                          `json:"operation_id,required"`
	JSON        settingOperationGetResponseJSON `json:"-"`
}

// settingOperationGetResponseJSON contains the JSON metadata for the struct
// [SettingOperationGetResponse]
type settingOperationGetResponseJSON struct {
	MitigationAction apijson.Field
	OperationID      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingOperationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationGetResponseJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation which supersedes a
// global schema validation setting just for this operation
//
//   - `"log"` - log request when request does not conform to schema for this
//     operation
//   - `"block"` - deny access to the site when request does not conform to schema
//     for this operation
//   - `"none"` - will skip mitigation for this operation
type SettingOperationGetResponseMitigationAction string

const (
	SettingOperationGetResponseMitigationActionLog   SettingOperationGetResponseMitigationAction = "log"
	SettingOperationGetResponseMitigationActionBlock SettingOperationGetResponseMitigationAction = "block"
	SettingOperationGetResponseMitigationActionNone  SettingOperationGetResponseMitigationAction = "none"
)

func (r SettingOperationGetResponseMitigationAction) IsKnown() bool {
	switch r {
	case SettingOperationGetResponseMitigationActionLog, SettingOperationGetResponseMitigationActionBlock, SettingOperationGetResponseMitigationActionNone:
		return true
	}
	return false
}

type SettingOperationUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// When set, this applies a mitigation action to this operation
	//
	//   - `"log"` - log request when request does not conform to schema for this
	//     operation
	//   - `"block"` - deny access to the site when request does not conform to schema
	//     for this operation
	//   - `"none"` - will skip mitigation for this operation
	//   - `null` - clears any mitigation action
	MitigationAction param.Field[SettingOperationUpdateParamsMitigationAction] `json:"mitigation_action,required"`
}

func (r SettingOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set, this applies a mitigation action to this operation
//
//   - `"log"` - log request when request does not conform to schema for this
//     operation
//   - `"block"` - deny access to the site when request does not conform to schema
//     for this operation
//   - `"none"` - will skip mitigation for this operation
//   - `null` - clears any mitigation action
type SettingOperationUpdateParamsMitigationAction string

const (
	SettingOperationUpdateParamsMitigationActionLog   SettingOperationUpdateParamsMitigationAction = "log"
	SettingOperationUpdateParamsMitigationActionBlock SettingOperationUpdateParamsMitigationAction = "block"
	SettingOperationUpdateParamsMitigationActionNone  SettingOperationUpdateParamsMitigationAction = "none"
)

func (r SettingOperationUpdateParamsMitigationAction) IsKnown() bool {
	switch r {
	case SettingOperationUpdateParamsMitigationActionLog, SettingOperationUpdateParamsMitigationActionBlock, SettingOperationUpdateParamsMitigationActionNone:
		return true
	}
	return false
}

type SettingOperationUpdateResponseEnvelope struct {
	Errors   api_gateway.Message            `json:"errors,required"`
	Messages api_gateway.Message            `json:"messages,required"`
	Result   SettingOperationUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SettingOperationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingOperationUpdateResponseEnvelopeJSON    `json:"-"`
}

// settingOperationUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOperationUpdateResponseEnvelope]
type settingOperationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOperationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingOperationUpdateResponseEnvelopeSuccess bool

const (
	SettingOperationUpdateResponseEnvelopeSuccessTrue SettingOperationUpdateResponseEnvelopeSuccess = true
)

func (r SettingOperationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingOperationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingOperationListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [SettingOperationListParams]'s query parameters as
// `url.Values`.
func (r SettingOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type SettingOperationDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOperationDeleteResponseEnvelope struct {
	Errors   api_gateway.Message            `json:"errors,required"`
	Messages api_gateway.Message            `json:"messages,required"`
	Result   SettingOperationDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SettingOperationDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingOperationDeleteResponseEnvelopeJSON    `json:"-"`
}

// settingOperationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOperationDeleteResponseEnvelope]
type settingOperationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOperationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingOperationDeleteResponseEnvelopeSuccess bool

const (
	SettingOperationDeleteResponseEnvelopeSuccessTrue SettingOperationDeleteResponseEnvelopeSuccess = true
)

func (r SettingOperationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingOperationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingOperationBulkEditParams struct {
	// Identifier.
	ZoneID param.Field[string]                           `path:"zone_id,required"`
	Body   map[string]SettingOperationBulkEditParamsBody `json:"body,required"`
}

func (r SettingOperationBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// Operation ID to mitigation action mappings
type SettingOperationBulkEditParamsBody struct {
	// Mitigation actions are as follows:
	//
	//   - `log` - log request when request does not conform to schema _ `block` - deny
	//     access to the site when request does not conform to schema _ `none` - skip
	//     running schema validation \* null - clears any existing per-operation setting
	MitigationAction param.Field[SettingOperationBulkEditParamsBodyMitigationAction] `json:"mitigation_action"`
}

func (r SettingOperationBulkEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mitigation actions are as follows:
//
//   - `log` - log request when request does not conform to schema _ `block` - deny
//     access to the site when request does not conform to schema _ `none` - skip
//     running schema validation \* null - clears any existing per-operation setting
type SettingOperationBulkEditParamsBodyMitigationAction string

const (
	SettingOperationBulkEditParamsBodyMitigationActionNone  SettingOperationBulkEditParamsBodyMitigationAction = "none"
	SettingOperationBulkEditParamsBodyMitigationActionLog   SettingOperationBulkEditParamsBodyMitigationAction = "log"
	SettingOperationBulkEditParamsBodyMitigationActionBlock SettingOperationBulkEditParamsBodyMitigationAction = "block"
)

func (r SettingOperationBulkEditParamsBodyMitigationAction) IsKnown() bool {
	switch r {
	case SettingOperationBulkEditParamsBodyMitigationActionNone, SettingOperationBulkEditParamsBodyMitigationActionLog, SettingOperationBulkEditParamsBodyMitigationActionBlock:
		return true
	}
	return false
}

type SettingOperationBulkEditResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// Operation ID to per operation setting mapping
	Result SettingOperationBulkEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SettingOperationBulkEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingOperationBulkEditResponseEnvelopeJSON    `json:"-"`
}

// settingOperationBulkEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOperationBulkEditResponseEnvelope]
type settingOperationBulkEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOperationBulkEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationBulkEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingOperationBulkEditResponseEnvelopeSuccess bool

const (
	SettingOperationBulkEditResponseEnvelopeSuccessTrue SettingOperationBulkEditResponseEnvelopeSuccess = true
)

func (r SettingOperationBulkEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingOperationBulkEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingOperationGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingOperationGetResponseEnvelope struct {
	Errors   api_gateway.Message         `json:"errors,required"`
	Messages api_gateway.Message         `json:"messages,required"`
	Result   SettingOperationGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SettingOperationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingOperationGetResponseEnvelopeJSON    `json:"-"`
}

// settingOperationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SettingOperationGetResponseEnvelope]
type settingOperationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingOperationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingOperationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingOperationGetResponseEnvelopeSuccess bool

const (
	SettingOperationGetResponseEnvelopeSuccessTrue SettingOperationGetResponseEnvelopeSuccess = true
)

func (r SettingOperationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingOperationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
