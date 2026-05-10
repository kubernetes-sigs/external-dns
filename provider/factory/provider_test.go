/*
Copyright 2026 The Kubernetes Authors.

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

package factory

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
)

func TestSelectProvider(t *testing.T) {
	tests := []struct {
		name          string
		cfg           *externaldns.Config
		expectedType  string
		expectedError string
	}{
		{
			name: "aws provider",
			cfg: &externaldns.Config{
				Provider: externaldns.ProviderAWS,
			},
			expectedType: "*aws.AWSProvider",
		},
		{
			name: "rfc2136 provider",
			cfg: &externaldns.Config{
				Provider:             externaldns.ProviderRFC2136,
				RFC2136TSIGSecretAlg: "hmac-sha256",
			},
			expectedType: "*rfc2136.rfc2136Provider",
		},
		{
			name: "gandi provider",
			cfg: &externaldns.Config{
				Provider: externaldns.ProviderGandi,
			},
			expectedError: "no environment variable GANDI_KEY or GANDI_PAT provided",
		},
		{
			name: "inmemory provider",
			cfg: &externaldns.Config{
				Provider: externaldns.ProviderInMemory,
			},
			expectedType: "*inmemory.InMemoryProvider",
		},
		{
			name: "oci provider instance principal without compartment OCID",
			cfg: &externaldns.Config{
				Provider:                 externaldns.ProviderOCI,
				OCIAuthInstancePrincipal: true,
				OCICompartmentOCID:       "",
			},
			expectedError: "instance principal authentication requested, but no compartment OCID provided",
		},
		{
			name: "oci provider without config file",
			cfg: &externaldns.Config{
				Provider:      externaldns.ProviderOCI,
				OCIConfigFile: "",
			},
			expectedError: "reading OCI config file",
		},
		{
			name: "coredns provider",
			cfg: &externaldns.Config{
				Provider: externaldns.ProviderCoreDNS,
			},
			expectedType: "coredns.coreDNSProvider",
		},
		{
			name: "pihole provider",
			cfg: &externaldns.Config{
				Provider:     externaldns.ProviderPihole,
				PiholeServer: "http://localhost:8080",
			},
			expectedType: "*pihole.PiholeProvider",
		},
		{
			name: "dnsimple provider",
			cfg: &externaldns.Config{
				Provider: externaldns.ProviderDNSimple,
			},
			expectedError: "no dnsimple oauth token provided",
		},
		{
			name: "unknown provider",
			cfg: &externaldns.Config{
				Provider: "unknown",
			},
			expectedError: "unknown dns provider: unknown",
		},
		{
			name: "inmemory cached provider",
			cfg: &externaldns.Config{
				Provider:          externaldns.ProviderInMemory,
				ProviderCacheTime: 10 * time.Millisecond,
			},
			expectedType: "*provider.CachedProvider",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			domainFilter := endpoint.NewDomainFilter([]string{"example.com"})

			p, err := Select(t.Context(), tt.cfg, domainFilter)

			if tt.expectedError != "" {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				require.NotNil(t, p)
				assert.Contains(t, reflect.TypeOf(p).String(), tt.expectedType)
			}
		})
	}
}

func TestKnownProviders(t *testing.T) {
	names := make([]string, 0, len(externaldns.ProviderNames))
	for _, name := range externaldns.ProviderNames {
		t.Run(name, func(t *testing.T) {
			names = append(names, name)
			_, ok := providers(name)
			assert.True(t, ok, "expected provider %s to be registered", name)
		})
	}
	assert.ElementsMatch(t, externaldns.ProviderNames, names)
}

func TestSelectProvider_Webhook(t *testing.T) {
	// Stand up a minimal HTTP server that returns a valid negotiation response.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/external.dns.webhook+json;version=1")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	cfg := &externaldns.Config{
		Provider:           externaldns.ProviderWebhook,
		WebhookProviderURL: srv.URL,
	}
	p, err := Select(t.Context(), cfg, nil)
	require.NoError(t, err)
	require.NotNil(t, p)
}
