// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// ModelService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
	Schema  *ModelSchemaService
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r *ModelService) {
	r = &ModelService{}
	r.Options = opts
	r.Schema = NewModelSchemaService(opts...)
	return
}

// Model Search
func (r *ModelService) List(ctx context.Context, params ModelListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[ModelListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai/models/search", params.AccountID)
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

// Model Search
func (r *ModelService) ListAutoPaging(ctx context.Context, params ModelListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[ModelListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type ModelListResponse = interface{}

type ModelListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter by Author
	Author param.Field[string] `query:"author"`
	// Filter to hide experimental models
	HideExperimental param.Field[bool]  `query:"hide_experimental"`
	Page             param.Field[int64] `query:"page"`
	PerPage          param.Field[int64] `query:"per_page"`
	// Search
	Search param.Field[string] `query:"search"`
	// Filter by Source Id
	Source param.Field[float64] `query:"source"`
	// Filter by Task Name
	Task param.Field[string] `query:"task"`
}

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
