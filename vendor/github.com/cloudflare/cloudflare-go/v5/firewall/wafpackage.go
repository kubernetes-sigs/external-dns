// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// WAFPackageService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFPackageService] method instead.
type WAFPackageService struct {
	Options []option.RequestOption
	Groups  *WAFPackageGroupService
	Rules   *WAFPackageRuleService
}

// NewWAFPackageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFPackageService(opts ...option.RequestOption) (r *WAFPackageService) {
	r = &WAFPackageService{}
	r.Options = opts
	r.Groups = NewWAFPackageGroupService(opts...)
	r.Rules = NewWAFPackageRuleService(opts...)
	return
}

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageService) List(ctx context.Context, params WAFPackageListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WAFPackageListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages", params.ZoneID)
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

// Fetches WAF packages for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageService) ListAutoPaging(ctx context.Context, params WAFPackageListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WAFPackageListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Fetches the details of a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageService) Get(ctx context.Context, packageID string, query WAFPackageGetParams, opts ...option.RequestOption) (res *WAFPackageGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s", query.ZoneID, packageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type WAFPackageListResponse = interface{}

type WAFPackageGetResponse struct {
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Errors interface{} `json:"errors"`
	// This field can have the runtime type of [[]shared.ResponseInfo].
	Messages interface{} `json:"messages"`
	// This field can have the runtime type of [interface{}].
	Result interface{} `json:"result"`
	// Defines whether the API call was successful.
	Success WAFPackageGetResponseSuccess `json:"success"`
	JSON    wafPackageGetResponseJSON    `json:"-"`
	union   WAFPackageGetResponseUnion
}

// wafPackageGetResponseJSON contains the JSON metadata for the struct
// [WAFPackageGetResponse]
type wafPackageGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r wafPackageGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = WAFPackageGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [WAFPackageGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are
// [WAFPackageGetResponseFirewallAPIResponseSingle], [WAFPackageGetResponseResult].
func (r WAFPackageGetResponse) AsUnion() WAFPackageGetResponseUnion {
	return r.union
}

// Union satisfied by [WAFPackageGetResponseFirewallAPIResponseSingle] or
// [WAFPackageGetResponseResult].
type WAFPackageGetResponseUnion interface {
	implementsWAFPackageGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageGetResponseFirewallAPIResponseSingle{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageGetResponseResult{}),
		},
	)
}

type WAFPackageGetResponseFirewallAPIResponseSingle struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFPackageGetResponseFirewallAPIResponseSingleSuccess `json:"success,required"`
	JSON    wafPackageGetResponseFirewallAPIResponseSingleJSON    `json:"-"`
}

// wafPackageGetResponseFirewallAPIResponseSingleJSON contains the JSON metadata
// for the struct [WAFPackageGetResponseFirewallAPIResponseSingle]
type wafPackageGetResponseFirewallAPIResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGetResponseFirewallAPIResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGetResponseFirewallAPIResponseSingleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageGetResponseFirewallAPIResponseSingle) implementsWAFPackageGetResponse() {}

// Defines whether the API call was successful.
type WAFPackageGetResponseFirewallAPIResponseSingleSuccess bool

const (
	WAFPackageGetResponseFirewallAPIResponseSingleSuccessTrue WAFPackageGetResponseFirewallAPIResponseSingleSuccess = true
)

func (r WAFPackageGetResponseFirewallAPIResponseSingleSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGetResponseFirewallAPIResponseSingleSuccessTrue:
		return true
	}
	return false
}

type WAFPackageGetResponseResult struct {
	Result interface{}                     `json:"result"`
	JSON   wafPackageGetResponseResultJSON `json:"-"`
}

// wafPackageGetResponseResultJSON contains the JSON metadata for the struct
// [WAFPackageGetResponseResult]
type wafPackageGetResponseResultJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGetResponseResultJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageGetResponseResult) implementsWAFPackageGetResponse() {}

// Defines whether the API call was successful.
type WAFPackageGetResponseSuccess bool

const (
	WAFPackageGetResponseSuccessTrue WAFPackageGetResponseSuccess = true
)

func (r WAFPackageGetResponseSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGetResponseSuccessTrue:
		return true
	}
	return false
}

type WAFPackageListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned packages.
	Direction param.Field[WAFPackageListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[WAFPackageListParamsMatch] `query:"match"`
	// The name of the WAF package.
	Name param.Field[string] `query:"name"`
	// The field used to sort returned packages.
	Order param.Field[WAFPackageListParamsOrder] `query:"order"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of packages per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WAFPackageListParams]'s query parameters as `url.Values`.
func (r WAFPackageListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The direction used to sort returned packages.
type WAFPackageListParamsDirection string

const (
	WAFPackageListParamsDirectionAsc  WAFPackageListParamsDirection = "asc"
	WAFPackageListParamsDirectionDesc WAFPackageListParamsDirection = "desc"
)

func (r WAFPackageListParamsDirection) IsKnown() bool {
	switch r {
	case WAFPackageListParamsDirectionAsc, WAFPackageListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type WAFPackageListParamsMatch string

const (
	WAFPackageListParamsMatchAny WAFPackageListParamsMatch = "any"
	WAFPackageListParamsMatchAll WAFPackageListParamsMatch = "all"
)

func (r WAFPackageListParamsMatch) IsKnown() bool {
	switch r {
	case WAFPackageListParamsMatchAny, WAFPackageListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned packages.
type WAFPackageListParamsOrder string

const (
	WAFPackageListParamsOrderName WAFPackageListParamsOrder = "name"
)

func (r WAFPackageListParamsOrder) IsKnown() bool {
	switch r {
	case WAFPackageListParamsOrderName:
		return true
	}
	return false
}

type WAFPackageGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}
