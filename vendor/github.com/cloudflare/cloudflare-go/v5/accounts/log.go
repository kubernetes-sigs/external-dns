// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// LogService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogService] method instead.
type LogService struct {
	Options []option.RequestOption
	Audit   *LogAuditService
}

// NewLogService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLogService(opts ...option.RequestOption) (r *LogService) {
	r = &LogService{}
	r.Options = opts
	r.Audit = NewLogAuditService(opts...)
	return
}
