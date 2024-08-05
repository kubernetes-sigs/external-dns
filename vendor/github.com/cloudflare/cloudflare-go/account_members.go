package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AccountMember is the definition of a member of an account.
type AccountMember struct {
	ID     string                   `json:"id"`
	Code   string                   `json:"code"`
	User   AccountMemberUserDetails `json:"user"`
	Status string                   `json:"status"`
	Roles  []AccountRole            `json:"roles"`
}

// AccountMemberUserDetails outlines all the personal information about
// a member.
type AccountMemberUserDetails struct {
	ID                             string `json:"id"`
	FirstName                      string `json:"first_name"`
	LastName                       string `json:"last_name"`
	Email                          string `json:"email"`
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	TwoFactorAuthenticationEnabled bool   `json:"two_factor_authentication_enabled"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	TwoFactorAuthenticationEnabled bool
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	TwoFactorAuthenticationEnabled bool
=======
	TwoFactorAuthenticationEnabled bool   `json:"two_factor_authentication_enabled"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	TwoFactorAuthenticationEnabled bool
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	TwoFactorAuthenticationEnabled bool
=======
	TwoFactorAuthenticationEnabled bool   `json:"two_factor_authentication_enabled"`
>>>>>>> 6b7ce455e (update vendored files)
}

// AccountMembersListResponse represents the response from the list
// account members endpoint.
type AccountMembersListResponse struct {
	Result []AccountMember `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccountMemberDetailResponse is the API response, containing a single
// account member.
type AccountMemberDetailResponse struct {
	Success  bool          `json:"success"`
	Errors   []string      `json:"errors"`
	Messages []string      `json:"messages"`
	Result   AccountMember `json:"result"`
}

// AccountMemberInvitation represents the invitation for a new member to
// the account.
type AccountMemberInvitation struct {
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	Status string   `json:"status,omitempty"`
}

// AccountMembers returns all members of an account.
//
// API reference: https://api.cloudflare.com/#accounts-list-accounts
func (api *API) AccountMembers(ctx context.Context, accountID string, pageOpts PaginationOptions) ([]AccountMember, ResultInfo, error) {
	if accountID == "" {
		return []AccountMember{}, ResultInfo{}, errors.New(errMissingAccountID)
	}

	v := url.Values{}
	if pageOpts.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(pageOpts.PerPage))
	}
	if pageOpts.Page > 0 {
		v.Set("page", strconv.Itoa(pageOpts.Page))
	}

	uri := fmt.Sprintf("/accounts/%s/members", accountID)
	if len(v) > 0 {
		uri = fmt.Sprintf("%s?%s", uri, v.Encode())
	}

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, err
	}

	var accountMemberListresponse AccountMembersListResponse
	err = json.Unmarshal(res, &accountMemberListresponse)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}

	return accountMemberListresponse.Result, accountMemberListresponse.ResultInfo, nil
}

// CreateAccountMemberWithStatus invites a new member to join an account, allowing setting the status.
//
// Refer to the API reference for valid statuses.
//
// API reference: https://api.cloudflare.com/#account-members-add-member
func (api *API) CreateAccountMemberWithStatus(ctx context.Context, accountID string, emailAddress string, roles []string, status string) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, errors.New(errMissingAccountID)
	}

	uri := fmt.Sprintf("/accounts/%s/members", accountID)

	var newMember = AccountMemberInvitation{
		Email:  emailAddress,
		Roles:  roles,
		Status: status,
	}
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, newMember)
	if err != nil {
		return AccountMember{}, err
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, errors.Wrap(err, errUnmarshalError)
	}

	return accountMemberListResponse.Result, nil
}

// CreateAccountMember invites a new member to join an account.
// The member will be placed into "pending" status and receive an email confirmation.
//
// API reference: https://api.cloudflare.com/#account-members-add-member
func (api *API) CreateAccountMember(ctx context.Context, accountID string, emailAddress string, roles []string) (AccountMember, error) {
	return api.CreateAccountMemberWithStatus(ctx, accountID, emailAddress, roles, "")
}

// DeleteAccountMember removes a member from an account.
//
// API reference: https://api.cloudflare.com/#account-members-remove-member
func (api *API) DeleteAccountMember(ctx context.Context, accountID string, userID string) error {
	if accountID == "" {
		return errors.New(errMissingAccountID)
	}

	uri := fmt.Sprintf("/accounts/%s/members/%s", accountID, userID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAccountMember modifies an existing account member.
//
// API reference: https://api.cloudflare.com/#account-members-update-member
func (api *API) UpdateAccountMember(ctx context.Context, accountID string, userID string, member AccountMember) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, errors.New(errMissingAccountID)
	}

	uri := fmt.Sprintf("/accounts/%s/members/%s", accountID, userID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, member)
	if err != nil {
		return AccountMember{}, err
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, errors.Wrap(err, errUnmarshalError)
	}

	return accountMemberListResponse.Result, nil
}

// AccountMember returns details of a single account member.
//
// API reference: https://api.cloudflare.com/#account-members-member-details
func (api *API) AccountMember(ctx context.Context, accountID string, memberID string) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, errors.New(errMissingAccountID)
	}

	uri := fmt.Sprintf(
		"/accounts/%s/members/%s",
		accountID,
		memberID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccountMember{}, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// AccountMember is the definition of a member of an account.
type AccountMember struct {
	ID     string                   `json:"id"`
	Code   string                   `json:"code"`
	User   AccountMemberUserDetails `json:"user"`
	Status string                   `json:"status"`
	Roles  []AccountRole            `json:"roles"`
}

// AccountMemberUserDetails outlines all the personal information about
// a member.
type AccountMemberUserDetails struct {
	ID                             string `json:"id"`
	FirstName                      string `json:"first_name"`
	LastName                       string `json:"last_name"`
	Email                          string `json:"email"`
	TwoFactorAuthenticationEnabled bool   `json:"two_factor_authentication_enabled"`
}

// AccountMembersListResponse represents the response from the list
// account members endpoint.
type AccountMembersListResponse struct {
	Result []AccountMember `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccountMemberDetailResponse is the API response, containing a single
// account member.
type AccountMemberDetailResponse struct {
	Success  bool          `json:"success"`
	Errors   []string      `json:"errors"`
	Messages []string      `json:"messages"`
	Result   AccountMember `json:"result"`
}

// AccountMemberInvitation represents the invitation for a new member to
// the account.
type AccountMemberInvitation struct {
	Email  string   `json:"email"`
	Roles  []string `json:"roles"`
	Status string   `json:"status,omitempty"`
}

// AccountMembers returns all members of an account.
//
// API reference: https://api.cloudflare.com/#accounts-list-accounts
func (api *API) AccountMembers(ctx context.Context, accountID string, pageOpts PaginationOptions) ([]AccountMember, ResultInfo, error) {
	if accountID == "" {
		return []AccountMember{}, ResultInfo{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/members", accountID), pageOpts)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, err
	}

	var accountMemberListresponse AccountMembersListResponse
	err = json.Unmarshal(res, &accountMemberListresponse)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accountMemberListresponse.Result, accountMemberListresponse.ResultInfo, nil
}

// CreateAccountMemberWithStatus invites a new member to join an account, allowing setting the status.
//
// Refer to the API reference for valid statuses.
//
// API reference: https://api.cloudflare.com/#account-members-add-member
func (api *API) CreateAccountMemberWithStatus(ctx context.Context, accountID string, emailAddress string, roles []string, status string) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/members", accountID)

	newMember := AccountMemberInvitation{
		Email:  emailAddress,
		Roles:  roles,
		Status: status,
	}
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, newMember)
	if err != nil {
		return AccountMember{}, err
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accountMemberListResponse.Result, nil
}

// CreateAccountMember invites a new member to join an account.
// The member will be placed into "pending" status and receive an email confirmation.
//
// API reference: https://api.cloudflare.com/#account-members-add-member
func (api *API) CreateAccountMember(ctx context.Context, accountID string, emailAddress string, roles []string) (AccountMember, error) {
	return api.CreateAccountMemberWithStatus(ctx, accountID, emailAddress, roles, "")
}

// DeleteAccountMember removes a member from an account.
//
// API reference: https://api.cloudflare.com/#account-members-remove-member
func (api *API) DeleteAccountMember(ctx context.Context, accountID string, userID string) error {
	if accountID == "" {
		return ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/members/%s", accountID, userID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAccountMember modifies an existing account member.
//
// API reference: https://api.cloudflare.com/#account-members-update-member
func (api *API) UpdateAccountMember(ctx context.Context, accountID string, userID string, member AccountMember) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/members/%s", accountID, userID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, member)
	if err != nil {
		return AccountMember{}, err
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accountMemberListResponse.Result, nil
}

// AccountMember returns details of a single account member.
//
// API reference: https://api.cloudflare.com/#account-members-member-details
func (api *API) AccountMember(ctx context.Context, accountID string, memberID string) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf(
		"/accounts/%s/members/%s",
		accountID,
		memberID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return AccountMember{}, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return AccountMember{}, errors.Wrap(err, errMakeRequestError)
=======
		return AccountMember{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
	}

	var accountMemberResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberResponse)
	if err != nil {
		return AccountMember{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"encoding/json"
=======
	"context"
	"errors"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// AccountMember is the definition of a member of an account.
type AccountMember struct {
	ID       string                   `json:"id"`
	Code     string                   `json:"code"`
	User     AccountMemberUserDetails `json:"user"`
	Status   string                   `json:"status"`
	Roles    []AccountRole            `json:"roles,omitempty"`
	Policies []Policy                 `json:"policies,omitempty"`
}

// AccountMemberUserDetails outlines all the personal information about
// a member.
type AccountMemberUserDetails struct {
	ID                             string `json:"id"`
	FirstName                      string `json:"first_name"`
	LastName                       string `json:"last_name"`
	Email                          string `json:"email"`
	TwoFactorAuthenticationEnabled bool   `json:"two_factor_authentication_enabled"`
}

// AccountMembersListResponse represents the response from the list
// account members endpoint.
type AccountMembersListResponse struct {
	Result []AccountMember `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccountMemberDetailResponse is the API response, containing a single
// account member.
type AccountMemberDetailResponse struct {
	Success  bool          `json:"success"`
	Errors   []string      `json:"errors"`
	Messages []string      `json:"messages"`
	Result   AccountMember `json:"result"`
}

// AccountMemberInvitation represents the invitation for a new member to
// the account.
type AccountMemberInvitation struct {
	Email    string   `json:"email"`
	Roles    []string `json:"roles,omitempty"`
	Policies []Policy `json:"policies,omitempty"`
	Status   string   `json:"status,omitempty"`
}

const errMissingMemberRolesOrPolicies = "account member must be created with roles or policies (not both)"

var ErrMissingMemberRolesOrPolicies = errors.New(errMissingMemberRolesOrPolicies)

type CreateAccountMemberParams struct {
	EmailAddress string
	Roles        []string
	Policies     []Policy
	Status       string
}

// AccountMembers returns all members of an account.
//
// API reference: https://api.cloudflare.com/#accounts-list-accounts
func (api *API) AccountMembers(ctx context.Context, accountID string, pageOpts PaginationOptions) ([]AccountMember, ResultInfo, error) {
	if accountID == "" {
		return []AccountMember{}, ResultInfo{}, ErrMissingAccountID
	}

	uri := buildURI(fmt.Sprintf("/accounts/%s/members", accountID), pageOpts)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, err
	}

	var accountMemberListresponse AccountMembersListResponse
	err = json.Unmarshal(res, &accountMemberListresponse)
	if err != nil {
		return []AccountMember{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accountMemberListresponse.Result, accountMemberListresponse.ResultInfo, nil
}

// CreateAccountMemberWithStatus invites a new member to join an account, allowing setting the status.
//
// Refer to the API reference for valid statuses.
//
// Deprecated: Use `CreateAccountMember` with a `Status` field instead.
//
// API reference: https://api.cloudflare.com/#account-members-add-member
func (api *API) CreateAccountMemberWithStatus(ctx context.Context, accountID string, emailAddress string, roles []string, status string) (AccountMember, error) {
	return api.CreateAccountMember(ctx, AccountIdentifier(accountID), CreateAccountMemberParams{
		EmailAddress: emailAddress,
		Roles:        roles,
		Status:       status,
	})
}

// CreateAccountMember invites a new member to join an account with roles.
// The member will be placed into "pending" status and receive an email confirmation.
// NOTE: If you are currently enrolled in Domain Scoped Roles, your roles will
// be converted to policies upon member invitation.
//
// API reference: https://api.cloudflare.com/#account-members-add-member
func (api *API) CreateAccountMember(ctx context.Context, rc *ResourceContainer, params CreateAccountMemberParams) (AccountMember, error) {
	if rc.Level != AccountRouteLevel {
		return AccountMember{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if rc.Identifier == "" {
		return AccountMember{}, ErrMissingAccountID
	}

	invite := AccountMemberInvitation{
		Email:  params.EmailAddress,
		Status: params.Status,
	}

	roles := []AccountRole{}
	for i := 0; i < len(params.Roles); i++ {
		roles = append(roles, AccountRole{ID: params.Roles[i]})
	}
	err := validateRolesAndPolicies(roles, params.Policies)
	if err != nil {
		return AccountMember{}, err
	}

	if params.Roles != nil {
		invite.Roles = params.Roles
	} else if params.Policies != nil {
		invite.Policies = params.Policies
	}

	uri := fmt.Sprintf("/accounts/%s/members", rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, invite)
	if err != nil {
		return AccountMember{}, err
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accountMemberListResponse.Result, nil
}

// DeleteAccountMember removes a member from an account.
//
// API reference: https://api.cloudflare.com/#account-members-remove-member
func (api *API) DeleteAccountMember(ctx context.Context, accountID string, userID string) error {
	if accountID == "" {
		return ErrMissingAccountID
	}

	uri := fmt.Sprintf("/accounts/%s/members/%s", accountID, userID)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	return nil
}

// UpdateAccountMember modifies an existing account member.
//
// API reference: https://api.cloudflare.com/#account-members-update-member
func (api *API) UpdateAccountMember(ctx context.Context, accountID string, userID string, member AccountMember) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, ErrMissingAccountID
	}

	err := validateRolesAndPolicies(member.Roles, member.Policies)
	if err != nil {
		return AccountMember{}, err
	}

	uri := fmt.Sprintf("/accounts/%s/members/%s", accountID, userID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, member)
	if err != nil {
		return AccountMember{}, err
	}

	var accountMemberListResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberListResponse)
	if err != nil {
		return AccountMember{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accountMemberListResponse.Result, nil
}

// AccountMember returns details of a single account member.
//
// API reference: https://api.cloudflare.com/#account-members-member-details
func (api *API) AccountMember(ctx context.Context, accountID string, memberID string) (AccountMember, error) {
	if accountID == "" {
		return AccountMember{}, ErrMissingAccountID
	}

	uri := fmt.Sprintf(
		"/accounts/%s/members/%s",
		accountID,
		memberID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccountMember{}, err
	}

	var accountMemberResponse AccountMemberDetailResponse
	err = json.Unmarshal(res, &accountMemberResponse)
	if err != nil {
<<<<<<< HEAD
		return AccountMember{}, errors.Wrap(err, errUnmarshalError)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		return AccountMember{}, errors.Wrap(err, errUnmarshalError)
=======
		return AccountMember{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}

	return accountMemberResponse.Result, nil
}

// validateRolesAndPolicies ensures either roles or policies are provided in
// CreateAccountMember requests, but not both.
func validateRolesAndPolicies(roles []AccountRole, policies []Policy) error {
	hasRoles := len(roles) > 0
	hasPolicies := len(policies) > 0
	hasRolesOrPolicies := hasRoles || hasPolicies
	hasRolesAndPolicies := hasRoles && hasPolicies
	hasCorrectPermissions := hasRolesOrPolicies && !hasRolesAndPolicies
	if !hasCorrectPermissions {
		return ErrMissingMemberRolesOrPolicies
	}
	return nil
}
