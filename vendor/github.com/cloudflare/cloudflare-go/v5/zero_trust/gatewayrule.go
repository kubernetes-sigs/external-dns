// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// GatewayRuleService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGatewayRuleService] method instead.
type GatewayRuleService struct {
	Options []option.RequestOption
}

// NewGatewayRuleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGatewayRuleService(opts ...option.RequestOption) (r *GatewayRuleService) {
	r = &GatewayRuleService{}
	r.Options = opts
	return
}

// Creates a new Zero Trust Gateway rule.
func (r *GatewayRuleService) New(ctx context.Context, params GatewayRuleNewParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured Zero Trust Gateway rule.
func (r *GatewayRuleService) Update(ctx context.Context, ruleID string, params GatewayRuleUpdateParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) List(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) (res *pagination.SinglePage[GatewayRule], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Fetches the Zero Trust Gateway rules for an account.
func (r *GatewayRuleService) ListAutoPaging(ctx context.Context, query GatewayRuleListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[GatewayRule] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Zero Trust Gateway rule.
func (r *GatewayRuleService) Delete(ctx context.Context, ruleID string, body GatewayRuleDeleteParams, opts ...option.RequestOption) (res *GatewayRuleDeleteResponse, err error) {
	var env GatewayRuleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single Zero Trust Gateway rule.
func (r *GatewayRuleService) Get(ctx context.Context, ruleID string, query GatewayRuleGetParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Resets the expiration of a Zero Trust Gateway Rule if its duration has elapsed
// and it has a default duration.
//
// The Zero Trust Gateway Rule must have values for both `expiration.expires_at`
// and `expiration.duration`.
func (r *GatewayRuleService) ResetExpiration(ctx context.Context, ruleID string, body GatewayRuleResetExpirationParams, opts ...option.RequestOption) (res *GatewayRule, err error) {
	var env GatewayRuleResetExpirationResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/gateway/rules/%s/reset_expiration", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSResolverSettingsV4 struct {
	// IPv4 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                    `json:"vnet_id"`
	JSON   dnsResolverSettingsV4JSON `json:"-"`
}

// dnsResolverSettingsV4JSON contains the JSON metadata for the struct
// [DNSResolverSettingsV4]
type dnsResolverSettingsV4JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DNSResolverSettingsV4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsResolverSettingsV4JSON) RawJSON() string {
	return r.raw
}

type DNSResolverSettingsV4Param struct {
	// IPv4 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r DNSResolverSettingsV4Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSResolverSettingsV6 struct {
	// IPv6 address of upstream resolver.
	IP string `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port int64 `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork bool `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID string                    `json:"vnet_id"`
	JSON   dnsResolverSettingsV6JSON `json:"-"`
}

// dnsResolverSettingsV6JSON contains the JSON metadata for the struct
// [DNSResolverSettingsV6]
type dnsResolverSettingsV6JSON struct {
	IP                         apijson.Field
	Port                       apijson.Field
	RouteThroughPrivateNetwork apijson.Field
	VnetID                     apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *DNSResolverSettingsV6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsResolverSettingsV6JSON) RawJSON() string {
	return r.raw
}

type DNSResolverSettingsV6Param struct {
	// IPv6 address of upstream resolver.
	IP param.Field[string] `json:"ip,required"`
	// A port number to use for upstream resolver. Defaults to 53 if unspecified.
	Port param.Field[int64] `json:"port"`
	// Whether to connect to this resolver over a private network. Must be set when
	// vnet_id is set.
	RouteThroughPrivateNetwork param.Field[bool] `json:"route_through_private_network"`
	// Optionally specify a virtual network for this resolver. Uses default virtual
	// network id if omitted.
	VnetID param.Field[string] `json:"vnet_id"`
}

func (r DNSResolverSettingsV6Param) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The protocol or layer to use.
type GatewayFilter string

const (
	GatewayFilterHTTP        GatewayFilter = "http"
	GatewayFilterDNS         GatewayFilter = "dns"
	GatewayFilterL4          GatewayFilter = "l4"
	GatewayFilterEgress      GatewayFilter = "egress"
	GatewayFilterDNSResolver GatewayFilter = "dns_resolver"
)

func (r GatewayFilter) IsKnown() bool {
	switch r {
	case GatewayFilterHTTP, GatewayFilterDNS, GatewayFilterL4, GatewayFilterEgress, GatewayFilterDNSResolver:
		return true
	}
	return false
}

type GatewayRule struct {
	// The action to perform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action GatewayRuleAction `json:"action,required"`
	// True if the rule is enabled.
	Enabled bool `json:"enabled,required"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters []GatewayFilter `json:"filters,required"`
	// The name of the rule.
	Name string `json:"name,required"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value. Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// docs on how to manage precedence via Terraform.
	Precedence int64 `json:"precedence,required"`
	// The wirefilter expression used for traffic matching. The API automatically
	// formats and sanitizes this expression. This returns a normalized version that
	// may differ from your input and cause Terraform state drift.
	Traffic string `json:"traffic,required"`
	// The API resource UUID.
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Date of deletion, if any.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// The description of the rule.
	Description string `json:"description"`
	// The wirefilter expression used for device posture check matching. The API
	// automatically formats and sanitizes this expression. This returns a normalized
	// version that may differ from your input and cause Terraform state drift.
	DevicePosture string `json:"device_posture"`
	// The expiration time stamp and default duration of a DNS policy. Takes precedence
	// over the policy's `schedule` configuration, if any.
	//
	// This does not apply to HTTP or network policies.
	Expiration GatewayRuleExpiration `json:"expiration,nullable"`
	// The wirefilter expression used for identity matching. The API automatically
	// formats and sanitizes this expression. This returns a normalized version that
	// may differ from your input and cause Terraform state drift.
	Identity string `json:"identity"`
	// The rule cannot be shared via the Orgs API
	NotSharable bool `json:"not_sharable"`
	// The rule was shared via the Orgs API and cannot be edited by the current account
	ReadOnly bool `json:"read_only"`
	// Additional settings that modify the rule's action.
	RuleSettings RuleSetting `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule Schedule `json:"schedule,nullable"`
	// account tag of account that created the rule
	SourceAccount string    `json:"source_account"`
	UpdatedAt     time.Time `json:"updated_at" format:"date-time"`
	// version number of the rule
	Version int64 `json:"version"`
	// Warning for a misconfigured rule, if any.
	WarningStatus string          `json:"warning_status,nullable"`
	JSON          gatewayRuleJSON `json:"-"`
}

// gatewayRuleJSON contains the JSON metadata for the struct [GatewayRule]
type gatewayRuleJSON struct {
	Action        apijson.Field
	Enabled       apijson.Field
	Filters       apijson.Field
	Name          apijson.Field
	Precedence    apijson.Field
	Traffic       apijson.Field
	ID            apijson.Field
	CreatedAt     apijson.Field
	DeletedAt     apijson.Field
	Description   apijson.Field
	DevicePosture apijson.Field
	Expiration    apijson.Field
	Identity      apijson.Field
	NotSharable   apijson.Field
	ReadOnly      apijson.Field
	RuleSettings  apijson.Field
	Schedule      apijson.Field
	SourceAccount apijson.Field
	UpdatedAt     apijson.Field
	Version       apijson.Field
	WarningStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *GatewayRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleAction string

const (
	GatewayRuleActionOn           GatewayRuleAction = "on"
	GatewayRuleActionOff          GatewayRuleAction = "off"
	GatewayRuleActionAllow        GatewayRuleAction = "allow"
	GatewayRuleActionBlock        GatewayRuleAction = "block"
	GatewayRuleActionScan         GatewayRuleAction = "scan"
	GatewayRuleActionNoscan       GatewayRuleAction = "noscan"
	GatewayRuleActionSafesearch   GatewayRuleAction = "safesearch"
	GatewayRuleActionYtrestricted GatewayRuleAction = "ytrestricted"
	GatewayRuleActionIsolate      GatewayRuleAction = "isolate"
	GatewayRuleActionNoisolate    GatewayRuleAction = "noisolate"
	GatewayRuleActionOverride     GatewayRuleAction = "override"
	GatewayRuleActionL4Override   GatewayRuleAction = "l4_override"
	GatewayRuleActionEgress       GatewayRuleAction = "egress"
	GatewayRuleActionResolve      GatewayRuleAction = "resolve"
	GatewayRuleActionQuarantine   GatewayRuleAction = "quarantine"
	GatewayRuleActionRedirect     GatewayRuleAction = "redirect"
)

func (r GatewayRuleAction) IsKnown() bool {
	switch r {
	case GatewayRuleActionOn, GatewayRuleActionOff, GatewayRuleActionAllow, GatewayRuleActionBlock, GatewayRuleActionScan, GatewayRuleActionNoscan, GatewayRuleActionSafesearch, GatewayRuleActionYtrestricted, GatewayRuleActionIsolate, GatewayRuleActionNoisolate, GatewayRuleActionOverride, GatewayRuleActionL4Override, GatewayRuleActionEgress, GatewayRuleActionResolve, GatewayRuleActionQuarantine, GatewayRuleActionRedirect:
		return true
	}
	return false
}

// The expiration time stamp and default duration of a DNS policy. Takes precedence
// over the policy's `schedule` configuration, if any.
//
// This does not apply to HTTP or network policies.
type GatewayRuleExpiration struct {
	// The time stamp at which the policy will expire and cease to be applied.
	//
	// Must adhere to RFC 3339 and include a UTC offset. Non-zero offsets are accepted
	// but will be converted to the equivalent value with offset zero (UTC+00:00) and
	// will be returned as time stamps with offset zero denoted by a trailing 'Z'.
	//
	// Policies with an expiration do not consider the timezone of clients they are
	// applied to, and expire "globally" at the point given by their `expires_at`
	// value.
	ExpiresAt time.Time `json:"expires_at,required" format:"date-time"`
	// The default duration a policy will be active in minutes. Must be set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration int64 `json:"duration"`
	// Whether the policy has expired.
	Expired bool                      `json:"expired"`
	JSON    gatewayRuleExpirationJSON `json:"-"`
}

// gatewayRuleExpirationJSON contains the JSON metadata for the struct
// [GatewayRuleExpiration]
type gatewayRuleExpirationJSON struct {
	ExpiresAt   apijson.Field
	Duration    apijson.Field
	Expired     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleExpiration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleExpirationJSON) RawJSON() string {
	return r.raw
}

// Additional settings that modify the rule's action.
type RuleSetting struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders map[string][]string `json:"add_headers,nullable"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass bool `json:"allow_child_bypass,nullable"`
	// Settings for the Audit SSH action.
	AuditSSH RuleSettingAuditSSH `json:"audit_ssh,nullable"`
	// Configure how browser isolation behaves.
	BISOAdminControls RuleSettingBISOAdminControls `json:"biso_admin_controls"`
	// Custom block page settings. If missing/null, blocking will use the the account
	// settings.
	BlockPage RuleSettingBlockPage `json:"block_page,nullable"`
	// Enable the custom block page.
	BlockPageEnabled bool `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason string `json:"block_reason,nullable"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule bool `json:"bypass_parent_rule,nullable"`
	// Configure how session check behaves.
	CheckSession RuleSettingCheckSession `json:"check_session,nullable"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
	// are set. DNS queries will route to the address closest to their origin. Only
	// valid when a rule's action is set to 'resolve'.
	DNSResolvers RuleSettingDNSResolvers `json:"dns_resolvers,nullable"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress RuleSettingEgress `json:"egress,nullable"`
	// Set to true, to ignore the category matches at CNAME domains in a response. If
	// unchecked, the categories in this rule will be checked against all the CNAME
	// domain categories in a response.
	IgnoreCNAMECategoryMatches bool `json:"ignore_cname_category_matches"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories bool `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds bool `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override RuleSettingL4override `json:"l4override,nullable"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings RuleSettingNotificationSettings `json:"notification_settings,nullable"`
	// Override matching DNS queries with a hostname.
	OverrideHost string `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs []string `json:"override_ips,nullable"`
	// Configure DLP payload logging.
	PayloadLog RuleSettingPayloadLog `json:"payload_log,nullable"`
	// Settings that apply to quarantine rules
	Quarantine RuleSettingQuarantine `json:"quarantine,nullable"`
	// Settings that apply to redirect rules
	Redirect RuleSettingRedirect `json:"redirect,nullable"`
	// Configure to forward the query to the internal DNS service, passing the
	// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
	// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
	// set to 'resolve'.
	ResolveDNSInternally RuleSettingResolveDNSInternally `json:"resolve_dns_internally,nullable"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when 'dns_resolvers' are specified or
	// 'resolve_dns_internally' is set. Only valid when a rule's action is set to
	// 'resolve'.
	ResolveDNSThroughCloudflare bool `json:"resolve_dns_through_cloudflare,nullable"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCERT RuleSettingUntrustedCERT `json:"untrusted_cert,nullable"`
	JSON          ruleSettingJSON          `json:"-"`
}

// ruleSettingJSON contains the JSON metadata for the struct [RuleSetting]
type ruleSettingJSON struct {
	AddHeaders                      apijson.Field
	AllowChildBypass                apijson.Field
	AuditSSH                        apijson.Field
	BISOAdminControls               apijson.Field
	BlockPage                       apijson.Field
	BlockPageEnabled                apijson.Field
	BlockReason                     apijson.Field
	BypassParentRule                apijson.Field
	CheckSession                    apijson.Field
	DNSResolvers                    apijson.Field
	Egress                          apijson.Field
	IgnoreCNAMECategoryMatches      apijson.Field
	InsecureDisableDNSSECValidation apijson.Field
	IPCategories                    apijson.Field
	IPIndicatorFeeds                apijson.Field
	L4override                      apijson.Field
	NotificationSettings            apijson.Field
	OverrideHost                    apijson.Field
	OverrideIPs                     apijson.Field
	PayloadLog                      apijson.Field
	Quarantine                      apijson.Field
	Redirect                        apijson.Field
	ResolveDNSInternally            apijson.Field
	ResolveDNSThroughCloudflare     apijson.Field
	UntrustedCERT                   apijson.Field
	raw                             string
	ExtraFields                     map[string]apijson.Field
}

func (r *RuleSetting) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingJSON) RawJSON() string {
	return r.raw
}

// Settings for the Audit SSH action.
type RuleSettingAuditSSH struct {
	// Enable to turn on SSH command logging.
	CommandLogging bool                    `json:"command_logging"`
	JSON           ruleSettingAuditSSHJSON `json:"-"`
}

// ruleSettingAuditSSHJSON contains the JSON metadata for the struct
// [RuleSettingAuditSSH]
type ruleSettingAuditSSHJSON struct {
	CommandLogging apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingAuditSSH) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingAuditSSHJSON) RawJSON() string {
	return r.raw
}

// Configure how browser isolation behaves.
type RuleSettingBISOAdminControls struct {
	// Configure whether copy is enabled or not. When set with "remote_only", copying
	// isolated content from the remote browser to the user's local clipboard is
	// disabled. When absent, copy is enabled. Only applies when `version == "v2"`.
	Copy RuleSettingBISOAdminControlsCopy `json:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	DCP bool `json:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	DD bool `json:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	DK bool `json:"dk"`
	// Configure whether downloading enabled or not. When set with "remote_only",
	// downloads are only available for viewing. Only applies when `version == "v2"`.
	Download RuleSettingBISOAdminControlsDownload `json:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	DP bool `json:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	DU bool `json:"du"`
	// Configure whether keyboard usage is enabled or not. When absent, keyboard usage
	// is enabled. Only applies when `version == "v2"`.
	Keyboard RuleSettingBISOAdminControlsKeyboard `json:"keyboard"`
	// Configure whether pasting is enabled or not. When set with "remote_only",
	// pasting content from the user's local clipboard into isolated pages is disabled.
	// When absent, paste is enabled. Only applies when `version == "v2"`.
	Paste RuleSettingBISOAdminControlsPaste `json:"paste"`
	// Configure whether printing is enabled or not. When absent, printing is enabled.
	// Only applies when `version == "v2"`.
	Printing RuleSettingBISOAdminControlsPrinting `json:"printing"`
	// Configure whether uploading is enabled or not. When absent, uploading is
	// enabled. Only applies when `version == "v2"`.
	Upload RuleSettingBISOAdminControlsUpload `json:"upload"`
	// Indicates which version of the browser isolation controls should apply.
	Version RuleSettingBISOAdminControlsVersion `json:"version"`
	JSON    ruleSettingBISOAdminControlsJSON    `json:"-"`
}

// ruleSettingBISOAdminControlsJSON contains the JSON metadata for the struct
// [RuleSettingBISOAdminControls]
type ruleSettingBISOAdminControlsJSON struct {
	Copy        apijson.Field
	DCP         apijson.Field
	DD          apijson.Field
	DK          apijson.Field
	Download    apijson.Field
	DP          apijson.Field
	DU          apijson.Field
	Keyboard    apijson.Field
	Paste       apijson.Field
	Printing    apijson.Field
	Upload      apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingBISOAdminControls) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingBISOAdminControlsJSON) RawJSON() string {
	return r.raw
}

// Configure whether copy is enabled or not. When set with "remote_only", copying
// isolated content from the remote browser to the user's local clipboard is
// disabled. When absent, copy is enabled. Only applies when `version == "v2"`.
type RuleSettingBISOAdminControlsCopy string

const (
	RuleSettingBISOAdminControlsCopyEnabled    RuleSettingBISOAdminControlsCopy = "enabled"
	RuleSettingBISOAdminControlsCopyDisabled   RuleSettingBISOAdminControlsCopy = "disabled"
	RuleSettingBISOAdminControlsCopyRemoteOnly RuleSettingBISOAdminControlsCopy = "remote_only"
)

func (r RuleSettingBISOAdminControlsCopy) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsCopyEnabled, RuleSettingBISOAdminControlsCopyDisabled, RuleSettingBISOAdminControlsCopyRemoteOnly:
		return true
	}
	return false
}

// Configure whether downloading enabled or not. When set with "remote_only",
// downloads are only available for viewing. Only applies when `version == "v2"`.
type RuleSettingBISOAdminControlsDownload string

const (
	RuleSettingBISOAdminControlsDownloadEnabled    RuleSettingBISOAdminControlsDownload = "enabled"
	RuleSettingBISOAdminControlsDownloadDisabled   RuleSettingBISOAdminControlsDownload = "disabled"
	RuleSettingBISOAdminControlsDownloadRemoteOnly RuleSettingBISOAdminControlsDownload = "remote_only"
)

func (r RuleSettingBISOAdminControlsDownload) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsDownloadEnabled, RuleSettingBISOAdminControlsDownloadDisabled, RuleSettingBISOAdminControlsDownloadRemoteOnly:
		return true
	}
	return false
}

// Configure whether keyboard usage is enabled or not. When absent, keyboard usage
// is enabled. Only applies when `version == "v2"`.
type RuleSettingBISOAdminControlsKeyboard string

const (
	RuleSettingBISOAdminControlsKeyboardEnabled  RuleSettingBISOAdminControlsKeyboard = "enabled"
	RuleSettingBISOAdminControlsKeyboardDisabled RuleSettingBISOAdminControlsKeyboard = "disabled"
)

func (r RuleSettingBISOAdminControlsKeyboard) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsKeyboardEnabled, RuleSettingBISOAdminControlsKeyboardDisabled:
		return true
	}
	return false
}

// Configure whether pasting is enabled or not. When set with "remote_only",
// pasting content from the user's local clipboard into isolated pages is disabled.
// When absent, paste is enabled. Only applies when `version == "v2"`.
type RuleSettingBISOAdminControlsPaste string

const (
	RuleSettingBISOAdminControlsPasteEnabled    RuleSettingBISOAdminControlsPaste = "enabled"
	RuleSettingBISOAdminControlsPasteDisabled   RuleSettingBISOAdminControlsPaste = "disabled"
	RuleSettingBISOAdminControlsPasteRemoteOnly RuleSettingBISOAdminControlsPaste = "remote_only"
)

func (r RuleSettingBISOAdminControlsPaste) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsPasteEnabled, RuleSettingBISOAdminControlsPasteDisabled, RuleSettingBISOAdminControlsPasteRemoteOnly:
		return true
	}
	return false
}

// Configure whether printing is enabled or not. When absent, printing is enabled.
// Only applies when `version == "v2"`.
type RuleSettingBISOAdminControlsPrinting string

const (
	RuleSettingBISOAdminControlsPrintingEnabled  RuleSettingBISOAdminControlsPrinting = "enabled"
	RuleSettingBISOAdminControlsPrintingDisabled RuleSettingBISOAdminControlsPrinting = "disabled"
)

func (r RuleSettingBISOAdminControlsPrinting) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsPrintingEnabled, RuleSettingBISOAdminControlsPrintingDisabled:
		return true
	}
	return false
}

// Configure whether uploading is enabled or not. When absent, uploading is
// enabled. Only applies when `version == "v2"`.
type RuleSettingBISOAdminControlsUpload string

const (
	RuleSettingBISOAdminControlsUploadEnabled  RuleSettingBISOAdminControlsUpload = "enabled"
	RuleSettingBISOAdminControlsUploadDisabled RuleSettingBISOAdminControlsUpload = "disabled"
)

func (r RuleSettingBISOAdminControlsUpload) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsUploadEnabled, RuleSettingBISOAdminControlsUploadDisabled:
		return true
	}
	return false
}

// Indicates which version of the browser isolation controls should apply.
type RuleSettingBISOAdminControlsVersion string

const (
	RuleSettingBISOAdminControlsVersionV1 RuleSettingBISOAdminControlsVersion = "v1"
	RuleSettingBISOAdminControlsVersionV2 RuleSettingBISOAdminControlsVersion = "v2"
)

func (r RuleSettingBISOAdminControlsVersion) IsKnown() bool {
	switch r {
	case RuleSettingBISOAdminControlsVersionV1, RuleSettingBISOAdminControlsVersionV2:
		return true
	}
	return false
}

// Custom block page settings. If missing/null, blocking will use the the account
// settings.
type RuleSettingBlockPage struct {
	// URI to which the user will be redirected
	TargetURI string `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext bool                     `json:"include_context"`
	JSON           ruleSettingBlockPageJSON `json:"-"`
}

// ruleSettingBlockPageJSON contains the JSON metadata for the struct
// [RuleSettingBlockPage]
type ruleSettingBlockPageJSON struct {
	TargetURI      apijson.Field
	IncludeContext apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingBlockPage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingBlockPageJSON) RawJSON() string {
	return r.raw
}

// Configure how session check behaves.
type RuleSettingCheckSession struct {
	// Configure how fresh the session needs to be to be considered valid. The API
	// automatically formats and sanitizes this expression. This returns a normalized
	// version that may differ from your input and cause Terraform state drift.
	Duration string `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce bool                        `json:"enforce"`
	JSON    ruleSettingCheckSessionJSON `json:"-"`
}

// ruleSettingCheckSessionJSON contains the JSON metadata for the struct
// [RuleSettingCheckSession]
type ruleSettingCheckSessionJSON struct {
	Duration    apijson.Field
	Enforce     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingCheckSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingCheckSessionJSON) RawJSON() string {
	return r.raw
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
// are set. DNS queries will route to the address closest to their origin. Only
// valid when a rule's action is set to 'resolve'.
type RuleSettingDNSResolvers struct {
	IPV4 []DNSResolverSettingsV4     `json:"ipv4"`
	IPV6 []DNSResolverSettingsV6     `json:"ipv6"`
	JSON ruleSettingDNSResolversJSON `json:"-"`
}

// ruleSettingDNSResolversJSON contains the JSON metadata for the struct
// [RuleSettingDNSResolvers]
type ruleSettingDNSResolversJSON struct {
	IPV4        apijson.Field
	IPV6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingDNSResolvers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingDNSResolversJSON) RawJSON() string {
	return r.raw
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type RuleSettingEgress struct {
	// The IPv4 address to be used for egress.
	IPV4 string `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback string `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 string                `json:"ipv6"`
	JSON ruleSettingEgressJSON `json:"-"`
}

// ruleSettingEgressJSON contains the JSON metadata for the struct
// [RuleSettingEgress]
type ruleSettingEgressJSON struct {
	IPV4         apijson.Field
	IPV4Fallback apijson.Field
	IPV6         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RuleSettingEgress) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingEgressJSON) RawJSON() string {
	return r.raw
}

// Send matching traffic to the supplied destination IP address and port.
type RuleSettingL4override struct {
	// IPv4 or IPv6 address.
	IP string `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port int64                     `json:"port"`
	JSON ruleSettingL4overrideJSON `json:"-"`
}

// ruleSettingL4overrideJSON contains the JSON metadata for the struct
// [RuleSettingL4override]
type ruleSettingL4overrideJSON struct {
	IP          apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingL4override) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingL4overrideJSON) RawJSON() string {
	return r.raw
}

// Configure a notification to display on the user's device when this rule is
// matched.
type RuleSettingNotificationSettings struct {
	// Set notification on
	Enabled bool `json:"enabled"`
	// If true, context information will be passed as query parameters
	IncludeContext bool `json:"include_context"`
	// Customize the message shown in the notification.
	Msg string `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL string                              `json:"support_url"`
	JSON       ruleSettingNotificationSettingsJSON `json:"-"`
}

// ruleSettingNotificationSettingsJSON contains the JSON metadata for the struct
// [RuleSettingNotificationSettings]
type ruleSettingNotificationSettingsJSON struct {
	Enabled        apijson.Field
	IncludeContext apijson.Field
	Msg            apijson.Field
	SupportURL     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleSettingNotificationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingNotificationSettingsJSON) RawJSON() string {
	return r.raw
}

// Configure DLP payload logging.
type RuleSettingPayloadLog struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled bool                      `json:"enabled"`
	JSON    ruleSettingPayloadLogJSON `json:"-"`
}

// ruleSettingPayloadLogJSON contains the JSON metadata for the struct
// [RuleSettingPayloadLog]
type ruleSettingPayloadLogJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingPayloadLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingPayloadLogJSON) RawJSON() string {
	return r.raw
}

// Settings that apply to quarantine rules
type RuleSettingQuarantine struct {
	// Types of files to sandbox.
	FileTypes []RuleSettingQuarantineFileType `json:"file_types"`
	JSON      ruleSettingQuarantineJSON       `json:"-"`
}

// ruleSettingQuarantineJSON contains the JSON metadata for the struct
// [RuleSettingQuarantine]
type ruleSettingQuarantineJSON struct {
	FileTypes   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingQuarantine) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingQuarantineJSON) RawJSON() string {
	return r.raw
}

type RuleSettingQuarantineFileType string

const (
	RuleSettingQuarantineFileTypeExe  RuleSettingQuarantineFileType = "exe"
	RuleSettingQuarantineFileTypePDF  RuleSettingQuarantineFileType = "pdf"
	RuleSettingQuarantineFileTypeDoc  RuleSettingQuarantineFileType = "doc"
	RuleSettingQuarantineFileTypeDocm RuleSettingQuarantineFileType = "docm"
	RuleSettingQuarantineFileTypeDocx RuleSettingQuarantineFileType = "docx"
	RuleSettingQuarantineFileTypeRtf  RuleSettingQuarantineFileType = "rtf"
	RuleSettingQuarantineFileTypePpt  RuleSettingQuarantineFileType = "ppt"
	RuleSettingQuarantineFileTypePptx RuleSettingQuarantineFileType = "pptx"
	RuleSettingQuarantineFileTypeXls  RuleSettingQuarantineFileType = "xls"
	RuleSettingQuarantineFileTypeXlsm RuleSettingQuarantineFileType = "xlsm"
	RuleSettingQuarantineFileTypeXlsx RuleSettingQuarantineFileType = "xlsx"
	RuleSettingQuarantineFileTypeZip  RuleSettingQuarantineFileType = "zip"
	RuleSettingQuarantineFileTypeRar  RuleSettingQuarantineFileType = "rar"
)

func (r RuleSettingQuarantineFileType) IsKnown() bool {
	switch r {
	case RuleSettingQuarantineFileTypeExe, RuleSettingQuarantineFileTypePDF, RuleSettingQuarantineFileTypeDoc, RuleSettingQuarantineFileTypeDocm, RuleSettingQuarantineFileTypeDocx, RuleSettingQuarantineFileTypeRtf, RuleSettingQuarantineFileTypePpt, RuleSettingQuarantineFileTypePptx, RuleSettingQuarantineFileTypeXls, RuleSettingQuarantineFileTypeXlsm, RuleSettingQuarantineFileTypeXlsx, RuleSettingQuarantineFileTypeZip, RuleSettingQuarantineFileTypeRar:
		return true
	}
	return false
}

// Settings that apply to redirect rules
type RuleSettingRedirect struct {
	// URI to which the user will be redirected
	TargetURI string `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext bool `json:"include_context"`
	// If true, the path and query parameters from the original request will be
	// appended to target_uri
	PreservePathAndQuery bool                    `json:"preserve_path_and_query"`
	JSON                 ruleSettingRedirectJSON `json:"-"`
}

// ruleSettingRedirectJSON contains the JSON metadata for the struct
// [RuleSettingRedirect]
type ruleSettingRedirectJSON struct {
	TargetURI            apijson.Field
	IncludeContext       apijson.Field
	PreservePathAndQuery apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RuleSettingRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingRedirectJSON) RawJSON() string {
	return r.raw
}

// Configure to forward the query to the internal DNS service, passing the
// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
// set to 'resolve'.
type RuleSettingResolveDNSInternally struct {
	// The fallback behavior to apply when the internal DNS response code is different
	// from 'NOERROR' or when the response data only contains CNAME records for 'A' or
	// 'AAAA' queries.
	Fallback RuleSettingResolveDNSInternallyFallback `json:"fallback"`
	// The internal DNS view identifier that's passed to the internal DNS service.
	ViewID string                              `json:"view_id"`
	JSON   ruleSettingResolveDNSInternallyJSON `json:"-"`
}

// ruleSettingResolveDNSInternallyJSON contains the JSON metadata for the struct
// [RuleSettingResolveDNSInternally]
type ruleSettingResolveDNSInternallyJSON struct {
	Fallback    apijson.Field
	ViewID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingResolveDNSInternally) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingResolveDNSInternallyJSON) RawJSON() string {
	return r.raw
}

// The fallback behavior to apply when the internal DNS response code is different
// from 'NOERROR' or when the response data only contains CNAME records for 'A' or
// 'AAAA' queries.
type RuleSettingResolveDNSInternallyFallback string

const (
	RuleSettingResolveDNSInternallyFallbackNone      RuleSettingResolveDNSInternallyFallback = "none"
	RuleSettingResolveDNSInternallyFallbackPublicDNS RuleSettingResolveDNSInternallyFallback = "public_dns"
)

func (r RuleSettingResolveDNSInternallyFallback) IsKnown() bool {
	switch r {
	case RuleSettingResolveDNSInternallyFallbackNone, RuleSettingResolveDNSInternallyFallbackPublicDNS:
		return true
	}
	return false
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type RuleSettingUntrustedCERT struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action RuleSettingUntrustedCERTAction `json:"action"`
	JSON   ruleSettingUntrustedCERTJSON   `json:"-"`
}

// ruleSettingUntrustedCERTJSON contains the JSON metadata for the struct
// [RuleSettingUntrustedCERT]
type ruleSettingUntrustedCERTJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleSettingUntrustedCERT) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleSettingUntrustedCERTJSON) RawJSON() string {
	return r.raw
}

// The action performed when an untrusted certificate is seen. The default action
// is an error with HTTP code 526.
type RuleSettingUntrustedCERTAction string

const (
	RuleSettingUntrustedCERTActionPassThrough RuleSettingUntrustedCERTAction = "pass_through"
	RuleSettingUntrustedCERTActionBlock       RuleSettingUntrustedCERTAction = "block"
	RuleSettingUntrustedCERTActionError       RuleSettingUntrustedCERTAction = "error"
)

func (r RuleSettingUntrustedCERTAction) IsKnown() bool {
	switch r {
	case RuleSettingUntrustedCERTActionPassThrough, RuleSettingUntrustedCERTActionBlock, RuleSettingUntrustedCERTActionError:
		return true
	}
	return false
}

// Additional settings that modify the rule's action.
type RuleSettingParam struct {
	// Add custom headers to allowed requests, in the form of key-value pairs. Keys are
	// header names, pointing to an array with its header value(s).
	AddHeaders param.Field[map[string][]string] `json:"add_headers"`
	// Set by parent MSP accounts to enable their children to bypass this rule.
	AllowChildBypass param.Field[bool] `json:"allow_child_bypass"`
	// Settings for the Audit SSH action.
	AuditSSH param.Field[RuleSettingAuditSSHParam] `json:"audit_ssh"`
	// Configure how browser isolation behaves.
	BISOAdminControls param.Field[RuleSettingBISOAdminControlsParam] `json:"biso_admin_controls"`
	// Custom block page settings. If missing/null, blocking will use the the account
	// settings.
	BlockPage param.Field[RuleSettingBlockPageParam] `json:"block_page"`
	// Enable the custom block page.
	BlockPageEnabled param.Field[bool] `json:"block_page_enabled"`
	// The text describing why this block occurred, displayed on the custom block page
	// (if enabled).
	BlockReason param.Field[string] `json:"block_reason"`
	// Set by children MSP accounts to bypass their parent's rules.
	BypassParentRule param.Field[bool] `json:"bypass_parent_rule"`
	// Configure how session check behaves.
	CheckSession param.Field[RuleSettingCheckSessionParam] `json:"check_session"`
	// Add your own custom resolvers to route queries that match the resolver policy.
	// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
	// are set. DNS queries will route to the address closest to their origin. Only
	// valid when a rule's action is set to 'resolve'.
	DNSResolvers param.Field[RuleSettingDNSResolversParam] `json:"dns_resolvers"`
	// Configure how Gateway Proxy traffic egresses. You can enable this setting for
	// rules with Egress actions and filters, or omit it to indicate local egress via
	// WARP IPs.
	Egress param.Field[RuleSettingEgressParam] `json:"egress"`
	// Set to true, to ignore the category matches at CNAME domains in a response. If
	// unchecked, the categories in this rule will be checked against all the CNAME
	// domain categories in a response.
	IgnoreCNAMECategoryMatches param.Field[bool] `json:"ignore_cname_category_matches"`
	// INSECURE - disable DNSSEC validation (for Allow actions).
	InsecureDisableDNSSECValidation param.Field[bool] `json:"insecure_disable_dnssec_validation"`
	// Set to true to enable IPs in DNS resolver category blocks. By default categories
	// only block based on domain names.
	IPCategories param.Field[bool] `json:"ip_categories"`
	// Set to true to include IPs in DNS resolver indicator feed blocks. By default
	// indicator feeds only block based on domain names.
	IPIndicatorFeeds param.Field[bool] `json:"ip_indicator_feeds"`
	// Send matching traffic to the supplied destination IP address and port.
	L4override param.Field[RuleSettingL4overrideParam] `json:"l4override"`
	// Configure a notification to display on the user's device when this rule is
	// matched.
	NotificationSettings param.Field[RuleSettingNotificationSettingsParam] `json:"notification_settings"`
	// Override matching DNS queries with a hostname.
	OverrideHost param.Field[string] `json:"override_host"`
	// Override matching DNS queries with an IP or set of IPs.
	OverrideIPs param.Field[[]string] `json:"override_ips"`
	// Configure DLP payload logging.
	PayloadLog param.Field[RuleSettingPayloadLogParam] `json:"payload_log"`
	// Settings that apply to quarantine rules
	Quarantine param.Field[RuleSettingQuarantineParam] `json:"quarantine"`
	// Settings that apply to redirect rules
	Redirect param.Field[RuleSettingRedirectParam] `json:"redirect"`
	// Configure to forward the query to the internal DNS service, passing the
	// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
	// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
	// set to 'resolve'.
	ResolveDNSInternally param.Field[RuleSettingResolveDNSInternallyParam] `json:"resolve_dns_internally"`
	// Enable to send queries that match the policy to Cloudflare's default 1.1.1.1 DNS
	// resolver. Cannot be set when 'dns_resolvers' are specified or
	// 'resolve_dns_internally' is set. Only valid when a rule's action is set to
	// 'resolve'.
	ResolveDNSThroughCloudflare param.Field[bool] `json:"resolve_dns_through_cloudflare"`
	// Configure behavior when an upstream cert is invalid or an SSL error occurs.
	UntrustedCERT param.Field[RuleSettingUntrustedCERTParam] `json:"untrusted_cert"`
}

func (r RuleSettingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for the Audit SSH action.
type RuleSettingAuditSSHParam struct {
	// Enable to turn on SSH command logging.
	CommandLogging param.Field[bool] `json:"command_logging"`
}

func (r RuleSettingAuditSSHParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how browser isolation behaves.
type RuleSettingBISOAdminControlsParam struct {
	// Configure whether copy is enabled or not. When set with "remote_only", copying
	// isolated content from the remote browser to the user's local clipboard is
	// disabled. When absent, copy is enabled. Only applies when `version == "v2"`.
	Copy param.Field[RuleSettingBISOAdminControlsCopy] `json:"copy"`
	// Set to false to enable copy-pasting. Only applies when `version == "v1"`.
	DCP param.Field[bool] `json:"dcp"`
	// Set to false to enable downloading. Only applies when `version == "v1"`.
	DD param.Field[bool] `json:"dd"`
	// Set to false to enable keyboard usage. Only applies when `version == "v1"`.
	DK param.Field[bool] `json:"dk"`
	// Configure whether downloading enabled or not. When set with "remote_only",
	// downloads are only available for viewing. Only applies when `version == "v2"`.
	Download param.Field[RuleSettingBISOAdminControlsDownload] `json:"download"`
	// Set to false to enable printing. Only applies when `version == "v1"`.
	DP param.Field[bool] `json:"dp"`
	// Set to false to enable uploading. Only applies when `version == "v1"`.
	DU param.Field[bool] `json:"du"`
	// Configure whether keyboard usage is enabled or not. When absent, keyboard usage
	// is enabled. Only applies when `version == "v2"`.
	Keyboard param.Field[RuleSettingBISOAdminControlsKeyboard] `json:"keyboard"`
	// Configure whether pasting is enabled or not. When set with "remote_only",
	// pasting content from the user's local clipboard into isolated pages is disabled.
	// When absent, paste is enabled. Only applies when `version == "v2"`.
	Paste param.Field[RuleSettingBISOAdminControlsPaste] `json:"paste"`
	// Configure whether printing is enabled or not. When absent, printing is enabled.
	// Only applies when `version == "v2"`.
	Printing param.Field[RuleSettingBISOAdminControlsPrinting] `json:"printing"`
	// Configure whether uploading is enabled or not. When absent, uploading is
	// enabled. Only applies when `version == "v2"`.
	Upload param.Field[RuleSettingBISOAdminControlsUpload] `json:"upload"`
	// Indicates which version of the browser isolation controls should apply.
	Version param.Field[RuleSettingBISOAdminControlsVersion] `json:"version"`
}

func (r RuleSettingBISOAdminControlsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Custom block page settings. If missing/null, blocking will use the the account
// settings.
type RuleSettingBlockPageParam struct {
	// URI to which the user will be redirected
	TargetURI param.Field[string] `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
}

func (r RuleSettingBlockPageParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how session check behaves.
type RuleSettingCheckSessionParam struct {
	// Configure how fresh the session needs to be to be considered valid. The API
	// automatically formats and sanitizes this expression. This returns a normalized
	// version that may differ from your input and cause Terraform state drift.
	Duration param.Field[string] `json:"duration"`
	// Set to true to enable session enforcement.
	Enforce param.Field[bool] `json:"enforce"`
}

func (r RuleSettingCheckSessionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Add your own custom resolvers to route queries that match the resolver policy.
// Cannot be used when 'resolve_dns_through_cloudflare' or 'resolve_dns_internally'
// are set. DNS queries will route to the address closest to their origin. Only
// valid when a rule's action is set to 'resolve'.
type RuleSettingDNSResolversParam struct {
	IPV4 param.Field[[]DNSResolverSettingsV4Param] `json:"ipv4"`
	IPV6 param.Field[[]DNSResolverSettingsV6Param] `json:"ipv6"`
}

func (r RuleSettingDNSResolversParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure how Gateway Proxy traffic egresses. You can enable this setting for
// rules with Egress actions and filters, or omit it to indicate local egress via
// WARP IPs.
type RuleSettingEgressParam struct {
	// The IPv4 address to be used for egress.
	IPV4 param.Field[string] `json:"ipv4"`
	// The fallback IPv4 address to be used for egress in the event of an error
	// egressing with the primary IPv4. Can be '0.0.0.0' to indicate local egress via
	// WARP IPs.
	IPV4Fallback param.Field[string] `json:"ipv4_fallback"`
	// The IPv6 range to be used for egress.
	IPV6 param.Field[string] `json:"ipv6"`
}

func (r RuleSettingEgressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Send matching traffic to the supplied destination IP address and port.
type RuleSettingL4overrideParam struct {
	// IPv4 or IPv6 address.
	IP param.Field[string] `json:"ip"`
	// A port number to use for TCP/UDP overrides.
	Port param.Field[int64] `json:"port"`
}

func (r RuleSettingL4overrideParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure a notification to display on the user's device when this rule is
// matched.
type RuleSettingNotificationSettingsParam struct {
	// Set notification on
	Enabled param.Field[bool] `json:"enabled"`
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
	// Customize the message shown in the notification.
	Msg param.Field[string] `json:"msg"`
	// Optional URL to direct users to additional information. If not set, the
	// notification will open a block page.
	SupportURL param.Field[string] `json:"support_url"`
}

func (r RuleSettingNotificationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure DLP payload logging.
type RuleSettingPayloadLogParam struct {
	// Set to true to enable DLP payload logging for this rule.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r RuleSettingPayloadLogParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings that apply to quarantine rules
type RuleSettingQuarantineParam struct {
	// Types of files to sandbox.
	FileTypes param.Field[[]RuleSettingQuarantineFileType] `json:"file_types"`
}

func (r RuleSettingQuarantineParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings that apply to redirect rules
type RuleSettingRedirectParam struct {
	// URI to which the user will be redirected
	TargetURI param.Field[string] `json:"target_uri,required" format:"uri"`
	// If true, context information will be passed as query parameters
	IncludeContext param.Field[bool] `json:"include_context"`
	// If true, the path and query parameters from the original request will be
	// appended to target_uri
	PreservePathAndQuery param.Field[bool] `json:"preserve_path_and_query"`
}

func (r RuleSettingRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure to forward the query to the internal DNS service, passing the
// specified 'view_id' as input. Cannot be set when 'dns_resolvers' are specified
// or 'resolve_dns_through_cloudflare' is set. Only valid when a rule's action is
// set to 'resolve'.
type RuleSettingResolveDNSInternallyParam struct {
	// The fallback behavior to apply when the internal DNS response code is different
	// from 'NOERROR' or when the response data only contains CNAME records for 'A' or
	// 'AAAA' queries.
	Fallback param.Field[RuleSettingResolveDNSInternallyFallback] `json:"fallback"`
	// The internal DNS view identifier that's passed to the internal DNS service.
	ViewID param.Field[string] `json:"view_id"`
}

func (r RuleSettingResolveDNSInternallyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure behavior when an upstream cert is invalid or an SSL error occurs.
type RuleSettingUntrustedCERTParam struct {
	// The action performed when an untrusted certificate is seen. The default action
	// is an error with HTTP code 526.
	Action param.Field[RuleSettingUntrustedCERTAction] `json:"action"`
}

func (r RuleSettingUntrustedCERTParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type Schedule struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri string `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon string `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat string `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun string `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu string `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone string `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue string `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed  string       `json:"wed"`
	JSON scheduleJSON `json:"-"`
}

// scheduleJSON contains the JSON metadata for the struct [Schedule]
type scheduleJSON struct {
	Fri         apijson.Field
	Mon         apijson.Field
	Sat         apijson.Field
	Sun         apijson.Field
	Thu         apijson.Field
	TimeZone    apijson.Field
	Tue         apijson.Field
	Wed         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Schedule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scheduleJSON) RawJSON() string {
	return r.raw
}

// The schedule for activating DNS policies. This does not apply to HTTP or network
// policies.
type ScheduleParam struct {
	// The time intervals when the rule will be active on Fridays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Fridays.
	Fri param.Field[string] `json:"fri"`
	// The time intervals when the rule will be active on Mondays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Mondays.
	Mon param.Field[string] `json:"mon"`
	// The time intervals when the rule will be active on Saturdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Saturdays.
	Sat param.Field[string] `json:"sat"`
	// The time intervals when the rule will be active on Sundays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Sundays.
	Sun param.Field[string] `json:"sun"`
	// The time intervals when the rule will be active on Thursdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Thursdays.
	Thu param.Field[string] `json:"thu"`
	// The time zone the rule will be evaluated against. If a
	// [valid time zone city name](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List)
	// is provided, Gateway will always use the current time at that time zone. If this
	// parameter is omitted, then Gateway will use the time zone inferred from the
	// user's source IP to evaluate the rule. If Gateway cannot determine the time zone
	// from the IP, we will fall back to the time zone of the user's connected data
	// center.
	TimeZone param.Field[string] `json:"time_zone"`
	// The time intervals when the rule will be active on Tuesdays, in increasing order
	// from 00:00-24:00. If this parameter is omitted, the rule will be deactivated on
	// Tuesdays.
	Tue param.Field[string] `json:"tue"`
	// The time intervals when the rule will be active on Wednesdays, in increasing
	// order from 00:00-24:00. If this parameter is omitted, the rule will be
	// deactivated on Wednesdays.
	Wed param.Field[string] `json:"wed"`
}

func (r ScheduleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleDeleteResponse = interface{}

type GatewayRuleNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The action to perform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[GatewayRuleNewParamsAction] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching. The API
	// automatically formats and sanitizes this expression. This returns a normalized
	// version that may differ from your input and cause Terraform state drift.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The expiration time stamp and default duration of a DNS policy. Takes precedence
	// over the policy's `schedule` configuration, if any.
	//
	// This does not apply to HTTP or network policies.
	Expiration param.Field[GatewayRuleNewParamsExpiration] `json:"expiration"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]GatewayFilter] `json:"filters"`
	// The wirefilter expression used for identity matching. The API automatically
	// formats and sanitizes this expression. This returns a normalized version that
	// may differ from your input and cause Terraform state drift.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value. Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// docs on how to manage precedence via Terraform.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[RuleSettingParam] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// The wirefilter expression used for traffic matching. The API automatically
	// formats and sanitizes this expression. This returns a normalized version that
	// may differ from your input and cause Terraform state drift.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleNewParamsAction string

const (
	GatewayRuleNewParamsActionOn           GatewayRuleNewParamsAction = "on"
	GatewayRuleNewParamsActionOff          GatewayRuleNewParamsAction = "off"
	GatewayRuleNewParamsActionAllow        GatewayRuleNewParamsAction = "allow"
	GatewayRuleNewParamsActionBlock        GatewayRuleNewParamsAction = "block"
	GatewayRuleNewParamsActionScan         GatewayRuleNewParamsAction = "scan"
	GatewayRuleNewParamsActionNoscan       GatewayRuleNewParamsAction = "noscan"
	GatewayRuleNewParamsActionSafesearch   GatewayRuleNewParamsAction = "safesearch"
	GatewayRuleNewParamsActionYtrestricted GatewayRuleNewParamsAction = "ytrestricted"
	GatewayRuleNewParamsActionIsolate      GatewayRuleNewParamsAction = "isolate"
	GatewayRuleNewParamsActionNoisolate    GatewayRuleNewParamsAction = "noisolate"
	GatewayRuleNewParamsActionOverride     GatewayRuleNewParamsAction = "override"
	GatewayRuleNewParamsActionL4Override   GatewayRuleNewParamsAction = "l4_override"
	GatewayRuleNewParamsActionEgress       GatewayRuleNewParamsAction = "egress"
	GatewayRuleNewParamsActionResolve      GatewayRuleNewParamsAction = "resolve"
	GatewayRuleNewParamsActionQuarantine   GatewayRuleNewParamsAction = "quarantine"
	GatewayRuleNewParamsActionRedirect     GatewayRuleNewParamsAction = "redirect"
)

func (r GatewayRuleNewParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleNewParamsActionOn, GatewayRuleNewParamsActionOff, GatewayRuleNewParamsActionAllow, GatewayRuleNewParamsActionBlock, GatewayRuleNewParamsActionScan, GatewayRuleNewParamsActionNoscan, GatewayRuleNewParamsActionSafesearch, GatewayRuleNewParamsActionYtrestricted, GatewayRuleNewParamsActionIsolate, GatewayRuleNewParamsActionNoisolate, GatewayRuleNewParamsActionOverride, GatewayRuleNewParamsActionL4Override, GatewayRuleNewParamsActionEgress, GatewayRuleNewParamsActionResolve, GatewayRuleNewParamsActionQuarantine, GatewayRuleNewParamsActionRedirect:
		return true
	}
	return false
}

// The expiration time stamp and default duration of a DNS policy. Takes precedence
// over the policy's `schedule` configuration, if any.
//
// This does not apply to HTTP or network policies.
type GatewayRuleNewParamsExpiration struct {
	// The time stamp at which the policy will expire and cease to be applied.
	//
	// Must adhere to RFC 3339 and include a UTC offset. Non-zero offsets are accepted
	// but will be converted to the equivalent value with offset zero (UTC+00:00) and
	// will be returned as time stamps with offset zero denoted by a trailing 'Z'.
	//
	// Policies with an expiration do not consider the timezone of clients they are
	// applied to, and expire "globally" at the point given by their `expires_at`
	// value.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// The default duration a policy will be active in minutes. Must be set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration param.Field[int64] `json:"duration"`
}

func (r GatewayRuleNewParamsExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayRuleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                           `json:"result"`
	JSON    gatewayRuleNewResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleNewResponseEnvelope]
type gatewayRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleNewResponseEnvelopeSuccess bool

const (
	GatewayRuleNewResponseEnvelopeSuccessTrue GatewayRuleNewResponseEnvelopeSuccess = true
)

func (r GatewayRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The action to perform when the associated traffic, identity, and device posture
	// expressions are either absent or evaluate to `true`.
	Action param.Field[GatewayRuleUpdateParamsAction] `json:"action,required"`
	// The name of the rule.
	Name param.Field[string] `json:"name,required"`
	// The description of the rule.
	Description param.Field[string] `json:"description"`
	// The wirefilter expression used for device posture check matching. The API
	// automatically formats and sanitizes this expression. This returns a normalized
	// version that may differ from your input and cause Terraform state drift.
	DevicePosture param.Field[string] `json:"device_posture"`
	// True if the rule is enabled.
	Enabled param.Field[bool] `json:"enabled"`
	// The expiration time stamp and default duration of a DNS policy. Takes precedence
	// over the policy's `schedule` configuration, if any.
	//
	// This does not apply to HTTP or network policies.
	Expiration param.Field[GatewayRuleUpdateParamsExpiration] `json:"expiration"`
	// The protocol or layer to evaluate the traffic, identity, and device posture
	// expressions.
	Filters param.Field[[]GatewayFilter] `json:"filters"`
	// The wirefilter expression used for identity matching. The API automatically
	// formats and sanitizes this expression. This returns a normalized version that
	// may differ from your input and cause Terraform state drift.
	Identity param.Field[string] `json:"identity"`
	// Precedence sets the order of your rules. Lower values indicate higher
	// precedence. At each processing phase, applicable rules are evaluated in
	// ascending order of this value. Refer to
	// [Order of enforcement](http://developers.cloudflare.com/learning-paths/secure-internet-traffic/understand-policies/order-of-enforcement/#manage-precedence-with-terraform)
	// docs on how to manage precedence via Terraform.
	Precedence param.Field[int64] `json:"precedence"`
	// Additional settings that modify the rule's action.
	RuleSettings param.Field[RuleSettingParam] `json:"rule_settings"`
	// The schedule for activating DNS policies. This does not apply to HTTP or network
	// policies.
	Schedule param.Field[ScheduleParam] `json:"schedule"`
	// The wirefilter expression used for traffic matching. The API automatically
	// formats and sanitizes this expression. This returns a normalized version that
	// may differ from your input and cause Terraform state drift.
	Traffic param.Field[string] `json:"traffic"`
}

func (r GatewayRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the associated traffic, identity, and device posture
// expressions are either absent or evaluate to `true`.
type GatewayRuleUpdateParamsAction string

const (
	GatewayRuleUpdateParamsActionOn           GatewayRuleUpdateParamsAction = "on"
	GatewayRuleUpdateParamsActionOff          GatewayRuleUpdateParamsAction = "off"
	GatewayRuleUpdateParamsActionAllow        GatewayRuleUpdateParamsAction = "allow"
	GatewayRuleUpdateParamsActionBlock        GatewayRuleUpdateParamsAction = "block"
	GatewayRuleUpdateParamsActionScan         GatewayRuleUpdateParamsAction = "scan"
	GatewayRuleUpdateParamsActionNoscan       GatewayRuleUpdateParamsAction = "noscan"
	GatewayRuleUpdateParamsActionSafesearch   GatewayRuleUpdateParamsAction = "safesearch"
	GatewayRuleUpdateParamsActionYtrestricted GatewayRuleUpdateParamsAction = "ytrestricted"
	GatewayRuleUpdateParamsActionIsolate      GatewayRuleUpdateParamsAction = "isolate"
	GatewayRuleUpdateParamsActionNoisolate    GatewayRuleUpdateParamsAction = "noisolate"
	GatewayRuleUpdateParamsActionOverride     GatewayRuleUpdateParamsAction = "override"
	GatewayRuleUpdateParamsActionL4Override   GatewayRuleUpdateParamsAction = "l4_override"
	GatewayRuleUpdateParamsActionEgress       GatewayRuleUpdateParamsAction = "egress"
	GatewayRuleUpdateParamsActionResolve      GatewayRuleUpdateParamsAction = "resolve"
	GatewayRuleUpdateParamsActionQuarantine   GatewayRuleUpdateParamsAction = "quarantine"
	GatewayRuleUpdateParamsActionRedirect     GatewayRuleUpdateParamsAction = "redirect"
)

func (r GatewayRuleUpdateParamsAction) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateParamsActionOn, GatewayRuleUpdateParamsActionOff, GatewayRuleUpdateParamsActionAllow, GatewayRuleUpdateParamsActionBlock, GatewayRuleUpdateParamsActionScan, GatewayRuleUpdateParamsActionNoscan, GatewayRuleUpdateParamsActionSafesearch, GatewayRuleUpdateParamsActionYtrestricted, GatewayRuleUpdateParamsActionIsolate, GatewayRuleUpdateParamsActionNoisolate, GatewayRuleUpdateParamsActionOverride, GatewayRuleUpdateParamsActionL4Override, GatewayRuleUpdateParamsActionEgress, GatewayRuleUpdateParamsActionResolve, GatewayRuleUpdateParamsActionQuarantine, GatewayRuleUpdateParamsActionRedirect:
		return true
	}
	return false
}

// The expiration time stamp and default duration of a DNS policy. Takes precedence
// over the policy's `schedule` configuration, if any.
//
// This does not apply to HTTP or network policies.
type GatewayRuleUpdateParamsExpiration struct {
	// The time stamp at which the policy will expire and cease to be applied.
	//
	// Must adhere to RFC 3339 and include a UTC offset. Non-zero offsets are accepted
	// but will be converted to the equivalent value with offset zero (UTC+00:00) and
	// will be returned as time stamps with offset zero denoted by a trailing 'Z'.
	//
	// Policies with an expiration do not consider the timezone of clients they are
	// applied to, and expire "globally" at the point given by their `expires_at`
	// value.
	ExpiresAt param.Field[time.Time] `json:"expires_at,required" format:"date-time"`
	// The default duration a policy will be active in minutes. Must be set in order to
	// use the `reset_expiration` endpoint on this rule.
	Duration param.Field[int64] `json:"duration"`
}

func (r GatewayRuleUpdateParamsExpiration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GatewayRuleUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                              `json:"result"`
	JSON    gatewayRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleUpdateResponseEnvelope]
type gatewayRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleUpdateResponseEnvelopeSuccess bool

const (
	GatewayRuleUpdateResponseEnvelopeSuccessTrue GatewayRuleUpdateResponseEnvelopeSuccess = true
)

func (r GatewayRuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRuleDeleteResponse                `json:"result"`
	JSON    gatewayRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleDeleteResponseEnvelope]
type gatewayRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleDeleteResponseEnvelopeSuccess bool

const (
	GatewayRuleDeleteResponseEnvelopeSuccessTrue GatewayRuleDeleteResponseEnvelopeSuccess = true
)

func (r GatewayRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayRuleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                           `json:"result"`
	JSON    gatewayRuleGetResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [GatewayRuleGetResponseEnvelope]
type gatewayRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleGetResponseEnvelopeSuccess bool

const (
	GatewayRuleGetResponseEnvelopeSuccessTrue GatewayRuleGetResponseEnvelopeSuccess = true
)

func (r GatewayRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type GatewayRuleResetExpirationParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type GatewayRuleResetExpirationResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success GatewayRuleResetExpirationResponseEnvelopeSuccess `json:"success,required"`
	Result  GatewayRule                                       `json:"result"`
	JSON    gatewayRuleResetExpirationResponseEnvelopeJSON    `json:"-"`
}

// gatewayRuleResetExpirationResponseEnvelopeJSON contains the JSON metadata for
// the struct [GatewayRuleResetExpirationResponseEnvelope]
type gatewayRuleResetExpirationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GatewayRuleResetExpirationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r gatewayRuleResetExpirationResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type GatewayRuleResetExpirationResponseEnvelopeSuccess bool

const (
	GatewayRuleResetExpirationResponseEnvelopeSuccessTrue GatewayRuleResetExpirationResponseEnvelopeSuccess = true
)

func (r GatewayRuleResetExpirationResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case GatewayRuleResetExpirationResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
