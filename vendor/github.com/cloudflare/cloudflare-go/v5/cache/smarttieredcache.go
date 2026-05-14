// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cache

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

// SmartTieredCacheService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSmartTieredCacheService] method instead.
type SmartTieredCacheService struct {
	Options []option.RequestOption
}

// NewSmartTieredCacheService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSmartTieredCacheService(opts ...option.RequestOption) (r *SmartTieredCacheService) {
	r = &SmartTieredCacheService{}
	r.Options = opts
	return
}

// Smart Tiered Cache dynamically selects the single closest upper tier for each of
// your website’s origins with no configuration required, using our in-house
// performance and routing data. Cloudflare collects latency data for each request
// to an origin, and uses the latency data to determine how well any upper-tier
// data center is connected with an origin. As a result, Cloudflare can select the
// data center with the lowest latency to be the upper-tier for an origin.
func (r *SmartTieredCacheService) Delete(ctx context.Context, body SmartTieredCacheDeleteParams, opts ...option.RequestOption) (res *SmartTieredCacheDeleteResponse, err error) {
	var env SmartTieredCacheDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Smart Tiered Cache dynamically selects the single closest upper tier for each of
// your website’s origins with no configuration required, using our in-house
// performance and routing data. Cloudflare collects latency data for each request
// to an origin, and uses the latency data to determine how well any upper-tier
// data center is connected with an origin. As a result, Cloudflare can select the
// data center with the lowest latency to be the upper-tier for an origin.
func (r *SmartTieredCacheService) Edit(ctx context.Context, params SmartTieredCacheEditParams, opts ...option.RequestOption) (res *SmartTieredCacheEditResponse, err error) {
	var env SmartTieredCacheEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Smart Tiered Cache dynamically selects the single closest upper tier for each of
// your website’s origins with no configuration required, using our in-house
// performance and routing data. Cloudflare collects latency data for each request
// to an origin, and uses the latency data to determine how well any upper-tier
// data center is connected with an origin. As a result, Cloudflare can select the
// data center with the lowest latency to be the upper-tier for an origin.
func (r *SmartTieredCacheService) Get(ctx context.Context, query SmartTieredCacheGetParams, opts ...option.RequestOption) (res *SmartTieredCacheGetResponse, err error) {
	var env SmartTieredCacheGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/tiered_cache_smart_topology_enable", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SmartTieredCacheDeleteResponse struct {
	// ID of the zone setting.
	ID SmartTieredCacheDeleteResponseID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       smartTieredCacheDeleteResponseJSON `json:"-"`
}

// smartTieredCacheDeleteResponseJSON contains the JSON metadata for the struct
// [SmartTieredCacheDeleteResponse]
type smartTieredCacheDeleteResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SmartTieredCacheDeleteResponseID string

const (
	SmartTieredCacheDeleteResponseIDTieredCacheSmartTopologyEnable SmartTieredCacheDeleteResponseID = "tiered_cache_smart_topology_enable"
)

func (r SmartTieredCacheDeleteResponseID) IsKnown() bool {
	switch r {
	case SmartTieredCacheDeleteResponseIDTieredCacheSmartTopologyEnable:
		return true
	}
	return false
}

type SmartTieredCacheEditResponse struct {
	// ID of the zone setting.
	ID SmartTieredCacheEditResponseID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value SmartTieredCacheEditResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                        `json:"modified_on,nullable" format:"date-time"`
	JSON       smartTieredCacheEditResponseJSON `json:"-"`
}

// smartTieredCacheEditResponseJSON contains the JSON metadata for the struct
// [SmartTieredCacheEditResponse]
type smartTieredCacheEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheEditResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SmartTieredCacheEditResponseID string

const (
	SmartTieredCacheEditResponseIDTieredCacheSmartTopologyEnable SmartTieredCacheEditResponseID = "tiered_cache_smart_topology_enable"
)

func (r SmartTieredCacheEditResponseID) IsKnown() bool {
	switch r {
	case SmartTieredCacheEditResponseIDTieredCacheSmartTopologyEnable:
		return true
	}
	return false
}

// The value of the feature
type SmartTieredCacheEditResponseValue string

const (
	SmartTieredCacheEditResponseValueOn  SmartTieredCacheEditResponseValue = "on"
	SmartTieredCacheEditResponseValueOff SmartTieredCacheEditResponseValue = "off"
)

func (r SmartTieredCacheEditResponseValue) IsKnown() bool {
	switch r {
	case SmartTieredCacheEditResponseValueOn, SmartTieredCacheEditResponseValueOff:
		return true
	}
	return false
}

type SmartTieredCacheGetResponse struct {
	// ID of the zone setting.
	ID SmartTieredCacheGetResponseID `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value SmartTieredCacheGetResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                       `json:"modified_on,nullable" format:"date-time"`
	JSON       smartTieredCacheGetResponseJSON `json:"-"`
}

// smartTieredCacheGetResponseJSON contains the JSON metadata for the struct
// [SmartTieredCacheGetResponse]
type smartTieredCacheGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheGetResponseJSON) RawJSON() string {
	return r.raw
}

// ID of the zone setting.
type SmartTieredCacheGetResponseID string

const (
	SmartTieredCacheGetResponseIDTieredCacheSmartTopologyEnable SmartTieredCacheGetResponseID = "tiered_cache_smart_topology_enable"
)

func (r SmartTieredCacheGetResponseID) IsKnown() bool {
	switch r {
	case SmartTieredCacheGetResponseIDTieredCacheSmartTopologyEnable:
		return true
	}
	return false
}

// The value of the feature
type SmartTieredCacheGetResponseValue string

const (
	SmartTieredCacheGetResponseValueOn  SmartTieredCacheGetResponseValue = "on"
	SmartTieredCacheGetResponseValueOff SmartTieredCacheGetResponseValue = "off"
)

func (r SmartTieredCacheGetResponseValue) IsKnown() bool {
	switch r {
	case SmartTieredCacheGetResponseValueOn, SmartTieredCacheGetResponseValueOff:
		return true
	}
	return false
}

type SmartTieredCacheDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartTieredCacheDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SmartTieredCacheDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  SmartTieredCacheDeleteResponse                `json:"result"`
	JSON    smartTieredCacheDeleteResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCacheDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCacheDeleteResponseEnvelope]
type smartTieredCacheDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCacheDeleteResponseEnvelopeSuccess bool

const (
	SmartTieredCacheDeleteResponseEnvelopeSuccessTrue SmartTieredCacheDeleteResponseEnvelopeSuccess = true
)

func (r SmartTieredCacheDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartTieredCacheDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartTieredCacheEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Enable or disable the Smart Tiered Cache
	Value param.Field[SmartTieredCacheEditParamsValue] `json:"value,required"`
}

func (r SmartTieredCacheEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Enable or disable the Smart Tiered Cache
type SmartTieredCacheEditParamsValue string

const (
	SmartTieredCacheEditParamsValueOn  SmartTieredCacheEditParamsValue = "on"
	SmartTieredCacheEditParamsValueOff SmartTieredCacheEditParamsValue = "off"
)

func (r SmartTieredCacheEditParamsValue) IsKnown() bool {
	switch r {
	case SmartTieredCacheEditParamsValueOn, SmartTieredCacheEditParamsValueOff:
		return true
	}
	return false
}

type SmartTieredCacheEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SmartTieredCacheEditResponseEnvelopeSuccess `json:"success,required"`
	Result  SmartTieredCacheEditResponse                `json:"result"`
	JSON    smartTieredCacheEditResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCacheEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCacheEditResponseEnvelope]
type smartTieredCacheEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCacheEditResponseEnvelopeSuccess bool

const (
	SmartTieredCacheEditResponseEnvelopeSuccessTrue SmartTieredCacheEditResponseEnvelopeSuccess = true
)

func (r SmartTieredCacheEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartTieredCacheEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SmartTieredCacheGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SmartTieredCacheGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SmartTieredCacheGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SmartTieredCacheGetResponse                `json:"result"`
	JSON    smartTieredCacheGetResponseEnvelopeJSON    `json:"-"`
}

// smartTieredCacheGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [SmartTieredCacheGetResponseEnvelope]
type smartTieredCacheGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SmartTieredCacheGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r smartTieredCacheGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SmartTieredCacheGetResponseEnvelopeSuccess bool

const (
	SmartTieredCacheGetResponseEnvelopeSuccessTrue SmartTieredCacheGetResponseEnvelopeSuccess = true
)

func (r SmartTieredCacheGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SmartTieredCacheGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
