// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schema_validation

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/api_gateway"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SettingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingService] method instead.
type SettingService struct {
	Options    []option.RequestOption
	Operations *SettingOperationService
}

// NewSettingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSettingService(opts ...option.RequestOption) (r *SettingService) {
	r = &SettingService{}
	r.Options = opts
	r.Operations = NewSettingOperationService(opts...)
	return
}

// Update global schema validation settings
func (r *SettingService) Update(ctx context.Context, params SettingUpdateParams, opts ...option.RequestOption) (res *SettingUpdateResponse, err error) {
	var env SettingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit global schema validation settings
func (r *SettingService) Edit(ctx context.Context, params SettingEditParams, opts ...option.RequestOption) (res *SettingEditResponse, err error) {
	var env SettingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get global schema validation settings
func (r *SettingService) Get(ctx context.Context, query SettingGetParams, opts ...option.RequestOption) (res *SettingGetResponse, err error) {
	var env SettingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/schema_validation/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingUpdateResponse struct {
	// The default mitigation action used
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	// - `none` - skip running schema validation
	ValidationDefaultMitigationAction SettingUpdateResponseValidationDefaultMitigationAction `json:"validation_default_mitigation_action,required"`
	// When not null, this overrides global both zone level and operation level
	// mitigation actions. This can serve as a quick way to disable schema validation
	// for the whole zone.
	//
	// - `"none"` will skip running schema validation entirely for the request
	ValidationOverrideMitigationAction SettingUpdateResponseValidationOverrideMitigationAction `json:"validation_override_mitigation_action"`
	JSON                               settingUpdateResponseJSON                               `json:"-"`
}

// settingUpdateResponseJSON contains the JSON metadata for the struct
// [SettingUpdateResponse]
type settingUpdateResponseJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SettingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The default mitigation action used
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
// - `none` - skip running schema validation
type SettingUpdateResponseValidationDefaultMitigationAction string

const (
	SettingUpdateResponseValidationDefaultMitigationActionNone  SettingUpdateResponseValidationDefaultMitigationAction = "none"
	SettingUpdateResponseValidationDefaultMitigationActionLog   SettingUpdateResponseValidationDefaultMitigationAction = "log"
	SettingUpdateResponseValidationDefaultMitigationActionBlock SettingUpdateResponseValidationDefaultMitigationAction = "block"
)

func (r SettingUpdateResponseValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingUpdateResponseValidationDefaultMitigationActionNone, SettingUpdateResponseValidationDefaultMitigationActionLog, SettingUpdateResponseValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When not null, this overrides global both zone level and operation level
// mitigation actions. This can serve as a quick way to disable schema validation
// for the whole zone.
//
// - `"none"` will skip running schema validation entirely for the request
type SettingUpdateResponseValidationOverrideMitigationAction string

const (
	SettingUpdateResponseValidationOverrideMitigationActionNone SettingUpdateResponseValidationOverrideMitigationAction = "none"
)

func (r SettingUpdateResponseValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingUpdateResponseValidationOverrideMitigationActionNone:
		return true
	}
	return false
}

type SettingEditResponse struct {
	// The default mitigation action used
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	// - `none` - skip running schema validation
	ValidationDefaultMitigationAction SettingEditResponseValidationDefaultMitigationAction `json:"validation_default_mitigation_action,required"`
	// When not null, this overrides global both zone level and operation level
	// mitigation actions. This can serve as a quick way to disable schema validation
	// for the whole zone.
	//
	// - `"none"` will skip running schema validation entirely for the request
	ValidationOverrideMitigationAction SettingEditResponseValidationOverrideMitigationAction `json:"validation_override_mitigation_action"`
	JSON                               settingEditResponseJSON                               `json:"-"`
}

// settingEditResponseJSON contains the JSON metadata for the struct
// [SettingEditResponse]
type settingEditResponseJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SettingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseJSON) RawJSON() string {
	return r.raw
}

// The default mitigation action used
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
// - `none` - skip running schema validation
type SettingEditResponseValidationDefaultMitigationAction string

const (
	SettingEditResponseValidationDefaultMitigationActionNone  SettingEditResponseValidationDefaultMitigationAction = "none"
	SettingEditResponseValidationDefaultMitigationActionLog   SettingEditResponseValidationDefaultMitigationAction = "log"
	SettingEditResponseValidationDefaultMitigationActionBlock SettingEditResponseValidationDefaultMitigationAction = "block"
)

func (r SettingEditResponseValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingEditResponseValidationDefaultMitigationActionNone, SettingEditResponseValidationDefaultMitigationActionLog, SettingEditResponseValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When not null, this overrides global both zone level and operation level
// mitigation actions. This can serve as a quick way to disable schema validation
// for the whole zone.
//
// - `"none"` will skip running schema validation entirely for the request
type SettingEditResponseValidationOverrideMitigationAction string

const (
	SettingEditResponseValidationOverrideMitigationActionNone SettingEditResponseValidationOverrideMitigationAction = "none"
)

func (r SettingEditResponseValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingEditResponseValidationOverrideMitigationActionNone:
		return true
	}
	return false
}

type SettingGetResponse struct {
	// The default mitigation action used
	//
	// Mitigation actions are as follows:
	//
	// - `log` - log request when request does not conform to schema
	// - `block` - deny access to the site when request does not conform to schema
	// - `none` - skip running schema validation
	ValidationDefaultMitigationAction SettingGetResponseValidationDefaultMitigationAction `json:"validation_default_mitigation_action,required"`
	// When not null, this overrides global both zone level and operation level
	// mitigation actions. This can serve as a quick way to disable schema validation
	// for the whole zone.
	//
	// - `"none"` will skip running schema validation entirely for the request
	ValidationOverrideMitigationAction SettingGetResponseValidationOverrideMitigationAction `json:"validation_override_mitigation_action"`
	JSON                               settingGetResponseJSON                               `json:"-"`
}

// settingGetResponseJSON contains the JSON metadata for the struct
// [SettingGetResponse]
type settingGetResponseJSON struct {
	ValidationDefaultMitigationAction  apijson.Field
	ValidationOverrideMitigationAction apijson.Field
	raw                                string
	ExtraFields                        map[string]apijson.Field
}

func (r *SettingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseJSON) RawJSON() string {
	return r.raw
}

// The default mitigation action used
//
// Mitigation actions are as follows:
//
// - `log` - log request when request does not conform to schema
// - `block` - deny access to the site when request does not conform to schema
// - `none` - skip running schema validation
type SettingGetResponseValidationDefaultMitigationAction string

const (
	SettingGetResponseValidationDefaultMitigationActionNone  SettingGetResponseValidationDefaultMitigationAction = "none"
	SettingGetResponseValidationDefaultMitigationActionLog   SettingGetResponseValidationDefaultMitigationAction = "log"
	SettingGetResponseValidationDefaultMitigationActionBlock SettingGetResponseValidationDefaultMitigationAction = "block"
)

func (r SettingGetResponseValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingGetResponseValidationDefaultMitigationActionNone, SettingGetResponseValidationDefaultMitigationActionLog, SettingGetResponseValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When not null, this overrides global both zone level and operation level
// mitigation actions. This can serve as a quick way to disable schema validation
// for the whole zone.
//
// - `"none"` will skip running schema validation entirely for the request
type SettingGetResponseValidationOverrideMitigationAction string

const (
	SettingGetResponseValidationOverrideMitigationActionNone SettingGetResponseValidationOverrideMitigationAction = "none"
)

func (r SettingGetResponseValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingGetResponseValidationOverrideMitigationActionNone:
		return true
	}
	return false
}

type SettingUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The default mitigation action used Mitigation actions are as follows:
	//
	// - `"log"` - log request when request does not conform to schema
	// - `"block"` - deny access to the site when request does not conform to schema
	// - `"none"` - skip running schema validation
	ValidationDefaultMitigationAction param.Field[SettingUpdateParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action,required"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `"none"` - skip running schema validation entirely for the request
	// - `null` - clears any existing override
	ValidationOverrideMitigationAction param.Field[SettingUpdateParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r SettingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default mitigation action used Mitigation actions are as follows:
//
// - `"log"` - log request when request does not conform to schema
// - `"block"` - deny access to the site when request does not conform to schema
// - `"none"` - skip running schema validation
type SettingUpdateParamsValidationDefaultMitigationAction string

const (
	SettingUpdateParamsValidationDefaultMitigationActionNone  SettingUpdateParamsValidationDefaultMitigationAction = "none"
	SettingUpdateParamsValidationDefaultMitigationActionLog   SettingUpdateParamsValidationDefaultMitigationAction = "log"
	SettingUpdateParamsValidationDefaultMitigationActionBlock SettingUpdateParamsValidationDefaultMitigationAction = "block"
)

func (r SettingUpdateParamsValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingUpdateParamsValidationDefaultMitigationActionNone, SettingUpdateParamsValidationDefaultMitigationActionLog, SettingUpdateParamsValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `"none"` - skip running schema validation entirely for the request
// - `null` - clears any existing override
type SettingUpdateParamsValidationOverrideMitigationAction string

const (
	SettingUpdateParamsValidationOverrideMitigationActionNone SettingUpdateParamsValidationOverrideMitigationAction = "none"
)

func (r SettingUpdateParamsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingUpdateParamsValidationOverrideMitigationActionNone:
		return true
	}
	return false
}

type SettingUpdateResponseEnvelope struct {
	Errors   api_gateway.Message   `json:"errors,required"`
	Messages api_gateway.Message   `json:"messages,required"`
	Result   SettingUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SettingUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingUpdateResponseEnvelopeJSON    `json:"-"`
}

// settingUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingUpdateResponseEnvelope]
type settingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingUpdateResponseEnvelopeSuccess bool

const (
	SettingUpdateResponseEnvelopeSuccessTrue SettingUpdateResponseEnvelopeSuccess = true
)

func (r SettingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The default mitigation action used Mitigation actions are as follows:
	//
	// - `"log"` - log request when request does not conform to schema
	// - `"block"` - deny access to the site when request does not conform to schema
	// - `"none"` - skip running schema validation
	ValidationDefaultMitigationAction param.Field[SettingEditParamsValidationDefaultMitigationAction] `json:"validation_default_mitigation_action"`
	// When set, this overrides both zone level and operation level mitigation actions.
	//
	// - `"none"` - skip running schema validation entirely for the request
	// - `null` - clears any existing override
	ValidationOverrideMitigationAction param.Field[SettingEditParamsValidationOverrideMitigationAction] `json:"validation_override_mitigation_action"`
}

func (r SettingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The default mitigation action used Mitigation actions are as follows:
//
// - `"log"` - log request when request does not conform to schema
// - `"block"` - deny access to the site when request does not conform to schema
// - `"none"` - skip running schema validation
type SettingEditParamsValidationDefaultMitigationAction string

const (
	SettingEditParamsValidationDefaultMitigationActionNone  SettingEditParamsValidationDefaultMitigationAction = "none"
	SettingEditParamsValidationDefaultMitigationActionLog   SettingEditParamsValidationDefaultMitigationAction = "log"
	SettingEditParamsValidationDefaultMitigationActionBlock SettingEditParamsValidationDefaultMitigationAction = "block"
)

func (r SettingEditParamsValidationDefaultMitigationAction) IsKnown() bool {
	switch r {
	case SettingEditParamsValidationDefaultMitigationActionNone, SettingEditParamsValidationDefaultMitigationActionLog, SettingEditParamsValidationDefaultMitigationActionBlock:
		return true
	}
	return false
}

// When set, this overrides both zone level and operation level mitigation actions.
//
// - `"none"` - skip running schema validation entirely for the request
// - `null` - clears any existing override
type SettingEditParamsValidationOverrideMitigationAction string

const (
	SettingEditParamsValidationOverrideMitigationActionNone SettingEditParamsValidationOverrideMitigationAction = "none"
)

func (r SettingEditParamsValidationOverrideMitigationAction) IsKnown() bool {
	switch r {
	case SettingEditParamsValidationOverrideMitigationActionNone:
		return true
	}
	return false
}

type SettingEditResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	Result   SettingEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success SettingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingEditResponseEnvelopeJSON    `json:"-"`
}

// settingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingEditResponseEnvelope]
type settingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingEditResponseEnvelopeSuccess bool

const (
	SettingEditResponseEnvelopeSuccessTrue SettingEditResponseEnvelopeSuccess = true
)

func (r SettingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingGetResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	Result   SettingGetResponse  `json:"result,required"`
	// Whether the API call was successful.
	Success SettingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    settingGetResponseEnvelopeJSON    `json:"-"`
}

// settingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingGetResponseEnvelope]
type settingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingGetResponseEnvelopeSuccess bool

const (
	SettingGetResponseEnvelopeSuccessTrue SettingGetResponseEnvelopeSuccess = true
)

func (r SettingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
