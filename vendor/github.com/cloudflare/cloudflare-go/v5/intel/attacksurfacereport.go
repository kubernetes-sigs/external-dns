// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AttackSurfaceReportService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackSurfaceReportService] method instead.
type AttackSurfaceReportService struct {
	Options    []option.RequestOption
	IssueTypes *AttackSurfaceReportIssueTypeService
	Issues     *AttackSurfaceReportIssueService
}

// NewAttackSurfaceReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAttackSurfaceReportService(opts ...option.RequestOption) (r *AttackSurfaceReportService) {
	r = &AttackSurfaceReportService{}
	r.Options = opts
	r.IssueTypes = NewAttackSurfaceReportIssueTypeService(opts...)
	r.Issues = NewAttackSurfaceReportIssueService(opts...)
	return
}
