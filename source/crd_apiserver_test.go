/*
Copyright 2018 The Kubernetes Authors.

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

package source

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

// TestCRDSourceAnnotationFilterAgainstAPIServer runs NewCRDSource against a fake API
// server, so informer and cache behave as in a cluster. The fake cache in crd_test.go is
// populated directly and misses what happens between the LIST response and the cache (#6728).
func TestCRDSourceAnnotationFilterAgainstAPIServer(t *testing.T) {
	matching := newFilterTestDNSEndpoint("matching", "matching.example.com", "192.0.2.1", map[string]string{"example.com/filter": "yes"})
	other := newFilterTestDNSEndpoint("other", "other.example.com", "192.0.2.2", map[string]string{"example.com/filter": "no"})
	unannotated := newFilterTestDNSEndpoint("unannotated", "unannotated.example.com", "192.0.2.3", nil)

	for _, tt := range []struct {
		title    string
		items    []apiv1alpha1.DNSEndpoint
		expected []string
	}{
		{
			title:    "only matching DNSEndpoints exist",
			items:    []apiv1alpha1.DNSEndpoint{matching},
			expected: []string{"matching.example.com"},
		},
		{
			title:    "a non-matching DNSEndpoint is listed first",
			items:    []apiv1alpha1.DNSEndpoint{other, matching},
			expected: []string{"matching.example.com"},
		},
		{
			title:    "a non-matching DNSEndpoint is listed last",
			items:    []apiv1alpha1.DNSEndpoint{matching, other},
			expected: []string{"matching.example.com"},
		},
		{
			title:    "an unannotated DNSEndpoint is listed alongside a matching one",
			items:    []apiv1alpha1.DNSEndpoint{unannotated, matching},
			expected: []string{"matching.example.com"},
		},
		{
			title:    "no DNSEndpoint matches",
			items:    []apiv1alpha1.DNSEndpoint{other, unannotated},
			expected: nil,
		},
	} {
		t.Run(tt.title, func(t *testing.T) {
			restConfig := startFakeDNSEndpointAPIServer(t, tt.items)

			selector, err := annotations.ParseFilter("example.com/filter=yes")
			require.NoError(t, err)

			ctx, cancel := context.WithTimeout(t.Context(), 30*time.Second)
			defer cancel()

			source, err := NewCRDSource(ctx, restConfig, &Config{AnnotationFilter: selector})
			require.NoError(t, err)

			endpoints, err := source.Endpoints(ctx)
			require.NoError(t, err)

			var got []string
			for _, ep := range endpoints {
				got = append(got, ep.DNSName)
			}
			assert.ElementsMatch(t, tt.expected, got)
		})
	}
}

// newFilterTestDNSEndpoint builds a DNSEndpoint holding a single A record.
func newFilterTestDNSEndpoint(name, dnsName, target string, anns map[string]string) apiv1alpha1.DNSEndpoint {
	return apiv1alpha1.DNSEndpoint{
		APIVersion: apiv1alpha1.GroupVersion.String(), Kind: "DNSEndpoint",
		Name:            name,
		Namespace:       "default",
		ResourceVersion: "1",
		Annotations:     anns,
		Spec: apiv1alpha1.DNSEndpointSpec{
			Endpoints: []*endpoint.Endpoint{{
				DNSName:    dnsName,
				RecordType: endpoint.RecordTypeA,
				RecordTTL:  60,
				Targets:    endpoint.Targets{target},
			}},
		},
	}
}

// startFakeDNSEndpointAPIServer serves the discovery and DNSEndpoint list/watch endpoints
// a controller-runtime cache needs. The watch stays open until the test finishes.
func startFakeDNSEndpointAPIServer(t *testing.T, items []apiv1alpha1.DNSEndpoint) *rest.Config {
	t.Helper()

	gv := apiv1alpha1.GroupVersion
	writeJSON := func(w http.ResponseWriter, payload any) {
		w.Header().Set("Content-Type", "application/json")
		require.NoError(t, json.NewEncoder(w).Encode(payload))
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/api":
			writeJSON(w, &metav1.APIVersions{Versions: []string{}})
		case r.URL.Path == "/apis":
			writeJSON(w, &metav1.APIGroupList{Groups: []metav1.APIGroup{{
				Name:             gv.Group,
				Versions:         []metav1.GroupVersionForDiscovery{{GroupVersion: gv.String(), Version: gv.Version}},
				PreferredVersion: metav1.GroupVersionForDiscovery{GroupVersion: gv.String(), Version: gv.Version},
			}}})
		case r.URL.Path == "/apis/"+gv.String():
			writeJSON(w, &metav1.APIResourceList{GroupVersion: gv.String(), APIResources: []metav1.APIResource{{
				Name:       "dnsendpoints",
				Namespaced: true,
				Kind:       "DNSEndpoint",
				Verbs:      metav1.Verbs{"get", "list", "watch"},
			}}})
		case r.URL.Query().Get("watch") == "true":
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
			<-r.Context().Done()
		default:
			writeJSON(w, &apiv1alpha1.DNSEndpointList{
				APIVersion: gv.String(), Kind: "DNSEndpointList",
				ResourceVersion: "1",
				Items:           items,
			})
		}
	}))
	t.Cleanup(srv.Close)

	return &rest.Config{Host: srv.URL}
}
