// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// WAFPackageGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFPackageGroupService] method instead.
type WAFPackageGroupService struct {
	Options []option.RequestOption
}

// NewWAFPackageGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFPackageGroupService(opts ...option.RequestOption) (r *WAFPackageGroupService) {
	r = &WAFPackageGroupService{}
	r.Options = opts
	return
}

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageGroupService) List(ctx context.Context, packageID string, params WAFPackageGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Group], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups", params.ZoneID, packageID)
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

// Fetches the WAF rule groups in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageGroupService) ListAutoPaging(ctx context.Context, packageID string, params WAFPackageGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Group] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, packageID, params, opts...))
}

// Updates a WAF rule group. You can update the state (`mode` parameter) of a rule
// group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageGroupService) Edit(ctx context.Context, packageID string, groupID string, params WAFPackageGroupEditParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env WAFPackageGroupEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", params.ZoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a WAF rule group.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageGroupService) Get(ctx context.Context, packageID string, groupID string, query WAFPackageGroupGetParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env WAFPackageGroupGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/groups/%s", query.ZoneID, packageID, groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Group struct {
	// Defines the unique identifier of the rule group.
	ID string `json:"id,required"`
	// Defines an informative summary of what the rule group does.
	Description string `json:"description,required,nullable"`
	// Defines the state of the rules contained in the rule group. When `on`, the rules
	// in the group are configurable/usable.
	Mode GroupMode `json:"mode,required"`
	// Defines the name of the rule group.
	Name string `json:"name,required"`
	// Defines the number of rules in the current rule group.
	RulesCount float64 `json:"rules_count,required"`
	// Defines the available states for the rule group.
	AllowedModes []GroupAllowedMode `json:"allowed_modes"`
	// Defines the number of rules within the group that have been modified from their
	// default configuration.
	ModifiedRulesCount float64 `json:"modified_rules_count"`
	// Defines the unique identifier of a WAF package.
	PackageID string    `json:"package_id"`
	JSON      groupJSON `json:"-"`
}

// groupJSON contains the JSON metadata for the struct [Group]
type groupJSON struct {
	ID                 apijson.Field
	Description        apijson.Field
	Mode               apijson.Field
	Name               apijson.Field
	RulesCount         apijson.Field
	AllowedModes       apijson.Field
	ModifiedRulesCount apijson.Field
	PackageID          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Group) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r groupJSON) RawJSON() string {
	return r.raw
}

// Defines the state of the rules contained in the rule group. When `on`, the rules
// in the group are configurable/usable.
type GroupMode string

const (
	GroupModeOn  GroupMode = "on"
	GroupModeOff GroupMode = "off"
)

func (r GroupMode) IsKnown() bool {
	switch r {
	case GroupModeOn, GroupModeOff:
		return true
	}
	return false
}

// Defines the state of the rules contained in the rule group. When `on`, the rules
// in the group are configurable/usable.
type GroupAllowedMode string

const (
	GroupAllowedModeOn  GroupAllowedMode = "on"
	GroupAllowedModeOff GroupAllowedMode = "off"
)

func (r GroupAllowedMode) IsKnown() bool {
	switch r {
	case GroupAllowedModeOn, GroupAllowedModeOff:
		return true
	}
	return false
}

type WAFPackageGroupListParams struct {
	// Defines an identifier of a schema.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines the direction used to sort returned rule groups.
	Direction param.Field[WAFPackageGroupListParamsDirection] `query:"direction"`
	// Defines the condition for search requirements. When set to `all`, all the search
	// requirements must match. When set to `any`, only one of the search requirements
	// has to match.
	Match param.Field[WAFPackageGroupListParamsMatch] `query:"match"`
	// Defines the state of the rules contained in the rule group. When `on`, the rules
	// in the group are configurable/usable.
	Mode param.Field[WAFPackageGroupListParamsMode] `query:"mode"`
	// Defines the name of the rule group.
	Name param.Field[string] `query:"name"`
	// Defines the field used to sort returned rule groups.
	Order param.Field[WAFPackageGroupListParamsOrder] `query:"order"`
	// Defines the page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Defines the number of rule groups per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Defines the number of rules in the current rule group.
	RulesCount param.Field[float64] `query:"rules_count"`
}

// URLQuery serializes [WAFPackageGroupListParams]'s query parameters as
// `url.Values`.
func (r WAFPackageGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Defines the direction used to sort returned rule groups.
type WAFPackageGroupListParamsDirection string

const (
	WAFPackageGroupListParamsDirectionAsc  WAFPackageGroupListParamsDirection = "asc"
	WAFPackageGroupListParamsDirectionDesc WAFPackageGroupListParamsDirection = "desc"
)

func (r WAFPackageGroupListParamsDirection) IsKnown() bool {
	switch r {
	case WAFPackageGroupListParamsDirectionAsc, WAFPackageGroupListParamsDirectionDesc:
		return true
	}
	return false
}

// Defines the condition for search requirements. When set to `all`, all the search
// requirements must match. When set to `any`, only one of the search requirements
// has to match.
type WAFPackageGroupListParamsMatch string

const (
	WAFPackageGroupListParamsMatchAny WAFPackageGroupListParamsMatch = "any"
	WAFPackageGroupListParamsMatchAll WAFPackageGroupListParamsMatch = "all"
)

func (r WAFPackageGroupListParamsMatch) IsKnown() bool {
	switch r {
	case WAFPackageGroupListParamsMatchAny, WAFPackageGroupListParamsMatchAll:
		return true
	}
	return false
}

// Defines the state of the rules contained in the rule group. When `on`, the rules
// in the group are configurable/usable.
type WAFPackageGroupListParamsMode string

const (
	WAFPackageGroupListParamsModeOn  WAFPackageGroupListParamsMode = "on"
	WAFPackageGroupListParamsModeOff WAFPackageGroupListParamsMode = "off"
)

func (r WAFPackageGroupListParamsMode) IsKnown() bool {
	switch r {
	case WAFPackageGroupListParamsModeOn, WAFPackageGroupListParamsModeOff:
		return true
	}
	return false
}

// Defines the field used to sort returned rule groups.
type WAFPackageGroupListParamsOrder string

const (
	WAFPackageGroupListParamsOrderMode       WAFPackageGroupListParamsOrder = "mode"
	WAFPackageGroupListParamsOrderRulesCount WAFPackageGroupListParamsOrder = "rules_count"
)

func (r WAFPackageGroupListParamsOrder) IsKnown() bool {
	switch r {
	case WAFPackageGroupListParamsOrderMode, WAFPackageGroupListParamsOrderRulesCount:
		return true
	}
	return false
}

type WAFPackageGroupEditParams struct {
	// Defines an identifier of a schema.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines the state of the rules contained in the rule group. When `on`, the rules
	// in the group are configurable/usable.
	Mode param.Field[WAFPackageGroupEditParamsMode] `json:"mode"`
}

func (r WAFPackageGroupEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the state of the rules contained in the rule group. When `on`, the rules
// in the group are configurable/usable.
type WAFPackageGroupEditParamsMode string

const (
	WAFPackageGroupEditParamsModeOn  WAFPackageGroupEditParamsMode = "on"
	WAFPackageGroupEditParamsModeOff WAFPackageGroupEditParamsMode = "off"
)

func (r WAFPackageGroupEditParamsMode) IsKnown() bool {
	switch r {
	case WAFPackageGroupEditParamsModeOn, WAFPackageGroupEditParamsModeOff:
		return true
	}
	return false
}

type WAFPackageGroupEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFPackageGroupEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafPackageGroupEditResponseEnvelopeJSON    `json:"-"`
}

// wafPackageGroupEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [WAFPackageGroupEditResponseEnvelope]
type wafPackageGroupEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGroupEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGroupEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFPackageGroupEditResponseEnvelopeSuccess bool

const (
	WAFPackageGroupEditResponseEnvelopeSuccessTrue WAFPackageGroupEditResponseEnvelopeSuccess = true
)

func (r WAFPackageGroupEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGroupEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WAFPackageGroupGetParams struct {
	// Defines an identifier of a schema.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WAFPackageGroupGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFPackageGroupGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafPackageGroupGetResponseEnvelopeJSON    `json:"-"`
}

// wafPackageGroupGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFPackageGroupGetResponseEnvelope]
type wafPackageGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFPackageGroupGetResponseEnvelopeSuccess bool

const (
	WAFPackageGroupGetResponseEnvelopeSuccessTrue WAFPackageGroupGetResponseEnvelopeSuccess = true
)

func (r WAFPackageGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFPackageGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
