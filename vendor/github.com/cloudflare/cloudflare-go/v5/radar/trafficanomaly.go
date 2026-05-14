// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// TrafficAnomalyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTrafficAnomalyService] method instead.
type TrafficAnomalyService struct {
	Options   []option.RequestOption
	Locations *TrafficAnomalyLocationService
}

// NewTrafficAnomalyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTrafficAnomalyService(opts ...option.RequestOption) (r *TrafficAnomalyService) {
	r = &TrafficAnomalyService{}
	r.Options = opts
	r.Locations = NewTrafficAnomalyLocationService(opts...)
	return
}

// Retrieves the latest Internet traffic anomalies, which are signals that might
// indicate an outage. These alerts are automatically detected by Radar and
// manually verified by our team.
func (r *TrafficAnomalyService) Get(ctx context.Context, query TrafficAnomalyGetParams, opts ...option.RequestOption) (res *TrafficAnomalyGetResponse, err error) {
	var env TrafficAnomalyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/traffic_anomalies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TrafficAnomalyGetResponse struct {
	TrafficAnomalies []TrafficAnomalyGetResponseTrafficAnomaly `json:"trafficAnomalies,required"`
	JSON             trafficAnomalyGetResponseJSON             `json:"-"`
}

// trafficAnomalyGetResponseJSON contains the JSON metadata for the struct
// [TrafficAnomalyGetResponse]
type trafficAnomalyGetResponseJSON struct {
	TrafficAnomalies apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaly struct {
	StartDate            string                                                   `json:"startDate,required"`
	Status               string                                                   `json:"status,required"`
	Type                 string                                                   `json:"type,required"`
	UUID                 string                                                   `json:"uuid,required"`
	ASNDetails           TrafficAnomalyGetResponseTrafficAnomaliesASNDetails      `json:"asnDetails"`
	EndDate              time.Time                                                `json:"endDate" format:"date-time"`
	LocationDetails      TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails `json:"locationDetails"`
	VisibleInDataSources []string                                                 `json:"visibleInDataSources"`
	JSON                 trafficAnomalyGetResponseTrafficAnomalyJSON              `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomalyJSON contains the JSON metadata for the
// struct [TrafficAnomalyGetResponseTrafficAnomaly]
type trafficAnomalyGetResponseTrafficAnomalyJSON struct {
	StartDate            apijson.Field
	Status               apijson.Field
	Type                 apijson.Field
	UUID                 apijson.Field
	ASNDetails           apijson.Field
	EndDate              apijson.Field
	LocationDetails      apijson.Field
	VisibleInDataSources apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomalyJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaliesASNDetails struct {
	ASN       string                                                       `json:"asn,required"`
	Name      string                                                       `json:"name,required"`
	Locations TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocations `json:"locations"`
	JSON      trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON      `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON contains the JSON
// metadata for the struct [TrafficAnomalyGetResponseTrafficAnomaliesASNDetails]
type trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON struct {
	ASN         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaliesASNDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomaliesASNDetailsJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocations struct {
	Code string                                                           `json:"code,required"`
	Name string                                                           `json:"name,required"`
	JSON trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationsJSON `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationsJSON contains the
// JSON metadata for the struct
// [TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocations]
type trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomaliesASNDetailsLocationsJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails struct {
	Code string                                                       `json:"code,required"`
	Name string                                                       `json:"name,required"`
	JSON trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON `json:"-"`
}

// trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON contains the JSON
// metadata for the struct
// [TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails]
type trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseTrafficAnomaliesLocationDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseTrafficAnomaliesLocationDetailsJSON) RawJSON() string {
	return r.raw
}

type TrafficAnomalyGetParams struct {
	// Filters results by Autonomous System. Specify a single Autonomous System Number
	// (ASN) as integer.
	ASN param.Field[int64] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[TrafficAnomalyGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64]                         `query:"offset"`
	Status param.Field[TrafficAnomalyGetParamsStatus] `query:"status"`
}

// URLQuery serializes [TrafficAnomalyGetParams]'s query parameters as
// `url.Values`.
func (r TrafficAnomalyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type TrafficAnomalyGetParamsFormat string

const (
	TrafficAnomalyGetParamsFormatJson TrafficAnomalyGetParamsFormat = "JSON"
	TrafficAnomalyGetParamsFormatCsv  TrafficAnomalyGetParamsFormat = "CSV"
)

func (r TrafficAnomalyGetParamsFormat) IsKnown() bool {
	switch r {
	case TrafficAnomalyGetParamsFormatJson, TrafficAnomalyGetParamsFormatCsv:
		return true
	}
	return false
}

type TrafficAnomalyGetParamsStatus string

const (
	TrafficAnomalyGetParamsStatusVerified   TrafficAnomalyGetParamsStatus = "VERIFIED"
	TrafficAnomalyGetParamsStatusUnverified TrafficAnomalyGetParamsStatus = "UNVERIFIED"
)

func (r TrafficAnomalyGetParamsStatus) IsKnown() bool {
	switch r {
	case TrafficAnomalyGetParamsStatusVerified, TrafficAnomalyGetParamsStatusUnverified:
		return true
	}
	return false
}

type TrafficAnomalyGetResponseEnvelope struct {
	Result  TrafficAnomalyGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    trafficAnomalyGetResponseEnvelopeJSON `json:"-"`
}

// trafficAnomalyGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [TrafficAnomalyGetResponseEnvelope]
type trafficAnomalyGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TrafficAnomalyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r trafficAnomalyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
