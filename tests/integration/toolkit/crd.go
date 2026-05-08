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

package toolkit

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
	"sigs.k8s.io/external-dns/internal/testutils"
)

const (
	dnsEndpointListKind = "DNSEndpointList"
	dnsEndpointResource = "dnsendpoints"
	initialEventsEndKey = "k8s.io/initial-events-end"
)

// crdClientGenerator wraps a StubClientGenerator and overrides RESTConfig to
// return a fake API server's config, enabling the CRD source's controller-runtime
// cache to connect without a real cluster.
type crdClientGenerator struct {
	testutils.StubClientGenerator
	restCfg *rest.Config
}

func (g crdClientGenerator) RESTConfig() (*rest.Config, error) {
	return g.restCfg, nil
}

// newCRDClientGenerator builds a crdClientGenerator backed by the given fake
// Kubernetes clientset (for non-CRD sources) and the REST config of the fake
// CRD API server.
func newCRDClientGenerator(k8sClient *fake.Clientset, restCfg *rest.Config) crdClientGenerator {
	return crdClientGenerator{
		StubClientGenerator: testutils.NewFakeClientGenerator(k8sClient),
		restCfg:             restCfg,
	}
}

// newFakeDNSEndpointServer starts a minimal in-process HTTP server that serves
// just enough of the Kubernetes API for the CRD source's controller-runtime cache
// to initialize and sync:
//
//   - /api, /apis, /apis/externaldns.k8s.io — API discovery
//   - /apis/externaldns.k8s.io/v1alpha1 — APIResourceList
//   - /apis/externaldns.k8s.io/v1alpha1/dnsendpoints — List + Watch
//   - /apis/externaldns.k8s.io/v1alpha1/namespaces/{ns}/dnsendpoints — same
//
// The server is closed automatically when ctx is canceled.
func newFakeDNSEndpointServer(ctx context.Context, dnsEndpoints []*apiv1alpha1.DNSEndpoint) *rest.Config {
	h := &dnsEndpointHandler{endpoints: dnsEndpoints}
	srv := httptest.NewServer(h)
	go func() {
		<-ctx.Done()
		srv.Close()
	}()
	return &rest.Config{Host: srv.URL}
}

// dnsEndpointHandler is the HTTP handler for the fake Kubernetes API server.
type dnsEndpointHandler struct {
	mu        sync.Mutex
	endpoints []*apiv1alpha1.DNSEndpoint
}

func (h *dnsEndpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	isWatch := r.URL.Query().Get("watch") == "true"

	switch {
	case path == "/api":
		h.serveAPIVersions(w)
	case path == "/apis":
		h.serveAPIGroupList(w)
	case path == "/apis/"+apiv1alpha1.GroupVersion.Group:
		h.serveAPIGroup(w)
	case path == "/apis/"+apiv1alpha1.GroupVersion.String():
		h.serveAPIResourceList(w)
	case isDNSEndpointPath(path) && isWatch:
		h.serveWatch(w, r)
	case isDNSEndpointPath(path):
		h.serveList(w, r)
	default:
		http.NotFound(w, r)
	}
}

func isDNSEndpointPath(path string) bool {
	return strings.HasPrefix(path, "/apis/"+apiv1alpha1.GroupVersion.String()) && 
	       strings.HasSuffix(path, "/"+dnsEndpointResource)
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *dnsEndpointHandler) serveAPIVersions(w http.ResponseWriter) {
	writeJSON(w, &metav1.APIVersions{
		TypeMeta:                   metav1.TypeMeta{APIVersion: "v1", Kind: "APIVersions"},
		ServerAddressByClientCIDRs: []metav1.ServerAddressByClientCIDR{},
		Versions:                   []string{},
	})
}

func (h *dnsEndpointHandler) serveAPIGroupList(w http.ResponseWriter) {
	gv := metav1.GroupVersionForDiscovery{
		GroupVersion: apiv1alpha1.GroupVersion.String(),
		Version:      apiv1alpha1.GroupVersion.Version,
	}
	writeJSON(w, &metav1.APIGroupList{
		TypeMeta: metav1.TypeMeta{APIVersion: "v1", Kind: "APIGroupList"},
		Groups: []metav1.APIGroup{
			{
				Name:             apiv1alpha1.GroupVersion.Group,
				Versions:         []metav1.GroupVersionForDiscovery{gv},
				PreferredVersion: gv,
			},
		},
	})
}

func (h *dnsEndpointHandler) serveAPIGroup(w http.ResponseWriter) {
	gv := metav1.GroupVersionForDiscovery{
		GroupVersion: apiv1alpha1.GroupVersion.String(),
		Version:      apiv1alpha1.GroupVersion.Version,
	}
	writeJSON(w, &metav1.APIGroup{
		TypeMeta:         metav1.TypeMeta{APIVersion: "v1", Kind: "APIGroup"},
		Name:             apiv1alpha1.GroupVersion.Group,
		Versions:         []metav1.GroupVersionForDiscovery{gv},
		PreferredVersion: gv,
	})
}

func (h *dnsEndpointHandler) serveAPIResourceList(w http.ResponseWriter) {
	writeJSON(w, &metav1.APIResourceList{
		TypeMeta:     metav1.TypeMeta{APIVersion: "v1", Kind: "APIResourceList"},
		GroupVersion: apiv1alpha1.GroupVersion.String(),
		APIResources: []metav1.APIResource{
			{
				Name:         dnsEndpointResource,
				SingularName: "dnsendpoint",
				Namespaced:   true,
				Kind:         apiv1alpha1.DNSEndpointKind,
				Verbs:        metav1.Verbs{"get", "list", "watch"},
			},
		},
	})
}

// filteredEndpoints returns copies of all endpoints matching the given namespace
// (or all endpoints if namespace is empty), with TypeMeta set for API responses.
func (h *dnsEndpointHandler) filteredEndpoints(ns string) []apiv1alpha1.DNSEndpoint {
	h.mu.Lock()
	all := h.endpoints
	h.mu.Unlock()

	result := make([]apiv1alpha1.DNSEndpoint, 0, len(all))
	for _, ep := range all {
		if ns != "" && ep.Namespace != ns {
			continue
		}
		item := *ep
		item.TypeMeta = metav1.TypeMeta{
			APIVersion: apiv1alpha1.GroupVersion.String(),
			Kind:       apiv1alpha1.DNSEndpointKind,
		}
		result = append(result, item)
	}
	return result
}

func (h *dnsEndpointHandler) serveList(w http.ResponseWriter, r *http.Request) {
	items := h.filteredEndpoints(extractNamespace(r.URL.Path))
	writeJSON(w, &apiv1alpha1.DNSEndpointList{
		TypeMeta: metav1.TypeMeta{APIVersion: apiv1alpha1.GroupVersion.String(), Kind: dnsEndpointListKind},
		ListMeta: metav1.ListMeta{ResourceVersion: "1"},
		Items:    items,
	})
}

// bookmarkEvent builds the BOOKMARK watch event object. When initialEventsEnd
// is true it carries the k8s.io/initial-events-end annotation required by the
// watchList protocol so the reflector marks the informer as synced.
func bookmarkEvent(initialEventsEnd bool) map[string]any {
	meta := map[string]any{"resourceVersion": "1"}
	if initialEventsEnd {
		meta["annotations"] = map[string]string{initialEventsEndKey: "true"}
	}
	return map[string]any{
		"type": "BOOKMARK",
		"object": map[string]any{
			"apiVersion": apiv1alpha1.GroupVersion.String(),
			"kind":       apiv1alpha1.DNSEndpointKind,
			"metadata":   meta,
		},
	}
}

// serveWatch streams watch events so the informer's initial sync completes.
//
// client-go v0.35+ uses "watchList": the reflector sends a watch request with
// sendInitialEvents=true and waits for a BOOKMARK carrying the annotation
// k8s.io/initial-events-end: "true" before it marks the informer as synced.
func (h *dnsEndpointHandler) serveWatch(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.WriteHeader(http.StatusOK)

	flusher, ok := w.(http.Flusher)
	if !ok {
		return
	}

	enc := json.NewEncoder(w)
	sendInitialEvents := r.URL.Query().Get("sendInitialEvents") == "true"

	if sendInitialEvents {
		for _, item := range h.filteredEndpoints(extractNamespace(r.URL.Path)) {
			if err := enc.Encode(map[string]any{"type": "ADDED", "object": item}); err != nil {
				return
			}
		}
	}
	if err := enc.Encode(bookmarkEvent(sendInitialEvents)); err != nil {
		return
	}
	flusher.Flush()

	// Block until the client closes the connection.
	<-r.Context().Done()
}

func extractNamespace(path string) string {
	const prefix = "/apis/externaldns.k8s.io/v1alpha1/namespaces/"
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	tail := strings.TrimPrefix(path, prefix)
	if idx := strings.Index(tail, "/"); idx > 0 {
		return tail[:idx]
	}
	return ""
}
