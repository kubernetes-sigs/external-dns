// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
)

// UserService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options       []option.RequestOption
	AuditLogs     *AuditLogService
	Billing       *BillingService
	Invites       *InviteService
	Organizations *OrganizationService
	Subscriptions *SubscriptionService
	Tokens        *TokenService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.AuditLogs = NewAuditLogService(opts...)
	r.Billing = NewBillingService(opts...)
	r.Invites = NewInviteService(opts...)
	r.Organizations = NewOrganizationService(opts...)
	r.Subscriptions = NewSubscriptionService(opts...)
	r.Tokens = NewTokenService(opts...)
	return
}

// Edit part of your user details.
func (r *UserService) Edit(ctx context.Context, body UserEditParams, opts ...option.RequestOption) (res *UserEditResponse, err error) {
	var env UserEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// User Details
func (r *UserService) Get(ctx context.Context, opts ...option.RequestOption) (res *UserGetResponse, err error) {
	var env UserGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserEditResponse struct {
	// Identifier of the user.
	ID string `json:"id"`
	// Lists the betas that the user is participating in.
	Betas []string `json:"betas"`
	// The country in which the user lives.
	Country string `json:"country,nullable"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// Indicates whether user has any business zones
	HasBusinessZones bool `json:"has_business_zones"`
	// Indicates whether user has any enterprise zones
	HasEnterpriseZones bool `json:"has_enterprise_zones"`
	// Indicates whether user has any pro zones
	HasProZones bool `json:"has_pro_zones"`
	// User's last name
	LastName      string         `json:"last_name,nullable"`
	Organizations []Organization `json:"organizations"`
	// Indicates whether user has been suspended
	Suspended bool `json:"suspended"`
	// User's telephone number
	Telephone string `json:"telephone,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool `json:"two_factor_authentication_enabled"`
	// Indicates whether two-factor authentication is required by one of the accounts
	// that the user is a member of.
	TwoFactorAuthenticationLocked bool `json:"two_factor_authentication_locked"`
	// The zipcode or postal code where the user lives.
	Zipcode string               `json:"zipcode,nullable"`
	JSON    userEditResponseJSON `json:"-"`
}

// userEditResponseJSON contains the JSON metadata for the struct
// [UserEditResponse]
type userEditResponseJSON struct {
	ID                             apijson.Field
	Betas                          apijson.Field
	Country                        apijson.Field
	FirstName                      apijson.Field
	HasBusinessZones               apijson.Field
	HasEnterpriseZones             apijson.Field
	HasProZones                    apijson.Field
	LastName                       apijson.Field
	Organizations                  apijson.Field
	Suspended                      apijson.Field
	Telephone                      apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	TwoFactorAuthenticationLocked  apijson.Field
	Zipcode                        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *UserEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseJSON) RawJSON() string {
	return r.raw
}

type UserGetResponse struct {
	// Identifier of the user.
	ID string `json:"id"`
	// Lists the betas that the user is participating in.
	Betas []string `json:"betas"`
	// The country in which the user lives.
	Country string `json:"country,nullable"`
	// User's first name
	FirstName string `json:"first_name,nullable"`
	// Indicates whether user has any business zones
	HasBusinessZones bool `json:"has_business_zones"`
	// Indicates whether user has any enterprise zones
	HasEnterpriseZones bool `json:"has_enterprise_zones"`
	// Indicates whether user has any pro zones
	HasProZones bool `json:"has_pro_zones"`
	// User's last name
	LastName      string         `json:"last_name,nullable"`
	Organizations []Organization `json:"organizations"`
	// Indicates whether user has been suspended
	Suspended bool `json:"suspended"`
	// User's telephone number
	Telephone string `json:"telephone,nullable"`
	// Indicates whether two-factor authentication is enabled for the user account.
	// Does not apply to API authentication.
	TwoFactorAuthenticationEnabled bool `json:"two_factor_authentication_enabled"`
	// Indicates whether two-factor authentication is required by one of the accounts
	// that the user is a member of.
	TwoFactorAuthenticationLocked bool `json:"two_factor_authentication_locked"`
	// The zipcode or postal code where the user lives.
	Zipcode string              `json:"zipcode,nullable"`
	JSON    userGetResponseJSON `json:"-"`
}

// userGetResponseJSON contains the JSON metadata for the struct [UserGetResponse]
type userGetResponseJSON struct {
	ID                             apijson.Field
	Betas                          apijson.Field
	Country                        apijson.Field
	FirstName                      apijson.Field
	HasBusinessZones               apijson.Field
	HasEnterpriseZones             apijson.Field
	HasProZones                    apijson.Field
	LastName                       apijson.Field
	Organizations                  apijson.Field
	Suspended                      apijson.Field
	Telephone                      apijson.Field
	TwoFactorAuthenticationEnabled apijson.Field
	TwoFactorAuthenticationLocked  apijson.Field
	Zipcode                        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *UserGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseJSON) RawJSON() string {
	return r.raw
}

type UserEditParams struct {
	// The country in which the user lives.
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"first_name"`
	// User's last name
	LastName param.Field[string] `json:"last_name"`
	// User's telephone number
	Telephone param.Field[string] `json:"telephone"`
	// The zipcode or postal code where the user lives.
	Zipcode param.Field[string] `json:"zipcode"`
}

func (r UserEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserEditResponseEnvelope struct {
	Errors   []UserEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserEditResponseEnvelopeSuccess `json:"success,required"`
	Result  UserEditResponse                `json:"result"`
	JSON    userEditResponseEnvelopeJSON    `json:"-"`
}

// userEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelope]
type userEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserEditResponseEnvelopeErrors struct {
	Code             int64                                `json:"code,required"`
	Message          string                               `json:"message,required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           UserEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             userEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// userEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelopeErrors]
type userEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserEditResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    userEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [UserEditResponseEnvelopeErrorsSource]
type userEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserEditResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           UserEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             userEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// userEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UserEditResponseEnvelopeMessages]
type userEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserEditResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    userEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [UserEditResponseEnvelopeMessagesSource]
type userEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserEditResponseEnvelopeSuccess bool

const (
	UserEditResponseEnvelopeSuccessTrue UserEditResponseEnvelopeSuccess = true
)

func (r UserEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type UserGetResponseEnvelope struct {
	Errors   []UserGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success UserGetResponseEnvelopeSuccess `json:"success,required"`
	Result  UserGetResponse                `json:"result"`
	JSON    userGetResponseEnvelopeJSON    `json:"-"`
}

// userGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelope]
type userGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type UserGetResponseEnvelopeErrors struct {
	Code             int64                               `json:"code,required"`
	Message          string                              `json:"message,required"`
	DocumentationURL string                              `json:"documentation_url"`
	Source           UserGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             userGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// userGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelopeErrors]
type userGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type UserGetResponseEnvelopeErrorsSource struct {
	Pointer string                                  `json:"pointer"`
	JSON    userGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// userGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [UserGetResponseEnvelopeErrorsSource]
type userGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type UserGetResponseEnvelopeMessages struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           UserGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             userGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// userGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [UserGetResponseEnvelopeMessages]
type userGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type UserGetResponseEnvelopeMessagesSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    userGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// userGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [UserGetResponseEnvelopeMessagesSource]
type userGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type UserGetResponseEnvelopeSuccess bool

const (
	UserGetResponseEnvelopeSuccessTrue UserGetResponseEnvelopeSuccess = true
)

func (r UserGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case UserGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
