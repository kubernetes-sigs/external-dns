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
	"github.com/cloudflare/cloudflare-go/v5/packages/pagination"
)

// AccessUserActiveSessionService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserActiveSessionService] method instead.
type AccessUserActiveSessionService struct {
	Options []option.RequestOption
}

// NewAccessUserActiveSessionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessUserActiveSessionService(opts ...option.RequestOption) (r *AccessUserActiveSessionService) {
	r = &AccessUserActiveSessionService{}
	r.Options = opts
	return
}

// Get active sessions for a single user.
func (r *AccessUserActiveSessionService) List(ctx context.Context, userID string, query AccessUserActiveSessionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccessUserActiveSessionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions", query.AccountID, userID)
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

// Get active sessions for a single user.
func (r *AccessUserActiveSessionService) ListAutoPaging(ctx context.Context, userID string, query AccessUserActiveSessionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessUserActiveSessionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, userID, query, opts...))
}

// Get an active session for a single user.
func (r *AccessUserActiveSessionService) Get(ctx context.Context, userID string, nonce string, query AccessUserActiveSessionGetParams, opts ...option.RequestOption) (res *AccessUserActiveSessionGetResponse, err error) {
	var env AccessUserActiveSessionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return
	}
	if nonce == "" {
		err = errors.New("missing required nonce parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users/%s/active_sessions/%s", query.AccountID, userID, nonce)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessUserActiveSessionListResponse struct {
	Expiration int64                                       `json:"expiration"`
	Metadata   AccessUserActiveSessionListResponseMetadata `json:"metadata"`
	Name       string                                      `json:"name"`
	JSON       accessUserActiveSessionListResponseJSON     `json:"-"`
}

// accessUserActiveSessionListResponseJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionListResponse]
type accessUserActiveSessionListResponseJSON struct {
	Expiration  apijson.Field
	Metadata    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListResponseMetadata struct {
	Apps    map[string]AccessUserActiveSessionListResponseMetadataApp `json:"apps"`
	Expires int64                                                     `json:"expires"`
	Iat     int64                                                     `json:"iat"`
	Nonce   string                                                    `json:"nonce"`
	TTL     int64                                                     `json:"ttl"`
	JSON    accessUserActiveSessionListResponseMetadataJSON           `json:"-"`
}

// accessUserActiveSessionListResponseMetadataJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionListResponseMetadata]
type accessUserActiveSessionListResponseMetadataJSON struct {
	Apps        apijson.Field
	Expires     apijson.Field
	Iat         apijson.Field
	Nonce       apijson.Field
	TTL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseMetadataJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListResponseMetadataApp struct {
	Hostname string                                             `json:"hostname"`
	Name     string                                             `json:"name"`
	Type     string                                             `json:"type"`
	UID      string                                             `json:"uid"`
	JSON     accessUserActiveSessionListResponseMetadataAppJSON `json:"-"`
}

// accessUserActiveSessionListResponseMetadataAppJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionListResponseMetadataApp]
type accessUserActiveSessionListResponseMetadataAppJSON struct {
	Hostname    apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionListResponseMetadataApp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionListResponseMetadataAppJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponse struct {
	AccountID          string                                                     `json:"account_id"`
	AuthStatus         string                                                     `json:"auth_status"`
	CommonName         string                                                     `json:"common_name"`
	DeviceID           string                                                     `json:"device_id"`
	DeviceSessions     map[string]AccessUserActiveSessionGetResponseDeviceSession `json:"device_sessions"`
	DevicePosture      map[string]AccessUserActiveSessionGetResponseDevicePosture `json:"devicePosture"`
	Email              string                                                     `json:"email"`
	Geo                UserPolicyCheckGeo                                         `json:"geo"`
	Iat                float64                                                    `json:"iat"`
	IdP                AccessUserActiveSessionGetResponseIdP                      `json:"idp"`
	IP                 string                                                     `json:"ip"`
	IsGateway          bool                                                       `json:"is_gateway"`
	IsWARP             bool                                                       `json:"is_warp"`
	IsActive           bool                                                       `json:"isActive"`
	MTLSAuth           AccessUserActiveSessionGetResponseMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                                                     `json:"service_token_id"`
	ServiceTokenStatus bool                                                       `json:"service_token_status"`
	UserUUID           string                                                     `json:"user_uuid"`
	Version            float64                                                    `json:"version"`
	JSON               accessUserActiveSessionGetResponseJSON                     `json:"-"`
}

// accessUserActiveSessionGetResponseJSON contains the JSON metadata for the struct
// [AccessUserActiveSessionGetResponse]
type accessUserActiveSessionGetResponseJSON struct {
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
	IsActive           apijson.Field
	MTLSAuth           apijson.Field
	ServiceTokenID     apijson.Field
	ServiceTokenStatus apijson.Field
	UserUUID           apijson.Field
	Version            apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseDeviceSession struct {
	LastAuthenticated float64                                             `json:"last_authenticated"`
	JSON              accessUserActiveSessionGetResponseDeviceSessionJSON `json:"-"`
}

// accessUserActiveSessionGetResponseDeviceSessionJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionGetResponseDeviceSession]
type accessUserActiveSessionGetResponseDeviceSessionJSON struct {
	LastAuthenticated apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseDeviceSession) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseDeviceSessionJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseDevicePosture struct {
	ID          string                                               `json:"id"`
	Check       AccessUserActiveSessionGetResponseDevicePostureCheck `json:"check"`
	Data        interface{}                                          `json:"data"`
	Description string                                               `json:"description"`
	Error       string                                               `json:"error"`
	RuleName    string                                               `json:"rule_name"`
	Success     bool                                                 `json:"success"`
	Timestamp   string                                               `json:"timestamp"`
	Type        string                                               `json:"type"`
	JSON        accessUserActiveSessionGetResponseDevicePostureJSON  `json:"-"`
}

// accessUserActiveSessionGetResponseDevicePostureJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionGetResponseDevicePosture]
type accessUserActiveSessionGetResponseDevicePostureJSON struct {
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

func (r *AccessUserActiveSessionGetResponseDevicePosture) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseDevicePostureJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseDevicePostureCheck struct {
	Exists bool                                                     `json:"exists"`
	Path   string                                                   `json:"path"`
	JSON   accessUserActiveSessionGetResponseDevicePostureCheckJSON `json:"-"`
}

// accessUserActiveSessionGetResponseDevicePostureCheckJSON contains the JSON
// metadata for the struct [AccessUserActiveSessionGetResponseDevicePostureCheck]
type accessUserActiveSessionGetResponseDevicePostureCheckJSON struct {
	Exists      apijson.Field
	Path        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseDevicePostureCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseDevicePostureCheckJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseIdP struct {
	ID   string                                    `json:"id"`
	Type string                                    `json:"type"`
	JSON accessUserActiveSessionGetResponseIdPJSON `json:"-"`
}

// accessUserActiveSessionGetResponseIdPJSON contains the JSON metadata for the
// struct [AccessUserActiveSessionGetResponseIdP]
type accessUserActiveSessionGetResponseIdPJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseIdP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseIdPJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseMTLSAuth struct {
	AuthStatus    string                                         `json:"auth_status"`
	CERTIssuerDn  string                                         `json:"cert_issuer_dn"`
	CERTIssuerSki string                                         `json:"cert_issuer_ski"`
	CERTPresented bool                                           `json:"cert_presented"`
	CERTSerial    string                                         `json:"cert_serial"`
	JSON          accessUserActiveSessionGetResponseMTLSAuthJSON `json:"-"`
}

// accessUserActiveSessionGetResponseMTLSAuthJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionGetResponseMTLSAuth]
type accessUserActiveSessionGetResponseMTLSAuthJSON struct {
	AuthStatus    apijson.Field
	CERTIssuerDn  apijson.Field
	CERTIssuerSki apijson.Field
	CERTPresented apijson.Field
	CERTSerial    apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseMTLSAuth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseMTLSAuthJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessUserActiveSessionGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type AccessUserActiveSessionGetResponseEnvelope struct {
	Errors   []AccessUserActiveSessionGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessUserActiveSessionGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success AccessUserActiveSessionGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AccessUserActiveSessionGetResponse                `json:"result"`
	JSON    accessUserActiveSessionGetResponseEnvelopeJSON    `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [AccessUserActiveSessionGetResponseEnvelope]
type accessUserActiveSessionGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseEnvelopeErrors struct {
	Code             int64                                                  `json:"code,required"`
	Message          string                                                 `json:"message,required"`
	DocumentationURL string                                                 `json:"documentation_url"`
	Source           AccessUserActiveSessionGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             accessUserActiveSessionGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessUserActiveSessionGetResponseEnvelopeErrors]
type accessUserActiveSessionGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                     `json:"pointer"`
	JSON    accessUserActiveSessionGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [AccessUserActiveSessionGetResponseEnvelopeErrorsSource]
type accessUserActiveSessionGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseEnvelopeMessages struct {
	Code             int64                                                    `json:"code,required"`
	Message          string                                                   `json:"message,required"`
	DocumentationURL string                                                   `json:"documentation_url"`
	Source           AccessUserActiveSessionGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             accessUserActiveSessionGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [AccessUserActiveSessionGetResponseEnvelopeMessages]
type accessUserActiveSessionGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type AccessUserActiveSessionGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                       `json:"pointer"`
	JSON    accessUserActiveSessionGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// accessUserActiveSessionGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [AccessUserActiveSessionGetResponseEnvelopeMessagesSource]
type accessUserActiveSessionGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessUserActiveSessionGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserActiveSessionGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type AccessUserActiveSessionGetResponseEnvelopeSuccess bool

const (
	AccessUserActiveSessionGetResponseEnvelopeSuccessTrue AccessUserActiveSessionGetResponseEnvelopeSuccess = true
)

func (r AccessUserActiveSessionGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessUserActiveSessionGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
