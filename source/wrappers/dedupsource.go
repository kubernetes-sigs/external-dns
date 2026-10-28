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
	"cmp"
	"context"
	"iter"
	"slices"
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
	byIdentifier := NewOrderedMap(len(endpoints))

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
		slices.SortStableFunc(eps, CompareByResource)

		winner, losers := eps[0], eps[1:]
		for _, looser := range losers {
			for _, ref := range looser.RefObjects() {
				winner.WithRefObject(ref)
			}

			deduplicatedEndpoints.AddWithLabels(1, looser.RecordType, endpointSource(looser))
		}

		result = append(result, winner)
	}

	return result, nil
}

func (ms *dedupSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("dedupSource: adding event handler")
	ms.source.AddEventHandler(ctx, handler)
}

func GetIdentifier(ep *endpoint.Endpoint) string {
	if len(ep.Targets) > 1 {
		ep.Targets = endpoint.NewTargets(ep.Targets...)
	}

	return strings.Join([]string{ep.RecordType, ep.DNSName, ep.SetIdentifier, ep.Targets.String()}, "/")
}

func CompareByResource(ae, be *endpoint.Endpoint) int {
	a, b := ae.Labels[endpoint.ResourceLabelKey], be.Labels[endpoint.ResourceLabelKey]
	if a == "" || b == "" {
		return cmp.Compare(b, a)
	}
	return cmp.Compare(a, b)
}

func NewOrderedMap(capacity int) *OrderedMap {
	return &OrderedMap{
		keys: make([]string, 0, capacity),
		m:    make(map[string][]*endpoint.Endpoint, capacity),
	}
}

type OrderedMap struct {
	keys []string
	m    map[string][]*endpoint.Endpoint
}

func (o *OrderedMap) Append(k string, v *endpoint.Endpoint) {
	if _, ok := o.m[k]; !ok {
		o.keys = append(o.keys, k)
	}
	o.m[k] = append(o.m[k], v)
}

// All makes OrderedMap rangeable: for k, v := range om.All()
func (o *OrderedMap) All() iter.Seq2[string, []*endpoint.Endpoint] {
	return func(yield func(string, []*endpoint.Endpoint) bool) {
		for _, k := range o.keys {
			if !yield(k, o.m[k]) {
				return
			}
		}
	}
}
