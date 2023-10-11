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
	"sort"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
)

// ConflictResolver is used to make a decision in case of two or more different kubernetes resources
// are trying to acquire same DNS name
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
	var min *endpoint.Endpoint
	for _, ep := range candidates {
		if min == nil || s.less(ep, min) {
			min = ep
		}
	}
	return min
}

// ResolveUpdate is invoked when dns name is already owned by "current" endpoint
// ResolveUpdate uses "current" record as base and updates it accordingly with new version of same resource
// if it doesn't exist then pick min
func (s PerResource) ResolveUpdate(current *endpoint.Endpoint, candidates []*endpoint.Endpoint) *endpoint.Endpoint {
	currentResource := current.Labels[endpoint.ResourceLabelKey] // resource which has already acquired the DNS
	// TODO: sort candidates only needed because we can still have two endpoints from same resource here. We sort for consistency
	// TODO: remove once single endpoint can have multiple targets
	sort.SliceStable(candidates, func(i, j int) bool {
		return s.less(candidates[i], candidates[j])
	})
	for _, ep := range candidates {
		if ep.Labels[endpoint.ResourceLabelKey] == currentResource {
			return ep
		}
	}
	return s.ResolveCreate(candidates)
}

// ResolveRecordTypes attempts to detect and resolve record type conflicts in desired
// endpoints for a domain. For eample if the there is more than 1 candidate and at lease one
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

	cname := false
	other := false
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

	// conflict was found, remove candiates of non-preferred record types
	if cname && other {
		log.Infof("Domain %s contains conflicting record type candidates; discarding CNAME record", key.dnsName)
		records := map[string]*domainEndpoints{}
		for recordType, recs := range row.records {
			// policy is to prefer the non-CNAME record types when a conflict is found
			if recordType == endpoint.RecordTypeCNAME {
				// discard candidates of conflicting records
				// keep currect so they can be deleted
				records[recordType] = &domainEndpoints{
					current:    recs.current,
					candidates: []*endpoint.Endpoint{},
				}
			} else {
				records[recordType] = recs
			}
		}

		return records
	}

	// no conflict, return all records types
	return row.records
}

// less returns true if endpoint x is less than y
func (s PerResource) less(x, y *endpoint.Endpoint) bool {
	return x.Targets.IsLess(y.Targets)
}

// TODO: with cross-resource/cross-cluster setup alternative variations of ConflictResolver can be used
