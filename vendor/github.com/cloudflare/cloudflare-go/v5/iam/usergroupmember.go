// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package iam

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
)

// UserGroupMemberService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserGroupMemberService] method instead.
type UserGroupMemberService struct {
	Options []option.RequestOption
}

// NewUserGroupMemberService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserGroupMemberService(opts ...option.RequestOption) (r *UserGroupMemberService) {
	r = &UserGroupMemberService{}
	r.Options = opts
	return
}

// Add members to a User Group.
func (r *UserGroupMemberService) New(ctx context.Context, userGroupID string, params UserGroupMemberNewParams, opts ...option.RequestOption) (res *UserGroupMemberNewResponse, err error) {
	var env UserGroupMemberNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members", params.AccountID, userGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Replace the set of members attached to a User Group.
func (r *UserGroupMemberService) Update(ctx context.Context, userGroupID string, params UserGroupMemberUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[UserGroupMemberUpdateResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members", params.AccountID, userGroupID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Replace the set of members attached to a User Group.
func (r *UserGroupMemberService) UpdateAutoPaging(ctx context.Context, userGroupID string, params UserGroupMemberUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[UserGroupMemberUpdateResponse] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, userGroupID, params, opts...))
}

// List all the members attached to a user group.
func (r *UserGroupMemberService) List(ctx context.Context, userGroupID string, params UserGroupMemberListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[UserGroupMemberListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members", params.AccountID, userGroupID)
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

// List all the members attached to a user group.
func (r *UserGroupMemberService) ListAutoPaging(ctx context.Context, userGroupID string, params UserGroupMemberListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[UserGroupMemberListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, userGroupID, params, opts...))
}

// Remove a member from User Group
func (r *UserGroupMemberService) Delete(ctx context.Context, userGroupID string, memberID string, body UserGroupMemberDeleteParams, opts ...option.RequestOption) (res *UserGroupMemberDeleteResponse, err error) {
	var env UserGroupMemberDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userGroupID == "" {
		err = errors.New("missing required user_group_id parameter")
		return
	}
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/iam/user_groups/%s/members/%s", body.AccountID, userGroupID, memberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Member attached to a User Group.
type UserGroupMemberNewResponse struct {
	// Account member identifier.
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberNewResponseStatus `json:"status"`
	JSON   userGroupMemberNewResponseJSON   `json:"-"`
}

// userGroupMemberNewResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberNewResponse]
type userGroupMemberNewResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberNewResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberNewResponseStatus string

const (
	UserGroupMemberNewResponseStatusAccepted UserGroupMemberNewResponseStatus = "accepted"
	UserGroupMemberNewResponseStatusPending  UserGroupMemberNewResponseStatus = "pending"
)

func (r UserGroupMemberNewResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberNewResponseStatusAccepted, UserGroupMemberNewResponseStatusPending:
		return true
	}
	return false
}

// Member attached to a User Group.
type UserGroupMemberUpdateResponse struct {
	// Account member identifier.
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberUpdateResponseStatus `json:"status"`
	JSON   userGroupMemberUpdateResponseJSON   `json:"-"`
}

// userGroupMemberUpdateResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberUpdateResponse]
type userGroupMemberUpdateResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberUpdateResponseStatus string

const (
	UserGroupMemberUpdateResponseStatusAccepted UserGroupMemberUpdateResponseStatus = "accepted"
	UserGroupMemberUpdateResponseStatusPending  UserGroupMemberUpdateResponseStatus = "pending"
)

func (r UserGroupMemberUpdateResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberUpdateResponseStatusAccepted, UserGroupMemberUpdateResponseStatusPending:
		return true
	}
	return false
}

// Member attached to a User Group.
type UserGroupMemberListResponse struct {
	// Account member identifier.
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberListResponseStatus `json:"status"`
	JSON   userGroupMemberListResponseJSON   `json:"-"`
}

// userGroupMemberListResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberListResponse]
type userGroupMemberListResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberListResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberListResponseStatus string

const (
	UserGroupMemberListResponseStatusAccepted UserGroupMemberListResponseStatus = "accepted"
	UserGroupMemberListResponseStatusPending  UserGroupMemberListResponseStatus = "pending"
)

func (r UserGroupMemberListResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberListResponseStatusAccepted, UserGroupMemberListResponseStatusPending:
		return true
	}
	return false
}

// Member attached to a User Group.
type UserGroupMemberDeleteResponse struct {
	// Account member identifier.
	ID string `json:"id,required"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The member's status in the account.
	Status UserGroupMemberDeleteResponseStatus `json:"status"`
	JSON   userGroupMemberDeleteResponseJSON   `json:"-"`
}

// userGroupMemberDeleteResponseJSON contains the JSON metadata for the struct
// [UserGroupMemberDeleteResponse]
type userGroupMemberDeleteResponseJSON struct {
	ID          apijson.Field
	Email       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The member's status in the account.
type UserGroupMemberDeleteResponseStatus string

const (
	UserGroupMemberDeleteResponseStatusAccepted UserGroupMemberDeleteResponseStatus = "accepted"
	UserGroupMemberDeleteResponseStatusPending  UserGroupMemberDeleteResponseStatus = "pending"
)

func (r UserGroupMemberDeleteResponseStatus) IsKnown() bool {
	switch r {
	case UserGroupMemberDeleteResponseStatusAccepted, UserGroupMemberDeleteResponseStatusPending:
		return true
	}
	return false
}

type UserGroupMemberNewParams struct {
	// Account identifier tag.
	AccountID param.Field[string]            `path:"account_id,required"`
	Body      []UserGroupMemberNewParamsBody `json:"body,required"`
}

func (r UserGroupMemberNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserGroupMemberNewParamsBody struct {
	// The identifier of an existing account Member.
	ID param.Field[string] `json:"id,required"`
}

func (r UserGroupMemberNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupMemberNewResponseEnvelope struct {
	Errors   []UserGroupMemberNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGroupMemberNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserGroupMemberNewResponseEnvelopeSuccess `json:"success,required"`
	// Member attached to a User Group.
	Result UserGroupMemberNewResponse             `json:"result"`
	JSON   userGroupMemberNewResponseEnvelopeJSON `json:"-"`
}

// userGroupMemberNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGroupMemberNewResponseEnvelope]
type userGroupMemberNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           UserGroupMemberNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupMemberNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupMemberNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [UserGroupMemberNewResponseEnvelopeErrors]
type userGroupMemberNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    userGroupMemberNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupMemberNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [UserGroupMemberNewResponseEnvelopeErrorsSource]
type userGroupMemberNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           UserGroupMemberNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupMemberNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupMemberNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserGroupMemberNewResponseEnvelopeMessages]
type userGroupMemberNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    userGroupMemberNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupMemberNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [UserGroupMemberNewResponseEnvelopeMessagesSource]
type userGroupMemberNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupMemberNewResponseEnvelopeSuccess bool

const (
	UserGroupMemberNewResponseEnvelopeSuccessTrue UserGroupMemberNewResponseEnvelopeSuccess = true
)

func (r UserGroupMemberNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupMemberNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGroupMemberUpdateParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Set/Replace members to a user group.
	Body []UserGroupMemberUpdateParamsBody `json:"body,required"`
}

func (r UserGroupMemberUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type UserGroupMemberUpdateParamsBody struct {
	// The identifier of an existing account Member.
	ID param.Field[string] `json:"id,required"`
}

func (r UserGroupMemberUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserGroupMemberListParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [UserGroupMemberListParams]'s query parameters as
// `url.Values`.
func (r UserGroupMemberListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type UserGroupMemberDeleteParams struct {
	// Account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type UserGroupMemberDeleteResponseEnvelope struct {
	Errors   []UserGroupMemberDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGroupMemberDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserGroupMemberDeleteResponseEnvelopeSuccess `json:"success,required"`
	// Member attached to a User Group.
	Result UserGroupMemberDeleteResponse             `json:"result"`
	JSON   userGroupMemberDeleteResponseEnvelopeJSON `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [UserGroupMemberDeleteResponseEnvelope]
type userGroupMemberDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           UserGroupMemberDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGroupMemberDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [UserGroupMemberDeleteResponseEnvelopeErrors]
type userGroupMemberDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [UserGroupMemberDeleteResponseEnvelopeErrorsSource]
type userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           UserGroupMemberDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGroupMemberDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [UserGroupMemberDeleteResponseEnvelopeMessages]
type userGroupMemberDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGroupMemberDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [UserGroupMemberDeleteResponseEnvelopeMessagesSource]
type userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGroupMemberDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGroupMemberDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGroupMemberDeleteResponseEnvelopeSuccess bool

const (
	UserGroupMemberDeleteResponseEnvelopeSuccessTrue UserGroupMemberDeleteResponseEnvelopeSuccess = true
)

func (r UserGroupMemberDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGroupMemberDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
