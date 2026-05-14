// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// AttackService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackService] method instead.
type AttackService struct {
	Options []option.RequestOption
	Layer3  *AttackLayer3Service
	Layer7  *AttackLayer7Service
}

// NewAttackService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAttackService(opts ...option.RequestOption) (r *AttackService) {
	r = &AttackService{}
	r.Options = opts
	r.Layer3 = NewAttackLayer3Service(opts...)
	r.Layer7 = NewAttackLayer7Service(opts...)
	return
}
