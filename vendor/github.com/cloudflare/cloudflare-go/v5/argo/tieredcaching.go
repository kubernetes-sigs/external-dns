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

// TieredCachingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTieredCachingService] method instead.
type TieredCachingService struct {
	Options []option.RequestOption
}

// NewTieredCachingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTieredCachingService(opts ...option.RequestOption) (r *TieredCachingService) {
	r = &TieredCachingService{}
	r.Options = opts
	return
}

// Tiered Cache works by dividing Cloudflare's data centers into a hierarchy of
// lower-tiers and upper-tiers. If content is not cached in lower-tier data centers
// (generally the ones closest to a visitor), the lower-tier must ask an upper-tier
// to see if it has the content. If the upper-tier does not have the content, only
// the upper-tier can ask the origin for content. This practice improves bandwidth
// efficiency by limiting the number of data centers that can ask the origin for
// content, which reduces origin load and makes websites more cost-effective to
// operate. Additionally, Tiered Cache concentrates connections to origin servers
// so they come from a small number of data centers rather than the full set of
// network locations. This results in fewer open connections using server
// resources.
func (r *TieredCachingService) Edit(ctx context.Context, params TieredCachingEditParams, opts ...option.RequestOption) (res *TieredCachingEditResponse, err error) {
	var env TieredCachingEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Tiered Cache works by dividing Cloudflare's data centers into a hierarchy of
// lower-tiers and upper-tiers. If content is not cached in lower-tier data centers
// (generally the ones closest to a visitor), the lower-tier must ask an upper-tier
// to see if it has the content. If the upper-tier does not have the content, only
// the upper-tier can ask the origin for content. This practice improves bandwidth
// efficiency by limiting the number of data centers that can ask the origin for
// content, which reduces origin load and makes websites more cost-effective to
// operate. Additionally, Tiered Cache concentrates connections to origin servers
// so they come from a small number of data centers rather than the full set of
// network locations. This results in fewer open connections using server
// resources.
func (r *TieredCachingService) Get(ctx context.Context, query TieredCachingGetParams, opts ...option.RequestOption) (res *TieredCachingGetResponse, err error) {
	var env TieredCachingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/argo/tiered_caching", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TieredCachingEditResponse struct {
	// ID of the zone setting.
	ID TieredCachingEditResponseID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value TieredCachingEditResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                     `json:"modified_on,nullable" format:"date-time"`
	JSON       tieredCachingEditResponseJSON `json:"-"`
}

// tieredCachingEditResponseJSON contains the JSON metadata for the struct
// [TieredCachingEditResponse]
type tieredCachingEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type TieredCachingEditResponseID string

const (
	TieredCachingEditResponseIDTieredCaching TieredCachingEditResponseID = "tiered_caching"
)

func (r TieredCachingEditResponseID) IsKnown() bool {
	switch r {
	case TieredCachingEditResponseIDTieredCaching:
		return true
	}
	return false
}

// The value of the feature
type TieredCachingEditResponseValue string

const (
	TieredCachingEditResponseValueOn  TieredCachingEditResponseValue = "on"
	TieredCachingEditResponseValueOff TieredCachingEditResponseValue = "off"
)

func (r TieredCachingEditResponseValue) IsKnown() bool {
	switch r {
	case TieredCachingEditResponseValueOn, TieredCachingEditResponseValueOff:
		return true
	}
	return false
}

type TieredCachingGetResponse struct {
	// ID of the zone setting.
	ID TieredCachingGetResponseID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value TieredCachingGetResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                    `json:"modified_on,nullable" format:"date-time"`
	JSON       tieredCachingGetResponseJSON `json:"-"`
}

// tieredCachingGetResponseJSON contains the JSON metadata for the struct
// [TieredCachingGetResponse]
type tieredCachingGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type TieredCachingGetResponseID string

const (
	TieredCachingGetResponseIDTieredCaching TieredCachingGetResponseID = "tiered_caching"
)

func (r TieredCachingGetResponseID) IsKnown() bool {
	switch r {
	case TieredCachingGetResponseIDTieredCaching:
		return true
	}
	return false
}

// The value of the feature
type TieredCachingGetResponseValue string

const (
	TieredCachingGetResponseValueOn  TieredCachingGetResponseValue = "on"
	TieredCachingGetResponseValueOff TieredCachingGetResponseValue = "off"
)

func (r TieredCachingGetResponseValue) IsKnown() bool {
	switch r {
	case TieredCachingGetResponseValueOn, TieredCachingGetResponseValueOff:
		return true
	}
	return false
}

type TieredCachingEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enables Tiered Caching.
	Value param.Field[TieredCachingEditParamsValue] `json:"value,required"`
}

func (r TieredCachingEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enables Tiered Caching.
type TieredCachingEditParamsValue string

const (
	TieredCachingEditParamsValueOn  TieredCachingEditParamsValue = "on"
	TieredCachingEditParamsValueOff TieredCachingEditParamsValue = "off"
)

func (r TieredCachingEditParamsValue) IsKnown() bool {
	switch r {
	case TieredCachingEditParamsValueOn, TieredCachingEditParamsValueOff:
		return true
	}
	return false
}

type TieredCachingEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TieredCachingEditResponseEnvelopeSuccess `json:"success,required"`
	Result  TieredCachingEditResponse                `json:"result"`
	JSON    tieredCachingEditResponseEnvelopeJSON    `json:"-"`
}

// tieredCachingEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [TieredCachingEditResponseEnvelope]
type tieredCachingEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TieredCachingEditResponseEnvelopeSuccess bool

const (
	TieredCachingEditResponseEnvelopeSuccessTrue TieredCachingEditResponseEnvelopeSuccess = true
)

func (r TieredCachingEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TieredCachingEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TieredCachingGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type TieredCachingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success TieredCachingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  TieredCachingGetResponse                `json:"result"`
	JSON    tieredCachingGetResponseEnvelopeJSON    `json:"-"`
}

// tieredCachingGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TieredCachingGetResponseEnvelope]
type tieredCachingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TieredCachingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tieredCachingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type TieredCachingGetResponseEnvelopeSuccess bool

const (
	TieredCachingGetResponseEnvelopeSuccessTrue TieredCachingGetResponseEnvelopeSuccess = true
)

func (r TieredCachingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TieredCachingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
