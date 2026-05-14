// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// LogService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r *LogService) {
	r = &LogService{}
	r.Options = opts
	return
}

// List Gateway Logs
func (r *LogService) List(ctx context.Context, gatewayID string, params LogListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[LogListResponse], err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs", params.AccountID, gatewayID)
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

// List Gateway Logs
func (r *LogService) ListAutoPaging(ctx context.Context, gatewayID string, params LogListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[LogListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, gatewayID, params, opts...))
}

// Delete Gateway Logs
func (r *LogService) Delete(ctx context.Context, gatewayID string, params LogDeleteParams, opts ...option.RequestOption) (res *LogDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if gatewayID == "" {
		err = errors.New("missing required gateway_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs", params.AccountID, gatewayID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Patch Gateway Log
func (r *LogService) Edit(ctx context.Context, gatewayID string, id string, params LogEditParams, opts ...option.RequestOption) (res *LogEditResponse, err error) {
	var env LogEditResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s", params.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Gateway Log Detail
func (r *LogService) Get(ctx context.Context, gatewayID string, id string, query LogGetParams, opts ...option.RequestOption) (res *LogGetResponse, err error) {
	var env LogGetResponseEnvelope
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get Gateway Log Request
func (r *LogService) Request(ctx context.Context, gatewayID string, id string, query LogRequestParams, opts ...option.RequestOption) (res *LogRequestResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s/request", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get Gateway Log Response
func (r *LogService) Response(ctx context.Context, gatewayID string, id string, query LogResponseParams, opts ...option.RequestOption) (res *LogResponseResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s/response", query.AccountID, gatewayID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogListResponse struct {
	ID                  string              `json:"id,required"`
	Cached              bool                `json:"cached,required"`
	CreatedAt           time.Time           `json:"created_at,required" format:"date-time"`
	Duration            int64               `json:"duration,required"`
	Model               string              `json:"model,required"`
	Path                string              `json:"path,required"`
	Provider            string              `json:"provider,required"`
	Success             bool                `json:"success,required"`
	TokensIn            int64               `json:"tokens_in,required,nullable"`
	TokensOut           int64               `json:"tokens_out,required,nullable"`
	Cost                float64             `json:"cost"`
	CustomCost          bool                `json:"custom_cost"`
	Metadata            string              `json:"metadata"`
	ModelType           string              `json:"model_type"`
	RequestContentType  string              `json:"request_content_type"`
	RequestType         string              `json:"request_type"`
	ResponseContentType string              `json:"response_content_type"`
	StatusCode          int64               `json:"status_code"`
	Step                int64               `json:"step"`
	JSON                logListResponseJSON `json:"-"`
}

// logListResponseJSON contains the JSON metadata for the struct [LogListResponse]
type logListResponseJSON struct {
	ID                  apijson.Field
	Cached              apijson.Field
	CreatedAt           apijson.Field
	Duration            apijson.Field
	Model               apijson.Field
	Path                apijson.Field
	Provider            apijson.Field
	Success             apijson.Field
	TokensIn            apijson.Field
	TokensOut           apijson.Field
	Cost                apijson.Field
	CustomCost          apijson.Field
	Metadata            apijson.Field
	ModelType           apijson.Field
	RequestContentType  apijson.Field
	RequestType         apijson.Field
	ResponseContentType apijson.Field
	StatusCode          apijson.Field
	Step                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *LogListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logListResponseJSON) RawJSON() string {
	return r.raw
}

type LogDeleteResponse struct {
	Success bool                  `json:"success,required"`
	JSON    logDeleteResponseJSON `json:"-"`
}

// logDeleteResponseJSON contains the JSON metadata for the struct
// [LogDeleteResponse]
type logDeleteResponseJSON struct {
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type LogEditResponse = interface{}

type LogGetResponse struct {
	ID                   string             `json:"id,required"`
	Cached               bool               `json:"cached,required"`
	CreatedAt            time.Time          `json:"created_at,required" format:"date-time"`
	Duration             int64              `json:"duration,required"`
	Model                string             `json:"model,required"`
	Path                 string             `json:"path,required"`
	Provider             string             `json:"provider,required"`
	Success              bool               `json:"success,required"`
	TokensIn             int64              `json:"tokens_in,required,nullable"`
	TokensOut            int64              `json:"tokens_out,required,nullable"`
	Cost                 float64            `json:"cost"`
	CustomCost           bool               `json:"custom_cost"`
	Metadata             string             `json:"metadata"`
	ModelType            string             `json:"model_type"`
	RequestContentType   string             `json:"request_content_type"`
	RequestHead          string             `json:"request_head"`
	RequestHeadComplete  bool               `json:"request_head_complete"`
	RequestSize          int64              `json:"request_size"`
	RequestType          string             `json:"request_type"`
	ResponseContentType  string             `json:"response_content_type"`
	ResponseHead         string             `json:"response_head"`
	ResponseHeadComplete bool               `json:"response_head_complete"`
	ResponseSize         int64              `json:"response_size"`
	StatusCode           int64              `json:"status_code"`
	Step                 int64              `json:"step"`
	JSON                 logGetResponseJSON `json:"-"`
}

// logGetResponseJSON contains the JSON metadata for the struct [LogGetResponse]
type logGetResponseJSON struct {
	ID                   apijson.Field
	Cached               apijson.Field
	CreatedAt            apijson.Field
	Duration             apijson.Field
	Model                apijson.Field
	Path                 apijson.Field
	Provider             apijson.Field
	Success              apijson.Field
	TokensIn             apijson.Field
	TokensOut            apijson.Field
	Cost                 apijson.Field
	CustomCost           apijson.Field
	Metadata             apijson.Field
	ModelType            apijson.Field
	RequestContentType   apijson.Field
	RequestHead          apijson.Field
	RequestHeadComplete  apijson.Field
	RequestSize          apijson.Field
	RequestType          apijson.Field
	ResponseContentType  apijson.Field
	ResponseHead         apijson.Field
	ResponseHeadComplete apijson.Field
	ResponseSize         apijson.Field
	StatusCode           apijson.Field
	Step                 apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *LogGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logGetResponseJSON) RawJSON() string {
	return r.raw
}

type LogRequestResponse = interface{}

type LogResponseResponse = interface{}

type LogListParams struct {
	AccountID           param.Field[string]                        `path:"account_id,required"`
	Cached              param.Field[bool]                          `query:"cached"`
	Direction           param.Field[LogListParamsDirection]        `query:"direction"`
	EndDate             param.Field[time.Time]                     `query:"end_date" format:"date-time"`
	Feedback            param.Field[LogListParamsFeedback]         `query:"feedback"`
	Filters             param.Field[[]LogListParamsFilter]         `query:"filters"`
	MaxCost             param.Field[float64]                       `query:"max_cost"`
	MaxDuration         param.Field[float64]                       `query:"max_duration"`
	MaxTokensIn         param.Field[float64]                       `query:"max_tokens_in"`
	MaxTokensOut        param.Field[float64]                       `query:"max_tokens_out"`
	MaxTotalTokens      param.Field[float64]                       `query:"max_total_tokens"`
	MetaInfo            param.Field[bool]                          `query:"meta_info"`
	MinCost             param.Field[float64]                       `query:"min_cost"`
	MinDuration         param.Field[float64]                       `query:"min_duration"`
	MinTokensIn         param.Field[float64]                       `query:"min_tokens_in"`
	MinTokensOut        param.Field[float64]                       `query:"min_tokens_out"`
	MinTotalTokens      param.Field[float64]                       `query:"min_total_tokens"`
	Model               param.Field[string]                        `query:"model"`
	ModelType           param.Field[string]                        `query:"model_type"`
	OrderBy             param.Field[LogListParamsOrderBy]          `query:"order_by"`
	OrderByDirection    param.Field[LogListParamsOrderByDirection] `query:"order_by_direction"`
	Page                param.Field[int64]                         `query:"page"`
	PerPage             param.Field[int64]                         `query:"per_page"`
	Provider            param.Field[string]                        `query:"provider"`
	RequestContentType  param.Field[string]                        `query:"request_content_type"`
	ResponseContentType param.Field[string]                        `query:"response_content_type"`
	Search              param.Field[string]                        `query:"search"`
	StartDate           param.Field[time.Time]                     `query:"start_date" format:"date-time"`
	Success             param.Field[bool]                          `query:"success"`
}

// URLQuery serializes [LogListParams]'s query parameters as `url.Values`.
func (r LogListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogListParamsDirection string

const (
	LogListParamsDirectionAsc  LogListParamsDirection = "asc"
	LogListParamsDirectionDesc LogListParamsDirection = "desc"
)

func (r LogListParamsDirection) IsKnown() bool {
	switch r {
	case LogListParamsDirectionAsc, LogListParamsDirectionDesc:
		return true
	}
	return false
}

type LogListParamsFeedback float64

const (
	LogListParamsFeedback0 LogListParamsFeedback = 0
	LogListParamsFeedback1 LogListParamsFeedback = 1
)

func (r LogListParamsFeedback) IsKnown() bool {
	switch r {
	case LogListParamsFeedback0, LogListParamsFeedback1:
		return true
	}
	return false
}

type LogListParamsFilter struct {
	Key      param.Field[LogListParamsFiltersKey]          `query:"key,required"`
	Operator param.Field[LogListParamsFiltersOperator]     `query:"operator,required"`
	Value    param.Field[[]LogListParamsFiltersValueUnion] `query:"value,required"`
}

// URLQuery serializes [LogListParamsFilter]'s query parameters as `url.Values`.
func (r LogListParamsFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogListParamsFiltersKey string

const (
	LogListParamsFiltersKeyID                  LogListParamsFiltersKey = "id"
	LogListParamsFiltersKeyCreatedAt           LogListParamsFiltersKey = "created_at"
	LogListParamsFiltersKeyRequestContentType  LogListParamsFiltersKey = "request_content_type"
	LogListParamsFiltersKeyResponseContentType LogListParamsFiltersKey = "response_content_type"
	LogListParamsFiltersKeyRequestType         LogListParamsFiltersKey = "request_type"
	LogListParamsFiltersKeySuccess             LogListParamsFiltersKey = "success"
	LogListParamsFiltersKeyCached              LogListParamsFiltersKey = "cached"
	LogListParamsFiltersKeyProvider            LogListParamsFiltersKey = "provider"
	LogListParamsFiltersKeyModel               LogListParamsFiltersKey = "model"
	LogListParamsFiltersKeyModelType           LogListParamsFiltersKey = "model_type"
	LogListParamsFiltersKeyCost                LogListParamsFiltersKey = "cost"
	LogListParamsFiltersKeyTokens              LogListParamsFiltersKey = "tokens"
	LogListParamsFiltersKeyTokensIn            LogListParamsFiltersKey = "tokens_in"
	LogListParamsFiltersKeyTokensOut           LogListParamsFiltersKey = "tokens_out"
	LogListParamsFiltersKeyDuration            LogListParamsFiltersKey = "duration"
	LogListParamsFiltersKeyFeedback            LogListParamsFiltersKey = "feedback"
	LogListParamsFiltersKeyEventID             LogListParamsFiltersKey = "event_id"
	LogListParamsFiltersKeyMetadataKey         LogListParamsFiltersKey = "metadata.key"
	LogListParamsFiltersKeyMetadataValue       LogListParamsFiltersKey = "metadata.value"
	LogListParamsFiltersKeyPromptsPromptID     LogListParamsFiltersKey = "prompts.prompt_id"
	LogListParamsFiltersKeyPromptsVersionID    LogListParamsFiltersKey = "prompts.version_id"
	LogListParamsFiltersKeyAuthentication      LogListParamsFiltersKey = "authentication"
	LogListParamsFiltersKeyWholesale           LogListParamsFiltersKey = "wholesale"
	LogListParamsFiltersKeyCompatibilityMode   LogListParamsFiltersKey = "compatibilityMode"
)

func (r LogListParamsFiltersKey) IsKnown() bool {
	switch r {
	case LogListParamsFiltersKeyID, LogListParamsFiltersKeyCreatedAt, LogListParamsFiltersKeyRequestContentType, LogListParamsFiltersKeyResponseContentType, LogListParamsFiltersKeyRequestType, LogListParamsFiltersKeySuccess, LogListParamsFiltersKeyCached, LogListParamsFiltersKeyProvider, LogListParamsFiltersKeyModel, LogListParamsFiltersKeyModelType, LogListParamsFiltersKeyCost, LogListParamsFiltersKeyTokens, LogListParamsFiltersKeyTokensIn, LogListParamsFiltersKeyTokensOut, LogListParamsFiltersKeyDuration, LogListParamsFiltersKeyFeedback, LogListParamsFiltersKeyEventID, LogListParamsFiltersKeyMetadataKey, LogListParamsFiltersKeyMetadataValue, LogListParamsFiltersKeyPromptsPromptID, LogListParamsFiltersKeyPromptsVersionID, LogListParamsFiltersKeyAuthentication, LogListParamsFiltersKeyWholesale, LogListParamsFiltersKeyCompatibilityMode:
		return true
	}
	return false
}

type LogListParamsFiltersOperator string

const (
	LogListParamsFiltersOperatorEq       LogListParamsFiltersOperator = "eq"
	LogListParamsFiltersOperatorNeq      LogListParamsFiltersOperator = "neq"
	LogListParamsFiltersOperatorContains LogListParamsFiltersOperator = "contains"
	LogListParamsFiltersOperatorLt       LogListParamsFiltersOperator = "lt"
	LogListParamsFiltersOperatorGt       LogListParamsFiltersOperator = "gt"
)

func (r LogListParamsFiltersOperator) IsKnown() bool {
	switch r {
	case LogListParamsFiltersOperatorEq, LogListParamsFiltersOperatorNeq, LogListParamsFiltersOperatorContains, LogListParamsFiltersOperatorLt, LogListParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type LogListParamsFiltersValueUnion interface {
	ImplementsLogListParamsFiltersValueUnion()
}

type LogListParamsOrderBy string

const (
	LogListParamsOrderByCreatedAt LogListParamsOrderBy = "created_at"
	LogListParamsOrderByProvider  LogListParamsOrderBy = "provider"
	LogListParamsOrderByModel     LogListParamsOrderBy = "model"
	LogListParamsOrderByModelType LogListParamsOrderBy = "model_type"
	LogListParamsOrderBySuccess   LogListParamsOrderBy = "success"
	LogListParamsOrderByCached    LogListParamsOrderBy = "cached"
)

func (r LogListParamsOrderBy) IsKnown() bool {
	switch r {
	case LogListParamsOrderByCreatedAt, LogListParamsOrderByProvider, LogListParamsOrderByModel, LogListParamsOrderByModelType, LogListParamsOrderBySuccess, LogListParamsOrderByCached:
		return true
	}
	return false
}

type LogListParamsOrderByDirection string

const (
	LogListParamsOrderByDirectionAsc  LogListParamsOrderByDirection = "asc"
	LogListParamsOrderByDirectionDesc LogListParamsOrderByDirection = "desc"
)

func (r LogListParamsOrderByDirection) IsKnown() bool {
	switch r {
	case LogListParamsOrderByDirectionAsc, LogListParamsOrderByDirectionDesc:
		return true
	}
	return false
}

type LogDeleteParams struct {
	AccountID        param.Field[string]                          `path:"account_id,required"`
	Filters          param.Field[[]LogDeleteParamsFilter]         `query:"filters"`
	Limit            param.Field[int64]                           `query:"limit"`
	OrderBy          param.Field[LogDeleteParamsOrderBy]          `query:"order_by"`
	OrderByDirection param.Field[LogDeleteParamsOrderByDirection] `query:"order_by_direction"`
}

// URLQuery serializes [LogDeleteParams]'s query parameters as `url.Values`.
func (r LogDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogDeleteParamsFilter struct {
	Key      param.Field[LogDeleteParamsFiltersKey]          `query:"key,required"`
	Operator param.Field[LogDeleteParamsFiltersOperator]     `query:"operator,required"`
	Value    param.Field[[]LogDeleteParamsFiltersValueUnion] `query:"value,required"`
}

// URLQuery serializes [LogDeleteParamsFilter]'s query parameters as `url.Values`.
func (r LogDeleteParamsFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LogDeleteParamsFiltersKey string

const (
	LogDeleteParamsFiltersKeyID                  LogDeleteParamsFiltersKey = "id"
	LogDeleteParamsFiltersKeyCreatedAt           LogDeleteParamsFiltersKey = "created_at"
	LogDeleteParamsFiltersKeyRequestContentType  LogDeleteParamsFiltersKey = "request_content_type"
	LogDeleteParamsFiltersKeyResponseContentType LogDeleteParamsFiltersKey = "response_content_type"
	LogDeleteParamsFiltersKeyRequestType         LogDeleteParamsFiltersKey = "request_type"
	LogDeleteParamsFiltersKeySuccess             LogDeleteParamsFiltersKey = "success"
	LogDeleteParamsFiltersKeyCached              LogDeleteParamsFiltersKey = "cached"
	LogDeleteParamsFiltersKeyProvider            LogDeleteParamsFiltersKey = "provider"
	LogDeleteParamsFiltersKeyModel               LogDeleteParamsFiltersKey = "model"
	LogDeleteParamsFiltersKeyModelType           LogDeleteParamsFiltersKey = "model_type"
	LogDeleteParamsFiltersKeyCost                LogDeleteParamsFiltersKey = "cost"
	LogDeleteParamsFiltersKeyTokens              LogDeleteParamsFiltersKey = "tokens"
	LogDeleteParamsFiltersKeyTokensIn            LogDeleteParamsFiltersKey = "tokens_in"
	LogDeleteParamsFiltersKeyTokensOut           LogDeleteParamsFiltersKey = "tokens_out"
	LogDeleteParamsFiltersKeyDuration            LogDeleteParamsFiltersKey = "duration"
	LogDeleteParamsFiltersKeyFeedback            LogDeleteParamsFiltersKey = "feedback"
	LogDeleteParamsFiltersKeyEventID             LogDeleteParamsFiltersKey = "event_id"
	LogDeleteParamsFiltersKeyMetadataKey         LogDeleteParamsFiltersKey = "metadata.key"
	LogDeleteParamsFiltersKeyMetadataValue       LogDeleteParamsFiltersKey = "metadata.value"
	LogDeleteParamsFiltersKeyPromptsPromptID     LogDeleteParamsFiltersKey = "prompts.prompt_id"
	LogDeleteParamsFiltersKeyPromptsVersionID    LogDeleteParamsFiltersKey = "prompts.version_id"
	LogDeleteParamsFiltersKeyAuthentication      LogDeleteParamsFiltersKey = "authentication"
	LogDeleteParamsFiltersKeyWholesale           LogDeleteParamsFiltersKey = "wholesale"
	LogDeleteParamsFiltersKeyCompatibilityMode   LogDeleteParamsFiltersKey = "compatibilityMode"
)

func (r LogDeleteParamsFiltersKey) IsKnown() bool {
	switch r {
	case LogDeleteParamsFiltersKeyID, LogDeleteParamsFiltersKeyCreatedAt, LogDeleteParamsFiltersKeyRequestContentType, LogDeleteParamsFiltersKeyResponseContentType, LogDeleteParamsFiltersKeyRequestType, LogDeleteParamsFiltersKeySuccess, LogDeleteParamsFiltersKeyCached, LogDeleteParamsFiltersKeyProvider, LogDeleteParamsFiltersKeyModel, LogDeleteParamsFiltersKeyModelType, LogDeleteParamsFiltersKeyCost, LogDeleteParamsFiltersKeyTokens, LogDeleteParamsFiltersKeyTokensIn, LogDeleteParamsFiltersKeyTokensOut, LogDeleteParamsFiltersKeyDuration, LogDeleteParamsFiltersKeyFeedback, LogDeleteParamsFiltersKeyEventID, LogDeleteParamsFiltersKeyMetadataKey, LogDeleteParamsFiltersKeyMetadataValue, LogDeleteParamsFiltersKeyPromptsPromptID, LogDeleteParamsFiltersKeyPromptsVersionID, LogDeleteParamsFiltersKeyAuthentication, LogDeleteParamsFiltersKeyWholesale, LogDeleteParamsFiltersKeyCompatibilityMode:
		return true
	}
	return false
}

type LogDeleteParamsFiltersOperator string

const (
	LogDeleteParamsFiltersOperatorEq       LogDeleteParamsFiltersOperator = "eq"
	LogDeleteParamsFiltersOperatorNeq      LogDeleteParamsFiltersOperator = "neq"
	LogDeleteParamsFiltersOperatorContains LogDeleteParamsFiltersOperator = "contains"
	LogDeleteParamsFiltersOperatorLt       LogDeleteParamsFiltersOperator = "lt"
	LogDeleteParamsFiltersOperatorGt       LogDeleteParamsFiltersOperator = "gt"
)

func (r LogDeleteParamsFiltersOperator) IsKnown() bool {
	switch r {
	case LogDeleteParamsFiltersOperatorEq, LogDeleteParamsFiltersOperatorNeq, LogDeleteParamsFiltersOperatorContains, LogDeleteParamsFiltersOperatorLt, LogDeleteParamsFiltersOperatorGt:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type LogDeleteParamsFiltersValueUnion interface {
	ImplementsLogDeleteParamsFiltersValueUnion()
}

type LogDeleteParamsOrderBy string

const (
	LogDeleteParamsOrderByCreatedAt LogDeleteParamsOrderBy = "created_at"
	LogDeleteParamsOrderByProvider  LogDeleteParamsOrderBy = "provider"
	LogDeleteParamsOrderByModel     LogDeleteParamsOrderBy = "model"
	LogDeleteParamsOrderByModelType LogDeleteParamsOrderBy = "model_type"
	LogDeleteParamsOrderBySuccess   LogDeleteParamsOrderBy = "success"
	LogDeleteParamsOrderByCached    LogDeleteParamsOrderBy = "cached"
	LogDeleteParamsOrderByCost      LogDeleteParamsOrderBy = "cost"
	LogDeleteParamsOrderByTokensIn  LogDeleteParamsOrderBy = "tokens_in"
	LogDeleteParamsOrderByTokensOut LogDeleteParamsOrderBy = "tokens_out"
	LogDeleteParamsOrderByDuration  LogDeleteParamsOrderBy = "duration"
	LogDeleteParamsOrderByFeedback  LogDeleteParamsOrderBy = "feedback"
)

func (r LogDeleteParamsOrderBy) IsKnown() bool {
	switch r {
	case LogDeleteParamsOrderByCreatedAt, LogDeleteParamsOrderByProvider, LogDeleteParamsOrderByModel, LogDeleteParamsOrderByModelType, LogDeleteParamsOrderBySuccess, LogDeleteParamsOrderByCached, LogDeleteParamsOrderByCost, LogDeleteParamsOrderByTokensIn, LogDeleteParamsOrderByTokensOut, LogDeleteParamsOrderByDuration, LogDeleteParamsOrderByFeedback:
		return true
	}
	return false
}

type LogDeleteParamsOrderByDirection string

const (
	LogDeleteParamsOrderByDirectionAsc  LogDeleteParamsOrderByDirection = "asc"
	LogDeleteParamsOrderByDirectionDesc LogDeleteParamsOrderByDirection = "desc"
)

func (r LogDeleteParamsOrderByDirection) IsKnown() bool {
	switch r {
	case LogDeleteParamsOrderByDirectionAsc, LogDeleteParamsOrderByDirectionDesc:
		return true
	}
	return false
}

type LogEditParams struct {
	AccountID param.Field[string]                                `path:"account_id,required"`
	Feedback  param.Field[float64]                               `json:"feedback"`
	Metadata  param.Field[map[string]LogEditParamsMetadataUnion] `json:"metadata"`
	Score     param.Field[float64]                               `json:"score"`
}

func (r LogEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type LogEditParamsMetadataUnion interface {
	ImplementsLogEditParamsMetadataUnion()
}

type LogEditResponseEnvelope struct {
	Result  LogEditResponse             `json:"result,required"`
	Success bool                        `json:"success,required"`
	JSON    logEditResponseEnvelopeJSON `json:"-"`
}

// logEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogEditResponseEnvelope]
type logEditResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type LogGetResponseEnvelope struct {
	Result  LogGetResponse             `json:"result,required"`
	Success bool                       `json:"success,required"`
	JSON    logGetResponseEnvelopeJSON `json:"-"`
}

// logGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LogGetResponseEnvelope]
type logGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LogGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LogRequestParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type LogResponseParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
