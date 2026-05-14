// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package firewall

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// WAFService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWAFService] method instead.
//
// Deprecated: WAF managed rules API is deprecated in favour of using the Ruleset
// Engine. See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#waf-managed-rules-apis-previous-version
// for full details.
type WAFService struct {
	Options   []option.RequestOption
	Overrides *WAFOverrideService
	Packages  *WAFPackageService
}

// NewWAFService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWAFService(opts ...option.RequestOption) (r *WAFService) {
	r = &WAFService{}
	r.Options = opts
	r.Overrides = NewWAFOverrideService(opts...)
	r.Packages = NewWAFPackageService(opts...)
	return
}
