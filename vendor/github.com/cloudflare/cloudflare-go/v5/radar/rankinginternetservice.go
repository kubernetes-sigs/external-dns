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

// RankingInternetServiceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRankingInternetServiceService] method instead.
type RankingInternetServiceService struct {
	Options []option.RequestOption
}

// NewRankingInternetServiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRankingInternetServiceService(opts ...option.RequestOption) (r *RankingInternetServiceService) {
	r = &RankingInternetServiceService{}
	r.Options = opts
	return
}

// Retrieves the list of Internet services categories.
func (r *RankingInternetServiceService) Categories(ctx context.Context, query RankingInternetServiceCategoriesParams, opts ...option.RequestOption) (res *RankingInternetServiceCategoriesResponse, err error) {
	var env RankingInternetServiceCategoriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/categories"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves Internet Services rank update changes over time.
func (r *RankingInternetServiceService) TimeseriesGroups(ctx context.Context, query RankingInternetServiceTimeseriesGroupsParams, opts ...option.RequestOption) (res *RankingInternetServiceTimeseriesGroupsResponse, err error) {
	var env RankingInternetServiceTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves top Internet services based on their rank.
func (r *RankingInternetServiceService) Top(ctx context.Context, query RankingInternetServiceTopParams, opts ...option.RequestOption) (res *RankingInternetServiceTopResponse, err error) {
	var env RankingInternetServiceTopResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ranking/internet_services/top"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RankingInternetServiceCategoriesResponse struct {
	Categories0 []RankingInternetServiceCategoriesResponseCategories0 `json:"categories_0,required"`
	JSON        rankingInternetServiceCategoriesResponseJSON          `json:"-"`
}

// rankingInternetServiceCategoriesResponseJSON contains the JSON metadata for the
// struct [RankingInternetServiceCategoriesResponse]
type rankingInternetServiceCategoriesResponseJSON struct {
	Categories0 apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceCategoriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceCategoriesResponseJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceCategoriesResponseCategories0 struct {
	Name string                                                  `json:"name,required"`
	JSON rankingInternetServiceCategoriesResponseCategories0JSON `json:"-"`
}

// rankingInternetServiceCategoriesResponseCategories0JSON contains the JSON
// metadata for the struct [RankingInternetServiceCategoriesResponseCategories0]
type rankingInternetServiceCategoriesResponseCategories0JSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceCategoriesResponseCategories0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceCategoriesResponseCategories0JSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   RankingInternetServiceTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 RankingInternetServiceTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   rankingInternetServiceTimeseriesGroupsResponseJSON   `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseJSON contains the JSON metadata
// for the struct [RankingInternetServiceTimeseriesGroupsResponse]
type rankingInternetServiceTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type RankingInternetServiceTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []RankingInternetServiceTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RankingInternetServiceTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RankingInternetServiceTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  rankingInternetServiceTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseMetaJSON contains the JSON
// metadata for the struct [RankingInternetServiceTimeseriesGroupsResponseMeta]
type rankingInternetServiceTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval string

const (
	RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneHour        RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneDay         RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneWeek        RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneMonth       RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r RankingInternetServiceTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneHour, RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneDay, RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneWeek, RankingInternetServiceTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                `json:"level,required"`
	JSON  rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfo]
type rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                      `json:"startDate,required" format:"date-time"`
	JSON            rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RankingInternetServiceTimeseriesGroupsResponseMetaDateRange]
type rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RankingInternetServiceTimeseriesGroupsResponseMetaNormalization string

const (
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationPercentage           RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationMin0Max              RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationMinMax               RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationRawValues            RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationPercentageChange     RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationRollingAverage       RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationRatio                RankingInternetServiceTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r RankingInternetServiceTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationPercentage, RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationMin0Max, RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationMinMax, RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationRawValues, RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationPercentageChange, RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationRollingAverage, RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, RankingInternetServiceTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type RankingInternetServiceTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  rankingInternetServiceTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseMetaUnitJSON contains the JSON
// metadata for the struct [RankingInternetServiceTimeseriesGroupsResponseMetaUnit]
type rankingInternetServiceTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                                                            `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]RankingInternetServiceTimeseriesGroupsResponseSerie0Union `json:"-,extras"`
	JSON        rankingInternetServiceTimeseriesGroupsResponseSerie0JSON               `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseSerie0JSON contains the JSON
// metadata for the struct [RankingInternetServiceTimeseriesGroupsResponseSerie0]
type rankingInternetServiceTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

// A numeric string.
//
// Union satisfied by [shared.UnionString] or [shared.UnionFloat].
type RankingInternetServiceTimeseriesGroupsResponseSerie0Union interface {
	ImplementsRankingInternetServiceTimeseriesGroupsResponseSerie0Union()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RankingInternetServiceTimeseriesGroupsResponseSerie0Union)(nil)).Elem(),
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

type RankingInternetServiceTopResponse struct {
	Meta RankingInternetServiceTopResponseMeta   `json:"meta,required"`
	Top0 []RankingInternetServiceTopResponseTop0 `json:"top_0,required"`
	JSON rankingInternetServiceTopResponseJSON   `json:"-"`
}

// rankingInternetServiceTopResponseJSON contains the JSON metadata for the struct
// [RankingInternetServiceTopResponse]
type rankingInternetServiceTopResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopResponseMeta struct {
	ConfidenceInfo RankingInternetServiceTopResponseMetaConfidenceInfo `json:"confidenceInfo,required,nullable"`
	DateRange      []RankingInternetServiceTopResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization RankingInternetServiceTopResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []RankingInternetServiceTopResponseMetaUnit `json:"units,required"`
	JSON  rankingInternetServiceTopResponseMetaJSON   `json:"-"`
}

// rankingInternetServiceTopResponseMetaJSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseMeta]
type rankingInternetServiceTopResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopResponseMetaConfidenceInfo struct {
	Annotations []RankingInternetServiceTopResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                   `json:"level,required"`
	JSON  rankingInternetServiceTopResponseMetaConfidenceInfoJSON `json:"-"`
}

// rankingInternetServiceTopResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RankingInternetServiceTopResponseMetaConfidenceInfo]
type rankingInternetServiceTopResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type RankingInternetServiceTopResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	LinkedURL       string                                                            `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                         `json:"startDate,required" format:"date-time"`
	JSON            rankingInternetServiceTopResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// rankingInternetServiceTopResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RankingInternetServiceTopResponseMetaConfidenceInfoAnnotation]
type rankingInternetServiceTopResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RankingInternetServiceTopResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      rankingInternetServiceTopResponseMetaDateRangeJSON `json:"-"`
}

// rankingInternetServiceTopResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RankingInternetServiceTopResponseMetaDateRange]
type rankingInternetServiceTopResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RankingInternetServiceTopResponseMetaNormalization string

const (
	RankingInternetServiceTopResponseMetaNormalizationPercentage           RankingInternetServiceTopResponseMetaNormalization = "PERCENTAGE"
	RankingInternetServiceTopResponseMetaNormalizationMin0Max              RankingInternetServiceTopResponseMetaNormalization = "MIN0_MAX"
	RankingInternetServiceTopResponseMetaNormalizationMinMax               RankingInternetServiceTopResponseMetaNormalization = "MIN_MAX"
	RankingInternetServiceTopResponseMetaNormalizationRawValues            RankingInternetServiceTopResponseMetaNormalization = "RAW_VALUES"
	RankingInternetServiceTopResponseMetaNormalizationPercentageChange     RankingInternetServiceTopResponseMetaNormalization = "PERCENTAGE_CHANGE"
	RankingInternetServiceTopResponseMetaNormalizationRollingAverage       RankingInternetServiceTopResponseMetaNormalization = "ROLLING_AVERAGE"
	RankingInternetServiceTopResponseMetaNormalizationOverlappedPercentage RankingInternetServiceTopResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	RankingInternetServiceTopResponseMetaNormalizationRatio                RankingInternetServiceTopResponseMetaNormalization = "RATIO"
)

func (r RankingInternetServiceTopResponseMetaNormalization) IsKnown() bool {
	switch r {
	case RankingInternetServiceTopResponseMetaNormalizationPercentage, RankingInternetServiceTopResponseMetaNormalizationMin0Max, RankingInternetServiceTopResponseMetaNormalizationMinMax, RankingInternetServiceTopResponseMetaNormalizationRawValues, RankingInternetServiceTopResponseMetaNormalizationPercentageChange, RankingInternetServiceTopResponseMetaNormalizationRollingAverage, RankingInternetServiceTopResponseMetaNormalizationOverlappedPercentage, RankingInternetServiceTopResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type RankingInternetServiceTopResponseMetaUnit struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  rankingInternetServiceTopResponseMetaUnitJSON `json:"-"`
}

// rankingInternetServiceTopResponseMetaUnitJSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseMetaUnit]
type rankingInternetServiceTopResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopResponseTop0 struct {
	Rank    int64                                     `json:"rank,required"`
	Service string                                    `json:"service,required"`
	JSON    rankingInternetServiceTopResponseTop0JSON `json:"-"`
}

// rankingInternetServiceTopResponseTop0JSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseTop0]
type rankingInternetServiceTopResponseTop0JSON struct {
	Rank        apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceCategoriesParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RankingInternetServiceCategoriesParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RankingInternetServiceCategoriesParams]'s query parameters
// as `url.Values`.
func (r RankingInternetServiceCategoriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type RankingInternetServiceCategoriesParamsFormat string

const (
	RankingInternetServiceCategoriesParamsFormatJson RankingInternetServiceCategoriesParamsFormat = "JSON"
	RankingInternetServiceCategoriesParamsFormatCsv  RankingInternetServiceCategoriesParamsFormat = "CSV"
)

func (r RankingInternetServiceCategoriesParamsFormat) IsKnown() bool {
	switch r {
	case RankingInternetServiceCategoriesParamsFormatJson, RankingInternetServiceCategoriesParamsFormatCsv:
		return true
	}
	return false
}

type RankingInternetServiceCategoriesResponseEnvelope struct {
	Result  RankingInternetServiceCategoriesResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    rankingInternetServiceCategoriesResponseEnvelopeJSON `json:"-"`
}

// rankingInternetServiceCategoriesResponseEnvelopeJSON contains the JSON metadata
// for the struct [RankingInternetServiceCategoriesResponseEnvelope]
type rankingInternetServiceCategoriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceCategoriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceCategoriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTimeseriesGroupsParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[RankingInternetServiceTimeseriesGroupsParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by Internet service category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RankingInternetServiceTimeseriesGroupsParams]'s query
// parameters as `url.Values`.
func (r RankingInternetServiceTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type RankingInternetServiceTimeseriesGroupsParamsFormat string

const (
	RankingInternetServiceTimeseriesGroupsParamsFormatJson RankingInternetServiceTimeseriesGroupsParamsFormat = "JSON"
	RankingInternetServiceTimeseriesGroupsParamsFormatCsv  RankingInternetServiceTimeseriesGroupsParamsFormat = "CSV"
)

func (r RankingInternetServiceTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case RankingInternetServiceTimeseriesGroupsParamsFormatJson, RankingInternetServiceTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type RankingInternetServiceTimeseriesGroupsResponseEnvelope struct {
	Result  RankingInternetServiceTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON contains the JSON
// metadata for the struct [RankingInternetServiceTimeseriesGroupsResponseEnvelope]
type rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RankingInternetServiceTopParams struct {
	// Filters results by the specified array of dates.
	Date param.Field[[]time.Time] `query:"date" format:"date"`
	// Format in which results will be returned.
	Format param.Field[RankingInternetServiceTopParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by Internet service category.
	ServiceCategory param.Field[[]string] `query:"serviceCategory"`
}

// URLQuery serializes [RankingInternetServiceTopParams]'s query parameters as
// `url.Values`.
func (r RankingInternetServiceTopParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type RankingInternetServiceTopParamsFormat string

const (
	RankingInternetServiceTopParamsFormatJson RankingInternetServiceTopParamsFormat = "JSON"
	RankingInternetServiceTopParamsFormatCsv  RankingInternetServiceTopParamsFormat = "CSV"
)

func (r RankingInternetServiceTopParamsFormat) IsKnown() bool {
	switch r {
	case RankingInternetServiceTopParamsFormatJson, RankingInternetServiceTopParamsFormatCsv:
		return true
	}
	return false
}

type RankingInternetServiceTopResponseEnvelope struct {
	Result  RankingInternetServiceTopResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    rankingInternetServiceTopResponseEnvelopeJSON `json:"-"`
}

// rankingInternetServiceTopResponseEnvelopeJSON contains the JSON metadata for the
// struct [RankingInternetServiceTopResponseEnvelope]
type rankingInternetServiceTopResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RankingInternetServiceTopResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rankingInternetServiceTopResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
