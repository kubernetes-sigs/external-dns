// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// CallService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallService] method instead.
type CallService struct {
	Options []option.RequestOption
	SFU     *SFUService
	TURN    *TURNService
}

// NewCallService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCallService(opts ...option.RequestOption) (r *CallService) {
	r = &CallService{}
	r.Options = opts
	r.SFU = NewSFUService(opts...)
	r.TURN = NewTURNService(opts...)
	return
}
