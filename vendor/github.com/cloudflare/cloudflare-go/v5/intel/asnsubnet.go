// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// ASNSubnetService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewASNSubnetService] method instead.
type ASNSubnetService struct {
	Options []option.RequestOption
}

// NewASNSubnetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewASNSubnetService(opts ...option.RequestOption) (r *ASNSubnetService) {
	r = &ASNSubnetService{}
	r.Options = opts
	return
}

// Get ASN Subnets.
func (r *ASNSubnetService) Get(ctx context.Context, asn shared.ASNParam, query ASNSubnetGetParams, opts ...option.RequestOption) (res *ASNSubnetGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/intel/asn/%v/subnets", query.AccountID, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ASNSubnetGetResponse struct {
	ASN shared.ASN `json:"asn"`
	// Total results returned based on your search parameters.
	Count        float64 `json:"count"`
	IPCountTotal int64   `json:"ip_count_total"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64                  `json:"per_page"`
	Subnets []string                 `json:"subnets"`
	JSON    asnSubnetGetResponseJSON `json:"-"`
}

// asnSubnetGetResponseJSON contains the JSON metadata for the struct
// [ASNSubnetGetResponse]
type asnSubnetGetResponseJSON struct {
	ASN          apijson.Field
	Count        apijson.Field
	IPCountTotal apijson.Field
	Page         apijson.Field
	PerPage      apijson.Field
	Subnets      apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ASNSubnetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnSubnetGetResponseJSON) RawJSON() string {
	return r.raw
}

type ASNSubnetGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}
