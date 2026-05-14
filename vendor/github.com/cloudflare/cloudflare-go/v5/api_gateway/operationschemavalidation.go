// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// OperationSchemaValidationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationSchemaValidationService] method instead.
type OperationSchemaValidationService struct {
	Options []option.RequestOption
}

// NewOperationSchemaValidationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewOperationSchemaValidationService(opts ...option.RequestOption) (r *OperationSchemaValidationService) {
	r = &OperationSchemaValidationService{}
	r.Options = opts
	return
}

// Updates operation-level schema validation settings on the zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *OperationSchemaValidationService) Update(ctx context.Context, operationID string, params OperationSchemaValidationUpdateParams, opts ...option.RequestOption) (res *OperationSchemaValidationUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/schema_validation", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Updates multiple operation-level schema validation settings on the zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *OperationSchemaValidationService) Edit(ctx context.Context, params OperationSchemaValidationEditParams, opts ...option.RequestOption) (res *SettingsMultipleRequest, err error) {
	var env OperationSchemaValidationEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/schema_validation", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves operation-level schema validation settings on the zone
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *OperationSchemaValidationService) Get(ctx context.Context, operationID string, query OperationSchemaValidationGetParams, opts ...option.RequestOption) (res *OperationSchemaValidationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s/schema_validation", query.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SettingsMultipleRequest map[string]SettingsMultipleRequestItem

// Operation ID to mitigation action mappings
type SettingsMultipleRequestItem struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction SettingsMultipleRequestItemMitigationAction `json:"mitigation_action,nullable"`
	JSON             settingsMultipleRequestItemJSON             `json:"-"`
}

// settingsMultipleRequestItemJSON contains the JSON metadata for the struct
// [SettingsMultipleRequestItem]
type settingsMultipleRequestItemJSON struct {
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingsMultipleRequestItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingsMultipleRequestItemJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type SettingsMultipleRequestItemMitigationAction string

const (
	SettingsMultipleRequestItemMitigationActionLog   SettingsMultipleRequestItemMitigationAction = "log"
	SettingsMultipleRequestItemMitigationActionBlock SettingsMultipleRequestItemMitigationAction = "block"
	SettingsMultipleRequestItemMitigationActionNone  SettingsMultipleRequestItemMitigationAction = "none"
)

func (r SettingsMultipleRequestItemMitigationAction) IsKnown() bool {
	switch r {
	case SettingsMultipleRequestItemMitigationActionLog, SettingsMultipleRequestItemMitigationActionBlock, SettingsMultipleRequestItemMitigationActionNone:
		return true
	}
	return false
}

type SettingsMultipleRequestParam map[string]SettingsMultipleRequestItemParam

// Operation ID to mitigation action mappings
type SettingsMultipleRequestItemParam struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction param.Field[SettingsMultipleRequestItemMitigationAction] `json:"mitigation_action"`
}

func (r SettingsMultipleRequestItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OperationSchemaValidationUpdateResponse struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction OperationSchemaValidationUpdateResponseMitigationAction `json:"mitigation_action,nullable"`
	// UUID.
	OperationID string                                      `json:"operation_id"`
	JSON        operationSchemaValidationUpdateResponseJSON `json:"-"`
}

// operationSchemaValidationUpdateResponseJSON contains the JSON metadata for the
// struct [OperationSchemaValidationUpdateResponse]
type operationSchemaValidationUpdateResponseJSON struct {
	MitigationAction apijson.Field
	OperationID      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationSchemaValidationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationSchemaValidationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type OperationSchemaValidationUpdateResponseMitigationAction string

const (
	OperationSchemaValidationUpdateResponseMitigationActionLog   OperationSchemaValidationUpdateResponseMitigationAction = "log"
	OperationSchemaValidationUpdateResponseMitigationActionBlock OperationSchemaValidationUpdateResponseMitigationAction = "block"
	OperationSchemaValidationUpdateResponseMitigationActionNone  OperationSchemaValidationUpdateResponseMitigationAction = "none"
)

func (r OperationSchemaValidationUpdateResponseMitigationAction) IsKnown() bool {
	switch r {
	case OperationSchemaValidationUpdateResponseMitigationActionLog, OperationSchemaValidationUpdateResponseMitigationActionBlock, OperationSchemaValidationUpdateResponseMitigationActionNone:
		return true
	}
	return false
}

type OperationSchemaValidationGetResponse struct {
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction OperationSchemaValidationGetResponseMitigationAction `json:"mitigation_action,nullable"`
	// UUID.
	OperationID string                                   `json:"operation_id"`
	JSON        operationSchemaValidationGetResponseJSON `json:"-"`
}

// operationSchemaValidationGetResponseJSON contains the JSON metadata for the
// struct [OperationSchemaValidationGetResponse]
type operationSchemaValidationGetResponseJSON struct {
	MitigationAction apijson.Field
	OperationID      apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationSchemaValidationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationSchemaValidationGetResponseJSON) RawJSON() string {
	return r.raw
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type OperationSchemaValidationGetResponseMitigationAction string

const (
	OperationSchemaValidationGetResponseMitigationActionLog   OperationSchemaValidationGetResponseMitigationAction = "log"
	OperationSchemaValidationGetResponseMitigationActionBlock OperationSchemaValidationGetResponseMitigationAction = "block"
	OperationSchemaValidationGetResponseMitigationActionNone  OperationSchemaValidationGetResponseMitigationAction = "none"
)

func (r OperationSchemaValidationGetResponseMitigationAction) IsKnown() bool {
	switch r {
	case OperationSchemaValidationGetResponseMitigationActionLog, OperationSchemaValidationGetResponseMitigationActionBlock, OperationSchemaValidationGetResponseMitigationActionNone:
		return true
	}
	return false
}

type OperationSchemaValidationUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// When set, this applies a mitigation action to this operation
	//
	//   - `log` log request when request does not conform to schema for this operation
	//   - `block` deny access to the site when request does not conform to schema for
	//     this operation
	//   - `none` will skip mitigation for this operation
	//   - `null` indicates that no operation level mitigation is in place, see Zone
	//     Level Schema Validation Settings for mitigation action that will be applied
	MitigationAction param.Field[OperationSchemaValidationUpdateParamsMitigationAction] `json:"mitigation_action"`
}

func (r OperationSchemaValidationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// When set, this applies a mitigation action to this operation
//
//   - `log` log request when request does not conform to schema for this operation
//   - `block` deny access to the site when request does not conform to schema for
//     this operation
//   - `none` will skip mitigation for this operation
//   - `null` indicates that no operation level mitigation is in place, see Zone
//     Level Schema Validation Settings for mitigation action that will be applied
type OperationSchemaValidationUpdateParamsMitigationAction string

const (
	OperationSchemaValidationUpdateParamsMitigationActionLog   OperationSchemaValidationUpdateParamsMitigationAction = "log"
	OperationSchemaValidationUpdateParamsMitigationActionBlock OperationSchemaValidationUpdateParamsMitigationAction = "block"
	OperationSchemaValidationUpdateParamsMitigationActionNone  OperationSchemaValidationUpdateParamsMitigationAction = "none"
)

func (r OperationSchemaValidationUpdateParamsMitigationAction) IsKnown() bool {
	switch r {
	case OperationSchemaValidationUpdateParamsMitigationActionLog, OperationSchemaValidationUpdateParamsMitigationActionBlock, OperationSchemaValidationUpdateParamsMitigationActionNone:
		return true
	}
	return false
}

type OperationSchemaValidationEditParams struct {
	// Identifier.
	ZoneID                  param.Field[string]          `path:"zone_id,required"`
	SettingsMultipleRequest SettingsMultipleRequestParam `json:"settings_multiple_request,required"`
}

func (r OperationSchemaValidationEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SettingsMultipleRequest)
}

type OperationSchemaValidationEditResponseEnvelope struct {
	Errors   Message                 `json:"errors,required"`
	Messages Message                 `json:"messages,required"`
	Result   SettingsMultipleRequest `json:"result,required"`
	// Whether the API call was successful.
	Success OperationSchemaValidationEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    operationSchemaValidationEditResponseEnvelopeJSON    `json:"-"`
}

// operationSchemaValidationEditResponseEnvelopeJSON contains the JSON metadata for
// the struct [OperationSchemaValidationEditResponseEnvelope]
type operationSchemaValidationEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationSchemaValidationEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationSchemaValidationEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationSchemaValidationEditResponseEnvelopeSuccess bool

const (
	OperationSchemaValidationEditResponseEnvelopeSuccessTrue OperationSchemaValidationEditResponseEnvelopeSuccess = true
)

func (r OperationSchemaValidationEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationSchemaValidationEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OperationSchemaValidationGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
