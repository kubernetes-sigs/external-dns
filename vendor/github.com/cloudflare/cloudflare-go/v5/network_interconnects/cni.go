// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package network_interconnects

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
)

// CNIService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCNIService] method instead.
type CNIService struct {
	Options []option.RequestOption
}

// NewCNIService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCNIService(opts ...option.RequestOption) (r *CNIService) {
	r = &CNIService{}
	r.Options = opts
	return
}

// Create a new CNI object
func (r *CNIService) New(ctx context.Context, params CNINewParams, opts ...option.RequestOption) (res *CNINewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Modify stored information about a CNI object
func (r *CNIService) Update(ctx context.Context, cni string, params CNIUpdateParams, opts ...option.RequestOption) (res *CNIUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cni == "" {
		err = errors.New("missing required cni parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis/%s", params.AccountID, cni)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// List existing CNI objects
func (r *CNIService) List(ctx context.Context, params CNIListParams, opts ...option.RequestOption) (res *CNIListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a specified CNI object
func (r *CNIService) Delete(ctx context.Context, cni string, body CNIDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cni == "" {
		err = errors.New("missing required cni parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis/%s", body.AccountID, cni)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get information about a CNI object
func (r *CNIService) Get(ctx context.Context, cni string, query CNIGetParams, opts ...option.RequestOption) (res *CNIGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if cni == "" {
		err = errors.New("missing required cni parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/cni/cnis/%s", query.AccountID, cni)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type CNINewResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Customer account tag
	Account string `json:"account,required"`
	// Customer end of the point-to-point link
	//
	// This should always be inside the same prefix as `p2p_ip`.
	CustIP string `json:"cust_ip,required" format:"A.B.C.D/N"`
	// Interconnect identifier hosting this CNI
	Interconnect string              `json:"interconnect,required"`
	Magic        CNINewResponseMagic `json:"magic,required"`
	// Cloudflare end of the point-to-point link
	P2pIP string             `json:"p2p_ip,required" format:"A.B.C.D/N"`
	BGP   CNINewResponseBGP  `json:"bgp"`
	JSON  cniNewResponseJSON `json:"-"`
}

// cniNewResponseJSON contains the JSON metadata for the struct [CNINewResponse]
type cniNewResponseJSON struct {
	ID           apijson.Field
	Account      apijson.Field
	CustIP       apijson.Field
	Interconnect apijson.Field
	Magic        apijson.Field
	P2pIP        apijson.Field
	BGP          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CNINewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniNewResponseJSON) RawJSON() string {
	return r.raw
}

type CNINewResponseMagic struct {
	ConduitName string                  `json:"conduit_name,required"`
	Description string                  `json:"description,required"`
	Mtu         int64                   `json:"mtu,required"`
	JSON        cniNewResponseMagicJSON `json:"-"`
}

// cniNewResponseMagicJSON contains the JSON metadata for the struct
// [CNINewResponseMagic]
type cniNewResponseMagicJSON struct {
	ConduitName apijson.Field
	Description apijson.Field
	Mtu         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CNINewResponseMagic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniNewResponseMagicJSON) RawJSON() string {
	return r.raw
}

type CNINewResponseBGP struct {
	// ASN used on the customer end of the BGP session
	CustomerASN int64 `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes []string `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                `json:"md5_key,nullable"`
	JSON   cniNewResponseBGPJSON `json:"-"`
}

// cniNewResponseBGPJSON contains the JSON metadata for the struct
// [CNINewResponseBGP]
type cniNewResponseBGPJSON struct {
	CustomerASN   apijson.Field
	ExtraPrefixes apijson.Field
	Md5Key        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CNINewResponseBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniNewResponseBGPJSON) RawJSON() string {
	return r.raw
}

type CNIUpdateResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Customer account tag
	Account string `json:"account,required"`
	// Customer end of the point-to-point link
	//
	// This should always be inside the same prefix as `p2p_ip`.
	CustIP string `json:"cust_ip,required" format:"A.B.C.D/N"`
	// Interconnect identifier hosting this CNI
	Interconnect string                 `json:"interconnect,required"`
	Magic        CNIUpdateResponseMagic `json:"magic,required"`
	// Cloudflare end of the point-to-point link
	P2pIP string                `json:"p2p_ip,required" format:"A.B.C.D/N"`
	BGP   CNIUpdateResponseBGP  `json:"bgp"`
	JSON  cniUpdateResponseJSON `json:"-"`
}

// cniUpdateResponseJSON contains the JSON metadata for the struct
// [CNIUpdateResponse]
type cniUpdateResponseJSON struct {
	ID           apijson.Field
	Account      apijson.Field
	CustIP       apijson.Field
	Interconnect apijson.Field
	Magic        apijson.Field
	P2pIP        apijson.Field
	BGP          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CNIUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type CNIUpdateResponseMagic struct {
	ConduitName string                     `json:"conduit_name,required"`
	Description string                     `json:"description,required"`
	Mtu         int64                      `json:"mtu,required"`
	JSON        cniUpdateResponseMagicJSON `json:"-"`
}

// cniUpdateResponseMagicJSON contains the JSON metadata for the struct
// [CNIUpdateResponseMagic]
type cniUpdateResponseMagicJSON struct {
	ConduitName apijson.Field
	Description apijson.Field
	Mtu         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CNIUpdateResponseMagic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniUpdateResponseMagicJSON) RawJSON() string {
	return r.raw
}

type CNIUpdateResponseBGP struct {
	// ASN used on the customer end of the BGP session
	CustomerASN int64 `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes []string `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                   `json:"md5_key,nullable"`
	JSON   cniUpdateResponseBGPJSON `json:"-"`
}

// cniUpdateResponseBGPJSON contains the JSON metadata for the struct
// [CNIUpdateResponseBGP]
type cniUpdateResponseBGPJSON struct {
	CustomerASN   apijson.Field
	ExtraPrefixes apijson.Field
	Md5Key        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CNIUpdateResponseBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniUpdateResponseBGPJSON) RawJSON() string {
	return r.raw
}

type CNIListResponse struct {
	Items []CNIListResponseItem `json:"items,required"`
	Next  int64                 `json:"next,nullable"`
	JSON  cniListResponseJSON   `json:"-"`
}

// cniListResponseJSON contains the JSON metadata for the struct [CNIListResponse]
type cniListResponseJSON struct {
	Items       apijson.Field
	Next        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CNIListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniListResponseJSON) RawJSON() string {
	return r.raw
}

type CNIListResponseItem struct {
	ID string `json:"id,required" format:"uuid"`
	// Customer account tag
	Account string `json:"account,required"`
	// Customer end of the point-to-point link
	//
	// This should always be inside the same prefix as `p2p_ip`.
	CustIP string `json:"cust_ip,required" format:"A.B.C.D/N"`
	// Interconnect identifier hosting this CNI
	Interconnect string                    `json:"interconnect,required"`
	Magic        CNIListResponseItemsMagic `json:"magic,required"`
	// Cloudflare end of the point-to-point link
	P2pIP string                  `json:"p2p_ip,required" format:"A.B.C.D/N"`
	BGP   CNIListResponseItemsBGP `json:"bgp"`
	JSON  cniListResponseItemJSON `json:"-"`
}

// cniListResponseItemJSON contains the JSON metadata for the struct
// [CNIListResponseItem]
type cniListResponseItemJSON struct {
	ID           apijson.Field
	Account      apijson.Field
	CustIP       apijson.Field
	Interconnect apijson.Field
	Magic        apijson.Field
	P2pIP        apijson.Field
	BGP          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CNIListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniListResponseItemJSON) RawJSON() string {
	return r.raw
}

type CNIListResponseItemsMagic struct {
	ConduitName string                        `json:"conduit_name,required"`
	Description string                        `json:"description,required"`
	Mtu         int64                         `json:"mtu,required"`
	JSON        cniListResponseItemsMagicJSON `json:"-"`
}

// cniListResponseItemsMagicJSON contains the JSON metadata for the struct
// [CNIListResponseItemsMagic]
type cniListResponseItemsMagicJSON struct {
	ConduitName apijson.Field
	Description apijson.Field
	Mtu         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CNIListResponseItemsMagic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniListResponseItemsMagicJSON) RawJSON() string {
	return r.raw
}

type CNIListResponseItemsBGP struct {
	// ASN used on the customer end of the BGP session
	CustomerASN int64 `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes []string `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                      `json:"md5_key,nullable"`
	JSON   cniListResponseItemsBGPJSON `json:"-"`
}

// cniListResponseItemsBGPJSON contains the JSON metadata for the struct
// [CNIListResponseItemsBGP]
type cniListResponseItemsBGPJSON struct {
	CustomerASN   apijson.Field
	ExtraPrefixes apijson.Field
	Md5Key        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CNIListResponseItemsBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniListResponseItemsBGPJSON) RawJSON() string {
	return r.raw
}

type CNIGetResponse struct {
	ID string `json:"id,required" format:"uuid"`
	// Customer account tag
	Account string `json:"account,required"`
	// Customer end of the point-to-point link
	//
	// This should always be inside the same prefix as `p2p_ip`.
	CustIP string `json:"cust_ip,required" format:"A.B.C.D/N"`
	// Interconnect identifier hosting this CNI
	Interconnect string              `json:"interconnect,required"`
	Magic        CNIGetResponseMagic `json:"magic,required"`
	// Cloudflare end of the point-to-point link
	P2pIP string             `json:"p2p_ip,required" format:"A.B.C.D/N"`
	BGP   CNIGetResponseBGP  `json:"bgp"`
	JSON  cniGetResponseJSON `json:"-"`
}

// cniGetResponseJSON contains the JSON metadata for the struct [CNIGetResponse]
type cniGetResponseJSON struct {
	ID           apijson.Field
	Account      apijson.Field
	CustIP       apijson.Field
	Interconnect apijson.Field
	Magic        apijson.Field
	P2pIP        apijson.Field
	BGP          apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CNIGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniGetResponseJSON) RawJSON() string {
	return r.raw
}

type CNIGetResponseMagic struct {
	ConduitName string                  `json:"conduit_name,required"`
	Description string                  `json:"description,required"`
	Mtu         int64                   `json:"mtu,required"`
	JSON        cniGetResponseMagicJSON `json:"-"`
}

// cniGetResponseMagicJSON contains the JSON metadata for the struct
// [CNIGetResponseMagic]
type cniGetResponseMagicJSON struct {
	ConduitName apijson.Field
	Description apijson.Field
	Mtu         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CNIGetResponseMagic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniGetResponseMagicJSON) RawJSON() string {
	return r.raw
}

type CNIGetResponseBGP struct {
	// ASN used on the customer end of the BGP session
	CustomerASN int64 `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes []string `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key string                `json:"md5_key,nullable"`
	JSON   cniGetResponseBGPJSON `json:"-"`
}

// cniGetResponseBGPJSON contains the JSON metadata for the struct
// [CNIGetResponseBGP]
type cniGetResponseBGPJSON struct {
	CustomerASN   apijson.Field
	ExtraPrefixes apijson.Field
	Md5Key        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *CNIGetResponseBGP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cniGetResponseBGPJSON) RawJSON() string {
	return r.raw
}

type CNINewParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
	// Customer account tag
	Account      param.Field[string]            `json:"account,required"`
	Interconnect param.Field[string]            `json:"interconnect,required"`
	Magic        param.Field[CNINewParamsMagic] `json:"magic,required"`
	BGP          param.Field[CNINewParamsBGP]   `json:"bgp"`
}

func (r CNINewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CNINewParamsMagic struct {
	ConduitName param.Field[string] `json:"conduit_name,required"`
	Description param.Field[string] `json:"description,required"`
	Mtu         param.Field[int64]  `json:"mtu,required"`
}

func (r CNINewParamsMagic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CNINewParamsBGP struct {
	// ASN used on the customer end of the BGP session
	CustomerASN param.Field[int64] `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes param.Field[[]string] `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key param.Field[string] `json:"md5_key"`
}

func (r CNINewParamsBGP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CNIUpdateParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
	ID        param.Field[string] `json:"id,required" format:"uuid"`
	// Customer account tag
	Account param.Field[string] `json:"account,required"`
	// Customer end of the point-to-point link
	//
	// This should always be inside the same prefix as `p2p_ip`.
	CustIP param.Field[string] `json:"cust_ip,required" format:"A.B.C.D/N"`
	// Interconnect identifier hosting this CNI
	Interconnect param.Field[string]               `json:"interconnect,required"`
	Magic        param.Field[CNIUpdateParamsMagic] `json:"magic,required"`
	// Cloudflare end of the point-to-point link
	P2pIP param.Field[string]             `json:"p2p_ip,required" format:"A.B.C.D/N"`
	BGP   param.Field[CNIUpdateParamsBGP] `json:"bgp"`
}

func (r CNIUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CNIUpdateParamsMagic struct {
	ConduitName param.Field[string] `json:"conduit_name,required"`
	Description param.Field[string] `json:"description,required"`
	Mtu         param.Field[int64]  `json:"mtu,required"`
}

func (r CNIUpdateParamsMagic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CNIUpdateParamsBGP struct {
	// ASN used on the customer end of the BGP session
	CustomerASN param.Field[int64] `json:"customer_asn,required"`
	// Extra set of static prefixes to advertise to the customer's end of the session
	ExtraPrefixes param.Field[[]string] `json:"extra_prefixes,required" format:"A.B.C.D/N"`
	// MD5 key to use for session authentication.
	//
	// Note that _this is not a security measure_. MD5 is not a valid security
	// mechanism, and the key is not treated as a secret value. This is _only_
	// supported for preventing misconfiguration, not for defending against malicious
	// attacks.
	//
	// The MD5 key, if set, must be of non-zero length and consist only of the
	// following types of character:
	//
	// - ASCII alphanumerics: `[a-zA-Z0-9]`
	// - Special characters in the set `'!@#$%^&*()+[]{}<>/.,;:_-~`= \|`
	//
	// In other words, MD5 keys may contain any printable ASCII character aside from
	// newline (0x0A), quotation mark (`"`), vertical tab (0x0B), carriage return
	// (0x0D), tab (0x09), form feed (0x0C), and the question mark (`?`). Requests
	// specifying an MD5 key with one or more of these disallowed characters will be
	// rejected.
	Md5Key param.Field[string] `json:"md5_key"`
}

func (r CNIUpdateParamsBGP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CNIListParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
	Cursor    param.Field[int64]  `query:"cursor"`
	Limit     param.Field[int64]  `query:"limit"`
	// If specified, only show CNIs associated with the specified slot
	Slot param.Field[string] `query:"slot"`
	// If specified, only show cnis associated with the specified tunnel id
	TunnelID param.Field[string] `query:"tunnel_id"`
}

// URLQuery serializes [CNIListParams]'s query parameters as `url.Values`.
func (r CNIListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type CNIDeleteParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
}

type CNIGetParams struct {
	// Customer account tag
	AccountID param.Field[string] `path:"account_id,required"`
}
