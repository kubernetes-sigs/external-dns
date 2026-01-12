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
	result := make([]*endpoint.Endpoint, 0)
	collected := make(map[string]struct{})

	endpoints, err := ms.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	for _, ep := range endpoints {
		if ep == nil {
			continue
		}

		// validate endpoint before normalization
		if ok := ep.CheckEndpoint(); !ok {
			log.Warnf("Skipping endpoint [%s:%s] due to invalid configuration [%s:%s]", ep.SetIdentifier, ep.DNSName, ep.RecordType, strings.Join(ep.Targets, ","))
			continue
		}

		if len(ep.Targets) > 1 {
			ep.Targets = endpoint.NewTargets(ep.Targets...)
		}

		identifier := strings.Join([]string{ep.RecordType, ep.DNSName, ep.SetIdentifier, ep.Targets.String()}, "/")

		if _, ok := collected[identifier]; ok {
			log.Debugf("Removing duplicate endpoint %s", ep)
			continue
		}

		collected[identifier] = struct{}{}
		result = append(result, ep)
	}

	return result, nil
}

func (ms *dedupSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("dedupSource: adding event handler")
	ms.source.AddEventHandler(ctx, handler)
}
