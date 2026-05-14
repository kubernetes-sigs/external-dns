// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// VerifiedBotService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerifiedBotService] method instead.
type VerifiedBotService struct {
	Options []option.RequestOption
	Top     *VerifiedBotTopService
}

// NewVerifiedBotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVerifiedBotService(opts ...option.RequestOption) (r *VerifiedBotService) {
	r = &VerifiedBotService{}
	r.Options = opts
	r.Top = NewVerifiedBotTopService(opts...)
	return
}
