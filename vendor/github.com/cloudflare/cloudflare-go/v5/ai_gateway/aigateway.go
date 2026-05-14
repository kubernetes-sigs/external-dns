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

// AIGatewayService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIGatewayService] method instead.
type AIGatewayService struct {
	Options         []option.RequestOption
	EvaluationTypes *EvaluationTypeService
	Logs            *LogService
	Datasets        *DatasetService
	Evaluations     *EvaluationService
	URLs            *URLService
}

// NewAIGatewayService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIGatewayService(opts ...option.RequestOption) (r *AIGatewayService) {
	r = &AIGatewayService{}
	r.Options = opts
	r.EvaluationTypes = NewEvaluationTypeService(opts...)
	r.Logs = NewLogService(opts...)
	r.Datasets = NewDatasetService(opts...)
	r.Evaluations = NewEvaluationService(opts...)
	r.URLs = NewURLService(opts...)
	return
}

// Create a new Gateway
func (r *AIGatewayService) New(ctx context.Context, params AIGatewayNewParams, opts ...option.RequestOption) (res *AIGatewayNewResponse, err error) {
	var env AIGatewayNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a Gateway
func (r *AIGatewayService) Update(ctx context.Context, id string, params AIGatewayUpdateParams, opts ...option.RequestOption) (res *AIGatewayUpdateResponse, err error) {
	var env AIGatewayUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Gateways
func (r *AIGatewayService) List(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AIGatewayListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways", params.AccountID)
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

// List Gateways
func (r *AIGatewayService) ListAutoPaging(ctx context.Context, params AIGatewayListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AIGatewayListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a Gateway
func (r *AIGatewayService) Delete(ctx context.Context, id string, body AIGatewayDeleteParams, opts ...option.RequestOption) (res *AIGatewayDeleteResponse, err error) {
	var env AIGatewayDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", body.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a Gateway
func (r *AIGatewayService) Get(ctx context.Context, id string, query AIGatewayGetParams, opts ...option.RequestOption) (res *AIGatewayGetResponse, err error) {
	var env AIGatewayGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s", query.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AIGatewayNewResponse struct {
	// gateway id
	ID                      string                                    `json:"id,required"`
	AccountID               string                                    `json:"account_id,required"`
	AccountTag              string                                    `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                      `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                     `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                      `json:"collect_logs,required"`
	CreatedAt               time.Time                                 `json:"created_at,required" format:"date-time"`
	InternalID              string                                    `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                 `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                     `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                     `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayNewResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                      `json:"authentication"`
	LogManagement           int64                                     `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayNewResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                      `json:"logpush"`
	LogpushPublicKey        string                                    `json:"logpush_public_key,nullable"`
	StoreID                 string                                    `json:"store_id,nullable"`
	JSON                    aiGatewayNewResponseJSON                  `json:"-"`
}

// aiGatewayNewResponseJSON contains the JSON metadata for the struct
// [AIGatewayNewResponse]
type aiGatewayNewResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	StoreID                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayNewResponseRateLimitingTechnique string

const (
	AIGatewayNewResponseRateLimitingTechniqueFixed   AIGatewayNewResponseRateLimitingTechnique = "fixed"
	AIGatewayNewResponseRateLimitingTechniqueSliding AIGatewayNewResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayNewResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseRateLimitingTechniqueFixed, AIGatewayNewResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayNewResponseLogManagementStrategy string

const (
	AIGatewayNewResponseLogManagementStrategyStopInserting AIGatewayNewResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayNewResponseLogManagementStrategyDeleteOldest  AIGatewayNewResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayNewResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayNewResponseLogManagementStrategyStopInserting, AIGatewayNewResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayUpdateResponse struct {
	// gateway id
	ID                      string                                       `json:"id,required"`
	AccountID               string                                       `json:"account_id,required"`
	AccountTag              string                                       `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                         `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                        `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                         `json:"collect_logs,required"`
	CreatedAt               time.Time                                    `json:"created_at,required" format:"date-time"`
	InternalID              string                                       `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                    `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                        `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                        `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayUpdateResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                         `json:"authentication"`
	LogManagement           int64                                        `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayUpdateResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                         `json:"logpush"`
	LogpushPublicKey        string                                       `json:"logpush_public_key,nullable"`
	StoreID                 string                                       `json:"store_id,nullable"`
	JSON                    aiGatewayUpdateResponseJSON                  `json:"-"`
}

// aiGatewayUpdateResponseJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponse]
type aiGatewayUpdateResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	StoreID                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateResponseRateLimitingTechnique string

const (
	AIGatewayUpdateResponseRateLimitingTechniqueFixed   AIGatewayUpdateResponseRateLimitingTechnique = "fixed"
	AIGatewayUpdateResponseRateLimitingTechniqueSliding AIGatewayUpdateResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayUpdateResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseRateLimitingTechniqueFixed, AIGatewayUpdateResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayUpdateResponseLogManagementStrategy string

const (
	AIGatewayUpdateResponseLogManagementStrategyStopInserting AIGatewayUpdateResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayUpdateResponseLogManagementStrategyDeleteOldest  AIGatewayUpdateResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayUpdateResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayUpdateResponseLogManagementStrategyStopInserting, AIGatewayUpdateResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayListResponse struct {
	// gateway id
	ID                      string                                     `json:"id,required"`
	AccountID               string                                     `json:"account_id,required"`
	AccountTag              string                                     `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                       `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                      `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                       `json:"collect_logs,required"`
	CreatedAt               time.Time                                  `json:"created_at,required" format:"date-time"`
	InternalID              string                                     `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                  `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                      `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                      `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayListResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                       `json:"authentication"`
	LogManagement           int64                                      `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayListResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                       `json:"logpush"`
	LogpushPublicKey        string                                     `json:"logpush_public_key,nullable"`
	StoreID                 string                                     `json:"store_id,nullable"`
	JSON                    aiGatewayListResponseJSON                  `json:"-"`
}

// aiGatewayListResponseJSON contains the JSON metadata for the struct
// [AIGatewayListResponse]
type aiGatewayListResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	StoreID                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayListResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListResponseRateLimitingTechnique string

const (
	AIGatewayListResponseRateLimitingTechniqueFixed   AIGatewayListResponseRateLimitingTechnique = "fixed"
	AIGatewayListResponseRateLimitingTechniqueSliding AIGatewayListResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayListResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayListResponseRateLimitingTechniqueFixed, AIGatewayListResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayListResponseLogManagementStrategy string

const (
	AIGatewayListResponseLogManagementStrategyStopInserting AIGatewayListResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayListResponseLogManagementStrategyDeleteOldest  AIGatewayListResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayListResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayListResponseLogManagementStrategyStopInserting, AIGatewayListResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayDeleteResponse struct {
	// gateway id
	ID                      string                                       `json:"id,required"`
	AccountID               string                                       `json:"account_id,required"`
	AccountTag              string                                       `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                         `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                        `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                         `json:"collect_logs,required"`
	CreatedAt               time.Time                                    `json:"created_at,required" format:"date-time"`
	InternalID              string                                       `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                    `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                        `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                        `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayDeleteResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                         `json:"authentication"`
	LogManagement           int64                                        `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayDeleteResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                         `json:"logpush"`
	LogpushPublicKey        string                                       `json:"logpush_public_key,nullable"`
	StoreID                 string                                       `json:"store_id,nullable"`
	JSON                    aiGatewayDeleteResponseJSON                  `json:"-"`
}

// aiGatewayDeleteResponseJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponse]
type aiGatewayDeleteResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	StoreID                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayDeleteResponseRateLimitingTechnique string

const (
	AIGatewayDeleteResponseRateLimitingTechniqueFixed   AIGatewayDeleteResponseRateLimitingTechnique = "fixed"
	AIGatewayDeleteResponseRateLimitingTechniqueSliding AIGatewayDeleteResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayDeleteResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseRateLimitingTechniqueFixed, AIGatewayDeleteResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayDeleteResponseLogManagementStrategy string

const (
	AIGatewayDeleteResponseLogManagementStrategyStopInserting AIGatewayDeleteResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayDeleteResponseLogManagementStrategyDeleteOldest  AIGatewayDeleteResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayDeleteResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayDeleteResponseLogManagementStrategyStopInserting, AIGatewayDeleteResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayGetResponse struct {
	// gateway id
	ID                      string                                    `json:"id,required"`
	AccountID               string                                    `json:"account_id,required"`
	AccountTag              string                                    `json:"account_tag,required"`
	CacheInvalidateOnUpdate bool                                      `json:"cache_invalidate_on_update,required"`
	CacheTTL                int64                                     `json:"cache_ttl,required,nullable"`
	CollectLogs             bool                                      `json:"collect_logs,required"`
	CreatedAt               time.Time                                 `json:"created_at,required" format:"date-time"`
	InternalID              string                                    `json:"internal_id,required" format:"uuid"`
	ModifiedAt              time.Time                                 `json:"modified_at,required" format:"date-time"`
	RateLimitingInterval    int64                                     `json:"rate_limiting_interval,required,nullable"`
	RateLimitingLimit       int64                                     `json:"rate_limiting_limit,required,nullable"`
	RateLimitingTechnique   AIGatewayGetResponseRateLimitingTechnique `json:"rate_limiting_technique,required"`
	Authentication          bool                                      `json:"authentication"`
	LogManagement           int64                                     `json:"log_management,nullable"`
	LogManagementStrategy   AIGatewayGetResponseLogManagementStrategy `json:"log_management_strategy,nullable"`
	Logpush                 bool                                      `json:"logpush"`
	LogpushPublicKey        string                                    `json:"logpush_public_key,nullable"`
	StoreID                 string                                    `json:"store_id,nullable"`
	JSON                    aiGatewayGetResponseJSON                  `json:"-"`
}

// aiGatewayGetResponseJSON contains the JSON metadata for the struct
// [AIGatewayGetResponse]
type aiGatewayGetResponseJSON struct {
	ID                      apijson.Field
	AccountID               apijson.Field
	AccountTag              apijson.Field
	CacheInvalidateOnUpdate apijson.Field
	CacheTTL                apijson.Field
	CollectLogs             apijson.Field
	CreatedAt               apijson.Field
	InternalID              apijson.Field
	ModifiedAt              apijson.Field
	RateLimitingInterval    apijson.Field
	RateLimitingLimit       apijson.Field
	RateLimitingTechnique   apijson.Field
	Authentication          apijson.Field
	LogManagement           apijson.Field
	LogManagementStrategy   apijson.Field
	Logpush                 apijson.Field
	LogpushPublicKey        apijson.Field
	StoreID                 apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AIGatewayGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetResponseRateLimitingTechnique string

const (
	AIGatewayGetResponseRateLimitingTechniqueFixed   AIGatewayGetResponseRateLimitingTechnique = "fixed"
	AIGatewayGetResponseRateLimitingTechniqueSliding AIGatewayGetResponseRateLimitingTechnique = "sliding"
)

func (r AIGatewayGetResponseRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseRateLimitingTechniqueFixed, AIGatewayGetResponseRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayGetResponseLogManagementStrategy string

const (
	AIGatewayGetResponseLogManagementStrategyStopInserting AIGatewayGetResponseLogManagementStrategy = "STOP_INSERTING"
	AIGatewayGetResponseLogManagementStrategyDeleteOldest  AIGatewayGetResponseLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayGetResponseLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayGetResponseLogManagementStrategyStopInserting, AIGatewayGetResponseLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// gateway id
	ID                      param.Field[string]                                  `json:"id,required"`
	CacheInvalidateOnUpdate param.Field[bool]                                    `json:"cache_invalidate_on_update,required"`
	CacheTTL                param.Field[int64]                                   `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]                                    `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]                                   `json:"rate_limiting_interval,required"`
	RateLimitingLimit       param.Field[int64]                                   `json:"rate_limiting_limit,required"`
	RateLimitingTechnique   param.Field[AIGatewayNewParamsRateLimitingTechnique] `json:"rate_limiting_technique,required"`
	Authentication          param.Field[bool]                                    `json:"authentication"`
	LogManagement           param.Field[int64]                                   `json:"log_management"`
	LogManagementStrategy   param.Field[AIGatewayNewParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                    `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                  `json:"logpush_public_key"`
}

func (r AIGatewayNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayNewParamsRateLimitingTechnique string

const (
	AIGatewayNewParamsRateLimitingTechniqueFixed   AIGatewayNewParamsRateLimitingTechnique = "fixed"
	AIGatewayNewParamsRateLimitingTechniqueSliding AIGatewayNewParamsRateLimitingTechnique = "sliding"
)

func (r AIGatewayNewParamsRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayNewParamsRateLimitingTechniqueFixed, AIGatewayNewParamsRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayNewParamsLogManagementStrategy string

const (
	AIGatewayNewParamsLogManagementStrategyStopInserting AIGatewayNewParamsLogManagementStrategy = "STOP_INSERTING"
	AIGatewayNewParamsLogManagementStrategyDeleteOldest  AIGatewayNewParamsLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayNewParamsLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayNewParamsLogManagementStrategyStopInserting, AIGatewayNewParamsLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayNewResponseEnvelope struct {
	Result  AIGatewayNewResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    aiGatewayNewResponseEnvelopeJSON `json:"-"`
}

// aiGatewayNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayNewResponseEnvelope]
type aiGatewayNewResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayUpdateParams struct {
	AccountID               param.Field[string]                                     `path:"account_id,required"`
	CacheInvalidateOnUpdate param.Field[bool]                                       `json:"cache_invalidate_on_update,required"`
	CacheTTL                param.Field[int64]                                      `json:"cache_ttl,required"`
	CollectLogs             param.Field[bool]                                       `json:"collect_logs,required"`
	RateLimitingInterval    param.Field[int64]                                      `json:"rate_limiting_interval,required"`
	RateLimitingLimit       param.Field[int64]                                      `json:"rate_limiting_limit,required"`
	RateLimitingTechnique   param.Field[AIGatewayUpdateParamsRateLimitingTechnique] `json:"rate_limiting_technique,required"`
	Authentication          param.Field[bool]                                       `json:"authentication"`
	LogManagement           param.Field[int64]                                      `json:"log_management"`
	LogManagementStrategy   param.Field[AIGatewayUpdateParamsLogManagementStrategy] `json:"log_management_strategy"`
	Logpush                 param.Field[bool]                                       `json:"logpush"`
	LogpushPublicKey        param.Field[string]                                     `json:"logpush_public_key"`
	StoreID                 param.Field[string]                                     `json:"store_id"`
}

func (r AIGatewayUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AIGatewayUpdateParamsRateLimitingTechnique string

const (
	AIGatewayUpdateParamsRateLimitingTechniqueFixed   AIGatewayUpdateParamsRateLimitingTechnique = "fixed"
	AIGatewayUpdateParamsRateLimitingTechniqueSliding AIGatewayUpdateParamsRateLimitingTechnique = "sliding"
)

func (r AIGatewayUpdateParamsRateLimitingTechnique) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsRateLimitingTechniqueFixed, AIGatewayUpdateParamsRateLimitingTechniqueSliding:
		return true
	}
	return false
}

type AIGatewayUpdateParamsLogManagementStrategy string

const (
	AIGatewayUpdateParamsLogManagementStrategyStopInserting AIGatewayUpdateParamsLogManagementStrategy = "STOP_INSERTING"
	AIGatewayUpdateParamsLogManagementStrategyDeleteOldest  AIGatewayUpdateParamsLogManagementStrategy = "DELETE_OLDEST"
)

func (r AIGatewayUpdateParamsLogManagementStrategy) IsKnown() bool {
	switch r {
	case AIGatewayUpdateParamsLogManagementStrategyStopInserting, AIGatewayUpdateParamsLogManagementStrategyDeleteOldest:
		return true
	}
	return false
}

type AIGatewayUpdateResponseEnvelope struct {
	Result  AIGatewayUpdateResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiGatewayUpdateResponseEnvelopeJSON `json:"-"`
}

// aiGatewayUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayUpdateResponseEnvelope]
type aiGatewayUpdateResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	Page      param.Field[int64]  `query:"page"`
	PerPage   param.Field[int64]  `query:"per_page"`
	// Search by id
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AIGatewayListParams]'s query parameters as `url.Values`.
func (r AIGatewayListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AIGatewayDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AIGatewayDeleteResponseEnvelope struct {
	Result  AIGatewayDeleteResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    aiGatewayDeleteResponseEnvelopeJSON `json:"-"`
}

// aiGatewayDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayDeleteResponseEnvelope]
type aiGatewayDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AIGatewayGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type AIGatewayGetResponseEnvelope struct {
	Result  AIGatewayGetResponse             `json:"result,required"`
	Success bool                             `json:"success,required"`
	JSON    aiGatewayGetResponseEnvelopeJSON `json:"-"`
}

// aiGatewayGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AIGatewayGetResponseEnvelope]
type aiGatewayGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AIGatewayGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aiGatewayGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
