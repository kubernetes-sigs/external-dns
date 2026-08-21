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

package registry

import (
	"context"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

// Registry tracks ownership of DNS records managed by external-dns.
type Registry interface {
	// Records returns all DNS records known to the registry, including ownership metadata.
	Records(ctx context.Context) ([]*endpoint.Endpoint, error)
	// ApplyChanges propagates the given changes to the DNS provider and updates ownership records accordingly.
	ApplyChanges(ctx context.Context, changes *plan.Changes) error
	// AdjustEndpoints normalises endpoints before they are processed by the planner.
	AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error)
	// GetDomainFilter returns the domain filter configured for the underlying provider.
	GetDomainFilter() endpoint.DomainFilterInterface
	// OwnerID returns the owner identifier used to claim DNS records.
	OwnerID() string
}
