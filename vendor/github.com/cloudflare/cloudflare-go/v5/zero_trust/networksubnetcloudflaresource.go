// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// NetworkSubnetCloudflareSourceService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkSubnetCloudflareSourceService] method instead.
type NetworkSubnetCloudflareSourceService struct {
	Options []option.RequestOption
}

// NewNetworkSubnetCloudflareSourceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNetworkSubnetCloudflareSourceService(opts ...option.RequestOption) (r *NetworkSubnetCloudflareSourceService) {
	r = &NetworkSubnetCloudflareSourceService{}
	r.Options = opts
	return
}

// Updates the Cloudflare Source subnet of the given address family
func (r *NetworkSubnetCloudflareSourceService) Update(ctx context.Context, addressFamily NetworkSubnetCloudflareSourceUpdateParamsAddressFamily, params NetworkSubnetCloudflareSourceUpdateParams, opts ...option.RequestOption) (res *NetworkSubnetCloudflareSourceUpdateResponse, err error) {
	var env NetworkSubnetCloudflareSourceUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/zerotrust/subnets/cloudflare_source/%v", params.AccountID, addressFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type NetworkSubnetCloudflareSourceUpdateResponse struct {
	// The UUID of the subnet.
	ID string `json:"id" format:"uuid"`
	// An optional description of the subnet.
	Comment string `json:"comment"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// If `true`, this is the default subnet for the account. There can only be one
	// default subnet per account.
	IsDefaultNetwork bool `json:"is_default_network"`
	// A user-friendly name for the subnet.
	Name string `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network string `json:"network"`
	// The type of subnet.
	SubnetType NetworkSubnetCloudflareSourceUpdateResponseSubnetType `json:"subnet_type"`
	JSON       networkSubnetCloudflareSourceUpdateResponseJSON       `json:"-"`
}

// networkSubnetCloudflareSourceUpdateResponseJSON contains the JSON metadata for
// the struct [NetworkSubnetCloudflareSourceUpdateResponse]
type networkSubnetCloudflareSourceUpdateResponseJSON struct {
	ID               apijson.Field
	Comment          apijson.Field
	CreatedAt        apijson.Field
	DeletedAt        apijson.Field
	IsDefaultNetwork apijson.Field
	Name             apijson.Field
	Network          apijson.Field
	SubnetType       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *NetworkSubnetCloudflareSourceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetCloudflareSourceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The type of subnet.
type NetworkSubnetCloudflareSourceUpdateResponseSubnetType string

const (
	NetworkSubnetCloudflareSourceUpdateResponseSubnetTypeCloudflareSource NetworkSubnetCloudflareSourceUpdateResponseSubnetType = "cloudflare_source"
)

func (r NetworkSubnetCloudflareSourceUpdateResponseSubnetType) IsKnown() bool {
	switch r {
	case NetworkSubnetCloudflareSourceUpdateResponseSubnetTypeCloudflareSource:
		return true
	}
	return false
}

type NetworkSubnetCloudflareSourceUpdateParams struct {
	// Cloudflare account ID
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description of the subnet.
	Comment param.Field[string] `json:"comment"`
	// A user-friendly name for the subnet.
	Name param.Field[string] `json:"name"`
	// The private IPv4 or IPv6 range defining the subnet, in CIDR notation.
	Network param.Field[string] `json:"network"`
}

func (r NetworkSubnetCloudflareSourceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// IP address family, either `v4` (IPv4) or `v6` (IPv6)
type NetworkSubnetCloudflareSourceUpdateParamsAddressFamily string

const (
	NetworkSubnetCloudflareSourceUpdateParamsAddressFamilyV4 NetworkSubnetCloudflareSourceUpdateParamsAddressFamily = "v4"
	NetworkSubnetCloudflareSourceUpdateParamsAddressFamilyV6 NetworkSubnetCloudflareSourceUpdateParamsAddressFamily = "v6"
)

func (r NetworkSubnetCloudflareSourceUpdateParamsAddressFamily) IsKnown() bool {
	switch r {
	case NetworkSubnetCloudflareSourceUpdateParamsAddressFamilyV4, NetworkSubnetCloudflareSourceUpdateParamsAddressFamilyV6:
		return true
	}
	return false
}

type NetworkSubnetCloudflareSourceUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo                       `json:"errors,required"`
	Messages []shared.ResponseInfo                       `json:"messages,required"`
	Result   NetworkSubnetCloudflareSourceUpdateResponse `json:"result,required"`
	// Whether the API call was successful
	Success NetworkSubnetCloudflareSourceUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    networkSubnetCloudflareSourceUpdateResponseEnvelopeJSON    `json:"-"`
}

// networkSubnetCloudflareSourceUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [NetworkSubnetCloudflareSourceUpdateResponseEnvelope]
type networkSubnetCloudflareSourceUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkSubnetCloudflareSourceUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r networkSubnetCloudflareSourceUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NetworkSubnetCloudflareSourceUpdateResponseEnvelopeSuccess bool

const (
	NetworkSubnetCloudflareSourceUpdateResponseEnvelopeSuccessTrue NetworkSubnetCloudflareSourceUpdateResponseEnvelopeSuccess = true
)

func (r NetworkSubnetCloudflareSourceUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NetworkSubnetCloudflareSourceUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
