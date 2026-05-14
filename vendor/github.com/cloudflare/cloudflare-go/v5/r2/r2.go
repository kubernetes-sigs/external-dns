// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// R2Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewR2Service] method instead.
type R2Service struct {
	Options              []option.RequestOption
	Buckets              *BucketService
	TemporaryCredentials *TemporaryCredentialService
	SuperSlurper         *SuperSlurperService
}

// NewR2Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewR2Service(opts ...option.RequestOption) (r *R2Service) {
	r = &R2Service{}
	r.Options = opts
	r.Buckets = NewBucketService(opts...)
	r.TemporaryCredentials = NewTemporaryCredentialService(opts...)
	r.SuperSlurper = NewSuperSlurperService(opts...)
	return
}
