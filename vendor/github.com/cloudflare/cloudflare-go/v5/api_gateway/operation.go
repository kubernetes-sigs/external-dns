// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/tidwall/gjson"
)

// OperationService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOperationService] method instead.
type OperationService struct {
	Options          []option.RequestOption
	SchemaValidation *OperationSchemaValidationService
}

// NewOperationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOperationService(opts ...option.RequestOption) (r *OperationService) {
	r = &OperationService{}
	r.Options = opts
	r.SchemaValidation = NewOperationSchemaValidationService(opts...)
	return
}

// Add one operation to a zone. Endpoints can contain path variables. Host, method,
// endpoint will be normalized to a canoncial form when creating an operation and
// must be unique on the zone. Inserting an operation that matches an existing one
// will return the record of the already existing operation and update its
// last_updated date.
func (r *OperationService) New(ctx context.Context, params OperationNewParams, opts ...option.RequestOption) (res *OperationNewResponse, err error) {
	var env OperationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/item", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve information about all operations on a zone
func (r *OperationService) List(ctx context.Context, params OperationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[OperationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations", params.ZoneID)
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

// Retrieve information about all operations on a zone
func (r *OperationService) ListAutoPaging(ctx context.Context, params OperationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[OperationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete an operation
func (r *OperationService) Delete(ctx context.Context, operationID string, body OperationDeleteParams, opts ...option.RequestOption) (res *OperationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", body.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add one or more operations to a zone. Endpoints can contain path variables.
// Host, method, endpoint will be normalized to a canoncial form when creating an
// operation and must be unique on the zone. Inserting an operation that matches an
// existing one will return the record of the already existing operation and update
// its last_updated date.
func (r *OperationService) BulkNew(ctx context.Context, params OperationBulkNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[OperationBulkNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations", params.ZoneID)
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

// Add one or more operations to a zone. Endpoints can contain path variables.
// Host, method, endpoint will be normalized to a canoncial form when creating an
// operation and must be unique on the zone. Inserting an operation that matches an
// existing one will return the record of the already existing operation and update
// its last_updated date.
func (r *OperationService) BulkNewAutoPaging(ctx context.Context, params OperationBulkNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[OperationBulkNewResponse] {
	return pagination.NewSinglePageAutoPager(r.BulkNew(ctx, params, opts...))
}

// Delete multiple operations
func (r *OperationService) BulkDelete(ctx context.Context, body OperationBulkDeleteParams, opts ...option.RequestOption) (res *OperationBulkDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve information about an operation
func (r *OperationService) Get(ctx context.Context, operationID string, params OperationGetParams, opts ...option.RequestOption) (res *OperationGetResponse, err error) {
	var env OperationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if operationID == "" {
		err = errors.New("missing required operation_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/operations/%s", params.ZoneID, operationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type OperationNewResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationNewResponseMethod `json:"method,required"`
	// UUID.
	OperationID string                       `json:"operation_id,required"`
	Features    OperationNewResponseFeatures `json:"features"`
	JSON        operationNewResponseJSON     `json:"-"`
}

// operationNewResponseJSON contains the JSON metadata for the struct
// [OperationNewResponse]
type operationNewResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationNewResponseMethod string

const (
	OperationNewResponseMethodGet     OperationNewResponseMethod = "GET"
	OperationNewResponseMethodPost    OperationNewResponseMethod = "POST"
	OperationNewResponseMethodHead    OperationNewResponseMethod = "HEAD"
	OperationNewResponseMethodOptions OperationNewResponseMethod = "OPTIONS"
	OperationNewResponseMethodPut     OperationNewResponseMethod = "PUT"
	OperationNewResponseMethodDelete  OperationNewResponseMethod = "DELETE"
	OperationNewResponseMethodConnect OperationNewResponseMethod = "CONNECT"
	OperationNewResponseMethodPatch   OperationNewResponseMethod = "PATCH"
	OperationNewResponseMethodTrace   OperationNewResponseMethod = "TRACE"
)

func (r OperationNewResponseMethod) IsKnown() bool {
	switch r {
	case OperationNewResponseMethodGet, OperationNewResponseMethodPost, OperationNewResponseMethodHead, OperationNewResponseMethodOptions, OperationNewResponseMethodPut, OperationNewResponseMethodDelete, OperationNewResponseMethodConnect, OperationNewResponseMethodPatch, OperationNewResponseMethodTrace:
		return true
	}
	return false
}

type OperationNewResponseFeatures struct {
	// This field can have the runtime type of
	// [OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing"`
	// This field can have the runtime type of
	// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals"`
	// This field can have the runtime type of
	// [OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas"`
	// This field can have the runtime type of
	// [OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{} `json:"schema_info"`
	// This field can have the runtime type of
	// [OperationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{}                      `json:"thresholds"`
	JSON       operationNewResponseFeaturesJSON `json:"-"`
	union      OperationNewResponseFeaturesUnion
}

// operationNewResponseFeaturesJSON contains the JSON metadata for the struct
// [OperationNewResponseFeatures]
type operationNewResponseFeaturesJSON struct {
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	ParameterSchemas    apijson.Field
	SchemaInfo          apijson.Field
	Thresholds          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r operationNewResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *OperationNewResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = OperationNewResponseFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OperationNewResponseFeaturesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [OperationNewResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r OperationNewResponseFeatures) AsUnion() OperationNewResponseFeaturesUnion {
	return r.union
}

// Union satisfied by
// [OperationNewResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals] or
// [OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
type OperationNewResponseFeaturesUnion interface {
	implementsOperationNewResponseFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationNewResponseFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationNewResponseFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds OperationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON contains the
// JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureThresholds]
type operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r OperationNewResponseFeaturesAPIShieldOperationFeatureThresholds) implementsOperationNewResponseFeatures() {
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
	// The total number of auth-ids seen across this calculation.
	AuthIDTokens int64 `json:"auth_id_tokens"`
	// The number of data points used for the threshold suggestion calculation.
	DataPoints  int64     `json:"data_points"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The p50 quantile of requests (in period_seconds).
	P50 int64 `json:"p50"`
	// The p90 quantile of requests (in period_seconds).
	P90 int64 `json:"p90"`
	// The p99 quantile of requests (in period_seconds).
	P99 int64 `json:"p99"`
	// The period over which this threshold is suggested.
	PeriodSeconds int64 `json:"period_seconds"`
	// The estimated number of requests covered by these calculations.
	Requests int64 `json:"requests"`
	// The suggested threshold in requests done by the same auth_id or period_seconds.
	SuggestedThreshold int64                                                                         `json:"suggested_threshold"`
	JSON               operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
	AuthIDTokens       apijson.Field
	DataPoints         apijson.Field
	LastUpdated        apijson.Field
	P50                apijson.Field
	P90                apijson.Field
	P99                apijson.Field
	PeriodSeconds      apijson.Field
	Requests           apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsOperationNewResponseFeatures() {
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                               `json:"responses,nullable"`
	JSON      operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains the
// JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRouting]
type operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRouting) implementsOperationNewResponseFeatures() {
}

// API Routing settings on endpoint.
type OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                                        `json:"route"`
	JSON  operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsOperationNewResponseFeatures() {
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                                     `json:"last_updated" format:"date-time"`
	SuggestedThreshold OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                           `json:"mean"`
	JSON operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains the
// JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo]
type operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo) implementsOperationNewResponseFeatures() {
}

type OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                                    `json:"name"`
	JSON operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, OperationNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type OperationListResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationListResponseMethod `json:"method,required"`
	// UUID.
	OperationID string                        `json:"operation_id,required"`
	Features    OperationListResponseFeatures `json:"features"`
	JSON        operationListResponseJSON     `json:"-"`
}

// operationListResponseJSON contains the JSON metadata for the struct
// [OperationListResponse]
type operationListResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationListResponseMethod string

const (
	OperationListResponseMethodGet     OperationListResponseMethod = "GET"
	OperationListResponseMethodPost    OperationListResponseMethod = "POST"
	OperationListResponseMethodHead    OperationListResponseMethod = "HEAD"
	OperationListResponseMethodOptions OperationListResponseMethod = "OPTIONS"
	OperationListResponseMethodPut     OperationListResponseMethod = "PUT"
	OperationListResponseMethodDelete  OperationListResponseMethod = "DELETE"
	OperationListResponseMethodConnect OperationListResponseMethod = "CONNECT"
	OperationListResponseMethodPatch   OperationListResponseMethod = "PATCH"
	OperationListResponseMethodTrace   OperationListResponseMethod = "TRACE"
)

func (r OperationListResponseMethod) IsKnown() bool {
	switch r {
	case OperationListResponseMethodGet, OperationListResponseMethodPost, OperationListResponseMethodHead, OperationListResponseMethodOptions, OperationListResponseMethodPut, OperationListResponseMethodDelete, OperationListResponseMethodConnect, OperationListResponseMethodPatch, OperationListResponseMethodTrace:
		return true
	}
	return false
}

type OperationListResponseFeatures struct {
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{} `json:"schema_info"`
	// This field can have the runtime type of
	// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{}                       `json:"thresholds"`
	JSON       operationListResponseFeaturesJSON `json:"-"`
	union      OperationListResponseFeaturesUnion
}

// operationListResponseFeaturesJSON contains the JSON metadata for the struct
// [OperationListResponseFeatures]
type operationListResponseFeaturesJSON struct {
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	ParameterSchemas    apijson.Field
	SchemaInfo          apijson.Field
	Thresholds          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r operationListResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *OperationListResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = OperationListResponseFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OperationListResponseFeaturesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r OperationListResponseFeatures) AsUnion() OperationListResponseFeaturesUnion {
	return r.union
}

// Union satisfied by
// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals] or
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
type OperationListResponseFeaturesUnion interface {
	implementsOperationListResponseFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationListResponseFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type OperationListResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON contains
// the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholds]
type operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureThresholds) implementsOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
	// The total number of auth-ids seen across this calculation.
	AuthIDTokens int64 `json:"auth_id_tokens"`
	// The number of data points used for the threshold suggestion calculation.
	DataPoints  int64     `json:"data_points"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The p50 quantile of requests (in period_seconds).
	P50 int64 `json:"p50"`
	// The p90 quantile of requests (in period_seconds).
	P90 int64 `json:"p90"`
	// The p99 quantile of requests (in period_seconds).
	P99 int64 `json:"p99"`
	// The period over which this threshold is suggested.
	PeriodSeconds int64 `json:"period_seconds"`
	// The estimated number of requests covered by these calculations.
	Requests int64 `json:"requests"`
	// The suggested threshold in requests done by the same auth_id or period_seconds.
	SuggestedThreshold int64                                                                          `json:"suggested_threshold"`
	JSON               operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
	AuthIDTokens       apijson.Field
	DataPoints         apijson.Field
	LastUpdated        apijson.Field
	P50                apijson.Field
	P90                apijson.Field
	P99                apijson.Field
	PeriodSeconds      apijson.Field
	Requests           apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                `json:"responses,nullable"`
	JSON      operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains
// the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting]
type operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureAPIRouting) implementsOperationListResponseFeatures() {
}

// API Routing settings on endpoint.
type OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                                         `json:"route"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                                      `json:"last_updated" format:"date-time"`
	SuggestedThreshold OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                            `json:"mean"`
	JSON operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                  `json:"upper"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                  `json:"upper"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                  `json:"upper"`
	JSON  operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains
// the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo]
type operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfo) implementsOperationListResponseFeatures() {
}

type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                                     `json:"name"`
	JSON operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, OperationListResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type OperationDeleteResponse struct {
	Errors   Message `json:"errors,required"`
	Messages Message `json:"messages,required"`
	// Whether the API call was successful.
	Success OperationDeleteResponseSuccess `json:"success,required"`
	JSON    operationDeleteResponseJSON    `json:"-"`
}

// operationDeleteResponseJSON contains the JSON metadata for the struct
// [OperationDeleteResponse]
type operationDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationDeleteResponseSuccess bool

const (
	OperationDeleteResponseSuccessTrue OperationDeleteResponseSuccess = true
)

func (r OperationDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case OperationDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type OperationBulkNewResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationBulkNewResponseMethod `json:"method,required"`
	// UUID.
	OperationID string                           `json:"operation_id,required"`
	Features    OperationBulkNewResponseFeatures `json:"features"`
	JSON        operationBulkNewResponseJSON     `json:"-"`
}

// operationBulkNewResponseJSON contains the JSON metadata for the struct
// [OperationBulkNewResponse]
type operationBulkNewResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationBulkNewResponseMethod string

const (
	OperationBulkNewResponseMethodGet     OperationBulkNewResponseMethod = "GET"
	OperationBulkNewResponseMethodPost    OperationBulkNewResponseMethod = "POST"
	OperationBulkNewResponseMethodHead    OperationBulkNewResponseMethod = "HEAD"
	OperationBulkNewResponseMethodOptions OperationBulkNewResponseMethod = "OPTIONS"
	OperationBulkNewResponseMethodPut     OperationBulkNewResponseMethod = "PUT"
	OperationBulkNewResponseMethodDelete  OperationBulkNewResponseMethod = "DELETE"
	OperationBulkNewResponseMethodConnect OperationBulkNewResponseMethod = "CONNECT"
	OperationBulkNewResponseMethodPatch   OperationBulkNewResponseMethod = "PATCH"
	OperationBulkNewResponseMethodTrace   OperationBulkNewResponseMethod = "TRACE"
)

func (r OperationBulkNewResponseMethod) IsKnown() bool {
	switch r {
	case OperationBulkNewResponseMethodGet, OperationBulkNewResponseMethodPost, OperationBulkNewResponseMethodHead, OperationBulkNewResponseMethodOptions, OperationBulkNewResponseMethodPut, OperationBulkNewResponseMethodDelete, OperationBulkNewResponseMethodConnect, OperationBulkNewResponseMethodPatch, OperationBulkNewResponseMethodTrace:
		return true
	}
	return false
}

type OperationBulkNewResponseFeatures struct {
	// This field can have the runtime type of
	// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing"`
	// This field can have the runtime type of
	// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals"`
	// This field can have the runtime type of
	// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas"`
	// This field can have the runtime type of
	// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{} `json:"schema_info"`
	// This field can have the runtime type of
	// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{}                          `json:"thresholds"`
	JSON       operationBulkNewResponseFeaturesJSON `json:"-"`
	union      OperationBulkNewResponseFeaturesUnion
}

// operationBulkNewResponseFeaturesJSON contains the JSON metadata for the struct
// [OperationBulkNewResponseFeatures]
type operationBulkNewResponseFeaturesJSON struct {
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	ParameterSchemas    apijson.Field
	SchemaInfo          apijson.Field
	Thresholds          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r operationBulkNewResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *OperationBulkNewResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = OperationBulkNewResponseFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OperationBulkNewResponseFeaturesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r OperationBulkNewResponseFeatures) AsUnion() OperationBulkNewResponseFeaturesUnion {
	return r.union
}

// Union satisfied by
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
// or [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
type OperationBulkNewResponseFeaturesUnion interface {
	implementsOperationBulkNewResponseFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationBulkNewResponseFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON contains
// the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholds]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholds) implementsOperationBulkNewResponseFeatures() {
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
	// The total number of auth-ids seen across this calculation.
	AuthIDTokens int64 `json:"auth_id_tokens"`
	// The number of data points used for the threshold suggestion calculation.
	DataPoints  int64     `json:"data_points"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The p50 quantile of requests (in period_seconds).
	P50 int64 `json:"p50"`
	// The p90 quantile of requests (in period_seconds).
	P90 int64 `json:"p90"`
	// The p99 quantile of requests (in period_seconds).
	P99 int64 `json:"p99"`
	// The period over which this threshold is suggested.
	PeriodSeconds int64 `json:"period_seconds"`
	// The estimated number of requests covered by these calculations.
	Requests int64 `json:"requests"`
	// The suggested threshold in requests done by the same auth_id or period_seconds.
	SuggestedThreshold int64                                                                             `json:"suggested_threshold"`
	JSON               operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
	AuthIDTokens       apijson.Field
	DataPoints         apijson.Field
	LastUpdated        apijson.Field
	P50                apijson.Field
	P90                apijson.Field
	P99                apijson.Field
	PeriodSeconds      apijson.Field
	Requests           apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsOperationBulkNewResponseFeatures() {
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                   `json:"responses,nullable"`
	JSON      operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains
// the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRouting]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRouting) implementsOperationBulkNewResponseFeatures() {
}

// API Routing settings on endpoint.
type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                                            `json:"route"`
	JSON  operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsOperationBulkNewResponseFeatures() {
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                                         `json:"last_updated" format:"date-time"`
	SuggestedThreshold OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                               `json:"mean"`
	JSON operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                     `json:"upper"`
	JSON  operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                     `json:"upper"`
	JSON  operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                     `json:"upper"`
	JSON  operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains
// the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfo) implementsOperationBulkNewResponseFeatures() {
}

type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                                        `json:"name"`
	JSON operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, OperationBulkNewResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type OperationBulkDeleteResponse struct {
	Errors   Message `json:"errors,required"`
	Messages Message `json:"messages,required"`
	// Whether the API call was successful.
	Success OperationBulkDeleteResponseSuccess `json:"success,required"`
	JSON    operationBulkDeleteResponseJSON    `json:"-"`
}

// operationBulkDeleteResponseJSON contains the JSON metadata for the struct
// [OperationBulkDeleteResponse]
type operationBulkDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationBulkDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationBulkDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationBulkDeleteResponseSuccess bool

const (
	OperationBulkDeleteResponseSuccessTrue OperationBulkDeleteResponseSuccess = true
)

func (r OperationBulkDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case OperationBulkDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type OperationGetResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method OperationGetResponseMethod `json:"method,required"`
	// UUID.
	OperationID string                       `json:"operation_id,required"`
	Features    OperationGetResponseFeatures `json:"features"`
	JSON        operationGetResponseJSON     `json:"-"`
}

// operationGetResponseJSON contains the JSON metadata for the struct
// [OperationGetResponse]
type operationGetResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseJSON) RawJSON() string {
	return r.raw
}

// The HTTP method used to access the endpoint.
type OperationGetResponseMethod string

const (
	OperationGetResponseMethodGet     OperationGetResponseMethod = "GET"
	OperationGetResponseMethodPost    OperationGetResponseMethod = "POST"
	OperationGetResponseMethodHead    OperationGetResponseMethod = "HEAD"
	OperationGetResponseMethodOptions OperationGetResponseMethod = "OPTIONS"
	OperationGetResponseMethodPut     OperationGetResponseMethod = "PUT"
	OperationGetResponseMethodDelete  OperationGetResponseMethod = "DELETE"
	OperationGetResponseMethodConnect OperationGetResponseMethod = "CONNECT"
	OperationGetResponseMethodPatch   OperationGetResponseMethod = "PATCH"
	OperationGetResponseMethodTrace   OperationGetResponseMethod = "TRACE"
)

func (r OperationGetResponseMethod) IsKnown() bool {
	switch r {
	case OperationGetResponseMethodGet, OperationGetResponseMethodPost, OperationGetResponseMethodHead, OperationGetResponseMethodOptions, OperationGetResponseMethodPut, OperationGetResponseMethodDelete, OperationGetResponseMethodConnect, OperationGetResponseMethodPatch, OperationGetResponseMethodTrace:
		return true
	}
	return false
}

type OperationGetResponseFeatures struct {
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{} `json:"schema_info"`
	// This field can have the runtime type of
	// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{}                      `json:"thresholds"`
	JSON       operationGetResponseFeaturesJSON `json:"-"`
	union      OperationGetResponseFeaturesUnion
}

// operationGetResponseFeaturesJSON contains the JSON metadata for the struct
// [OperationGetResponseFeatures]
type operationGetResponseFeaturesJSON struct {
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	ParameterSchemas    apijson.Field
	SchemaInfo          apijson.Field
	Thresholds          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r operationGetResponseFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *OperationGetResponseFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = OperationGetResponseFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OperationGetResponseFeaturesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r OperationGetResponseFeatures) AsUnion() OperationGetResponseFeaturesUnion {
	return r.union
}

// Union satisfied by
// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds],
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas],
// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting],
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals] or
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo].
type OperationGetResponseFeaturesUnion interface {
	implementsOperationGetResponseFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OperationGetResponseFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON contains the
// JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds]
type operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureThresholds) implementsOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
	// The total number of auth-ids seen across this calculation.
	AuthIDTokens int64 `json:"auth_id_tokens"`
	// The number of data points used for the threshold suggestion calculation.
	DataPoints  int64     `json:"data_points"`
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The p50 quantile of requests (in period_seconds).
	P50 int64 `json:"p50"`
	// The p90 quantile of requests (in period_seconds).
	P90 int64 `json:"p90"`
	// The p99 quantile of requests (in period_seconds).
	P99 int64 `json:"p99"`
	// The period over which this threshold is suggested.
	PeriodSeconds int64 `json:"period_seconds"`
	// The estimated number of requests covered by these calculations.
	Requests int64 `json:"requests"`
	// The suggested threshold in requests done by the same auth_id or period_seconds.
	SuggestedThreshold int64                                                                         `json:"suggested_threshold"`
	JSON               operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
	AuthIDTokens       apijson.Field
	DataPoints         apijson.Field
	LastUpdated        apijson.Field
	P50                apijson.Field
	P90                apijson.Field
	P99                apijson.Field
	PeriodSeconds      apijson.Field
	Requests           apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas]
type operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemas) implementsOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                               `json:"responses,nullable"`
	JSON      operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON contains the
// JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting]
type operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRouting) implementsOperationGetResponseFeatures() {
}

// API Routing settings on endpoint.
type OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                                        `json:"route"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                                     `json:"last_updated" format:"date-time"`
	SuggestedThreshold OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                           `json:"mean"`
	JSON operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                 `json:"upper"`
	JSON  operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON contains the
// JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo]
type operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfo) implementsOperationGetResponseFeatures() {
}

type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                                    `json:"name"`
	JSON operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, OperationGetResponseFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type OperationNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint param.Field[string] `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method param.Field[OperationNewParamsMethod] `json:"method,required"`
}

func (r OperationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method used to access the endpoint.
type OperationNewParamsMethod string

const (
	OperationNewParamsMethodGet     OperationNewParamsMethod = "GET"
	OperationNewParamsMethodPost    OperationNewParamsMethod = "POST"
	OperationNewParamsMethodHead    OperationNewParamsMethod = "HEAD"
	OperationNewParamsMethodOptions OperationNewParamsMethod = "OPTIONS"
	OperationNewParamsMethodPut     OperationNewParamsMethod = "PUT"
	OperationNewParamsMethodDelete  OperationNewParamsMethod = "DELETE"
	OperationNewParamsMethodConnect OperationNewParamsMethod = "CONNECT"
	OperationNewParamsMethodPatch   OperationNewParamsMethod = "PATCH"
	OperationNewParamsMethodTrace   OperationNewParamsMethod = "TRACE"
)

func (r OperationNewParamsMethod) IsKnown() bool {
	switch r {
	case OperationNewParamsMethodGet, OperationNewParamsMethodPost, OperationNewParamsMethodHead, OperationNewParamsMethodOptions, OperationNewParamsMethodPut, OperationNewParamsMethodDelete, OperationNewParamsMethodConnect, OperationNewParamsMethodPatch, OperationNewParamsMethodTrace:
		return true
	}
	return false
}

type OperationNewResponseEnvelope struct {
	Errors   Message              `json:"errors,required"`
	Messages Message              `json:"messages,required"`
	Result   OperationNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success OperationNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    operationNewResponseEnvelopeJSON    `json:"-"`
}

// operationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OperationNewResponseEnvelope]
type operationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationNewResponseEnvelopeSuccess bool

const (
	OperationNewResponseEnvelopeSuccessTrue OperationNewResponseEnvelopeSuccess = true
)

func (r OperationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OperationListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Direction to order results.
	Direction param.Field[OperationListParamsDirection] `query:"direction"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]OperationListParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Field to order by. When requesting a feature, the feature keys are available for
	// ordering as well, e.g., `thresholds.suggested_threshold`.
	Order param.Field[OperationListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [OperationListParams]'s query parameters as `url.Values`.
func (r OperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type OperationListParamsDirection string

const (
	OperationListParamsDirectionAsc  OperationListParamsDirection = "asc"
	OperationListParamsDirectionDesc OperationListParamsDirection = "desc"
)

func (r OperationListParamsDirection) IsKnown() bool {
	switch r {
	case OperationListParamsDirectionAsc, OperationListParamsDirectionDesc:
		return true
	}
	return false
}

type OperationListParamsFeature string

const (
	OperationListParamsFeatureThresholds       OperationListParamsFeature = "thresholds"
	OperationListParamsFeatureParameterSchemas OperationListParamsFeature = "parameter_schemas"
	OperationListParamsFeatureSchemaInfo       OperationListParamsFeature = "schema_info"
)

func (r OperationListParamsFeature) IsKnown() bool {
	switch r {
	case OperationListParamsFeatureThresholds, OperationListParamsFeatureParameterSchemas, OperationListParamsFeatureSchemaInfo:
		return true
	}
	return false
}

// Field to order by. When requesting a feature, the feature keys are available for
// ordering as well, e.g., `thresholds.suggested_threshold`.
type OperationListParamsOrder string

const (
	OperationListParamsOrderMethod        OperationListParamsOrder = "method"
	OperationListParamsOrderHost          OperationListParamsOrder = "host"
	OperationListParamsOrderEndpoint      OperationListParamsOrder = "endpoint"
	OperationListParamsOrderThresholdsKey OperationListParamsOrder = "thresholds.$key"
)

func (r OperationListParamsOrder) IsKnown() bool {
	switch r {
	case OperationListParamsOrderMethod, OperationListParamsOrderHost, OperationListParamsOrderEndpoint, OperationListParamsOrderThresholdsKey:
		return true
	}
	return false
}

type OperationDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OperationBulkNewParams struct {
	// Identifier.
	ZoneID param.Field[string]          `path:"zone_id,required"`
	Body   []OperationBulkNewParamsBody `json:"body,required"`
}

func (r OperationBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OperationBulkNewParamsBody struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint param.Field[string] `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method param.Field[OperationBulkNewParamsBodyMethod] `json:"method,required"`
}

func (r OperationBulkNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The HTTP method used to access the endpoint.
type OperationBulkNewParamsBodyMethod string

const (
	OperationBulkNewParamsBodyMethodGet     OperationBulkNewParamsBodyMethod = "GET"
	OperationBulkNewParamsBodyMethodPost    OperationBulkNewParamsBodyMethod = "POST"
	OperationBulkNewParamsBodyMethodHead    OperationBulkNewParamsBodyMethod = "HEAD"
	OperationBulkNewParamsBodyMethodOptions OperationBulkNewParamsBodyMethod = "OPTIONS"
	OperationBulkNewParamsBodyMethodPut     OperationBulkNewParamsBodyMethod = "PUT"
	OperationBulkNewParamsBodyMethodDelete  OperationBulkNewParamsBodyMethod = "DELETE"
	OperationBulkNewParamsBodyMethodConnect OperationBulkNewParamsBodyMethod = "CONNECT"
	OperationBulkNewParamsBodyMethodPatch   OperationBulkNewParamsBodyMethod = "PATCH"
	OperationBulkNewParamsBodyMethodTrace   OperationBulkNewParamsBodyMethod = "TRACE"
)

func (r OperationBulkNewParamsBodyMethod) IsKnown() bool {
	switch r {
	case OperationBulkNewParamsBodyMethodGet, OperationBulkNewParamsBodyMethodPost, OperationBulkNewParamsBodyMethodHead, OperationBulkNewParamsBodyMethodOptions, OperationBulkNewParamsBodyMethodPut, OperationBulkNewParamsBodyMethodDelete, OperationBulkNewParamsBodyMethodConnect, OperationBulkNewParamsBodyMethodPatch, OperationBulkNewParamsBodyMethodTrace:
		return true
	}
	return false
}

type OperationBulkDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type OperationGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]OperationGetParamsFeature] `query:"feature"`
}

// URLQuery serializes [OperationGetParams]'s query parameters as `url.Values`.
func (r OperationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OperationGetParamsFeature string

const (
	OperationGetParamsFeatureThresholds       OperationGetParamsFeature = "thresholds"
	OperationGetParamsFeatureParameterSchemas OperationGetParamsFeature = "parameter_schemas"
	OperationGetParamsFeatureSchemaInfo       OperationGetParamsFeature = "schema_info"
)

func (r OperationGetParamsFeature) IsKnown() bool {
	switch r {
	case OperationGetParamsFeatureThresholds, OperationGetParamsFeatureParameterSchemas, OperationGetParamsFeatureSchemaInfo:
		return true
	}
	return false
}

type OperationGetResponseEnvelope struct {
	Errors   Message              `json:"errors,required"`
	Messages Message              `json:"messages,required"`
	Result   OperationGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success OperationGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    operationGetResponseEnvelopeJSON    `json:"-"`
}

// operationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [OperationGetResponseEnvelope]
type operationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OperationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r operationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OperationGetResponseEnvelopeSuccess bool

const (
	OperationGetResponseEnvelopeSuccessTrue OperationGetResponseEnvelopeSuccess = true
)

func (r OperationGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OperationGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
