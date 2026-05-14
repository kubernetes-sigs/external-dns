// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

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
)

// EntityLocationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityLocationService] method instead.
type EntityLocationService struct {
	Options []option.RequestOption
}

// NewEntityLocationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntityLocationService(opts ...option.RequestOption) (r *EntityLocationService) {
	r = &EntityLocationService{}
	r.Options = opts
	return
}

// Retrieves a list of locations.
func (r *EntityLocationService) List(ctx context.Context, query EntityLocationListParams, opts ...option.RequestOption) (res *EntityLocationListResponse, err error) {
	var env EntityLocationListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/entities/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested location information. (A confidence level below `5`
// indicates a low level of confidence in the traffic data - normally this happens
// because Cloudflare has a small amount of traffic from/to this location).
func (r *EntityLocationService) Get(ctx context.Context, location string, query EntityLocationGetParams, opts ...option.RequestOption) (res *EntityLocationGetResponse, err error) {
	var env EntityLocationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if location == "" {
		err = errors.New("missing required location parameter")
		return
	}
	path := fmt.Sprintf("radar/entities/locations/%s", location)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EntityLocationListResponse struct {
	Locations []EntityLocationListResponseLocation `json:"locations,required"`
	JSON      entityLocationListResponseJSON       `json:"-"`
}

// entityLocationListResponseJSON contains the JSON metadata for the struct
// [EntityLocationListResponse]
type entityLocationListResponseJSON struct {
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityLocationListResponseJSON) RawJSON() string {
	return r.raw
}

type EntityLocationListResponseLocation struct {
	Alpha2 string `json:"alpha2,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                 `json:"longitude,required"`
	Name      string                                 `json:"name,required"`
	JSON      entityLocationListResponseLocationJSON `json:"-"`
}

// entityLocationListResponseLocationJSON contains the JSON metadata for the struct
// [EntityLocationListResponseLocation]
type entityLocationListResponseLocationJSON struct {
	Alpha2      apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityLocationListResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityLocationListResponseLocationJSON) RawJSON() string {
	return r.raw
}

type EntityLocationGetResponse struct {
	Location EntityLocationGetResponseLocation `json:"location,required"`
	JSON     entityLocationGetResponseJSON     `json:"-"`
}

// entityLocationGetResponseJSON contains the JSON metadata for the struct
// [EntityLocationGetResponse]
type entityLocationGetResponseJSON struct {
	Location    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type EntityLocationGetResponseLocation struct {
	Alpha2          string `json:"alpha2,required"`
	ConfidenceLevel int64  `json:"confidenceLevel,required"`
	// A numeric string.
	Latitude string `json:"latitude,required"`
	// A numeric string.
	Longitude string                                `json:"longitude,required"`
	Name      string                                `json:"name,required"`
	Region    string                                `json:"region,required"`
	Subregion string                                `json:"subregion,required"`
	JSON      entityLocationGetResponseLocationJSON `json:"-"`
}

// entityLocationGetResponseLocationJSON contains the JSON metadata for the struct
// [EntityLocationGetResponseLocation]
type entityLocationGetResponseLocationJSON struct {
	Alpha2          apijson.Field
	ConfidenceLevel apijson.Field
	Latitude        apijson.Field
	Longitude       apijson.Field
	Name            apijson.Field
	Region          apijson.Field
	Subregion       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *EntityLocationGetResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityLocationGetResponseLocationJSON) RawJSON() string {
	return r.raw
}

type EntityLocationListParams struct {
	// Format in which results will be returned.
	Format param.Field[EntityLocationListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [EntityLocationListParams]'s query parameters as
// `url.Values`.
func (r EntityLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityLocationListParamsFormat string

const (
	EntityLocationListParamsFormatJson EntityLocationListParamsFormat = "JSON"
	EntityLocationListParamsFormatCsv  EntityLocationListParamsFormat = "CSV"
)

func (r EntityLocationListParamsFormat) IsKnown() bool {
	switch r {
	case EntityLocationListParamsFormatJson, EntityLocationListParamsFormatCsv:
		return true
	}
	return false
}

type EntityLocationListResponseEnvelope struct {
	Result  EntityLocationListResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    entityLocationListResponseEnvelopeJSON `json:"-"`
}

// entityLocationListResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityLocationListResponseEnvelope]
type entityLocationListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityLocationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityLocationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EntityLocationGetParams struct {
	// Format in which results will be returned.
	Format param.Field[EntityLocationGetParamsFormat] `query:"format"`
}

// URLQuery serializes [EntityLocationGetParams]'s query parameters as
// `url.Values`.
func (r EntityLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type EntityLocationGetParamsFormat string

const (
	EntityLocationGetParamsFormatJson EntityLocationGetParamsFormat = "JSON"
	EntityLocationGetParamsFormatCsv  EntityLocationGetParamsFormat = "CSV"
)

func (r EntityLocationGetParamsFormat) IsKnown() bool {
	switch r {
	case EntityLocationGetParamsFormatJson, EntityLocationGetParamsFormatCsv:
		return true
	}
	return false
}

type EntityLocationGetResponseEnvelope struct {
	Result  EntityLocationGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    entityLocationGetResponseEnvelopeJSON `json:"-"`
}

// entityLocationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EntityLocationGetResponseEnvelope]
type entityLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EntityLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r entityLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
