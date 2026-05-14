// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ZeroTrustService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZeroTrustService] method instead.
type ZeroTrustService struct {
	Options              []option.RequestOption
	Devices              *DeviceService
	IdentityProviders    *IdentityProviderService
	Organizations        *OrganizationService
	Seats                *SeatService
	Access               *AccessService
	DEX                  *DEXService
	Tunnels              *TunnelService
	ConnectivitySettings *ConnectivitySettingService
	DLP                  *DLPService
	Gateway              *GatewayService
	Networks             *NetworkService
	RiskScoring          *RiskScoringService
}

// NewZeroTrustService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZeroTrustService(opts ...option.RequestOption) (r *ZeroTrustService) {
	r = &ZeroTrustService{}
	r.Options = opts
	r.Devices = NewDeviceService(opts...)
	r.IdentityProviders = NewIdentityProviderService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.Seats = NewSeatService(opts...)
	r.Access = NewAccessService(opts...)
	r.DEX = NewDEXService(opts...)
	r.Tunnels = NewTunnelService(opts...)
	r.ConnectivitySettings = NewConnectivitySettingService(opts...)
	r.DLP = NewDLPService(opts...)
	r.Gateway = NewGatewayService(opts...)
	r.Networks = NewNetworkService(opts...)
	r.RiskScoring = NewRiskScoringService(opts...)
	return
}
