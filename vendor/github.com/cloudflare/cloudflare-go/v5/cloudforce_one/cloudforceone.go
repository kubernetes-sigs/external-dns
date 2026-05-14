// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// CloudforceOneService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCloudforceOneService] method instead.
type CloudforceOneService struct {
	Options       []option.RequestOption
	Scans         *ScanService
	BinaryStorage *BinaryStorageService
	Requests      *RequestService
	ThreatEvents  *ThreatEventService
}

// NewCloudforceOneService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCloudforceOneService(opts ...option.RequestOption) (r *CloudforceOneService) {
	r = &CloudforceOneService{}
	r.Options = opts
	r.Scans = NewScanService(opts...)
	r.BinaryStorage = NewBinaryStorageService(opts...)
	r.Requests = NewRequestService(opts...)
	r.ThreatEvents = NewThreatEventService(opts...)
	return
}
