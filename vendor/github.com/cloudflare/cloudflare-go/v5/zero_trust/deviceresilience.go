// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DeviceResilienceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceResilienceService] method instead.
type DeviceResilienceService struct {
	Options            []option.RequestOption
	GlobalWARPOverride *DeviceResilienceGlobalWARPOverrideService
}

// NewDeviceResilienceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceResilienceService(opts ...option.RequestOption) (r *DeviceResilienceService) {
	r = &DeviceResilienceService{}
	r.Options = opts
	r.GlobalWARPOverride = NewDeviceResilienceGlobalWARPOverrideService(opts...)
	return
}
