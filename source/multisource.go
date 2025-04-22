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
	"strings"

	log "github.com/sirupsen/logrus"
	"sigs.k8s.io/external-dns/endpoint"
)

// multiSource is a Source that merges the endpoints of its nested Sources.
type multiSource struct {
	children            []Source
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

		if hasDefaultTargets {
			for i := range endpoints {
				hasSourceTargets := len(endpoints[i].Targets) > 0

				if !ms.forceDefaultTargets && hasSourceTargets {
					// New behavior (default): Source targets exist, use them and ignore defaults.
					// Log a warning every time this happens if defaults are configured.
					log.Warnf("Source provided targets for %q (%s), ignoring default targets [%s] due to new behavior. Use --force-default-targets to revert to old behavior.", endpoints[i].DNSName, endpoints[i].RecordType, strings.Join(ms.defaultTargets, ", "))
					result = append(result, endpoints[i])
					continue
				} else if ms.forceDefaultTargets || !hasSourceTargets {
					// Old behavior (forced via flag) OR New behavior (source targets are empty): Apply default targets.
					eps := endpointsForHostname(endpoints[i].DNSName, ms.defaultTargets, endpoints[i].RecordTTL, endpoints[i].ProviderSpecific, endpoints[i].SetIdentifier, "")
					for _, ep := range eps {
						ep.Labels = endpoints[i].Labels
					}
					result = append(result, eps...)
				} else {
					// This case should logically not be reached given the conditions above, but handles completeness.
					result = append(result, endpoints[i])
				}
			}
		} else {
			// No default targets configured, just append source endpoints.
			result = append(result, endpoints...)
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
func NewMultiSource(children []Source, defaultTargets []string, forceDefaultTargets bool) Source {
	return &multiSource{children: children, defaultTargets: defaultTargets, forceDefaultTargets: forceDefaultTargets}
}
