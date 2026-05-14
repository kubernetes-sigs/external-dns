// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package durable_objects

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// DurableObjectService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDurableObjectService] method instead.
type DurableObjectService struct {
	Options    []option.RequestOption
	Namespaces *NamespaceService
}

// NewDurableObjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDurableObjectService(opts ...option.RequestOption) (r *DurableObjectService) {
	r = &DurableObjectService{}
	r.Options = opts
	r.Namespaces = NewNamespaceService(opts...)
	return
}
