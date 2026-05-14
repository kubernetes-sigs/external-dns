// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_hostnames

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// CertificatePackService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCertificatePackService] method instead.
type CertificatePackService struct {
	Options      []option.RequestOption
	Certificates *CertificatePackCertificateService
}

// NewCertificatePackService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCertificatePackService(opts ...option.RequestOption) (r *CertificatePackService) {
	r = &CertificatePackService{}
	r.Options = opts
	r.Certificates = NewCertificatePackCertificateService(opts...)
	return
}
