// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SuperSlurperService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperSlurperService] method instead.
type SuperSlurperService struct {
	Options              []option.RequestOption
	Jobs                 *SuperSlurperJobService
	ConnectivityPrecheck *SuperSlurperConnectivityPrecheckService
}

// NewSuperSlurperService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSuperSlurperService(opts ...option.RequestOption) (r *SuperSlurperService) {
	r = &SuperSlurperService{}
	r.Options = opts
	r.Jobs = NewSuperSlurperJobService(opts...)
	r.ConnectivityPrecheck = NewSuperSlurperConnectivityPrecheckService(opts...)
	return
}
