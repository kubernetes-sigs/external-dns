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

// WAFOverrideService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFOverrideService] method instead.
type WAFOverrideService struct {
	Options []option.RequestOption
}

// NewWAFOverrideService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFOverrideService(opts ...option.RequestOption) (r *WAFOverrideService) {
	r = &WAFOverrideService{}
	r.Options = opts
	return
}

// Creates a URI-based WAF override for a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFOverrideService) New(ctx context.Context, params WAFOverrideNewParams, opts ...option.RequestOption) (res *Override, err error) {
	var env WAFOverrideNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFOverrideService) Update(ctx context.Context, overridesID string, params WAFOverrideUpdateParams, opts ...option.RequestOption) (res *Override, err error) {
	var env WAFOverrideUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", params.ZoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFOverrideService) List(ctx context.Context, params WAFOverrideListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Override], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides", params.ZoneID)
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

// Fetches the URI-based WAF overrides in a zone.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFOverrideService) ListAutoPaging(ctx context.Context, params WAFOverrideListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Override] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFOverrideService) Delete(ctx context.Context, overridesID string, body WAFOverrideDeleteParams, opts ...option.RequestOption) (res *WAFOverrideDeleteResponse, err error) {
	var env WAFOverrideDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", body.ZoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a URI-based WAF override.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFOverrideService) Get(ctx context.Context, overridesID string, query WAFOverrideGetParams, opts ...option.RequestOption) (res *Override, err error) {
	var env WAFOverrideGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if overridesID == "" {
		err = errors.New("missing required overrides_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/overrides/%s", query.ZoneID, overridesID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Override struct {
	// The unique identifier of the WAF override.
	ID string `json:"id"`
	// An informative summary of the current URI-based WAF override.
	Description string `json:"description,nullable"`
	// An object that allows you to enable or disable WAF rule groups for the current
	// WAF override. Each key of this object must be the ID of a WAF rule group, and
	// each value must be a valid WAF action (usually `default` or `disable`). When
	// creating a new URI-based WAF override, you must provide a `groups` object or a
	// `rules` object.
	Groups map[string]interface{} `json:"groups"`
	// When true, indicates that the rule is currently paused.
	Paused bool `json:"paused"`
	// The relative priority of the current URI-based WAF override when multiple
	// overrides match a single URL. A lower number indicates higher priority. Higher
	// priority overrides may overwrite values set by lower priority overrides.
	Priority float64 `json:"priority"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction RewriteAction `json:"rewrite_action"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules WAFRule `json:"rules"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs []OverrideURL `json:"urls"`
	JSON overrideJSON  `json:"-"`
}

// overrideJSON contains the JSON metadata for the struct [Override]
type overrideJSON struct {
	ID            apijson.Field
	Description   apijson.Field
	Groups        apijson.Field
	Paused        apijson.Field
	Priority      apijson.Field
	RewriteAction apijson.Field
	Rules         apijson.Field
	URLs          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r overrideJSON) RawJSON() string {
	return r.raw
}

type OverrideURL = string

type OverrideURLParam = string

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type RewriteAction struct {
	// The WAF rule action to apply.
	Block RewriteActionBlock `json:"block"`
	// The WAF rule action to apply.
	Challenge RewriteActionChallenge `json:"challenge"`
	// The WAF rule action to apply.
	Default RewriteActionDefault `json:"default"`
	// The WAF rule action to apply.
	Disable RewriteActionDisable `json:"disable"`
	// The WAF rule action to apply.
	Simulate RewriteActionSimulate `json:"simulate"`
	JSON     rewriteActionJSON     `json:"-"`
}

// rewriteActionJSON contains the JSON metadata for the struct [RewriteAction]
type rewriteActionJSON struct {
	Block       apijson.Field
	Challenge   apijson.Field
	Default     apijson.Field
	Disable     apijson.Field
	Simulate    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteActionJSON) RawJSON() string {
	return r.raw
}

// The WAF rule action to apply.
type RewriteActionBlock string

const (
	RewriteActionBlockChallenge RewriteActionBlock = "challenge"
	RewriteActionBlockBlock     RewriteActionBlock = "block"
	RewriteActionBlockSimulate  RewriteActionBlock = "simulate"
	RewriteActionBlockDisable   RewriteActionBlock = "disable"
	RewriteActionBlockDefault   RewriteActionBlock = "default"
)

func (r RewriteActionBlock) IsKnown() bool {
	switch r {
	case RewriteActionBlockChallenge, RewriteActionBlockBlock, RewriteActionBlockSimulate, RewriteActionBlockDisable, RewriteActionBlockDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionChallenge string

const (
	RewriteActionChallengeChallenge RewriteActionChallenge = "challenge"
	RewriteActionChallengeBlock     RewriteActionChallenge = "block"
	RewriteActionChallengeSimulate  RewriteActionChallenge = "simulate"
	RewriteActionChallengeDisable   RewriteActionChallenge = "disable"
	RewriteActionChallengeDefault   RewriteActionChallenge = "default"
)

func (r RewriteActionChallenge) IsKnown() bool {
	switch r {
	case RewriteActionChallengeChallenge, RewriteActionChallengeBlock, RewriteActionChallengeSimulate, RewriteActionChallengeDisable, RewriteActionChallengeDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionDefault string

const (
	RewriteActionDefaultChallenge RewriteActionDefault = "challenge"
	RewriteActionDefaultBlock     RewriteActionDefault = "block"
	RewriteActionDefaultSimulate  RewriteActionDefault = "simulate"
	RewriteActionDefaultDisable   RewriteActionDefault = "disable"
	RewriteActionDefaultDefault   RewriteActionDefault = "default"
)

func (r RewriteActionDefault) IsKnown() bool {
	switch r {
	case RewriteActionDefaultChallenge, RewriteActionDefaultBlock, RewriteActionDefaultSimulate, RewriteActionDefaultDisable, RewriteActionDefaultDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionDisable string

const (
	RewriteActionDisableChallenge RewriteActionDisable = "challenge"
	RewriteActionDisableBlock     RewriteActionDisable = "block"
	RewriteActionDisableSimulate  RewriteActionDisable = "simulate"
	RewriteActionDisableDisable   RewriteActionDisable = "disable"
	RewriteActionDisableDefault   RewriteActionDisable = "default"
)

func (r RewriteActionDisable) IsKnown() bool {
	switch r {
	case RewriteActionDisableChallenge, RewriteActionDisableBlock, RewriteActionDisableSimulate, RewriteActionDisableDisable, RewriteActionDisableDefault:
		return true
	}
	return false
}

// The WAF rule action to apply.
type RewriteActionSimulate string

const (
	RewriteActionSimulateChallenge RewriteActionSimulate = "challenge"
	RewriteActionSimulateBlock     RewriteActionSimulate = "block"
	RewriteActionSimulateSimulate  RewriteActionSimulate = "simulate"
	RewriteActionSimulateDisable   RewriteActionSimulate = "disable"
	RewriteActionSimulateDefault   RewriteActionSimulate = "default"
)

func (r RewriteActionSimulate) IsKnown() bool {
	switch r {
	case RewriteActionSimulateChallenge, RewriteActionSimulateBlock, RewriteActionSimulateSimulate, RewriteActionSimulateDisable, RewriteActionSimulateDefault:
		return true
	}
	return false
}

// Specifies that, when a WAF rule matches, its configured action will be replaced
// by the action configured in this object.
type RewriteActionParam struct {
	// The WAF rule action to apply.
	Block param.Field[RewriteActionBlock] `json:"block"`
	// The WAF rule action to apply.
	Challenge param.Field[RewriteActionChallenge] `json:"challenge"`
	// The WAF rule action to apply.
	Default param.Field[RewriteActionDefault] `json:"default"`
	// The WAF rule action to apply.
	Disable param.Field[RewriteActionDisable] `json:"disable"`
	// The WAF rule action to apply.
	Simulate param.Field[RewriteActionSimulate] `json:"simulate"`
}

func (r RewriteActionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WAFRule map[string]WAFRuleItem

// The WAF rule action to apply.
type WAFRuleItem string

const (
	WAFRuleItemChallenge WAFRuleItem = "challenge"
	WAFRuleItemBlock     WAFRuleItem = "block"
	WAFRuleItemSimulate  WAFRuleItem = "simulate"
	WAFRuleItemDisable   WAFRuleItem = "disable"
	WAFRuleItemDefault   WAFRuleItem = "default"
)

func (r WAFRuleItem) IsKnown() bool {
	switch r {
	case WAFRuleItemChallenge, WAFRuleItemBlock, WAFRuleItemSimulate, WAFRuleItemDisable, WAFRuleItemDefault:
		return true
	}
	return false
}

type WAFRuleParam map[string]WAFRuleItem

type WAFOverrideDeleteResponse struct {
	// The unique identifier of the WAF override.
	ID   string                        `json:"id"`
	JSON wafOverrideDeleteResponseJSON `json:"-"`
}

// wafOverrideDeleteResponseJSON contains the JSON metadata for the struct
// [WAFOverrideDeleteResponse]
type wafOverrideDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type WAFOverrideNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]OverrideURLParam] `json:"urls,required"`
}

func (r WAFOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WAFOverrideNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Override              `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFOverrideNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafOverrideNewResponseEnvelopeJSON    `json:"-"`
}

// wafOverrideNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideNewResponseEnvelope]
type wafOverrideNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFOverrideNewResponseEnvelopeSuccess bool

const (
	WAFOverrideNewResponseEnvelopeSuccessTrue WAFOverrideNewResponseEnvelopeSuccess = true
)

func (r WAFOverrideNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFOverrideNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WAFOverrideUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines an identifier.
	ID param.Field[string] `json:"id,required"`
	// Specifies that, when a WAF rule matches, its configured action will be replaced
	// by the action configured in this object.
	RewriteAction param.Field[RewriteActionParam] `json:"rewrite_action,required"`
	// An object that allows you to override the action of specific WAF rules. Each key
	// of this object must be the ID of a WAF rule, and each value must be a valid WAF
	// action. Unless you are disabling a rule, ensure that you also enable the rule
	// group that this WAF rule belongs to. When creating a new URI-based WAF override,
	// you must provide a `groups` object or a `rules` object.
	Rules param.Field[WAFRuleParam] `json:"rules,required"`
	// The URLs to include in the current WAF override. You can use wildcards. Each
	// entered URL will be escaped before use, which means you can only use simple
	// wildcard patterns.
	URLs param.Field[[]OverrideURLParam] `json:"urls,required"`
}

func (r WAFOverrideUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type WAFOverrideUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Override              `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFOverrideUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafOverrideUpdateResponseEnvelopeJSON    `json:"-"`
}

// wafOverrideUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideUpdateResponseEnvelope]
type wafOverrideUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFOverrideUpdateResponseEnvelopeSuccess bool

const (
	WAFOverrideUpdateResponseEnvelopeSuccessTrue WAFOverrideUpdateResponseEnvelopeSuccess = true
)

func (r WAFOverrideUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFOverrideUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WAFOverrideListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// The number of WAF overrides per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [WAFOverrideListParams]'s query parameters as `url.Values`.
func (r WAFOverrideListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type WAFOverrideDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WAFOverrideDeleteResponseEnvelope struct {
	Result WAFOverrideDeleteResponse             `json:"result"`
	JSON   wafOverrideDeleteResponseEnvelopeJSON `json:"-"`
}

// wafOverrideDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideDeleteResponseEnvelope]
type wafOverrideDeleteResponseEnvelopeJSON struct {
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WAFOverrideGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WAFOverrideGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   Override              `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFOverrideGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafOverrideGetResponseEnvelopeJSON    `json:"-"`
}

// wafOverrideGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFOverrideGetResponseEnvelope]
type wafOverrideGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFOverrideGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafOverrideGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFOverrideGetResponseEnvelopeSuccess bool

const (
	WAFOverrideGetResponseEnvelopeSuccessTrue WAFOverrideGetResponseEnvelopeSuccess = true
)

func (r WAFOverrideGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFOverrideGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
