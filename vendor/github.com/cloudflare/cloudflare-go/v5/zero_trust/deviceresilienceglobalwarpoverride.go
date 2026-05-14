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

// DeviceResilienceGlobalWARPOverrideService contains methods and other services
// that help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceResilienceGlobalWARPOverrideService] method instead.
type DeviceResilienceGlobalWARPOverrideService struct {
	Options []option.RequestOption
}

// NewDeviceResilienceGlobalWARPOverrideService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDeviceResilienceGlobalWARPOverrideService(opts ...option.RequestOption) (r *DeviceResilienceGlobalWARPOverrideService) {
	r = &DeviceResilienceGlobalWARPOverrideService{}
	r.Options = opts
	return
}

// Sets the Global WARP override state.
func (r *DeviceResilienceGlobalWARPOverrideService) New(ctx context.Context, params DeviceResilienceGlobalWARPOverrideNewParams, opts ...option.RequestOption) (res *DeviceResilienceGlobalWARPOverrideNewResponse, err error) {
	var env DeviceResilienceGlobalWARPOverrideNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/resilience/disconnect", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch the Global WARP override state.
func (r *DeviceResilienceGlobalWARPOverrideService) Get(ctx context.Context, query DeviceResilienceGlobalWARPOverrideGetParams, opts ...option.RequestOption) (res *DeviceResilienceGlobalWARPOverrideGetResponse, err error) {
	var env DeviceResilienceGlobalWARPOverrideGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/resilience/disconnect", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceResilienceGlobalWARPOverrideNewResponse struct {
	// Disconnects all devices on the account using Global WARP override.
	Disconnect bool `json:"disconnect"`
	// When the Global WARP override state was updated.
	Timestamp time.Time                                         `json:"timestamp" format:"date-time"`
	JSON      deviceResilienceGlobalWARPOverrideNewResponseJSON `json:"-"`
}

// deviceResilienceGlobalWARPOverrideNewResponseJSON contains the JSON metadata for
// the struct [DeviceResilienceGlobalWARPOverrideNewResponse]
type deviceResilienceGlobalWARPOverrideNewResponseJSON struct {
	Disconnect  apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceResilienceGlobalWARPOverrideNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceResilienceGlobalWARPOverrideNewResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceResilienceGlobalWARPOverrideGetResponse struct {
	// Disconnects all devices on the account using Global WARP override.
	Disconnect bool `json:"disconnect"`
	// When the Global WARP override state was updated.
	Timestamp time.Time                                         `json:"timestamp" format:"date-time"`
	JSON      deviceResilienceGlobalWARPOverrideGetResponseJSON `json:"-"`
}

// deviceResilienceGlobalWARPOverrideGetResponseJSON contains the JSON metadata for
// the struct [DeviceResilienceGlobalWARPOverrideGetResponse]
type deviceResilienceGlobalWARPOverrideGetResponseJSON struct {
	Disconnect  apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceResilienceGlobalWARPOverrideGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceResilienceGlobalWARPOverrideGetResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceResilienceGlobalWARPOverrideNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Disconnects all devices on the account using Global WARP override.
	Disconnect param.Field[bool] `json:"disconnect,required"`
	// Reasoning for setting the Global WARP override state. This will be surfaced in
	// the audit log.
	Justification param.Field[string] `json:"justification"`
}

func (r DeviceResilienceGlobalWARPOverrideNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DeviceResilienceGlobalWARPOverrideNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo                         `json:"errors,required"`
	Messages []shared.ResponseInfo                         `json:"messages,required"`
	Result   DeviceResilienceGlobalWARPOverrideNewResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceResilienceGlobalWARPOverrideNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceResilienceGlobalWARPOverrideNewResponseEnvelopeJSON    `json:"-"`
}

// deviceResilienceGlobalWARPOverrideNewResponseEnvelopeJSON contains the JSON
// metadata for the struct [DeviceResilienceGlobalWARPOverrideNewResponseEnvelope]
type deviceResilienceGlobalWARPOverrideNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceResilienceGlobalWARPOverrideNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceResilienceGlobalWARPOverrideNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceResilienceGlobalWARPOverrideNewResponseEnvelopeSuccess bool

const (
	DeviceResilienceGlobalWARPOverrideNewResponseEnvelopeSuccessTrue DeviceResilienceGlobalWARPOverrideNewResponseEnvelopeSuccess = true
)

func (r DeviceResilienceGlobalWARPOverrideNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceResilienceGlobalWARPOverrideNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DeviceResilienceGlobalWARPOverrideGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceResilienceGlobalWARPOverrideGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                         `json:"errors,required"`
	Messages []shared.ResponseInfo                         `json:"messages,required"`
	Result   DeviceResilienceGlobalWARPOverrideGetResponse `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success DeviceResilienceGlobalWARPOverrideGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    deviceResilienceGlobalWARPOverrideGetResponseEnvelopeJSON    `json:"-"`
}

// deviceResilienceGlobalWARPOverrideGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [DeviceResilienceGlobalWARPOverrideGetResponseEnvelope]
type deviceResilienceGlobalWARPOverrideGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceResilienceGlobalWARPOverrideGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceResilienceGlobalWARPOverrideGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DeviceResilienceGlobalWARPOverrideGetResponseEnvelopeSuccess bool

const (
	DeviceResilienceGlobalWARPOverrideGetResponseEnvelopeSuccessTrue DeviceResilienceGlobalWARPOverrideGetResponseEnvelopeSuccess = true
)

func (r DeviceResilienceGlobalWARPOverrideGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DeviceResilienceGlobalWARPOverrideGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
