package cloudflare

import (
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
	}

	return accessIdentityProviderResponse.Result, nil
}
