// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package turnstile

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// TurnstileService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTurnstileService] method instead.
type TurnstileService struct {
	Options []option.RequestOption
	Widgets *WidgetService
}

// NewTurnstileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTurnstileService(opts ...option.RequestOption) (r *TurnstileService) {
	r = &TurnstileService{}
	r.Options = opts
	r.Widgets = NewWidgetService(opts...)
	return
}
