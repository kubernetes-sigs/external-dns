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

package fakes

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
)

type DNSEndpointClient struct {
	endpoints   []apiv1alpha1.DNSEndpoint
	namespace   string
	apiVersion  string
	kind        string
	returnError bool
}

// watchTrackingClient wraps fakeDNSEndpointClient to track watch calls
type WatchTrackingClient struct {
	*DNSEndpointClient
	watchCalled bool
}

// NewFakeDNSEndpointClient returns the concrete fake type for tests that need it.
// For backwards compatibility, accepts a single endpoint.
func NewFakeDNSEndpointClient(
	endpoint *apiv1alpha1.DNSEndpoint,
	namespace, apiVersion, kind string,
) *DNSEndpointClient {
	return &DNSEndpointClient{
		endpoints:  []apiv1alpha1.DNSEndpoint{*endpoint},
		namespace:  namespace,
		apiVersion: apiVersion,
		kind:       kind,
	}
}

// NewFakeDNSEndpointClientWithList returns a fake client with multiple endpoints.
func NewFakeDNSEndpointClientWithList(
	list *apiv1alpha1.DNSEndpointList,
	namespace, apiVersion, kind string,
) *DNSEndpointClient {
	return &DNSEndpointClient{
		endpoints:  list.Items,
		namespace:  namespace,
		apiVersion: apiVersion,
		kind:       kind,
	}
}

func (f *DNSEndpointClient) Get(_ context.Context, namespace, name string) (*apiv1alpha1.DNSEndpoint, error) {
	if f.returnError {
		return nil, fmt.Errorf("error getting DNSEndpoint")
	}
	for i := range f.endpoints {
		if f.endpoints[i].Namespace == namespace && f.endpoints[i].Name == name {
			return &f.endpoints[i], nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func (f *DNSEndpointClient) List(_ context.Context, namespace string, _ *metav1.ListOptions) (*apiv1alpha1.DNSEndpointList, error) {
	if f.returnError {
		return nil, fmt.Errorf("error listing DNSEndpoints")
	}
	// Filter by namespace if specified
	if namespace == "" {
		return &apiv1alpha1.DNSEndpointList{Items: f.endpoints}, nil
	}
	var filtered []apiv1alpha1.DNSEndpoint
	for _, ep := range f.endpoints {
		if ep.Namespace == namespace {
			filtered = append(filtered, ep)
		}
	}
	return &apiv1alpha1.DNSEndpointList{Items: filtered}, nil
}

func (f *DNSEndpointClient) UpdateStatus(_ context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint) (*apiv1alpha1.DNSEndpoint, error) {
	if f.returnError {
		return nil, fmt.Errorf("error updating status")
	}
	for i := range f.endpoints {
		if f.endpoints[i].Name == dnsEndpoint.Name && f.endpoints[i].Namespace == dnsEndpoint.Namespace {
			f.endpoints[i].Status.ObservedGeneration = dnsEndpoint.Status.ObservedGeneration
			return &f.endpoints[i], nil
		}
	}
	return dnsEndpoint, nil
}

func (f *DNSEndpointClient) Watch(_ context.Context, _ string, _ *metav1.ListOptions) (watch.Interface, error) {
	if f.returnError {
		return nil, fmt.Errorf("error watching")
	}
	return watch.NewFake(), nil
}

func (w *WatchTrackingClient) Watch(
	ctx context.Context,
	namespace string,
	opts *metav1.ListOptions) (watch.Interface, error) {
	w.watchCalled = true
	return w.DNSEndpointClient.Watch(ctx, namespace, opts)
}

func (w *WatchTrackingClient) WatchCalled() bool {
	return w.watchCalled
}
