// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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
)

// OrganizationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrganizationService] method instead.
type OrganizationService struct {
	Options []option.RequestOption
	DOH     *OrganizationDOHService
}

// NewOrganizationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrganizationService(opts ...option.RequestOption) (r *OrganizationService) {
	r = &OrganizationService{}
	r.Options = opts
	r.DOH = NewOrganizationDOHService(opts...)
	return
}

// Sets up a Zero Trust organization for your account or zone.
func (r *OrganizationService) New(ctx context.Context, params OrganizationNewParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates the configuration for your Zero Trust organization.
func (r *OrganizationService) Update(ctx context.Context, params OrganizationUpdateParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Returns the configuration for your Zero Trust organization.
func (r *OrganizationService) List(ctx context.Context, query OrganizationListParams, opts ...option.RequestOption) (res *Organization, err error) {
	var env OrganizationListResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Revokes a user's access across all applications.
func (r *OrganizationService) RevokeUsers(ctx context.Context, params OrganizationRevokeUsersParams, opts ...option.RequestOption) (res *OrganizationRevokeUsersResponse, err error) {
	var env OrganizationRevokeUsersResponseEnvelope
	opts = append(r.Options[:], opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/organizations/revoke_user", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoginDesign struct {
	// The background color on your login page.
	BackgroundColor string `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText string `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText string `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath string `json:"logo_path"`
	// The text color on your login page.
	TextColor string          `json:"text_color"`
	JSON      loginDesignJSON `json:"-"`
}

// loginDesignJSON contains the JSON metadata for the struct [LoginDesign]
type loginDesignJSON struct {
	BackgroundColor apijson.Field
	FooterText      apijson.Field
	HeaderText      apijson.Field
	LogoPath        apijson.Field
	TextColor       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LoginDesign) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loginDesignJSON) RawJSON() string {
	return r.raw
}

type LoginDesignParam struct {
	// The background color on your login page.
	BackgroundColor param.Field[string] `json:"background_color"`
	// The text at the bottom of your login page.
	FooterText param.Field[string] `json:"footer_text"`
	// The text at the top of your login page.
	HeaderText param.Field[string] `json:"header_text"`
	// The URL of the logo on your login page.
	LogoPath param.Field[string] `json:"logo_path"`
	// The text color on your login page.
	TextColor param.Field[string] `json:"text_color"`
}

func (r LoginDesignParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type Organization struct {
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP bool `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain string `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity bool                    `json:"auto_redirect_to_identity"`
	CreatedAt              time.Time               `json:"created_at" format:"date-time"`
	CustomPages            OrganizationCustomPages `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUIReadOnly bool        `json:"is_ui_read_only"`
	LoginDesign  LoginDesign `json:"login_design"`
	// The name of your Zero Trust organization.
	Name string `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration string `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason string    `json:"ui_read_only_toggle_reason"`
	UpdatedAt              time.Time `json:"updated_at" format:"date-time"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime string `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration string           `json:"warp_auth_session_duration"`
	JSON                    organizationJSON `json:"-"`
}

// organizationJSON contains the JSON metadata for the struct [Organization]
type organizationJSON struct {
	AllowAuthenticateViaWARP       apijson.Field
	AuthDomain                     apijson.Field
	AutoRedirectToIdentity         apijson.Field
	CreatedAt                      apijson.Field
	CustomPages                    apijson.Field
	IsUIReadOnly                   apijson.Field
	LoginDesign                    apijson.Field
	Name                           apijson.Field
	SessionDuration                apijson.Field
	UIReadOnlyToggleReason         apijson.Field
	UpdatedAt                      apijson.Field
	UserSeatExpirationInactiveTime apijson.Field
	WARPAuthSessionDuration        apijson.Field
	raw                            string
	ExtraFields                    map[string]apijson.Field
}

func (r *Organization) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationJSON) RawJSON() string {
	return r.raw
}

type OrganizationCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden string `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied string                      `json:"identity_denied"`
	JSON           organizationCustomPagesJSON `json:"-"`
}

// organizationCustomPagesJSON contains the JSON metadata for the struct
// [OrganizationCustomPages]
type organizationCustomPagesJSON struct {
	Forbidden      apijson.Field
	IdentityDenied apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrganizationCustomPages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationCustomPagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationRevokeUsersResponse bool

const (
	OrganizationRevokeUsersResponseTrue  OrganizationRevokeUsersResponse = true
	OrganizationRevokeUsersResponseFalse OrganizationRevokeUsersResponse = false
)

func (r OrganizationRevokeUsersResponse) IsKnown() bool {
	switch r {
	case OrganizationRevokeUsersResponseTrue, OrganizationRevokeUsersResponseFalse:
		return true
	}
	return false
}

type OrganizationNewParams struct {
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain,required"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool] `json:"auto_redirect_to_identity"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUIReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r OrganizationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationNewResponseEnvelope struct {
	Errors   []OrganizationNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OrganizationNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OrganizationNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Organization                           `json:"result"`
	JSON    organizationNewResponseEnvelopeJSON    `json:"-"`
}

// organizationNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationNewResponseEnvelope]
type organizationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           OrganizationNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationNewResponseEnvelopeErrors]
type organizationNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    organizationNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OrganizationNewResponseEnvelopeErrorsSource]
type organizationNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           OrganizationNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OrganizationNewResponseEnvelopeMessages]
type organizationNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationNewResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    organizationNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [OrganizationNewResponseEnvelopeMessagesSource]
type organizationNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationNewResponseEnvelopeSuccess bool

const (
	OrganizationNewResponseEnvelopeSuccessTrue OrganizationNewResponseEnvelopeSuccess = true
)

func (r OrganizationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to true, users can authenticate via WARP for any application in your
	// organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWARP param.Field[bool] `json:"allow_authenticate_via_warp"`
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain param.Field[string] `json:"auth_domain"`
	// When set to `true`, users skip the identity provider selection step during
	// login.
	AutoRedirectToIdentity param.Field[bool]                                `json:"auto_redirect_to_identity"`
	CustomPages            param.Field[OrganizationUpdateParamsCustomPages] `json:"custom_pages"`
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// Updates may only be made via the API or Terraform for this account when enabled.
	IsUIReadOnly param.Field[bool]             `json:"is_ui_read_only"`
	LoginDesign  param.Field[LoginDesignParam] `json:"login_design"`
	// The name of your Zero Trust organization.
	Name param.Field[string] `json:"name"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m,
	// h.
	SessionDuration param.Field[string] `json:"session_duration"`
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason param.Field[string] `json:"ui_read_only_toggle_reason"`
	// The amount of time a user seat is inactive before it expires. When the user seat
	// exceeds the set time of inactivity, the user is removed as an active seat and no
	// longer counts against your Teams seat count. Minimum value for this setting is 1
	// month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are:
	// `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime param.Field[string] `json:"user_seat_expiration_inactive_time"`
	// The amount of time that tokens issued for applications will be valid. Must be in
	// the format `30m` or `2h45m`. Valid time units are: m, h.
	WARPAuthSessionDuration param.Field[string] `json:"warp_auth_session_duration"`
}

func (r OrganizationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateParamsCustomPages struct {
	// The uid of the custom page to use when a user is denied access after failing a
	// non-identity rule.
	Forbidden param.Field[string] `json:"forbidden"`
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied param.Field[string] `json:"identity_denied"`
}

func (r OrganizationUpdateParamsCustomPages) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrganizationUpdateResponseEnvelope struct {
	Errors   []OrganizationUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OrganizationUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OrganizationUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  Organization                              `json:"result"`
	JSON    organizationUpdateResponseEnvelopeJSON    `json:"-"`
}

// organizationUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationUpdateResponseEnvelope]
type organizationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           OrganizationUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationUpdateResponseEnvelopeErrors]
type organizationUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    organizationUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [OrganizationUpdateResponseEnvelopeErrorsSource]
type organizationUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           OrganizationUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [OrganizationUpdateResponseEnvelopeMessages]
type organizationUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    organizationUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [OrganizationUpdateResponseEnvelopeMessagesSource]
type organizationUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationUpdateResponseEnvelopeSuccess bool

const (
	OrganizationUpdateResponseEnvelopeSuccessTrue OrganizationUpdateResponseEnvelopeSuccess = true
)

func (r OrganizationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type OrganizationListResponseEnvelope struct {
	Errors   []OrganizationListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []OrganizationListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success OrganizationListResponseEnvelopeSuccess `json:"success,required"`
	Result  Organization                            `json:"result"`
	JSON    organizationListResponseEnvelopeJSON    `json:"-"`
}

// organizationListResponseEnvelopeJSON contains the JSON metadata for the struct
// [OrganizationListResponseEnvelope]
type organizationListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           OrganizationListResponseEnvelopeErrorsSource `json:"source"`
	JSON             organizationListResponseEnvelopeErrorsJSON   `json:"-"`
}

// organizationListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [OrganizationListResponseEnvelopeErrors]
type organizationListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    organizationListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// organizationListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [OrganizationListResponseEnvelopeErrorsSource]
type organizationListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           OrganizationListResponseEnvelopeMessagesSource `json:"source"`
	JSON             organizationListResponseEnvelopeMessagesJSON   `json:"-"`
}

// organizationListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [OrganizationListResponseEnvelopeMessages]
type organizationListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type OrganizationListResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    organizationListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// organizationListResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [OrganizationListResponseEnvelopeMessagesSource]
type organizationListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type OrganizationListResponseEnvelopeSuccess bool

const (
	OrganizationListResponseEnvelopeSuccessTrue OrganizationListResponseEnvelopeSuccess = true
)

func (r OrganizationListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type OrganizationRevokeUsersParams struct {
	// The email of the user to revoke.
	Email param.Field[string] `json:"email,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// When set to `true`, all devices associated with the user will be revoked.
	QueryDevices param.Field[bool] `query:"devices"`
	// When set to `true`, all devices associated with the user will be revoked.
	BodyDevices param.Field[bool] `json:"devices"`
	// The uuid of the user to revoke.
	UserUID param.Field[string] `json:"user_uid"`
	// When set to `true`, the user will be required to re-authenticate to WARP for all
	// Gateway policies that enforce a WARP client session duration. When `false`, the
	// user’s WARP session will remain active
	WARPSessionReauth param.Field[bool] `json:"warp_session_reauth"`
}

func (r OrganizationRevokeUsersParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrganizationRevokeUsersParams]'s query parameters as
// `url.Values`.
func (r OrganizationRevokeUsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type OrganizationRevokeUsersResponseEnvelope struct {
	Result  OrganizationRevokeUsersResponse                `json:"result"`
	Success OrganizationRevokeUsersResponseEnvelopeSuccess `json:"success"`
	JSON    organizationRevokeUsersResponseEnvelopeJSON    `json:"-"`
}

// organizationRevokeUsersResponseEnvelopeJSON contains the JSON metadata for the
// struct [OrganizationRevokeUsersResponseEnvelope]
type organizationRevokeUsersResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrganizationRevokeUsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r organizationRevokeUsersResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type OrganizationRevokeUsersResponseEnvelopeSuccess bool

const (
	OrganizationRevokeUsersResponseEnvelopeSuccessTrue  OrganizationRevokeUsersResponseEnvelopeSuccess = true
	OrganizationRevokeUsersResponseEnvelopeSuccessFalse OrganizationRevokeUsersResponseEnvelopeSuccess = false
)

func (r OrganizationRevokeUsersResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case OrganizationRevokeUsersResponseEnvelopeSuccessTrue, OrganizationRevokeUsersResponseEnvelopeSuccessFalse:
		return true
	}
	return false
}
