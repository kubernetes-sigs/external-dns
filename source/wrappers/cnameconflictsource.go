/*
Copyright 2026 The Kubernetes Authors.

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

// cnameConflictSource reports CNAME/DNAME conflicts for names within the domain filter.
// Endpoints pass through unchanged.
type cnameConflictSource struct {
	source       source.Source
	domainFilter endpoint.DomainFilterInterface
}

// NewCNAMEConflictSource creates a new cnameConflictSource wrapping the provided Source.
// A nil domainFilter puts every endpoint in scope.
func NewCNAMEConflictSource(source source.Source, domainFilter endpoint.DomainFilterInterface) source.Source {
	return &cnameConflictSource{source: source, domainFilter: domainFilter}
}

func (s *cnameConflictSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := s.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}
	s.reportConflicts(endpoints)
	return endpoints, nil
}

// reportConflicts warns about CNAME/DNAME endpoints with the same name and set identifier
// but different targets.
func (s *cnameConflictSource) reportConflicts(endpoints []*endpoint.Endpoint) {
	lastTargets := make(map[string]string) // recordType/DNSName/SetIdentifier -> last target seen
	for _, ep := range endpoints {
		if ep == nil || len(ep.Targets) == 0 || (ep.RecordType != endpoint.RecordTypeCNAME && ep.RecordType != endpoint.RecordTypeDNAME) {
			continue
		}
		if s.domainFilter != nil && !s.domainFilter.Match(ep.DNSName) {
			continue
		}
		// Lowercase the name like DomainFilter.Match does, so names that differ only in case share a key.
		key := ep.RecordType + "/" + strings.ToLower(ep.DNSName) + "/" + ep.SetIdentifier
		if last, ok := lastTargets[key]; ok && last != ep.Targets[0] {
			// Sort the pair so the same two targets always print the same way.
			log.Warnf("Only one %[1]s per name — %[2]s %[1]s %[3]s and %[2]s %[1]s %[4]s is invalid DNS. A resolver wouldn't know which canonical name to follow.", ep.RecordType, ep.DNSName, min(last, ep.Targets[0]), max(last, ep.Targets[0]))
			invalidEndpoints.AddWithLabels(1, ep.RecordType, endpointSource(ep))
		}
		lastTargets[key] = ep.Targets[0]
	}
}

func (s *cnameConflictSource) AddEventHandler(ctx context.Context, handler func()) {
	s.source.AddEventHandler(ctx, handler)
}
