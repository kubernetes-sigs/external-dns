// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// BotWebCrawlerService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotWebCrawlerService] method instead.
type BotWebCrawlerService struct {
	Options []option.RequestOption
}

// NewBotWebCrawlerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotWebCrawlerService(opts ...option.RequestOption) (r *BotWebCrawlerService) {
	r = &BotWebCrawlerService{}
	r.Options = opts
	return
}

// Retrieves an aggregated summary of HTTP requests from crawlers, grouped by the
// specified dimension.
func (r *BotWebCrawlerService) Summary(ctx context.Context, dimension BotWebCrawlerSummaryParamsDimension, query BotWebCrawlerSummaryParams, opts ...option.RequestOption) (res *BotWebCrawlerSummaryResponse, err error) {
	var env BotWebCrawlerSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/bots/crawlers/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests from crawlers, grouped by chosen the
// specified dimension over time.
func (r *BotWebCrawlerService) TimeseriesGroups(ctx context.Context, dimension BotWebCrawlerTimeseriesGroupsParamsDimension, query BotWebCrawlerTimeseriesGroupsParams, opts ...option.RequestOption) (res *BotWebCrawlerTimeseriesGroupsResponse, err error) {
	var env BotWebCrawlerTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/bots/crawlers/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BotWebCrawlerSummaryResponse struct {
	// Metadata for the results.
	Meta     BotWebCrawlerSummaryResponseMeta `json:"meta,required"`
	Summary0 map[string]string                `json:"summary_0,required"`
	JSON     botWebCrawlerSummaryResponseJSON `json:"-"`
}

// botWebCrawlerSummaryResponseJSON contains the JSON metadata for the struct
// [BotWebCrawlerSummaryResponse]
type botWebCrawlerSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type BotWebCrawlerSummaryResponseMeta struct {
	ConfidenceInfo BotWebCrawlerSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []BotWebCrawlerSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization BotWebCrawlerSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []BotWebCrawlerSummaryResponseMetaUnit `json:"units,required"`
	JSON  botWebCrawlerSummaryResponseMetaJSON   `json:"-"`
}

// botWebCrawlerSummaryResponseMetaJSON contains the JSON metadata for the struct
// [BotWebCrawlerSummaryResponseMeta]
type botWebCrawlerSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BotWebCrawlerSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BotWebCrawlerSummaryResponseMetaConfidenceInfo struct {
	Annotations []BotWebCrawlerSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                              `json:"level,required"`
	JSON  botWebCrawlerSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// botWebCrawlerSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [BotWebCrawlerSummaryResponseMetaConfidenceInfo]
type botWebCrawlerSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type BotWebCrawlerSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	LinkedURL       string                                                       `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                    `json:"startDate,required" format:"date-time"`
	JSON            botWebCrawlerSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// botWebCrawlerSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [BotWebCrawlerSummaryResponseMetaConfidenceInfoAnnotation]
type botWebCrawlerSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BotWebCrawlerSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BotWebCrawlerSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      botWebCrawlerSummaryResponseMetaDateRangeJSON `json:"-"`
}

// botWebCrawlerSummaryResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [BotWebCrawlerSummaryResponseMetaDateRange]
type botWebCrawlerSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type BotWebCrawlerSummaryResponseMetaNormalization string

const (
	BotWebCrawlerSummaryResponseMetaNormalizationPercentage           BotWebCrawlerSummaryResponseMetaNormalization = "PERCENTAGE"
	BotWebCrawlerSummaryResponseMetaNormalizationMin0Max              BotWebCrawlerSummaryResponseMetaNormalization = "MIN0_MAX"
	BotWebCrawlerSummaryResponseMetaNormalizationMinMax               BotWebCrawlerSummaryResponseMetaNormalization = "MIN_MAX"
	BotWebCrawlerSummaryResponseMetaNormalizationRawValues            BotWebCrawlerSummaryResponseMetaNormalization = "RAW_VALUES"
	BotWebCrawlerSummaryResponseMetaNormalizationPercentageChange     BotWebCrawlerSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	BotWebCrawlerSummaryResponseMetaNormalizationRollingAverage       BotWebCrawlerSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	BotWebCrawlerSummaryResponseMetaNormalizationOverlappedPercentage BotWebCrawlerSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	BotWebCrawlerSummaryResponseMetaNormalizationRatio                BotWebCrawlerSummaryResponseMetaNormalization = "RATIO"
)

func (r BotWebCrawlerSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case BotWebCrawlerSummaryResponseMetaNormalizationPercentage, BotWebCrawlerSummaryResponseMetaNormalizationMin0Max, BotWebCrawlerSummaryResponseMetaNormalizationMinMax, BotWebCrawlerSummaryResponseMetaNormalizationRawValues, BotWebCrawlerSummaryResponseMetaNormalizationPercentageChange, BotWebCrawlerSummaryResponseMetaNormalizationRollingAverage, BotWebCrawlerSummaryResponseMetaNormalizationOverlappedPercentage, BotWebCrawlerSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type BotWebCrawlerSummaryResponseMetaUnit struct {
	Name  string                                   `json:"name,required"`
	Value string                                   `json:"value,required"`
	JSON  botWebCrawlerSummaryResponseMetaUnitJSON `json:"-"`
}

// botWebCrawlerSummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [BotWebCrawlerSummaryResponseMetaUnit]
type botWebCrawlerSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type BotWebCrawlerTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   BotWebCrawlerTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 BotWebCrawlerTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   botWebCrawlerTimeseriesGroupsResponseJSON   `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [BotWebCrawlerTimeseriesGroupsResponse]
type botWebCrawlerTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type BotWebCrawlerTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []BotWebCrawlerTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization BotWebCrawlerTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []BotWebCrawlerTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  botWebCrawlerTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the
// struct [BotWebCrawlerTimeseriesGroupsResponseMeta]
type botWebCrawlerTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BotWebCrawlerTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval string

const (
	BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneHour        BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneDay         BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneWeek        BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneMonth       BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r BotWebCrawlerTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneHour, BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneDay, BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneWeek, BotWebCrawlerTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                       `json:"level,required"`
	JSON  botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfo]
type botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                             `json:"startDate,required" format:"date-time"`
	JSON            botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BotWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BotWebCrawlerTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      botWebCrawlerTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [BotWebCrawlerTimeseriesGroupsResponseMetaDateRange]
type botWebCrawlerTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type BotWebCrawlerTimeseriesGroupsResponseMetaNormalization string

const (
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationPercentage           BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationMin0Max              BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationMinMax               BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationRawValues            BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationPercentageChange     BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationRollingAverage       BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationRatio                BotWebCrawlerTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r BotWebCrawlerTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationPercentage, BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationMin0Max, BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationMinMax, BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationRawValues, BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationPercentageChange, BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationRollingAverage, BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, BotWebCrawlerTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type BotWebCrawlerTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                            `json:"name,required"`
	Value string                                            `json:"value,required"`
	JSON  botWebCrawlerTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for
// the struct [BotWebCrawlerTimeseriesGroupsResponseMetaUnit]
type botWebCrawlerTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type BotWebCrawlerTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                                     `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                             `json:"-,extras"`
	JSON        botWebCrawlerTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for
// the struct [BotWebCrawlerTimeseriesGroupsResponseSerie0]
type botWebCrawlerTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BotWebCrawlerSummaryParams struct {
	// Filters results by bot operator.
	BotOperator param.Field[[]string] `query:"botOperator"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BotWebCrawlerSummaryParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [BotWebCrawlerSummaryParams]'s query parameters as
// `url.Values`.
func (r BotWebCrawlerSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type BotWebCrawlerSummaryParamsDimension string

const (
	BotWebCrawlerSummaryParamsDimensionUserAgent       BotWebCrawlerSummaryParamsDimension = "USER_AGENT"
	BotWebCrawlerSummaryParamsDimensionReferer         BotWebCrawlerSummaryParamsDimension = "REFERER"
	BotWebCrawlerSummaryParamsDimensionCrawlReferRatio BotWebCrawlerSummaryParamsDimension = "CRAWL_REFER_RATIO"
)

func (r BotWebCrawlerSummaryParamsDimension) IsKnown() bool {
	switch r {
	case BotWebCrawlerSummaryParamsDimensionUserAgent, BotWebCrawlerSummaryParamsDimensionReferer, BotWebCrawlerSummaryParamsDimensionCrawlReferRatio:
		return true
	}
	return false
}

// Format in which results will be returned.
type BotWebCrawlerSummaryParamsFormat string

const (
	BotWebCrawlerSummaryParamsFormatJson BotWebCrawlerSummaryParamsFormat = "JSON"
	BotWebCrawlerSummaryParamsFormatCsv  BotWebCrawlerSummaryParamsFormat = "CSV"
)

func (r BotWebCrawlerSummaryParamsFormat) IsKnown() bool {
	switch r {
	case BotWebCrawlerSummaryParamsFormatJson, BotWebCrawlerSummaryParamsFormatCsv:
		return true
	}
	return false
}

type BotWebCrawlerSummaryResponseEnvelope struct {
	Result  BotWebCrawlerSummaryResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    botWebCrawlerSummaryResponseEnvelopeJSON `json:"-"`
}

// botWebCrawlerSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [BotWebCrawlerSummaryResponseEnvelope]
type botWebCrawlerSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BotWebCrawlerTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[BotWebCrawlerTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Filters results by bot operator.
	BotOperator param.Field[[]string] `query:"botOperator"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[BotWebCrawlerTimeseriesGroupsParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [BotWebCrawlerTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r BotWebCrawlerTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type BotWebCrawlerTimeseriesGroupsParamsDimension string

const (
	BotWebCrawlerTimeseriesGroupsParamsDimensionUserAgent       BotWebCrawlerTimeseriesGroupsParamsDimension = "USER_AGENT"
	BotWebCrawlerTimeseriesGroupsParamsDimensionReferer         BotWebCrawlerTimeseriesGroupsParamsDimension = "REFERER"
	BotWebCrawlerTimeseriesGroupsParamsDimensionCrawlReferRatio BotWebCrawlerTimeseriesGroupsParamsDimension = "CRAWL_REFER_RATIO"
)

func (r BotWebCrawlerTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case BotWebCrawlerTimeseriesGroupsParamsDimensionUserAgent, BotWebCrawlerTimeseriesGroupsParamsDimensionReferer, BotWebCrawlerTimeseriesGroupsParamsDimensionCrawlReferRatio:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BotWebCrawlerTimeseriesGroupsParamsAggInterval string

const (
	BotWebCrawlerTimeseriesGroupsParamsAggInterval15m BotWebCrawlerTimeseriesGroupsParamsAggInterval = "15m"
	BotWebCrawlerTimeseriesGroupsParamsAggInterval1h  BotWebCrawlerTimeseriesGroupsParamsAggInterval = "1h"
	BotWebCrawlerTimeseriesGroupsParamsAggInterval1d  BotWebCrawlerTimeseriesGroupsParamsAggInterval = "1d"
	BotWebCrawlerTimeseriesGroupsParamsAggInterval1w  BotWebCrawlerTimeseriesGroupsParamsAggInterval = "1w"
)

func (r BotWebCrawlerTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case BotWebCrawlerTimeseriesGroupsParamsAggInterval15m, BotWebCrawlerTimeseriesGroupsParamsAggInterval1h, BotWebCrawlerTimeseriesGroupsParamsAggInterval1d, BotWebCrawlerTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type BotWebCrawlerTimeseriesGroupsParamsFormat string

const (
	BotWebCrawlerTimeseriesGroupsParamsFormatJson BotWebCrawlerTimeseriesGroupsParamsFormat = "JSON"
	BotWebCrawlerTimeseriesGroupsParamsFormatCsv  BotWebCrawlerTimeseriesGroupsParamsFormat = "CSV"
)

func (r BotWebCrawlerTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case BotWebCrawlerTimeseriesGroupsParamsFormatJson, BotWebCrawlerTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type BotWebCrawlerTimeseriesGroupsResponseEnvelope struct {
	Result  BotWebCrawlerTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    botWebCrawlerTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// botWebCrawlerTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for
// the struct [BotWebCrawlerTimeseriesGroupsResponseEnvelope]
type botWebCrawlerTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotWebCrawlerTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botWebCrawlerTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
