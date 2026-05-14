// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package web3

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// Web3Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWeb3Service] method instead.
type Web3Service struct {
	Options   []option.RequestOption
	Hostnames *HostnameService
}

// NewWeb3Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWeb3Service(opts ...option.RequestOption) (r *Web3Service) {
	r = &Web3Service{}
	r.Options = opts
	r.Hostnames = NewHostnameService(opts...)
	return
}
