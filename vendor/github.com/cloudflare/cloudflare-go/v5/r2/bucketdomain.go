// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// BucketDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBucketDomainService] method instead.
type BucketDomainService struct {
	Options []option.RequestOption
	Custom  *BucketDomainCustomService
	Managed *BucketDomainManagedService
}

// NewBucketDomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBucketDomainService(opts ...option.RequestOption) (r *BucketDomainService) {
	r = &BucketDomainService{}
	r.Options = opts
	r.Custom = NewBucketDomainCustomService(opts...)
	r.Managed = NewBucketDomainManagedService(opts...)
	return
}
