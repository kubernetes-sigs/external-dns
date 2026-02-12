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

package crd

import (
	"context"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"

	apiv1alpha1 "sigs.k8s.io/external-dns/apis/v1alpha1"
)

// DNSEndpointClient provides CRUD operations for DNSEndpoint custom resources.
// This is a repository/data access layer with no business logic.
type DNSEndpointClient interface {
	// Get retrieves a single DNSEndpoint by namespace and name
	Get(ctx context.Context, namespace, name string) (*apiv1alpha1.DNSEndpoint, error)

	// List retrieves all DNSEndpoints matching the given options
	List(ctx context.Context, namespace string, opts *metav1.ListOptions) (*apiv1alpha1.DNSEndpointList, error)

	// UpdateStatus updates the status subresource of a DNSEndpoint
	UpdateStatus(ctx context.Context, dnsEndpoint *apiv1alpha1.DNSEndpoint) (*apiv1alpha1.DNSEndpoint, error)

	// Watch returns a watch interface for DNSEndpoint changes
	Watch(ctx context.Context, namespace string, opts *metav1.ListOptions) (watch.Interface, error)
}

// dnsEndpointClient implements DNSEndpointClient interface
type dnsEndpointClient struct {
	restClient rest.Interface
	resource   string
	codec      runtime.ParameterCodec
}

// NewDNSEndpointClient creates a new DNSEndpointClient.
// Parameters:
//   - restClient: Kubernetes REST client configured for the DNSEndpoint API group
//   - kind: The Kind name (e.g., "DNSEndpoint") - will be pluralized to resource name
func NewDNSEndpointClient(
	restClient rest.Interface,
	kind string) DNSEndpointClient {
	return &dnsEndpointClient{
		restClient: restClient,
		resource:   strings.ToLower(kind) + "s", // e.g., "DNSEndpoint" -> "dnsendpoints"
		codec:      metav1.ParameterCodec,
	}
}

// Get retrieves a single DNSEndpoint by namespace and name
func (c *dnsEndpointClient) Get(
	ctx context.Context,
	namespace, name string) (*apiv1alpha1.DNSEndpoint, error) {
	result := &apiv1alpha1.DNSEndpoint{}
	return result, c.restClient.Get().
		Namespace(namespace).
		Resource(c.resource).
		Name(name).
		Do(ctx).
		Into(result)
}

// List retrieves all DNSEndpoints matching the given options
func (c *dnsEndpointClient) List(
	ctx context.Context,
	namespace string,
	opts *metav1.ListOptions) (*apiv1alpha1.DNSEndpointList, error) {
	result := &apiv1alpha1.DNSEndpointList{}
	return result, c.restClient.Get().
		Namespace(namespace).
		Resource(c.resource).
		VersionedParams(opts, c.codec).
		Do(ctx).
		Into(result)
}

// UpdateStatus updates the status subresource of a DNSEndpoint
func (c *dnsEndpointClient) UpdateStatus(
	ctx context.Context,
	dnsEndpoint *apiv1alpha1.DNSEndpoint) (*apiv1alpha1.DNSEndpoint, error) {
	result := &apiv1alpha1.DNSEndpoint{}
	return result, c.restClient.Put().
		Namespace(dnsEndpoint.Namespace).
		Resource(c.resource).
		Name(dnsEndpoint.Name).
		SubResource("status").
		Body(dnsEndpoint).
		Do(ctx).
		Into(result)
}

// Watch returns a watch interface for DNSEndpoint changes
func (c *dnsEndpointClient) Watch(
	ctx context.Context,
	ns string,
	opts *metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.restClient.Get().
		Namespace(ns).
		Resource(c.resource).
		VersionedParams(opts, c.codec).
		Watch(ctx)
}
