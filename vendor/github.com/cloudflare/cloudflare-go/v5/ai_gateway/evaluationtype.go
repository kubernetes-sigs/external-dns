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

// EvaluationTypeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvaluationTypeService] method instead.
type EvaluationTypeService struct {
	Options []option.RequestOption
}

// NewEvaluationTypeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvaluationTypeService(opts ...option.RequestOption) (r *EvaluationTypeService) {
	r = &EvaluationTypeService{}
	r.Options = opts
	return
}

// List Evaluators
func (r *EvaluationTypeService) List(ctx context.Context, params EvaluationTypeListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[EvaluationTypeListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/evaluation-types", params.AccountID)
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

// List Evaluators
func (r *EvaluationTypeService) ListAutoPaging(ctx context.Context, params EvaluationTypeListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[EvaluationTypeListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type EvaluationTypeListResponse struct {
	ID          string                         `json:"id,required"`
	CreatedAt   time.Time                      `json:"created_at,required" format:"date-time"`
	Description string                         `json:"description,required"`
	Enable      bool                           `json:"enable,required"`
	Mandatory   bool                           `json:"mandatory,required"`
	ModifiedAt  time.Time                      `json:"modified_at,required" format:"date-time"`
	Name        string                         `json:"name,required"`
	Type        string                         `json:"type,required"`
	JSON        evaluationTypeListResponseJSON `json:"-"`
}

// evaluationTypeListResponseJSON contains the JSON metadata for the struct
// [EvaluationTypeListResponse]
type evaluationTypeListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Enable      apijson.Field
	Mandatory   apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EvaluationTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r evaluationTypeListResponseJSON) RawJSON() string {
	return r.raw
}

type EvaluationTypeListParams struct {
	AccountID        param.Field[string]                                   `path:"account_id,required"`
	OrderBy          param.Field[string]                                   `query:"order_by"`
	OrderByDirection param.Field[EvaluationTypeListParamsOrderByDirection] `query:"order_by_direction"`
	Page             param.Field[int64]                                    `query:"page"`
	PerPage          param.Field[int64]                                    `query:"per_page"`
}

// URLQuery serializes [EvaluationTypeListParams]'s query parameters as
// `url.Values`.
func (r EvaluationTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type EvaluationTypeListParamsOrderByDirection string

const (
	EvaluationTypeListParamsOrderByDirectionAsc  EvaluationTypeListParamsOrderByDirection = "asc"
	EvaluationTypeListParamsOrderByDirectionDesc EvaluationTypeListParamsOrderByDirection = "desc"
)

func (r EvaluationTypeListParamsOrderByDirection) IsKnown() bool {
	switch r {
	case EvaluationTypeListParamsOrderByDirectionAsc, EvaluationTypeListParamsOrderByDirectionDesc:
		return true
	}
	return false
}
