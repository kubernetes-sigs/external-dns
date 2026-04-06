/*
Copyright 2025 The Kubernetes Authors.

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

package txt

import (
	"maps"

	"sigs.k8s.io/external-dns/endpoint"
)

func newEndpointWithOwner(dnsName, target, recordType, ownerID string) *endpoint.Endpoint {
	return newEndpointWithOwnerAndLabels(dnsName, target, recordType, ownerID, nil)
}

func newMultiTargetEndpointWithOwner(dnsName string, targets endpoint.Targets, recordType, ownerID string) *endpoint.Endpoint {
	return newMultiTargetEndpointWithOwnerAndLabels(dnsName, targets, recordType, ownerID, nil)
}

func newTXTEndpointWithOwnedRecord(dnsName, target, ownedRecord string) *endpoint.Endpoint {
	return newEndpointWithOwnerAndLabels(dnsName, target, endpoint.RecordTypeTXT, "", endpoint.Labels{endpoint.OwnedRecordLabelKey: ownedRecord})
}

func newMultiTargetEndpointWithOwnerAndLabels(dnsName string, targets endpoint.Targets, recordType, ownerID string, labels endpoint.Labels) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, recordType, targets...)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	maps.Copy(e.Labels, labels)
	return e
}

func newEndpointWithOwnerAndLabels(dnsName, target, recordType, ownerID string, labels endpoint.Labels) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, recordType, target)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	maps.Copy(e.Labels, labels)
	return e
}

func newCNAMEEndpointWithOwnerResource(dnsName, target, ownerID, resource string) *endpoint.Endpoint {
	e := endpoint.NewEndpoint(dnsName, endpoint.RecordTypeCNAME, target)
	e.Labels[endpoint.OwnerLabelKey] = ownerID
	e.Labels[endpoint.ResourceLabelKey] = resource
	return e
}

// This is primarily used to prevent data races when running tests in parallel (t.Parallel).
func cloneEndpointsWithOpts(list []*endpoint.Endpoint, opt ...func(*endpoint.Endpoint)) []*endpoint.Endpoint {
	cloned := make([]*endpoint.Endpoint, len(list))
	for i, e := range list {
		cloned[i] = cloneEndpointWithOpts(e, opt...)
	}
	return cloned
}

func cloneEndpointWithOpts(e *endpoint.Endpoint, opt ...func(*endpoint.Endpoint)) *endpoint.Endpoint {
	targets := make(endpoint.Targets, len(e.Targets))
	copy(targets, e.Targets)

	// SameEndpoints treats nil and empty maps/slices as different.
	// To avoid introducing unintended differences, we retain nil when original is nil.
	var labels endpoint.Labels
	if e.Labels != nil {
		labels = make(endpoint.Labels, len(e.Labels))
		maps.Copy(labels, e.Labels)
	}

	var providerSpecific endpoint.ProviderSpecific
	if e.ProviderSpecific != nil {
		providerSpecific = make(endpoint.ProviderSpecific, len(e.ProviderSpecific))
		for i, p := range e.ProviderSpecific {
			providerSpecific[i] = p
		}
	}

	ttl := e.RecordTTL

	ep := &endpoint.Endpoint{
		DNSName:          e.DNSName,
		Targets:          targets,
		RecordType:       e.RecordType,
		RecordTTL:        ttl,
		Labels:           labels,
		ProviderSpecific: providerSpecific,
		SetIdentifier:    e.SetIdentifier,
	}
	for _, o := range opt {
		o(ep)
	}
	return ep
}
