// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package argo

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ArgoService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewArgoService] method instead.
type ArgoService struct {
	Options       []option.RequestOption
	SmartRouting  *SmartRoutingService
	TieredCaching *TieredCachingService
}

// NewArgoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewArgoService(opts ...option.RequestOption) (r *ArgoService) {
	r = &ArgoService{}
	r.Options = opts
	r.SmartRouting = NewSmartRoutingService(opts...)
	r.TieredCaching = NewTieredCachingService(opts...)
	return
}
