package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

type AccessUserActiveSessionsResponse struct {
	Result     []AccessUserActiveSessionResult `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

type AccessUserActiveSessionResult struct {
	Expiration int64                           `json:"expiration"`
	Metadata   AccessUserActiveSessionMetadata `json:"metadata"`
	Name       string                          `json:"name"`
}

type AccessUserActiveSessionMetadata struct {
	Apps    map[string]AccessUserActiveSessionMetadataApp `json:"apps"`
	Expires int64                                         `json:"expires"`
	IAT     int64                                         `json:"iat"`
	Nonce   string                                        `json:"nonce"`
	TTL     int64                                         `json:"ttl"`
}

type AccessUserActiveSessionMetadataApp struct {
	Hostname string `json:"hostname"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	UID      string `json:"uid"`
}

type AccessUserDevicePosture struct {
	Check       AccessUserDevicePostureCheck `json:"check"`
	Data        map[string]interface{}       `json:"data"`
	Description string                       `json:"description"`
	Error       string                       `json:"error"`
	ID          string                       `json:"id"`
	RuleName    string                       `json:"rule_name"`
	Success     *bool                        `json:"success"`
	Timestamp   string                       `json:"timestamp"`
	Type        string                       `json:"type"`
}

type AccessUserDeviceSession struct {
	LastAuthenticated int64 `json:"last_authenticated"`
}

type AccessUserFailedLoginsResponse struct {
	Result     []AccessUserFailedLoginResult `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

type AccessUserFailedLoginResult struct {
	Expiration int64                         `json:"expiration"`
	Metadata   AccessUserFailedLoginMetadata `json:"metadata"`
}

type AccessUserFailedLoginMetadata struct {
	AppName   string `json:"app_name"`
	Aud       string `json:"aud"`
	Datetime  string `json:"datetime"`
	RayID     string `json:"ray_id"`
	UserEmail string `json:"user_email"`
	UserUUID  string `json:"user_uuid"`
}

type AccessUserLastSeenIdentityResponse struct {
	Result     AccessUserLastSeenIdentityResult `json:"result"`
	ResultInfo ResultInfo                       `json:"result_info"`
	Response
}

type AccessUserLastSeenIdentityResult struct {
	AccountID          string                             `json:"account_id"`
	AuthStatus         string                             `json:"auth_status"`
	CommonName         string                             `json:"common_name"`
	DeviceID           string                             `json:"device_id"`
	DevicePosture      map[string]AccessUserDevicePosture `json:"devicePosture"`
	DeviceSessions     map[string]AccessUserDeviceSession `json:"device_sessions"`
	Email              string                             `json:"email"`
	Geo                AccessUserIdentityGeo              `json:"geo"`
	IAT                int64                              `json:"iat"`
	IDP                AccessUserIDP                      `json:"idp"`
	IP                 string                             `json:"ip"`
	IsGateway          *bool                              `json:"is_gateway"`
	IsWarp             *bool                              `json:"is_warp"`
	MtlsAuth           AccessUserMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                             `json:"service_token_id"`
	ServiceTokenStatus *bool                              `json:"service_token_status"`
	UserUUID           string                             `json:"user_uuid"`
	Version            int                                `json:"version"`
}

type AccessUserLastSeenIdentitySessionResponse struct {
	Result     GetAccessUserLastSeenIdentityResult `json:"result"`
	ResultInfo ResultInfo                          `json:"result_info"`
	Response
}

type GetAccessUserLastSeenIdentityResult struct {
	AccountID          string                             `json:"account_id"`
	AuthStatus         string                             `json:"auth_status"`
	CommonName         string                             `json:"common_name"`
	DevicePosture      map[string]AccessUserDevicePosture `json:"devicePosture"`
	DeviceSessions     map[string]AccessUserDeviceSession `json:"device_sessions"`
	DeviceID           string                             `json:"device_id"`
	Email              string                             `json:"email"`
	Geo                AccessUserIdentityGeo              `json:"geo"`
	IAT                int64                              `json:"iat"`
	IDP                AccessUserIDP                      `json:"idp"`
	IP                 string                             `json:"ip"`
	IsGateway          *bool                              `json:"is_gateway"`
	IsWarp             *bool                              `json:"is_warp"`
	MtlsAuth           AccessUserMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                             `json:"service_token_id"`
	ServiceTokenStatus *bool                              `json:"service_token_status"`
	UserUUID           string                             `json:"user_uuid"`
	Version            int                                `json:"version"`
}

type AccessUserDevicePostureCheck struct {
	Exists *bool  `json:"exists"`
	Path   string `json:"path"`
}

type AccessUserIdentityGeo struct {
	Country string `json:"country"`
}

type AccessUserIDP struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type AccessUserMTLSAuth struct {
	AuthStatus    string `json:"auth_status"`
	CertIssuerDN  string `json:"cert_issuer_dn"`
	CertIssuerSKI string `json:"cert_issuer_ski"`
	CertPresented *bool  `json:"cert_presented"`
	CertSerial    string `json:"cert_serial"`
}

type AccessUserListResponse struct {
	Result     []AccessUser `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

type AccessUser struct {
	ID                  string `json:"id"`
	AccessSeat          *bool  `json:"access_seat"`
	ActiveDeviceCount   int    `json:"active_device_count"`
	CreatedAt           string `json:"created_at"`
	Email               string `json:"email"`
	GatewaySeat         *bool  `json:"gateway_seat"`
	LastSuccessfulLogin string `json:"last_successful_login"`
	Name                string `json:"name"`
	SeatUID             string `json:"seat_uid"`
	UID                 string `json:"uid"`
	UpdatedAt           string `json:"updated_at"`
}

type AccessUserParams struct {
	ResultInfo
}

type GetAccessUserSingleActiveSessionResponse struct {
	Result     GetAccessUserSingleActiveSessionResult `json:"result"`
	ResultInfo `json:"result_info"`
	Response
}

type GetAccessUserSingleActiveSessionResult struct {
	AccountID          string                             `json:"account_id"`
	AuthStatus         string                             `json:"auth_status"`
	CommonName         string                             `json:"common_name"`
	DevicePosture      map[string]AccessUserDevicePosture `json:"devicePosture"`
	DeviceSessions     map[string]AccessUserDeviceSession `json:"device_sessions"`
	DeviceID           string                             `json:"device_id"`
	Email              string                             `json:"email"`
	Geo                AccessUserIdentityGeo              `json:"geo"`
	IAT                int64                              `json:"iat"`
	IDP                AccessUserIDP                      `json:"idp"`
	IP                 string                             `json:"ip"`
	IsGateway          *bool                              `json:"is_gateway"`
	IsWarp             *bool                              `json:"is_warp"`
	MtlsAuth           AccessUserMTLSAuth                 `json:"mtls_auth"`
	ServiceTokenID     string                             `json:"service_token_id"`
	ServiceTokenStatus *bool                              `json:"service_token_status"`
	UserUUID           string                             `json:"user_uuid"`
	Version            int                                `json:"version"`
	IsActive           *bool                              `json:"isActive"`
}

// ListAccessUsers returns a list of users for a single cloudflare access/zerotrust account.
//
// API documentation: https://developers.cloudflare.com/api/operations/zero-trust-users-get-users
func (api *API) ListAccessUsers(ctx context.Context, rc *ResourceContainer, params AccessUserParams) ([]AccessUser, *ResultInfo, error) {
	if rc.Level != AccountRouteLevel {
		return []AccessUser{}, &ResultInfo{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	baseURL := fmt.Sprintf("/%s/%s/access/users", rc.Level, rc.Identifier)

	autoPaginate := true
	if params.PerPage >= 1 || params.Page >= 1 {
		autoPaginate = false
	}

	if params.PerPage < 1 {
		params.PerPage = 25
	}

	if params.Page < 1 {
		params.Page = 1
	}

	var accessUsers []AccessUser
	var resultInfo *ResultInfo = nil

	for {
		uri := buildURI(baseURL, params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []AccessUser{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}
		var r AccessUserListResponse
		resultInfo = &r.ResultInfo

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []AccessUser{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		accessUsers = append(accessUsers, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return accessUsers, resultInfo, nil
}

// GetAccessUserActiveSessions returns a list of active sessions for an user.
//
// API documentation: https://developers.cloudflare.com/api/operations/zero-trust-users-get-active-sessions
func (api *API) GetAccessUserActiveSessions(ctx context.Context, rc *ResourceContainer, userID string) ([]AccessUserActiveSessionResult, error) {
	if rc.Level != AccountRouteLevel {
		return []AccessUserActiveSessionResult{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/active_sessions",
		rc.Level,
		rc.Identifier,
		userID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessUserActiveSessionResult{}, err
	}

	var accessUserActiveSessionsResponse AccessUserActiveSessionsResponse
	err = json.Unmarshal(res, &accessUserActiveSessionsResponse)
	if err != nil {
		return []AccessUserActiveSessionResult{}, err
	}
	return accessUserActiveSessionsResponse.Result, nil
}

// GetAccessUserSingleActiveSession returns a single active session for a user.
//
// API documentation: https://developers.cloudflare.com/api/operations/zero-trust-users-get-active-session
func (api *API) GetAccessUserSingleActiveSession(ctx context.Context, rc *ResourceContainer, userID string, sessionID string) (GetAccessUserSingleActiveSessionResult, error) {
	if rc.Level != AccountRouteLevel {
		return GetAccessUserSingleActiveSessionResult{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/active_sessions/%s",
		rc.Level,
		rc.Identifier,
		userID,
		sessionID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return GetAccessUserSingleActiveSessionResult{}, err
	}

	var accessUserActiveSingleSessionsResponse GetAccessUserSingleActiveSessionResponse
	err = json.Unmarshal(res, &accessUserActiveSingleSessionsResponse)
	if err != nil {
		return GetAccessUserSingleActiveSessionResult{}, err
	}
	return accessUserActiveSingleSessionsResponse.Result, nil
}

// GetAccessUserFailedLogins returns a list of failed logins for a user.
//
// API documentation: https://developers.cloudflare.com/api/operations/zero-trust-users-get-failed-logins
func (api *API) GetAccessUserFailedLogins(ctx context.Context, rc *ResourceContainer, userID string) ([]AccessUserFailedLoginResult, error) {
	if rc.Level != AccountRouteLevel {
		return []AccessUserFailedLoginResult{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/failed_logins",
		rc.Level,
		rc.Identifier,
		userID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessUserFailedLoginResult{}, err
	}

	var accessUserFailedLoginsResponse AccessUserFailedLoginsResponse
	err = json.Unmarshal(res, &accessUserFailedLoginsResponse)
	if err != nil {
		return []AccessUserFailedLoginResult{}, err
	}
	return accessUserFailedLoginsResponse.Result, nil
}

// GetAccessUserLastSeenIdentity returns the last seen identity for a user.
//
// API documentation: https://developers.cloudflare.com/api/operations/zero-trust-users-get-last-seen-identity
func (api *API) GetAccessUserLastSeenIdentity(ctx context.Context, rc *ResourceContainer, userID string) (GetAccessUserLastSeenIdentityResult, error) {
	if rc.Level != AccountRouteLevel {
		return GetAccessUserLastSeenIdentityResult{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/users/%s/last_seen_identity",
		rc.Level,
		rc.Identifier,
		userID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return GetAccessUserLastSeenIdentityResult{}, err
	}

	var accessUserLastSeenIdentityResponse AccessUserLastSeenIdentitySessionResponse
	err = json.Unmarshal(res, &accessUserLastSeenIdentityResponse)
	if err != nil {
		return GetAccessUserLastSeenIdentityResult{}, err
	}
	return accessUserLastSeenIdentityResponse.Result, nil
}
