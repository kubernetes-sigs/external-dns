// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
	"github.com/tidwall/gjson"
)

// AccessRuleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessRuleService] method instead.
type AccessRuleService struct {
	Options []option.RequestOption
}

// NewAccessRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessRuleService(opts ...option.RequestOption) (r *AccessRuleService) {
	r = &AccessRuleService{}
	r.Options = opts
	return
}

// Creates a new IP Access rule for an account or zone. The rule will apply to all
// zones in the account or zone.
//
// Note: To create an IP Access rule that applies to a single zone, refer to the
// [IP Access rules for a zone](#ip-access-rules-for-a-zone) endpoints.
func (r *AccessRuleService) New(ctx context.Context, params AccessRuleNewParams, opts ...option.RequestOption) (res *AccessRuleNewResponse, err error) {
	var env AccessRuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches IP Access rules of an account or zone. These rules apply to all the
// zones in the account or zone. You can filter the results using several optional
// parameters.
func (r *AccessRuleService) List(ctx context.Context, params AccessRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[AccessRuleListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules", accountOrZone, accountOrZoneID)
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

// Fetches IP Access rules of an account or zone. These rules apply to all the
// zones in the account or zone. You can filter the results using several optional
// parameters.
func (r *AccessRuleService) ListAutoPaging(ctx context.Context, params AccessRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[AccessRuleListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes an existing IP Access rule defined.
//
// Note: This operation will affect all zones in the account or zone.
func (r *AccessRuleService) Delete(ctx context.Context, ruleID string, body AccessRuleDeleteParams, opts ...option.RequestOption) (res *AccessRuleDeleteResponse, err error) {
	var env AccessRuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules/%s", accountOrZone, accountOrZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an IP Access rule defined.
//
// Note: This operation will affect all zones in the account or zone.
func (r *AccessRuleService) Edit(ctx context.Context, ruleID string, params AccessRuleEditParams, opts ...option.RequestOption) (res *AccessRuleEditResponse, err error) {
	var env AccessRuleEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules/%s", accountOrZone, accountOrZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of an IP Access rule defined.
func (r *AccessRuleService) Get(ctx context.Context, ruleID string, query AccessRuleGetParams, opts ...option.RequestOption) (res *AccessRuleGetResponse, err error) {
	var env AccessRuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/firewall/access_rules/rules/%s", accountOrZone, accountOrZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessRuleCIDRConfiguration struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target AccessRuleCIDRConfigurationTarget `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value string                          `json:"value"`
	JSON  accessRuleCIDRConfigurationJSON `json:"-"`
}

// accessRuleCIDRConfigurationJSON contains the JSON metadata for the struct
// [AccessRuleCIDRConfiguration]
type accessRuleCIDRConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleCIDRConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleCIDRConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleCIDRConfiguration) implementsAccessRuleNewResponseConfiguration() {}

func (r AccessRuleCIDRConfiguration) implementsAccessRuleListResponseConfiguration() {}

func (r AccessRuleCIDRConfiguration) implementsAccessRuleEditResponseConfiguration() {}

func (r AccessRuleCIDRConfiguration) implementsAccessRuleGetResponseConfiguration() {}

// The configuration target. You must set the target to `ip_range` when specifying
// an IP address range in the rule.
type AccessRuleCIDRConfigurationTarget string

const (
	AccessRuleCIDRConfigurationTargetIPRange AccessRuleCIDRConfigurationTarget = "ip_range"
)

func (r AccessRuleCIDRConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleCIDRConfigurationTargetIPRange:
		return true
	}
	return false
}

type AccessRuleCIDRConfigurationParam struct {
	// The configuration target. You must set the target to `ip_range` when specifying
	// an IP address range in the rule.
	Target param.Field[AccessRuleCIDRConfigurationTarget] `json:"target"`
	// The IP address range to match. You can only use prefix lengths `/16` and `/24`
	// for IPv4 ranges, and prefix lengths `/32`, `/48`, and `/64` for IPv6 ranges.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleCIDRConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleCIDRConfigurationParam) implementsAccessRuleNewParamsConfigurationUnion() {}

func (r AccessRuleCIDRConfigurationParam) implementsAccessRuleEditParamsConfigurationUnion() {}

func (r AccessRuleCIDRConfigurationParam) implementsUARuleUpdateParamsConfigurationUnion() {}

type AccessRuleIPConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target AccessRuleIPConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                        `json:"value"`
	JSON  accessRuleIPConfigurationJSON `json:"-"`
}

// accessRuleIPConfigurationJSON contains the JSON metadata for the struct
// [AccessRuleIPConfiguration]
type accessRuleIPConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleIPConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleIPConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r AccessRuleIPConfiguration) implementsAccessRuleNewResponseConfiguration() {}

func (r AccessRuleIPConfiguration) implementsAccessRuleListResponseConfiguration() {}

func (r AccessRuleIPConfiguration) implementsAccessRuleEditResponseConfiguration() {}

func (r AccessRuleIPConfiguration) implementsAccessRuleGetResponseConfiguration() {}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleIPConfigurationTarget string

const (
	AccessRuleIPConfigurationTargetIP AccessRuleIPConfigurationTarget = "ip"
)

func (r AccessRuleIPConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleIPConfigurationTargetIP:
		return true
	}
	return false
}

type AccessRuleIPConfigurationParam struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleIPConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleIPConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleIPConfigurationParam) implementsAccessRuleNewParamsConfigurationUnion() {}

func (r AccessRuleIPConfigurationParam) implementsAccessRuleEditParamsConfigurationUnion() {}

func (r AccessRuleIPConfigurationParam) implementsUARuleUpdateParamsConfigurationUnion() {}

type ASNConfiguration struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target ASNConfigurationTarget `json:"target"`
	// The AS number to match.
	Value string               `json:"value"`
	JSON  asnConfigurationJSON `json:"-"`
}

// asnConfigurationJSON contains the JSON metadata for the struct
// [ASNConfiguration]
type asnConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r ASNConfiguration) implementsAccessRuleNewResponseConfiguration() {}

func (r ASNConfiguration) implementsAccessRuleListResponseConfiguration() {}

func (r ASNConfiguration) implementsAccessRuleEditResponseConfiguration() {}

func (r ASNConfiguration) implementsAccessRuleGetResponseConfiguration() {}

// The configuration target. You must set the target to `asn` when specifying an
// Autonomous System Number (ASN) in the rule.
type ASNConfigurationTarget string

const (
	ASNConfigurationTargetASN ASNConfigurationTarget = "asn"
)

func (r ASNConfigurationTarget) IsKnown() bool {
	switch r {
	case ASNConfigurationTargetASN:
		return true
	}
	return false
}

type ASNConfigurationParam struct {
	// The configuration target. You must set the target to `asn` when specifying an
	// Autonomous System Number (ASN) in the rule.
	Target param.Field[ASNConfigurationTarget] `json:"target"`
	// The AS number to match.
	Value param.Field[string] `json:"value"`
}

func (r ASNConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ASNConfigurationParam) implementsAccessRuleNewParamsConfigurationUnion() {}

func (r ASNConfigurationParam) implementsAccessRuleEditParamsConfigurationUnion() {}

func (r ASNConfigurationParam) implementsUARuleUpdateParamsConfigurationUnion() {}

type CountryConfiguration struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target CountryConfigurationTarget `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value string                   `json:"value"`
	JSON  countryConfigurationJSON `json:"-"`
}

// countryConfigurationJSON contains the JSON metadata for the struct
// [CountryConfiguration]
type countryConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CountryConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r countryConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r CountryConfiguration) implementsAccessRuleNewResponseConfiguration() {}

func (r CountryConfiguration) implementsAccessRuleListResponseConfiguration() {}

func (r CountryConfiguration) implementsAccessRuleEditResponseConfiguration() {}

func (r CountryConfiguration) implementsAccessRuleGetResponseConfiguration() {}

// The configuration target. You must set the target to `country` when specifying a
// country code in the rule.
type CountryConfigurationTarget string

const (
	CountryConfigurationTargetCountry CountryConfigurationTarget = "country"
)

func (r CountryConfigurationTarget) IsKnown() bool {
	switch r {
	case CountryConfigurationTargetCountry:
		return true
	}
	return false
}

type CountryConfigurationParam struct {
	// The configuration target. You must set the target to `country` when specifying a
	// country code in the rule.
	Target param.Field[CountryConfigurationTarget] `json:"target"`
	// The two-letter ISO-3166-1 alpha-2 code to match. For more information, refer to
	// [IP Access rules: Parameters](https://developers.cloudflare.com/waf/tools/ip-access-rules/parameters/#country).
	Value param.Field[string] `json:"value"`
}

func (r CountryConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CountryConfigurationParam) implementsAccessRuleNewParamsConfigurationUnion() {}

func (r CountryConfigurationParam) implementsAccessRuleEditParamsConfigurationUnion() {}

func (r CountryConfigurationParam) implementsUARuleUpdateParamsConfigurationUnion() {}

type IPV6Configuration struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target IPV6ConfigurationTarget `json:"target"`
	// The IPv6 address to match.
	Value string                `json:"value"`
	JSON  ipv6ConfigurationJSON `json:"-"`
}

// ipv6ConfigurationJSON contains the JSON metadata for the struct
// [IPV6Configuration]
type ipv6ConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPV6Configuration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipv6ConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r IPV6Configuration) implementsAccessRuleNewResponseConfiguration() {}

func (r IPV6Configuration) implementsAccessRuleListResponseConfiguration() {}

func (r IPV6Configuration) implementsAccessRuleEditResponseConfiguration() {}

func (r IPV6Configuration) implementsAccessRuleGetResponseConfiguration() {}

// The configuration target. You must set the target to `ip6` when specifying an
// IPv6 address in the rule.
type IPV6ConfigurationTarget string

const (
	IPV6ConfigurationTargetIp6 IPV6ConfigurationTarget = "ip6"
)

func (r IPV6ConfigurationTarget) IsKnown() bool {
	switch r {
	case IPV6ConfigurationTargetIp6:
		return true
	}
	return false
}

type IPV6ConfigurationParam struct {
	// The configuration target. You must set the target to `ip6` when specifying an
	// IPv6 address in the rule.
	Target param.Field[IPV6ConfigurationTarget] `json:"target"`
	// The IPv6 address to match.
	Value param.Field[string] `json:"value"`
}

func (r IPV6ConfigurationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r IPV6ConfigurationParam) implementsAccessRuleNewParamsConfigurationUnion() {}

func (r IPV6ConfigurationParam) implementsAccessRuleEditParamsConfigurationUnion() {}

func (r IPV6ConfigurationParam) implementsUARuleUpdateParamsConfigurationUnion() {}

type AccessRuleNewResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []AccessRuleNewResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration AccessRuleNewResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode AccessRuleNewResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string `json:"notes"`
	// All zones owned by the user will have the rule applied.
	Scope AccessRuleNewResponseScope `json:"scope"`
	JSON  accessRuleNewResponseJSON  `json:"-"`
}

// accessRuleNewResponseJSON contains the JSON metadata for the struct
// [AccessRuleNewResponse]
type accessRuleNewResponseJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	Scope         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type AccessRuleNewResponseAllowedMode string

const (
	AccessRuleNewResponseAllowedModeBlock            AccessRuleNewResponseAllowedMode = "block"
	AccessRuleNewResponseAllowedModeChallenge        AccessRuleNewResponseAllowedMode = "challenge"
	AccessRuleNewResponseAllowedModeWhitelist        AccessRuleNewResponseAllowedMode = "whitelist"
	AccessRuleNewResponseAllowedModeJSChallenge      AccessRuleNewResponseAllowedMode = "js_challenge"
	AccessRuleNewResponseAllowedModeManagedChallenge AccessRuleNewResponseAllowedMode = "managed_challenge"
)

func (r AccessRuleNewResponseAllowedMode) IsKnown() bool {
	switch r {
	case AccessRuleNewResponseAllowedModeBlock, AccessRuleNewResponseAllowedModeChallenge, AccessRuleNewResponseAllowedModeWhitelist, AccessRuleNewResponseAllowedModeJSChallenge, AccessRuleNewResponseAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
type AccessRuleNewResponseConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target AccessRuleNewResponseConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                 `json:"value"`
	JSON  accessRuleNewResponseConfigurationJSON `json:"-"`
	union AccessRuleNewResponseConfigurationUnion
}

// accessRuleNewResponseConfigurationJSON contains the JSON metadata for the struct
// [AccessRuleNewResponseConfiguration]
type accessRuleNewResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accessRuleNewResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessRuleNewResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = AccessRuleNewResponseConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessRuleNewResponseConfigurationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [AccessRuleIPConfiguration],
// [IPV6Configuration], [AccessRuleCIDRConfiguration], [ASNConfiguration],
// [CountryConfiguration].
func (r AccessRuleNewResponseConfiguration) AsUnion() AccessRuleNewResponseConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by [AccessRuleIPConfiguration], [IPV6Configuration],
// [AccessRuleCIDRConfiguration], [ASNConfiguration] or [CountryConfiguration].
type AccessRuleNewResponseConfigurationUnion interface {
	implementsAccessRuleNewResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessRuleNewResponseConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CountryConfiguration{}),
		},
	)
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleNewResponseConfigurationTarget string

const (
	AccessRuleNewResponseConfigurationTargetIP      AccessRuleNewResponseConfigurationTarget = "ip"
	AccessRuleNewResponseConfigurationTargetIp6     AccessRuleNewResponseConfigurationTarget = "ip6"
	AccessRuleNewResponseConfigurationTargetIPRange AccessRuleNewResponseConfigurationTarget = "ip_range"
	AccessRuleNewResponseConfigurationTargetASN     AccessRuleNewResponseConfigurationTarget = "asn"
	AccessRuleNewResponseConfigurationTargetCountry AccessRuleNewResponseConfigurationTarget = "country"
)

func (r AccessRuleNewResponseConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewResponseConfigurationTargetIP, AccessRuleNewResponseConfigurationTargetIp6, AccessRuleNewResponseConfigurationTargetIPRange, AccessRuleNewResponseConfigurationTargetASN, AccessRuleNewResponseConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleNewResponseMode string

const (
	AccessRuleNewResponseModeBlock            AccessRuleNewResponseMode = "block"
	AccessRuleNewResponseModeChallenge        AccessRuleNewResponseMode = "challenge"
	AccessRuleNewResponseModeWhitelist        AccessRuleNewResponseMode = "whitelist"
	AccessRuleNewResponseModeJSChallenge      AccessRuleNewResponseMode = "js_challenge"
	AccessRuleNewResponseModeManagedChallenge AccessRuleNewResponseMode = "managed_challenge"
)

func (r AccessRuleNewResponseMode) IsKnown() bool {
	switch r {
	case AccessRuleNewResponseModeBlock, AccessRuleNewResponseModeChallenge, AccessRuleNewResponseModeWhitelist, AccessRuleNewResponseModeJSChallenge, AccessRuleNewResponseModeManagedChallenge:
		return true
	}
	return false
}

// All zones owned by the user will have the rule applied.
type AccessRuleNewResponseScope struct {
	// Defines an identifier.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// Defines the scope of the rule.
	Type AccessRuleNewResponseScopeType `json:"type"`
	JSON accessRuleNewResponseScopeJSON `json:"-"`
}

// accessRuleNewResponseScopeJSON contains the JSON metadata for the struct
// [AccessRuleNewResponseScope]
type accessRuleNewResponseScopeJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleNewResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleNewResponseScopeJSON) RawJSON() string {
	return r.raw
}

// Defines the scope of the rule.
type AccessRuleNewResponseScopeType string

const (
	AccessRuleNewResponseScopeTypeUser         AccessRuleNewResponseScopeType = "user"
	AccessRuleNewResponseScopeTypeOrganization AccessRuleNewResponseScopeType = "organization"
)

func (r AccessRuleNewResponseScopeType) IsKnown() bool {
	switch r {
	case AccessRuleNewResponseScopeTypeUser, AccessRuleNewResponseScopeTypeOrganization:
		return true
	}
	return false
}

type AccessRuleListResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []AccessRuleListResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration AccessRuleListResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode AccessRuleListResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string `json:"notes"`
	// All zones owned by the user will have the rule applied.
	Scope AccessRuleListResponseScope `json:"scope"`
	JSON  accessRuleListResponseJSON  `json:"-"`
}

// accessRuleListResponseJSON contains the JSON metadata for the struct
// [AccessRuleListResponse]
type accessRuleListResponseJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	Scope         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleListResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type AccessRuleListResponseAllowedMode string

const (
	AccessRuleListResponseAllowedModeBlock            AccessRuleListResponseAllowedMode = "block"
	AccessRuleListResponseAllowedModeChallenge        AccessRuleListResponseAllowedMode = "challenge"
	AccessRuleListResponseAllowedModeWhitelist        AccessRuleListResponseAllowedMode = "whitelist"
	AccessRuleListResponseAllowedModeJSChallenge      AccessRuleListResponseAllowedMode = "js_challenge"
	AccessRuleListResponseAllowedModeManagedChallenge AccessRuleListResponseAllowedMode = "managed_challenge"
)

func (r AccessRuleListResponseAllowedMode) IsKnown() bool {
	switch r {
	case AccessRuleListResponseAllowedModeBlock, AccessRuleListResponseAllowedModeChallenge, AccessRuleListResponseAllowedModeWhitelist, AccessRuleListResponseAllowedModeJSChallenge, AccessRuleListResponseAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
type AccessRuleListResponseConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target AccessRuleListResponseConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                  `json:"value"`
	JSON  accessRuleListResponseConfigurationJSON `json:"-"`
	union AccessRuleListResponseConfigurationUnion
}

// accessRuleListResponseConfigurationJSON contains the JSON metadata for the
// struct [AccessRuleListResponseConfiguration]
type accessRuleListResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accessRuleListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessRuleListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = AccessRuleListResponseConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessRuleListResponseConfigurationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [AccessRuleIPConfiguration],
// [IPV6Configuration], [AccessRuleCIDRConfiguration], [ASNConfiguration],
// [CountryConfiguration].
func (r AccessRuleListResponseConfiguration) AsUnion() AccessRuleListResponseConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by [AccessRuleIPConfiguration], [IPV6Configuration],
// [AccessRuleCIDRConfiguration], [ASNConfiguration] or [CountryConfiguration].
type AccessRuleListResponseConfigurationUnion interface {
	implementsAccessRuleListResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessRuleListResponseConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CountryConfiguration{}),
		},
	)
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleListResponseConfigurationTarget string

const (
	AccessRuleListResponseConfigurationTargetIP      AccessRuleListResponseConfigurationTarget = "ip"
	AccessRuleListResponseConfigurationTargetIp6     AccessRuleListResponseConfigurationTarget = "ip6"
	AccessRuleListResponseConfigurationTargetIPRange AccessRuleListResponseConfigurationTarget = "ip_range"
	AccessRuleListResponseConfigurationTargetASN     AccessRuleListResponseConfigurationTarget = "asn"
	AccessRuleListResponseConfigurationTargetCountry AccessRuleListResponseConfigurationTarget = "country"
)

func (r AccessRuleListResponseConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleListResponseConfigurationTargetIP, AccessRuleListResponseConfigurationTargetIp6, AccessRuleListResponseConfigurationTargetIPRange, AccessRuleListResponseConfigurationTargetASN, AccessRuleListResponseConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleListResponseMode string

const (
	AccessRuleListResponseModeBlock            AccessRuleListResponseMode = "block"
	AccessRuleListResponseModeChallenge        AccessRuleListResponseMode = "challenge"
	AccessRuleListResponseModeWhitelist        AccessRuleListResponseMode = "whitelist"
	AccessRuleListResponseModeJSChallenge      AccessRuleListResponseMode = "js_challenge"
	AccessRuleListResponseModeManagedChallenge AccessRuleListResponseMode = "managed_challenge"
)

func (r AccessRuleListResponseMode) IsKnown() bool {
	switch r {
	case AccessRuleListResponseModeBlock, AccessRuleListResponseModeChallenge, AccessRuleListResponseModeWhitelist, AccessRuleListResponseModeJSChallenge, AccessRuleListResponseModeManagedChallenge:
		return true
	}
	return false
}

// All zones owned by the user will have the rule applied.
type AccessRuleListResponseScope struct {
	// Defines an identifier.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// Defines the scope of the rule.
	Type AccessRuleListResponseScopeType `json:"type"`
	JSON accessRuleListResponseScopeJSON `json:"-"`
}

// accessRuleListResponseScopeJSON contains the JSON metadata for the struct
// [AccessRuleListResponseScope]
type accessRuleListResponseScopeJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleListResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleListResponseScopeJSON) RawJSON() string {
	return r.raw
}

// Defines the scope of the rule.
type AccessRuleListResponseScopeType string

const (
	AccessRuleListResponseScopeTypeUser         AccessRuleListResponseScopeType = "user"
	AccessRuleListResponseScopeTypeOrganization AccessRuleListResponseScopeType = "organization"
)

func (r AccessRuleListResponseScopeType) IsKnown() bool {
	switch r {
	case AccessRuleListResponseScopeTypeUser, AccessRuleListResponseScopeTypeOrganization:
		return true
	}
	return false
}

type AccessRuleDeleteResponse struct {
	// Defines an identifier.
	ID   string                       `json:"id,required"`
	JSON accessRuleDeleteResponseJSON `json:"-"`
}

// accessRuleDeleteResponseJSON contains the JSON metadata for the struct
// [AccessRuleDeleteResponse]
type accessRuleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessRuleEditResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []AccessRuleEditResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration AccessRuleEditResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode AccessRuleEditResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string `json:"notes"`
	// All zones owned by the user will have the rule applied.
	Scope AccessRuleEditResponseScope `json:"scope"`
	JSON  accessRuleEditResponseJSON  `json:"-"`
}

// accessRuleEditResponseJSON contains the JSON metadata for the struct
// [AccessRuleEditResponse]
type accessRuleEditResponseJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	Scope         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type AccessRuleEditResponseAllowedMode string

const (
	AccessRuleEditResponseAllowedModeBlock            AccessRuleEditResponseAllowedMode = "block"
	AccessRuleEditResponseAllowedModeChallenge        AccessRuleEditResponseAllowedMode = "challenge"
	AccessRuleEditResponseAllowedModeWhitelist        AccessRuleEditResponseAllowedMode = "whitelist"
	AccessRuleEditResponseAllowedModeJSChallenge      AccessRuleEditResponseAllowedMode = "js_challenge"
	AccessRuleEditResponseAllowedModeManagedChallenge AccessRuleEditResponseAllowedMode = "managed_challenge"
)

func (r AccessRuleEditResponseAllowedMode) IsKnown() bool {
	switch r {
	case AccessRuleEditResponseAllowedModeBlock, AccessRuleEditResponseAllowedModeChallenge, AccessRuleEditResponseAllowedModeWhitelist, AccessRuleEditResponseAllowedModeJSChallenge, AccessRuleEditResponseAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
type AccessRuleEditResponseConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target AccessRuleEditResponseConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                  `json:"value"`
	JSON  accessRuleEditResponseConfigurationJSON `json:"-"`
	union AccessRuleEditResponseConfigurationUnion
}

// accessRuleEditResponseConfigurationJSON contains the JSON metadata for the
// struct [AccessRuleEditResponseConfiguration]
type accessRuleEditResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accessRuleEditResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessRuleEditResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = AccessRuleEditResponseConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessRuleEditResponseConfigurationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [AccessRuleIPConfiguration],
// [IPV6Configuration], [AccessRuleCIDRConfiguration], [ASNConfiguration],
// [CountryConfiguration].
func (r AccessRuleEditResponseConfiguration) AsUnion() AccessRuleEditResponseConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by [AccessRuleIPConfiguration], [IPV6Configuration],
// [AccessRuleCIDRConfiguration], [ASNConfiguration] or [CountryConfiguration].
type AccessRuleEditResponseConfigurationUnion interface {
	implementsAccessRuleEditResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessRuleEditResponseConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CountryConfiguration{}),
		},
	)
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleEditResponseConfigurationTarget string

const (
	AccessRuleEditResponseConfigurationTargetIP      AccessRuleEditResponseConfigurationTarget = "ip"
	AccessRuleEditResponseConfigurationTargetIp6     AccessRuleEditResponseConfigurationTarget = "ip6"
	AccessRuleEditResponseConfigurationTargetIPRange AccessRuleEditResponseConfigurationTarget = "ip_range"
	AccessRuleEditResponseConfigurationTargetASN     AccessRuleEditResponseConfigurationTarget = "asn"
	AccessRuleEditResponseConfigurationTargetCountry AccessRuleEditResponseConfigurationTarget = "country"
)

func (r AccessRuleEditResponseConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditResponseConfigurationTargetIP, AccessRuleEditResponseConfigurationTargetIp6, AccessRuleEditResponseConfigurationTargetIPRange, AccessRuleEditResponseConfigurationTargetASN, AccessRuleEditResponseConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleEditResponseMode string

const (
	AccessRuleEditResponseModeBlock            AccessRuleEditResponseMode = "block"
	AccessRuleEditResponseModeChallenge        AccessRuleEditResponseMode = "challenge"
	AccessRuleEditResponseModeWhitelist        AccessRuleEditResponseMode = "whitelist"
	AccessRuleEditResponseModeJSChallenge      AccessRuleEditResponseMode = "js_challenge"
	AccessRuleEditResponseModeManagedChallenge AccessRuleEditResponseMode = "managed_challenge"
)

func (r AccessRuleEditResponseMode) IsKnown() bool {
	switch r {
	case AccessRuleEditResponseModeBlock, AccessRuleEditResponseModeChallenge, AccessRuleEditResponseModeWhitelist, AccessRuleEditResponseModeJSChallenge, AccessRuleEditResponseModeManagedChallenge:
		return true
	}
	return false
}

// All zones owned by the user will have the rule applied.
type AccessRuleEditResponseScope struct {
	// Defines an identifier.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// Defines the scope of the rule.
	Type AccessRuleEditResponseScopeType `json:"type"`
	JSON accessRuleEditResponseScopeJSON `json:"-"`
}

// accessRuleEditResponseScopeJSON contains the JSON metadata for the struct
// [AccessRuleEditResponseScope]
type accessRuleEditResponseScopeJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleEditResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleEditResponseScopeJSON) RawJSON() string {
	return r.raw
}

// Defines the scope of the rule.
type AccessRuleEditResponseScopeType string

const (
	AccessRuleEditResponseScopeTypeUser         AccessRuleEditResponseScopeType = "user"
	AccessRuleEditResponseScopeTypeOrganization AccessRuleEditResponseScopeType = "organization"
)

func (r AccessRuleEditResponseScopeType) IsKnown() bool {
	switch r {
	case AccessRuleEditResponseScopeTypeUser, AccessRuleEditResponseScopeTypeOrganization:
		return true
	}
	return false
}

type AccessRuleGetResponse struct {
	// The unique identifier of the IP Access rule.
	ID string `json:"id,required"`
	// The available actions that a rule can apply to a matched request.
	AllowedModes []AccessRuleGetResponseAllowedMode `json:"allowed_modes,required"`
	// The rule configuration.
	Configuration AccessRuleGetResponseConfiguration `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode AccessRuleGetResponseMode `json:"mode,required"`
	// The timestamp of when the rule was created.
	CreatedOn time.Time `json:"created_on" format:"date-time"`
	// The timestamp of when the rule was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes string `json:"notes"`
	// All zones owned by the user will have the rule applied.
	Scope AccessRuleGetResponseScope `json:"scope"`
	JSON  accessRuleGetResponseJSON  `json:"-"`
}

// accessRuleGetResponseJSON contains the JSON metadata for the struct
// [AccessRuleGetResponse]
type accessRuleGetResponseJSON struct {
	ID            apijson.Field
	AllowedModes  apijson.Field
	Configuration apijson.Field
	Mode          apijson.Field
	CreatedOn     apijson.Field
	ModifiedOn    apijson.Field
	Notes         apijson.Field
	Scope         apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

// The action to apply to a matched request.
type AccessRuleGetResponseAllowedMode string

const (
	AccessRuleGetResponseAllowedModeBlock            AccessRuleGetResponseAllowedMode = "block"
	AccessRuleGetResponseAllowedModeChallenge        AccessRuleGetResponseAllowedMode = "challenge"
	AccessRuleGetResponseAllowedModeWhitelist        AccessRuleGetResponseAllowedMode = "whitelist"
	AccessRuleGetResponseAllowedModeJSChallenge      AccessRuleGetResponseAllowedMode = "js_challenge"
	AccessRuleGetResponseAllowedModeManagedChallenge AccessRuleGetResponseAllowedMode = "managed_challenge"
)

func (r AccessRuleGetResponseAllowedMode) IsKnown() bool {
	switch r {
	case AccessRuleGetResponseAllowedModeBlock, AccessRuleGetResponseAllowedModeChallenge, AccessRuleGetResponseAllowedModeWhitelist, AccessRuleGetResponseAllowedModeJSChallenge, AccessRuleGetResponseAllowedModeManagedChallenge:
		return true
	}
	return false
}

// The rule configuration.
type AccessRuleGetResponseConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target AccessRuleGetResponseConfigurationTarget `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value string                                 `json:"value"`
	JSON  accessRuleGetResponseConfigurationJSON `json:"-"`
	union AccessRuleGetResponseConfigurationUnion
}

// accessRuleGetResponseConfigurationJSON contains the JSON metadata for the struct
// [AccessRuleGetResponseConfiguration]
type accessRuleGetResponseConfigurationJSON struct {
	Target      apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r accessRuleGetResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

func (r *AccessRuleGetResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	*r = AccessRuleGetResponseConfiguration{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [AccessRuleGetResponseConfigurationUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [AccessRuleIPConfiguration],
// [IPV6Configuration], [AccessRuleCIDRConfiguration], [ASNConfiguration],
// [CountryConfiguration].
func (r AccessRuleGetResponseConfiguration) AsUnion() AccessRuleGetResponseConfigurationUnion {
	return r.union
}

// The rule configuration.
//
// Union satisfied by [AccessRuleIPConfiguration], [IPV6Configuration],
// [AccessRuleCIDRConfiguration], [ASNConfiguration] or [CountryConfiguration].
type AccessRuleGetResponseConfigurationUnion interface {
	implementsAccessRuleGetResponseConfiguration()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessRuleGetResponseConfigurationUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleIPConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPV6Configuration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AccessRuleCIDRConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ASNConfiguration{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(CountryConfiguration{}),
		},
	)
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleGetResponseConfigurationTarget string

const (
	AccessRuleGetResponseConfigurationTargetIP      AccessRuleGetResponseConfigurationTarget = "ip"
	AccessRuleGetResponseConfigurationTargetIp6     AccessRuleGetResponseConfigurationTarget = "ip6"
	AccessRuleGetResponseConfigurationTargetIPRange AccessRuleGetResponseConfigurationTarget = "ip_range"
	AccessRuleGetResponseConfigurationTargetASN     AccessRuleGetResponseConfigurationTarget = "asn"
	AccessRuleGetResponseConfigurationTargetCountry AccessRuleGetResponseConfigurationTarget = "country"
)

func (r AccessRuleGetResponseConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleGetResponseConfigurationTargetIP, AccessRuleGetResponseConfigurationTargetIp6, AccessRuleGetResponseConfigurationTargetIPRange, AccessRuleGetResponseConfigurationTargetASN, AccessRuleGetResponseConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleGetResponseMode string

const (
	AccessRuleGetResponseModeBlock            AccessRuleGetResponseMode = "block"
	AccessRuleGetResponseModeChallenge        AccessRuleGetResponseMode = "challenge"
	AccessRuleGetResponseModeWhitelist        AccessRuleGetResponseMode = "whitelist"
	AccessRuleGetResponseModeJSChallenge      AccessRuleGetResponseMode = "js_challenge"
	AccessRuleGetResponseModeManagedChallenge AccessRuleGetResponseMode = "managed_challenge"
)

func (r AccessRuleGetResponseMode) IsKnown() bool {
	switch r {
	case AccessRuleGetResponseModeBlock, AccessRuleGetResponseModeChallenge, AccessRuleGetResponseModeWhitelist, AccessRuleGetResponseModeJSChallenge, AccessRuleGetResponseModeManagedChallenge:
		return true
	}
	return false
}

// All zones owned by the user will have the rule applied.
type AccessRuleGetResponseScope struct {
	// Defines an identifier.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// Defines the scope of the rule.
	Type AccessRuleGetResponseScopeType `json:"type"`
	JSON accessRuleGetResponseScopeJSON `json:"-"`
}

// accessRuleGetResponseScopeJSON contains the JSON metadata for the struct
// [AccessRuleGetResponseScope]
type accessRuleGetResponseScopeJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleGetResponseScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleGetResponseScopeJSON) RawJSON() string {
	return r.raw
}

// Defines the scope of the rule.
type AccessRuleGetResponseScopeType string

const (
	AccessRuleGetResponseScopeTypeUser         AccessRuleGetResponseScopeType = "user"
	AccessRuleGetResponseScopeTypeOrganization AccessRuleGetResponseScopeType = "organization"
)

func (r AccessRuleGetResponseScopeType) IsKnown() bool {
	switch r {
	case AccessRuleGetResponseScopeTypeUser, AccessRuleGetResponseScopeTypeOrganization:
		return true
	}
	return false
}

type AccessRuleNewParams struct {
	// The rule configuration.
	Configuration param.Field[AccessRuleNewParamsConfigurationUnion] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleNewParamsMode] `json:"mode,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccessRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type AccessRuleNewParamsConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleNewParamsConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleNewParamsConfiguration) implementsAccessRuleNewParamsConfigurationUnion() {}

// The rule configuration.
//
// Satisfied by [firewall.AccessRuleIPConfigurationParam],
// [firewall.IPV6ConfigurationParam], [firewall.AccessRuleCIDRConfigurationParam],
// [firewall.ASNConfigurationParam], [firewall.CountryConfigurationParam],
// [AccessRuleNewParamsConfiguration].
type AccessRuleNewParamsConfigurationUnion interface {
	implementsAccessRuleNewParamsConfigurationUnion()
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleNewParamsConfigurationTarget string

const (
	AccessRuleNewParamsConfigurationTargetIP      AccessRuleNewParamsConfigurationTarget = "ip"
	AccessRuleNewParamsConfigurationTargetIp6     AccessRuleNewParamsConfigurationTarget = "ip6"
	AccessRuleNewParamsConfigurationTargetIPRange AccessRuleNewParamsConfigurationTarget = "ip_range"
	AccessRuleNewParamsConfigurationTargetASN     AccessRuleNewParamsConfigurationTarget = "asn"
	AccessRuleNewParamsConfigurationTargetCountry AccessRuleNewParamsConfigurationTarget = "country"
)

func (r AccessRuleNewParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsConfigurationTargetIP, AccessRuleNewParamsConfigurationTargetIp6, AccessRuleNewParamsConfigurationTargetIPRange, AccessRuleNewParamsConfigurationTargetASN, AccessRuleNewParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleNewParamsMode string

const (
	AccessRuleNewParamsModeBlock            AccessRuleNewParamsMode = "block"
	AccessRuleNewParamsModeChallenge        AccessRuleNewParamsMode = "challenge"
	AccessRuleNewParamsModeWhitelist        AccessRuleNewParamsMode = "whitelist"
	AccessRuleNewParamsModeJSChallenge      AccessRuleNewParamsMode = "js_challenge"
	AccessRuleNewParamsModeManagedChallenge AccessRuleNewParamsMode = "managed_challenge"
)

func (r AccessRuleNewParamsMode) IsKnown() bool {
	switch r {
	case AccessRuleNewParamsModeBlock, AccessRuleNewParamsModeChallenge, AccessRuleNewParamsModeWhitelist, AccessRuleNewParamsModeJSChallenge, AccessRuleNewParamsModeManagedChallenge:
		return true
	}
	return false
}

type AccessRuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AccessRuleNewResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success AccessRuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleNewResponseEnvelopeJSON    `json:"-"`
}

// accessRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleNewResponseEnvelope]
type accessRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AccessRuleNewResponseEnvelopeSuccess bool

const (
	AccessRuleNewResponseEnvelopeSuccessTrue AccessRuleNewResponseEnvelopeSuccess = true
)

func (r AccessRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessRuleListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID        param.Field[string]                            `path:"zone_id"`
	Configuration param.Field[AccessRuleListParamsConfiguration] `query:"configuration"`
	// Defines the direction used to sort returned rules.
	Direction param.Field[AccessRuleListParamsDirection] `query:"direction"`
	// Defines the search requirements. When set to `all`, all the search requirements
	// must match. When set to `any`, only one of the search requirements has to match.
	Match param.Field[AccessRuleListParamsMatch] `query:"match"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleListParamsMode] `query:"mode"`
	// Defines the string to search for in the notes of existing IP Access rules.
	// Notes: For example, the string 'attack' would match IP Access rules with notes
	// 'Attack 26/02' and 'Attack 27/02'. The search is case insensitive.
	Notes param.Field[string] `query:"notes"`
	// Defines the field used to sort returned rules.
	Order param.Field[AccessRuleListParamsOrder] `query:"order"`
	// Defines the requested page within paginated list of results.
	Page param.Field[float64] `query:"page"`
	// Defines the maximum number of results requested.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccessRuleListParams]'s query parameters as `url.Values`.
func (r AccessRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type AccessRuleListParamsConfiguration struct {
	// Defines the target to search in existing rules.
	Target param.Field[AccessRuleListParamsConfigurationTarget] `query:"target"`
	// Defines the target value to search for in existing rules: an IP address, an IP
	// address range, or a country code, depending on the provided
	// `configuration.target`. Notes: You can search for a single IPv4 address, an IP
	// address range with a subnet of '/16' or '/24', or a two-letter ISO-3166-1
	// alpha-2 country code.
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [AccessRuleListParamsConfiguration]'s query parameters as
// `url.Values`.
func (r AccessRuleListParamsConfiguration) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Defines the target to search in existing rules.
type AccessRuleListParamsConfigurationTarget string

const (
	AccessRuleListParamsConfigurationTargetIP      AccessRuleListParamsConfigurationTarget = "ip"
	AccessRuleListParamsConfigurationTargetIPRange AccessRuleListParamsConfigurationTarget = "ip_range"
	AccessRuleListParamsConfigurationTargetASN     AccessRuleListParamsConfigurationTarget = "asn"
	AccessRuleListParamsConfigurationTargetCountry AccessRuleListParamsConfigurationTarget = "country"
)

func (r AccessRuleListParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleListParamsConfigurationTargetIP, AccessRuleListParamsConfigurationTargetIPRange, AccessRuleListParamsConfigurationTargetASN, AccessRuleListParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// Defines the direction used to sort returned rules.
type AccessRuleListParamsDirection string

const (
	AccessRuleListParamsDirectionAsc  AccessRuleListParamsDirection = "asc"
	AccessRuleListParamsDirectionDesc AccessRuleListParamsDirection = "desc"
)

func (r AccessRuleListParamsDirection) IsKnown() bool {
	switch r {
	case AccessRuleListParamsDirectionAsc, AccessRuleListParamsDirectionDesc:
		return true
	}
	return false
}

// Defines the search requirements. When set to `all`, all the search requirements
// must match. When set to `any`, only one of the search requirements has to match.
type AccessRuleListParamsMatch string

const (
	AccessRuleListParamsMatchAny AccessRuleListParamsMatch = "any"
	AccessRuleListParamsMatchAll AccessRuleListParamsMatch = "all"
)

func (r AccessRuleListParamsMatch) IsKnown() bool {
	switch r {
	case AccessRuleListParamsMatchAny, AccessRuleListParamsMatchAll:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleListParamsMode string

const (
	AccessRuleListParamsModeBlock            AccessRuleListParamsMode = "block"
	AccessRuleListParamsModeChallenge        AccessRuleListParamsMode = "challenge"
	AccessRuleListParamsModeWhitelist        AccessRuleListParamsMode = "whitelist"
	AccessRuleListParamsModeJSChallenge      AccessRuleListParamsMode = "js_challenge"
	AccessRuleListParamsModeManagedChallenge AccessRuleListParamsMode = "managed_challenge"
)

func (r AccessRuleListParamsMode) IsKnown() bool {
	switch r {
	case AccessRuleListParamsModeBlock, AccessRuleListParamsModeChallenge, AccessRuleListParamsModeWhitelist, AccessRuleListParamsModeJSChallenge, AccessRuleListParamsModeManagedChallenge:
		return true
	}
	return false
}

// Defines the field used to sort returned rules.
type AccessRuleListParamsOrder string

const (
	AccessRuleListParamsOrderConfigurationTarget AccessRuleListParamsOrder = "configuration.target"
	AccessRuleListParamsOrderConfigurationValue  AccessRuleListParamsOrder = "configuration.value"
	AccessRuleListParamsOrderMode                AccessRuleListParamsOrder = "mode"
)

func (r AccessRuleListParamsOrder) IsKnown() bool {
	switch r {
	case AccessRuleListParamsOrderConfigurationTarget, AccessRuleListParamsOrderConfigurationValue, AccessRuleListParamsOrderMode:
		return true
	}
	return false
}

type AccessRuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessRuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   AccessRuleDeleteResponse `json:"result,required,nullable"`
	// Defines whether the API call was successful.
	Success AccessRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleDeleteResponseEnvelope]
type accessRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AccessRuleDeleteResponseEnvelopeSuccess bool

const (
	AccessRuleDeleteResponseEnvelopeSuccessTrue AccessRuleDeleteResponseEnvelopeSuccess = true
)

func (r AccessRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessRuleEditParams struct {
	// The rule configuration.
	Configuration param.Field[AccessRuleEditParamsConfigurationUnion] `json:"configuration,required"`
	// The action to apply to a matched request.
	Mode param.Field[AccessRuleEditParamsMode] `json:"mode,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative summary of the rule, typically used as a reminder or explanation.
	Notes param.Field[string] `json:"notes"`
}

func (r AccessRuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rule configuration.
type AccessRuleEditParamsConfiguration struct {
	// The configuration target. You must set the target to `ip` when specifying an IP
	// address in the rule.
	Target param.Field[AccessRuleEditParamsConfigurationTarget] `json:"target"`
	// The IP address to match. This address will be compared to the IP address of
	// incoming requests.
	Value param.Field[string] `json:"value"`
}

func (r AccessRuleEditParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessRuleEditParamsConfiguration) implementsAccessRuleEditParamsConfigurationUnion() {}

// The rule configuration.
//
// Satisfied by [firewall.AccessRuleIPConfigurationParam],
// [firewall.IPV6ConfigurationParam], [firewall.AccessRuleCIDRConfigurationParam],
// [firewall.ASNConfigurationParam], [firewall.CountryConfigurationParam],
// [AccessRuleEditParamsConfiguration].
type AccessRuleEditParamsConfigurationUnion interface {
	implementsAccessRuleEditParamsConfigurationUnion()
}

// The configuration target. You must set the target to `ip` when specifying an IP
// address in the rule.
type AccessRuleEditParamsConfigurationTarget string

const (
	AccessRuleEditParamsConfigurationTargetIP      AccessRuleEditParamsConfigurationTarget = "ip"
	AccessRuleEditParamsConfigurationTargetIp6     AccessRuleEditParamsConfigurationTarget = "ip6"
	AccessRuleEditParamsConfigurationTargetIPRange AccessRuleEditParamsConfigurationTarget = "ip_range"
	AccessRuleEditParamsConfigurationTargetASN     AccessRuleEditParamsConfigurationTarget = "asn"
	AccessRuleEditParamsConfigurationTargetCountry AccessRuleEditParamsConfigurationTarget = "country"
)

func (r AccessRuleEditParamsConfigurationTarget) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsConfigurationTargetIP, AccessRuleEditParamsConfigurationTargetIp6, AccessRuleEditParamsConfigurationTargetIPRange, AccessRuleEditParamsConfigurationTargetASN, AccessRuleEditParamsConfigurationTargetCountry:
		return true
	}
	return false
}

// The action to apply to a matched request.
type AccessRuleEditParamsMode string

const (
	AccessRuleEditParamsModeBlock            AccessRuleEditParamsMode = "block"
	AccessRuleEditParamsModeChallenge        AccessRuleEditParamsMode = "challenge"
	AccessRuleEditParamsModeWhitelist        AccessRuleEditParamsMode = "whitelist"
	AccessRuleEditParamsModeJSChallenge      AccessRuleEditParamsMode = "js_challenge"
	AccessRuleEditParamsModeManagedChallenge AccessRuleEditParamsMode = "managed_challenge"
)

func (r AccessRuleEditParamsMode) IsKnown() bool {
	switch r {
	case AccessRuleEditParamsModeBlock, AccessRuleEditParamsModeChallenge, AccessRuleEditParamsModeWhitelist, AccessRuleEditParamsModeJSChallenge, AccessRuleEditParamsModeManagedChallenge:
		return true
	}
	return false
}

type AccessRuleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   AccessRuleEditResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success AccessRuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleEditResponseEnvelopeJSON    `json:"-"`
}

// accessRuleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleEditResponseEnvelope]
type accessRuleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AccessRuleEditResponseEnvelopeSuccess bool

const (
	AccessRuleEditResponseEnvelopeSuccessTrue AccessRuleEditResponseEnvelopeSuccess = true
)

func (r AccessRuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessRuleGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessRuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   AccessRuleGetResponse `json:"result,required"`
	// Defines whether the API call was successful.
	Success AccessRuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessRuleGetResponseEnvelopeJSON    `json:"-"`
}

// accessRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessRuleGetResponseEnvelope]
type accessRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Defines whether the API call was successful.
type AccessRuleGetResponseEnvelopeSuccess bool

const (
	AccessRuleGetResponseEnvelopeSuccessTrue AccessRuleGetResponseEnvelopeSuccess = true
)

func (r AccessRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
