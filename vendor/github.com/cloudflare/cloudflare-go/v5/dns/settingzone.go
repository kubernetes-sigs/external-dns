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

// SettingZoneService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSettingZoneService] method instead.
type SettingZoneService struct {
	Options []option.RequestOption
}

// NewSettingZoneService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSettingZoneService(opts ...option.RequestOption) (r *SettingZoneService) {
	r = &SettingZoneService{}
	r.Options = opts
	return
}

// Update DNS settings for a zone
func (r *SettingZoneService) Edit(ctx context.Context, params SettingZoneEditParams, opts ...option.RequestOption) (res *SettingZoneEditResponse, err error) {
	var env SettingZoneEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Show DNS settings for a zone
func (r *SettingZoneService) Get(ctx context.Context, query SettingZoneGetParams, opts ...option.RequestOption) (res *SettingZoneGetResponse, err error) {
	var env SettingZoneGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SettingZoneEditResponse struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS SettingZoneEditResponseInternalDNS `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers SettingZoneEditResponseNameservers `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA SettingZoneEditResponseSOA `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingZoneEditResponseZoneMode `json:"zone_mode"`
	JSON     settingZoneEditResponseJSON     `json:"-"`
}

// settingZoneEditResponseJSON contains the JSON metadata for the struct
// [SettingZoneEditResponse]
type settingZoneEditResponseJSON struct {
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

func (r *SettingZoneEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type SettingZoneEditResponseInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                                 `json:"reference_zone_id"`
	JSON            settingZoneEditResponseInternalDNSJSON `json:"-"`
}

// settingZoneEditResponseInternalDNSJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseInternalDNS]
type settingZoneEditResponseInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingZoneEditResponseInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseInternalDNSJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type SettingZoneEditResponseNameservers struct {
	// Nameserver type
	Type SettingZoneEditResponseNameserversType `json:"type,required"`
	// Configured nameserver set to be used for this zone
	NSSet int64                                  `json:"ns_set"`
	JSON  settingZoneEditResponseNameserversJSON `json:"-"`
}

// settingZoneEditResponseNameserversJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseNameservers]
type settingZoneEditResponseNameserversJSON struct {
	Type        apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type SettingZoneEditResponseNameserversType string

const (
	SettingZoneEditResponseNameserversTypeCloudflareStandard SettingZoneEditResponseNameserversType = "cloudflare.standard"
	SettingZoneEditResponseNameserversTypeCustomAccount      SettingZoneEditResponseNameserversType = "custom.account"
	SettingZoneEditResponseNameserversTypeCustomTenant       SettingZoneEditResponseNameserversType = "custom.tenant"
	SettingZoneEditResponseNameserversTypeCustomZone         SettingZoneEditResponseNameserversType = "custom.zone"
)

func (r SettingZoneEditResponseNameserversType) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseNameserversTypeCloudflareStandard, SettingZoneEditResponseNameserversTypeCustomAccount, SettingZoneEditResponseNameserversTypeCustomTenant, SettingZoneEditResponseNameserversTypeCustomZone:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingZoneEditResponseSOA struct {
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
	TTL  float64                        `json:"ttl,required"`
	JSON settingZoneEditResponseSOAJSON `json:"-"`
}

// settingZoneEditResponseSOAJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseSOA]
type settingZoneEditResponseSOAJSON struct {
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

func (r *SettingZoneEditResponseSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingZoneEditResponseZoneMode string

const (
	SettingZoneEditResponseZoneModeStandard SettingZoneEditResponseZoneMode = "standard"
	SettingZoneEditResponseZoneModeCDNOnly  SettingZoneEditResponseZoneMode = "cdn_only"
	SettingZoneEditResponseZoneModeDNSOnly  SettingZoneEditResponseZoneMode = "dns_only"
)

func (r SettingZoneEditResponseZoneMode) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseZoneModeStandard, SettingZoneEditResponseZoneModeCDNOnly, SettingZoneEditResponseZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingZoneGetResponse struct {
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs bool `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS bool `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS SettingZoneGetResponseInternalDNS `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider bool `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers SettingZoneGetResponseNameservers `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL float64 `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides bool `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA SettingZoneGetResponseSOA `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode SettingZoneGetResponseZoneMode `json:"zone_mode"`
	JSON     settingZoneGetResponseJSON     `json:"-"`
}

// settingZoneGetResponseJSON contains the JSON metadata for the struct
// [SettingZoneGetResponse]
type settingZoneGetResponseJSON struct {
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

func (r *SettingZoneGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseJSON) RawJSON() string {
	return r.raw
}

// Settings for this internal zone.
type SettingZoneGetResponseInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID string                                `json:"reference_zone_id"`
	JSON            settingZoneGetResponseInternalDNSJSON `json:"-"`
}

// settingZoneGetResponseInternalDNSJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseInternalDNS]
type settingZoneGetResponseInternalDNSJSON struct {
	ReferenceZoneID apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SettingZoneGetResponseInternalDNS) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseInternalDNSJSON) RawJSON() string {
	return r.raw
}

// Settings determining the nameservers through which the zone should be available.
type SettingZoneGetResponseNameservers struct {
	// Nameserver type
	Type SettingZoneGetResponseNameserversType `json:"type,required"`
	// Configured nameserver set to be used for this zone
	NSSet int64                                 `json:"ns_set"`
	JSON  settingZoneGetResponseNameserversJSON `json:"-"`
}

// settingZoneGetResponseNameserversJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseNameservers]
type settingZoneGetResponseNameserversJSON struct {
	Type        apijson.Field
	NSSet       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseNameservers) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseNameserversJSON) RawJSON() string {
	return r.raw
}

// Nameserver type
type SettingZoneGetResponseNameserversType string

const (
	SettingZoneGetResponseNameserversTypeCloudflareStandard SettingZoneGetResponseNameserversType = "cloudflare.standard"
	SettingZoneGetResponseNameserversTypeCustomAccount      SettingZoneGetResponseNameserversType = "custom.account"
	SettingZoneGetResponseNameserversTypeCustomTenant       SettingZoneGetResponseNameserversType = "custom.tenant"
	SettingZoneGetResponseNameserversTypeCustomZone         SettingZoneGetResponseNameserversType = "custom.zone"
)

func (r SettingZoneGetResponseNameserversType) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseNameserversTypeCloudflareStandard, SettingZoneGetResponseNameserversTypeCustomAccount, SettingZoneGetResponseNameserversTypeCustomTenant, SettingZoneGetResponseNameserversTypeCustomZone:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingZoneGetResponseSOA struct {
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
	TTL  float64                       `json:"ttl,required"`
	JSON settingZoneGetResponseSOAJSON `json:"-"`
}

// settingZoneGetResponseSOAJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseSOA]
type settingZoneGetResponseSOAJSON struct {
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

func (r *SettingZoneGetResponseSOA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseSOAJSON) RawJSON() string {
	return r.raw
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingZoneGetResponseZoneMode string

const (
	SettingZoneGetResponseZoneModeStandard SettingZoneGetResponseZoneMode = "standard"
	SettingZoneGetResponseZoneModeCDNOnly  SettingZoneGetResponseZoneMode = "cdn_only"
	SettingZoneGetResponseZoneModeDNSOnly  SettingZoneGetResponseZoneMode = "dns_only"
)

func (r SettingZoneGetResponseZoneMode) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseZoneModeStandard, SettingZoneGetResponseZoneModeCDNOnly, SettingZoneGetResponseZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingZoneEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether to flatten all CNAME records in the zone. Note that, due to DNS
	// limitations, a CNAME record at the zone apex will always be flattened.
	FlattenAllCNAMEs param.Field[bool] `json:"flatten_all_cnames"`
	// Whether to enable Foundation DNS Advanced Nameservers on the zone.
	FoundationDNS param.Field[bool] `json:"foundation_dns"`
	// Settings for this internal zone.
	InternalDNS param.Field[SettingZoneEditParamsInternalDNS] `json:"internal_dns"`
	// Whether to enable multi-provider DNS, which causes Cloudflare to activate the
	// zone even when non-Cloudflare NS records exist, and to respect NS records at the
	// zone apex during outbound zone transfers.
	MultiProvider param.Field[bool] `json:"multi_provider"`
	// Settings determining the nameservers through which the zone should be available.
	Nameservers param.Field[SettingZoneEditParamsNameservers] `json:"nameservers"`
	// The time to live (TTL) of the zone's nameserver (NS) records.
	NSTTL param.Field[float64] `json:"ns_ttl"`
	// Allows a Secondary DNS zone to use (proxied) override records and CNAME
	// flattening at the zone apex.
	SecondaryOverrides param.Field[bool] `json:"secondary_overrides"`
	// Components of the zone's SOA record.
	SOA param.Field[SettingZoneEditParamsSOA] `json:"soa"`
	// Whether the zone mode is a regular or CDN/DNS only zone.
	ZoneMode param.Field[SettingZoneEditParamsZoneMode] `json:"zone_mode"`
}

func (r SettingZoneEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings for this internal zone.
type SettingZoneEditParamsInternalDNS struct {
	// The ID of the zone to fallback to.
	ReferenceZoneID param.Field[string] `json:"reference_zone_id"`
}

func (r SettingZoneEditParamsInternalDNS) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Settings determining the nameservers through which the zone should be available.
type SettingZoneEditParamsNameservers struct {
	// Nameserver type
	Type param.Field[SettingZoneEditParamsNameserversType] `json:"type,required"`
	// Configured nameserver set to be used for this zone
	NSSet param.Field[int64] `json:"ns_set"`
}

func (r SettingZoneEditParamsNameservers) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nameserver type
type SettingZoneEditParamsNameserversType string

const (
	SettingZoneEditParamsNameserversTypeCloudflareStandard SettingZoneEditParamsNameserversType = "cloudflare.standard"
	SettingZoneEditParamsNameserversTypeCustomAccount      SettingZoneEditParamsNameserversType = "custom.account"
	SettingZoneEditParamsNameserversTypeCustomTenant       SettingZoneEditParamsNameserversType = "custom.tenant"
	SettingZoneEditParamsNameserversTypeCustomZone         SettingZoneEditParamsNameserversType = "custom.zone"
)

func (r SettingZoneEditParamsNameserversType) IsKnown() bool {
	switch r {
	case SettingZoneEditParamsNameserversTypeCloudflareStandard, SettingZoneEditParamsNameserversTypeCustomAccount, SettingZoneEditParamsNameserversTypeCustomTenant, SettingZoneEditParamsNameserversTypeCustomZone:
		return true
	}
	return false
}

// Components of the zone's SOA record.
type SettingZoneEditParamsSOA struct {
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

func (r SettingZoneEditParamsSOA) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether the zone mode is a regular or CDN/DNS only zone.
type SettingZoneEditParamsZoneMode string

const (
	SettingZoneEditParamsZoneModeStandard SettingZoneEditParamsZoneMode = "standard"
	SettingZoneEditParamsZoneModeCDNOnly  SettingZoneEditParamsZoneMode = "cdn_only"
	SettingZoneEditParamsZoneModeDNSOnly  SettingZoneEditParamsZoneMode = "dns_only"
)

func (r SettingZoneEditParamsZoneMode) IsKnown() bool {
	switch r {
	case SettingZoneEditParamsZoneModeStandard, SettingZoneEditParamsZoneModeCDNOnly, SettingZoneEditParamsZoneModeDNSOnly:
		return true
	}
	return false
}

type SettingZoneEditResponseEnvelope struct {
	Errors   []SettingZoneEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZoneEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingZoneEditResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingZoneEditResponse                `json:"result"`
	JSON    settingZoneEditResponseEnvelopeJSON    `json:"-"`
}

// settingZoneEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZoneEditResponseEnvelope]
type settingZoneEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           SettingZoneEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingZoneEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingZoneEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZoneEditResponseEnvelopeErrors]
type settingZoneEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    settingZoneEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingZoneEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingZoneEditResponseEnvelopeErrorsSource]
type settingZoneEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           SettingZoneEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingZoneEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingZoneEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingZoneEditResponseEnvelopeMessages]
type settingZoneEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingZoneEditResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    settingZoneEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingZoneEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SettingZoneEditResponseEnvelopeMessagesSource]
type settingZoneEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingZoneEditResponseEnvelopeSuccess bool

const (
	SettingZoneEditResponseEnvelopeSuccessTrue SettingZoneEditResponseEnvelopeSuccess = true
)

func (r SettingZoneEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingZoneEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SettingZoneGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type SettingZoneGetResponseEnvelope struct {
	Errors   []SettingZoneGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SettingZoneGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success SettingZoneGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SettingZoneGetResponse                `json:"result"`
	JSON    settingZoneGetResponseEnvelopeJSON    `json:"-"`
}

// settingZoneGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SettingZoneGetResponseEnvelope]
type settingZoneGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeErrors struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           SettingZoneGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             settingZoneGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// settingZoneGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [SettingZoneGetResponseEnvelopeErrors]
type settingZoneGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeErrorsSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    settingZoneGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// settingZoneGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [SettingZoneGetResponseEnvelopeErrorsSource]
type settingZoneGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeMessages struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           SettingZoneGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             settingZoneGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// settingZoneGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [SettingZoneGetResponseEnvelopeMessages]
type settingZoneGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type SettingZoneGetResponseEnvelopeMessagesSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    settingZoneGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// settingZoneGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [SettingZoneGetResponseEnvelopeMessagesSource]
type settingZoneGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SettingZoneGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r settingZoneGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type SettingZoneGetResponseEnvelopeSuccess bool

const (
	SettingZoneGetResponseEnvelopeSuccessTrue SettingZoneGetResponseEnvelopeSuccess = true
)

func (r SettingZoneGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SettingZoneGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
