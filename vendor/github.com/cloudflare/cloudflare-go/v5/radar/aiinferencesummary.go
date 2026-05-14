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

// AIInferenceSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIInferenceSummaryService] method instead.
type AIInferenceSummaryService struct {
	Options []option.RequestOption
}

// NewAIInferenceSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIInferenceSummaryService(opts ...option.RequestOption) (r *AIInferenceSummaryService) {
	r = &AIInferenceSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of unique accounts by model.
func (r *AIInferenceSummaryService) Model(ctx context.Context, query AIInferenceSummaryModelParams, opts ...option.RequestOption) (res *AIInferenceSummaryModelResponse, err error) {
	var env AIInferenceSummaryModelResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/model"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of unique accounts by task.
func (r *AIInferenceSummaryService) Task(ctx context.Context, query AIInferenceSummaryTaskParams, opts ...option.RequestOption) (res *AIInferenceSummaryTaskResponse, err error) {
	var env AIInferenceSummaryTaskResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/inference/summary/task"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIInferenceSummaryModelResponse struct {
	// Metadata for the results.
	Meta     AIInferenceSummaryModelResponseMeta `json:"meta,required"`
	Summary0 map[string]string                   `json:"summary_0,required"`
	JSON     aiInferenceSummaryModelResponseJSON `json:"-"`
}

// aiInferenceSummaryModelResponseJSON contains the JSON metadata for the struct
// [AIInferenceSummaryModelResponse]
type aiInferenceSummaryModelResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIInferenceSummaryModelResponseMeta struct {
	ConfidenceInfo AIInferenceSummaryModelResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIInferenceSummaryModelResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIInferenceSummaryModelResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIInferenceSummaryModelResponseMetaUnit `json:"units,required"`
	JSON  aiInferenceSummaryModelResponseMetaJSON   `json:"-"`
}

// aiInferenceSummaryModelResponseMetaJSON contains the JSON metadata for the
// struct [AIInferenceSummaryModelResponseMeta]
type aiInferenceSummaryModelResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                 `json:"level,required"`
	JSON  aiInferenceSummaryModelResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiInferenceSummaryModelResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIInferenceSummaryModelResponseMetaConfidenceInfo]
type aiInferenceSummaryModelResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	LinkedURL       string                                                          `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                       `json:"startDate,required" format:"date-time"`
	JSON            aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation]
type aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceSummaryModelResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceSummaryModelResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceSummaryModelResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIInferenceSummaryModelResponseMetaDateRange]
type aiInferenceSummaryModelResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIInferenceSummaryModelResponseMetaNormalization string

const (
	AIInferenceSummaryModelResponseMetaNormalizationPercentage           AIInferenceSummaryModelResponseMetaNormalization = "PERCENTAGE"
	AIInferenceSummaryModelResponseMetaNormalizationMin0Max              AIInferenceSummaryModelResponseMetaNormalization = "MIN0_MAX"
	AIInferenceSummaryModelResponseMetaNormalizationMinMax               AIInferenceSummaryModelResponseMetaNormalization = "MIN_MAX"
	AIInferenceSummaryModelResponseMetaNormalizationRawValues            AIInferenceSummaryModelResponseMetaNormalization = "RAW_VALUES"
	AIInferenceSummaryModelResponseMetaNormalizationPercentageChange     AIInferenceSummaryModelResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIInferenceSummaryModelResponseMetaNormalizationRollingAverage       AIInferenceSummaryModelResponseMetaNormalization = "ROLLING_AVERAGE"
	AIInferenceSummaryModelResponseMetaNormalizationOverlappedPercentage AIInferenceSummaryModelResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIInferenceSummaryModelResponseMetaNormalizationRatio                AIInferenceSummaryModelResponseMetaNormalization = "RATIO"
)

func (r AIInferenceSummaryModelResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIInferenceSummaryModelResponseMetaNormalizationPercentage, AIInferenceSummaryModelResponseMetaNormalizationMin0Max, AIInferenceSummaryModelResponseMetaNormalizationMinMax, AIInferenceSummaryModelResponseMetaNormalizationRawValues, AIInferenceSummaryModelResponseMetaNormalizationPercentageChange, AIInferenceSummaryModelResponseMetaNormalizationRollingAverage, AIInferenceSummaryModelResponseMetaNormalizationOverlappedPercentage, AIInferenceSummaryModelResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIInferenceSummaryModelResponseMetaUnit struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  aiInferenceSummaryModelResponseMetaUnitJSON `json:"-"`
}

// aiInferenceSummaryModelResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIInferenceSummaryModelResponseMetaUnit]
type aiInferenceSummaryModelResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponse struct {
	// Metadata for the results.
	Meta     AIInferenceSummaryTaskResponseMeta `json:"meta,required"`
	Summary0 map[string]string                  `json:"summary_0,required"`
	JSON     aiInferenceSummaryTaskResponseJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseJSON contains the JSON metadata for the struct
// [AIInferenceSummaryTaskResponse]
type aiInferenceSummaryTaskResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIInferenceSummaryTaskResponseMeta struct {
	ConfidenceInfo AIInferenceSummaryTaskResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIInferenceSummaryTaskResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIInferenceSummaryTaskResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIInferenceSummaryTaskResponseMetaUnit `json:"units,required"`
	JSON  aiInferenceSummaryTaskResponseMetaJSON   `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaJSON contains the JSON metadata for the struct
// [AIInferenceSummaryTaskResponseMeta]
type aiInferenceSummaryTaskResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponseMetaConfidenceInfo struct {
	Annotations []AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                                `json:"level,required"`
	JSON  aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIInferenceSummaryTaskResponseMetaConfidenceInfo]
type aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                           `json:"isInstantaneous,required"`
	LinkedURL       string                                                         `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                      `json:"startDate,required" format:"date-time"`
	JSON            aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation]
type aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIInferenceSummaryTaskResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      aiInferenceSummaryTaskResponseMetaDateRangeJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIInferenceSummaryTaskResponseMetaDateRange]
type aiInferenceSummaryTaskResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIInferenceSummaryTaskResponseMetaNormalization string

const (
	AIInferenceSummaryTaskResponseMetaNormalizationPercentage           AIInferenceSummaryTaskResponseMetaNormalization = "PERCENTAGE"
	AIInferenceSummaryTaskResponseMetaNormalizationMin0Max              AIInferenceSummaryTaskResponseMetaNormalization = "MIN0_MAX"
	AIInferenceSummaryTaskResponseMetaNormalizationMinMax               AIInferenceSummaryTaskResponseMetaNormalization = "MIN_MAX"
	AIInferenceSummaryTaskResponseMetaNormalizationRawValues            AIInferenceSummaryTaskResponseMetaNormalization = "RAW_VALUES"
	AIInferenceSummaryTaskResponseMetaNormalizationPercentageChange     AIInferenceSummaryTaskResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIInferenceSummaryTaskResponseMetaNormalizationRollingAverage       AIInferenceSummaryTaskResponseMetaNormalization = "ROLLING_AVERAGE"
	AIInferenceSummaryTaskResponseMetaNormalizationOverlappedPercentage AIInferenceSummaryTaskResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIInferenceSummaryTaskResponseMetaNormalizationRatio                AIInferenceSummaryTaskResponseMetaNormalization = "RATIO"
)

func (r AIInferenceSummaryTaskResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIInferenceSummaryTaskResponseMetaNormalizationPercentage, AIInferenceSummaryTaskResponseMetaNormalizationMin0Max, AIInferenceSummaryTaskResponseMetaNormalizationMinMax, AIInferenceSummaryTaskResponseMetaNormalizationRawValues, AIInferenceSummaryTaskResponseMetaNormalizationPercentageChange, AIInferenceSummaryTaskResponseMetaNormalizationRollingAverage, AIInferenceSummaryTaskResponseMetaNormalizationOverlappedPercentage, AIInferenceSummaryTaskResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIInferenceSummaryTaskResponseMetaUnit struct {
	Name  string                                     `json:"name,required"`
	Value string                                     `json:"value,required"`
	JSON  aiInferenceSummaryTaskResponseMetaUnitJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIInferenceSummaryTaskResponseMetaUnit]
type aiInferenceSummaryTaskResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryModelParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceSummaryModelParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceSummaryModelParams]'s query parameters as
// `url.Values`.
func (r AIInferenceSummaryModelParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AIInferenceSummaryModelParamsFormat string

const (
	AIInferenceSummaryModelParamsFormatJson AIInferenceSummaryModelParamsFormat = "JSON"
	AIInferenceSummaryModelParamsFormatCsv  AIInferenceSummaryModelParamsFormat = "CSV"
)

func (r AIInferenceSummaryModelParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceSummaryModelParamsFormatJson, AIInferenceSummaryModelParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceSummaryModelResponseEnvelope struct {
	Result  AIInferenceSummaryModelResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    aiInferenceSummaryModelResponseEnvelopeJSON `json:"-"`
}

// aiInferenceSummaryModelResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIInferenceSummaryModelResponseEnvelope]
type aiInferenceSummaryModelResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryModelResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryModelResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIInferenceSummaryTaskParams struct {
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format in which results will be returned.
	Format param.Field[AIInferenceSummaryTaskParamsFormat] `query:"format"`
	// Limits the number of objects per group to the top items within the specified
	// time range. When item count exceeds the limit, extra items appear grouped under
	// an "other" category.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [AIInferenceSummaryTaskParams]'s query parameters as
// `url.Values`.
func (r AIInferenceSummaryTaskParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AIInferenceSummaryTaskParamsFormat string

const (
	AIInferenceSummaryTaskParamsFormatJson AIInferenceSummaryTaskParamsFormat = "JSON"
	AIInferenceSummaryTaskParamsFormatCsv  AIInferenceSummaryTaskParamsFormat = "CSV"
)

func (r AIInferenceSummaryTaskParamsFormat) IsKnown() bool {
	switch r {
	case AIInferenceSummaryTaskParamsFormatJson, AIInferenceSummaryTaskParamsFormatCsv:
		return true
	}
	return false
}

type AIInferenceSummaryTaskResponseEnvelope struct {
	Result  AIInferenceSummaryTaskResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    aiInferenceSummaryTaskResponseEnvelopeJSON `json:"-"`
}

// aiInferenceSummaryTaskResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIInferenceSummaryTaskResponseEnvelope]
type aiInferenceSummaryTaskResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIInferenceSummaryTaskResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiInferenceSummaryTaskResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
