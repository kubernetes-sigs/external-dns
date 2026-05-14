// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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

// ZoneTransferACLService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferACLService] method instead.
type ZoneTransferACLService struct {
	Options []option.RequestOption
}

// NewZoneTransferACLService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneTransferACLService(opts ...option.RequestOption) (r *ZoneTransferACLService) {
	r = &ZoneTransferACLService{}
	r.Options = opts
	return
}

// Create ACL.
func (r *ZoneTransferACLService) New(ctx context.Context, params ZoneTransferACLNewParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env ZoneTransferACLNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify ACL.
func (r *ZoneTransferACLService) Update(ctx context.Context, aclID string, params ZoneTransferACLUpdateParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env ZoneTransferACLUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", params.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List ACLs.
func (r *ZoneTransferACLService) List(ctx context.Context, query ZoneTransferACLListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ACL], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls", query.AccountID)
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

// List ACLs.
func (r *ZoneTransferACLService) ListAutoPaging(ctx context.Context, query ZoneTransferACLListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ACL] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete ACL.
func (r *ZoneTransferACLService) Delete(ctx context.Context, aclID string, body ZoneTransferACLDeleteParams, opts ...option.RequestOption) (res *ZoneTransferACLDeleteResponse, err error) {
	var env ZoneTransferACLDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", body.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get ACL.
func (r *ZoneTransferACLService) Get(ctx context.Context, aclID string, query ZoneTransferACLGetParams, opts ...option.RequestOption) (res *ACL, err error) {
	var env ZoneTransferACLGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if aclID == "" {
		err = errors.New("missing required acl_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/secondary_dns/acls/%s", query.AccountID, aclID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ACL struct {
	ID string `json:"id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange string `json:"ip_range,required"`
	// The name of the acl.
	Name string  `json:"name,required"`
	JSON aclJSON `json:"-"`
}

// aclJSON contains the JSON metadata for the struct [ACL]
type aclJSON struct {
	ID          apijson.Field
	IPRange     apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ACL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r aclJSON) RawJSON() string {
	return r.raw
}

type ACLParam struct {
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r ACLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferACLDeleteResponse struct {
	ID   string                            `json:"id"`
	JSON zoneTransferACLDeleteResponseJSON `json:"-"`
}

// zoneTransferACLDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneTransferACLDeleteResponse]
type zoneTransferACLDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Allowed IPv4/IPv6 address range of primary or secondary nameservers. This will
	// be applied for the entire account. The IP range is used to allow additional
	// NOTIFY IPs for secondary zones and IPs Cloudflare allows AXFR/IXFR requests from
	// for primary zones. CIDRs are limited to a maximum of /24 for IPv4 and /64 for
	// IPv6 respectively.
	IPRange param.Field[string] `json:"ip_range,required"`
	// The name of the acl.
	Name param.Field[string] `json:"name,required"`
}

func (r ZoneTransferACLNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferACLNewResponseEnvelope struct {
	Errors   []ZoneTransferACLNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferACLNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferACLNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ACL                                       `json:"result"`
	JSON    zoneTransferACLNewResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferACLNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneTransferACLNewResponseEnvelope]
type zoneTransferACLNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ZoneTransferACLNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferACLNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferACLNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneTransferACLNewResponseEnvelopeErrors]
type zoneTransferACLNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    zoneTransferACLNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferACLNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ZoneTransferACLNewResponseEnvelopeErrorsSource]
type zoneTransferACLNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ZoneTransferACLNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferACLNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferACLNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneTransferACLNewResponseEnvelopeMessages]
type zoneTransferACLNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    zoneTransferACLNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferACLNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ZoneTransferACLNewResponseEnvelopeMessagesSource]
type zoneTransferACLNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferACLNewResponseEnvelopeSuccess bool

const (
	ZoneTransferACLNewResponseEnvelopeSuccessTrue ZoneTransferACLNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferACLNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferACLNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferACLUpdateParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	ACL       ACLParam            `json:"acl,required"`
}

func (r ZoneTransferACLUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ACL)
}

type ZoneTransferACLUpdateResponseEnvelope struct {
	Errors   []ZoneTransferACLUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferACLUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferACLUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ACL                                          `json:"result"`
	JSON    zoneTransferACLUpdateResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferACLUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferACLUpdateResponseEnvelope]
type zoneTransferACLUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLUpdateResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ZoneTransferACLUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferACLUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferACLUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneTransferACLUpdateResponseEnvelopeErrors]
type zoneTransferACLUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    zoneTransferACLUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferACLUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ZoneTransferACLUpdateResponseEnvelopeErrorsSource]
type zoneTransferACLUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLUpdateResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ZoneTransferACLUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferACLUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferACLUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneTransferACLUpdateResponseEnvelopeMessages]
type zoneTransferACLUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    zoneTransferACLUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferACLUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneTransferACLUpdateResponseEnvelopeMessagesSource]
type zoneTransferACLUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferACLUpdateResponseEnvelopeSuccess bool

const (
	ZoneTransferACLUpdateResponseEnvelopeSuccessTrue ZoneTransferACLUpdateResponseEnvelopeSuccess = true
)

func (r ZoneTransferACLUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferACLUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferACLListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferACLDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferACLDeleteResponseEnvelope struct {
	Errors   []ZoneTransferACLDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferACLDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferACLDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferACLDeleteResponse                `json:"result"`
	JSON    zoneTransferACLDeleteResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferACLDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferACLDeleteResponseEnvelope]
type zoneTransferACLDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ZoneTransferACLDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferACLDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferACLDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneTransferACLDeleteResponseEnvelopeErrors]
type zoneTransferACLDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    zoneTransferACLDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferACLDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ZoneTransferACLDeleteResponseEnvelopeErrorsSource]
type zoneTransferACLDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ZoneTransferACLDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferACLDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferACLDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneTransferACLDeleteResponseEnvelopeMessages]
type zoneTransferACLDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    zoneTransferACLDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferACLDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneTransferACLDeleteResponseEnvelopeMessagesSource]
type zoneTransferACLDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferACLDeleteResponseEnvelopeSuccess bool

const (
	ZoneTransferACLDeleteResponseEnvelopeSuccessTrue ZoneTransferACLDeleteResponseEnvelopeSuccess = true
)

func (r ZoneTransferACLDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferACLDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferACLGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZoneTransferACLGetResponseEnvelope struct {
	Errors   []ZoneTransferACLGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneTransferACLGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneTransferACLGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ACL                                       `json:"result"`
	JSON    zoneTransferACLGetResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferACLGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneTransferACLGetResponseEnvelope]
type zoneTransferACLGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ZoneTransferACLGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneTransferACLGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneTransferACLGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneTransferACLGetResponseEnvelopeErrors]
type zoneTransferACLGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    zoneTransferACLGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneTransferACLGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ZoneTransferACLGetResponseEnvelopeErrorsSource]
type zoneTransferACLGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ZoneTransferACLGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneTransferACLGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneTransferACLGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneTransferACLGetResponseEnvelopeMessages]
type zoneTransferACLGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneTransferACLGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferACLGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    zoneTransferACLGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneTransferACLGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ZoneTransferACLGetResponseEnvelopeMessagesSource]
type zoneTransferACLGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferACLGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferACLGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneTransferACLGetResponseEnvelopeSuccess bool

const (
	ZoneTransferACLGetResponseEnvelopeSuccessTrue ZoneTransferACLGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferACLGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferACLGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
