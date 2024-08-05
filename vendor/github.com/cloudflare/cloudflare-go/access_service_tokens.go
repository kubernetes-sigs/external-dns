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
	"time"
)

// AccessServiceToken represents an Access Service Token.
type AccessServiceToken struct {
	ClientID  string     `json:"client_id"`
	CreatedAt *time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// AccessServiceTokenUpdateResponse represents the response from the API
// when a new Service Token is updated. This base struct is also used in the
// Create as they are very similar responses.
type AccessServiceTokenUpdateResponse struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	ClientID  string     `json:"client_id"`
}

// AccessServiceTokenRefreshResponse represents the response from the API
// when an existing service token is refreshed to last longer.
type AccessServiceTokenRefreshResponse struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	ClientID  string     `json:"client_id"`
}

// AccessServiceTokenCreateResponse is the same API response as the Update
// operation with the exception that the `ClientSecret` is present in a
// Create operation.
type AccessServiceTokenCreateResponse struct {
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	ExpiresAt    *time.Time `json:"expires_at"`
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	ClientID     string     `json:"client_id"`
	ClientSecret string     `json:"client_secret"`
}

// AccessServiceTokensListResponse represents the response from the list
// Access Service Tokens endpoint.
type AccessServiceTokensListResponse struct {
	Result []AccessServiceToken `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccessServiceTokensDetailResponse is the API response, containing a single
// Access Service Token.
type AccessServiceTokensDetailResponse struct {
	Success  bool               `json:"success"`
	Errors   []string           `json:"errors"`
	Messages []string           `json:"messages"`
	Result   AccessServiceToken `json:"result"`
}

// AccessServiceTokensCreationDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensCreationDetailResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   AccessServiceTokenCreateResponse `json:"result"`
}

// AccessServiceTokensUpdateDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensUpdateDetailResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   AccessServiceTokenUpdateResponse `json:"result"`
}

// AccessServiceTokensRefreshDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensRefreshDetailResponse struct {
	Success  bool                              `json:"success"`
	Errors   []string                          `json:"errors"`
	Messages []string                          `json:"messages"`
	Result   AccessServiceTokenRefreshResponse `json:"result"`
}

// AccessServiceTokens returns all Access Service Tokens for an account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-list-access-service-tokens
func (api *API) AccessServiceTokens(ctx context.Context, accountID string) ([]AccessServiceToken, ResultInfo, error) {
	return api.accessServiceTokens(ctx, accountID, AccountRouteRoot)
}

// ZoneLevelAccessServiceTokens returns all Access Service Tokens for a zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-list-access-service-tokens
func (api *API) ZoneLevelAccessServiceTokens(ctx context.Context, zoneID string) ([]AccessServiceToken, ResultInfo, error) {
	return api.accessServiceTokens(ctx, zoneID, ZoneRouteRoot)
}

func (api *API) accessServiceTokens(ctx context.Context, id string, routeRoot RouteRoot) ([]AccessServiceToken, ResultInfo, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessServiceToken{}, ResultInfo{}, err
	}

	var accessServiceTokensListResponse AccessServiceTokensListResponse
	err = json.Unmarshal(res, &accessServiceTokensListResponse)
	if err != nil {
		return []AccessServiceToken{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokensListResponse.Result, accessServiceTokensListResponse.ResultInfo, nil
}

// CreateAccessServiceToken creates a new Access Service Token for an account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-create-access-service-token
func (api *API) CreateAccessServiceToken(ctx context.Context, accountID, name string) (AccessServiceTokenCreateResponse, error) {
	return api.createAccessServiceToken(ctx, accountID, name, AccountRouteRoot)
}

// CreateZoneLevelAccessServiceToken creates a new Access Service Token for a zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-create-access-service-token
func (api *API) CreateZoneLevelAccessServiceToken(ctx context.Context, zoneID, name string) (AccessServiceTokenCreateResponse, error) {
	return api.createAccessServiceToken(ctx, zoneID, name, ZoneRouteRoot)
}

func (api *API) createAccessServiceToken(ctx context.Context, id, name string, routeRoot RouteRoot) (AccessServiceTokenCreateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens", routeRoot, id)
	marshalledName, _ := json.Marshal(struct {
		Name string `json:"name"`
	}{name})

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, marshalledName)

	if err != nil {
		return AccessServiceTokenCreateResponse{}, err
	}

	var accessServiceTokenCreation AccessServiceTokensCreationDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenCreation)
	if err != nil {
		return AccessServiceTokenCreateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenCreation.Result, nil
}

// UpdateAccessServiceToken updates an existing Access Service Token for an
// account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-update-access-service-token
func (api *API) UpdateAccessServiceToken(ctx context.Context, accountID, uuid, name string) (AccessServiceTokenUpdateResponse, error) {
	return api.updateAccessServiceToken(ctx, accountID, uuid, name, AccountRouteRoot)
}

// UpdateZoneLevelAccessServiceToken updates an existing Access Service Token for a
// zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-update-access-service-token
func (api *API) UpdateZoneLevelAccessServiceToken(ctx context.Context, zoneID, uuid, name string) (AccessServiceTokenUpdateResponse, error) {
	return api.updateAccessServiceToken(ctx, zoneID, uuid, name, ZoneRouteRoot)
}

func (api *API) updateAccessServiceToken(ctx context.Context, id, uuid, name string, routeRoot RouteRoot) (AccessServiceTokenUpdateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s", routeRoot, id, uuid)

	marshalledName, _ := json.Marshal(struct {
		Name string `json:"name"`
	}{name})

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, marshalledName)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, err
	}

	var accessServiceTokenUpdate AccessServiceTokensUpdateDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenUpdate)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenUpdate.Result, nil
}

// DeleteAccessServiceToken removes an existing Access Service Token for an
// account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-delete-access-service-token
func (api *API) DeleteAccessServiceToken(ctx context.Context, accountID, uuid string) (AccessServiceTokenUpdateResponse, error) {
	return api.deleteAccessServiceToken(ctx, accountID, uuid, AccountRouteRoot)
}

// DeleteZoneLevelAccessServiceToken removes an existing Access Service Token for a
// zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-delete-access-service-token
func (api *API) DeleteZoneLevelAccessServiceToken(ctx context.Context, zoneID, uuid string) (AccessServiceTokenUpdateResponse, error) {
	return api.deleteAccessServiceToken(ctx, zoneID, uuid, ZoneRouteRoot)
}

func (api *API) deleteAccessServiceToken(ctx context.Context, id, uuid string, routeRoot RouteRoot) (AccessServiceTokenUpdateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s", routeRoot, id, uuid)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// AccessServiceToken represents an Access Service Token.
type AccessServiceToken struct {
	ClientID  string     `json:"client_id"`
	CreatedAt *time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

// AccessServiceTokenUpdateResponse represents the response from the API
// when a new Service Token is updated. This base struct is also used in the
// Create as they are very similar responses.
type AccessServiceTokenUpdateResponse struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	ClientID  string     `json:"client_id"`
}

// AccessServiceTokenCreateResponse is the same API response as the Update
// operation with the exception that the `ClientSecret` is present in a
// Create operation.
type AccessServiceTokenCreateResponse struct {
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	ExpiresAt    *time.Time `json:"expires_at"`
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	ClientID     string     `json:"client_id"`
	ClientSecret string     `json:"client_secret"`
}

// AccessServiceTokensListResponse represents the response from the list
// Access Service Tokens endpoint.
type AccessServiceTokensListResponse struct {
	Result []AccessServiceToken `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccessServiceTokensDetailResponse is the API response, containing a single
// Access Service Token.
type AccessServiceTokensDetailResponse struct {
	Success  bool               `json:"success"`
	Errors   []string           `json:"errors"`
	Messages []string           `json:"messages"`
	Result   AccessServiceToken `json:"result"`
}

// AccessServiceTokensCreationDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensCreationDetailResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   AccessServiceTokenCreateResponse `json:"result"`
}

// AccessServiceTokensUpdateDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensUpdateDetailResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   AccessServiceTokenUpdateResponse `json:"result"`
}

// AccessServiceTokens returns all Access Service Tokens for an account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-list-access-service-tokens
func (api *API) AccessServiceTokens(ctx context.Context, accountID string) ([]AccessServiceToken, ResultInfo, error) {
	return api.accessServiceTokens(ctx, accountID, AccountRouteRoot)
}

// ZoneLevelAccessServiceTokens returns all Access Service Tokens for a zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-list-access-service-tokens
func (api *API) ZoneLevelAccessServiceTokens(ctx context.Context, zoneID string) ([]AccessServiceToken, ResultInfo, error) {
	return api.accessServiceTokens(ctx, zoneID, ZoneRouteRoot)
}

func (api *API) accessServiceTokens(ctx context.Context, id string, routeRoot RouteRoot) ([]AccessServiceToken, ResultInfo, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessServiceToken{}, ResultInfo{}, err
	}

	var accessServiceTokensListResponse AccessServiceTokensListResponse
	err = json.Unmarshal(res, &accessServiceTokensListResponse)
	if err != nil {
		return []AccessServiceToken{}, ResultInfo{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessServiceTokensListResponse.Result, accessServiceTokensListResponse.ResultInfo, nil
}

// CreateAccessServiceToken creates a new Access Service Token for an account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-create-access-service-token
func (api *API) CreateAccessServiceToken(ctx context.Context, accountID, name string) (AccessServiceTokenCreateResponse, error) {
	return api.createAccessServiceToken(ctx, accountID, name, AccountRouteRoot)
}

// CreateZoneLevelAccessServiceToken creates a new Access Service Token for a zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-create-access-service-token
func (api *API) CreateZoneLevelAccessServiceToken(ctx context.Context, zoneID, name string) (AccessServiceTokenCreateResponse, error) {
	return api.createAccessServiceToken(ctx, zoneID, name, ZoneRouteRoot)
}

func (api *API) createAccessServiceToken(ctx context.Context, id, name string, routeRoot RouteRoot) (AccessServiceTokenCreateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens", routeRoot, id)
	marshalledName, _ := json.Marshal(struct {
		Name string `json:"name"`
	}{name})

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, marshalledName)

	if err != nil {
		return AccessServiceTokenCreateResponse{}, err
	}

	var accessServiceTokenCreation AccessServiceTokensCreationDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenCreation)
	if err != nil {
		return AccessServiceTokenCreateResponse{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessServiceTokenCreation.Result, nil
}

// UpdateAccessServiceToken updates an existing Access Service Token for an
// account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-update-access-service-token
func (api *API) UpdateAccessServiceToken(ctx context.Context, accountID, uuid, name string) (AccessServiceTokenUpdateResponse, error) {
	return api.updateAccessServiceToken(ctx, accountID, uuid, name, AccountRouteRoot)
}

// UpdateZoneLevelAccessServiceToken updates an existing Access Service Token for a
// zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-update-access-service-token
func (api *API) UpdateZoneLevelAccessServiceToken(ctx context.Context, zoneID, uuid, name string) (AccessServiceTokenUpdateResponse, error) {
	return api.updateAccessServiceToken(ctx, zoneID, uuid, name, ZoneRouteRoot)
}

func (api *API) updateAccessServiceToken(ctx context.Context, id, uuid, name string, routeRoot RouteRoot) (AccessServiceTokenUpdateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s", routeRoot, id, uuid)

	marshalledName, _ := json.Marshal(struct {
		Name string `json:"name"`
	}{name})

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, marshalledName)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, err
	}

	var accessServiceTokenUpdate AccessServiceTokensUpdateDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenUpdate)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessServiceTokenUpdate.Result, nil
}

// DeleteAccessServiceToken removes an existing Access Service Token for an
// account.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-delete-access-service-token
func (api *API) DeleteAccessServiceToken(ctx context.Context, accountID, uuid string) (AccessServiceTokenUpdateResponse, error) {
	return api.deleteAccessServiceToken(ctx, accountID, uuid, AccountRouteRoot)
}

// DeleteZoneLevelAccessServiceToken removes an existing Access Service Token for a
// zone.
//
// API reference: https://api.cloudflare.com/#zone-level-access-service-tokens-delete-access-service-token
func (api *API) DeleteZoneLevelAccessServiceToken(ctx context.Context, zoneID, uuid string) (AccessServiceTokenUpdateResponse, error) {
	return api.deleteAccessServiceToken(ctx, zoneID, uuid, ZoneRouteRoot)
}

func (api *API) deleteAccessServiceToken(ctx context.Context, id, uuid string, routeRoot RouteRoot) (AccessServiceTokenUpdateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s", routeRoot, id, uuid)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return AccessServiceTokenUpdateResponse{}, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return AccessServiceTokenUpdateResponse{}, errors.Wrap(err, errMakeRequestError)
=======
		return AccessServiceTokenUpdateResponse{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
	}

	var accessServiceTokenUpdate AccessServiceTokensUpdateDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenUpdate)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenUpdate.Result, nil
}

// RefreshAccessServiceToken updates the expiry of an Access Service Token
// in place.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-refresh-a-service-token
func (api *API) RefreshAccessServiceToken(ctx context.Context, rc *ResourceContainer, id string) (AccessServiceTokenRefreshResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s/refresh", rc.Level, rc.Identifier, id)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return AccessServiceTokenRefreshResponse{}, err
	}

	var accessServiceTokenRefresh AccessServiceTokensRefreshDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenRefresh)
	if err != nil {
		return AccessServiceTokenRefreshResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenRefresh.Result, nil
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
	"time"

	"github.com/goccy/go-json"
)

var (
	ErrMissingServiceTokenUUID = errors.New("missing required service token UUID")
)

// AccessServiceToken represents an Access Service Token.
type AccessServiceToken struct {
	ClientID  string     `json:"client_id"`
	CreatedAt *time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
	Duration  string     `json:"duration,omitempty"`
}

// AccessServiceTokenUpdateResponse represents the response from the API
// when a new Service Token is updated. This base struct is also used in the
// Create as they are very similar responses.
type AccessServiceTokenUpdateResponse struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	ClientID  string     `json:"client_id"`
	Duration  string     `json:"duration,omitempty"`
}

// AccessServiceTokenRefreshResponse represents the response from the API
// when an existing service token is refreshed to last longer.
type AccessServiceTokenRefreshResponse struct {
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	ExpiresAt *time.Time `json:"expires_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	ClientID  string     `json:"client_id"`
	Duration  string     `json:"duration,omitempty"`
}

// AccessServiceTokenCreateResponse is the same API response as the Update
// operation with the exception that the `ClientSecret` is present in a
// Create operation.
type AccessServiceTokenCreateResponse struct {
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	ExpiresAt    *time.Time `json:"expires_at"`
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	ClientID     string     `json:"client_id"`
	ClientSecret string     `json:"client_secret"`
	Duration     string     `json:"duration,omitempty"`
}

// AccessServiceTokenRotateResponse is the same API response as the Create
// operation.
type AccessServiceTokenRotateResponse struct {
	CreatedAt    *time.Time `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
	ExpiresAt    *time.Time `json:"expires_at"`
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	ClientID     string     `json:"client_id"`
	ClientSecret string     `json:"client_secret"`
	Duration     string     `json:"duration,omitempty"`
}

// AccessServiceTokensListResponse represents the response from the list
// Access Service Tokens endpoint.
type AccessServiceTokensListResponse struct {
	Result []AccessServiceToken `json:"result"`
	Response
	ResultInfo `json:"result_info"`
}

// AccessServiceTokensDetailResponse is the API response, containing a single
// Access Service Token.
type AccessServiceTokensDetailResponse struct {
	Success  bool               `json:"success"`
	Errors   []string           `json:"errors"`
	Messages []string           `json:"messages"`
	Result   AccessServiceToken `json:"result"`
}

// AccessServiceTokensCreationDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensCreationDetailResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   AccessServiceTokenCreateResponse `json:"result"`
}

// AccessServiceTokensUpdateDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensUpdateDetailResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   AccessServiceTokenUpdateResponse `json:"result"`
}

// AccessServiceTokensRefreshDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensRefreshDetailResponse struct {
	Success  bool                              `json:"success"`
	Errors   []string                          `json:"errors"`
	Messages []string                          `json:"messages"`
	Result   AccessServiceTokenRefreshResponse `json:"result"`
}

// AccessServiceTokensRotateSecretDetailResponse is the API response, containing a
// single Access Service Token.
type AccessServiceTokensRotateSecretDetailResponse struct {
	Success  bool                             `json:"success"`
	Errors   []string                         `json:"errors"`
	Messages []string                         `json:"messages"`
	Result   AccessServiceTokenRotateResponse `json:"result"`
}

type ListAccessServiceTokensParams struct{}

type CreateAccessServiceTokenParams struct {
	Name     string `json:"name"`
	Duration string `json:"duration,omitempty"`
}

type UpdateAccessServiceTokenParams struct {
	Name     string `json:"name"`
	UUID     string `json:"-"`
	Duration string `json:"duration,omitempty"`
}

func (api *API) ListAccessServiceTokens(ctx context.Context, rc *ResourceContainer, params ListAccessServiceTokensParams) ([]AccessServiceToken, ResultInfo, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens", rc.Level, rc.Identifier)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessServiceToken{}, ResultInfo{}, err
	}

	var accessServiceTokensListResponse AccessServiceTokensListResponse
	err = json.Unmarshal(res, &accessServiceTokensListResponse)
	if err != nil {
		return []AccessServiceToken{}, ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokensListResponse.Result, accessServiceTokensListResponse.ResultInfo, nil
}

func (api *API) CreateAccessServiceToken(ctx context.Context, rc *ResourceContainer, params CreateAccessServiceTokenParams) (AccessServiceTokenCreateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens", rc.Level, rc.Identifier)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)

	if err != nil {
		return AccessServiceTokenCreateResponse{}, err
	}

	var accessServiceTokenCreation AccessServiceTokensCreationDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenCreation)
	if err != nil {
		return AccessServiceTokenCreateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenCreation.Result, nil
}

func (api *API) UpdateAccessServiceToken(ctx context.Context, rc *ResourceContainer, params UpdateAccessServiceTokenParams) (AccessServiceTokenUpdateResponse, error) {
	if params.UUID == "" {
		return AccessServiceTokenUpdateResponse{}, ErrMissingServiceTokenUUID
	}

	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s", rc.Level, rc.Identifier, params.UUID)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, err
	}

	var accessServiceTokenUpdate AccessServiceTokensUpdateDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenUpdate)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenUpdate.Result, nil
}

func (api *API) DeleteAccessServiceToken(ctx context.Context, rc *ResourceContainer, uuid string) (AccessServiceTokenUpdateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s", rc.Level, rc.Identifier, uuid)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, err
	}

	var accessServiceTokenUpdate AccessServiceTokensUpdateDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenUpdate)
	if err != nil {
		return AccessServiceTokenUpdateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenUpdate.Result, nil
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// RefreshAccessServiceToken updates the expiry of an Access Service Token
// in place.
//
// API reference: https://api.cloudflare.com/#access-service-tokens-refresh-a-service-token
func (api *API) RefreshAccessServiceToken(ctx context.Context, rc *ResourceContainer, id string) (AccessServiceTokenRefreshResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s/refresh", rc.Level, rc.Identifier, id)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return AccessServiceTokenRefreshResponse{}, err
	}

	var accessServiceTokenRefresh AccessServiceTokensRefreshDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenRefresh)
	if err != nil {
		return AccessServiceTokenRefreshResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenRefresh.Result, nil
}

// RotateAccessServiceToken rotates the client secret of an Access Service
// Token in place.
// API reference: https://api.cloudflare.com/#access-service-tokens-rotate-a-service-token
func (api *API) RotateAccessServiceToken(ctx context.Context, rc *ResourceContainer, id string) (AccessServiceTokenRotateResponse, error) {
	uri := fmt.Sprintf("/%s/%s/access/service_tokens/%s/rotate", rc.Level, rc.Identifier, id)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return AccessServiceTokenRotateResponse{}, err
	}

	var accessServiceTokenRotate AccessServiceTokensRotateSecretDetailResponse
	err = json.Unmarshal(res, &accessServiceTokenRotate)
	if err != nil {
		return AccessServiceTokenRotateResponse{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessServiceTokenRotate.Result, nil
}
