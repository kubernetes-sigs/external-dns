// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// DeviceUnrevokeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceUnrevokeService] method instead.
type DeviceUnrevokeService struct {
	Options []option.RequestOption
}

// NewDeviceUnrevokeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceUnrevokeService(opts ...option.RequestOption) (r *DeviceUnrevokeService) {
	r = &DeviceUnrevokeService{}
	r.Options = opts
	return
}

// Unrevokes a list of devices. Not supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled.
//
// **Deprecated**: please use POST
// /accounts/{account_id}/devices/registrations/unrevoke instead.
//
// Deprecated: deprecated
func (r *DeviceUnrevokeService) New(ctx context.Context, params DeviceUnrevokeNewParams, opts ...option.RequestOption) (res *interface{}, err error) {
	var env DeviceUnrevokeNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/unrevoke", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceUnrevokeNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of Registration IDs to unrevoke.
	Body []string `json:"body,required"`
}

func (r DeviceUnrevokeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type DeviceUnrevokeNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   interface{}           `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceUnrevokeNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceUnrevokeNewResponseEnvelopeJSON    `json:"-"`
}

// deviceUnrevokeNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceUnrevokeNewResponseEnvelope]
type deviceUnrevokeNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceUnrevokeNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceUnrevokeNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceUnrevokeNewResponseEnvelopeSuccess bool

const (
	DeviceUnrevokeNewResponseEnvelopeSuccessTrue DeviceUnrevokeNewResponseEnvelopeSuccess = true
)

func (r DeviceUnrevokeNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceUnrevokeNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
