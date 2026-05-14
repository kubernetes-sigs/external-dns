// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// IntelService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntelService] method instead.
type IntelService struct {
	Options             []option.RequestOption
	ASN                 *ASNService
	DNS                 *DNSService
	Domains             *DomainService
	DomainHistory       *DomainHistoryService
	IPs                 *IPService
	IPLists             *IPListService
	Miscategorizations  *MiscategorizationService
	Whois               *WhoisService
	IndicatorFeeds      *IndicatorFeedService
	Sinkholes           *SinkholeService
	AttackSurfaceReport *AttackSurfaceReportService
}

// NewIntelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIntelService(opts ...option.RequestOption) (r *IntelService) {
	r = &IntelService{}
	r.Options = opts
	r.ASN = NewASNService(opts...)
	r.DNS = NewDNSService(opts...)
	r.Domains = NewDomainService(opts...)
	r.DomainHistory = NewDomainHistoryService(opts...)
	r.IPs = NewIPService(opts...)
	r.IPLists = NewIPListService(opts...)
	r.Miscategorizations = NewMiscategorizationService(opts...)
	r.Whois = NewWhoisService(opts...)
	r.IndicatorFeeds = NewIndicatorFeedService(opts...)
	r.Sinkholes = NewSinkholeService(opts...)
	r.AttackSurfaceReport = NewAttackSurfaceReportService(opts...)
	return
}
