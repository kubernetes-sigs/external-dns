// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package security_center

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SecurityCenterService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecurityCenterService] method instead.
type SecurityCenterService struct {
	Options  []option.RequestOption
	Insights *InsightService
}

// NewSecurityCenterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecurityCenterService(opts ...option.RequestOption) (r *SecurityCenterService) {
	r = &SecurityCenterService{}
	r.Options = opts
	r.Insights = NewInsightService(opts...)
	return
}
