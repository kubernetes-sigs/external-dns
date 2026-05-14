// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"errors"
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

// BotService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotService] method instead.
type BotService struct {
	Options     []option.RequestOption
	WebCrawlers *BotWebCrawlerService
}

// NewBotService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewBotService(opts ...option.RequestOption) (r *BotService) {
	r = &BotService{}
	r.Options = opts
	r.WebCrawlers = NewBotWebCrawlerService(opts...)
	return
}

// Retrieves a list of bots.
func (r *BotService) List(ctx context.Context, query BotListParams, opts ...option.RequestOption) (res *BotListResponse, err error) {
	var env BotListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the requested bot information.
func (r *BotService) Get(ctx context.Context, botSlug string, query BotGetParams, opts ...option.RequestOption) (res *BotGetResponse, err error) {
	var env BotGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if botSlug == "" {
		err = errors.New("missing required bot_slug parameter")
		return
	}
	path := fmt.Sprintf("radar/bots/%s", botSlug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves an aggregated summary of bots HTTP requests grouped by the specified
// dimension.
func (r *BotService) Summary(ctx context.Context, dimension BotSummaryParamsDimension, query BotSummaryParams, opts ...option.RequestOption) (res *BotSummaryResponse, err error) {
	var env BotSummaryResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/bots/summary/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves bots HTTP request volume over time.
func (r *BotService) Timeseries(ctx context.Context, query BotTimeseriesParams, opts ...option.RequestOption) (res *BotTimeseriesResponse, err error) {
	var env BotTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/bots/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves the distribution of HTTP requests from bots, grouped by chosen the
// specified dimension over time.
func (r *BotService) TimeseriesGroups(ctx context.Context, dimension BotTimeseriesGroupsParamsDimension, query BotTimeseriesGroupsParams, opts ...option.RequestOption) (res *BotTimeseriesGroupsResponse, err error) {
	var env BotTimeseriesGroupsResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/bots/timeseries_groups/%v", dimension)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BotListResponse struct {
	Bots []BotListResponseBot `json:"bots,required"`
	JSON botListResponseJSON  `json:"-"`
}

// botListResponseJSON contains the JSON metadata for the struct [BotListResponse]
type botListResponseJSON struct {
	Bots        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botListResponseJSON) RawJSON() string {
	return r.raw
}

type BotListResponseBot struct {
	// The category of the bot.
	Category string `json:"category,required"`
	// A summary for the bot (e.g., purpose).
	Description string `json:"description,required"`
	// The name of the bot.
	Name string `json:"name,required"`
	// The organization that owns and operates the bot.
	Operator string `json:"operator,required"`
	// A kebab-case identifier derived from the bot name.
	Slug              string                 `json:"slug,required"`
	UserAgentPatterns []string               `json:"userAgentPatterns,required"`
	JSON              botListResponseBotJSON `json:"-"`
}

// botListResponseBotJSON contains the JSON metadata for the struct
// [BotListResponseBot]
type botListResponseBotJSON struct {
	Category          apijson.Field
	Description       apijson.Field
	Name              apijson.Field
	Operator          apijson.Field
	Slug              apijson.Field
	UserAgentPatterns apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BotListResponseBot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botListResponseBotJSON) RawJSON() string {
	return r.raw
}

type BotGetResponse struct {
	Bot  BotGetResponseBot  `json:"bot,required"`
	JSON botGetResponseJSON `json:"-"`
}

// botGetResponseJSON contains the JSON metadata for the struct [BotGetResponse]
type botGetResponseJSON struct {
	Bot         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botGetResponseJSON) RawJSON() string {
	return r.raw
}

type BotGetResponseBot struct {
	// The category of the bot.
	Category string `json:"category,required"`
	// A summary for the bot (e.g., purpose).
	Description string `json:"description,required"`
	// The name of the bot.
	Name string `json:"name,required"`
	// The organization that owns and operates the bot.
	Operator string `json:"operator,required"`
	// The link to the bot documentation.
	OperatorURL string `json:"operatorUrl,required"`
	// A kebab-case identifier derived from the bot name.
	Slug              string                `json:"slug,required"`
	UserAgentPatterns []string              `json:"userAgentPatterns,required"`
	UserAgents        []string              `json:"userAgents,required"`
	JSON              botGetResponseBotJSON `json:"-"`
}

// botGetResponseBotJSON contains the JSON metadata for the struct
// [BotGetResponseBot]
type botGetResponseBotJSON struct {
	Category          apijson.Field
	Description       apijson.Field
	Name              apijson.Field
	Operator          apijson.Field
	OperatorURL       apijson.Field
	Slug              apijson.Field
	UserAgentPatterns apijson.Field
	UserAgents        apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *BotGetResponseBot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botGetResponseBotJSON) RawJSON() string {
	return r.raw
}

type BotSummaryResponse struct {
	// Metadata for the results.
	Meta     BotSummaryResponseMeta `json:"meta,required"`
	Summary0 map[string]string      `json:"summary_0,required"`
	JSON     botSummaryResponseJSON `json:"-"`
}

// botSummaryResponseJSON contains the JSON metadata for the struct
// [BotSummaryResponse]
type botSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botSummaryResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type BotSummaryResponseMeta struct {
	ConfidenceInfo BotSummaryResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []BotSummaryResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization BotSummaryResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []BotSummaryResponseMetaUnit `json:"units,required"`
	JSON  botSummaryResponseMetaJSON   `json:"-"`
}

// botSummaryResponseMetaJSON contains the JSON metadata for the struct
// [BotSummaryResponseMeta]
type botSummaryResponseMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BotSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botSummaryResponseMetaJSON) RawJSON() string {
	return r.raw
}

type BotSummaryResponseMetaConfidenceInfo struct {
	Annotations []BotSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                    `json:"level,required"`
	JSON  botSummaryResponseMetaConfidenceInfoJSON `json:"-"`
}

// botSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [BotSummaryResponseMetaConfidenceInfo]
type botSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botSummaryResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type BotSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                               `json:"isInstantaneous,required"`
	LinkedURL       string                                             `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                          `json:"startDate,required" format:"date-time"`
	JSON            botSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// botSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [BotSummaryResponseMetaConfidenceInfoAnnotation]
type botSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BotSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botSummaryResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BotSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                           `json:"startTime,required" format:"date-time"`
	JSON      botSummaryResponseMetaDateRangeJSON `json:"-"`
}

// botSummaryResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [BotSummaryResponseMetaDateRange]
type botSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botSummaryResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type BotSummaryResponseMetaNormalization string

const (
	BotSummaryResponseMetaNormalizationPercentage           BotSummaryResponseMetaNormalization = "PERCENTAGE"
	BotSummaryResponseMetaNormalizationMin0Max              BotSummaryResponseMetaNormalization = "MIN0_MAX"
	BotSummaryResponseMetaNormalizationMinMax               BotSummaryResponseMetaNormalization = "MIN_MAX"
	BotSummaryResponseMetaNormalizationRawValues            BotSummaryResponseMetaNormalization = "RAW_VALUES"
	BotSummaryResponseMetaNormalizationPercentageChange     BotSummaryResponseMetaNormalization = "PERCENTAGE_CHANGE"
	BotSummaryResponseMetaNormalizationRollingAverage       BotSummaryResponseMetaNormalization = "ROLLING_AVERAGE"
	BotSummaryResponseMetaNormalizationOverlappedPercentage BotSummaryResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	BotSummaryResponseMetaNormalizationRatio                BotSummaryResponseMetaNormalization = "RATIO"
)

func (r BotSummaryResponseMetaNormalization) IsKnown() bool {
	switch r {
	case BotSummaryResponseMetaNormalizationPercentage, BotSummaryResponseMetaNormalizationMin0Max, BotSummaryResponseMetaNormalizationMinMax, BotSummaryResponseMetaNormalizationRawValues, BotSummaryResponseMetaNormalizationPercentageChange, BotSummaryResponseMetaNormalizationRollingAverage, BotSummaryResponseMetaNormalizationOverlappedPercentage, BotSummaryResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type BotSummaryResponseMetaUnit struct {
	Name  string                         `json:"name,required"`
	Value string                         `json:"value,required"`
	JSON  botSummaryResponseMetaUnitJSON `json:"-"`
}

// botSummaryResponseMetaUnitJSON contains the JSON metadata for the struct
// [BotSummaryResponseMetaUnit]
type botSummaryResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotSummaryResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botSummaryResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type BotTimeseriesResponse struct {
	// Metadata for the results.
	Meta        BotTimeseriesResponseMeta        `json:"meta,required"`
	ExtraFields map[string]BotTimeseriesResponse `json:"-,extras"`
	JSON        botTimeseriesResponseJSON        `json:"-"`
}

// botTimeseriesResponseJSON contains the JSON metadata for the struct
// [BotTimeseriesResponse]
type botTimeseriesResponseJSON struct {
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type BotTimeseriesResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    BotTimeseriesResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo BotTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []BotTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization BotTimeseriesResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []BotTimeseriesResponseMetaUnit `json:"units,required"`
	JSON  botTimeseriesResponseMetaJSON   `json:"-"`
}

// botTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [BotTimeseriesResponseMeta]
type botTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BotTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BotTimeseriesResponseMetaAggInterval string

const (
	BotTimeseriesResponseMetaAggIntervalFifteenMinutes BotTimeseriesResponseMetaAggInterval = "FIFTEEN_MINUTES"
	BotTimeseriesResponseMetaAggIntervalOneHour        BotTimeseriesResponseMetaAggInterval = "ONE_HOUR"
	BotTimeseriesResponseMetaAggIntervalOneDay         BotTimeseriesResponseMetaAggInterval = "ONE_DAY"
	BotTimeseriesResponseMetaAggIntervalOneWeek        BotTimeseriesResponseMetaAggInterval = "ONE_WEEK"
	BotTimeseriesResponseMetaAggIntervalOneMonth       BotTimeseriesResponseMetaAggInterval = "ONE_MONTH"
)

func (r BotTimeseriesResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case BotTimeseriesResponseMetaAggIntervalFifteenMinutes, BotTimeseriesResponseMetaAggIntervalOneHour, BotTimeseriesResponseMetaAggIntervalOneDay, BotTimeseriesResponseMetaAggIntervalOneWeek, BotTimeseriesResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type BotTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []BotTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                       `json:"level,required"`
	JSON  botTimeseriesResponseMetaConfidenceInfoJSON `json:"-"`
}

// botTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [BotTimeseriesResponseMetaConfidenceInfo]
type botTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type BotTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                  `json:"isInstantaneous,required"`
	LinkedURL       string                                                `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                             `json:"startDate,required" format:"date-time"`
	JSON            botTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// botTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON metadata
// for the struct [BotTimeseriesResponseMetaConfidenceInfoAnnotation]
type botTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BotTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BotTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                              `json:"startTime,required" format:"date-time"`
	JSON      botTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// botTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the struct
// [BotTimeseriesResponseMetaDateRange]
type botTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type BotTimeseriesResponseMetaNormalization string

const (
	BotTimeseriesResponseMetaNormalizationPercentage           BotTimeseriesResponseMetaNormalization = "PERCENTAGE"
	BotTimeseriesResponseMetaNormalizationMin0Max              BotTimeseriesResponseMetaNormalization = "MIN0_MAX"
	BotTimeseriesResponseMetaNormalizationMinMax               BotTimeseriesResponseMetaNormalization = "MIN_MAX"
	BotTimeseriesResponseMetaNormalizationRawValues            BotTimeseriesResponseMetaNormalization = "RAW_VALUES"
	BotTimeseriesResponseMetaNormalizationPercentageChange     BotTimeseriesResponseMetaNormalization = "PERCENTAGE_CHANGE"
	BotTimeseriesResponseMetaNormalizationRollingAverage       BotTimeseriesResponseMetaNormalization = "ROLLING_AVERAGE"
	BotTimeseriesResponseMetaNormalizationOverlappedPercentage BotTimeseriesResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	BotTimeseriesResponseMetaNormalizationRatio                BotTimeseriesResponseMetaNormalization = "RATIO"
)

func (r BotTimeseriesResponseMetaNormalization) IsKnown() bool {
	switch r {
	case BotTimeseriesResponseMetaNormalizationPercentage, BotTimeseriesResponseMetaNormalizationMin0Max, BotTimeseriesResponseMetaNormalizationMinMax, BotTimeseriesResponseMetaNormalizationRawValues, BotTimeseriesResponseMetaNormalizationPercentageChange, BotTimeseriesResponseMetaNormalizationRollingAverage, BotTimeseriesResponseMetaNormalizationOverlappedPercentage, BotTimeseriesResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type BotTimeseriesResponseMetaUnit struct {
	Name  string                            `json:"name,required"`
	Value string                            `json:"value,required"`
	JSON  botTimeseriesResponseMetaUnitJSON `json:"-"`
}

// botTimeseriesResponseMetaUnitJSON contains the JSON metadata for the struct
// [BotTimeseriesResponseMetaUnit]
type botTimeseriesResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type BotTimeseriesGroupsResponse struct {
	// Metadata for the results.
	Meta   BotTimeseriesGroupsResponseMeta   `json:"meta,required"`
	Serie0 BotTimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   botTimeseriesGroupsResponseJSON   `json:"-"`
}

// botTimeseriesGroupsResponseJSON contains the JSON metadata for the struct
// [BotTimeseriesGroupsResponse]
type botTimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseJSON) RawJSON() string {
	return r.raw
}

// Metadata for the results.
type BotTimeseriesGroupsResponseMeta struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval    BotTimeseriesGroupsResponseMetaAggInterval    `json:"aggInterval,required"`
	ConfidenceInfo BotTimeseriesGroupsResponseMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      []BotTimeseriesGroupsResponseMetaDateRange    `json:"dateRange,required"`
	// Timestamp of the last dataset update.
	LastUpdated time.Time `json:"lastUpdated,required" format:"date-time"`
	// Normalization method applied to the results. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization BotTimeseriesGroupsResponseMetaNormalization `json:"normalization,required"`
	// Measurement units for the results.
	Units []BotTimeseriesGroupsResponseMetaUnit `json:"units,required"`
	JSON  botTimeseriesGroupsResponseMetaJSON   `json:"-"`
}

// botTimeseriesGroupsResponseMetaJSON contains the JSON metadata for the struct
// [BotTimeseriesGroupsResponseMeta]
type botTimeseriesGroupsResponseMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	Units          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *BotTimeseriesGroupsResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseMetaJSON) RawJSON() string {
	return r.raw
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BotTimeseriesGroupsResponseMetaAggInterval string

const (
	BotTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes BotTimeseriesGroupsResponseMetaAggInterval = "FIFTEEN_MINUTES"
	BotTimeseriesGroupsResponseMetaAggIntervalOneHour        BotTimeseriesGroupsResponseMetaAggInterval = "ONE_HOUR"
	BotTimeseriesGroupsResponseMetaAggIntervalOneDay         BotTimeseriesGroupsResponseMetaAggInterval = "ONE_DAY"
	BotTimeseriesGroupsResponseMetaAggIntervalOneWeek        BotTimeseriesGroupsResponseMetaAggInterval = "ONE_WEEK"
	BotTimeseriesGroupsResponseMetaAggIntervalOneMonth       BotTimeseriesGroupsResponseMetaAggInterval = "ONE_MONTH"
)

func (r BotTimeseriesGroupsResponseMetaAggInterval) IsKnown() bool {
	switch r {
	case BotTimeseriesGroupsResponseMetaAggIntervalFifteenMinutes, BotTimeseriesGroupsResponseMetaAggIntervalOneHour, BotTimeseriesGroupsResponseMetaAggIntervalOneDay, BotTimeseriesGroupsResponseMetaAggIntervalOneWeek, BotTimeseriesGroupsResponseMetaAggIntervalOneMonth:
		return true
	}
	return false
}

type BotTimeseriesGroupsResponseMetaConfidenceInfo struct {
	Annotations []BotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation `json:"annotations,required"`
	// Provides an indication of how much confidence Cloudflare has in the data.
	Level int64                                             `json:"level,required"`
	JSON  botTimeseriesGroupsResponseMetaConfidenceInfoJSON `json:"-"`
}

// botTimeseriesGroupsResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [BotTimeseriesGroupsResponseMetaConfidenceInfo]
type botTimeseriesGroupsResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesGroupsResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

// Annotation associated with the result (e.g. outage or other type of event).
type BotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation struct {
	DataSource  string    `json:"dataSource,required"`
	Description string    `json:"description,required"`
	EndDate     time.Time `json:"endDate,required" format:"date-time"`
	EventType   string    `json:"eventType,required"`
	// Whether event is a single point in time or a time range.
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	LinkedURL       string                                                      `json:"linkedUrl,required" format:"uri"`
	StartDate       time.Time                                                   `json:"startDate,required" format:"date-time"`
	JSON            botTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// botTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [BotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation]
type botTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *BotTimeseriesGroupsResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type BotTimeseriesGroupsResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      botTimeseriesGroupsResponseMetaDateRangeJSON `json:"-"`
}

// botTimeseriesGroupsResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [BotTimeseriesGroupsResponseMetaDateRange]
type botTimeseriesGroupsResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesGroupsResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

// Normalization method applied to the results. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type BotTimeseriesGroupsResponseMetaNormalization string

const (
	BotTimeseriesGroupsResponseMetaNormalizationPercentage           BotTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE"
	BotTimeseriesGroupsResponseMetaNormalizationMin0Max              BotTimeseriesGroupsResponseMetaNormalization = "MIN0_MAX"
	BotTimeseriesGroupsResponseMetaNormalizationMinMax               BotTimeseriesGroupsResponseMetaNormalization = "MIN_MAX"
	BotTimeseriesGroupsResponseMetaNormalizationRawValues            BotTimeseriesGroupsResponseMetaNormalization = "RAW_VALUES"
	BotTimeseriesGroupsResponseMetaNormalizationPercentageChange     BotTimeseriesGroupsResponseMetaNormalization = "PERCENTAGE_CHANGE"
	BotTimeseriesGroupsResponseMetaNormalizationRollingAverage       BotTimeseriesGroupsResponseMetaNormalization = "ROLLING_AVERAGE"
	BotTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage BotTimeseriesGroupsResponseMetaNormalization = "OVERLAPPED_PERCENTAGE"
	BotTimeseriesGroupsResponseMetaNormalizationRatio                BotTimeseriesGroupsResponseMetaNormalization = "RATIO"
)

func (r BotTimeseriesGroupsResponseMetaNormalization) IsKnown() bool {
	switch r {
	case BotTimeseriesGroupsResponseMetaNormalizationPercentage, BotTimeseriesGroupsResponseMetaNormalizationMin0Max, BotTimeseriesGroupsResponseMetaNormalizationMinMax, BotTimeseriesGroupsResponseMetaNormalizationRawValues, BotTimeseriesGroupsResponseMetaNormalizationPercentageChange, BotTimeseriesGroupsResponseMetaNormalizationRollingAverage, BotTimeseriesGroupsResponseMetaNormalizationOverlappedPercentage, BotTimeseriesGroupsResponseMetaNormalizationRatio:
		return true
	}
	return false
}

type BotTimeseriesGroupsResponseMetaUnit struct {
	Name  string                                  `json:"name,required"`
	Value string                                  `json:"value,required"`
	JSON  botTimeseriesGroupsResponseMetaUnitJSON `json:"-"`
}

// botTimeseriesGroupsResponseMetaUnitJSON contains the JSON metadata for the
// struct [BotTimeseriesGroupsResponseMetaUnit]
type botTimeseriesGroupsResponseMetaUnitJSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesGroupsResponseMetaUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseMetaUnitJSON) RawJSON() string {
	return r.raw
}

type BotTimeseriesGroupsResponseSerie0 struct {
	Timestamps  []time.Time                           `json:"timestamps,required" format:"date-time"`
	ExtraFields map[string][]string                   `json:"-,extras"`
	JSON        botTimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// botTimeseriesGroupsResponseSerie0JSON contains the JSON metadata for the struct
// [BotTimeseriesGroupsResponseSerie0]
type botTimeseriesGroupsResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type BotListParams struct {
	// Filters results by bot category.
	BotCategory param.Field[BotListParamsBotCategory] `query:"botCategory"`
	// Filters results by bot operator.
	BotOperator param.Field[string] `query:"botOperator"`
	// Filters results by bot verification status.
	BotVerificationStatus param.Field[BotListParamsBotVerificationStatus] `query:"botVerificationStatus"`
	// Format in which results will be returned.
	Format param.Field[BotListParamsFormat] `query:"format"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Skips the specified number of objects before fetching the results.
	Offset param.Field[int64] `query:"offset"`
}

// URLQuery serializes [BotListParams]'s query parameters as `url.Values`.
func (r BotListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filters results by bot category.
type BotListParamsBotCategory string

const (
	BotListParamsBotCategorySearchEngineCrawler      BotListParamsBotCategory = "SEARCH_ENGINE_CRAWLER"
	BotListParamsBotCategorySearchEngineOptimization BotListParamsBotCategory = "SEARCH_ENGINE_OPTIMIZATION"
	BotListParamsBotCategoryMonitoringAndAnalytics   BotListParamsBotCategory = "MONITORING_AND_ANALYTICS"
	BotListParamsBotCategoryAdvertisingAndMarketing  BotListParamsBotCategory = "ADVERTISING_AND_MARKETING"
	BotListParamsBotCategorySocialMediaMarketing     BotListParamsBotCategory = "SOCIAL_MEDIA_MARKETING"
	BotListParamsBotCategoryPagePreview              BotListParamsBotCategory = "PAGE_PREVIEW"
	BotListParamsBotCategoryAcademicResearch         BotListParamsBotCategory = "ACADEMIC_RESEARCH"
	BotListParamsBotCategorySecurity                 BotListParamsBotCategory = "SECURITY"
	BotListParamsBotCategoryAccessibility            BotListParamsBotCategory = "ACCESSIBILITY"
	BotListParamsBotCategoryWebhooks                 BotListParamsBotCategory = "WEBHOOKS"
	BotListParamsBotCategoryFeedFetcher              BotListParamsBotCategory = "FEED_FETCHER"
	BotListParamsBotCategoryAICrawler                BotListParamsBotCategory = "AI_CRAWLER"
	BotListParamsBotCategoryAggregator               BotListParamsBotCategory = "AGGREGATOR"
	BotListParamsBotCategoryAIAssistant              BotListParamsBotCategory = "AI_ASSISTANT"
	BotListParamsBotCategoryAISearch                 BotListParamsBotCategory = "AI_SEARCH"
	BotListParamsBotCategoryArchiver                 BotListParamsBotCategory = "ARCHIVER"
)

func (r BotListParamsBotCategory) IsKnown() bool {
	switch r {
	case BotListParamsBotCategorySearchEngineCrawler, BotListParamsBotCategorySearchEngineOptimization, BotListParamsBotCategoryMonitoringAndAnalytics, BotListParamsBotCategoryAdvertisingAndMarketing, BotListParamsBotCategorySocialMediaMarketing, BotListParamsBotCategoryPagePreview, BotListParamsBotCategoryAcademicResearch, BotListParamsBotCategorySecurity, BotListParamsBotCategoryAccessibility, BotListParamsBotCategoryWebhooks, BotListParamsBotCategoryFeedFetcher, BotListParamsBotCategoryAICrawler, BotListParamsBotCategoryAggregator, BotListParamsBotCategoryAIAssistant, BotListParamsBotCategoryAISearch, BotListParamsBotCategoryArchiver:
		return true
	}
	return false
}

// Filters results by bot verification status.
type BotListParamsBotVerificationStatus string

const (
	BotListParamsBotVerificationStatusVerified BotListParamsBotVerificationStatus = "VERIFIED"
)

func (r BotListParamsBotVerificationStatus) IsKnown() bool {
	switch r {
	case BotListParamsBotVerificationStatusVerified:
		return true
	}
	return false
}

// Format in which results will be returned.
type BotListParamsFormat string

const (
	BotListParamsFormatJson BotListParamsFormat = "JSON"
	BotListParamsFormatCsv  BotListParamsFormat = "CSV"
)

func (r BotListParamsFormat) IsKnown() bool {
	switch r {
	case BotListParamsFormatJson, BotListParamsFormatCsv:
		return true
	}
	return false
}

type BotListResponseEnvelope struct {
	Result  BotListResponse             `json:"result,required"`
	Success bool                        `json:"success,required"`
	JSON    botListResponseEnvelopeJSON `json:"-"`
}

// botListResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotListResponseEnvelope]
type botListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BotGetParams struct {
	// Format in which results will be returned.
	Format param.Field[BotGetParamsFormat] `query:"format"`
}

// URLQuery serializes [BotGetParams]'s query parameters as `url.Values`.
func (r BotGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Format in which results will be returned.
type BotGetParamsFormat string

const (
	BotGetParamsFormatJson BotGetParamsFormat = "JSON"
	BotGetParamsFormatCsv  BotGetParamsFormat = "CSV"
)

func (r BotGetParamsFormat) IsKnown() bool {
	switch r {
	case BotGetParamsFormatJson, BotGetParamsFormatCsv:
		return true
	}
	return false
}

type BotGetResponseEnvelope struct {
	Result  BotGetResponse             `json:"result,required"`
	Success bool                       `json:"success,required"`
	JSON    botGetResponseEnvelopeJSON `json:"-"`
}

// botGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotGetResponseEnvelope]
type botGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BotSummaryParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot name.
	Bot param.Field[[]string] `query:"bot"`
	// Filters results by bot category.
	BotCategory param.Field[[]BotSummaryParamsBotCategory] `query:"botCategory"`
	// Filters results by bot operator.
	BotOperator param.Field[[]string] `query:"botOperator"`
	// Filters results by bot verification status (Verified vs. Unverified).
	BotVerificationStatus param.Field[[]BotSummaryParamsBotVerificationStatus] `query:"botVerificationStatus"`
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
	Format param.Field[BotSummaryParamsFormat] `query:"format"`
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

// URLQuery serializes [BotSummaryParams]'s query parameters as `url.Values`.
func (r BotSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type BotSummaryParamsDimension string

const (
	BotSummaryParamsDimensionBot         BotSummaryParamsDimension = "BOT"
	BotSummaryParamsDimensionBotOperator BotSummaryParamsDimension = "BOT_OPERATOR"
	BotSummaryParamsDimensionBotCategory BotSummaryParamsDimension = "BOT_CATEGORY"
)

func (r BotSummaryParamsDimension) IsKnown() bool {
	switch r {
	case BotSummaryParamsDimensionBot, BotSummaryParamsDimensionBotOperator, BotSummaryParamsDimensionBotCategory:
		return true
	}
	return false
}

// The category of the bot.
type BotSummaryParamsBotCategory string

const (
	BotSummaryParamsBotCategorySearchEngineCrawler      BotSummaryParamsBotCategory = "SEARCH_ENGINE_CRAWLER"
	BotSummaryParamsBotCategorySearchEngineOptimization BotSummaryParamsBotCategory = "SEARCH_ENGINE_OPTIMIZATION"
	BotSummaryParamsBotCategoryMonitoringAndAnalytics   BotSummaryParamsBotCategory = "MONITORING_AND_ANALYTICS"
	BotSummaryParamsBotCategoryAdvertisingAndMarketing  BotSummaryParamsBotCategory = "ADVERTISING_AND_MARKETING"
	BotSummaryParamsBotCategorySocialMediaMarketing     BotSummaryParamsBotCategory = "SOCIAL_MEDIA_MARKETING"
	BotSummaryParamsBotCategoryPagePreview              BotSummaryParamsBotCategory = "PAGE_PREVIEW"
	BotSummaryParamsBotCategoryAcademicResearch         BotSummaryParamsBotCategory = "ACADEMIC_RESEARCH"
	BotSummaryParamsBotCategorySecurity                 BotSummaryParamsBotCategory = "SECURITY"
	BotSummaryParamsBotCategoryAccessibility            BotSummaryParamsBotCategory = "ACCESSIBILITY"
	BotSummaryParamsBotCategoryWebhooks                 BotSummaryParamsBotCategory = "WEBHOOKS"
	BotSummaryParamsBotCategoryFeedFetcher              BotSummaryParamsBotCategory = "FEED_FETCHER"
	BotSummaryParamsBotCategoryAICrawler                BotSummaryParamsBotCategory = "AI_CRAWLER"
	BotSummaryParamsBotCategoryAggregator               BotSummaryParamsBotCategory = "AGGREGATOR"
	BotSummaryParamsBotCategoryAIAssistant              BotSummaryParamsBotCategory = "AI_ASSISTANT"
	BotSummaryParamsBotCategoryAISearch                 BotSummaryParamsBotCategory = "AI_SEARCH"
	BotSummaryParamsBotCategoryArchiver                 BotSummaryParamsBotCategory = "ARCHIVER"
)

func (r BotSummaryParamsBotCategory) IsKnown() bool {
	switch r {
	case BotSummaryParamsBotCategorySearchEngineCrawler, BotSummaryParamsBotCategorySearchEngineOptimization, BotSummaryParamsBotCategoryMonitoringAndAnalytics, BotSummaryParamsBotCategoryAdvertisingAndMarketing, BotSummaryParamsBotCategorySocialMediaMarketing, BotSummaryParamsBotCategoryPagePreview, BotSummaryParamsBotCategoryAcademicResearch, BotSummaryParamsBotCategorySecurity, BotSummaryParamsBotCategoryAccessibility, BotSummaryParamsBotCategoryWebhooks, BotSummaryParamsBotCategoryFeedFetcher, BotSummaryParamsBotCategoryAICrawler, BotSummaryParamsBotCategoryAggregator, BotSummaryParamsBotCategoryAIAssistant, BotSummaryParamsBotCategoryAISearch, BotSummaryParamsBotCategoryArchiver:
		return true
	}
	return false
}

// The category of the bot.
type BotSummaryParamsBotVerificationStatus string

const (
	BotSummaryParamsBotVerificationStatusVerified BotSummaryParamsBotVerificationStatus = "VERIFIED"
)

func (r BotSummaryParamsBotVerificationStatus) IsKnown() bool {
	switch r {
	case BotSummaryParamsBotVerificationStatusVerified:
		return true
	}
	return false
}

// Format in which results will be returned.
type BotSummaryParamsFormat string

const (
	BotSummaryParamsFormatJson BotSummaryParamsFormat = "JSON"
	BotSummaryParamsFormatCsv  BotSummaryParamsFormat = "CSV"
)

func (r BotSummaryParamsFormat) IsKnown() bool {
	switch r {
	case BotSummaryParamsFormatJson, BotSummaryParamsFormatCsv:
		return true
	}
	return false
}

type BotSummaryResponseEnvelope struct {
	Result  BotSummaryResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    botSummaryResponseEnvelopeJSON `json:"-"`
}

// botSummaryResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotSummaryResponseEnvelope]
type botSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botSummaryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BotTimeseriesParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[BotTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot name.
	Bot param.Field[[]string] `query:"bot"`
	// Filters results by bot category.
	BotCategory param.Field[[]BotTimeseriesParamsBotCategory] `query:"botCategory"`
	// Filters results by bot operator.
	BotOperator param.Field[[]string] `query:"botOperator"`
	// Filters results by bot verification status (Verified vs. Unverified).
	BotVerificationStatus param.Field[[]BotTimeseriesParamsBotVerificationStatus] `query:"botVerificationStatus"`
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
	Format param.Field[BotTimeseriesParamsFormat] `query:"format"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [BotTimeseriesParams]'s query parameters as `url.Values`.
func (r BotTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BotTimeseriesParamsAggInterval string

const (
	BotTimeseriesParamsAggInterval15m BotTimeseriesParamsAggInterval = "15m"
	BotTimeseriesParamsAggInterval1h  BotTimeseriesParamsAggInterval = "1h"
	BotTimeseriesParamsAggInterval1d  BotTimeseriesParamsAggInterval = "1d"
	BotTimeseriesParamsAggInterval1w  BotTimeseriesParamsAggInterval = "1w"
)

func (r BotTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case BotTimeseriesParamsAggInterval15m, BotTimeseriesParamsAggInterval1h, BotTimeseriesParamsAggInterval1d, BotTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

// The category of the bot.
type BotTimeseriesParamsBotCategory string

const (
	BotTimeseriesParamsBotCategorySearchEngineCrawler      BotTimeseriesParamsBotCategory = "SEARCH_ENGINE_CRAWLER"
	BotTimeseriesParamsBotCategorySearchEngineOptimization BotTimeseriesParamsBotCategory = "SEARCH_ENGINE_OPTIMIZATION"
	BotTimeseriesParamsBotCategoryMonitoringAndAnalytics   BotTimeseriesParamsBotCategory = "MONITORING_AND_ANALYTICS"
	BotTimeseriesParamsBotCategoryAdvertisingAndMarketing  BotTimeseriesParamsBotCategory = "ADVERTISING_AND_MARKETING"
	BotTimeseriesParamsBotCategorySocialMediaMarketing     BotTimeseriesParamsBotCategory = "SOCIAL_MEDIA_MARKETING"
	BotTimeseriesParamsBotCategoryPagePreview              BotTimeseriesParamsBotCategory = "PAGE_PREVIEW"
	BotTimeseriesParamsBotCategoryAcademicResearch         BotTimeseriesParamsBotCategory = "ACADEMIC_RESEARCH"
	BotTimeseriesParamsBotCategorySecurity                 BotTimeseriesParamsBotCategory = "SECURITY"
	BotTimeseriesParamsBotCategoryAccessibility            BotTimeseriesParamsBotCategory = "ACCESSIBILITY"
	BotTimeseriesParamsBotCategoryWebhooks                 BotTimeseriesParamsBotCategory = "WEBHOOKS"
	BotTimeseriesParamsBotCategoryFeedFetcher              BotTimeseriesParamsBotCategory = "FEED_FETCHER"
	BotTimeseriesParamsBotCategoryAICrawler                BotTimeseriesParamsBotCategory = "AI_CRAWLER"
	BotTimeseriesParamsBotCategoryAggregator               BotTimeseriesParamsBotCategory = "AGGREGATOR"
	BotTimeseriesParamsBotCategoryAIAssistant              BotTimeseriesParamsBotCategory = "AI_ASSISTANT"
	BotTimeseriesParamsBotCategoryAISearch                 BotTimeseriesParamsBotCategory = "AI_SEARCH"
	BotTimeseriesParamsBotCategoryArchiver                 BotTimeseriesParamsBotCategory = "ARCHIVER"
)

func (r BotTimeseriesParamsBotCategory) IsKnown() bool {
	switch r {
	case BotTimeseriesParamsBotCategorySearchEngineCrawler, BotTimeseriesParamsBotCategorySearchEngineOptimization, BotTimeseriesParamsBotCategoryMonitoringAndAnalytics, BotTimeseriesParamsBotCategoryAdvertisingAndMarketing, BotTimeseriesParamsBotCategorySocialMediaMarketing, BotTimeseriesParamsBotCategoryPagePreview, BotTimeseriesParamsBotCategoryAcademicResearch, BotTimeseriesParamsBotCategorySecurity, BotTimeseriesParamsBotCategoryAccessibility, BotTimeseriesParamsBotCategoryWebhooks, BotTimeseriesParamsBotCategoryFeedFetcher, BotTimeseriesParamsBotCategoryAICrawler, BotTimeseriesParamsBotCategoryAggregator, BotTimeseriesParamsBotCategoryAIAssistant, BotTimeseriesParamsBotCategoryAISearch, BotTimeseriesParamsBotCategoryArchiver:
		return true
	}
	return false
}

// The category of the bot.
type BotTimeseriesParamsBotVerificationStatus string

const (
	BotTimeseriesParamsBotVerificationStatusVerified BotTimeseriesParamsBotVerificationStatus = "VERIFIED"
)

func (r BotTimeseriesParamsBotVerificationStatus) IsKnown() bool {
	switch r {
	case BotTimeseriesParamsBotVerificationStatusVerified:
		return true
	}
	return false
}

// Format in which results will be returned.
type BotTimeseriesParamsFormat string

const (
	BotTimeseriesParamsFormatJson BotTimeseriesParamsFormat = "JSON"
	BotTimeseriesParamsFormatCsv  BotTimeseriesParamsFormat = "CSV"
)

func (r BotTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case BotTimeseriesParamsFormatJson, BotTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type BotTimeseriesResponseEnvelope struct {
	Result  BotTimeseriesResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    botTimeseriesResponseEnvelopeJSON `json:"-"`
}

// botTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [BotTimeseriesResponseEnvelope]
type botTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type BotTimeseriesGroupsParams struct {
	// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
	// Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[BotTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot name.
	Bot param.Field[[]string] `query:"bot"`
	// Filters results by bot category.
	BotCategory param.Field[[]BotTimeseriesGroupsParamsBotCategory] `query:"botCategory"`
	// Filters results by bot operator.
	BotOperator param.Field[[]string] `query:"botOperator"`
	// Filters results by bot verification status (Verified vs. Unverified).
	BotVerificationStatus param.Field[[]BotTimeseriesGroupsParamsBotVerificationStatus] `query:"botVerificationStatus"`
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
	Format param.Field[BotTimeseriesGroupsParamsFormat] `query:"format"`
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

// URLQuery serializes [BotTimeseriesGroupsParams]'s query parameters as
// `url.Values`.
func (r BotTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Specifies the attribute by which to group the results.
type BotTimeseriesGroupsParamsDimension string

const (
	BotTimeseriesGroupsParamsDimensionBot         BotTimeseriesGroupsParamsDimension = "BOT"
	BotTimeseriesGroupsParamsDimensionBotOperator BotTimeseriesGroupsParamsDimension = "BOT_OPERATOR"
	BotTimeseriesGroupsParamsDimensionBotCategory BotTimeseriesGroupsParamsDimension = "BOT_CATEGORY"
)

func (r BotTimeseriesGroupsParamsDimension) IsKnown() bool {
	switch r {
	case BotTimeseriesGroupsParamsDimensionBot, BotTimeseriesGroupsParamsDimensionBotOperator, BotTimeseriesGroupsParamsDimensionBotCategory:
		return true
	}
	return false
}

// Aggregation interval of the results (e.g., in 15 minutes or 1 hour intervals).
// Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type BotTimeseriesGroupsParamsAggInterval string

const (
	BotTimeseriesGroupsParamsAggInterval15m BotTimeseriesGroupsParamsAggInterval = "15m"
	BotTimeseriesGroupsParamsAggInterval1h  BotTimeseriesGroupsParamsAggInterval = "1h"
	BotTimeseriesGroupsParamsAggInterval1d  BotTimeseriesGroupsParamsAggInterval = "1d"
	BotTimeseriesGroupsParamsAggInterval1w  BotTimeseriesGroupsParamsAggInterval = "1w"
)

func (r BotTimeseriesGroupsParamsAggInterval) IsKnown() bool {
	switch r {
	case BotTimeseriesGroupsParamsAggInterval15m, BotTimeseriesGroupsParamsAggInterval1h, BotTimeseriesGroupsParamsAggInterval1d, BotTimeseriesGroupsParamsAggInterval1w:
		return true
	}
	return false
}

// The category of the bot.
type BotTimeseriesGroupsParamsBotCategory string

const (
	BotTimeseriesGroupsParamsBotCategorySearchEngineCrawler      BotTimeseriesGroupsParamsBotCategory = "SEARCH_ENGINE_CRAWLER"
	BotTimeseriesGroupsParamsBotCategorySearchEngineOptimization BotTimeseriesGroupsParamsBotCategory = "SEARCH_ENGINE_OPTIMIZATION"
	BotTimeseriesGroupsParamsBotCategoryMonitoringAndAnalytics   BotTimeseriesGroupsParamsBotCategory = "MONITORING_AND_ANALYTICS"
	BotTimeseriesGroupsParamsBotCategoryAdvertisingAndMarketing  BotTimeseriesGroupsParamsBotCategory = "ADVERTISING_AND_MARKETING"
	BotTimeseriesGroupsParamsBotCategorySocialMediaMarketing     BotTimeseriesGroupsParamsBotCategory = "SOCIAL_MEDIA_MARKETING"
	BotTimeseriesGroupsParamsBotCategoryPagePreview              BotTimeseriesGroupsParamsBotCategory = "PAGE_PREVIEW"
	BotTimeseriesGroupsParamsBotCategoryAcademicResearch         BotTimeseriesGroupsParamsBotCategory = "ACADEMIC_RESEARCH"
	BotTimeseriesGroupsParamsBotCategorySecurity                 BotTimeseriesGroupsParamsBotCategory = "SECURITY"
	BotTimeseriesGroupsParamsBotCategoryAccessibility            BotTimeseriesGroupsParamsBotCategory = "ACCESSIBILITY"
	BotTimeseriesGroupsParamsBotCategoryWebhooks                 BotTimeseriesGroupsParamsBotCategory = "WEBHOOKS"
	BotTimeseriesGroupsParamsBotCategoryFeedFetcher              BotTimeseriesGroupsParamsBotCategory = "FEED_FETCHER"
	BotTimeseriesGroupsParamsBotCategoryAICrawler                BotTimeseriesGroupsParamsBotCategory = "AI_CRAWLER"
	BotTimeseriesGroupsParamsBotCategoryAggregator               BotTimeseriesGroupsParamsBotCategory = "AGGREGATOR"
	BotTimeseriesGroupsParamsBotCategoryAIAssistant              BotTimeseriesGroupsParamsBotCategory = "AI_ASSISTANT"
	BotTimeseriesGroupsParamsBotCategoryAISearch                 BotTimeseriesGroupsParamsBotCategory = "AI_SEARCH"
	BotTimeseriesGroupsParamsBotCategoryArchiver                 BotTimeseriesGroupsParamsBotCategory = "ARCHIVER"
)

func (r BotTimeseriesGroupsParamsBotCategory) IsKnown() bool {
	switch r {
	case BotTimeseriesGroupsParamsBotCategorySearchEngineCrawler, BotTimeseriesGroupsParamsBotCategorySearchEngineOptimization, BotTimeseriesGroupsParamsBotCategoryMonitoringAndAnalytics, BotTimeseriesGroupsParamsBotCategoryAdvertisingAndMarketing, BotTimeseriesGroupsParamsBotCategorySocialMediaMarketing, BotTimeseriesGroupsParamsBotCategoryPagePreview, BotTimeseriesGroupsParamsBotCategoryAcademicResearch, BotTimeseriesGroupsParamsBotCategorySecurity, BotTimeseriesGroupsParamsBotCategoryAccessibility, BotTimeseriesGroupsParamsBotCategoryWebhooks, BotTimeseriesGroupsParamsBotCategoryFeedFetcher, BotTimeseriesGroupsParamsBotCategoryAICrawler, BotTimeseriesGroupsParamsBotCategoryAggregator, BotTimeseriesGroupsParamsBotCategoryAIAssistant, BotTimeseriesGroupsParamsBotCategoryAISearch, BotTimeseriesGroupsParamsBotCategoryArchiver:
		return true
	}
	return false
}

// The category of the bot.
type BotTimeseriesGroupsParamsBotVerificationStatus string

const (
	BotTimeseriesGroupsParamsBotVerificationStatusVerified BotTimeseriesGroupsParamsBotVerificationStatus = "VERIFIED"
)

func (r BotTimeseriesGroupsParamsBotVerificationStatus) IsKnown() bool {
	switch r {
	case BotTimeseriesGroupsParamsBotVerificationStatusVerified:
		return true
	}
	return false
}

// Format in which results will be returned.
type BotTimeseriesGroupsParamsFormat string

const (
	BotTimeseriesGroupsParamsFormatJson BotTimeseriesGroupsParamsFormat = "JSON"
	BotTimeseriesGroupsParamsFormatCsv  BotTimeseriesGroupsParamsFormat = "CSV"
)

func (r BotTimeseriesGroupsParamsFormat) IsKnown() bool {
	switch r {
	case BotTimeseriesGroupsParamsFormatJson, BotTimeseriesGroupsParamsFormatCsv:
		return true
	}
	return false
}

type BotTimeseriesGroupsResponseEnvelope struct {
	Result  BotTimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    botTimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// botTimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata for the
// struct [BotTimeseriesGroupsResponseEnvelope]
type botTimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BotTimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r botTimeseriesGroupsResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
