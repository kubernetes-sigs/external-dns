// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package acm

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ACMService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewACMService] method instead.
type ACMService struct {
	Options  []option.RequestOption
	TotalTLS *TotalTLSService
}

// NewACMService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewACMService(opts ...option.RequestOption) (r *ACMService) {
	r = &ACMService{}
	r.Options = opts
	r.TotalTLS = NewTotalTLSService(opts...)
	return
}
