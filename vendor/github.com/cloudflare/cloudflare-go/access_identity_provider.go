package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AccessIdentityProvider is the structure of the provider object.
type AccessIdentityProvider struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	ID     string                              `json:"id,omitemtpy"`
	Name   string                              `json:"name"`
	Type   string                              `json:"type"`
	Config AccessIdentityProviderConfiguration `json:"config"`
}

// AccessIdentityProviderConfiguration is the combined structure of *all*
// identity provider configuration fields. This is done to simplify the use of
// Access products and their relationship to each other.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/
type AccessIdentityProviderConfiguration struct {
	AppsDomain         string   `json:"apps_domain,omitempty"`
	Attributes         []string `json:"attributes,omitempty"`
	AuthURL            string   `json:"auth_url,omitempty"`
	CentrifyAccount    string   `json:"centrify_account,omitempty"`
	CentrifyAppID      string   `json:"centrify_app_id,omitempty"`
	CertsURL           string   `json:"certs_url,omitempty"`
	ClientID           string   `json:"client_id,omitempty"`
	ClientSecret       string   `json:"client_secret,omitempty"`
	DirectoryID        string   `json:"directory_id,omitempty"`
	EmailAttributeName string   `json:"email_attribute_name,omitempty"`
	IdpPublicCert      string   `json:"idp_public_cert,omitempty"`
	IssuerURL          string   `json:"issuer_url,omitempty"`
	OktaAccount        string   `json:"okta_account,omitempty"`
	OneloginAccount    string   `json:"onelogin_account,omitempty"`
	RedirectURL        string   `json:"redirect_url,omitempty"`
	SignRequest        bool     `json:"sign_request,omitempty"`
	SsoTargetURL       string   `json:"sso_target_url,omitempty"`
	SupportGroups      bool     `json:"support_groups,omitempty"`
	TokenURL           string   `json:"token_url,omitempty"`
}

// AccessIdentityProvidersListResponse is the API response for multiple
// Access Identity Providers.
type AccessIdentityProvidersListResponse struct {
	Response
	Result   []AccessIdentityProvider         `json:"result"`
}

// AccessIdentityProviderListResponse is the API response for a single
// Access Identity Provider.
type AccessIdentityProviderListResponse struct {
	Response
	Result   AccessIdentityProvider           `json:"result"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	ID     string      `json:"id,omitemtpy"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
||||||| parent of 5ce8c7613 (update vendored files)
	ID     string      `json:"id,omitemtpy"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
=======
	ID     string                              `json:"id,omitemtpy"`
	Name   string                              `json:"name"`
	Type   string                              `json:"type"`
	Config AccessIdentityProviderConfiguration `json:"config"`
>>>>>>> 5ce8c7613 (update vendored files)
}

// AccessIdentityProviderConfiguration is the combined structure of *all*
// identity provider configuration fields. This is done to simplify the use of
// Access products and their relationship to each other.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/
type AccessIdentityProviderConfiguration struct {
	AppsDomain         string   `json:"apps_domain,omitempty"`
	Attributes         []string `json:"attributes,omitempty"`
	AuthURL            string   `json:"auth_url,omitempty"`
	CentrifyAccount    string   `json:"centrify_account,omitempty"`
	CentrifyAppID      string   `json:"centrify_app_id,omitempty"`
	CertsURL           string   `json:"certs_url,omitempty"`
	ClientID           string   `json:"client_id,omitempty"`
	ClientSecret       string   `json:"client_secret,omitempty"`
	DirectoryID        string   `json:"directory_id,omitempty"`
	EmailAttributeName string   `json:"email_attribute_name,omitempty"`
	IdpPublicCert      string   `json:"idp_public_cert,omitempty"`
	IssuerURL          string   `json:"issuer_url,omitempty"`
	OktaAccount        string   `json:"okta_account,omitempty"`
	OneloginAccount    string   `json:"onelogin_account,omitempty"`
	RedirectURL        string   `json:"redirect_url,omitempty"`
	SignRequest        bool     `json:"sign_request,omitempty"`
	SsoTargetURL       string   `json:"sso_target_url,omitempty"`
	SupportGroups      bool     `json:"support_groups,omitempty"`
	TokenURL           string   `json:"token_url,omitempty"`
}

// AccessIdentityProvidersListResponse is the API response for multiple
// Access Identity Providers.
type AccessIdentityProvidersListResponse struct {
	Response
	Result   []AccessIdentityProvider         `json:"result"`
}

// AccessIdentityProviderListResponse is the API response for a single
// Access Identity Provider.
type AccessIdentityProviderListResponse struct {
<<<<<<< HEAD
	Success  bool                   `json:"success"`
	Errors   []string               `json:"errors"`
	Messages []string               `json:"messages"`
	Result   AccessIdentityProvider `json:"result"`
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	Success  bool                   `json:"success"`
	Errors   []string               `json:"errors"`
	Messages []string               `json:"messages"`
	Result   AccessIdentityProvider `json:"result"`
=======
	Response
	Result   AccessIdentityProvider           `json:"result"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	ID     string      `json:"id,omitemtpy"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
||||||| parent of 6b7ce455e (update vendored files)
	ID     string      `json:"id,omitemtpy"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
=======
	ID     string                              `json:"id,omitempty"`
	Name   string                              `json:"name"`
	Type   string                              `json:"type"`
	Config AccessIdentityProviderConfiguration `json:"config"`
>>>>>>> 6b7ce455e (update vendored files)
}

// AccessIdentityProviderConfiguration is the combined structure of *all*
// identity provider configuration fields. This is done to simplify the use of
// Access products and their relationship to each other.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/
type AccessIdentityProviderConfiguration struct {
	APIToken           string   `json:"api_token,omitempty"`
	AppsDomain         string   `json:"apps_domain,omitempty"`
	Attributes         []string `json:"attributes,omitempty"`
	AuthURL            string   `json:"auth_url,omitempty"`
	CentrifyAccount    string   `json:"centrify_account,omitempty"`
	CentrifyAppID      string   `json:"centrify_app_id,omitempty"`
	CertsURL           string   `json:"certs_url,omitempty"`
	ClientID           string   `json:"client_id,omitempty"`
	ClientSecret       string   `json:"client_secret,omitempty"`
	DirectoryID        string   `json:"directory_id,omitempty"`
	EmailAttributeName string   `json:"email_attribute_name,omitempty"`
	IdpPublicCert      string   `json:"idp_public_cert,omitempty"`
	IssuerURL          string   `json:"issuer_url,omitempty"`
	OktaAccount        string   `json:"okta_account,omitempty"`
	OneloginAccount    string   `json:"onelogin_account,omitempty"`
	RedirectURL        string   `json:"redirect_url,omitempty"`
	SignRequest        bool     `json:"sign_request,omitempty"`
	SsoTargetURL       string   `json:"sso_target_url,omitempty"`
	SupportGroups      bool     `json:"support_groups,omitempty"`
	TokenURL           string   `json:"token_url,omitempty"`
}

// AccessIdentityProvidersListResponse is the API response for multiple
// Access Identity Providers.
type AccessIdentityProvidersListResponse struct {
	Response
	Result []AccessIdentityProvider `json:"result"`
}

// AccessIdentityProviderListResponse is the API response for a single
// Access Identity Provider.
type AccessIdentityProviderListResponse struct {
<<<<<<< HEAD
	Success  bool                   `json:"success"`
	Errors   []string               `json:"errors"`
	Messages []string               `json:"messages"`
	Result   AccessIdentityProvider `json:"result"`
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	Success  bool                   `json:"success"`
	Errors   []string               `json:"errors"`
	Messages []string               `json:"messages"`
	Result   AccessIdentityProvider `json:"result"`
=======
	Response
	Result AccessIdentityProvider `json:"result"`
>>>>>>> 6b7ce455e (update vendored files)
}

// AccessIdentityProviders returns all Access Identity Providers for an
// account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-list-access-identity-providers
func (api *API) AccessIdentityProviders(ctx context.Context, accountID string) ([]AccessIdentityProvider, error) {
	return api.accessIdentityProviders(ctx, accountID, AccountRouteRoot)
}

// ZoneLevelAccessIdentityProviders returns all Access Identity Providers for an
// account.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-list-access-identity-providers
func (api *API) ZoneLevelAccessIdentityProviders(ctx context.Context, zoneID string) ([]AccessIdentityProvider, error) {
	return api.accessIdentityProviders(ctx, zoneID, ZoneRouteRoot)
}

func (api *API) accessIdentityProviders(ctx context.Context, id string, routeRoot RouteRoot) ([]AccessIdentityProvider, error) {
	uri := fmt.Sprintf("/%s/%s/access/identity_providers", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProvidersListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return []AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// AccessIdentityProviderDetails returns a single Access Identity
// Provider for an account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-access-identity-providers-details
func (api *API) AccessIdentityProviderDetails(ctx context.Context, accountID, identityProviderID string) (AccessIdentityProvider, error) {
	return api.accessIdentityProviderDetails(ctx, accountID, identityProviderID, AccountRouteRoot)
}

// ZoneLevelAccessIdentityProviderDetails returns a single zone level Access Identity
// Provider for an account.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-access-identity-providers-details
func (api *API) ZoneLevelAccessIdentityProviderDetails(ctx context.Context, zoneID, identityProviderID string) (AccessIdentityProvider, error) {
	return api.accessIdentityProviderDetails(ctx, zoneID, identityProviderID, ZoneRouteRoot)
}

func (api *API) accessIdentityProviderDetails(ctx context.Context, id string, identityProviderID string, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/identity_providers/%s",
		routeRoot,
		id,
		identityProviderID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// CreateAccessIdentityProvider creates a new Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) CreateAccessIdentityProvider(ctx context.Context, accountID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.createAccessIdentityProvider(ctx, accountID, identityProviderConfiguration, AccountRouteRoot)
}

// CreateZoneLevelAccessIdentityProvider creates a new zone level Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-create-access-identity-provider
func (api *API) CreateZoneLevelAccessIdentityProvider(ctx context.Context, zoneID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.createAccessIdentityProvider(ctx, zoneID, identityProviderConfiguration, ZoneRouteRoot)
}

func (api *API) createAccessIdentityProvider(ctx context.Context, id string, identityProviderConfiguration AccessIdentityProvider, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf("/%s/%s/access/identity_providers", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// UpdateAccessIdentityProvider updates an existing Access Identity
// Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) UpdateAccessIdentityProvider(ctx context.Context, accountID, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.updateAccessIdentityProvider(ctx, accountID, identityProviderUUID, identityProviderConfiguration, AccountRouteRoot)
}

// UpdateZoneLevelAccessIdentityProvider updates an existing zone level Access Identity
// Provider.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-update-access-identity-provider
func (api *API) UpdateZoneLevelAccessIdentityProvider(ctx context.Context, zoneID, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.updateAccessIdentityProvider(ctx, zoneID, identityProviderUUID, identityProviderConfiguration, ZoneRouteRoot)
}

func (api *API) updateAccessIdentityProvider(ctx context.Context, id string, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/identity_providers/%s",
		routeRoot,
		id,
		identityProviderUUID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// DeleteAccessIdentityProvider deletes an Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) DeleteAccessIdentityProvider(ctx context.Context, accountID, identityProviderUUID string) (AccessIdentityProvider, error) {
	return api.deleteAccessIdentityProvider(ctx, accountID, identityProviderUUID, AccountRouteRoot)
}

// DeleteZoneLevelAccessIdentityProvider deletes a zone level Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-delete-access-identity-provider
func (api *API) DeleteZoneLevelAccessIdentityProvider(ctx context.Context, zoneID, identityProviderUUID string) (AccessIdentityProvider, error) {
	return api.deleteAccessIdentityProvider(ctx, zoneID, identityProviderUUID, ZoneRouteRoot)
}

func (api *API) deleteAccessIdentityProvider(ctx context.Context, id string, identityProviderUUID string, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/identity_providers/%s",
		routeRoot,
		id,
		identityProviderUUID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return AccessIdentityProvider{}, err
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
>>>>>>> 4d7e5ad26 (update vendored files)
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// AccessIdentityProvider is the structure of the provider object.
type AccessIdentityProvider struct {
	ID     string                              `json:"id,omitempty"`
	Name   string                              `json:"name"`
	Type   string                              `json:"type"`
	Config AccessIdentityProviderConfiguration `json:"config"`
}

// AccessIdentityProviderConfiguration is the combined structure of *all*
// identity provider configuration fields. This is done to simplify the use of
// Access products and their relationship to each other.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/
type AccessIdentityProviderConfiguration struct {
	APIToken           string   `json:"api_token,omitempty"`
	AppsDomain         string   `json:"apps_domain,omitempty"`
	Attributes         []string `json:"attributes,omitempty"`
	AuthURL            string   `json:"auth_url,omitempty"`
	CentrifyAccount    string   `json:"centrify_account,omitempty"`
	CentrifyAppID      string   `json:"centrify_app_id,omitempty"`
	CertsURL           string   `json:"certs_url,omitempty"`
	ClientID           string   `json:"client_id,omitempty"`
	ClientSecret       string   `json:"client_secret,omitempty"`
	DirectoryID        string   `json:"directory_id,omitempty"`
	EmailAttributeName string   `json:"email_attribute_name,omitempty"`
	IdpPublicCert      string   `json:"idp_public_cert,omitempty"`
	IssuerURL          string   `json:"issuer_url,omitempty"`
	OktaAccount        string   `json:"okta_account,omitempty"`
	OneloginAccount    string   `json:"onelogin_account,omitempty"`
	RedirectURL        string   `json:"redirect_url,omitempty"`
	SignRequest        bool     `json:"sign_request,omitempty"`
	SsoTargetURL       string   `json:"sso_target_url,omitempty"`
	SupportGroups      bool     `json:"support_groups,omitempty"`
	TokenURL           string   `json:"token_url,omitempty"`
	PKCEEnabled        *bool    `json:"pkce_enabled,omitempty"`
}

// AccessIdentityProvidersListResponse is the API response for multiple
// Access Identity Providers.
type AccessIdentityProvidersListResponse struct {
	Response
	Result []AccessIdentityProvider `json:"result"`
}

// AccessIdentityProviderListResponse is the API response for a single
// Access Identity Provider.
type AccessIdentityProviderListResponse struct {
	Response
	Result AccessIdentityProvider `json:"result"`
}

// AccessIdentityProviders returns all Access Identity Providers for an
// account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-list-access-identity-providers
func (api *API) AccessIdentityProviders(ctx context.Context, accountID string) ([]AccessIdentityProvider, error) {
	return api.accessIdentityProviders(ctx, accountID, AccountRouteRoot)
}

// ZoneLevelAccessIdentityProviders returns all Access Identity Providers for an
// account.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-list-access-identity-providers
func (api *API) ZoneLevelAccessIdentityProviders(ctx context.Context, zoneID string) ([]AccessIdentityProvider, error) {
	return api.accessIdentityProviders(ctx, zoneID, ZoneRouteRoot)
}

func (api *API) accessIdentityProviders(ctx context.Context, id string, routeRoot RouteRoot) ([]AccessIdentityProvider, error) {
	uri := fmt.Sprintf("/%s/%s/access/identity_providers", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProvidersListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return []AccessIdentityProvider{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessIdentityProviderResponse.Result, nil
}

// AccessIdentityProviderDetails returns a single Access Identity
// Provider for an account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-access-identity-providers-details
func (api *API) AccessIdentityProviderDetails(ctx context.Context, accountID, identityProviderID string) (AccessIdentityProvider, error) {
	return api.accessIdentityProviderDetails(ctx, accountID, identityProviderID, AccountRouteRoot)
}

// ZoneLevelAccessIdentityProviderDetails returns a single zone level Access Identity
// Provider for an account.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-access-identity-providers-details
func (api *API) ZoneLevelAccessIdentityProviderDetails(ctx context.Context, zoneID, identityProviderID string) (AccessIdentityProvider, error) {
	return api.accessIdentityProviderDetails(ctx, zoneID, identityProviderID, ZoneRouteRoot)
}

func (api *API) accessIdentityProviderDetails(ctx context.Context, id string, identityProviderID string, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/identity_providers/%s",
		routeRoot,
		id,
		identityProviderID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessIdentityProviderResponse.Result, nil
}

// CreateAccessIdentityProvider creates a new Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) CreateAccessIdentityProvider(ctx context.Context, accountID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.createAccessIdentityProvider(ctx, accountID, identityProviderConfiguration, AccountRouteRoot)
}

// CreateZoneLevelAccessIdentityProvider creates a new zone level Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-create-access-identity-provider
func (api *API) CreateZoneLevelAccessIdentityProvider(ctx context.Context, zoneID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.createAccessIdentityProvider(ctx, zoneID, identityProviderConfiguration, ZoneRouteRoot)
}

func (api *API) createAccessIdentityProvider(ctx context.Context, id string, identityProviderConfiguration AccessIdentityProvider, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf("/%s/%s/access/identity_providers", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessIdentityProviderResponse.Result, nil
}

// UpdateAccessIdentityProvider updates an existing Access Identity
// Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) UpdateAccessIdentityProvider(ctx context.Context, accountID, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.updateAccessIdentityProvider(ctx, accountID, identityProviderUUID, identityProviderConfiguration, AccountRouteRoot)
}

// UpdateZoneLevelAccessIdentityProvider updates an existing zone level Access Identity
// Provider.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-update-access-identity-provider
func (api *API) UpdateZoneLevelAccessIdentityProvider(ctx context.Context, zoneID, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	return api.updateAccessIdentityProvider(ctx, zoneID, identityProviderUUID, identityProviderConfiguration, ZoneRouteRoot)
}

func (api *API) updateAccessIdentityProvider(ctx context.Context, id string, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/identity_providers/%s",
		routeRoot,
		id,
		identityProviderUUID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, err
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessIdentityProviderResponse.Result, nil
}

// DeleteAccessIdentityProvider deletes an Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) DeleteAccessIdentityProvider(ctx context.Context, accountID, identityProviderUUID string) (AccessIdentityProvider, error) {
	return api.deleteAccessIdentityProvider(ctx, accountID, identityProviderUUID, AccountRouteRoot)
}

// DeleteZoneLevelAccessIdentityProvider deletes a zone level Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#zone-level-access-identity-providers-delete-access-identity-provider
func (api *API) DeleteZoneLevelAccessIdentityProvider(ctx context.Context, zoneID, identityProviderUUID string) (AccessIdentityProvider, error) {
	return api.deleteAccessIdentityProvider(ctx, zoneID, identityProviderUUID, ZoneRouteRoot)
}

func (api *API) deleteAccessIdentityProvider(ctx context.Context, id string, identityProviderUUID string, routeRoot RouteRoot) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/identity_providers/%s",
		routeRoot,
		id,
		identityProviderUUID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
<<<<<<< HEAD
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
=======
		return AccessIdentityProvider{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// AccessIdentityProvider is the structure of the provider object.
type AccessIdentityProvider struct {
	ID     string      `json:"id,omitemtpy"`
	Name   string      `json:"name"`
	Type   string      `json:"type"`
	Config interface{} `json:"config"`
}

// AccessAzureADConfiguration is the representation of the Azure AD identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/azuread/
type AccessAzureADConfiguration struct {
	ClientID      string `json:"client_id"`
	ClientSecret  string `json:"client_secret"`
	DirectoryID   string `json:"directory_id"`
	SupportGroups bool   `json:"support_groups"`
}

// AccessCentrifyConfiguration is the representation of the Centrify identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/centrify/
type AccessCentrifyConfiguration struct {
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	CentrifyAccount string `json:"centrify_account"`
	CentrifyAppID   string `json:"centrify_app_id"`
}

// AccessCentrifySAMLConfiguration is the representation of the Centrify
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/saml-centrify/
type AccessCentrifySAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// AccessFacebookConfiguration is the representation of the Facebook identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/facebook-login/
type AccessFacebookConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// AccessGSuiteConfiguration is the representation of the GSuite identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/gsuite/
type AccessGSuiteConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AppsDomain   string `json:"apps_domain"`
}

// AccessGenericOIDCConfiguration is the representation of the generic OpenID
// Connect (OIDC) connector.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/generic-oidc/
type AccessGenericOIDCConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AuthURL      string `json:"auth_url"`
	TokenURL     string `json:"token_url"`
	CertsURL     string `json:"certs_url"`
}

// AccessGitHubConfiguration is the representation of the GitHub identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/github/
type AccessGitHubConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// AccessGoogleConfiguration is the representation of the Google identity
// provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/google/
type AccessGoogleConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// AccessJumpCloudSAMLConfiguration is the representation of the Jump Cloud
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/jumpcloud-saml/
type AccessJumpCloudSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// AccessOktaConfiguration is the representation of the Okta identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/okta/
type AccessOktaConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	OktaAccount  string `json:"okta_account"`
}

// AccessOktaSAMLConfiguration is the representation of the Okta identity
// provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/saml-okta/
type AccessOktaSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// AccessOneTimePinConfiguration is the representation of the default One Time
// Pin identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/one-time-pin/
type AccessOneTimePinConfiguration struct{}

// AccessOneLoginOIDCConfiguration is the representation of the OneLogin
// OpenID connector as an identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/onelogin-oidc/
type AccessOneLoginOIDCConfiguration struct {
	ClientID        string `json:"client_id"`
	ClientSecret    string `json:"client_secret"`
	OneloginAccount string `json:"onelogin_account"`
}

// AccessOneLoginSAMLConfiguration is the representation of the OneLogin
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/onelogin-saml/
type AccessOneLoginSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// AccessPingSAMLConfiguration is the representation of the Ping identity
// provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/ping-saml/
type AccessPingSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// AccessYandexConfiguration is the representation of the Yandex identity provider.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/yandex/
type AccessYandexConfiguration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// AccessADSAMLConfiguration is the representation of the Active Directory
// identity provider using SAML.
//
// API reference: https://developers.cloudflare.com/access/configuring-identity-providers/adfs/
type AccessADSAMLConfiguration struct {
	IssuerURL          string   `json:"issuer_url"`
	SsoTargetURL       string   `json:"sso_target_url"`
	Attributes         []string `json:"attributes"`
	EmailAttributeName string   `json:"email_attribute_name"`
	SignRequest        bool     `json:"sign_request"`
	IdpPublicCert      string   `json:"idp_public_cert"`
}

// AccessIdentityProvidersListResponse is the API response for multiple
// Access Identity Providers.
type AccessIdentityProvidersListResponse struct {
	Success  bool                     `json:"success"`
	Errors   []string                 `json:"errors"`
	Messages []string                 `json:"messages"`
	Result   []AccessIdentityProvider `json:"result"`
}

// AccessIdentityProviderListResponse is the API response for a single
// Access Identity Provider.
type AccessIdentityProviderListResponse struct {
	Success  bool                   `json:"success"`
	Errors   []string               `json:"errors"`
	Messages []string               `json:"messages"`
	Result   AccessIdentityProvider `json:"result"`
}

// AccessIdentityProviders returns all Access Identity Providers for an
// account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-list-access-identity-providers
func (api *API) AccessIdentityProviders(accountID string) ([]AccessIdentityProvider, error) {
	uri := "/accounts/" + accountID + "/access/identity_providers"

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProvidersListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return []AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// AccessIdentityProviderDetails returns a single Access Identity
// Provider for an account.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-access-identity-providers-details
func (api *API) AccessIdentityProviderDetails(accountID, identityProviderID string) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/identity_providers/%s",
		accountID,
		identityProviderID,
	)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// CreateAccessIdentityProvider creates a new Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) CreateAccessIdentityProvider(accountID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	uri := "/accounts/" + accountID + "/access/identity_providers"

	res, err := api.makeRequest("POST", uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// UpdateAccessIdentityProvider updates an existing Access Identity
// Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) UpdateAccessIdentityProvider(accountID, identityProviderUUID string, identityProviderConfiguration AccessIdentityProvider) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/identity_providers/%s",
		accountID,
		identityProviderUUID,
	)

	res, err := api.makeRequest("PUT", uri, identityProviderConfiguration)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessIdentityProviderResponse.Result, nil
}

// DeleteAccessIdentityProvider deletes an Access Identity Provider.
//
// API reference: https://api.cloudflare.com/#access-identity-providers-create-access-identity-provider
func (api *API) DeleteAccessIdentityProvider(accountID, identityProviderUUID string) (AccessIdentityProvider, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/identity_providers/%s",
		accountID,
		identityProviderUUID,
	)

	res, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessIdentityProviderResponse AccessIdentityProviderListResponse
	err = json.Unmarshal(res, &accessIdentityProviderResponse)
	if err != nil {
		return AccessIdentityProvider{}, errors.Wrap(err, errUnmarshalError)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	}

	return accessIdentityProviderResponse.Result, nil
}
