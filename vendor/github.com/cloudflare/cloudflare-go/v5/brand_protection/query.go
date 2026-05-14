// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

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
)

// QueryService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryService] method instead.
type QueryService struct {
	Options []option.RequestOption
}

// NewQueryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueryService(opts ...option.RequestOption) (r *QueryService) {
	r = &QueryService{}
	r.Options = opts
	return
}

// Return a success message after creating new saved string queries
func (r *QueryService) New(ctx context.Context, params QueryNewParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/queries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Return a success message after deleting saved string queries by ID
func (r *QueryService) Delete(ctx context.Context, params QueryDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/queries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return
}

type QueryNewParams struct {
	AccountID     param.Field[string]      `path:"account_id,required"`
	ID            param.Field[string]      `query:"id"`
	QueryScan     param.Field[bool]        `query:"scan"`
	QueryTag      param.Field[string]      `query:"tag"`
	MaxTime       param.Field[time.Time]   `json:"max_time" format:"date-time"`
	MinTime       param.Field[time.Time]   `json:"min_time" format:"date-time"`
	BodyScan      param.Field[bool]        `json:"scan"`
	StringMatches param.Field[interface{}] `json:"string_matches"`
	BodyTag       param.Field[string]      `json:"tag"`
}

func (r QueryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [QueryNewParams]'s query parameters as `url.Values`.
func (r QueryNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type QueryDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	ID        param.Field[string] `query:"id"`
	Scan      param.Field[bool]   `query:"scan"`
	Tag       param.Field[string] `query:"tag"`
}

// URLQuery serializes [QueryDeleteParams]'s query parameters as `url.Values`.
func (r QueryDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
