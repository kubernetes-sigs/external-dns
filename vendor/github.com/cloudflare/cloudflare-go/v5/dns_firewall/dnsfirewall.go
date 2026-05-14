// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// DNSFirewallService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSFirewallService] method instead.
type DNSFirewallService struct {
	Options    []option.RequestOption
	Analytics  *AnalyticsService
	ReverseDNS *ReverseDNSService
}

// NewDNSFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSFirewallService(opts ...option.RequestOption) (r *DNSFirewallService) {
	r = &DNSFirewallService{}
	r.Options = opts
	r.Analytics = NewAnalyticsService(opts...)
	r.ReverseDNS = NewReverseDNSService(opts...)
	return
}

// Create a DNS Firewall cluster
func (r *DNSFirewallService) New(ctx context.Context, params DNSFirewallNewParams, opts ...option.RequestOption) (res *DNSFirewallNewResponse, err error) {
	var env DNSFirewallNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List DNS Firewall clusters for an account
func (r *DNSFirewallService) List(ctx context.Context, params DNSFirewallListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DNSFirewallListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall", params.AccountID)
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

// List DNS Firewall clusters for an account
func (r *DNSFirewallService) ListAutoPaging(ctx context.Context, params DNSFirewallListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DNSFirewallListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a DNS Firewall cluster
func (r *DNSFirewallService) Delete(ctx context.Context, dnsFirewallID string, body DNSFirewallDeleteParams, opts ...option.RequestOption) (res *DNSFirewallDeleteResponse, err error) {
	var env DNSFirewallDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", body.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify the configuration of a DNS Firewall cluster
func (r *DNSFirewallService) Edit(ctx context.Context, dnsFirewallID string, params DNSFirewallEditParams, opts ...option.RequestOption) (res *DNSFirewallEditResponse, err error) {
	var env DNSFirewallEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", params.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show a single DNS Firewall cluster for an account
func (r *DNSFirewallService) Get(ctx context.Context, dnsFirewallID string, query DNSFirewallGetParams, opts ...option.RequestOption) (res *DNSFirewallGetResponse, err error) {
	var env DNSFirewallGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if dnsFirewallID == "" {
		err = errors.New("missing required dns_firewall_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s", query.AccountID, dnsFirewallID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Attack mitigation settings
type AttackMitigation struct {
	// When enabled, automatically mitigate random-prefix attacks to protect upstream
	// DNS servers
	Enabled bool `json:"enabled"`
	// Only mitigate attacks when upstream servers seem unhealthy
	OnlyWhenUpstreamUnhealthy bool                 `json:"only_when_upstream_unhealthy"`
	JSON                      attackMitigationJSON `json:"-"`
}

// attackMitigationJSON contains the JSON metadata for the struct
// [AttackMitigation]
type attackMitigationJSON struct {
	Enabled                   apijson.Field
	OnlyWhenUpstreamUnhealthy apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *AttackMitigation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackMitigationJSON) RawJSON() string {
	return r.raw
}

// Attack mitigation settings
type AttackMitigationParam struct {
	// When enabled, automatically mitigate random-prefix attacks to protect upstream
	// DNS servers
	Enabled param.Field[bool] `json:"enabled"`
	// Only mitigate attacks when upstream servers seem unhealthy
	OnlyWhenUpstreamUnhealthy param.Field[bool] `json:"only_when_upstream_unhealthy"`
}

func (r AttackMitigationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallIPs = string

type UpstreamIPs = string

type UpstreamIPsParam = string

type DNSFirewallNewResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall cluster name
	Name string `json:"name,required"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings
	AttackMitigation AttackMitigation           `json:"attack_mitigation,nullable"`
	JSON             dnsFirewallNewResponseJSON `json:"-"`
}

// dnsFirewallNewResponseJSON contains the JSON metadata for the struct
// [DNSFirewallNewResponse]
type dnsFirewallNewResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DNSFirewallNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallNewResponseJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallListResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall cluster name
	Name string `json:"name,required"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings
	AttackMitigation AttackMitigation            `json:"attack_mitigation,nullable"`
	JSON             dnsFirewallListResponseJSON `json:"-"`
}

// dnsFirewallListResponseJSON contains the JSON metadata for the struct
// [DNSFirewallListResponse]
type dnsFirewallListResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DNSFirewallListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallListResponseJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallDeleteResponse struct {
	// Identifier.
	ID   string                        `json:"id"`
	JSON dnsFirewallDeleteResponseJSON `json:"-"`
}

// dnsFirewallDeleteResponseJSON contains the JSON metadata for the struct
// [DNSFirewallDeleteResponse]
type dnsFirewallDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallEditResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall cluster name
	Name string `json:"name,required"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings
	AttackMitigation AttackMitigation            `json:"attack_mitigation,nullable"`
	JSON             dnsFirewallEditResponseJSON `json:"-"`
}

// dnsFirewallEditResponseJSON contains the JSON metadata for the struct
// [DNSFirewallEditResponse]
type dnsFirewallEditResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DNSFirewallEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallEditResponseJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallGetResponse struct {
	// Identifier.
	ID string `json:"id,required"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests bool          `json:"deprecate_any_requests,required"`
	DNSFirewallIPs       []FirewallIPs `json:"dns_firewall_ips,required" format:"ipv4"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	ECSFallback bool `json:"ecs_fallback,required"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL float64 `json:"maximum_cache_ttl,required"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL float64 `json:"minimum_cache_ttl,required"`
	// Last modification of DNS Firewall cluster
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// DNS Firewall cluster name
	Name string `json:"name,required"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL float64 `json:"negative_cache_ttl,required,nullable"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit float64 `json:"ratelimit,required,nullable"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries     float64       `json:"retries,required"`
	UpstreamIPs []UpstreamIPs `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings
	AttackMitigation AttackMitigation           `json:"attack_mitigation,nullable"`
	JSON             dnsFirewallGetResponseJSON `json:"-"`
}

// dnsFirewallGetResponseJSON contains the JSON metadata for the struct
// [DNSFirewallGetResponse]
type dnsFirewallGetResponseJSON struct {
	ID                   apijson.Field
	DeprecateAnyRequests apijson.Field
	DNSFirewallIPs       apijson.Field
	ECSFallback          apijson.Field
	MaximumCacheTTL      apijson.Field
	MinimumCacheTTL      apijson.Field
	ModifiedOn           apijson.Field
	Name                 apijson.Field
	NegativeCacheTTL     apijson.Field
	Ratelimit            apijson.Field
	Retries              apijson.Field
	UpstreamIPs          apijson.Field
	AttackMitigation     apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DNSFirewallGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallGetResponseJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// DNS Firewall cluster name
	Name        param.Field[string]             `json:"name,required"`
	UpstreamIPs param.Field[[]UpstreamIPsParam] `json:"upstream_ips,required" format:"ipv4"`
	// Attack mitigation settings
	AttackMitigation param.Field[AttackMitigationParam] `json:"attack_mitigation"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests param.Field[bool] `json:"deprecate_any_requests"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	ECSFallback param.Field[bool] `json:"ecs_fallback"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL param.Field[float64] `json:"maximum_cache_ttl"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL param.Field[float64] `json:"minimum_cache_ttl"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL param.Field[float64] `json:"negative_cache_ttl"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries param.Field[float64] `json:"retries"`
}

func (r DNSFirewallNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSFirewallNewResponseEnvelope struct {
	Errors   []DNSFirewallNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DNSFirewallNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSFirewallNewResponse                `json:"result"`
	JSON    dnsFirewallNewResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallNewResponseEnvelope]
type dnsFirewallNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallNewResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DNSFirewallNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsFirewallNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsFirewallNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallNewResponseEnvelopeErrors]
type dnsFirewallNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallNewResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dnsFirewallNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsFirewallNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DNSFirewallNewResponseEnvelopeErrorsSource]
type dnsFirewallNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallNewResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DNSFirewallNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsFirewallNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsFirewallNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallNewResponseEnvelopeMessages]
type dnsFirewallNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallNewResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dnsFirewallNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsFirewallNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DNSFirewallNewResponseEnvelopeMessagesSource]
type dnsFirewallNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSFirewallNewResponseEnvelopeSuccess bool

const (
	DNSFirewallNewResponseEnvelopeSuccessTrue DNSFirewallNewResponseEnvelopeSuccess = true
)

func (r DNSFirewallNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSFirewallNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSFirewallListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page"`
	// Number of clusters per page
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [DNSFirewallListParams]'s query parameters as `url.Values`.
func (r DNSFirewallListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DNSFirewallDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DNSFirewallDeleteResponseEnvelope struct {
	Errors   []DNSFirewallDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DNSFirewallDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSFirewallDeleteResponse                `json:"result"`
	JSON    dnsFirewallDeleteResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallDeleteResponseEnvelope]
type dnsFirewallDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallDeleteResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DNSFirewallDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsFirewallDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallDeleteResponseEnvelopeErrors]
type dnsFirewallDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dnsFirewallDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DNSFirewallDeleteResponseEnvelopeErrorsSource]
type dnsFirewallDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallDeleteResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           DNSFirewallDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsFirewallDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallDeleteResponseEnvelopeMessages]
type dnsFirewallDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    dnsFirewallDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsFirewallDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DNSFirewallDeleteResponseEnvelopeMessagesSource]
type dnsFirewallDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSFirewallDeleteResponseEnvelopeSuccess bool

const (
	DNSFirewallDeleteResponseEnvelopeSuccessTrue DNSFirewallDeleteResponseEnvelopeSuccess = true
)

func (r DNSFirewallDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSFirewallDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSFirewallEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// Attack mitigation settings
	AttackMitigation param.Field[AttackMitigationParam] `json:"attack_mitigation"`
	// Whether to refuse to answer queries for the ANY type
	DeprecateAnyRequests param.Field[bool] `json:"deprecate_any_requests"`
	// Whether to forward client IP (resolver) subnet if no EDNS Client Subnet is sent
	ECSFallback param.Field[bool] `json:"ecs_fallback"`
	// Maximum DNS cache TTL This setting sets an upper bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Higher TTLs will be
	// decreased to the maximum defined here for caching purposes.
	MaximumCacheTTL param.Field[float64] `json:"maximum_cache_ttl"`
	// Minimum DNS cache TTL This setting sets a lower bound on DNS TTLs for purposes
	// of caching between DNS Firewall and the upstream servers. Lower TTLs will be
	// increased to the minimum defined here for caching purposes.
	MinimumCacheTTL param.Field[float64] `json:"minimum_cache_ttl"`
	// DNS Firewall cluster name
	Name param.Field[string] `json:"name"`
	// Negative DNS cache TTL This setting controls how long DNS Firewall should cache
	// negative responses (e.g., NXDOMAIN) from the upstream servers.
	NegativeCacheTTL param.Field[float64] `json:"negative_cache_ttl"`
	// Ratelimit in queries per second per datacenter (applies to DNS queries sent to
	// the upstream nameservers configured on the cluster)
	Ratelimit param.Field[float64] `json:"ratelimit"`
	// Number of retries for fetching DNS responses from upstream nameservers (not
	// counting the initial attempt)
	Retries     param.Field[float64]            `json:"retries"`
	UpstreamIPs param.Field[[]UpstreamIPsParam] `json:"upstream_ips" format:"ipv4"`
}

func (r DNSFirewallEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DNSFirewallEditResponseEnvelope struct {
	Errors   []DNSFirewallEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DNSFirewallEditResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSFirewallEditResponse                `json:"result"`
	JSON    dnsFirewallEditResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallEditResponseEnvelope]
type dnsFirewallEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallEditResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DNSFirewallEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsFirewallEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsFirewallEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallEditResponseEnvelopeErrors]
type dnsFirewallEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallEditResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dnsFirewallEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsFirewallEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DNSFirewallEditResponseEnvelopeErrorsSource]
type dnsFirewallEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallEditResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DNSFirewallEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsFirewallEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsFirewallEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallEditResponseEnvelopeMessages]
type dnsFirewallEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallEditResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dnsFirewallEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsFirewallEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DNSFirewallEditResponseEnvelopeMessagesSource]
type dnsFirewallEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSFirewallEditResponseEnvelopeSuccess bool

const (
	DNSFirewallEditResponseEnvelopeSuccessTrue DNSFirewallEditResponseEnvelopeSuccess = true
)

func (r DNSFirewallEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSFirewallEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DNSFirewallGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type DNSFirewallGetResponseEnvelope struct {
	Errors   []DNSFirewallGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DNSFirewallGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DNSFirewallGetResponse                `json:"result"`
	JSON    dnsFirewallGetResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DNSFirewallGetResponseEnvelope]
type dnsFirewallGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           DNSFirewallGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dnsFirewallGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dnsFirewallGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DNSFirewallGetResponseEnvelopeErrors]
type dnsFirewallGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    dnsFirewallGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dnsFirewallGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DNSFirewallGetResponseEnvelopeErrorsSource]
type dnsFirewallGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           DNSFirewallGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dnsFirewallGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dnsFirewallGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DNSFirewallGetResponseEnvelopeMessages]
type dnsFirewallGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DNSFirewallGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    dnsFirewallGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dnsFirewallGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DNSFirewallGetResponseEnvelopeMessagesSource]
type dnsFirewallGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsFirewallGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DNSFirewallGetResponseEnvelopeSuccess bool

const (
	DNSFirewallGetResponseEnvelopeSuccessTrue DNSFirewallGetResponseEnvelopeSuccess = true
)

func (r DNSFirewallGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DNSFirewallGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
