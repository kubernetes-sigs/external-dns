// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// AccessApplicationPolicyTestService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationPolicyTestService] method instead.
type AccessApplicationPolicyTestService struct {
	Options []option.RequestOption
	Users   *AccessApplicationPolicyTestUserService
}

// NewAccessApplicationPolicyTestService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessApplicationPolicyTestService(opts ...option.RequestOption) (r *AccessApplicationPolicyTestService) {
	r = &AccessApplicationPolicyTestService{}
	r.Options = opts
	r.Users = NewAccessApplicationPolicyTestUserService(opts...)
	return
}

// Starts an Access policy test.
func (r *AccessApplicationPolicyTestService) New(ctx context.Context, params AccessApplicationPolicyTestNewParams, opts ...option.RequestOption) (res *AccessApplicationPolicyTestNewResponse, err error) {
	var env AccessApplicationPolicyTestNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the current status of a given Access policy test.
func (r *AccessApplicationPolicyTestService) Get(ctx context.Context, policyTestID string, query AccessApplicationPolicyTestGetParams, opts ...option.RequestOption) (res *AccessApplicationPolicyTestGetResponse, err error) {
	var env AccessApplicationPolicyTestGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if policyTestID == "" {
		err = errors.New("missing required policy_test_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/policy-tests/%s", query.AccountID, policyTestID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessApplicationPolicyTestNewResponse struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The status of the policy test request.
	Status AccessApplicationPolicyTestNewResponseStatus `json:"status"`
	JSON   accessApplicationPolicyTestNewResponseJSON   `json:"-"`
}

// accessApplicationPolicyTestNewResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyTestNewResponse]
type accessApplicationPolicyTestNewResponseJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test request.
type AccessApplicationPolicyTestNewResponseStatus string

const (
	AccessApplicationPolicyTestNewResponseStatusSuccess AccessApplicationPolicyTestNewResponseStatus = "success"
)

func (r AccessApplicationPolicyTestNewResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestNewResponseStatusSuccess:
		return true
	}
	return false
}

type AccessApplicationPolicyTestGetResponse struct {
	// The UUID of the policy test.
	ID string `json:"id"`
	// The percentage of (processed) users approved based on policy evaluation results.
	PercentApproved int64 `json:"percent_approved"`
	// The percentage of (processed) users blocked based on policy evaluation results.
	PercentBlocked int64 `json:"percent_blocked"`
	// The percentage of (processed) users errored based on policy evaluation results.
	PercentErrored int64 `json:"percent_errored"`
	// The percentage of users processed so far (of the entire user base).
	PercentUsersProcessed int64 `json:"percent_users_processed"`
	// The status of the policy test.
	Status AccessApplicationPolicyTestGetResponseStatus `json:"status"`
	// The total number of users in the user base.
	TotalUsers int64 `json:"total_users"`
	// The number of (processed) users approved based on policy evaluation results.
	UsersApproved int64 `json:"users_approved"`
	// The number of (processed) users blocked based on policy evaluation results.
	UsersBlocked int64 `json:"users_blocked"`
	// The number of (processed) users errored based on policy evaluation results.
	UsersErrored int64                                      `json:"users_errored"`
	JSON         accessApplicationPolicyTestGetResponseJSON `json:"-"`
}

// accessApplicationPolicyTestGetResponseJSON contains the JSON metadata for the
// struct [AccessApplicationPolicyTestGetResponse]
type accessApplicationPolicyTestGetResponseJSON struct {
	ID                    apijson.Field
	PercentApproved       apijson.Field
	PercentBlocked        apijson.Field
	PercentErrored        apijson.Field
	PercentUsersProcessed apijson.Field
	Status                apijson.Field
	TotalUsers            apijson.Field
	UsersApproved         apijson.Field
	UsersBlocked          apijson.Field
	UsersErrored          apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseJSON) RawJSON() string {
	return r.raw
}

// The status of the policy test.
type AccessApplicationPolicyTestGetResponseStatus string

const (
	AccessApplicationPolicyTestGetResponseStatusBlocked      AccessApplicationPolicyTestGetResponseStatus = "blocked"
	AccessApplicationPolicyTestGetResponseStatusProcessing   AccessApplicationPolicyTestGetResponseStatus = "processing"
	AccessApplicationPolicyTestGetResponseStatusExceededTime AccessApplicationPolicyTestGetResponseStatus = "exceeded time"
	AccessApplicationPolicyTestGetResponseStatusComplete     AccessApplicationPolicyTestGetResponseStatus = "complete"
)

func (r AccessApplicationPolicyTestGetResponseStatus) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestGetResponseStatusBlocked, AccessApplicationPolicyTestGetResponseStatusProcessing, AccessApplicationPolicyTestGetResponseStatusExceededTime, AccessApplicationPolicyTestGetResponseStatusComplete:
		return true
	}
	return false
}

type AccessApplicationPolicyTestNewParams struct {
	// Identifier.
	AccountID param.Field[string]                                            `path:"account_id,required"`
	Policies  param.Field[[]AccessApplicationPolicyTestNewParamsPolicyUnion] `json:"policies"`
}

func (r AccessApplicationPolicyTestNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The UUID of the reusable policy you wish to test
//
// Satisfied by [zero_trust.AccessApplicationPolicyTestNewParamsPoliciesObject],
// [shared.UnionString].
type AccessApplicationPolicyTestNewParamsPolicyUnion interface {
	ImplementsAccessApplicationPolicyTestNewParamsPolicyUnion()
}

type AccessApplicationPolicyTestNewParamsPoliciesObject struct {
	// The action Access will take if a user matches this policy. Infrastructure
	// application policies can only use the Allow action.
	Decision param.Field[Decision] `json:"decision,required"`
	// Rules evaluated with an OR logical operator. A user needs to meet only one of
	// the Include rules.
	Include param.Field[[]AccessRuleUnionParam] `json:"include,required"`
	// The name of the Access policy.
	Name param.Field[string] `json:"name,required"`
	// Administrators who can approve a temporary authentication request.
	ApprovalGroups param.Field[[]ApprovalGroupParam] `json:"approval_groups"`
	// Requires the user to request access from an administrator at the start of each
	// session.
	ApprovalRequired param.Field[bool] `json:"approval_required"`
	// Rules evaluated with a NOT logical operator. To match the policy, a user cannot
	// meet any of the Exclude rules.
	Exclude param.Field[[]AccessRuleUnionParam] `json:"exclude"`
	// Require this application to be served in an isolated browser for users matching
	// this policy. 'Client Web Isolation' must be on for the account in order to use
	// this feature.
	IsolationRequired param.Field[bool] `json:"isolation_required"`
	// A custom message that will appear on the purpose justification screen.
	PurposeJustificationPrompt param.Field[string] `json:"purpose_justification_prompt"`
	// Require users to enter a justification when they log in to the application.
	PurposeJustificationRequired param.Field[bool] `json:"purpose_justification_required"`
	// Rules evaluated with an AND logical operator. To match the policy, a user must
	// meet all of the Require rules.
	Require param.Field[[]AccessRuleUnionParam] `json:"require"`
	// The amount of time that tokens issued for the application will be valid. Must be
	// in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s,
	// m, h.
	SessionDuration param.Field[string] `json:"session_duration"`
}

func (r AccessApplicationPolicyTestNewParamsPoliciesObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccessApplicationPolicyTestNewParamsPoliciesObject) ImplementsAccessApplicationPolicyTestNewParamsPolicyUnion() {
}

type AccessApplicationPolicyTestNewResponseEnvelope struct {
	Errors   []AccessApplicationPolicyTestNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyTestNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationPolicyTestNewResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationPolicyTestNewResponse                `json:"result"`
	JSON    accessApplicationPolicyTestNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyTestNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyTestNewResponseEnvelope]
type accessApplicationPolicyTestNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestNewResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessApplicationPolicyTestNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationPolicyTestNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationPolicyTestNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyTestNewResponseEnvelopeErrors]
type accessApplicationPolicyTestNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessApplicationPolicyTestNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationPolicyTestNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyTestNewResponseEnvelopeErrorsSource]
type accessApplicationPolicyTestNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestNewResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AccessApplicationPolicyTestNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationPolicyTestNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationPolicyTestNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyTestNewResponseEnvelopeMessages]
type accessApplicationPolicyTestNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    accessApplicationPolicyTestNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationPolicyTestNewResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyTestNewResponseEnvelopeMessagesSource]
type accessApplicationPolicyTestNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationPolicyTestNewResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyTestNewResponseEnvelopeSuccessTrue AccessApplicationPolicyTestNewResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyTestNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationPolicyTestGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessApplicationPolicyTestGetResponseEnvelope struct {
	Errors   []AccessApplicationPolicyTestGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationPolicyTestGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationPolicyTestGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationPolicyTestGetResponse                `json:"result"`
	JSON    accessApplicationPolicyTestGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationPolicyTestGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AccessApplicationPolicyTestGetResponseEnvelope]
type accessApplicationPolicyTestGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestGetResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code,required"`
	Message          string                                                     `json:"message,required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           AccessApplicationPolicyTestGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationPolicyTestGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationPolicyTestGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyTestGetResponseEnvelopeErrors]
type accessApplicationPolicyTestGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    accessApplicationPolicyTestGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationPolicyTestGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessApplicationPolicyTestGetResponseEnvelopeErrorsSource]
type accessApplicationPolicyTestGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestGetResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code,required"`
	Message          string                                                       `json:"message,required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           AccessApplicationPolicyTestGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationPolicyTestGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationPolicyTestGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessApplicationPolicyTestGetResponseEnvelopeMessages]
type accessApplicationPolicyTestGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationPolicyTestGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    accessApplicationPolicyTestGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationPolicyTestGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessApplicationPolicyTestGetResponseEnvelopeMessagesSource]
type accessApplicationPolicyTestGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationPolicyTestGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationPolicyTestGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationPolicyTestGetResponseEnvelopeSuccess bool

const (
	AccessApplicationPolicyTestGetResponseEnvelopeSuccessTrue AccessApplicationPolicyTestGetResponseEnvelopeSuccess = true
)

func (r AccessApplicationPolicyTestGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationPolicyTestGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
