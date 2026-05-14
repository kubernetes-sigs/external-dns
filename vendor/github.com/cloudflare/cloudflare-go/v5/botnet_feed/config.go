// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package botnet_feed

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ConfigService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigService] method instead.
type ConfigService struct {
	Options []option.RequestOption
	ASN     *ConfigASNService
}

// NewConfigService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConfigService(opts ...option.RequestOption) (r *ConfigService) {
	r = &ConfigService{}
	r.Options = opts
	r.ASN = NewConfigASNService(opts...)
	return
}
