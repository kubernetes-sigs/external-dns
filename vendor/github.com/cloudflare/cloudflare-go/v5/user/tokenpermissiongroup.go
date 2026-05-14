// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
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

// Find all available permission groups for API Tokens
func (r *TokenPermissionGroupService) List(ctx context.Context, query TokenPermissionGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[TokenPermissionGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/tokens/permission_groups"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Find all available permission groups for API Tokens
func (r *TokenPermissionGroupService) ListAutoPaging(ctx context.Context, query TokenPermissionGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TokenPermissionGroupListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
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

type TokenPermissionGroupListParams struct {
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
