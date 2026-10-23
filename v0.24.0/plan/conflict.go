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

package plan

import (
	"slices"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

// ConflictResolver is used to make a decision in case of two or more different kubernetes resources
// are trying to acquire the same DNS name
type ConflictResolver interface {
	ResolveCreate(candidates []*endpoint.Endpoint) *endpoint.Endpoint
	ResolveUpdate(current *endpoint.Endpoint, candidates []*endpoint.Endpoint) *endpoint.Endpoint
	ResolveRecordTypes(key planKey, row *planTableRow) map[string]*domainEndpoints
}

// PerResource allows only one resource to own a given dns name
type PerResource struct{}

// ResolveCreate is invoked when dns name is not owned by any resource
// ResolveCreate takes "minimal" (string comparison of Target) endpoint to acquire the DNS record
func (s PerResource) ResolveCreate(candidates []*endpoint.Endpoint) *endpoint.Endpoint {
	return slices.MinFunc(candidates, compareEndpoints)
}

// ResolveUpdate is invoked when dns name is already owned by "current" endpoint
// ResolveUpdate uses "current" record as base and updates it accordingly with new version of same resource
// if it doesn't exist then pick min
func (s PerResource) ResolveUpdate(current *endpoint.Endpoint, candidates []*endpoint.Endpoint) *endpoint.Endpoint {
	currentResource := current.Labels[endpoint.ResourceLabelKey] // resource which has already acquired the DNS
	slices.SortStableFunc(candidates, compareEndpoints)
	for _, ep := range candidates {
		if ep.Labels[endpoint.ResourceLabelKey] == currentResource {
			return ep
		}
	}
	return s.ResolveCreate(candidates)
}

// ResolveRecordTypes attempts to detect and resolve record type conflicts in desired
// endpoints for a domain. For example if there is more than 1 candidate and at least one
// of them is a CNAME. Per [RFC 1034 3.6.2] domains that contain a CNAME can not contain any
// other record types. The default policy will prefer A and AAAA record types when a conflict is
// detected (consistent with [endpoint.Targets.Less]).
//
// [RFC 1034 3.6.2]: https://datatracker.ietf.org/doc/html/rfc1034#autoid-15
func (s PerResource) ResolveRecordTypes(key planKey, row *planTableRow) map[string]*domainEndpoints {
	// no conflicts if only a single desired record type for the domain
	if len(row.candidates) <= 1 {
		return row.records
	}

	cname, other := false, false
	for _, c := range row.candidates {
		if c.RecordType == endpoint.RecordTypeCNAME {
			cname = true
		} else {
			other = true
		}
		if cname && other {
			break
		}
	}

	if !cname || !other {
		return row.records
	}

	// conflict was found: prefer non-CNAME record types, discard CNAME candidates
	// but keep current CNAME so it can be deleted
	// TODO: emit metric
	log.Warnf("Domain %s contains conflicting record type candidates; discarding CNAME record", key.dnsName)
	records := make(map[string]*domainEndpoints, len(row.records))
	for recordType, recs := range row.records {
		if recordType == endpoint.RecordTypeCNAME {
			records[recordType] = &domainEndpoints{current: recs.current, candidates: []*endpoint.Endpoint{}}
			continue
		}
		records[recordType] = recs
	}
	return records
}

// compareEndpoints compares two endpoints by their targets for use in sort/min operations.
func compareEndpoints(a, b *endpoint.Endpoint) int {
	if a.Targets.IsLess(b.Targets) {
		return -1
	}
	if b.Targets.IsLess(a.Targets) {
		return 1
	}
	return 0
}

// TODO: with cross-resource/cross-cluster setup alternative variations of ConflictResolver can be used
