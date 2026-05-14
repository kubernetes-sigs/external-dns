// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

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

// RegionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegionService] method instead.
type RegionService struct {
	Options []option.RequestOption
}

// NewRegionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRegionService(opts ...option.RequestOption) (r *RegionService) {
	r = &RegionService{}
	r.Options = opts
	return
}

// List all region mappings.
func (r *RegionService) List(ctx context.Context, params RegionListParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env RegionListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/regions", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a single region mapping.
func (r *RegionService) Get(ctx context.Context, regionID RegionGetParamsRegionID, query RegionGetParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env RegionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/regions/%v", query.AccountID, regionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RegionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Two-letter alpha-2 country code followed in ISO 3166-1.
	CountryCodeA2 param.Field[string] `query:"country_code_a2"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCode param.Field[string] `query:"subdivision_code"`
	// Two-letter subdivision code followed in ISO 3166-2.
	SubdivisionCodeA2 param.Field[string] `query:"subdivision_code_a2"`
}

// URLQuery serializes [RegionListParams]'s query parameters as `url.Values`.
func (r RegionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RegionListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Whether the API call was successful.
	Success RegionListResponseEnvelopeSuccess `json:"success,required"`
	JSON    regionListResponseEnvelopeJSON    `json:"-"`
}

// regionListResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegionListResponseEnvelope]
type regionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegionListResponseEnvelopeSuccess bool

const (
	RegionListResponseEnvelopeSuccessTrue RegionListResponseEnvelopeSuccess = true
)

func (r RegionListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RegionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A list of Cloudflare regions. WNAM: Western North America, ENAM: Eastern North
// America, WEU: Western Europe, EEU: Eastern Europe, NSAM: Northern South America,
// SSAM: Southern South America, OC: Oceania, ME: Middle East, NAF: North Africa,
// SAF: South Africa, SAS: Southern Asia, SEAS: South East Asia, NEAS: North East
// Asia).
type RegionGetParamsRegionID string

const (
	RegionGetParamsRegionIDWnam RegionGetParamsRegionID = "WNAM"
	RegionGetParamsRegionIDEnam RegionGetParamsRegionID = "ENAM"
	RegionGetParamsRegionIDWeu  RegionGetParamsRegionID = "WEU"
	RegionGetParamsRegionIDEeu  RegionGetParamsRegionID = "EEU"
	RegionGetParamsRegionIDNsam RegionGetParamsRegionID = "NSAM"
	RegionGetParamsRegionIDSsam RegionGetParamsRegionID = "SSAM"
	RegionGetParamsRegionIDOc   RegionGetParamsRegionID = "OC"
	RegionGetParamsRegionIDMe   RegionGetParamsRegionID = "ME"
	RegionGetParamsRegionIDNaf  RegionGetParamsRegionID = "NAF"
	RegionGetParamsRegionIDSaf  RegionGetParamsRegionID = "SAF"
	RegionGetParamsRegionIDSas  RegionGetParamsRegionID = "SAS"
	RegionGetParamsRegionIDSeas RegionGetParamsRegionID = "SEAS"
	RegionGetParamsRegionIDNeas RegionGetParamsRegionID = "NEAS"
)

func (r RegionGetParamsRegionID) IsKnown() bool {
	switch r {
	case RegionGetParamsRegionIDWnam, RegionGetParamsRegionIDEnam, RegionGetParamsRegionIDWeu, RegionGetParamsRegionIDEeu, RegionGetParamsRegionIDNsam, RegionGetParamsRegionIDSsam, RegionGetParamsRegionIDOc, RegionGetParamsRegionIDMe, RegionGetParamsRegionIDNaf, RegionGetParamsRegionIDSaf, RegionGetParamsRegionIDSas, RegionGetParamsRegionIDSeas, RegionGetParamsRegionIDNeas:
		return true
	}
	return false
}

type RegionGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// A list of countries and subdivisions mapped to a region.
	Result interface{} `json:"result,required"`
	// Whether the API call was successful.
	Success RegionGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    regionGetResponseEnvelopeJSON    `json:"-"`
}

// regionGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RegionGetResponseEnvelope]
type regionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RegionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r regionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RegionGetResponseEnvelopeSuccess bool

const (
	RegionGetResponseEnvelopeSuccessTrue RegionGetResponseEnvelopeSuccess = true
)

func (r RegionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RegionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
