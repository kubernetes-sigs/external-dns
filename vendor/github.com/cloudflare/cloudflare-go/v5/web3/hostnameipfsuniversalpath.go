// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// HostnameIPFSUniversalPathService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameIPFSUniversalPathService] method instead.
type HostnameIPFSUniversalPathService struct {
	Options      []option.RequestOption
	ContentLists *HostnameIPFSUniversalPathContentListService
}

// NewHostnameIPFSUniversalPathService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHostnameIPFSUniversalPathService(opts ...option.RequestOption) (r *HostnameIPFSUniversalPathService) {
	r = &HostnameIPFSUniversalPathService{}
	r.Options = opts
	r.ContentLists = NewHostnameIPFSUniversalPathContentListService(opts...)
	return
}
