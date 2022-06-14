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
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	// OwnerID returns the owner identifier used to claim DNS records.
	OwnerID() string
||||||| parent of e93f1e928 (UPSTREAM 2811: Handle the migration to the new TXT format - create missing records)
||||||| parent of 8fa7f7d99 (UPSTREAM 2811: Handle the migration to the new TXT format: create missing records)
=======
	MissingRecords() []*endpoint.Endpoint
>>>>>>> 8fa7f7d99 (UPSTREAM 2811: Handle the migration to the new TXT format: create missing records)
}

//TODO(ideahitme): consider moving this to Plan
func filterOwnedRecords(ownerID string, eps []*endpoint.Endpoint) []*endpoint.Endpoint {
	filtered := []*endpoint.Endpoint{}
	for _, ep := range eps {
		if endpointOwner, ok := ep.Labels[endpoint.OwnerLabelKey]; !ok || endpointOwner != ownerID {
			log.Debugf(`Skipping endpoint %v because owner id does not match, found: "%s", required: "%s"`, ep, endpointOwner, ownerID)
			continue
		}
		filtered = append(filtered, ep)
	}
	return filtered
=======
	MissingRecords() []*endpoint.Endpoint
||||||| parent of 959bf0129 (Revert "UPSTREAM 2811: Handle the migration to the new TXT format - create missing records")
	MissingRecords() []*endpoint.Endpoint
=======
>>>>>>> 959bf0129 (Revert "UPSTREAM 2811: Handle the migration to the new TXT format - create missing records")
}

//TODO(ideahitme): consider moving this to Plan
func filterOwnedRecords(ownerID string, eps []*endpoint.Endpoint) []*endpoint.Endpoint {
	filtered := []*endpoint.Endpoint{}
	for _, ep := range eps {
		if endpointOwner, ok := ep.Labels[endpoint.OwnerLabelKey]; !ok || endpointOwner != ownerID {
			log.Debugf(`Skipping endpoint %v because owner id does not match, found: "%s", required: "%s"`, ep, endpointOwner, ownerID)
			continue
		}
		filtered = append(filtered, ep)
	}
	return filtered
>>>>>>> e93f1e928 (UPSTREAM 2811: Handle the migration to the new TXT format - create missing records)
}
