// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_network_monitoring

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

// VPCFlowTokenService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVPCFlowTokenService] method instead.
type VPCFlowTokenService struct {
	Options []option.RequestOption
}

// NewVPCFlowTokenService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVPCFlowTokenService(opts ...option.RequestOption) (r *VPCFlowTokenService) {
	r = &VPCFlowTokenService{}
	r.Options = opts
	return
}

// Generate authentication token for VPC flow logs export.
func (r *VPCFlowTokenService) New(ctx context.Context, body VPCFlowTokenNewParams, opts ...option.RequestOption) (res *string, err error) {
	var env VPCFlowTokenNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/mnm/vpc-flows/token", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type VPCFlowTokenNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type VPCFlowTokenNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Authentication token to be used for VPC Flows export authentication.
	Result string `json:"result,required"`
	// Whether the API call was successful
	Success VPCFlowTokenNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    vpcFlowTokenNewResponseEnvelopeJSON    `json:"-"`
}

// vpcFlowTokenNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [VPCFlowTokenNewResponseEnvelope]
type vpcFlowTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VPCFlowTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r vpcFlowTokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type VPCFlowTokenNewResponseEnvelopeSuccess bool

const (
	VPCFlowTokenNewResponseEnvelopeSuccessTrue VPCFlowTokenNewResponseEnvelopeSuccess = true
)

func (r VPCFlowTokenNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VPCFlowTokenNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
