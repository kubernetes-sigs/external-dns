// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// LeakedCredentialService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLeakedCredentialService] method instead.
type LeakedCredentialService struct {
	Options          []option.RequestOption
	Summary          *LeakedCredentialSummaryService
	TimeseriesGroups *LeakedCredentialTimeseriesGroupService
}

// NewLeakedCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLeakedCredentialService(opts ...option.RequestOption) (r *LeakedCredentialService) {
	r = &LeakedCredentialService{}
	r.Options = opts
	r.Summary = NewLeakedCredentialSummaryService(opts...)
	r.TimeseriesGroups = NewLeakedCredentialTimeseriesGroupService(opts...)
	return
}
