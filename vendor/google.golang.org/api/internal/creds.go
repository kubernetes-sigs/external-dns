// Copyright 2017 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"context"
<<<<<<< HEAD
	"crypto/tls"
	"encoding/json"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/api/internal/impersonate"

	"golang.org/x/oauth2/google"
)

const quotaProjectEnvVar = "GOOGLE_CLOUD_QUOTA_PROJECT"

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	creds, err := baseCreds(ctx, ds)
	if err != nil {
		return nil, err
	}
	if ds.ImpersonationConfig != nil {
		return impersonateCredentials(ctx, creds, ds)
	}
	return creds, nil
}

func baseCreds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	if ds.InternalCredentials != nil {
		return ds.InternalCredentials, nil
	}
	if ds.Credentials != nil {
		return ds.Credentials, nil
	}
	if ds.CredentialsJSON != nil {
		return credentialsFromJSON(ctx, ds.CredentialsJSON, ds)
	}
	if ds.CredentialsFile != "" {
		data, err := ioutil.ReadFile(ds.CredentialsFile)
		if err != nil {
			return nil, fmt.Errorf("cannot read credentials file: %v", err)
		}
		return credentialsFromJSON(ctx, data, ds)
	}
	if ds.TokenSource != nil {
		return &google.Credentials{TokenSource: ds.TokenSource}, nil
	}
	cred, err := google.FindDefaultCredentials(ctx, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}
	if len(cred.JSON) > 0 {
		return credentialsFromJSON(ctx, cred.JSON, ds)
	}
	// For GAE and GCE, the JSON is empty so return the default credentials directly.
	return cred, nil
}

// JSON key file type.
const (
	serviceAccountKey = "service_account"
)

// credentialsFromJSON returns a google.Credentials from the JSON data
//
// - A self-signed JWT flow will be executed if the following conditions are
// met:
//
//	(1) At least one of the following is true:
//	    (a) No scope is provided
//	    (b) Scope for self-signed JWT flow is enabled
//	    (c) Audiences are explicitly provided by users
//	(2) No service account impersontation
//
// - Otherwise, executes standard OAuth 2.0 flow
// More details: google.aip.dev/auth/4111
func credentialsFromJSON(ctx context.Context, data []byte, ds *DialSettings) (*google.Credentials, error) {
	var params google.CredentialsParams
	params.Scopes = ds.GetScopes()

	// Determine configurations for the OAuth2 transport, which is separate from the API transport.
	// The OAuth2 transport and endpoint will be configured for mTLS if applicable.
	clientCertSource, oauth2Endpoint, err := GetClientCertificateSourceAndEndpoint(oauth2DialSettings(ds))
	if err != nil {
		return nil, err
	}
	params.TokenURL = oauth2Endpoint
	if clientCertSource != nil {
		tlsConfig := &tls.Config{
			GetClientCertificate: clientCertSource,
		}
		ctx = context.WithValue(ctx, oauth2.HTTPClient, customHTTPClient(tlsConfig))
	}

	// By default, a standard OAuth 2.0 token source is created
	cred, err := google.CredentialsFromJSONWithParams(ctx, data, params)
	if err != nil {
		return nil, err
	}

	// Override the token source to use self-signed JWT if conditions are met
	isJWTFlow, err := isSelfSignedJWTFlow(data, ds)
	if err != nil {
		return nil, err
	}
	if isJWTFlow {
		ts, err := selfSignedJWTTokenSource(data, ds)
		if err != nil {
			return nil, err
		}
		cred.TokenSource = ts
	}

	return cred, err
}

func isSelfSignedJWTFlow(data []byte, ds *DialSettings) (bool, error) {
	if (ds.EnableJwtWithScope || ds.HasCustomAudience()) &&
		ds.ImpersonationConfig == nil {
		// Check if JSON is a service account and if so create a self-signed JWT.
		var f struct {
			Type string `json:"type"`
			// The rest JSON fields are omitted because they are not used.
		}
		if err := json.Unmarshal(data, &f); err != nil {
			return false, err
		}
		return f.Type == serviceAccountKey, nil
	}
	return false, nil
}

func selfSignedJWTTokenSource(data []byte, ds *DialSettings) (oauth2.TokenSource, error) {
	if len(ds.GetScopes()) > 0 && !ds.HasCustomAudience() {
		// Scopes are preferred in self-signed JWT unless the scope is not available
		// or a custom audience is used.
		return google.JWTAccessTokenSourceWithScope(data, ds.GetScopes()...)
	} else if ds.GetAudience() != "" {
		// Fallback to audience if scope is not provided
		return google.JWTAccessTokenSourceFromJSON(data, ds.GetAudience())
	} else {
		return nil, errors.New("neither scopes or audience are available for the self-signed JWT")
	}
}

// GetQuotaProject retrieves quota project with precedence being: client option,
// environment variable, creds file.
func GetQuotaProject(creds *google.Credentials, clientOpt string) string {
	if clientOpt != "" {
		return clientOpt
	}
	if env := os.Getenv(quotaProjectEnvVar); env != "" {
		return env
	}
	if creds == nil {
		return ""
	}
	var v struct {
		QuotaProject string `json:"quota_project_id"`
	}
	if err := json.Unmarshal(creds.JSON, &v); err != nil {
		return ""
	}
	return v.QuotaProject
}

func impersonateCredentials(ctx context.Context, creds *google.Credentials, ds *DialSettings) (*google.Credentials, error) {
	if len(ds.ImpersonationConfig.Scopes) == 0 {
		ds.ImpersonationConfig.Scopes = ds.GetScopes()
	}
	ts, err := impersonate.TokenSource(ctx, creds.TokenSource, ds.ImpersonationConfig)
	if err != nil {
		return nil, err
	}
	return &google.Credentials{
		TokenSource: ts,
		ProjectID:   creds.ProjectID,
	}, nil
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	"errors"
>>>>>>> 5ce8c7613 (update vendored files)
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
	"google.golang.org/api/internal/impersonate"

	"golang.org/x/oauth2/google"
)

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	creds, err := baseCreds(ctx, ds)
	if err != nil {
		return nil, err
	}
	if ds.ImpersonationConfig != nil {
		return impersonateCredentials(ctx, creds, ds)
	}
	return creds, nil
}

func baseCreds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	if ds.InternalCredentials != nil {
		return ds.InternalCredentials, nil
	}
	if ds.Credentials != nil {
		return ds.Credentials, nil
	}
	if ds.CredentialsJSON != nil {
		return credentialsFromJSON(ctx, ds.CredentialsJSON, ds)
	}
	if ds.CredentialsFile != "" {
		data, err := ioutil.ReadFile(ds.CredentialsFile)
		if err != nil {
			return nil, fmt.Errorf("cannot read credentials file: %v", err)
		}
		return credentialsFromJSON(ctx, data, ds)
	}
	if ds.TokenSource != nil {
		return &google.Credentials{TokenSource: ds.TokenSource}, nil
	}
	cred, err := google.FindDefaultCredentials(ctx, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}
	if len(cred.JSON) > 0 {
		return credentialsFromJSON(ctx, cred.JSON, ds)
	}
	// For GAE and GCE, the JSON is empty so return the default credentials directly.
	return cred, nil
}

// JSON key file type.
const (
	serviceAccountKey = "service_account"
)

// credentialsFromJSON returns a google.Credentials from the JSON data
//
// - A self-signed JWT flow will be executed if the following conditions are
// met:
//   (1) At least one of the following is true:
//       (a) No scope is provided
//       (b) Scope for self-signed JWT flow is enabled
//       (c) Audiences are explicitly provided by users
//   (2) No service account impersontation
//
// - Otherwise, executes standard OAuth 2.0 flow
// More details: google.aip.dev/auth/4111
func credentialsFromJSON(ctx context.Context, data []byte, ds *DialSettings) (*google.Credentials, error) {
	// By default, a standard OAuth 2.0 token source is created
	cred, err := google.CredentialsFromJSON(ctx, data, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}

	// Override the token source to use self-signed JWT if conditions are met
	isJWTFlow, err := isSelfSignedJWTFlow(data, ds)
	if err != nil {
		return nil, err
	}
	if isJWTFlow {
		ts, err := selfSignedJWTTokenSource(data, ds)
		if err != nil {
			return nil, err
		}
		cred.TokenSource = ts
	}

	return cred, err
}

func isSelfSignedJWTFlow(data []byte, ds *DialSettings) (bool, error) {
	if (ds.EnableJwtWithScope || ds.HasCustomAudience()) &&
		ds.ImpersonationConfig == nil {
		// Check if JSON is a service account and if so create a self-signed JWT.
		var f struct {
			Type string `json:"type"`
			// The rest JSON fields are omitted because they are not used.
		}
		if err := json.Unmarshal(data, &f); err != nil {
			return false, err
		}
		return f.Type == serviceAccountKey, nil
	}
	return false, nil
}

func selfSignedJWTTokenSource(data []byte, ds *DialSettings) (oauth2.TokenSource, error) {
	if len(ds.GetScopes()) > 0 && !ds.HasCustomAudience() {
		// Scopes are preferred in self-signed JWT unless the scope is not available
		// or a custom audience is used.
		return google.JWTAccessTokenSourceWithScope(data, ds.GetScopes()...)
	} else if ds.GetAudience() != "" {
		// Fallback to audience if scope is not provided
		return google.JWTAccessTokenSourceFromJSON(data, ds.GetAudience())
	} else {
		return nil, errors.New("neither scopes or audience are available for the self-signed JWT")
	}
<<<<<<< HEAD
	return google.JWTAccessTokenSourceFromJSON(data, audience)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	return google.JWTAccessTokenSourceFromJSON(data, audience)
=======
}

// QuotaProjectFromCreds returns the quota project from the JSON blob in the provided credentials.
//
// NOTE(cbro): consider promoting this to a field on google.Credentials.
func QuotaProjectFromCreds(cred *google.Credentials) string {
	var v struct {
		QuotaProject string `json:"quota_project_id"`
	}
	if err := json.Unmarshal(cred.JSON, &v); err != nil {
		return ""
	}
	return v.QuotaProject
}

func impersonateCredentials(ctx context.Context, creds *google.Credentials, ds *DialSettings) (*google.Credentials, error) {
	if len(ds.ImpersonationConfig.Scopes) == 0 {
		ds.ImpersonationConfig.Scopes = ds.GetScopes()
	}
	ts, err := impersonate.TokenSource(ctx, creds.TokenSource, ds.ImpersonationConfig)
	if err != nil {
		return nil, err
	}
	return &google.Credentials{
		TokenSource: ts,
		ProjectID:   creds.ProjectID,
	}, nil
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	"errors"
>>>>>>> 6b7ce455e (update vendored files)
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
	"google.golang.org/api/internal/impersonate"

	"golang.org/x/oauth2/google"
)

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	creds, err := baseCreds(ctx, ds)
	if err != nil {
		return nil, err
	}
	if ds.ImpersonationConfig != nil {
		return impersonateCredentials(ctx, creds, ds)
	}
	return creds, nil
}

func baseCreds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	if ds.InternalCredentials != nil {
		return ds.InternalCredentials, nil
	}
	if ds.Credentials != nil {
		return ds.Credentials, nil
	}
	if ds.CredentialsJSON != nil {
		return credentialsFromJSON(ctx, ds.CredentialsJSON, ds)
	}
	if ds.CredentialsFile != "" {
		data, err := ioutil.ReadFile(ds.CredentialsFile)
		if err != nil {
			return nil, fmt.Errorf("cannot read credentials file: %v", err)
		}
		return credentialsFromJSON(ctx, data, ds)
	}
	if ds.TokenSource != nil {
		return &google.Credentials{TokenSource: ds.TokenSource}, nil
	}
	cred, err := google.FindDefaultCredentials(ctx, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}
	if len(cred.JSON) > 0 {
		return credentialsFromJSON(ctx, cred.JSON, ds)
	}
	// For GAE and GCE, the JSON is empty so return the default credentials directly.
	return cred, nil
}

// JSON key file type.
const (
	serviceAccountKey = "service_account"
)

// credentialsFromJSON returns a google.Credentials from the JSON data
//
// - A self-signed JWT flow will be executed if the following conditions are
// met:
//   (1) At least one of the following is true:
//       (a) No scope is provided
//       (b) Scope for self-signed JWT flow is enabled
//       (c) Audiences are explicitly provided by users
//   (2) No service account impersontation
//
// - Otherwise, executes standard OAuth 2.0 flow
// More details: google.aip.dev/auth/4111
func credentialsFromJSON(ctx context.Context, data []byte, ds *DialSettings) (*google.Credentials, error) {
	// By default, a standard OAuth 2.0 token source is created
	cred, err := google.CredentialsFromJSON(ctx, data, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}

	// Override the token source to use self-signed JWT if conditions are met
	isJWTFlow, err := isSelfSignedJWTFlow(data, ds)
	if err != nil {
		return nil, err
	}
	if isJWTFlow {
		ts, err := selfSignedJWTTokenSource(data, ds)
		if err != nil {
			return nil, err
		}
		cred.TokenSource = ts
	}

	return cred, err
}

func isSelfSignedJWTFlow(data []byte, ds *DialSettings) (bool, error) {
	if (ds.EnableJwtWithScope || ds.HasCustomAudience()) &&
		ds.ImpersonationConfig == nil {
		// Check if JSON is a service account and if so create a self-signed JWT.
		var f struct {
			Type string `json:"type"`
			// The rest JSON fields are omitted because they are not used.
		}
		if err := json.Unmarshal(data, &f); err != nil {
			return false, err
		}
		return f.Type == serviceAccountKey, nil
	}
	return false, nil
}

func selfSignedJWTTokenSource(data []byte, ds *DialSettings) (oauth2.TokenSource, error) {
	if len(ds.GetScopes()) > 0 && !ds.HasCustomAudience() {
		// Scopes are preferred in self-signed JWT unless the scope is not available
		// or a custom audience is used.
		return google.JWTAccessTokenSourceWithScope(data, ds.GetScopes()...)
	} else if ds.GetAudience() != "" {
		// Fallback to audience if scope is not provided
		return google.JWTAccessTokenSourceFromJSON(data, ds.GetAudience())
	} else {
		return nil, errors.New("neither scopes or audience are available for the self-signed JWT")
	}
<<<<<<< HEAD
	return google.JWTAccessTokenSourceFromJSON(data, audience)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	return google.JWTAccessTokenSourceFromJSON(data, audience)
=======
}

// QuotaProjectFromCreds returns the quota project from the JSON blob in the provided credentials.
//
// NOTE(cbro): consider promoting this to a field on google.Credentials.
func QuotaProjectFromCreds(cred *google.Credentials) string {
	var v struct {
		QuotaProject string `json:"quota_project_id"`
	}
	if err := json.Unmarshal(cred.JSON, &v); err != nil {
		return ""
	}
	return v.QuotaProject
}

func impersonateCredentials(ctx context.Context, creds *google.Credentials, ds *DialSettings) (*google.Credentials, error) {
	if len(ds.ImpersonationConfig.Scopes) == 0 {
		ds.ImpersonationConfig.Scopes = ds.GetScopes()
	}
	ts, err := impersonate.TokenSource(ctx, creds.TokenSource, ds.ImpersonationConfig)
	if err != nil {
		return nil, err
	}
	return &google.Credentials{
		TokenSource: ts,
		ProjectID:   creds.ProjectID,
	}, nil
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"errors"
>>>>>>> 4d7e5ad26 (update vendored files)
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"
	"google.golang.org/api/internal/impersonate"

	"golang.org/x/oauth2/google"
)

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	creds, err := baseCreds(ctx, ds)
	if err != nil {
		return nil, err
	}
	if ds.ImpersonationConfig != nil {
		return impersonateCredentials(ctx, creds, ds)
	}
	return creds, nil
}

func baseCreds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	if ds.InternalCredentials != nil {
		return ds.InternalCredentials, nil
	}
	if ds.Credentials != nil {
		return ds.Credentials, nil
	}
	if ds.CredentialsJSON != nil {
		return credentialsFromJSON(ctx, ds.CredentialsJSON, ds)
	}
	if ds.CredentialsFile != "" {
		data, err := ioutil.ReadFile(ds.CredentialsFile)
		if err != nil {
			return nil, fmt.Errorf("cannot read credentials file: %v", err)
		}
		return credentialsFromJSON(ctx, data, ds)
	}
	if ds.TokenSource != nil {
		return &google.Credentials{TokenSource: ds.TokenSource}, nil
	}
	cred, err := google.FindDefaultCredentials(ctx, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}
	if len(cred.JSON) > 0 {
		return credentialsFromJSON(ctx, cred.JSON, ds)
	}
	// For GAE and GCE, the JSON is empty so return the default credentials directly.
	return cred, nil
}

// JSON key file type.
const (
	serviceAccountKey = "service_account"
)

// credentialsFromJSON returns a google.Credentials from the JSON data
//
// - A self-signed JWT flow will be executed if the following conditions are
// met:
//   (1) At least one of the following is true:
//       (a) No scope is provided
//       (b) Scope for self-signed JWT flow is enabled
//       (c) Audiences are explicitly provided by users
//   (2) No service account impersontation
//
// - Otherwise, executes standard OAuth 2.0 flow
// More details: google.aip.dev/auth/4111
func credentialsFromJSON(ctx context.Context, data []byte, ds *DialSettings) (*google.Credentials, error) {
	// By default, a standard OAuth 2.0 token source is created
	cred, err := google.CredentialsFromJSON(ctx, data, ds.GetScopes()...)
	if err != nil {
		return nil, err
	}

	// Override the token source to use self-signed JWT if conditions are met
	isJWTFlow, err := isSelfSignedJWTFlow(data, ds)
	if err != nil {
		return nil, err
	}
	if isJWTFlow {
		ts, err := selfSignedJWTTokenSource(data, ds)
		if err != nil {
			return nil, err
		}
		cred.TokenSource = ts
	}

	return cred, err
}

func isSelfSignedJWTFlow(data []byte, ds *DialSettings) (bool, error) {
	if (ds.EnableJwtWithScope || ds.HasCustomAudience()) &&
		ds.ImpersonationConfig == nil {
		// Check if JSON is a service account and if so create a self-signed JWT.
		var f struct {
			Type string `json:"type"`
			// The rest JSON fields are omitted because they are not used.
		}
		if err := json.Unmarshal(data, &f); err != nil {
			return false, err
		}
		return f.Type == serviceAccountKey, nil
	}
	return false, nil
}

func selfSignedJWTTokenSource(data []byte, ds *DialSettings) (oauth2.TokenSource, error) {
	if len(ds.GetScopes()) > 0 && !ds.HasCustomAudience() {
		// Scopes are preferred in self-signed JWT unless the scope is not available
		// or a custom audience is used.
		return google.JWTAccessTokenSourceWithScope(data, ds.GetScopes()...)
	} else if ds.GetAudience() != "" {
		// Fallback to audience if scope is not provided
		return google.JWTAccessTokenSourceFromJSON(data, ds.GetAudience())
	} else {
		return nil, errors.New("neither scopes or audience are available for the self-signed JWT")
	}
<<<<<<< HEAD
	return google.JWTAccessTokenSourceFromJSON(data, audience)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	return google.JWTAccessTokenSourceFromJSON(data, audience)
=======
}

// QuotaProjectFromCreds returns the quota project from the JSON blob in the provided credentials.
//
// NOTE(cbro): consider promoting this to a field on google.Credentials.
func QuotaProjectFromCreds(cred *google.Credentials) string {
	var v struct {
		QuotaProject string `json:"quota_project_id"`
	}
	if err := json.Unmarshal(cred.JSON, &v); err != nil {
		return ""
	}
	return v.QuotaProject
}

func impersonateCredentials(ctx context.Context, creds *google.Credentials, ds *DialSettings) (*google.Credentials, error) {
	if len(ds.ImpersonationConfig.Scopes) == 0 {
		ds.ImpersonationConfig.Scopes = ds.GetScopes()
	}
	ts, err := impersonate.TokenSource(ctx, creds.TokenSource, ds.ImpersonationConfig)
	if err != nil {
		return nil, err
	}
	return &google.Credentials{
		TokenSource: ts,
		ProjectID:   creds.ProjectID,
	}, nil
>>>>>>> 4d7e5ad26 (update vendored files)
}

// oauth2DialSettings returns the settings to be used by the OAuth2 transport, which is separate from the API transport.
func oauth2DialSettings(ds *DialSettings) *DialSettings {
	var ods DialSettings
	ods.DefaultEndpoint = google.Endpoint.TokenURL
	ods.DefaultMTLSEndpoint = google.MTLSTokenURL
	ods.ClientCertSource = ds.ClientCertSource
	return &ods
}

// customHTTPClient constructs an HTTPClient using the provided tlsConfig, to support mTLS.
func customHTTPClient(tlsConfig *tls.Config) *http.Client {
	trans := baseTransport()
	trans.TLSClientConfig = tlsConfig
	return &http.Client{Transport: trans}
}

func baseTransport() *http.Transport {
	return &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"encoding/json"
	"fmt"
	"io/ioutil"

	"golang.org/x/oauth2"

	"golang.org/x/oauth2/google"
)

// Creds returns credential information obtained from DialSettings, or if none, then
// it returns default credential information.
func Creds(ctx context.Context, ds *DialSettings) (*google.Credentials, error) {
	if ds.Credentials != nil {
		return ds.Credentials, nil
	}
	if ds.CredentialsJSON != nil {
		return credentialsFromJSON(ctx, ds.CredentialsJSON, ds.Endpoint, ds.Scopes, ds.Audiences)
	}
	if ds.CredentialsFile != "" {
		data, err := ioutil.ReadFile(ds.CredentialsFile)
		if err != nil {
			return nil, fmt.Errorf("cannot read credentials file: %v", err)
		}
		return credentialsFromJSON(ctx, data, ds.Endpoint, ds.Scopes, ds.Audiences)
	}
	if ds.TokenSource != nil {
		return &google.Credentials{TokenSource: ds.TokenSource}, nil
	}
	cred, err := google.FindDefaultCredentials(ctx, ds.Scopes...)
	if err != nil {
		return nil, err
	}
	if len(cred.JSON) > 0 {
		return credentialsFromJSON(ctx, cred.JSON, ds.Endpoint, ds.Scopes, ds.Audiences)
	}
	// For GAE and GCE, the JSON is empty so return the default credentials directly.
	return cred, nil
}

// JSON key file type.
const (
	serviceAccountKey = "service_account"
)

// credentialsFromJSON returns a google.Credentials based on the input.
//
// - If the JSON is a service account and no scopes provided, returns self-signed JWT auth flow
// - Otherwise, returns OAuth 2.0 flow.
func credentialsFromJSON(ctx context.Context, data []byte, endpoint string, scopes []string, audiences []string) (*google.Credentials, error) {
	cred, err := google.CredentialsFromJSON(ctx, data, scopes...)
	if err != nil {
		return nil, err
	}
	if len(data) > 0 && len(scopes) == 0 {
		var f struct {
			Type string `json:"type"`
			// The rest JSON fields are omitted because they are not used.
		}
		if err := json.Unmarshal(cred.JSON, &f); err != nil {
			return nil, err
		}
		if f.Type == serviceAccountKey {
			ts, err := selfSignedJWTTokenSource(data, endpoint, audiences)
			if err != nil {
				return nil, err
			}
			cred.TokenSource = ts
		}
	}
	return cred, err
}

func selfSignedJWTTokenSource(data []byte, endpoint string, audiences []string) (oauth2.TokenSource, error) {
	// Use the API endpoint as the default audience
	audience := endpoint
	if len(audiences) > 0 {
		// TODO(shinfan): Update golang oauth to support multiple audiences.
		if len(audiences) > 1 {
			return nil, fmt.Errorf("multiple audiences support is not implemented")
		}
		audience = audiences[0]
	}
	return google.JWTAccessTokenSourceFromJSON(data, audience)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
