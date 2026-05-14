// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

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

// UserGroupService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserGroupService] method instead.
type UserGroupService struct {
	Options []option.RequestOption
	Members *UserGroupMemberService
}

// NewUserGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserGroupService(opts ...option.RequestOption) (r *UserGroupService) {
	r = &UserGroupService{}
	r.Options = opts
	r.Members = NewUserGroupMemberService(opts...)
	return
}

// Create a new user group under the specified account.
func (r *UserGroupService) New(ctx context.Context, params UserGroupNewParams, opts ...option.RequestOption) (res *UserGroupNewResponse, err error) {
	var env UserGroupNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an existing user group.
func (r *UserGroupService) Update(ctx context.Context, userGroupID string, params UserGroupUpdateParams, opts ...option.RequestOption) (res *UserGroupUpdateResponse, err error) {
	var env UserGroupUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s", params.AccountID, userGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all the user groups for an account.
func (r *UserGroupService) List(ctx context.Context, params UserGroupListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UserGroupListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups", params.AccountID)
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

// List all the user groups for an account.
func (r *UserGroupService) ListAutoPaging(ctx context.Context, params UserGroupListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UserGroupListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a user group from an account.
func (r *UserGroupService) Delete(ctx context.Context, userGroupID string, body UserGroupDeleteParams, opts ...option.RequestOption) (res *UserGroupDeleteResponse, err error) {
	var env UserGroupDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s", body.AccountID, userGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific user group in an account.
func (r *UserGroupService) Get(ctx context.Context, userGroupID string, query UserGroupGetParams, opts ...option.RequestOption) (res *UserGroupGetResponse, err error) {
	var env UserGroupGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s", query.AccountID, userGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A group of policies resources.
type UserGroupNewResponse struct {
	// User Group identifier tag.
	ID string `json:"id,required"`
	// Timestamp for the creation of the user group
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Last time the user group was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Name of the user group.
	Name string `json:"name,required"`
	// Policies attached to the User group
	Policies []UserGroupNewResponsePolicy `json:"policies"`
	JSON     userGroupNewResponseJSON     `json:"-"`
}

// userGroupNewResponseJSON contains the JSON metadata for the struct
// [UserGroupNewResponse]
type userGroupNewResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

// Policy
type UserGroupNewResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access UserGroupNewResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []UserGroupNewResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []UserGroupNewResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           userGroupNewResponsePolicyJSON              `json:"-"`
}

// userGroupNewResponsePolicyJSON contains the JSON metadata for the struct
// [UserGroupNewResponsePolicy]
type userGroupNewResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupNewResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type UserGroupNewResponsePoliciesAccess string

const (
	UserGroupNewResponsePoliciesAccessAllow UserGroupNewResponsePoliciesAccess = "allow"
	UserGroupNewResponsePoliciesAccessDeny  UserGroupNewResponsePoliciesAccess = "deny"
)

func (r UserGroupNewResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case UserGroupNewResponsePoliciesAccessAllow, UserGroupNewResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type UserGroupNewResponsePoliciesPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta UserGroupNewResponsePoliciesPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                                          `json:"name"`
	JSON userGroupNewResponsePoliciesPermissionGroupJSON `json:"-"`
}

// userGroupNewResponsePoliciesPermissionGroupJSON contains the JSON metadata for
// the struct [UserGroupNewResponsePoliciesPermissionGroup]
type userGroupNewResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type UserGroupNewResponsePoliciesPermissionGroupsMeta struct {
	Key   string                                               `json:"key"`
	Value string                                               `json:"value"`
	JSON  userGroupNewResponsePoliciesPermissionGroupsMetaJSON `json:"-"`
}

// userGroupNewResponsePoliciesPermissionGroupsMetaJSON contains the JSON metadata
// for the struct [UserGroupNewResponsePoliciesPermissionGroupsMeta]
type userGroupNewResponsePoliciesPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponsePoliciesPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponsePoliciesPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type UserGroupNewResponsePoliciesResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []UserGroupNewResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta UserGroupNewResponsePoliciesResourceGroupsMeta `json:"meta"`
	// Name of the resource group.
	Name string                                        `json:"name"`
	JSON userGroupNewResponsePoliciesResourceGroupJSON `json:"-"`
}

// userGroupNewResponsePoliciesResourceGroupJSON contains the JSON metadata for the
// struct [UserGroupNewResponsePoliciesResourceGroup]
type userGroupNewResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type UserGroupNewResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []UserGroupNewResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    userGroupNewResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// userGroupNewResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata
// for the struct [UserGroupNewResponsePoliciesResourceGroupsScope]
type userGroupNewResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type UserGroupNewResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                    `json:"key,required"`
	JSON userGroupNewResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// userGroupNewResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [UserGroupNewResponsePoliciesResourceGroupsScopeObject]
type userGroupNewResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type UserGroupNewResponsePoliciesResourceGroupsMeta struct {
	Key   string                                             `json:"key"`
	Value string                                             `json:"value"`
	JSON  userGroupNewResponsePoliciesResourceGroupsMetaJSON `json:"-"`
}

// userGroupNewResponsePoliciesResourceGroupsMetaJSON contains the JSON metadata
// for the struct [UserGroupNewResponsePoliciesResourceGroupsMeta]
type userGroupNewResponsePoliciesResourceGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponsePoliciesResourceGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponsePoliciesResourceGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of policies resources.
type UserGroupUpdateResponse struct {
	// User Group identifier tag.
	ID string `json:"id,required"`
	// Timestamp for the creation of the user group
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Last time the user group was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Name of the user group.
	Name string `json:"name,required"`
	// Policies attached to the User group
	Policies []UserGroupUpdateResponsePolicy `json:"policies"`
	JSON     userGroupUpdateResponseJSON     `json:"-"`
}

// userGroupUpdateResponseJSON contains the JSON metadata for the struct
// [UserGroupUpdateResponse]
type userGroupUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// Policy
type UserGroupUpdateResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access UserGroupUpdateResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []UserGroupUpdateResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []UserGroupUpdateResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           userGroupUpdateResponsePolicyJSON              `json:"-"`
}

// userGroupUpdateResponsePolicyJSON contains the JSON metadata for the struct
// [UserGroupUpdateResponsePolicy]
type userGroupUpdateResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupUpdateResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type UserGroupUpdateResponsePoliciesAccess string

const (
	UserGroupUpdateResponsePoliciesAccessAllow UserGroupUpdateResponsePoliciesAccess = "allow"
	UserGroupUpdateResponsePoliciesAccessDeny  UserGroupUpdateResponsePoliciesAccess = "deny"
)

func (r UserGroupUpdateResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case UserGroupUpdateResponsePoliciesAccessAllow, UserGroupUpdateResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type UserGroupUpdateResponsePoliciesPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta UserGroupUpdateResponsePoliciesPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                                             `json:"name"`
	JSON userGroupUpdateResponsePoliciesPermissionGroupJSON `json:"-"`
}

// userGroupUpdateResponsePoliciesPermissionGroupJSON contains the JSON metadata
// for the struct [UserGroupUpdateResponsePoliciesPermissionGroup]
type userGroupUpdateResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type UserGroupUpdateResponsePoliciesPermissionGroupsMeta struct {
	Key   string                                                  `json:"key"`
	Value string                                                  `json:"value"`
	JSON  userGroupUpdateResponsePoliciesPermissionGroupsMetaJSON `json:"-"`
}

// userGroupUpdateResponsePoliciesPermissionGroupsMetaJSON contains the JSON
// metadata for the struct [UserGroupUpdateResponsePoliciesPermissionGroupsMeta]
type userGroupUpdateResponsePoliciesPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponsePoliciesPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponsePoliciesPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type UserGroupUpdateResponsePoliciesResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []UserGroupUpdateResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta UserGroupUpdateResponsePoliciesResourceGroupsMeta `json:"meta"`
	// Name of the resource group.
	Name string                                           `json:"name"`
	JSON userGroupUpdateResponsePoliciesResourceGroupJSON `json:"-"`
}

// userGroupUpdateResponsePoliciesResourceGroupJSON contains the JSON metadata for
// the struct [UserGroupUpdateResponsePoliciesResourceGroup]
type userGroupUpdateResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type UserGroupUpdateResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []UserGroupUpdateResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    userGroupUpdateResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// userGroupUpdateResponsePoliciesResourceGroupsScopeJSON contains the JSON
// metadata for the struct [UserGroupUpdateResponsePoliciesResourceGroupsScope]
type userGroupUpdateResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type UserGroupUpdateResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                       `json:"key,required"`
	JSON userGroupUpdateResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// userGroupUpdateResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct
// [UserGroupUpdateResponsePoliciesResourceGroupsScopeObject]
type userGroupUpdateResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type UserGroupUpdateResponsePoliciesResourceGroupsMeta struct {
	Key   string                                                `json:"key"`
	Value string                                                `json:"value"`
	JSON  userGroupUpdateResponsePoliciesResourceGroupsMetaJSON `json:"-"`
}

// userGroupUpdateResponsePoliciesResourceGroupsMetaJSON contains the JSON metadata
// for the struct [UserGroupUpdateResponsePoliciesResourceGroupsMeta]
type userGroupUpdateResponsePoliciesResourceGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponsePoliciesResourceGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponsePoliciesResourceGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of policies resources.
type UserGroupListResponse struct {
	// User Group identifier tag.
	ID string `json:"id,required"`
	// Timestamp for the creation of the user group
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Last time the user group was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Name of the user group.
	Name string `json:"name,required"`
	// Policies attached to the User group
	Policies []UserGroupListResponsePolicy `json:"policies"`
	JSON     userGroupListResponseJSON     `json:"-"`
}

// userGroupListResponseJSON contains the JSON metadata for the struct
// [UserGroupListResponse]
type userGroupListResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponseJSON) RawJSON() string {
	return r.raw
}

// Policy
type UserGroupListResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access UserGroupListResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []UserGroupListResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []UserGroupListResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           userGroupListResponsePolicyJSON              `json:"-"`
}

// userGroupListResponsePolicyJSON contains the JSON metadata for the struct
// [UserGroupListResponsePolicy]
type userGroupListResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupListResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type UserGroupListResponsePoliciesAccess string

const (
	UserGroupListResponsePoliciesAccessAllow UserGroupListResponsePoliciesAccess = "allow"
	UserGroupListResponsePoliciesAccessDeny  UserGroupListResponsePoliciesAccess = "deny"
)

func (r UserGroupListResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case UserGroupListResponsePoliciesAccessAllow, UserGroupListResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type UserGroupListResponsePoliciesPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta UserGroupListResponsePoliciesPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                                           `json:"name"`
	JSON userGroupListResponsePoliciesPermissionGroupJSON `json:"-"`
}

// userGroupListResponsePoliciesPermissionGroupJSON contains the JSON metadata for
// the struct [UserGroupListResponsePoliciesPermissionGroup]
type userGroupListResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupListResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type UserGroupListResponsePoliciesPermissionGroupsMeta struct {
	Key   string                                                `json:"key"`
	Value string                                                `json:"value"`
	JSON  userGroupListResponsePoliciesPermissionGroupsMetaJSON `json:"-"`
}

// userGroupListResponsePoliciesPermissionGroupsMetaJSON contains the JSON metadata
// for the struct [UserGroupListResponsePoliciesPermissionGroupsMeta]
type userGroupListResponsePoliciesPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupListResponsePoliciesPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponsePoliciesPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type UserGroupListResponsePoliciesResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []UserGroupListResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta UserGroupListResponsePoliciesResourceGroupsMeta `json:"meta"`
	// Name of the resource group.
	Name string                                         `json:"name"`
	JSON userGroupListResponsePoliciesResourceGroupJSON `json:"-"`
}

// userGroupListResponsePoliciesResourceGroupJSON contains the JSON metadata for
// the struct [UserGroupListResponsePoliciesResourceGroup]
type userGroupListResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupListResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type UserGroupListResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []UserGroupListResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    userGroupListResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// userGroupListResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata
// for the struct [UserGroupListResponsePoliciesResourceGroupsScope]
type userGroupListResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupListResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type UserGroupListResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                     `json:"key,required"`
	JSON userGroupListResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// userGroupListResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [UserGroupListResponsePoliciesResourceGroupsScopeObject]
type userGroupListResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupListResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type UserGroupListResponsePoliciesResourceGroupsMeta struct {
	Key   string                                              `json:"key"`
	Value string                                              `json:"value"`
	JSON  userGroupListResponsePoliciesResourceGroupsMetaJSON `json:"-"`
}

// userGroupListResponsePoliciesResourceGroupsMetaJSON contains the JSON metadata
// for the struct [UserGroupListResponsePoliciesResourceGroupsMeta]
type userGroupListResponsePoliciesResourceGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupListResponsePoliciesResourceGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupListResponsePoliciesResourceGroupsMetaJSON) RawJSON() string {
	return r.raw
}

type UserGroupDeleteResponse struct {
	// Identifier
	ID   string                      `json:"id,required"`
	JSON userGroupDeleteResponseJSON `json:"-"`
}

// userGroupDeleteResponseJSON contains the JSON metadata for the struct
// [UserGroupDeleteResponse]
type userGroupDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// A group of policies resources.
type UserGroupGetResponse struct {
	// User Group identifier tag.
	ID string `json:"id,required"`
	// Timestamp for the creation of the user group
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Last time the user group was modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// Name of the user group.
	Name string `json:"name,required"`
	// Policies attached to the User group
	Policies []UserGroupGetResponsePolicy `json:"policies"`
	JSON     userGroupGetResponseJSON     `json:"-"`
}

// userGroupGetResponseJSON contains the JSON metadata for the struct
// [UserGroupGetResponse]
type userGroupGetResponseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

// Policy
type UserGroupGetResponsePolicy struct {
	// Policy identifier.
	ID string `json:"id"`
	// Allow or deny operations against the resources.
	Access UserGroupGetResponsePoliciesAccess `json:"access"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups []UserGroupGetResponsePoliciesPermissionGroup `json:"permission_groups"`
	// A list of resource groups that the policy applies to.
	ResourceGroups []UserGroupGetResponsePoliciesResourceGroup `json:"resource_groups"`
	JSON           userGroupGetResponsePolicyJSON              `json:"-"`
}

// userGroupGetResponsePolicyJSON contains the JSON metadata for the struct
// [UserGroupGetResponsePolicy]
type userGroupGetResponsePolicyJSON struct {
	ID               apijson.Field
	Access           apijson.Field
	PermissionGroups apijson.Field
	ResourceGroups   apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupGetResponsePolicy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponsePolicyJSON) RawJSON() string {
	return r.raw
}

// Allow or deny operations against the resources.
type UserGroupGetResponsePoliciesAccess string

const (
	UserGroupGetResponsePoliciesAccessAllow UserGroupGetResponsePoliciesAccess = "allow"
	UserGroupGetResponsePoliciesAccessDeny  UserGroupGetResponsePoliciesAccess = "deny"
)

func (r UserGroupGetResponsePoliciesAccess) IsKnown() bool {
	switch r {
	case UserGroupGetResponsePoliciesAccessAllow, UserGroupGetResponsePoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type UserGroupGetResponsePoliciesPermissionGroup struct {
	// Identifier of the permission group.
	ID string `json:"id,required"`
	// Attributes associated to the permission group.
	Meta UserGroupGetResponsePoliciesPermissionGroupsMeta `json:"meta"`
	// Name of the permission group.
	Name string                                          `json:"name"`
	JSON userGroupGetResponsePoliciesPermissionGroupJSON `json:"-"`
}

// userGroupGetResponsePoliciesPermissionGroupJSON contains the JSON metadata for
// the struct [UserGroupGetResponsePoliciesPermissionGroup]
type userGroupGetResponsePoliciesPermissionGroupJSON struct {
	ID          apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponsePoliciesPermissionGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponsePoliciesPermissionGroupJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the permission group.
type UserGroupGetResponsePoliciesPermissionGroupsMeta struct {
	Key   string                                               `json:"key"`
	Value string                                               `json:"value"`
	JSON  userGroupGetResponsePoliciesPermissionGroupsMetaJSON `json:"-"`
}

// userGroupGetResponsePoliciesPermissionGroupsMetaJSON contains the JSON metadata
// for the struct [UserGroupGetResponsePoliciesPermissionGroupsMeta]
type userGroupGetResponsePoliciesPermissionGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponsePoliciesPermissionGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponsePoliciesPermissionGroupsMetaJSON) RawJSON() string {
	return r.raw
}

// A group of scoped resources.
type UserGroupGetResponsePoliciesResourceGroup struct {
	// Identifier of the resource group.
	ID string `json:"id,required"`
	// The scope associated to the resource group
	Scope []UserGroupGetResponsePoliciesResourceGroupsScope `json:"scope,required"`
	// Attributes associated to the resource group.
	Meta UserGroupGetResponsePoliciesResourceGroupsMeta `json:"meta"`
	// Name of the resource group.
	Name string                                        `json:"name"`
	JSON userGroupGetResponsePoliciesResourceGroupJSON `json:"-"`
}

// userGroupGetResponsePoliciesResourceGroupJSON contains the JSON metadata for the
// struct [UserGroupGetResponsePoliciesResourceGroup]
type userGroupGetResponsePoliciesResourceGroupJSON struct {
	ID          apijson.Field
	Scope       apijson.Field
	Meta        apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponsePoliciesResourceGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponsePoliciesResourceGroupJSON) RawJSON() string {
	return r.raw
}

// A scope is a combination of scope objects which provides additional context.
type UserGroupGetResponsePoliciesResourceGroupsScope struct {
	// This is a combination of pre-defined resource name and identifier (like Account
	// ID etc.)
	Key string `json:"key,required"`
	// A list of scope objects for additional context.
	Objects []UserGroupGetResponsePoliciesResourceGroupsScopeObject `json:"objects,required"`
	JSON    userGroupGetResponsePoliciesResourceGroupsScopeJSON     `json:"-"`
}

// userGroupGetResponsePoliciesResourceGroupsScopeJSON contains the JSON metadata
// for the struct [UserGroupGetResponsePoliciesResourceGroupsScope]
type userGroupGetResponsePoliciesResourceGroupsScopeJSON struct {
	Key         apijson.Field
	Objects     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponsePoliciesResourceGroupsScope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponsePoliciesResourceGroupsScopeJSON) RawJSON() string {
	return r.raw
}

// A scope object represents any resource that can have actions applied against
// invite.
type UserGroupGetResponsePoliciesResourceGroupsScopeObject struct {
	// This is a combination of pre-defined resource name and identifier (like Zone ID
	// etc.)
	Key  string                                                    `json:"key,required"`
	JSON userGroupGetResponsePoliciesResourceGroupsScopeObjectJSON `json:"-"`
}

// userGroupGetResponsePoliciesResourceGroupsScopeObjectJSON contains the JSON
// metadata for the struct [UserGroupGetResponsePoliciesResourceGroupsScopeObject]
type userGroupGetResponsePoliciesResourceGroupsScopeObjectJSON struct {
	Key         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponsePoliciesResourceGroupsScopeObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponsePoliciesResourceGroupsScopeObjectJSON) RawJSON() string {
	return r.raw
}

// Attributes associated to the resource group.
type UserGroupGetResponsePoliciesResourceGroupsMeta struct {
	Key   string                                             `json:"key"`
	Value string                                             `json:"value"`
	JSON  userGroupGetResponsePoliciesResourceGroupsMetaJSON `json:"-"`
}

// userGroupGetResponsePoliciesResourceGroupsMetaJSON contains the JSON metadata
// for the struct [UserGroupGetResponsePoliciesResourceGroupsMeta]
type userGroupGetResponsePoliciesResourceGroupsMetaJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponsePoliciesResourceGroupsMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponsePoliciesResourceGroupsMetaJSON) RawJSON() string {
	return r.raw
}

type UserGroupNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the User group.
	Name param.Field[string] `json:"name,required"`
	// Policies attached to the User group
	Policies param.Field[[]UserGroupNewParamsPolicy] `json:"policies,required"`
}

func (r UserGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupNewParamsPolicy struct {
	// Allow or deny operations against the resources.
	Access param.Field[UserGroupNewParamsPoliciesAccess] `json:"access,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]UserGroupNewParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A set of resource groups that are specified to the policy.
	ResourceGroups param.Field[[]UserGroupNewParamsPoliciesResourceGroup] `json:"resource_groups,required"`
}

func (r UserGroupNewParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type UserGroupNewParamsPoliciesAccess string

const (
	UserGroupNewParamsPoliciesAccessAllow UserGroupNewParamsPoliciesAccess = "allow"
	UserGroupNewParamsPoliciesAccessDeny  UserGroupNewParamsPoliciesAccess = "deny"
)

func (r UserGroupNewParamsPoliciesAccess) IsKnown() bool {
	switch r {
	case UserGroupNewParamsPoliciesAccessAllow, UserGroupNewParamsPoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type UserGroupNewParamsPoliciesPermissionGroup struct {
	// Permission Group identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r UserGroupNewParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type UserGroupNewParamsPoliciesResourceGroup struct {
	// Resource Group identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r UserGroupNewParamsPoliciesResourceGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupNewResponseEnvelope struct {
	Errors   []UserGroupNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGroupNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserGroupNewResponseEnvelopeSuccess `json:"success,required"`
	// A group of policies resources.
	Result UserGroupNewResponse             `json:"result"`
	JSON   userGroupNewResponseEnvelopeJSON `json:"-"`
}

// userGroupNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGroupNewResponseEnvelope]
type userGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupNewResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           UserGroupNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserGroupNewResponseEnvelopeErrors]
type userGroupNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupNewResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    userGroupNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [UserGroupNewResponseEnvelopeErrorsSource]
type userGroupNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupNewResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           UserGroupNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserGroupNewResponseEnvelopeMessages]
type userGroupNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupNewResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    userGroupNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [UserGroupNewResponseEnvelopeMessagesSource]
type userGroupNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupNewResponseEnvelopeSuccess bool

const (
	UserGroupNewResponseEnvelopeSuccessTrue UserGroupNewResponseEnvelopeSuccess = true
)

func (r UserGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGroupUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Name of the User group.
	Name param.Field[string] `json:"name"`
	// Policies attached to the User group
	Policies param.Field[[]UserGroupUpdateParamsPolicy] `json:"policies"`
}

func (r UserGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupUpdateParamsPolicy struct {
	// Policy identifier.
	ID param.Field[string] `json:"id,required"`
	// Allow or deny operations against the resources.
	Access param.Field[UserGroupUpdateParamsPoliciesAccess] `json:"access,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]UserGroupUpdateParamsPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A set of resource groups that are specified to the policy.
	ResourceGroups param.Field[[]UserGroupUpdateParamsPoliciesResourceGroup] `json:"resource_groups,required"`
}

func (r UserGroupUpdateParamsPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type UserGroupUpdateParamsPoliciesAccess string

const (
	UserGroupUpdateParamsPoliciesAccessAllow UserGroupUpdateParamsPoliciesAccess = "allow"
	UserGroupUpdateParamsPoliciesAccessDeny  UserGroupUpdateParamsPoliciesAccess = "deny"
)

func (r UserGroupUpdateParamsPoliciesAccess) IsKnown() bool {
	switch r {
	case UserGroupUpdateParamsPoliciesAccessAllow, UserGroupUpdateParamsPoliciesAccessDeny:
		return true
	}
	return false
}

// A named group of permissions that map to a group of operations against
// resources.
type UserGroupUpdateParamsPoliciesPermissionGroup struct {
	// Permission Group identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r UserGroupUpdateParamsPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type UserGroupUpdateParamsPoliciesResourceGroup struct {
	// Resource Group identifier tag.
	ID param.Field[string] `json:"id,required"`
}

func (r UserGroupUpdateParamsPoliciesResourceGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupUpdateResponseEnvelope struct {
	Errors   []UserGroupUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGroupUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserGroupUpdateResponseEnvelopeSuccess `json:"success,required"`
	// A group of policies resources.
	Result UserGroupUpdateResponse             `json:"result"`
	JSON   userGroupUpdateResponseEnvelopeJSON `json:"-"`
}

// userGroupUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGroupUpdateResponseEnvelope]
type userGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupUpdateResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           UserGroupUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserGroupUpdateResponseEnvelopeErrors]
type userGroupUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    userGroupUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [UserGroupUpdateResponseEnvelopeErrorsSource]
type userGroupUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupUpdateResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           UserGroupUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserGroupUpdateResponseEnvelopeMessages]
type userGroupUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    userGroupUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [UserGroupUpdateResponseEnvelopeMessagesSource]
type userGroupUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupUpdateResponseEnvelopeSuccess bool

const (
	UserGroupUpdateResponseEnvelopeSuccessTrue UserGroupUpdateResponseEnvelopeSuccess = true
)

func (r UserGroupUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGroupListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// ID of the user group to be fetched.
	ID param.Field[string] `query:"id"`
	// The sort order of returned user groups by name. Default sort order is ascending.
	// To switch to descending, set this parameter to "desc"
	Direction param.Field[string] `query:"direction"`
	// A string used for searching for user groups containing that substring.
	FuzzyName param.Field[string] `query:"fuzzyName"`
	// Name of the user group to be fetched.
	Name param.Field[string] `query:"name"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserGroupListParams]'s query parameters as `url.Values`.
func (r UserGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserGroupDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type UserGroupDeleteResponseEnvelope struct {
	Errors   []UserGroupDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGroupDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserGroupDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  UserGroupDeleteResponse                `json:"result,nullable"`
	JSON    userGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// userGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGroupDeleteResponseEnvelope]
type userGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupDeleteResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           UserGroupDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserGroupDeleteResponseEnvelopeErrors]
type userGroupDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    userGroupDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [UserGroupDeleteResponseEnvelopeErrorsSource]
type userGroupDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupDeleteResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           UserGroupDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserGroupDeleteResponseEnvelopeMessages]
type userGroupDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    userGroupDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [UserGroupDeleteResponseEnvelopeMessagesSource]
type userGroupDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupDeleteResponseEnvelopeSuccess bool

const (
	UserGroupDeleteResponseEnvelopeSuccessTrue UserGroupDeleteResponseEnvelopeSuccess = true
)

func (r UserGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGroupGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type UserGroupGetResponseEnvelope struct {
	Errors   []UserGroupGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGroupGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserGroupGetResponseEnvelopeSuccess `json:"success,required"`
	// A group of policies resources.
	Result UserGroupGetResponse             `json:"result"`
	JSON   userGroupGetResponseEnvelopeJSON `json:"-"`
}

// userGroupGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGroupGetResponseEnvelope]
type userGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupGetResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           UserGroupGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserGroupGetResponseEnvelopeErrors]
type userGroupGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupGetResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    userGroupGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [UserGroupGetResponseEnvelopeErrorsSource]
type userGroupGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupGetResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           UserGroupGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [UserGroupGetResponseEnvelopeMessages]
type userGroupGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupGetResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    userGroupGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [UserGroupGetResponseEnvelopeMessagesSource]
type userGroupGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupGetResponseEnvelopeSuccess bool

const (
	UserGroupGetResponseEnvelopeSuccessTrue UserGroupGetResponseEnvelopeSuccess = true
)

func (r UserGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
