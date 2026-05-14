// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai

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

// FinetunePublicService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFinetunePublicService] method instead.
type FinetunePublicService struct {
	Options []option.RequestOption
}

// NewFinetunePublicService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFinetunePublicService(opts ...option.RequestOption) (r *FinetunePublicService) {
	r = &FinetunePublicService{}
	r.Options = opts
	return
}

// List Public Finetunes
func (r *FinetunePublicService) List(ctx context.Context, params FinetunePublicListParams, opts ...option.RequestOption) (res *pagination.SinglePage[FinetunePublicListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/finetunes/public", params.AccountID)
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

// List Public Finetunes
func (r *FinetunePublicService) ListAutoPaging(ctx context.Context, params FinetunePublicListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FinetunePublicListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type FinetunePublicListResponse struct {
	ID          string                         `json:"id,required" format:"uuid"`
	CreatedAt   time.Time                      `json:"created_at,required" format:"date-time"`
	Model       string                         `json:"model,required"`
	ModifiedAt  time.Time                      `json:"modified_at,required" format:"date-time"`
	Name        string                         `json:"name,required"`
	Public      bool                           `json:"public,required"`
	Description string                         `json:"description"`
	JSON        finetunePublicListResponseJSON `json:"-"`
}

// finetunePublicListResponseJSON contains the JSON metadata for the struct
// [FinetunePublicListResponse]
type finetunePublicListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Model       apijson.Field
	ModifiedAt  apijson.Field
	Name        apijson.Field
	Public      apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FinetunePublicListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r finetunePublicListResponseJSON) RawJSON() string {
	return r.raw
}

type FinetunePublicListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Pagination Limit
	Limit param.Field[float64] `query:"limit"`
	// Pagination Offset
	Offset param.Field[float64] `query:"offset"`
	// Order By Column Name
	OrderBy param.Field[string] `query:"orderBy"`
}

// URLQuery serializes [FinetunePublicListParams]'s query parameters as
// `url.Values`.
func (r FinetunePublicListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
