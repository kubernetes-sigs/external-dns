// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// FirewallService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFirewallService] method instead.
type FirewallService struct {
	Options   []option.RequestOption
	Lockdowns *LockdownService
	// Deprecated: The Firewall Rules API is deprecated in favour of using the Ruleset
	// Engine. See
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#firewall-rules-api-and-filters-api
	// for full details.
	Rules       *RuleService
	AccessRules *AccessRuleService
	UARules     *UARuleService
	// Deprecated: WAF managed rules API is deprecated in favour of using the Ruleset
	// Engine. See
	// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#waf-managed-rules-apis-previous-version
	// for full details.
	WAF *WAFService
}

// NewFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallService(opts ...option.RequestOption) (r *FirewallService) {
	r = &FirewallService{}
	r.Options = opts
	r.Lockdowns = NewLockdownService(opts...)
	r.Rules = NewRuleService(opts...)
	r.AccessRules = NewAccessRuleService(opts...)
	r.UARules = NewUARuleService(opts...)
	r.WAF = NewWAFService(opts...)
	return
}
