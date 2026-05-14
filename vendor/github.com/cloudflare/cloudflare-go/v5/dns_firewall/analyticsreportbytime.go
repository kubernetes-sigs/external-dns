// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AnalyticsReportBytimeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsReportBytimeService] method instead.
type AnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewAnalyticsReportBytimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsReportBytimeService(opts ...option.RequestOption) (r *AnalyticsReportBytimeService) {
	r = &AnalyticsReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AnalyticsReportBytimeService) Get(ctx context.Context, dnsFirewallID string, params AnalyticsReportBytimeGetParams, opts ...option.RequestOption) (res *dns.ByTime, err error) {
	var env AnalyticsReportBytimeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", params.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AnalyticsReportBytimeGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
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
	// Unit of time to group data by.
	TimeDelta param.Field[AnalyticsReportBytimeGetParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsReportBytimeGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsReportBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Unit of time to group data by.
type AnalyticsReportBytimeGetParamsTimeDelta string

const (
	AnalyticsReportBytimeGetParamsTimeDeltaAll        AnalyticsReportBytimeGetParamsTimeDelta = "all"
	AnalyticsReportBytimeGetParamsTimeDeltaAuto       AnalyticsReportBytimeGetParamsTimeDelta = "auto"
	AnalyticsReportBytimeGetParamsTimeDeltaYear       AnalyticsReportBytimeGetParamsTimeDelta = "year"
	AnalyticsReportBytimeGetParamsTimeDeltaQuarter    AnalyticsReportBytimeGetParamsTimeDelta = "quarter"
	AnalyticsReportBytimeGetParamsTimeDeltaMonth      AnalyticsReportBytimeGetParamsTimeDelta = "month"
	AnalyticsReportBytimeGetParamsTimeDeltaWeek       AnalyticsReportBytimeGetParamsTimeDelta = "week"
	AnalyticsReportBytimeGetParamsTimeDeltaDay        AnalyticsReportBytimeGetParamsTimeDelta = "day"
	AnalyticsReportBytimeGetParamsTimeDeltaHour       AnalyticsReportBytimeGetParamsTimeDelta = "hour"
	AnalyticsReportBytimeGetParamsTimeDeltaDekaminute AnalyticsReportBytimeGetParamsTimeDelta = "dekaminute"
	AnalyticsReportBytimeGetParamsTimeDeltaMinute     AnalyticsReportBytimeGetParamsTimeDelta = "minute"
)

func (r AnalyticsReportBytimeGetParamsTimeDelta) IsKnown() bool {
	switch r {
	case AnalyticsReportBytimeGetParamsTimeDeltaAll, AnalyticsReportBytimeGetParamsTimeDeltaAuto, AnalyticsReportBytimeGetParamsTimeDeltaYear, AnalyticsReportBytimeGetParamsTimeDeltaQuarter, AnalyticsReportBytimeGetParamsTimeDeltaMonth, AnalyticsReportBytimeGetParamsTimeDeltaWeek, AnalyticsReportBytimeGetParamsTimeDeltaDay, AnalyticsReportBytimeGetParamsTimeDeltaHour, AnalyticsReportBytimeGetParamsTimeDeltaDekaminute, AnalyticsReportBytimeGetParamsTimeDeltaMinute:
		return true
	}
	return false
}

type AnalyticsReportBytimeGetResponseEnvelope struct {
	Errors   []AnalyticsReportBytimeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsReportBytimeGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AnalyticsReportBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	Result  dns.ByTime                                      `json:"result"`
	JSON    analyticsReportBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsReportBytimeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsReportBytimeGetResponseEnvelope]
type analyticsReportBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportBytimeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportBytimeGetResponseEnvelopeErrors struct {
	Code             int64                                                `json:"code,required"`
	Message          string                                               `json:"message,required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           AnalyticsReportBytimeGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             analyticsReportBytimeGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// analyticsReportBytimeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AnalyticsReportBytimeGetResponseEnvelopeErrors]
type analyticsReportBytimeGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsReportBytimeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportBytimeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportBytimeGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    analyticsReportBytimeGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// analyticsReportBytimeGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AnalyticsReportBytimeGetResponseEnvelopeErrorsSource]
type analyticsReportBytimeGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportBytimeGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportBytimeGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportBytimeGetResponseEnvelopeMessages struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AnalyticsReportBytimeGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             analyticsReportBytimeGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// analyticsReportBytimeGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AnalyticsReportBytimeGetResponseEnvelopeMessages]
type analyticsReportBytimeGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AnalyticsReportBytimeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportBytimeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportBytimeGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    analyticsReportBytimeGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// analyticsReportBytimeGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [AnalyticsReportBytimeGetResponseEnvelopeMessagesSource]
type analyticsReportBytimeGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportBytimeGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportBytimeGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AnalyticsReportBytimeGetResponseEnvelopeSuccess bool

const (
	AnalyticsReportBytimeGetResponseEnvelopeSuccessTrue AnalyticsReportBytimeGetResponseEnvelopeSuccess = true
)

func (r AnalyticsReportBytimeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsReportBytimeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
