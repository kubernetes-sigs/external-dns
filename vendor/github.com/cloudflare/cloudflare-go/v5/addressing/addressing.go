// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AddressingService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressingService] method instead.
type AddressingService struct {
	Options           []option.RequestOption
	RegionalHostnames *RegionalHostnameService
	Services          *ServiceService
	AddressMaps       *AddressMapService
	LOADocuments      *LOADocumentService
	Prefixes          *PrefixService
}

// NewAddressingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressingService(opts ...option.RequestOption) (r *AddressingService) {
	r = &AddressingService{}
	r.Options = opts
	r.RegionalHostnames = NewRegionalHostnameService(opts...)
	r.Services = NewServiceService(opts...)
	r.AddressMaps = NewAddressMapService(opts...)
	r.LOADocuments = NewLOADocumentService(opts...)
	r.Prefixes = NewPrefixService(opts...)
	return
}
