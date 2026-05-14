// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// CtService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCtService] method instead.
type CtService struct {
	Options     []option.RequestOption
	Authorities *CtAuthorityService
	Logs        *CtLogService
}

// NewCtService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCtService(opts ...option.RequestOption) (r *CtService) {
	r = &CtService{}
	r.Options = opts
	r.Authorities = NewCtAuthorityService(opts...)
	r.Logs = NewCtLogService(opts...)
	return
}

// Retrieves an aggregated summary of certificates grouped by the specified
// dimension.
func (r *CtService) Summary(ctx context.Context, dimension CtSummaryParamsDimension, query CtSummaryParams, opts ...option.RequestOption) (res *CtSummaryResponse, err error) {
	var env CtSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ct/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves certificate volume over time.
func (r *CtService) Timeseries(ctx context.Context, query CtTimeseriesParams, opts ...option.RequestOption) (res *CtTimeseriesResponse, err error) {
	var env CtTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ct/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of certificates grouped by chosen the specified
// dimension over time.
func (r *CtService) TimeseriesGroups(ctx context.Context, dimension CtTimeseriesGroupsParamsDimension, query CtTimeseriesGroupsParams, opts ...option.RequestOption) (res *CtTimeseriesGroupsResponse, err error) {
	var env CtTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ct/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type CtSummaryResponse struct {
	// Metadata for the results.
	Meta     CtSummaryResponseMeta          `json:"meta,required"`
	Summary0 CtSummaryResponseSummary0Union `json:"summary_0,required"`
	JSON     ctSummaryResponseJSON          `json:"-"`
}

// ctSummaryResponseJSON contains the JSON metadata for the struct
// [CtSummaryResponse]
type ctSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type CtSummaryResponseMeta struct {
	ConfidenceInfo CtSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []CtSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization CtSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []CtSummaryResponseMetaUnit `json:"units,required"`
	JSON  ctSummaryResponseMetaJSON   `json:"-"`
}

// ctSummaryResponseMetaJSON contains the JSON metadata for the struct
// [CtSummaryResponseMeta]
type ctSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CtSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type CtSummaryResponseMetaConfidenceInfo struct {
	Annotations []CtSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                   `json:"level,required"`
	JSON  ctSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// ctSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [CtSummaryResponseMetaConfidenceInfo]
type ctSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type CtSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                         `json:"startDate,required" format:"date-time"`
	JSON            ctSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// ctSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata for
// the struct [CtSummaryResponseMetaConfidenceInfoAnnotation]
type ctSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *CtSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type CtSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                          `json:"startTime,required" format:"date-time"`
	JSON      ctSummaryResponseMetaDateRangeJSON `json:"-"`
}

// ctSummaryResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [CtSummaryResponseMetaDateRange]
type ctSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CtSummaryResponseMetaNormalization string

const (
	CtSummaryResponseMetaNormalizationPercentage           CtSummaryResponseMetaNormalization = "PERCENTAGE"
	CtSummaryResponseMetaNormalizationMin0Max              CtSummaryResponseMetaNormalization = "MIN0_MAX"
	CtSummaryResponseMetaNormalizationMinMax               CtSummaryResponseMetaNormalization = "MIN_MAX"
	CtSummaryResponseMetaNormalizationRawValues            CtSummaryResponseMetaNormalization = "RAW_VALUES"
	CtSummaryResponseMetaNormalizationPercentageChange     CtSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	CtSummaryResponseMetaNormalizationRollingAverage       CtSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	CtSummaryResponseMetaNormalizationOverlappedPercentage CtSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	CtSummaryResponseMetaNormalizationRatio                CtSummaryResponseMetaNormalization = "RATIO"
)

func (r CtSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case CtSummaryResponseMetaNormalizationPercentage, CtSummaryResponseMetaNormalizationMin0Max, CtSummaryResponseMetaNormalizationMinMax, CtSummaryResponseMetaNormalizationRawValues, CtSummaryResponseMetaNormalizationPercentageChange, CtSummaryResponseMetaNormalizationRollingAverage, CtSummaryResponseMetaNormalizationOverlappedPercentage, CtSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type CtSummaryResponseMetaUnit struct {
	Name  string                        `json:"name,required"`
	Value string                        `json:"value,required"`
	JSON  ctSummaryResponseMetaUnitJSON `json:"-"`
}

// ctSummaryResponseMetaUnitJSON contains the JSON metadata for the struct
// [CtSummaryResponseMetaUnit]
type ctSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [CtSummaryResponseSummary0Map],
// [CtSummaryResponseSummary0Object], [CtSummaryResponseSummary0Object],
// [CtSummaryResponseSummary0Object], [CtSummaryResponseSummary0Object],
// [CtSummaryResponseSummary0Object], [CtSummaryResponseSummary0Object] or
// [CtSummaryResponseSummary0Object].
type CtSummaryResponseSummary0Union interface {
	implementsCtSummaryResponseSummary0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CtSummaryResponseSummary0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Map{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtSummaryResponseSummary0Object{}),
		},
	)
}

type CtSummaryResponseSummary0Map map[string]string

func (r CtSummaryResponseSummary0Map) implementsCtSummaryResponseSummary0Union() {}

type CtSummaryResponseSummary0Object struct {
	Rfc6962 string                              `json:"rfc6962,required"`
	Static  string                              `json:"static,required"`
	JSON    ctSummaryResponseSummary0ObjectJSON `json:"-"`
}

// ctSummaryResponseSummary0ObjectJSON contains the JSON metadata for the struct
// [CtSummaryResponseSummary0Object]
type ctSummaryResponseSummary0ObjectJSON struct {
	Rfc6962     apijson.Field
	Static      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtSummaryResponseSummary0Object) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseSummary0ObjectJSON) RawJSON() string {
	return r.raw
}

func (r CtSummaryResponseSummary0Object) implementsCtSummaryResponseSummary0Union() {}

type CtTimeseriesResponse struct {
	// Metadata for the results.
	Meta        CtTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]CtTimeseriesResponse `json:"-,extras"`
	JSON        ctTimeseriesResponseJSON        `json:"-"`
}

// ctTimeseriesResponseJSON contains the JSON metadata for the struct
// [CtTimeseriesResponse]
type ctTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type CtTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    CtTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo CtTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []CtTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization CtTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []CtTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  ctTimeseriesResponseMetaJSON   `json:"-"`
}

// ctTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [CtTimeseriesResponseMeta]
type ctTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CtTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CtTimeseriesResponseMetaAggInterval string

const (
	CtTimeseriesResponseMetaAggIntervalFifteenMinutes CtTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	CtTimeseriesResponseMetaAggIntervalOneHour        CtTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	CtTimeseriesResponseMetaAggIntervalOneDay         CtTimeseriesResponseMetaAggInterval = "ONE_DAY"
	CtTimeseriesResponseMetaAggIntervalOneWeek        CtTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	CtTimeseriesResponseMetaAggIntervalOneMonth       CtTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r CtTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case CtTimeseriesResponseMetaAggIntervalFifteenMinutes, CtTimeseriesResponseMetaAggIntervalOneHour, CtTimeseriesResponseMetaAggIntervalOneDay, CtTimeseriesResponseMetaAggIntervalOneWeek, CtTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type CtTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []CtTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                      `json:"level,required"`
	JSON  ctTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// ctTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [CtTimeseriesResponseMetaConfidenceInfo]
type ctTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type CtTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                 `json:"isInstantaneous,required"`
	LinkedURL       string                                               `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                            `json:"startDate,required" format:"date-time"`
	JSON            ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [CtTimeseriesResponseMetaConfidenceInfoAnnotation]
type ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *CtTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type CtTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                             `json:"startTime,required" format:"date-time"`
	JSON      ctTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// ctTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [CtTimeseriesResponseMetaDateRange]
type ctTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CtTimeseriesResponseMetaNormalization string

const (
	CtTimeseriesResponseMetaNormalizationPercentage           CtTimeseriesResponseMetaNormalization = "PERCENTAGE"
	CtTimeseriesResponseMetaNormalizationMin0Max              CtTimeseriesResponseMetaNormalization = "MIN0_MAX"
	CtTimeseriesResponseMetaNormalizationMinMax               CtTimeseriesResponseMetaNormalization = "MIN_MAX"
	CtTimeseriesResponseMetaNormalizationRawValues            CtTimeseriesResponseMetaNormalization = "RAW_VALUES"
	CtTimeseriesResponseMetaNormalizationPercentageChange     CtTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	CtTimeseriesResponseMetaNormalizationRollingAverage       CtTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	CtTimeseriesResponseMetaNormalizationOverlappedPercentage CtTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	CtTimeseriesResponseMetaNormalizationRatio                CtTimeseriesResponseMetaNormalization = "RATIO"
)

func (r CtTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case CtTimeseriesResponseMetaNormalizationPercentage, CtTimeseriesResponseMetaNormalizationMin0Max, CtTimeseriesResponseMetaNormalizationMinMax, CtTimeseriesResponseMetaNormalizationRawValues, CtTimeseriesResponseMetaNormalizationPercentageChange, CtTimeseriesResponseMetaNormalizationRollingAverage, CtTimeseriesResponseMetaNormalizationOverlappedPercentage, CtTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type CtTimeseriesResponseMetaUnit struct {
	Name  string                           `json:"name,required"`
	Value string                           `json:"value,required"`
	JSON  ctTimeseriesResponseMetaUnitJSON `json:"-"`
}

// ctTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [CtTimeseriesResponseMetaUnit]
type ctTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type CtTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   CtTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 CtTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   ctTimeseriesGroupsResponseJSON   `json:"-"`
}

// ctTimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [CtTimeseriesGroupsResponse]
type ctTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type CtTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    CtTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo CtTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []CtTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization CtTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []CtTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  ctTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// ctTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the struct
// [CtTimeseriesGroupsResponseMeta]
type ctTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CtTimeseriesGroupsResponseMetaAggInterval string

const (
	CtTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes CtTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	CtTimeseriesGroupsResponseMetaAggIntervalOneHour        CtTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	CtTimeseriesGroupsResponseMetaAggIntervalOneDay         CtTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	CtTimeseriesGroupsResponseMetaAggIntervalOneWeek        CtTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	CtTimeseriesGroupsResponseMetaAggIntervalOneMonth       CtTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r CtTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, CtTimeseriesGroupsResponseMetaAggIntervalOneHour, CtTimeseriesGroupsResponseMetaAggIntervalOneDay, CtTimeseriesGroupsResponseMetaAggIntervalOneWeek, CtTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type CtTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []CtTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                            `json:"level,required"`
	JSON  ctTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [CtTimeseriesGroupsResponseMetaConfidenceInfo]
type ctTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type CtTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	LinkedURL       string                                                     `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                  `json:"startDate,required" format:"date-time"`
	JSON            ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [CtTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *CtTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type CtTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      ctTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [CtTimeseriesGroupsResponseMetaDateRange]
type ctTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CtTimeseriesGroupsResponseMetaNormalization string

const (
	CtTimeseriesGroupsResponseMetaNormalizationPercentage           CtTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	CtTimeseriesGroupsResponseMetaNormalizationMin0Max              CtTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	CtTimeseriesGroupsResponseMetaNormalizationMinMax               CtTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	CtTimeseriesGroupsResponseMetaNormalizationRawValues            CtTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	CtTimeseriesGroupsResponseMetaNormalizationPercentageChange     CtTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	CtTimeseriesGroupsResponseMetaNormalizationRollingAverage       CtTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	CtTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage CtTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	CtTimeseriesGroupsResponseMetaNormalizationRatio                CtTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r CtTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsResponseMetaNormalizationPercentage, CtTimeseriesGroupsResponseMetaNormalizationMin0Max, CtTimeseriesGroupsResponseMetaNormalizationMinMax, CtTimeseriesGroupsResponseMetaNormalizationRawValues, CtTimeseriesGroupsResponseMetaNormalizationPercentageChange, CtTimeseriesGroupsResponseMetaNormalizationRollingAverage, CtTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, CtTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type CtTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                 `json:"name,required"`
	Value string                                 `json:"value,required"`
	JSON  ctTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// ctTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for the struct
// [CtTimeseriesGroupsResponseMetaUnit]
type ctTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type CtTimeseriesGroupsResponseSerie0 struct {
	// This field can have the runtime type of [[]string].
	Certificate interface{} `json:"CERTIFICATE"`
	// This field can have the runtime type of [[]string].
	Domain interface{} `json:"domain"`
	// This field can have the runtime type of [[]string].
	Dsa interface{} `json:"DSA"`
	// This field can have the runtime type of [[]string].
	Ecdsa interface{} `json:"ECDSA"`
	// This field can have the runtime type of [[]string].
	Expired interface{} `json:"EXPIRED"`
	// This field can have the runtime type of [[]string].
	Extended interface{} `json:"extended"`
	// This field can have the runtime type of [[]string].
	Gt121d interface{} `json:"gt_121d"`
	// This field can have the runtime type of [[]string].
	Gt16dLte31d interface{} `json:"gt_16d_lte_31d"`
	// This field can have the runtime type of [[]string].
	Gt31dLte91d interface{} `json:"gt_31d_lte_91d"`
	// This field can have the runtime type of [[]string].
	Gt3dLte16d interface{} `json:"gt_3d_lte_16d"`
	// This field can have the runtime type of [[]string].
	Gt91dLte121d interface{} `json:"gt_91d_lte_121d"`
	// This field can have the runtime type of [[]string].
	Lte3d interface{} `json:"lte_3d"`
	// This field can have the runtime type of [[]string].
	Negative interface{} `json:"NEGATIVE"`
	// This field can have the runtime type of [[]string].
	Organization interface{} `json:"organization"`
	// This field can have the runtime type of [[]string].
	Positive interface{} `json:"POSITIVE"`
	// This field can have the runtime type of [[]string].
	Precertificate interface{} `json:"PRECERTIFICATE"`
	// This field can have the runtime type of [[]string].
	Rfc6962 interface{} `json:"rfc6962"`
	// This field can have the runtime type of [[]string].
	RSA interface{} `json:"RSA"`
	// This field can have the runtime type of [[]string].
	Static interface{} `json:"static"`
	// This field can have the runtime type of [[]time.Time].
	Timestamps interface{} `json:"timestamps"`
	// This field can have the runtime type of [[]string].
	Unknown interface{} `json:"unknown"`
	// This field can have the runtime type of [[]string].
	Valid interface{}                          `json:"VALID"`
	JSON  ctTimeseriesGroupsResponseSerie0JSON `json:"-"`
	union CtTimeseriesGroupsResponseSerie0Union
}

// ctTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the struct
// [CtTimeseriesGroupsResponseSerie0]
type ctTimeseriesGroupsResponseSerie0JSON struct {
	Certificate    apijson.Field
	Domain         apijson.Field
	Dsa            apijson.Field
	Ecdsa          apijson.Field
	Expired        apijson.Field
	Extended       apijson.Field
	Gt121d         apijson.Field
	Gt16dLte31d    apijson.Field
	Gt31dLte91d    apijson.Field
	Gt3dLte16d     apijson.Field
	Gt91dLte121d   apijson.Field
	Lte3d          apijson.Field
	Negative       apijson.Field
	Organization   apijson.Field
	Positive       apijson.Field
	Precertificate apijson.Field
	Rfc6962        apijson.Field
	RSA            apijson.Field
	Static         apijson.Field
	Timestamps     apijson.Field
	Unknown        apijson.Field
	Valid          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r ctTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

func (r *CtTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	*r = CtTimeseriesGroupsResponseSerie0{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [CtTimeseriesGroupsResponseSerie0Union] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [CtTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object].
func (r CtTimeseriesGroupsResponseSerie0) AsUnion() CtTimeseriesGroupsResponseSerie0Union {
	return r.union
}

// Union satisfied by
// [CtTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object],
// [CtTimeseriesGroupsResponseSerie0Object] or
// [CtTimeseriesGroupsResponseSerie0Object].
type CtTimeseriesGroupsResponseSerie0Union interface {
	implementsCtTimeseriesGroupsResponseSerie0()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*CtTimeseriesGroupsResponseSerie0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0Object{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CtTimeseriesGroupsResponseSerie0Object{}),
		},
	)
}

type CtTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55 struct {
	Timestamps  []time.Time                                                                          `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                                                  `json:"-,extras"`
	JSON        ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON `json:"-"`
}

// ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON
// contains the JSON metadata for the struct
// [CtTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55]
type ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55JSON) RawJSON() string {
	return r.raw
}

func (r CtTimeseriesGroupsResponseSerie0UnnamedSchemaRef7826220e105d84352ba1108d9ed88e55) implementsCtTimeseriesGroupsResponseSerie0() {
}

type CtTimeseriesGroupsResponseSerie0Object struct {
	Rfc6962 []string                                   `json:"rfc6962,required"`
	Static  []string                                   `json:"static,required"`
	JSON    ctTimeseriesGroupsResponseSerie0ObjectJSON `json:"-"`
}

// ctTimeseriesGroupsResponseSerie0ObjectJSON contains the JSON metadata for the
// struct [CtTimeseriesGroupsResponseSerie0Object]
type ctTimeseriesGroupsResponseSerie0ObjectJSON struct {
	Rfc6962     apijson.Field
	Static      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponseSerie0Object) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseSerie0ObjectJSON) RawJSON() string {
	return r.raw
}

func (r CtTimeseriesGroupsResponseSerie0Object) implementsCtTimeseriesGroupsResponseSerie0() {}

type CtSummaryParams struct {
	// Filters results by certificate authority.
	CA param.Field[[]string] `query:"ca"`
	// Filters results by certificate authority owner.
	CAOwner param.Field[[]string] `query:"caOwner"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by certificate duration.
	Duration param.Field[[]CtSummaryParamsDuration] `query:"duration"`
	// Filters results by entry type (certificate vs. pre-certificate).
	EntryType param.Field[[]CtSummaryParamsEntryType] `query:"entryType"`
	// Filters results by expiration status (expired vs. valid).
	ExpirationStatus param.Field[[]CtSummaryParamsExpirationStatus] `query:"expirationStatus"`
	// Format in which results will be returned.
	Format param.Field[CtSummaryParamsFormat] `query:"format"`
	// Filters results based on whether the certificates are bound to specific IP
	// addresses.
	HasIPs param.Field[[]bool] `query:"hasIps"`
	// Filters results based on whether the certificates contain wildcard domains.
	HasWildcards param.Field[[]bool] `query:"hasWildcards"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by certificate log.
	Log param.Field[[]string] `query:"log"`
	// Filters results by certificate log API (RFC6962 vs. static).
	LogAPI param.Field[[]CtSummaryParamsLogAPI] `query:"logApi"`
	// Filters results by certificate log operator.
	LogOperator param.Field[[]string] `query:"logOperator"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[CtSummaryParamsNormalization] `query:"normalization"`
	// Filters results by public key algorithm.
	PublicKeyAlgorithm param.Field[[]CtSummaryParamsPublicKeyAlgorithm] `query:"publicKeyAlgorithm"`
	// Filters results by signature algorithm.
	SignatureAlgorithm param.Field[[]CtSummaryParamsSignatureAlgorithm] `query:"signatureAlgorithm"`
	// Filters results by top-level domain.
	Tld param.Field[[]string] `query:"tld"`
	// Specifies whether to filter out duplicate certificates and pre-certificates. Set
	// to true for unique entries only.
	UniqueEntries param.Field[[]CtSummaryParamsUniqueEntry] `query:"uniqueEntries"`
	// Filters results by validation level.
	ValidationLevel param.Field[[]CtSummaryParamsValidationLevel] `query:"validationLevel"`
}

// URLQuery serializes [CtSummaryParams]'s query parameters as `url.Values`.
func (r CtSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the certificate attribute by which to group the results.
type CtSummaryParamsDimension string

const (
	CtSummaryParamsDimensionCA                 CtSummaryParamsDimension = "CA"
	CtSummaryParamsDimensionCAOwner            CtSummaryParamsDimension = "CA_OWNER"
	CtSummaryParamsDimensionDuration           CtSummaryParamsDimension = "DURATION"
	CtSummaryParamsDimensionEntryType          CtSummaryParamsDimension = "ENTRY_TYPE"
	CtSummaryParamsDimensionExpirationStatus   CtSummaryParamsDimension = "EXPIRATION_STATUS"
	CtSummaryParamsDimensionHasIPs             CtSummaryParamsDimension = "HAS_IPS"
	CtSummaryParamsDimensionHasWildcards       CtSummaryParamsDimension = "HAS_WILDCARDS"
	CtSummaryParamsDimensionLog                CtSummaryParamsDimension = "LOG"
	CtSummaryParamsDimensionLogAPI             CtSummaryParamsDimension = "LOG_API"
	CtSummaryParamsDimensionLogOperator        CtSummaryParamsDimension = "LOG_OPERATOR"
	CtSummaryParamsDimensionPublicKeyAlgorithm CtSummaryParamsDimension = "PUBLIC_KEY_ALGORITHM"
	CtSummaryParamsDimensionSignatureAlgorithm CtSummaryParamsDimension = "SIGNATURE_ALGORITHM"
	CtSummaryParamsDimensionTld                CtSummaryParamsDimension = "TLD"
	CtSummaryParamsDimensionValidationLevel    CtSummaryParamsDimension = "VALIDATION_LEVEL"
)

func (r CtSummaryParamsDimension) IsKnown() bool {
	switch r {
	case CtSummaryParamsDimensionCA, CtSummaryParamsDimensionCAOwner, CtSummaryParamsDimensionDuration, CtSummaryParamsDimensionEntryType, CtSummaryParamsDimensionExpirationStatus, CtSummaryParamsDimensionHasIPs, CtSummaryParamsDimensionHasWildcards, CtSummaryParamsDimensionLog, CtSummaryParamsDimensionLogAPI, CtSummaryParamsDimensionLogOperator, CtSummaryParamsDimensionPublicKeyAlgorithm, CtSummaryParamsDimensionSignatureAlgorithm, CtSummaryParamsDimensionTld, CtSummaryParamsDimensionValidationLevel:
		return true
	}
	return false
}

type CtSummaryParamsDuration string

const (
	CtSummaryParamsDurationLte3D         CtSummaryParamsDuration = "LTE_3D"
	CtSummaryParamsDurationGt3DLte7D     CtSummaryParamsDuration = "GT_3D_LTE_7D"
	CtSummaryParamsDurationGt7DLte10D    CtSummaryParamsDuration = "GT_7D_LTE_10D"
	CtSummaryParamsDurationGt10DLte47D   CtSummaryParamsDuration = "GT_10D_LTE_47D"
	CtSummaryParamsDurationGt47DLte100D  CtSummaryParamsDuration = "GT_47D_LTE_100D"
	CtSummaryParamsDurationGt100DLte200D CtSummaryParamsDuration = "GT_100D_LTE_200D"
	CtSummaryParamsDurationGt200D        CtSummaryParamsDuration = "GT_200D"
)

func (r CtSummaryParamsDuration) IsKnown() bool {
	switch r {
	case CtSummaryParamsDurationLte3D, CtSummaryParamsDurationGt3DLte7D, CtSummaryParamsDurationGt7DLte10D, CtSummaryParamsDurationGt10DLte47D, CtSummaryParamsDurationGt47DLte100D, CtSummaryParamsDurationGt100DLte200D, CtSummaryParamsDurationGt200D:
		return true
	}
	return false
}

type CtSummaryParamsEntryType string

const (
	CtSummaryParamsEntryTypePrecertificate CtSummaryParamsEntryType = "PRECERTIFICATE"
	CtSummaryParamsEntryTypeCertificate    CtSummaryParamsEntryType = "CERTIFICATE"
)

func (r CtSummaryParamsEntryType) IsKnown() bool {
	switch r {
	case CtSummaryParamsEntryTypePrecertificate, CtSummaryParamsEntryTypeCertificate:
		return true
	}
	return false
}

type CtSummaryParamsExpirationStatus string

const (
	CtSummaryParamsExpirationStatusExpired CtSummaryParamsExpirationStatus = "EXPIRED"
	CtSummaryParamsExpirationStatusValid   CtSummaryParamsExpirationStatus = "VALID"
)

func (r CtSummaryParamsExpirationStatus) IsKnown() bool {
	switch r {
	case CtSummaryParamsExpirationStatusExpired, CtSummaryParamsExpirationStatusValid:
		return true
	}
	return false
}

// Format in which results will be returned.
type CtSummaryParamsFormat string

const (
	CtSummaryParamsFormatJson CtSummaryParamsFormat = "JSON"
	CtSummaryParamsFormatCsv  CtSummaryParamsFormat = "CSV"
)

func (r CtSummaryParamsFormat) IsKnown() bool {
	switch r {
	case CtSummaryParamsFormatJson, CtSummaryParamsFormatCsv:
		return true
	}
	return false
}

type CtSummaryParamsLogAPI string

const (
	CtSummaryParamsLogAPIRfc6962 CtSummaryParamsLogAPI = "RFC6962"
	CtSummaryParamsLogAPIStatic  CtSummaryParamsLogAPI = "STATIC"
)

func (r CtSummaryParamsLogAPI) IsKnown() bool {
	switch r {
	case CtSummaryParamsLogAPIRfc6962, CtSummaryParamsLogAPIStatic:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CtSummaryParamsNormalization string

const (
	CtSummaryParamsNormalizationRawValues  CtSummaryParamsNormalization = "RAW_VALUES"
	CtSummaryParamsNormalizationPercentage CtSummaryParamsNormalization = "PERCENTAGE"
)

func (r CtSummaryParamsNormalization) IsKnown() bool {
	switch r {
	case CtSummaryParamsNormalizationRawValues, CtSummaryParamsNormalizationPercentage:
		return true
	}
	return false
}

type CtSummaryParamsPublicKeyAlgorithm string

const (
	CtSummaryParamsPublicKeyAlgorithmDsa   CtSummaryParamsPublicKeyAlgorithm = "DSA"
	CtSummaryParamsPublicKeyAlgorithmEcdsa CtSummaryParamsPublicKeyAlgorithm = "ECDSA"
	CtSummaryParamsPublicKeyAlgorithmRSA   CtSummaryParamsPublicKeyAlgorithm = "RSA"
)

func (r CtSummaryParamsPublicKeyAlgorithm) IsKnown() bool {
	switch r {
	case CtSummaryParamsPublicKeyAlgorithmDsa, CtSummaryParamsPublicKeyAlgorithmEcdsa, CtSummaryParamsPublicKeyAlgorithmRSA:
		return true
	}
	return false
}

type CtSummaryParamsSignatureAlgorithm string

const (
	CtSummaryParamsSignatureAlgorithmDsaSha1     CtSummaryParamsSignatureAlgorithm = "DSA_SHA_1"
	CtSummaryParamsSignatureAlgorithmDsaSha256   CtSummaryParamsSignatureAlgorithm = "DSA_SHA_256"
	CtSummaryParamsSignatureAlgorithmEcdsaSha1   CtSummaryParamsSignatureAlgorithm = "ECDSA_SHA_1"
	CtSummaryParamsSignatureAlgorithmEcdsaSha256 CtSummaryParamsSignatureAlgorithm = "ECDSA_SHA_256"
	CtSummaryParamsSignatureAlgorithmEcdsaSha384 CtSummaryParamsSignatureAlgorithm = "ECDSA_SHA_384"
	CtSummaryParamsSignatureAlgorithmEcdsaSha512 CtSummaryParamsSignatureAlgorithm = "ECDSA_SHA_512"
	CtSummaryParamsSignatureAlgorithmPssSha256   CtSummaryParamsSignatureAlgorithm = "PSS_SHA_256"
	CtSummaryParamsSignatureAlgorithmPssSha384   CtSummaryParamsSignatureAlgorithm = "PSS_SHA_384"
	CtSummaryParamsSignatureAlgorithmPssSha512   CtSummaryParamsSignatureAlgorithm = "PSS_SHA_512"
	CtSummaryParamsSignatureAlgorithmRSAMd2      CtSummaryParamsSignatureAlgorithm = "RSA_MD2"
	CtSummaryParamsSignatureAlgorithmRSAMd5      CtSummaryParamsSignatureAlgorithm = "RSA_MD5"
	CtSummaryParamsSignatureAlgorithmRSASha1     CtSummaryParamsSignatureAlgorithm = "RSA_SHA_1"
	CtSummaryParamsSignatureAlgorithmRSASha256   CtSummaryParamsSignatureAlgorithm = "RSA_SHA_256"
	CtSummaryParamsSignatureAlgorithmRSASha384   CtSummaryParamsSignatureAlgorithm = "RSA_SHA_384"
	CtSummaryParamsSignatureAlgorithmRSASha512   CtSummaryParamsSignatureAlgorithm = "RSA_SHA_512"
)

func (r CtSummaryParamsSignatureAlgorithm) IsKnown() bool {
	switch r {
	case CtSummaryParamsSignatureAlgorithmDsaSha1, CtSummaryParamsSignatureAlgorithmDsaSha256, CtSummaryParamsSignatureAlgorithmEcdsaSha1, CtSummaryParamsSignatureAlgorithmEcdsaSha256, CtSummaryParamsSignatureAlgorithmEcdsaSha384, CtSummaryParamsSignatureAlgorithmEcdsaSha512, CtSummaryParamsSignatureAlgorithmPssSha256, CtSummaryParamsSignatureAlgorithmPssSha384, CtSummaryParamsSignatureAlgorithmPssSha512, CtSummaryParamsSignatureAlgorithmRSAMd2, CtSummaryParamsSignatureAlgorithmRSAMd5, CtSummaryParamsSignatureAlgorithmRSASha1, CtSummaryParamsSignatureAlgorithmRSASha256, CtSummaryParamsSignatureAlgorithmRSASha384, CtSummaryParamsSignatureAlgorithmRSASha512:
		return true
	}
	return false
}

type CtSummaryParamsUniqueEntry string

const (
	CtSummaryParamsUniqueEntryTrue  CtSummaryParamsUniqueEntry = "true"
	CtSummaryParamsUniqueEntryFalse CtSummaryParamsUniqueEntry = "false"
)

func (r CtSummaryParamsUniqueEntry) IsKnown() bool {
	switch r {
	case CtSummaryParamsUniqueEntryTrue, CtSummaryParamsUniqueEntryFalse:
		return true
	}
	return false
}

type CtSummaryParamsValidationLevel string

const (
	CtSummaryParamsValidationLevelDomain       CtSummaryParamsValidationLevel = "DOMAIN"
	CtSummaryParamsValidationLevelOrganization CtSummaryParamsValidationLevel = "ORGANIZATION"
	CtSummaryParamsValidationLevelExtended     CtSummaryParamsValidationLevel = "EXTENDED"
)

func (r CtSummaryParamsValidationLevel) IsKnown() bool {
	switch r {
	case CtSummaryParamsValidationLevelDomain, CtSummaryParamsValidationLevelOrganization, CtSummaryParamsValidationLevelExtended:
		return true
	}
	return false
}

type CtSummaryResponseEnvelope struct {
	Result  CtSummaryResponse             `json:"result,required"`
	Success bool                          `json:"success,required"`
	JSON    ctSummaryResponseEnvelopeJSON `json:"-"`
}

// ctSummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [CtSummaryResponseEnvelope]
type ctSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CtTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[CtTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by certificate authority.
	CA param.Field[[]string] `query:"ca"`
	// Filters results by certificate authority owner.
	CAOwner param.Field[[]string] `query:"caOwner"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by certificate duration.
	Duration param.Field[[]CtTimeseriesParamsDuration] `query:"duration"`
	// Filters results by entry type (certificate vs. pre-certificate).
	EntryType param.Field[[]CtTimeseriesParamsEntryType] `query:"entryType"`
	// Filters results by expiration status (expired vs. valid).
	ExpirationStatus param.Field[[]CtTimeseriesParamsExpirationStatus] `query:"expirationStatus"`
	// Format in which results will be returned.
	Format param.Field[CtTimeseriesParamsFormat] `query:"format"`
	// Filters results based on whether the certificates are bound to specific IP
	// addresses.
	HasIPs param.Field[[]bool] `query:"hasIps"`
	// Filters results based on whether the certificates contain wildcard domains.
	HasWildcards param.Field[[]bool] `query:"hasWildcards"`
	// Filters results by certificate log.
	Log param.Field[[]string] `query:"log"`
	// Filters results by certificate log API (RFC6962 vs. static).
	LogAPI param.Field[[]CtTimeseriesParamsLogAPI] `query:"logApi"`
	// Filters results by certificate log operator.
	LogOperator param.Field[[]string] `query:"logOperator"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by public key algorithm.
	PublicKeyAlgorithm param.Field[[]CtTimeseriesParamsPublicKeyAlgorithm] `query:"publicKeyAlgorithm"`
	// Filters results by signature algorithm.
	SignatureAlgorithm param.Field[[]CtTimeseriesParamsSignatureAlgorithm] `query:"signatureAlgorithm"`
	// Filters results by top-level domain.
	Tld param.Field[[]string] `query:"tld"`
	// Specifies whether to filter out duplicate certificates and pre-certificates. Set
	// to true for unique entries only.
	UniqueEntries param.Field[[]CtTimeseriesParamsUniqueEntry] `query:"uniqueEntries"`
	// Filters results by validation level.
	ValidationLevel param.Field[[]CtTimeseriesParamsValidationLevel] `query:"validationLevel"`
}

// URLQuery serializes [CtTimeseriesParams]'s query parameters as `url.Values`.
func (r CtTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CtTimeseriesParamsAggInterval string

const (
	CtTimeseriesParamsAggInterval15m CtTimeseriesParamsAggInterval = "15m"
	CtTimeseriesParamsAggInterval1h  CtTimeseriesParamsAggInterval = "1h"
	CtTimeseriesParamsAggInterval1d  CtTimeseriesParamsAggInterval = "1d"
	CtTimeseriesParamsAggInterval1w  CtTimeseriesParamsAggInterval = "1w"
)

func (r CtTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsAggInterval15m, CtTimeseriesParamsAggInterval1h, CtTimeseriesParamsAggInterval1d, CtTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type CtTimeseriesParamsDuration string

const (
	CtTimeseriesParamsDurationLte3D         CtTimeseriesParamsDuration = "LTE_3D"
	CtTimeseriesParamsDurationGt3DLte7D     CtTimeseriesParamsDuration = "GT_3D_LTE_7D"
	CtTimeseriesParamsDurationGt7DLte10D    CtTimeseriesParamsDuration = "GT_7D_LTE_10D"
	CtTimeseriesParamsDurationGt10DLte47D   CtTimeseriesParamsDuration = "GT_10D_LTE_47D"
	CtTimeseriesParamsDurationGt47DLte100D  CtTimeseriesParamsDuration = "GT_47D_LTE_100D"
	CtTimeseriesParamsDurationGt100DLte200D CtTimeseriesParamsDuration = "GT_100D_LTE_200D"
	CtTimeseriesParamsDurationGt200D        CtTimeseriesParamsDuration = "GT_200D"
)

func (r CtTimeseriesParamsDuration) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsDurationLte3D, CtTimeseriesParamsDurationGt3DLte7D, CtTimeseriesParamsDurationGt7DLte10D, CtTimeseriesParamsDurationGt10DLte47D, CtTimeseriesParamsDurationGt47DLte100D, CtTimeseriesParamsDurationGt100DLte200D, CtTimeseriesParamsDurationGt200D:
		return true
	}
	return false
}

type CtTimeseriesParamsEntryType string

const (
	CtTimeseriesParamsEntryTypePrecertificate CtTimeseriesParamsEntryType = "PRECERTIFICATE"
	CtTimeseriesParamsEntryTypeCertificate    CtTimeseriesParamsEntryType = "CERTIFICATE"
)

func (r CtTimeseriesParamsEntryType) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsEntryTypePrecertificate, CtTimeseriesParamsEntryTypeCertificate:
		return true
	}
	return false
}

type CtTimeseriesParamsExpirationStatus string

const (
	CtTimeseriesParamsExpirationStatusExpired CtTimeseriesParamsExpirationStatus = "EXPIRED"
	CtTimeseriesParamsExpirationStatusValid   CtTimeseriesParamsExpirationStatus = "VALID"
)

func (r CtTimeseriesParamsExpirationStatus) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsExpirationStatusExpired, CtTimeseriesParamsExpirationStatusValid:
		return true
	}
	return false
}

// Format in which results will be returned.
type CtTimeseriesParamsFormat string

const (
	CtTimeseriesParamsFormatJson CtTimeseriesParamsFormat = "JSON"
	CtTimeseriesParamsFormatCsv  CtTimeseriesParamsFormat = "CSV"
)

func (r CtTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsFormatJson, CtTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type CtTimeseriesParamsLogAPI string

const (
	CtTimeseriesParamsLogAPIRfc6962 CtTimeseriesParamsLogAPI = "RFC6962"
	CtTimeseriesParamsLogAPIStatic  CtTimeseriesParamsLogAPI = "STATIC"
)

func (r CtTimeseriesParamsLogAPI) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsLogAPIRfc6962, CtTimeseriesParamsLogAPIStatic:
		return true
	}
	return false
}

type CtTimeseriesParamsPublicKeyAlgorithm string

const (
	CtTimeseriesParamsPublicKeyAlgorithmDsa   CtTimeseriesParamsPublicKeyAlgorithm = "DSA"
	CtTimeseriesParamsPublicKeyAlgorithmEcdsa CtTimeseriesParamsPublicKeyAlgorithm = "ECDSA"
	CtTimeseriesParamsPublicKeyAlgorithmRSA   CtTimeseriesParamsPublicKeyAlgorithm = "RSA"
)

func (r CtTimeseriesParamsPublicKeyAlgorithm) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsPublicKeyAlgorithmDsa, CtTimeseriesParamsPublicKeyAlgorithmEcdsa, CtTimeseriesParamsPublicKeyAlgorithmRSA:
		return true
	}
	return false
}

type CtTimeseriesParamsSignatureAlgorithm string

const (
	CtTimeseriesParamsSignatureAlgorithmDsaSha1     CtTimeseriesParamsSignatureAlgorithm = "DSA_SHA_1"
	CtTimeseriesParamsSignatureAlgorithmDsaSha256   CtTimeseriesParamsSignatureAlgorithm = "DSA_SHA_256"
	CtTimeseriesParamsSignatureAlgorithmEcdsaSha1   CtTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_1"
	CtTimeseriesParamsSignatureAlgorithmEcdsaSha256 CtTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_256"
	CtTimeseriesParamsSignatureAlgorithmEcdsaSha384 CtTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_384"
	CtTimeseriesParamsSignatureAlgorithmEcdsaSha512 CtTimeseriesParamsSignatureAlgorithm = "ECDSA_SHA_512"
	CtTimeseriesParamsSignatureAlgorithmPssSha256   CtTimeseriesParamsSignatureAlgorithm = "PSS_SHA_256"
	CtTimeseriesParamsSignatureAlgorithmPssSha384   CtTimeseriesParamsSignatureAlgorithm = "PSS_SHA_384"
	CtTimeseriesParamsSignatureAlgorithmPssSha512   CtTimeseriesParamsSignatureAlgorithm = "PSS_SHA_512"
	CtTimeseriesParamsSignatureAlgorithmRSAMd2      CtTimeseriesParamsSignatureAlgorithm = "RSA_MD2"
	CtTimeseriesParamsSignatureAlgorithmRSAMd5      CtTimeseriesParamsSignatureAlgorithm = "RSA_MD5"
	CtTimeseriesParamsSignatureAlgorithmRSASha1     CtTimeseriesParamsSignatureAlgorithm = "RSA_SHA_1"
	CtTimeseriesParamsSignatureAlgorithmRSASha256   CtTimeseriesParamsSignatureAlgorithm = "RSA_SHA_256"
	CtTimeseriesParamsSignatureAlgorithmRSASha384   CtTimeseriesParamsSignatureAlgorithm = "RSA_SHA_384"
	CtTimeseriesParamsSignatureAlgorithmRSASha512   CtTimeseriesParamsSignatureAlgorithm = "RSA_SHA_512"
)

func (r CtTimeseriesParamsSignatureAlgorithm) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsSignatureAlgorithmDsaSha1, CtTimeseriesParamsSignatureAlgorithmDsaSha256, CtTimeseriesParamsSignatureAlgorithmEcdsaSha1, CtTimeseriesParamsSignatureAlgorithmEcdsaSha256, CtTimeseriesParamsSignatureAlgorithmEcdsaSha384, CtTimeseriesParamsSignatureAlgorithmEcdsaSha512, CtTimeseriesParamsSignatureAlgorithmPssSha256, CtTimeseriesParamsSignatureAlgorithmPssSha384, CtTimeseriesParamsSignatureAlgorithmPssSha512, CtTimeseriesParamsSignatureAlgorithmRSAMd2, CtTimeseriesParamsSignatureAlgorithmRSAMd5, CtTimeseriesParamsSignatureAlgorithmRSASha1, CtTimeseriesParamsSignatureAlgorithmRSASha256, CtTimeseriesParamsSignatureAlgorithmRSASha384, CtTimeseriesParamsSignatureAlgorithmRSASha512:
		return true
	}
	return false
}

type CtTimeseriesParamsUniqueEntry string

const (
	CtTimeseriesParamsUniqueEntryTrue  CtTimeseriesParamsUniqueEntry = "true"
	CtTimeseriesParamsUniqueEntryFalse CtTimeseriesParamsUniqueEntry = "false"
)

func (r CtTimeseriesParamsUniqueEntry) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsUniqueEntryTrue, CtTimeseriesParamsUniqueEntryFalse:
		return true
	}
	return false
}

type CtTimeseriesParamsValidationLevel string

const (
	CtTimeseriesParamsValidationLevelDomain       CtTimeseriesParamsValidationLevel = "DOMAIN"
	CtTimeseriesParamsValidationLevelOrganization CtTimeseriesParamsValidationLevel = "ORGANIZATION"
	CtTimeseriesParamsValidationLevelExtended     CtTimeseriesParamsValidationLevel = "EXTENDED"
)

func (r CtTimeseriesParamsValidationLevel) IsKnown() bool {
	switch r {
	case CtTimeseriesParamsValidationLevelDomain, CtTimeseriesParamsValidationLevelOrganization, CtTimeseriesParamsValidationLevelExtended:
		return true
	}
	return false
}

type CtTimeseriesResponseEnvelope struct {
	Result  CtTimeseriesResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    ctTimeseriesResponseEnvelopeJSON `json:"-"`
}

// ctTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [CtTimeseriesResponseEnvelope]
type ctTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CtTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[CtTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Filters results by certificate authority.
	CA param.Field[[]string] `query:"ca"`
	// Filters results by certificate authority owner.
	CAOwner param.Field[[]string] `query:"caOwner"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by certificate duration.
	Duration param.Field[[]CtTimeseriesGroupsParamsDuration] `query:"duration"`
	// Filters results by entry type (certificate vs. pre-certificate).
	EntryType param.Field[[]CtTimeseriesGroupsParamsEntryType] `query:"entryType"`
	// Filters results by expiration status (expired vs. valid).
	ExpirationStatus param.Field[[]CtTimeseriesGroupsParamsExpirationStatus] `query:"expirationStatus"`
	// Format in which results will be returned.
	Format param.Field[CtTimeseriesGroupsParamsFormat] `query:"format"`
	// Filters results based on whether the certificates are bound to specific IP
	// addresses.
	HasIPs param.Field[[]bool] `query:"hasIps"`
	// Filters results based on whether the certificates contain wildcard domains.
	HasWildcards param.Field[[]bool] `query:"hasWildcards"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by certificate log.
	Log param.Field[[]string] `query:"log"`
	// Filters results by certificate log API (RFC6962 vs. static).
	LogAPI param.Field[[]CtTimeseriesGroupsParamsLogAPI] `query:"logApi"`
	// Filters results by certificate log operator.
	LogOperator param.Field[[]string] `query:"logOperator"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[CtTimeseriesGroupsParamsNormalization] `query:"normalization"`
	// Filters results by public key algorithm.
	PublicKeyAlgorithm param.Field[[]CtTimeseriesGroupsParamsPublicKeyAlgorithm] `query:"publicKeyAlgorithm"`
	// Filters results by signature algorithm.
	SignatureAlgorithm param.Field[[]CtTimeseriesGroupsParamsSignatureAlgorithm] `query:"signatureAlgorithm"`
	// Filters results by top-level domain.
	Tld param.Field[[]string] `query:"tld"`
	// Specifies whether to filter out duplicate certificates and pre-certificates. Set
	// to true for unique entries only.
	UniqueEntries param.Field[[]CtTimeseriesGroupsParamsUniqueEntry] `query:"uniqueEntries"`
	// Filters results by validation level.
	ValidationLevel param.Field[[]CtTimeseriesGroupsParamsValidationLevel] `query:"validationLevel"`
}

// URLQuery serializes [CtTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r CtTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the certificate attribute by which to group the results.
type CtTimeseriesGroupsParamsDimension string

const (
	CtTimeseriesGroupsParamsDimensionCA                 CtTimeseriesGroupsParamsDimension = "CA"
	CtTimeseriesGroupsParamsDimensionCAOwner            CtTimeseriesGroupsParamsDimension = "CA_OWNER"
	CtTimeseriesGroupsParamsDimensionDuration           CtTimeseriesGroupsParamsDimension = "DURATION"
	CtTimeseriesGroupsParamsDimensionEntryType          CtTimeseriesGroupsParamsDimension = "ENTRY_TYPE"
	CtTimeseriesGroupsParamsDimensionExpirationStatus   CtTimeseriesGroupsParamsDimension = "EXPIRATION_STATUS"
	CtTimeseriesGroupsParamsDimensionHasIPs             CtTimeseriesGroupsParamsDimension = "HAS_IPS"
	CtTimeseriesGroupsParamsDimensionHasWildcards       CtTimeseriesGroupsParamsDimension = "HAS_WILDCARDS"
	CtTimeseriesGroupsParamsDimensionLog                CtTimeseriesGroupsParamsDimension = "LOG"
	CtTimeseriesGroupsParamsDimensionLogAPI             CtTimeseriesGroupsParamsDimension = "LOG_API"
	CtTimeseriesGroupsParamsDimensionLogOperator        CtTimeseriesGroupsParamsDimension = "LOG_OPERATOR"
	CtTimeseriesGroupsParamsDimensionPublicKeyAlgorithm CtTimeseriesGroupsParamsDimension = "PUBLIC_KEY_ALGORITHM"
	CtTimeseriesGroupsParamsDimensionSignatureAlgorithm CtTimeseriesGroupsParamsDimension = "SIGNATURE_ALGORITHM"
	CtTimeseriesGroupsParamsDimensionTld                CtTimeseriesGroupsParamsDimension = "TLD"
	CtTimeseriesGroupsParamsDimensionValidationLevel    CtTimeseriesGroupsParamsDimension = "VALIDATION_LEVEL"
)

func (r CtTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsDimensionCA, CtTimeseriesGroupsParamsDimensionCAOwner, CtTimeseriesGroupsParamsDimensionDuration, CtTimeseriesGroupsParamsDimensionEntryType, CtTimeseriesGroupsParamsDimensionExpirationStatus, CtTimeseriesGroupsParamsDimensionHasIPs, CtTimeseriesGroupsParamsDimensionHasWildcards, CtTimeseriesGroupsParamsDimensionLog, CtTimeseriesGroupsParamsDimensionLogAPI, CtTimeseriesGroupsParamsDimensionLogOperator, CtTimeseriesGroupsParamsDimensionPublicKeyAlgorithm, CtTimeseriesGroupsParamsDimensionSignatureAlgorithm, CtTimeseriesGroupsParamsDimensionTld, CtTimeseriesGroupsParamsDimensionValidationLevel:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type CtTimeseriesGroupsParamsAggInterval string

const (
	CtTimeseriesGroupsParamsAggInterval15m CtTimeseriesGroupsParamsAggInterval = "15m"
	CtTimeseriesGroupsParamsAggInterval1h  CtTimeseriesGroupsParamsAggInterval = "1h"
	CtTimeseriesGroupsParamsAggInterval1d  CtTimeseriesGroupsParamsAggInterval = "1d"
	CtTimeseriesGroupsParamsAggInterval1w  CtTimeseriesGroupsParamsAggInterval = "1w"
)

func (r CtTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsAggInterval15m, CtTimeseriesGroupsParamsAggInterval1h, CtTimeseriesGroupsParamsAggInterval1d, CtTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsDuration string

const (
	CtTimeseriesGroupsParamsDurationLte3D         CtTimeseriesGroupsParamsDuration = "LTE_3D"
	CtTimeseriesGroupsParamsDurationGt3DLte7D     CtTimeseriesGroupsParamsDuration = "GT_3D_LTE_7D"
	CtTimeseriesGroupsParamsDurationGt7DLte10D    CtTimeseriesGroupsParamsDuration = "GT_7D_LTE_10D"
	CtTimeseriesGroupsParamsDurationGt10DLte47D   CtTimeseriesGroupsParamsDuration = "GT_10D_LTE_47D"
	CtTimeseriesGroupsParamsDurationGt47DLte100D  CtTimeseriesGroupsParamsDuration = "GT_47D_LTE_100D"
	CtTimeseriesGroupsParamsDurationGt100DLte200D CtTimeseriesGroupsParamsDuration = "GT_100D_LTE_200D"
	CtTimeseriesGroupsParamsDurationGt200D        CtTimeseriesGroupsParamsDuration = "GT_200D"
)

func (r CtTimeseriesGroupsParamsDuration) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsDurationLte3D, CtTimeseriesGroupsParamsDurationGt3DLte7D, CtTimeseriesGroupsParamsDurationGt7DLte10D, CtTimeseriesGroupsParamsDurationGt10DLte47D, CtTimeseriesGroupsParamsDurationGt47DLte100D, CtTimeseriesGroupsParamsDurationGt100DLte200D, CtTimeseriesGroupsParamsDurationGt200D:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsEntryType string

const (
	CtTimeseriesGroupsParamsEntryTypePrecertificate CtTimeseriesGroupsParamsEntryType = "PRECERTIFICATE"
	CtTimeseriesGroupsParamsEntryTypeCertificate    CtTimeseriesGroupsParamsEntryType = "CERTIFICATE"
)

func (r CtTimeseriesGroupsParamsEntryType) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsEntryTypePrecertificate, CtTimeseriesGroupsParamsEntryTypeCertificate:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsExpirationStatus string

const (
	CtTimeseriesGroupsParamsExpirationStatusExpired CtTimeseriesGroupsParamsExpirationStatus = "EXPIRED"
	CtTimeseriesGroupsParamsExpirationStatusValid   CtTimeseriesGroupsParamsExpirationStatus = "VALID"
)

func (r CtTimeseriesGroupsParamsExpirationStatus) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsExpirationStatusExpired, CtTimeseriesGroupsParamsExpirationStatusValid:
		return true
	}
	return false
}

// Format in which results will be returned.
type CtTimeseriesGroupsParamsFormat string

const (
	CtTimeseriesGroupsParamsFormatJson CtTimeseriesGroupsParamsFormat = "JSON"
	CtTimeseriesGroupsParamsFormatCsv  CtTimeseriesGroupsParamsFormat = "CSV"
)

func (r CtTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsFormatJson, CtTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsLogAPI string

const (
	CtTimeseriesGroupsParamsLogAPIRfc6962 CtTimeseriesGroupsParamsLogAPI = "RFC6962"
	CtTimeseriesGroupsParamsLogAPIStatic  CtTimeseriesGroupsParamsLogAPI = "STATIC"
)

func (r CtTimeseriesGroupsParamsLogAPI) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsLogAPIRfc6962, CtTimeseriesGroupsParamsLogAPIStatic:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type CtTimeseriesGroupsParamsNormalization string

const (
	CtTimeseriesGroupsParamsNormalizationRawValues  CtTimeseriesGroupsParamsNormalization = "RAW_VALUES"
	CtTimeseriesGroupsParamsNormalizationPercentage CtTimeseriesGroupsParamsNormalization = "PERCENTAGE"
)

func (r CtTimeseriesGroupsParamsNormalization) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsNormalizationRawValues, CtTimeseriesGroupsParamsNormalizationPercentage:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsPublicKeyAlgorithm string

const (
	CtTimeseriesGroupsParamsPublicKeyAlgorithmDsa   CtTimeseriesGroupsParamsPublicKeyAlgorithm = "DSA"
	CtTimeseriesGroupsParamsPublicKeyAlgorithmEcdsa CtTimeseriesGroupsParamsPublicKeyAlgorithm = "ECDSA"
	CtTimeseriesGroupsParamsPublicKeyAlgorithmRSA   CtTimeseriesGroupsParamsPublicKeyAlgorithm = "RSA"
)

func (r CtTimeseriesGroupsParamsPublicKeyAlgorithm) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsPublicKeyAlgorithmDsa, CtTimeseriesGroupsParamsPublicKeyAlgorithmEcdsa, CtTimeseriesGroupsParamsPublicKeyAlgorithmRSA:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsSignatureAlgorithm string

const (
	CtTimeseriesGroupsParamsSignatureAlgorithmDsaSha1     CtTimeseriesGroupsParamsSignatureAlgorithm = "DSA_SHA_1"
	CtTimeseriesGroupsParamsSignatureAlgorithmDsaSha256   CtTimeseriesGroupsParamsSignatureAlgorithm = "DSA_SHA_256"
	CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha1   CtTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_1"
	CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha256 CtTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_256"
	CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha384 CtTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_384"
	CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha512 CtTimeseriesGroupsParamsSignatureAlgorithm = "ECDSA_SHA_512"
	CtTimeseriesGroupsParamsSignatureAlgorithmPssSha256   CtTimeseriesGroupsParamsSignatureAlgorithm = "PSS_SHA_256"
	CtTimeseriesGroupsParamsSignatureAlgorithmPssSha384   CtTimeseriesGroupsParamsSignatureAlgorithm = "PSS_SHA_384"
	CtTimeseriesGroupsParamsSignatureAlgorithmPssSha512   CtTimeseriesGroupsParamsSignatureAlgorithm = "PSS_SHA_512"
	CtTimeseriesGroupsParamsSignatureAlgorithmRSAMd2      CtTimeseriesGroupsParamsSignatureAlgorithm = "RSA_MD2"
	CtTimeseriesGroupsParamsSignatureAlgorithmRSAMd5      CtTimeseriesGroupsParamsSignatureAlgorithm = "RSA_MD5"
	CtTimeseriesGroupsParamsSignatureAlgorithmRSASha1     CtTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_1"
	CtTimeseriesGroupsParamsSignatureAlgorithmRSASha256   CtTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_256"
	CtTimeseriesGroupsParamsSignatureAlgorithmRSASha384   CtTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_384"
	CtTimeseriesGroupsParamsSignatureAlgorithmRSASha512   CtTimeseriesGroupsParamsSignatureAlgorithm = "RSA_SHA_512"
)

func (r CtTimeseriesGroupsParamsSignatureAlgorithm) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsSignatureAlgorithmDsaSha1, CtTimeseriesGroupsParamsSignatureAlgorithmDsaSha256, CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha1, CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha256, CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha384, CtTimeseriesGroupsParamsSignatureAlgorithmEcdsaSha512, CtTimeseriesGroupsParamsSignatureAlgorithmPssSha256, CtTimeseriesGroupsParamsSignatureAlgorithmPssSha384, CtTimeseriesGroupsParamsSignatureAlgorithmPssSha512, CtTimeseriesGroupsParamsSignatureAlgorithmRSAMd2, CtTimeseriesGroupsParamsSignatureAlgorithmRSAMd5, CtTimeseriesGroupsParamsSignatureAlgorithmRSASha1, CtTimeseriesGroupsParamsSignatureAlgorithmRSASha256, CtTimeseriesGroupsParamsSignatureAlgorithmRSASha384, CtTimeseriesGroupsParamsSignatureAlgorithmRSASha512:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsUniqueEntry string

const (
	CtTimeseriesGroupsParamsUniqueEntryTrue  CtTimeseriesGroupsParamsUniqueEntry = "true"
	CtTimeseriesGroupsParamsUniqueEntryFalse CtTimeseriesGroupsParamsUniqueEntry = "false"
)

func (r CtTimeseriesGroupsParamsUniqueEntry) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsUniqueEntryTrue, CtTimeseriesGroupsParamsUniqueEntryFalse:
		return true
	}
	return false
}

type CtTimeseriesGroupsParamsValidationLevel string

const (
	CtTimeseriesGroupsParamsValidationLevelDomain       CtTimeseriesGroupsParamsValidationLevel = "DOMAIN"
	CtTimeseriesGroupsParamsValidationLevelOrganization CtTimeseriesGroupsParamsValidationLevel = "ORGANIZATION"
	CtTimeseriesGroupsParamsValidationLevelExtended     CtTimeseriesGroupsParamsValidationLevel = "EXTENDED"
)

func (r CtTimeseriesGroupsParamsValidationLevel) IsKnown() bool {
	switch r {
	case CtTimeseriesGroupsParamsValidationLevelDomain, CtTimeseriesGroupsParamsValidationLevelOrganization, CtTimeseriesGroupsParamsValidationLevelExtended:
		return true
	}
	return false
}

type CtTimeseriesGroupsResponseEnvelope struct {
	Result  CtTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    ctTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// ctTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for the struct
// [CtTimeseriesGroupsResponseEnvelope]
type ctTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CtTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ctTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
