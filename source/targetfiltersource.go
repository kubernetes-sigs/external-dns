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

package source

import (
	"context"

	"sigs.k8s.io/external-dns/endpoint"
)

// targetFilterSource is a Source that removes endpoints matching the target filter from its wrapped source.
type targetFilterSource struct {
	source       Source
	targetFilter endpoint.TargetFilterInterface
}

// NewTargetFilterSource creates a new targetFilterSource wrapping the provided Source.
func NewTargetFilterSource(source Source, targetFilter endpoint.TargetFilterInterface) Source {
	return &targetFilterSource{source: source, targetFilter: targetFilter}
}

// Endpoints collects endpoints from its wrapped source and returns
// them without targets matching the target filter.
func (ms *targetFilterSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}

	endpoints, err := ms.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	for _, ep := range endpoints {
		filteredTargets := []string{}

		for _, t := range ep.Targets {
			if ms.targetFilter.Match(t) {
				filteredTargets = append(filteredTargets, t)
			}
		}

		ep.Targets = filteredTargets

		result = append(result, ep)
	}

	return result, nil
}

func (ms *targetFilterSource) AddEventHandler(ctx context.Context, handler func()) {
	ms.source.AddEventHandler(ctx, handler)
}
