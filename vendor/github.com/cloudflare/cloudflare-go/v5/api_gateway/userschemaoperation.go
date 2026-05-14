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

// UserSchemaOperationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserSchemaOperationService] method instead.
type UserSchemaOperationService struct {
	Options []option.RequestOption
}

// NewUserSchemaOperationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserSchemaOperationService(opts ...option.RequestOption) (r *UserSchemaOperationService) {
	r = &UserSchemaOperationService{}
	r.Options = opts
	return
}

// Retrieves all operations from the schema. Operations that already exist in API
// Shield Endpoint Management will be returned as full operations.
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaOperationService) List(ctx context.Context, schemaID string, params UserSchemaOperationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UserSchemaOperationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if schemaID == "" {
		err = errors.New("missing required schema_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/user_schemas/%s/operations", params.ZoneID, schemaID)
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

// Retrieves all operations from the schema. Operations that already exist in API
// Shield Endpoint Management will be returned as full operations.
//
// Deprecated: Use
// [Schema Validation API](https://developers.cloudflare.com/api/resources/schema_validation/)
// instead.
func (r *UserSchemaOperationService) ListAutoPaging(ctx context.Context, schemaID string, params UserSchemaOperationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UserSchemaOperationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, schemaID, params, opts...))
}

type UserSchemaOperationListResponse struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method UserSchemaOperationListResponseMethod `json:"method,required"`
	// This field can have the runtime type of
	// [UserSchemaOperationListResponseAPIShieldOperationFeatures].
	Features    interface{} `json:"features"`
	LastUpdated time.Time   `json:"last_updated" format:"date-time"`
	// UUID.
	OperationID string                              `json:"operation_id"`
	JSON        userSchemaOperationListResponseJSON `json:"-"`
	union       UserSchemaOperationListResponseUnion
}

// userSchemaOperationListResponseJSON contains the JSON metadata for the struct
// [UserSchemaOperationListResponse]
type userSchemaOperationListResponseJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	Features    apijson.Field
	LastUpdated apijson.Field
	OperationID apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r userSchemaOperationListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *UserSchemaOperationListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = UserSchemaOperationListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [UserSchemaOperationListResponseUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSchemaOperationListResponseAPIShieldOperation],
// [UserSchemaOperationListResponseAPIShieldBasicOperation].
func (r UserSchemaOperationListResponse) AsUnion() UserSchemaOperationListResponseUnion {
	return r.union
}

// Union satisfied by [UserSchemaOperationListResponseAPIShieldOperation] or
// [UserSchemaOperationListResponseAPIShieldBasicOperation].
type UserSchemaOperationListResponseUnion interface {
	implementsUserSchemaOperationListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSchemaOperationListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldOperation{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldBasicOperation{}),
		},
	)
}

type UserSchemaOperationListResponseAPIShieldOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host        string    `json:"host,required" format:"hostname"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The HTTP method used to access the endpoint.
	Method UserSchemaOperationListResponseAPIShieldOperationMethod `json:"method,required"`
	// UUID.
	OperationID string                                                    `json:"operation_id,required"`
	Features    UserSchemaOperationListResponseAPIShieldOperationFeatures `json:"features"`
	JSON        userSchemaOperationListResponseAPIShieldOperationJSON     `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationJSON contains the JSON metadata
// for the struct [UserSchemaOperationListResponseAPIShieldOperation]
type userSchemaOperationListResponseAPIShieldOperationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	LastUpdated apijson.Field
	Method      apijson.Field
	OperationID apijson.Field
	Features    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldOperation) implementsUserSchemaOperationListResponse() {
}

// The HTTP method used to access the endpoint.
type UserSchemaOperationListResponseAPIShieldOperationMethod string

const (
	UserSchemaOperationListResponseAPIShieldOperationMethodGet     UserSchemaOperationListResponseAPIShieldOperationMethod = "GET"
	UserSchemaOperationListResponseAPIShieldOperationMethodPost    UserSchemaOperationListResponseAPIShieldOperationMethod = "POST"
	UserSchemaOperationListResponseAPIShieldOperationMethodHead    UserSchemaOperationListResponseAPIShieldOperationMethod = "HEAD"
	UserSchemaOperationListResponseAPIShieldOperationMethodOptions UserSchemaOperationListResponseAPIShieldOperationMethod = "OPTIONS"
	UserSchemaOperationListResponseAPIShieldOperationMethodPut     UserSchemaOperationListResponseAPIShieldOperationMethod = "PUT"
	UserSchemaOperationListResponseAPIShieldOperationMethodDelete  UserSchemaOperationListResponseAPIShieldOperationMethod = "DELETE"
	UserSchemaOperationListResponseAPIShieldOperationMethodConnect UserSchemaOperationListResponseAPIShieldOperationMethod = "CONNECT"
	UserSchemaOperationListResponseAPIShieldOperationMethodPatch   UserSchemaOperationListResponseAPIShieldOperationMethod = "PATCH"
	UserSchemaOperationListResponseAPIShieldOperationMethodTrace   UserSchemaOperationListResponseAPIShieldOperationMethod = "TRACE"
)

func (r UserSchemaOperationListResponseAPIShieldOperationMethod) IsKnown() bool {
	switch r {
	case UserSchemaOperationListResponseAPIShieldOperationMethodGet, UserSchemaOperationListResponseAPIShieldOperationMethodPost, UserSchemaOperationListResponseAPIShieldOperationMethodHead, UserSchemaOperationListResponseAPIShieldOperationMethodOptions, UserSchemaOperationListResponseAPIShieldOperationMethodPut, UserSchemaOperationListResponseAPIShieldOperationMethodDelete, UserSchemaOperationListResponseAPIShieldOperationMethodConnect, UserSchemaOperationListResponseAPIShieldOperationMethodPatch, UserSchemaOperationListResponseAPIShieldOperationMethodTrace:
		return true
	}
	return false
}

type UserSchemaOperationListResponseAPIShieldOperationFeatures struct {
	// This field can have the runtime type of
	// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting].
	APIRouting interface{} `json:"api_routing"`
	// This field can have the runtime type of
	// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals].
	ConfidenceIntervals interface{} `json:"confidence_intervals"`
	// This field can have the runtime type of
	// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas].
	ParameterSchemas interface{} `json:"parameter_schemas"`
	// This field can have the runtime type of
	// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo].
	SchemaInfo interface{} `json:"schema_info"`
	// This field can have the runtime type of
	// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds].
	Thresholds interface{}                                                   `json:"thresholds"`
	JSON       userSchemaOperationListResponseAPIShieldOperationFeaturesJSON `json:"-"`
	union      UserSchemaOperationListResponseAPIShieldOperationFeaturesUnion
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesJSON contains the JSON
// metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeatures]
type userSchemaOperationListResponseAPIShieldOperationFeaturesJSON struct {
	APIRouting          apijson.Field
	ConfidenceIntervals apijson.Field
	ParameterSchemas    apijson.Field
	SchemaInfo          apijson.Field
	Thresholds          apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesJSON) RawJSON() string {
	return r.raw
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeatures) UnmarshalJSON(data []byte) (err error) {
	*r = UserSchemaOperationListResponseAPIShieldOperationFeatures{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds],
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas],
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRouting],
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervals],
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfo].
func (r UserSchemaOperationListResponseAPIShieldOperationFeatures) AsUnion() UserSchemaOperationListResponseAPIShieldOperationFeaturesUnion {
	return r.union
}

// Union satisfied by
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds],
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas],
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRouting],
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervals]
// or
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfo].
type UserSchemaOperationListResponseAPIShieldOperationFeaturesUnion interface {
	implementsUserSchemaOperationListResponseAPIShieldOperationFeatures()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserSchemaOperationListResponseAPIShieldOperationFeaturesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRouting{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervals{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfo{}),
		},
	)
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds struct {
	Thresholds UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds `json:"thresholds"`
	JSON       userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsJSON       `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsJSON struct {
	Thresholds  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholds) implementsUserSchemaOperationListResponseAPIShieldOperationFeatures() {
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds struct {
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
	SuggestedThreshold int64                                                                                                      `json:"suggested_threshold"`
	JSON               userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON struct {
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

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholds) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureThresholdsThresholdsJSON) RawJSON() string {
	return r.raw
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas struct {
	ParameterSchemas UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas `json:"parameter_schemas,required"`
	JSON             userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasJSON             `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasJSON struct {
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemas) implementsUserSchemaOperationListResponseAPIShieldOperationFeatures() {
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An operation schema object containing a response.
	ParameterSchemas UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas `json:"parameter_schemas"`
	JSON             userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON             `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON struct {
	LastUpdated      apijson.Field
	ParameterSchemas apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

// An operation schema object containing a response.
type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas struct {
	// An array containing the learned parameter schemas.
	Parameters []interface{} `json:"parameters"`
	// An empty response object. This field is required to yield a valid operation
	// schema.
	Responses interface{}                                                                                                                            `json:"responses,nullable"`
	JSON      userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON struct {
	Parameters  apijson.Field
	Responses   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemas) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureParameterSchemasParameterSchemasParameterSchemasJSON) RawJSON() string {
	return r.raw
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRouting struct {
	// API Routing settings on endpoint.
	APIRouting UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting `json:"api_routing"`
	JSON       userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingJSON       `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRouting]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingJSON struct {
	APIRouting  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRouting) implementsUserSchemaOperationListResponseAPIShieldOperationFeatures() {
}

// API Routing settings on endpoint.
type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting struct {
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// Target route.
	Route string                                                                                                     `json:"route"`
	JSON  userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON struct {
	LastUpdated apijson.Field
	Route       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRouting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureAPIRoutingAPIRoutingJSON) RawJSON() string {
	return r.raw
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervals struct {
	ConfidenceIntervals UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals `json:"confidence_intervals"`
	JSON                userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON                `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervals]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON struct {
	ConfidenceIntervals apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervals) implementsUserSchemaOperationListResponseAPIShieldOperationFeatures() {
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals struct {
	LastUpdated        time.Time                                                                                                                                  `json:"last_updated" format:"date-time"`
	SuggestedThreshold UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold `json:"suggested_threshold"`
	JSON               userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON               `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON struct {
	LastUpdated        apijson.Field
	SuggestedThreshold apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold struct {
	ConfidenceIntervals UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals `json:"confidence_intervals"`
	// Suggested threshold.
	Mean float64                                                                                                                                        `json:"mean"`
	JSON userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON struct {
	ConfidenceIntervals apijson.Field
	Mean                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThreshold) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdJSON) RawJSON() string {
	return r.raw
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals struct {
	// Upper and lower bound for percentile estimate
	P90 UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 `json:"p90"`
	// Upper and lower bound for percentile estimate
	P95 UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 `json:"p95"`
	// Upper and lower bound for percentile estimate
	P99  UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99  `json:"p99"`
	JSON userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON struct {
	P90         apijson.Field
	P95         apijson.Field
	P99         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsJSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                                              `json:"upper"`
	JSON  userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP90JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                                              `json:"upper"`
	JSON  userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP95JSON) RawJSON() string {
	return r.raw
}

// Upper and lower bound for percentile estimate
type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99 struct {
	// Lower bound for percentile estimate
	Lower float64 `json:"lower"`
	// Upper bound for percentile estimate
	Upper float64                                                                                                                                                              `json:"upper"`
	JSON  userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON struct {
	Lower       apijson.Field
	Upper       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureConfidenceIntervalsConfidenceIntervalsSuggestedThresholdConfidenceIntervalsP99JSON) RawJSON() string {
	return r.raw
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfo struct {
	SchemaInfo UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo `json:"schema_info"`
	JSON       userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoJSON       `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfo]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoJSON struct {
	SchemaInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfo) implementsUserSchemaOperationListResponseAPIShieldOperationFeatures() {
}

type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo struct {
	// Schema active on endpoint.
	ActiveSchema UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema `json:"active_schema"`
	// True if a Cloudflare-provided learned schema is available for this endpoint.
	LearnedAvailable bool `json:"learned_available"`
	// Action taken on requests failing validation.
	MitigationAction UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction `json:"mitigation_action,nullable"`
	JSON             userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON             `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON struct {
	ActiveSchema     apijson.Field
	LearnedAvailable apijson.Field
	MitigationAction apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoJSON) RawJSON() string {
	return r.raw
}

// Schema active on endpoint.
type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema struct {
	// UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// True if schema is Cloudflare-provided.
	IsLearned bool `json:"is_learned"`
	// Schema file name.
	Name string                                                                                                                 `json:"name"`
	JSON userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON `json:"-"`
}

// userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON
// contains the JSON metadata for the struct
// [UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema]
type userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	IsLearned   apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoActiveSchemaJSON) RawJSON() string {
	return r.raw
}

// Action taken on requests failing validation.
type UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction string

const (
	UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone  UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "none"
	UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog   UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "log"
	UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction = "block"
)

func (r UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationAction) IsKnown() bool {
	switch r {
	case UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionNone, UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionLog, UserSchemaOperationListResponseAPIShieldOperationFeaturesAPIShieldOperationFeatureSchemaInfoSchemaInfoMitigationActionBlock:
		return true
	}
	return false
}

type UserSchemaOperationListResponseAPIShieldBasicOperation struct {
	// The endpoint which can contain path parameter templates in curly braces, each
	// will be replaced from left to right with {varN}, starting with {var1}, during
	// insertion. This will further be Cloudflare-normalized upon insertion. See:
	// https://developers.cloudflare.com/rules/normalization/how-it-works/.
	Endpoint string `json:"endpoint,required" format:"uri-template"`
	// RFC3986-compliant host.
	Host string `json:"host,required" format:"hostname"`
	// The HTTP method used to access the endpoint.
	Method UserSchemaOperationListResponseAPIShieldBasicOperationMethod `json:"method,required"`
	JSON   userSchemaOperationListResponseAPIShieldBasicOperationJSON   `json:"-"`
}

// userSchemaOperationListResponseAPIShieldBasicOperationJSON contains the JSON
// metadata for the struct [UserSchemaOperationListResponseAPIShieldBasicOperation]
type userSchemaOperationListResponseAPIShieldBasicOperationJSON struct {
	Endpoint    apijson.Field
	Host        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserSchemaOperationListResponseAPIShieldBasicOperation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSchemaOperationListResponseAPIShieldBasicOperationJSON) RawJSON() string {
	return r.raw
}

func (r UserSchemaOperationListResponseAPIShieldBasicOperation) implementsUserSchemaOperationListResponse() {
}

// The HTTP method used to access the endpoint.
type UserSchemaOperationListResponseAPIShieldBasicOperationMethod string

const (
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodGet     UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "GET"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodPost    UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "POST"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodHead    UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "HEAD"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodOptions UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "OPTIONS"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodPut     UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "PUT"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodDelete  UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "DELETE"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodConnect UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "CONNECT"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodPatch   UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "PATCH"
	UserSchemaOperationListResponseAPIShieldBasicOperationMethodTrace   UserSchemaOperationListResponseAPIShieldBasicOperationMethod = "TRACE"
)

func (r UserSchemaOperationListResponseAPIShieldBasicOperationMethod) IsKnown() bool {
	switch r {
	case UserSchemaOperationListResponseAPIShieldBasicOperationMethodGet, UserSchemaOperationListResponseAPIShieldBasicOperationMethodPost, UserSchemaOperationListResponseAPIShieldBasicOperationMethodHead, UserSchemaOperationListResponseAPIShieldBasicOperationMethodOptions, UserSchemaOperationListResponseAPIShieldBasicOperationMethodPut, UserSchemaOperationListResponseAPIShieldBasicOperationMethodDelete, UserSchemaOperationListResponseAPIShieldBasicOperationMethodConnect, UserSchemaOperationListResponseAPIShieldBasicOperationMethodPatch, UserSchemaOperationListResponseAPIShieldBasicOperationMethodTrace:
		return true
	}
	return false
}

// The HTTP method used to access the endpoint.
type UserSchemaOperationListResponseMethod string

const (
	UserSchemaOperationListResponseMethodGet     UserSchemaOperationListResponseMethod = "GET"
	UserSchemaOperationListResponseMethodPost    UserSchemaOperationListResponseMethod = "POST"
	UserSchemaOperationListResponseMethodHead    UserSchemaOperationListResponseMethod = "HEAD"
	UserSchemaOperationListResponseMethodOptions UserSchemaOperationListResponseMethod = "OPTIONS"
	UserSchemaOperationListResponseMethodPut     UserSchemaOperationListResponseMethod = "PUT"
	UserSchemaOperationListResponseMethodDelete  UserSchemaOperationListResponseMethod = "DELETE"
	UserSchemaOperationListResponseMethodConnect UserSchemaOperationListResponseMethod = "CONNECT"
	UserSchemaOperationListResponseMethodPatch   UserSchemaOperationListResponseMethod = "PATCH"
	UserSchemaOperationListResponseMethodTrace   UserSchemaOperationListResponseMethod = "TRACE"
)

func (r UserSchemaOperationListResponseMethod) IsKnown() bool {
	switch r {
	case UserSchemaOperationListResponseMethodGet, UserSchemaOperationListResponseMethodPost, UserSchemaOperationListResponseMethodHead, UserSchemaOperationListResponseMethodOptions, UserSchemaOperationListResponseMethodPut, UserSchemaOperationListResponseMethodDelete, UserSchemaOperationListResponseMethodConnect, UserSchemaOperationListResponseMethodPatch, UserSchemaOperationListResponseMethodTrace:
		return true
	}
	return false
}

type UserSchemaOperationListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Filter results to only include endpoints containing this pattern.
	Endpoint param.Field[string] `query:"endpoint"`
	// Add feature(s) to the results. The feature name that is given here corresponds
	// to the resulting feature object. Have a look at the top-level object description
	// for more details on the specific meaning.
	Feature param.Field[[]UserSchemaOperationListParamsFeature] `query:"feature"`
	// Filter results to only include the specified hosts.
	Host param.Field[[]string] `query:"host"`
	// Filter results to only include the specified HTTP methods.
	Method param.Field[[]string] `query:"method"`
	// Filter results by whether operations exist in API Shield Endpoint Management or
	// not. `new` will just return operations from the schema that do not exist in API
	// Shield Endpoint Management. `existing` will just return operations from the
	// schema that already exist in API Shield Endpoint Management.
	OperationStatus param.Field[UserSchemaOperationListParamsOperationStatus] `query:"operation_status"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [UserSchemaOperationListParams]'s query parameters as
// `url.Values`.
func (r UserSchemaOperationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserSchemaOperationListParamsFeature string

const (
	UserSchemaOperationListParamsFeatureThresholds       UserSchemaOperationListParamsFeature = "thresholds"
	UserSchemaOperationListParamsFeatureParameterSchemas UserSchemaOperationListParamsFeature = "parameter_schemas"
	UserSchemaOperationListParamsFeatureSchemaInfo       UserSchemaOperationListParamsFeature = "schema_info"
)

func (r UserSchemaOperationListParamsFeature) IsKnown() bool {
	switch r {
	case UserSchemaOperationListParamsFeatureThresholds, UserSchemaOperationListParamsFeatureParameterSchemas, UserSchemaOperationListParamsFeatureSchemaInfo:
		return true
	}
	return false
}

// Filter results by whether operations exist in API Shield Endpoint Management or
// not. `new` will just return operations from the schema that do not exist in API
// Shield Endpoint Management. `existing` will just return operations from the
// schema that already exist in API Shield Endpoint Management.
type UserSchemaOperationListParamsOperationStatus string

const (
	UserSchemaOperationListParamsOperationStatusNew      UserSchemaOperationListParamsOperationStatus = "new"
	UserSchemaOperationListParamsOperationStatusExisting UserSchemaOperationListParamsOperationStatus = "existing"
)

func (r UserSchemaOperationListParamsOperationStatus) IsKnown() bool {
	switch r {
	case UserSchemaOperationListParamsOperationStatusNew, UserSchemaOperationListParamsOperationStatusExisting:
		return true
	}
	return false
}
