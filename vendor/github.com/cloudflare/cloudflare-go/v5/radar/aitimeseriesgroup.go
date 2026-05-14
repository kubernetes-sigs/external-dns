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

// AITimeseriesGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITimeseriesGroupService] method instead.
type AITimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewAITimeseriesGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITimeseriesGroupService(opts ...option.RequestOption) (r *AITimeseriesGroupService) {
	r = &AITimeseriesGroupService{}
	r.Options = opts
	return
}

// Retrieves an aggregated summary of AI bots HTTP requests grouped by the
// specified dimension.
func (r *AITimeseriesGroupService) Summary(ctx context.Context, dimension AITimeseriesGroupSummaryParamsDimension, query AITimeseriesGroupSummaryParams, opts ...option.RequestOption) (res *AITimeseriesGroupSummaryResponse, err error) {
	var env AITimeseriesGroupSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ai/bots/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves AI bots HTTP request volume over time.
func (r *AITimeseriesGroupService) Timeseries(ctx context.Context, query AITimeseriesGroupTimeseriesParams, opts ...option.RequestOption) (res *AITimeseriesGroupTimeseriesResponse, err error) {
	var env AITimeseriesGroupTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests from AI bots, grouped by chosen the
// specified dimension over time.
func (r *AITimeseriesGroupService) TimeseriesGroups(ctx context.Context, dimension AITimeseriesGroupTimeseriesGroupsParamsDimension, query AITimeseriesGroupTimeseriesGroupsParams, opts ...option.RequestOption) (res *AITimeseriesGroupTimeseriesGroupsResponse, err error) {
	var env AITimeseriesGroupTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/ai/bots/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of traffic by AI user agent over time.
//
// Deprecated: deprecated
func (r *AITimeseriesGroupService) UserAgent(ctx context.Context, query AITimeseriesGroupUserAgentParams, opts ...option.RequestOption) (res *AITimeseriesGroupUserAgentResponse, err error) {
	var env AITimeseriesGroupUserAgentResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/timeseries_groups/user_agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AITimeseriesGroupSummaryResponse struct {
	// Metadata for the results.
	Meta     AITimeseriesGroupSummaryResponseMeta `json:"meta,required"`
	Summary0 map[string]string                    `json:"summary_0,required"`
	JSON     aiTimeseriesGroupSummaryResponseJSON `json:"-"`
}

// aiTimeseriesGroupSummaryResponseJSON contains the JSON metadata for the struct
// [AITimeseriesGroupSummaryResponse]
type aiTimeseriesGroupSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AITimeseriesGroupSummaryResponseMeta struct {
	ConfidenceInfo AITimeseriesGroupSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AITimeseriesGroupSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AITimeseriesGroupSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AITimeseriesGroupSummaryResponseMetaUnit `json:"units,required"`
	JSON  aiTimeseriesGroupSummaryResponseMetaJSON   `json:"-"`
}

// aiTimeseriesGroupSummaryResponseMetaJSON contains the JSON metadata for the
// struct [AITimeseriesGroupSummaryResponseMeta]
type aiTimeseriesGroupSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AITimeseriesGroupSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupSummaryResponseMetaConfidenceInfo struct {
	Annotations []AITimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                  `json:"level,required"`
	JSON  aiTimeseriesGroupSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiTimeseriesGroupSummaryResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AITimeseriesGroupSummaryResponseMetaConfidenceInfo]
type aiTimeseriesGroupSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AITimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	LinkedURL       string                                                           `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                        `json:"startDate,required" format:"date-time"`
	JSON            aiTimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiTimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AITimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotation]
type aiTimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AITimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      aiTimeseriesGroupSummaryResponseMetaDateRangeJSON `json:"-"`
}

// aiTimeseriesGroupSummaryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AITimeseriesGroupSummaryResponseMetaDateRange]
type aiTimeseriesGroupSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AITimeseriesGroupSummaryResponseMetaNormalization string

const (
	AITimeseriesGroupSummaryResponseMetaNormalizationPercentage           AITimeseriesGroupSummaryResponseMetaNormalization = "PERCENTAGE"
	AITimeseriesGroupSummaryResponseMetaNormalizationMin0Max              AITimeseriesGroupSummaryResponseMetaNormalization = "MIN0_MAX"
	AITimeseriesGroupSummaryResponseMetaNormalizationMinMax               AITimeseriesGroupSummaryResponseMetaNormalization = "MIN_MAX"
	AITimeseriesGroupSummaryResponseMetaNormalizationRawValues            AITimeseriesGroupSummaryResponseMetaNormalization = "RAW_VALUES"
	AITimeseriesGroupSummaryResponseMetaNormalizationPercentageChange     AITimeseriesGroupSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AITimeseriesGroupSummaryResponseMetaNormalizationRollingAverage       AITimeseriesGroupSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	AITimeseriesGroupSummaryResponseMetaNormalizationOverlappedPercentage AITimeseriesGroupSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AITimeseriesGroupSummaryResponseMetaNormalizationRatio                AITimeseriesGroupSummaryResponseMetaNormalization = "RATIO"
)

func (r AITimeseriesGroupSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AITimeseriesGroupSummaryResponseMetaNormalizationPercentage, AITimeseriesGroupSummaryResponseMetaNormalizationMin0Max, AITimeseriesGroupSummaryResponseMetaNormalizationMinMax, AITimeseriesGroupSummaryResponseMetaNormalizationRawValues, AITimeseriesGroupSummaryResponseMetaNormalizationPercentageChange, AITimeseriesGroupSummaryResponseMetaNormalizationRollingAverage, AITimeseriesGroupSummaryResponseMetaNormalizationOverlappedPercentage, AITimeseriesGroupSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AITimeseriesGroupSummaryResponseMetaUnit struct {
	Name  string                                       `json:"name,required"`
	Value string                                       `json:"value,required"`
	JSON  aiTimeseriesGroupSummaryResponseMetaUnitJSON `json:"-"`
}

// aiTimeseriesGroupSummaryResponseMetaUnitJSON contains the JSON metadata for the
// struct [AITimeseriesGroupSummaryResponseMetaUnit]
type aiTimeseriesGroupSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupTimeseriesResponse struct {
	// Metadata for the results.
	Meta        AITimeseriesGroupTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]AITimeseriesGroupTimeseriesResponse `json:"-,extras"`
	JSON        aiTimeseriesGroupTimeseriesResponseJSON        `json:"-"`
}

// aiTimeseriesGroupTimeseriesResponseJSON contains the JSON metadata for the
// struct [AITimeseriesGroupTimeseriesResponse]
type aiTimeseriesGroupTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AITimeseriesGroupTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AITimeseriesGroupTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AITimeseriesGroupTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AITimeseriesGroupTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AITimeseriesGroupTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AITimeseriesGroupTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  aiTimeseriesGroupTimeseriesResponseMetaJSON   `json:"-"`
}

// aiTimeseriesGroupTimeseriesResponseMetaJSON contains the JSON metadata for the
// struct [AITimeseriesGroupTimeseriesResponseMeta]
type aiTimeseriesGroupTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupTimeseriesResponseMetaAggInterval string

const (
	AITimeseriesGroupTimeseriesResponseMetaAggIntervalFifteenMinutes AITimeseriesGroupTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneHour        AITimeseriesGroupTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneDay         AITimeseriesGroupTimeseriesResponseMetaAggInterval = "ONE_DAY"
	AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneWeek        AITimeseriesGroupTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneMonth       AITimeseriesGroupTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r AITimeseriesGroupTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesResponseMetaAggIntervalFifteenMinutes, AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneHour, AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneDay, AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneWeek, AITimeseriesGroupTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AITimeseriesGroupTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []AITimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                     `json:"level,required"`
	JSON  aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AITimeseriesGroupTimeseriesResponseMetaConfidenceInfo]
type aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AITimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                `json:"isInstantaneous,required"`
	LinkedURL       string                                                              `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                           `json:"startDate,required" format:"date-time"`
	JSON            aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AITimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotation]
type aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AITimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      aiTimeseriesGroupTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AITimeseriesGroupTimeseriesResponseMetaDateRange]
type aiTimeseriesGroupTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AITimeseriesGroupTimeseriesResponseMetaNormalization string

const (
	AITimeseriesGroupTimeseriesResponseMetaNormalizationPercentage           AITimeseriesGroupTimeseriesResponseMetaNormalization = "PERCENTAGE"
	AITimeseriesGroupTimeseriesResponseMetaNormalizationMin0Max              AITimeseriesGroupTimeseriesResponseMetaNormalization = "MIN0_MAX"
	AITimeseriesGroupTimeseriesResponseMetaNormalizationMinMax               AITimeseriesGroupTimeseriesResponseMetaNormalization = "MIN_MAX"
	AITimeseriesGroupTimeseriesResponseMetaNormalizationRawValues            AITimeseriesGroupTimeseriesResponseMetaNormalization = "RAW_VALUES"
	AITimeseriesGroupTimeseriesResponseMetaNormalizationPercentageChange     AITimeseriesGroupTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AITimeseriesGroupTimeseriesResponseMetaNormalizationRollingAverage       AITimeseriesGroupTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	AITimeseriesGroupTimeseriesResponseMetaNormalizationOverlappedPercentage AITimeseriesGroupTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AITimeseriesGroupTimeseriesResponseMetaNormalizationRatio                AITimeseriesGroupTimeseriesResponseMetaNormalization = "RATIO"
)

func (r AITimeseriesGroupTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesResponseMetaNormalizationPercentage, AITimeseriesGroupTimeseriesResponseMetaNormalizationMin0Max, AITimeseriesGroupTimeseriesResponseMetaNormalizationMinMax, AITimeseriesGroupTimeseriesResponseMetaNormalizationRawValues, AITimeseriesGroupTimeseriesResponseMetaNormalizationPercentageChange, AITimeseriesGroupTimeseriesResponseMetaNormalizationRollingAverage, AITimeseriesGroupTimeseriesResponseMetaNormalizationOverlappedPercentage, AITimeseriesGroupTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AITimeseriesGroupTimeseriesResponseMetaUnit struct {
	Name  string                                          `json:"name,required"`
	Value string                                          `json:"value,required"`
	JSON  aiTimeseriesGroupTimeseriesResponseMetaUnitJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesResponseMetaUnitJSON contains the JSON metadata for
// the struct [AITimeseriesGroupTimeseriesResponseMetaUnit]
type aiTimeseriesGroupTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   AITimeseriesGroupTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 AITimeseriesGroupTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   aiTimeseriesGroupTimeseriesGroupsResponseJSON   `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [AITimeseriesGroupTimeseriesGroupsResponse]
type aiTimeseriesGroupTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AITimeseriesGroupTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AITimeseriesGroupTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AITimeseriesGroupTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  aiTimeseriesGroupTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseMetaJSON contains the JSON metadata for
// the struct [AITimeseriesGroupTimeseriesGroupsResponseMeta]
type aiTimeseriesGroupTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval string

const (
	AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneHour        AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneDay         AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneWeek        AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneMonth       AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r AITimeseriesGroupTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneHour, AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneDay, AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneWeek, AITimeseriesGroupTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                           `json:"level,required"`
	JSON  aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfo]
type aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                      `json:"isInstantaneous,required"`
	LinkedURL       string                                                                    `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                 `json:"startDate,required" format:"date-time"`
	JSON            aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AITimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      aiTimeseriesGroupTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [AITimeseriesGroupTimeseriesGroupsResponseMetaDateRange]
type aiTimeseriesGroupTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization string

const (
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationPercentage           AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationMin0Max              AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationMinMax               AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationRawValues            AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationPercentageChange     AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationRollingAverage       AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationRatio                AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r AITimeseriesGroupTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationPercentage, AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationMin0Max, AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationMinMax, AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationRawValues, AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationPercentageChange, AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationRollingAverage, AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, AITimeseriesGroupTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AITimeseriesGroupTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                                `json:"name,required"`
	Value string                                                `json:"value,required"`
	JSON  aiTimeseriesGroupTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata
// for the struct [AITimeseriesGroupTimeseriesGroupsResponseMetaUnit]
type aiTimeseriesGroupTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                                         `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                 `json:"-,extras"`
	JSON        aiTimeseriesGroupTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseSerie0JSON contains the JSON metadata
// for the struct [AITimeseriesGroupTimeseriesGroupsResponseSerie0]
type aiTimeseriesGroupTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentResponse struct {
	// Metadata for the results.
	Meta   AITimeseriesGroupUserAgentResponseMeta   `json:"meta,required"`
	Serie0 AITimeseriesGroupUserAgentResponseSerie0 `json:"serie_0,required"`
	JSON   aiTimeseriesGroupUserAgentResponseJSON   `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseJSON contains the JSON metadata for the struct
// [AITimeseriesGroupUserAgentResponse]
type aiTimeseriesGroupUserAgentResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AITimeseriesGroupUserAgentResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AITimeseriesGroupUserAgentResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AITimeseriesGroupUserAgentResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AITimeseriesGroupUserAgentResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AITimeseriesGroupUserAgentResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AITimeseriesGroupUserAgentResponseMetaUnit `json:"units,required"`
	JSON  aiTimeseriesGroupUserAgentResponseMetaJSON   `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaJSON contains the JSON metadata for the
// struct [AITimeseriesGroupUserAgentResponseMeta]
type aiTimeseriesGroupUserAgentResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupUserAgentResponseMetaAggInterval string

const (
	AITimeseriesGroupUserAgentResponseMetaAggIntervalFifteenMinutes AITimeseriesGroupUserAgentResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneHour        AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_HOUR"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneDay         AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_DAY"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneWeek        AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_WEEK"
	AITimeseriesGroupUserAgentResponseMetaAggIntervalOneMonth       AITimeseriesGroupUserAgentResponseMetaAggInterval = "ONE_MONTH"
)

func (r AITimeseriesGroupUserAgentResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentResponseMetaAggIntervalFifteenMinutes, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneHour, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneDay, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneWeek, AITimeseriesGroupUserAgentResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AITimeseriesGroupUserAgentResponseMetaConfidenceInfo struct {
	Annotations []AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                    `json:"level,required"`
	JSON  aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [AITimeseriesGroupUserAgentResponseMetaConfidenceInfo]
type aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                          `json:"startDate,required" format:"date-time"`
	JSON            aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation]
type aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AITimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [AITimeseriesGroupUserAgentResponseMetaDateRange]
type aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AITimeseriesGroupUserAgentResponseMetaNormalization string

const (
	AITimeseriesGroupUserAgentResponseMetaNormalizationPercentage           AITimeseriesGroupUserAgentResponseMetaNormalization = "PERCENTAGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationMin0Max              AITimeseriesGroupUserAgentResponseMetaNormalization = "MIN0_MAX"
	AITimeseriesGroupUserAgentResponseMetaNormalizationMinMax               AITimeseriesGroupUserAgentResponseMetaNormalization = "MIN_MAX"
	AITimeseriesGroupUserAgentResponseMetaNormalizationRawValues            AITimeseriesGroupUserAgentResponseMetaNormalization = "RAW_VALUES"
	AITimeseriesGroupUserAgentResponseMetaNormalizationPercentageChange     AITimeseriesGroupUserAgentResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationRollingAverage       AITimeseriesGroupUserAgentResponseMetaNormalization = "ROLLING_AVERAGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationOverlappedPercentage AITimeseriesGroupUserAgentResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AITimeseriesGroupUserAgentResponseMetaNormalizationRatio                AITimeseriesGroupUserAgentResponseMetaNormalization = "RATIO"
)

func (r AITimeseriesGroupUserAgentResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentResponseMetaNormalizationPercentage, AITimeseriesGroupUserAgentResponseMetaNormalizationMin0Max, AITimeseriesGroupUserAgentResponseMetaNormalizationMinMax, AITimeseriesGroupUserAgentResponseMetaNormalizationRawValues, AITimeseriesGroupUserAgentResponseMetaNormalizationPercentageChange, AITimeseriesGroupUserAgentResponseMetaNormalizationRollingAverage, AITimeseriesGroupUserAgentResponseMetaNormalizationOverlappedPercentage, AITimeseriesGroupUserAgentResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AITimeseriesGroupUserAgentResponseMetaUnit struct {
	Name  string                                         `json:"name,required"`
	Value string                                         `json:"value,required"`
	JSON  aiTimeseriesGroupUserAgentResponseMetaUnitJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseMetaUnitJSON contains the JSON metadata for
// the struct [AITimeseriesGroupUserAgentResponseMetaUnit]
type aiTimeseriesGroupUserAgentResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentResponseSerie0 struct {
	Timestamps  []time.Time                                  `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                          `json:"-,extras"`
	JSON        aiTimeseriesGroupUserAgentResponseSerie0JSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseSerie0JSON contains the JSON metadata for the
// struct [AITimeseriesGroupUserAgentResponseSerie0]
type aiTimeseriesGroupUserAgentResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// Filters results by bot crawl purpose.
	CrawlPurpose param.Field[[]string] `query:"crawlPurpose"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AITimeseriesGroupSummaryParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AITimeseriesGroupSummaryParams]'s query parameters as
// `url.Values`.
func (r AITimeseriesGroupSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AITimeseriesGroupSummaryParamsDimension string

const (
	AITimeseriesGroupSummaryParamsDimensionUserAgent    AITimeseriesGroupSummaryParamsDimension = "USER_AGENT"
	AITimeseriesGroupSummaryParamsDimensionCrawlPurpose AITimeseriesGroupSummaryParamsDimension = "CRAWL_PURPOSE"
)

func (r AITimeseriesGroupSummaryParamsDimension) IsKnown() bool {
	switch r {
	case AITimeseriesGroupSummaryParamsDimensionUserAgent, AITimeseriesGroupSummaryParamsDimensionCrawlPurpose:
		return true
	}
	return false
}

// Format in which results will be returned.
type AITimeseriesGroupSummaryParamsFormat string

const (
	AITimeseriesGroupSummaryParamsFormatJson AITimeseriesGroupSummaryParamsFormat = "JSON"
	AITimeseriesGroupSummaryParamsFormatCsv  AITimeseriesGroupSummaryParamsFormat = "CSV"
)

func (r AITimeseriesGroupSummaryParamsFormat) IsKnown() bool {
	switch r {
	case AITimeseriesGroupSummaryParamsFormatJson, AITimeseriesGroupSummaryParamsFormatCsv:
		return true
	}
	return false
}

type AITimeseriesGroupSummaryResponseEnvelope struct {
	Result  AITimeseriesGroupSummaryResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    aiTimeseriesGroupSummaryResponseEnvelopeJSON `json:"-"`
}

// aiTimeseriesGroupSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [AITimeseriesGroupSummaryResponseEnvelope]
type aiTimeseriesGroupSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AITimeseriesGroupTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// Filters results by bot crawl purpose.
	CrawlPurpose param.Field[[]string] `query:"crawlPurpose"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AITimeseriesGroupTimeseriesParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by user agent.
	UserAgent param.Field[[]string] `query:"userAgent"`
}

// URLQuery serializes [AITimeseriesGroupTimeseriesParams]'s query parameters as
// `url.Values`.
func (r AITimeseriesGroupTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupTimeseriesParamsAggInterval string

const (
	AITimeseriesGroupTimeseriesParamsAggInterval15m AITimeseriesGroupTimeseriesParamsAggInterval = "15m"
	AITimeseriesGroupTimeseriesParamsAggInterval1h  AITimeseriesGroupTimeseriesParamsAggInterval = "1h"
	AITimeseriesGroupTimeseriesParamsAggInterval1d  AITimeseriesGroupTimeseriesParamsAggInterval = "1d"
	AITimeseriesGroupTimeseriesParamsAggInterval1w  AITimeseriesGroupTimeseriesParamsAggInterval = "1w"
)

func (r AITimeseriesGroupTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesParamsAggInterval15m, AITimeseriesGroupTimeseriesParamsAggInterval1h, AITimeseriesGroupTimeseriesParamsAggInterval1d, AITimeseriesGroupTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AITimeseriesGroupTimeseriesParamsFormat string

const (
	AITimeseriesGroupTimeseriesParamsFormatJson AITimeseriesGroupTimeseriesParamsFormat = "JSON"
	AITimeseriesGroupTimeseriesParamsFormatCsv  AITimeseriesGroupTimeseriesParamsFormat = "CSV"
)

func (r AITimeseriesGroupTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesParamsFormatJson, AITimeseriesGroupTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type AITimeseriesGroupTimeseriesResponseEnvelope struct {
	Result  AITimeseriesGroupTimeseriesResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    aiTimeseriesGroupTimeseriesResponseEnvelopeJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [AITimeseriesGroupTimeseriesResponseEnvelope]
type aiTimeseriesGroupTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AITimeseriesGroupTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// Filters results by bot crawl purpose.
	CrawlPurpose param.Field[[]string] `query:"crawlPurpose"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AITimeseriesGroupTimeseriesGroupsParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[AITimeseriesGroupTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [AITimeseriesGroupTimeseriesGroupsParams]'s query parameters
// as `url.Values`.
func (r AITimeseriesGroupTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type AITimeseriesGroupTimeseriesGroupsParamsDimension string

const (
	AITimeseriesGroupTimeseriesGroupsParamsDimensionUserAgent    AITimeseriesGroupTimeseriesGroupsParamsDimension = "USER_AGENT"
	AITimeseriesGroupTimeseriesGroupsParamsDimensionCrawlPurpose AITimeseriesGroupTimeseriesGroupsParamsDimension = "CRAWL_PURPOSE"
)

func (r AITimeseriesGroupTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesGroupsParamsDimensionUserAgent, AITimeseriesGroupTimeseriesGroupsParamsDimensionCrawlPurpose:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupTimeseriesGroupsParamsAggInterval string

const (
	AITimeseriesGroupTimeseriesGroupsParamsAggInterval15m AITimeseriesGroupTimeseriesGroupsParamsAggInterval = "15m"
	AITimeseriesGroupTimeseriesGroupsParamsAggInterval1h  AITimeseriesGroupTimeseriesGroupsParamsAggInterval = "1h"
	AITimeseriesGroupTimeseriesGroupsParamsAggInterval1d  AITimeseriesGroupTimeseriesGroupsParamsAggInterval = "1d"
	AITimeseriesGroupTimeseriesGroupsParamsAggInterval1w  AITimeseriesGroupTimeseriesGroupsParamsAggInterval = "1w"
)

func (r AITimeseriesGroupTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesGroupsParamsAggInterval15m, AITimeseriesGroupTimeseriesGroupsParamsAggInterval1h, AITimeseriesGroupTimeseriesGroupsParamsAggInterval1d, AITimeseriesGroupTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AITimeseriesGroupTimeseriesGroupsParamsFormat string

const (
	AITimeseriesGroupTimeseriesGroupsParamsFormatJson AITimeseriesGroupTimeseriesGroupsParamsFormat = "JSON"
	AITimeseriesGroupTimeseriesGroupsParamsFormatCsv  AITimeseriesGroupTimeseriesGroupsParamsFormat = "CSV"
)

func (r AITimeseriesGroupTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesGroupsParamsFormatJson, AITimeseriesGroupTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AITimeseriesGroupTimeseriesGroupsParamsNormalization string

const (
	AITimeseriesGroupTimeseriesGroupsParamsNormalizationPercentageChange AITimeseriesGroupTimeseriesGroupsParamsNormalization = "PERCENTAGE_CHANGE"
	AITimeseriesGroupTimeseriesGroupsParamsNormalizationMin0Max          AITimeseriesGroupTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)

func (r AITimeseriesGroupTimeseriesGroupsParamsNormalization) IsKnown() bool {
	switch r {
	case AITimeseriesGroupTimeseriesGroupsParamsNormalizationPercentageChange, AITimeseriesGroupTimeseriesGroupsParamsNormalizationMin0Max:
		return true
	}
	return false
}

type AITimeseriesGroupTimeseriesGroupsResponseEnvelope struct {
	Result  AITimeseriesGroupTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    aiTimeseriesGroupTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// aiTimeseriesGroupTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata
// for the struct [AITimeseriesGroupTimeseriesGroupsResponseEnvelope]
type aiTimeseriesGroupTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AITimeseriesGroupUserAgentParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AITimeseriesGroupUserAgentParamsAggInterval] `query:"aggInterval"`
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
	Format param.Field[AITimeseriesGroupUserAgentParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AITimeseriesGroupUserAgentParams]'s query parameters as
// `url.Values`.
func (r AITimeseriesGroupUserAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AITimeseriesGroupUserAgentParamsAggInterval string

const (
	AITimeseriesGroupUserAgentParamsAggInterval15m AITimeseriesGroupUserAgentParamsAggInterval = "15m"
	AITimeseriesGroupUserAgentParamsAggInterval1h  AITimeseriesGroupUserAgentParamsAggInterval = "1h"
	AITimeseriesGroupUserAgentParamsAggInterval1d  AITimeseriesGroupUserAgentParamsAggInterval = "1d"
	AITimeseriesGroupUserAgentParamsAggInterval1w  AITimeseriesGroupUserAgentParamsAggInterval = "1w"
)

func (r AITimeseriesGroupUserAgentParamsAggInterval) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentParamsAggInterval15m, AITimeseriesGroupUserAgentParamsAggInterval1h, AITimeseriesGroupUserAgentParamsAggInterval1d, AITimeseriesGroupUserAgentParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AITimeseriesGroupUserAgentParamsFormat string

const (
	AITimeseriesGroupUserAgentParamsFormatJson AITimeseriesGroupUserAgentParamsFormat = "JSON"
	AITimeseriesGroupUserAgentParamsFormatCsv  AITimeseriesGroupUserAgentParamsFormat = "CSV"
)

func (r AITimeseriesGroupUserAgentParamsFormat) IsKnown() bool {
	switch r {
	case AITimeseriesGroupUserAgentParamsFormatJson, AITimeseriesGroupUserAgentParamsFormatCsv:
		return true
	}
	return false
}

type AITimeseriesGroupUserAgentResponseEnvelope struct {
	Result  AITimeseriesGroupUserAgentResponse             `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    aiTimeseriesGroupUserAgentResponseEnvelopeJSON `json:"-"`
}

// aiTimeseriesGroupUserAgentResponseEnvelopeJSON contains the JSON metadata for
// the struct [AITimeseriesGroupUserAgentResponseEnvelope]
type aiTimeseriesGroupUserAgentResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AITimeseriesGroupUserAgentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiTimeseriesGroupUserAgentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
