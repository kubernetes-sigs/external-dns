package cloudflare

import (
	"context"
<<<<<<< HEAD
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AccessMutualTLSCertificate is the structure of a single Access Mutual TLS
// certificate.
type AccessMutualTLSCertificate struct {
	ID                  string    `json:"id,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	ExpiresOn           time.Time `json:"expires_on,omitempty"`
	Name                string    `json:"name,omitempty"`
	Fingerprint         string    `json:"fingerprint,omitempty"`
	Certificate         string    `json:"certificate,omitempty"`
	AssociatedHostnames []string  `json:"associated_hostnames,omitempty"`
}

// AccessMutualTLSCertificateListResponse is the API response for all Access
// Mutual TLS certificates.
type AccessMutualTLSCertificateListResponse struct {
	Response
	Result []AccessMutualTLSCertificate `json:"result"`
}

// AccessMutualTLSCertificateDetailResponse is the API response for a single
// Access Mutual TLS certificate.
type AccessMutualTLSCertificateDetailResponse struct {
	Response
	Result AccessMutualTLSCertificate `json:"result"`
}

// AccessMutualTLSCertificates returns all Access TLS certificates for the account
// level.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-properties
func (api *API) AccessMutualTLSCertificates(ctx context.Context, accountID string) ([]AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificates(ctx, accountID, AccountRouteRoot)
}

// ZoneAccessMutualTLSCertificates returns all Access TLS certificates for the
// zone level.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-properties
func (api *API) ZoneAccessMutualTLSCertificates(ctx context.Context, zoneID string) ([]AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificates(ctx, zoneID, ZoneRouteRoot)
}

func (api *API) accessMutualTLSCertificates(ctx context.Context, id string, routeRoot RouteRoot) ([]AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates",
		routeRoot,
		id,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessMutualTLSCertificate{}, err
	}

	var accessMutualTLSCertificateListResponse AccessMutualTLSCertificateListResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateListResponse)
	if err != nil {
		return []AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateListResponse.Result, nil
}

// AccessMutualTLSCertificate returns a single account level Access Mutual TLS
// certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-access-certificate-details
func (api *API) AccessMutualTLSCertificate(ctx context.Context, accountID, certificateID string) (AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificate(ctx, accountID, certificateID, AccountRouteRoot)
}

// ZoneAccessMutualTLSCertificate returns a single zone level Access Mutual TLS
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-access-certificate-details
func (api *API) ZoneAccessMutualTLSCertificate(ctx context.Context, zoneID, certificateID string) (AccessMutualTLSCertificate, error) {
	return api.accessMutualTLSCertificate(ctx, zoneID, certificateID, ZoneRouteRoot)
}

func (api *API) accessMutualTLSCertificate(ctx context.Context, id, certificateID string, routeRoot RouteRoot) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		routeRoot,
		id,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessMutualTLSCertificate{}, err
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// CreateAccessMutualTLSCertificate creates an account level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-create-access-certificate
func (api *API) CreateAccessMutualTLSCertificate(ctx context.Context, accountID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.createAccessMutualTLSCertificate(ctx, accountID, certificate, AccountRouteRoot)
}

// CreateZoneAccessMutualTLSCertificate creates a zone level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-create-access-certificate
func (api *API) CreateZoneAccessMutualTLSCertificate(ctx context.Context, zoneID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.createAccessMutualTLSCertificate(ctx, zoneID, certificate, ZoneRouteRoot)
}

func (api *API) createAccessMutualTLSCertificate(ctx context.Context, id string, certificate AccessMutualTLSCertificate, routeRoot RouteRoot) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates",
		routeRoot,
		id,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, certificate)
	if err != nil {
		return AccessMutualTLSCertificate{}, err
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// UpdateAccessMutualTLSCertificate updates an account level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-update-access-certificate
func (api *API) UpdateAccessMutualTLSCertificate(ctx context.Context, accountID, certificateID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.updateAccessMutualTLSCertificate(ctx, accountID, certificateID, certificate, AccountRouteRoot)
}

// UpdateZoneAccessMutualTLSCertificate updates a zone level Access TLS Mutual
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-update-access-certificate
func (api *API) UpdateZoneAccessMutualTLSCertificate(ctx context.Context, zoneID, certificateID string, certificate AccessMutualTLSCertificate) (AccessMutualTLSCertificate, error) {
	return api.updateAccessMutualTLSCertificate(ctx, zoneID, certificateID, certificate, ZoneRouteRoot)
}

func (api *API) updateAccessMutualTLSCertificate(ctx context.Context, id string, certificateID string, certificate AccessMutualTLSCertificate, routeRoot RouteRoot) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		routeRoot,
		id,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, certificate)
	if err != nil {
		return AccessMutualTLSCertificate{}, err
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// DeleteAccessMutualTLSCertificate destroys an account level Access Mutual
// TLS certificate.
//
// API reference: https://api.cloudflare.com/#access-mutual-tls-authentication-update-access-certificate
func (api *API) DeleteAccessMutualTLSCertificate(ctx context.Context, accountID, certificateID string) error {
	return api.deleteAccessMutualTLSCertificate(ctx, accountID, certificateID, AccountRouteRoot)
}

// DeleteZoneAccessMutualTLSCertificate destroys a zone level Access Mutual TLS
// certificate.
//
// API reference: https://api.cloudflare.com/#zone-level-access-mutual-tls-authentication-update-access-certificate
func (api *API) DeleteZoneAccessMutualTLSCertificate(ctx context.Context, zoneID, certificateID string) error {
	return api.deleteAccessMutualTLSCertificate(ctx, zoneID, certificateID, ZoneRouteRoot)
}

func (api *API) deleteAccessMutualTLSCertificate(ctx context.Context, id, certificateID string, routeRoot RouteRoot) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		routeRoot,
		id,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// AccessMutualTLSCertificate is the structure of a single Access Mutual TLS
// certificate.
type AccessMutualTLSCertificate struct {
	ID                  string    `json:"id,omitempty"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
	ExpiresOn           time.Time `json:"expires_on,omitempty"`
	Name                string    `json:"name,omitempty"`
	Fingerprint         string    `json:"fingerprint,omitempty"`
	Certificate         string    `json:"certificate,omitempty"`
	AssociatedHostnames []string  `json:"associated_hostnames,omitempty"`
}

// AccessMutualTLSCertificateListResponse is the API response for all Access
// Mutual TLS certificates.
type AccessMutualTLSCertificateListResponse struct {
	Response
	Result     []AccessMutualTLSCertificate `json:"result"`
	ResultInfo `json:"result_info"`
}

// AccessMutualTLSCertificateDetailResponse is the API response for a single
// Access Mutual TLS certificate.
type AccessMutualTLSCertificateDetailResponse struct {
	Response
	Result AccessMutualTLSCertificate `json:"result"`
}

type ListAccessMutualTLSCertificatesParams struct {
	ResultInfo
}

type CreateAccessMutualTLSCertificateParams struct {
	ExpiresOn           time.Time `json:"expires_on,omitempty"`
	Name                string    `json:"name,omitempty"`
	Fingerprint         string    `json:"fingerprint,omitempty"`
	Certificate         string    `json:"certificate,omitempty"`
	AssociatedHostnames []string  `json:"associated_hostnames,omitempty"`
}

type UpdateAccessMutualTLSCertificateParams struct {
	ID                  string    `json:"-"`
	ExpiresOn           time.Time `json:"expires_on,omitempty"`
	Name                string    `json:"name,omitempty"`
	Fingerprint         string    `json:"fingerprint,omitempty"`
	Certificate         string    `json:"certificate,omitempty"`
	AssociatedHostnames []string  `json:"associated_hostnames,omitempty"`
}

type AccessMutualTLSHostnameSettings struct {
	ChinaNetwork                *bool  `json:"china_network,omitempty"`
	ClientCertificateForwarding *bool  `json:"client_certificate_forwarding,omitempty"`
	Hostname                    string `json:"hostname,omitempty"`
}

type GetAccessMutualTLSHostnameSettingsResponse struct {
	Response
	Result []AccessMutualTLSHostnameSettings `json:"result"`
}

type UpdateAccessMutualTLSHostnameSettingsParams struct {
	Settings []AccessMutualTLSHostnameSettings `json:"settings,omitempty"`
}

// ListAccessMutualTLSCertificates returns all Access TLS certificates
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-list-mtls-certificates
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-list-mtls-certificates
func (api *API) ListAccessMutualTLSCertificates(ctx context.Context, rc *ResourceContainer, params ListAccessMutualTLSCertificatesParams) ([]AccessMutualTLSCertificate, *ResultInfo, error) {
	baseURL := fmt.Sprintf(
		"/%s/%s/access/certificates",
		rc.Level,
		rc.Identifier,
	)

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

	var accessCertificates []AccessMutualTLSCertificate
	var r AccessMutualTLSCertificateListResponse

	for {
		uri := buildURI(baseURL, params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []AccessMutualTLSCertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []AccessMutualTLSCertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		accessCertificates = append(accessCertificates, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return accessCertificates, &r.ResultInfo, nil
}

// GetAccessMutualTLSCertificate returns a single Access Mutual TLS
// certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-get-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-get-an-mtls-certificate
func (api *API) GetAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		rc.Level,
		rc.Identifier,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// CreateAccessMutualTLSCertificate creates an Access TLS Mutual
// certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-add-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-add-an-mtls-certificate
func (api *API) CreateAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, params CreateAccessMutualTLSCertificateParams) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, params)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// UpdateAccessMutualTLSCertificate updates an account level Access TLS Mutual
// certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-update-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-update-an-mtls-certificate
func (api *API) UpdateAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, params UpdateAccessMutualTLSCertificateParams) (AccessMutualTLSCertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		rc.Level,
		rc.Identifier,
		params.ID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return AccessMutualTLSCertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSCertificateDetailResponse.Result, nil
}

// DeleteAccessMutualTLSCertificate destroys an Access Mutual
// TLS certificate.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-delete-an-mtls-certificate
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-delete-an-mtls-certificate
func (api *API) DeleteAccessMutualTLSCertificate(ctx context.Context, rc *ResourceContainer, certificateID string) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/%s",
		rc.Level,
		rc.Identifier,
		certificateID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSCertificateDetailResponse AccessMutualTLSCertificateDetailResponse
	err = json.Unmarshal(res, &accessMutualTLSCertificateDetailResponse)
	if err != nil {
		return fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return nil
}

// GetAccessMutualTLSHostnameSettings returns all Access mTLS hostname settings.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-update-an-mtls-certificate-settings
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-list-mtls-certificates-hostname-settings
func (api *API) GetAccessMutualTLSHostnameSettings(ctx context.Context, rc *ResourceContainer) ([]AccessMutualTLSHostnameSettings, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/settings",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessMutualTLSHostnameSettings{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSHostnameSettingsResponse GetAccessMutualTLSHostnameSettingsResponse
	err = json.Unmarshal(res, &accessMutualTLSHostnameSettingsResponse)
	if err != nil {
		return []AccessMutualTLSHostnameSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSHostnameSettingsResponse.Result, nil
}

// UpdateAccessMutualTLSHostnameSettings updates Access mTLS certificate hostname settings.
//
// Account API Reference: https://developers.cloudflare.com/api/operations/access-mtls-authentication-update-an-mtls-certificate-settings
// Zone API Reference: https://developers.cloudflare.com/api/operations/zone-level-access-mtls-authentication-update-an-mtls-certificate-settings
func (api *API) UpdateAccessMutualTLSHostnameSettings(ctx context.Context, rc *ResourceContainer, params UpdateAccessMutualTLSHostnameSettingsParams) ([]AccessMutualTLSHostnameSettings, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/certificates/settings",
		rc.Level,
		rc.Identifier,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return []AccessMutualTLSHostnameSettings{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var accessMutualTLSHostnameSettingsResponse GetAccessMutualTLSHostnameSettingsResponse
	err = json.Unmarshal(res, &accessMutualTLSHostnameSettingsResponse)
	if err != nil {
		return []AccessMutualTLSHostnameSettings{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessMutualTLSHostnameSettingsResponse.Result, nil
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}
