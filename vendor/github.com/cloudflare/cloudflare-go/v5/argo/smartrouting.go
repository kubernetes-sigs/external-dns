// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// SmartRoutingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmartRoutingService] method instead.
type SmartRoutingService struct {
	Options []option.RequestOption
}

// NewSmartRoutingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSmartRoutingService(opts ...option.RequestOption) (r *SmartRoutingService) {
	r = &SmartRoutingService{}
	r.Options = opts
	return
}

// Configures the value of the Argo Smart Routing enablement setting.
func (r *SmartRoutingService) Edit(ctx context.Context, params SmartRoutingEditParams, opts ...option.RequestOption) (res *SmartRoutingEditResponse, err error) {
	var env SmartRoutingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/smart_routing", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the value of Argo Smart Routing enablement setting.
func (r *SmartRoutingService) Get(ctx context.Context, query SmartRoutingGetParams, opts ...option.RequestOption) (res *SmartRoutingGetResponse, err error) {
	var env SmartRoutingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/smart_routing", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SmartRoutingEditResponse struct {
	// Specifies the identifier of the Argo Smart Routing setting.
	ID string `json:"id,required"`
	// Specifies if the setting is editable.
	Editable bool `json:"editable,required"`
	// Specifies the enablement value of Argo Smart Routing.
	Value SmartRoutingEditResponseValue `json:"value,required"`
	// Specifies the time when the setting was last modified.
	ModifiedOn time.Time                    `json:"modified_on" format:"date-time"`
	JSON       smartRoutingEditResponseJSON `json:"-"`
}

// smartRoutingEditResponseJSON contains the JSON metadata for the struct
// [SmartRoutingEditResponse]
type smartRoutingEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartRoutingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartRoutingEditResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the enablement value of Argo Smart Routing.
type SmartRoutingEditResponseValue string

const (
	SmartRoutingEditResponseValueOn  SmartRoutingEditResponseValue = "on"
	SmartRoutingEditResponseValueOff SmartRoutingEditResponseValue = "off"
)

func (r SmartRoutingEditResponseValue) IsKnown() bool {
	switch r {
	case SmartRoutingEditResponseValueOn, SmartRoutingEditResponseValueOff:
		return true
	}
	return false
}

type SmartRoutingGetResponse struct {
	// Specifies the identifier of the Argo Smart Routing setting.
	ID string `json:"id,required"`
	// Specifies if the setting is editable.
	Editable bool `json:"editable,required"`
	// Specifies the enablement value of Argo Smart Routing.
	Value SmartRoutingGetResponseValue `json:"value,required"`
	// Specifies the time when the setting was last modified.
	ModifiedOn time.Time                   `json:"modified_on" format:"date-time"`
	JSON       smartRoutingGetResponseJSON `json:"-"`
}

// smartRoutingGetResponseJSON contains the JSON metadata for the struct
// [SmartRoutingGetResponse]
type smartRoutingGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartRoutingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartRoutingGetResponseJSON) RawJSON() string {
	return r.raw
}

// Specifies the enablement value of Argo Smart Routing.
type SmartRoutingGetResponseValue string

const (
	SmartRoutingGetResponseValueOn  SmartRoutingGetResponseValue = "on"
	SmartRoutingGetResponseValueOff SmartRoutingGetResponseValue = "off"
)

func (r SmartRoutingGetResponseValue) IsKnown() bool {
	switch r {
	case SmartRoutingGetResponseValueOn, SmartRoutingGetResponseValueOff:
		return true
	}
	return false
}

type SmartRoutingEditParams struct {
	// Specifies the zone associated with the API call.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Specifies the enablement value of Argo Smart Routing.
	Value param.Field[SmartRoutingEditParamsValue] `json:"value,required"`
}

func (r SmartRoutingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specifies the enablement value of Argo Smart Routing.
type SmartRoutingEditParamsValue string

const (
	SmartRoutingEditParamsValueOn  SmartRoutingEditParamsValue = "on"
	SmartRoutingEditParamsValueOff SmartRoutingEditParamsValue = "off"
)

func (r SmartRoutingEditParamsValue) IsKnown() bool {
	switch r {
	case SmartRoutingEditParamsValueOn, SmartRoutingEditParamsValueOff:
		return true
	}
	return false
}

type SmartRoutingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   SmartRoutingEditResponse `json:"result,required"`
	// Describes a successful API response.
	Success SmartRoutingEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartRoutingEditResponseEnvelopeJSON    `json:"-"`
}

// smartRoutingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SmartRoutingEditResponseEnvelope]
type smartRoutingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartRoutingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartRoutingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Describes a successful API response.
type SmartRoutingEditResponseEnvelopeSuccess bool

const (
	SmartRoutingEditResponseEnvelopeSuccessTrue SmartRoutingEditResponseEnvelopeSuccess = true
)

func (r SmartRoutingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartRoutingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartRoutingGetParams struct {
	// Specifies the zone associated with the API call.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartRoutingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo   `json:"errors,required"`
	Messages []shared.ResponseInfo   `json:"messages,required"`
	Result   SmartRoutingGetResponse `json:"result,required"`
	// Describes a successful API response.
	Success SmartRoutingGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    smartRoutingGetResponseEnvelopeJSON    `json:"-"`
}

// smartRoutingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SmartRoutingGetResponseEnvelope]
type smartRoutingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartRoutingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartRoutingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Describes a successful API response.
type SmartRoutingGetResponseEnvelopeSuccess bool

const (
	SmartRoutingGetResponseEnvelopeSuccessTrue SmartRoutingGetResponseEnvelopeSuccess = true
)

func (r SmartRoutingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartRoutingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
