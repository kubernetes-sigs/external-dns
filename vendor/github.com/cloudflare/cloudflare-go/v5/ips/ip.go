// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ips

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/tidwall/gjson"
)

// IPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIPService] method instead.
type IPService struct {
	Options []option.RequestOption
}

// NewIPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIPService(opts ...option.RequestOption) (r *IPService) {
	r = &IPService{}
	r.Options = opts
	return
}

// Get IPs used on the Cloudflare/JD Cloud network, see
// https://www.cloudflare.com/ips for Cloudflare IPs or
// https://developers.cloudflare.com/china-network/reference/infrastructure/ for JD
// Cloud IPs.
func (r *IPService) List(ctx context.Context, query IPListParams, opts ...option.RequestOption) (res *IPListResponse, err error) {
	var env IPListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IPs []IPsItem

type IPsItem struct {
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// An IPv4 or IPv6 address.
	IP   string      `json:"ip"`
	JSON ipsItemJSON `json:"-"`
}

// ipsItemJSON contains the JSON metadata for the struct [IPsItem]
type ipsItemJSON struct {
	CreatedAt   apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPsItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipsItemJSON) RawJSON() string {
	return r.raw
}

type IPListResponse struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// This field can have the runtime type of [[]string].
	IPV4CIDRs interface{} `json:"ipv4_cidrs"`
	// This field can have the runtime type of [[]string].
	IPV6CIDRs interface{} `json:"ipv6_cidrs"`
	// This field can have the runtime type of [[]string].
	JDCloudCIDRs interface{}        `json:"jdcloud_cidrs"`
	JSON         ipListResponseJSON `json:"-"`
	union        IPListResponseUnion
}

// ipListResponseJSON contains the JSON metadata for the struct [IPListResponse]
type ipListResponseJSON struct {
	Etag         apijson.Field
	IPV4CIDRs    apijson.Field
	IPV6CIDRs    apijson.Field
	JDCloudCIDRs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r ipListResponseJSON) RawJSON() string {
	return r.raw
}

func (r *IPListResponse) UnmarshalJSON(data []byte) (err error) {
	*r = IPListResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [IPListResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [IPListResponsePublicIPIPs],
// [IPListResponsePublicIPIPsJDCloud].
func (r IPListResponse) AsUnion() IPListResponseUnion {
	return r.union
}

// Union satisfied by [IPListResponsePublicIPIPs] or
// [IPListResponsePublicIPIPsJDCloud].
type IPListResponseUnion interface {
	implementsIPListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*IPListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListResponsePublicIPIPs{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(IPListResponsePublicIPIPsJDCloud{}),
		},
	)
}

type IPListResponsePublicIPIPs struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	IPV4CIDRs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6CIDRs []string                      `json:"ipv6_cidrs"`
	JSON      ipListResponsePublicIpiPsJSON `json:"-"`
}

// ipListResponsePublicIpiPsJSON contains the JSON metadata for the struct
// [IPListResponsePublicIPIPs]
type ipListResponsePublicIpiPsJSON struct {
	Etag        apijson.Field
	IPV4CIDRs   apijson.Field
	IPV6CIDRs   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponsePublicIPIPs) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponsePublicIpiPsJSON) RawJSON() string {
	return r.raw
}

func (r IPListResponsePublicIPIPs) implementsIPListResponse() {}

type IPListResponsePublicIPIPsJDCloud struct {
	// A digest of the IP data. Useful for determining if the data has changed.
	Etag string `json:"etag"`
	// List of Cloudflare IPv4 CIDR addresses.
	IPV4CIDRs []string `json:"ipv4_cidrs"`
	// List of Cloudflare IPv6 CIDR addresses.
	IPV6CIDRs []string `json:"ipv6_cidrs"`
	// List IPv4 and IPv6 CIDRs, only populated if `?networks=jdcloud` is used.
	JDCloudCIDRs []string                             `json:"jdcloud_cidrs"`
	JSON         ipListResponsePublicIpiPsJDCloudJSON `json:"-"`
}

// ipListResponsePublicIpiPsJDCloudJSON contains the JSON metadata for the struct
// [IPListResponsePublicIPIPsJDCloud]
type ipListResponsePublicIpiPsJDCloudJSON struct {
	Etag         apijson.Field
	IPV4CIDRs    apijson.Field
	IPV6CIDRs    apijson.Field
	JDCloudCIDRs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IPListResponsePublicIPIPsJDCloud) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponsePublicIpiPsJDCloudJSON) RawJSON() string {
	return r.raw
}

func (r IPListResponsePublicIPIPsJDCloud) implementsIPListResponse() {}

type IPListParams struct {
	// Specified as `jdcloud` to list IPs used by JD Cloud data centers.
	Networks param.Field[string] `query:"networks"`
}

// URLQuery serializes [IPListParams]'s query parameters as `url.Values`.
func (r IPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type IPListResponseEnvelope struct {
	Errors   []IPListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IPListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success IPListResponseEnvelopeSuccess `json:"success,required"`
	Result  IPListResponse                `json:"result"`
	JSON    ipListResponseEnvelopeJSON    `json:"-"`
}

// ipListResponseEnvelopeJSON contains the JSON metadata for the struct
// [IPListResponseEnvelope]
type ipListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IPListResponseEnvelopeErrors struct {
	Code             int64                              `json:"code,required"`
	Message          string                             `json:"message,required"`
	DocumentationURL string                             `json:"documentation_url"`
	Source           IPListResponseEnvelopeErrorsSource `json:"source"`
	JSON             ipListResponseEnvelopeErrorsJSON   `json:"-"`
}

// ipListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [IPListResponseEnvelopeErrors]
type ipListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IPListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IPListResponseEnvelopeErrorsSource struct {
	Pointer string                                 `json:"pointer"`
	JSON    ipListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ipListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the struct
// [IPListResponseEnvelopeErrorsSource]
type ipListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type IPListResponseEnvelopeMessages struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           IPListResponseEnvelopeMessagesSource `json:"source"`
	JSON             ipListResponseEnvelopeMessagesJSON   `json:"-"`
}

// ipListResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [IPListResponseEnvelopeMessages]
type ipListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *IPListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type IPListResponseEnvelopeMessagesSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    ipListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ipListResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [IPListResponseEnvelopeMessagesSource]
type ipListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IPListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ipListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type IPListResponseEnvelopeSuccess bool

const (
	IPListResponseEnvelopeSuccessTrue IPListResponseEnvelopeSuccess = true
)

func (r IPListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case IPListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
