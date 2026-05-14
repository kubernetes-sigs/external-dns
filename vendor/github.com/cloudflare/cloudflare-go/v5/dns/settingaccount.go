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
)

// SettingAccountService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingAccountService] method instead.
type SettingAccountService struct {
	Options []option.RequestOption
	Views   *SettingAccountViewService
}

// NewSettingAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingAccountService(opts ...option.RequestOption) (r *SettingAccountService) {
	r = &SettingAccountService{}
	r.Options = opts
	r.Views = NewSettingAccountViewService(opts...)
	return
}

// Update DNS settings for an account
func (r *SettingAccountService) Edit(ctx context.Context, params SettingAccountEditParams, opts ...option.RequestOption) (res *SettingAccountEditResponse, err error) {
	var env SettingAccountEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show DNS settings for an account
func (r *SettingAccountService) Get(ctx context.Context, query SettingAccountGetParams, opts ...option.RequestOption) (res *SettingAccountGetResponse, err error) {
	var env SettingAccountGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dns_settings", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingAccountEditResponse struct {
	ZoneDefaults SettingAccountEditResponseZoneDefaults `json:"zone_defaults"`
	JSON         settingAccountEditResponseJSON         `json:"-"`
}

// settingAccountEditResponseJSON contains the JSON metadata for the struct
// [SettingAccountEditResponse]
type settingAccountEditResponseJSON struct {
	ZoneDefaults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAccountEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountEditResponseZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS SettingAccountEditResponseZoneDefaultsInternalDNS `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers SettingAccountEditResponseZoneDefaultsNameservers `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA SettingAccountEditResponseZoneDefaultsSOA `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingAccountEditResponseZoneDefaultsZoneMode `json:"zone_mode"`
	JSON     settingAccountEditResponseZoneDefaultsJSON     `json:"-"`
}

// settingAccountEditResponseZoneDefaultsJSON contains the JSON metadata for the
// struct [SettingAccountEditResponseZoneDefaults]
type settingAccountEditResponseZoneDefaultsJSON struct {
	FlattenAllCNAMEs   apijson.Field
	FoundationDNS      apijson.Field
	InternalDNS        apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NSTTL              apijson.Field
	SecondaryOverrides apijson.Field
	SOA                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAccountEditResponseZoneDefaults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseZoneDefaultsJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type SettingAccountEditResponseZoneDefaultsInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                                                `json:"reference_zone_id"`
	JSON            settingAccountEditResponseZoneDefaultsInternalDNSJSON `json:"-"`
}

// settingAccountEditResponseZoneDefaultsInternalDNSJSON contains the JSON metadata
// for the struct [SettingAccountEditResponseZoneDefaultsInternalDNS]
type settingAccountEditResponseZoneDefaultsInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingAccountEditResponseZoneDefaultsInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseZoneDefaultsInternalDNSJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type SettingAccountEditResponseZoneDefaultsNameservers struct {
	// Nameserver type
	Type SettingAccountEditResponseZoneDefaultsNameserversType `json:"type,required"`
	JSON settingAccountEditResponseZoneDefaultsNameserversJSON `json:"-"`
}

// settingAccountEditResponseZoneDefaultsNameserversJSON contains the JSON metadata
// for the struct [SettingAccountEditResponseZoneDefaultsNameservers]
type settingAccountEditResponseZoneDefaultsNameserversJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountEditResponseZoneDefaultsNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseZoneDefaultsNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type SettingAccountEditResponseZoneDefaultsNameserversType string

const (
	SettingAccountEditResponseZoneDefaultsNameserversTypeCloudflareStandard       SettingAccountEditResponseZoneDefaultsNameserversType = "cloudflare.standard"
	SettingAccountEditResponseZoneDefaultsNameserversTypeCloudflareStandardRandom SettingAccountEditResponseZoneDefaultsNameserversType = "cloudflare.standard.random"
	SettingAccountEditResponseZoneDefaultsNameserversTypeCustomAccount            SettingAccountEditResponseZoneDefaultsNameserversType = "custom.account"
	SettingAccountEditResponseZoneDefaultsNameserversTypeCustomTenant             SettingAccountEditResponseZoneDefaultsNameserversType = "custom.tenant"
)

func (r SettingAccountEditResponseZoneDefaultsNameserversType) IsKnown() bool {
	switch r {
	case SettingAccountEditResponseZoneDefaultsNameserversTypeCloudflareStandard, SettingAccountEditResponseZoneDefaultsNameserversTypeCloudflareStandardRandom, SettingAccountEditResponseZoneDefaultsNameserversTypeCustomAccount, SettingAccountEditResponseZoneDefaultsNameserversTypeCustomTenant:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingAccountEditResponseZoneDefaultsSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL float64 `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	MNAME string `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME string `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL  float64                                       `json:"ttl,required"`
	JSON settingAccountEditResponseZoneDefaultsSOAJSON `json:"-"`
}

// settingAccountEditResponseZoneDefaultsSOAJSON contains the JSON metadata for the
// struct [SettingAccountEditResponseZoneDefaultsSOA]
type settingAccountEditResponseZoneDefaultsSOAJSON struct {
	Expire      apijson.Field
	MinTTL      apijson.Field
	MNAME       apijson.Field
	Refresh     apijson.Field
	Retry       apijson.Field
	RNAME       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountEditResponseZoneDefaultsSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseZoneDefaultsSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingAccountEditResponseZoneDefaultsZoneMode string

const (
	SettingAccountEditResponseZoneDefaultsZoneModeStandard SettingAccountEditResponseZoneDefaultsZoneMode = "standard"
	SettingAccountEditResponseZoneDefaultsZoneModeCDNOnly  SettingAccountEditResponseZoneDefaultsZoneMode = "cdn_only"
	SettingAccountEditResponseZoneDefaultsZoneModeDNSOnly  SettingAccountEditResponseZoneDefaultsZoneMode = "dns_only"
)

func (r SettingAccountEditResponseZoneDefaultsZoneMode) IsKnown() bool {
	switch r {
	case SettingAccountEditResponseZoneDefaultsZoneModeStandard, SettingAccountEditResponseZoneDefaultsZoneModeCDNOnly, SettingAccountEditResponseZoneDefaultsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingAccountGetResponse struct {
	ZoneDefaults SettingAccountGetResponseZoneDefaults `json:"zone_defaults"`
	JSON         settingAccountGetResponseJSON         `json:"-"`
}

// settingAccountGetResponseJSON contains the JSON metadata for the struct
// [SettingAccountGetResponse]
type settingAccountGetResponseJSON struct {
	ZoneDefaults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *SettingAccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseJSON) RawJSON() string {
	return r.raw
}

type SettingAccountGetResponseZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS SettingAccountGetResponseZoneDefaultsInternalDNS `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers SettingAccountGetResponseZoneDefaultsNameservers `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA SettingAccountGetResponseZoneDefaultsSOA `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingAccountGetResponseZoneDefaultsZoneMode `json:"zone_mode"`
	JSON     settingAccountGetResponseZoneDefaultsJSON     `json:"-"`
}

// settingAccountGetResponseZoneDefaultsJSON contains the JSON metadata for the
// struct [SettingAccountGetResponseZoneDefaults]
type settingAccountGetResponseZoneDefaultsJSON struct {
	FlattenAllCNAMEs   apijson.Field
	FoundationDNS      apijson.Field
	InternalDNS        apijson.Field
	MultiProvider      apijson.Field
	Nameservers        apijson.Field
	NSTTL              apijson.Field
	SecondaryOverrides apijson.Field
	SOA                apijson.Field
	ZoneMode           apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SettingAccountGetResponseZoneDefaults) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseZoneDefaultsJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type SettingAccountGetResponseZoneDefaultsInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                                               `json:"reference_zone_id"`
	JSON            settingAccountGetResponseZoneDefaultsInternalDNSJSON `json:"-"`
}

// settingAccountGetResponseZoneDefaultsInternalDNSJSON contains the JSON metadata
// for the struct [SettingAccountGetResponseZoneDefaultsInternalDNS]
type settingAccountGetResponseZoneDefaultsInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingAccountGetResponseZoneDefaultsInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseZoneDefaultsInternalDNSJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type SettingAccountGetResponseZoneDefaultsNameservers struct {
	// Nameserver type
	Type SettingAccountGetResponseZoneDefaultsNameserversType `json:"type,required"`
	JSON settingAccountGetResponseZoneDefaultsNameserversJSON `json:"-"`
}

// settingAccountGetResponseZoneDefaultsNameserversJSON contains the JSON metadata
// for the struct [SettingAccountGetResponseZoneDefaultsNameservers]
type settingAccountGetResponseZoneDefaultsNameserversJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountGetResponseZoneDefaultsNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseZoneDefaultsNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type SettingAccountGetResponseZoneDefaultsNameserversType string

const (
	SettingAccountGetResponseZoneDefaultsNameserversTypeCloudflareStandard       SettingAccountGetResponseZoneDefaultsNameserversType = "cloudflare.standard"
	SettingAccountGetResponseZoneDefaultsNameserversTypeCloudflareStandardRandom SettingAccountGetResponseZoneDefaultsNameserversType = "cloudflare.standard.random"
	SettingAccountGetResponseZoneDefaultsNameserversTypeCustomAccount            SettingAccountGetResponseZoneDefaultsNameserversType = "custom.account"
	SettingAccountGetResponseZoneDefaultsNameserversTypeCustomTenant             SettingAccountGetResponseZoneDefaultsNameserversType = "custom.tenant"
)

func (r SettingAccountGetResponseZoneDefaultsNameserversType) IsKnown() bool {
	switch r {
	case SettingAccountGetResponseZoneDefaultsNameserversTypeCloudflareStandard, SettingAccountGetResponseZoneDefaultsNameserversTypeCloudflareStandardRandom, SettingAccountGetResponseZoneDefaultsNameserversTypeCustomAccount, SettingAccountGetResponseZoneDefaultsNameserversTypeCustomTenant:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingAccountGetResponseZoneDefaultsSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire float64 `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL float64 `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	MNAME string `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh float64 `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry float64 `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME string `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL  float64                                      `json:"ttl,required"`
	JSON settingAccountGetResponseZoneDefaultsSOAJSON `json:"-"`
}

// settingAccountGetResponseZoneDefaultsSOAJSON contains the JSON metadata for the
// struct [SettingAccountGetResponseZoneDefaultsSOA]
type settingAccountGetResponseZoneDefaultsSOAJSON struct {
	Expire      apijson.Field
	MinTTL      apijson.Field
	MNAME       apijson.Field
	Refresh     apijson.Field
	Retry       apijson.Field
	RNAME       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountGetResponseZoneDefaultsSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseZoneDefaultsSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingAccountGetResponseZoneDefaultsZoneMode string

const (
	SettingAccountGetResponseZoneDefaultsZoneModeStandard SettingAccountGetResponseZoneDefaultsZoneMode = "standard"
	SettingAccountGetResponseZoneDefaultsZoneModeCDNOnly  SettingAccountGetResponseZoneDefaultsZoneMode = "cdn_only"
	SettingAccountGetResponseZoneDefaultsZoneModeDNSOnly  SettingAccountGetResponseZoneDefaultsZoneMode = "dns_only"
)

func (r SettingAccountGetResponseZoneDefaultsZoneMode) IsKnown() bool {
	switch r {
	case SettingAccountGetResponseZoneDefaultsZoneModeStandard, SettingAccountGetResponseZoneDefaultsZoneModeCDNOnly, SettingAccountGetResponseZoneDefaultsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingAccountEditParams struct {
	// Identifier.
	AccountID    param.Field[string]                               `path:"account_id,required"`
	ZoneDefaults param.Field[SettingAccountEditParamsZoneDefaults] `json:"zone_defaults"`
}

func (r SettingAccountEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SettingAccountEditParamsZoneDefaults struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs param.Field[bool] `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS param.Field[bool] `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS param.Field[SettingAccountEditParamsZoneDefaultsInternalDNS] `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider param.Field[bool] `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[SettingAccountEditParamsZoneDefaultsNameservers] `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL param.Field[float64] `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA param.Field[SettingAccountEditParamsZoneDefaultsSOA] `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode param.Field[SettingAccountEditParamsZoneDefaultsZoneMode] `json:"zone_mode"`
}

func (r SettingAccountEditParamsZoneDefaults) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for this internal zone.
type SettingAccountEditParamsZoneDefaultsInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID param.Field[string] `json:"reference_zone_id"`
}

func (r SettingAccountEditParamsZoneDefaultsInternalDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings determining the nameservers through which the zone should be available.
type SettingAccountEditParamsZoneDefaultsNameservers struct {
	// Nameserver type
	Type param.Field[SettingAccountEditParamsZoneDefaultsNameserversType] `json:"type,required"`
}

func (r SettingAccountEditParamsZoneDefaultsNameservers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nameserver type
type SettingAccountEditParamsZoneDefaultsNameserversType string

const (
	SettingAccountEditParamsZoneDefaultsNameserversTypeCloudflareStandard       SettingAccountEditParamsZoneDefaultsNameserversType = "cloudflare.standard"
	SettingAccountEditParamsZoneDefaultsNameserversTypeCloudflareStandardRandom SettingAccountEditParamsZoneDefaultsNameserversType = "cloudflare.standard.random"
	SettingAccountEditParamsZoneDefaultsNameserversTypeCustomAccount            SettingAccountEditParamsZoneDefaultsNameserversType = "custom.account"
	SettingAccountEditParamsZoneDefaultsNameserversTypeCustomTenant             SettingAccountEditParamsZoneDefaultsNameserversType = "custom.tenant"
)

func (r SettingAccountEditParamsZoneDefaultsNameserversType) IsKnown() bool {
	switch r {
	case SettingAccountEditParamsZoneDefaultsNameserversTypeCloudflareStandard, SettingAccountEditParamsZoneDefaultsNameserversTypeCloudflareStandardRandom, SettingAccountEditParamsZoneDefaultsNameserversTypeCustomAccount, SettingAccountEditParamsZoneDefaultsNameserversTypeCustomTenant:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingAccountEditParamsZoneDefaultsSOA struct {
	// Time in seconds of being unable to query the primary server after which
	// secondary servers should stop serving the zone.
	Expire param.Field[float64] `json:"expire,required"`
	// The time to live (TTL) for negative caching of records within the zone.
	MinTTL param.Field[float64] `json:"min_ttl,required"`
	// The primary nameserver, which may be used for outbound zone transfers.
	MNAME param.Field[string] `json:"mname,required"`
	// Time in seconds after which secondary servers should re-check the SOA record to
	// see if the zone has been updated.
	Refresh param.Field[float64] `json:"refresh,required"`
	// Time in seconds after which secondary servers should retry queries after the
	// primary server was unresponsive.
	Retry param.Field[float64] `json:"retry,required"`
	// The email address of the zone administrator, with the first label representing
	// the local part of the email address.
	RNAME param.Field[string] `json:"rname,required"`
	// The time to live (TTL) of the SOA record itself.
	TTL param.Field[float64] `json:"ttl,required"`
}

func (r SettingAccountEditParamsZoneDefaultsSOA) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingAccountEditParamsZoneDefaultsZoneMode string

const (
	SettingAccountEditParamsZoneDefaultsZoneModeStandard SettingAccountEditParamsZoneDefaultsZoneMode = "standard"
	SettingAccountEditParamsZoneDefaultsZoneModeCDNOnly  SettingAccountEditParamsZoneDefaultsZoneMode = "cdn_only"
	SettingAccountEditParamsZoneDefaultsZoneModeDNSOnly  SettingAccountEditParamsZoneDefaultsZoneMode = "dns_only"
)

func (r SettingAccountEditParamsZoneDefaultsZoneMode) IsKnown() bool {
	switch r {
	case SettingAccountEditParamsZoneDefaultsZoneModeStandard, SettingAccountEditParamsZoneDefaultsZoneModeCDNOnly, SettingAccountEditParamsZoneDefaultsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingAccountEditResponseEnvelope struct {
	Errors   []SettingAccountEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAccountEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingAccountEditResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingAccountEditResponse                `json:"result"`
	JSON    settingAccountEditResponseEnvelopeJSON    `json:"-"`
}

// settingAccountEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingAccountEditResponseEnvelope]
type settingAccountEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountEditResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           SettingAccountEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingAccountEditResponseEnvelopeErrors]
type settingAccountEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountEditResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    settingAccountEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [SettingAccountEditResponseEnvelopeErrorsSource]
type settingAccountEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountEditResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           SettingAccountEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountEditResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [SettingAccountEditResponseEnvelopeMessages]
type settingAccountEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountEditResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    settingAccountEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingAccountEditResponseEnvelopeMessagesSource]
type settingAccountEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountEditResponseEnvelopeSuccess bool

const (
	SettingAccountEditResponseEnvelopeSuccessTrue SettingAccountEditResponseEnvelopeSuccess = true
)

func (r SettingAccountEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingAccountGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SettingAccountGetResponseEnvelope struct {
	Errors   []SettingAccountGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingAccountGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingAccountGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingAccountGetResponse                `json:"result"`
	JSON    settingAccountGetResponseEnvelopeJSON    `json:"-"`
}

// settingAccountGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingAccountGetResponseEnvelope]
type settingAccountGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingAccountGetResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           SettingAccountGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingAccountGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingAccountGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingAccountGetResponseEnvelopeErrors]
type settingAccountGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingAccountGetResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    settingAccountGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingAccountGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingAccountGetResponseEnvelopeErrorsSource]
type settingAccountGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingAccountGetResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           SettingAccountGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingAccountGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingAccountGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingAccountGetResponseEnvelopeMessages]
type settingAccountGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingAccountGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingAccountGetResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    settingAccountGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingAccountGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [SettingAccountGetResponseEnvelopeMessagesSource]
type settingAccountGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingAccountGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingAccountGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingAccountGetResponseEnvelopeSuccess bool

const (
	SettingAccountGetResponseEnvelopeSuccessTrue SettingAccountGetResponseEnvelopeSuccess = true
)

func (r SettingAccountGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingAccountGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
