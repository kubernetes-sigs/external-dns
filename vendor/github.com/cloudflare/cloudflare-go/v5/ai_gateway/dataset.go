// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// DatasetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetService] method instead.
type DatasetService struct {
	Options []option.RequestOption
}

// NewDatasetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDatasetService(opts ...option.RequestOption) (r *DatasetService) {
	r = &DatasetService{}
	r.Options = opts
	return
}

// Create a new Dataset
func (r *DatasetService) New(ctx context.Context, gatewayID string, params DatasetNewParams, opts ...option.RequestOption) (res *DatasetNewResponse, err error) {
	var env DatasetNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets", params.AccountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Dataset
func (r *DatasetService) Update(ctx context.Context, gatewayID string, id string, params DatasetUpdateParams, opts ...option.RequestOption) (res *DatasetUpdateResponse, err error) {
	var env DatasetUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets/%s", params.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Datasets
func (r *DatasetService) List(ctx context.Context, gatewayID string, params DatasetListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DatasetListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets", params.AccountID, gatewayID)
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

// List Datasets
func (r *DatasetService) ListAutoPaging(ctx context.Context, gatewayID string, params DatasetListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DatasetListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, gatewayID, params, opts...))
}

// Delete a Dataset
func (r *DatasetService) Delete(ctx context.Context, gatewayID string, id string, body DatasetDeleteParams, opts ...option.RequestOption) (res *DatasetDeleteResponse, err error) {
	var env DatasetDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets/%s", body.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a Dataset
func (r *DatasetService) Get(ctx context.Context, gatewayID string, id string, query DatasetGetParams, opts ...option.RequestOption) (res *DatasetGetResponse, err error) {
	var env DatasetGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/datasets/%s", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DatasetNewResponse struct {
	ID         string                     `json:"id,required"`
	AccountID  string                     `json:"account_id,required"`
	AccountTag string                     `json:"account_tag,required"`
	CreatedAt  time.Time                  `json:"created_at,required" format:"date-time"`
	Enable     bool                       `json:"enable,required"`
	Filters    []DatasetNewResponseFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                 `json:"gateway_id,required"`
	ModifiedAt time.Time              `json:"modified_at,required" format:"date-time"`
	Name       string                 `json:"name,required"`
	JSON       datasetNewResponseJSON `json:"-"`
}

// datasetNewResponseJSON contains the JSON metadata for the struct
// [DatasetNewResponse]
type datasetNewResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetNewResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetNewResponseFilter struct {
	Key      DatasetNewResponseFiltersKey          `json:"key,required"`
	Operator DatasetNewResponseFiltersOperator     `json:"operator,required"`
	Value    []DatasetNewResponseFiltersValueUnion `json:"value,required"`
	JSON     datasetNewResponseFilterJSON          `json:"-"`
}

// datasetNewResponseFilterJSON contains the JSON metadata for the struct
// [DatasetNewResponseFilter]
type datasetNewResponseFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetNewResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetNewResponseFilterJSON) RawJSON() string {
	return r.raw
}

type DatasetNewResponseFiltersKey string

const (
	DatasetNewResponseFiltersKeyCreatedAt           DatasetNewResponseFiltersKey = "created_at"
	DatasetNewResponseFiltersKeyRequestContentType  DatasetNewResponseFiltersKey = "request_content_type"
	DatasetNewResponseFiltersKeyResponseContentType DatasetNewResponseFiltersKey = "response_content_type"
	DatasetNewResponseFiltersKeySuccess             DatasetNewResponseFiltersKey = "success"
	DatasetNewResponseFiltersKeyCached              DatasetNewResponseFiltersKey = "cached"
	DatasetNewResponseFiltersKeyProvider            DatasetNewResponseFiltersKey = "provider"
	DatasetNewResponseFiltersKeyModel               DatasetNewResponseFiltersKey = "model"
	DatasetNewResponseFiltersKeyCost                DatasetNewResponseFiltersKey = "cost"
	DatasetNewResponseFiltersKeyTokens              DatasetNewResponseFiltersKey = "tokens"
	DatasetNewResponseFiltersKeyTokensIn            DatasetNewResponseFiltersKey = "tokens_in"
	DatasetNewResponseFiltersKeyTokensOut           DatasetNewResponseFiltersKey = "tokens_out"
	DatasetNewResponseFiltersKeyDuration            DatasetNewResponseFiltersKey = "duration"
	DatasetNewResponseFiltersKeyFeedback            DatasetNewResponseFiltersKey = "feedback"
)

func (r DatasetNewResponseFiltersKey) IsKnown() bool {
	switch r {
	case DatasetNewResponseFiltersKeyCreatedAt, DatasetNewResponseFiltersKeyRequestContentType, DatasetNewResponseFiltersKeyResponseContentType, DatasetNewResponseFiltersKeySuccess, DatasetNewResponseFiltersKeyCached, DatasetNewResponseFiltersKeyProvider, DatasetNewResponseFiltersKeyModel, DatasetNewResponseFiltersKeyCost, DatasetNewResponseFiltersKeyTokens, DatasetNewResponseFiltersKeyTokensIn, DatasetNewResponseFiltersKeyTokensOut, DatasetNewResponseFiltersKeyDuration, DatasetNewResponseFiltersKeyFeedback:
		return true
	}
	return false
}

type DatasetNewResponseFiltersOperator string

const (
	DatasetNewResponseFiltersOperatorEq       DatasetNewResponseFiltersOperator = "eq"
	DatasetNewResponseFiltersOperatorContains DatasetNewResponseFiltersOperator = "contains"
	DatasetNewResponseFiltersOperatorLt       DatasetNewResponseFiltersOperator = "lt"
	DatasetNewResponseFiltersOperatorGt       DatasetNewResponseFiltersOperator = "gt"
)

func (r DatasetNewResponseFiltersOperator) IsKnown() bool {
	switch r {
	case DatasetNewResponseFiltersOperatorEq, DatasetNewResponseFiltersOperatorContains, DatasetNewResponseFiltersOperatorLt, DatasetNewResponseFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type DatasetNewResponseFiltersValueUnion interface {
	ImplementsDatasetNewResponseFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatasetNewResponseFiltersValueUnion)(nil)).Elem(),
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

type DatasetUpdateResponse struct {
	ID         string                        `json:"id,required"`
	AccountID  string                        `json:"account_id,required"`
	AccountTag string                        `json:"account_tag,required"`
	CreatedAt  time.Time                     `json:"created_at,required" format:"date-time"`
	Enable     bool                          `json:"enable,required"`
	Filters    []DatasetUpdateResponseFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                    `json:"gateway_id,required"`
	ModifiedAt time.Time                 `json:"modified_at,required" format:"date-time"`
	Name       string                    `json:"name,required"`
	JSON       datasetUpdateResponseJSON `json:"-"`
}

// datasetUpdateResponseJSON contains the JSON metadata for the struct
// [DatasetUpdateResponse]
type datasetUpdateResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetUpdateResponseFilter struct {
	Key      DatasetUpdateResponseFiltersKey          `json:"key,required"`
	Operator DatasetUpdateResponseFiltersOperator     `json:"operator,required"`
	Value    []DatasetUpdateResponseFiltersValueUnion `json:"value,required"`
	JSON     datasetUpdateResponseFilterJSON          `json:"-"`
}

// datasetUpdateResponseFilterJSON contains the JSON metadata for the struct
// [DatasetUpdateResponseFilter]
type datasetUpdateResponseFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetUpdateResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetUpdateResponseFilterJSON) RawJSON() string {
	return r.raw
}

type DatasetUpdateResponseFiltersKey string

const (
	DatasetUpdateResponseFiltersKeyCreatedAt           DatasetUpdateResponseFiltersKey = "created_at"
	DatasetUpdateResponseFiltersKeyRequestContentType  DatasetUpdateResponseFiltersKey = "request_content_type"
	DatasetUpdateResponseFiltersKeyResponseContentType DatasetUpdateResponseFiltersKey = "response_content_type"
	DatasetUpdateResponseFiltersKeySuccess             DatasetUpdateResponseFiltersKey = "success"
	DatasetUpdateResponseFiltersKeyCached              DatasetUpdateResponseFiltersKey = "cached"
	DatasetUpdateResponseFiltersKeyProvider            DatasetUpdateResponseFiltersKey = "provider"
	DatasetUpdateResponseFiltersKeyModel               DatasetUpdateResponseFiltersKey = "model"
	DatasetUpdateResponseFiltersKeyCost                DatasetUpdateResponseFiltersKey = "cost"
	DatasetUpdateResponseFiltersKeyTokens              DatasetUpdateResponseFiltersKey = "tokens"
	DatasetUpdateResponseFiltersKeyTokensIn            DatasetUpdateResponseFiltersKey = "tokens_in"
	DatasetUpdateResponseFiltersKeyTokensOut           DatasetUpdateResponseFiltersKey = "tokens_out"
	DatasetUpdateResponseFiltersKeyDuration            DatasetUpdateResponseFiltersKey = "duration"
	DatasetUpdateResponseFiltersKeyFeedback            DatasetUpdateResponseFiltersKey = "feedback"
)

func (r DatasetUpdateResponseFiltersKey) IsKnown() bool {
	switch r {
	case DatasetUpdateResponseFiltersKeyCreatedAt, DatasetUpdateResponseFiltersKeyRequestContentType, DatasetUpdateResponseFiltersKeyResponseContentType, DatasetUpdateResponseFiltersKeySuccess, DatasetUpdateResponseFiltersKeyCached, DatasetUpdateResponseFiltersKeyProvider, DatasetUpdateResponseFiltersKeyModel, DatasetUpdateResponseFiltersKeyCost, DatasetUpdateResponseFiltersKeyTokens, DatasetUpdateResponseFiltersKeyTokensIn, DatasetUpdateResponseFiltersKeyTokensOut, DatasetUpdateResponseFiltersKeyDuration, DatasetUpdateResponseFiltersKeyFeedback:
		return true
	}
	return false
}

type DatasetUpdateResponseFiltersOperator string

const (
	DatasetUpdateResponseFiltersOperatorEq       DatasetUpdateResponseFiltersOperator = "eq"
	DatasetUpdateResponseFiltersOperatorContains DatasetUpdateResponseFiltersOperator = "contains"
	DatasetUpdateResponseFiltersOperatorLt       DatasetUpdateResponseFiltersOperator = "lt"
	DatasetUpdateResponseFiltersOperatorGt       DatasetUpdateResponseFiltersOperator = "gt"
)

func (r DatasetUpdateResponseFiltersOperator) IsKnown() bool {
	switch r {
	case DatasetUpdateResponseFiltersOperatorEq, DatasetUpdateResponseFiltersOperatorContains, DatasetUpdateResponseFiltersOperatorLt, DatasetUpdateResponseFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type DatasetUpdateResponseFiltersValueUnion interface {
	ImplementsDatasetUpdateResponseFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatasetUpdateResponseFiltersValueUnion)(nil)).Elem(),
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

type DatasetListResponse struct {
	ID         string                      `json:"id,required"`
	AccountID  string                      `json:"account_id,required"`
	AccountTag string                      `json:"account_tag,required"`
	CreatedAt  time.Time                   `json:"created_at,required" format:"date-time"`
	Enable     bool                        `json:"enable,required"`
	Filters    []DatasetListResponseFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                  `json:"gateway_id,required"`
	ModifiedAt time.Time               `json:"modified_at,required" format:"date-time"`
	Name       string                  `json:"name,required"`
	JSON       datasetListResponseJSON `json:"-"`
}

// datasetListResponseJSON contains the JSON metadata for the struct
// [DatasetListResponse]
type datasetListResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetListResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetListResponseFilter struct {
	Key      DatasetListResponseFiltersKey          `json:"key,required"`
	Operator DatasetListResponseFiltersOperator     `json:"operator,required"`
	Value    []DatasetListResponseFiltersValueUnion `json:"value,required"`
	JSON     datasetListResponseFilterJSON          `json:"-"`
}

// datasetListResponseFilterJSON contains the JSON metadata for the struct
// [DatasetListResponseFilter]
type datasetListResponseFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetListResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetListResponseFilterJSON) RawJSON() string {
	return r.raw
}

type DatasetListResponseFiltersKey string

const (
	DatasetListResponseFiltersKeyCreatedAt           DatasetListResponseFiltersKey = "created_at"
	DatasetListResponseFiltersKeyRequestContentType  DatasetListResponseFiltersKey = "request_content_type"
	DatasetListResponseFiltersKeyResponseContentType DatasetListResponseFiltersKey = "response_content_type"
	DatasetListResponseFiltersKeySuccess             DatasetListResponseFiltersKey = "success"
	DatasetListResponseFiltersKeyCached              DatasetListResponseFiltersKey = "cached"
	DatasetListResponseFiltersKeyProvider            DatasetListResponseFiltersKey = "provider"
	DatasetListResponseFiltersKeyModel               DatasetListResponseFiltersKey = "model"
	DatasetListResponseFiltersKeyCost                DatasetListResponseFiltersKey = "cost"
	DatasetListResponseFiltersKeyTokens              DatasetListResponseFiltersKey = "tokens"
	DatasetListResponseFiltersKeyTokensIn            DatasetListResponseFiltersKey = "tokens_in"
	DatasetListResponseFiltersKeyTokensOut           DatasetListResponseFiltersKey = "tokens_out"
	DatasetListResponseFiltersKeyDuration            DatasetListResponseFiltersKey = "duration"
	DatasetListResponseFiltersKeyFeedback            DatasetListResponseFiltersKey = "feedback"
)

func (r DatasetListResponseFiltersKey) IsKnown() bool {
	switch r {
	case DatasetListResponseFiltersKeyCreatedAt, DatasetListResponseFiltersKeyRequestContentType, DatasetListResponseFiltersKeyResponseContentType, DatasetListResponseFiltersKeySuccess, DatasetListResponseFiltersKeyCached, DatasetListResponseFiltersKeyProvider, DatasetListResponseFiltersKeyModel, DatasetListResponseFiltersKeyCost, DatasetListResponseFiltersKeyTokens, DatasetListResponseFiltersKeyTokensIn, DatasetListResponseFiltersKeyTokensOut, DatasetListResponseFiltersKeyDuration, DatasetListResponseFiltersKeyFeedback:
		return true
	}
	return false
}

type DatasetListResponseFiltersOperator string

const (
	DatasetListResponseFiltersOperatorEq       DatasetListResponseFiltersOperator = "eq"
	DatasetListResponseFiltersOperatorContains DatasetListResponseFiltersOperator = "contains"
	DatasetListResponseFiltersOperatorLt       DatasetListResponseFiltersOperator = "lt"
	DatasetListResponseFiltersOperatorGt       DatasetListResponseFiltersOperator = "gt"
)

func (r DatasetListResponseFiltersOperator) IsKnown() bool {
	switch r {
	case DatasetListResponseFiltersOperatorEq, DatasetListResponseFiltersOperatorContains, DatasetListResponseFiltersOperatorLt, DatasetListResponseFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type DatasetListResponseFiltersValueUnion interface {
	ImplementsDatasetListResponseFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatasetListResponseFiltersValueUnion)(nil)).Elem(),
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

type DatasetDeleteResponse struct {
	ID         string                        `json:"id,required"`
	AccountID  string                        `json:"account_id,required"`
	AccountTag string                        `json:"account_tag,required"`
	CreatedAt  time.Time                     `json:"created_at,required" format:"date-time"`
	Enable     bool                          `json:"enable,required"`
	Filters    []DatasetDeleteResponseFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                    `json:"gateway_id,required"`
	ModifiedAt time.Time                 `json:"modified_at,required" format:"date-time"`
	Name       string                    `json:"name,required"`
	JSON       datasetDeleteResponseJSON `json:"-"`
}

// datasetDeleteResponseJSON contains the JSON metadata for the struct
// [DatasetDeleteResponse]
type datasetDeleteResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetDeleteResponseFilter struct {
	Key      DatasetDeleteResponseFiltersKey          `json:"key,required"`
	Operator DatasetDeleteResponseFiltersOperator     `json:"operator,required"`
	Value    []DatasetDeleteResponseFiltersValueUnion `json:"value,required"`
	JSON     datasetDeleteResponseFilterJSON          `json:"-"`
}

// datasetDeleteResponseFilterJSON contains the JSON metadata for the struct
// [DatasetDeleteResponseFilter]
type datasetDeleteResponseFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetDeleteResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetDeleteResponseFilterJSON) RawJSON() string {
	return r.raw
}

type DatasetDeleteResponseFiltersKey string

const (
	DatasetDeleteResponseFiltersKeyCreatedAt           DatasetDeleteResponseFiltersKey = "created_at"
	DatasetDeleteResponseFiltersKeyRequestContentType  DatasetDeleteResponseFiltersKey = "request_content_type"
	DatasetDeleteResponseFiltersKeyResponseContentType DatasetDeleteResponseFiltersKey = "response_content_type"
	DatasetDeleteResponseFiltersKeySuccess             DatasetDeleteResponseFiltersKey = "success"
	DatasetDeleteResponseFiltersKeyCached              DatasetDeleteResponseFiltersKey = "cached"
	DatasetDeleteResponseFiltersKeyProvider            DatasetDeleteResponseFiltersKey = "provider"
	DatasetDeleteResponseFiltersKeyModel               DatasetDeleteResponseFiltersKey = "model"
	DatasetDeleteResponseFiltersKeyCost                DatasetDeleteResponseFiltersKey = "cost"
	DatasetDeleteResponseFiltersKeyTokens              DatasetDeleteResponseFiltersKey = "tokens"
	DatasetDeleteResponseFiltersKeyTokensIn            DatasetDeleteResponseFiltersKey = "tokens_in"
	DatasetDeleteResponseFiltersKeyTokensOut           DatasetDeleteResponseFiltersKey = "tokens_out"
	DatasetDeleteResponseFiltersKeyDuration            DatasetDeleteResponseFiltersKey = "duration"
	DatasetDeleteResponseFiltersKeyFeedback            DatasetDeleteResponseFiltersKey = "feedback"
)

func (r DatasetDeleteResponseFiltersKey) IsKnown() bool {
	switch r {
	case DatasetDeleteResponseFiltersKeyCreatedAt, DatasetDeleteResponseFiltersKeyRequestContentType, DatasetDeleteResponseFiltersKeyResponseContentType, DatasetDeleteResponseFiltersKeySuccess, DatasetDeleteResponseFiltersKeyCached, DatasetDeleteResponseFiltersKeyProvider, DatasetDeleteResponseFiltersKeyModel, DatasetDeleteResponseFiltersKeyCost, DatasetDeleteResponseFiltersKeyTokens, DatasetDeleteResponseFiltersKeyTokensIn, DatasetDeleteResponseFiltersKeyTokensOut, DatasetDeleteResponseFiltersKeyDuration, DatasetDeleteResponseFiltersKeyFeedback:
		return true
	}
	return false
}

type DatasetDeleteResponseFiltersOperator string

const (
	DatasetDeleteResponseFiltersOperatorEq       DatasetDeleteResponseFiltersOperator = "eq"
	DatasetDeleteResponseFiltersOperatorContains DatasetDeleteResponseFiltersOperator = "contains"
	DatasetDeleteResponseFiltersOperatorLt       DatasetDeleteResponseFiltersOperator = "lt"
	DatasetDeleteResponseFiltersOperatorGt       DatasetDeleteResponseFiltersOperator = "gt"
)

func (r DatasetDeleteResponseFiltersOperator) IsKnown() bool {
	switch r {
	case DatasetDeleteResponseFiltersOperatorEq, DatasetDeleteResponseFiltersOperatorContains, DatasetDeleteResponseFiltersOperatorLt, DatasetDeleteResponseFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type DatasetDeleteResponseFiltersValueUnion interface {
	ImplementsDatasetDeleteResponseFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatasetDeleteResponseFiltersValueUnion)(nil)).Elem(),
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

type DatasetGetResponse struct {
	ID         string                     `json:"id,required"`
	AccountID  string                     `json:"account_id,required"`
	AccountTag string                     `json:"account_tag,required"`
	CreatedAt  time.Time                  `json:"created_at,required" format:"date-time"`
	Enable     bool                       `json:"enable,required"`
	Filters    []DatasetGetResponseFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                 `json:"gateway_id,required"`
	ModifiedAt time.Time              `json:"modified_at,required" format:"date-time"`
	Name       string                 `json:"name,required"`
	JSON       datasetGetResponseJSON `json:"-"`
}

// datasetGetResponseJSON contains the JSON metadata for the struct
// [DatasetGetResponse]
type datasetGetResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Enable      apijson.Field
	Filters     apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetGetResponseJSON) RawJSON() string {
	return r.raw
}

type DatasetGetResponseFilter struct {
	Key      DatasetGetResponseFiltersKey          `json:"key,required"`
	Operator DatasetGetResponseFiltersOperator     `json:"operator,required"`
	Value    []DatasetGetResponseFiltersValueUnion `json:"value,required"`
	JSON     datasetGetResponseFilterJSON          `json:"-"`
}

// datasetGetResponseFilterJSON contains the JSON metadata for the struct
// [DatasetGetResponseFilter]
type datasetGetResponseFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetGetResponseFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetGetResponseFilterJSON) RawJSON() string {
	return r.raw
}

type DatasetGetResponseFiltersKey string

const (
	DatasetGetResponseFiltersKeyCreatedAt           DatasetGetResponseFiltersKey = "created_at"
	DatasetGetResponseFiltersKeyRequestContentType  DatasetGetResponseFiltersKey = "request_content_type"
	DatasetGetResponseFiltersKeyResponseContentType DatasetGetResponseFiltersKey = "response_content_type"
	DatasetGetResponseFiltersKeySuccess             DatasetGetResponseFiltersKey = "success"
	DatasetGetResponseFiltersKeyCached              DatasetGetResponseFiltersKey = "cached"
	DatasetGetResponseFiltersKeyProvider            DatasetGetResponseFiltersKey = "provider"
	DatasetGetResponseFiltersKeyModel               DatasetGetResponseFiltersKey = "model"
	DatasetGetResponseFiltersKeyCost                DatasetGetResponseFiltersKey = "cost"
	DatasetGetResponseFiltersKeyTokens              DatasetGetResponseFiltersKey = "tokens"
	DatasetGetResponseFiltersKeyTokensIn            DatasetGetResponseFiltersKey = "tokens_in"
	DatasetGetResponseFiltersKeyTokensOut           DatasetGetResponseFiltersKey = "tokens_out"
	DatasetGetResponseFiltersKeyDuration            DatasetGetResponseFiltersKey = "duration"
	DatasetGetResponseFiltersKeyFeedback            DatasetGetResponseFiltersKey = "feedback"
)

func (r DatasetGetResponseFiltersKey) IsKnown() bool {
	switch r {
	case DatasetGetResponseFiltersKeyCreatedAt, DatasetGetResponseFiltersKeyRequestContentType, DatasetGetResponseFiltersKeyResponseContentType, DatasetGetResponseFiltersKeySuccess, DatasetGetResponseFiltersKeyCached, DatasetGetResponseFiltersKeyProvider, DatasetGetResponseFiltersKeyModel, DatasetGetResponseFiltersKeyCost, DatasetGetResponseFiltersKeyTokens, DatasetGetResponseFiltersKeyTokensIn, DatasetGetResponseFiltersKeyTokensOut, DatasetGetResponseFiltersKeyDuration, DatasetGetResponseFiltersKeyFeedback:
		return true
	}
	return false
}

type DatasetGetResponseFiltersOperator string

const (
	DatasetGetResponseFiltersOperatorEq       DatasetGetResponseFiltersOperator = "eq"
	DatasetGetResponseFiltersOperatorContains DatasetGetResponseFiltersOperator = "contains"
	DatasetGetResponseFiltersOperatorLt       DatasetGetResponseFiltersOperator = "lt"
	DatasetGetResponseFiltersOperatorGt       DatasetGetResponseFiltersOperator = "gt"
)

func (r DatasetGetResponseFiltersOperator) IsKnown() bool {
	switch r {
	case DatasetGetResponseFiltersOperatorEq, DatasetGetResponseFiltersOperatorContains, DatasetGetResponseFiltersOperatorLt, DatasetGetResponseFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type DatasetGetResponseFiltersValueUnion interface {
	ImplementsDatasetGetResponseFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DatasetGetResponseFiltersValueUnion)(nil)).Elem(),
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

type DatasetNewParams struct {
	AccountID param.Field[string]                   `path:"account_id,required"`
	Enable    param.Field[bool]                     `json:"enable,required"`
	Filters   param.Field[[]DatasetNewParamsFilter] `json:"filters,required"`
	Name      param.Field[string]                   `json:"name,required"`
}

func (r DatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetNewParamsFilter struct {
	Key      param.Field[DatasetNewParamsFiltersKey]          `json:"key,required"`
	Operator param.Field[DatasetNewParamsFiltersOperator]     `json:"operator,required"`
	Value    param.Field[[]DatasetNewParamsFiltersValueUnion] `json:"value,required"`
}

func (r DatasetNewParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetNewParamsFiltersKey string

const (
	DatasetNewParamsFiltersKeyCreatedAt           DatasetNewParamsFiltersKey = "created_at"
	DatasetNewParamsFiltersKeyRequestContentType  DatasetNewParamsFiltersKey = "request_content_type"
	DatasetNewParamsFiltersKeyResponseContentType DatasetNewParamsFiltersKey = "response_content_type"
	DatasetNewParamsFiltersKeySuccess             DatasetNewParamsFiltersKey = "success"
	DatasetNewParamsFiltersKeyCached              DatasetNewParamsFiltersKey = "cached"
	DatasetNewParamsFiltersKeyProvider            DatasetNewParamsFiltersKey = "provider"
	DatasetNewParamsFiltersKeyModel               DatasetNewParamsFiltersKey = "model"
	DatasetNewParamsFiltersKeyCost                DatasetNewParamsFiltersKey = "cost"
	DatasetNewParamsFiltersKeyTokens              DatasetNewParamsFiltersKey = "tokens"
	DatasetNewParamsFiltersKeyTokensIn            DatasetNewParamsFiltersKey = "tokens_in"
	DatasetNewParamsFiltersKeyTokensOut           DatasetNewParamsFiltersKey = "tokens_out"
	DatasetNewParamsFiltersKeyDuration            DatasetNewParamsFiltersKey = "duration"
	DatasetNewParamsFiltersKeyFeedback            DatasetNewParamsFiltersKey = "feedback"
)

func (r DatasetNewParamsFiltersKey) IsKnown() bool {
	switch r {
	case DatasetNewParamsFiltersKeyCreatedAt, DatasetNewParamsFiltersKeyRequestContentType, DatasetNewParamsFiltersKeyResponseContentType, DatasetNewParamsFiltersKeySuccess, DatasetNewParamsFiltersKeyCached, DatasetNewParamsFiltersKeyProvider, DatasetNewParamsFiltersKeyModel, DatasetNewParamsFiltersKeyCost, DatasetNewParamsFiltersKeyTokens, DatasetNewParamsFiltersKeyTokensIn, DatasetNewParamsFiltersKeyTokensOut, DatasetNewParamsFiltersKeyDuration, DatasetNewParamsFiltersKeyFeedback:
		return true
	}
	return false
}

type DatasetNewParamsFiltersOperator string

const (
	DatasetNewParamsFiltersOperatorEq       DatasetNewParamsFiltersOperator = "eq"
	DatasetNewParamsFiltersOperatorContains DatasetNewParamsFiltersOperator = "contains"
	DatasetNewParamsFiltersOperatorLt       DatasetNewParamsFiltersOperator = "lt"
	DatasetNewParamsFiltersOperatorGt       DatasetNewParamsFiltersOperator = "gt"
)

func (r DatasetNewParamsFiltersOperator) IsKnown() bool {
	switch r {
	case DatasetNewParamsFiltersOperatorEq, DatasetNewParamsFiltersOperatorContains, DatasetNewParamsFiltersOperatorLt, DatasetNewParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type DatasetNewParamsFiltersValueUnion interface {
	ImplementsDatasetNewParamsFiltersValueUnion()
}

type DatasetNewResponseEnvelope struct {
	Result  DatasetNewResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    datasetNewResponseEnvelopeJSON `json:"-"`
}

// datasetNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetNewResponseEnvelope]
type datasetNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DatasetUpdateParams struct {
	AccountID param.Field[string]                      `path:"account_id,required"`
	Enable    param.Field[bool]                        `json:"enable,required"`
	Filters   param.Field[[]DatasetUpdateParamsFilter] `json:"filters,required"`
	Name      param.Field[string]                      `json:"name,required"`
}

func (r DatasetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetUpdateParamsFilter struct {
	Key      param.Field[DatasetUpdateParamsFiltersKey]          `json:"key,required"`
	Operator param.Field[DatasetUpdateParamsFiltersOperator]     `json:"operator,required"`
	Value    param.Field[[]DatasetUpdateParamsFiltersValueUnion] `json:"value,required"`
}

func (r DatasetUpdateParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetUpdateParamsFiltersKey string

const (
	DatasetUpdateParamsFiltersKeyCreatedAt           DatasetUpdateParamsFiltersKey = "created_at"
	DatasetUpdateParamsFiltersKeyRequestContentType  DatasetUpdateParamsFiltersKey = "request_content_type"
	DatasetUpdateParamsFiltersKeyResponseContentType DatasetUpdateParamsFiltersKey = "response_content_type"
	DatasetUpdateParamsFiltersKeySuccess             DatasetUpdateParamsFiltersKey = "success"
	DatasetUpdateParamsFiltersKeyCached              DatasetUpdateParamsFiltersKey = "cached"
	DatasetUpdateParamsFiltersKeyProvider            DatasetUpdateParamsFiltersKey = "provider"
	DatasetUpdateParamsFiltersKeyModel               DatasetUpdateParamsFiltersKey = "model"
	DatasetUpdateParamsFiltersKeyCost                DatasetUpdateParamsFiltersKey = "cost"
	DatasetUpdateParamsFiltersKeyTokens              DatasetUpdateParamsFiltersKey = "tokens"
	DatasetUpdateParamsFiltersKeyTokensIn            DatasetUpdateParamsFiltersKey = "tokens_in"
	DatasetUpdateParamsFiltersKeyTokensOut           DatasetUpdateParamsFiltersKey = "tokens_out"
	DatasetUpdateParamsFiltersKeyDuration            DatasetUpdateParamsFiltersKey = "duration"
	DatasetUpdateParamsFiltersKeyFeedback            DatasetUpdateParamsFiltersKey = "feedback"
)

func (r DatasetUpdateParamsFiltersKey) IsKnown() bool {
	switch r {
	case DatasetUpdateParamsFiltersKeyCreatedAt, DatasetUpdateParamsFiltersKeyRequestContentType, DatasetUpdateParamsFiltersKeyResponseContentType, DatasetUpdateParamsFiltersKeySuccess, DatasetUpdateParamsFiltersKeyCached, DatasetUpdateParamsFiltersKeyProvider, DatasetUpdateParamsFiltersKeyModel, DatasetUpdateParamsFiltersKeyCost, DatasetUpdateParamsFiltersKeyTokens, DatasetUpdateParamsFiltersKeyTokensIn, DatasetUpdateParamsFiltersKeyTokensOut, DatasetUpdateParamsFiltersKeyDuration, DatasetUpdateParamsFiltersKeyFeedback:
		return true
	}
	return false
}

type DatasetUpdateParamsFiltersOperator string

const (
	DatasetUpdateParamsFiltersOperatorEq       DatasetUpdateParamsFiltersOperator = "eq"
	DatasetUpdateParamsFiltersOperatorContains DatasetUpdateParamsFiltersOperator = "contains"
	DatasetUpdateParamsFiltersOperatorLt       DatasetUpdateParamsFiltersOperator = "lt"
	DatasetUpdateParamsFiltersOperatorGt       DatasetUpdateParamsFiltersOperator = "gt"
)

func (r DatasetUpdateParamsFiltersOperator) IsKnown() bool {
	switch r {
	case DatasetUpdateParamsFiltersOperatorEq, DatasetUpdateParamsFiltersOperatorContains, DatasetUpdateParamsFiltersOperatorLt, DatasetUpdateParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type DatasetUpdateParamsFiltersValueUnion interface {
	ImplementsDatasetUpdateParamsFiltersValueUnion()
}

type DatasetUpdateResponseEnvelope struct {
	Result  DatasetUpdateResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    datasetUpdateResponseEnvelopeJSON `json:"-"`
}

// datasetUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetUpdateResponseEnvelope]
type datasetUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DatasetListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Enable    param.Field[bool]   `query:"enable"`
	Name      param.Field[string] `query:"name"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id, name, filters
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [DatasetListParams]'s query parameters as `url.Values`.
func (r DatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DatasetDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DatasetDeleteResponseEnvelope struct {
	Result  DatasetDeleteResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    datasetDeleteResponseEnvelopeJSON `json:"-"`
}

// datasetDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetDeleteResponseEnvelope]
type datasetDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DatasetGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DatasetGetResponseEnvelope struct {
	Result  DatasetGetResponse             `json:"result,required"`
	Success bool                           `json:"success,required"`
	JSON    datasetGetResponseEnvelopeJSON `json:"-"`
}

// datasetGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DatasetGetResponseEnvelope]
type datasetGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
