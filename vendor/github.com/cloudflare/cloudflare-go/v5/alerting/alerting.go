// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AlertingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAlertingService] method instead.
type AlertingService struct {
	Options         []option.RequestOption
	AvailableAlerts *AvailableAlertService
	Destinations    *DestinationService
	History         *HistoryService
	Policies        *PolicyService
}

// NewAlertingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAlertingService(opts ...option.RequestOption) (r *AlertingService) {
	r = &AlertingService{}
	r.Options = opts
	r.AvailableAlerts = NewAvailableAlertService(opts...)
	r.Destinations = NewDestinationService(opts...)
	r.History = NewHistoryService(opts...)
	r.Policies = NewPolicyService(opts...)
	return
}
