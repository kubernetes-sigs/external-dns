// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

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

// SearchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Search for Load Balancing resources.
func (r *SearchService) List(ctx context.Context, params SearchListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[SearchListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/search", params.AccountID)
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

// Search for Load Balancing resources.
func (r *SearchService) ListAutoPaging(ctx context.Context, params SearchListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[SearchListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

type SearchListResponse struct {
	// A list of resources matching the search query.
	Resources []SearchListResponseResource `json:"resources"`
	JSON      searchListResponseJSON       `json:"-"`
}

// searchListResponseJSON contains the JSON metadata for the struct
// [SearchListResponse]
type searchListResponseJSON struct {
	Resources   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SearchListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchListResponseJSON) RawJSON() string {
	return r.raw
}

// A reference to a load balancer resource.
type SearchListResponseResource struct {
	// When listed as a reference, the type (direction) of the reference.
	ReferenceType SearchListResponseResourcesReferenceType `json:"reference_type"`
	// A list of references to (referrer) or from (referral) this resource.
	References []interface{} `json:"references"`
	ResourceID string        `json:"resource_id"`
	// The human-identifiable name of the resource.
	ResourceName string `json:"resource_name"`
	// The type of the resource.
	ResourceType SearchListResponseResourcesResourceType `json:"resource_type"`
	JSON         searchListResponseResourceJSON          `json:"-"`
}

// searchListResponseResourceJSON contains the JSON metadata for the struct
// [SearchListResponseResource]
type searchListResponseResourceJSON struct {
	ReferenceType apijson.Field
	References    apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SearchListResponseResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r searchListResponseResourceJSON) RawJSON() string {
	return r.raw
}

// When listed as a reference, the type (direction) of the reference.
type SearchListResponseResourcesReferenceType string

const (
	SearchListResponseResourcesReferenceTypeReferral SearchListResponseResourcesReferenceType = "referral"
	SearchListResponseResourcesReferenceTypeReferrer SearchListResponseResourcesReferenceType = "referrer"
)

func (r SearchListResponseResourcesReferenceType) IsKnown() bool {
	switch r {
	case SearchListResponseResourcesReferenceTypeReferral, SearchListResponseResourcesReferenceTypeReferrer:
		return true
	}
	return false
}

// The type of the resource.
type SearchListResponseResourcesResourceType string

const (
	SearchListResponseResourcesResourceTypeLoadBalancer SearchListResponseResourcesResourceType = "load_balancer"
	SearchListResponseResourcesResourceTypeMonitor      SearchListResponseResourcesResourceType = "monitor"
	SearchListResponseResourcesResourceTypePool         SearchListResponseResourcesResourceType = "pool"
)

func (r SearchListResponseResourcesResourceType) IsKnown() bool {
	switch r {
	case SearchListResponseResourcesResourceTypeLoadBalancer, SearchListResponseResourcesResourceTypeMonitor, SearchListResponseResourcesResourceTypePool:
		return true
	}
	return false
}

type SearchListParams struct {
	// Identifier.
	AccountID param.Field[string]  `path:"account_id,required"`
	Page      param.Field[float64] `query:"page"`
	PerPage   param.Field[float64] `query:"per_page"`
	// Search query term.
	Query param.Field[string] `query:"query"`
	// The type of references to include. "\*" to include both referral and referrer
	// references. "" to not include any reference information.
	References param.Field[SearchListParamsReferences] `query:"references"`
}

// URLQuery serializes [SearchListParams]'s query parameters as `url.Values`.
func (r SearchListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The type of references to include. "\*" to include both referral and referrer
// references. "" to not include any reference information.
type SearchListParamsReferences string

const (
	SearchListParamsReferencesEmpty    SearchListParamsReferences = ""
	SearchListParamsReferencesStar     SearchListParamsReferences = "*"
	SearchListParamsReferencesReferral SearchListParamsReferences = "referral"
	SearchListParamsReferencesReferrer SearchListParamsReferences = "referrer"
)

func (r SearchListParamsReferences) IsKnown() bool {
	switch r {
	case SearchListParamsReferencesEmpty, SearchListParamsReferencesStar, SearchListParamsReferencesReferral, SearchListParamsReferencesReferrer:
		return true
	}
	return false
}
