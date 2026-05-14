// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v5/filters"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/rate_limits"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Create one or more firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) New(ctx context.Context, params RuleNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// Create one or more firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) NewAutoPaging(ctx context.Context, params RuleNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FirewallRule] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, params, opts...))
}

// Updates an existing firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Update(ctx context.Context, ruleID string, params RuleUpdateParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	var env RuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", params.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) List(ctx context.Context, params RuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
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

// Fetches firewall rules in a zone. You can filter the results using several
// optional parameters.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) ListAutoPaging(ctx context.Context, params RuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[FirewallRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Delete(ctx context.Context, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	var env RuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", body.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkDelete(ctx context.Context, body RuleBulkDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", body.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Deletes existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkDeleteAutoPaging(ctx context.Context, body RuleBulkDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FirewallRule] {
	return pagination.NewSinglePageAutoPager(r.BulkDelete(ctx, body, opts...))
}

// Updates the priority of existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkEdit(ctx context.Context, params RuleBulkEditParams, opts ...option.RequestOption) (res *pagination.SinglePage[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPatch, path, params, &res, opts...)
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

// Updates the priority of existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkEditAutoPaging(ctx context.Context, params RuleBulkEditParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FirewallRule] {
	return pagination.NewSinglePageAutoPager(r.BulkEdit(ctx, params, opts...))
}

// Updates one or more existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkUpdate(ctx context.Context, params RuleBulkUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Updates one or more existing firewall rules.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) BulkUpdateAutoPaging(ctx context.Context, params RuleBulkUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FirewallRule] {
	return pagination.NewSinglePageAutoPager(r.BulkUpdate(ctx, params, opts...))
}

// Updates the priority of an existing firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Edit(ctx context.Context, ruleID string, body RuleEditParams, opts ...option.RequestOption) (res *pagination.SinglePage[FirewallRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", body.ZoneID, ruleID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPatch, path, body, &res, opts...)
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

// Updates the priority of an existing firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) EditAutoPaging(ctx context.Context, ruleID string, body RuleEditParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[FirewallRule] {
	return pagination.NewSinglePageAutoPager(r.Edit(ctx, ruleID, body, opts...))
}

// Fetches the details of a firewall rule.
//
// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
// for full details.
func (r *RuleService) Get(ctx context.Context, ruleID string, query RuleGetParams, opts ...option.RequestOption) (res *FirewallRule, err error) {
	var env RuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/rules/%s", query.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeletedFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id,required"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool              `json:"deleted,required"`
	JSON    deletedFilterJSON `json:"-"`
}

// deletedFilterJSON contains the JSON metadata for the struct [DeletedFilter]
type deletedFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeletedFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deletedFilterJSON) RawJSON() string {
	return r.raw
}

func (r DeletedFilter) ImplementsFirewallRuleFilter() {}

type FirewallRule struct {
	// The unique identifier of the firewall rule.
	ID string `json:"id"`
	// The action to apply to a matched request. The `log` action is only available on
	// an Enterprise plan.
	Action rate_limits.Action `json:"action"`
	// An informative summary of the firewall rule.
	Description string             `json:"description"`
	Filter      FirewallRuleFilter `json:"filter"`
	// When true, indicates that the firewall rule is currently paused.
	Paused bool `json:"paused"`
	// The priority of the rule. Optional value used to define the processing order. A
	// lower number indicates a higher priority. If not provided, rules with a defined
	// priority will be processed before rules without a priority.
	Priority float64   `json:"priority"`
	Products []Product `json:"products"`
	// A short reference tag. Allows you to select related firewall rules.
	Ref  string           `json:"ref"`
	JSON firewallRuleJSON `json:"-"`
}

// firewallRuleJSON contains the JSON metadata for the struct [FirewallRule]
type firewallRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Description apijson.Field
	Filter      apijson.Field
	Paused      apijson.Field
	Priority    apijson.Field
	Products    apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallRuleJSON) RawJSON() string {
	return r.raw
}

type FirewallRuleFilter struct {
	// The unique identifier of the filter.
	ID string `json:"id"`
	// When true, indicates that the firewall rule was deleted.
	Deleted bool `json:"deleted"`
	// An informative summary of the filter.
	Description string `json:"description"`
	// The filter expression. For more information, refer to
	// [Expressions](https://developers.cloudflare.com/ruleset-engine/rules-language/expressions/).
	Expression string `json:"expression"`
	// When true, indicates that the filter is currently paused.
	Paused bool `json:"paused"`
	// A short reference tag. Allows you to select related filters.
	Ref   string                 `json:"ref"`
	JSON  firewallRuleFilterJSON `json:"-"`
	union FirewallRuleFilterUnion
}

// firewallRuleFilterJSON contains the JSON metadata for the struct
// [FirewallRuleFilter]
type firewallRuleFilterJSON struct {
	ID          apijson.Field
	Deleted     apijson.Field
	Description apijson.Field
	Expression  apijson.Field
	Paused      apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r firewallRuleFilterJSON) RawJSON() string {
	return r.raw
}

func (r *FirewallRuleFilter) UnmarshalJSON(data []byte) (err error) {
	*r = FirewallRuleFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [FirewallRuleFilterUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [filters.FirewallFilter],
// [DeletedFilter].
func (r FirewallRuleFilter) AsUnion() FirewallRuleFilterUnion {
	return r.union
}

// Union satisfied by [filters.FirewallFilter] or [DeletedFilter].
type FirewallRuleFilterUnion interface {
	ImplementsFirewallRuleFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*FirewallRuleFilterUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(filters.FirewallFilter{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(DeletedFilter{}),
		},
	)
}

// A list of products to bypass for a request when using the `bypass` action.
type Product string

const (
	ProductZoneLockdown  Product = "zoneLockdown"
	ProductUABlock       Product = "uaBlock"
	ProductBIC           Product = "bic"
	ProductHot           Product = "hot"
	ProductSecurityLevel Product = "securityLevel"
	ProductRateLimit     Product = "rateLimit"
	ProductWAF           Product = "waf"
)

func (r Product) IsKnown() bool {
	switch r {
	case ProductZoneLockdown, ProductUABlock, ProductBIC, ProductHot, ProductSecurityLevel, ProductRateLimit, ProductWAF:
		return true
	}
	return false
}

type RuleNewParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RuleNewParamsAction]         `json:"action,required"`
	Filter param.Field[filters.FirewallFilterParam] `json:"filter,required"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RuleNewParamsAction struct {
	// The action to perform.
	Mode param.Field[RuleNewParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RuleNewParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RuleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RuleNewParamsActionMode string

const (
	RuleNewParamsActionModeSimulate         RuleNewParamsActionMode = "simulate"
	RuleNewParamsActionModeBan              RuleNewParamsActionMode = "ban"
	RuleNewParamsActionModeChallenge        RuleNewParamsActionMode = "challenge"
	RuleNewParamsActionModeJSChallenge      RuleNewParamsActionMode = "js_challenge"
	RuleNewParamsActionModeManagedChallenge RuleNewParamsActionMode = "managed_challenge"
)

func (r RuleNewParamsActionMode) IsKnown() bool {
	switch r {
	case RuleNewParamsActionModeSimulate, RuleNewParamsActionModeBan, RuleNewParamsActionModeChallenge, RuleNewParamsActionModeJSChallenge, RuleNewParamsActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RuleNewParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RuleNewParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The action to perform when the threshold of matched traffic within the
	// configured period is exceeded.
	Action param.Field[RuleUpdateParamsAction]      `json:"action,required"`
	Filter param.Field[filters.FirewallFilterParam] `json:"filter,required"`
}

func (r RuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the threshold of matched traffic within the
// configured period is exceeded.
type RuleUpdateParamsAction struct {
	// The action to perform.
	Mode param.Field[RuleUpdateParamsActionMode] `json:"mode"`
	// A custom content type and reponse to return when the threshold is exceeded. The
	// custom response configured in this object will override the custom error for the
	// zone. This object is optional. Notes: If you omit this object, Cloudflare will
	// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
	// or "js_challenge", Cloudflare will use the zone challenge pages and you should
	// not provide the "response" object.
	Response param.Field[RuleUpdateParamsActionResponse] `json:"response"`
	// The time in seconds during which Cloudflare will perform the mitigation action.
	// Must be an integer value greater than or equal to the period. Notes: If "mode"
	// is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the
	// zone's Challenge Passage time and you should not provide this value.
	Timeout param.Field[float64] `json:"timeout"`
}

func (r RuleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform.
type RuleUpdateParamsActionMode string

const (
	RuleUpdateParamsActionModeSimulate         RuleUpdateParamsActionMode = "simulate"
	RuleUpdateParamsActionModeBan              RuleUpdateParamsActionMode = "ban"
	RuleUpdateParamsActionModeChallenge        RuleUpdateParamsActionMode = "challenge"
	RuleUpdateParamsActionModeJSChallenge      RuleUpdateParamsActionMode = "js_challenge"
	RuleUpdateParamsActionModeManagedChallenge RuleUpdateParamsActionMode = "managed_challenge"
)

func (r RuleUpdateParamsActionMode) IsKnown() bool {
	switch r {
	case RuleUpdateParamsActionModeSimulate, RuleUpdateParamsActionModeBan, RuleUpdateParamsActionModeChallenge, RuleUpdateParamsActionModeJSChallenge, RuleUpdateParamsActionModeManagedChallenge:
		return true
	}
	return false
}

// A custom content type and reponse to return when the threshold is exceeded. The
// custom response configured in this object will override the custom error for the
// zone. This object is optional. Notes: If you omit this object, Cloudflare will
// use the default HTML error page. If "mode" is "challenge", "managed_challenge",
// or "js_challenge", Cloudflare will use the zone challenge pages and you should
// not provide the "response" object.
type RuleUpdateParamsActionResponse struct {
	// The response body to return. The value must conform to the configured content
	// type.
	Body param.Field[string] `json:"body"`
	// The content type of the body. Must be one of the following: `text/plain`,
	// `text/xml`, or `application/json`.
	ContentType param.Field[string] `json:"content_type"`
}

func (r RuleUpdateParamsActionResponse) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallRule          `json:"result,required"`
	// Defines whether the API call was successful.
	Success RuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// ruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleUpdateResponseEnvelope]
type ruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type RuleUpdateResponseEnvelopeSuccess bool

const (
	RuleUpdateResponseEnvelopeSuccessTrue RuleUpdateResponseEnvelopeSuccess = true
)

func (r RuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleListParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The unique identifier of the firewall rule.
	ID param.Field[string] `query:"id"`
	// The action to search for. Must be an exact match.
	Action param.Field[string] `query:"action"`
	// A case-insensitive string to find in the description.
	Description param.Field[string] `query:"description"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// When true, indicates that the firewall rule is currently paused.
	Paused param.Field[bool] `query:"paused"`
	// Number of firewall rules per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type RuleDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallRule          `json:"result,required"`
	// Defines whether the API call was successful.
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleBulkDeleteParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleBulkEditParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r RuleBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleBulkUpdateParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	Body   interface{}         `json:"body,required"`
}

func (r RuleBulkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleEditParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleGetParams struct {
	// Defines an identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   FirewallRule          `json:"result,required"`
	// Defines whether the API call was successful.
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
