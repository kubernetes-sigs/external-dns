// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DLPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPService] method instead.
type DLPService struct {
	Options     []option.RequestOption
	Datasets    *DLPDatasetService
	Patterns    *DLPPatternService
	PayloadLogs *DLPPayloadLogService
	Email       *DLPEmailService
	Profiles    *DLPProfileService
	Limits      *DLPLimitService
	Entries     *DLPEntryService
}

// NewDLPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDLPService(opts ...option.RequestOption) (r *DLPService) {
	r = &DLPService{}
	r.Options = opts
	r.Datasets = NewDLPDatasetService(opts...)
	r.Patterns = NewDLPPatternService(opts...)
	r.PayloadLogs = NewDLPPayloadLogService(opts...)
	r.Email = NewDLPEmailService(opts...)
	r.Profiles = NewDLPProfileService(opts...)
	r.Limits = NewDLPLimitService(opts...)
	r.Entries = NewDLPEntryService(opts...)
	return
}
