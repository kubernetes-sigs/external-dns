// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// NetworkService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkService] method instead.
type NetworkService struct {
	Options         []option.RequestOption
	Routes          *NetworkRouteService
	VirtualNetworks *NetworkVirtualNetworkService
	Subnets         *NetworkSubnetService
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r *NetworkService) {
	r = &NetworkService{}
	r.Options = opts
	r.Routes = NewNetworkRouteService(opts...)
	r.VirtualNetworks = NewNetworkVirtualNetworkService(opts...)
	r.Subnets = NewNetworkSubnetService(opts...)
	return
}
