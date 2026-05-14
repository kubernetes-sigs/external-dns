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

// NetflowTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetflowTopService] method instead.
type NetflowTopService struct {
	Options []option.RequestOption
}

// NewNetflowTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetflowTopService(opts ...option.RequestOption) (r *NetflowTopService) {
	r = &NetflowTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by network traffic (NetFlows).
func (r *NetflowTopService) Ases(ctx context.Context, query NetflowTopAsesParams, opts ...option.RequestOption) (res *NetflowTopAsesResponse, err error) {
	var env NetflowTopAsesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations by network traffic (NetFlows).
func (r *NetflowTopService) Locations(ctx context.Context, query NetflowTopLocationsParams, opts ...option.RequestOption) (res *NetflowTopLocationsResponse, err error) {
	var env NetflowTopLocationsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetflowTopAsesResponse struct {
	// Metadata for the results.
	Meta NetflowTopAsesResponseMeta   `json:"meta,required"`
	Top0 []NetflowTopAsesResponseTop0 `json:"top_0,required"`
	JSON netflowTopAsesResponseJSON   `json:"-"`
}

// netflowTopAsesResponseJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponse]
type netflowTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetflowTopAsesResponseMeta struct {
	ConfidenceInfo NetflowTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []NetflowTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetflowTopAsesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetflowTopAsesResponseMetaUnit `json:"units,required"`
	JSON  netflowTopAsesResponseMetaJSON   `json:"-"`
}

// netflowTopAsesResponseMetaJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseMeta]
type netflowTopAsesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetflowTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesResponseMetaConfidenceInfo struct {
	Annotations []NetflowTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                        `json:"level,required"`
	JSON  netflowTopAsesResponseMetaConfidenceInfoJSON `json:"-"`
}

// netflowTopAsesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [NetflowTopAsesResponseMetaConfidenceInfo]
type netflowTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetflowTopAsesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	LinkedURL       string                                                 `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                              `json:"startDate,required" format:"date-time"`
	JSON            netflowTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netflowTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [NetflowTopAsesResponseMetaConfidenceInfoAnnotation]
type netflowTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NetflowTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      netflowTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// netflowTopAsesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetflowTopAsesResponseMetaDateRange]
type netflowTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetflowTopAsesResponseMetaNormalization string

const (
	NetflowTopAsesResponseMetaNormalizationPercentage           NetflowTopAsesResponseMetaNormalization = "PERCENTAGE"
	NetflowTopAsesResponseMetaNormalizationMin0Max              NetflowTopAsesResponseMetaNormalization = "MIN0_MAX"
	NetflowTopAsesResponseMetaNormalizationMinMax               NetflowTopAsesResponseMetaNormalization = "MIN_MAX"
	NetflowTopAsesResponseMetaNormalizationRawValues            NetflowTopAsesResponseMetaNormalization = "RAW_VALUES"
	NetflowTopAsesResponseMetaNormalizationPercentageChange     NetflowTopAsesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetflowTopAsesResponseMetaNormalizationRollingAverage       NetflowTopAsesResponseMetaNormalization = "ROLLING_AVERAGE"
	NetflowTopAsesResponseMetaNormalizationOverlappedPercentage NetflowTopAsesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetflowTopAsesResponseMetaNormalizationRatio                NetflowTopAsesResponseMetaNormalization = "RATIO"
)

func (r NetflowTopAsesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetflowTopAsesResponseMetaNormalizationPercentage, NetflowTopAsesResponseMetaNormalizationMin0Max, NetflowTopAsesResponseMetaNormalizationMinMax, NetflowTopAsesResponseMetaNormalizationRawValues, NetflowTopAsesResponseMetaNormalizationPercentageChange, NetflowTopAsesResponseMetaNormalizationRollingAverage, NetflowTopAsesResponseMetaNormalizationOverlappedPercentage, NetflowTopAsesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetflowTopAsesResponseMetaUnit struct {
	Name  string                             `json:"name,required"`
	Value string                             `json:"value,required"`
	JSON  netflowTopAsesResponseMetaUnitJSON `json:"-"`
}

// netflowTopAsesResponseMetaUnitJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseMetaUnit]
type netflowTopAsesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesResponseTop0 struct {
	ClientASN    float64 `json:"clientASN,required"`
	ClientAsName string  `json:"clientASName,required"`
	// A numeric string.
	Value string                         `json:"value,required"`
	JSON  netflowTopAsesResponseTop0JSON `json:"-"`
}

// netflowTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseTop0]
type netflowTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *NetflowTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponse struct {
	// Metadata for the results.
	Meta NetflowTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []NetflowTopLocationsResponseTop0 `json:"top_0,required"`
	JSON netflowTopLocationsResponseJSON   `json:"-"`
}

// netflowTopLocationsResponseJSON contains the JSON metadata for the struct
// [NetflowTopLocationsResponse]
type netflowTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type NetflowTopLocationsResponseMeta struct {
	ConfidenceInfo NetflowTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []NetflowTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization NetflowTopLocationsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []NetflowTopLocationsResponseMetaUnit `json:"units,required"`
	JSON  netflowTopLocationsResponseMetaJSON   `json:"-"`
}

// netflowTopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [NetflowTopLocationsResponseMeta]
type netflowTopLocationsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []NetflowTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  netflowTopLocationsResponseMetaConfidenceInfoJSON `json:"-"`
}

// netflowTopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [NetflowTopLocationsResponseMetaConfidenceInfo]
type netflowTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type NetflowTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            netflowTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// netflowTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [NetflowTopLocationsResponseMetaConfidenceInfoAnnotation]
type netflowTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EndDate         apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	LinkedURL       apijson.Field
	StartDate       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      netflowTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// netflowTopLocationsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [NetflowTopLocationsResponseMetaDateRange]
type netflowTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type NetflowTopLocationsResponseMetaNormalization string

const (
	NetflowTopLocationsResponseMetaNormalizationPercentage           NetflowTopLocationsResponseMetaNormalization = "PERCENTAGE"
	NetflowTopLocationsResponseMetaNormalizationMin0Max              NetflowTopLocationsResponseMetaNormalization = "MIN0_MAX"
	NetflowTopLocationsResponseMetaNormalizationMinMax               NetflowTopLocationsResponseMetaNormalization = "MIN_MAX"
	NetflowTopLocationsResponseMetaNormalizationRawValues            NetflowTopLocationsResponseMetaNormalization = "RAW_VALUES"
	NetflowTopLocationsResponseMetaNormalizationPercentageChange     NetflowTopLocationsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	NetflowTopLocationsResponseMetaNormalizationRollingAverage       NetflowTopLocationsResponseMetaNormalization = "ROLLING_AVERAGE"
	NetflowTopLocationsResponseMetaNormalizationOverlappedPercentage NetflowTopLocationsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	NetflowTopLocationsResponseMetaNormalizationRatio                NetflowTopLocationsResponseMetaNormalization = "RATIO"
)

func (r NetflowTopLocationsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case NetflowTopLocationsResponseMetaNormalizationPercentage, NetflowTopLocationsResponseMetaNormalizationMin0Max, NetflowTopLocationsResponseMetaNormalizationMinMax, NetflowTopLocationsResponseMetaNormalizationRawValues, NetflowTopLocationsResponseMetaNormalizationPercentageChange, NetflowTopLocationsResponseMetaNormalizationRollingAverage, NetflowTopLocationsResponseMetaNormalizationOverlappedPercentage, NetflowTopLocationsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type NetflowTopLocationsResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  netflowTopLocationsResponseMetaUnitJSON `json:"-"`
}

// netflowTopLocationsResponseMetaUnitJSON contains the JSON metadata for the
// struct [NetflowTopLocationsResponseMetaUnit]
type netflowTopLocationsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                              `json:"value,required"`
	JSON  netflowTopLocationsResponseTop0JSON `json:"-"`
}

// netflowTopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [NetflowTopLocationsResponseTop0]
type netflowTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type NetflowTopAsesParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[NetflowTopAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowTopAsesParams]'s query parameters as `url.Values`.
func (r NetflowTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetflowTopAsesParamsFormat string

const (
	NetflowTopAsesParamsFormatJson NetflowTopAsesParamsFormat = "JSON"
	NetflowTopAsesParamsFormatCsv  NetflowTopAsesParamsFormat = "CSV"
)

func (r NetflowTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case NetflowTopAsesParamsFormatJson, NetflowTopAsesParamsFormatCsv:
		return true
	}
	return false
}

type NetflowTopAsesResponseEnvelope struct {
	Result  NetflowTopAsesResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    netflowTopAsesResponseEnvelopeJSON `json:"-"`
}

// netflowTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [NetflowTopAsesResponseEnvelope]
type netflowTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type NetflowTopLocationsParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[NetflowTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [NetflowTopLocationsParams]'s query parameters as
// `url.Values`.
func (r NetflowTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type NetflowTopLocationsParamsFormat string

const (
	NetflowTopLocationsParamsFormatJson NetflowTopLocationsParamsFormat = "JSON"
	NetflowTopLocationsParamsFormatCsv  NetflowTopLocationsParamsFormat = "CSV"
)

func (r NetflowTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case NetflowTopLocationsParamsFormatJson, NetflowTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

type NetflowTopLocationsResponseEnvelope struct {
	Result  NetflowTopLocationsResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    netflowTopLocationsResponseEnvelopeJSON `json:"-"`
}

// netflowTopLocationsResponseEnvelopeJSON contains the JSON metadata for the
// struct [NetflowTopLocationsResponseEnvelope]
type netflowTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetflowTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r netflowTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
