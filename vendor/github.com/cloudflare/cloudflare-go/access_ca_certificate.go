package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// AccessCACertificate is the structure of the CA certificate used for
// short lived certificates.
type AccessCACertificate struct {
	ID        string `json:"id"`
	Aud       string `json:"aud"`
	PublicKey string `json:"public_key"`
}

// AccessCACertificateListResponse represents the response of all CA
// certificates within Access.
type AccessCACertificateListResponse struct {
	Response
	Result []AccessCACertificate `json:"result"`
}

// AccessCACertificateResponse represents the response of a single CA
// certificate.
type AccessCACertificateResponse struct {
	Response
	Result AccessCACertificate `json:"result"`
}

// AccessCACertificates returns all CA certificates within Access.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-list-short-lived-certificates
func (api *API) AccessCACertificates(accountID string) ([]AccessCACertificate, error) {
	uri := fmt.Sprintf("/accounts/%s/access/apps/ca", accountID)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return []AccessCACertificate{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessCAListResponse AccessCACertificateListResponse
	err = json.Unmarshal(res, &accessCAListResponse)
	if err != nil {
		return []AccessCACertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessCAListResponse.Result, nil
}

// AccessCACertificate returns a single CA certificate associated with an Access
// Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-short-lived-certificate-details
func (api *API) AccessCACertificate(accountID, applicationID string) (AccessCACertificate, error) {
	uri := fmt.Sprintf("/accounts/%s/access/apps/%s/ca", accountID, applicationID)

	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return AccessCACertificate{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessCAResponse AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCAResponse)
	if err != nil {
		return AccessCACertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessCAResponse.Result, nil
}

// CreateAccessCACertificate creates a new CA certificate for an Access
// Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-create-short-lived-certificate
func (api *API) CreateAccessCACertificate(accountID, applicationID string) (AccessCACertificate, error) {
	uri := fmt.Sprintf(
		"/accounts/%s/access/apps/%s/ca",
		accountID,
		applicationID,
	)

	res, err := api.makeRequest("POST", uri, nil)
	if err != nil {
		return AccessCACertificate{}, errors.Wrap(err, errMakeRequestError)
	}

	var accessCACertificate AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCACertificate)
	if err != nil {
		return AccessCACertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessCACertificate.Result, nil
}

// DeleteAccessCACertificate deletes an Access CA certificate on a defined
// Access Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-delete-access-certificate
func (api *API) DeleteAccessCACertificate(accountID, applicationID string) error {
	uri := fmt.Sprintf(
		"/accounts/%s/access/apps/%s/ca",
		accountID,
		applicationID,
	)

	_, err := api.makeRequest("DELETE", uri, nil)
	if err != nil {
		return errors.Wrap(err, errMakeRequestError)
||||||| parent of 6b7ce455e (update vendored files)
=======
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// AccessCACertificate is the structure of the CA certificate used for
// short lived certificates.
type AccessCACertificate struct {
	ID        string `json:"id"`
	Aud       string `json:"aud"`
	PublicKey string `json:"public_key"`
}

// AccessCACertificateListResponse represents the response of all CA
// certificates within Access.
type AccessCACertificateListResponse struct {
	Response
	Result []AccessCACertificate `json:"result"`
}

// AccessCACertificateResponse represents the response of a single CA
// certificate.
type AccessCACertificateResponse struct {
	Response
	Result AccessCACertificate `json:"result"`
}

// AccessCACertificates returns all CA certificates within Access.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-list-short-lived-certificates
func (api *API) AccessCACertificates(ctx context.Context, accountID string) ([]AccessCACertificate, error) {
	return api.accessCACertificates(ctx, accountID, AccountRouteRoot)
}

// ZoneLevelAccessCACertificates returns all zone level CA certificates within Access.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-list-short-lived-certificates
func (api *API) ZoneLevelAccessCACertificates(ctx context.Context, zoneID string) ([]AccessCACertificate, error) {
	return api.accessCACertificates(ctx, zoneID, ZoneRouteRoot)
}

func (api *API) accessCACertificates(ctx context.Context, id string, routeRoot RouteRoot) ([]AccessCACertificate, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/ca", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessCACertificate{}, err
	}

	var accessCAListResponse AccessCACertificateListResponse
	err = json.Unmarshal(res, &accessCAListResponse)
	if err != nil {
		return []AccessCACertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessCAListResponse.Result, nil
}

// AccessCACertificate returns a single CA certificate associated with an Access
// Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-short-lived-certificate-details
func (api *API) AccessCACertificate(ctx context.Context, accountID, applicationID string) (AccessCACertificate, error) {
	return api.accessCACertificate(ctx, accountID, applicationID, AccountRouteRoot)
}

// ZoneLevelAccessCACertificate returns a single zone level CA certificate associated with an Access
// Application.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-short-lived-certificate-details
func (api *API) ZoneLevelAccessCACertificate(ctx context.Context, zoneID, applicationID string) (AccessCACertificate, error) {
	return api.accessCACertificate(ctx, zoneID, applicationID, ZoneRouteRoot)
}

func (api *API) accessCACertificate(ctx context.Context, id, applicationID string, routeRoot RouteRoot) (AccessCACertificate, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/%s/ca", routeRoot, id, applicationID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessCACertificate{}, err
	}

	var accessCAResponse AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCAResponse)
	if err != nil {
		return AccessCACertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessCAResponse.Result, nil
}

// CreateAccessCACertificate creates a new CA certificate for an Access
// Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-create-short-lived-certificate
func (api *API) CreateAccessCACertificate(ctx context.Context, accountID, applicationID string) (AccessCACertificate, error) {
	return api.createAccessCACertificate(ctx, accountID, applicationID, AccountRouteRoot)
}

// CreateZoneLevelAccessCACertificate creates a new zone level CA certificate for an Access
// Application.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-create-short-lived-certificate
func (api *API) CreateZoneLevelAccessCACertificate(ctx context.Context, zoneID string, applicationID string) (AccessCACertificate, error) {
	return api.createAccessCACertificate(ctx, zoneID, applicationID, ZoneRouteRoot)
}

func (api *API) createAccessCACertificate(ctx context.Context, id string, applicationID string, routeRoot RouteRoot) (AccessCACertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/apps/%s/ca",
		routeRoot,
		id,
		applicationID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return AccessCACertificate{}, err
	}

	var accessCACertificate AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCACertificate)
	if err != nil {
		return AccessCACertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessCACertificate.Result, nil
}

// DeleteAccessCACertificate deletes an Access CA certificate on a defined
// Access Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-delete-access-certificate
func (api *API) DeleteAccessCACertificate(ctx context.Context, accountID, applicationID string) error {
	return api.deleteAccessCACertificate(ctx, accountID, applicationID, AccountRouteRoot)
}

// DeleteZoneLevelAccessCACertificate deletes a zone level Access CA certificate on a defined
// Access Application.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-delete-access-certificate
func (api *API) DeleteZoneLevelAccessCACertificate(ctx context.Context, zoneID, applicationID string) error {
	return api.deleteAccessCACertificate(ctx, zoneID, applicationID, ZoneRouteRoot)
}

func (api *API) deleteAccessCACertificate(ctx context.Context, id string, applicationID string, routeRoot RouteRoot) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/apps/%s/ca",
		routeRoot,
		id,
		applicationID,
	)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// AccessCACertificate is the structure of the CA certificate used for
// short lived certificates.
type AccessCACertificate struct {
	ID        string `json:"id"`
	Aud       string `json:"aud"`
	PublicKey string `json:"public_key"`
}

// AccessCACertificateListResponse represents the response of all CA
// certificates within Access.
type AccessCACertificateListResponse struct {
	Response
	Result []AccessCACertificate `json:"result"`
}

// AccessCACertificateResponse represents the response of a single CA
// certificate.
type AccessCACertificateResponse struct {
	Response
	Result AccessCACertificate `json:"result"`
}

// AccessCACertificates returns all CA certificates within Access.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-list-short-lived-certificates
func (api *API) AccessCACertificates(ctx context.Context, accountID string) ([]AccessCACertificate, error) {
	return api.accessCACertificates(ctx, accountID, AccountRouteRoot)
}

// ZoneLevelAccessCACertificates returns all zone level CA certificates within Access.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-list-short-lived-certificates
func (api *API) ZoneLevelAccessCACertificates(ctx context.Context, zoneID string) ([]AccessCACertificate, error) {
	return api.accessCACertificates(ctx, zoneID, ZoneRouteRoot)
}

func (api *API) accessCACertificates(ctx context.Context, id string, routeRoot RouteRoot) ([]AccessCACertificate, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/ca", routeRoot, id)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return []AccessCACertificate{}, err
	}

	var accessCAListResponse AccessCACertificateListResponse
	err = json.Unmarshal(res, &accessCAListResponse)
	if err != nil {
		return []AccessCACertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessCAListResponse.Result, nil
}

// AccessCACertificate returns a single CA certificate associated with an Access
// Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-short-lived-certificate-details
func (api *API) AccessCACertificate(ctx context.Context, accountID, applicationID string) (AccessCACertificate, error) {
	return api.accessCACertificate(ctx, accountID, applicationID, AccountRouteRoot)
}

// ZoneLevelAccessCACertificate returns a single zone level CA certificate associated with an Access
// Application.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-short-lived-certificate-details
func (api *API) ZoneLevelAccessCACertificate(ctx context.Context, zoneID, applicationID string) (AccessCACertificate, error) {
	return api.accessCACertificate(ctx, zoneID, applicationID, ZoneRouteRoot)
}

func (api *API) accessCACertificate(ctx context.Context, id, applicationID string, routeRoot RouteRoot) (AccessCACertificate, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/%s/ca", routeRoot, id, applicationID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessCACertificate{}, err
	}

	var accessCAResponse AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCAResponse)
	if err != nil {
		return AccessCACertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessCAResponse.Result, nil
}

// CreateAccessCACertificate creates a new CA certificate for an Access
// Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-create-short-lived-certificate
func (api *API) CreateAccessCACertificate(ctx context.Context, accountID, applicationID string) (AccessCACertificate, error) {
	return api.createAccessCACertificate(ctx, accountID, applicationID, AccountRouteRoot)
}

// CreateZoneLevelAccessCACertificate creates a new zone level CA certificate for an Access
// Application.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-create-short-lived-certificate
func (api *API) CreateZoneLevelAccessCACertificate(ctx context.Context, zoneID string, applicationID string) (AccessCACertificate, error) {
	return api.createAccessCACertificate(ctx, zoneID, applicationID, ZoneRouteRoot)
}

func (api *API) createAccessCACertificate(ctx context.Context, id string, applicationID string, routeRoot RouteRoot) (AccessCACertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/apps/%s/ca",
		routeRoot,
		id,
		applicationID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return AccessCACertificate{}, err
	}

	var accessCACertificate AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCACertificate)
	if err != nil {
		return AccessCACertificate{}, errors.Wrap(err, errUnmarshalError)
	}

	return accessCACertificate.Result, nil
}

// DeleteAccessCACertificate deletes an Access CA certificate on a defined
// Access Application.
//
// API reference: https://api.cloudflare.com/#access-short-lived-certificates-delete-access-certificate
func (api *API) DeleteAccessCACertificate(ctx context.Context, accountID, applicationID string) error {
	return api.deleteAccessCACertificate(ctx, accountID, applicationID, AccountRouteRoot)
}

// DeleteZoneLevelAccessCACertificate deletes a zone level Access CA certificate on a defined
// Access Application.
//
// API reference: https://api.cloudflare.com/#zone-level-access-short-lived-certificates-delete-access-certificate
func (api *API) DeleteZoneLevelAccessCACertificate(ctx context.Context, zoneID, applicationID string) error {
	return api.deleteAccessCACertificate(ctx, zoneID, applicationID, ZoneRouteRoot)
}

func (api *API) deleteAccessCACertificate(ctx context.Context, id string, applicationID string, routeRoot RouteRoot) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/apps/%s/ca",
		routeRoot,
		id,
		applicationID,
	)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"context"
	"fmt"
	"net/http"

	"github.com/goccy/go-json"
)

// AccessCACertificate is the structure of the CA certificate used for
// short-lived certificates.
type AccessCACertificate struct {
	ID        string `json:"id"`
	Aud       string `json:"aud"`
	PublicKey string `json:"public_key"`
}

// AccessCACertificateListResponse represents the response of all CA
// certificates within Access.
type AccessCACertificateListResponse struct {
	Response
	Result []AccessCACertificate `json:"result"`
	ResultInfo
}

// AccessCACertificateResponse represents the response of a single CA
// certificate.
type AccessCACertificateResponse struct {
	Response
	Result AccessCACertificate `json:"result"`
}

type ListAccessCACertificatesParams struct {
	ResultInfo
}

type CreateAccessCACertificateParams struct {
	ApplicationID string
}

// ListAccessCACertificates returns all AccessCACertificate within Access.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-short-lived-certificate-c-as-list-short-lived-certificate-c-as
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-short-lived-certificate-c-as-list-short-lived-certificate-c-as
func (api *API) ListAccessCACertificates(ctx context.Context, rc *ResourceContainer, params ListAccessCACertificatesParams) ([]AccessCACertificate, *ResultInfo, error) {
	baseURL := fmt.Sprintf("/%s/%s/access/apps/ca", rc.Level, rc.Identifier)

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

	var accessCACertificates []AccessCACertificate
	var r AccessCACertificateListResponse

	for {
		uri := buildURI(baseURL, params)
		res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
		if err != nil {
			return []AccessCACertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
		}

		err = json.Unmarshal(res, &r)
		if err != nil {
			return []AccessCACertificate{}, &ResultInfo{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
		}
		accessCACertificates = append(accessCACertificates, r.Result...)
		params.ResultInfo = r.ResultInfo.Next()
		if params.ResultInfo.Done() || !autoPaginate {
			break
		}
	}

	return accessCACertificates, &r.ResultInfo, nil
}

// GetAccessCACertificate returns a single CA certificate associated within
// Access.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-short-lived-certificate-c-as-get-a-short-lived-certificate-ca
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-short-lived-certificate-c-as-get-a-short-lived-certificate-ca
func (api *API) GetAccessCACertificate(ctx context.Context, rc *ResourceContainer, applicationID string) (AccessCACertificate, error) {
	uri := fmt.Sprintf("/%s/%s/access/apps/%s/ca", rc.Level, rc.Identifier, applicationID)

	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AccessCACertificate{}, err
	}

	var accessCAResponse AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCAResponse)
	if err != nil {
		return AccessCACertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessCAResponse.Result, nil
}

// CreateAccessCACertificate creates a new CA certificate for an AccessApplication.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-short-lived-certificate-c-as-create-a-short-lived-certificate-ca
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-short-lived-certificate-c-as-create-a-short-lived-certificate-ca
func (api *API) CreateAccessCACertificate(ctx context.Context, rc *ResourceContainer, params CreateAccessCACertificateParams) (AccessCACertificate, error) {
	uri := fmt.Sprintf(
		"/%s/%s/access/apps/%s/ca",
		rc.Level,
		rc.Identifier,
		params.ApplicationID,
	)

	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, nil)
	if err != nil {
		return AccessCACertificate{}, err
	}

	var accessCACertificate AccessCACertificateResponse
	err = json.Unmarshal(res, &accessCACertificate)
	if err != nil {
		return AccessCACertificate{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return accessCACertificate.Result, nil
}

// DeleteAccessCACertificate deletes an Access CA certificate on a defined
// AccessApplication.
//
// Account API reference: https://developers.cloudflare.com/api/operations/access-short-lived-certificate-c-as-delete-a-short-lived-certificate-ca
// Zone API reference: https://developers.cloudflare.com/api/operations/zone-level-access-short-lived-certificate-c-as-delete-a-short-lived-certificate-ca
func (api *API) DeleteAccessCACertificate(ctx context.Context, rc *ResourceContainer, applicationID string) error {
	uri := fmt.Sprintf(
		"/%s/%s/access/apps/%s/ca",
		rc.Level,
		rc.Identifier,
		applicationID,
	)

	_, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return err
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}

	return nil
}
