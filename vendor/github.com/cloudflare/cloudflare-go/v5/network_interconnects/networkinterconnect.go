// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_interconnects

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// NetworkInterconnectService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkInterconnectService] method instead.
type NetworkInterconnectService struct {
	Options       []option.RequestOption
	CNIs          *CNIService
	Interconnects *InterconnectService
	Settings      *SettingService
	Slots         *SlotService
}

// NewNetworkInterconnectService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNetworkInterconnectService(opts ...option.RequestOption) (r *NetworkInterconnectService) {
	r = &NetworkInterconnectService{}
	r.Options = opts
	r.CNIs = NewCNIService(opts...)
	r.Interconnects = NewInterconnectService(opts...)
	r.Settings = NewSettingService(opts...)
	r.Slots = NewSlotService(opts...)
	return
}
