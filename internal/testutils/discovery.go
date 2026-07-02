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

package testutils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
)

// NewFakeExternalDNSDiscoveryServer starts an httptest.Server advertising the given resources
// under externaldns.k8s.io — enough for crcache.New and the informer to start. Any other path
// returns 500 so LIST fails and the cache never syncs, letting callers exercise sync-failure
// paths without a real API server.
func NewFakeExternalDNSDiscoveryServer(t *testing.T, resources ...metav1.APIResource) *httptest.Server {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		encode := func(v any) {
			if err := json.NewEncoder(w).Encode(v); err != nil {
				t.Errorf("fakeDiscoveryServer: json.Encode %s: %v", r.URL.Path, err)
			}
		}
		switch r.URL.Path {
		case "/api":
			encode(metav1.APIVersions{
				TypeMeta: metav1.TypeMeta{Kind: "APIVersions", APIVersion: "v1"},
				Versions: []string{"v1"},
			})
		case "/apis":
			encode(metav1.APIGroupList{
				TypeMeta: metav1.TypeMeta{Kind: "APIGroupList", APIVersion: "v1"},
				Groups: []metav1.APIGroup{{
					Name:             apiv1alpha1.GroupVersion.Group,
					Versions:         []metav1.GroupVersionForDiscovery{{GroupVersion: apiv1alpha1.GroupVersion.String(), Version: apiv1alpha1.GroupVersion.Version}},
					PreferredVersion: metav1.GroupVersionForDiscovery{GroupVersion: apiv1alpha1.GroupVersion.String(), Version: apiv1alpha1.GroupVersion.Version},
				}},
			})
		case "/apis/" + apiv1alpha1.GroupVersion.String():
			encode(metav1.APIResourceList{
				TypeMeta:     metav1.TypeMeta{Kind: "APIResourceList", APIVersion: "v1"},
				GroupVersion: apiv1alpha1.GroupVersion.String(),
				APIResources: resources,
			})
		default:
			// Causes the informer's LIST to fail so the cache never syncs.
			http.Error(w, `{"kind":"Status","apiVersion":"v1","status":"Failure","code":500}`, http.StatusInternalServerError)
		}
	}))
	t.Cleanup(srv.Close)
	return srv
}
