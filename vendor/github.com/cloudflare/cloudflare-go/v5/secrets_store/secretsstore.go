// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package secrets_store

import (
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// SecretsStoreService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecretsStoreService] method instead.
type SecretsStoreService struct {
	Options []option.RequestOption
	Stores  *StoreService
	Quota   *QuotaService
}

// NewSecretsStoreService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSecretsStoreService(opts ...option.RequestOption) (r *SecretsStoreService) {
	r = &SecretsStoreService{}
	r.Options = opts
	r.Stores = NewStoreService(opts...)
	r.Quota = NewQuotaService(opts...)
	return
}
