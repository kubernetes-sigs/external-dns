// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"context"
	"errors"
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

// AnalyticsEventBytimeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsEventBytimeService] method instead.
type AnalyticsEventBytimeService struct {
	Options []option.RequestOption
}

// NewAnalyticsEventBytimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsEventBytimeService(opts ...option.RequestOption) (r *AnalyticsEventBytimeService) {
	r = &AnalyticsEventBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
func (r *AnalyticsEventBytimeService) Get(ctx context.Context, params AnalyticsEventBytimeGetParams, opts ...option.RequestOption) (res *AnalyticsEventBytimeGetResponse, err error) {
	var env AnalyticsEventBytimeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/bytime", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnalyticsEventBytimeGetResponse struct {
	// List of columns returned by the analytics query.
	Data []AnalyticsEventBytimeGetResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum result for each selected metrics across all data.
	Max map[string]float64 `json:"max,required"`
	// Minimum result for each selected metrics across all data.
	Min   map[string]float64                   `json:"min,required"`
	Query AnalyticsEventBytimeGetResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total result for each selected metrics across all data.
	Totals map[string]float64 `json:"totals,required"`
	// List of time interval buckets: [start, end]
	TimeIntervals [][]time.Time                       `json:"time_intervals" format:"date-time"`
	JSON          analyticsEventBytimeGetResponseJSON `json:"-"`
}

// analyticsEventBytimeGetResponseJSON contains the JSON metadata for the struct
// [AnalyticsEventBytimeGetResponse]
type analyticsEventBytimeGetResponseJSON struct {
	Data          apijson.Field
	DataLag       apijson.Field
	Max           apijson.Field
	Min           apijson.Field
	Query         apijson.Field
	Rows          apijson.Field
	Totals        apijson.Field
	TimeIntervals apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventBytimeGetResponseData struct {
	Dimensions []string                                        `json:"dimensions"`
	Metrics    AnalyticsEventBytimeGetResponseDataMetricsUnion `json:"metrics"`
	JSON       analyticsEventBytimeGetResponseDataJSON         `json:"-"`
}

// analyticsEventBytimeGetResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsEventBytimeGetResponseData]
type analyticsEventBytimeGetResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [AnalyticsEventBytimeGetResponseDataMetricsArray] or
// [AnalyticsEventBytimeGetResponseDataMetricsArray].
type AnalyticsEventBytimeGetResponseDataMetricsUnion interface {
	implementsAnalyticsEventBytimeGetResponseDataMetricsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsEventBytimeGetResponseDataMetricsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AnalyticsEventBytimeGetResponseDataMetricsArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AnalyticsEventBytimeGetResponseDataMetricsArray{}),
		},
	)
}

type AnalyticsEventBytimeGetResponseDataMetricsArray []float64

func (r AnalyticsEventBytimeGetResponseDataMetricsArray) implementsAnalyticsEventBytimeGetResponseDataMetricsUnion() {
}

type AnalyticsEventBytimeGetResponseQuery struct {
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions []Dimension `json:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | \>       | Greater Than             | %3E         |
	// | \<       | Less Than                | %3C         |
	// | \>=      | Greater than or equal to | %3E%3D      |
	// | \<=      | Less than or equal to    | %3C%3D      |
	Filters string `json:"filters"`
	// Limit number of returned metrics.
	Limit float64 `json:"limit"`
	// One or more metrics to compute. Options are:
	//
	// | Metric         | Name                                | Example | Unit                  |
	// | -------------- | ----------------------------------- | ------- | --------------------- |
	// | count          | Count of total events               | 1000    | Count                 |
	// | bytesIngress   | Sum of ingress bytes                | 1000    | Sum                   |
	// | bytesEgress    | Sum of egress bytes                 | 1000    | Sum                   |
	// | durationAvg    | Average connection duration         | 1.0     | Time in milliseconds  |
	// | durationMedian | Median connection duration          | 1.0     | Time in milliseconds  |
	// | duration90th   | 90th percentile connection duration | 1.0     | Time in milliseconds  |
	// | duration99th   | 99th percentile connection duration | 1.0     | Time in milliseconds. |
	Metrics []AnalyticsEventBytimeGetResponseQueryMetric `json:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since time.Time `json:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort []string `json:"sort"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until time.Time                                `json:"until" format:"date-time"`
	JSON  analyticsEventBytimeGetResponseQueryJSON `json:"-"`
}

// analyticsEventBytimeGetResponseQueryJSON contains the JSON metadata for the
// struct [AnalyticsEventBytimeGetResponseQuery]
type analyticsEventBytimeGetResponseQueryJSON struct {
	Dimensions  apijson.Field
	Filters     apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Sort        apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseQueryJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventBytimeGetResponseQueryMetric string

const (
	AnalyticsEventBytimeGetResponseQueryMetricCount          AnalyticsEventBytimeGetResponseQueryMetric = "count"
	AnalyticsEventBytimeGetResponseQueryMetricBytesIngress   AnalyticsEventBytimeGetResponseQueryMetric = "bytesIngress"
	AnalyticsEventBytimeGetResponseQueryMetricBytesEgress    AnalyticsEventBytimeGetResponseQueryMetric = "bytesEgress"
	AnalyticsEventBytimeGetResponseQueryMetricDurationAvg    AnalyticsEventBytimeGetResponseQueryMetric = "durationAvg"
	AnalyticsEventBytimeGetResponseQueryMetricDurationMedian AnalyticsEventBytimeGetResponseQueryMetric = "durationMedian"
	AnalyticsEventBytimeGetResponseQueryMetricDuration90th   AnalyticsEventBytimeGetResponseQueryMetric = "duration90th"
	AnalyticsEventBytimeGetResponseQueryMetricDuration99th   AnalyticsEventBytimeGetResponseQueryMetric = "duration99th"
)

func (r AnalyticsEventBytimeGetResponseQueryMetric) IsKnown() bool {
	switch r {
	case AnalyticsEventBytimeGetResponseQueryMetricCount, AnalyticsEventBytimeGetResponseQueryMetricBytesIngress, AnalyticsEventBytimeGetResponseQueryMetricBytesEgress, AnalyticsEventBytimeGetResponseQueryMetricDurationAvg, AnalyticsEventBytimeGetResponseQueryMetricDurationMedian, AnalyticsEventBytimeGetResponseQueryMetricDuration90th, AnalyticsEventBytimeGetResponseQueryMetricDuration99th:
		return true
	}
	return false
}

type AnalyticsEventBytimeGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Used to select time series resolution.
	TimeDelta param.Field[AnalyticsEventBytimeGetParamsTimeDelta] `query:"time_delta,required"`
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions param.Field[[]Dimension] `query:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | \>       | Greater Than             | %3E         |
	// | \<       | Less Than                | %3C         |
	// | \>=      | Greater than or equal to | %3E%3D      |
	// | \<=      | Less than or equal to    | %3C%3D      |
	Filters param.Field[string] `query:"filters"`
	// One or more metrics to compute. Options are:
	//
	// | Metric         | Name                                | Example | Unit                  |
	// | -------------- | ----------------------------------- | ------- | --------------------- |
	// | count          | Count of total events               | 1000    | Count                 |
	// | bytesIngress   | Sum of ingress bytes                | 1000    | Sum                   |
	// | bytesEgress    | Sum of egress bytes                 | 1000    | Sum                   |
	// | durationAvg    | Average connection duration         | 1.0     | Time in milliseconds  |
	// | durationMedian | Median connection duration          | 1.0     | Time in milliseconds  |
	// | duration90th   | 90th percentile connection duration | 1.0     | Time in milliseconds  |
	// | duration99th   | 99th percentile connection duration | 1.0     | Time in milliseconds. |
	Metrics param.Field[[]AnalyticsEventBytimeGetParamsMetric] `query:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort param.Field[[]string] `query:"sort"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsEventBytimeGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsEventBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Used to select time series resolution.
type AnalyticsEventBytimeGetParamsTimeDelta string

const (
	AnalyticsEventBytimeGetParamsTimeDeltaYear       AnalyticsEventBytimeGetParamsTimeDelta = "year"
	AnalyticsEventBytimeGetParamsTimeDeltaQuarter    AnalyticsEventBytimeGetParamsTimeDelta = "quarter"
	AnalyticsEventBytimeGetParamsTimeDeltaMonth      AnalyticsEventBytimeGetParamsTimeDelta = "month"
	AnalyticsEventBytimeGetParamsTimeDeltaWeek       AnalyticsEventBytimeGetParamsTimeDelta = "week"
	AnalyticsEventBytimeGetParamsTimeDeltaDay        AnalyticsEventBytimeGetParamsTimeDelta = "day"
	AnalyticsEventBytimeGetParamsTimeDeltaHour       AnalyticsEventBytimeGetParamsTimeDelta = "hour"
	AnalyticsEventBytimeGetParamsTimeDeltaDekaminute AnalyticsEventBytimeGetParamsTimeDelta = "dekaminute"
	AnalyticsEventBytimeGetParamsTimeDeltaMinute     AnalyticsEventBytimeGetParamsTimeDelta = "minute"
)

func (r AnalyticsEventBytimeGetParamsTimeDelta) IsKnown() bool {
	switch r {
	case AnalyticsEventBytimeGetParamsTimeDeltaYear, AnalyticsEventBytimeGetParamsTimeDeltaQuarter, AnalyticsEventBytimeGetParamsTimeDeltaMonth, AnalyticsEventBytimeGetParamsTimeDeltaWeek, AnalyticsEventBytimeGetParamsTimeDeltaDay, AnalyticsEventBytimeGetParamsTimeDeltaHour, AnalyticsEventBytimeGetParamsTimeDeltaDekaminute, AnalyticsEventBytimeGetParamsTimeDeltaMinute:
		return true
	}
	return false
}

type AnalyticsEventBytimeGetParamsMetric string

const (
	AnalyticsEventBytimeGetParamsMetricCount          AnalyticsEventBytimeGetParamsMetric = "count"
	AnalyticsEventBytimeGetParamsMetricBytesIngress   AnalyticsEventBytimeGetParamsMetric = "bytesIngress"
	AnalyticsEventBytimeGetParamsMetricBytesEgress    AnalyticsEventBytimeGetParamsMetric = "bytesEgress"
	AnalyticsEventBytimeGetParamsMetricDurationAvg    AnalyticsEventBytimeGetParamsMetric = "durationAvg"
	AnalyticsEventBytimeGetParamsMetricDurationMedian AnalyticsEventBytimeGetParamsMetric = "durationMedian"
	AnalyticsEventBytimeGetParamsMetricDuration90th   AnalyticsEventBytimeGetParamsMetric = "duration90th"
	AnalyticsEventBytimeGetParamsMetricDuration99th   AnalyticsEventBytimeGetParamsMetric = "duration99th"
)

func (r AnalyticsEventBytimeGetParamsMetric) IsKnown() bool {
	switch r {
	case AnalyticsEventBytimeGetParamsMetricCount, AnalyticsEventBytimeGetParamsMetricBytesIngress, AnalyticsEventBytimeGetParamsMetricBytesEgress, AnalyticsEventBytimeGetParamsMetricDurationAvg, AnalyticsEventBytimeGetParamsMetricDurationMedian, AnalyticsEventBytimeGetParamsMetricDuration90th, AnalyticsEventBytimeGetParamsMetricDuration99th:
		return true
	}
	return false
}

type AnalyticsEventBytimeGetResponseEnvelope struct {
	Errors   []AnalyticsEventBytimeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsEventBytimeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AnalyticsEventBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AnalyticsEventBytimeGetResponse                `json:"result"`
	JSON    analyticsEventBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsEventBytimeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsEventBytimeGetResponseEnvelope]
type analyticsEventBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventBytimeGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           AnalyticsEventBytimeGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             analyticsEventBytimeGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// analyticsEventBytimeGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AnalyticsEventBytimeGetResponseEnvelopeErrors]
type analyticsEventBytimeGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventBytimeGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    analyticsEventBytimeGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// analyticsEventBytimeGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AnalyticsEventBytimeGetResponseEnvelopeErrorsSource]
type analyticsEventBytimeGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventBytimeGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           AnalyticsEventBytimeGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             analyticsEventBytimeGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// analyticsEventBytimeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AnalyticsEventBytimeGetResponseEnvelopeMessages]
type analyticsEventBytimeGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventBytimeGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    analyticsEventBytimeGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// analyticsEventBytimeGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AnalyticsEventBytimeGetResponseEnvelopeMessagesSource]
type analyticsEventBytimeGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventBytimeGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventBytimeGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AnalyticsEventBytimeGetResponseEnvelopeSuccess bool

const (
	AnalyticsEventBytimeGetResponseEnvelopeSuccessTrue AnalyticsEventBytimeGetResponseEnvelopeSuccess = true
)

func (r AnalyticsEventBytimeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsEventBytimeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
