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
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// PrefixBGPPrefixService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPrefixBGPPrefixService] method instead.
type PrefixBGPPrefixService struct {
	Options []option.RequestOption
}

// NewPrefixBGPPrefixService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrefixBGPPrefixService(opts ...option.RequestOption) (r *PrefixBGPPrefixService) {
	r = &PrefixBGPPrefixService{}
	r.Options = opts
	return
}

// Create a BGP prefix, controlling the BGP advertisement status of a specific
// subnet. When created, BGP prefixes are initially withdrawn, and can be
// advertised with the Update BGP Prefix API.
func (r *PrefixBGPPrefixService) New(ctx context.Context, prefixID string, params PrefixBGPPrefixNewParams, opts ...option.RequestOption) (res *BGPPrefix, err error) {
	var env PrefixBGPPrefixNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes", params.AccountID, prefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *PrefixBGPPrefixService) List(ctx context.Context, prefixID string, query PrefixBGPPrefixListParams, opts ...option.RequestOption) (res *pagination.SinglePage[BGPPrefix], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes", query.AccountID, prefixID)
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

// List all BGP Prefixes within the specified IP Prefix. BGP Prefixes are used to
// control which specific subnets are advertised to the Internet. It is possible to
// advertise subnets more specific than an IP Prefix by creating more specific BGP
// Prefixes.
func (r *PrefixBGPPrefixService) ListAutoPaging(ctx context.Context, prefixID string, query PrefixBGPPrefixListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[BGPPrefix] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, prefixID, query, opts...))
}

// Update the properties of a BGP Prefix, such as the on demand advertisement
// status (advertised or withdrawn).
func (r *PrefixBGPPrefixService) Edit(ctx context.Context, prefixID string, bgpPrefixID string, params PrefixBGPPrefixEditParams, opts ...option.RequestOption) (res *BGPPrefix, err error) {
	var env PrefixBGPPrefixEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bgpPrefixID == "" {
		err = errors.New("missing required bgp_prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", params.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieve a single BGP Prefix according to its identifier
func (r *PrefixBGPPrefixService) Get(ctx context.Context, prefixID string, bgpPrefixID string, query PrefixBGPPrefixGetParams, opts ...option.RequestOption) (res *BGPPrefix, err error) {
	var env PrefixBGPPrefixGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if prefixID == "" {
		err = errors.New("missing required prefix_id parameter")
		return
	}
	if bgpPrefixID == "" {
		err = errors.New("missing required bgp_prefix_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/addressing/prefixes/%s/bgp/prefixes/%s", query.AccountID, prefixID, bgpPrefixID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BGPPrefix struct {
	// Identifier of BGP Prefix.
	ID string `json:"id"`
	// Autonomous System Number (ASN) the prefix will be advertised under.
	ASN int64 `json:"asn,nullable"`
	// Number of times to prepend the Cloudflare ASN to the BGP AS-Path attribute
	ASNPrependCount int64 `json:"asn_prepend_count"`
	// Determines if Cloudflare advertises a BYOIP BGP prefix even when there is no
	// matching BGP prefix in the Magic routing table. When true, Cloudflare will
	// automatically withdraw the BGP prefix when there are no matching BGP routes, and
	// will resume advertising when there is at least one matching BGP route.
	AutoAdvertiseWithdraw bool                   `json:"auto_advertise_withdraw"`
	BGPSignalOpts         BGPPrefixBGPSignalOpts `json:"bgp_signal_opts"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR       string            `json:"cidr"`
	CreatedAt  time.Time         `json:"created_at" format:"date-time"`
	ModifiedAt time.Time         `json:"modified_at" format:"date-time"`
	OnDemand   BGPPrefixOnDemand `json:"on_demand"`
	JSON       bgpPrefixJSON     `json:"-"`
}

// bgpPrefixJSON contains the JSON metadata for the struct [BGPPrefix]
type bgpPrefixJSON struct {
	ID                    apijson.Field
	ASN                   apijson.Field
	ASNPrependCount       apijson.Field
	AutoAdvertiseWithdraw apijson.Field
	BGPSignalOpts         apijson.Field
	CIDR                  apijson.Field
	CreatedAt             apijson.Field
	ModifiedAt            apijson.Field
	OnDemand              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *BGPPrefix) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpPrefixJSON) RawJSON() string {
	return r.raw
}

type BGPPrefixBGPSignalOpts struct {
	// Whether control of advertisement of the prefix to the Internet is enabled to be
	// performed via BGP signal
	Enabled bool `json:"enabled"`
	// Last time BGP signaling control was toggled. This field is null if BGP signaling
	// has never been enabled.
	ModifiedAt time.Time                  `json:"modified_at,nullable" format:"date-time"`
	JSON       bgpPrefixBGPSignalOptsJSON `json:"-"`
}

// bgpPrefixBGPSignalOptsJSON contains the JSON metadata for the struct
// [BGPPrefixBGPSignalOpts]
type bgpPrefixBGPSignalOptsJSON struct {
	Enabled     apijson.Field
	ModifiedAt  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BGPPrefixBGPSignalOpts) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpPrefixBGPSignalOptsJSON) RawJSON() string {
	return r.raw
}

type BGPPrefixOnDemand struct {
	// Prefix advertisement status to the Internet. This field is only not 'null' if on
	// demand is enabled.
	Advertised bool `json:"advertised,nullable"`
	// Last time the advertisement status was changed. This field is only not 'null' if
	// on demand is enabled.
	AdvertisedModifiedAt time.Time `json:"advertised_modified_at,nullable" format:"date-time"`
	// Whether advertisement of the prefix to the Internet may be dynamically enabled
	// or disabled.
	OnDemandEnabled bool `json:"on_demand_enabled"`
	// Whether advertisement status of the prefix is locked, meaning it cannot be
	// changed.
	OnDemandLocked bool                  `json:"on_demand_locked"`
	JSON           bgpPrefixOnDemandJSON `json:"-"`
}

// bgpPrefixOnDemandJSON contains the JSON metadata for the struct
// [BGPPrefixOnDemand]
type bgpPrefixOnDemandJSON struct {
	Advertised           apijson.Field
	AdvertisedModifiedAt apijson.Field
	OnDemandEnabled      apijson.Field
	OnDemandLocked       apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *BGPPrefixOnDemand) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r bgpPrefixOnDemandJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixNewParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// IP Prefix in Classless Inter-Domain Routing format.
	CIDR param.Field[string] `json:"cidr"`
}

func (r PrefixBGPPrefixNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPPrefixNewResponseEnvelope struct {
	Errors   []PrefixBGPPrefixNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPPrefixNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success PrefixBGPPrefixNewResponseEnvelopeSuccess `json:"success,required"`
	Result  BGPPrefix                                 `json:"result"`
	JSON    prefixBGPPrefixNewResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPPrefixNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixBGPPrefixNewResponseEnvelope]
type prefixBGPPrefixNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           PrefixBGPPrefixNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             prefixBGPPrefixNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// prefixBGPPrefixNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixNewResponseEnvelopeErrors]
type prefixBGPPrefixNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixBGPPrefixNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    prefixBGPPrefixNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// prefixBGPPrefixNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [PrefixBGPPrefixNewResponseEnvelopeErrorsSource]
type prefixBGPPrefixNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           PrefixBGPPrefixNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             prefixBGPPrefixNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// prefixBGPPrefixNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPPrefixNewResponseEnvelopeMessages]
type prefixBGPPrefixNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixBGPPrefixNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    prefixBGPPrefixNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// prefixBGPPrefixNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [PrefixBGPPrefixNewResponseEnvelopeMessagesSource]
type prefixBGPPrefixNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrefixBGPPrefixNewResponseEnvelopeSuccess bool

const (
	PrefixBGPPrefixNewResponseEnvelopeSuccessTrue PrefixBGPPrefixNewResponseEnvelopeSuccess = true
)

func (r PrefixBGPPrefixNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPPrefixNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixBGPPrefixListParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPPrefixEditParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
	// Number of times to prepend the Cloudflare ASN to the BGP AS-Path attribute
	ASNPrependCount param.Field[int64] `json:"asn_prepend_count"`
	// Determines if Cloudflare advertises a BYOIP BGP prefix even when there is no
	// matching BGP prefix in the Magic routing table. When true, Cloudflare will
	// automatically withdraw the BGP prefix when there are no matching BGP routes, and
	// will resume advertising when there is at least one matching BGP route.
	AutoAdvertiseWithdraw param.Field[bool]                              `json:"auto_advertise_withdraw"`
	OnDemand              param.Field[PrefixBGPPrefixEditParamsOnDemand] `json:"on_demand"`
}

func (r PrefixBGPPrefixEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPPrefixEditParamsOnDemand struct {
	Advertised param.Field[bool] `json:"advertised"`
}

func (r PrefixBGPPrefixEditParamsOnDemand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrefixBGPPrefixEditResponseEnvelope struct {
	Errors   []PrefixBGPPrefixEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPPrefixEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success PrefixBGPPrefixEditResponseEnvelopeSuccess `json:"success,required"`
	Result  BGPPrefix                                  `json:"result"`
	JSON    prefixBGPPrefixEditResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixEditResponseEnvelope]
type prefixBGPPrefixEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           PrefixBGPPrefixEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             prefixBGPPrefixEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixEditResponseEnvelopeErrors]
type prefixBGPPrefixEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    prefixBGPPrefixEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [PrefixBGPPrefixEditResponseEnvelopeErrorsSource]
type prefixBGPPrefixEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           PrefixBGPPrefixEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             prefixBGPPrefixEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPPrefixEditResponseEnvelopeMessages]
type prefixBGPPrefixEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixEditResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    prefixBGPPrefixEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// prefixBGPPrefixEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [PrefixBGPPrefixEditResponseEnvelopeMessagesSource]
type prefixBGPPrefixEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrefixBGPPrefixEditResponseEnvelopeSuccess bool

const (
	PrefixBGPPrefixEditResponseEnvelopeSuccessTrue PrefixBGPPrefixEditResponseEnvelopeSuccess = true
)

func (r PrefixBGPPrefixEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPPrefixEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PrefixBGPPrefixGetParams struct {
	// Identifier of a Cloudflare account.
	AccountID param.Field[string] `path:"account_id,required"`
}

type PrefixBGPPrefixGetResponseEnvelope struct {
	Errors   []PrefixBGPPrefixGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PrefixBGPPrefixGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success PrefixBGPPrefixGetResponseEnvelopeSuccess `json:"success,required"`
	Result  BGPPrefix                                 `json:"result"`
	JSON    prefixBGPPrefixGetResponseEnvelopeJSON    `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PrefixBGPPrefixGetResponseEnvelope]
type prefixBGPPrefixGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           PrefixBGPPrefixGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             prefixBGPPrefixGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PrefixBGPPrefixGetResponseEnvelopeErrors]
type prefixBGPPrefixGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    prefixBGPPrefixGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [PrefixBGPPrefixGetResponseEnvelopeErrorsSource]
type prefixBGPPrefixGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           PrefixBGPPrefixGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             prefixBGPPrefixGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [PrefixBGPPrefixGetResponseEnvelopeMessages]
type prefixBGPPrefixGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type PrefixBGPPrefixGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    prefixBGPPrefixGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// prefixBGPPrefixGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [PrefixBGPPrefixGetResponseEnvelopeMessagesSource]
type prefixBGPPrefixGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrefixBGPPrefixGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r prefixBGPPrefixGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PrefixBGPPrefixGetResponseEnvelopeSuccess bool

const (
	PrefixBGPPrefixGetResponseEnvelopeSuccessTrue PrefixBGPPrefixGetResponseEnvelopeSuccess = true
)

func (r PrefixBGPPrefixGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PrefixBGPPrefixGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
