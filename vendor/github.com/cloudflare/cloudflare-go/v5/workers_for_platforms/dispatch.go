// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DispatchService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDispatchService] method instead.
type DispatchService struct {
	Options    []option.RequestOption
	Namespaces *DispatchNamespaceService
}

// NewDispatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDispatchService(opts ...option.RequestOption) (r *DispatchService) {
	r = &DispatchService{}
	r.Options = opts
	r.Namespaces = NewDispatchNamespaceService(opts...)
	return
}
