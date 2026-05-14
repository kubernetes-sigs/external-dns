// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package botnet_feed

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// BotnetFeedService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBotnetFeedService] method instead.
type BotnetFeedService struct {
	Options []option.RequestOption
	ASN     *ASNService
	Configs *ConfigService
}

// NewBotnetFeedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBotnetFeedService(opts ...option.RequestOption) (r *BotnetFeedService) {
	r = &BotnetFeedService{}
	r.Options = opts
	r.ASN = NewASNService(opts...)
	r.Configs = NewConfigService(opts...)
	return
}
