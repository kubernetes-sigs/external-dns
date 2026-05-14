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

// WAFPackageRuleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFPackageRuleService] method instead.
type WAFPackageRuleService struct {
	Options []option.RequestOption
}

// NewWAFPackageRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWAFPackageRuleService(opts ...option.RequestOption) (r *WAFPackageRuleService) {
	r = &WAFPackageRuleService{}
	r.Options = opts
	return
}

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageRuleService) List(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[WAFPackageRuleListResponse], err error) {
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
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules", params.ZoneID, packageID)
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

// Fetches WAF rules in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageRuleService) ListAutoPaging(ctx context.Context, packageID string, params WAFPackageRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[WAFPackageRuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, packageID, params, opts...))
}

// Updates a WAF rule. You can only update the mode/action of the rule.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageRuleService) Edit(ctx context.Context, packageID string, ruleID string, params WAFPackageRuleEditParams, opts ...option.RequestOption) (res *WAFPackageRuleEditResponse, err error) {
	var env WAFPackageRuleEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", params.ZoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a WAF rule in a WAF package.
//
// **Note:** Applies only to the
// [previous version of WAF managed rules](https://developers.cloudflare.com/support/firewall/managed-rules-web-application-firewall-waf/understanding-waf-managed-rules-web-application-firewall/).
//
// Deprecated: deprecated
func (r *WAFPackageRuleService) Get(ctx context.Context, packageID string, ruleID string, query WAFPackageRuleGetParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env WAFPackageRuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if packageID == "" {
		err = errors.New("missing required package_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/firewall/waf/packages/%s/rules/%s", query.ZoneID, packageID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
// when evaluating the request. Applies to anomaly detection WAF rules.
type AllowedModesAnomaly string

const (
	AllowedModesAnomalyOn  AllowedModesAnomaly = "on"
	AllowedModesAnomalyOff AllowedModesAnomaly = "off"
)

func (r AllowedModesAnomaly) IsKnown() bool {
	switch r {
	case AllowedModesAnomalyOn, AllowedModesAnomalyOff:
		return true
	}
	return false
}

// Defines the rule group to which the current WAF rule belongs.
type WAFRuleGroup struct {
	// Defines the unique identifier of the rule group.
	ID string `json:"id"`
	// Defines the name of the rule group.
	Name string           `json:"name"`
	JSON wafRuleGroupJSON `json:"-"`
}

// wafRuleGroupJSON contains the JSON metadata for the struct [WAFRuleGroup]
type wafRuleGroupJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFRuleGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafRuleGroupJSON) RawJSON() string {
	return r.raw
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WAFPackageRuleListResponse struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// This field can have the runtime type of [[]AllowedModesAnomaly],
	// [[]WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode],
	// [[]WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode].
	AllowedModes interface{} `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
	// when evaluating the request. Applies to anomaly detection WAF rules.
	Mode AllowedModesAnomaly `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string `json:"priority,required"`
	// Defines the default action/mode of a rule.
	DefaultMode WAFPackageRuleListResponseDefaultMode `json:"default_mode"`
	JSON        wafPackageRuleListResponseJSON        `json:"-"`
	union       WAFPackageRuleListResponseUnion
}

// wafPackageRuleListResponseJSON contains the JSON metadata for the struct
// [WAFPackageRuleListResponse]
type wafPackageRuleListResponseJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	DefaultMode  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r wafPackageRuleListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = WAFPackageRuleListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [WAFPackageRuleListResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [WAFPackageRuleListResponseWAFManagedRulesAnomalyRule],
// [WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule],
// [WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule].
func (r WAFPackageRuleListResponse) AsUnion() WAFPackageRuleListResponseUnion {
	return r.union
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by [WAFPackageRuleListResponseWAFManagedRulesAnomalyRule],
// [WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule] or
// [WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule].
type WAFPackageRuleListResponseUnion interface {
	implementsWAFPackageRuleListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageRuleListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleListResponseWAFManagedRulesAnomalyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule{}),
		},
	)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WAFPackageRuleListResponseWAFManagedRulesAnomalyRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []AllowedModesAnomaly `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
	// when evaluating the request. Applies to anomaly detection WAF rules.
	Mode AllowedModesAnomaly `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                                   `json:"priority,required"`
	JSON     wafPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// wafPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON contains the JSON
// metadata for the struct [WAFPackageRuleListResponseWAFManagedRulesAnomalyRule]
type wafPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WAFPackageRuleListResponseWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleListResponseWAFManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleListResponseWAFManagedRulesAnomalyRule) implementsWAFPackageRuleListResponse() {
}

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the list of possible actions of the WAF rule when it is triggered.
	AllowedModes []WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// Defines the default action/mode of a rule.
	DefaultMode WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// Defines the action that the current WAF rule will perform when triggered.
	// Applies to traditional (deny) WAF rules.
	Mode WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                                           `json:"priority,required"`
	JSON     wafPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// wafPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON contains the
// JSON metadata for the struct
// [WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule]
type wafPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	DefaultMode  apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRule) implementsWAFPackageRuleListResponse() {
}

// Defines the action that the current WAF rule will perform when triggered.
// Applies to traditional (deny) WAF rules.
type WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

func (r WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge:
		return true
	}
	return false
}

// Defines the default action/mode of a rule.
type WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

func (r WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge:
		return true
	}
	return false
}

// Defines the action that the current WAF rule will perform when triggered.
// Applies to traditional (deny) WAF rules.
type WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode string

const (
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeDefault   WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "default"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeDisable   WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "disable"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeSimulate  WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeBlock     WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "block"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeChallenge WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

func (r WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeDefault, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeDisable, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeSimulate, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeBlock, WAFPackageRuleListResponseWAFManagedRulesTraditionalDenyRuleModeChallenge:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                                            `json:"priority,required"`
	JSON     wafPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// wafPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON contains the
// JSON metadata for the struct
// [WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule]
type wafPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRule) implementsWAFPackageRuleListResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn  WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

func (r WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn, WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff:
		return true
	}
	return false
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleModeOn  WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleModeOff WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

func (r WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleModeOn, WAFPackageRuleListResponseWAFManagedRulesTraditionalAllowRuleModeOff:
		return true
	}
	return false
}

// Defines the default action/mode of a rule.
type WAFPackageRuleListResponseDefaultMode string

const (
	WAFPackageRuleListResponseDefaultModeDisable   WAFPackageRuleListResponseDefaultMode = "disable"
	WAFPackageRuleListResponseDefaultModeSimulate  WAFPackageRuleListResponseDefaultMode = "simulate"
	WAFPackageRuleListResponseDefaultModeBlock     WAFPackageRuleListResponseDefaultMode = "block"
	WAFPackageRuleListResponseDefaultModeChallenge WAFPackageRuleListResponseDefaultMode = "challenge"
)

func (r WAFPackageRuleListResponseDefaultMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListResponseDefaultModeDisable, WAFPackageRuleListResponseDefaultModeSimulate, WAFPackageRuleListResponseDefaultModeBlock, WAFPackageRuleListResponseDefaultModeChallenge:
		return true
	}
	return false
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WAFPackageRuleEditResponse struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// This field can have the runtime type of [[]AllowedModesAnomaly],
	// [[]WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode],
	// [[]WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode].
	AllowedModes interface{} `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
	// when evaluating the request. Applies to anomaly detection WAF rules.
	Mode AllowedModesAnomaly `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string `json:"priority,required"`
	// Defines the default action/mode of a rule.
	DefaultMode WAFPackageRuleEditResponseDefaultMode `json:"default_mode"`
	JSON        wafPackageRuleEditResponseJSON        `json:"-"`
	union       WAFPackageRuleEditResponseUnion
}

// wafPackageRuleEditResponseJSON contains the JSON metadata for the struct
// [WAFPackageRuleEditResponse]
type wafPackageRuleEditResponseJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	DefaultMode  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r wafPackageRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

func (r *WAFPackageRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	*r = WAFPackageRuleEditResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [WAFPackageRuleEditResponseUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are
// [WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule],
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule],
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule].
func (r WAFPackageRuleEditResponse) AsUnion() WAFPackageRuleEditResponseUnion {
	return r.union
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
//
// Union satisfied by [WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule],
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule] or
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule].
type WAFPackageRuleEditResponseUnion interface {
	implementsWAFPackageRuleEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*WAFPackageRuleEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule{}),
		},
	)
}

// When triggered, anomaly detection WAF rules contribute to an overall threat
// score that will determine if a request is considered malicious. You can
// configure the total scoring threshold through the 'sensitivity' property of the
// WAF package.
type WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule. Applies to anomaly
	// detection WAF rules.
	AllowedModes []AllowedModesAnomaly `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// Defines the mode anomaly. When set to `on`, the current WAF rule will be used
	// when evaluating the request. Applies to anomaly detection WAF rules.
	Mode AllowedModesAnomaly `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                                   `json:"priority,required"`
	JSON     wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON contains the JSON
// metadata for the struct [WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule]
type wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesAnomalyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleEditResponseWAFManagedRulesAnomalyRule) implementsWAFPackageRuleEditResponse() {
}

// When triggered, traditional WAF rules cause the firewall to immediately act upon
// the request based on the configuration of the rule. A 'deny' rule will
// immediately respond to the request based on the configured rule action/mode (for
// example, 'block') and no other rules will be processed.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the list of possible actions of the WAF rule when it is triggered.
	AllowedModes []WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode `json:"allowed_modes,required"`
	// Defines the default action/mode of a rule.
	DefaultMode WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode `json:"default_mode,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// Defines the action that the current WAF rule will perform when triggered.
	// Applies to traditional (deny) WAF rules.
	Mode WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                                           `json:"priority,required"`
	JSON     wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON contains the
// JSON metadata for the struct
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule]
type wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	DefaultMode  apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRule) implementsWAFPackageRuleEditResponse() {
}

// Defines the action that the current WAF rule will perform when triggered.
// Applies to traditional (deny) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "default"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "disable"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate  WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "simulate"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock     WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "block"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode = "challenge"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDefault, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeDisable, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeSimulate, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeBlock, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleAllowedModeChallenge:
		return true
	}
	return false
}

// Defines the default action/mode of a rule.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "disable"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate  WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "simulate"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock     WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "block"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode = "challenge"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeDisable, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeSimulate, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeBlock, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleDefaultModeChallenge:
		return true
	}
	return false
}

// Defines the action that the current WAF rule will perform when triggered.
// Applies to traditional (deny) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDefault   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "default"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDisable   WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "disable"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeSimulate  WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "simulate"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeBlock     WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "block"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeChallenge WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode = "challenge"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDefault, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeDisable, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeSimulate, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeBlock, WAFPackageRuleEditResponseWAFManagedRulesTraditionalDenyRuleModeChallenge:
		return true
	}
	return false
}

// When triggered, traditional WAF rules cause the firewall to immediately act on
// the request based on the rule configuration. An 'allow' rule will immediately
// allow the request and no other rules will be processed.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule struct {
	// Defines the unique identifier of the WAF rule.
	ID string `json:"id,required"`
	// Defines the available modes for the current WAF rule.
	AllowedModes []WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode `json:"allowed_modes,required"`
	// Defines the public description of the WAF rule.
	Description string `json:"description,required"`
	// Defines the rule group to which the current WAF rule belongs.
	Group WAFRuleGroup `json:"group,required"`
	// When set to `on`, the current rule will be used when evaluating the request.
	// Applies to traditional (allow) WAF rules.
	Mode WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode `json:"mode,required"`
	// Defines the unique identifier of a WAF package.
	PackageID string `json:"package_id,required"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority string                                                            `json:"priority,required"`
	JSON     wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON `json:"-"`
}

// wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON contains the
// JSON metadata for the struct
// [WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule]
type wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON struct {
	ID           apijson.Field
	AllowedModes apijson.Field
	Description  apijson.Field
	Group        apijson.Field
	Mode         apijson.Field
	PackageID    apijson.Field
	Priority     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleJSON) RawJSON() string {
	return r.raw
}

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRule) implementsWAFPackageRuleEditResponse() {
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn  WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode = "off"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOn, WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleAllowedModeOff:
		return true
	}
	return false
}

// When set to `on`, the current rule will be used when evaluating the request.
// Applies to traditional (allow) WAF rules.
type WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode string

const (
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOn  WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "on"
	WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOff WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode = "off"
)

func (r WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOn, WAFPackageRuleEditResponseWAFManagedRulesTraditionalAllowRuleModeOff:
		return true
	}
	return false
}

// Defines the default action/mode of a rule.
type WAFPackageRuleEditResponseDefaultMode string

const (
	WAFPackageRuleEditResponseDefaultModeDisable   WAFPackageRuleEditResponseDefaultMode = "disable"
	WAFPackageRuleEditResponseDefaultModeSimulate  WAFPackageRuleEditResponseDefaultMode = "simulate"
	WAFPackageRuleEditResponseDefaultModeBlock     WAFPackageRuleEditResponseDefaultMode = "block"
	WAFPackageRuleEditResponseDefaultModeChallenge WAFPackageRuleEditResponseDefaultMode = "challenge"
)

func (r WAFPackageRuleEditResponseDefaultMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseDefaultModeDisable, WAFPackageRuleEditResponseDefaultModeSimulate, WAFPackageRuleEditResponseDefaultModeBlock, WAFPackageRuleEditResponseDefaultModeChallenge:
		return true
	}
	return false
}

type WAFPackageRuleListParams struct {
	// Defines an identifier of a schema.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines the public description of the WAF rule.
	Description param.Field[string] `query:"description"`
	// Defines the direction used to sort returned rules.
	Direction param.Field[WAFPackageRuleListParamsDirection] `query:"direction"`
	// Defines the unique identifier of the rule group.
	GroupID param.Field[string] `query:"group_id"`
	// Defines the search requirements. When set to `all`, all the search requirements
	// must match. When set to `any`, only one of the search requirements has to match.
	Match param.Field[WAFPackageRuleListParamsMatch] `query:"match"`
	// Defines the action/mode a rule has been overridden to perform.
	Mode param.Field[WAFPackageRuleListParamsMode] `query:"mode"`
	// Defines the field used to sort returned rules.
	Order param.Field[WAFPackageRuleListParamsOrder] `query:"order"`
	// Defines the page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Defines the number of rules per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Defines the order in which the individual WAF rule is executed within its rule
	// group.
	Priority param.Field[string] `query:"priority"`
}

// URLQuery serializes [WAFPackageRuleListParams]'s query parameters as
// `url.Values`.
func (r WAFPackageRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Defines the direction used to sort returned rules.
type WAFPackageRuleListParamsDirection string

const (
	WAFPackageRuleListParamsDirectionAsc  WAFPackageRuleListParamsDirection = "asc"
	WAFPackageRuleListParamsDirectionDesc WAFPackageRuleListParamsDirection = "desc"
)

func (r WAFPackageRuleListParamsDirection) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsDirectionAsc, WAFPackageRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// Defines the search requirements. When set to `all`, all the search requirements
// must match. When set to `any`, only one of the search requirements has to match.
type WAFPackageRuleListParamsMatch string

const (
	WAFPackageRuleListParamsMatchAny WAFPackageRuleListParamsMatch = "any"
	WAFPackageRuleListParamsMatchAll WAFPackageRuleListParamsMatch = "all"
)

func (r WAFPackageRuleListParamsMatch) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsMatchAny, WAFPackageRuleListParamsMatchAll:
		return true
	}
	return false
}

// Defines the action/mode a rule has been overridden to perform.
type WAFPackageRuleListParamsMode string

const (
	WAFPackageRuleListParamsModeDis WAFPackageRuleListParamsMode = "DIS"
	WAFPackageRuleListParamsModeChl WAFPackageRuleListParamsMode = "CHL"
	WAFPackageRuleListParamsModeBlk WAFPackageRuleListParamsMode = "BLK"
	WAFPackageRuleListParamsModeSim WAFPackageRuleListParamsMode = "SIM"
)

func (r WAFPackageRuleListParamsMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsModeDis, WAFPackageRuleListParamsModeChl, WAFPackageRuleListParamsModeBlk, WAFPackageRuleListParamsModeSim:
		return true
	}
	return false
}

// Defines the field used to sort returned rules.
type WAFPackageRuleListParamsOrder string

const (
	WAFPackageRuleListParamsOrderPriority    WAFPackageRuleListParamsOrder = "priority"
	WAFPackageRuleListParamsOrderGroupID     WAFPackageRuleListParamsOrder = "group_id"
	WAFPackageRuleListParamsOrderDescription WAFPackageRuleListParamsOrder = "description"
)

func (r WAFPackageRuleListParamsOrder) IsKnown() bool {
	switch r {
	case WAFPackageRuleListParamsOrderPriority, WAFPackageRuleListParamsOrderGroupID, WAFPackageRuleListParamsOrderDescription:
		return true
	}
	return false
}

type WAFPackageRuleEditParams struct {
	// Defines an identifier of a schema.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Defines the mode/action of the rule when triggered. You must use a value from
	// the `allowed_modes` array of the current rule.
	Mode param.Field[WAFPackageRuleEditParamsMode] `json:"mode"`
}

func (r WAFPackageRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Defines the mode/action of the rule when triggered. You must use a value from
// the `allowed_modes` array of the current rule.
type WAFPackageRuleEditParamsMode string

const (
	WAFPackageRuleEditParamsModeDefault   WAFPackageRuleEditParamsMode = "default"
	WAFPackageRuleEditParamsModeDisable   WAFPackageRuleEditParamsMode = "disable"
	WAFPackageRuleEditParamsModeSimulate  WAFPackageRuleEditParamsMode = "simulate"
	WAFPackageRuleEditParamsModeBlock     WAFPackageRuleEditParamsMode = "block"
	WAFPackageRuleEditParamsModeChallenge WAFPackageRuleEditParamsMode = "challenge"
	WAFPackageRuleEditParamsModeOn        WAFPackageRuleEditParamsMode = "on"
	WAFPackageRuleEditParamsModeOff       WAFPackageRuleEditParamsMode = "off"
)

func (r WAFPackageRuleEditParamsMode) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditParamsModeDefault, WAFPackageRuleEditParamsModeDisable, WAFPackageRuleEditParamsModeSimulate, WAFPackageRuleEditParamsModeBlock, WAFPackageRuleEditParamsModeChallenge, WAFPackageRuleEditParamsModeOn, WAFPackageRuleEditParamsModeOff:
		return true
	}
	return false
}

type WAFPackageRuleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// When triggered, anomaly detection WAF rules contribute to an overall threat
	// score that will determine if a request is considered malicious. You can
	// configure the total scoring threshold through the 'sensitivity' property of the
	// WAF package.
	Result WAFPackageRuleEditResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFPackageRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafPackageRuleEditResponseEnvelopeJSON    `json:"-"`
}

// wafPackageRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFPackageRuleEditResponseEnvelope]
type wafPackageRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFPackageRuleEditResponseEnvelopeSuccess bool

const (
	WAFPackageRuleEditResponseEnvelopeSuccessTrue WAFPackageRuleEditResponseEnvelopeSuccess = true
)

func (r WAFPackageRuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFPackageRuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type WAFPackageRuleGetParams struct {
	// Defines an identifier of a schema.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type WAFPackageRuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required"`
	// Defines whether the API call was successful.
	Success WAFPackageRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    wafPackageRuleGetResponseEnvelopeJSON    `json:"-"`
}

// wafPackageRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WAFPackageRuleGetResponseEnvelope]
type wafPackageRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WAFPackageRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r wafPackageRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type WAFPackageRuleGetResponseEnvelopeSuccess bool

const (
	WAFPackageRuleGetResponseEnvelopeSuccessTrue WAFPackageRuleGetResponseEnvelopeSuccess = true
)

func (r WAFPackageRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WAFPackageRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
