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

// AnalyticsEventSummaryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsEventSummaryService] method instead.
type AnalyticsEventSummaryService struct {
	Options []option.RequestOption
}

// NewAnalyticsEventSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsEventSummaryService(opts ...option.RequestOption) (r *AnalyticsEventSummaryService) {
	r = &AnalyticsEventSummaryService{}
	r.Options = opts
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
func (r *AnalyticsEventSummaryService) Get(ctx context.Context, params AnalyticsEventSummaryGetParams, opts ...option.RequestOption) (res *AnalyticsEventSummaryGetResponse, err error) {
	var env AnalyticsEventSummaryGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/summary", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnalyticsEventSummaryGetResponse struct {
	// List of columns returned by the analytics query.
	Data []AnalyticsEventSummaryGetResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum result for each selected metrics across all data.
	Max map[string]float64 `json:"max,required"`
	// Minimum result for each selected metrics across all data.
	Min   map[string]float64                    `json:"min,required"`
	Query AnalyticsEventSummaryGetResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total result for each selected metrics across all data.
	Totals map[string]float64 `json:"totals,required"`
	// List of time interval buckets: [start, end]
	TimeIntervals [][]time.Time                        `json:"time_intervals" format:"date-time"`
	JSON          analyticsEventSummaryGetResponseJSON `json:"-"`
}

// analyticsEventSummaryGetResponseJSON contains the JSON metadata for the struct
// [AnalyticsEventSummaryGetResponse]
type analyticsEventSummaryGetResponseJSON struct {
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

func (r *AnalyticsEventSummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventSummaryGetResponseData struct {
	Dimensions []string                                         `json:"dimensions"`
	Metrics    AnalyticsEventSummaryGetResponseDataMetricsUnion `json:"metrics"`
	JSON       analyticsEventSummaryGetResponseDataJSON         `json:"-"`
}

// analyticsEventSummaryGetResponseDataJSON contains the JSON metadata for the
// struct [AnalyticsEventSummaryGetResponseData]
type analyticsEventSummaryGetResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventSummaryGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseDataJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [AnalyticsEventSummaryGetResponseDataMetricsArray] or
// [AnalyticsEventSummaryGetResponseDataMetricsArray].
type AnalyticsEventSummaryGetResponseDataMetricsUnion interface {
	implementsAnalyticsEventSummaryGetResponseDataMetricsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsEventSummaryGetResponseDataMetricsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AnalyticsEventSummaryGetResponseDataMetricsArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AnalyticsEventSummaryGetResponseDataMetricsArray{}),
		},
	)
}

type AnalyticsEventSummaryGetResponseDataMetricsArray []float64

func (r AnalyticsEventSummaryGetResponseDataMetricsArray) implementsAnalyticsEventSummaryGetResponseDataMetricsUnion() {
}

type AnalyticsEventSummaryGetResponseQuery struct {
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
	Metrics []AnalyticsEventSummaryGetResponseQueryMetric `json:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since time.Time `json:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort []string `json:"sort"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until time.Time                                 `json:"until" format:"date-time"`
	JSON  analyticsEventSummaryGetResponseQueryJSON `json:"-"`
}

// analyticsEventSummaryGetResponseQueryJSON contains the JSON metadata for the
// struct [AnalyticsEventSummaryGetResponseQuery]
type analyticsEventSummaryGetResponseQueryJSON struct {
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

func (r *AnalyticsEventSummaryGetResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseQueryJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventSummaryGetResponseQueryMetric string

const (
	AnalyticsEventSummaryGetResponseQueryMetricCount          AnalyticsEventSummaryGetResponseQueryMetric = "count"
	AnalyticsEventSummaryGetResponseQueryMetricBytesIngress   AnalyticsEventSummaryGetResponseQueryMetric = "bytesIngress"
	AnalyticsEventSummaryGetResponseQueryMetricBytesEgress    AnalyticsEventSummaryGetResponseQueryMetric = "bytesEgress"
	AnalyticsEventSummaryGetResponseQueryMetricDurationAvg    AnalyticsEventSummaryGetResponseQueryMetric = "durationAvg"
	AnalyticsEventSummaryGetResponseQueryMetricDurationMedian AnalyticsEventSummaryGetResponseQueryMetric = "durationMedian"
	AnalyticsEventSummaryGetResponseQueryMetricDuration90th   AnalyticsEventSummaryGetResponseQueryMetric = "duration90th"
	AnalyticsEventSummaryGetResponseQueryMetricDuration99th   AnalyticsEventSummaryGetResponseQueryMetric = "duration99th"
)

func (r AnalyticsEventSummaryGetResponseQueryMetric) IsKnown() bool {
	switch r {
	case AnalyticsEventSummaryGetResponseQueryMetricCount, AnalyticsEventSummaryGetResponseQueryMetricBytesIngress, AnalyticsEventSummaryGetResponseQueryMetricBytesEgress, AnalyticsEventSummaryGetResponseQueryMetricDurationAvg, AnalyticsEventSummaryGetResponseQueryMetricDurationMedian, AnalyticsEventSummaryGetResponseQueryMetricDuration90th, AnalyticsEventSummaryGetResponseQueryMetricDuration99th:
		return true
	}
	return false
}

type AnalyticsEventSummaryGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	Metrics param.Field[[]AnalyticsEventSummaryGetParamsMetric] `query:"metrics"`
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

// URLQuery serializes [AnalyticsEventSummaryGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsEventSummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AnalyticsEventSummaryGetParamsMetric string

const (
	AnalyticsEventSummaryGetParamsMetricCount          AnalyticsEventSummaryGetParamsMetric = "count"
	AnalyticsEventSummaryGetParamsMetricBytesIngress   AnalyticsEventSummaryGetParamsMetric = "bytesIngress"
	AnalyticsEventSummaryGetParamsMetricBytesEgress    AnalyticsEventSummaryGetParamsMetric = "bytesEgress"
	AnalyticsEventSummaryGetParamsMetricDurationAvg    AnalyticsEventSummaryGetParamsMetric = "durationAvg"
	AnalyticsEventSummaryGetParamsMetricDurationMedian AnalyticsEventSummaryGetParamsMetric = "durationMedian"
	AnalyticsEventSummaryGetParamsMetricDuration90th   AnalyticsEventSummaryGetParamsMetric = "duration90th"
	AnalyticsEventSummaryGetParamsMetricDuration99th   AnalyticsEventSummaryGetParamsMetric = "duration99th"
)

func (r AnalyticsEventSummaryGetParamsMetric) IsKnown() bool {
	switch r {
	case AnalyticsEventSummaryGetParamsMetricCount, AnalyticsEventSummaryGetParamsMetricBytesIngress, AnalyticsEventSummaryGetParamsMetricBytesEgress, AnalyticsEventSummaryGetParamsMetricDurationAvg, AnalyticsEventSummaryGetParamsMetricDurationMedian, AnalyticsEventSummaryGetParamsMetricDuration90th, AnalyticsEventSummaryGetParamsMetricDuration99th:
		return true
	}
	return false
}

type AnalyticsEventSummaryGetResponseEnvelope struct {
	Errors   []AnalyticsEventSummaryGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsEventSummaryGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AnalyticsEventSummaryGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AnalyticsEventSummaryGetResponse                `json:"result"`
	JSON    analyticsEventSummaryGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsEventSummaryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsEventSummaryGetResponseEnvelope]
type analyticsEventSummaryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventSummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventSummaryGetResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AnalyticsEventSummaryGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             analyticsEventSummaryGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// analyticsEventSummaryGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AnalyticsEventSummaryGetResponseEnvelopeErrors]
type analyticsEventSummaryGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsEventSummaryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventSummaryGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    analyticsEventSummaryGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// analyticsEventSummaryGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AnalyticsEventSummaryGetResponseEnvelopeErrorsSource]
type analyticsEventSummaryGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventSummaryGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventSummaryGetResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AnalyticsEventSummaryGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             analyticsEventSummaryGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// analyticsEventSummaryGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AnalyticsEventSummaryGetResponseEnvelopeMessages]
type analyticsEventSummaryGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsEventSummaryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsEventSummaryGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    analyticsEventSummaryGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// analyticsEventSummaryGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AnalyticsEventSummaryGetResponseEnvelopeMessagesSource]
type analyticsEventSummaryGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventSummaryGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AnalyticsEventSummaryGetResponseEnvelopeSuccess bool

const (
	AnalyticsEventSummaryGetResponseEnvelopeSuccessTrue AnalyticsEventSummaryGetResponseEnvelopeSuccess = true
)

func (r AnalyticsEventSummaryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsEventSummaryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
