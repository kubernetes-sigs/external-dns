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

package targetfrom

import (
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/dynamic"
	fakedynamic "k8s.io/client-go/dynamic/fake"
	clienttesting "k8s.io/client-go/testing"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source/annotations"
)

var (
	gatewayGVR = schema.GroupVersionResource{Group: "gateway.networking.k8s.io", Version: "v1", Resource: "gateways"}
	serviceGVR = schema.GroupVersionResource{Version: "v1", Resource: "services"}
	gaGVR      = schema.GroupVersionResource{Group: "aga.k8s.aws", Version: "v1beta1", Resource: "globalaccelerators"}
)

const addressesTmpl = "{{ range .status.addresses }}{{ .value }},{{ end }}"

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		value   string
		want    Reference
		wantErr bool
	}{
		{
			value: "gateways.gateway.networking.k8s.io/istio-ingress/central#" + addressesTmpl,
			want: Reference{
				Resource:  schema.GroupResource{Group: "gateway.networking.k8s.io", Resource: "gateways"},
				Namespace: "istio-ingress", Name: "central", Template: addressesTmpl,
			},
		},
		{
			value: "services/shared-lb#{{ .spec.clusterIP }}",
			want:  Reference{Resource: schema.GroupResource{Resource: "services"}, Name: "shared-lb", Template: "{{ .spec.clusterIP }}"},
		},
		{
			// Only the first '#' splits; the template keeps any later '#' and '/'.
			value: `services/a#{{ "x#y/z" }}`,
			want:  Reference{Resource: schema.GroupResource{Resource: "services"}, Name: "a", Template: `{{ "x#y/z" }}`},
		},
		{value: "services/a", wantErr: true},
		{value: "services/a#", wantErr: true},
		{value: "services/a#  ", wantErr: true},
		{value: "a#{{ . }}", wantErr: true},
		{value: "/a#{{ . }}", wantErr: true},
		{value: "services/#{{ . }}", wantErr: true},
		{value: "services//a#{{ . }}", wantErr: true},
		{value: "services/ns/a/b#{{ . }}", wantErr: true},
	} {
		t.Run(tt.value, func(t *testing.T) {
			got, err := Parse(tt.value)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func newObject(gvr schema.GroupVersionResource, kind, namespace, name string, status map[string]any) *unstructured.Unstructured {
	u := &unstructured.Unstructured{Object: map[string]any{"status": status}}
	u.SetAPIVersion(gvr.GroupVersion().String())
	u.SetKind(kind)
	u.SetNamespace(namespace)
	u.SetName(name)
	return u
}

// newTestResolver creates objs through their GVR: the fake client's tracker
// guesses resources from kinds, and turns Gateway into "gatewaies".
func newTestResolver(t *testing.T, objs ...*unstructured.Unstructured) (*Resolver, *fakedynamic.FakeDynamicClient) {
	t.Helper()
	disc := &fakediscovery.FakeDiscovery{Fake: &clienttesting.Fake{Resources: []*metav1.APIResourceList{
		{GroupVersion: "v1", APIResources: []metav1.APIResource{{Name: "services", Kind: "Service", Namespaced: true}}},
		{GroupVersion: "gateway.networking.k8s.io/v1", APIResources: []metav1.APIResource{{Name: "gateways", Kind: "Gateway", Namespaced: true}}},
		{GroupVersion: "aga.k8s.aws/v1beta1", APIResources: []metav1.APIResource{{Name: "globalaccelerators", Kind: "GlobalAccelerator"}}},
	}}}
	client := fakedynamic.NewSimpleDynamicClientWithCustomListKinds(runtime.NewScheme(), map[schema.GroupVersionResource]string{
		gatewayGVR: "GatewayList",
		serviceGVR: "ServiceList",
		gaGVR:      "GlobalAcceleratorList",
	})
	for _, obj := range objs {
		gvr := map[string]schema.GroupVersionResource{"Gateway": gatewayGVR, "Service": serviceGVR, "GlobalAccelerator": gaGVR}[obj.GetKind()]
		require.NoError(t, client.Tracker().Create(gvr, obj, obj.GetNamespace()))
	}
	r := NewResolver(t.Context(), func() (dynamic.Interface, discovery.DiscoveryInterface, error) {
		return client, disc, nil
	})
	return r, client
}

func TestResolve(t *testing.T) {
	gw := newObject(gatewayGVR, "Gateway", "istio-ingress", "central", map[string]any{
		"addresses": []any{
			map[string]any{"value": "1.2.3.4"},
			map[string]any{"value": "lb.example.com"},
			map[string]any{"value": "1.2.3.4"},
		},
	})
	pending := newObject(gatewayGVR, "Gateway", "istio-ingress", "pending", map[string]any{})
	svc := newObject(serviceGVR, "Service", "apps", "shared-lb", map[string]any{
		"loadBalancer": map[string]any{"ingress": []any{map[string]any{"ip": "10.0.0.1"}}},
	})
	ga := newObject(gaGVR, "GlobalAccelerator", "", "accel", map[string]any{"dnsName": "a1.awsglobalaccelerator.com"})

	for _, tt := range []struct {
		name    string
		value   string
		want    endpoint.Targets
		wantErr string
	}{
		{
			name:  "cross-namespace gateway addresses, deduplicated",
			value: "gateways.gateway.networking.k8s.io/istio-ingress/central#" + addressesTmpl,
			want:  endpoint.Targets{"1.2.3.4", "lb.example.com"},
		},
		{
			name:  "core group, namespace defaults to the owner's",
			value: "services/shared-lb#{{ range .status.loadBalancer.ingress }}{{ .ip }},{{ end }}",
			want:  endpoint.Targets{"10.0.0.1"},
		},
		{
			name:  "cluster-scoped resource",
			value: "globalaccelerators.aga.k8s.aws/accel#{{ .status.dnsName }}",
			want:  endpoint.Targets{"a1.awsglobalaccelerator.com"},
		},
		{
			name:    "cluster-scoped resource with a namespace",
			value:   "globalaccelerators.aga.k8s.aws/apps/accel#{{ .status.dnsName }}",
			wantErr: "cluster-scoped",
		},
		{
			name:    "object not found",
			value:   "services/missing#{{ .status }}",
			wantErr: "not found",
		},
		{
			name:    "unknown resource",
			value:   "widgets.example.com/a#{{ .status }}",
			wantErr: "no matches",
		},
		{
			name:    "missing key is an error, not an empty value",
			value:   "gateways.gateway.networking.k8s.io/istio-ingress/pending#{{ .status.addresses }}",
			wantErr: "map has no entry for key",
		},
		{
			name:    "empty render, field not populated yet",
			value:   "gateways.gateway.networking.k8s.io/istio-ingress/pending#{{ range .status }}{{ . }},{{ end }}",
			wantErr: "rendered no targets",
		},
		{
			name:    "blank template",
			value:   "gateways.gateway.networking.k8s.io/istio-ingress/central#  ",
			wantErr: "empty template",
		},
		{
			name:    "rendered whitespace is unresolved",
			value:   "gateways.gateway.networking.k8s.io/istio-ingress/central#{{ \" \" }}",
			wantErr: "rendered no targets",
		},
		{
			name:    "non-scalar render",
			value:   "gateways.gateway.networking.k8s.io/istio-ingress/central#{{ .status.addresses }}",
			wantErr: "neither an IP address nor a hostname",
		},
		{
			name:    "template parse error",
			value:   "gateways.gateway.networking.k8s.io/istio-ingress/central#{{ .status",
			wantErr: "unclosed action",
		},
		{
			name:    "no template",
			value:   "services/shared-lb",
			wantErr: "missing '#<template>'",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			r, _ := newTestResolver(t, gw, pending, svc, ga)
			got, err := r.Resolve(t.Context(), "apps", tt.value)
			if tt.wantErr != "" {
				require.ErrorContains(t, err, tt.wantErr)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResolveNilResolver(t *testing.T) {
	var r *Resolver
	_, err := r.Resolve(t.Context(), "apps", "services/a#{{ .x }}")
	require.Error(t, err)

	targets, ok := r.TargetsFromAnnotations(t.Context(), "ingress", &metav1.ObjectMeta{
		Annotations: map[string]string{annotations.TargetFromKey: "services/a#{{ .x }}"},
	})
	assert.True(t, ok, "a set annotation must block the native-discovery fallback")
	assert.Empty(t, targets)

	r.AddEventHandler(func() {})
}

func TestResolveClientError(t *testing.T) {
	r := NewResolver(t.Context(), func() (dynamic.Interface, discovery.DiscoveryInterface, error) {
		return nil, nil, errors.New("no kubeconfig")
	})
	_, err := r.Resolve(t.Context(), "apps", "services/a#{{ .x }}")
	require.ErrorContains(t, err, "no kubeconfig")
}

func TestTargetsFromAnnotations(t *testing.T) {
	svc := newObject(serviceGVR, "Service", "apps", "shared-lb", map[string]any{"ip": "10.0.0.1"})
	r, _ := newTestResolver(t, svc)

	_, ok := r.TargetsFromAnnotations(t.Context(), "ingress", &metav1.ObjectMeta{Namespace: "apps"})
	assert.False(t, ok)

	targets, ok := r.TargetsFromAnnotations(t.Context(), "ingress", &metav1.ObjectMeta{
		Namespace:   "apps",
		Annotations: map[string]string{annotations.TargetFromKey: "services/shared-lb#{{ .status.ip }}"},
	})
	assert.True(t, ok)
	assert.Equal(t, endpoint.Targets{"10.0.0.1"}, targets)

	targets, ok = r.TargetsFromAnnotations(t.Context(), "ingress", &metav1.ObjectMeta{
		Namespace:   "apps",
		Annotations: map[string]string{annotations.TargetFromKey: "services/missing#{{ .status.ip }}"},
	})
	assert.True(t, ok)
	assert.Empty(t, targets)
}

// After the first direct read, the informer serves reads and its events
// trigger the registered handlers.
func TestResolverInformer(t *testing.T) {
	svc := newObject(serviceGVR, "Service", "apps", "shared-lb", map[string]any{"ip": "10.0.0.1"})
	r, client := newTestResolver(t, svc)

	var calls atomic.Int32
	r.AddEventHandler(func() { calls.Add(1) })

	value := "services/shared-lb#{{ .status.ip }}"
	got, err := r.Resolve(t.Context(), "apps", value)
	require.NoError(t, err)
	assert.Equal(t, endpoint.Targets{"10.0.0.1"}, got)

	inf := r.informers[informerKey{gvr: serviceGVR, namespace: "apps"}]
	require.NotNil(t, inf)
	require.Eventually(t, inf.Informer().HasSynced, 5*time.Second, 10*time.Millisecond)

	updated := svc.DeepCopy()
	updated.Object["status"] = map[string]any{"ip": "10.0.0.2"}
	_, err = client.Resource(serviceGVR).Namespace("apps").Update(t.Context(), updated, metav1.UpdateOptions{})
	require.NoError(t, err)

	require.Eventually(t, func() bool {
		got, err := r.Resolve(t.Context(), "apps", value)
		return err == nil && len(got) == 1 && got[0] == "10.0.0.2"
	}, 5*time.Second, 10*time.Millisecond)
	assert.GreaterOrEqual(t, calls.Load(), int32(2), "initial add and update should both trigger a sync")
}
