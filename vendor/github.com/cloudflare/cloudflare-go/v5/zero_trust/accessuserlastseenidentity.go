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

// AccessUserLastSeenIdentityService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserLastSeenIdentityService] method instead.
type AccessUserLastSeenIdentityService struct {
	Options []option.RequestOption
}

// NewAccessUserLastSeenIdentityService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessUserLastSeenIdentityService(opts ...option.RequestOption) (r *AccessUserLastSeenIdentityService) {
	r = &AccessUserLastSeenIdentityService{}
	r.Options = opts
	return
}

// Get last seen identity for a single user.
func (r *AccessUserLastSeenIdentityService) Get(ctx context.Context, userID string, query AccessUserLastSeenIdentityGetParams, opts ...option.RequestOption) (res *Identity, err error) {
	var env AccessUserLastSeenIdentityGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/last_seen_identity", query.AccountID, userID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Identity struct {
	AccountID          string                           `json:"account_id"`
	AuthStatus         string                           `json:"auth_status"`
	CommonName         string                           `json:"common_name"`
	DeviceID           string                           `json:"device_id"`
	DeviceSessions     map[string]IdentityDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]IdentityDevicePosture `json:"devicePosture"`
	Email              string                           `json:"email"`
	Geo                UserPolicyCheckGeo               `json:"geo"`
	Iat                float64                          `json:"iat"`
	IdP                IdentityIdP                      `json:"idp"`
	IP                 string                           `json:"ip"`
	IsGateway          bool                             `json:"is_gateway"`
	IsWARP             bool                             `json:"is_warp"`
	MTLSAuth           IdentityMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                           `json:"service_token_id"`
	ServiceTokenStatus bool                             `json:"service_token_status"`
	UserUUID           string                           `json:"user_uuid"`
	Version            float64                          `json:"version"`
	JSON               identityJSON                     `json:"-"`
}

// identityJSON contains the JSON metadata for the struct [Identity]
type identityJSON struct {
	AccountID          apijson.Field
	AuthStatus         apijson.Field
	CommonName         apijson.Field
	DeviceID           apijson.Field
	DeviceSessions     apijson.Field
	DevicePosture      apijson.Field
	Email              apijson.Field
	Geo                apijson.Field
	Iat                apijson.Field
	IdP                apijson.Field
	IP                 apijson.Field
	IsGateway          apijson.Field
	IsWARP             apijson.Field
	MTLSAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUUID           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *Identity) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityJSON) RawJSON() string {
	return r.raw
}

type IdentityDeviceSession struct {
	LastAuthenticated float64                   `json:"last_authenticated"`
	JSON              identityDeviceSessionJSON `json:"-"`
}

// identityDeviceSessionJSON contains the JSON metadata for the struct
// [IdentityDeviceSession]
type identityDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *IdentityDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type IdentityDevicePosture struct {
	ID          string                     `json:"id"`
	Check       IdentityDevicePostureCheck `json:"check"`
	Data        interface{}                `json:"data"`
	Description string                     `json:"description"`
	Error       string                     `json:"error"`
	RuleName    string                     `json:"rule_name"`
	Success     bool                       `json:"success"`
	Timestamp   string                     `json:"timestamp"`
	Type        string                     `json:"type"`
	JSON        identityDevicePostureJSON  `json:"-"`
}

// identityDevicePostureJSON contains the JSON metadata for the struct
// [IdentityDevicePosture]
type identityDevicePostureJSON struct {
	ID          apijson.Field
	Check       apijson.Field
	Data        apijson.Field
	Description apijson.Field
	Error       apijson.Field
	RuleName    apijson.Field
	Success     apijson.Field
	Timestamp   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityDevicePostureJSON) RawJSON() string {
	return r.raw
}

type IdentityDevicePostureCheck struct {
	Exists bool                           `json:"exists"`
	Path   string                         `json:"path"`
	JSON   identityDevicePostureCheckJSON `json:"-"`
}

// identityDevicePostureCheckJSON contains the JSON metadata for the struct
// [IdentityDevicePostureCheck]
type identityDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type IdentityIdP struct {
	ID   string          `json:"id"`
	Type string          `json:"type"`
	JSON identityIdPJSON `json:"-"`
}

// identityIdPJSON contains the JSON metadata for the struct [IdentityIdP]
type identityIdPJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IdentityIdP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityIdPJSON) RawJSON() string {
	return r.raw
}

type IdentityMTLSAuth struct {
	AuthStatus    string               `json:"auth_status"`
	CERTIssuerDn  string               `json:"cert_issuer_dn"`
	CERTIssuerSki string               `json:"cert_issuer_ski"`
	CERTPresented bool                 `json:"cert_presented"`
	CERTSerial    string               `json:"cert_serial"`
	JSON          identityMTLSAuthJSON `json:"-"`
}

// identityMTLSAuthJSON contains the JSON metadata for the struct
// [IdentityMTLSAuth]
type identityMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CERTIssuerDn  apijson.Field
	CERTIssuerSki apijson.Field
	CERTPresented apijson.Field
	CERTSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *IdentityMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r identityMTLSAuthJSON) RawJSON() string {
	return r.raw
}

type AccessUserLastSeenIdentityGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessUserLastSeenIdentityGetResponseEnvelope struct {
	Errors   []AccessUserLastSeenIdentityGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserLastSeenIdentityGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessUserLastSeenIdentityGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Identity                                             `json:"result"`
	JSON    accessUserLastSeenIdentityGetResponseEnvelopeJSON    `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserLastSeenIdentityGetResponseEnvelope]
type accessUserLastSeenIdentityGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserLastSeenIdentityGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserLastSeenIdentityGetResponseEnvelopeErrors struct {
	Code             int64                                                     `json:"code,required"`
	Message          string                                                    `json:"message,required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           AccessUserLastSeenIdentityGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityGetResponseEnvelopeErrors]
type accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserLastSeenIdentityGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserLastSeenIdentityGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    accessUserLastSeenIdentityGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [AccessUserLastSeenIdentityGetResponseEnvelopeErrorsSource]
type accessUserLastSeenIdentityGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserLastSeenIdentityGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserLastSeenIdentityGetResponseEnvelopeMessages struct {
	Code             int64                                                       `json:"code,required"`
	Message          string                                                      `json:"message,required"`
	DocumentationURL string                                                      `json:"documentation_url"`
	Source           AccessUserLastSeenIdentityGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserLastSeenIdentityGetResponseEnvelopeMessages]
type accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserLastSeenIdentityGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserLastSeenIdentityGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                          `json:"pointer"`
	JSON    accessUserLastSeenIdentityGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserLastSeenIdentityGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [AccessUserLastSeenIdentityGetResponseEnvelopeMessagesSource]
type accessUserLastSeenIdentityGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserLastSeenIdentityGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserLastSeenIdentityGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserLastSeenIdentityGetResponseEnvelopeSuccess bool

const (
	AccessUserLastSeenIdentityGetResponseEnvelopeSuccessTrue AccessUserLastSeenIdentityGetResponseEnvelopeSuccess = true
)

func (r AccessUserLastSeenIdentityGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserLastSeenIdentityGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
