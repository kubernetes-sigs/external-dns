// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// WorkersForPlatformService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkersForPlatformService] method instead.
type WorkersForPlatformService struct {
	Options  []option.RequestOption
	Dispatch *DispatchService
}

// NewWorkersForPlatformService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkersForPlatformService(opts ...option.RequestOption) (r *WorkersForPlatformService) {
	r = &WorkersForPlatformService{}
	r.Options = opts
	r.Dispatch = NewDispatchService(opts...)
	return
}
