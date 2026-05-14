// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DeviceDeviceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceDeviceService] method instead.
type DeviceDeviceService struct {
	Options []option.RequestOption
}

// NewDeviceDeviceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceDeviceService(opts ...option.RequestOption) (r *DeviceDeviceService) {
	r = &DeviceDeviceService{}
	r.Options = opts
	return
}

// Lists WARP devices.
func (r *DeviceDeviceService) List(ctx context.Context, params DeviceDeviceListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[DeviceDeviceListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/physical-devices", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists WARP devices.
func (r *DeviceDeviceService) ListAutoPaging(ctx context.Context, params DeviceDeviceListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[DeviceDeviceListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a WARP device.
func (r *DeviceDeviceService) Delete(ctx context.Context, deviceID string, body DeviceDeviceDeleteParams, opts ...option.RequestOption) (res *DeviceDeviceDeleteResponse, err error) {
	var env DeviceDeviceDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/physical-devices/%s", body.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single WARP device.
func (r *DeviceDeviceService) Get(ctx context.Context, deviceID string, query DeviceDeviceGetParams, opts ...option.RequestOption) (res *DeviceDeviceGetResponse, err error) {
	var env DeviceDeviceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/physical-devices/%s", query.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes all WARP registrations associated with the specified device.
func (r *DeviceDeviceService) Revoke(ctx context.Context, deviceID string, body DeviceDeviceRevokeParams, opts ...option.RequestOption) (res *DeviceDeviceRevokeResponse, err error) {
	var env DeviceDeviceRevokeResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if deviceID == "" {
		err = errors.New("missing required device_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/physical-devices/%s/revoke", body.AccountID, deviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A WARP Device.
type DeviceDeviceListResponse struct {
	// The unique ID of the device.
	ID string `json:"id,required"`
	// The number of active registrations for the device. Active registrations are
	// those which haven't been revoked or deleted.
	ActiveRegistrations int64 `json:"active_registrations,required"`
	// The RFC3339 timestamp when the device was created.
	CreatedAt string `json:"created_at,required"`
	// The RFC3339 timestamp when the device was last seen.
	LastSeenAt string `json:"last_seen_at,required,nullable"`
	// The name of the device.
	Name string `json:"name,required"`
	// The RFC3339 timestamp when the device was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// Version of the WARP client.
	ClientVersion string `json:"client_version,nullable"`
	// The RFC3339 timestamp when the device was deleted.
	DeletedAt string `json:"deleted_at,nullable"`
	// The device operating system.
	DeviceType string `json:"device_type,nullable"`
	// A string that uniquely identifies the hardware or virtual machine (VM).
	HardwareID string `json:"hardware_id,nullable"`
	// The last user to use the WARP device.
	LastSeenUser DeviceDeviceListResponseLastSeenUser `json:"last_seen_user,nullable"`
	// The device MAC address.
	MacAddress string `json:"mac_address,nullable"`
	// The device manufacturer.
	Manufacturer string `json:"manufacturer,nullable"`
	// The model name of the device.
	Model string `json:"model,nullable"`
	// The device operating system version number.
	OSVersion string `json:"os_version,nullable"`
	// Additional operating system version data. For macOS or iOS, the Product Version
	// Extra. For Linux, the kernel release version.
	OSVersionExtra string `json:"os_version_extra,nullable"`
	// The public IP address of the WARP client.
	PublicIP string `json:"public_ip,nullable"`
	// The device serial number.
	SerialNumber string                       `json:"serial_number,nullable"`
	JSON         deviceDeviceListResponseJSON `json:"-"`
}

// deviceDeviceListResponseJSON contains the JSON metadata for the struct
// [DeviceDeviceListResponse]
type deviceDeviceListResponseJSON struct {
	ID                  apijson.Field
	ActiveRegistrations apijson.Field
	CreatedAt           apijson.Field
	LastSeenAt          apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	ClientVersion       apijson.Field
	DeletedAt           apijson.Field
	DeviceType          apijson.Field
	HardwareID          apijson.Field
	LastSeenUser        apijson.Field
	MacAddress          apijson.Field
	Manufacturer        apijson.Field
	Model               apijson.Field
	OSVersion           apijson.Field
	OSVersionExtra      apijson.Field
	PublicIP            apijson.Field
	SerialNumber        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DeviceDeviceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceListResponseJSON) RawJSON() string {
	return r.raw
}

// The last user to use the WARP device.
type DeviceDeviceListResponseLastSeenUser struct {
	// UUID.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                                   `json:"name"`
	JSON deviceDeviceListResponseLastSeenUserJSON `json:"-"`
}

// deviceDeviceListResponseLastSeenUserJSON contains the JSON metadata for the
// struct [DeviceDeviceListResponseLastSeenUser]
type deviceDeviceListResponseLastSeenUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceListResponseLastSeenUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceListResponseLastSeenUserJSON) RawJSON() string {
	return r.raw
}

type DeviceDeviceDeleteResponse = interface{}

// A WARP Device.
type DeviceDeviceGetResponse struct {
	// The unique ID of the device.
	ID string `json:"id,required"`
	// The number of active registrations for the device. Active registrations are
	// those which haven't been revoked or deleted.
	ActiveRegistrations int64 `json:"active_registrations,required"`
	// The RFC3339 timestamp when the device was created.
	CreatedAt string `json:"created_at,required"`
	// The RFC3339 timestamp when the device was last seen.
	LastSeenAt string `json:"last_seen_at,required,nullable"`
	// The name of the device.
	Name string `json:"name,required"`
	// The RFC3339 timestamp when the device was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// Version of the WARP client.
	ClientVersion string `json:"client_version,nullable"`
	// The RFC3339 timestamp when the device was deleted.
	DeletedAt string `json:"deleted_at,nullable"`
	// The device operating system.
	DeviceType string `json:"device_type,nullable"`
	// A string that uniquely identifies the hardware or virtual machine (VM).
	HardwareID string `json:"hardware_id,nullable"`
	// The last user to use the WARP device.
	LastSeenUser DeviceDeviceGetResponseLastSeenUser `json:"last_seen_user,nullable"`
	// The device MAC address.
	MacAddress string `json:"mac_address,nullable"`
	// The device manufacturer.
	Manufacturer string `json:"manufacturer,nullable"`
	// The model name of the device.
	Model string `json:"model,nullable"`
	// The device operating system version number.
	OSVersion string `json:"os_version,nullable"`
	// Additional operating system version data. For macOS or iOS, the Product Version
	// Extra. For Linux, the kernel release version.
	OSVersionExtra string `json:"os_version_extra,nullable"`
	// The public IP address of the WARP client.
	PublicIP string `json:"public_ip,nullable"`
	// The device serial number.
	SerialNumber string                      `json:"serial_number,nullable"`
	JSON         deviceDeviceGetResponseJSON `json:"-"`
}

// deviceDeviceGetResponseJSON contains the JSON metadata for the struct
// [DeviceDeviceGetResponse]
type deviceDeviceGetResponseJSON struct {
	ID                  apijson.Field
	ActiveRegistrations apijson.Field
	CreatedAt           apijson.Field
	LastSeenAt          apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	ClientVersion       apijson.Field
	DeletedAt           apijson.Field
	DeviceType          apijson.Field
	HardwareID          apijson.Field
	LastSeenUser        apijson.Field
	MacAddress          apijson.Field
	Manufacturer        apijson.Field
	Model               apijson.Field
	OSVersion           apijson.Field
	OSVersionExtra      apijson.Field
	PublicIP            apijson.Field
	SerialNumber        apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *DeviceDeviceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceGetResponseJSON) RawJSON() string {
	return r.raw
}

// The last user to use the WARP device.
type DeviceDeviceGetResponseLastSeenUser struct {
	// UUID.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                                  `json:"name"`
	JSON deviceDeviceGetResponseLastSeenUserJSON `json:"-"`
}

// deviceDeviceGetResponseLastSeenUserJSON contains the JSON metadata for the
// struct [DeviceDeviceGetResponseLastSeenUser]
type deviceDeviceGetResponseLastSeenUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceGetResponseLastSeenUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceGetResponseLastSeenUserJSON) RawJSON() string {
	return r.raw
}

type DeviceDeviceRevokeResponse = interface{}

type DeviceDeviceListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter by a one or more device IDs.
	ID param.Field[[]string] `query:"id"`
	// Include or exclude devices with active registrations. The default is "only" -
	// return only devices with active registrations.
	ActiveRegistrations param.Field[DeviceDeviceListParamsActiveRegistrations] `query:"active_registrations"`
	// Opaque token indicating the starting position when requesting the next set of
	// records. A cursor value can be obtained from the result_info.cursor field in the
	// response.
	Cursor       param.Field[string]                             `query:"cursor"`
	Include      param.Field[string]                             `query:"include"`
	LastSeenUser param.Field[DeviceDeviceListParamsLastSeenUser] `query:"last_seen_user"`
	// The maximum number of devices to return in a single response.
	PerPage param.Field[int64] `query:"per_page"`
	// Search by device details.
	Search param.Field[string] `query:"search"`
	// Filter by the last_seen timestamp - returns only devices last seen after this
	// timestamp.
	SeenAfter param.Field[string] `query:"seen_after"`
	// Filter by the last_seen timestamp - returns only devices last seen before this
	// timestamp.
	SeenBefore param.Field[string] `query:"seen_before"`
	// The device field to order results by.
	SortBy param.Field[DeviceDeviceListParamsSortBy] `query:"sort_by"`
	// Sort direction.
	SortOrder param.Field[DeviceDeviceListParamsSortOrder] `query:"sort_order"`
}

// URLQuery serializes [DeviceDeviceListParams]'s query parameters as `url.Values`.
func (r DeviceDeviceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Include or exclude devices with active registrations. The default is "only" -
// return only devices with active registrations.
type DeviceDeviceListParamsActiveRegistrations string

const (
	DeviceDeviceListParamsActiveRegistrationsInclude DeviceDeviceListParamsActiveRegistrations = "include"
	DeviceDeviceListParamsActiveRegistrationsOnly    DeviceDeviceListParamsActiveRegistrations = "only"
	DeviceDeviceListParamsActiveRegistrationsExclude DeviceDeviceListParamsActiveRegistrations = "exclude"
)

func (r DeviceDeviceListParamsActiveRegistrations) IsKnown() bool {
	switch r {
	case DeviceDeviceListParamsActiveRegistrationsInclude, DeviceDeviceListParamsActiveRegistrationsOnly, DeviceDeviceListParamsActiveRegistrationsExclude:
		return true
	}
	return false
}

type DeviceDeviceListParamsLastSeenUser struct {
	// Filter by the last seen user's email.
	Email param.Field[string] `query:"email"`
}

// URLQuery serializes [DeviceDeviceListParamsLastSeenUser]'s query parameters as
// `url.Values`.
func (r DeviceDeviceListParamsLastSeenUser) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The device field to order results by.
type DeviceDeviceListParamsSortBy string

const (
	DeviceDeviceListParamsSortByName                DeviceDeviceListParamsSortBy = "name"
	DeviceDeviceListParamsSortByID                  DeviceDeviceListParamsSortBy = "id"
	DeviceDeviceListParamsSortByClientVersion       DeviceDeviceListParamsSortBy = "client_version"
	DeviceDeviceListParamsSortByLastSeenUserEmail   DeviceDeviceListParamsSortBy = "last_seen_user.email"
	DeviceDeviceListParamsSortByLastSeenAt          DeviceDeviceListParamsSortBy = "last_seen_at"
	DeviceDeviceListParamsSortByActiveRegistrations DeviceDeviceListParamsSortBy = "active_registrations"
	DeviceDeviceListParamsSortByCreatedAt           DeviceDeviceListParamsSortBy = "created_at"
)

func (r DeviceDeviceListParamsSortBy) IsKnown() bool {
	switch r {
	case DeviceDeviceListParamsSortByName, DeviceDeviceListParamsSortByID, DeviceDeviceListParamsSortByClientVersion, DeviceDeviceListParamsSortByLastSeenUserEmail, DeviceDeviceListParamsSortByLastSeenAt, DeviceDeviceListParamsSortByActiveRegistrations, DeviceDeviceListParamsSortByCreatedAt:
		return true
	}
	return false
}

// Sort direction.
type DeviceDeviceListParamsSortOrder string

const (
	DeviceDeviceListParamsSortOrderAsc  DeviceDeviceListParamsSortOrder = "asc"
	DeviceDeviceListParamsSortOrderDesc DeviceDeviceListParamsSortOrder = "desc"
)

func (r DeviceDeviceListParamsSortOrder) IsKnown() bool {
	switch r {
	case DeviceDeviceListParamsSortOrderAsc, DeviceDeviceListParamsSortOrderDesc:
		return true
	}
	return false
}

type DeviceDeviceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDeviceDeleteResponseEnvelope struct {
	Errors   []DeviceDeviceDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDeviceDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success,required"`
	Result  DeviceDeviceDeleteResponse             `json:"result,nullable"`
	JSON    deviceDeviceDeleteResponseEnvelopeJSON `json:"-"`
}

// deviceDeviceDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDeviceDeleteResponseEnvelope]
type deviceDeviceDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeviceDeleteResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    deviceDeviceDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDeviceDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDeviceDeleteResponseEnvelopeErrors]
type deviceDeviceDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeviceDeleteResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    deviceDeviceDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDeviceDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDeviceDeleteResponseEnvelopeMessages]
type deviceDeviceDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDeviceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDeviceGetResponseEnvelope struct {
	Errors   []DeviceDeviceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDeviceGetResponseEnvelopeMessages `json:"messages,required"`
	// A WARP Device.
	Result DeviceDeviceGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success bool                                `json:"success,required"`
	JSON    deviceDeviceGetResponseEnvelopeJSON `json:"-"`
}

// deviceDeviceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDeviceGetResponseEnvelope]
type deviceDeviceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeviceGetResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    deviceDeviceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDeviceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDeviceGetResponseEnvelopeErrors]
type deviceDeviceGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeviceGetResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    deviceDeviceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDeviceGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DeviceDeviceGetResponseEnvelopeMessages]
type deviceDeviceGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceDeviceRevokeParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceDeviceRevokeResponseEnvelope struct {
	Errors   []DeviceDeviceRevokeResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceDeviceRevokeResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                                   `json:"success,required"`
	Result  DeviceDeviceRevokeResponse             `json:"result,nullable"`
	JSON    deviceDeviceRevokeResponseEnvelopeJSON `json:"-"`
}

// deviceDeviceRevokeResponseEnvelopeJSON contains the JSON metadata for the struct
// [DeviceDeviceRevokeResponseEnvelope]
type deviceDeviceRevokeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceRevokeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceRevokeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeviceRevokeResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    deviceDeviceRevokeResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceDeviceRevokeResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DeviceDeviceRevokeResponseEnvelopeErrors]
type deviceDeviceRevokeResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceRevokeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceRevokeResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceDeviceRevokeResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    deviceDeviceRevokeResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceDeviceRevokeResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceDeviceRevokeResponseEnvelopeMessages]
type deviceDeviceRevokeResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceDeviceRevokeResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceDeviceRevokeResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
