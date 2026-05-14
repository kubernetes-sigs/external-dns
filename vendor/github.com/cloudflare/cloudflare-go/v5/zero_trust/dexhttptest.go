// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DEXHTTPTestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXHTTPTestService] method instead.
type DEXHTTPTestService struct {
	Options     []option.RequestOption
	Percentiles *DEXHTTPTestPercentileService
}

// NewDEXHTTPTestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDEXHTTPTestService(opts ...option.RequestOption) (r *DEXHTTPTestService) {
	r = &DEXHTTPTestService{}
	r.Options = opts
	r.Percentiles = NewDEXHTTPTestPercentileService(opts...)
	return
}

// Get test details and aggregate performance metrics for an http test for a given
// time period between 1 hour and 7 days.
func (r *DEXHTTPTestService) Get(ctx context.Context, testID string, params DEXHTTPTestGetParams, opts ...option.RequestOption) (res *HTTPDetails, err error) {
	var env DexhttpTestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/http-tests/%s", params.AccountID, testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPDetails struct {
	// The url of the HTTP synthetic application test
	Host            string                       `json:"host"`
	HTTPStats       HTTPDetailsHTTPStats         `json:"httpStats,nullable"`
	HTTPStatsByColo []HTTPDetailsHTTPStatsByColo `json:"httpStatsByColo"`
	// The interval at which the HTTP synthetic application test is set to run.
	Interval string          `json:"interval"`
	Kind     HTTPDetailsKind `json:"kind"`
	// The HTTP method to use when running the test
	Method string `json:"method"`
	// The name of the HTTP synthetic application test
	Name           string                     `json:"name"`
	TargetPolicies []DigitalExperienceMonitor `json:"target_policies,nullable"`
	Targeted       bool                       `json:"targeted"`
	JSON           httpDetailsJSON            `json:"-"`
}

// httpDetailsJSON contains the JSON metadata for the struct [HTTPDetails]
type httpDetailsJSON struct {
	Host            apijson.Field
	HTTPStats       apijson.Field
	HTTPStatsByColo apijson.Field
	Interval        apijson.Field
	Kind            apijson.Field
	Method          apijson.Field
	Name            apijson.Field
	TargetPolicies  apijson.Field
	Targeted        apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *HTTPDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStats struct {
	AvailabilityPct      HTTPDetailsHTTPStatsAvailabilityPct  `json:"availabilityPct,required"`
	DNSResponseTimeMs    TestStatOverTime                     `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []HTTPDetailsHTTPStatsHTTPStatusCode `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  TestStatOverTime                     `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs TestStatOverTime                     `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                    `json:"uniqueDevicesTotal,required"`
	JSON               httpDetailsHTTPStatsJSON `json:"-"`
}

// httpDetailsHTTPStatsJSON contains the JSON metadata for the struct
// [HTTPDetailsHTTPStats]
type httpDetailsHTTPStatsJSON struct {
	AvailabilityPct      apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStatsAvailabilityPct struct {
	Slots []HTTPDetailsHTTPStatsAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                 `json:"min,nullable"`
	JSON httpDetailsHTTPStatsAvailabilityPctJSON `json:"-"`
}

// httpDetailsHTTPStatsAvailabilityPctJSON contains the JSON metadata for the
// struct [HTTPDetailsHTTPStatsAvailabilityPct]
type httpDetailsHTTPStatsAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStatsAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStatsAvailabilityPctSlot struct {
	Timestamp string                                      `json:"timestamp,required"`
	Value     float64                                     `json:"value,required"`
	JSON      httpDetailsHTTPStatsAvailabilityPctSlotJSON `json:"-"`
}

// httpDetailsHTTPStatsAvailabilityPctSlotJSON contains the JSON metadata for the
// struct [HTTPDetailsHTTPStatsAvailabilityPctSlot]
type httpDetailsHTTPStatsAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStatsAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStatsHTTPStatusCode struct {
	Status200 int64                                  `json:"status200,required"`
	Status300 int64                                  `json:"status300,required"`
	Status400 int64                                  `json:"status400,required"`
	Status500 int64                                  `json:"status500,required"`
	Timestamp string                                 `json:"timestamp,required"`
	JSON      httpDetailsHTTPStatsHTTPStatusCodeJSON `json:"-"`
}

// httpDetailsHTTPStatsHTTPStatusCodeJSON contains the JSON metadata for the struct
// [HTTPDetailsHTTPStatsHTTPStatusCode]
type httpDetailsHTTPStatsHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStatsHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsHTTPStatusCodeJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStatsByColo struct {
	AvailabilityPct      HTTPDetailsHTTPStatsByColoAvailabilityPct  `json:"availabilityPct,required"`
	Colo                 string                                     `json:"colo,required"`
	DNSResponseTimeMs    TestStatOverTime                           `json:"dnsResponseTimeMs,required"`
	HTTPStatusCode       []HTTPDetailsHTTPStatsByColoHTTPStatusCode `json:"httpStatusCode,required"`
	ResourceFetchTimeMs  TestStatOverTime                           `json:"resourceFetchTimeMs,required"`
	ServerResponseTimeMs TestStatOverTime                           `json:"serverResponseTimeMs,required"`
	// Count of unique devices that have run this test in the given time period
	UniqueDevicesTotal int64                          `json:"uniqueDevicesTotal,required"`
	JSON               httpDetailsHTTPStatsByColoJSON `json:"-"`
}

// httpDetailsHTTPStatsByColoJSON contains the JSON metadata for the struct
// [HTTPDetailsHTTPStatsByColo]
type httpDetailsHTTPStatsByColoJSON struct {
	AvailabilityPct      apijson.Field
	Colo                 apijson.Field
	DNSResponseTimeMs    apijson.Field
	HTTPStatusCode       apijson.Field
	ResourceFetchTimeMs  apijson.Field
	ServerResponseTimeMs apijson.Field
	UniqueDevicesTotal   apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStatsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsByColoJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStatsByColoAvailabilityPct struct {
	Slots []HTTPDetailsHTTPStatsByColoAvailabilityPctSlot `json:"slots,required"`
	// average observed in the time period
	Avg float64 `json:"avg,nullable"`
	// highest observed in the time period
	Max float64 `json:"max,nullable"`
	// lowest observed in the time period
	Min  float64                                       `json:"min,nullable"`
	JSON httpDetailsHTTPStatsByColoAvailabilityPctJSON `json:"-"`
}

// httpDetailsHTTPStatsByColoAvailabilityPctJSON contains the JSON metadata for the
// struct [HTTPDetailsHTTPStatsByColoAvailabilityPct]
type httpDetailsHTTPStatsByColoAvailabilityPctJSON struct {
	Slots       apijson.Field
	Avg         apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStatsByColoAvailabilityPct) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsByColoAvailabilityPctJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStatsByColoAvailabilityPctSlot struct {
	Timestamp string                                            `json:"timestamp,required"`
	Value     float64                                           `json:"value,required"`
	JSON      httpDetailsHTTPStatsByColoAvailabilityPctSlotJSON `json:"-"`
}

// httpDetailsHTTPStatsByColoAvailabilityPctSlotJSON contains the JSON metadata for
// the struct [HTTPDetailsHTTPStatsByColoAvailabilityPctSlot]
type httpDetailsHTTPStatsByColoAvailabilityPctSlotJSON struct {
	Timestamp   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStatsByColoAvailabilityPctSlot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsByColoAvailabilityPctSlotJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsHTTPStatsByColoHTTPStatusCode struct {
	Status200 int64                                        `json:"status200,required"`
	Status300 int64                                        `json:"status300,required"`
	Status400 int64                                        `json:"status400,required"`
	Status500 int64                                        `json:"status500,required"`
	Timestamp string                                       `json:"timestamp,required"`
	JSON      httpDetailsHTTPStatsByColoHTTPStatusCodeJSON `json:"-"`
}

// httpDetailsHTTPStatsByColoHTTPStatusCodeJSON contains the JSON metadata for the
// struct [HTTPDetailsHTTPStatsByColoHTTPStatusCode]
type httpDetailsHTTPStatsByColoHTTPStatusCodeJSON struct {
	Status200   apijson.Field
	Status300   apijson.Field
	Status400   apijson.Field
	Status500   apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPDetailsHTTPStatsByColoHTTPStatusCode) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpDetailsHTTPStatsByColoHTTPStatusCodeJSON) RawJSON() string {
	return r.raw
}

type HTTPDetailsKind string

const (
	HTTPDetailsKindHTTP HTTPDetailsKind = "http"
)

func (r HTTPDetailsKind) IsKnown() bool {
	switch r {
	case HTTPDetailsKindHTTP:
		return true
	}
	return false
}

type DEXHTTPTestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Start time for aggregate metrics in ISO ms
	From param.Field[string] `query:"from,required"`
	// Time interval for aggregate time slots.
	Interval param.Field[DexhttpTestGetParamsInterval] `query:"interval,required"`
	// End time for aggregate metrics in ISO ms
	To param.Field[string] `query:"to,required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
}

// URLQuery serializes [DEXHTTPTestGetParams]'s query parameters as `url.Values`.
func (r DEXHTTPTestGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Time interval for aggregate time slots.
type DexhttpTestGetParamsInterval string

const (
	DexhttpTestGetParamsIntervalMinute DexhttpTestGetParamsInterval = "minute"
	DexhttpTestGetParamsIntervalHour   DexhttpTestGetParamsInterval = "hour"
)

func (r DexhttpTestGetParamsInterval) IsKnown() bool {
	switch r {
	case DexhttpTestGetParamsIntervalMinute, DexhttpTestGetParamsIntervalHour:
		return true
	}
	return false
}

type DexhttpTestGetResponseEnvelope struct {
	Errors   []DexhttpTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DexhttpTestGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DexhttpTestGetResponseEnvelopeSuccess `json:"success,required"`
	Result  HTTPDetails                           `json:"result"`
	JSON    dexhttpTestGetResponseEnvelopeJSON    `json:"-"`
}

// dexhttpTestGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DexhttpTestGetResponseEnvelope]
type dexhttpTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DexhttpTestGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DexhttpTestGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexhttpTestGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexhttpTestGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DexhttpTestGetResponseEnvelopeErrors]
type dexhttpTestGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DexhttpTestGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dexhttpTestGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexhttpTestGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DexhttpTestGetResponseEnvelopeErrorsSource]
type dexhttpTestGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DexhttpTestGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DexhttpTestGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexhttpTestGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexhttpTestGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DexhttpTestGetResponseEnvelopeMessages]
type dexhttpTestGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DexhttpTestGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dexhttpTestGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexhttpTestGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DexhttpTestGetResponseEnvelopeMessagesSource]
type dexhttpTestGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DexhttpTestGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexhttpTestGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DexhttpTestGetResponseEnvelopeSuccess bool

const (
	DexhttpTestGetResponseEnvelopeSuccessTrue DexhttpTestGetResponseEnvelopeSuccess = true
)

func (r DexhttpTestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DexhttpTestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
