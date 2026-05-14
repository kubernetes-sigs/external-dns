// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package addressing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/ips"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// AddressMapService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressMapService] method instead.
type AddressMapService struct {
	Options  []option.RequestOption
	Accounts *AddressMapAccountService
	IPs      *AddressMapIPService
	Zones    *AddressMapZoneService
}

// NewAddressMapService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAddressMapService(opts ...option.RequestOption) (r *AddressMapService) {
	r = &AddressMapService{}
	r.Options = opts
	r.Accounts = NewAddressMapAccountService(opts...)
	r.IPs = NewAddressMapIPService(opts...)
	r.Zones = NewAddressMapZoneService(opts...)
	return
}

// Create a new address map under the account.
func (r *AddressMapService) New(ctx context.Context, params AddressMapNewParams, opts ...option.RequestOption) (res *AddressMapNewResponse, err error) {
	var env AddressMapNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all address maps owned by the account.
func (r *AddressMapService) List(ctx context.Context, query AddressMapListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AddressMap], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps", query.AccountID)
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

// List all address maps owned by the account.
func (r *AddressMapService) ListAutoPaging(ctx context.Context, query AddressMapListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AddressMap] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a particular address map owned by the account. An Address Map must be
// disabled before it can be deleted.
func (r *AddressMapService) Delete(ctx context.Context, addressMapID string, body AddressMapDeleteParams, opts ...option.RequestOption) (res *AddressMapDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", body.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Modify properties of an address map owned by the account.
func (r *AddressMapService) Edit(ctx context.Context, addressMapID string, params AddressMapEditParams, opts ...option.RequestOption) (res *AddressMap, err error) {
	var env AddressMapEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", params.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a particular address map owned by the account.
func (r *AddressMapService) Get(ctx context.Context, addressMapID string, query AddressMapGetParams, opts ...option.RequestOption) (res *AddressMapGetResponse, err error) {
	var env AddressMapGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if addressMapID == "" {
		err = errors.New("missing required address_map_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s", query.AccountID, addressMapID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AddressMap struct {
	// Identifier of an Address Map.
	ID string `json:"id"`
	// If set to false, then the Address Map cannot be deleted via API. This is true
	// for Cloudflare-managed maps.
	CanDelete bool `json:"can_delete"`
	// If set to false, then the IPs on the Address Map cannot be modified via the API.
	// This is true for Cloudflare-managed maps.
	CanModifyIPs bool      `json:"can_modify_ips"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSNI string `json:"default_sni,nullable"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description string `json:"description,nullable"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled    bool           `json:"enabled,nullable"`
	ModifiedAt time.Time      `json:"modified_at" format:"date-time"`
	JSON       addressMapJSON `json:"-"`
}

// addressMapJSON contains the JSON metadata for the struct [AddressMap]
type addressMapJSON struct {
	ID           apijson.Field
	CanDelete    apijson.Field
	CanModifyIPs apijson.Field
	CreatedAt    apijson.Field
	DefaultSNI   apijson.Field
	Description  apijson.Field
	Enabled      apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressMap) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapJSON) RawJSON() string {
	return r.raw
}

// The type of the membership.
type Kind string

const (
	KindZone    Kind = "zone"
	KindAccount Kind = "account"
)

func (r Kind) IsKnown() bool {
	switch r {
	case KindZone, KindAccount:
		return true
	}
	return false
}

type AddressMapNewResponse struct {
	// Identifier of an Address Map.
	ID string `json:"id"`
	// If set to false, then the Address Map cannot be deleted via API. This is true
	// for Cloudflare-managed maps.
	CanDelete bool `json:"can_delete"`
	// If set to false, then the IPs on the Address Map cannot be modified via the API.
	// This is true for Cloudflare-managed maps.
	CanModifyIPs bool      `json:"can_modify_ips"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSNI string `json:"default_sni,nullable"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description string `json:"description,nullable"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled bool `json:"enabled,nullable"`
	// The set of IPs on the Address Map.
	IPs ips.IPs `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressMapNewResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                         `json:"modified_at" format:"date-time"`
	JSON        addressMapNewResponseJSON         `json:"-"`
}

// addressMapNewResponseJSON contains the JSON metadata for the struct
// [AddressMapNewResponse]
type addressMapNewResponseJSON struct {
	ID           apijson.Field
	CanDelete    apijson.Field
	CanModifyIPs apijson.Field
	CreatedAt    apijson.Field
	DefaultSNI   apijson.Field
	Description  apijson.Field
	Enabled      apijson.Field
	IPs          apijson.Field
	Memberships  apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressMapNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The identifier for the membership (eg. a zone or account tag).
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind Kind                                `json:"kind"`
	JSON addressMapNewResponseMembershipJSON `json:"-"`
}

// addressMapNewResponseMembershipJSON contains the JSON metadata for the struct
// [AddressMapNewResponseMembership]
type addressMapNewResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapNewResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseMembershipJSON) RawJSON() string {
	return r.raw
}

type AddressMapDeleteResponse struct {
	Errors   []AddressMapDeleteResponseError   `json:"errors,required"`
	Messages []AddressMapDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful.
	Success    AddressMapDeleteResponseSuccess    `json:"success,required"`
	ResultInfo AddressMapDeleteResponseResultInfo `json:"result_info"`
	JSON       addressMapDeleteResponseJSON       `json:"-"`
}

// addressMapDeleteResponseJSON contains the JSON metadata for the struct
// [AddressMapDeleteResponse]
type addressMapDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AddressMapDeleteResponseError struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           AddressMapDeleteResponseErrorsSource `json:"source"`
	JSON             addressMapDeleteResponseErrorJSON    `json:"-"`
}

// addressMapDeleteResponseErrorJSON contains the JSON metadata for the struct
// [AddressMapDeleteResponseError]
type addressMapDeleteResponseErrorJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type AddressMapDeleteResponseErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    addressMapDeleteResponseErrorsSourceJSON `json:"-"`
}

// addressMapDeleteResponseErrorsSourceJSON contains the JSON metadata for the
// struct [AddressMapDeleteResponseErrorsSource]
type addressMapDeleteResponseErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapDeleteResponseErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AddressMapDeleteResponseMessage struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           AddressMapDeleteResponseMessagesSource `json:"source"`
	JSON             addressMapDeleteResponseMessageJSON    `json:"-"`
}

// addressMapDeleteResponseMessageJSON contains the JSON metadata for the struct
// [AddressMapDeleteResponseMessage]
type addressMapDeleteResponseMessageJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

type AddressMapDeleteResponseMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    addressMapDeleteResponseMessagesSourceJSON `json:"-"`
}

// addressMapDeleteResponseMessagesSourceJSON contains the JSON metadata for the
// struct [AddressMapDeleteResponseMessagesSource]
type addressMapDeleteResponseMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapDeleteResponseMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AddressMapDeleteResponseSuccess bool

const (
	AddressMapDeleteResponseSuccessTrue AddressMapDeleteResponseSuccess = true
)

func (r AddressMapDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case AddressMapDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type AddressMapDeleteResponseResultInfo struct {
	// Total number of results for the requested service.
	Count float64 `json:"count"`
	// Current page within paginated list of results.
	Page float64 `json:"page"`
	// Number of results per page of results.
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters.
	TotalCount float64                                `json:"total_count"`
	JSON       addressMapDeleteResponseResultInfoJSON `json:"-"`
}

// addressMapDeleteResponseResultInfoJSON contains the JSON metadata for the struct
// [AddressMapDeleteResponseResultInfo]
type addressMapDeleteResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapDeleteResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapDeleteResponseResultInfoJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponse struct {
	// Identifier of an Address Map.
	ID string `json:"id"`
	// If set to false, then the Address Map cannot be deleted via API. This is true
	// for Cloudflare-managed maps.
	CanDelete bool `json:"can_delete"`
	// If set to false, then the IPs on the Address Map cannot be modified via the API.
	// This is true for Cloudflare-managed maps.
	CanModifyIPs bool      `json:"can_modify_ips"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSNI string `json:"default_sni,nullable"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description string `json:"description,nullable"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled bool `json:"enabled,nullable"`
	// The set of IPs on the Address Map.
	IPs ips.IPs `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships []AddressMapGetResponseMembership `json:"memberships"`
	ModifiedAt  time.Time                         `json:"modified_at" format:"date-time"`
	JSON        addressMapGetResponseJSON         `json:"-"`
}

// addressMapGetResponseJSON contains the JSON metadata for the struct
// [AddressMapGetResponse]
type addressMapGetResponseJSON struct {
	ID           apijson.Field
	CanDelete    apijson.Field
	CanModifyIPs apijson.Field
	CreatedAt    apijson.Field
	DefaultSNI   apijson.Field
	Description  apijson.Field
	Enabled      apijson.Field
	IPs          apijson.Field
	Memberships  apijson.Field
	ModifiedAt   apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AddressMapGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponseMembership struct {
	// Controls whether the membership can be deleted via the API or not.
	CanDelete bool      `json:"can_delete"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The identifier for the membership (eg. a zone or account tag).
	Identifier string `json:"identifier"`
	// The type of the membership.
	Kind Kind                                `json:"kind"`
	JSON addressMapGetResponseMembershipJSON `json:"-"`
}

// addressMapGetResponseMembershipJSON contains the JSON metadata for the struct
// [AddressMapGetResponseMembership]
type addressMapGetResponseMembershipJSON struct {
	CanDelete   apijson.Field
	CreatedAt   apijson.Field
	Identifier  apijson.Field
	Kind        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapGetResponseMembership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseMembershipJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool]     `json:"enabled"`
	IPs     param.Field[[]string] `json:"ips"`
	// Zones and Accounts which will be assigned IPs on this Address Map. A zone
	// membership will take priority over an account membership.
	Memberships param.Field[[]AddressMapNewParamsMembership] `json:"memberships"`
}

func (r AddressMapNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressMapNewParamsMembership struct {
	// The identifier for the membership (eg. a zone or account tag).
	Identifier param.Field[string] `json:"identifier"`
	// The type of the membership.
	Kind param.Field[Kind] `json:"kind"`
}

func (r AddressMapNewParamsMembership) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressMapNewResponseEnvelope struct {
	Errors   []AddressMapNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressMapNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AddressMapNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AddressMapNewResponse                `json:"result"`
	JSON    addressMapNewResponseEnvelopeJSON    `json:"-"`
}

// addressMapNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapNewResponseEnvelope]
type addressMapNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AddressMapNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             addressMapNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// addressMapNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressMapNewResponseEnvelopeErrors]
type addressMapNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    addressMapNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// addressMapNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AddressMapNewResponseEnvelopeErrorsSource]
type addressMapNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AddressMapNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             addressMapNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// addressMapNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressMapNewResponseEnvelopeMessages]
type addressMapNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AddressMapNewResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    addressMapNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// addressMapNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AddressMapNewResponseEnvelopeMessagesSource]
type addressMapNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AddressMapNewResponseEnvelopeSuccess bool

const (
	AddressMapNewResponseEnvelopeSuccessTrue AddressMapNewResponseEnvelopeSuccess = true
)

func (r AddressMapNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapListParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapDeleteParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapEditParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// If you have legacy TLS clients which do not send the TLS server name indicator,
	// then you can specify one default SNI on the map. If Cloudflare receives a TLS
	// handshake from a client without an SNI, it will respond with the default SNI on
	// those IPs. The default SNI can be any valid zone or subdomain owned by the
	// account.
	DefaultSNI param.Field[string] `json:"default_sni"`
	// An optional description field which may be used to describe the types of IPs or
	// zones on the map.
	Description param.Field[string] `json:"description"`
	// Whether the Address Map is enabled or not. Cloudflare's DNS will not respond
	// with IP addresses on an Address Map until the map is enabled.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AddressMapEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressMapEditResponseEnvelope struct {
	Errors   []AddressMapEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressMapEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AddressMapEditResponseEnvelopeSuccess `json:"success,required"`
	Result  AddressMap                            `json:"result"`
	JSON    addressMapEditResponseEnvelopeJSON    `json:"-"`
}

// addressMapEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapEditResponseEnvelope]
type addressMapEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressMapEditResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           AddressMapEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             addressMapEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// addressMapEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressMapEditResponseEnvelopeErrors]
type addressMapEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressMapEditResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    addressMapEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// addressMapEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [AddressMapEditResponseEnvelopeErrorsSource]
type addressMapEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AddressMapEditResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           AddressMapEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             addressMapEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// addressMapEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressMapEditResponseEnvelopeMessages]
type addressMapEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AddressMapEditResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    addressMapEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// addressMapEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AddressMapEditResponseEnvelopeMessagesSource]
type addressMapEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AddressMapEditResponseEnvelopeSuccess bool

const (
	AddressMapEditResponseEnvelopeSuccessTrue AddressMapEditResponseEnvelopeSuccess = true
)

func (r AddressMapEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressMapGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressMapGetResponseEnvelope struct {
	Errors   []AddressMapGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AddressMapGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AddressMapGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AddressMapGetResponse                `json:"result"`
	JSON    addressMapGetResponseEnvelopeJSON    `json:"-"`
}

// addressMapGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressMapGetResponseEnvelope]
type addressMapGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           AddressMapGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             addressMapGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// addressMapGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AddressMapGetResponseEnvelopeErrors]
type addressMapGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    addressMapGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// addressMapGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [AddressMapGetResponseEnvelopeErrorsSource]
type addressMapGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           AddressMapGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             addressMapGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// addressMapGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AddressMapGetResponseEnvelopeMessages]
type addressMapGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AddressMapGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AddressMapGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    addressMapGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// addressMapGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [AddressMapGetResponseEnvelopeMessagesSource]
type addressMapGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressMapGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressMapGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AddressMapGetResponseEnvelopeSuccess bool

const (
	AddressMapGetResponseEnvelopeSuccessTrue AddressMapGetResponseEnvelopeSuccess = true
)

func (r AddressMapGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressMapGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
