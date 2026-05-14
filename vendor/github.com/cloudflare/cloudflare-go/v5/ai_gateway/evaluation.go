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

// EvaluationService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvaluationService] method instead.
type EvaluationService struct {
	Options []option.RequestOption
}

// NewEvaluationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvaluationService(opts ...option.RequestOption) (r *EvaluationService) {
	r = &EvaluationService{}
	r.Options = opts
	return
}

// Create a new Evaluation
func (r *EvaluationService) New(ctx context.Context, gatewayID string, params EvaluationNewParams, opts ...option.RequestOption) (res *EvaluationNewResponse, err error) {
	var env EvaluationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations", params.AccountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Evaluations
func (r *EvaluationService) List(ctx context.Context, gatewayID string, params EvaluationListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[EvaluationListResponse], err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations", params.AccountID, gatewayID)
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

// List Evaluations
func (r *EvaluationService) ListAutoPaging(ctx context.Context, gatewayID string, params EvaluationListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[EvaluationListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, gatewayID, params, opts...))
}

// Delete a Evaluation
func (r *EvaluationService) Delete(ctx context.Context, gatewayID string, id string, body EvaluationDeleteParams, opts ...option.RequestOption) (res *EvaluationDeleteResponse, err error) {
	var env EvaluationDeleteResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations/%s", body.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a Evaluation
func (r *EvaluationService) Get(ctx context.Context, gatewayID string, id string, query EvaluationGetParams, opts ...option.RequestOption) (res *EvaluationGetResponse, err error) {
	var env EvaluationGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/evaluations/%s", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EvaluationNewResponse struct {
	ID         string                         `json:"id,required"`
	AccountID  string                         `json:"account_id,required"`
	AccountTag string                         `json:"account_tag,required"`
	CreatedAt  time.Time                      `json:"created_at,required" format:"date-time"`
	Datasets   []EvaluationNewResponseDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                        `json:"gateway_id,required"`
	ModifiedAt time.Time                     `json:"modified_at,required" format:"date-time"`
	Name       string                        `json:"name,required"`
	Processed  bool                          `json:"processed,required"`
	Results    []EvaluationNewResponseResult `json:"results,required"`
	TotalLogs  float64                       `json:"total_logs,required"`
	JSON       evaluationNewResponseJSON     `json:"-"`
}

// evaluationNewResponseJSON contains the JSON metadata for the struct
// [EvaluationNewResponse]
type evaluationNewResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationNewResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationNewResponseDataset struct {
	ID         string                                `json:"id,required"`
	AccountID  string                                `json:"account_id,required"`
	AccountTag string                                `json:"account_tag,required"`
	CreatedAt  time.Time                             `json:"created_at,required" format:"date-time"`
	Enable     bool                                  `json:"enable,required"`
	Filters    []EvaluationNewResponseDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                           `json:"gateway_id,required"`
	ModifiedAt time.Time                        `json:"modified_at,required" format:"date-time"`
	Name       string                           `json:"name,required"`
	JSON       evaluationNewResponseDatasetJSON `json:"-"`
}

// evaluationNewResponseDatasetJSON contains the JSON metadata for the struct
// [EvaluationNewResponseDataset]
type evaluationNewResponseDatasetJSON struct {
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

func (r *EvaluationNewResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationNewResponseDatasetJSON) RawJSON() string {
	return r.raw
}

type EvaluationNewResponseDatasetsFilter struct {
	Key      EvaluationNewResponseDatasetsFiltersKey          `json:"key,required"`
	Operator EvaluationNewResponseDatasetsFiltersOperator     `json:"operator,required"`
	Value    []EvaluationNewResponseDatasetsFiltersValueUnion `json:"value,required"`
	JSON     evaluationNewResponseDatasetsFilterJSON          `json:"-"`
}

// evaluationNewResponseDatasetsFilterJSON contains the JSON metadata for the
// struct [EvaluationNewResponseDatasetsFilter]
type evaluationNewResponseDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationNewResponseDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationNewResponseDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type EvaluationNewResponseDatasetsFiltersKey string

const (
	EvaluationNewResponseDatasetsFiltersKeyCreatedAt           EvaluationNewResponseDatasetsFiltersKey = "created_at"
	EvaluationNewResponseDatasetsFiltersKeyRequestContentType  EvaluationNewResponseDatasetsFiltersKey = "request_content_type"
	EvaluationNewResponseDatasetsFiltersKeyResponseContentType EvaluationNewResponseDatasetsFiltersKey = "response_content_type"
	EvaluationNewResponseDatasetsFiltersKeySuccess             EvaluationNewResponseDatasetsFiltersKey = "success"
	EvaluationNewResponseDatasetsFiltersKeyCached              EvaluationNewResponseDatasetsFiltersKey = "cached"
	EvaluationNewResponseDatasetsFiltersKeyProvider            EvaluationNewResponseDatasetsFiltersKey = "provider"
	EvaluationNewResponseDatasetsFiltersKeyModel               EvaluationNewResponseDatasetsFiltersKey = "model"
	EvaluationNewResponseDatasetsFiltersKeyCost                EvaluationNewResponseDatasetsFiltersKey = "cost"
	EvaluationNewResponseDatasetsFiltersKeyTokens              EvaluationNewResponseDatasetsFiltersKey = "tokens"
	EvaluationNewResponseDatasetsFiltersKeyTokensIn            EvaluationNewResponseDatasetsFiltersKey = "tokens_in"
	EvaluationNewResponseDatasetsFiltersKeyTokensOut           EvaluationNewResponseDatasetsFiltersKey = "tokens_out"
	EvaluationNewResponseDatasetsFiltersKeyDuration            EvaluationNewResponseDatasetsFiltersKey = "duration"
	EvaluationNewResponseDatasetsFiltersKeyFeedback            EvaluationNewResponseDatasetsFiltersKey = "feedback"
)

func (r EvaluationNewResponseDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case EvaluationNewResponseDatasetsFiltersKeyCreatedAt, EvaluationNewResponseDatasetsFiltersKeyRequestContentType, EvaluationNewResponseDatasetsFiltersKeyResponseContentType, EvaluationNewResponseDatasetsFiltersKeySuccess, EvaluationNewResponseDatasetsFiltersKeyCached, EvaluationNewResponseDatasetsFiltersKeyProvider, EvaluationNewResponseDatasetsFiltersKeyModel, EvaluationNewResponseDatasetsFiltersKeyCost, EvaluationNewResponseDatasetsFiltersKeyTokens, EvaluationNewResponseDatasetsFiltersKeyTokensIn, EvaluationNewResponseDatasetsFiltersKeyTokensOut, EvaluationNewResponseDatasetsFiltersKeyDuration, EvaluationNewResponseDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type EvaluationNewResponseDatasetsFiltersOperator string

const (
	EvaluationNewResponseDatasetsFiltersOperatorEq       EvaluationNewResponseDatasetsFiltersOperator = "eq"
	EvaluationNewResponseDatasetsFiltersOperatorContains EvaluationNewResponseDatasetsFiltersOperator = "contains"
	EvaluationNewResponseDatasetsFiltersOperatorLt       EvaluationNewResponseDatasetsFiltersOperator = "lt"
	EvaluationNewResponseDatasetsFiltersOperatorGt       EvaluationNewResponseDatasetsFiltersOperator = "gt"
)

func (r EvaluationNewResponseDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case EvaluationNewResponseDatasetsFiltersOperatorEq, EvaluationNewResponseDatasetsFiltersOperatorContains, EvaluationNewResponseDatasetsFiltersOperatorLt, EvaluationNewResponseDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type EvaluationNewResponseDatasetsFiltersValueUnion interface {
	ImplementsEvaluationNewResponseDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluationNewResponseDatasetsFiltersValueUnion)(nil)).Elem(),
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

type EvaluationNewResponseResult struct {
	ID                string                          `json:"id,required"`
	CreatedAt         time.Time                       `json:"created_at,required" format:"date-time"`
	EvaluationID      string                          `json:"evaluation_id,required"`
	EvaluationTypeID  string                          `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                       `json:"modified_at,required" format:"date-time"`
	Result            string                          `json:"result,required"`
	Status            float64                         `json:"status,required"`
	StatusDescription string                          `json:"status_description,required"`
	TotalLogs         float64                         `json:"total_logs,required"`
	JSON              evaluationNewResponseResultJSON `json:"-"`
}

// evaluationNewResponseResultJSON contains the JSON metadata for the struct
// [EvaluationNewResponseResult]
type evaluationNewResponseResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EvaluationNewResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationNewResponseResultJSON) RawJSON() string {
	return r.raw
}

type EvaluationListResponse struct {
	ID         string                          `json:"id,required"`
	AccountID  string                          `json:"account_id,required"`
	AccountTag string                          `json:"account_tag,required"`
	CreatedAt  time.Time                       `json:"created_at,required" format:"date-time"`
	Datasets   []EvaluationListResponseDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                         `json:"gateway_id,required"`
	ModifiedAt time.Time                      `json:"modified_at,required" format:"date-time"`
	Name       string                         `json:"name,required"`
	Processed  bool                           `json:"processed,required"`
	Results    []EvaluationListResponseResult `json:"results,required"`
	TotalLogs  float64                        `json:"total_logs,required"`
	JSON       evaluationListResponseJSON     `json:"-"`
}

// evaluationListResponseJSON contains the JSON metadata for the struct
// [EvaluationListResponse]
type evaluationListResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationListResponseDataset struct {
	ID         string                                 `json:"id,required"`
	AccountID  string                                 `json:"account_id,required"`
	AccountTag string                                 `json:"account_tag,required"`
	CreatedAt  time.Time                              `json:"created_at,required" format:"date-time"`
	Enable     bool                                   `json:"enable,required"`
	Filters    []EvaluationListResponseDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                            `json:"gateway_id,required"`
	ModifiedAt time.Time                         `json:"modified_at,required" format:"date-time"`
	Name       string                            `json:"name,required"`
	JSON       evaluationListResponseDatasetJSON `json:"-"`
}

// evaluationListResponseDatasetJSON contains the JSON metadata for the struct
// [EvaluationListResponseDataset]
type evaluationListResponseDatasetJSON struct {
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

func (r *EvaluationListResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseDatasetJSON) RawJSON() string {
	return r.raw
}

type EvaluationListResponseDatasetsFilter struct {
	Key      EvaluationListResponseDatasetsFiltersKey          `json:"key,required"`
	Operator EvaluationListResponseDatasetsFiltersOperator     `json:"operator,required"`
	Value    []EvaluationListResponseDatasetsFiltersValueUnion `json:"value,required"`
	JSON     evaluationListResponseDatasetsFilterJSON          `json:"-"`
}

// evaluationListResponseDatasetsFilterJSON contains the JSON metadata for the
// struct [EvaluationListResponseDatasetsFilter]
type evaluationListResponseDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationListResponseDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type EvaluationListResponseDatasetsFiltersKey string

const (
	EvaluationListResponseDatasetsFiltersKeyCreatedAt           EvaluationListResponseDatasetsFiltersKey = "created_at"
	EvaluationListResponseDatasetsFiltersKeyRequestContentType  EvaluationListResponseDatasetsFiltersKey = "request_content_type"
	EvaluationListResponseDatasetsFiltersKeyResponseContentType EvaluationListResponseDatasetsFiltersKey = "response_content_type"
	EvaluationListResponseDatasetsFiltersKeySuccess             EvaluationListResponseDatasetsFiltersKey = "success"
	EvaluationListResponseDatasetsFiltersKeyCached              EvaluationListResponseDatasetsFiltersKey = "cached"
	EvaluationListResponseDatasetsFiltersKeyProvider            EvaluationListResponseDatasetsFiltersKey = "provider"
	EvaluationListResponseDatasetsFiltersKeyModel               EvaluationListResponseDatasetsFiltersKey = "model"
	EvaluationListResponseDatasetsFiltersKeyCost                EvaluationListResponseDatasetsFiltersKey = "cost"
	EvaluationListResponseDatasetsFiltersKeyTokens              EvaluationListResponseDatasetsFiltersKey = "tokens"
	EvaluationListResponseDatasetsFiltersKeyTokensIn            EvaluationListResponseDatasetsFiltersKey = "tokens_in"
	EvaluationListResponseDatasetsFiltersKeyTokensOut           EvaluationListResponseDatasetsFiltersKey = "tokens_out"
	EvaluationListResponseDatasetsFiltersKeyDuration            EvaluationListResponseDatasetsFiltersKey = "duration"
	EvaluationListResponseDatasetsFiltersKeyFeedback            EvaluationListResponseDatasetsFiltersKey = "feedback"
)

func (r EvaluationListResponseDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case EvaluationListResponseDatasetsFiltersKeyCreatedAt, EvaluationListResponseDatasetsFiltersKeyRequestContentType, EvaluationListResponseDatasetsFiltersKeyResponseContentType, EvaluationListResponseDatasetsFiltersKeySuccess, EvaluationListResponseDatasetsFiltersKeyCached, EvaluationListResponseDatasetsFiltersKeyProvider, EvaluationListResponseDatasetsFiltersKeyModel, EvaluationListResponseDatasetsFiltersKeyCost, EvaluationListResponseDatasetsFiltersKeyTokens, EvaluationListResponseDatasetsFiltersKeyTokensIn, EvaluationListResponseDatasetsFiltersKeyTokensOut, EvaluationListResponseDatasetsFiltersKeyDuration, EvaluationListResponseDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type EvaluationListResponseDatasetsFiltersOperator string

const (
	EvaluationListResponseDatasetsFiltersOperatorEq       EvaluationListResponseDatasetsFiltersOperator = "eq"
	EvaluationListResponseDatasetsFiltersOperatorContains EvaluationListResponseDatasetsFiltersOperator = "contains"
	EvaluationListResponseDatasetsFiltersOperatorLt       EvaluationListResponseDatasetsFiltersOperator = "lt"
	EvaluationListResponseDatasetsFiltersOperatorGt       EvaluationListResponseDatasetsFiltersOperator = "gt"
)

func (r EvaluationListResponseDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case EvaluationListResponseDatasetsFiltersOperatorEq, EvaluationListResponseDatasetsFiltersOperatorContains, EvaluationListResponseDatasetsFiltersOperatorLt, EvaluationListResponseDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type EvaluationListResponseDatasetsFiltersValueUnion interface {
	ImplementsEvaluationListResponseDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluationListResponseDatasetsFiltersValueUnion)(nil)).Elem(),
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

type EvaluationListResponseResult struct {
	ID                string                           `json:"id,required"`
	CreatedAt         time.Time                        `json:"created_at,required" format:"date-time"`
	EvaluationID      string                           `json:"evaluation_id,required"`
	EvaluationTypeID  string                           `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                        `json:"modified_at,required" format:"date-time"`
	Result            string                           `json:"result,required"`
	Status            float64                          `json:"status,required"`
	StatusDescription string                           `json:"status_description,required"`
	TotalLogs         float64                          `json:"total_logs,required"`
	JSON              evaluationListResponseResultJSON `json:"-"`
}

// evaluationListResponseResultJSON contains the JSON metadata for the struct
// [EvaluationListResponseResult]
type evaluationListResponseResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EvaluationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationListResponseResultJSON) RawJSON() string {
	return r.raw
}

type EvaluationDeleteResponse struct {
	ID         string                            `json:"id,required"`
	AccountID  string                            `json:"account_id,required"`
	AccountTag string                            `json:"account_tag,required"`
	CreatedAt  time.Time                         `json:"created_at,required" format:"date-time"`
	Datasets   []EvaluationDeleteResponseDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                           `json:"gateway_id,required"`
	ModifiedAt time.Time                        `json:"modified_at,required" format:"date-time"`
	Name       string                           `json:"name,required"`
	Processed  bool                             `json:"processed,required"`
	Results    []EvaluationDeleteResponseResult `json:"results,required"`
	TotalLogs  float64                          `json:"total_logs,required"`
	JSON       evaluationDeleteResponseJSON     `json:"-"`
}

// evaluationDeleteResponseJSON contains the JSON metadata for the struct
// [EvaluationDeleteResponse]
type evaluationDeleteResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationDeleteResponseDataset struct {
	ID         string                                   `json:"id,required"`
	AccountID  string                                   `json:"account_id,required"`
	AccountTag string                                   `json:"account_tag,required"`
	CreatedAt  time.Time                                `json:"created_at,required" format:"date-time"`
	Enable     bool                                     `json:"enable,required"`
	Filters    []EvaluationDeleteResponseDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                              `json:"gateway_id,required"`
	ModifiedAt time.Time                           `json:"modified_at,required" format:"date-time"`
	Name       string                              `json:"name,required"`
	JSON       evaluationDeleteResponseDatasetJSON `json:"-"`
}

// evaluationDeleteResponseDatasetJSON contains the JSON metadata for the struct
// [EvaluationDeleteResponseDataset]
type evaluationDeleteResponseDatasetJSON struct {
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

func (r *EvaluationDeleteResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationDeleteResponseDatasetJSON) RawJSON() string {
	return r.raw
}

type EvaluationDeleteResponseDatasetsFilter struct {
	Key      EvaluationDeleteResponseDatasetsFiltersKey          `json:"key,required"`
	Operator EvaluationDeleteResponseDatasetsFiltersOperator     `json:"operator,required"`
	Value    []EvaluationDeleteResponseDatasetsFiltersValueUnion `json:"value,required"`
	JSON     evaluationDeleteResponseDatasetsFilterJSON          `json:"-"`
}

// evaluationDeleteResponseDatasetsFilterJSON contains the JSON metadata for the
// struct [EvaluationDeleteResponseDatasetsFilter]
type evaluationDeleteResponseDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationDeleteResponseDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationDeleteResponseDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type EvaluationDeleteResponseDatasetsFiltersKey string

const (
	EvaluationDeleteResponseDatasetsFiltersKeyCreatedAt           EvaluationDeleteResponseDatasetsFiltersKey = "created_at"
	EvaluationDeleteResponseDatasetsFiltersKeyRequestContentType  EvaluationDeleteResponseDatasetsFiltersKey = "request_content_type"
	EvaluationDeleteResponseDatasetsFiltersKeyResponseContentType EvaluationDeleteResponseDatasetsFiltersKey = "response_content_type"
	EvaluationDeleteResponseDatasetsFiltersKeySuccess             EvaluationDeleteResponseDatasetsFiltersKey = "success"
	EvaluationDeleteResponseDatasetsFiltersKeyCached              EvaluationDeleteResponseDatasetsFiltersKey = "cached"
	EvaluationDeleteResponseDatasetsFiltersKeyProvider            EvaluationDeleteResponseDatasetsFiltersKey = "provider"
	EvaluationDeleteResponseDatasetsFiltersKeyModel               EvaluationDeleteResponseDatasetsFiltersKey = "model"
	EvaluationDeleteResponseDatasetsFiltersKeyCost                EvaluationDeleteResponseDatasetsFiltersKey = "cost"
	EvaluationDeleteResponseDatasetsFiltersKeyTokens              EvaluationDeleteResponseDatasetsFiltersKey = "tokens"
	EvaluationDeleteResponseDatasetsFiltersKeyTokensIn            EvaluationDeleteResponseDatasetsFiltersKey = "tokens_in"
	EvaluationDeleteResponseDatasetsFiltersKeyTokensOut           EvaluationDeleteResponseDatasetsFiltersKey = "tokens_out"
	EvaluationDeleteResponseDatasetsFiltersKeyDuration            EvaluationDeleteResponseDatasetsFiltersKey = "duration"
	EvaluationDeleteResponseDatasetsFiltersKeyFeedback            EvaluationDeleteResponseDatasetsFiltersKey = "feedback"
)

func (r EvaluationDeleteResponseDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case EvaluationDeleteResponseDatasetsFiltersKeyCreatedAt, EvaluationDeleteResponseDatasetsFiltersKeyRequestContentType, EvaluationDeleteResponseDatasetsFiltersKeyResponseContentType, EvaluationDeleteResponseDatasetsFiltersKeySuccess, EvaluationDeleteResponseDatasetsFiltersKeyCached, EvaluationDeleteResponseDatasetsFiltersKeyProvider, EvaluationDeleteResponseDatasetsFiltersKeyModel, EvaluationDeleteResponseDatasetsFiltersKeyCost, EvaluationDeleteResponseDatasetsFiltersKeyTokens, EvaluationDeleteResponseDatasetsFiltersKeyTokensIn, EvaluationDeleteResponseDatasetsFiltersKeyTokensOut, EvaluationDeleteResponseDatasetsFiltersKeyDuration, EvaluationDeleteResponseDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type EvaluationDeleteResponseDatasetsFiltersOperator string

const (
	EvaluationDeleteResponseDatasetsFiltersOperatorEq       EvaluationDeleteResponseDatasetsFiltersOperator = "eq"
	EvaluationDeleteResponseDatasetsFiltersOperatorContains EvaluationDeleteResponseDatasetsFiltersOperator = "contains"
	EvaluationDeleteResponseDatasetsFiltersOperatorLt       EvaluationDeleteResponseDatasetsFiltersOperator = "lt"
	EvaluationDeleteResponseDatasetsFiltersOperatorGt       EvaluationDeleteResponseDatasetsFiltersOperator = "gt"
)

func (r EvaluationDeleteResponseDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case EvaluationDeleteResponseDatasetsFiltersOperatorEq, EvaluationDeleteResponseDatasetsFiltersOperatorContains, EvaluationDeleteResponseDatasetsFiltersOperatorLt, EvaluationDeleteResponseDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type EvaluationDeleteResponseDatasetsFiltersValueUnion interface {
	ImplementsEvaluationDeleteResponseDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluationDeleteResponseDatasetsFiltersValueUnion)(nil)).Elem(),
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

type EvaluationDeleteResponseResult struct {
	ID                string                             `json:"id,required"`
	CreatedAt         time.Time                          `json:"created_at,required" format:"date-time"`
	EvaluationID      string                             `json:"evaluation_id,required"`
	EvaluationTypeID  string                             `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                          `json:"modified_at,required" format:"date-time"`
	Result            string                             `json:"result,required"`
	Status            float64                            `json:"status,required"`
	StatusDescription string                             `json:"status_description,required"`
	TotalLogs         float64                            `json:"total_logs,required"`
	JSON              evaluationDeleteResponseResultJSON `json:"-"`
}

// evaluationDeleteResponseResultJSON contains the JSON metadata for the struct
// [EvaluationDeleteResponseResult]
type evaluationDeleteResponseResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EvaluationDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationDeleteResponseResultJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetResponse struct {
	ID         string                         `json:"id,required"`
	AccountID  string                         `json:"account_id,required"`
	AccountTag string                         `json:"account_tag,required"`
	CreatedAt  time.Time                      `json:"created_at,required" format:"date-time"`
	Datasets   []EvaluationGetResponseDataset `json:"datasets,required"`
	// gateway id
	GatewayID  string                        `json:"gateway_id,required"`
	ModifiedAt time.Time                     `json:"modified_at,required" format:"date-time"`
	Name       string                        `json:"name,required"`
	Processed  bool                          `json:"processed,required"`
	Results    []EvaluationGetResponseResult `json:"results,required"`
	TotalLogs  float64                       `json:"total_logs,required"`
	JSON       evaluationGetResponseJSON     `json:"-"`
}

// evaluationGetResponseJSON contains the JSON metadata for the struct
// [EvaluationGetResponse]
type evaluationGetResponseJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	AccountTag  apijson.Field
	CreatedAt   apijson.Field
	Datasets    apijson.Field
	GatewayID   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Processed   apijson.Field
	Results     apijson.Field
	TotalLogs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetResponseDataset struct {
	ID         string                                `json:"id,required"`
	AccountID  string                                `json:"account_id,required"`
	AccountTag string                                `json:"account_tag,required"`
	CreatedAt  time.Time                             `json:"created_at,required" format:"date-time"`
	Enable     bool                                  `json:"enable,required"`
	Filters    []EvaluationGetResponseDatasetsFilter `json:"filters,required"`
	// gateway id
	GatewayID  string                           `json:"gateway_id,required"`
	ModifiedAt time.Time                        `json:"modified_at,required" format:"date-time"`
	Name       string                           `json:"name,required"`
	JSON       evaluationGetResponseDatasetJSON `json:"-"`
}

// evaluationGetResponseDatasetJSON contains the JSON metadata for the struct
// [EvaluationGetResponseDataset]
type evaluationGetResponseDatasetJSON struct {
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

func (r *EvaluationGetResponseDataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseDatasetJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetResponseDatasetsFilter struct {
	Key      EvaluationGetResponseDatasetsFiltersKey          `json:"key,required"`
	Operator EvaluationGetResponseDatasetsFiltersOperator     `json:"operator,required"`
	Value    []EvaluationGetResponseDatasetsFiltersValueUnion `json:"value,required"`
	JSON     evaluationGetResponseDatasetsFilterJSON          `json:"-"`
}

// evaluationGetResponseDatasetsFilterJSON contains the JSON metadata for the
// struct [EvaluationGetResponseDatasetsFilter]
type evaluationGetResponseDatasetsFilterJSON struct {
	Key         apijson.Field
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetResponseDatasetsFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseDatasetsFilterJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetResponseDatasetsFiltersKey string

const (
	EvaluationGetResponseDatasetsFiltersKeyCreatedAt           EvaluationGetResponseDatasetsFiltersKey = "created_at"
	EvaluationGetResponseDatasetsFiltersKeyRequestContentType  EvaluationGetResponseDatasetsFiltersKey = "request_content_type"
	EvaluationGetResponseDatasetsFiltersKeyResponseContentType EvaluationGetResponseDatasetsFiltersKey = "response_content_type"
	EvaluationGetResponseDatasetsFiltersKeySuccess             EvaluationGetResponseDatasetsFiltersKey = "success"
	EvaluationGetResponseDatasetsFiltersKeyCached              EvaluationGetResponseDatasetsFiltersKey = "cached"
	EvaluationGetResponseDatasetsFiltersKeyProvider            EvaluationGetResponseDatasetsFiltersKey = "provider"
	EvaluationGetResponseDatasetsFiltersKeyModel               EvaluationGetResponseDatasetsFiltersKey = "model"
	EvaluationGetResponseDatasetsFiltersKeyCost                EvaluationGetResponseDatasetsFiltersKey = "cost"
	EvaluationGetResponseDatasetsFiltersKeyTokens              EvaluationGetResponseDatasetsFiltersKey = "tokens"
	EvaluationGetResponseDatasetsFiltersKeyTokensIn            EvaluationGetResponseDatasetsFiltersKey = "tokens_in"
	EvaluationGetResponseDatasetsFiltersKeyTokensOut           EvaluationGetResponseDatasetsFiltersKey = "tokens_out"
	EvaluationGetResponseDatasetsFiltersKeyDuration            EvaluationGetResponseDatasetsFiltersKey = "duration"
	EvaluationGetResponseDatasetsFiltersKeyFeedback            EvaluationGetResponseDatasetsFiltersKey = "feedback"
)

func (r EvaluationGetResponseDatasetsFiltersKey) IsKnown() bool {
	switch r {
	case EvaluationGetResponseDatasetsFiltersKeyCreatedAt, EvaluationGetResponseDatasetsFiltersKeyRequestContentType, EvaluationGetResponseDatasetsFiltersKeyResponseContentType, EvaluationGetResponseDatasetsFiltersKeySuccess, EvaluationGetResponseDatasetsFiltersKeyCached, EvaluationGetResponseDatasetsFiltersKeyProvider, EvaluationGetResponseDatasetsFiltersKeyModel, EvaluationGetResponseDatasetsFiltersKeyCost, EvaluationGetResponseDatasetsFiltersKeyTokens, EvaluationGetResponseDatasetsFiltersKeyTokensIn, EvaluationGetResponseDatasetsFiltersKeyTokensOut, EvaluationGetResponseDatasetsFiltersKeyDuration, EvaluationGetResponseDatasetsFiltersKeyFeedback:
		return true
	}
	return false
}

type EvaluationGetResponseDatasetsFiltersOperator string

const (
	EvaluationGetResponseDatasetsFiltersOperatorEq       EvaluationGetResponseDatasetsFiltersOperator = "eq"
	EvaluationGetResponseDatasetsFiltersOperatorContains EvaluationGetResponseDatasetsFiltersOperator = "contains"
	EvaluationGetResponseDatasetsFiltersOperatorLt       EvaluationGetResponseDatasetsFiltersOperator = "lt"
	EvaluationGetResponseDatasetsFiltersOperatorGt       EvaluationGetResponseDatasetsFiltersOperator = "gt"
)

func (r EvaluationGetResponseDatasetsFiltersOperator) IsKnown() bool {
	switch r {
	case EvaluationGetResponseDatasetsFiltersOperatorEq, EvaluationGetResponseDatasetsFiltersOperatorContains, EvaluationGetResponseDatasetsFiltersOperatorLt, EvaluationGetResponseDatasetsFiltersOperatorGt:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type EvaluationGetResponseDatasetsFiltersValueUnion interface {
	ImplementsEvaluationGetResponseDatasetsFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*EvaluationGetResponseDatasetsFiltersValueUnion)(nil)).Elem(),
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

type EvaluationGetResponseResult struct {
	ID                string                          `json:"id,required"`
	CreatedAt         time.Time                       `json:"created_at,required" format:"date-time"`
	EvaluationID      string                          `json:"evaluation_id,required"`
	EvaluationTypeID  string                          `json:"evaluation_type_id,required"`
	ModifiedAt        time.Time                       `json:"modified_at,required" format:"date-time"`
	Result            string                          `json:"result,required"`
	Status            float64                         `json:"status,required"`
	StatusDescription string                          `json:"status_description,required"`
	TotalLogs         float64                         `json:"total_logs,required"`
	JSON              evaluationGetResponseResultJSON `json:"-"`
}

// evaluationGetResponseResultJSON contains the JSON metadata for the struct
// [EvaluationGetResponseResult]
type evaluationGetResponseResultJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	EvaluationID      apijson.Field
	EvaluationTypeID  apijson.Field
	ModifiedAt        apijson.Field
	Result            apijson.Field
	Status            apijson.Field
	StatusDescription apijson.Field
	TotalLogs         apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *EvaluationGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseResultJSON) RawJSON() string {
	return r.raw
}

type EvaluationNewParams struct {
	AccountID         param.Field[string]   `path:"account_id,required"`
	DatasetIDs        param.Field[[]string] `json:"dataset_ids,required"`
	EvaluationTypeIDs param.Field[[]string] `json:"evaluation_type_ids,required"`
	Name              param.Field[string]   `json:"name,required"`
}

func (r EvaluationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EvaluationNewResponseEnvelope struct {
	Result  EvaluationNewResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    evaluationNewResponseEnvelopeJSON `json:"-"`
}

// evaluationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [EvaluationNewResponseEnvelope]
type evaluationNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EvaluationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Name      param.Field[string] `query:"name"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	Processed param.Field[bool]   `query:"processed"`
	// Search by id, name
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [EvaluationListParams]'s query parameters as `url.Values`.
func (r EvaluationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EvaluationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type EvaluationDeleteResponseEnvelope struct {
	Result  EvaluationDeleteResponse             `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    evaluationDeleteResponseEnvelopeJSON `json:"-"`
}

// evaluationDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [EvaluationDeleteResponseEnvelope]
type evaluationDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EvaluationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type EvaluationGetResponseEnvelope struct {
	Result  EvaluationGetResponse             `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    evaluationGetResponseEnvelopeJSON `json:"-"`
}

// evaluationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [EvaluationGetResponseEnvelope]
type evaluationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
