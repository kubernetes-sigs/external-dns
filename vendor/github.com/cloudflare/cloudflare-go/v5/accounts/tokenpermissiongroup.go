// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// TokenPermissionGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTokenPermissionGroupService] method instead.
type TokenPermissionGroupService struct {
	Options []option.RequestOption
}

// NewTokenPermissionGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTokenPermissionGroupService(opts ...option.RequestOption) (r *TokenPermissionGroupService) {
	r = &TokenPermissionGroupService{}
	r.Options = opts
	return
}

// Find all available permission groups for Account Owned API Tokens
func (r *TokenPermissionGroupService) List(ctx context.Context, params TokenPermissionGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TokenPermissionGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/permission_groups", params.AccountID)
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

// Find all available permission groups for Account Owned API Tokens
func (r *TokenPermissionGroupService) ListAutoPaging(ctx context.Context, params TokenPermissionGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TokenPermissionGroupListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Find all available permission groups for Account Owned API Tokens
func (r *TokenPermissionGroupService) Get(ctx context.Context, params TokenPermissionGroupGetParams, opts ...option.RequestOption) (res *[]TokenPermissionGroupGetResponse, err error) {
	var env TokenPermissionGroupGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/tokens/permission_groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TokenPermissionGroupListResponse struct {
	// Public ID.
	ID string `json:"id"`
	// Permission Group Name
	Name string `json:"name"`
	// Resources to which the Permission Group is scoped
	Scopes []TokenPermissionGroupListResponseScope `json:"scopes"`
	JSON   tokenPermissionGroupListResponseJSON    `json:"-"`
}

// tokenPermissionGroupListResponseJSON contains the JSON metadata for the struct
// [TokenPermissionGroupListResponse]
type tokenPermissionGroupListResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Scopes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupListResponseScope string

const (
	TokenPermissionGroupListResponseScopeComCloudflareAPIAccount     TokenPermissionGroupListResponseScope = "com.cloudflare.api.account"
	TokenPermissionGroupListResponseScopeComCloudflareAPIAccountZone TokenPermissionGroupListResponseScope = "com.cloudflare.api.account.zone"
	TokenPermissionGroupListResponseScopeComCloudflareAPIUser        TokenPermissionGroupListResponseScope = "com.cloudflare.api.user"
	TokenPermissionGroupListResponseScopeComCloudflareEdgeR2Bucket   TokenPermissionGroupListResponseScope = "com.cloudflare.edge.r2.bucket"
)

func (r TokenPermissionGroupListResponseScope) IsKnown() bool {
	switch r {
	case TokenPermissionGroupListResponseScopeComCloudflareAPIAccount, TokenPermissionGroupListResponseScopeComCloudflareAPIAccountZone, TokenPermissionGroupListResponseScopeComCloudflareAPIUser, TokenPermissionGroupListResponseScopeComCloudflareEdgeR2Bucket:
		return true
	}
	return false
}

type TokenPermissionGroupGetResponse struct {
	// Public ID.
	ID string `json:"id"`
	// Permission Group Name
	Name string `json:"name"`
	// Resources to which the Permission Group is scoped
	Scopes []TokenPermissionGroupGetResponseScope `json:"scopes"`
	JSON   tokenPermissionGroupGetResponseJSON    `json:"-"`
}

// tokenPermissionGroupGetResponseJSON contains the JSON metadata for the struct
// [TokenPermissionGroupGetResponse]
type tokenPermissionGroupGetResponseJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Scopes      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupGetResponseScope string

const (
	TokenPermissionGroupGetResponseScopeComCloudflareAPIAccount     TokenPermissionGroupGetResponseScope = "com.cloudflare.api.account"
	TokenPermissionGroupGetResponseScopeComCloudflareAPIAccountZone TokenPermissionGroupGetResponseScope = "com.cloudflare.api.account.zone"
	TokenPermissionGroupGetResponseScopeComCloudflareAPIUser        TokenPermissionGroupGetResponseScope = "com.cloudflare.api.user"
	TokenPermissionGroupGetResponseScopeComCloudflareEdgeR2Bucket   TokenPermissionGroupGetResponseScope = "com.cloudflare.edge.r2.bucket"
)

func (r TokenPermissionGroupGetResponseScope) IsKnown() bool {
	switch r {
	case TokenPermissionGroupGetResponseScopeComCloudflareAPIAccount, TokenPermissionGroupGetResponseScopeComCloudflareAPIAccountZone, TokenPermissionGroupGetResponseScopeComCloudflareAPIUser, TokenPermissionGroupGetResponseScopeComCloudflareEdgeR2Bucket:
		return true
	}
	return false
}

type TokenPermissionGroupListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter by the name of the permission group. The value must be URL-encoded.
	Name param.Field[string] `query:"name"`
	// Filter by the scope of the permission group. The value must be URL-encoded.
	Scope param.Field[string] `query:"scope"`
}

// URLQuery serializes [TokenPermissionGroupListParams]'s query parameters as
// `url.Values`.
func (r TokenPermissionGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TokenPermissionGroupGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter by the name of the permission group. The value must be URL-encoded.
	Name param.Field[string] `query:"name"`
	// Filter by the scope of the permission group. The value must be URL-encoded.
	Scope param.Field[string] `query:"scope"`
}

// URLQuery serializes [TokenPermissionGroupGetParams]'s query parameters as
// `url.Values`.
func (r TokenPermissionGroupGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type TokenPermissionGroupGetResponseEnvelope struct {
	Errors   []TokenPermissionGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []TokenPermissionGroupGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success    TokenPermissionGroupGetResponseEnvelopeSuccess    `json:"success,required"`
	Result     []TokenPermissionGroupGetResponse                 `json:"result"`
	ResultInfo TokenPermissionGroupGetResponseEnvelopeResultInfo `json:"result_info"`
	JSON       tokenPermissionGroupGetResponseEnvelopeJSON       `json:"-"`
}

// tokenPermissionGroupGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [TokenPermissionGroupGetResponseEnvelope]
type tokenPermissionGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupGetResponseEnvelopeErrors struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           TokenPermissionGroupGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             tokenPermissionGroupGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// tokenPermissionGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [TokenPermissionGroupGetResponseEnvelopeErrors]
type tokenPermissionGroupGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenPermissionGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    tokenPermissionGroupGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// tokenPermissionGroupGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [TokenPermissionGroupGetResponseEnvelopeErrorsSource]
type tokenPermissionGroupGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupGetResponseEnvelopeMessages struct {
	Code             int64                                                 `json:"code,required"`
	Message          string                                                `json:"message,required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           TokenPermissionGroupGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             tokenPermissionGroupGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// tokenPermissionGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [TokenPermissionGroupGetResponseEnvelopeMessages]
type tokenPermissionGroupGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenPermissionGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type TokenPermissionGroupGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    tokenPermissionGroupGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// tokenPermissionGroupGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [TokenPermissionGroupGetResponseEnvelopeMessagesSource]
type tokenPermissionGroupGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type TokenPermissionGroupGetResponseEnvelopeSuccess bool

const (
	TokenPermissionGroupGetResponseEnvelopeSuccessTrue TokenPermissionGroupGetResponseEnvelopeSuccess = true
)

func (r TokenPermissionGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case TokenPermissionGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type TokenPermissionGroupGetResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       tokenPermissionGroupGetResponseEnvelopeResultInfoJSON `json:"-"`
}

// tokenPermissionGroupGetResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [TokenPermissionGroupGetResponseEnvelopeResultInfo]
type tokenPermissionGroupGetResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPermissionGroupGetResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPermissionGroupGetResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
