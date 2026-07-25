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

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/source"
)

// cnameConflictSource is a Source that warns about conflicting CNAME endpoints, i.e.
// endpoints sharing the same DNS name and set identifier but pointing at different
// targets. Per the DNS spec only one CNAME is allowed per name, so such desired state
// is invalid and the planner will resolve it by picking a single candidate.
//
// The warning is only emitted for DNS names that match the configured domain filter.
// Names outside the domain filter are never managed by this ExternalDNS instance
// (e.g. Istio or Ingress hosts used purely for in-cluster routing), so conflicts on
// them are logged at debug level instead to avoid noise.
//
// Endpoints are returned unmodified.
type cnameConflictSource struct {
	source       source.Source
	domainFilter endpoint.DomainFilterInterface
}

// NewCNAMEConflictSource creates a new cnameConflictSource wrapping the provided Source.
// A nil domainFilter matches all domains.
func NewCNAMEConflictSource(source source.Source, domainFilter endpoint.DomainFilterInterface) source.Source {
	return &cnameConflictSource{source: source, domainFilter: domainFilter}
}

// Endpoints collects endpoints from its wrapped source and returns them unmodified,
// warning about CNAME conflicts on DNS names matching the domain filter.
func (s *cnameConflictSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := s.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}
	s.warnOnCNAMEConflicts(endpoints)
	return endpoints, nil
}

// warnOnCNAMEConflicts logs a warning for every CNAME endpoint whose DNS name and set
// identifier were already claimed by a CNAME endpoint with a different target. Conflicts
// on DNS names not matching the domain filter are logged at debug level only.
func (s *cnameConflictSource) warnOnCNAMEConflicts(endpoints []*endpoint.Endpoint) {
	cnameTargets := make(map[string]string) // DNSName+SetIdentifier -> first target seen
	for _, ep := range endpoints {
		if ep == nil || ep.RecordType != endpoint.RecordTypeCNAME || len(ep.Targets) == 0 {
			continue
		}
		key := ep.DNSName + "/" + ep.SetIdentifier
		first, seen := cnameTargets[key]
		if !seen {
			cnameTargets[key] = ep.Targets[0]
			continue
		}
		if first == ep.Targets[0] {
			continue
		}
		if s.domainFilter != nil && !s.domainFilter.Match(ep.DNSName) {
			log.Debugf("Skipping CNAME conflict warning for %s: it does not match the domain filter", ep.DNSName)
			continue
		}
		// This will be caught by the provider when it tries to create the record, but log a warning here to make it more obvious.
		// TODO: add metric for CNAME conflicts
		log.Warnf("Only one CNAME per name — %s CNAME %s and %s CNAME %s is invalid DNS. A resolver wouldn't know which canonical name to follow.", ep.DNSName, first, ep.DNSName, ep.Targets[0])
	}
}

func (s *cnameConflictSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("cnameConflictSource: adding event handler")
	s.source.AddEventHandler(ctx, handler)
}
