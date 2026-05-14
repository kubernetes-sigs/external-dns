// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package accounts

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v5/shared"
)

// MemberService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemberService] method instead.
type MemberService struct {
	Options []option.RequestOption
}

// NewMemberService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemberService(opts ...option.RequestOption) (r *MemberService) {
	r = &MemberService{}
	r.Options = opts
	return
}

// Add a user to the list of members for this account.
func (r *MemberService) New(ctx context.Context, params MemberNewParams, opts ...option.RequestOption) (res *shared.Member, err error) {
	var env MemberNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Modify an account member.
func (r *MemberService) Update(ctx context.Context, memberID string, params MemberUpdateParams, opts ...option.RequestOption) (res *shared.Member, err error) {
	var env MemberUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", params.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all members of an account.
func (r *MemberService) List(ctx context.Context, params MemberListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[shared.Member], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members", params.AccountID)
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

// List all members of an account.
func (r *MemberService) ListAutoPaging(ctx context.Context, params MemberListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[shared.Member] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Remove a member from an account.
func (r *MemberService) Delete(ctx context.Context, memberID string, body MemberDeleteParams, opts ...option.RequestOption) (res *MemberDeleteResponse, err error) {
	var env MemberDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", body.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get information about a specific member of an account.
func (r *MemberService) Get(ctx context.Context, memberID string, query MemberGetParams, opts ...option.RequestOption) (res *shared.Member, err error) {
	var env MemberGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/members/%s", query.AccountID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Whether the user is a member of the organization or has an invitation pending.
type Status string

const (
	StatusMember  Status = "member"
	StatusInvited Status = "invited"
)

func (r Status) IsKnown() bool {
	switch r {
	case StatusMember, StatusInvited:
		return true
	}
	return false
}

type MemberDeleteResponse struct {
	// Identifier
	ID   string                   `json:"id,required"`
	JSON memberDeleteResponseJSON `json:"-"`
}

// memberDeleteResponseJSON contains the JSON metadata for the struct
// [MemberDeleteResponse]
type memberDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type MemberNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string]      `path:"account_id,required"`
	Body      MemberNewParamsBodyUnion `json:"body,required"`
}

func (r MemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MemberNewParamsBody struct {
	// The contact email address of the user.
	Email    param.Field[string]                    `json:"email,required"`
	Policies param.Field[interface{}]               `json:"policies"`
	Roles    param.Field[interface{}]               `json:"roles"`
	Status   param.Field[MemberNewParamsBodyStatus] `json:"status"`
}

func (r MemberNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberNewParamsBody) implementsMemberNewParamsBodyUnion() {}

// Satisfied by [accounts.MemberNewParamsBodyIAMCreateMemberWithRoles],
// [accounts.MemberNewParamsBodyIAMCreateMemberWithPolicies],
// [MemberNewParamsBody].
type MemberNewParamsBodyUnion interface {
	implementsMemberNewParamsBodyUnion()
}

type MemberNewParamsBodyIAMCreateMemberWithRoles struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of roles associated with this member.
	Roles  param.Field[[]string]                                          `json:"roles,required"`
	Status param.Field[MemberNewParamsBodyIAMCreateMemberWithRolesStatus] `json:"status"`
}

func (r MemberNewParamsBodyIAMCreateMemberWithRoles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberNewParamsBodyIAMCreateMemberWithRoles) implementsMemberNewParamsBodyUnion() {}

type MemberNewParamsBodyIAMCreateMemberWithRolesStatus string

const (
	MemberNewParamsBodyIAMCreateMemberWithRolesStatusAccepted MemberNewParamsBodyIAMCreateMemberWithRolesStatus = "accepted"
	MemberNewParamsBodyIAMCreateMemberWithRolesStatusPending  MemberNewParamsBodyIAMCreateMemberWithRolesStatus = "pending"
)

func (r MemberNewParamsBodyIAMCreateMemberWithRolesStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyIAMCreateMemberWithRolesStatusAccepted, MemberNewParamsBodyIAMCreateMemberWithRolesStatusPending:
		return true
	}
	return false
}

type MemberNewParamsBodyIAMCreateMemberWithPolicies struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Array of policies associated with this member.
	Policies param.Field[[]MemberNewParamsBodyIAMCreateMemberWithPoliciesPolicy] `json:"policies,required"`
	Status   param.Field[MemberNewParamsBodyIAMCreateMemberWithPoliciesStatus]   `json:"status"`
}

func (r MemberNewParamsBodyIAMCreateMemberWithPolicies) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberNewParamsBodyIAMCreateMemberWithPolicies) implementsMemberNewParamsBodyUnion() {}

type MemberNewParamsBodyIAMCreateMemberWithPoliciesPolicy struct {
	// Allow or deny operations against the resources.
	Access param.Field[MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccess] `json:"access,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource groups that the policy applies to.
	ResourceGroups param.Field[[]MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesResourceGroup] `json:"resource_groups,required"`
}

func (r MemberNewParamsBodyIAMCreateMemberWithPoliciesPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccess string

const (
	MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccessAllow MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccess = "allow"
	MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccessDeny  MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccess = "deny"
)

func (r MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccess) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccessAllow, MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesAccessDeny:
		return true
	}
	return false
}

// A group of permissions.
type MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesPermissionGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesResourceGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberNewParamsBodyIAMCreateMemberWithPoliciesPoliciesResourceGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberNewParamsBodyIAMCreateMemberWithPoliciesStatus string

const (
	MemberNewParamsBodyIAMCreateMemberWithPoliciesStatusAccepted MemberNewParamsBodyIAMCreateMemberWithPoliciesStatus = "accepted"
	MemberNewParamsBodyIAMCreateMemberWithPoliciesStatusPending  MemberNewParamsBodyIAMCreateMemberWithPoliciesStatus = "pending"
)

func (r MemberNewParamsBodyIAMCreateMemberWithPoliciesStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyIAMCreateMemberWithPoliciesStatusAccepted, MemberNewParamsBodyIAMCreateMemberWithPoliciesStatusPending:
		return true
	}
	return false
}

type MemberNewParamsBodyStatus string

const (
	MemberNewParamsBodyStatusAccepted MemberNewParamsBodyStatus = "accepted"
	MemberNewParamsBodyStatusPending  MemberNewParamsBodyStatus = "pending"
)

func (r MemberNewParamsBodyStatus) IsKnown() bool {
	switch r {
	case MemberNewParamsBodyStatusAccepted, MemberNewParamsBodyStatusPending:
		return true
	}
	return false
}

type MemberNewResponseEnvelope struct {
	Errors   []MemberNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success MemberNewResponseEnvelopeSuccess `json:"success,required"`
	Result  shared.Member                    `json:"result"`
	JSON    memberNewResponseEnvelopeJSON    `json:"-"`
}

// memberNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelope]
type memberNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           MemberNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             memberNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// memberNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelopeErrors]
type memberNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    memberNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// memberNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [MemberNewResponseEnvelopeErrorsSource]
type memberNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           MemberNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             memberNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// memberNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MemberNewResponseEnvelopeMessages]
type memberNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MemberNewResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    memberNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// memberNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [MemberNewResponseEnvelopeMessagesSource]
type memberNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MemberNewResponseEnvelopeSuccess bool

const (
	MemberNewResponseEnvelopeSuccessTrue MemberNewResponseEnvelopeSuccess = true
)

func (r MemberNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string]         `path:"account_id,required"`
	Body      MemberUpdateParamsBodyUnion `json:"body,required"`
}

func (r MemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type MemberUpdateParamsBody struct {
	Policies param.Field[interface{}] `json:"policies"`
	Roles    param.Field[interface{}] `json:"roles"`
	User     param.Field[interface{}] `json:"user"`
}

func (r MemberUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberUpdateParamsBody) implementsMemberUpdateParamsBodyUnion() {}

// Satisfied by [accounts.MemberUpdateParamsBodyIAMUpdateMemberWithRoles],
// [accounts.MemberUpdateParamsBodyIAMUpdateMemberWithPolicies],
// [MemberUpdateParamsBody].
type MemberUpdateParamsBodyUnion interface {
	implementsMemberUpdateParamsBodyUnion()
}

type MemberUpdateParamsBodyIAMUpdateMemberWithRoles struct {
	// Roles assigned to this member.
	Roles param.Field[[]shared.RoleParam] `json:"roles"`
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithRoles) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithRoles) implementsMemberUpdateParamsBodyUnion() {}

// A member's status in the account.
type MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatus string

const (
	MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatusAccepted MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatus = "accepted"
	MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatusPending  MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatus = "pending"
)

func (r MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatus) IsKnown() bool {
	switch r {
	case MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatusAccepted, MemberUpdateParamsBodyIAMUpdateMemberWithRolesStatusPending:
		return true
	}
	return false
}

// Details of the user associated to the membership.
type MemberUpdateParamsBodyIAMUpdateMemberWithRolesUser struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
	// Identifier
	ID param.Field[string] `json:"id"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithRolesUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type MemberUpdateParamsBodyIAMUpdateMemberWithPolicies struct {
	// Array of policies associated with this member.
	Policies param.Field[[]MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPolicy] `json:"policies,required"`
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithPolicies) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithPolicies) implementsMemberUpdateParamsBodyUnion() {}

type MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPolicy struct {
	// Allow or deny operations against the resources.
	Access param.Field[MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccess] `json:"access,required"`
	// A set of permission groups that are specified to the policy.
	PermissionGroups param.Field[[]MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesPermissionGroup] `json:"permission_groups,required"`
	// A list of resource groups that the policy applies to.
	ResourceGroups param.Field[[]MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesResourceGroup] `json:"resource_groups,required"`
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPolicy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Allow or deny operations against the resources.
type MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccess string

const (
	MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccessAllow MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccess = "allow"
	MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccessDeny  MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccess = "deny"
)

func (r MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccess) IsKnown() bool {
	switch r {
	case MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccessAllow, MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesAccessDeny:
		return true
	}
	return false
}

// A group of permissions.
type MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesPermissionGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesPermissionGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A group of scoped resources.
type MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesResourceGroup struct {
	// Identifier of the group.
	ID param.Field[string] `json:"id,required"`
}

func (r MemberUpdateParamsBodyIAMUpdateMemberWithPoliciesPoliciesResourceGroup) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A member's status in the account.
type MemberUpdateParamsBodyStatus string

const (
	MemberUpdateParamsBodyStatusAccepted MemberUpdateParamsBodyStatus = "accepted"
	MemberUpdateParamsBodyStatusPending  MemberUpdateParamsBodyStatus = "pending"
)

func (r MemberUpdateParamsBodyStatus) IsKnown() bool {
	switch r {
	case MemberUpdateParamsBodyStatusAccepted, MemberUpdateParamsBodyStatusPending:
		return true
	}
	return false
}

type MemberUpdateResponseEnvelope struct {
	Errors   []MemberUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success MemberUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  shared.Member                       `json:"result"`
	JSON    memberUpdateResponseEnvelopeJSON    `json:"-"`
}

// memberUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberUpdateResponseEnvelope]
type memberUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           MemberUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             memberUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// memberUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberUpdateResponseEnvelopeErrors]
type memberUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    memberUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// memberUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [MemberUpdateResponseEnvelopeErrorsSource]
type memberUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           MemberUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             memberUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// memberUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MemberUpdateResponseEnvelopeMessages]
type memberUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MemberUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    memberUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// memberUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [MemberUpdateResponseEnvelopeMessagesSource]
type memberUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MemberUpdateResponseEnvelopeSuccess bool

const (
	MemberUpdateResponseEnvelopeSuccessTrue MemberUpdateResponseEnvelopeSuccess = true
)

func (r MemberUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Direction to order results.
	Direction param.Field[MemberListParamsDirection] `query:"direction"`
	// Field to order results by.
	Order param.Field[MemberListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// A member's status in the account.
	Status param.Field[MemberListParamsStatus] `query:"status"`
}

// URLQuery serializes [MemberListParams]'s query parameters as `url.Values`.
func (r MemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order results.
type MemberListParamsDirection string

const (
	MemberListParamsDirectionAsc  MemberListParamsDirection = "asc"
	MemberListParamsDirectionDesc MemberListParamsDirection = "desc"
)

func (r MemberListParamsDirection) IsKnown() bool {
	switch r {
	case MemberListParamsDirectionAsc, MemberListParamsDirectionDesc:
		return true
	}
	return false
}

// Field to order results by.
type MemberListParamsOrder string

const (
	MemberListParamsOrderUserFirstName MemberListParamsOrder = "user.first_name"
	MemberListParamsOrderUserLastName  MemberListParamsOrder = "user.last_name"
	MemberListParamsOrderUserEmail     MemberListParamsOrder = "user.email"
	MemberListParamsOrderStatus        MemberListParamsOrder = "status"
)

func (r MemberListParamsOrder) IsKnown() bool {
	switch r {
	case MemberListParamsOrderUserFirstName, MemberListParamsOrderUserLastName, MemberListParamsOrderUserEmail, MemberListParamsOrderStatus:
		return true
	}
	return false
}

// A member's status in the account.
type MemberListParamsStatus string

const (
	MemberListParamsStatusAccepted MemberListParamsStatus = "accepted"
	MemberListParamsStatusPending  MemberListParamsStatus = "pending"
	MemberListParamsStatusRejected MemberListParamsStatus = "rejected"
)

func (r MemberListParamsStatus) IsKnown() bool {
	switch r {
	case MemberListParamsStatusAccepted, MemberListParamsStatusPending, MemberListParamsStatusRejected:
		return true
	}
	return false
}

type MemberDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MemberDeleteResponseEnvelope struct {
	Errors   []MemberDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success MemberDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  MemberDeleteResponse                `json:"result,nullable"`
	JSON    memberDeleteResponseEnvelopeJSON    `json:"-"`
}

// memberDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelope]
type memberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeErrors struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           MemberDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             memberDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// memberDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberDeleteResponseEnvelopeErrors]
type memberDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    memberDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// memberDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [MemberDeleteResponseEnvelopeErrorsSource]
type memberDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeMessages struct {
	Code             int64                                      `json:"code,required"`
	Message          string                                     `json:"message,required"`
	DocumentationURL string                                     `json:"documentation_url"`
	Source           MemberDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             memberDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// memberDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [MemberDeleteResponseEnvelopeMessages]
type memberDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MemberDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                         `json:"pointer"`
	JSON    memberDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// memberDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [MemberDeleteResponseEnvelopeMessagesSource]
type memberDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MemberDeleteResponseEnvelopeSuccess bool

const (
	MemberDeleteResponseEnvelopeSuccessTrue MemberDeleteResponseEnvelopeSuccess = true
)

func (r MemberDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type MemberGetParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type MemberGetResponseEnvelope struct {
	Errors   []MemberGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []MemberGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success MemberGetResponseEnvelopeSuccess `json:"success,required"`
	Result  shared.Member                    `json:"result"`
	JSON    memberGetResponseEnvelopeJSON    `json:"-"`
}

// memberGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelope]
type memberGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           MemberGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             memberGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// memberGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelopeErrors]
type memberGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    memberGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// memberGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [MemberGetResponseEnvelopeErrorsSource]
type memberGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           MemberGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             memberGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// memberGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [MemberGetResponseEnvelopeMessages]
type memberGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type MemberGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    memberGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// memberGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [MemberGetResponseEnvelopeMessagesSource]
type memberGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MemberGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r memberGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type MemberGetResponseEnvelopeSuccess bool

const (
	MemberGetResponseEnvelopeSuccessTrue MemberGetResponseEnvelopeSuccess = true
)

func (r MemberGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case MemberGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
