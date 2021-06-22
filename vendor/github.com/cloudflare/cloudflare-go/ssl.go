package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// ZoneCustomSSL represents custom SSL certificate metadata.
type ZoneCustomSSL struct {
	ID              string                       `json:"id"`
	Hosts           []string                     `json:"hosts"`
	Issuer          string                       `json:"issuer"`
	Signature       string                       `json:"signature"`
	Status          string                       `json:"status"`
	BundleMethod    string                       `json:"bundle_method"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions"`
	ZoneID          string                       `json:"zone_id"`
	UploadedOn      time.Time                    `json:"uploaded_on"`
	ModifiedOn      time.Time                    `json:"modified_on"`
	ExpiresOn       time.Time                    `json:"expires_on"`
	Priority        int                          `json:"priority"`
	KeylessServer   KeylessSSL                   `json:"keyless_server"`
}

// ZoneCustomSSLGeoRestrictions represents the parameter to create or update
// geographic restrictions on a custom ssl certificate.
type ZoneCustomSSLGeoRestrictions struct {
	Label string `json:"label"`
}

// zoneCustomSSLResponse represents the response from the zone SSL details endpoint.
type zoneCustomSSLResponse struct {
	Response
	Result ZoneCustomSSL `json:"result"`
}

// zoneCustomSSLsResponse represents the response from the zone SSL list endpoint.
type zoneCustomSSLsResponse struct {
	Response
	Result []ZoneCustomSSL `json:"result"`
}

// ZoneCustomSSLOptions represents the parameters to create or update an existing
// custom SSL configuration.
type ZoneCustomSSLOptions struct {
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	Certificate     string                        `json:"certificate"`
	PrivateKey      string                        `json:"private_key"`
	BundleMethod    string                        `json:"bundle_method,omitempty"`
	GeoRestrictions *ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                        `json:"type,omitempty"`
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	Certificate     string                       `json:"certificate"`
	PrivateKey      string                       `json:"private_key"`
	BundleMethod    string                       `json:"bundle_method,omitempty"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                       `json:"type,omitempty"`
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	Certificate     string                       `json:"certificate"`
	PrivateKey      string                       `json:"private_key"`
	BundleMethod    string                       `json:"bundle_method,omitempty"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                       `json:"type,omitempty"`
=======
	Certificate     string                        `json:"certificate"`
	PrivateKey      string                        `json:"private_key"`
	BundleMethod    string                        `json:"bundle_method,omitempty"`
	GeoRestrictions *ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                        `json:"type,omitempty"`
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	Certificate     string                       `json:"certificate"`
	PrivateKey      string                       `json:"private_key"`
	BundleMethod    string                       `json:"bundle_method,omitempty"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                       `json:"type,omitempty"`
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	Certificate     string                       `json:"certificate"`
	PrivateKey      string                       `json:"private_key"`
	BundleMethod    string                       `json:"bundle_method,omitempty"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                       `json:"type,omitempty"`
=======
	Certificate     string                        `json:"certificate"`
	PrivateKey      string                        `json:"private_key"`
	BundleMethod    string                        `json:"bundle_method,omitempty"`
	GeoRestrictions *ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                        `json:"type,omitempty"`
>>>>>>> 6b7ce455e (update vendored files)
}

// ZoneCustomSSLPriority represents a certificate's ID and priority. It is a
// subset of ZoneCustomSSL used for patch requests.
type ZoneCustomSSLPriority struct {
	ID       string `json:"ID"`
	Priority int    `json:"priority"`
}

// CreateSSL allows you to add a custom SSL certificate to the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-create-ssl-configuration
func (api *API) CreateSSL(ctx context.Context, zoneID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, options)
	if err != nil {
		return ZoneCustomSSL{}, err
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListSSL lists the custom certificates for the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-list-ssl-configurations
func (api *API) ListSSL(ctx context.Context, zoneID string) ([]ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r zoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SSLDetails returns the configuration details for a custom SSL certificate.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-ssl-configuration-details
func (api *API) SSLDetails(ctx context.Context, zoneID, certificateID string) (ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/%s", zoneID, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneCustomSSL{}, err
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateSSL updates (replaces) a custom SSL certificate.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-update-ssl-configuration
func (api *API) UpdateSSL(ctx context.Context, zoneID, certificateID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/%s", zoneID, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, options)
	if err != nil {
		return ZoneCustomSSL{}, err
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ReprioritizeSSL allows you to change the priority (which is served for a given
// request) of custom SSL certificates associated with the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-re-prioritize-ssl-certificates
func (api *API) ReprioritizeSSL(ctx context.Context, zoneID string, p []ZoneCustomSSLPriority) ([]ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/prioritize", zoneID)
	params := struct {
		Certificates []ZoneCustomSSLPriority `json:"certificates"`
	}{
		Certificates: p,
	}
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return nil, err
	}
	var r zoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteSSL deletes a custom SSL certificate from the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-delete-an-ssl-certificate
func (api *API) DeleteSSL(ctx context.Context, zoneID, certificateID string) error {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/%s", zoneID, certificateID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
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

// ZoneCustomSSL represents custom SSL certificate metadata.
type ZoneCustomSSL struct {
	ID              string                       `json:"id"`
	Hosts           []string                     `json:"hosts"`
	Issuer          string                       `json:"issuer"`
	Signature       string                       `json:"signature"`
	Status          string                       `json:"status"`
	BundleMethod    string                       `json:"bundle_method"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions"`
	ZoneID          string                       `json:"zone_id"`
	UploadedOn      time.Time                    `json:"uploaded_on"`
	ModifiedOn      time.Time                    `json:"modified_on"`
	ExpiresOn       time.Time                    `json:"expires_on"`
	Priority        int                          `json:"priority"`
	KeylessServer   KeylessSSL                   `json:"keyless_server"`
}

// ZoneCustomSSLGeoRestrictions represents the parameter to create or update
// geographic restrictions on a custom ssl certificate.
type ZoneCustomSSLGeoRestrictions struct {
	Label string `json:"label"`
}

// zoneCustomSSLResponse represents the response from the zone SSL details endpoint.
type zoneCustomSSLResponse struct {
	Response
	Result ZoneCustomSSL `json:"result"`
}

// zoneCustomSSLsResponse represents the response from the zone SSL list endpoint.
type zoneCustomSSLsResponse struct {
	Response
	Result []ZoneCustomSSL `json:"result"`
}

// ZoneCustomSSLOptions represents the parameters to create or update an existing
// custom SSL configuration.
type ZoneCustomSSLOptions struct {
	Certificate     string                        `json:"certificate"`
	PrivateKey      string                        `json:"private_key"`
	BundleMethod    string                        `json:"bundle_method,omitempty"`
	GeoRestrictions *ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                        `json:"type,omitempty"`
}

// ZoneCustomSSLPriority represents a certificate's ID and priority. It is a
// subset of ZoneCustomSSL used for patch requests.
type ZoneCustomSSLPriority struct {
	ID       string `json:"ID"`
	Priority int    `json:"priority"`
}

// CreateSSL allows you to add a custom SSL certificate to the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-create-ssl-configuration
func (api *API) CreateSSL(ctx context.Context, zoneID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodPost, uri, options)
	if err != nil {
		return ZoneCustomSSL{}, err
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ListSSL lists the custom certificates for the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-list-ssl-configurations
func (api *API) ListSSL(ctx context.Context, zoneID string) ([]ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}
	var r zoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// SSLDetails returns the configuration details for a custom SSL certificate.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-ssl-configuration-details
func (api *API) SSLDetails(ctx context.Context, zoneID, certificateID string) (ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/%s", zoneID, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return ZoneCustomSSL{}, err
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// UpdateSSL updates (replaces) a custom SSL certificate.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-update-ssl-configuration
func (api *API) UpdateSSL(ctx context.Context, zoneID, certificateID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/%s", zoneID, certificateID)
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, options)
	if err != nil {
		return ZoneCustomSSL{}, err
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// ReprioritizeSSL allows you to change the priority (which is served for a given
// request) of custom SSL certificates associated with the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-re-prioritize-ssl-certificates
func (api *API) ReprioritizeSSL(ctx context.Context, zoneID string, p []ZoneCustomSSLPriority) ([]ZoneCustomSSL, error) {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/prioritize", zoneID)
	params := struct {
		Certificates []ZoneCustomSSLPriority `json:"certificates"`
	}{
		Certificates: p,
	}
	res, err := api.makeRequestContext(ctx, http.MethodPut, uri, params)
	if err != nil {
		return nil, err
	}
	var r zoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// DeleteSSL deletes a custom SSL certificate from the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-delete-an-ssl-certificate
<<<<<<< HEAD
func (api *API) DeleteSSL(zoneID, certificateID string) error {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
func (api *API) DeleteSSL(zoneID, certificateID string) error {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
=======
func (api *API) DeleteSSL(ctx context.Context, zoneID, certificateID string) error {
	uri := fmt.Sprintf("/zones/%s/custom_certificates/%s", zoneID, certificateID)
	if _, err := api.makeRequestContext(ctx, http.MethodDelete, uri, nil); err != nil {
		return err
>>>>>>> 4d7e5ad26 (update vendored files)
	}
	return nil
}

// SSLValidationRecord displays Domain Control Validation tokens.
type SSLValidationRecord struct {
	CnameTarget string `json:"cname_target,omitempty"`
	CnameName   string `json:"cname,omitempty"`

	TxtName  string `json:"txt_name,omitempty"`
	TxtValue string `json:"txt_value,omitempty"`

	HTTPUrl  string `json:"http_url,omitempty"`
	HTTPBody string `json:"http_body,omitempty"`

	Emails []string `json:"emails,omitempty"`
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// ZoneCustomSSL represents custom SSL certificate metadata.
type ZoneCustomSSL struct {
	ID              string                       `json:"id"`
	Hosts           []string                     `json:"hosts"`
	Issuer          string                       `json:"issuer"`
	Signature       string                       `json:"signature"`
	Status          string                       `json:"status"`
	BundleMethod    string                       `json:"bundle_method"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions"`
	ZoneID          string                       `json:"zone_id"`
	UploadedOn      time.Time                    `json:"uploaded_on"`
	ModifiedOn      time.Time                    `json:"modified_on"`
	ExpiresOn       time.Time                    `json:"expires_on"`
	Priority        int                          `json:"priority"`
	KeylessServer   KeylessSSL                   `json:"keyless_server"`
}

// ZoneCustomSSLGeoRestrictions represents the parameter to create or update
// geographic restrictions on a custom ssl certificate.
type ZoneCustomSSLGeoRestrictions struct {
	Label string `json:"label"`
}

// zoneCustomSSLResponse represents the response from the zone SSL details endpoint.
type zoneCustomSSLResponse struct {
	Response
	Result ZoneCustomSSL `json:"result"`
}

// zoneCustomSSLsResponse represents the response from the zone SSL list endpoint.
type zoneCustomSSLsResponse struct {
	Response
	Result []ZoneCustomSSL `json:"result"`
}

// ZoneCustomSSLOptions represents the parameters to create or update an existing
// custom SSL configuration.
type ZoneCustomSSLOptions struct {
	Certificate     string                       `json:"certificate"`
	PrivateKey      string                       `json:"private_key"`
	BundleMethod    string                       `json:"bundle_method,omitempty"`
	GeoRestrictions ZoneCustomSSLGeoRestrictions `json:"geo_restrictions,omitempty"`
	Type            string                       `json:"type,omitempty"`
}

// ZoneCustomSSLPriority represents a certificate's ID and priority. It is a
// subset of ZoneCustomSSL used for patch requests.
type ZoneCustomSSLPriority struct {
	ID       string `json:"ID"`
	Priority int    `json:"priority"`
}

// CreateSSL allows you to add a custom SSL certificate to the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-create-ssl-configuration
func (api *API) CreateSSL(zoneID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates"
	res, err := api.makeRequest("POST", uri, options)
	if err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errMakeRequestError)
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ListSSL lists the custom certificates for the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-list-ssl-configurations
func (api *API) ListSSL(zoneID string) ([]ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r zoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SSLDetails returns the configuration details for a custom SSL certificate.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-ssl-configuration-details
func (api *API) SSLDetails(zoneID, certificateID string) (ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errMakeRequestError)
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// UpdateSSL updates (replaces) a custom SSL certificate.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-update-ssl-configuration
func (api *API) UpdateSSL(zoneID, certificateID string, options ZoneCustomSSLOptions) (ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	res, err := api.makeRequest("PATCH", uri, options)
	if err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errMakeRequestError)
	}
	var r zoneCustomSSLResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return ZoneCustomSSL{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// ReprioritizeSSL allows you to change the priority (which is served for a given
// request) of custom SSL certificates associated with the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-re-prioritize-ssl-certificates
func (api *API) ReprioritizeSSL(zoneID string, p []ZoneCustomSSLPriority) ([]ZoneCustomSSL, error) {
	uri := "/zones/" + zoneID + "/custom_certificates/prioritize"
	params := struct {
		Certificates []ZoneCustomSSLPriority `json:"certificates"`
	}{
		Certificates: p,
	}
	res, err := api.makeRequest("PUT", uri, params)
	if err != nil {
		return nil, errors.Wrap(err, errMakeRequestError)
	}
	var r zoneCustomSSLsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return nil, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// DeleteSSL deletes a custom SSL certificate from the given zone.
//
// API reference: https://api.cloudflare.com/#custom-ssl-for-a-zone-delete-an-ssl-certificate
func (api *API) DeleteSSL(zoneID, certificateID string) error {
	uri := "/zones/" + zoneID + "/custom_certificates/" + certificateID
	if _, err := api.makeRequest("DELETE", uri, nil); err != nil {
		return errors.Wrap(err, errMakeRequestError)
	}
	return nil
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
