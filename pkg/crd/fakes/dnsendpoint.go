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
	endpoints   *apiv1alpha1.DNSEndpoint
	namespace   string
	apiVersion  string
	kind        string
	returnError bool
}

// NewFakeDNSEndpointClient returns the concrete fake type for tests that need it
func NewFakeDNSEndpointClient(
	endpoints *apiv1alpha1.DNSEndpoint,
	namespace, apiVersion, kind string,
) *DNSEndpointClient {
	return &DNSEndpointClient{
		endpoints:  endpoints,
		namespace:  namespace,
		apiVersion: apiVersion,
		kind:       kind,
	}
}

func (f *DNSEndpointClient) Get(_ context.Context, namespace, name string) (*apiv1alpha1.DNSEndpoint, error) {
	if f.returnError {
		return nil, fmt.Errorf("error getting DNSEndpoint")
	}
	if f.endpoints.Namespace == namespace && f.endpoints.Name == name {
		return f.endpoints, nil
	}
	return nil, fmt.Errorf("not found")
}

func (f *DNSEndpointClient) List(_ context.Context, namespace string, _ *metav1.ListOptions) (*apiv1alpha1.DNSEndpointList, error) {
	if f.returnError {
		return nil, fmt.Errorf("error listing DNSEndpoints")
	}
	// Return empty list if namespace doesn't match
	if namespace != "" && f.endpoints.Namespace != namespace {
		return &apiv1alpha1.DNSEndpointList{}, nil
	}
	return &apiv1alpha1.DNSEndpointList{
		Items: []apiv1alpha1.DNSEndpoint{*f.endpoints},
	}, nil
}

func (f *DNSEndpointClient) UpdateStatus(_ context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint) (*apiv1alpha1.DNSEndpoint, error) {
	if f.returnError {
		return nil, fmt.Errorf("error updating status")
	}
	f.endpoints.Status.ObservedGeneration = dnsEndpoint.Status.ObservedGeneration
	return f.endpoints, nil
}

func (f *DNSEndpointClient) Watch(_ context.Context, _ string, _ *metav1.ListOptions) (watch.Interface, error) {
	if f.returnError {
		return nil, fmt.Errorf("error watching")
	}
	return watch.NewFake(), nil
}
