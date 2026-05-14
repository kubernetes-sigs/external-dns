// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/tidwall/gjson"
)

type ASN = int64

type ASNParam = int64

type AuditLog struct {
	// A string that uniquely identifies the audit log.
	ID     string         `json:"id"`
	Action AuditLogAction `json:"action"`
	Actor  AuditLogActor  `json:"actor"`
	// The source of the event.
	Interface string `json:"interface"`
	// An object which can lend more context to the action being logged. This is a
	// flexible value and varies between different actions.
	Metadata interface{} `json:"metadata"`
	// The new value of the resource that was modified.
	NewValue string `json:"newValue"`
	// The value of the resource before it was modified.
	OldValue string           `json:"oldValue"`
	Owner    AuditLogOwner    `json:"owner"`
	Resource AuditLogResource `json:"resource"`
	// A UTC RFC3339 timestamp that specifies when the action being logged occured.
	When time.Time    `json:"when" format:"date-time"`
	JSON auditLogJSON `json:"-"`
}

// auditLogJSON contains the JSON metadata for the struct [AuditLog]
type auditLogJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Actor       apijson.Field
	Interface   apijson.Field
	Metadata    apijson.Field
	NewValue    apijson.Field
	OldValue    apijson.Field
	Owner       apijson.Field
	Resource    apijson.Field
	When        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLog) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogJSON) RawJSON() string {
	return r.raw
}

type AuditLogAction struct {
	// A boolean that indicates if the action attempted was successful.
	Result bool `json:"result"`
	// A short string that describes the action that was performed.
	Type string             `json:"type"`
	JSON auditLogActionJSON `json:"-"`
}

// auditLogActionJSON contains the JSON metadata for the struct [AuditLogAction]
type auditLogActionJSON struct {
	Result      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogActionJSON) RawJSON() string {
	return r.raw
}

type AuditLogActor struct {
	// The ID of the actor that performed the action. If a user performed the action,
	// this will be their User ID.
	ID string `json:"id"`
	// The email of the user that performed the action.
	Email string `json:"email" format:"email"`
	// The IP address of the request that performed the action.
	IP string `json:"ip"`
	// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
	Type AuditLogActorType `json:"type"`
	JSON auditLogActorJSON `json:"-"`
}

// auditLogActorJSON contains the JSON metadata for the struct [AuditLogActor]
type auditLogActorJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	IP          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogActor) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogActorJSON) RawJSON() string {
	return r.raw
}

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
type AuditLogActorType string

const (
	AuditLogActorTypeUser       AuditLogActorType = "user"
	AuditLogActorTypeAdmin      AuditLogActorType = "admin"
	AuditLogActorTypeCloudflare AuditLogActorType = "Cloudflare"
)

func (r AuditLogActorType) IsKnown() bool {
	switch r {
	case AuditLogActorTypeUser, AuditLogActorTypeAdmin, AuditLogActorTypeCloudflare:
		return true
	}
	return false
}

type AuditLogOwner struct {
	// Identifier
	ID   string            `json:"id"`
	JSON auditLogOwnerJSON `json:"-"`
}

// auditLogOwnerJSON contains the JSON metadata for the struct [AuditLogOwner]
type auditLogOwnerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogOwnerJSON) RawJSON() string {
	return r.raw
}

type AuditLogResource struct {
	// An identifier for the resource that was affected by the action.
	ID string `json:"id"`
	// A short string that describes the resource that was affected by the action.
	Type string               `json:"type"`
	JSON auditLogResourceJSON `json:"-"`
}

// auditLogResourceJSON contains the JSON metadata for the struct
// [AuditLogResource]
type auditLogResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AuditLogResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r auditLogResourceJSON) RawJSON() string {
	return r.raw
}

// The Certificate Authority that will issue the certificate
type CertificateCA string

const (
	CertificateCADigicert    CertificateCA = "digicert"
	CertificateCAGoogle      CertificateCA = "google"
	CertificateCALetsEncrypt CertificateCA = "lets_encrypt"
	CertificateCASSLCom      CertificateCA = "ssl_com"
)

func (r CertificateCA) IsKnown() bool {
	switch r {
	case CertificateCADigicert, CertificateCAGoogle, CertificateCALetsEncrypt, CertificateCASSLCom:
		return true
	}
	return false
}

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
type CertificateRequestType string

const (
	CertificateRequestTypeOriginRSA          CertificateRequestType = "origin-rsa"
	CertificateRequestTypeOriginECC          CertificateRequestType = "origin-ecc"
	CertificateRequestTypeKeylessCertificate CertificateRequestType = "keyless-certificate"
)

func (r CertificateRequestType) IsKnown() bool {
	switch r {
	case CertificateRequestTypeOriginRSA, CertificateRequestTypeOriginECC, CertificateRequestTypeKeylessCertificate:
		return true
	}
	return false
}

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
type CloudflareTunnel struct {
	// UUID of the tunnel.
	ID string `json:"id" format:"uuid"`
	// Cloudflare account ID
	AccountTag string `json:"account_tag"`
	// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
	// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
	// tunnel on the Zero Trust dashboard.
	ConfigSrc CloudflareTunnelConfigSrc `json:"config_src"`
	// The Cloudflare Tunnel connections between your origin and Cloudflare's edge.
	//
	// Deprecated: This field will start returning an empty array. To fetch the
	// connections of a given tunnel, please use the dedicated endpoint
	// `/accounts/{account_id}/{tunnel_type}/{tunnel_id}/connections`
	Connections []CloudflareTunnelConnection `json:"connections"`
	// Timestamp of when the tunnel established at least one connection to Cloudflare's
	// edge. If `null`, the tunnel is inactive.
	ConnsActiveAt time.Time `json:"conns_active_at" format:"date-time"`
	// Timestamp of when the tunnel became inactive (no connections to Cloudflare's
	// edge). If `null`, the tunnel is active.
	ConnsInactiveAt time.Time `json:"conns_inactive_at" format:"date-time"`
	// Timestamp of when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Timestamp of when the resource was deleted. If `null`, the resource has not been
	// deleted.
	DeletedAt time.Time `json:"deleted_at" format:"date-time"`
	// Metadata associated with the tunnel.
	Metadata interface{} `json:"metadata"`
	// A user-friendly name for a tunnel.
	Name string `json:"name"`
	// If `true`, the tunnel can be configured remotely from the Zero Trust dashboard.
	// If `false`, the tunnel must be configured locally on the origin machine.
	//
	// Deprecated: Use the config_src field instead.
	RemoteConfig bool `json:"remote_config"`
	// The status of the tunnel. Valid values are `inactive` (tunnel has never been
	// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
	// state), `healthy` (tunnel is active and able to serve traffic), or `down`
	// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
	Status CloudflareTunnelStatus `json:"status"`
	// The type of tunnel.
	TunType CloudflareTunnelTunType `json:"tun_type"`
	JSON    cloudflareTunnelJSON    `json:"-"`
}

// cloudflareTunnelJSON contains the JSON metadata for the struct
// [CloudflareTunnel]
type cloudflareTunnelJSON struct {
	ID              apijson.Field
	AccountTag      apijson.Field
	ConfigSrc       apijson.Field
	Connections     apijson.Field
	ConnsActiveAt   apijson.Field
	ConnsInactiveAt apijson.Field
	CreatedAt       apijson.Field
	DeletedAt       apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	RemoteConfig    apijson.Field
	Status          apijson.Field
	TunType         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *CloudflareTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudflareTunnelJSON) RawJSON() string {
	return r.raw
}

func (r CloudflareTunnel) ImplementsTunnelListResponse() {}

func (r CloudflareTunnel) ImplementsTunnelCloudflaredNewResponse() {}

func (r CloudflareTunnel) ImplementsTunnelCloudflaredListResponse() {}

func (r CloudflareTunnel) ImplementsTunnelCloudflaredDeleteResponse() {}

func (r CloudflareTunnel) ImplementsTunnelCloudflaredEditResponse() {}

func (r CloudflareTunnel) ImplementsTunnelCloudflaredGetResponse() {}

func (r CloudflareTunnel) ImplementsTunnelWARPConnectorNewResponse() {}

func (r CloudflareTunnel) ImplementsTunnelWARPConnectorListResponse() {}

func (r CloudflareTunnel) ImplementsTunnelWARPConnectorDeleteResponse() {}

func (r CloudflareTunnel) ImplementsTunnelWARPConnectorEditResponse() {}

func (r CloudflareTunnel) ImplementsTunnelWARPConnectorGetResponse() {}

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
type CloudflareTunnelConfigSrc string

const (
	CloudflareTunnelConfigSrcLocal      CloudflareTunnelConfigSrc = "local"
	CloudflareTunnelConfigSrcCloudflare CloudflareTunnelConfigSrc = "cloudflare"
)

func (r CloudflareTunnelConfigSrc) IsKnown() bool {
	switch r {
	case CloudflareTunnelConfigSrcLocal, CloudflareTunnelConfigSrcCloudflare:
		return true
	}
	return false
}

type CloudflareTunnelConnection struct {
	// UUID of the Cloudflare Tunnel connection.
	ID string `json:"id" format:"uuid"`
	// UUID of the Cloudflare Tunnel connector.
	ClientID string `json:"client_id" format:"uuid"`
	// The cloudflared version used to establish this connection.
	ClientVersion string `json:"client_version"`
	// The Cloudflare data center used for this connection.
	ColoName string `json:"colo_name"`
	// Cloudflare continues to track connections for several minutes after they
	// disconnect. This is an optimization to improve latency and reliability of
	// reconnecting. If `true`, the connection has disconnected but is still being
	// tracked. If `false`, the connection is actively serving traffic.
	IsPendingReconnect bool `json:"is_pending_reconnect"`
	// Timestamp of when the connection was established.
	OpenedAt time.Time `json:"opened_at" format:"date-time"`
	// The public IP address of the host running cloudflared.
	OriginIP string `json:"origin_ip"`
	// UUID of the Cloudflare Tunnel connection.
	UUID string                         `json:"uuid" format:"uuid"`
	JSON cloudflareTunnelConnectionJSON `json:"-"`
}

// cloudflareTunnelConnectionJSON contains the JSON metadata for the struct
// [CloudflareTunnelConnection]
type cloudflareTunnelConnectionJSON struct {
	ID                 apijson.Field
	ClientID           apijson.Field
	ClientVersion      apijson.Field
	ColoName           apijson.Field
	IsPendingReconnect apijson.Field
	OpenedAt           apijson.Field
	OriginIP           apijson.Field
	UUID               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *CloudflareTunnelConnection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r cloudflareTunnelConnectionJSON) RawJSON() string {
	return r.raw
}

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
type CloudflareTunnelStatus string

const (
	CloudflareTunnelStatusInactive CloudflareTunnelStatus = "inactive"
	CloudflareTunnelStatusDegraded CloudflareTunnelStatus = "degraded"
	CloudflareTunnelStatusHealthy  CloudflareTunnelStatus = "healthy"
	CloudflareTunnelStatusDown     CloudflareTunnelStatus = "down"
)

func (r CloudflareTunnelStatus) IsKnown() bool {
	switch r {
	case CloudflareTunnelStatusInactive, CloudflareTunnelStatusDegraded, CloudflareTunnelStatusHealthy, CloudflareTunnelStatusDown:
		return true
	}
	return false
}

// The type of tunnel.
type CloudflareTunnelTunType string

const (
	CloudflareTunnelTunTypeCfdTunnel     CloudflareTunnelTunType = "cfd_tunnel"
	CloudflareTunnelTunTypeWARPConnector CloudflareTunnelTunType = "warp_connector"
	CloudflareTunnelTunTypeWARP          CloudflareTunnelTunType = "warp"
	CloudflareTunnelTunTypeMagic         CloudflareTunnelTunType = "magic"
	CloudflareTunnelTunTypeIPSec         CloudflareTunnelTunType = "ip_sec"
	CloudflareTunnelTunTypeGRE           CloudflareTunnelTunType = "gre"
	CloudflareTunnelTunTypeCNI           CloudflareTunnelTunType = "cni"
)

func (r CloudflareTunnelTunType) IsKnown() bool {
	switch r {
	case CloudflareTunnelTunTypeCfdTunnel, CloudflareTunnelTunTypeWARPConnector, CloudflareTunnelTunTypeWARP, CloudflareTunnelTunTypeMagic, CloudflareTunnelTunTypeIPSec, CloudflareTunnelTunTypeGRE, CloudflareTunnelTunTypeCNI:
		return true
	}
	return false
}

type ErrorData struct {
	Code             int64           `json:"code"`
	DocumentationURL string          `json:"documentation_url"`
	Message          string          `json:"message"`
	Source           ErrorDataSource `json:"source"`
	JSON             errorDataJSON   `json:"-"`
}

// errorDataJSON contains the JSON metadata for the struct [ErrorData]
type errorDataJSON struct {
	Code             apijson.Field
	DocumentationURL apijson.Field
	Message          apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ErrorData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorDataJSON) RawJSON() string {
	return r.raw
}

type ErrorDataSource struct {
	Pointer string              `json:"pointer"`
	JSON    errorDataSourceJSON `json:"-"`
}

// errorDataSourceJSON contains the JSON metadata for the struct [ErrorDataSource]
type errorDataSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ErrorDataSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r errorDataSourceJSON) RawJSON() string {
	return r.raw
}

type Member struct {
	// Membership identifier tag.
	ID string `json:"id"`
	// Access policy for the membership
	Policies []MemberPolicy `json:"policies"`
	// Roles assigned to this Member.
	Roles []Role `json:"roles"`
	// A member's status in the account.
	Status MemberStatus `json:"status"`
	// Details of the user associated to the membership.
	User MemberUser `json:"user"`
	JSON memberJSON `json:"-"`
}

// memberJSON contains the JSON metadata for the struct [Member]
type memberJSON struct {
	ID          apijson.Field
	Policies    apijson.Field
	Roles       apijson.Field
	Status      apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Member) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberJSON) RawJSON() string {
	return r.raw
}

type MemberPolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access MemberPoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []MemberPoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []MemberPoliciesResourceGroup `json:"resource_groups"`
	JSON           memberPolicyJSON              `json:"-"`
}

// memberPolicyJSON contains the JSON metadata for the struct [MemberPolicy]
type memberPolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type MemberPoliciesAccess string

const (
	MemberPoliciesAccessAllow MemberPoliciesAccess = "allow"
	MemberPoliciesAccessDeny  MemberPoliciesAccess = "deny"
)

func (r MemberPoliciesAccess) IsKnown() bool {
	switch r {
	case MemberPoliciesAccessAllow, MemberPoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type MemberPoliciesPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta MemberPoliciesPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                            `json:"name"`
	JSON memberPoliciesPermissionGroupJSON `json:"-"`
}

// memberPoliciesPermissionGroupJSON contains the JSON metadata for the struct
// [MemberPoliciesPermissionGroup]
type memberPoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberPoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type MemberPoliciesPermissionGroupsMeta struct {
	Key   string                                 `json:"key"`
	Value string                                 `json:"value"`
	JSON  memberPoliciesPermissionGroupsMetaJSON `json:"-"`
}

// memberPoliciesPermissionGroupsMetaJSON contains the JSON metadata for the struct
// [MemberPoliciesPermissionGroupsMeta]
type memberPoliciesPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberPoliciesPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPoliciesPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type MemberPoliciesResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []MemberPoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta MemberPoliciesResourceGroupsMeta `json:"meta"`
	// Name of the resource group.
	Name string                          `json:"name"`
	JSON memberPoliciesResourceGroupJSON `json:"-"`
}

// memberPoliciesResourceGroupJSON contains the JSON metadata for the struct
// [MemberPoliciesResourceGroup]
type memberPoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberPoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type MemberPoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []MemberPoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    memberPoliciesResourceGroupsScopeJSON     `json:"-"`
}

// memberPoliciesResourceGroupsScopeJSON contains the JSON metadata for the struct
// [MemberPoliciesResourceGroupsScope]
type memberPoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberPoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type MemberPoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                      `json:"key,required"`
	JSON memberPoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// memberPoliciesResourceGroupsScopeObjectJSON contains the JSON metadata for the
// struct [MemberPoliciesResourceGroupsScopeObject]
type memberPoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberPoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type MemberPoliciesResourceGroupsMeta struct {
	Key   string                               `json:"key"`
	Value string                               `json:"value"`
	JSON  memberPoliciesResourceGroupsMetaJSON `json:"-"`
}

// memberPoliciesResourceGroupsMetaJSON contains the JSON metadata for the struct
// [MemberPoliciesResourceGroupsMeta]
type memberPoliciesResourceGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberPoliciesResourceGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberPoliciesResourceGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A member's status in the account.
type MemberStatus string

const (
	MemberStatusAccepted MemberStatus = "accepted"
	MemberStatusPending  MemberStatus = "pending"
)

func (r MemberStatus) IsKnown() bool {
	switch r {
	case MemberStatusAccepted, MemberStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type MemberUser struct {
	// The contact email address of the user.
	Email string `json:"email,required"`
	// Identifier
	ID string `json:"id"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// User's last name
	LastName string `json:"last_name,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool           `json:"two_factor_authentication_enabled"`
	JSON                           memberUserJSON `json:"-"`
}

// memberUserJSON contains the JSON metadata for the struct [MemberUser]
type memberUserJSON struct {
	Email                          apijson.Field
	ID                             apijson.Field
	FirstName                      apijson.Field
	LastName                       apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *MemberUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUserJSON) RawJSON() string {
	return r.raw
}

type Permission = string

type PermissionGrant struct {
	Read  bool                `json:"read"`
	Write bool                `json:"write"`
	JSON  permissionGrantJSON `json:"-"`
}

// permissionGrantJSON contains the JSON metadata for the struct [PermissionGrant]
type permissionGrantJSON struct {
	Read        apijson.Field
	Write       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PermissionGrant) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r permissionGrantJSON) RawJSON() string {
	return r.raw
}

type PermissionGrantParam struct {
	Read  param.Field[bool] `json:"read"`
	Write param.Field[bool] `json:"write"`
}

func (r PermissionGrantParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The rate plan applied to the subscription.
type RatePlan struct {
	// The ID of the rate plan.
	ID RatePlanID `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency string `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged bool `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract bool `json:"is_contract"`
	// The full name of the rate plan.
	PublicName string `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope string `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets []string     `json:"sets"`
	JSON ratePlanJSON `json:"-"`
}

// ratePlanJSON contains the JSON metadata for the struct [RatePlan]
type ratePlanJSON struct {
	ID                apijson.Field
	Currency          apijson.Field
	ExternallyManaged apijson.Field
	IsContract        apijson.Field
	PublicName        apijson.Field
	Scope             apijson.Field
	Sets              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *RatePlan) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ratePlanJSON) RawJSON() string {
	return r.raw
}

// The ID of the rate plan.
type RatePlanID string

const (
	RatePlanIDFree               RatePlanID = "free"
	RatePlanIDLite               RatePlanID = "lite"
	RatePlanIDPro                RatePlanID = "pro"
	RatePlanIDProPlus            RatePlanID = "pro_plus"
	RatePlanIDBusiness           RatePlanID = "business"
	RatePlanIDEnterprise         RatePlanID = "enterprise"
	RatePlanIDPartnersFree       RatePlanID = "partners_free"
	RatePlanIDPartnersPro        RatePlanID = "partners_pro"
	RatePlanIDPartnersBusiness   RatePlanID = "partners_business"
	RatePlanIDPartnersEnterprise RatePlanID = "partners_enterprise"
)

func (r RatePlanID) IsKnown() bool {
	switch r {
	case RatePlanIDFree, RatePlanIDLite, RatePlanIDPro, RatePlanIDProPlus, RatePlanIDBusiness, RatePlanIDEnterprise, RatePlanIDPartnersFree, RatePlanIDPartnersPro, RatePlanIDPartnersBusiness, RatePlanIDPartnersEnterprise:
		return true
	}
	return false
}

// The rate plan applied to the subscription.
type RatePlanParam struct {
	// The ID of the rate plan.
	ID param.Field[RatePlanID] `json:"id"`
	// The currency applied to the rate plan subscription.
	Currency param.Field[string] `json:"currency"`
	// Whether this rate plan is managed externally from Cloudflare.
	ExternallyManaged param.Field[bool] `json:"externally_managed"`
	// Whether a rate plan is enterprise-based (or newly adopted term contract).
	IsContract param.Field[bool] `json:"is_contract"`
	// The full name of the rate plan.
	PublicName param.Field[string] `json:"public_name"`
	// The scope that this rate plan applies to.
	Scope param.Field[string] `json:"scope"`
	// The list of sets this rate plan applies to.
	Sets param.Field[[]string] `json:"sets"`
}

func (r RatePlanParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ResponseInfo struct {
	Code             int64              `json:"code,required"`
	Message          string             `json:"message,required"`
	DocumentationURL string             `json:"documentation_url"`
	Source           ResponseInfoSource `json:"source"`
	JSON             responseInfoJSON   `json:"-"`
}

// responseInfoJSON contains the JSON metadata for the struct [ResponseInfo]
type responseInfoJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ResponseInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseInfoJSON) RawJSON() string {
	return r.raw
}

type ResponseInfoSource struct {
	Pointer string                 `json:"pointer"`
	JSON    responseInfoSourceJSON `json:"-"`
}

// responseInfoSourceJSON contains the JSON metadata for the struct
// [ResponseInfoSource]
type responseInfoSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseInfoSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r responseInfoSourceJSON) RawJSON() string {
	return r.raw
}

type Role struct {
	// Role identifier tag.
	ID string `json:"id,required"`
	// Description of role's permissions.
	Description string `json:"description,required"`
	// Role name.
	Name        string          `json:"name,required"`
	Permissions RolePermissions `json:"permissions,required"`
	JSON        roleJSON        `json:"-"`
}

// roleJSON contains the JSON metadata for the struct [Role]
type roleJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Role) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r roleJSON) RawJSON() string {
	return r.raw
}

type RolePermissions struct {
	Analytics    PermissionGrant     `json:"analytics"`
	Billing      PermissionGrant     `json:"billing"`
	CachePurge   PermissionGrant     `json:"cache_purge"`
	DNS          PermissionGrant     `json:"dns"`
	DNSRecords   PermissionGrant     `json:"dns_records"`
	LB           PermissionGrant     `json:"lb"`
	Logs         PermissionGrant     `json:"logs"`
	Organization PermissionGrant     `json:"organization"`
	SSL          PermissionGrant     `json:"ssl"`
	WAF          PermissionGrant     `json:"waf"`
	ZoneSettings PermissionGrant     `json:"zone_settings"`
	Zones        PermissionGrant     `json:"zones"`
	JSON         rolePermissionsJSON `json:"-"`
}

// rolePermissionsJSON contains the JSON metadata for the struct [RolePermissions]
type rolePermissionsJSON struct {
	Analytics    apijson.Field
	Billing      apijson.Field
	CachePurge   apijson.Field
	DNS          apijson.Field
	DNSRecords   apijson.Field
	LB           apijson.Field
	Logs         apijson.Field
	Organization apijson.Field
	SSL          apijson.Field
	WAF          apijson.Field
	ZoneSettings apijson.Field
	Zones        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RolePermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rolePermissionsJSON) RawJSON() string {
	return r.raw
}

type RoleParam struct {
	// Role identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r RoleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RolePermissionsParam struct {
	Analytics    param.Field[PermissionGrantParam] `json:"analytics"`
	Billing      param.Field[PermissionGrantParam] `json:"billing"`
	CachePurge   param.Field[PermissionGrantParam] `json:"cache_purge"`
	DNS          param.Field[PermissionGrantParam] `json:"dns"`
	DNSRecords   param.Field[PermissionGrantParam] `json:"dns_records"`
	LB           param.Field[PermissionGrantParam] `json:"lb"`
	Logs         param.Field[PermissionGrantParam] `json:"logs"`
	Organization param.Field[PermissionGrantParam] `json:"organization"`
	SSL          param.Field[PermissionGrantParam] `json:"ssl"`
	WAF          param.Field[PermissionGrantParam] `json:"waf"`
	ZoneSettings param.Field[PermissionGrantParam] `json:"zone_settings"`
	Zones        param.Field[PermissionGrantParam] `json:"zones"`
}

func (r RolePermissionsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Direction to order DNS records in.
type SortDirection string

const (
	SortDirectionAsc  SortDirection = "asc"
	SortDirectionDesc SortDirection = "desc"
)

func (r SortDirection) IsKnown() bool {
	switch r {
	case SortDirectionAsc, SortDirectionDesc:
		return true
	}
	return false
}

type Subscription struct {
	// Subscription identifier tag.
	ID string `json:"id"`
	// The monetary unit in which pricing information is displayed.
	Currency string `json:"currency"`
	// The end of the current period and also when the next billing is due.
	CurrentPeriodEnd time.Time `json:"current_period_end" format:"date-time"`
	// When the current billing period started. May match initial_period_start if this
	// is the first period.
	CurrentPeriodStart time.Time `json:"current_period_start" format:"date-time"`
	// How often the subscription is renewed automatically.
	Frequency SubscriptionFrequency `json:"frequency"`
	// The price of the subscription that will be billed, in US dollars.
	Price float64 `json:"price"`
	// The rate plan applied to the subscription.
	RatePlan RatePlan `json:"rate_plan"`
	// The state that the subscription is in.
	State SubscriptionState `json:"state"`
	JSON  subscriptionJSON  `json:"-"`
}

// subscriptionJSON contains the JSON metadata for the struct [Subscription]
type subscriptionJSON struct {
	ID                 apijson.Field
	Currency           apijson.Field
	CurrentPeriodEnd   apijson.Field
	CurrentPeriodStart apijson.Field
	Frequency          apijson.Field
	Price              apijson.Field
	RatePlan           apijson.Field
	State              apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Subscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r subscriptionJSON) RawJSON() string {
	return r.raw
}

// How often the subscription is renewed automatically.
type SubscriptionFrequency string

const (
	SubscriptionFrequencyWeekly    SubscriptionFrequency = "weekly"
	SubscriptionFrequencyMonthly   SubscriptionFrequency = "monthly"
	SubscriptionFrequencyQuarterly SubscriptionFrequency = "quarterly"
	SubscriptionFrequencyYearly    SubscriptionFrequency = "yearly"
)

func (r SubscriptionFrequency) IsKnown() bool {
	switch r {
	case SubscriptionFrequencyWeekly, SubscriptionFrequencyMonthly, SubscriptionFrequencyQuarterly, SubscriptionFrequencyYearly:
		return true
	}
	return false
}

// The state that the subscription is in.
type SubscriptionState string

const (
	SubscriptionStateTrial           SubscriptionState = "Trial"
	SubscriptionStateProvisioned     SubscriptionState = "Provisioned"
	SubscriptionStatePaid            SubscriptionState = "Paid"
	SubscriptionStateAwaitingPayment SubscriptionState = "AwaitingPayment"
	SubscriptionStateCancelled       SubscriptionState = "Cancelled"
	SubscriptionStateFailed          SubscriptionState = "Failed"
	SubscriptionStateExpired         SubscriptionState = "Expired"
)

func (r SubscriptionState) IsKnown() bool {
	switch r {
	case SubscriptionStateTrial, SubscriptionStateProvisioned, SubscriptionStatePaid, SubscriptionStateAwaitingPayment, SubscriptionStateCancelled, SubscriptionStateFailed, SubscriptionStateExpired:
		return true
	}
	return false
}

type SubscriptionParam struct {
	// How often the subscription is renewed automatically.
	Frequency param.Field[SubscriptionFrequency] `json:"frequency"`
	// The rate plan applied to the subscription.
	RatePlan param.Field[RatePlanParam] `json:"rate_plan"`
}

func (r SubscriptionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Token struct {
	// Token identifier tag.
	ID        string         `json:"id"`
	Condition TokenCondition `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The time on which the token was created.
	IssuedOn time.Time `json:"issued_on" format:"date-time"`
	// Last time the token was used.
	LastUsedOn time.Time `json:"last_used_on" format:"date-time"`
	// Last time the token was modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// Token name.
	Name string `json:"name"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore time.Time `json:"not_before" format:"date-time"`
	// List of access policies assigned to the token.
	Policies []TokenPolicy `json:"policies"`
	// Status of the token.
	Status TokenStatus `json:"status"`
	JSON   tokenJSON   `json:"-"`
}

// tokenJSON contains the JSON metadata for the struct [Token]
type tokenJSON struct {
	ID          apijson.Field
	Condition   apijson.Field
	ExpiresOn   apijson.Field
	IssuedOn    apijson.Field
	LastUsedOn  apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	NotBefore   apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Token) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenJSON) RawJSON() string {
	return r.raw
}

type TokenCondition struct {
	// Client IP restrictions.
	RequestIP TokenConditionRequestIP `json:"request_ip"`
	JSON      tokenConditionJSON      `json:"-"`
}

// tokenConditionJSON contains the JSON metadata for the struct [TokenCondition]
type tokenConditionJSON struct {
	RequestIP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenCondition) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenConditionJSON) RawJSON() string {
	return r.raw
}

// Client IP restrictions.
type TokenConditionRequestIP struct {
	// List of IPv4/IPv6 CIDR addresses.
	In []TokenConditionCIDRList `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn []TokenConditionCIDRList    `json:"not_in"`
	JSON  tokenConditionRequestIPJSON `json:"-"`
}

// tokenConditionRequestIPJSON contains the JSON metadata for the struct
// [TokenConditionRequestIP]
type tokenConditionRequestIPJSON struct {
	In          apijson.Field
	NotIn       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenConditionRequestIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenConditionRequestIPJSON) RawJSON() string {
	return r.raw
}

// Status of the token.
type TokenStatus string

const (
	TokenStatusActive   TokenStatus = "active"
	TokenStatusDisabled TokenStatus = "disabled"
	TokenStatusExpired  TokenStatus = "expired"
)

func (r TokenStatus) IsKnown() bool {
	switch r {
	case TokenStatusActive, TokenStatusDisabled, TokenStatusExpired:
		return true
	}
	return false
}

type TokenParam struct {
	Condition param.Field[TokenConditionParam] `json:"condition"`
	// The expiration time on or after which the JWT MUST NOT be accepted for
	// processing.
	ExpiresOn param.Field[time.Time] `json:"expires_on" format:"date-time"`
	// Token name.
	Name param.Field[string] `json:"name"`
	// The time before which the token MUST NOT be accepted for processing.
	NotBefore param.Field[time.Time] `json:"not_before" format:"date-time"`
	// List of access policies assigned to the token.
	Policies param.Field[[]TokenPolicyParam] `json:"policies"`
	// Status of the token.
	Status param.Field[TokenStatus] `json:"status"`
}

func (r TokenParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenConditionParam struct {
	// Client IP restrictions.
	RequestIP param.Field[TokenConditionRequestIPParam] `json:"request_ip"`
}

func (r TokenConditionParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Client IP restrictions.
type TokenConditionRequestIPParam struct {
	// List of IPv4/IPv6 CIDR addresses.
	In param.Field[[]TokenConditionCIDRListParam] `json:"in"`
	// List of IPv4/IPv6 CIDR addresses.
	NotIn param.Field[[]TokenConditionCIDRListParam] `json:"not_in"`
}

func (r TokenConditionRequestIPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type TokenConditionCIDRList = string

type TokenConditionCIDRListParam = string

type TokenPolicy struct {
	// Policy identifier.
	ID string `json:"id,required"`
	// Allow or deny operations against the resources.
	Effect TokenPolicyEffect `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []TokenPolicyPermissionGroup `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources TokenPolicyResourcesUnion `json:"resources,required"`
	JSON      tokenPolicyJSON           `json:"-"`
}

// tokenPolicyJSON contains the JSON metadata for the struct [TokenPolicy]
type tokenPolicyJSON struct {
	ID               apijson.Field
	Effect           apijson.Field
	PermissionGroups apijson.Field
	Resources        apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TokenPolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type TokenPolicyEffect string

const (
	TokenPolicyEffectAllow TokenPolicyEffect = "allow"
	TokenPolicyEffectDeny  TokenPolicyEffect = "deny"
)

func (r TokenPolicyEffect) IsKnown() bool {
	switch r {
	case TokenPolicyEffectAllow, TokenPolicyEffectDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type TokenPolicyPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta TokenPolicyPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                         `json:"name"`
	JSON tokenPolicyPermissionGroupJSON `json:"-"`
}

// tokenPolicyPermissionGroupJSON contains the JSON metadata for the struct
// [TokenPolicyPermissionGroup]
type tokenPolicyPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPolicyPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPolicyPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type TokenPolicyPermissionGroupsMeta struct {
	Key   string                              `json:"key"`
	Value string                              `json:"value"`
	JSON  tokenPolicyPermissionGroupsMetaJSON `json:"-"`
}

// tokenPolicyPermissionGroupsMetaJSON contains the JSON metadata for the struct
// [TokenPolicyPermissionGroupsMeta]
type tokenPolicyPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenPolicyPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenPolicyPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A list of resource names that the policy applies to.
//
// Union satisfied by [TokenPolicyResourcesIAMResourcesTypeObjectString] or
// [TokenPolicyResourcesIAMResourcesTypeObjectNested].
type TokenPolicyResourcesUnion interface {
	implementsTokenPolicyResourcesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*TokenPolicyResourcesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenPolicyResourcesIAMResourcesTypeObjectString{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(TokenPolicyResourcesIAMResourcesTypeObjectNested{}),
		},
	)
}

type TokenPolicyResourcesIAMResourcesTypeObjectString map[string]string

func (r TokenPolicyResourcesIAMResourcesTypeObjectString) implementsTokenPolicyResourcesUnion() {}

type TokenPolicyResourcesIAMResourcesTypeObjectNested map[string]map[string]string

func (r TokenPolicyResourcesIAMResourcesTypeObjectNested) implementsTokenPolicyResourcesUnion() {}

type TokenPolicyParam struct {
	// Allow or deny operations against the resources.
	Effect param.Field[TokenPolicyEffect] `json:"effect,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]TokenPolicyPermissionGroupParam] `json:"permission_groups,required"`
	// A list of resource names that the policy applies to.
	Resources param.Field[TokenPolicyResourcesUnionParam] `json:"resources,required"`
}

func (r TokenPolicyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A named group of permissions that map to a group of operations against
// resources.
type TokenPolicyPermissionGroupParam struct {
	// Identifier of the permission group.
	ID param.Field[string] `json:"id,required"`
	// Attributes associated to the permission group.
	Meta param.Field[TokenPolicyPermissionGroupsMetaParam] `json:"meta"`
}

func (r TokenPolicyPermissionGroupParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Attributes associated to the permission group.
type TokenPolicyPermissionGroupsMetaParam struct {
	Key   param.Field[string] `json:"key"`
	Value param.Field[string] `json:"value"`
}

func (r TokenPolicyPermissionGroupsMetaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A list of resource names that the policy applies to.
//
// Satisfied by [shared.TokenPolicyResourcesIAMResourcesTypeObjectStringParam],
// [shared.TokenPolicyResourcesIAMResourcesTypeObjectNestedParam].
type TokenPolicyResourcesUnionParam interface {
	implementsTokenPolicyResourcesUnionParam()
}

type TokenPolicyResourcesIAMResourcesTypeObjectStringParam map[string]string

func (r TokenPolicyResourcesIAMResourcesTypeObjectStringParam) implementsTokenPolicyResourcesUnionParam() {
}

type TokenPolicyResourcesIAMResourcesTypeObjectNestedParam map[string]map[string]string

func (r TokenPolicyResourcesIAMResourcesTypeObjectNestedParam) implementsTokenPolicyResourcesUnionParam() {
}

type TokenValue = string
