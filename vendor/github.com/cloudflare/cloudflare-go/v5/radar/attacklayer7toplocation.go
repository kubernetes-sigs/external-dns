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

// AttackLayer7TopLocationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackLayer7TopLocationService] method instead.
type AttackLayer7TopLocationService struct {
	Options []option.RequestOption
}

// NewAttackLayer7TopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackLayer7TopLocationService(opts ...option.RequestOption) (r *AttackLayer7TopLocationService) {
	r = &AttackLayer7TopLocationService{}
	r.Options = opts
	return
}

// Retrieves the top origin locations of layer 7 attacks. Values are percentages of
// the total layer 7 attacks, with the origin location determined by the client IP
// address.
func (r *AttackLayer7TopLocationService) Origin(ctx context.Context, query AttackLayer7TopLocationOriginParams, opts ...option.RequestOption) (res *AttackLayer7TopLocationOriginResponse, err error) {
	var env AttackLayer7TopLocationOriginResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/locations/origin"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top target locations of and by layer 7 attacks. Values are a
// percentage out of the total layer 7 attacks. The target location is determined
// by the attacked zone's billing country, when available.
func (r *AttackLayer7TopLocationService) Target(ctx context.Context, query AttackLayer7TopLocationTargetParams, opts ...option.RequestOption) (res *AttackLayer7TopLocationTargetResponse, err error) {
	var env AttackLayer7TopLocationTargetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/locations/target"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackLayer7TopLocationOriginResponse struct {
	// Metadata for the results.
	Meta AttackLayer7TopLocationOriginResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopLocationOriginResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopLocationOriginResponseJSON   `json:"-"`
}

// attackLayer7TopLocationOriginResponseJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationOriginResponse]
type attackLayer7TopLocationOriginResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TopLocationOriginResponseMeta struct {
	ConfidenceInfo AttackLayer7TopLocationOriginResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer7TopLocationOriginResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TopLocationOriginResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TopLocationOriginResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TopLocationOriginResponseMetaJSON   `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationOriginResponseMeta]
type attackLayer7TopLocationOriginResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7TopLocationOriginResponseMetaConfidenceInfo]
type attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopLocationOriginResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7TopLocationOriginResponseMetaDateRange]
type attackLayer7TopLocationOriginResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopLocationOriginResponseMetaNormalization string

const (
	AttackLayer7TopLocationOriginResponseMetaNormalizationPercentage           AttackLayer7TopLocationOriginResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TopLocationOriginResponseMetaNormalizationMin0Max              AttackLayer7TopLocationOriginResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TopLocationOriginResponseMetaNormalizationMinMax               AttackLayer7TopLocationOriginResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TopLocationOriginResponseMetaNormalizationRawValues            AttackLayer7TopLocationOriginResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TopLocationOriginResponseMetaNormalizationPercentageChange     AttackLayer7TopLocationOriginResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TopLocationOriginResponseMetaNormalizationRollingAverage       AttackLayer7TopLocationOriginResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TopLocationOriginResponseMetaNormalizationOverlappedPercentage AttackLayer7TopLocationOriginResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TopLocationOriginResponseMetaNormalizationRatio                AttackLayer7TopLocationOriginResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TopLocationOriginResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginResponseMetaNormalizationPercentage, AttackLayer7TopLocationOriginResponseMetaNormalizationMin0Max, AttackLayer7TopLocationOriginResponseMetaNormalizationMinMax, AttackLayer7TopLocationOriginResponseMetaNormalizationRawValues, AttackLayer7TopLocationOriginResponseMetaNormalizationPercentageChange, AttackLayer7TopLocationOriginResponseMetaNormalizationRollingAverage, AttackLayer7TopLocationOriginResponseMetaNormalizationOverlappedPercentage, AttackLayer7TopLocationOriginResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TopLocationOriginResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  attackLayer7TopLocationOriginResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer7TopLocationOriginResponseMetaUnit]
type attackLayer7TopLocationOriginResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginResponseTop0 struct {
	OriginCountryAlpha2 string                                        `json:"originCountryAlpha2,required"`
	OriginCountryName   string                                        `json:"originCountryName,required"`
	Rank                float64                                       `json:"rank,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer7TopLocationOriginResponseTop0JSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationOriginResponseTop0]
type attackLayer7TopLocationOriginResponseTop0JSON struct {
	OriginCountryAlpha2 apijson.Field
	OriginCountryName   apijson.Field
	Rank                apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponse struct {
	// Metadata for the results.
	Meta AttackLayer7TopLocationTargetResponseMeta   `json:"meta,required"`
	Top0 []AttackLayer7TopLocationTargetResponseTop0 `json:"top_0,required"`
	JSON attackLayer7TopLocationTargetResponseJSON   `json:"-"`
}

// attackLayer7TopLocationTargetResponseJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationTargetResponse]
type attackLayer7TopLocationTargetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AttackLayer7TopLocationTargetResponseMeta struct {
	ConfidenceInfo AttackLayer7TopLocationTargetResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []AttackLayer7TopLocationTargetResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AttackLayer7TopLocationTargetResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AttackLayer7TopLocationTargetResponseMetaUnit `json:"units,required"`
	JSON  attackLayer7TopLocationTargetResponseMetaJSON   `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaJSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationTargetResponseMeta]
type attackLayer7TopLocationTargetResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseMetaConfidenceInfo struct {
	Annotations []AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [AttackLayer7TopLocationTargetResponseMetaConfidenceInfo]
type attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation]
type attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AttackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      attackLayer7TopLocationTargetResponseMetaDateRangeJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AttackLayer7TopLocationTargetResponseMetaDateRange]
type attackLayer7TopLocationTargetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AttackLayer7TopLocationTargetResponseMetaNormalization string

const (
	AttackLayer7TopLocationTargetResponseMetaNormalizationPercentage           AttackLayer7TopLocationTargetResponseMetaNormalization = "PERCENTAGE"
	AttackLayer7TopLocationTargetResponseMetaNormalizationMin0Max              AttackLayer7TopLocationTargetResponseMetaNormalization = "MIN0_MAX"
	AttackLayer7TopLocationTargetResponseMetaNormalizationMinMax               AttackLayer7TopLocationTargetResponseMetaNormalization = "MIN_MAX"
	AttackLayer7TopLocationTargetResponseMetaNormalizationRawValues            AttackLayer7TopLocationTargetResponseMetaNormalization = "RAW_VALUES"
	AttackLayer7TopLocationTargetResponseMetaNormalizationPercentageChange     AttackLayer7TopLocationTargetResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AttackLayer7TopLocationTargetResponseMetaNormalizationRollingAverage       AttackLayer7TopLocationTargetResponseMetaNormalization = "ROLLING_AVERAGE"
	AttackLayer7TopLocationTargetResponseMetaNormalizationOverlappedPercentage AttackLayer7TopLocationTargetResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AttackLayer7TopLocationTargetResponseMetaNormalizationRatio                AttackLayer7TopLocationTargetResponseMetaNormalization = "RATIO"
)

func (r AttackLayer7TopLocationTargetResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationTargetResponseMetaNormalizationPercentage, AttackLayer7TopLocationTargetResponseMetaNormalizationMin0Max, AttackLayer7TopLocationTargetResponseMetaNormalizationMinMax, AttackLayer7TopLocationTargetResponseMetaNormalizationRawValues, AttackLayer7TopLocationTargetResponseMetaNormalizationPercentageChange, AttackLayer7TopLocationTargetResponseMetaNormalizationRollingAverage, AttackLayer7TopLocationTargetResponseMetaNormalizationOverlappedPercentage, AttackLayer7TopLocationTargetResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AttackLayer7TopLocationTargetResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  attackLayer7TopLocationTargetResponseMetaUnitJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseMetaUnitJSON contains the JSON metadata for
// the struct [AttackLayer7TopLocationTargetResponseMetaUnit]
type attackLayer7TopLocationTargetResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetResponseTop0 struct {
	Rank                float64                                       `json:"rank,required"`
	TargetCountryAlpha2 string                                        `json:"targetCountryAlpha2,required"`
	TargetCountryName   string                                        `json:"targetCountryName,required"`
	Value               string                                        `json:"value,required"`
	JSON                attackLayer7TopLocationTargetResponseTop0JSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseTop0JSON contains the JSON metadata for the
// struct [AttackLayer7TopLocationTargetResponseTop0]
type attackLayer7TopLocationTargetResponseTop0JSON struct {
	Rank                apijson.Field
	TargetCountryAlpha2 apijson.Field
	TargetCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationOriginParams struct {
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
	Format param.Field[AttackLayer7TopLocationOriginParamsFormat] `query:"format"`
	// Filters results by HTTP method.
	HTTPMethod param.Field[[]AttackLayer7TopLocationOriginParamsHTTPMethod] `query:"httpMethod"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]AttackLayer7TopLocationOriginParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]AttackLayer7TopLocationOriginParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TopLocationOriginParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopLocationOriginParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopLocationOriginParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7TopLocationOriginParamsFormat string

const (
	AttackLayer7TopLocationOriginParamsFormatJson AttackLayer7TopLocationOriginParamsFormat = "JSON"
	AttackLayer7TopLocationOriginParamsFormatCsv  AttackLayer7TopLocationOriginParamsFormat = "CSV"
)

func (r AttackLayer7TopLocationOriginParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginParamsFormatJson, AttackLayer7TopLocationOriginParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopLocationOriginParamsHTTPMethod string

const (
	AttackLayer7TopLocationOriginParamsHTTPMethodGet             AttackLayer7TopLocationOriginParamsHTTPMethod = "GET"
	AttackLayer7TopLocationOriginParamsHTTPMethodPost            AttackLayer7TopLocationOriginParamsHTTPMethod = "POST"
	AttackLayer7TopLocationOriginParamsHTTPMethodDelete          AttackLayer7TopLocationOriginParamsHTTPMethod = "DELETE"
	AttackLayer7TopLocationOriginParamsHTTPMethodPut             AttackLayer7TopLocationOriginParamsHTTPMethod = "PUT"
	AttackLayer7TopLocationOriginParamsHTTPMethodHead            AttackLayer7TopLocationOriginParamsHTTPMethod = "HEAD"
	AttackLayer7TopLocationOriginParamsHTTPMethodPurge           AttackLayer7TopLocationOriginParamsHTTPMethod = "PURGE"
	AttackLayer7TopLocationOriginParamsHTTPMethodOptions         AttackLayer7TopLocationOriginParamsHTTPMethod = "OPTIONS"
	AttackLayer7TopLocationOriginParamsHTTPMethodPropfind        AttackLayer7TopLocationOriginParamsHTTPMethod = "PROPFIND"
	AttackLayer7TopLocationOriginParamsHTTPMethodMkcol           AttackLayer7TopLocationOriginParamsHTTPMethod = "MKCOL"
	AttackLayer7TopLocationOriginParamsHTTPMethodPatch           AttackLayer7TopLocationOriginParamsHTTPMethod = "PATCH"
	AttackLayer7TopLocationOriginParamsHTTPMethodACL             AttackLayer7TopLocationOriginParamsHTTPMethod = "ACL"
	AttackLayer7TopLocationOriginParamsHTTPMethodBcopy           AttackLayer7TopLocationOriginParamsHTTPMethod = "BCOPY"
	AttackLayer7TopLocationOriginParamsHTTPMethodBdelete         AttackLayer7TopLocationOriginParamsHTTPMethod = "BDELETE"
	AttackLayer7TopLocationOriginParamsHTTPMethodBmove           AttackLayer7TopLocationOriginParamsHTTPMethod = "BMOVE"
	AttackLayer7TopLocationOriginParamsHTTPMethodBpropfind       AttackLayer7TopLocationOriginParamsHTTPMethod = "BPROPFIND"
	AttackLayer7TopLocationOriginParamsHTTPMethodBproppatch      AttackLayer7TopLocationOriginParamsHTTPMethod = "BPROPPATCH"
	AttackLayer7TopLocationOriginParamsHTTPMethodCheckin         AttackLayer7TopLocationOriginParamsHTTPMethod = "CHECKIN"
	AttackLayer7TopLocationOriginParamsHTTPMethodCheckout        AttackLayer7TopLocationOriginParamsHTTPMethod = "CHECKOUT"
	AttackLayer7TopLocationOriginParamsHTTPMethodConnect         AttackLayer7TopLocationOriginParamsHTTPMethod = "CONNECT"
	AttackLayer7TopLocationOriginParamsHTTPMethodCopy            AttackLayer7TopLocationOriginParamsHTTPMethod = "COPY"
	AttackLayer7TopLocationOriginParamsHTTPMethodLabel           AttackLayer7TopLocationOriginParamsHTTPMethod = "LABEL"
	AttackLayer7TopLocationOriginParamsHTTPMethodLock            AttackLayer7TopLocationOriginParamsHTTPMethod = "LOCK"
	AttackLayer7TopLocationOriginParamsHTTPMethodMerge           AttackLayer7TopLocationOriginParamsHTTPMethod = "MERGE"
	AttackLayer7TopLocationOriginParamsHTTPMethodMkactivity      AttackLayer7TopLocationOriginParamsHTTPMethod = "MKACTIVITY"
	AttackLayer7TopLocationOriginParamsHTTPMethodMkworkspace     AttackLayer7TopLocationOriginParamsHTTPMethod = "MKWORKSPACE"
	AttackLayer7TopLocationOriginParamsHTTPMethodMove            AttackLayer7TopLocationOriginParamsHTTPMethod = "MOVE"
	AttackLayer7TopLocationOriginParamsHTTPMethodNotify          AttackLayer7TopLocationOriginParamsHTTPMethod = "NOTIFY"
	AttackLayer7TopLocationOriginParamsHTTPMethodOrderpatch      AttackLayer7TopLocationOriginParamsHTTPMethod = "ORDERPATCH"
	AttackLayer7TopLocationOriginParamsHTTPMethodPoll            AttackLayer7TopLocationOriginParamsHTTPMethod = "POLL"
	AttackLayer7TopLocationOriginParamsHTTPMethodProppatch       AttackLayer7TopLocationOriginParamsHTTPMethod = "PROPPATCH"
	AttackLayer7TopLocationOriginParamsHTTPMethodReport          AttackLayer7TopLocationOriginParamsHTTPMethod = "REPORT"
	AttackLayer7TopLocationOriginParamsHTTPMethodSearch          AttackLayer7TopLocationOriginParamsHTTPMethod = "SEARCH"
	AttackLayer7TopLocationOriginParamsHTTPMethodSubscribe       AttackLayer7TopLocationOriginParamsHTTPMethod = "SUBSCRIBE"
	AttackLayer7TopLocationOriginParamsHTTPMethodTrace           AttackLayer7TopLocationOriginParamsHTTPMethod = "TRACE"
	AttackLayer7TopLocationOriginParamsHTTPMethodUncheckout      AttackLayer7TopLocationOriginParamsHTTPMethod = "UNCHECKOUT"
	AttackLayer7TopLocationOriginParamsHTTPMethodUnlock          AttackLayer7TopLocationOriginParamsHTTPMethod = "UNLOCK"
	AttackLayer7TopLocationOriginParamsHTTPMethodUnsubscribe     AttackLayer7TopLocationOriginParamsHTTPMethod = "UNSUBSCRIBE"
	AttackLayer7TopLocationOriginParamsHTTPMethodUpdate          AttackLayer7TopLocationOriginParamsHTTPMethod = "UPDATE"
	AttackLayer7TopLocationOriginParamsHTTPMethodVersioncontrol  AttackLayer7TopLocationOriginParamsHTTPMethod = "VERSIONCONTROL"
	AttackLayer7TopLocationOriginParamsHTTPMethodBaselinecontrol AttackLayer7TopLocationOriginParamsHTTPMethod = "BASELINECONTROL"
	AttackLayer7TopLocationOriginParamsHTTPMethodXmsenumatts     AttackLayer7TopLocationOriginParamsHTTPMethod = "XMSENUMATTS"
	AttackLayer7TopLocationOriginParamsHTTPMethodRpcOutData      AttackLayer7TopLocationOriginParamsHTTPMethod = "RPC_OUT_DATA"
	AttackLayer7TopLocationOriginParamsHTTPMethodRpcInData       AttackLayer7TopLocationOriginParamsHTTPMethod = "RPC_IN_DATA"
	AttackLayer7TopLocationOriginParamsHTTPMethodJson            AttackLayer7TopLocationOriginParamsHTTPMethod = "JSON"
	AttackLayer7TopLocationOriginParamsHTTPMethodCook            AttackLayer7TopLocationOriginParamsHTTPMethod = "COOK"
	AttackLayer7TopLocationOriginParamsHTTPMethodTrack           AttackLayer7TopLocationOriginParamsHTTPMethod = "TRACK"
)

func (r AttackLayer7TopLocationOriginParamsHTTPMethod) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginParamsHTTPMethodGet, AttackLayer7TopLocationOriginParamsHTTPMethodPost, AttackLayer7TopLocationOriginParamsHTTPMethodDelete, AttackLayer7TopLocationOriginParamsHTTPMethodPut, AttackLayer7TopLocationOriginParamsHTTPMethodHead, AttackLayer7TopLocationOriginParamsHTTPMethodPurge, AttackLayer7TopLocationOriginParamsHTTPMethodOptions, AttackLayer7TopLocationOriginParamsHTTPMethodPropfind, AttackLayer7TopLocationOriginParamsHTTPMethodMkcol, AttackLayer7TopLocationOriginParamsHTTPMethodPatch, AttackLayer7TopLocationOriginParamsHTTPMethodACL, AttackLayer7TopLocationOriginParamsHTTPMethodBcopy, AttackLayer7TopLocationOriginParamsHTTPMethodBdelete, AttackLayer7TopLocationOriginParamsHTTPMethodBmove, AttackLayer7TopLocationOriginParamsHTTPMethodBpropfind, AttackLayer7TopLocationOriginParamsHTTPMethodBproppatch, AttackLayer7TopLocationOriginParamsHTTPMethodCheckin, AttackLayer7TopLocationOriginParamsHTTPMethodCheckout, AttackLayer7TopLocationOriginParamsHTTPMethodConnect, AttackLayer7TopLocationOriginParamsHTTPMethodCopy, AttackLayer7TopLocationOriginParamsHTTPMethodLabel, AttackLayer7TopLocationOriginParamsHTTPMethodLock, AttackLayer7TopLocationOriginParamsHTTPMethodMerge, AttackLayer7TopLocationOriginParamsHTTPMethodMkactivity, AttackLayer7TopLocationOriginParamsHTTPMethodMkworkspace, AttackLayer7TopLocationOriginParamsHTTPMethodMove, AttackLayer7TopLocationOriginParamsHTTPMethodNotify, AttackLayer7TopLocationOriginParamsHTTPMethodOrderpatch, AttackLayer7TopLocationOriginParamsHTTPMethodPoll, AttackLayer7TopLocationOriginParamsHTTPMethodProppatch, AttackLayer7TopLocationOriginParamsHTTPMethodReport, AttackLayer7TopLocationOriginParamsHTTPMethodSearch, AttackLayer7TopLocationOriginParamsHTTPMethodSubscribe, AttackLayer7TopLocationOriginParamsHTTPMethodTrace, AttackLayer7TopLocationOriginParamsHTTPMethodUncheckout, AttackLayer7TopLocationOriginParamsHTTPMethodUnlock, AttackLayer7TopLocationOriginParamsHTTPMethodUnsubscribe, AttackLayer7TopLocationOriginParamsHTTPMethodUpdate, AttackLayer7TopLocationOriginParamsHTTPMethodVersioncontrol, AttackLayer7TopLocationOriginParamsHTTPMethodBaselinecontrol, AttackLayer7TopLocationOriginParamsHTTPMethodXmsenumatts, AttackLayer7TopLocationOriginParamsHTTPMethodRpcOutData, AttackLayer7TopLocationOriginParamsHTTPMethodRpcInData, AttackLayer7TopLocationOriginParamsHTTPMethodJson, AttackLayer7TopLocationOriginParamsHTTPMethodCook, AttackLayer7TopLocationOriginParamsHTTPMethodTrack:
		return true
	}
	return false
}

type AttackLayer7TopLocationOriginParamsHTTPVersion string

const (
	AttackLayer7TopLocationOriginParamsHTTPVersionHttPv1 AttackLayer7TopLocationOriginParamsHTTPVersion = "HTTPv1"
	AttackLayer7TopLocationOriginParamsHTTPVersionHttPv2 AttackLayer7TopLocationOriginParamsHTTPVersion = "HTTPv2"
	AttackLayer7TopLocationOriginParamsHTTPVersionHttPv3 AttackLayer7TopLocationOriginParamsHTTPVersion = "HTTPv3"
)

func (r AttackLayer7TopLocationOriginParamsHTTPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginParamsHTTPVersionHttPv1, AttackLayer7TopLocationOriginParamsHTTPVersionHttPv2, AttackLayer7TopLocationOriginParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type AttackLayer7TopLocationOriginParamsIPVersion string

const (
	AttackLayer7TopLocationOriginParamsIPVersionIPv4 AttackLayer7TopLocationOriginParamsIPVersion = "IPv4"
	AttackLayer7TopLocationOriginParamsIPVersionIPv6 AttackLayer7TopLocationOriginParamsIPVersion = "IPv6"
)

func (r AttackLayer7TopLocationOriginParamsIPVersion) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginParamsIPVersionIPv4, AttackLayer7TopLocationOriginParamsIPVersionIPv6:
		return true
	}
	return false
}

type AttackLayer7TopLocationOriginParamsMitigationProduct string

const (
	AttackLayer7TopLocationOriginParamsMitigationProductDDoS               AttackLayer7TopLocationOriginParamsMitigationProduct = "DDOS"
	AttackLayer7TopLocationOriginParamsMitigationProductWAF                AttackLayer7TopLocationOriginParamsMitigationProduct = "WAF"
	AttackLayer7TopLocationOriginParamsMitigationProductBotManagement      AttackLayer7TopLocationOriginParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TopLocationOriginParamsMitigationProductAccessRules        AttackLayer7TopLocationOriginParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TopLocationOriginParamsMitigationProductIPReputation       AttackLayer7TopLocationOriginParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TopLocationOriginParamsMitigationProductAPIShield          AttackLayer7TopLocationOriginParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TopLocationOriginParamsMitigationProductDataLossPrevention AttackLayer7TopLocationOriginParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TopLocationOriginParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationOriginParamsMitigationProductDDoS, AttackLayer7TopLocationOriginParamsMitigationProductWAF, AttackLayer7TopLocationOriginParamsMitigationProductBotManagement, AttackLayer7TopLocationOriginParamsMitigationProductAccessRules, AttackLayer7TopLocationOriginParamsMitigationProductIPReputation, AttackLayer7TopLocationOriginParamsMitigationProductAPIShield, AttackLayer7TopLocationOriginParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7TopLocationOriginResponseEnvelope struct {
	Result  AttackLayer7TopLocationOriginResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer7TopLocationOriginResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopLocationOriginResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7TopLocationOriginResponseEnvelope]
type attackLayer7TopLocationOriginResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationOriginResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationOriginResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AttackLayer7TopLocationTargetParams struct {
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
	Format param.Field[AttackLayer7TopLocationTargetParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters the results by layer 7 mitigation product.
	MitigationProduct param.Field[[]AttackLayer7TopLocationTargetParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AttackLayer7TopLocationTargetParams]'s query parameters as
// `url.Values`.
func (r AttackLayer7TopLocationTargetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AttackLayer7TopLocationTargetParamsFormat string

const (
	AttackLayer7TopLocationTargetParamsFormatJson AttackLayer7TopLocationTargetParamsFormat = "JSON"
	AttackLayer7TopLocationTargetParamsFormatCsv  AttackLayer7TopLocationTargetParamsFormat = "CSV"
)

func (r AttackLayer7TopLocationTargetParamsFormat) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationTargetParamsFormatJson, AttackLayer7TopLocationTargetParamsFormatCsv:
		return true
	}
	return false
}

type AttackLayer7TopLocationTargetParamsMitigationProduct string

const (
	AttackLayer7TopLocationTargetParamsMitigationProductDDoS               AttackLayer7TopLocationTargetParamsMitigationProduct = "DDOS"
	AttackLayer7TopLocationTargetParamsMitigationProductWAF                AttackLayer7TopLocationTargetParamsMitigationProduct = "WAF"
	AttackLayer7TopLocationTargetParamsMitigationProductBotManagement      AttackLayer7TopLocationTargetParamsMitigationProduct = "BOT_MANAGEMENT"
	AttackLayer7TopLocationTargetParamsMitigationProductAccessRules        AttackLayer7TopLocationTargetParamsMitigationProduct = "ACCESS_RULES"
	AttackLayer7TopLocationTargetParamsMitigationProductIPReputation       AttackLayer7TopLocationTargetParamsMitigationProduct = "IP_REPUTATION"
	AttackLayer7TopLocationTargetParamsMitigationProductAPIShield          AttackLayer7TopLocationTargetParamsMitigationProduct = "API_SHIELD"
	AttackLayer7TopLocationTargetParamsMitigationProductDataLossPrevention AttackLayer7TopLocationTargetParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

func (r AttackLayer7TopLocationTargetParamsMitigationProduct) IsKnown() bool {
	switch r {
	case AttackLayer7TopLocationTargetParamsMitigationProductDDoS, AttackLayer7TopLocationTargetParamsMitigationProductWAF, AttackLayer7TopLocationTargetParamsMitigationProductBotManagement, AttackLayer7TopLocationTargetParamsMitigationProductAccessRules, AttackLayer7TopLocationTargetParamsMitigationProductIPReputation, AttackLayer7TopLocationTargetParamsMitigationProductAPIShield, AttackLayer7TopLocationTargetParamsMitigationProductDataLossPrevention:
		return true
	}
	return false
}

type AttackLayer7TopLocationTargetResponseEnvelope struct {
	Result  AttackLayer7TopLocationTargetResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    attackLayer7TopLocationTargetResponseEnvelopeJSON `json:"-"`
}

// attackLayer7TopLocationTargetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AttackLayer7TopLocationTargetResponseEnvelope]
type attackLayer7TopLocationTargetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackLayer7TopLocationTargetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackLayer7TopLocationTargetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
