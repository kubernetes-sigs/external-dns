// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"github.com/cloudflare/cloudflare-go/v5/internal/apierror"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

type Error = apierror.Error

// This is an alias to an internal type.
type ASN = shared.ASN

// This is an alias to an internal type.
type ASNParam = shared.ASNParam

// This is an alias to an internal type.
type AuditLog = shared.AuditLog

// This is an alias to an internal type.
type AuditLogAction = shared.AuditLogAction

// This is an alias to an internal type.
type AuditLogActor = shared.AuditLogActor

// The type of actor, whether a User, Cloudflare Admin, or an Automated System.
//
// This is an alias to an internal type.
type AuditLogActorType = shared.AuditLogActorType

// This is an alias to an internal value.
const AuditLogActorTypeUser = shared.AuditLogActorTypeUser

// This is an alias to an internal value.
const AuditLogActorTypeAdmin = shared.AuditLogActorTypeAdmin

// This is an alias to an internal value.
const AuditLogActorTypeCloudflare = shared.AuditLogActorTypeCloudflare

// This is an alias to an internal type.
type AuditLogOwner = shared.AuditLogOwner

// This is an alias to an internal type.
type AuditLogResource = shared.AuditLogResource

// The Certificate Authority that will issue the certificate
//
// This is an alias to an internal type.
type CertificateCA = shared.CertificateCA

// This is an alias to an internal value.
const CertificateCADigicert = shared.CertificateCADigicert

// This is an alias to an internal value.
const CertificateCAGoogle = shared.CertificateCAGoogle

// This is an alias to an internal value.
const CertificateCALetsEncrypt = shared.CertificateCALetsEncrypt

// This is an alias to an internal value.
const CertificateCASSLCom = shared.CertificateCASSLCom

// Signature type desired on certificate ("origin-rsa" (rsa), "origin-ecc" (ecdsa),
// or "keyless-certificate" (for Keyless SSL servers).
//
// This is an alias to an internal type.
type CertificateRequestType = shared.CertificateRequestType

// This is an alias to an internal value.
const CertificateRequestTypeOriginRSA = shared.CertificateRequestTypeOriginRSA

// This is an alias to an internal value.
const CertificateRequestTypeOriginECC = shared.CertificateRequestTypeOriginECC

// This is an alias to an internal value.
const CertificateRequestTypeKeylessCertificate = shared.CertificateRequestTypeKeylessCertificate

// A Cloudflare Tunnel that connects your origin to Cloudflare's edge.
//
// This is an alias to an internal type.
type CloudflareTunnel = shared.CloudflareTunnel

// Indicates if this is a locally or remotely configured tunnel. If `local`, manage
// the tunnel using a YAML file on the origin machine. If `cloudflare`, manage the
// tunnel on the Zero Trust dashboard.
//
// This is an alias to an internal type.
type CloudflareTunnelConfigSrc = shared.CloudflareTunnelConfigSrc

// This is an alias to an internal value.
const CloudflareTunnelConfigSrcLocal = shared.CloudflareTunnelConfigSrcLocal

// This is an alias to an internal value.
const CloudflareTunnelConfigSrcCloudflare = shared.CloudflareTunnelConfigSrcCloudflare

// This is an alias to an internal type.
type CloudflareTunnelConnection = shared.CloudflareTunnelConnection

// The status of the tunnel. Valid values are `inactive` (tunnel has never been
// run), `degraded` (tunnel is active and able to serve traffic but in an unhealthy
// state), `healthy` (tunnel is active and able to serve traffic), or `down`
// (tunnel can not serve traffic as it has no connections to the Cloudflare Edge).
//
// This is an alias to an internal type.
type CloudflareTunnelStatus = shared.CloudflareTunnelStatus

// This is an alias to an internal value.
const CloudflareTunnelStatusInactive = shared.CloudflareTunnelStatusInactive

// This is an alias to an internal value.
const CloudflareTunnelStatusDegraded = shared.CloudflareTunnelStatusDegraded

// This is an alias to an internal value.
const CloudflareTunnelStatusHealthy = shared.CloudflareTunnelStatusHealthy

// This is an alias to an internal value.
const CloudflareTunnelStatusDown = shared.CloudflareTunnelStatusDown

// The type of tunnel.
//
// This is an alias to an internal type.
type CloudflareTunnelTunType = shared.CloudflareTunnelTunType

// This is an alias to an internal value.
const CloudflareTunnelTunTypeCfdTunnel = shared.CloudflareTunnelTunTypeCfdTunnel

// This is an alias to an internal value.
const CloudflareTunnelTunTypeWARPConnector = shared.CloudflareTunnelTunTypeWARPConnector

// This is an alias to an internal value.
const CloudflareTunnelTunTypeWARP = shared.CloudflareTunnelTunTypeWARP

// This is an alias to an internal value.
const CloudflareTunnelTunTypeMagic = shared.CloudflareTunnelTunTypeMagic

// This is an alias to an internal value.
const CloudflareTunnelTunTypeIPSec = shared.CloudflareTunnelTunTypeIPSec

// This is an alias to an internal value.
const CloudflareTunnelTunTypeGRE = shared.CloudflareTunnelTunTypeGRE

// This is an alias to an internal value.
const CloudflareTunnelTunTypeCNI = shared.CloudflareTunnelTunTypeCNI

// This is an alias to an internal type.
type ErrorData = shared.ErrorData

// This is an alias to an internal type.
type ErrorDataSource = shared.ErrorDataSource

// This is an alias to an internal type.
type Member = shared.Member

// This is an alias to an internal type.
type MemberPolicy = shared.MemberPolicy

// Allow or deny operations against the resources.
//
// This is an alias to an internal type.
type MemberPoliciesAccess = shared.MemberPoliciesAccess

// This is an alias to an internal value.
const MemberPoliciesAccessAllow = shared.MemberPoliciesAccessAllow

// This is an alias to an internal value.
const MemberPoliciesAccessDeny = shared.MemberPoliciesAccessDeny

// A named group of permissions that map to a group of operations against
// resources.
//
// This is an alias to an internal type.
type MemberPoliciesPermissionGroup = shared.MemberPoliciesPermissionGroup

// Attributes associated to the permission group.
//
// This is an alias to an internal type.
type MemberPoliciesPermissionGroupsMeta = shared.MemberPoliciesPermissionGroupsMeta

// A group of scoped resources.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroup = shared.MemberPoliciesResourceGroup

// A scope is a combination of scope objects which provides additional context.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroupsScope = shared.MemberPoliciesResourceGroupsScope

// A scope object represents any resource that can have actions applied against
// invite.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroupsScopeObject = shared.MemberPoliciesResourceGroupsScopeObject

// Attributes associated to the resource group.
//
// This is an alias to an internal type.
type MemberPoliciesResourceGroupsMeta = shared.MemberPoliciesResourceGroupsMeta

// A member's status in the account.
//
// This is an alias to an internal type.
type MemberStatus = shared.MemberStatus

// This is an alias to an internal value.
const MemberStatusAccepted = shared.MemberStatusAccepted

// This is an alias to an internal value.
const MemberStatusPending = shared.MemberStatusPending

// Details of the user associated to the membership.
//
// This is an alias to an internal type.
type MemberUser = shared.MemberUser

// This is an alias to an internal type.
type Permission = shared.Permission

// This is an alias to an internal type.
type PermissionGrant = shared.PermissionGrant

// This is an alias to an internal type.
type PermissionGrantParam = shared.PermissionGrantParam

// The rate plan applied to the subscription.
//
// This is an alias to an internal type.
type RatePlan = shared.RatePlan

// The ID of the rate plan.
//
// This is an alias to an internal type.
type RatePlanID = shared.RatePlanID

// This is an alias to an internal value.
const RatePlanIDFree = shared.RatePlanIDFree

// This is an alias to an internal value.
const RatePlanIDLite = shared.RatePlanIDLite

// This is an alias to an internal value.
const RatePlanIDPro = shared.RatePlanIDPro

// This is an alias to an internal value.
const RatePlanIDProPlus = shared.RatePlanIDProPlus

// This is an alias to an internal value.
const RatePlanIDBusiness = shared.RatePlanIDBusiness

// This is an alias to an internal value.
const RatePlanIDEnterprise = shared.RatePlanIDEnterprise

// This is an alias to an internal value.
const RatePlanIDPartnersFree = shared.RatePlanIDPartnersFree

// This is an alias to an internal value.
const RatePlanIDPartnersPro = shared.RatePlanIDPartnersPro

// This is an alias to an internal value.
const RatePlanIDPartnersBusiness = shared.RatePlanIDPartnersBusiness

// This is an alias to an internal value.
const RatePlanIDPartnersEnterprise = shared.RatePlanIDPartnersEnterprise

// The rate plan applied to the subscription.
//
// This is an alias to an internal type.
type RatePlanParam = shared.RatePlanParam

// This is an alias to an internal type.
type ResponseInfo = shared.ResponseInfo

// This is an alias to an internal type.
type ResponseInfoSource = shared.ResponseInfoSource

// This is an alias to an internal type.
type Role = shared.Role

// This is an alias to an internal type.
type RolePermissions = shared.RolePermissions

// This is an alias to an internal type.
type RoleParam = shared.RoleParam

// This is an alias to an internal type.
type RolePermissionsParam = shared.RolePermissionsParam

// Direction to order DNS records in.
//
// This is an alias to an internal type.
type SortDirection = shared.SortDirection

// This is an alias to an internal value.
const SortDirectionAsc = shared.SortDirectionAsc

// This is an alias to an internal value.
const SortDirectionDesc = shared.SortDirectionDesc

// This is an alias to an internal type.
type Subscription = shared.Subscription

// How often the subscription is renewed automatically.
//
// This is an alias to an internal type.
type SubscriptionFrequency = shared.SubscriptionFrequency

// This is an alias to an internal value.
const SubscriptionFrequencyWeekly = shared.SubscriptionFrequencyWeekly

// This is an alias to an internal value.
const SubscriptionFrequencyMonthly = shared.SubscriptionFrequencyMonthly

// This is an alias to an internal value.
const SubscriptionFrequencyQuarterly = shared.SubscriptionFrequencyQuarterly

// This is an alias to an internal value.
const SubscriptionFrequencyYearly = shared.SubscriptionFrequencyYearly

// The state that the subscription is in.
//
// This is an alias to an internal type.
type SubscriptionState = shared.SubscriptionState

// This is an alias to an internal value.
const SubscriptionStateTrial = shared.SubscriptionStateTrial

// This is an alias to an internal value.
const SubscriptionStateProvisioned = shared.SubscriptionStateProvisioned

// This is an alias to an internal value.
const SubscriptionStatePaid = shared.SubscriptionStatePaid

// This is an alias to an internal value.
const SubscriptionStateAwaitingPayment = shared.SubscriptionStateAwaitingPayment

// This is an alias to an internal value.
const SubscriptionStateCancelled = shared.SubscriptionStateCancelled

// This is an alias to an internal value.
const SubscriptionStateFailed = shared.SubscriptionStateFailed

// This is an alias to an internal value.
const SubscriptionStateExpired = shared.SubscriptionStateExpired

// This is an alias to an internal type.
type SubscriptionParam = shared.SubscriptionParam

// This is an alias to an internal type.
type Token = shared.Token

// This is an alias to an internal type.
type TokenCondition = shared.TokenCondition

// Client IP restrictions.
//
// This is an alias to an internal type.
type TokenConditionRequestIP = shared.TokenConditionRequestIP

// Status of the token.
//
// This is an alias to an internal type.
type TokenStatus = shared.TokenStatus

// This is an alias to an internal value.
const TokenStatusActive = shared.TokenStatusActive

// This is an alias to an internal value.
const TokenStatusDisabled = shared.TokenStatusDisabled

// This is an alias to an internal value.
const TokenStatusExpired = shared.TokenStatusExpired

// This is an alias to an internal type.
type TokenParam = shared.TokenParam

// This is an alias to an internal type.
type TokenConditionParam = shared.TokenConditionParam

// Client IP restrictions.
//
// This is an alias to an internal type.
type TokenConditionRequestIPParam = shared.TokenConditionRequestIPParam

// IPv4/IPv6 CIDR.
//
// This is an alias to an internal type.
type TokenConditionCIDRList = shared.TokenConditionCIDRList

// IPv4/IPv6 CIDR.
//
// This is an alias to an internal type.
type TokenConditionCIDRListParam = shared.TokenConditionCIDRListParam

// This is an alias to an internal type.
type TokenPolicy = shared.TokenPolicy

// Allow or deny operations against the resources.
//
// This is an alias to an internal type.
type TokenPolicyEffect = shared.TokenPolicyEffect

// This is an alias to an internal value.
const TokenPolicyEffectAllow = shared.TokenPolicyEffectAllow

// This is an alias to an internal value.
const TokenPolicyEffectDeny = shared.TokenPolicyEffectDeny

// A named group of permissions that map to a group of operations against
// resources.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroup = shared.TokenPolicyPermissionGroup

// Attributes associated to the permission group.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroupsMeta = shared.TokenPolicyPermissionGroupsMeta

// A list of resource names that the policy applies to.
//
// This is an alias to an internal type.
type TokenPolicyResourcesUnion = shared.TokenPolicyResourcesUnion

// Map of simple string resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectString = shared.TokenPolicyResourcesIAMResourcesTypeObjectString

// Map of nested resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectNested = shared.TokenPolicyResourcesIAMResourcesTypeObjectNested

// This is an alias to an internal type.
type TokenPolicyParam = shared.TokenPolicyParam

// A named group of permissions that map to a group of operations against
// resources.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroupParam = shared.TokenPolicyPermissionGroupParam

// Attributes associated to the permission group.
//
// This is an alias to an internal type.
type TokenPolicyPermissionGroupsMetaParam = shared.TokenPolicyPermissionGroupsMetaParam

// A list of resource names that the policy applies to.
//
// This is an alias to an internal type.
type TokenPolicyResourcesUnionParam = shared.TokenPolicyResourcesUnionParam

// Map of simple string resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectStringParam = shared.TokenPolicyResourcesIAMResourcesTypeObjectStringParam

// Map of nested resource permissions
//
// This is an alias to an internal type.
type TokenPolicyResourcesIAMResourcesTypeObjectNestedParam = shared.TokenPolicyResourcesIAMResourcesTypeObjectNestedParam

// The token value.
//
// This is an alias to an internal type.
type TokenValue = shared.TokenValue
