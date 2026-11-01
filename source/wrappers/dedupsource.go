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

package wrappers

import (
	"context"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

// dedupSource is a Source that removes duplicate endpoints from its wrapped source.
type dedupSource struct {
	source source.Source
}

// NewDedupSource creates a new dedupSource wrapping the provided Source.
func NewDedupSource(source source.Source) source.Source {
	return &dedupSource{source: source}
}

// Endpoints collects endpoints from its wrapped source and returns them without duplicates.
func (ms *dedupSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Debug("dedupSource: collecting endpoints and removing duplicates")
	resetMetrics()

	endpoints, err := ms.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*endpoint.Endpoint, 0, len(endpoints))
	byIdentifier := endpoint.NewOrderedMap[string](len(endpoints))

	for _, ep := range endpoints {
		if ep == nil {
			continue
		}

		// validate endpoint before normalization
		if ok := ep.CheckEndpoint(); !ok {
			log.Warnf("Skipping endpoint [%s:%s] due to invalid configuration [%s:%s]", ep.SetIdentifier, ep.DNSName, ep.RecordType, strings.Join(ep.Targets, ","))
			invalidEndpoints.AddWithLabels(1, ep.RecordType, endpointSource(ep))
			continue
		}

		byIdentifier.Append(GetIdentifier(ep), ep)
	}

	for _, eps := range byIdentifier.All() {
		endpoint.SortByResource(eps)

		winner, losers := eps[0], eps[1:]
		endpoint.AbsorbRefObjects(winner, losers...)
		for _, loser := range losers {
			deduplicatedEndpoints.AddWithLabels(1, loser.RecordType, endpointSource(loser))
		}

		result = append(result, winner)
	}

	return result, nil
}

func (ms *dedupSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("dedupSource: adding event handler")
	ms.source.AddEventHandler(ctx, handler)
}

// GetIdentifier returns the key under which endpoints are considered duplicates.
func GetIdentifier(ep *endpoint.Endpoint) string {
	if len(ep.Targets) > 1 {
		ep.Targets = endpoint.NewTargets(ep.Targets...)
	}

	return strings.Join([]string{ep.RecordType, ep.DNSName, ep.SetIdentifier, ep.Targets.String()}, "/")
}
