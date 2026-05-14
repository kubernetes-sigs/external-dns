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

// AccessApplicationUserPolicyCheckService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessApplicationUserPolicyCheckService] method instead.
type AccessApplicationUserPolicyCheckService struct {
	Options []option.RequestOption
}

// NewAccessApplicationUserPolicyCheckService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessApplicationUserPolicyCheckService(opts ...option.RequestOption) (r *AccessApplicationUserPolicyCheckService) {
	r = &AccessApplicationUserPolicyCheckService{}
	r.Options = opts
	return
}

// Tests if a specific user has permission to access an application.
func (r *AccessApplicationUserPolicyCheckService) List(ctx context.Context, appID AppIDParam, query AccessApplicationUserPolicyCheckListParams, opts ...option.RequestOption) (res *AccessApplicationUserPolicyCheckListResponse, err error) {
	var env AccessApplicationUserPolicyCheckListResponseEnvelope
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
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/user_policy_checks", accountOrZone, accountOrZoneID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserPolicyCheckGeo struct {
	Country string                 `json:"country"`
	JSON    userPolicyCheckGeoJSON `json:"-"`
}

// userPolicyCheckGeoJSON contains the JSON metadata for the struct
// [UserPolicyCheckGeo]
type userPolicyCheckGeoJSON struct {
	Country     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserPolicyCheckGeo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userPolicyCheckGeoJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListResponse struct {
	AppState     AccessApplicationUserPolicyCheckListResponseAppState     `json:"app_state"`
	UserIdentity AccessApplicationUserPolicyCheckListResponseUserIdentity `json:"user_identity"`
	JSON         accessApplicationUserPolicyCheckListResponseJSON         `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseJSON contains the JSON metadata for
// the struct [AccessApplicationUserPolicyCheckListResponse]
type accessApplicationUserPolicyCheckListResponseJSON struct {
	AppState     apijson.Field
	UserIdentity apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListResponseAppState struct {
	// UUID.
	AppUID   string                                                   `json:"app_uid"`
	AUD      string                                                   `json:"aud"`
	Hostname string                                                   `json:"hostname"`
	Name     string                                                   `json:"name"`
	Policies []interface{}                                            `json:"policies"`
	Status   string                                                   `json:"status"`
	JSON     accessApplicationUserPolicyCheckListResponseAppStateJSON `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseAppStateJSON contains the JSON
// metadata for the struct [AccessApplicationUserPolicyCheckListResponseAppState]
type accessApplicationUserPolicyCheckListResponseAppStateJSON struct {
	AppUID      apijson.Field
	AUD         apijson.Field
	Hostname    apijson.Field
	Name        apijson.Field
	Policies    apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponseAppState) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseAppStateJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListResponseUserIdentity struct {
	ID             string             `json:"id"`
	AccountID      string             `json:"account_id"`
	DeviceSessions interface{}        `json:"device_sessions"`
	Email          string             `json:"email"`
	Geo            UserPolicyCheckGeo `json:"geo"`
	Iat            int64              `json:"iat"`
	IsGateway      bool               `json:"is_gateway"`
	IsWARP         bool               `json:"is_warp"`
	Name           string             `json:"name"`
	// UUID.
	UserUUID string                                                       `json:"user_uuid"`
	Version  int64                                                        `json:"version"`
	JSON     accessApplicationUserPolicyCheckListResponseUserIdentityJSON `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseUserIdentityJSON contains the JSON
// metadata for the struct
// [AccessApplicationUserPolicyCheckListResponseUserIdentity]
type accessApplicationUserPolicyCheckListResponseUserIdentityJSON struct {
	ID             apijson.Field
	AccountID      apijson.Field
	DeviceSessions apijson.Field
	Email          apijson.Field
	Geo            apijson.Field
	Iat            apijson.Field
	IsGateway      apijson.Field
	IsWARP         apijson.Field
	Name           apijson.Field
	UserUUID       apijson.Field
	Version        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponseUserIdentity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseUserIdentityJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationUserPolicyCheckListResponseEnvelope struct {
	Errors   []AccessApplicationUserPolicyCheckListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationUserPolicyCheckListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessApplicationUserPolicyCheckListResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessApplicationUserPolicyCheckListResponse                `json:"result"`
	JSON    accessApplicationUserPolicyCheckListResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseEnvelopeJSON contains the JSON
// metadata for the struct [AccessApplicationUserPolicyCheckListResponseEnvelope]
type accessApplicationUserPolicyCheckListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListResponseEnvelopeErrors struct {
	Code             int64                                                            `json:"code,required"`
	Message          string                                                           `json:"message,required"`
	DocumentationURL string                                                           `json:"documentation_url"`
	Source           AccessApplicationUserPolicyCheckListResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessApplicationUserPolicyCheckListResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [AccessApplicationUserPolicyCheckListResponseEnvelopeErrors]
type accessApplicationUserPolicyCheckListResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListResponseEnvelopeErrorsSource struct {
	Pointer string                                                               `json:"pointer"`
	JSON    accessApplicationUserPolicyCheckListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseEnvelopeErrorsSourceJSON contains
// the JSON metadata for the struct
// [AccessApplicationUserPolicyCheckListResponseEnvelopeErrorsSource]
type accessApplicationUserPolicyCheckListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListResponseEnvelopeMessages struct {
	Code             int64                                                              `json:"code,required"`
	Message          string                                                             `json:"message,required"`
	DocumentationURL string                                                             `json:"documentation_url"`
	Source           AccessApplicationUserPolicyCheckListResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessApplicationUserPolicyCheckListResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [AccessApplicationUserPolicyCheckListResponseEnvelopeMessages]
type accessApplicationUserPolicyCheckListResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationUserPolicyCheckListResponseEnvelopeMessagesSource struct {
	Pointer string                                                                 `json:"pointer"`
	JSON    accessApplicationUserPolicyCheckListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessApplicationUserPolicyCheckListResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [AccessApplicationUserPolicyCheckListResponseEnvelopeMessagesSource]
type accessApplicationUserPolicyCheckListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationUserPolicyCheckListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationUserPolicyCheckListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessApplicationUserPolicyCheckListResponseEnvelopeSuccess bool

const (
	AccessApplicationUserPolicyCheckListResponseEnvelopeSuccessTrue AccessApplicationUserPolicyCheckListResponseEnvelopeSuccess = true
)

func (r AccessApplicationUserPolicyCheckListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationUserPolicyCheckListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
