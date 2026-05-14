// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ssl

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// UniversalService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUniversalService] method instead.
type UniversalService struct {
	Options  []option.RequestOption
	Settings *UniversalSettingService
}

// NewUniversalService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUniversalService(opts ...option.RequestOption) (r *UniversalService) {
	r = &UniversalService{}
	r.Options = opts
	r.Settings = NewUniversalSettingService(opts...)
	return
}
