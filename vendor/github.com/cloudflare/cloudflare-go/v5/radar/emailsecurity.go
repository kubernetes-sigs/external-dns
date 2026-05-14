// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// EmailSecurityService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailSecurityService] method instead.
type EmailSecurityService struct {
	Options          []option.RequestOption
	Top              *EmailSecurityTopService
	Summary          *EmailSecuritySummaryService
	TimeseriesGroups *EmailSecurityTimeseriesGroupService
}

// NewEmailSecurityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailSecurityService(opts ...option.RequestOption) (r *EmailSecurityService) {
	r = &EmailSecurityService{}
	r.Options = opts
	r.Top = NewEmailSecurityTopService(opts...)
	r.Summary = NewEmailSecuritySummaryService(opts...)
	r.TimeseriesGroups = NewEmailSecurityTimeseriesGroupService(opts...)
	return
}
