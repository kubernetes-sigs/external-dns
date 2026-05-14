// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package brand_protection

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// MatchService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMatchService] method instead.
type MatchService struct {
	Options []option.RequestOption
}

// NewMatchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMatchService(opts ...option.RequestOption) (r *MatchService) {
	r = &MatchService{}
	r.Options = opts
	return
}

// Return matches as CSV for string queries based on ID
func (r *MatchService) Download(ctx context.Context, params MatchDownloadParams, opts ...option.RequestOption) (res *MatchDownloadResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/matches/download", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Return matches for string queries based on ID
func (r *MatchService) Get(ctx context.Context, params MatchGetParams, opts ...option.RequestOption) (res *MatchGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/brand-protection/matches", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type MatchDownloadResponse struct {
	Matches []map[string]interface{}  `json:"matches"`
	Total   int64                     `json:"total"`
	JSON    matchDownloadResponseJSON `json:"-"`
}

// matchDownloadResponseJSON contains the JSON metadata for the struct
// [MatchDownloadResponse]
type matchDownloadResponseJSON struct {
	Matches     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatchDownloadResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matchDownloadResponseJSON) RawJSON() string {
	return r.raw
}

type MatchGetResponse struct {
	Matches []map[string]interface{} `json:"matches"`
	Total   int64                    `json:"total"`
	JSON    matchGetResponseJSON     `json:"-"`
}

// matchGetResponseJSON contains the JSON metadata for the struct
// [MatchGetResponse]
type matchGetResponseJSON struct {
	Matches     apijson.Field
	Total       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatchGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matchGetResponseJSON) RawJSON() string {
	return r.raw
}

type MatchDownloadParams struct {
	AccountID       param.Field[string] `path:"account_id,required"`
	ID              param.Field[string] `query:"id"`
	IncludeDomainID param.Field[bool]   `query:"include_domain_id"`
	Limit           param.Field[int64]  `query:"limit"`
	Offset          param.Field[int64]  `query:"offset"`
}

// URLQuery serializes [MatchDownloadParams]'s query parameters as `url.Values`.
func (r MatchDownloadParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MatchGetParams struct {
	AccountID       param.Field[string] `path:"account_id,required"`
	ID              param.Field[string] `query:"id"`
	IncludeDomainID param.Field[bool]   `query:"include_domain_id"`
	Limit           param.Field[int64]  `query:"limit"`
	Offset          param.Field[int64]  `query:"offset"`
}

// URLQuery serializes [MatchGetParams]'s query parameters as `url.Values`.
func (r MatchGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
