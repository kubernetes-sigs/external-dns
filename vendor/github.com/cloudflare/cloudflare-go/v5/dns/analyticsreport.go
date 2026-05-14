// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// AnalyticsReportService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsReportService] method instead.
type AnalyticsReportService struct {
	Options []option.RequestOption
	Bytimes *AnalyticsReportBytimeService
}

// NewAnalyticsReportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsReportService(opts ...option.RequestOption) (r *AnalyticsReportService) {
	r = &AnalyticsReportService{}
	r.Options = opts
	r.Bytimes = NewAnalyticsReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AnalyticsReportService) Get(ctx context.Context, params AnalyticsReportGetParams, opts ...option.RequestOption) (res *Report, err error) {
	var env AnalyticsReportGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_analytics/report", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Report struct {
	// Array with one row per combination of dimension values.
	Data []ReportData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{} `json:"min,required"`
	Query ReportQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{} `json:"totals,required"`
	JSON   reportJSON  `json:"-"`
}

// reportJSON contains the JSON metadata for the struct [Report]
type reportJSON struct {
	Data        apijson.Field
	DataLag     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	Query       apijson.Field
	Rows        apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Report) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportJSON) RawJSON() string {
	return r.raw
}

type ReportData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64      `json:"metrics,required"`
	JSON    reportDataJSON `json:"-"`
}

// reportDataJSON contains the JSON metadata for the struct [ReportData]
type reportDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportDataJSON) RawJSON() string {
	return r.raw
}

type ReportQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string        `json:"sort"`
	JSON reportQueryJSON `json:"-"`
}

// reportQueryJSON contains the JSON metadata for the struct [ReportQuery]
type reportQueryJSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReportQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reportQueryJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsReportGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AnalyticsReportGetResponseEnvelope struct {
	Errors   []AnalyticsReportGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsReportGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AnalyticsReportGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Report                                    `json:"result"`
	JSON    analyticsReportGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsReportGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnalyticsReportGetResponseEnvelope]
type analyticsReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           AnalyticsReportGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             analyticsReportGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// analyticsReportGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AnalyticsReportGetResponseEnvelopeErrors]
type analyticsReportGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    analyticsReportGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// analyticsReportGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [AnalyticsReportGetResponseEnvelopeErrorsSource]
type analyticsReportGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           AnalyticsReportGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             analyticsReportGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// analyticsReportGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AnalyticsReportGetResponseEnvelopeMessages]
type analyticsReportGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    analyticsReportGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// analyticsReportGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [AnalyticsReportGetResponseEnvelopeMessagesSource]
type analyticsReportGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AnalyticsReportGetResponseEnvelopeSuccess bool

const (
	AnalyticsReportGetResponseEnvelopeSuccessTrue AnalyticsReportGetResponseEnvelopeSuccess = true
)

func (r AnalyticsReportGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsReportGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
