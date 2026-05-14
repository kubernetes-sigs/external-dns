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

// DNSTopService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSTopService] method instead.
type DNSTopService struct {
	Options []option.RequestOption
}

// NewDNSTopService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDNSTopService(opts ...option.RequestOption) (r *DNSTopService) {
	r = &DNSTopService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems by DNS queries made to 1.1.1.1 DNS
// resolver.
func (r *DNSTopService) Ases(ctx context.Context, query DNSTopAsesParams, opts ...option.RequestOption) (res *DNSTopAsesResponse, err error) {
	var env DNSTopAsesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top locations by DNS queries made to 1.1.1.1 DNS resolver.
func (r *DNSTopService) Locations(ctx context.Context, query DNSTopLocationsParams, opts ...option.RequestOption) (res *DNSTopLocationsResponse, err error) {
	var env DNSTopLocationsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/dns/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSTopAsesResponse struct {
	// Metadata for the results.
	Meta DNSTopAsesResponseMeta   `json:"meta,required"`
	Top0 []DNSTopAsesResponseTop0 `json:"top_0,required"`
	JSON dnsTopAsesResponseJSON   `json:"-"`
}

// dnsTopAsesResponseJSON contains the JSON metadata for the struct
// [DNSTopAsesResponse]
type dnsTopAsesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTopAsesResponseMeta struct {
	ConfidenceInfo DNSTopAsesResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []DNSTopAsesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTopAsesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSTopAsesResponseMetaUnit `json:"units,required"`
	JSON  dnsTopAsesResponseMetaJSON   `json:"-"`
}

// dnsTopAsesResponseMetaJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseMeta]
type dnsTopAsesResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTopAsesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSTopAsesResponseMetaConfidenceInfo struct {
	Annotations []DNSTopAsesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                    `json:"level,required"`
	JSON  dnsTopAsesResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTopAsesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTopAsesResponseMetaConfidenceInfo]
type dnsTopAsesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTopAsesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                          `json:"startDate,required" format:"date-time"`
	JSON            dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [DNSTopAsesResponseMetaConfidenceInfoAnnotation]
type dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSTopAsesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSTopAsesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                           `json:"startTime,required" format:"date-time"`
	JSON      dnsTopAsesResponseMetaDateRangeJSON `json:"-"`
}

// dnsTopAsesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseMetaDateRange]
type dnsTopAsesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTopAsesResponseMetaNormalization string

const (
	DNSTopAsesResponseMetaNormalizationPercentage           DNSTopAsesResponseMetaNormalization = "PERCENTAGE"
	DNSTopAsesResponseMetaNormalizationMin0Max              DNSTopAsesResponseMetaNormalization = "MIN0_MAX"
	DNSTopAsesResponseMetaNormalizationMinMax               DNSTopAsesResponseMetaNormalization = "MIN_MAX"
	DNSTopAsesResponseMetaNormalizationRawValues            DNSTopAsesResponseMetaNormalization = "RAW_VALUES"
	DNSTopAsesResponseMetaNormalizationPercentageChange     DNSTopAsesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTopAsesResponseMetaNormalizationRollingAverage       DNSTopAsesResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTopAsesResponseMetaNormalizationOverlappedPercentage DNSTopAsesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTopAsesResponseMetaNormalizationRatio                DNSTopAsesResponseMetaNormalization = "RATIO"
)

func (r DNSTopAsesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTopAsesResponseMetaNormalizationPercentage, DNSTopAsesResponseMetaNormalizationMin0Max, DNSTopAsesResponseMetaNormalizationMinMax, DNSTopAsesResponseMetaNormalizationRawValues, DNSTopAsesResponseMetaNormalizationPercentageChange, DNSTopAsesResponseMetaNormalizationRollingAverage, DNSTopAsesResponseMetaNormalizationOverlappedPercentage, DNSTopAsesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTopAsesResponseMetaUnit struct {
	Name  string                         `json:"name,required"`
	Value string                         `json:"value,required"`
	JSON  dnsTopAsesResponseMetaUnitJSON `json:"-"`
}

// dnsTopAsesResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseMetaUnit]
type dnsTopAsesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTopAsesResponseTop0 struct {
	ClientASN    int64  `json:"clientASN,required"`
	ClientAsName string `json:"clientASName,required"`
	// A numeric string.
	Value string                     `json:"value,required"`
	JSON  dnsTopAsesResponseTop0JSON `json:"-"`
}

// dnsTopAsesResponseTop0JSON contains the JSON metadata for the struct
// [DNSTopAsesResponseTop0]
type dnsTopAsesResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DNSTopAsesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseTop0JSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsResponse struct {
	// Metadata for the results.
	Meta DNSTopLocationsResponseMeta   `json:"meta,required"`
	Top0 []DNSTopLocationsResponseTop0 `json:"top_0,required"`
	JSON dnsTopLocationsResponseJSON   `json:"-"`
}

// dnsTopLocationsResponseJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponse]
type dnsTopLocationsResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type DNSTopLocationsResponseMeta struct {
	ConfidenceInfo DNSTopLocationsResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []DNSTopLocationsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization DNSTopLocationsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []DNSTopLocationsResponseMetaUnit `json:"units,required"`
	JSON  dnsTopLocationsResponseMetaJSON   `json:"-"`
}

// dnsTopLocationsResponseMetaJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseMeta]
type dnsTopLocationsResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaJSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsResponseMetaConfidenceInfo struct {
	Annotations []DNSTopLocationsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                         `json:"level,required"`
	JSON  dnsTopLocationsResponseMetaConfidenceInfoJSON `json:"-"`
}

// dnsTopLocationsResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [DNSTopLocationsResponseMetaConfidenceInfo]
type dnsTopLocationsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type DNSTopLocationsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	LinkedURL       string                                                  `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                               `json:"startDate,required" format:"date-time"`
	JSON            dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [DNSTopLocationsResponseMetaConfidenceInfoAnnotation]
type dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *DNSTopLocationsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      dnsTopLocationsResponseMetaDateRangeJSON `json:"-"`
}

// dnsTopLocationsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [DNSTopLocationsResponseMetaDateRange]
type dnsTopLocationsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type DNSTopLocationsResponseMetaNormalization string

const (
	DNSTopLocationsResponseMetaNormalizationPercentage           DNSTopLocationsResponseMetaNormalization = "PERCENTAGE"
	DNSTopLocationsResponseMetaNormalizationMin0Max              DNSTopLocationsResponseMetaNormalization = "MIN0_MAX"
	DNSTopLocationsResponseMetaNormalizationMinMax               DNSTopLocationsResponseMetaNormalization = "MIN_MAX"
	DNSTopLocationsResponseMetaNormalizationRawValues            DNSTopLocationsResponseMetaNormalization = "RAW_VALUES"
	DNSTopLocationsResponseMetaNormalizationPercentageChange     DNSTopLocationsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	DNSTopLocationsResponseMetaNormalizationRollingAverage       DNSTopLocationsResponseMetaNormalization = "ROLLING_AVERAGE"
	DNSTopLocationsResponseMetaNormalizationOverlappedPercentage DNSTopLocationsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	DNSTopLocationsResponseMetaNormalizationRatio                DNSTopLocationsResponseMetaNormalization = "RATIO"
)

func (r DNSTopLocationsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case DNSTopLocationsResponseMetaNormalizationPercentage, DNSTopLocationsResponseMetaNormalizationMin0Max, DNSTopLocationsResponseMetaNormalizationMinMax, DNSTopLocationsResponseMetaNormalizationRawValues, DNSTopLocationsResponseMetaNormalizationPercentageChange, DNSTopLocationsResponseMetaNormalizationRollingAverage, DNSTopLocationsResponseMetaNormalizationOverlappedPercentage, DNSTopLocationsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type DNSTopLocationsResponseMetaUnit struct {
	Name  string                              `json:"name,required"`
	Value string                              `json:"value,required"`
	JSON  dnsTopLocationsResponseMetaUnitJSON `json:"-"`
}

// dnsTopLocationsResponseMetaUnitJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseMetaUnit]
type dnsTopLocationsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsResponseTop0 struct {
	ClientCountryAlpha2 string `json:"clientCountryAlpha2,required"`
	ClientCountryName   string `json:"clientCountryName,required"`
	// A numeric string.
	Value string                          `json:"value,required"`
	JSON  dnsTopLocationsResponseTop0JSON `json:"-"`
}

// dnsTopLocationsResponseTop0JSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseTop0]
type dnsTopLocationsResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DNSTopLocationsResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseTop0JSON) RawJSON() string {
	return r.raw
}

type DNSTopAsesParams struct {
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
	// Filters results by domain name.
	Domain param.Field[[]string] `query:"domain"`
	// Format in which results will be returned.
	Format param.Field[DNSTopAsesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [DNSTopAsesParams]'s query parameters as `url.Values`.
func (r DNSTopAsesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSTopAsesParamsFormat string

const (
	DNSTopAsesParamsFormatJson DNSTopAsesParamsFormat = "JSON"
	DNSTopAsesParamsFormatCsv  DNSTopAsesParamsFormat = "CSV"
)

func (r DNSTopAsesParamsFormat) IsKnown() bool {
	switch r {
	case DNSTopAsesParamsFormatJson, DNSTopAsesParamsFormatCsv:
		return true
	}
	return false
}

type DNSTopAsesResponseEnvelope struct {
	Result  DNSTopAsesResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    dnsTopAsesResponseEnvelopeJSON `json:"-"`
}

// dnsTopAsesResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTopAsesResponseEnvelope]
type dnsTopAsesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopAsesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopAsesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSTopLocationsParams struct {
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
	// Filters results by domain name.
	Domain param.Field[[]string] `query:"domain"`
	// Format in which results will be returned.
	Format param.Field[DNSTopLocationsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [DNSTopLocationsParams]'s query parameters as `url.Values`.
func (r DNSTopLocationsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type DNSTopLocationsParamsFormat string

const (
	DNSTopLocationsParamsFormatJson DNSTopLocationsParamsFormat = "JSON"
	DNSTopLocationsParamsFormatCsv  DNSTopLocationsParamsFormat = "CSV"
)

func (r DNSTopLocationsParamsFormat) IsKnown() bool {
	switch r {
	case DNSTopLocationsParamsFormatJson, DNSTopLocationsParamsFormatCsv:
		return true
	}
	return false
}

type DNSTopLocationsResponseEnvelope struct {
	Result  DNSTopLocationsResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    dnsTopLocationsResponseEnvelopeJSON `json:"-"`
}

// dnsTopLocationsResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSTopLocationsResponseEnvelope]
type dnsTopLocationsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSTopLocationsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsTopLocationsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
