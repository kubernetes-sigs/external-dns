// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ScanService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScanService] method instead.
type ScanService struct {
	Options []option.RequestOption
	Results *ScanResultService
	Config  *ScanConfigService
}

// NewScanService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewScanService(opts ...option.RequestOption) (r *ScanService) {
	r = &ScanService{}
	r.Options = opts
	r.Results = NewScanResultService(opts...)
	r.Config = NewScanConfigService(opts...)
	return
}
