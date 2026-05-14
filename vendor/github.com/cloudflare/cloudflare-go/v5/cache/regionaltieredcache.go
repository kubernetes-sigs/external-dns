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

// RegionalTieredCacheService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionalTieredCacheService] method instead.
type RegionalTieredCacheService struct {
	Options []option.RequestOption
}

// NewRegionalTieredCacheService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegionalTieredCacheService(opts ...option.RequestOption) (r *RegionalTieredCacheService) {
	r = &RegionalTieredCacheService{}
	r.Options = opts
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *RegionalTieredCacheService) Edit(ctx context.Context, params RegionalTieredCacheEditParams, opts ...option.RequestOption) (res *RegionalTieredCacheEditResponse, err error) {
	var env RegionalTieredCacheEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Instructs Cloudflare to check a regional hub data center on the way to your
// upper tier. This can help improve performance for smart and custom tiered cache
// topologies.
func (r *RegionalTieredCacheService) Get(ctx context.Context, query RegionalTieredCacheGetParams, opts ...option.RequestOption) (res *RegionalTieredCacheGetResponse, err error) {
	var env RegionalTieredCacheGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/cache/regional_tiered_cache", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// ID of the zone setting.
type RegionalTieredCache string

const (
	RegionalTieredCacheTcRegional RegionalTieredCache = "tc_regional"
)

func (r RegionalTieredCache) IsKnown() bool {
	switch r {
	case RegionalTieredCacheTcRegional:
		return true
	}
	return false
}

type RegionalTieredCacheEditResponse struct {
	// ID of the zone setting.
	ID RegionalTieredCache `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value RegionalTieredCacheEditResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                           `json:"modified_on,nullable" format:"date-time"`
	JSON       regionalTieredCacheEditResponseJSON `json:"-"`
}

// regionalTieredCacheEditResponseJSON contains the JSON metadata for the struct
// [RegionalTieredCacheEditResponse]
type regionalTieredCacheEditResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheEditResponseJSON) RawJSON() string {
	return r.raw
}

// The value of the feature
type RegionalTieredCacheEditResponseValue string

const (
	RegionalTieredCacheEditResponseValueOn  RegionalTieredCacheEditResponseValue = "on"
	RegionalTieredCacheEditResponseValueOff RegionalTieredCacheEditResponseValue = "off"
)

func (r RegionalTieredCacheEditResponseValue) IsKnown() bool {
	switch r {
	case RegionalTieredCacheEditResponseValueOn, RegionalTieredCacheEditResponseValueOff:
		return true
	}
	return false
}

type RegionalTieredCacheGetResponse struct {
	// ID of the zone setting.
	ID RegionalTieredCache `json:"id,required"`
	// Whether the setting is editable
	Editable bool `json:"editable,required"`
	// The value of the feature
	Value RegionalTieredCacheGetResponseValue `json:"value,required"`
	// Last time this setting was modified.
	ModifiedOn time.Time                          `json:"modified_on,nullable" format:"date-time"`
	JSON       regionalTieredCacheGetResponseJSON `json:"-"`
}

// regionalTieredCacheGetResponseJSON contains the JSON metadata for the struct
// [RegionalTieredCacheGetResponse]
type regionalTieredCacheGetResponseJSON struct {
	ID          apijson.Field
	Editable    apijson.Field
	Value       apijson.Field
	ModifiedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheGetResponseJSON) RawJSON() string {
	return r.raw
}

// The value of the feature
type RegionalTieredCacheGetResponseValue string

const (
	RegionalTieredCacheGetResponseValueOn  RegionalTieredCacheGetResponseValue = "on"
	RegionalTieredCacheGetResponseValueOff RegionalTieredCacheGetResponseValue = "off"
)

func (r RegionalTieredCacheGetResponseValue) IsKnown() bool {
	switch r {
	case RegionalTieredCacheGetResponseValueOn, RegionalTieredCacheGetResponseValueOff:
		return true
	}
	return false
}

type RegionalTieredCacheEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Value of the Regional Tiered Cache zone setting.
	Value param.Field[RegionalTieredCacheEditParamsValue] `json:"value,required"`
}

func (r RegionalTieredCacheEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Value of the Regional Tiered Cache zone setting.
type RegionalTieredCacheEditParamsValue string

const (
	RegionalTieredCacheEditParamsValueOn  RegionalTieredCacheEditParamsValue = "on"
	RegionalTieredCacheEditParamsValueOff RegionalTieredCacheEditParamsValue = "off"
)

func (r RegionalTieredCacheEditParamsValue) IsKnown() bool {
	switch r {
	case RegionalTieredCacheEditParamsValueOn, RegionalTieredCacheEditParamsValueOff:
		return true
	}
	return false
}

type RegionalTieredCacheEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RegionalTieredCacheEditResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalTieredCacheEditResponse                `json:"result"`
	JSON    regionalTieredCacheEditResponseEnvelopeJSON    `json:"-"`
}

// regionalTieredCacheEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalTieredCacheEditResponseEnvelope]
type regionalTieredCacheEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalTieredCacheEditResponseEnvelopeSuccess bool

const (
	RegionalTieredCacheEditResponseEnvelopeSuccessTrue RegionalTieredCacheEditResponseEnvelopeSuccess = true
)

func (r RegionalTieredCacheEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalTieredCacheEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegionalTieredCacheGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RegionalTieredCacheGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success RegionalTieredCacheGetResponseEnvelopeSuccess `json:"success,required"`
	Result  RegionalTieredCacheGetResponse                `json:"result"`
	JSON    regionalTieredCacheGetResponseEnvelopeJSON    `json:"-"`
}

// regionalTieredCacheGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RegionalTieredCacheGetResponseEnvelope]
type regionalTieredCacheGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionalTieredCacheGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionalTieredCacheGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RegionalTieredCacheGetResponseEnvelopeSuccess bool

const (
	RegionalTieredCacheGetResponseEnvelopeSuccessTrue RegionalTieredCacheGetResponseEnvelopeSuccess = true
)

func (r RegionalTieredCacheGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionalTieredCacheGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
