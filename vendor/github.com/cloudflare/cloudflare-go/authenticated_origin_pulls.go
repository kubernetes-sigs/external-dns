package cloudflare

import (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

// AuthenticatedOriginPulls represents global AuthenticatedOriginPulls (tls_client_auth) metadata.
type AuthenticatedOriginPulls struct {
	ID         string    `json:"id"`
	Value      string    `json:"value"`
	Editable   bool      `json:"editable"`
	ModifiedOn time.Time `json:"modified_on"`
}

// AuthenticatedOriginPullsResponse represents the response from the global AuthenticatedOriginPulls (tls_client_auth) details endpoint.
type AuthenticatedOriginPullsResponse struct {
	Response
	Result AuthenticatedOriginPulls `json:"result"`
}

// GetAuthenticatedOriginPullsStatus returns the configuration details for global AuthenticatedOriginPulls (tls_client_auth).
//
// API reference: https://api.cloudflare.com/#zone-settings-get-tls-client-auth-setting
func (api *API) GetAuthenticatedOriginPullsStatus(zoneID string) (AuthenticatedOriginPulls, error) {
	uri := "/zones/" + zoneID + "/settings/tls_client_auth"
	res, err := api.makeRequest("GET", uri, nil)
	if err != nil {
		return AuthenticatedOriginPulls{}, errors.Wrap(err, errMakeRequestError)
	}
	var r AuthenticatedOriginPullsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AuthenticatedOriginPulls{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SetAuthenticatedOriginPullsStatus toggles whether global AuthenticatedOriginPulls is enabled for the zone.
//
// API reference: https://api.cloudflare.com/#zone-settings-change-tls-client-auth-setting
func (api *API) SetAuthenticatedOriginPullsStatus(zoneID string, enable bool) (AuthenticatedOriginPulls, error) {
	uri := "/zones/" + zoneID + "/settings/tls_client_auth"
	var val string
	if enable {
		val = "on"
	} else {
		val = "off"
	}
	params := struct {
		Value string `json:"value"`
	}{
		Value: val,
	}
	res, err := api.makeRequest("PATCH", uri, params)
	if err != nil {
		return AuthenticatedOriginPulls{}, errors.Wrap(err, errMakeRequestError)
||||||| parent of 6b7ce455e (update vendored files)
=======
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// AuthenticatedOriginPulls represents global AuthenticatedOriginPulls (tls_client_auth) metadata.
type AuthenticatedOriginPulls struct {
	ID         string    `json:"id"`
	Value      string    `json:"value"`
	Editable   bool      `json:"editable"`
	ModifiedOn time.Time `json:"modified_on"`
}

// AuthenticatedOriginPullsResponse represents the response from the global AuthenticatedOriginPulls (tls_client_auth) details endpoint.
type AuthenticatedOriginPullsResponse struct {
	Response
	Result AuthenticatedOriginPulls `json:"result"`
}

// GetAuthenticatedOriginPullsStatus returns the configuration details for global AuthenticatedOriginPulls (tls_client_auth).
//
// API reference: https://api.cloudflare.com/#zone-settings-get-tls-client-auth-setting
func (api *API) GetAuthenticatedOriginPullsStatus(ctx context.Context, zoneID string) (AuthenticatedOriginPulls, error) {
	uri := fmt.Sprintf("/zones/%s/settings/tls_client_auth", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AuthenticatedOriginPulls{}, err
	}
	var r AuthenticatedOriginPullsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AuthenticatedOriginPulls{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// SetAuthenticatedOriginPullsStatus toggles whether global AuthenticatedOriginPulls is enabled for the zone.
//
// API reference: https://api.cloudflare.com/#zone-settings-change-tls-client-auth-setting
func (api *API) SetAuthenticatedOriginPullsStatus(ctx context.Context, zoneID string, enable bool) (AuthenticatedOriginPulls, error) {
	uri := fmt.Sprintf("/zones/%s/settings/tls_client_auth", zoneID)
	var val string
	if enable {
		val = "on"
	} else {
		val = "off"
	}
	params := struct {
		Value string `json:"value"`
	}{
		Value: val,
	}
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AuthenticatedOriginPulls{}, err
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// AuthenticatedOriginPulls represents global AuthenticatedOriginPulls (tls_client_auth) metadata.
type AuthenticatedOriginPulls struct {
	ID         string    `json:"id"`
	Value      string    `json:"value"`
	Editable   bool      `json:"editable"`
	ModifiedOn time.Time `json:"modified_on"`
}

// AuthenticatedOriginPullsResponse represents the response from the global AuthenticatedOriginPulls (tls_client_auth) details endpoint.
type AuthenticatedOriginPullsResponse struct {
	Response
	Result AuthenticatedOriginPulls `json:"result"`
}

// GetAuthenticatedOriginPullsStatus returns the configuration details for global AuthenticatedOriginPulls (tls_client_auth).
//
// API reference: https://api.cloudflare.com/#zone-settings-get-tls-client-auth-setting
func (api *API) GetAuthenticatedOriginPullsStatus(ctx context.Context, zoneID string) (AuthenticatedOriginPulls, error) {
	uri := fmt.Sprintf("/zones/%s/settings/tls_client_auth", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AuthenticatedOriginPulls{}, err
	}
	var r AuthenticatedOriginPullsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AuthenticatedOriginPulls{}, errors.Wrap(err, errUnmarshalError)
	}
	return r.Result, nil
}

// SetAuthenticatedOriginPullsStatus toggles whether global AuthenticatedOriginPulls is enabled for the zone.
//
// API reference: https://api.cloudflare.com/#zone-settings-change-tls-client-auth-setting
func (api *API) SetAuthenticatedOriginPullsStatus(ctx context.Context, zoneID string, enable bool) (AuthenticatedOriginPulls, error) {
	uri := fmt.Sprintf("/zones/%s/settings/tls_client_auth", zoneID)
	var val string
	if enable {
		val = "on"
	} else {
		val = "off"
	}
	params := struct {
		Value string `json:"value"`
	}{
		Value: val,
	}
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AuthenticatedOriginPulls{}, err
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

// AuthenticatedOriginPulls represents global AuthenticatedOriginPulls (tls_client_auth) metadata.
type AuthenticatedOriginPulls struct {
	ID         string    `json:"id"`
	Value      string    `json:"value"`
	Editable   bool      `json:"editable"`
	ModifiedOn time.Time `json:"modified_on"`
}

// AuthenticatedOriginPullsResponse represents the response from the global AuthenticatedOriginPulls (tls_client_auth) details endpoint.
type AuthenticatedOriginPullsResponse struct {
	Response
	Result AuthenticatedOriginPulls `json:"result"`
}

// GetAuthenticatedOriginPullsStatus returns the configuration details for global AuthenticatedOriginPulls (tls_client_auth).
//
// API reference: https://api.cloudflare.com/#zone-settings-get-tls-client-auth-setting
func (api *API) GetAuthenticatedOriginPullsStatus(ctx context.Context, zoneID string) (AuthenticatedOriginPulls, error) {
	uri := fmt.Sprintf("/zones/%s/settings/tls_client_auth", zoneID)
	res, err := api.makeRequestContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return AuthenticatedOriginPulls{}, err
	}
	var r AuthenticatedOriginPullsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AuthenticatedOriginPulls{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}

// SetAuthenticatedOriginPullsStatus toggles whether global AuthenticatedOriginPulls is enabled for the zone.
//
// API reference: https://api.cloudflare.com/#zone-settings-change-tls-client-auth-setting
func (api *API) SetAuthenticatedOriginPullsStatus(ctx context.Context, zoneID string, enable bool) (AuthenticatedOriginPulls, error) {
	uri := fmt.Sprintf("/zones/%s/settings/tls_client_auth", zoneID)
	var val string
	if enable {
		val = "on"
	} else {
		val = "off"
	}
	params := struct {
		Value string `json:"value"`
	}{
		Value: val,
	}
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return AuthenticatedOriginPulls{}, err
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	}
	var r AuthenticatedOriginPullsResponse
	if err := json.Unmarshal(res, &r); err != nil {
		return AuthenticatedOriginPulls{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}
	return r.Result, nil
}
