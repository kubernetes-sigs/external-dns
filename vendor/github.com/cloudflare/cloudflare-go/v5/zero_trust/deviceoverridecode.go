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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DeviceOverrideCodeService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceOverrideCodeService] method instead.
type DeviceOverrideCodeService struct {
	Options []option.RequestOption
}

// NewDeviceOverrideCodeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceOverrideCodeService(opts ...option.RequestOption) (r *DeviceOverrideCodeService) {
	r = &DeviceOverrideCodeService{}
	r.Options = opts
	return
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration. Not
// supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled for the account. **Deprecated:** please use GET
// /accounts/{account_id}/devices/registrations/{registration_id}/override_codes
// instead.
//
// Deprecated: deprecated
func (r *DeviceOverrideCodeService) List(ctx context.Context, deviceID string, query DeviceOverrideCodeListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DeviceOverrideCodeListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/%s/override_codes", query.AccountID, deviceID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Fetches a one-time use admin override code for a device. This relies on the
// **Admin Override** setting being enabled in your device configuration. Not
// supported when
// [multi-user mode](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/warp/deployment/mdm-deployment/windows-multiuser/)
// is enabled for the account. **Deprecated:** please use GET
// /accounts/{account_id}/devices/registrations/{registration_id}/override_codes
// instead.
//
// Deprecated: deprecated
func (r *DeviceOverrideCodeService) ListAutoPaging(ctx context.Context, deviceID string, query DeviceOverrideCodeListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DeviceOverrideCodeListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, deviceID, query, opts...))
}

// Fetches one-time use admin override codes for a registration. This relies on the
// **Admin Override** setting being enabled in your device configuration.
func (r *DeviceOverrideCodeService) Get(ctx context.Context, registrationID string, query DeviceOverrideCodeGetParams, opts ...option.RequestOption) (res *DeviceOverrideCodeGetResponse, err error) {
	var env DeviceOverrideCodeGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if registrationID == "" {
		err = errors.New("missing required registration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations/%s/override_codes", query.AccountID, registrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DeviceOverrideCodeListResponse = interface{}

type DeviceOverrideCodeGetResponse struct {
	DisableForTime map[string]string                 `json:"disable_for_time"`
	JSON           deviceOverrideCodeGetResponseJSON `json:"-"`
}

// deviceOverrideCodeGetResponseJSON contains the JSON metadata for the struct
// [DeviceOverrideCodeGetResponse]
type deviceOverrideCodeGetResponseJSON struct {
	DisableForTime apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseJSON) RawJSON() string {
	return r.raw
}

type DeviceOverrideCodeListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceOverrideCodeGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceOverrideCodeGetResponseEnvelope struct {
	Errors   []DeviceOverrideCodeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceOverrideCodeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceOverrideCodeGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success,required"`
	JSON    deviceOverrideCodeGetResponseEnvelopeJSON `json:"-"`
}

// deviceOverrideCodeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceOverrideCodeGetResponseEnvelope]
type deviceOverrideCodeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceOverrideCodeGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceOverrideCodeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceOverrideCodeGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeviceOverrideCodeGetResponseEnvelopeErrors]
type deviceOverrideCodeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceOverrideCodeGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    deviceOverrideCodeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceOverrideCodeGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceOverrideCodeGetResponseEnvelopeMessages]
type deviceOverrideCodeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceOverrideCodeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceOverrideCodeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
