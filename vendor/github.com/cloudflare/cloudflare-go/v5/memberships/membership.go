// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package memberships

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/accounts"
	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// MembershipService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMembershipService] method instead.
type MembershipService struct {
	Options []option.RequestOption
}

// NewMembershipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMembershipService(opts ...option.RequestOption) (r *MembershipService) {
	r = &MembershipService{}
	r.Options = opts
	return
}

// Accept or reject this account invitation.
func (r *MembershipService) Update(ctx context.Context, membershipID string, body MembershipUpdateParams, opts ...option.RequestOption) (res *MembershipUpdateResponse, err error) {
	var env MembershipUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List memberships of accounts the user can access.
func (r *MembershipService) List(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Membership], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "memberships"
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

// List memberships of accounts the user can access.
func (r *MembershipService) ListAutoPaging(ctx context.Context, query MembershipListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Membership] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, query, opts...))
}

// Remove the associated member from an account.
func (r *MembershipService) Delete(ctx context.Context, membershipID string, opts ...option.RequestOption) (res *MembershipDeleteResponse, err error) {
	var env MembershipDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a specific membership.
func (r *MembershipService) Get(ctx context.Context, membershipID string, opts ...option.RequestOption) (res *MembershipGetResponse, err error) {
	var env MembershipGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if membershipID == "" {
		err = errors.New("missing required membership_id parameter")
		return
	}
	path := fmt.Sprintf("memberships/%s", membershipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Membership struct {
	// Membership identifier tag.
	ID      string           `json:"id"`
	Account accounts.Account `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// All access permissions for the user at the account.
	Permissions MembershipPermissions `json:"permissions"`
	// List of role names the membership has for this account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status MembershipStatus `json:"status"`
	JSON   membershipJSON   `json:"-"`
}

// membershipJSON contains the JSON metadata for the struct [Membership]
type membershipJSON struct {
	ID               apijson.Field
	Account          apijson.Field
	APIAccessEnabled apijson.Field
	Permissions      apijson.Field
	Roles            apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Membership) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipJSON) RawJSON() string {
	return r.raw
}

// All access permissions for the user at the account.
type MembershipPermissions struct {
	Analytics    shared.PermissionGrant    `json:"analytics"`
	Billing      shared.PermissionGrant    `json:"billing"`
	CachePurge   shared.PermissionGrant    `json:"cache_purge"`
	DNS          shared.PermissionGrant    `json:"dns"`
	DNSRecords   shared.PermissionGrant    `json:"dns_records"`
	LB           shared.PermissionGrant    `json:"lb"`
	Logs         shared.PermissionGrant    `json:"logs"`
	Organization shared.PermissionGrant    `json:"organization"`
	SSL          shared.PermissionGrant    `json:"ssl"`
	WAF          shared.PermissionGrant    `json:"waf"`
	ZoneSettings shared.PermissionGrant    `json:"zone_settings"`
	Zones        shared.PermissionGrant    `json:"zones"`
	JSON         membershipPermissionsJSON `json:"-"`
}

// membershipPermissionsJSON contains the JSON metadata for the struct
// [MembershipPermissions]
type membershipPermissionsJSON struct {
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

func (r *MembershipPermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipPermissionsJSON) RawJSON() string {
	return r.raw
}

// Status of this membership.
type MembershipStatus string

const (
	MembershipStatusAccepted MembershipStatus = "accepted"
	MembershipStatusPending  MembershipStatus = "pending"
	MembershipStatusRejected MembershipStatus = "rejected"
)

func (r MembershipStatus) IsKnown() bool {
	switch r {
	case MembershipStatusAccepted, MembershipStatusPending, MembershipStatusRejected:
		return true
	}
	return false
}

type MembershipUpdateResponse struct {
	// Membership identifier tag.
	ID      string           `json:"id"`
	Account accounts.Account `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// All access permissions for the user at the account.
	Permissions MembershipUpdateResponsePermissions `json:"permissions"`
	// Access policy for the membership
	Policies []MembershipUpdateResponsePolicy `json:"policies"`
	// List of role names the membership has for this account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status MembershipUpdateResponseStatus `json:"status"`
	JSON   membershipUpdateResponseJSON   `json:"-"`
}

// membershipUpdateResponseJSON contains the JSON metadata for the struct
// [MembershipUpdateResponse]
type membershipUpdateResponseJSON struct {
	ID               apijson.Field
	Account          apijson.Field
	APIAccessEnabled apijson.Field
	Permissions      apijson.Field
	Policies         apijson.Field
	Roles            apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// All access permissions for the user at the account.
type MembershipUpdateResponsePermissions struct {
	Analytics    shared.PermissionGrant                  `json:"analytics"`
	Billing      shared.PermissionGrant                  `json:"billing"`
	CachePurge   shared.PermissionGrant                  `json:"cache_purge"`
	DNS          shared.PermissionGrant                  `json:"dns"`
	DNSRecords   shared.PermissionGrant                  `json:"dns_records"`
	LB           shared.PermissionGrant                  `json:"lb"`
	Logs         shared.PermissionGrant                  `json:"logs"`
	Organization shared.PermissionGrant                  `json:"organization"`
	SSL          shared.PermissionGrant                  `json:"ssl"`
	WAF          shared.PermissionGrant                  `json:"waf"`
	ZoneSettings shared.PermissionGrant                  `json:"zone_settings"`
	Zones        shared.PermissionGrant                  `json:"zones"`
	JSON         membershipUpdateResponsePermissionsJSON `json:"-"`
}

// membershipUpdateResponsePermissionsJSON contains the JSON metadata for the
// struct [MembershipUpdateResponsePermissions]
type membershipUpdateResponsePermissionsJSON struct {
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

func (r *MembershipUpdateResponsePermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePermissionsJSON) RawJSON() string {
	return r.raw
}

type MembershipUpdateResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access MembershipUpdateResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []MembershipUpdateResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []MembershipUpdateResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           membershipUpdateResponsePolicyJSON              `json:"-"`
}

// membershipUpdateResponsePolicyJSON contains the JSON metadata for the struct
// [MembershipUpdateResponsePolicy]
type membershipUpdateResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipUpdateResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type MembershipUpdateResponsePoliciesAccess string

const (
	MembershipUpdateResponsePoliciesAccessAllow MembershipUpdateResponsePoliciesAccess = "allow"
	MembershipUpdateResponsePoliciesAccessDeny  MembershipUpdateResponsePoliciesAccess = "deny"
)

func (r MembershipUpdateResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case MembershipUpdateResponsePoliciesAccessAllow, MembershipUpdateResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type MembershipUpdateResponsePoliciesPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta MembershipUpdateResponsePoliciesPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                                              `json:"name"`
	JSON membershipUpdateResponsePoliciesPermissionGroupJSON `json:"-"`
}

// membershipUpdateResponsePoliciesPermissionGroupJSON contains the JSON metadata
// for the struct [MembershipUpdateResponsePoliciesPermissionGroup]
type membershipUpdateResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type MembershipUpdateResponsePoliciesPermissionGroupsMeta struct {
	Key   string                                                   `json:"key"`
	Value string                                                   `json:"value"`
	JSON  membershipUpdateResponsePoliciesPermissionGroupsMetaJSON `json:"-"`
}

// membershipUpdateResponsePoliciesPermissionGroupsMetaJSON contains the JSON
// metadata for the struct [MembershipUpdateResponsePoliciesPermissionGroupsMeta]
type membershipUpdateResponsePoliciesPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponsePoliciesPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePoliciesPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type MembershipUpdateResponsePoliciesResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []MembershipUpdateResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta MembershipUpdateResponsePoliciesResourceGroupsMeta `json:"meta"`
	// Name of the resource group.
	Name string                                            `json:"name"`
	JSON membershipUpdateResponsePoliciesResourceGroupJSON `json:"-"`
}

// membershipUpdateResponsePoliciesResourceGroupJSON contains the JSON metadata for
// the struct [MembershipUpdateResponsePoliciesResourceGroup]
type membershipUpdateResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type MembershipUpdateResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []MembershipUpdateResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    membershipUpdateResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// membershipUpdateResponsePoliciesResourceGroupsScopeJSON contains the JSON
// metadata for the struct [MembershipUpdateResponsePoliciesResourceGroupsScope]
type membershipUpdateResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type MembershipUpdateResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                        `json:"key,required"`
	JSON membershipUpdateResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// membershipUpdateResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct
// [MembershipUpdateResponsePoliciesResourceGroupsScopeObject]
type membershipUpdateResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type MembershipUpdateResponsePoliciesResourceGroupsMeta struct {
	Key   string                                                 `json:"key"`
	Value string                                                 `json:"value"`
	JSON  membershipUpdateResponsePoliciesResourceGroupsMetaJSON `json:"-"`
}

// membershipUpdateResponsePoliciesResourceGroupsMetaJSON contains the JSON
// metadata for the struct [MembershipUpdateResponsePoliciesResourceGroupsMeta]
type membershipUpdateResponsePoliciesResourceGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponsePoliciesResourceGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponsePoliciesResourceGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// Status of this membership.
type MembershipUpdateResponseStatus string

const (
	MembershipUpdateResponseStatusAccepted MembershipUpdateResponseStatus = "accepted"
	MembershipUpdateResponseStatusPending  MembershipUpdateResponseStatus = "pending"
	MembershipUpdateResponseStatusRejected MembershipUpdateResponseStatus = "rejected"
)

func (r MembershipUpdateResponseStatus) IsKnown() bool {
	switch r {
	case MembershipUpdateResponseStatusAccepted, MembershipUpdateResponseStatusPending, MembershipUpdateResponseStatusRejected:
		return true
	}
	return false
}

type MembershipDeleteResponse struct {
	// Membership identifier tag.
	ID   string                       `json:"id"`
	JSON membershipDeleteResponseJSON `json:"-"`
}

// membershipDeleteResponseJSON contains the JSON metadata for the struct
// [MembershipDeleteResponse]
type membershipDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MembershipGetResponse struct {
	// Membership identifier tag.
	ID      string           `json:"id"`
	Account accounts.Account `json:"account"`
	// Enterprise only. Indicates whether or not API access is enabled specifically for
	// this user on a given account.
	APIAccessEnabled bool `json:"api_access_enabled,nullable"`
	// All access permissions for the user at the account.
	Permissions MembershipGetResponsePermissions `json:"permissions"`
	// Access policy for the membership
	Policies []MembershipGetResponsePolicy `json:"policies"`
	// List of role names the membership has for this account.
	Roles []string `json:"roles"`
	// Status of this membership.
	Status MembershipGetResponseStatus `json:"status"`
	JSON   membershipGetResponseJSON   `json:"-"`
}

// membershipGetResponseJSON contains the JSON metadata for the struct
// [MembershipGetResponse]
type membershipGetResponseJSON struct {
	ID               apijson.Field
	Account          apijson.Field
	APIAccessEnabled apijson.Field
	Permissions      apijson.Field
	Policies         apijson.Field
	Roles            apijson.Field
	Status           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponseJSON) RawJSON() string {
	return r.raw
}

// All access permissions for the user at the account.
type MembershipGetResponsePermissions struct {
	Analytics    shared.PermissionGrant               `json:"analytics"`
	Billing      shared.PermissionGrant               `json:"billing"`
	CachePurge   shared.PermissionGrant               `json:"cache_purge"`
	DNS          shared.PermissionGrant               `json:"dns"`
	DNSRecords   shared.PermissionGrant               `json:"dns_records"`
	LB           shared.PermissionGrant               `json:"lb"`
	Logs         shared.PermissionGrant               `json:"logs"`
	Organization shared.PermissionGrant               `json:"organization"`
	SSL          shared.PermissionGrant               `json:"ssl"`
	WAF          shared.PermissionGrant               `json:"waf"`
	ZoneSettings shared.PermissionGrant               `json:"zone_settings"`
	Zones        shared.PermissionGrant               `json:"zones"`
	JSON         membershipGetResponsePermissionsJSON `json:"-"`
}

// membershipGetResponsePermissionsJSON contains the JSON metadata for the struct
// [MembershipGetResponsePermissions]
type membershipGetResponsePermissionsJSON struct {
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

func (r *MembershipGetResponsePermissions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePermissionsJSON) RawJSON() string {
	return r.raw
}

type MembershipGetResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access MembershipGetResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []MembershipGetResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []MembershipGetResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           membershipGetResponsePolicyJSON              `json:"-"`
}

// membershipGetResponsePolicyJSON contains the JSON metadata for the struct
// [MembershipGetResponsePolicy]
type membershipGetResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipGetResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type MembershipGetResponsePoliciesAccess string

const (
	MembershipGetResponsePoliciesAccessAllow MembershipGetResponsePoliciesAccess = "allow"
	MembershipGetResponsePoliciesAccessDeny  MembershipGetResponsePoliciesAccess = "deny"
)

func (r MembershipGetResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case MembershipGetResponsePoliciesAccessAllow, MembershipGetResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type MembershipGetResponsePoliciesPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta MembershipGetResponsePoliciesPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                                           `json:"name"`
	JSON membershipGetResponsePoliciesPermissionGroupJSON `json:"-"`
}

// membershipGetResponsePoliciesPermissionGroupJSON contains the JSON metadata for
// the struct [MembershipGetResponsePoliciesPermissionGroup]
type membershipGetResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type MembershipGetResponsePoliciesPermissionGroupsMeta struct {
	Key   string                                                `json:"key"`
	Value string                                                `json:"value"`
	JSON  membershipGetResponsePoliciesPermissionGroupsMetaJSON `json:"-"`
}

// membershipGetResponsePoliciesPermissionGroupsMetaJSON contains the JSON metadata
// for the struct [MembershipGetResponsePoliciesPermissionGroupsMeta]
type membershipGetResponsePoliciesPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponsePoliciesPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePoliciesPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type MembershipGetResponsePoliciesResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []MembershipGetResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta MembershipGetResponsePoliciesResourceGroupsMeta `json:"meta"`
	// Name of the resource group.
	Name string                                         `json:"name"`
	JSON membershipGetResponsePoliciesResourceGroupJSON `json:"-"`
}

// membershipGetResponsePoliciesResourceGroupJSON contains the JSON metadata for
// the struct [MembershipGetResponsePoliciesResourceGroup]
type membershipGetResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type MembershipGetResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []MembershipGetResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    membershipGetResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// membershipGetResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata
// for the struct [MembershipGetResponsePoliciesResourceGroupsScope]
type membershipGetResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type MembershipGetResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                     `json:"key,required"`
	JSON membershipGetResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// membershipGetResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [MembershipGetResponsePoliciesResourceGroupsScopeObject]
type membershipGetResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type MembershipGetResponsePoliciesResourceGroupsMeta struct {
	Key   string                                              `json:"key"`
	Value string                                              `json:"value"`
	JSON  membershipGetResponsePoliciesResourceGroupsMetaJSON `json:"-"`
}

// membershipGetResponsePoliciesResourceGroupsMetaJSON contains the JSON metadata
// for the struct [MembershipGetResponsePoliciesResourceGroupsMeta]
type membershipGetResponsePoliciesResourceGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponsePoliciesResourceGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponsePoliciesResourceGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// Status of this membership.
type MembershipGetResponseStatus string

const (
	MembershipGetResponseStatusAccepted MembershipGetResponseStatus = "accepted"
	MembershipGetResponseStatusPending  MembershipGetResponseStatus = "pending"
	MembershipGetResponseStatusRejected MembershipGetResponseStatus = "rejected"
)

func (r MembershipGetResponseStatus) IsKnown() bool {
	switch r {
	case MembershipGetResponseStatusAccepted, MembershipGetResponseStatusPending, MembershipGetResponseStatusRejected:
		return true
	}
	return false
}

type MembershipUpdateParams struct {
	// Whether to accept or reject this account invitation.
	Status param.Field[MembershipUpdateParamsStatus] `json:"status,required"`
}

func (r MembershipUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to accept or reject this account invitation.
type MembershipUpdateParamsStatus string

const (
	MembershipUpdateParamsStatusAccepted MembershipUpdateParamsStatus = "accepted"
	MembershipUpdateParamsStatusRejected MembershipUpdateParamsStatus = "rejected"
)

func (r MembershipUpdateParamsStatus) IsKnown() bool {
	switch r {
	case MembershipUpdateParamsStatusAccepted, MembershipUpdateParamsStatusRejected:
		return true
	}
	return false
}

type MembershipUpdateResponseEnvelope struct {
	Errors   []MembershipUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MembershipUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success MembershipUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  MembershipUpdateResponse                `json:"result"`
	JSON    membershipUpdateResponseEnvelopeJSON    `json:"-"`
}

// membershipUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MembershipUpdateResponseEnvelope]
type membershipUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MembershipUpdateResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           MembershipUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             membershipUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// membershipUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MembershipUpdateResponseEnvelopeErrors]
type membershipUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MembershipUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    membershipUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// membershipUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [MembershipUpdateResponseEnvelopeErrorsSource]
type membershipUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MembershipUpdateResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           MembershipUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             membershipUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// membershipUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MembershipUpdateResponseEnvelopeMessages]
type membershipUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MembershipUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    membershipUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// membershipUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [MembershipUpdateResponseEnvelopeMessagesSource]
type membershipUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MembershipUpdateResponseEnvelopeSuccess bool

const (
	MembershipUpdateResponseEnvelopeSuccessTrue MembershipUpdateResponseEnvelopeSuccess = true
)

func (r MembershipUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MembershipUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MembershipListParams struct {
	Account param.Field[MembershipListParamsAccount] `query:"account"`
	// Direction to order memberships.
	Direction param.Field[MembershipListParamsDirection] `query:"direction"`
	// Account name
	Name param.Field[string] `query:"name"`
	// Field to order memberships by.
	Order param.Field[MembershipListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of memberships per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Status of this membership.
	Status param.Field[MembershipListParamsStatus] `query:"status"`
}

// URLQuery serializes [MembershipListParams]'s query parameters as `url.Values`.
func (r MembershipListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type MembershipListParamsAccount struct {
	// Account name
	Name param.Field[string] `query:"name"`
}

// URLQuery serializes [MembershipListParamsAccount]'s query parameters as
// `url.Values`.
func (r MembershipListParamsAccount) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order memberships.
type MembershipListParamsDirection string

const (
	MembershipListParamsDirectionAsc  MembershipListParamsDirection = "asc"
	MembershipListParamsDirectionDesc MembershipListParamsDirection = "desc"
)

func (r MembershipListParamsDirection) IsKnown() bool {
	switch r {
	case MembershipListParamsDirectionAsc, MembershipListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order memberships by.
type MembershipListParamsOrder string

const (
	MembershipListParamsOrderID          MembershipListParamsOrder = "id"
	MembershipListParamsOrderAccountName MembershipListParamsOrder = "account.name"
	MembershipListParamsOrderStatus      MembershipListParamsOrder = "status"
)

func (r MembershipListParamsOrder) IsKnown() bool {
	switch r {
	case MembershipListParamsOrderID, MembershipListParamsOrderAccountName, MembershipListParamsOrderStatus:
		return true
	}
	return false
}

// Status of this membership.
type MembershipListParamsStatus string

const (
	MembershipListParamsStatusAccepted MembershipListParamsStatus = "accepted"
	MembershipListParamsStatusPending  MembershipListParamsStatus = "pending"
	MembershipListParamsStatusRejected MembershipListParamsStatus = "rejected"
)

func (r MembershipListParamsStatus) IsKnown() bool {
	switch r {
	case MembershipListParamsStatusAccepted, MembershipListParamsStatusPending, MembershipListParamsStatusRejected:
		return true
	}
	return false
}

type MembershipDeleteResponseEnvelope struct {
	Errors   []MembershipDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MembershipDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success MembershipDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  MembershipDeleteResponse                `json:"result"`
	JSON    membershipDeleteResponseEnvelopeJSON    `json:"-"`
}

// membershipDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MembershipDeleteResponseEnvelope]
type membershipDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MembershipDeleteResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           MembershipDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             membershipDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// membershipDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MembershipDeleteResponseEnvelopeErrors]
type membershipDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MembershipDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    membershipDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// membershipDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [MembershipDeleteResponseEnvelopeErrorsSource]
type membershipDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MembershipDeleteResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           MembershipDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             membershipDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// membershipDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MembershipDeleteResponseEnvelopeMessages]
type membershipDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MembershipDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    membershipDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// membershipDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [MembershipDeleteResponseEnvelopeMessagesSource]
type membershipDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MembershipDeleteResponseEnvelopeSuccess bool

const (
	MembershipDeleteResponseEnvelopeSuccessTrue MembershipDeleteResponseEnvelopeSuccess = true
)

func (r MembershipDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MembershipDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MembershipGetResponseEnvelope struct {
	Errors   []MembershipGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MembershipGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success MembershipGetResponseEnvelopeSuccess `json:"success,required"`
	Result  MembershipGetResponse                `json:"result"`
	JSON    membershipGetResponseEnvelopeJSON    `json:"-"`
}

// membershipGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MembershipGetResponseEnvelope]
type membershipGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MembershipGetResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           MembershipGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             membershipGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// membershipGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [MembershipGetResponseEnvelopeErrors]
type membershipGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MembershipGetResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    membershipGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// membershipGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [MembershipGetResponseEnvelopeErrorsSource]
type membershipGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MembershipGetResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           MembershipGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             membershipGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// membershipGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MembershipGetResponseEnvelopeMessages]
type membershipGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MembershipGetResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    membershipGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// membershipGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [MembershipGetResponseEnvelopeMessagesSource]
type membershipGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MembershipGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r membershipGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MembershipGetResponseEnvelopeSuccess bool

const (
	MembershipGetResponseEnvelopeSuccessTrue MembershipGetResponseEnvelopeSuccess = true
)

func (r MembershipGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MembershipGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
