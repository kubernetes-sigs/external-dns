// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zones

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
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// ZoneService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneService] method instead.
type ZoneService struct {
	Options         []option.RequestOption
	ActivationCheck *ActivationCheckService
	Settings        *SettingService
	// Deprecated: Use DNS settings API instead.
	CustomNameservers *CustomNameserverService
	Holds             *HoldService
	Subscriptions     *SubscriptionService
	Plans             *PlanService
	RatePlans         *RatePlanService
}

// NewZoneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneService(opts ...option.RequestOption) (r *ZoneService) {
	r = &ZoneService{}
	r.Options = opts
	r.ActivationCheck = NewActivationCheckService(opts...)
	r.Settings = NewSettingService(opts...)
	r.CustomNameservers = NewCustomNameserverService(opts...)
	r.Holds = NewHoldService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Plans = NewPlanService(opts...)
	r.RatePlans = NewRatePlanService(opts...)
	return
}

// Create Zone
func (r *ZoneService) New(ctx context.Context, body ZoneNewParams, opts ...option.RequestOption) (res *Zone, err error) {
	var env ZoneNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "zones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists, searches, sorts, and filters your zones. Listing zones across more than
// 500 accounts is currently not allowed.
func (r *ZoneService) List(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Zone], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "zones"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// Lists, searches, sorts, and filters your zones. Listing zones across more than
// 500 accounts is currently not allowed.
func (r *ZoneService) ListAutoPaging(ctx context.Context, query ZoneListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Zone] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Deletes an existing zone.
func (r *ZoneService) Delete(ctx context.Context, body ZoneDeleteParams, opts ...option.RequestOption) (res *ZoneDeleteResponse, err error) {
	var env ZoneDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edits a zone. Only one zone property can be changed at a time.
func (r *ZoneService) Edit(ctx context.Context, params ZoneEditParams, opts ...option.RequestOption) (res *Zone, err error) {
	var env ZoneEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Zone Details
func (r *ZoneService) Get(ctx context.Context, query ZoneGetParams, opts ...option.RequestOption) (res *Zone, err error) {
	var env ZoneGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup.
type Type string

const (
	TypeFull      Type = "full"
	TypePartial   Type = "partial"
	TypeSecondary Type = "secondary"
	TypeInternal  Type = "internal"
)

func (r Type) IsKnown() bool {
	switch r {
	case TypeFull, TypePartial, TypeSecondary, TypeInternal:
		return true
	}
	return false
}

type Zone struct {
	// Identifier
	ID string `json:"id,required"`
	// The account the zone belongs to.
	Account ZoneAccount `json:"account,required"`
	// The last time proof of ownership was detected and the zone was made active.
	ActivatedOn time.Time `json:"activated_on,required,nullable" format:"date-time"`
	// When the zone was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The interval (in seconds) from when development mode expires (positive integer)
	// or last expired (negative integer) for the domain. If development mode has never
	// been enabled, this value is 0.
	DevelopmentMode float64 `json:"development_mode,required"`
	// Metadata about the zone.
	Meta ZoneMeta `json:"meta,required"`
	// When the zone was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The domain name.
	Name string `json:"name,required"`
	// The name servers Cloudflare assigns to a zone.
	NameServers []string `json:"name_servers,required" format:"hostname"`
	// DNS host at the time of switching to Cloudflare.
	OriginalDnshost string `json:"original_dnshost,required,nullable"`
	// Original name servers before moving to Cloudflare.
	OriginalNameServers []string `json:"original_name_servers,required,nullable" format:"hostname"`
	// Registrar for the domain at the time of switching to Cloudflare.
	OriginalRegistrar string `json:"original_registrar,required,nullable"`
	// The owner of the zone.
	Owner ZoneOwner `json:"owner,required"`
	// A Zones subscription information.
	//
	// Deprecated: deprecated
	Plan ZonePlan `json:"plan,required"`
	// Allows the customer to use a custom apex. _Tenants Only Configuration_.
	CNAMESuffix string `json:"cname_suffix"`
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused bool `json:"paused"`
	// Legacy permissions based on legacy user membership information.
	//
	// Deprecated: deprecated
	Permissions []string `json:"permissions"`
	// The zone status on Cloudflare.
	Status ZoneStatus `json:"status"`
	// The root organizational unit that this zone belongs to (such as a tenant or
	// organization).
	Tenant ZoneTenant `json:"tenant"`
	// The immediate parent organizational unit that this zone belongs to (such as
	// under a tenant or sub-organization).
	TenantUnit ZoneTenantUnit `json:"tenant_unit"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type Type `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers []string `json:"vanity_name_servers" format:"hostname"`
	// Verification key for partial zone setup.
	VerificationKey string   `json:"verification_key"`
	JSON            zoneJSON `json:"-"`
}

// zoneJSON contains the JSON metadata for the struct [Zone]
type zoneJSON struct {
	ID                  apijson.Field
	Account             apijson.Field
	ActivatedOn         apijson.Field
	CreatedOn           apijson.Field
	DevelopmentMode     apijson.Field
	Meta                apijson.Field
	ModifiedOn          apijson.Field
	Name                apijson.Field
	NameServers         apijson.Field
	OriginalDnshost     apijson.Field
	OriginalNameServers apijson.Field
	OriginalRegistrar   apijson.Field
	Owner               apijson.Field
	Plan                apijson.Field
	CNAMESuffix         apijson.Field
	Paused              apijson.Field
	Permissions         apijson.Field
	Status              apijson.Field
	Tenant              apijson.Field
	TenantUnit          apijson.Field
	Type                apijson.Field
	VanityNameServers   apijson.Field
	VerificationKey     apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Zone) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneJSON) RawJSON() string {
	return r.raw
}

// The account the zone belongs to.
type ZoneAccount struct {
	// Identifier
	ID string `json:"id"`
	// The name of the account.
	Name string          `json:"name"`
	JSON zoneAccountJSON `json:"-"`
}

// zoneAccountJSON contains the JSON metadata for the struct [ZoneAccount]
type zoneAccountJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAccount) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAccountJSON) RawJSON() string {
	return r.raw
}

// Metadata about the zone.
type ZoneMeta struct {
	// The zone is only configured for CDN.
	CDNOnly bool `json:"cdn_only"`
	// Number of Custom Certificates the zone can have.
	CustomCertificateQuota int64 `json:"custom_certificate_quota"`
	// The zone is only configured for DNS.
	DNSOnly bool `json:"dns_only"`
	// The zone is setup with Foundation DNS.
	FoundationDNS bool `json:"foundation_dns"`
	// Number of Page Rules a zone can have.
	PageRuleQuota int64 `json:"page_rule_quota"`
	// The zone has been flagged for phishing.
	PhishingDetected bool         `json:"phishing_detected"`
	Step             int64        `json:"step"`
	JSON             zoneMetaJSON `json:"-"`
}

// zoneMetaJSON contains the JSON metadata for the struct [ZoneMeta]
type zoneMetaJSON struct {
	CDNOnly                apijson.Field
	CustomCertificateQuota apijson.Field
	DNSOnly                apijson.Field
	FoundationDNS          apijson.Field
	PageRuleQuota          apijson.Field
	PhishingDetected       apijson.Field
	Step                   apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *ZoneMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneMetaJSON) RawJSON() string {
	return r.raw
}

// The owner of the zone.
type ZoneOwner struct {
	// Identifier
	ID string `json:"id"`
	// Name of the owner.
	Name string `json:"name"`
	// The type of owner.
	Type string        `json:"type"`
	JSON zoneOwnerJSON `json:"-"`
}

// zoneOwnerJSON contains the JSON metadata for the struct [ZoneOwner]
type zoneOwnerJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneOwnerJSON) RawJSON() string {
	return r.raw
}

// A Zones subscription information.
//
// Deprecated: deprecated
type ZonePlan struct {
	// Identifier
	ID string `json:"id"`
	// States if the subscription can be activated.
	CanSubscribe bool `json:"can_subscribe"`
	// The denomination of the customer.
	Currency string `json:"currency"`
	// If this Zone is managed by another company.
	ExternallyManaged bool `json:"externally_managed"`
	// How often the customer is billed.
	Frequency string `json:"frequency"`
	// States if the subscription active.
	IsSubscribed bool `json:"is_subscribed"`
	// If the legacy discount applies to this Zone.
	LegacyDiscount bool `json:"legacy_discount"`
	// The legacy name of the plan.
	LegacyID string `json:"legacy_id"`
	// Name of the owner.
	Name string `json:"name"`
	// How much the customer is paying.
	Price float64      `json:"price"`
	JSON  zonePlanJSON `json:"-"`
}

// zonePlanJSON contains the JSON metadata for the struct [ZonePlan]
type zonePlanJSON struct {
	ID                apijson.Field
	CanSubscribe      apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	Frequency         apijson.Field
	IsSubscribed      apijson.Field
	LegacyDiscount    apijson.Field
	LegacyID          apijson.Field
	Name              apijson.Field
	Price             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ZonePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonePlanJSON) RawJSON() string {
	return r.raw
}

// The zone status on Cloudflare.
type ZoneStatus string

const (
	ZoneStatusInitializing ZoneStatus = "initializing"
	ZoneStatusPending      ZoneStatus = "pending"
	ZoneStatusActive       ZoneStatus = "active"
	ZoneStatusMoved        ZoneStatus = "moved"
)

func (r ZoneStatus) IsKnown() bool {
	switch r {
	case ZoneStatusInitializing, ZoneStatusPending, ZoneStatusActive, ZoneStatusMoved:
		return true
	}
	return false
}

// The root organizational unit that this zone belongs to (such as a tenant or
// organization).
type ZoneTenant struct {
	// Identifier
	ID string `json:"id"`
	// The name of the Tenant account.
	Name string         `json:"name"`
	JSON zoneTenantJSON `json:"-"`
}

// zoneTenantJSON contains the JSON metadata for the struct [ZoneTenant]
type zoneTenantJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTenant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTenantJSON) RawJSON() string {
	return r.raw
}

// The immediate parent organizational unit that this zone belongs to (such as
// under a tenant or sub-organization).
type ZoneTenantUnit struct {
	// Identifier
	ID   string             `json:"id"`
	JSON zoneTenantUnitJSON `json:"-"`
}

// zoneTenantUnitJSON contains the JSON metadata for the struct [ZoneTenantUnit]
type zoneTenantUnitJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTenantUnit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTenantUnitJSON) RawJSON() string {
	return r.raw
}

type ZoneDeleteResponse struct {
	// Identifier
	ID   string                 `json:"id,required"`
	JSON zoneDeleteResponseJSON `json:"-"`
}

// zoneDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneDeleteResponse]
type zoneDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneNewParams struct {
	Account param.Field[ZoneNewParamsAccount] `json:"account,required"`
	// The domain name.
	Name param.Field[string] `json:"name,required"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup.
	Type param.Field[Type] `json:"type"`
}

func (r ZoneNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneNewParamsAccount struct {
	// Identifier
	ID param.Field[string] `json:"id"`
}

func (r ZoneNewParamsAccount) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                        `json:"success,required"`
	Result  Zone                        `json:"result"`
	JSON    zoneNewResponseEnvelopeJSON `json:"-"`
}

// zoneNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneNewResponseEnvelope]
type zoneNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneListParams struct {
	Account param.Field[ZoneListParamsAccount] `query:"account"`
	// Direction to order zones.
	Direction param.Field[ZoneListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any).
	Match param.Field[ZoneListParamsMatch] `query:"match"`
	// A domain name. Optional filter operators can be provided to extend refine the
	// search:
	//
	// - `equal` (default)
	// - `not_equal`
	// - `starts_with`
	// - `ends_with`
	// - `contains`
	// - `starts_with_case_sensitive`
	// - `ends_with_case_sensitive`
	// - `contains_case_sensitive`
	Name param.Field[string] `query:"name"`
	// Field to order zones by.
	Order param.Field[ZoneListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of zones per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Specify a zone status to filter by.
	Status param.Field[ZoneListParamsStatus] `query:"status"`
}

// URLQuery serializes [ZoneListParams]'s query parameters as `url.Values`.
func (r ZoneListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ZoneListParamsAccount struct {
	// Filter by an account ID.
	ID param.Field[string] `query:"id"`
	// An account Name. Optional filter operators can be provided to extend refine the
	// search:
	//
	// - `equal` (default)
	// - `not_equal`
	// - `starts_with`
	// - `ends_with`
	// - `contains`
	// - `starts_with_case_sensitive`
	// - `ends_with_case_sensitive`
	// - `contains_case_sensitive`
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [ZoneListParamsAccount]'s query parameters as `url.Values`.
func (r ZoneListParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order zones.
type ZoneListParamsDirection string

const (
	ZoneListParamsDirectionAsc  ZoneListParamsDirection = "asc"
	ZoneListParamsDirectionDesc ZoneListParamsDirection = "desc"
)

func (r ZoneListParamsDirection) IsKnown() bool {
	switch r {
	case ZoneListParamsDirectionAsc, ZoneListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any).
type ZoneListParamsMatch string

const (
	ZoneListParamsMatchAny ZoneListParamsMatch = "any"
	ZoneListParamsMatchAll ZoneListParamsMatch = "all"
)

func (r ZoneListParamsMatch) IsKnown() bool {
	switch r {
	case ZoneListParamsMatchAny, ZoneListParamsMatchAll:
		return true
	}
	return false
}

// Field to order zones by.
type ZoneListParamsOrder string

const (
	ZoneListParamsOrderName        ZoneListParamsOrder = "name"
	ZoneListParamsOrderStatus      ZoneListParamsOrder = "status"
	ZoneListParamsOrderAccountID   ZoneListParamsOrder = "account.id"
	ZoneListParamsOrderAccountName ZoneListParamsOrder = "account.name"
	ZoneListParamsOrderPlanID      ZoneListParamsOrder = "plan.id"
)

func (r ZoneListParamsOrder) IsKnown() bool {
	switch r {
	case ZoneListParamsOrderName, ZoneListParamsOrderStatus, ZoneListParamsOrderAccountID, ZoneListParamsOrderAccountName, ZoneListParamsOrderPlanID:
		return true
	}
	return false
}

// Specify a zone status to filter by.
type ZoneListParamsStatus string

const (
	ZoneListParamsStatusInitializing ZoneListParamsStatus = "initializing"
	ZoneListParamsStatusPending      ZoneListParamsStatus = "pending"
	ZoneListParamsStatusActive       ZoneListParamsStatus = "active"
	ZoneListParamsStatusMoved        ZoneListParamsStatus = "moved"
)

func (r ZoneListParamsStatus) IsKnown() bool {
	switch r {
	case ZoneListParamsStatusInitializing, ZoneListParamsStatusPending, ZoneListParamsStatusActive, ZoneListParamsStatusMoved:
		return true
	}
	return false
}

type ZoneDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                           `json:"success,required"`
	Result  ZoneDeleteResponse             `json:"result,nullable"`
	JSON    zoneDeleteResponseEnvelopeJSON `json:"-"`
}

// zoneDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneDeleteResponseEnvelope]
type zoneDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Indicates whether the zone is only using Cloudflare DNS services. A true value
	// means the zone will not receive security or performance benefits.
	Paused param.Field[bool] `json:"paused"`
	// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
	// typically a partner-hosted zone or a CNAME setup. This parameter is only
	// available to Enterprise customers or if it has been explicitly enabled on a
	// zone.
	Type param.Field[ZoneEditParamsType] `json:"type"`
	// An array of domains used for custom name servers. This is only available for
	// Business and Enterprise plans.
	VanityNameServers param.Field[[]string] `json:"vanity_name_servers" format:"hostname"`
}

func (r ZoneEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A full zone implies that DNS is hosted with Cloudflare. A partial zone is
// typically a partner-hosted zone or a CNAME setup. This parameter is only
// available to Enterprise customers or if it has been explicitly enabled on a
// zone.
type ZoneEditParamsType string

const (
	ZoneEditParamsTypeFull      ZoneEditParamsType = "full"
	ZoneEditParamsTypePartial   ZoneEditParamsType = "partial"
	ZoneEditParamsTypeSecondary ZoneEditParamsType = "secondary"
	ZoneEditParamsTypeInternal  ZoneEditParamsType = "internal"
)

func (r ZoneEditParamsType) IsKnown() bool {
	switch r {
	case ZoneEditParamsTypeFull, ZoneEditParamsTypePartial, ZoneEditParamsTypeSecondary, ZoneEditParamsTypeInternal:
		return true
	}
	return false
}

type ZoneEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                         `json:"success,required"`
	Result  Zone                         `json:"result"`
	JSON    zoneEditResponseEnvelopeJSON `json:"-"`
}

// zoneEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneEditResponseEnvelope]
type zoneEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful.
	Success bool                        `json:"success,required"`
	Result  Zone                        `json:"result"`
	JSON    zoneGetResponseEnvelopeJSON `json:"-"`
}

// zoneGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneGetResponseEnvelope]
type zoneGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
