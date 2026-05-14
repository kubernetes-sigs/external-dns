// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v5/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v5/internal/param"
	"github.com/cloudflare/cloudflare-go/v5/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// InviteService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInviteService] method instead.
type InviteService struct {
	Options []option.RequestOption
}

// NewInviteService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewInviteService(opts ...option.RequestOption) (r *InviteService) {
	r = &InviteService{}
	r.Options = opts
	return
}

// Lists all invitations associated with my user.
func (r *InviteService) List(ctx context.Context, opts ...option.RequestOption) (res *pagination.SinglePage[Invite], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/invites"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all invitations associated with my user.
func (r *InviteService) ListAutoPaging(ctx context.Context, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Invite] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, opts...))
}

// Responds to an invitation.
func (r *InviteService) Edit(ctx context.Context, inviteID string, body InviteEditParams, opts ...option.RequestOption) (res *Invite, err error) {
	var env InviteEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return
	}
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets the details of an invitation.
func (r *InviteService) Get(ctx context.Context, inviteID string, opts ...option.RequestOption) (res *Invite, err error) {
	var env InviteGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if inviteID == "" {
		err = errors.New("missing required invite_id parameter")
		return
	}
	path := fmt.Sprintf("user/invites/%s", inviteID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Invite struct {
	// ID of the user to add to the organization.
	InvitedMemberID string `json:"invited_member_id,required,nullable"`
	// ID of the organization the user will be added to.
	OrganizationID string `json:"organization_id,required"`
	// Invite identifier tag.
	ID string `json:"id"`
	// When the invite is no longer active.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The email address of the user who created the invite.
	InvitedBy string `json:"invited_by"`
	// Email address of the user to add to the organization.
	InvitedMemberEmail string `json:"invited_member_email"`
	// When the invite was sent.
	InvitedOn                        time.Time `json:"invited_on" format:"date-time"`
	OrganizationIsEnforcingTwofactor bool      `json:"organization_is_enforcing_twofactor"`
	// Organization name.
	OrganizationName string `json:"organization_name"`
	// List of role names the membership has for this account.
	Roles []string `json:"roles"`
	// Current status of the invitation.
	Status InviteStatus `json:"status"`
	JSON   inviteJSON   `json:"-"`
}

// inviteJSON contains the JSON metadata for the struct [Invite]
type inviteJSON struct {
	InvitedMemberID                  apijson.Field
	OrganizationID                   apijson.Field
	ID                               apijson.Field
	ExpiresOn                        apijson.Field
	InvitedBy                        apijson.Field
	InvitedMemberEmail               apijson.Field
	InvitedOn                        apijson.Field
	OrganizationIsEnforcingTwofactor apijson.Field
	OrganizationName                 apijson.Field
	Roles                            apijson.Field
	Status                           apijson.Field
	raw                              string
	ExtraFields                      map[string]apijson.Field
}

func (r *Invite) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteJSON) RawJSON() string {
	return r.raw
}

// Current status of the invitation.
type InviteStatus string

const (
	InviteStatusPending  InviteStatus = "pending"
	InviteStatusAccepted InviteStatus = "accepted"
	InviteStatusRejected InviteStatus = "rejected"
	InviteStatusExpired  InviteStatus = "expired"
)

func (r InviteStatus) IsKnown() bool {
	switch r {
	case InviteStatusPending, InviteStatusAccepted, InviteStatusRejected, InviteStatusExpired:
		return true
	}
	return false
}

type InviteEditParams struct {
	// Status of your response to the invitation (rejected or accepted).
	Status param.Field[InviteEditParamsStatus] `json:"status,required"`
}

func (r InviteEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Status of your response to the invitation (rejected or accepted).
type InviteEditParamsStatus string

const (
	InviteEditParamsStatusAccepted InviteEditParamsStatus = "accepted"
	InviteEditParamsStatusRejected InviteEditParamsStatus = "rejected"
)

func (r InviteEditParamsStatus) IsKnown() bool {
	switch r {
	case InviteEditParamsStatusAccepted, InviteEditParamsStatusRejected:
		return true
	}
	return false
}

type InviteEditResponseEnvelope struct {
	Errors   []InviteEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []InviteEditResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success InviteEditResponseEnvelopeSuccess `json:"success,required"`
	Result  Invite                            `json:"result"`
	JSON    inviteEditResponseEnvelopeJSON    `json:"-"`
}

// inviteEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [InviteEditResponseEnvelope]
type inviteEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InviteEditResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           InviteEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             inviteEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// inviteEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [InviteEditResponseEnvelopeErrors]
type inviteEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InviteEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InviteEditResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    inviteEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// inviteEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [InviteEditResponseEnvelopeErrorsSource]
type inviteEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InviteEditResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           InviteEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             inviteEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// inviteEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [InviteEditResponseEnvelopeMessages]
type inviteEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InviteEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InviteEditResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    inviteEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// inviteEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [InviteEditResponseEnvelopeMessagesSource]
type inviteEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InviteEditResponseEnvelopeSuccess bool

const (
	InviteEditResponseEnvelopeSuccessTrue InviteEditResponseEnvelopeSuccess = true
)

func (r InviteEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InviteEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type InviteGetResponseEnvelope struct {
	Errors   []InviteGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []InviteGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success InviteGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Invite                           `json:"result"`
	JSON    inviteGetResponseEnvelopeJSON    `json:"-"`
}

// inviteGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [InviteGetResponseEnvelope]
type inviteGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type InviteGetResponseEnvelopeErrors struct {
	Code             int64                                 `json:"code,required"`
	Message          string                                `json:"message,required"`
	DocumentationURL string                                `json:"documentation_url"`
	Source           InviteGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             inviteGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// inviteGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [InviteGetResponseEnvelopeErrors]
type inviteGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InviteGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type InviteGetResponseEnvelopeErrorsSource struct {
	Pointer string                                    `json:"pointer"`
	JSON    inviteGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// inviteGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [InviteGetResponseEnvelopeErrorsSource]
type inviteGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type InviteGetResponseEnvelopeMessages struct {
	Code             int64                                   `json:"code,required"`
	Message          string                                  `json:"message,required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           InviteGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             inviteGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// inviteGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [InviteGetResponseEnvelopeMessages]
type inviteGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *InviteGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type InviteGetResponseEnvelopeMessagesSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    inviteGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// inviteGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [InviteGetResponseEnvelopeMessagesSource]
type inviteGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *InviteGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r inviteGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type InviteGetResponseEnvelopeSuccess bool

const (
	InviteGetResponseEnvelopeSuccessTrue InviteGetResponseEnvelopeSuccess = true
)

func (r InviteGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case InviteGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
