// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// EmailRoutingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailRoutingService] method instead.
type EmailRoutingService struct {
	Options          []option.RequestOption
	Summary          *EmailRoutingSummaryService
	TimeseriesGroups *EmailRoutingTimeseriesGroupService
}

// NewEmailRoutingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailRoutingService(opts ...option.RequestOption) (r *EmailRoutingService) {
	r = &EmailRoutingService{}
	r.Options = opts
	r.Summary = NewEmailRoutingSummaryService(opts...)
	r.TimeseriesGroups = NewEmailRoutingTimeseriesGroupService(opts...)
	return
}
