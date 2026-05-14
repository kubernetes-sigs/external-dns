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

// AnnotationOutageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnnotationOutageService] method instead.
type AnnotationOutageService struct {
	Options []option.RequestOption
}

// NewAnnotationOutageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnnotationOutageService(opts ...option.RequestOption) (r *AnnotationOutageService) {
	r = &AnnotationOutageService{}
	r.Options = opts
	return
}

// Retrieves the latest Internet outages and anomalies.
func (r *AnnotationOutageService) Get(ctx context.Context, query AnnotationOutageGetParams, opts ...option.RequestOption) (res *AnnotationOutageGetResponse, err error) {
	var env AnnotationOutageGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/annotations/outages"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the number of outages by location.
func (r *AnnotationOutageService) Locations(ctx context.Context, query AnnotationOutageLocationsParams, opts ...option.RequestOption) (res *AnnotationOutageLocationsResponse, err error) {
	var env AnnotationOutageLocationsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/annotations/outages/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnnotationOutageGetResponse struct {
	Annotations []AnnotationOutageGetResponseAnnotation `json:"annotations,required"`
	JSON        annotationOutageGetResponseJSON         `json:"-"`
}

// annotationOutageGetResponseJSON contains the JSON metadata for the struct
// [AnnotationOutageGetResponse]
type annotationOutageGetResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotation struct {
	ID               string                                                  `json:"id,required"`
	ASNs             []int64                                                 `json:"asns,required"`
	ASNsDetails      []AnnotationOutageGetResponseAnnotationsASNsDetail      `json:"asnsDetails,required"`
	DataSource       string                                                  `json:"dataSource,required"`
	EventType        string                                                  `json:"eventType,required"`
	Locations        []string                                                `json:"locations,required"`
	LocationsDetails []AnnotationOutageGetResponseAnnotationsLocationsDetail `json:"locationsDetails,required"`
	Outage           AnnotationOutageGetResponseAnnotationsOutage            `json:"outage,required"`
	StartDate        time.Time                                               `json:"startDate,required" format:"date-time"`
	Description      string                                                  `json:"description"`
	EndDate          time.Time                                               `json:"endDate" format:"date-time"`
	LinkedURL        string                                                  `json:"linkedUrl"`
	Scope            string                                                  `json:"scope"`
	JSON             annotationOutageGetResponseAnnotationJSON               `json:"-"`
}

// annotationOutageGetResponseAnnotationJSON contains the JSON metadata for the
// struct [AnnotationOutageGetResponseAnnotation]
type annotationOutageGetResponseAnnotationJSON struct {
	ID               apijson.Field
	ASNs             apijson.Field
	ASNsDetails      apijson.Field
	DataSource       apijson.Field
	EventType        apijson.Field
	Locations        apijson.Field
	LocationsDetails apijson.Field
	Outage           apijson.Field
	StartDate        apijson.Field
	Description      apijson.Field
	EndDate          apijson.Field
	LinkedURL        apijson.Field
	Scope            apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsASNsDetail struct {
	ASN       string                                                     `json:"asn,required"`
	Name      string                                                     `json:"name,required"`
	Locations AnnotationOutageGetResponseAnnotationsASNsDetailsLocations `json:"locations"`
	JSON      annotationOutageGetResponseAnnotationsASNsDetailJSON       `json:"-"`
}

// annotationOutageGetResponseAnnotationsASNsDetailJSON contains the JSON metadata
// for the struct [AnnotationOutageGetResponseAnnotationsASNsDetail]
type annotationOutageGetResponseAnnotationsASNsDetailJSON struct {
	ASN         apijson.Field
	Name        apijson.Field
	Locations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsASNsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsASNsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsASNsDetailsLocations struct {
	Code string                                                         `json:"code,required"`
	Name string                                                         `json:"name,required"`
	JSON annotationOutageGetResponseAnnotationsASNsDetailsLocationsJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsASNsDetailsLocationsJSON contains the JSON
// metadata for the struct
// [AnnotationOutageGetResponseAnnotationsASNsDetailsLocations]
type annotationOutageGetResponseAnnotationsASNsDetailsLocationsJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsASNsDetailsLocations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsASNsDetailsLocationsJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsLocationsDetail struct {
	Code string                                                    `json:"code,required"`
	Name string                                                    `json:"name,required"`
	JSON annotationOutageGetResponseAnnotationsLocationsDetailJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsLocationsDetailJSON contains the JSON
// metadata for the struct [AnnotationOutageGetResponseAnnotationsLocationsDetail]
type annotationOutageGetResponseAnnotationsLocationsDetailJSON struct {
	Code        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsLocationsDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsLocationsDetailJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetResponseAnnotationsOutage struct {
	OutageCause string                                           `json:"outageCause,required"`
	OutageType  string                                           `json:"outageType,required"`
	JSON        annotationOutageGetResponseAnnotationsOutageJSON `json:"-"`
}

// annotationOutageGetResponseAnnotationsOutageJSON contains the JSON metadata for
// the struct [AnnotationOutageGetResponseAnnotationsOutage]
type annotationOutageGetResponseAnnotationsOutageJSON struct {
	OutageCause apijson.Field
	OutageType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseAnnotationsOutage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseAnnotationsOutageJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsResponse struct {
	Annotations []AnnotationOutageLocationsResponseAnnotation `json:"annotations,required"`
	JSON        annotationOutageLocationsResponseJSON         `json:"-"`
}

// annotationOutageLocationsResponseJSON contains the JSON metadata for the struct
// [AnnotationOutageLocationsResponse]
type annotationOutageLocationsResponseJSON struct {
	Annotations apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsResponseAnnotation struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                                          `json:"value,required"`
	JSON  annotationOutageLocationsResponseAnnotationJSON `json:"-"`
}

// annotationOutageLocationsResponseAnnotationJSON contains the JSON metadata for
// the struct [AnnotationOutageLocationsResponseAnnotation]
type annotationOutageLocationsResponseAnnotationJSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponseAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseAnnotationJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageGetParams struct {
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
	Format param.Field[AnnotationOutageGetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify an alpha-2 location code.
	Location param.Field[string] `query:"location"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [AnnotationOutageGetParams]'s query parameters as
// `url.Values`.
func (r AnnotationOutageGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AnnotationOutageGetParamsFormat string

const (
	AnnotationOutageGetParamsFormatJson AnnotationOutageGetParamsFormat = "JSON"
	AnnotationOutageGetParamsFormatCsv  AnnotationOutageGetParamsFormat = "CSV"
)

func (r AnnotationOutageGetParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationOutageGetParamsFormatJson, AnnotationOutageGetParamsFormatCsv:
		return true
	}
	return false
}

type AnnotationOutageGetResponseEnvelope struct {
	Result  AnnotationOutageGetResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    annotationOutageGetResponseEnvelopeJSON `json:"-"`
}

// annotationOutageGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnnotationOutageGetResponseEnvelope]
type annotationOutageGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnnotationOutageLocationsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range.
	DateRange param.Field[string] `query:"dateRange"`
	// Start of the date range (inclusive).
	DateStart param.Field[time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AnnotationOutageLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
}

// URLQuery serializes [AnnotationOutageLocationsParams]'s query parameters as
// `url.Values`.
func (r AnnotationOutageLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AnnotationOutageLocationsParamsFormat string

const (
	AnnotationOutageLocationsParamsFormatJson AnnotationOutageLocationsParamsFormat = "JSON"
	AnnotationOutageLocationsParamsFormatCsv  AnnotationOutageLocationsParamsFormat = "CSV"
)

func (r AnnotationOutageLocationsParamsFormat) IsKnown() bool {
	switch r {
	case AnnotationOutageLocationsParamsFormatJson, AnnotationOutageLocationsParamsFormatCsv:
		return true
	}
	return false
}

type AnnotationOutageLocationsResponseEnvelope struct {
	Result  AnnotationOutageLocationsResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    annotationOutageLocationsResponseEnvelopeJSON `json:"-"`
}

// annotationOutageLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnnotationOutageLocationsResponseEnvelope]
type annotationOutageLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnnotationOutageLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r annotationOutageLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
