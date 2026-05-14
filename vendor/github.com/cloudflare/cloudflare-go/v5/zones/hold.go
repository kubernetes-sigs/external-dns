// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// HoldService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHoldService] method instead.
type HoldService struct {
	Options []option.RequestOption
}

// NewHoldService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHoldService(opts ...option.RequestOption) (r *HoldService) {
	r = &HoldService{}
	r.Options = opts
	return
}

// Enforce a zone hold on the zone, blocking the creation and activation of zones
// with this zone's hostname.
func (r *HoldService) New(ctx context.Context, params HoldNewParams, opts ...option.RequestOption) (res *ZoneHold, err error) {
	var env HoldNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Stop enforcement of a zone hold on the zone, permanently or temporarily,
// allowing the creation and activation of zones with this zone's hostname.
func (r *HoldService) Delete(ctx context.Context, params HoldDeleteParams, opts ...option.RequestOption) (res *ZoneHold, err error) {
	var env HoldDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update the `hold_after` and/or `include_subdomains` values on an existing zone
// hold. The hold is enabled if the `hold_after` date-time value is in the past.
func (r *HoldService) Edit(ctx context.Context, params HoldEditParams, opts ...option.RequestOption) (res *ZoneHold, err error) {
	var env HoldEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve whether the zone is subject to a zone hold, and metadata about the
// hold.
func (r *HoldService) Get(ctx context.Context, query HoldGetParams, opts ...option.RequestOption) (res *ZoneHold, err error) {
	var env HoldGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/hold", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneHold struct {
	Hold              bool         `json:"hold"`
	HoldAfter         string       `json:"hold_after"`
	IncludeSubdomains string       `json:"include_subdomains"`
	JSON              zoneHoldJSON `json:"-"`
}

// zoneHoldJSON contains the JSON metadata for the struct [ZoneHold]
type zoneHoldJSON struct {
	Hold              apijson.Field
	HoldAfter         apijson.Field
	IncludeSubdomains apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZoneHold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneHoldJSON) RawJSON() string {
	return r.raw
}

type HoldNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If provided, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `query:"include_subdomains"`
}

// URLQuery serializes [HoldNewParams]'s query parameters as `url.Values`.
func (r HoldNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HoldNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZoneHold              `json:"result,required"`
	// Whether the API call was successful
	Success HoldNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    holdNewResponseEnvelopeJSON    `json:"-"`
}

// holdNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldNewResponseEnvelope]
type holdNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HoldNewResponseEnvelopeSuccess bool

const (
	HoldNewResponseEnvelopeSuccessTrue HoldNewResponseEnvelopeSuccess = true
)

func (r HoldNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HoldNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HoldDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If `hold_after` is provided, the hold will be temporarily disabled, then
	// automatically re-enabled by the system at the time specified in this
	// RFC3339-formatted timestamp. Otherwise, the hold will be disabled indefinitely.
	HoldAfter param.Field[string] `query:"hold_after"`
}

// URLQuery serializes [HoldDeleteParams]'s query parameters as `url.Values`.
func (r HoldDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HoldDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZoneHold              `json:"result,required"`
	// Whether the API call was successful
	Success HoldDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    holdDeleteResponseEnvelopeJSON    `json:"-"`
}

// holdDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldDeleteResponseEnvelope]
type holdDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HoldDeleteResponseEnvelopeSuccess bool

const (
	HoldDeleteResponseEnvelopeSuccessTrue HoldDeleteResponseEnvelopeSuccess = true
)

func (r HoldDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HoldDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HoldEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// If `hold_after` is provided and future-dated, the hold will be temporarily
	// disabled, then automatically re-enabled by the system at the time specified in
	// this RFC3339-formatted timestamp. A past-dated `hold_after` value will have no
	// effect on an existing, enabled hold. Providing an empty string will set its
	// value to the current time.
	HoldAfter param.Field[string] `json:"hold_after"`
	// If `true`, the zone hold will extend to block any subdomain of the given zone,
	// as well as SSL4SaaS Custom Hostnames. For example, a zone hold on a zone with
	// the hostname 'example.com' and include_subdomains=true will block 'example.com',
	// 'staging.example.com', 'api.staging.example.com', etc.
	IncludeSubdomains param.Field[bool] `json:"include_subdomains"`
}

func (r HoldEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HoldEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZoneHold              `json:"result,required"`
	// Whether the API call was successful
	Success HoldEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    holdEditResponseEnvelopeJSON    `json:"-"`
}

// holdEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldEditResponseEnvelope]
type holdEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HoldEditResponseEnvelopeSuccess bool

const (
	HoldEditResponseEnvelopeSuccessTrue HoldEditResponseEnvelopeSuccess = true
)

func (r HoldEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HoldEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HoldGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HoldGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   ZoneHold              `json:"result,required"`
	// Whether the API call was successful
	Success HoldGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    holdGetResponseEnvelopeJSON    `json:"-"`
}

// holdGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HoldGetResponseEnvelope]
type holdGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HoldGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r holdGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HoldGetResponseEnvelopeSuccess bool

const (
	HoldGetResponseEnvelopeSuccessTrue HoldGetResponseEnvelopeSuccess = true
)

func (r HoldGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HoldGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
