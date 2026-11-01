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
	"reflect"
	"slices"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

// multiSource is a Source that merges the endpoints of its nested Sources.
type multiSource struct {
	children            []source.Source
	defaultTargets      []string
	forceDefaultTargets bool
}

// Endpoints collects endpoints of all nested Sources and returns them in a single slice.
func (ms *multiSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	log.Debugf("multiSource: collecting endpoints from %d child sources and removing duplicates", len(ms.children))
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
			if !isAddressType(ep.RecordType) {
				result = append(result, ep)
				continue
			}

			hasSourceTargets := len(ep.Targets) > 0

			if ms.forceDefaultTargets || !hasSourceTargets {
				eps := endpoint.EndpointsForHostname(ep.DNSName, ms.defaultTargets, ep.RecordTTL, ep.ProviderSpecific, ep.SetIdentifier, "")
				eps = slices.DeleteFunc(eps, func(e *endpoint.Endpoint) bool { return e.RecordType != ep.RecordType })
				if len(eps) == 0 {
					// Keeping the source targets avoids deleting a live record because of a flag mismatch.
					if hasSourceTargets {
						log.Warnf("No default target [%s] fits %s record %q, keeping source targets", strings.Join(ms.defaultTargets, ", "), ep.RecordType, ep.DNSName)
						result = append(result, ep)
					} else {
						log.Warnf("No default target [%s] fits %s record %q, skipping it", strings.Join(ms.defaultTargets, ", "), ep.RecordType, ep.DNSName)
					}
					continue
				}
				for _, e := range eps {
					e.Labels = ep.Labels
					// Without the refs, events and the DNSEndpoint status can't tell which object these records came from.
					for _, ref := range ep.RefObjects() {
						e.WithRefObject(ref)
					}
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

// isAddressType reports whether default targets can stand in for the endpoint's targets.
func isAddressType(recordType string) bool {
	switch recordType {
	case endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME:
		return true
	}
	return false
}

func (ms *multiSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debugf("multiSource: adding event handler for %d child sources", len(ms.children))
	for _, s := range ms.children {
		log.Debugf("multiSource: adding event handler for child %q", reflect.TypeOf(s).String())
		s.AddEventHandler(ctx, handler)
	}
}

// NewMultiSource creates a new multiSource.
func NewMultiSource(children []source.Source, defaultTargets []string, forceDefaultTargets bool) source.Source {
	return &multiSource{children: children, defaultTargets: defaultTargets, forceDefaultTargets: forceDefaultTargets}
}
