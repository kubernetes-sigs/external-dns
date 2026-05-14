// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// IdentityProviderSCIMService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIdentityProviderSCIMService] method instead.
type IdentityProviderSCIMService struct {
	Options []option.RequestOption
	Groups  *IdentityProviderSCIMGroupService
	Users   *IdentityProviderSCIMUserService
}

// NewIdentityProviderSCIMService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIdentityProviderSCIMService(opts ...option.RequestOption) (r *IdentityProviderSCIMService) {
	r = &IdentityProviderSCIMService{}
	r.Options = opts
	r.Groups = NewIdentityProviderSCIMGroupService(opts...)
	r.Users = NewIdentityProviderSCIMUserService(opts...)
	return
}
