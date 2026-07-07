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
	"fmt"
	"strings"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/internal/sets"
	"sigs.k8s.io/external-dns/source"
	sourcetemplate "sigs.k8s.io/external-dns/source/template"
)

// acmeChallengePrefix is the well-known DNS-01 challenge label defined by RFC 8555.
const acmeChallengePrefix = "_acme-challenge."

// acmeDelegationSource wraps a Source to append an ACME DNS-01 delegation CNAME
// (_acme-challenge.<hostname>) for every A/AAAA/CNAME endpoint, pointing at a
// target rendered from --acme-cname-delegation-target-template. Hostnames can be
// restricted with --acme-cname-delegation-domain-filter, and the CNAME TTL set
// with --acme-cname-delegation-ttl.
type acmeDelegationSource struct {
	source       source.Source
	tmpl         *template.Template
	domainFilter *endpoint.DomainFilter
	// ttl for the generated CNAMEs; 0 leaves the TTL unconfigured so --min-ttl
	// or the provider default applies.
	ttl endpoint.TTL
}

// acmeDelegationTemplateData is the data available to
// --acme-cname-delegation-target-template.
type acmeDelegationTemplateData struct {
	// Hostname is the source hostname as discovered, including a leading "*." for wildcards.
	Hostname string
	// HostnameWithoutWildcard is the hostname with any leading "*." stripped.
	HostnameWithoutWildcard string
}

// NewACMEDelegationSource creates a source wrapper that generates ACME DNS-01
// delegation CNAME records for the hostnames of the wrapped source.
func NewACMEDelegationSource(src source.Source, targetTemplate string, domainFilter []string, ttl time.Duration) (source.Source, error) {
	tmpl, err := sourcetemplate.Parse(targetTemplate)
	if err != nil {
		return nil, fmt.Errorf("failed to parse --acme-cname-delegation-target-template: %w", err)
	}
	return &acmeDelegationSource{
		source:       src,
		tmpl:         tmpl,
		domainFilter: endpoint.NewDomainFilter(domainFilter),
		ttl:          endpoint.TTL(ttl.Seconds()),
	}, nil
}

// Endpoints collects endpoints from its wrapped source and appends a delegation
// CNAME for every eligible hostname that does not already have one.
func (s *acmeDelegationSource) Endpoints(ctx context.Context) ([]*endpoint.Endpoint, error) {
	endpoints, err := s.source.Endpoints(ctx)
	if err != nil {
		return nil, err
	}

	existing := sets.New[string]()
	for _, ep := range endpoints {
		existing.Insert(strings.ToLower(ep.DNSName))
	}

	seen := sets.New[string]()
	var additional []*endpoint.Endpoint
	for _, ep := range endpoints {
		if !supportsACMEDelegation(ep) || !s.domainFilter.Match(ep.DNSName) {
			continue
		}

		base := strings.TrimPrefix(ep.DNSName, "*.")
		challengeName := acmeChallengePrefix + base
		key := strings.ToLower(challengeName)
		if seen.Has(key) {
			continue
		}
		seen.Insert(key)
		if existing.Has(key) {
			log.Debugf("ACME delegation: skipping %s, an endpoint with this name is already defined by a source", challengeName)
			continue
		}

		target, err := s.renderTarget(ep.DNSName, base)
		if err != nil {
			log.Warnf("ACME delegation: %v", err)
			continue
		}

		cname := endpoint.NewEndpointWithTTL(challengeName, endpoint.RecordTypeCNAME, s.ttl, target)
		if cname == nil {
			log.Warnf("ACME delegation: skipping invalid delegation record %s", challengeName)
			continue
		}
		if resource, ok := ep.Labels[endpoint.ResourceLabelKey]; ok {
			cname.WithLabel(endpoint.ResourceLabelKey, resource)
		}
		for _, ref := range ep.RefObjects() {
			cname.WithRefObject(ref)
		}
		additional = append(additional, cname)
	}

	if len(additional) > 0 {
		log.Debugf("ACME delegation: generated %d delegation CNAME endpoints", len(additional))
		endpoints = append(endpoints, additional...)
	}
	return endpoints, nil
}

func (s *acmeDelegationSource) AddEventHandler(ctx context.Context, handler func()) {
	log.Debug("acmeDelegationSource: adding event handler")
	s.source.AddEventHandler(ctx, handler)
}

// supportsACMEDelegation returns true if the endpoint is eligible for delegation
// CNAME generation. Only hostname-bearing address records (A, AAAA, CNAME) may
// need certificates; hostnames already carrying the challenge prefix are skipped
// to keep explicitly managed delegation records untouched.
func supportsACMEDelegation(ep *endpoint.Endpoint) bool {
	switch ep.RecordType {
	case endpoint.RecordTypeA, endpoint.RecordTypeAAAA, endpoint.RecordTypeCNAME:
		return !strings.HasPrefix(ep.DNSName, acmeChallengePrefix)
	default:
		return false
	}
}

// renderTarget executes the target template for a hostname and returns the
// resulting CNAME target.
func (s *acmeDelegationSource) renderTarget(hostname, base string) (string, error) {
	var buf strings.Builder
	data := acmeDelegationTemplateData{Hostname: hostname, HostnameWithoutWildcard: base}
	if err := s.tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute target template for %q: %w", hostname, err)
	}
	target := strings.TrimSpace(buf.String())
	if target == "" {
		return "", fmt.Errorf("skipping %q, target template produced an empty target", hostname)
	}
	return target, nil
}
