// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// ObservabilityTelemetryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityTelemetryService] method instead.
type ObservabilityTelemetryService struct {
	Options []option.RequestOption
}

// NewObservabilityTelemetryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservabilityTelemetryService(opts ...option.RequestOption) (r *ObservabilityTelemetryService) {
	r = &ObservabilityTelemetryService{}
	r.Options = opts
	return
}

// List all the keys in your telemetry events.
func (r *ObservabilityTelemetryService) Keys(ctx context.Context, params ObservabilityTelemetryKeysParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityTelemetryKeysResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/keys", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// List all the keys in your telemetry events.
func (r *ObservabilityTelemetryService) KeysAutoPaging(ctx context.Context, params ObservabilityTelemetryKeysParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityTelemetryKeysResponse] {
	return pagination.NewSinglePageAutoPager(r.Keys(ctx, params, opts...))
}

// Runs a temporary or saved query
func (r *ObservabilityTelemetryService) Query(ctx context.Context, params ObservabilityTelemetryQueryParams, opts ...option.RequestOption) (res *ObservabilityTelemetryQueryResponse, err error) {
	var env ObservabilityTelemetryQueryResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/query", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List unique values found in your events
func (r *ObservabilityTelemetryService) Values(ctx context.Context, params ObservabilityTelemetryValuesParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityTelemetryValuesResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/values", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// List unique values found in your events
func (r *ObservabilityTelemetryService) ValuesAutoPaging(ctx context.Context, params ObservabilityTelemetryValuesParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityTelemetryValuesResponse] {
	return pagination.NewSinglePageAutoPager(r.Values(ctx, params, opts...))
}

type ObservabilityTelemetryKeysResponse struct {
	Key        string                                 `json:"key,required"`
	LastSeenAt float64                                `json:"lastSeenAt,required"`
	Type       ObservabilityTelemetryKeysResponseType `json:"type,required"`
	JSON       observabilityTelemetryKeysResponseJSON `json:"-"`
}

// observabilityTelemetryKeysResponseJSON contains the JSON metadata for the struct
// [ObservabilityTelemetryKeysResponse]
type observabilityTelemetryKeysResponseJSON struct {
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryKeysResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryKeysResponseType string

const (
	ObservabilityTelemetryKeysResponseTypeString  ObservabilityTelemetryKeysResponseType = "string"
	ObservabilityTelemetryKeysResponseTypeBoolean ObservabilityTelemetryKeysResponseType = "boolean"
	ObservabilityTelemetryKeysResponseTypeNumber  ObservabilityTelemetryKeysResponseType = "number"
)

func (r ObservabilityTelemetryKeysResponseType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysResponseTypeString, ObservabilityTelemetryKeysResponseTypeBoolean, ObservabilityTelemetryKeysResponseTypeNumber:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponse struct {
	// A Workers Observability Query Object
	Run ObservabilityTelemetryQueryResponseRun `json:"run,required"`
	// The statistics object contains information about query performance from the
	// database, it does not include any network latency
	Statistics   ObservabilityTelemetryQueryResponseStatistics              `json:"statistics,required"`
	Calculations []ObservabilityTelemetryQueryResponseCalculation           `json:"calculations"`
	Compare      []ObservabilityTelemetryQueryResponseCompare               `json:"compare"`
	Events       ObservabilityTelemetryQueryResponseEvents                  `json:"events"`
	Invocations  map[string][]ObservabilityTelemetryQueryResponseInvocation `json:"invocations"`
	Patterns     []ObservabilityTelemetryQueryResponsePattern               `json:"patterns"`
	JSON         observabilityTelemetryQueryResponseJSON                    `json:"-"`
}

// observabilityTelemetryQueryResponseJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponse]
type observabilityTelemetryQueryResponseJSON struct {
	Run          apijson.Field
	Statistics   apijson.Field
	Calculations apijson.Field
	Compare      apijson.Field
	Events       apijson.Field
	Invocations  apijson.Field
	Patterns     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseJSON) RawJSON() string {
	return r.raw
}

// A Workers Observability Query Object
type ObservabilityTelemetryQueryResponseRun struct {
	ID        string `json:"id,required"`
	AccountID string `json:"accountId,required"`
	Dry       bool   `json:"dry,required"`
	// Deprecated: deprecated
	EnvironmentID string                                          `json:"environmentId,required"`
	Granularity   float64                                         `json:"granularity,required"`
	Query         ObservabilityTelemetryQueryResponseRunQuery     `json:"query,required"`
	Status        ObservabilityTelemetryQueryResponseRunStatus    `json:"status,required"`
	Timeframe     ObservabilityTelemetryQueryResponseRunTimeframe `json:"timeframe,required"`
	UserID        string                                          `json:"userId,required"`
	// Deprecated: deprecated
	WorkspaceID string                                           `json:"workspaceId,required"`
	Created     string                                           `json:"created"`
	Statistics  ObservabilityTelemetryQueryResponseRunStatistics `json:"statistics"`
	Updated     string                                           `json:"updated"`
	JSON        observabilityTelemetryQueryResponseRunJSON       `json:"-"`
}

// observabilityTelemetryQueryResponseRunJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponseRun]
type observabilityTelemetryQueryResponseRunJSON struct {
	ID            apijson.Field
	AccountID     apijson.Field
	Dry           apijson.Field
	EnvironmentID apijson.Field
	Granularity   apijson.Field
	Query         apijson.Field
	Status        apijson.Field
	Timeframe     apijson.Field
	UserID        apijson.Field
	WorkspaceID   apijson.Field
	Created       apijson.Field
	Statistics    apijson.Field
	Updated       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQuery struct {
	// ID of the query
	ID          string `json:"id,required"`
	Created     string `json:"created,required"`
	Description string `json:"description,required,nullable"`
	// ID of your environment
	EnvironmentID string `json:"environmentId,required"`
	// Flag for alerts automatically created
	Generated bool `json:"generated,required,nullable"`
	// Query name
	Name       string                                                `json:"name,required,nullable"`
	Parameters ObservabilityTelemetryQueryResponseRunQueryParameters `json:"parameters,required"`
	Updated    string                                                `json:"updated,required"`
	UserID     string                                                `json:"userId,required"`
	// ID of your workspace
	WorkspaceID string                                          `json:"workspaceId,required"`
	JSON        observabilityTelemetryQueryResponseRunQueryJSON `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseRunQuery]
type observabilityTelemetryQueryResponseRunQueryJSON struct {
	ID            apijson.Field
	Created       apijson.Field
	Description   apijson.Field
	EnvironmentID apijson.Field
	Generated     apijson.Field
	Name          apijson.Field
	Parameters    apijson.Field
	Updated       apijson.Field
	UserID        apijson.Field
	WorkspaceID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations []ObservabilityTelemetryQueryResponseRunQueryParametersCalculation `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets []string `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination `json:"filterCombination"`
	// Configure the Filters to apply to the query.
	Filters []ObservabilityTelemetryQueryResponseRunQueryParametersFilter `json:"filters"`
	// Define how to group the results of the query.
	GroupBys []ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings []ObservabilityTelemetryQueryResponseRunQueryParametersHaving `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit int64 `json:"limit"`
	// Define an expression to search using full-text search.
	Needle ObservabilityTelemetryQueryResponseRunQueryParametersNeedle `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy `json:"orderBy"`
	JSON    observabilityTelemetryQueryResponseRunQueryParametersJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseRunQueryParameters]
type observabilityTelemetryQueryResponseRunQueryParametersJSON struct {
	Calculations      apijson.Field
	Datasets          apijson.Field
	FilterCombination apijson.Field
	Filters           apijson.Field
	GroupBys          apijson.Field
	Havings           apijson.Field
	Limit             apijson.Field
	Needle            apijson.Field
	OrderBy           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersCalculation struct {
	Operator ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator `json:"operator,required"`
	Alias    string                                                                    `json:"alias"`
	Key      string                                                                    `json:"key"`
	KeyType  ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType  `json:"keyType"`
	JSON     observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON      `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersCalculation]
type observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON struct {
	Operator    apijson.Field
	Alias       apijson.Field
	Key         apijson.Field
	KeyType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorUniq              ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "uniq"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCount             ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "count"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMax               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "max"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMin               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "min"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSum               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "sum"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvg               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "avg"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedian            ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "median"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001              ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p001"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p01"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p05"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p10"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p25"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p75"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p90"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p95"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p99"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999              ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p999"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddev            ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "stddev"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVariance          ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "variance"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountDistinct     ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountUppercase    ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "COUNT"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMaxUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "MAX"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMinUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "MIN"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSumUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "SUM"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvgUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "AVG"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedianUppercase   ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "MEDIAN"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001Uppercase     ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P001"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P01"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P05"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P10"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P25"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P75"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P90"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P95"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P99"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999Uppercase     ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P999"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddevUppercase   ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "STDDEV"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVarianceUppercase ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorUniq, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCount, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMax, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMin, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSum, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvg, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedian, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddev, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVariance, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountDistinct, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMaxUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMinUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSumUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvgUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedianUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddevUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeString  ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType = "string"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeNumber  ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType = "number"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeBoolean ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeString, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeNumber, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAnd          ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "and"
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOr           ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "or"
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAndUppercase ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "AND"
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOrUppercase  ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAnd, ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOr, ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAndUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersFilter struct {
	Key       string                                                                 `json:"key,required"`
	Operation ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation  `json:"operation,required"`
	Type      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType       `json:"type,required"`
	Value     ObservabilityTelemetryQueryResponseRunQueryParametersFiltersValueUnion `json:"value"`
	JSON      observabilityTelemetryQueryResponseRunQueryParametersFilterJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersFilterJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersFilter]
type observabilityTelemetryQueryResponseRunQueryParametersFilterJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersFilterJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludes            ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "includes"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIncludes         ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "not_includes"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWith          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "starts_with"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationRegex               ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "regex"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExists              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "exists"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIsNull              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "is_null"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIn                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "in"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIn               ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "not_in"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEq                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "eq"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNeq                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "neq"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGt                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "gt"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGte                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "gte"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLt                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "lt"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLte                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "lte"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEquals              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotEquals           ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "!="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreater             ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = ">"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreaterOrEquals     ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = ">="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLess                ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "<"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLessOrEquals        ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "<="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludesUppercase   ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "INCLUDES"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotInclude      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationMatchRegex          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExistsUppercase     ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "EXISTS"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotExist        ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationInUppercase         ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "IN"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotInUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "NOT_IN"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWithUppercase ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludes, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIncludes, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWith, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationRegex, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExists, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIsNull, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIn, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIn, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEq, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNeq, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGt, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGte, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLt, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLte, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreater, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreaterOrEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLess, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLessOrEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludesUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotInclude, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationMatchRegex, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExistsUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotExist, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationInUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotInUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeString  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType = "string"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeNumber  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType = "number"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeBoolean ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType = "boolean"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeString, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeNumber, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeBoolean:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseRunQueryParametersFiltersValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy struct {
	Type  ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType `json:"type,required"`
	Value string                                                            `json:"value,required"`
	JSON  observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON  `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy]
type observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeString  ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType = "string"
	ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeNumber  ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType = "number"
	ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeBoolean ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType = "boolean"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeString, ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeNumber, ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersHaving struct {
	Key       string                                                                `json:"key,required"`
	Operation ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation `json:"operation,required"`
	Value     float64                                                               `json:"value,required"`
	JSON      observabilityTelemetryQueryResponseRunQueryParametersHavingJSON       `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersHavingJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersHaving]
type observabilityTelemetryQueryResponseRunQueryParametersHavingJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersHaving) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersHavingJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationEq  ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "eq"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationNeq ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "neq"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGt  ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "gt"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGte ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "gte"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLt  ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "lt"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLte ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "lte"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationEq, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationNeq, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGt, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGte, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLt, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityTelemetryQueryResponseRunQueryParametersNeedle struct {
	Value     ObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion `json:"value,required"`
	IsRegex   bool                                                                  `json:"isRegex"`
	MatchCase bool                                                                  `json:"matchCase"`
	JSON      observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON       `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersNeedle]
type observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON struct {
	Value       apijson.Field
	IsRegex     apijson.Field
	MatchCase   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersNeedle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseRunQueryParametersNeedleValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Configure the order of the results returned by the query.
type ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value string `json:"value,required"`
	// Set the order of the results
	Order ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder `json:"order"`
	JSON  observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON  `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy]
type observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON struct {
	Value       apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON) RawJSON() string {
	return r.raw
}

// Set the order of the results
type ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderAsc  ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder = "asc"
	ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderDesc ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder = "desc"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderAsc, ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderDesc:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunStatus string

const (
	ObservabilityTelemetryQueryResponseRunStatusStarted   ObservabilityTelemetryQueryResponseRunStatus = "STARTED"
	ObservabilityTelemetryQueryResponseRunStatusCompleted ObservabilityTelemetryQueryResponseRunStatus = "COMPLETED"
)

func (r ObservabilityTelemetryQueryResponseRunStatus) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunStatusStarted, ObservabilityTelemetryQueryResponseRunStatusCompleted:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunTimeframe struct {
	// Set the start time for your query using UNIX time in milliseconds.
	From float64 `json:"from,required"`
	// Set the end time for your query using UNIX time in milliseconds.
	To   float64                                             `json:"to,required"`
	JSON observabilityTelemetryQueryResponseRunTimeframeJSON `json:"-"`
}

// observabilityTelemetryQueryResponseRunTimeframeJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseRunTimeframe]
type observabilityTelemetryQueryResponseRunTimeframeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunTimeframe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunTimeframeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read,required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed,required"`
	// Number of rows scanned from the table.
	RowsRead float64                                              `json:"rows_read,required"`
	JSON     observabilityTelemetryQueryResponseRunStatisticsJSON `json:"-"`
}

// observabilityTelemetryQueryResponseRunStatisticsJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseRunStatistics]
type observabilityTelemetryQueryResponseRunStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunStatisticsJSON) RawJSON() string {
	return r.raw
}

// The statistics object contains information about query performance from the
// database, it does not include any network latency
type ObservabilityTelemetryQueryResponseStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read,required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed,required"`
	// Number of rows scanned from the table.
	RowsRead float64                                           `json:"rows_read,required"`
	JSON     observabilityTelemetryQueryResponseStatisticsJSON `json:"-"`
}

// observabilityTelemetryQueryResponseStatisticsJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseStatistics]
type observabilityTelemetryQueryResponseStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseStatisticsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculation struct {
	Aggregates  []ObservabilityTelemetryQueryResponseCalculationsAggregate `json:"aggregates,required"`
	Calculation string                                                     `json:"calculation,required"`
	Series      []ObservabilityTelemetryQueryResponseCalculationsSery      `json:"series,required"`
	Alias       string                                                     `json:"alias"`
	JSON        observabilityTelemetryQueryResponseCalculationJSON         `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseCalculation]
type observabilityTelemetryQueryResponseCalculationJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsAggregate struct {
	Count          float64                                                          `json:"count,required"`
	Interval       float64                                                          `json:"interval,required"`
	SampleInterval float64                                                          `json:"sampleInterval,required"`
	Value          float64                                                          `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup `json:"groups"`
	JSON           observabilityTelemetryQueryResponseCalculationsAggregateJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsAggregateJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsAggregate]
type observabilityTelemetryQueryResponseCalculationsAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup struct {
	Key   string                                                                    `json:"key,required"`
	Value ObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup]
type observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseCalculationsSery struct {
	Data []ObservabilityTelemetryQueryResponseCalculationsSeriesData `json:"data,required"`
	Time string                                                      `json:"time,required"`
	JSON observabilityTelemetryQueryResponseCalculationsSeryJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsSeryJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseCalculationsSery]
type observabilityTelemetryQueryResponseCalculationsSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsSeriesData struct {
	Count          float64                                                          `json:"count,required"`
	FirstSeen      string                                                           `json:"firstSeen,required"`
	Interval       float64                                                          `json:"interval,required"`
	LastSeen       string                                                           `json:"lastSeen,required"`
	SampleInterval float64                                                          `json:"sampleInterval,required"`
	Value          float64                                                          `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup `json:"groups"`
	JSON           observabilityTelemetryQueryResponseCalculationsSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsSeriesDataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsSeriesData]
type observabilityTelemetryQueryResponseCalculationsSeriesDataJSON struct {
	Count          apijson.Field
	FirstSeen      apijson.Field
	Interval       apijson.Field
	LastSeen       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup struct {
	Key   string                                                                    `json:"key,required"`
	Value ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup]
type observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseCompare struct {
	Aggregates  []ObservabilityTelemetryQueryResponseCompareAggregate `json:"aggregates,required"`
	Calculation string                                                `json:"calculation,required"`
	Series      []ObservabilityTelemetryQueryResponseCompareSery      `json:"series,required"`
	Alias       string                                                `json:"alias"`
	JSON        observabilityTelemetryQueryResponseCompareJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCompareJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseCompare]
type observabilityTelemetryQueryResponseCompareJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareAggregate struct {
	Count          float64                                                     `json:"count,required"`
	Interval       float64                                                     `json:"interval,required"`
	SampleInterval float64                                                     `json:"sampleInterval,required"`
	Value          float64                                                     `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryResponseCompareAggregatesGroup `json:"groups"`
	JSON           observabilityTelemetryQueryResponseCompareAggregateJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseCompareAggregateJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseCompareAggregate]
type observabilityTelemetryQueryResponseCompareAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareAggregatesGroup struct {
	Key   string                                                               `json:"key,required"`
	Value ObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryResponseCompareAggregatesGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCompareAggregatesGroupJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCompareAggregatesGroup]
type observabilityTelemetryQueryResponseCompareAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseCompareSery struct {
	Data []ObservabilityTelemetryQueryResponseCompareSeriesData `json:"data,required"`
	Time string                                                 `json:"time,required"`
	JSON observabilityTelemetryQueryResponseCompareSeryJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseCompareSeryJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseCompareSery]
type observabilityTelemetryQueryResponseCompareSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareSeriesData struct {
	Count          float64                                                     `json:"count,required"`
	FirstSeen      string                                                      `json:"firstSeen,required"`
	Interval       float64                                                     `json:"interval,required"`
	LastSeen       string                                                      `json:"lastSeen,required"`
	SampleInterval float64                                                     `json:"sampleInterval,required"`
	Value          float64                                                     `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryResponseCompareSeriesDataGroup `json:"groups"`
	JSON           observabilityTelemetryQueryResponseCompareSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseCompareSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseCompareSeriesData]
type observabilityTelemetryQueryResponseCompareSeriesDataJSON struct {
	Count          apijson.Field
	FirstSeen      apijson.Field
	Interval       apijson.Field
	LastSeen       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareSeriesDataGroup struct {
	Key   string                                                               `json:"key,required"`
	Value ObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCompareSeriesDataGroup]
type observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseEvents struct {
	Count  float64                                          `json:"count"`
	Events []ObservabilityTelemetryQueryResponseEventsEvent `json:"events"`
	Fields []ObservabilityTelemetryQueryResponseEventsField `json:"fields"`
	Series []ObservabilityTelemetryQueryResponseEventsSery  `json:"series"`
	JSON   observabilityTelemetryQueryResponseEventsJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseEventsJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponseEvents]
type observabilityTelemetryQueryResponseEventsJSON struct {
	Count       apijson.Field
	Events      apijson.Field
	Fields      apijson.Field
	Series      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEvents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsJSON) RawJSON() string {
	return r.raw
}

// The data structure of a telemetry event
type ObservabilityTelemetryQueryResponseEventsEvent struct {
	Metadata  ObservabilityTelemetryQueryResponseEventsEventsMetadata `json:"$metadata,required"`
	Dataset   string                                                  `json:"dataset,required"`
	Source    interface{}                                             `json:"source,required"`
	Timestamp int64                                                   `json:"timestamp,required"`
	// Cloudflare Workers event information enriches your logs so you can easily
	// identify and debug issues.
	Workers ObservabilityTelemetryQueryResponseEventsEventsWorkers `json:"$workers"`
	JSON    observabilityTelemetryQueryResponseEventsEventJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseEventsEvent]
type observabilityTelemetryQueryResponseEventsEventJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsEventsMetadata struct {
	ID              string                                                      `json:"id,required"`
	Account         string                                                      `json:"account"`
	CloudService    string                                                      `json:"cloudService"`
	ColdStart       int64                                                       `json:"coldStart"`
	Cost            int64                                                       `json:"cost"`
	Duration        int64                                                       `json:"duration"`
	EndTime         int64                                                       `json:"endTime"`
	Error           string                                                      `json:"error"`
	ErrorTemplate   string                                                      `json:"errorTemplate"`
	Fingerprint     string                                                      `json:"fingerprint"`
	Level           string                                                      `json:"level"`
	Message         string                                                      `json:"message"`
	MessageTemplate string                                                      `json:"messageTemplate"`
	MetricName      string                                                      `json:"metricName"`
	Origin          string                                                      `json:"origin"`
	ParentSpanID    string                                                      `json:"parentSpanId"`
	Provider        string                                                      `json:"provider"`
	Region          string                                                      `json:"region"`
	RequestID       string                                                      `json:"requestId"`
	Service         string                                                      `json:"service"`
	SpanID          string                                                      `json:"spanId"`
	SpanName        string                                                      `json:"spanName"`
	StackID         string                                                      `json:"stackId"`
	StartTime       int64                                                       `json:"startTime"`
	StatusCode      int64                                                       `json:"statusCode"`
	TraceDuration   int64                                                       `json:"traceDuration"`
	TraceID         string                                                      `json:"traceId"`
	Trigger         string                                                      `json:"trigger"`
	Type            string                                                      `json:"type"`
	URL             string                                                      `json:"url"`
	JSON            observabilityTelemetryQueryResponseEventsEventsMetadataJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventsMetadataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsEventsMetadata]
type observabilityTelemetryQueryResponseEventsEventsMetadataJSON struct {
	ID              apijson.Field
	Account         apijson.Field
	CloudService    apijson.Field
	ColdStart       apijson.Field
	Cost            apijson.Field
	Duration        apijson.Field
	EndTime         apijson.Field
	Error           apijson.Field
	ErrorTemplate   apijson.Field
	Fingerprint     apijson.Field
	Level           apijson.Field
	Message         apijson.Field
	MessageTemplate apijson.Field
	MetricName      apijson.Field
	Origin          apijson.Field
	ParentSpanID    apijson.Field
	Provider        apijson.Field
	Region          apijson.Field
	RequestID       apijson.Field
	Service         apijson.Field
	SpanID          apijson.Field
	SpanName        apijson.Field
	StackID         apijson.Field
	StartTime       apijson.Field
	StatusCode      apijson.Field
	TraceDuration   apijson.Field
	TraceID         apijson.Field
	Trigger         apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsEventsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventsMetadataJSON) RawJSON() string {
	return r.raw
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
type ObservabilityTelemetryQueryResponseEventsEventsWorkers struct {
	EventType  ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType `json:"eventType,required"`
	Outcome    string                                                          `json:"outcome,required"`
	RequestID  string                                                          `json:"requestId,required"`
	ScriptName string                                                          `json:"scriptName,required"`
	CPUTimeMs  float64                                                         `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of
	// [map[string]ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion].
	Event          interface{}                                                          `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel `json:"executionModel"`
	// This field can have the runtime type of
	// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                                `json:"scriptVersion"`
	Truncated     bool                                                       `json:"truncated"`
	WallTimeMs    float64                                                    `json:"wallTimeMs"`
	JSON          observabilityTelemetryQueryResponseEventsEventsWorkersJSON `json:"-"`
	union         ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion
}

// observabilityTelemetryQueryResponseEventsEventsWorkersJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseEventsEventsWorkers]
type observabilityTelemetryQueryResponseEventsEventsWorkersJSON struct {
	EventType                apijson.Field
	Outcome                  apijson.Field
	RequestID                apijson.Field
	ScriptName               apijson.Field
	CPUTimeMs                apijson.Field
	DiagnosticsChannelEvents apijson.Field
	DispatchNamespace        apijson.Field
	Entrypoint               apijson.Field
	Event                    apijson.Field
	ExecutionModel           apijson.Field
	ScriptVersion            apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilityTelemetryQueryResponseEventsEventsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityTelemetryQueryResponseEventsEventsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityTelemetryQueryResponseEventsEventsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject],
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject].
func (r ObservabilityTelemetryQueryResponseEventsEventsWorkers) AsUnion() ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
//
// Union satisfied by
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject] or
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject].
type ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion interface {
	implementsObservabilityTelemetryQueryResponseEventsEventsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseEventsEventsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseEventsEventsWorkersObject{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObject struct {
	EventType      ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType             `json:"eventType,required"`
	Outcome        string                                                                            `json:"outcome,required"`
	RequestID      string                                                                            `json:"requestId,required"`
	ScriptName     string                                                                            `json:"scriptName,required"`
	Entrypoint     string                                                                            `json:"entrypoint"`
	Event          map[string]ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel        `json:"executionModel"`
	ScriptVersion  ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion         `json:"scriptVersion"`
	Truncated      bool                                                                              `json:"truncated"`
	JSON           observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON                  `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject]
type observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON struct {
	EventType      apijson.Field
	Outcome        apijson.Field
	RequestID      apijson.Field
	ScriptName     apijson.Field
	Entrypoint     apijson.Field
	Event          apijson.Field
	ExecutionModel apijson.Field
	ScriptVersion  apijson.Field
	Truncated      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsEventsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObject) implementsObservabilityTelemetryQueryResponseEventsEventsWorkers() {
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeFetch     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "fetch"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeScheduled ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "scheduled"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeAlarm     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "alarm"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeCron      ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "cron"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeQueue     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "queue"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeEmail     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "email"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeTail      ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "tail"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeRpc       ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "rpc"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeWebsocket ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "websocket"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeUnknown   ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeFetch, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeScheduled, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeAlarm, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeCron, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeQueue, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeEmail, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeTail, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeRpc, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeWebsocket, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or [ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMap].
type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion interface {
	ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMap{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMap map[string]ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMap) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventUnion() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or [ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMap].
type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMap{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMap map[string]ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMap) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapUnionItem() {
}

// Union satisfied by
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArray],
// [shared.UnionString], [shared.UnionFloat] or [shared.UnionBool].
type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArray []ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArray) ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapUnionItem() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem interface {
	ImplementsObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelDurableObject ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelStateless     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelDurableObject, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion struct {
	ID      string                                                                        `json:"id"`
	Message string                                                                        `json:"message"`
	Tag     string                                                                        `json:"tag"`
	JSON    observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion]
type observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeFetch     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "fetch"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeScheduled ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "scheduled"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeAlarm     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "alarm"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeCron      ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "cron"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeQueue     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "queue"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeEmail     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "email"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeTail      ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "tail"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeRpc       ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "rpc"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeWebsocket ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "websocket"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeUnknown   ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeFetch, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeScheduled, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeAlarm, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeCron, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeQueue, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeEmail, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeTail, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeRpc, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeWebsocket, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelDurableObject ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelStateless     ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelDurableObject, ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEventsField struct {
	Key  string                                             `json:"key,required"`
	Type string                                             `json:"type,required"`
	JSON observabilityTelemetryQueryResponseEventsFieldJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsFieldJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseEventsField]
type observabilityTelemetryQueryResponseEventsFieldJSON struct {
	Key         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsFieldJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsSery struct {
	Data []ObservabilityTelemetryQueryResponseEventsSeriesData `json:"data,required"`
	Time string                                                `json:"time,required"`
	JSON observabilityTelemetryQueryResponseEventsSeryJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseEventsSeryJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseEventsSery]
type observabilityTelemetryQueryResponseEventsSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsSeriesData struct {
	Aggregates     ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates `json:"aggregates,required"`
	Count          float64                                                       `json:"count,required"`
	Interval       float64                                                       `json:"interval,required"`
	SampleInterval float64                                                       `json:"sampleInterval,required"`
	Errors         float64                                                       `json:"errors"`
	// Groups in the query results.
	Groups map[string]ObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion `json:"groups"`
	JSON   observabilityTelemetryQueryResponseEventsSeriesDataJSON                   `json:"-"`
}

// observabilityTelemetryQueryResponseEventsSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseEventsSeriesData]
type observabilityTelemetryQueryResponseEventsSeriesDataJSON struct {
	Aggregates     apijson.Field
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Errors         apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates struct {
	// Deprecated: deprecated
	Count int64 `json:"_count,required"`
	// Deprecated: deprecated
	FirstSeen string `json:"_firstSeen,required"`
	// Deprecated: deprecated
	Interval int64 `json:"_interval,required"`
	// Deprecated: deprecated
	LastSeen string `json:"_lastSeen,required"`
	// Deprecated: deprecated
	Bin  interface{}                                                       `json:"bin"`
	JSON observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates]
type observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON struct {
	Count       apijson.Field
	FirstSeen   apijson.Field
	Interval    apijson.Field
	LastSeen    apijson.Field
	Bin         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion interface {
	ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// The data structure of a telemetry event
type ObservabilityTelemetryQueryResponseInvocation struct {
	Metadata  ObservabilityTelemetryQueryResponseInvocationsMetadata `json:"$metadata,required"`
	Dataset   string                                                 `json:"dataset,required"`
	Source    interface{}                                            `json:"source,required"`
	Timestamp int64                                                  `json:"timestamp,required"`
	// Cloudflare Workers event information enriches your logs so you can easily
	// identify and debug issues.
	Workers ObservabilityTelemetryQueryResponseInvocationsWorkers `json:"$workers"`
	JSON    observabilityTelemetryQueryResponseInvocationJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseInvocation]
type observabilityTelemetryQueryResponseInvocationJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseInvocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseInvocationsMetadata struct {
	ID              string                                                     `json:"id,required"`
	Account         string                                                     `json:"account"`
	CloudService    string                                                     `json:"cloudService"`
	ColdStart       int64                                                      `json:"coldStart"`
	Cost            int64                                                      `json:"cost"`
	Duration        int64                                                      `json:"duration"`
	EndTime         int64                                                      `json:"endTime"`
	Error           string                                                     `json:"error"`
	ErrorTemplate   string                                                     `json:"errorTemplate"`
	Fingerprint     string                                                     `json:"fingerprint"`
	Level           string                                                     `json:"level"`
	Message         string                                                     `json:"message"`
	MessageTemplate string                                                     `json:"messageTemplate"`
	MetricName      string                                                     `json:"metricName"`
	Origin          string                                                     `json:"origin"`
	ParentSpanID    string                                                     `json:"parentSpanId"`
	Provider        string                                                     `json:"provider"`
	Region          string                                                     `json:"region"`
	RequestID       string                                                     `json:"requestId"`
	Service         string                                                     `json:"service"`
	SpanID          string                                                     `json:"spanId"`
	SpanName        string                                                     `json:"spanName"`
	StackID         string                                                     `json:"stackId"`
	StartTime       int64                                                      `json:"startTime"`
	StatusCode      int64                                                      `json:"statusCode"`
	TraceDuration   int64                                                      `json:"traceDuration"`
	TraceID         string                                                     `json:"traceId"`
	Trigger         string                                                     `json:"trigger"`
	Type            string                                                     `json:"type"`
	URL             string                                                     `json:"url"`
	JSON            observabilityTelemetryQueryResponseInvocationsMetadataJSON `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationsMetadataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseInvocationsMetadata]
type observabilityTelemetryQueryResponseInvocationsMetadataJSON struct {
	ID              apijson.Field
	Account         apijson.Field
	CloudService    apijson.Field
	ColdStart       apijson.Field
	Cost            apijson.Field
	Duration        apijson.Field
	EndTime         apijson.Field
	Error           apijson.Field
	ErrorTemplate   apijson.Field
	Fingerprint     apijson.Field
	Level           apijson.Field
	Message         apijson.Field
	MessageTemplate apijson.Field
	MetricName      apijson.Field
	Origin          apijson.Field
	ParentSpanID    apijson.Field
	Provider        apijson.Field
	Region          apijson.Field
	RequestID       apijson.Field
	Service         apijson.Field
	SpanID          apijson.Field
	SpanName        apijson.Field
	StackID         apijson.Field
	StartTime       apijson.Field
	StatusCode      apijson.Field
	TraceDuration   apijson.Field
	TraceID         apijson.Field
	Trigger         apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseInvocationsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationsMetadataJSON) RawJSON() string {
	return r.raw
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
type ObservabilityTelemetryQueryResponseInvocationsWorkers struct {
	EventType  ObservabilityTelemetryQueryResponseInvocationsWorkersEventType `json:"eventType,required"`
	Outcome    string                                                         `json:"outcome,required"`
	RequestID  string                                                         `json:"requestId,required"`
	ScriptName string                                                         `json:"scriptName,required"`
	CPUTimeMs  float64                                                        `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilityTelemetryQueryResponseInvocationsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of
	// [map[string]ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion].
	Event          interface{}                                                         `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel `json:"executionModel"`
	// This field can have the runtime type of
	// [ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                               `json:"scriptVersion"`
	Truncated     bool                                                      `json:"truncated"`
	WallTimeMs    float64                                                   `json:"wallTimeMs"`
	JSON          observabilityTelemetryQueryResponseInvocationsWorkersJSON `json:"-"`
	union         ObservabilityTelemetryQueryResponseInvocationsWorkersUnion
}

// observabilityTelemetryQueryResponseInvocationsWorkersJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseInvocationsWorkers]
type observabilityTelemetryQueryResponseInvocationsWorkersJSON struct {
	EventType                apijson.Field
	Outcome                  apijson.Field
	RequestID                apijson.Field
	ScriptName               apijson.Field
	CPUTimeMs                apijson.Field
	DiagnosticsChannelEvents apijson.Field
	DispatchNamespace        apijson.Field
	Entrypoint               apijson.Field
	Event                    apijson.Field
	ExecutionModel           apijson.Field
	ScriptVersion            apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilityTelemetryQueryResponseInvocationsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityTelemetryQueryResponseInvocationsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityTelemetryQueryResponseInvocationsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityTelemetryQueryResponseInvocationsWorkersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObject],
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObject].
func (r ObservabilityTelemetryQueryResponseInvocationsWorkers) AsUnion() ObservabilityTelemetryQueryResponseInvocationsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
//
// Union satisfied by [ObservabilityTelemetryQueryResponseInvocationsWorkersObject]
// or [ObservabilityTelemetryQueryResponseInvocationsWorkersObject].
type ObservabilityTelemetryQueryResponseInvocationsWorkersUnion interface {
	implementsObservabilityTelemetryQueryResponseInvocationsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseInvocationsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseInvocationsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseInvocationsWorkersObject{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObject struct {
	EventType      ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType             `json:"eventType,required"`
	Outcome        string                                                                           `json:"outcome,required"`
	RequestID      string                                                                           `json:"requestId,required"`
	ScriptName     string                                                                           `json:"scriptName,required"`
	Entrypoint     string                                                                           `json:"entrypoint"`
	Event          map[string]ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel        `json:"executionModel"`
	ScriptVersion  ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion         `json:"scriptVersion"`
	Truncated      bool                                                                             `json:"truncated"`
	JSON           observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON                  `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObject]
type observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON struct {
	EventType      apijson.Field
	Outcome        apijson.Field
	RequestID      apijson.Field
	ScriptName     apijson.Field
	Entrypoint     apijson.Field
	Event          apijson.Field
	ExecutionModel apijson.Field
	ScriptVersion  apijson.Field
	Truncated      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseInvocationsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObject) implementsObservabilityTelemetryQueryResponseInvocationsWorkers() {
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeFetch     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "fetch"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeScheduled ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "scheduled"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeAlarm     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "alarm"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeCron      ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "cron"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeQueue     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "queue"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeEmail     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "email"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeTail      ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "tail"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeRpc       ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "rpc"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeWebsocket ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "websocket"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeUnknown   ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeFetch, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeScheduled, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeAlarm, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeCron, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeQueue, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeEmail, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeTail, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeRpc, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeWebsocket, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or [ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMap].
type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion interface {
	ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMap{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMap map[string]ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMap) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventUnion() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or [ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMap].
type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMap{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMap map[string]ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMap) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapUnionItem() {
}

// Union satisfied by
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArray],
// [shared.UnionString], [shared.UnionFloat] or [shared.UnionBool].
type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArray []ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArrayUnionItem

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArray) ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapUnionItem() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArrayUnionItem interface {
	ImplementsObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArrayUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventMapMapArrayUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelDurableObject ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelStateless     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelDurableObject, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion struct {
	ID      string                                                                       `json:"id"`
	Message string                                                                       `json:"message"`
	Tag     string                                                                       `json:"tag"`
	JSON    observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion]
type observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersEventType string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeFetch     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "fetch"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeScheduled ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "scheduled"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeAlarm     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "alarm"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeCron      ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "cron"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeQueue     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "queue"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeEmail     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "email"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeTail      ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "tail"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeRpc       ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "rpc"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeWebsocket ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "websocket"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeUnknown   ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeFetch, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeScheduled, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeAlarm, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeCron, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeQueue, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeEmail, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeTail, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeRpc, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeWebsocket, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelDurableObject ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelStateless     ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelDurableObject, ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponsePattern struct {
	Count   float64                                           `json:"count,required"`
	Pattern string                                            `json:"pattern,required"`
	Series  []ObservabilityTelemetryQueryResponsePatternsSery `json:"series,required"`
	Service string                                            `json:"service,required"`
	JSON    observabilityTelemetryQueryResponsePatternJSON    `json:"-"`
}

// observabilityTelemetryQueryResponsePatternJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponsePattern]
type observabilityTelemetryQueryResponsePatternJSON struct {
	Count       apijson.Field
	Pattern     apijson.Field
	Series      apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponsePattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponsePatternJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponsePatternsSery struct {
	Data ObservabilityTelemetryQueryResponsePatternsSeriesData `json:"data,required"`
	Time string                                                `json:"time,required"`
	JSON observabilityTelemetryQueryResponsePatternsSeryJSON   `json:"-"`
}

// observabilityTelemetryQueryResponsePatternsSeryJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponsePatternsSery]
type observabilityTelemetryQueryResponsePatternsSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponsePatternsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponsePatternsSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponsePatternsSeriesData struct {
	Count          float64                                                      `json:"count,required"`
	Interval       float64                                                      `json:"interval,required"`
	SampleInterval float64                                                      `json:"sampleInterval,required"`
	Value          float64                                                      `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryResponsePatternsSeriesDataGroup `json:"groups"`
	JSON           observabilityTelemetryQueryResponsePatternsSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryResponsePatternsSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponsePatternsSeriesData]
type observabilityTelemetryQueryResponsePatternsSeriesDataJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponsePatternsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponsePatternsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponsePatternsSeriesDataGroup struct {
	Key   string                                                                `json:"key,required"`
	Value ObservabilityTelemetryQueryResponsePatternsSeriesDataGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryResponsePatternsSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponsePatternsSeriesDataGroupJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponsePatternsSeriesDataGroup]
type observabilityTelemetryQueryResponsePatternsSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponsePatternsSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponsePatternsSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponsePatternsSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponsePatternsSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponsePatternsSeriesDataGroupsValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryValuesResponse struct {
	Dataset string                                         `json:"dataset,required"`
	Key     string                                         `json:"key,required"`
	Type    ObservabilityTelemetryValuesResponseType       `json:"type,required"`
	Value   ObservabilityTelemetryValuesResponseValueUnion `json:"value,required"`
	JSON    observabilityTelemetryValuesResponseJSON       `json:"-"`
}

// observabilityTelemetryValuesResponseJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryValuesResponse]
type observabilityTelemetryValuesResponseJSON struct {
	Dataset     apijson.Field
	Key         apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryValuesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryValuesResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryValuesResponseType string

const (
	ObservabilityTelemetryValuesResponseTypeString  ObservabilityTelemetryValuesResponseType = "string"
	ObservabilityTelemetryValuesResponseTypeBoolean ObservabilityTelemetryValuesResponseType = "boolean"
	ObservabilityTelemetryValuesResponseTypeNumber  ObservabilityTelemetryValuesResponseType = "number"
)

func (r ObservabilityTelemetryValuesResponseType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesResponseTypeString, ObservabilityTelemetryValuesResponseTypeBoolean, ObservabilityTelemetryValuesResponseTypeNumber:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryValuesResponseValueUnion interface {
	ImplementsObservabilityTelemetryValuesResponseValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryValuesResponseValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityTelemetryKeysParams struct {
	AccountID param.Field[string]                                   `path:"account_id,required"`
	Datasets  param.Field[[]string]                                 `json:"datasets"`
	Filters   param.Field[[]ObservabilityTelemetryKeysParamsFilter] `json:"filters"`
	// Search for a specific substring in the keys.
	KeyNeedle param.Field[ObservabilityTelemetryKeysParamsKeyNeedle] `json:"keyNeedle"`
	Limit     param.Field[float64]                                   `json:"limit"`
	// Search for a specific substring in the event.
	Needle    param.Field[ObservabilityTelemetryKeysParamsNeedle]    `json:"needle"`
	Timeframe param.Field[ObservabilityTelemetryKeysParamsTimeframe] `json:"timeframe"`
}

func (r ObservabilityTelemetryKeysParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryKeysParamsFilter struct {
	Key       param.Field[string]                                            `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryKeysParamsFiltersOperation]  `json:"operation,required"`
	Type      param.Field[ObservabilityTelemetryKeysParamsFiltersType]       `json:"type,required"`
	Value     param.Field[ObservabilityTelemetryKeysParamsFiltersValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryKeysParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryKeysParamsFiltersOperation string

const (
	ObservabilityTelemetryKeysParamsFiltersOperationIncludes            ObservabilityTelemetryKeysParamsFiltersOperation = "includes"
	ObservabilityTelemetryKeysParamsFiltersOperationNotIncludes         ObservabilityTelemetryKeysParamsFiltersOperation = "not_includes"
	ObservabilityTelemetryKeysParamsFiltersOperationStartsWith          ObservabilityTelemetryKeysParamsFiltersOperation = "starts_with"
	ObservabilityTelemetryKeysParamsFiltersOperationRegex               ObservabilityTelemetryKeysParamsFiltersOperation = "regex"
	ObservabilityTelemetryKeysParamsFiltersOperationExists              ObservabilityTelemetryKeysParamsFiltersOperation = "exists"
	ObservabilityTelemetryKeysParamsFiltersOperationIsNull              ObservabilityTelemetryKeysParamsFiltersOperation = "is_null"
	ObservabilityTelemetryKeysParamsFiltersOperationIn                  ObservabilityTelemetryKeysParamsFiltersOperation = "in"
	ObservabilityTelemetryKeysParamsFiltersOperationNotIn               ObservabilityTelemetryKeysParamsFiltersOperation = "not_in"
	ObservabilityTelemetryKeysParamsFiltersOperationEq                  ObservabilityTelemetryKeysParamsFiltersOperation = "eq"
	ObservabilityTelemetryKeysParamsFiltersOperationNeq                 ObservabilityTelemetryKeysParamsFiltersOperation = "neq"
	ObservabilityTelemetryKeysParamsFiltersOperationGt                  ObservabilityTelemetryKeysParamsFiltersOperation = "gt"
	ObservabilityTelemetryKeysParamsFiltersOperationGte                 ObservabilityTelemetryKeysParamsFiltersOperation = "gte"
	ObservabilityTelemetryKeysParamsFiltersOperationLt                  ObservabilityTelemetryKeysParamsFiltersOperation = "lt"
	ObservabilityTelemetryKeysParamsFiltersOperationLte                 ObservabilityTelemetryKeysParamsFiltersOperation = "lte"
	ObservabilityTelemetryKeysParamsFiltersOperationEquals              ObservabilityTelemetryKeysParamsFiltersOperation = "="
	ObservabilityTelemetryKeysParamsFiltersOperationNotEquals           ObservabilityTelemetryKeysParamsFiltersOperation = "!="
	ObservabilityTelemetryKeysParamsFiltersOperationGreater             ObservabilityTelemetryKeysParamsFiltersOperation = ">"
	ObservabilityTelemetryKeysParamsFiltersOperationGreaterOrEquals     ObservabilityTelemetryKeysParamsFiltersOperation = ">="
	ObservabilityTelemetryKeysParamsFiltersOperationLess                ObservabilityTelemetryKeysParamsFiltersOperation = "<"
	ObservabilityTelemetryKeysParamsFiltersOperationLessOrEquals        ObservabilityTelemetryKeysParamsFiltersOperation = "<="
	ObservabilityTelemetryKeysParamsFiltersOperationIncludesUppercase   ObservabilityTelemetryKeysParamsFiltersOperation = "INCLUDES"
	ObservabilityTelemetryKeysParamsFiltersOperationDoesNotInclude      ObservabilityTelemetryKeysParamsFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryKeysParamsFiltersOperationMatchRegex          ObservabilityTelemetryKeysParamsFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryKeysParamsFiltersOperationExistsUppercase     ObservabilityTelemetryKeysParamsFiltersOperation = "EXISTS"
	ObservabilityTelemetryKeysParamsFiltersOperationDoesNotExist        ObservabilityTelemetryKeysParamsFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryKeysParamsFiltersOperationInUppercase         ObservabilityTelemetryKeysParamsFiltersOperation = "IN"
	ObservabilityTelemetryKeysParamsFiltersOperationNotInUppercase      ObservabilityTelemetryKeysParamsFiltersOperation = "NOT_IN"
	ObservabilityTelemetryKeysParamsFiltersOperationStartsWithUppercase ObservabilityTelemetryKeysParamsFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryKeysParamsFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersOperationIncludes, ObservabilityTelemetryKeysParamsFiltersOperationNotIncludes, ObservabilityTelemetryKeysParamsFiltersOperationStartsWith, ObservabilityTelemetryKeysParamsFiltersOperationRegex, ObservabilityTelemetryKeysParamsFiltersOperationExists, ObservabilityTelemetryKeysParamsFiltersOperationIsNull, ObservabilityTelemetryKeysParamsFiltersOperationIn, ObservabilityTelemetryKeysParamsFiltersOperationNotIn, ObservabilityTelemetryKeysParamsFiltersOperationEq, ObservabilityTelemetryKeysParamsFiltersOperationNeq, ObservabilityTelemetryKeysParamsFiltersOperationGt, ObservabilityTelemetryKeysParamsFiltersOperationGte, ObservabilityTelemetryKeysParamsFiltersOperationLt, ObservabilityTelemetryKeysParamsFiltersOperationLte, ObservabilityTelemetryKeysParamsFiltersOperationEquals, ObservabilityTelemetryKeysParamsFiltersOperationNotEquals, ObservabilityTelemetryKeysParamsFiltersOperationGreater, ObservabilityTelemetryKeysParamsFiltersOperationGreaterOrEquals, ObservabilityTelemetryKeysParamsFiltersOperationLess, ObservabilityTelemetryKeysParamsFiltersOperationLessOrEquals, ObservabilityTelemetryKeysParamsFiltersOperationIncludesUppercase, ObservabilityTelemetryKeysParamsFiltersOperationDoesNotInclude, ObservabilityTelemetryKeysParamsFiltersOperationMatchRegex, ObservabilityTelemetryKeysParamsFiltersOperationExistsUppercase, ObservabilityTelemetryKeysParamsFiltersOperationDoesNotExist, ObservabilityTelemetryKeysParamsFiltersOperationInUppercase, ObservabilityTelemetryKeysParamsFiltersOperationNotInUppercase, ObservabilityTelemetryKeysParamsFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryKeysParamsFiltersType string

const (
	ObservabilityTelemetryKeysParamsFiltersTypeString  ObservabilityTelemetryKeysParamsFiltersType = "string"
	ObservabilityTelemetryKeysParamsFiltersTypeNumber  ObservabilityTelemetryKeysParamsFiltersType = "number"
	ObservabilityTelemetryKeysParamsFiltersTypeBoolean ObservabilityTelemetryKeysParamsFiltersType = "boolean"
)

func (r ObservabilityTelemetryKeysParamsFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersTypeString, ObservabilityTelemetryKeysParamsFiltersTypeNumber, ObservabilityTelemetryKeysParamsFiltersTypeBoolean:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeysParamsFiltersValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsFiltersValueUnion()
}

// Search for a specific substring in the keys.
type ObservabilityTelemetryKeysParamsKeyNeedle struct {
	Value     param.Field[ObservabilityTelemetryKeysParamsKeyNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                                `json:"isRegex"`
	MatchCase param.Field[bool]                                                `json:"matchCase"`
}

func (r ObservabilityTelemetryKeysParamsKeyNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeysParamsKeyNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion()
}

// Search for a specific substring in the event.
type ObservabilityTelemetryKeysParamsNeedle struct {
	Value     param.Field[ObservabilityTelemetryKeysParamsNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                             `json:"isRegex"`
	MatchCase param.Field[bool]                                             `json:"matchCase"`
}

func (r ObservabilityTelemetryKeysParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeysParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()
}

type ObservabilityTelemetryKeysParamsTimeframe struct {
	From param.Field[float64] `json:"from,required"`
	To   param.Field[float64] `json:"to,required"`
}

func (r ObservabilityTelemetryKeysParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParams struct {
	AccountID       param.Field[string]                                       `path:"account_id,required"`
	QueryID         param.Field[string]                                       `json:"queryId,required"`
	Timeframe       param.Field[ObservabilityTelemetryQueryParamsTimeframe]   `json:"timeframe,required"`
	Chart           param.Field[bool]                                         `json:"chart"`
	Compare         param.Field[bool]                                         `json:"compare"`
	Dry             param.Field[bool]                                         `json:"dry"`
	Granularity     param.Field[float64]                                      `json:"granularity"`
	IgnoreSeries    param.Field[bool]                                         `json:"ignoreSeries"`
	Limit           param.Field[float64]                                      `json:"limit"`
	Offset          param.Field[string]                                       `json:"offset"`
	OffsetBy        param.Field[float64]                                      `json:"offsetBy"`
	OffsetDirection param.Field[string]                                       `json:"offsetDirection"`
	Parameters      param.Field[ObservabilityTelemetryQueryParamsParameters]  `json:"parameters"`
	PatternType     param.Field[ObservabilityTelemetryQueryParamsPatternType] `json:"patternType"`
	View            param.Field[ObservabilityTelemetryQueryParamsView]        `json:"view"`
}

func (r ObservabilityTelemetryQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsTimeframe struct {
	From param.Field[float64] `json:"from,required"`
	To   param.Field[float64] `json:"to,required"`
}

func (r ObservabilityTelemetryQueryParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations param.Field[[]ObservabilityTelemetryQueryParamsParametersCalculation] `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets param.Field[[]string] `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFilterCombination] `json:"filterCombination"`
	// Configure the Filters to apply to the query.
	Filters param.Field[[]ObservabilityTelemetryQueryParamsParametersFilter] `json:"filters"`
	// Define how to group the results of the query.
	GroupBys param.Field[[]ObservabilityTelemetryQueryParamsParametersGroupBy] `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings param.Field[[]ObservabilityTelemetryQueryParamsParametersHaving] `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit param.Field[int64] `json:"limit"`
	// Define an expression to search using full-text search.
	Needle param.Field[ObservabilityTelemetryQueryParamsParametersNeedle] `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy param.Field[ObservabilityTelemetryQueryParamsParametersOrderBy] `json:"orderBy"`
}

func (r ObservabilityTelemetryQueryParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParametersCalculation struct {
	Operator param.Field[ObservabilityTelemetryQueryParamsParametersCalculationsOperator] `json:"operator,required"`
	Alias    param.Field[string]                                                          `json:"alias"`
	Key      param.Field[string]                                                          `json:"key"`
	KeyType  param.Field[ObservabilityTelemetryQueryParamsParametersCalculationsKeyType]  `json:"keyType"`
}

func (r ObservabilityTelemetryQueryParamsParametersCalculation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParametersCalculationsOperator string

const (
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorUniq              ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "uniq"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCount             ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "count"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMax               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "max"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMin               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "min"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSum               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "sum"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvg               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "avg"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedian            ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "median"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001              ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p001"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p01"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p05"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p10"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p25"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p75"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p90"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p95"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p99"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999              ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p999"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddev            ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "stddev"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVariance          ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "variance"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountDistinct     ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountUppercase    ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "COUNT"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMaxUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "MAX"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMinUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "MIN"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSumUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "SUM"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvgUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "AVG"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedianUppercase   ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "MEDIAN"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001Uppercase     ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P001"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P01"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P05"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P10"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P25"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P75"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P90"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P95"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P99"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999Uppercase     ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P999"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddevUppercase   ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "STDDEV"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVarianceUppercase ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityTelemetryQueryParamsParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersCalculationsOperatorUniq, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCount, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMax, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMin, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSum, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvg, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedian, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddev, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVariance, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountDistinct, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMaxUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMinUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSumUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvgUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedianUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddevUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersCalculationsKeyType string

const (
	ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeString  ObservabilityTelemetryQueryParamsParametersCalculationsKeyType = "string"
	ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeNumber  ObservabilityTelemetryQueryParamsParametersCalculationsKeyType = "number"
	ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeBoolean ObservabilityTelemetryQueryParamsParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeString, ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeNumber, ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityTelemetryQueryParamsParametersFilterCombination string

const (
	ObservabilityTelemetryQueryParamsParametersFilterCombinationAnd          ObservabilityTelemetryQueryParamsParametersFilterCombination = "and"
	ObservabilityTelemetryQueryParamsParametersFilterCombinationOr           ObservabilityTelemetryQueryParamsParametersFilterCombination = "or"
	ObservabilityTelemetryQueryParamsParametersFilterCombinationAndUppercase ObservabilityTelemetryQueryParamsParametersFilterCombination = "AND"
	ObservabilityTelemetryQueryParamsParametersFilterCombinationOrUppercase  ObservabilityTelemetryQueryParamsParametersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryParamsParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFilterCombinationAnd, ObservabilityTelemetryQueryParamsParametersFilterCombinationOr, ObservabilityTelemetryQueryParamsParametersFilterCombinationAndUppercase, ObservabilityTelemetryQueryParamsParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersFilter struct {
	Key       param.Field[string]                                                       `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersFiltersOperation]  `json:"operation,required"`
	Type      param.Field[ObservabilityTelemetryQueryParamsParametersFiltersType]       `json:"type,required"`
	Value     param.Field[ObservabilityTelemetryQueryParamsParametersFiltersValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryQueryParamsParametersFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParametersFiltersOperation string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludes            ObservabilityTelemetryQueryParamsParametersFiltersOperation = "includes"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIncludes         ObservabilityTelemetryQueryParamsParametersFiltersOperation = "not_includes"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWith          ObservabilityTelemetryQueryParamsParametersFiltersOperation = "starts_with"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationRegex               ObservabilityTelemetryQueryParamsParametersFiltersOperation = "regex"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationExists              ObservabilityTelemetryQueryParamsParametersFiltersOperation = "exists"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIsNull              ObservabilityTelemetryQueryParamsParametersFiltersOperation = "is_null"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIn                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "in"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIn               ObservabilityTelemetryQueryParamsParametersFiltersOperation = "not_in"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationEq                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "eq"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNeq                 ObservabilityTelemetryQueryParamsParametersFiltersOperation = "neq"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGt                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "gt"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGte                 ObservabilityTelemetryQueryParamsParametersFiltersOperation = "gte"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLt                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "lt"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLte                 ObservabilityTelemetryQueryParamsParametersFiltersOperation = "lte"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationEquals              ObservabilityTelemetryQueryParamsParametersFiltersOperation = "="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotEquals           ObservabilityTelemetryQueryParamsParametersFiltersOperation = "!="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGreater             ObservabilityTelemetryQueryParamsParametersFiltersOperation = ">"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGreaterOrEquals     ObservabilityTelemetryQueryParamsParametersFiltersOperation = ">="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLess                ObservabilityTelemetryQueryParamsParametersFiltersOperation = "<"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLessOrEquals        ObservabilityTelemetryQueryParamsParametersFiltersOperation = "<="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludesUppercase   ObservabilityTelemetryQueryParamsParametersFiltersOperation = "INCLUDES"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotInclude      ObservabilityTelemetryQueryParamsParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationMatchRegex          ObservabilityTelemetryQueryParamsParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationExistsUppercase     ObservabilityTelemetryQueryParamsParametersFiltersOperation = "EXISTS"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotExist        ObservabilityTelemetryQueryParamsParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationInUppercase         ObservabilityTelemetryQueryParamsParametersFiltersOperation = "IN"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotInUppercase      ObservabilityTelemetryQueryParamsParametersFiltersOperation = "NOT_IN"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWithUppercase ObservabilityTelemetryQueryParamsParametersFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludes, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIncludes, ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWith, ObservabilityTelemetryQueryParamsParametersFiltersOperationRegex, ObservabilityTelemetryQueryParamsParametersFiltersOperationExists, ObservabilityTelemetryQueryParamsParametersFiltersOperationIsNull, ObservabilityTelemetryQueryParamsParametersFiltersOperationIn, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIn, ObservabilityTelemetryQueryParamsParametersFiltersOperationEq, ObservabilityTelemetryQueryParamsParametersFiltersOperationNeq, ObservabilityTelemetryQueryParamsParametersFiltersOperationGt, ObservabilityTelemetryQueryParamsParametersFiltersOperationGte, ObservabilityTelemetryQueryParamsParametersFiltersOperationLt, ObservabilityTelemetryQueryParamsParametersFiltersOperationLte, ObservabilityTelemetryQueryParamsParametersFiltersOperationEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationGreater, ObservabilityTelemetryQueryParamsParametersFiltersOperationGreaterOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationLess, ObservabilityTelemetryQueryParamsParametersFiltersOperationLessOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludesUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotInclude, ObservabilityTelemetryQueryParamsParametersFiltersOperationMatchRegex, ObservabilityTelemetryQueryParamsParametersFiltersOperationExistsUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotExist, ObservabilityTelemetryQueryParamsParametersFiltersOperationInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersFiltersType string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersTypeString  ObservabilityTelemetryQueryParamsParametersFiltersType = "string"
	ObservabilityTelemetryQueryParamsParametersFiltersTypeNumber  ObservabilityTelemetryQueryParamsParametersFiltersType = "number"
	ObservabilityTelemetryQueryParamsParametersFiltersTypeBoolean ObservabilityTelemetryQueryParamsParametersFiltersType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersTypeString, ObservabilityTelemetryQueryParamsParametersFiltersTypeNumber, ObservabilityTelemetryQueryParamsParametersFiltersTypeBoolean:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryQueryParamsParametersFiltersValueUnion interface {
	ImplementsObservabilityTelemetryQueryParamsParametersFiltersValueUnion()
}

type ObservabilityTelemetryQueryParamsParametersGroupBy struct {
	Type  param.Field[ObservabilityTelemetryQueryParamsParametersGroupBysType] `json:"type,required"`
	Value param.Field[string]                                                  `json:"value,required"`
}

func (r ObservabilityTelemetryQueryParamsParametersGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParametersGroupBysType string

const (
	ObservabilityTelemetryQueryParamsParametersGroupBysTypeString  ObservabilityTelemetryQueryParamsParametersGroupBysType = "string"
	ObservabilityTelemetryQueryParamsParametersGroupBysTypeNumber  ObservabilityTelemetryQueryParamsParametersGroupBysType = "number"
	ObservabilityTelemetryQueryParamsParametersGroupBysTypeBoolean ObservabilityTelemetryQueryParamsParametersGroupBysType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersGroupBysTypeString, ObservabilityTelemetryQueryParamsParametersGroupBysTypeNumber, ObservabilityTelemetryQueryParamsParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersHaving struct {
	Key       param.Field[string]                                                      `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersHavingsOperation] `json:"operation,required"`
	Value     param.Field[float64]                                                     `json:"value,required"`
}

func (r ObservabilityTelemetryQueryParamsParametersHaving) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParametersHavingsOperation string

const (
	ObservabilityTelemetryQueryParamsParametersHavingsOperationEq  ObservabilityTelemetryQueryParamsParametersHavingsOperation = "eq"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationNeq ObservabilityTelemetryQueryParamsParametersHavingsOperation = "neq"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationGt  ObservabilityTelemetryQueryParamsParametersHavingsOperation = "gt"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationGte ObservabilityTelemetryQueryParamsParametersHavingsOperation = "gte"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationLt  ObservabilityTelemetryQueryParamsParametersHavingsOperation = "lt"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationLte ObservabilityTelemetryQueryParamsParametersHavingsOperation = "lte"
)

func (r ObservabilityTelemetryQueryParamsParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersHavingsOperationEq, ObservabilityTelemetryQueryParamsParametersHavingsOperationNeq, ObservabilityTelemetryQueryParamsParametersHavingsOperationGt, ObservabilityTelemetryQueryParamsParametersHavingsOperationGte, ObservabilityTelemetryQueryParamsParametersHavingsOperationLt, ObservabilityTelemetryQueryParamsParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityTelemetryQueryParamsParametersNeedle struct {
	Value     param.Field[ObservabilityTelemetryQueryParamsParametersNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                                        `json:"isRegex"`
	MatchCase param.Field[bool]                                                        `json:"matchCase"`
}

func (r ObservabilityTelemetryQueryParamsParametersNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryQueryParamsParametersNeedleValueUnion interface {
	ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion()
}

// Configure the order of the results returned by the query.
type ObservabilityTelemetryQueryParamsParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value param.Field[string] `json:"value,required"`
	// Set the order of the results
	Order param.Field[ObservabilityTelemetryQueryParamsParametersOrderByOrder] `json:"order"`
}

func (r ObservabilityTelemetryQueryParamsParametersOrderBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set the order of the results
type ObservabilityTelemetryQueryParamsParametersOrderByOrder string

const (
	ObservabilityTelemetryQueryParamsParametersOrderByOrderAsc  ObservabilityTelemetryQueryParamsParametersOrderByOrder = "asc"
	ObservabilityTelemetryQueryParamsParametersOrderByOrderDesc ObservabilityTelemetryQueryParamsParametersOrderByOrder = "desc"
)

func (r ObservabilityTelemetryQueryParamsParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersOrderByOrderAsc, ObservabilityTelemetryQueryParamsParametersOrderByOrderDesc:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsPatternType string

const (
	ObservabilityTelemetryQueryParamsPatternTypeMessage ObservabilityTelemetryQueryParamsPatternType = "message"
	ObservabilityTelemetryQueryParamsPatternTypeError   ObservabilityTelemetryQueryParamsPatternType = "error"
)

func (r ObservabilityTelemetryQueryParamsPatternType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsPatternTypeMessage, ObservabilityTelemetryQueryParamsPatternTypeError:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsView string

const (
	ObservabilityTelemetryQueryParamsViewTraces       ObservabilityTelemetryQueryParamsView = "traces"
	ObservabilityTelemetryQueryParamsViewEvents       ObservabilityTelemetryQueryParamsView = "events"
	ObservabilityTelemetryQueryParamsViewCalculations ObservabilityTelemetryQueryParamsView = "calculations"
	ObservabilityTelemetryQueryParamsViewInvocations  ObservabilityTelemetryQueryParamsView = "invocations"
	ObservabilityTelemetryQueryParamsViewRequests     ObservabilityTelemetryQueryParamsView = "requests"
	ObservabilityTelemetryQueryParamsViewPatterns     ObservabilityTelemetryQueryParamsView = "patterns"
)

func (r ObservabilityTelemetryQueryParamsView) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsViewTraces, ObservabilityTelemetryQueryParamsViewEvents, ObservabilityTelemetryQueryParamsViewCalculations, ObservabilityTelemetryQueryParamsViewInvocations, ObservabilityTelemetryQueryParamsViewRequests, ObservabilityTelemetryQueryParamsViewPatterns:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEnvelope struct {
	Errors   []ObservabilityTelemetryQueryResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ObservabilityTelemetryQueryResponseEnvelopeMessages `json:"messages,required"`
	Result   ObservabilityTelemetryQueryResponse                   `json:"result,required"`
	Success  ObservabilityTelemetryQueryResponseEnvelopeSuccess    `json:"success,required"`
	JSON     observabilityTelemetryQueryResponseEnvelopeJSON       `json:"-"`
}

// observabilityTelemetryQueryResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseEnvelope]
type observabilityTelemetryQueryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEnvelopeErrors struct {
	Message string                                                `json:"message,required"`
	JSON    observabilityTelemetryQueryResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseEnvelopeErrors]
type observabilityTelemetryQueryResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEnvelopeMessages struct {
	Message ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage `json:"message,required"`
	JSON    observabilityTelemetryQueryResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseEnvelopeMessages]
type observabilityTelemetryQueryResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage string

const (
	ObservabilityTelemetryQueryResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEnvelopeSuccess bool

const (
	ObservabilityTelemetryQueryResponseEnvelopeSuccessTrue ObservabilityTelemetryQueryResponseEnvelopeSuccess = true
)

func (r ObservabilityTelemetryQueryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParams struct {
	AccountID param.Field[string]                                      `path:"account_id,required"`
	Datasets  param.Field[[]string]                                    `json:"datasets,required"`
	Key       param.Field[string]                                      `json:"key,required"`
	Timeframe param.Field[ObservabilityTelemetryValuesParamsTimeframe] `json:"timeframe,required"`
	Type      param.Field[ObservabilityTelemetryValuesParamsType]      `json:"type,required"`
	Filters   param.Field[[]ObservabilityTelemetryValuesParamsFilter]  `json:"filters"`
	Limit     param.Field[float64]                                     `json:"limit"`
	// Search for a specific substring in the event.
	Needle param.Field[ObservabilityTelemetryValuesParamsNeedle] `json:"needle"`
}

func (r ObservabilityTelemetryValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValuesParamsTimeframe struct {
	From param.Field[float64] `json:"from,required"`
	To   param.Field[float64] `json:"to,required"`
}

func (r ObservabilityTelemetryValuesParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValuesParamsType string

const (
	ObservabilityTelemetryValuesParamsTypeString  ObservabilityTelemetryValuesParamsType = "string"
	ObservabilityTelemetryValuesParamsTypeBoolean ObservabilityTelemetryValuesParamsType = "boolean"
	ObservabilityTelemetryValuesParamsTypeNumber  ObservabilityTelemetryValuesParamsType = "number"
)

func (r ObservabilityTelemetryValuesParamsType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsTypeString, ObservabilityTelemetryValuesParamsTypeBoolean, ObservabilityTelemetryValuesParamsTypeNumber:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParamsFilter struct {
	Key       param.Field[string]                                              `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryValuesParamsFiltersOperation]  `json:"operation,required"`
	Type      param.Field[ObservabilityTelemetryValuesParamsFiltersType]       `json:"type,required"`
	Value     param.Field[ObservabilityTelemetryValuesParamsFiltersValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryValuesParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValuesParamsFiltersOperation string

const (
	ObservabilityTelemetryValuesParamsFiltersOperationIncludes            ObservabilityTelemetryValuesParamsFiltersOperation = "includes"
	ObservabilityTelemetryValuesParamsFiltersOperationNotIncludes         ObservabilityTelemetryValuesParamsFiltersOperation = "not_includes"
	ObservabilityTelemetryValuesParamsFiltersOperationStartsWith          ObservabilityTelemetryValuesParamsFiltersOperation = "starts_with"
	ObservabilityTelemetryValuesParamsFiltersOperationRegex               ObservabilityTelemetryValuesParamsFiltersOperation = "regex"
	ObservabilityTelemetryValuesParamsFiltersOperationExists              ObservabilityTelemetryValuesParamsFiltersOperation = "exists"
	ObservabilityTelemetryValuesParamsFiltersOperationIsNull              ObservabilityTelemetryValuesParamsFiltersOperation = "is_null"
	ObservabilityTelemetryValuesParamsFiltersOperationIn                  ObservabilityTelemetryValuesParamsFiltersOperation = "in"
	ObservabilityTelemetryValuesParamsFiltersOperationNotIn               ObservabilityTelemetryValuesParamsFiltersOperation = "not_in"
	ObservabilityTelemetryValuesParamsFiltersOperationEq                  ObservabilityTelemetryValuesParamsFiltersOperation = "eq"
	ObservabilityTelemetryValuesParamsFiltersOperationNeq                 ObservabilityTelemetryValuesParamsFiltersOperation = "neq"
	ObservabilityTelemetryValuesParamsFiltersOperationGt                  ObservabilityTelemetryValuesParamsFiltersOperation = "gt"
	ObservabilityTelemetryValuesParamsFiltersOperationGte                 ObservabilityTelemetryValuesParamsFiltersOperation = "gte"
	ObservabilityTelemetryValuesParamsFiltersOperationLt                  ObservabilityTelemetryValuesParamsFiltersOperation = "lt"
	ObservabilityTelemetryValuesParamsFiltersOperationLte                 ObservabilityTelemetryValuesParamsFiltersOperation = "lte"
	ObservabilityTelemetryValuesParamsFiltersOperationEquals              ObservabilityTelemetryValuesParamsFiltersOperation = "="
	ObservabilityTelemetryValuesParamsFiltersOperationNotEquals           ObservabilityTelemetryValuesParamsFiltersOperation = "!="
	ObservabilityTelemetryValuesParamsFiltersOperationGreater             ObservabilityTelemetryValuesParamsFiltersOperation = ">"
	ObservabilityTelemetryValuesParamsFiltersOperationGreaterOrEquals     ObservabilityTelemetryValuesParamsFiltersOperation = ">="
	ObservabilityTelemetryValuesParamsFiltersOperationLess                ObservabilityTelemetryValuesParamsFiltersOperation = "<"
	ObservabilityTelemetryValuesParamsFiltersOperationLessOrEquals        ObservabilityTelemetryValuesParamsFiltersOperation = "<="
	ObservabilityTelemetryValuesParamsFiltersOperationIncludesUppercase   ObservabilityTelemetryValuesParamsFiltersOperation = "INCLUDES"
	ObservabilityTelemetryValuesParamsFiltersOperationDoesNotInclude      ObservabilityTelemetryValuesParamsFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryValuesParamsFiltersOperationMatchRegex          ObservabilityTelemetryValuesParamsFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryValuesParamsFiltersOperationExistsUppercase     ObservabilityTelemetryValuesParamsFiltersOperation = "EXISTS"
	ObservabilityTelemetryValuesParamsFiltersOperationDoesNotExist        ObservabilityTelemetryValuesParamsFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryValuesParamsFiltersOperationInUppercase         ObservabilityTelemetryValuesParamsFiltersOperation = "IN"
	ObservabilityTelemetryValuesParamsFiltersOperationNotInUppercase      ObservabilityTelemetryValuesParamsFiltersOperation = "NOT_IN"
	ObservabilityTelemetryValuesParamsFiltersOperationStartsWithUppercase ObservabilityTelemetryValuesParamsFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryValuesParamsFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersOperationIncludes, ObservabilityTelemetryValuesParamsFiltersOperationNotIncludes, ObservabilityTelemetryValuesParamsFiltersOperationStartsWith, ObservabilityTelemetryValuesParamsFiltersOperationRegex, ObservabilityTelemetryValuesParamsFiltersOperationExists, ObservabilityTelemetryValuesParamsFiltersOperationIsNull, ObservabilityTelemetryValuesParamsFiltersOperationIn, ObservabilityTelemetryValuesParamsFiltersOperationNotIn, ObservabilityTelemetryValuesParamsFiltersOperationEq, ObservabilityTelemetryValuesParamsFiltersOperationNeq, ObservabilityTelemetryValuesParamsFiltersOperationGt, ObservabilityTelemetryValuesParamsFiltersOperationGte, ObservabilityTelemetryValuesParamsFiltersOperationLt, ObservabilityTelemetryValuesParamsFiltersOperationLte, ObservabilityTelemetryValuesParamsFiltersOperationEquals, ObservabilityTelemetryValuesParamsFiltersOperationNotEquals, ObservabilityTelemetryValuesParamsFiltersOperationGreater, ObservabilityTelemetryValuesParamsFiltersOperationGreaterOrEquals, ObservabilityTelemetryValuesParamsFiltersOperationLess, ObservabilityTelemetryValuesParamsFiltersOperationLessOrEquals, ObservabilityTelemetryValuesParamsFiltersOperationIncludesUppercase, ObservabilityTelemetryValuesParamsFiltersOperationDoesNotInclude, ObservabilityTelemetryValuesParamsFiltersOperationMatchRegex, ObservabilityTelemetryValuesParamsFiltersOperationExistsUppercase, ObservabilityTelemetryValuesParamsFiltersOperationDoesNotExist, ObservabilityTelemetryValuesParamsFiltersOperationInUppercase, ObservabilityTelemetryValuesParamsFiltersOperationNotInUppercase, ObservabilityTelemetryValuesParamsFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParamsFiltersType string

const (
	ObservabilityTelemetryValuesParamsFiltersTypeString  ObservabilityTelemetryValuesParamsFiltersType = "string"
	ObservabilityTelemetryValuesParamsFiltersTypeNumber  ObservabilityTelemetryValuesParamsFiltersType = "number"
	ObservabilityTelemetryValuesParamsFiltersTypeBoolean ObservabilityTelemetryValuesParamsFiltersType = "boolean"
)

func (r ObservabilityTelemetryValuesParamsFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersTypeString, ObservabilityTelemetryValuesParamsFiltersTypeNumber, ObservabilityTelemetryValuesParamsFiltersTypeBoolean:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryValuesParamsFiltersValueUnion interface {
	ImplementsObservabilityTelemetryValuesParamsFiltersValueUnion()
}

// Search for a specific substring in the event.
type ObservabilityTelemetryValuesParamsNeedle struct {
	Value     param.Field[ObservabilityTelemetryValuesParamsNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                               `json:"isRegex"`
	MatchCase param.Field[bool]                                               `json:"matchCase"`
}

func (r ObservabilityTelemetryValuesParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryValuesParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion()
}
