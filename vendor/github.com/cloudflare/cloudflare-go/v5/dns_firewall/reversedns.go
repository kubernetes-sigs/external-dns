// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// ReverseDNSService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReverseDNSService] method instead.
type ReverseDNSService struct {
	Options []option.RequestOption
}

// NewReverseDNSService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReverseDNSService(opts ...option.RequestOption) (r *ReverseDNSService) {
	r = &ReverseDNSService{}
	r.Options = opts
	return
}

// Update reverse DNS configuration (PTR records) for a DNS Firewall cluster
func (r *ReverseDNSService) Edit(ctx context.Context, dnsFirewallID string, params ReverseDNSEditParams, opts ...option.RequestOption) (res *ReverseDNSEditResponse, err error) {
	var env ReverseDNSEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/reverse_dns", params.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show reverse DNS configuration (PTR records) for a DNS Firewall cluster
func (r *ReverseDNSService) Get(ctx context.Context, dnsFirewallID string, query ReverseDNSGetParams, opts ...option.RequestOption) (res *ReverseDNSGetResponse, err error) {
	var env ReverseDNSGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/reverse_dns", query.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ReverseDNSEditResponse struct {
	// Map of cluster IP addresses to PTR record contents
	PTR  map[string]string          `json:"ptr,required"`
	JSON reverseDNSEditResponseJSON `json:"-"`
}

// reverseDNSEditResponseJSON contains the JSON metadata for the struct
// [ReverseDNSEditResponse]
type reverseDNSEditResponseJSON struct {
	PTR         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSGetResponse struct {
	// Map of cluster IP addresses to PTR record contents
	PTR  map[string]string         `json:"ptr,required"`
	JSON reverseDNSGetResponseJSON `json:"-"`
}

// reverseDNSGetResponseJSON contains the JSON metadata for the struct
// [ReverseDNSGetResponse]
type reverseDNSGetResponseJSON struct {
	PTR         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Map of cluster IP addresses to PTR record contents
	PTR param.Field[map[string]string] `json:"ptr"`
}

func (r ReverseDNSEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ReverseDNSEditResponseEnvelope struct {
	Errors   []ReverseDNSEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ReverseDNSEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ReverseDNSEditResponseEnvelopeSuccess `json:"success,required"`
	Result  ReverseDNSEditResponse                `json:"result"`
	JSON    reverseDNSEditResponseEnvelopeJSON    `json:"-"`
}

// reverseDNSEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ReverseDNSEditResponseEnvelope]
type reverseDNSEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSEditResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           ReverseDNSEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             reverseDNSEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// reverseDNSEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ReverseDNSEditResponseEnvelopeErrors]
type reverseDNSEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ReverseDNSEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSEditResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    reverseDNSEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// reverseDNSEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ReverseDNSEditResponseEnvelopeErrorsSource]
type reverseDNSEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSEditResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ReverseDNSEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             reverseDNSEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// reverseDNSEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ReverseDNSEditResponseEnvelopeMessages]
type reverseDNSEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ReverseDNSEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSEditResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    reverseDNSEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// reverseDNSEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ReverseDNSEditResponseEnvelopeMessagesSource]
type reverseDNSEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ReverseDNSEditResponseEnvelopeSuccess bool

const (
	ReverseDNSEditResponseEnvelopeSuccessTrue ReverseDNSEditResponseEnvelopeSuccess = true
)

func (r ReverseDNSEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ReverseDNSEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ReverseDNSGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ReverseDNSGetResponseEnvelope struct {
	Errors   []ReverseDNSGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ReverseDNSGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ReverseDNSGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ReverseDNSGetResponse                `json:"result"`
	JSON    reverseDNSGetResponseEnvelopeJSON    `json:"-"`
}

// reverseDNSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ReverseDNSGetResponseEnvelope]
type reverseDNSGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           ReverseDNSGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             reverseDNSGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// reverseDNSGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ReverseDNSGetResponseEnvelopeErrors]
type reverseDNSGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ReverseDNSGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    reverseDNSGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// reverseDNSGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [ReverseDNSGetResponseEnvelopeErrorsSource]
type reverseDNSGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           ReverseDNSGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             reverseDNSGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// reverseDNSGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ReverseDNSGetResponseEnvelopeMessages]
type reverseDNSGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ReverseDNSGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ReverseDNSGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    reverseDNSGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// reverseDNSGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [ReverseDNSGetResponseEnvelopeMessagesSource]
type reverseDNSGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseDNSGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseDNSGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ReverseDNSGetResponseEnvelopeSuccess bool

const (
	ReverseDNSGetResponseEnvelopeSuccessTrue ReverseDNSGetResponseEnvelopeSuccess = true
)

func (r ReverseDNSGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ReverseDNSGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
