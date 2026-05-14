// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// RankingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRankingService] method instead.
type RankingService struct {
	Options          []option.RequestOption
	Domain           *RankingDomainService
	InternetServices *RankingInternetServiceService
}

// NewRankingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRankingService(opts ...option.RequestOption) (r *RankingService) {
	r = &RankingService{}
	r.Options = opts
	r.Domain = NewRankingDomainService(opts...)
	r.InternetServices = NewRankingInternetServiceService(opts...)
	return
}

// Retrieves domains rank over time.
func (r *RankingService) TimeseriesGroups(ctx context.Context, query RankingTimeseriesGroupsParams, opts ...option.RequestOption) (res *RankingTimeseriesGroupsResponse, err error) {
	var env RankingTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the top or trending domains based on their rank. Popular domains are
// domains of broad appeal based on how people use the Internet. Trending domains
// are domains that are generating a surge in interest. For more information on top
// domains, see https://blog.cloudflare.com/radar-domain-rankings/.
func (r *RankingService) Top(ctx context.Context, query RankingTopParams, opts ...option.RequestOption) (res *RankingTopResponse, err error) {
	var env RankingTopResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RankingTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   RankingTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 RankingTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   rankingTimeseriesGroupsResponseJSON   `json:"-"`
}

// rankingTimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [RankingTimeseriesGroupsResponse]
type rankingTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RankingTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RankingTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RankingTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RankingTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RankingTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RankingTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  rankingTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// rankingTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the
// struct [RankingTimeseriesGroupsResponseMeta]
type rankingTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RankingTimeseriesGroupsResponseMetaAggInterval string

const (
	RankingTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes RankingTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	RankingTimeseriesGroupsResponseMetaAggIntervalOneHour        RankingTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	RankingTimeseriesGroupsResponseMetaAggIntervalOneDay         RankingTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	RankingTimeseriesGroupsResponseMetaAggIntervalOneWeek        RankingTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	RankingTimeseriesGroupsResponseMetaAggIntervalOneMonth       RankingTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r RankingTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case RankingTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, RankingTimeseriesGroupsResponseMetaAggIntervalOneHour, RankingTimeseriesGroupsResponseMetaAggIntervalOneDay, RankingTimeseriesGroupsResponseMetaAggIntervalOneWeek, RankingTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RankingTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []RankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  rankingTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// rankingTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RankingTimeseriesGroupsResponseMetaConfidenceInfo]
type rankingTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            rankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// rankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type rankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RankingTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      rankingTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// rankingTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RankingTimeseriesGroupsResponseMetaDateRange]
type rankingTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RankingTimeseriesGroupsResponseMetaNormalization string

const (
	RankingTimeseriesGroupsResponseMetaNormalizationPercentage           RankingTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	RankingTimeseriesGroupsResponseMetaNormalizationMin0Max              RankingTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	RankingTimeseriesGroupsResponseMetaNormalizationMinMax               RankingTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	RankingTimeseriesGroupsResponseMetaNormalizationRawValues            RankingTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	RankingTimeseriesGroupsResponseMetaNormalizationPercentageChange     RankingTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	RankingTimeseriesGroupsResponseMetaNormalizationRollingAverage       RankingTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	RankingTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage RankingTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RankingTimeseriesGroupsResponseMetaNormalizationRatio                RankingTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r RankingTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case RankingTimeseriesGroupsResponseMetaNormalizationPercentage, RankingTimeseriesGroupsResponseMetaNormalizationMin0Max, RankingTimeseriesGroupsResponseMetaNormalizationMinMax, RankingTimeseriesGroupsResponseMetaNormalizationRawValues, RankingTimeseriesGroupsResponseMetaNormalizationPercentageChange, RankingTimeseriesGroupsResponseMetaNormalizationRollingAverage, RankingTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, RankingTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type RankingTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  rankingTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// rankingTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for the
// struct [RankingTimeseriesGroupsResponseMetaUnit]
type rankingTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RankingTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                                             `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]RankingTimeseriesGroupsResponseSerie0Union `json:"-,extras"`
	JSON        rankingTimeseriesGroupsResponseSerie0JSON               `json:"-"`
}

// rankingTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the
// struct [RankingTimeseriesGroupsResponseSerie0]
type rankingTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

// A numeric string.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RankingTimeseriesGroupsResponseSerie0Union interface {
	ImplementsRankingTimeseriesGroupsResponseSerie0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RankingTimeseriesGroupsResponseSerie0Union)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
	)
}

type RankingTopResponse struct {
	Meta RankingTopResponseMeta   `json:"meta,required"`
	Top0 []RankingTopResponseTop0 `json:"top_0,required"`
	JSON rankingTopResponseJSON   `json:"-"`
}

// rankingTopResponseJSON contains the JSON metadata for the struct
// [RankingTopResponse]
type rankingTopResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseJSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseMeta struct {
	ConfidenceInfo RankingTopResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RankingTopResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RankingTopResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RankingTopResponseMetaUnit `json:"units,required"`
	JSON  rankingTopResponseMetaJSON   `json:"-"`
}

// rankingTopResponseMetaJSON contains the JSON metadata for the struct
// [RankingTopResponseMeta]
type rankingTopResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RankingTopResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseMetaConfidenceInfo struct {
	Annotations []RankingTopResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                    `json:"level,required"`
	JSON  rankingTopResponseMetaConfidenceInfoJSON `json:"-"`
}

// rankingTopResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [RankingTopResponseMetaConfidenceInfo]
type rankingTopResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RankingTopResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                          `json:"startDate,required" format:"date-time"`
	JSON            rankingTopResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// rankingTopResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [RankingTopResponseMetaConfidenceInfoAnnotation]
type rankingTopResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RankingTopResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                           `json:"startTime,required" format:"date-time"`
	JSON      rankingTopResponseMetaDateRangeJSON `json:"-"`
}

// rankingTopResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [RankingTopResponseMetaDateRange]
type rankingTopResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RankingTopResponseMetaNormalization string

const (
	RankingTopResponseMetaNormalizationPercentage           RankingTopResponseMetaNormalization = "PERCENTAGE"
	RankingTopResponseMetaNormalizationMin0Max              RankingTopResponseMetaNormalization = "MIN0_MAX"
	RankingTopResponseMetaNormalizationMinMax               RankingTopResponseMetaNormalization = "MIN_MAX"
	RankingTopResponseMetaNormalizationRawValues            RankingTopResponseMetaNormalization = "RAW_VALUES"
	RankingTopResponseMetaNormalizationPercentageChange     RankingTopResponseMetaNormalization = "PERCENTAGE_CHANGE"
	RankingTopResponseMetaNormalizationRollingAverage       RankingTopResponseMetaNormalization = "ROLLING_AVERAGE"
	RankingTopResponseMetaNormalizationOverlappedPercentage RankingTopResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RankingTopResponseMetaNormalizationRatio                RankingTopResponseMetaNormalization = "RATIO"
)

func (r RankingTopResponseMetaNormalization) IsKnown() bool {
	switch r {
	case RankingTopResponseMetaNormalizationPercentage, RankingTopResponseMetaNormalizationMin0Max, RankingTopResponseMetaNormalizationMinMax, RankingTopResponseMetaNormalizationRawValues, RankingTopResponseMetaNormalizationPercentageChange, RankingTopResponseMetaNormalizationRollingAverage, RankingTopResponseMetaNormalizationOverlappedPercentage, RankingTopResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type RankingTopResponseMetaUnit struct {
	Name  string                         `json:"name,required"`
	Value string                         `json:"value,required"`
	JSON  rankingTopResponseMetaUnitJSON `json:"-"`
}

// rankingTopResponseMetaUnitJSON contains the JSON metadata for the struct
// [RankingTopResponseMetaUnit]
type rankingTopResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseTop0 struct {
	Categories []RankingTopResponseTop0Category `json:"categories,required"`
	Domain     string                           `json:"domain,required"`
	Rank       int64                            `json:"rank,required"`
	// Only available in TRENDING rankings.
	PctRankChange float64                    `json:"pctRankChange"`
	JSON          rankingTopResponseTop0JSON `json:"-"`
}

// rankingTopResponseTop0JSON contains the JSON metadata for the struct
// [RankingTopResponseTop0]
type rankingTopResponseTop0JSON struct {
	Categories    apijson.Field
	Domain        apijson.Field
	Rank          apijson.Field
	PctRankChange apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RankingTopResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RankingTopResponseTop0Category struct {
	ID              float64                            `json:"id,required"`
	Name            string                             `json:"name,required"`
	SuperCategoryID float64                            `json:"superCategoryId,required"`
	JSON            rankingTopResponseTop0CategoryJSON `json:"-"`
}

// rankingTopResponseTop0CategoryJSON contains the JSON metadata for the struct
// [RankingTopResponseTop0Category]
type rankingTopResponseTop0CategoryJSON struct {
	ID              apijson.Field
	Name            apijson.Field
	SuperCategoryID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RankingTopResponseTop0Category) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseTop0CategoryJSON) RawJSON() string {
	return r.raw
}

type RankingTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Filters results by domain name. Specify a comma-separated list of domain names.
	Domains param.Field[[]string] `query:"domains"`
	// Format in which results will be returned.
	Format param.Field[RankingTimeseriesGroupsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RankingTimeseriesGroupsParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RankingTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r RankingTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type RankingTimeseriesGroupsParamsFormat string

const (
	RankingTimeseriesGroupsParamsFormatJson RankingTimeseriesGroupsParamsFormat = "JSON"
	RankingTimeseriesGroupsParamsFormatCsv  RankingTimeseriesGroupsParamsFormat = "CSV"
)

func (r RankingTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RankingTimeseriesGroupsParamsFormatJson, RankingTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

// The ranking type.
type RankingTimeseriesGroupsParamsRankingType string

const (
	RankingTimeseriesGroupsParamsRankingTypePopular        RankingTimeseriesGroupsParamsRankingType = "POPULAR"
	RankingTimeseriesGroupsParamsRankingTypeTrendingRise   RankingTimeseriesGroupsParamsRankingType = "TRENDING_RISE"
	RankingTimeseriesGroupsParamsRankingTypeTrendingSteady RankingTimeseriesGroupsParamsRankingType = "TRENDING_STEADY"
)

func (r RankingTimeseriesGroupsParamsRankingType) IsKnown() bool {
	switch r {
	case RankingTimeseriesGroupsParamsRankingTypePopular, RankingTimeseriesGroupsParamsRankingTypeTrendingRise, RankingTimeseriesGroupsParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}

type RankingTimeseriesGroupsResponseEnvelope struct {
	Result  RankingTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    rankingTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// rankingTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for the
// struct [RankingTimeseriesGroupsResponseEnvelope]
type rankingTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RankingTopParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Filters results by domain category.
	DomainCategory param.Field[[]string] `query:"domainCategory"`
	// Format in which results will be returned.
	Format param.Field[RankingTopParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 location
	// codes.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// The ranking type.
	RankingType param.Field[RankingTopParamsRankingType] `query:"rankingType"`
}

// URLQuery serializes [RankingTopParams]'s query parameters as `url.Values`.
func (r RankingTopParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type RankingTopParamsFormat string

const (
	RankingTopParamsFormatJson RankingTopParamsFormat = "JSON"
	RankingTopParamsFormatCsv  RankingTopParamsFormat = "CSV"
)

func (r RankingTopParamsFormat) IsKnown() bool {
	switch r {
	case RankingTopParamsFormatJson, RankingTopParamsFormatCsv:
		return true
	}
	return false
}

// The ranking type.
type RankingTopParamsRankingType string

const (
	RankingTopParamsRankingTypePopular        RankingTopParamsRankingType = "POPULAR"
	RankingTopParamsRankingTypeTrendingRise   RankingTopParamsRankingType = "TRENDING_RISE"
	RankingTopParamsRankingTypeTrendingSteady RankingTopParamsRankingType = "TRENDING_STEADY"
)

func (r RankingTopParamsRankingType) IsKnown() bool {
	switch r {
	case RankingTopParamsRankingTypePopular, RankingTopParamsRankingTypeTrendingRise, RankingTopParamsRankingTypeTrendingSteady:
		return true
	}
	return false
}

type RankingTopResponseEnvelope struct {
	Result  RankingTopResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    rankingTopResponseEnvelopeJSON `json:"-"`
}

// rankingTopResponseEnvelopeJSON contains the JSON metadata for the struct
// [RankingTopResponseEnvelope]
type rankingTopResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingTopResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingTopResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
