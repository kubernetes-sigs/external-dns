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

// AIInferenceTimeseriesGroupSummaryService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIInferenceTimeseriesGroupSummaryService] method instead.
type AIInferenceTimeseriesGroupSummaryService struct {
	Options []option.RequestOption
}

// NewAIInferenceTimeseriesGroupSummaryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIInferenceTimeseriesGroupSummaryService(opts ...option.RequestOption) (r *AIInferenceTimeseriesGroupSummaryService) {
	r = &AIInferenceTimeseriesGroupSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of unique accounts by model over time.
func (r *AIInferenceTimeseriesGroupSummaryService) Model(ctx context.Context, query AIInferenceTimeseriesGroupSummaryModelParams, opts ...option.RequestOption) (res *AIInferenceTimeseriesGroupSummaryModelResponse, err error) {
	var env AIInferenceTimeseriesGroupSummaryModelResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of unique accounts by task over time.
func (r *AIInferenceTimeseriesGroupSummaryService) Task(ctx context.Context, query AIInferenceTimeseriesGroupSummaryTaskParams, opts ...option.RequestOption) (res *AIInferenceTimeseriesGroupSummaryTaskResponse, err error) {
	var env AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/timeseries_groups/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIInferenceTimeseriesGroupSummaryModelResponse struct {
	// Metadata for the results.
	Meta   AIInferenceTimeseriesGroupSummaryModelResponseMeta   `json:"meta,required"`
	Serie0 AIInferenceTimeseriesGroupSummaryModelResponseSerie0 `json:"serie_0,required"`
	JSON   aiInferenceTimeseriesGroupSummaryModelResponseJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseJSON contains the JSON metadata
// for the struct [AIInferenceTimeseriesGroupSummaryModelResponse]
type aiInferenceTimeseriesGroupSummaryModelResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIInferenceTimeseriesGroupSummaryModelResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIInferenceTimeseriesGroupSummaryModelResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIInferenceTimeseriesGroupSummaryModelResponseMetaUnit `json:"units,required"`
	JSON  aiInferenceTimeseriesGroupSummaryModelResponseMetaJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseMetaJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryModelResponseMeta]
type aiInferenceTimeseriesGroupSummaryModelResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval string

const (
	AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalFifteenMinutes AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneHour        AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval = "ONE_HOUR"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneDay         AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval = "ONE_DAY"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneWeek        AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval = "ONE_WEEK"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneMonth       AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIInferenceTimeseriesGroupSummaryModelResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalFifteenMinutes, AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneHour, AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneDay, AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneWeek, AIInferenceTimeseriesGroupSummaryModelResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                                `json:"level,required"`
	JSON  aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfo]
type aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                      `json:"startDate,required" format:"date-time"`
	JSON            aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotation]
type aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryModelResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceTimeseriesGroupSummaryModelResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [AIInferenceTimeseriesGroupSummaryModelResponseMetaDateRange]
type aiInferenceTimeseriesGroupSummaryModelResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization string

const (
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationPercentage           AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "PERCENTAGE"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationMin0Max              AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "MIN0_MAX"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationMinMax               AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "MIN_MAX"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationRawValues            AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "RAW_VALUES"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationPercentageChange     AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationRollingAverage       AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "ROLLING_AVERAGE"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationOverlappedPercentage AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationRatio                AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization = "RATIO"
)

func (r AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationPercentage, AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationMin0Max, AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationMinMax, AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationRawValues, AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationPercentageChange, AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationRollingAverage, AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationOverlappedPercentage, AIInferenceTimeseriesGroupSummaryModelResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryModelResponseMetaUnit struct {
	Name  string                                                     `json:"name,required"`
	Value string                                                     `json:"value,required"`
	JSON  aiInferenceTimeseriesGroupSummaryModelResponseMetaUnitJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseMetaUnitJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryModelResponseMetaUnit]
type aiInferenceTimeseriesGroupSummaryModelResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryModelResponseSerie0 struct {
	Timestamps  []time.Time                                              `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                      `json:"-,extras"`
	JSON        aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryModelResponseSerie0]
type aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryTaskResponse struct {
	// Metadata for the results.
	Meta   AIInferenceTimeseriesGroupSummaryTaskResponseMeta   `json:"meta,required"`
	Serie0 AIInferenceTimeseriesGroupSummaryTaskResponseSerie0 `json:"serie_0,required"`
	JSON   aiInferenceTimeseriesGroupSummaryTaskResponseJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseJSON contains the JSON metadata for
// the struct [AIInferenceTimeseriesGroupSummaryTaskResponse]
type aiInferenceTimeseriesGroupSummaryTaskResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIInferenceTimeseriesGroupSummaryTaskResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIInferenceTimeseriesGroupSummaryTaskResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIInferenceTimeseriesGroupSummaryTaskResponseMetaUnit `json:"units,required"`
	JSON  aiInferenceTimeseriesGroupSummaryTaskResponseMetaJSON   `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseMetaJSON contains the JSON metadata
// for the struct [AIInferenceTimeseriesGroupSummaryTaskResponseMeta]
type aiInferenceTimeseriesGroupSummaryTaskResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval string

const (
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalFifteenMinutes AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval = "FIFTEEN_MINUTES"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneHour        AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval = "ONE_HOUR"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneDay         AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval = "ONE_DAY"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneWeek        AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval = "ONE_WEEK"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneMonth       AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval = "ONE_MONTH"
)

func (r AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalFifteenMinutes, AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneHour, AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneDay, AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneWeek, AIInferenceTimeseriesGroupSummaryTaskResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                               `json:"level,required"`
	JSON  aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfo]
type aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                                     `json:"startDate,required" format:"date-time"`
	JSON            aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotation]
type aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryTaskResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceTimeseriesGroupSummaryTaskResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [AIInferenceTimeseriesGroupSummaryTaskResponseMetaDateRange]
type aiInferenceTimeseriesGroupSummaryTaskResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization string

const (
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationPercentage           AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "PERCENTAGE"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationMin0Max              AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "MIN0_MAX"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationMinMax               AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "MIN_MAX"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationRawValues            AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "RAW_VALUES"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationPercentageChange     AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationRollingAverage       AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "ROLLING_AVERAGE"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationOverlappedPercentage AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationRatio                AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization = "RATIO"
)

func (r AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationPercentage, AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationMin0Max, AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationMinMax, AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationRawValues, AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationPercentageChange, AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationRollingAverage, AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationOverlappedPercentage, AIInferenceTimeseriesGroupSummaryTaskResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryTaskResponseMetaUnit struct {
	Name  string                                                    `json:"name,required"`
	Value string                                                    `json:"value,required"`
	JSON  aiInferenceTimeseriesGroupSummaryTaskResponseMetaUnitJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseMetaUnitJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryTaskResponseMetaUnit]
type aiInferenceTimeseriesGroupSummaryTaskResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryTaskResponseSerie0 struct {
	Timestamps  []time.Time                                             `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                                     `json:"-,extras"`
	JSON        aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryTaskResponseSerie0]
type aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryModelParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIInferenceTimeseriesGroupSummaryModelParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceTimeseriesGroupSummaryModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceTimeseriesGroupSummaryModelParams]'s query
// parameters as `url.Values`.
func (r AIInferenceTimeseriesGroupSummaryModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupSummaryModelParamsAggInterval string

const (
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval15m AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "15m"
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1h  AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "1h"
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1d  AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "1d"
	AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1w  AIInferenceTimeseriesGroupSummaryModelParamsAggInterval = "1w"
)

func (r AIInferenceTimeseriesGroupSummaryModelParamsAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryModelParamsAggInterval15m, AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1h, AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1d, AIInferenceTimeseriesGroupSummaryModelParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIInferenceTimeseriesGroupSummaryModelParamsFormat string

const (
	AIInferenceTimeseriesGroupSummaryModelParamsFormatJson AIInferenceTimeseriesGroupSummaryModelParamsFormat = "JSON"
	AIInferenceTimeseriesGroupSummaryModelParamsFormatCsv  AIInferenceTimeseriesGroupSummaryModelParamsFormat = "CSV"
)

func (r AIInferenceTimeseriesGroupSummaryModelParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryModelParamsFormatJson, AIInferenceTimeseriesGroupSummaryModelParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryModelResponseEnvelope struct {
	Result  AIInferenceTimeseriesGroupSummaryModelResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryModelResponseEnvelope]
type aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryModelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIInferenceTimeseriesGroupSummaryTaskParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceTimeseriesGroupSummaryTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceTimeseriesGroupSummaryTaskParams]'s query
// parameters as `url.Values`.
func (r AIInferenceTimeseriesGroupSummaryTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval string

const (
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval15m AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "15m"
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1h  AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "1h"
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1d  AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "1d"
	AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1w  AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval = "1w"
)

func (r AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval15m, AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1h, AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1d, AIInferenceTimeseriesGroupSummaryTaskParamsAggInterval1w:
		return true
	}
	return false
}

// Format in which results will be returned.
type AIInferenceTimeseriesGroupSummaryTaskParamsFormat string

const (
	AIInferenceTimeseriesGroupSummaryTaskParamsFormatJson AIInferenceTimeseriesGroupSummaryTaskParamsFormat = "JSON"
	AIInferenceTimeseriesGroupSummaryTaskParamsFormatCsv  AIInferenceTimeseriesGroupSummaryTaskParamsFormat = "CSV"
)

func (r AIInferenceTimeseriesGroupSummaryTaskParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceTimeseriesGroupSummaryTaskParamsFormatJson, AIInferenceTimeseriesGroupSummaryTaskParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope struct {
	Result  AIInferenceTimeseriesGroupSummaryTaskResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON `json:"-"`
}

// aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON contains the JSON
// metadata for the struct [AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope]
type aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceTimeseriesGroupSummaryTaskResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceTimeseriesGroupSummaryTaskResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
