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

// DeviceRegistrationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceRegistrationService] method instead.
type DeviceRegistrationService struct {
	Options []option.RequestOption
}

// NewDeviceRegistrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceRegistrationService(opts ...option.RequestOption) (r *DeviceRegistrationService) {
	r = &DeviceRegistrationService{}
	r.Options = opts
	return
}

// Lists WARP registrations.
func (r *DeviceRegistrationService) List(ctx context.Context, params DeviceRegistrationListParams, opts ...option.RequestOption) (res *pagination.CursorPagination[DeviceRegistrationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations", params.AccountID)
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

// Lists WARP registrations.
func (r *DeviceRegistrationService) ListAutoPaging(ctx context.Context, params DeviceRegistrationListParams, opts ...option.RequestOption) *pagination.CursorPaginationAutoPager[DeviceRegistrationListResponse] {
	return pagination.NewCursorPaginationAutoPager(r.List(ctx, params, opts...))
}

// Deletes a WARP registration.
func (r *DeviceRegistrationService) Delete(ctx context.Context, registrationID string, body DeviceRegistrationDeleteParams, opts ...option.RequestOption) (res *DeviceRegistrationDeleteResponse, err error) {
	var env DeviceRegistrationDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if registrationID == "" {
		err = errors.New("missing required registration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations/%s", body.AccountID, registrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a list of WARP registrations.
func (r *DeviceRegistrationService) BulkDelete(ctx context.Context, params DeviceRegistrationBulkDeleteParams, opts ...option.RequestOption) (res *DeviceRegistrationBulkDeleteResponse, err error) {
	var env DeviceRegistrationBulkDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single WARP registration.
func (r *DeviceRegistrationService) Get(ctx context.Context, registrationID string, query DeviceRegistrationGetParams, opts ...option.RequestOption) (res *DeviceRegistrationGetResponse, err error) {
	var env DeviceRegistrationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if registrationID == "" {
		err = errors.New("missing required registration_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations/%s", query.AccountID, registrationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes a list of WARP registrations.
func (r *DeviceRegistrationService) Revoke(ctx context.Context, params DeviceRegistrationRevokeParams, opts ...option.RequestOption) (res *DeviceRegistrationRevokeResponse, err error) {
	var env DeviceRegistrationRevokeResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations/revoke", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Unrevokes a list of WARP registrations.
func (r *DeviceRegistrationService) Unrevoke(ctx context.Context, params DeviceRegistrationUnrevokeParams, opts ...option.RequestOption) (res *DeviceRegistrationUnrevokeResponse, err error) {
	var env DeviceRegistrationUnrevokeResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/devices/registrations/unrevoke", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A WARP configuration tied to a single user. Multiple registrations can be
// created from a single WARP device.
type DeviceRegistrationListResponse struct {
	// The ID of the registration.
	ID string `json:"id,required"`
	// The RFC3339 timestamp when the registration was created.
	CreatedAt string `json:"created_at,required"`
	// Device details embedded inside of a registration.
	Device DeviceRegistrationListResponseDevice `json:"device,required"`
	// The public key used to connect to the Cloudflare network.
	Key string `json:"key,required"`
	// The RFC3339 timestamp when the registration was last seen.
	LastSeenAt string `json:"last_seen_at,required"`
	// The RFC3339 timestamp when the registration was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// The RFC3339 timestamp when the registration was deleted.
	DeletedAt string `json:"deleted_at,nullable"`
	// The type of encryption key used by the WARP client for the active key. Currently
	// 'curve25519' for WireGuard and 'secp256r1' for MASQUE.
	KeyType string `json:"key_type,nullable"`
	// The RFC3339 timestamp when the registration was revoked.
	RevokedAt string `json:"revoked_at,nullable"`
	// Type of the tunnel - wireguard or masque.
	TunnelType string                             `json:"tunnel_type,nullable"`
	User       DeviceRegistrationListResponseUser `json:"user"`
	JSON       deviceRegistrationListResponseJSON `json:"-"`
}

// deviceRegistrationListResponseJSON contains the JSON metadata for the struct
// [DeviceRegistrationListResponse]
type deviceRegistrationListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Device      apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	UpdatedAt   apijson.Field
	DeletedAt   apijson.Field
	KeyType     apijson.Field
	RevokedAt   apijson.Field
	TunnelType  apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationListResponseJSON) RawJSON() string {
	return r.raw
}

// Device details embedded inside of a registration.
type DeviceRegistrationListResponseDevice struct {
	// The ID of the device.
	ID string `json:"id,required"`
	// The name of the device.
	Name string `json:"name,required"`
	// Version of the WARP client.
	ClientVersion string                                   `json:"client_version"`
	JSON          deviceRegistrationListResponseDeviceJSON `json:"-"`
}

// deviceRegistrationListResponseDeviceJSON contains the JSON metadata for the
// struct [DeviceRegistrationListResponseDevice]
type deviceRegistrationListResponseDeviceJSON struct {
	ID            apijson.Field
	Name          apijson.Field
	ClientVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeviceRegistrationListResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationListResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationListResponseUser struct {
	// UUID.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                                 `json:"name"`
	JSON deviceRegistrationListResponseUserJSON `json:"-"`
}

// deviceRegistrationListResponseUserJSON contains the JSON metadata for the struct
// [DeviceRegistrationListResponseUser]
type deviceRegistrationListResponseUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationListResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationListResponseUserJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationDeleteResponse = interface{}

type DeviceRegistrationBulkDeleteResponse = interface{}

// A WARP configuration tied to a single user. Multiple registrations can be
// created from a single WARP device.
type DeviceRegistrationGetResponse struct {
	// The ID of the registration.
	ID string `json:"id,required"`
	// The RFC3339 timestamp when the registration was created.
	CreatedAt string `json:"created_at,required"`
	// Device details embedded inside of a registration.
	Device DeviceRegistrationGetResponseDevice `json:"device,required"`
	// The public key used to connect to the Cloudflare network.
	Key string `json:"key,required"`
	// The RFC3339 timestamp when the registration was last seen.
	LastSeenAt string `json:"last_seen_at,required"`
	// The RFC3339 timestamp when the registration was last updated.
	UpdatedAt string `json:"updated_at,required"`
	// The RFC3339 timestamp when the registration was deleted.
	DeletedAt string `json:"deleted_at,nullable"`
	// The type of encryption key used by the WARP client for the active key. Currently
	// 'curve25519' for WireGuard and 'secp256r1' for MASQUE.
	KeyType string `json:"key_type,nullable"`
	// The RFC3339 timestamp when the registration was revoked.
	RevokedAt string `json:"revoked_at,nullable"`
	// Type of the tunnel - wireguard or masque.
	TunnelType string                            `json:"tunnel_type,nullable"`
	User       DeviceRegistrationGetResponseUser `json:"user"`
	JSON       deviceRegistrationGetResponseJSON `json:"-"`
}

// deviceRegistrationGetResponseJSON contains the JSON metadata for the struct
// [DeviceRegistrationGetResponse]
type deviceRegistrationGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Device      apijson.Field
	Key         apijson.Field
	LastSeenAt  apijson.Field
	UpdatedAt   apijson.Field
	DeletedAt   apijson.Field
	KeyType     apijson.Field
	RevokedAt   apijson.Field
	TunnelType  apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationGetResponseJSON) RawJSON() string {
	return r.raw
}

// Device details embedded inside of a registration.
type DeviceRegistrationGetResponseDevice struct {
	// The ID of the device.
	ID string `json:"id,required"`
	// The name of the device.
	Name string `json:"name,required"`
	// Version of the WARP client.
	ClientVersion string                                  `json:"client_version"`
	JSON          deviceRegistrationGetResponseDeviceJSON `json:"-"`
}

// deviceRegistrationGetResponseDeviceJSON contains the JSON metadata for the
// struct [DeviceRegistrationGetResponseDevice]
type deviceRegistrationGetResponseDeviceJSON struct {
	ID            apijson.Field
	Name          apijson.Field
	ClientVersion apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DeviceRegistrationGetResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationGetResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationGetResponseUser struct {
	// UUID.
	ID string `json:"id"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The enrolled device user's name.
	Name string                                `json:"name"`
	JSON deviceRegistrationGetResponseUserJSON `json:"-"`
}

// deviceRegistrationGetResponseUserJSON contains the JSON metadata for the struct
// [DeviceRegistrationGetResponseUser]
type deviceRegistrationGetResponseUserJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationGetResponseUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationGetResponseUserJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationRevokeResponse = interface{}

type DeviceRegistrationUnrevokeResponse = interface{}

type DeviceRegistrationListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Filter by registration ID.
	ID param.Field[[]string] `query:"id"`
	// Opaque token indicating the starting position when requesting the next set of
	// records. A cursor value can be obtained from the result_info.cursor field in the
	// response.
	Cursor  param.Field[string]                             `query:"cursor"`
	Device  param.Field[DeviceRegistrationListParamsDevice] `query:"device"`
	Include param.Field[string]                             `query:"include"`
	// The maximum number of devices to return in a single response.
	PerPage param.Field[int64] `query:"per_page"`
	// Filter by registration details.
	Search param.Field[string] `query:"search"`
	// Filter by the last_seen timestamp - returns only registrations last seen after
	// this timestamp.
	SeenAfter param.Field[string] `query:"seen_after"`
	// Filter by the last_seen timestamp - returns only registrations last seen before
	// this timestamp.
	SeenBefore param.Field[string] `query:"seen_before"`
	// The registration field to order results by.
	SortBy param.Field[DeviceRegistrationListParamsSortBy] `query:"sort_by"`
	// Sort direction.
	SortOrder param.Field[DeviceRegistrationListParamsSortOrder] `query:"sort_order"`
	// Filter by registration status. Defaults to 'active'.
	Status param.Field[DeviceRegistrationListParamsStatus] `query:"status"`
	User   param.Field[DeviceRegistrationListParamsUser]   `query:"user"`
}

// URLQuery serializes [DeviceRegistrationListParams]'s query parameters as
// `url.Values`.
func (r DeviceRegistrationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeviceRegistrationListParamsDevice struct {
	// Filter by WARP device ID.
	ID param.Field[string] `query:"id"`
}

// URLQuery serializes [DeviceRegistrationListParamsDevice]'s query parameters as
// `url.Values`.
func (r DeviceRegistrationListParamsDevice) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The registration field to order results by.
type DeviceRegistrationListParamsSortBy string

const (
	DeviceRegistrationListParamsSortByID         DeviceRegistrationListParamsSortBy = "id"
	DeviceRegistrationListParamsSortByUserName   DeviceRegistrationListParamsSortBy = "user.name"
	DeviceRegistrationListParamsSortByUserEmail  DeviceRegistrationListParamsSortBy = "user.email"
	DeviceRegistrationListParamsSortByLastSeenAt DeviceRegistrationListParamsSortBy = "last_seen_at"
	DeviceRegistrationListParamsSortByCreatedAt  DeviceRegistrationListParamsSortBy = "created_at"
)

func (r DeviceRegistrationListParamsSortBy) IsKnown() bool {
	switch r {
	case DeviceRegistrationListParamsSortByID, DeviceRegistrationListParamsSortByUserName, DeviceRegistrationListParamsSortByUserEmail, DeviceRegistrationListParamsSortByLastSeenAt, DeviceRegistrationListParamsSortByCreatedAt:
		return true
	}
	return false
}

// Sort direction.
type DeviceRegistrationListParamsSortOrder string

const (
	DeviceRegistrationListParamsSortOrderAsc  DeviceRegistrationListParamsSortOrder = "asc"
	DeviceRegistrationListParamsSortOrderDesc DeviceRegistrationListParamsSortOrder = "desc"
)

func (r DeviceRegistrationListParamsSortOrder) IsKnown() bool {
	switch r {
	case DeviceRegistrationListParamsSortOrderAsc, DeviceRegistrationListParamsSortOrderDesc:
		return true
	}
	return false
}

// Filter by registration status. Defaults to 'active'.
type DeviceRegistrationListParamsStatus string

const (
	DeviceRegistrationListParamsStatusActive  DeviceRegistrationListParamsStatus = "active"
	DeviceRegistrationListParamsStatusAll     DeviceRegistrationListParamsStatus = "all"
	DeviceRegistrationListParamsStatusRevoked DeviceRegistrationListParamsStatus = "revoked"
)

func (r DeviceRegistrationListParamsStatus) IsKnown() bool {
	switch r {
	case DeviceRegistrationListParamsStatusActive, DeviceRegistrationListParamsStatusAll, DeviceRegistrationListParamsStatusRevoked:
		return true
	}
	return false
}

type DeviceRegistrationListParamsUser struct {
	// Filter by user ID.
	ID param.Field[[]string] `query:"id"`
}

// URLQuery serializes [DeviceRegistrationListParamsUser]'s query parameters as
// `url.Values`.
func (r DeviceRegistrationListParamsUser) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeviceRegistrationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceRegistrationDeleteResponseEnvelope struct {
	Errors   []DeviceRegistrationDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceRegistrationDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                                         `json:"success,required"`
	Result  DeviceRegistrationDeleteResponse             `json:"result,nullable"`
	JSON    deviceRegistrationDeleteResponseEnvelopeJSON `json:"-"`
}

// deviceRegistrationDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceRegistrationDeleteResponseEnvelope]
type deviceRegistrationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    deviceRegistrationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceRegistrationDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceRegistrationDeleteResponseEnvelopeErrors]
type deviceRegistrationDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    deviceRegistrationDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceRegistrationDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DeviceRegistrationDeleteResponseEnvelopeMessages]
type deviceRegistrationDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationBulkDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of registration IDs to delete.
	ID param.Field[[]string] `query:"id,required"`
}

// URLQuery serializes [DeviceRegistrationBulkDeleteParams]'s query parameters as
// `url.Values`.
func (r DeviceRegistrationBulkDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeviceRegistrationBulkDeleteResponseEnvelope struct {
	Errors   []DeviceRegistrationBulkDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceRegistrationBulkDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceRegistrationBulkDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success bool `json:"success,required"`
	// V4 public API Pagination/Cursor info.
	ResultInfo DeviceRegistrationBulkDeleteResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceRegistrationBulkDeleteResponseEnvelopeJSON       `json:"-"`
}

// deviceRegistrationBulkDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [DeviceRegistrationBulkDeleteResponseEnvelope]
type deviceRegistrationBulkDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationBulkDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationBulkDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationBulkDeleteResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    deviceRegistrationBulkDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceRegistrationBulkDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DeviceRegistrationBulkDeleteResponseEnvelopeErrors]
type deviceRegistrationBulkDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationBulkDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationBulkDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationBulkDeleteResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    deviceRegistrationBulkDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceRegistrationBulkDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DeviceRegistrationBulkDeleteResponseEnvelopeMessages]
type deviceRegistrationBulkDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationBulkDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationBulkDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// V4 public API Pagination/Cursor info.
type DeviceRegistrationBulkDeleteResponseEnvelopeResultInfo struct {
	// Number of records in the response.
	Count int64 `json:"count,required"`
	// Opaque token to request the next set of records.
	Cursor string `json:"cursor,required"`
	// The limit for the number of records in the response.
	PerPage int64 `json:"per_page,required"`
	// Total number of records available.
	TotalCount int64                                                      `json:"total_count,nullable"`
	JSON       deviceRegistrationBulkDeleteResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceRegistrationBulkDeleteResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DeviceRegistrationBulkDeleteResponseEnvelopeResultInfo]
type deviceRegistrationBulkDeleteResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationBulkDeleteResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationBulkDeleteResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DeviceRegistrationGetResponseEnvelope struct {
	Errors   []DeviceRegistrationGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceRegistrationGetResponseEnvelopeMessages `json:"messages,required"`
	// A WARP configuration tied to a single user. Multiple registrations can be
	// created from a single WARP device.
	Result DeviceRegistrationGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success bool                                      `json:"success,required"`
	JSON    deviceRegistrationGetResponseEnvelopeJSON `json:"-"`
}

// deviceRegistrationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceRegistrationGetResponseEnvelope]
type deviceRegistrationGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    deviceRegistrationGetResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceRegistrationGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DeviceRegistrationGetResponseEnvelopeErrors]
type deviceRegistrationGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    deviceRegistrationGetResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceRegistrationGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DeviceRegistrationGetResponseEnvelopeMessages]
type deviceRegistrationGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationRevokeParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of registration IDs to revoke.
	ID param.Field[[]string] `query:"id,required"`
}

// URLQuery serializes [DeviceRegistrationRevokeParams]'s query parameters as
// `url.Values`.
func (r DeviceRegistrationRevokeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeviceRegistrationRevokeResponseEnvelope struct {
	Errors   []DeviceRegistrationRevokeResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceRegistrationRevokeResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceRegistrationRevokeResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success bool `json:"success,required"`
	// V4 public API Pagination/Cursor info.
	ResultInfo DeviceRegistrationRevokeResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceRegistrationRevokeResponseEnvelopeJSON       `json:"-"`
}

// deviceRegistrationRevokeResponseEnvelopeJSON contains the JSON metadata for the
// struct [DeviceRegistrationRevokeResponseEnvelope]
type deviceRegistrationRevokeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationRevokeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationRevokeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationRevokeResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    deviceRegistrationRevokeResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceRegistrationRevokeResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceRegistrationRevokeResponseEnvelopeErrors]
type deviceRegistrationRevokeResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationRevokeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationRevokeResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationRevokeResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    deviceRegistrationRevokeResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceRegistrationRevokeResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DeviceRegistrationRevokeResponseEnvelopeMessages]
type deviceRegistrationRevokeResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationRevokeResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationRevokeResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// V4 public API Pagination/Cursor info.
type DeviceRegistrationRevokeResponseEnvelopeResultInfo struct {
	// Number of records in the response.
	Count int64 `json:"count,required"`
	// Opaque token to request the next set of records.
	Cursor string `json:"cursor,required"`
	// The limit for the number of records in the response.
	PerPage int64 `json:"per_page,required"`
	// Total number of records available.
	TotalCount int64                                                  `json:"total_count,nullable"`
	JSON       deviceRegistrationRevokeResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceRegistrationRevokeResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DeviceRegistrationRevokeResponseEnvelopeResultInfo]
type deviceRegistrationRevokeResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationRevokeResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationRevokeResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type DeviceRegistrationUnrevokeParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// A list of registration IDs to unrevoke.
	ID param.Field[[]string] `query:"id,required"`
}

// URLQuery serializes [DeviceRegistrationUnrevokeParams]'s query parameters as
// `url.Values`.
func (r DeviceRegistrationUnrevokeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DeviceRegistrationUnrevokeResponseEnvelope struct {
	Errors   []DeviceRegistrationUnrevokeResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DeviceRegistrationUnrevokeResponseEnvelopeMessages `json:"messages,required"`
	Result   DeviceRegistrationUnrevokeResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success bool `json:"success,required"`
	// V4 public API Pagination/Cursor info.
	ResultInfo DeviceRegistrationUnrevokeResponseEnvelopeResultInfo `json:"result_info"`
	JSON       deviceRegistrationUnrevokeResponseEnvelopeJSON       `json:"-"`
}

// deviceRegistrationUnrevokeResponseEnvelopeJSON contains the JSON metadata for
// the struct [DeviceRegistrationUnrevokeResponseEnvelope]
type deviceRegistrationUnrevokeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationUnrevokeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationUnrevokeResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationUnrevokeResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    deviceRegistrationUnrevokeResponseEnvelopeErrorsJSON `json:"-"`
}

// deviceRegistrationUnrevokeResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DeviceRegistrationUnrevokeResponseEnvelopeErrors]
type deviceRegistrationUnrevokeResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationUnrevokeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationUnrevokeResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// A message which can be returned in either the 'errors' or 'messages' fields in a
// v4 API response.
type DeviceRegistrationUnrevokeResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    deviceRegistrationUnrevokeResponseEnvelopeMessagesJSON `json:"-"`
}

// deviceRegistrationUnrevokeResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DeviceRegistrationUnrevokeResponseEnvelopeMessages]
type deviceRegistrationUnrevokeResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationUnrevokeResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationUnrevokeResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// V4 public API Pagination/Cursor info.
type DeviceRegistrationUnrevokeResponseEnvelopeResultInfo struct {
	// Number of records in the response.
	Count int64 `json:"count,required"`
	// Opaque token to request the next set of records.
	Cursor string `json:"cursor,required"`
	// The limit for the number of records in the response.
	PerPage int64 `json:"per_page,required"`
	// Total number of records available.
	TotalCount int64                                                    `json:"total_count,nullable"`
	JSON       deviceRegistrationUnrevokeResponseEnvelopeResultInfoJSON `json:"-"`
}

// deviceRegistrationUnrevokeResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [DeviceRegistrationUnrevokeResponseEnvelopeResultInfo]
type deviceRegistrationUnrevokeResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Cursor      apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DeviceRegistrationUnrevokeResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r deviceRegistrationUnrevokeResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
