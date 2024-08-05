package cloudflare

import (
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TeamsRuleSettings struct {
	// list of ipv4 or ipv6 ips to override with, when action is set to dns override
	OverrideIPs []string `json:"override_ips"`

	// show this string at block page caused by this rule
	BlockReason string `json:"block_reason"`

	// host name to override with when action is set to dns override. Can not be used with OverrideIPs
	OverrideHost string `json:"override_host"`

	// settings for browser isolation actions
	BISOAdminControls *TeamsBISOAdminControlSettings `json:"biso_admin_controls"`

	// settings for l4(network) level overrides
	L4Override *TeamsL4OverrideSettings `json:"l4override"`

	// settings for adding headers to http requests
	AddHeaders http.Header `json:"add_headers"`

	// settings for session check in allow action
	CheckSession *TeamsCheckSessionSettings `json:"check_session"`

	// Enable block page on rules with action block
	BlockPageEnabled bool `json:"block_page_enabled"`

	// whether to disable dnssec validation for allow action
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`
}

// TeamsL4OverrideSettings used in l4 filter type rule with action set to override.
type TeamsL4OverrideSettings struct {
	IP   string `json:"ip,omitempty"`
	Port int    `json:"port,omitempty"`
}

type TeamsBISOAdminControlSettings struct {
	DisablePrinting  bool `json:"dp"`
	DisableCopyPaste bool `json:"dcp"`
	DisableDownload  bool `json:"dd"`
	DisableUpload    bool `json:"du"`
	DisableKeyboard  bool `json:"dk"`
}

type TeamsCheckSessionSettings struct {
	Enforce  bool     `json:"enforce"`
	Duration Duration `json:"duration"`
}

type TeamsFilterType string

type TeamsGatewayAction string

const (
	HttpFilter TeamsFilterType = "http"
	DnsFilter  TeamsFilterType = "dns"
	L4Filter   TeamsFilterType = "l4"
)

const (
	Allow        TeamsGatewayAction = "allow"
	Block        TeamsGatewayAction = "block"
	SafeSearch   TeamsGatewayAction = "safesearch"
	YTRestricted TeamsGatewayAction = "ytrestricted"
	On           TeamsGatewayAction = "on"
	Off          TeamsGatewayAction = "off"
	Scan         TeamsGatewayAction = "scan"
	NoScan       TeamsGatewayAction = "noscan"
	Isolate      TeamsGatewayAction = "isolate"
	NoIsolate    TeamsGatewayAction = "noisolate"
	Override     TeamsGatewayAction = "override"
	L4Override   TeamsGatewayAction = "l4_override"
)

func TeamsRulesActionValues() []string {
	return []string{
		string(Allow),
		string(Block),
		string(SafeSearch),
		string(YTRestricted),
		string(On),
		string(Off),
		string(Scan),
		string(NoScan),
		string(Isolate),
		string(NoIsolate),
		string(Override),
		string(L4Override),
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

type TeamsRuleSettings struct {
	// list of ipv4 or ipv6 ips to override with, when action is set to dns override
	OverrideIPs []string `json:"override_ips"`

	// show this string at block page caused by this rule
	BlockReason string `json:"block_reason"`

	// host name to override with when action is set to dns override. Can not be used with OverrideIPs
	OverrideHost string `json:"override_host"`

	// settings for browser isolation actions
	BISOAdminControls *TeamsBISOAdminControlSettings `json:"biso_admin_controls"`

	// settings for l4(network) level overrides
	L4Override *TeamsL4OverrideSettings `json:"l4override"`

	// settings for adding headers to http requests
	AddHeaders http.Header `json:"add_headers"`

	// settings for session check in allow action
	CheckSession *TeamsCheckSessionSettings `json:"check_session"`

	// Enable block page on rules with action block
	BlockPageEnabled bool `json:"block_page_enabled"`

	// whether to disable dnssec validation for allow action
	InsecureDisableDNSSECValidation bool `json:"insecure_disable_dnssec_validation"`

	// settings for rules with egress action
	EgressSettings *EgressSettings `json:"egress"`

	// DLP payload logging configuration
	PayloadLog *TeamsDlpPayloadLogSettings `json:"payload_log"`

	//AuditSsh Settings
	AuditSSH *AuditSSHRuleSettings `json:"audit_ssh"`

	// Turns on ip category based filter on dns if the rule contains dns category checks
	IPCategories bool `json:"ip_categories"`

	// Allow parent MSP accounts to enable bypass their children's rules. Do not set them for non MSP accounts.
	AllowChildBypass *bool `json:"allow_child_bypass,omitempty"`

	// Allow child MSP accounts to bypass their parent's rules. Do not set them for non MSP accounts.
	BypassParentRule *bool `json:"bypass_parent_rule,omitempty"`

	// Action taken when an untrusted origin certificate error occurs in a http allow rule
	UntrustedCertSettings *UntrustedCertSettings `json:"untrusted_cert"`

	// Specifies that a resolver policy should use Cloudflare's DNS Resolver.
	ResolveDnsThroughCloudflare *bool `json:"resolve_dns_through_cloudflare,omitempty"`

	// Resolver policy settings.
	DnsResolverSettings *TeamsDnsResolverSettings `json:"dns_resolvers,omitempty"`

	NotificationSettings *TeamsNotificationSettings `json:"notification_settings"`
}

type TeamsGatewayUntrustedCertAction string

const (
	UntrustedCertPassthrough TeamsGatewayUntrustedCertAction = "pass_through"
	UntrustedCertBlock       TeamsGatewayUntrustedCertAction = "block"
	UntrustedCertError       TeamsGatewayUntrustedCertAction = "error"
)

type UntrustedCertSettings struct {
	Action TeamsGatewayUntrustedCertAction `json:"action"`
}

type TeamsNotificationSettings struct {
	Enabled    *bool  `json:"enabled,omitempty"`
	Message    string `json:"msg"`
	SupportURL string `json:"support_url"`
}

type AuditSSHRuleSettings struct {
	CommandLogging bool `json:"command_logging"`
}

type EgressSettings struct {
	Ipv6Range    string `json:"ipv6"`
	Ipv4         string `json:"ipv4"`
	Ipv4Fallback string `json:"ipv4_fallback"`
}

// TeamsL4OverrideSettings used in l4 filter type rule with action set to override.
type TeamsL4OverrideSettings struct {
	IP   string `json:"ip,omitempty"`
	Port int    `json:"port,omitempty"`
}

type TeamsBISOAdminControlSettings struct {
	DisablePrinting             bool `json:"dp"`
	DisableCopyPaste            bool `json:"dcp"`
	DisableDownload             bool `json:"dd"`
	DisableUpload               bool `json:"du"`
	DisableKeyboard             bool `json:"dk"`
	DisableClipboardRedirection bool `json:"dcr"`
}

type TeamsCheckSessionSettings struct {
	Enforce  bool     `json:"enforce"`
	Duration Duration `json:"duration"`
}

type (
	TeamsDnsResolverSettings struct {
		V4Resolvers []TeamsDnsResolverAddressV4 `json:"ipv4,omitempty"`
		V6Resolvers []TeamsDnsResolverAddressV6 `json:"ipv6,omitempty"`
	}

	TeamsDnsResolverAddressV4 struct {
		TeamsDnsResolverAddress
	}

	TeamsDnsResolverAddressV6 struct {
		TeamsDnsResolverAddress
	}

	TeamsDnsResolverAddress struct {
		IP                         string `json:"ip"`
		Port                       *int   `json:"port,omitempty"`
		VnetID                     string `json:"vnet_id,omitempty"`
		RouteThroughPrivateNetwork *bool  `json:"route_through_private_network,omitempty"`
	}
)

type TeamsDlpPayloadLogSettings struct {
	Enabled bool `json:"enabled"`
}

type TeamsFilterType string

type TeamsGatewayAction string

const (
	HttpFilter        TeamsFilterType = "http"
	DnsFilter         TeamsFilterType = "dns"
	L4Filter          TeamsFilterType = "l4"
	EgressFilter      TeamsFilterType = "egress"
	DnsResolverFilter TeamsFilterType = "dns_resolver"
)

const (
	Allow        TeamsGatewayAction = "allow"        // dns|http|l4
	Block        TeamsGatewayAction = "block"        // dns|http|l4
	SafeSearch   TeamsGatewayAction = "safesearch"   // dns
	YTRestricted TeamsGatewayAction = "ytrestricted" // dns
	On           TeamsGatewayAction = "on"           // http
	Off          TeamsGatewayAction = "off"          // http
	Scan         TeamsGatewayAction = "scan"         // http
	NoScan       TeamsGatewayAction = "noscan"       // http
	Isolate      TeamsGatewayAction = "isolate"      // http
	NoIsolate    TeamsGatewayAction = "noisolate"    // http
	Override     TeamsGatewayAction = "override"     // http
	L4Override   TeamsGatewayAction = "l4_override"  // l4
	Egress       TeamsGatewayAction = "egress"       // egress
	AuditSSH     TeamsGatewayAction = "audit_ssh"    // l4
	Resolve      TeamsGatewayAction = "resolve"      // resolve
)

func TeamsRulesActionValues() []string {
	return []string{
		string(Allow),
		string(Block),
		string(SafeSearch),
		string(YTRestricted),
		string(On),
		string(Off),
		string(Scan),
		string(NoScan),
		string(Isolate),
		string(NoIsolate),
		string(Override),
		string(L4Override),
		string(Egress),
		string(AuditSSH),
		string(Resolve),
	}
}

func TeamsRulesUntrustedCertActionValues() []string {
	return []string{
		string(UntrustedCertPassthrough),
		string(UntrustedCertBlock),
		string(UntrustedCertError),
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
}

// TeamsRule represents an Teams wirefilter rule.
type TeamsRule struct {
	ID            string             `json:"id,omitempty"`
	CreatedAt     *time.Time         `json:"created_at,omitempty"`
	UpdatedAt     *time.Time         `json:"updated_at,omitempty"`
	DeletedAt     *time.Time         `json:"deleted_at,omitempty"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Precedence    uint64             `json:"precedence"`
	Enabled       bool               `json:"enabled"`
	Action        TeamsGatewayAction `json:"action"`
	Filters       []TeamsFilterType  `json:"filters"`
	Traffic       string             `json:"traffic"`
	Identity      string             `json:"identity"`
	DevicePosture string             `json:"device_posture"`
	Version       uint64             `json:"version"`
	RuleSettings  TeamsRuleSettings  `json:"rule_settings,omitempty"`
}

// TeamsRuleResponse is the API response, containing a single rule.
type TeamsRuleResponse struct {
	Response
	Result TeamsRule `json:"result"`
}

// TeamsRuleResponse is the API response, containing an array of rules.
type TeamsRulesResponse struct {
	Response
	Result []TeamsRule `json:"result"`
}

// TeamsRulePatchRequest is used to patch an existing rule.
type TeamsRulePatchRequest struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Precedence   uint64             `json:"precedence"`
	Enabled      bool               `json:"enabled"`
	Action       TeamsGatewayAction `json:"action"`
	RuleSettings TeamsRuleSettings  `json:"rule_settings,omitempty"`
}

// TeamsRules returns all rules within an account.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsRules(ctx context.Context, accountID string) ([]TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []TeamsRule{}, err
	}

	var teamsRulesResponse TeamsRulesResponse
	err = json.Unmarshal(res, &teamsRulesResponse)
	if err != nil {
		return []TeamsRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsRulesResponse.Result, nil
}

// TeamsRule returns the rule with rule ID in the URL.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsRule(ctx context.Context, accountID string, ruleId string) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleId)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsCreateRule creates a rule with wirefilter expression.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsCreateRule(ctx context.Context, accountID string, rule TeamsRule) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules", accountID)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, rule)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsUpdateRule updates a rule with wirefilter expression.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsUpdateRule(ctx context.Context, accountID string, ruleId string, rule TeamsRule) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleId)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, rule)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsPatchRule patches a rule associated values.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsPatchRule(ctx context.Context, accountID string, ruleId string, rule TeamsRulePatchRequest) (TeamsRule, error) {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleId)

	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, rule)
	if err != nil {
		return TeamsRule{}, err
	}

	var teamsRuleResponse TeamsRuleResponse
	err = json.Unmarshal(res, &teamsRuleResponse)
	if err != nil {
		return TeamsRule{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return teamsRuleResponse.Result, nil
}

// TeamsDeleteRule deletes a rule.
//
// API reference: https://api.cloudflare.com/#teams-rules-properties
func (api *API) TeamsDeleteRule(ctx context.Context, accountID string, ruleId string) error {
	uri := fmt.Sprintf("/accounts/%s/gateway/rules/%s", accountID, ruleId)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}
