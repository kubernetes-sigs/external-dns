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

// AIBotSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIBotSummaryService] method instead.
type AIBotSummaryService struct {
	Options []option.RequestOption
}

// NewAIBotSummaryService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIBotSummaryService(opts ...option.RequestOption) (r *AIBotSummaryService) {
	r = &AIBotSummaryService{}
	r.Options = opts
	return
}

// Retrieves the distribution of traffic by AI user agent.
//
// Deprecated: deprecated
func (r *AIBotSummaryService) UserAgent(ctx context.Context, query AIBotSummaryUserAgentParams, opts ...option.RequestOption) (res *AIBotSummaryUserAgentResponse, err error) {
	var env AIBotSummaryUserAgentResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/ai/bots/summary/user_agent"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIBotSummaryUserAgentResponse struct {
	// Metadata for the results.
	Meta     AIBotSummaryUserAgentResponseMeta `json:"meta,required"`
	Summary0 map[string]string                 `json:"summary_0,required"`
	JSON     aiBotSummaryUserAgentResponseJSON `json:"-"`
}

// aiBotSummaryUserAgentResponseJSON contains the JSON metadata for the struct
// [AIBotSummaryUserAgentResponse]
type aiBotSummaryUserAgentResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryUserAgentResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryUserAgentResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type AIBotSummaryUserAgentResponseMeta struct {
	ConfidenceInfo AIBotSummaryUserAgentResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []AIBotSummaryUserAgentResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization AIBotSummaryUserAgentResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []AIBotSummaryUserAgentResponseMetaUnit `json:"units,required"`
	JSON  aiBotSummaryUserAgentResponseMetaJSON   `json:"-"`
}

// aiBotSummaryUserAgentResponseMetaJSON contains the JSON metadata for the struct
// [AIBotSummaryUserAgentResponseMeta]
type aiBotSummaryUserAgentResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AIBotSummaryUserAgentResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryUserAgentResponseMetaJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryUserAgentResponseMetaConfidenceInfo struct {
	Annotations []AIBotSummaryUserAgentResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                               `json:"level,required"`
	JSON  aiBotSummaryUserAgentResponseMetaConfidenceInfoJSON `json:"-"`
}

// aiBotSummaryUserAgentResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [AIBotSummaryUserAgentResponseMetaConfidenceInfo]
type aiBotSummaryUserAgentResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryUserAgentResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryUserAgentResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type AIBotSummaryUserAgentResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                          `json:"isInstantaneous,required"`
	LinkedURL       string                                                        `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                     `json:"startDate,required" format:"date-time"`
	JSON            aiBotSummaryUserAgentResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// aiBotSummaryUserAgentResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [AIBotSummaryUserAgentResponseMetaConfidenceInfoAnnotation]
type aiBotSummaryUserAgentResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *AIBotSummaryUserAgentResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryUserAgentResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryUserAgentResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                      `json:"startTime,required" format:"date-time"`
	JSON      aiBotSummaryUserAgentResponseMetaDateRangeJSON `json:"-"`
}

// aiBotSummaryUserAgentResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [AIBotSummaryUserAgentResponseMetaDateRange]
type aiBotSummaryUserAgentResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryUserAgentResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryUserAgentResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type AIBotSummaryUserAgentResponseMetaNormalization string

const (
	AIBotSummaryUserAgentResponseMetaNormalizationPercentage           AIBotSummaryUserAgentResponseMetaNormalization = "PERCENTAGE"
	AIBotSummaryUserAgentResponseMetaNormalizationMin0Max              AIBotSummaryUserAgentResponseMetaNormalization = "MIN0_MAX"
	AIBotSummaryUserAgentResponseMetaNormalizationMinMax               AIBotSummaryUserAgentResponseMetaNormalization = "MIN_MAX"
	AIBotSummaryUserAgentResponseMetaNormalizationRawValues            AIBotSummaryUserAgentResponseMetaNormalization = "RAW_VALUES"
	AIBotSummaryUserAgentResponseMetaNormalizationPercentageChange     AIBotSummaryUserAgentResponseMetaNormalization = "PERCENTAGE_CHANGE"
	AIBotSummaryUserAgentResponseMetaNormalizationRollingAverage       AIBotSummaryUserAgentResponseMetaNormalization = "ROLLING_AVERAGE"
	AIBotSummaryUserAgentResponseMetaNormalizationOverlappedPercentage AIBotSummaryUserAgentResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	AIBotSummaryUserAgentResponseMetaNormalizationRatio                AIBotSummaryUserAgentResponseMetaNormalization = "RATIO"
)

func (r AIBotSummaryUserAgentResponseMetaNormalization) IsKnown() bool {
	switch r {
	case AIBotSummaryUserAgentResponseMetaNormalizationPercentage, AIBotSummaryUserAgentResponseMetaNormalizationMin0Max, AIBotSummaryUserAgentResponseMetaNormalizationMinMax, AIBotSummaryUserAgentResponseMetaNormalizationRawValues, AIBotSummaryUserAgentResponseMetaNormalizationPercentageChange, AIBotSummaryUserAgentResponseMetaNormalizationRollingAverage, AIBotSummaryUserAgentResponseMetaNormalizationOverlappedPercentage, AIBotSummaryUserAgentResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type AIBotSummaryUserAgentResponseMetaUnit struct {
	Name  string                                    `json:"name,required"`
	Value string                                    `json:"value,required"`
	JSON  aiBotSummaryUserAgentResponseMetaUnitJSON `json:"-"`
}

// aiBotSummaryUserAgentResponseMetaUnitJSON contains the JSON metadata for the
// struct [AIBotSummaryUserAgentResponseMetaUnit]
type aiBotSummaryUserAgentResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryUserAgentResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryUserAgentResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type AIBotSummaryUserAgentParams struct {
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
	Format param.Field[AIBotSummaryUserAgentParamsFormat] `query:"format"`
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

// URLQuery serializes [AIBotSummaryUserAgentParams]'s query parameters as
// `url.Values`.
func (r AIBotSummaryUserAgentParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type AIBotSummaryUserAgentParamsFormat string

const (
	AIBotSummaryUserAgentParamsFormatJson AIBotSummaryUserAgentParamsFormat = "JSON"
	AIBotSummaryUserAgentParamsFormatCsv  AIBotSummaryUserAgentParamsFormat = "CSV"
)

func (r AIBotSummaryUserAgentParamsFormat) IsKnown() bool {
	switch r {
	case AIBotSummaryUserAgentParamsFormatJson, AIBotSummaryUserAgentParamsFormatCsv:
		return true
	}
	return false
}

type AIBotSummaryUserAgentResponseEnvelope struct {
	Result  AIBotSummaryUserAgentResponse             `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    aiBotSummaryUserAgentResponseEnvelopeJSON `json:"-"`
}

// aiBotSummaryUserAgentResponseEnvelopeJSON contains the JSON metadata for the
// struct [AIBotSummaryUserAgentResponseEnvelope]
type aiBotSummaryUserAgentResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIBotSummaryUserAgentResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiBotSummaryUserAgentResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
