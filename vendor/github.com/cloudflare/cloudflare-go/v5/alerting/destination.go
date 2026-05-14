// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DestinationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDestinationService] method instead.
type DestinationService struct {
	Options   []option.RequestOption
	Eligible  *DestinationEligibleService
	Pagerduty *DestinationPagerdutyService
	Webhooks  *DestinationWebhookService
}

// NewDestinationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDestinationService(opts ...option.RequestOption) (r *DestinationService) {
	r = &DestinationService{}
	r.Options = opts
	r.Eligible = NewDestinationEligibleService(opts...)
	r.Pagerduty = NewDestinationPagerdutyService(opts...)
	r.Webhooks = NewDestinationWebhookService(opts...)
	return
}
