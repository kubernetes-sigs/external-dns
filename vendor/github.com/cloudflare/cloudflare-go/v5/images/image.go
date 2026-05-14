// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package images

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ImageService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageService] method instead.
type ImageService struct {
	Options []option.RequestOption
	V1      *V1Service
	V2      *V2Service
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r *ImageService) {
	r = &ImageService{}
	r.Options = opts
	r.V1 = NewV1Service(opts...)
	r.V2 = NewV2Service(opts...)
	return
}
