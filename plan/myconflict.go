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

	"github.com/kubernetes-incubator/external-dns/endpoint"
)

// PerResource allows only one resource to own a given dns name
type PreferMax struct{}

func (s PreferMax) ResolveCreate(candidates []*endpoint.Endpoint) *endpoint.Endpoint {
	var min *endpoint.Endpoint
	for _, ep := range candidates {
		if min == nil {
			min = ep
			continue
		}
		if min.RecordType == endpoint.RecordTypeA && ep.RecordType == endpoint.RecordTypeCNAME {
			min = ep
		}
		if min.RecordType == endpoint.RecordTypeA && ep.RecordType == endpoint.RecordTypeA && ep.Targets.IsLess(min.Targets) {
			min = ep
		}
	}
	return min
}

func (s PreferMax) ResolveUpdate(current *endpoint.Endpoint, candidates []*endpoint.Endpoint) *endpoint.Endpoint {
	currentResource := current.Labels[endpoint.ResourceLabelKey] // resource which has already acquired the DNS
	// TODO: sort candidates only needed because we can still have two endpoints from same resource here. We sort for consistency
	// TODO: remove once single endpoint can have multiple targets
	sort.SliceStable(candidates, func(i, j int) bool {
		return s.less(candidates[i], candidates[j])
	})
	var resourceCandidates []*endpoint.Endpoint
	for _, ep := range candidates {
		if ep.Labels[endpoint.ResourceLabelKey] == currentResource && current.RecordType == ep.RecordType {
			resourceCandidates = append(resourceCandidates, ep)
		}
	}
	if len(resourceCandidates) > 0 {
		return s.ResolveCreate(resourceCandidates)
	}
	return s.ResolveCreate(candidates)
}

// less returns true if endpoint x is less than y
func (s PreferMax) less(x, y *endpoint.Endpoint) bool {
	return x.Targets.IsLess(y.Targets)
}

// TODO: with cross-resource/cross-cluster setup alternative variations of ConflictResolver can be used
