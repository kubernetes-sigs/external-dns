// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_nameservers

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

// CustomNameserverService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomNameserverService] method instead.
type CustomNameserverService struct {
	Options []option.RequestOption
}

// NewCustomNameserverService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomNameserverService(opts ...option.RequestOption) (r *CustomNameserverService) {
	r = &CustomNameserverService{}
	r.Options = opts
	return
}

// Add Account Custom Nameserver
func (r *CustomNameserverService) New(ctx context.Context, params CustomNameserverNewParams, opts ...option.RequestOption) (res *CustomNameserver, err error) {
	var env CustomNameserverNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Account Custom Nameserver
func (r *CustomNameserverService) Delete(ctx context.Context, customNSID string, body CustomNameserverDeleteParams, opts ...option.RequestOption) (res *pagination.SinglePage[string], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if customNSID == "" {
		err = errors.New("missing required custom_ns_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns/%s", body.AccountID, customNSID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodDelete, path, nil, &res, opts...)
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

// Delete Account Custom Nameserver
func (r *CustomNameserverService) DeleteAutoPaging(ctx context.Context, customNSID string, body CustomNameserverDeleteParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[string] {
	return pagination.NewSinglePageAutoPager(r.Delete(ctx, customNSID, body, opts...))
}

// List an account's custom nameservers.
func (r *CustomNameserverService) Get(ctx context.Context, query CustomNameserverGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[CustomNameserver], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/custom_ns", query.AccountID)
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

// List an account's custom nameservers.
func (r *CustomNameserverService) GetAutoPaging(ctx context.Context, query CustomNameserverGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[CustomNameserver] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

// A single account custom nameserver.
type CustomNameserver struct {
	// A and AAAA records associated with the nameserver.
	DNSRecords []CustomNameserverDNSRecord `json:"dns_records,required"`
	// The FQDN of the name server.
	NSName string `json:"ns_name,required" format:"hostname"`
	// Verification status of the nameserver.
	//
	// Deprecated: deprecated
	Status CustomNameserverStatus `json:"status,required"`
	// Identifier.
	ZoneTag string `json:"zone_tag,required"`
	// The number of the set that this name server belongs to.
	NSSet float64              `json:"ns_set"`
	JSON  customNameserverJSON `json:"-"`
}

// customNameserverJSON contains the JSON metadata for the struct
// [CustomNameserver]
type customNameserverJSON struct {
	DNSRecords  apijson.Field
	NSName      apijson.Field
	Status      apijson.Field
	ZoneTag     apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserver) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverDNSRecord struct {
	// DNS record type.
	Type CustomNameserverDNSRecordsType `json:"type"`
	// DNS record contents (an IPv4 or IPv6 address).
	Value string                        `json:"value"`
	JSON  customNameserverDNSRecordJSON `json:"-"`
}

// customNameserverDNSRecordJSON contains the JSON metadata for the struct
// [CustomNameserverDNSRecord]
type customNameserverDNSRecordJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverDNSRecord) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverDNSRecordJSON) RawJSON() string {
	return r.raw
}

// DNS record type.
type CustomNameserverDNSRecordsType string

const (
	CustomNameserverDNSRecordsTypeA    CustomNameserverDNSRecordsType = "A"
	CustomNameserverDNSRecordsTypeAAAA CustomNameserverDNSRecordsType = "AAAA"
)

func (r CustomNameserverDNSRecordsType) IsKnown() bool {
	switch r {
	case CustomNameserverDNSRecordsTypeA, CustomNameserverDNSRecordsTypeAAAA:
		return true
	}
	return false
}

// Verification status of the nameserver.
type CustomNameserverStatus string

const (
	CustomNameserverStatusMoved    CustomNameserverStatus = "moved"
	CustomNameserverStatusPending  CustomNameserverStatus = "pending"
	CustomNameserverStatusVerified CustomNameserverStatus = "verified"
)

func (r CustomNameserverStatus) IsKnown() bool {
	switch r {
	case CustomNameserverStatusMoved, CustomNameserverStatusPending, CustomNameserverStatusVerified:
		return true
	}
	return false
}

type CustomNameserverNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// The FQDN of the name server.
	NSName param.Field[string] `json:"ns_name,required" format:"hostname"`
	// The number of the set that this name server belongs to.
	NSSet param.Field[float64] `json:"ns_set"`
}

func (r CustomNameserverNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CustomNameserverNewResponseEnvelope struct {
	Errors   []CustomNameserverNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []CustomNameserverNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success CustomNameserverNewResponseEnvelopeSuccess `json:"success,required"`
	// A single account custom nameserver.
	Result CustomNameserver                        `json:"result"`
	JSON   customNameserverNewResponseEnvelopeJSON `json:"-"`
}

// customNameserverNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [CustomNameserverNewResponseEnvelope]
type customNameserverNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverNewResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           CustomNameserverNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             customNameserverNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// customNameserverNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [CustomNameserverNewResponseEnvelopeErrors]
type customNameserverNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverNewResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    customNameserverNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// customNameserverNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [CustomNameserverNewResponseEnvelopeErrorsSource]
type customNameserverNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverNewResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           CustomNameserverNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             customNameserverNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// customNameserverNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [CustomNameserverNewResponseEnvelopeMessages]
type customNameserverNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type CustomNameserverNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    customNameserverNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// customNameserverNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [CustomNameserverNewResponseEnvelopeMessagesSource]
type customNameserverNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomNameserverNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r customNameserverNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type CustomNameserverNewResponseEnvelopeSuccess bool

const (
	CustomNameserverNewResponseEnvelopeSuccessTrue CustomNameserverNewResponseEnvelopeSuccess = true
)

func (r CustomNameserverNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case CustomNameserverNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type CustomNameserverDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type CustomNameserverGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}
