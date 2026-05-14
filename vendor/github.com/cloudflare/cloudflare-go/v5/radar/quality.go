// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// QualityService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualityService] method instead.
type QualityService struct {
	Options []option.RequestOption
	IQI     *QualityIQIService
	Speed   *QualitySpeedService
}

// NewQualityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQualityService(opts ...option.RequestOption) (r *QualityService) {
	r = &QualityService{}
	r.Options = opts
	r.IQI = NewQualityIQIService(opts...)
	r.Speed = NewQualitySpeedService(opts...)
	return
}
