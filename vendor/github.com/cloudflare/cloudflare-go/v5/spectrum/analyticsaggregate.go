// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AnalyticsAggregateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnalyticsAggregateService] method instead.
type AnalyticsAggregateService struct {
	Options  []option.RequestOption
	Currents *AnalyticsAggregateCurrentService
}

// NewAnalyticsAggregateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsAggregateService(opts ...option.RequestOption) (r *AnalyticsAggregateService) {
	r = &AnalyticsAggregateService{}
	r.Options = opts
	r.Currents = NewAnalyticsAggregateCurrentService(opts...)
	return
}
