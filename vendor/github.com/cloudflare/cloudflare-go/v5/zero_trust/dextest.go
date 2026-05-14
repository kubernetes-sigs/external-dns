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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DEXTestService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXTestService] method instead.
type DEXTestService struct {
	Options       []option.RequestOption
	UniqueDevices *DEXTestUniqueDeviceService
}

// NewDEXTestService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXTestService(opts ...option.RequestOption) (r *DEXTestService) {
	r = &DEXTestService{}
	r.Options = opts
	r.UniqueDevices = NewDEXTestUniqueDeviceService(opts...)
	return
}

// List DEX tests with overview metrics
func (r *DEXTestService) List(ctx context.Context, params DEXTestListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[Tests], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/tests/overview", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List DEX tests with overview metrics
func (r *DEXTestService) ListAutoPaging(ctx context.Context, params DEXTestListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[Tests] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type AggregateTimePeriod struct {
	Units AggregateTimePeriodUnits `json:"units,required"`
	Value int64                    `json:"value,required"`
	JSON  aggregateTimePeriodJSON  `json:"-"`
}

// aggregateTimePeriodJSON contains the JSON metadata for the struct
// [AggregateTimePeriod]
type aggregateTimePeriodJSON struct {
	Units       apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AggregateTimePeriod) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aggregateTimePeriodJSON) RawJSON() string {
	return r.raw
}

type AggregateTimePeriodUnits string

const (
	AggregateTimePeriodUnitsHours    AggregateTimePeriodUnits = "hours"
	AggregateTimePeriodUnitsDays     AggregateTimePeriodUnits = "days"
	AggregateTimePeriodUnitsTestRuns AggregateTimePeriodUnits = "testRuns"
)

func (r AggregateTimePeriodUnits) IsKnown() bool {
	switch r {
	case AggregateTimePeriodUnitsHours, AggregateTimePeriodUnitsDays, AggregateTimePeriodUnitsTestRuns:
		return true
	}
	return false
}

type Tests struct {
	OverviewMetrics TestsOverviewMetrics `json:"overviewMetrics,required"`
	// array of test results objects.
	Tests []TestsTest `json:"tests,required"`
	JSON  testsJSON   `json:"-"`
}

// testsJSON contains the JSON metadata for the struct [Tests]
type testsJSON struct {
	OverviewMetrics apijson.Field
	Tests           apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Tests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsJSON) RawJSON() string {
	return r.raw
}

type TestsOverviewMetrics struct {
	// number of tests.
	TestsTotal int64 `json:"testsTotal,required"`
	// percentage availability for all HTTP test results in response
	AvgHTTPAvailabilityPct float64 `json:"avgHttpAvailabilityPct,nullable"`
	// percentage availability for all traceroutes results in response
	AvgTracerouteAvailabilityPct float64                  `json:"avgTracerouteAvailabilityPct,nullable"`
	JSON                         testsOverviewMetricsJSON `json:"-"`
}

// testsOverviewMetricsJSON contains the JSON metadata for the struct
// [TestsOverviewMetrics]
type testsOverviewMetricsJSON struct {
	TestsTotal                   apijson.Field
	AvgHTTPAvailabilityPct       apijson.Field
	AvgTracerouteAvailabilityPct apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *TestsOverviewMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsOverviewMetricsJSON) RawJSON() string {
	return r.raw
}

type TestsTest struct {
	// API Resource UUID tag.
	ID string `json:"id,required"`
	// date the test was created.
	Created string `json:"created,required"`
	// the test description defined during configuration
	Description string `json:"description,required"`
	// if true, then the test will run on targeted devices. Else, the test will not
	// run.
	Enabled bool   `json:"enabled,required"`
	Host    string `json:"host,required"`
	// The interval at which the synthetic application test is set to run.
	Interval string `json:"interval,required"`
	// test type, http or traceroute
	Kind TestsTestsKind `json:"kind,required"`
	// name given to this test
	Name              string                        `json:"name,required"`
	Updated           string                        `json:"updated,required"`
	HTTPResults       TestsTestsHTTPResults         `json:"httpResults,nullable"`
	HTTPResultsByColo []TestsTestsHTTPResultsByColo `json:"httpResultsByColo"`
	// for HTTP, the method to use when running the test
	Method                  string                              `json:"method"`
	TargetPolicies          []DigitalExperienceMonitor          `json:"target_policies,nullable"`
	Targeted                bool                                `json:"targeted"`
	TracerouteResults       TestsTestsTracerouteResults         `json:"tracerouteResults,nullable"`
	TracerouteResultsByColo []TestsTestsTracerouteResultsByColo `json:"tracerouteResultsByColo"`
	JSON                    testsTestJSON                       `json:"-"`
}

// testsTestJSON contains the JSON metadata for the struct [TestsTest]
type testsTestJSON struct {
	ID                      apijson.Field
	Created                 apijson.Field
	Description             apijson.Field
	Enabled                 apijson.Field
	Host                    apijson.Field
	Interval                apijson.Field
	Kind                    apijson.Field
	Name                    apijson.Field
	Updated                 apijson.Field
	HTTPResults             apijson.Field
	HTTPResultsByColo       apijson.Field
	Method                  apijson.Field
	TargetPolicies          apijson.Field
	Targeted                apijson.Field
	TracerouteResults       apijson.Field
	TracerouteResultsByColo apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TestsTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestJSON) RawJSON() string {
	return r.raw
}

// test type, http or traceroute
type TestsTestsKind string

const (
	TestsTestsKindHTTP       TestsTestsKind = "http"
	TestsTestsKindTraceroute TestsTestsKind = "traceroute"
)

func (r TestsTestsKind) IsKnown() bool {
	switch r {
	case TestsTestsKindHTTP, TestsTestsKindTraceroute:
		return true
	}
	return false
}

type TestsTestsHTTPResults struct {
	ResourceFetchTime TestsTestsHTTPResultsResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              testsTestsHTTPResultsJSON              `json:"-"`
}

// testsTestsHTTPResultsJSON contains the JSON metadata for the struct
// [TestsTestsHTTPResults]
type testsTestsHTTPResultsJSON struct {
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestsTestsHTTPResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsResourceFetchTime struct {
	History  []TestsTestsHTTPResultsResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                           `json:"avgMs,nullable"`
	OverTime TestsTestsHTTPResultsResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     testsTestsHTTPResultsResourceFetchTimeJSON      `json:"-"`
}

// testsTestsHTTPResultsResourceFetchTimeJSON contains the JSON metadata for the
// struct [TestsTestsHTTPResultsResourceFetchTime]
type testsTestsHTTPResultsResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsResourceFetchTimeHistory struct {
	TimePeriod AggregateTimePeriod                               `json:"timePeriod,required"`
	AvgMs      int64                                             `json:"avgMs,nullable"`
	DeltaPct   float64                                           `json:"deltaPct,nullable"`
	JSON       testsTestsHTTPResultsResourceFetchTimeHistoryJSON `json:"-"`
}

// testsTestsHTTPResultsResourceFetchTimeHistoryJSON contains the JSON metadata for
// the struct [TestsTestsHTTPResultsResourceFetchTimeHistory]
type testsTestsHTTPResultsResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsResourceFetchTimeOverTime struct {
	TimePeriod AggregateTimePeriod                                   `json:"timePeriod,required"`
	Values     []TestsTestsHTTPResultsResourceFetchTimeOverTimeValue `json:"values,required"`
	JSON       testsTestsHTTPResultsResourceFetchTimeOverTimeJSON    `json:"-"`
}

// testsTestsHTTPResultsResourceFetchTimeOverTimeJSON contains the JSON metadata
// for the struct [TestsTestsHTTPResultsResourceFetchTimeOverTime]
type testsTestsHTTPResultsResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                   `json:"avgMs,required"`
	Timestamp string                                                  `json:"timestamp,required"`
	JSON      testsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// testsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON contains the JSON
// metadata for the struct [TestsTestsHTTPResultsResourceFetchTimeOverTimeValue]
type testsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsByColo struct {
	// Cloudflare colo
	Colo              string                                       `json:"colo,required"`
	ResourceFetchTime TestsTestsHTTPResultsByColoResourceFetchTime `json:"resourceFetchTime,required"`
	JSON              testsTestsHTTPResultsByColoJSON              `json:"-"`
}

// testsTestsHTTPResultsByColoJSON contains the JSON metadata for the struct
// [TestsTestsHTTPResultsByColo]
type testsTestsHTTPResultsByColoJSON struct {
	Colo              apijson.Field
	ResourceFetchTime apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsByColoJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsByColoResourceFetchTime struct {
	History  []TestsTestsHTTPResultsByColoResourceFetchTimeHistory `json:"history,required"`
	AvgMs    int64                                                 `json:"avgMs,nullable"`
	OverTime TestsTestsHTTPResultsByColoResourceFetchTimeOverTime  `json:"overTime,nullable"`
	JSON     testsTestsHTTPResultsByColoResourceFetchTimeJSON      `json:"-"`
}

// testsTestsHTTPResultsByColoResourceFetchTimeJSON contains the JSON metadata for
// the struct [TestsTestsHTTPResultsByColoResourceFetchTime]
type testsTestsHTTPResultsByColoResourceFetchTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsByColoResourceFetchTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsByColoResourceFetchTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsByColoResourceFetchTimeHistory struct {
	TimePeriod AggregateTimePeriod                                     `json:"timePeriod,required"`
	AvgMs      int64                                                   `json:"avgMs,nullable"`
	DeltaPct   float64                                                 `json:"deltaPct,nullable"`
	JSON       testsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON `json:"-"`
}

// testsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON contains the JSON
// metadata for the struct [TestsTestsHTTPResultsByColoResourceFetchTimeHistory]
type testsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsByColoResourceFetchTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsByColoResourceFetchTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsByColoResourceFetchTimeOverTime struct {
	TimePeriod AggregateTimePeriod                                         `json:"timePeriod,required"`
	Values     []TestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue `json:"values,required"`
	JSON       testsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON    `json:"-"`
}

// testsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON contains the JSON
// metadata for the struct [TestsTestsHTTPResultsByColoResourceFetchTimeOverTime]
type testsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsByColoResourceFetchTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsByColoResourceFetchTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue struct {
	AvgMs     int64                                                         `json:"avgMs,required"`
	Timestamp string                                                        `json:"timestamp,required"`
	JSON      testsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON `json:"-"`
}

// testsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON contains the JSON
// metadata for the struct
// [TestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue]
type testsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsHTTPResultsByColoResourceFetchTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsHTTPResultsByColoResourceFetchTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResults struct {
	RoundTripTime TestsTestsTracerouteResultsRoundTripTime `json:"roundTripTime,required"`
	JSON          testsTestsTracerouteResultsJSON          `json:"-"`
}

// testsTestsTracerouteResultsJSON contains the JSON metadata for the struct
// [TestsTestsTracerouteResults]
type testsTestsTracerouteResultsJSON struct {
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TestsTestsTracerouteResults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsRoundTripTime struct {
	History  []TestsTestsTracerouteResultsRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                             `json:"avgMs,nullable"`
	OverTime TestsTestsTracerouteResultsRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     testsTestsTracerouteResultsRoundTripTimeJSON      `json:"-"`
}

// testsTestsTracerouteResultsRoundTripTimeJSON contains the JSON metadata for the
// struct [TestsTestsTracerouteResultsRoundTripTime]
type testsTestsTracerouteResultsRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsRoundTripTimeHistory struct {
	TimePeriod AggregateTimePeriod                                 `json:"timePeriod,required"`
	AvgMs      int64                                               `json:"avgMs,nullable"`
	DeltaPct   float64                                             `json:"deltaPct,nullable"`
	JSON       testsTestsTracerouteResultsRoundTripTimeHistoryJSON `json:"-"`
}

// testsTestsTracerouteResultsRoundTripTimeHistoryJSON contains the JSON metadata
// for the struct [TestsTestsTracerouteResultsRoundTripTimeHistory]
type testsTestsTracerouteResultsRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsRoundTripTimeOverTime struct {
	TimePeriod AggregateTimePeriod                                     `json:"timePeriod,required"`
	Values     []TestsTestsTracerouteResultsRoundTripTimeOverTimeValue `json:"values,required"`
	JSON       testsTestsTracerouteResultsRoundTripTimeOverTimeJSON    `json:"-"`
}

// testsTestsTracerouteResultsRoundTripTimeOverTimeJSON contains the JSON metadata
// for the struct [TestsTestsTracerouteResultsRoundTripTimeOverTime]
type testsTestsTracerouteResultsRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                     `json:"avgMs,required"`
	Timestamp string                                                    `json:"timestamp,required"`
	JSON      testsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON `json:"-"`
}

// testsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON contains the JSON
// metadata for the struct [TestsTestsTracerouteResultsRoundTripTimeOverTimeValue]
type testsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsRoundTripTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsByColo struct {
	// Cloudflare colo
	Colo          string                                         `json:"colo,required"`
	RoundTripTime TestsTestsTracerouteResultsByColoRoundTripTime `json:"roundTripTime,required"`
	JSON          testsTestsTracerouteResultsByColoJSON          `json:"-"`
}

// testsTestsTracerouteResultsByColoJSON contains the JSON metadata for the struct
// [TestsTestsTracerouteResultsByColo]
type testsTestsTracerouteResultsByColoJSON struct {
	Colo          apijson.Field
	RoundTripTime apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsByColo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsByColoJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsByColoRoundTripTime struct {
	History  []TestsTestsTracerouteResultsByColoRoundTripTimeHistory `json:"history,required"`
	AvgMs    int64                                                   `json:"avgMs,nullable"`
	OverTime TestsTestsTracerouteResultsByColoRoundTripTimeOverTime  `json:"overTime,nullable"`
	JSON     testsTestsTracerouteResultsByColoRoundTripTimeJSON      `json:"-"`
}

// testsTestsTracerouteResultsByColoRoundTripTimeJSON contains the JSON metadata
// for the struct [TestsTestsTracerouteResultsByColoRoundTripTime]
type testsTestsTracerouteResultsByColoRoundTripTimeJSON struct {
	History     apijson.Field
	AvgMs       apijson.Field
	OverTime    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsByColoRoundTripTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsByColoRoundTripTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsByColoRoundTripTimeHistory struct {
	TimePeriod AggregateTimePeriod                                       `json:"timePeriod,required"`
	AvgMs      int64                                                     `json:"avgMs,nullable"`
	DeltaPct   float64                                                   `json:"deltaPct,nullable"`
	JSON       testsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON `json:"-"`
}

// testsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON contains the JSON
// metadata for the struct [TestsTestsTracerouteResultsByColoRoundTripTimeHistory]
type testsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON struct {
	TimePeriod  apijson.Field
	AvgMs       apijson.Field
	DeltaPct    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsByColoRoundTripTimeHistory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsByColoRoundTripTimeHistoryJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsByColoRoundTripTimeOverTime struct {
	TimePeriod AggregateTimePeriod                                           `json:"timePeriod,required"`
	Values     []TestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue `json:"values,required"`
	JSON       testsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON    `json:"-"`
}

// testsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON contains the JSON
// metadata for the struct [TestsTestsTracerouteResultsByColoRoundTripTimeOverTime]
type testsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON struct {
	TimePeriod  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsByColoRoundTripTimeOverTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsByColoRoundTripTimeOverTimeJSON) RawJSON() string {
	return r.raw
}

type TestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue struct {
	AvgMs     int64                                                           `json:"avgMs,required"`
	Timestamp string                                                          `json:"timestamp,required"`
	JSON      testsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON `json:"-"`
}

// testsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON contains the
// JSON metadata for the struct
// [TestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue]
type testsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON struct {
	AvgMs       apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TestsTestsTracerouteResultsByColoRoundTripTimeOverTimeValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r testsTestsTracerouteResultsByColoRoundTripTimeOverTimeValueJSON) RawJSON() string {
	return r.raw
}

type DEXTestListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Optionally filter result stats to a Cloudflare colo. Cannot be used in
	// combination with deviceId param.
	Colo param.Field[string] `query:"colo"`
	// Optionally filter result stats to a specific device(s). Cannot be used in
	// combination with colo param.
	DeviceID param.Field[[]string] `query:"deviceId"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page"`
	// Optionally filter results by test name
	TestName param.Field[string] `query:"testName"`
}

// URLQuery serializes [DEXTestListParams]'s query parameters as `url.Values`.
func (r DEXTestListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
