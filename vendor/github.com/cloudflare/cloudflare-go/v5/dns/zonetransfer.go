// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ZoneTransferService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferService] method instead.
type ZoneTransferService struct {
	Options   []option.RequestOption
	ForceAXFR *ZoneTransferForceAXFRService
	Incoming  *ZoneTransferIncomingService
	Outgoing  *ZoneTransferOutgoingService
	ACLs      *ZoneTransferACLService
	Peers     *ZoneTransferPeerService
	TSIGs     *ZoneTransferTSIGService
}

// NewZoneTransferService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneTransferService(opts ...option.RequestOption) (r *ZoneTransferService) {
	r = &ZoneTransferService{}
	r.Options = opts
	r.ForceAXFR = NewZoneTransferForceAXFRService(opts...)
	r.Incoming = NewZoneTransferIncomingService(opts...)
	r.Outgoing = NewZoneTransferOutgoingService(opts...)
	r.ACLs = NewZoneTransferACLService(opts...)
	r.Peers = NewZoneTransferPeerService(opts...)
	r.TSIGs = NewZoneTransferTSIGService(opts...)
	return
}
