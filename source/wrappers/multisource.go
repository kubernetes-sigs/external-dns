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

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"

	log "github.com/sirupsen/logrus"
)

// multiSource is a Source that merges the endpoints of its nested Sources.
type multiSource struct {
	children            []source.Source
	defaultTargets      []string
	forceDefaultTargets bool
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (ms *multiSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	result := []*endpoint.Endpoint{}
	hasDefaultTargets := len(ms.defaultTargets) > 0

	for _, s := range ms.children {
		endpoints, err := s.Endpoints(ctx)
		if err != nil {
			return nil, err
		}

		if !hasDefaultTargets {
			result = append(result, endpoints...)
			continue
		}

		for _, ep := range endpoints {
			hasSourceTargets := len(ep.Targets) > 0

			if ms.forceDefaultTargets || !hasSourceTargets {
				eps := source.EndpointsForHostname(ep.DNSName, ms.defaultTargets, ep.RecordTTL, ep.ProviderSpecific, ep.SetIdentifier, "")
				for _, e := range eps {
					e.Labels = ep.Labels
				}
				result = append(result, eps...)
				continue
			}

			log.Warnf("Source provided targets for %q (%s), ignoring default targets [%s] due to new behavior. Use --force-default-targets to revert to old behavior.", ep.DNSName, ep.RecordType, strings.Join(ms.defaultTargets, ", "))
			result = append(result, ep)
		}
	}

	return result, nil
}

func (ms *multiSource) AddEventHandler(ctx context.Context, handler func()) {
	for _, s := range ms.children {
		s.AddEventHandler(ctx, handler)
	}
}

// NewMultiSource creates a new multiSource.
func NewMultiSource(children []source.Source, defaultTargets []string, forceDefaultTargets bool) source.Source {
	return &multiSource{children: children, defaultTargets: defaultTargets, forceDefaultTargets: forceDefaultTargets}
}
