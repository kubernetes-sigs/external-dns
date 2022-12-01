/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/Azure/go-autorest/autorest/adal"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/assert"
)

func TestGetAzureEnvironmentConfig(t *testing.T) {
	tmp, err := ioutil.TempFile("", "azureconf")
	if err != nil {
		t.Errorf("couldn't write temp file %v", err)
	}
	defer os.Remove(tmp.Name())

	tests := map[string]struct {
		cloud string
		err   error
	}{
		"AzureChinaCloud":   {"AzureChinaCloud", nil},
		"AzureGermanCloud":  {"AzureGermanCloud", nil},
		"AzurePublicCloud":  {"", nil},
		"AzureUSGovernment": {"AzureUSGovernmentCloud", nil},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, _ = tmp.Seek(0, 0)
			_, _ = tmp.Write([]byte(fmt.Sprintf(`{"cloud": "%s"}`, test.cloud)))
			got, err := getConfig(tmp.Name(), "", "")
			if err != nil {
				t.Errorf("got unexpected err %v", err)
			}

			if test.cloud == "" {
				test.cloud = "AzurePublicCloud"
			}
			want, err := azure.EnvironmentFromName(test.cloud)
			if err != nil {
				t.Errorf("couldn't get azure environment from provided name %v", err)
			}

			if !reflect.DeepEqual(want, got.Environment) {
				t.Errorf("got %v, want %v", got.Environment, want)
			}
		})
	}
}

func populateFederatedToken(t *testing.T, filename string, content string) {
	t.Helper()

	f, err := os.Create(filename)
	if err != nil {
		assert.FailNow(t, err.Error())
	}

	if _, err := io.WriteString(f, content); err != nil {
		assert.FailNow(t, err.Error())
	}

	if err := f.Close(); err != nil {
		assert.FailNow(t, err.Error())
	}
}

func TestGetAccessTokenWorkloadIdentity(t *testing.T) {
	// Create a file that will be used to store a federated token
	f, err := os.CreateTemp("", "")
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	defer os.Remove(f.Name())

	// Close the file to simplify logic within populateFederatedToken helper
	if err := f.Close(); err != nil {
		assert.FailNow(t, err.Error())
	}

	// The initial federated token is never used, so we don't care about the value yet
	// Though, it's a requirement from adal to have a non-empty value set
	populateFederatedToken(t, f.Name(), "random-jwt")

	// Envs are described here: https://azure.github.io/azure-workload-identity/docs/installation/mutating-admission-webhook.html
	t.Setenv("AZURE_TENANT_ID", "fakeTenantID")
	t.Setenv("AZURE_CLIENT_ID", "fakeClientID")
	t.Setenv("AZURE_FEDERATED_TOKEN_FILE", f.Name())

	t.Run("token refresh", func(t *testing.T) {
		// Basically, we want one token to be exchanged for the other (key and value respectively)
		tokens := map[string]string{
			"initialFederatedToken":   "initialAccessToken",
			"refreshedFederatedToken": "refreshedAccessToken",
		}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				assert.FailNow(t, err.Error())
			}

			w.Header().Set("Content-Type", "application/json")
			receivedFederatedToken := r.FormValue("client_assertion")
			accessToken := adal.Token{AccessToken: tokens[receivedFederatedToken]}

			if err := json.NewEncoder(w).Encode(accessToken); err != nil {
				assert.FailNow(t, err.Error())
			}

			// Expected format: http://<server>/<tenant-ID>/oauth2/token?api-version=1.0
			assert.Contains(t, r.RequestURI, os.Getenv("AZURE_TENANT_ID"), "URI should contain the tenant ID exposed through env variable")

			assert.Equal(t, os.Getenv("AZURE_CLIENT_ID"), r.FormValue("client_id"), "client_id should match the value exposed through env variable")
		}))
		defer ts.Close()

		env := azure.Environment{ActiveDirectoryEndpoint: ts.URL, ResourceManagerEndpoint: ts.URL}

		cfg := config{
			UseWorkloadIdentityExtension: true,
		}

		token, err := getAccessToken(cfg, env)
		assert.NoError(t, err)

		for federatedToken, accessToken := range tokens {
			populateFederatedToken(t, f.Name(), federatedToken)
			assert.NoError(t, token.Refresh(), "Token refresh failed")
			assert.Equal(t, accessToken, token.Token().AccessToken, "Access token should have been set to a value returned by the webserver")
		}
	})

	t.Run("clientID overrides through UserAssignedIdentityID section", func(t *testing.T) {
		cfg := config{
			UseWorkloadIdentityExtension: true,
			UserAssignedIdentityID:       "overridenClientID",
		}

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				assert.FailNow(t, err.Error())
			}

			w.Header().Set("Content-Type", "application/json")
			accessToken := adal.Token{AccessToken: "abc"}

			if err := json.NewEncoder(w).Encode(accessToken); err != nil {
				assert.FailNow(t, err.Error())
			}

			assert.Equal(t, cfg.UserAssignedIdentityID, r.FormValue("client_id"), "client_id should match the value passed through managedIdentity section")
		}))
		defer ts.Close()

		env := azure.Environment{ActiveDirectoryEndpoint: ts.URL, ResourceManagerEndpoint: ts.URL}

		token, err := getAccessToken(cfg, env)
		assert.NoError(t, err)

		assert.NoError(t, token.Refresh(), "Token refresh failed")
	})
}
