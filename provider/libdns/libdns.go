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

// Package libdns adapts https://github.com/libdns modules to the ExternalDNS
// Provider interface. A single generic adapter serves any registered libdns
// module; concrete modules are wired in behind the `libdns` build tag (see
// registry.go) so the default binary pulls in no vendor SDK.
package libdns

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/libdns/libdns"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// libdnsClient is the minimal set of libdns interfaces every registered module
// must satisfy. RecordSetter and ZoneLister are optional and detected at
// runtime via type assertion.
type libdnsClient interface {
	libdns.RecordGetter
	libdns.RecordAppender
	libdns.RecordDeleter
}

// clientFactory builds a libdns module client from the JSON config blob passed
// via --libdns-config.
type clientFactory func(config string) (libdnsClient, error)

// Provider adapts a libdns module to the ExternalDNS Provider interface.
type Provider struct {
	provider.BaseProvider
	client       libdnsClient
	zones        []string
	domainFilter *endpoint.DomainFilter
	dryRun       bool
}

// New builds a libdns-backed provider for the module named by --libdns-provider.
func New(ctx context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	factory, ok := modules[cfg.LibdnsProvider]
	if !ok {
		return nil, fmt.Errorf("unknown libdns provider %q (available: %s); rebuild with -tags libdns to enable libdns modules", cfg.LibdnsProvider, strings.Join(moduleNames(), ", "))
	}

	client, err := factory(cfg.LibdnsConfig)
	if err != nil {
		return nil, fmt.Errorf("configuring libdns provider %q: %w", cfg.LibdnsProvider, err)
	}

	zones, err := resolveZones(ctx, client, cfg.DomainFilter)
	if err != nil {
		return nil, err
	}

	return &Provider{
		client:       client,
		zones:        zones,
		domainFilter: domainFilter,
		dryRun:       cfg.DryRun,
	}, nil
}

// resolveZones determines the managed zones. --domain-filter is the primary
// source (works for every module). When it is empty and the module implements
// ZoneLister, zones are auto-discovered; otherwise --domain-filter is required.
func resolveZones(ctx context.Context, client libdnsClient, domainFilter []string) ([]string, error) {
	var zones []string
	for _, z := range domainFilter {
		if z = strings.TrimSuffix(strings.TrimSpace(z), "."); z != "" {
			zones = append(zones, z)
		}
	}
	if len(zones) > 0 {
		return zones, nil
	}

	lister, ok := client.(libdns.ZoneLister)
	if !ok {
		return nil, fmt.Errorf("libdns provider requires --domain-filter: the module does not support zone discovery")
	}
	discovered, err := lister.ListZones(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing libdns zones: %w", err)
	}
	for _, z := range discovered {
		zones = append(zones, strings.TrimSuffix(z.Name, "."))
	}
	if len(zones) == 0 {
		return nil, fmt.Errorf("libdns module returned no zones; set --domain-filter")
	}
	return zones, nil
}

// Records returns the current records across all managed zones.
func (p *Provider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	for _, zone := range p.zones {
		records, err := p.client.GetRecords(ctx, ensureTrailingDot(zone))
		if err != nil {
			return nil, provider.NewSoftError(fmt.Errorf("listing records for zone %q: %w", zone, err))
		}
		endpoints = append(endpoints, groupRecords(zone, records)...)
	}
	return endpoints, nil
}

// groupRecords folds libdns records into endpoints, grouping targets sharing the
// same (name, type) RRset.
func groupRecords(zone string, records []libdns.Record) []*endpoint.Endpoint {
	type key struct{ name, rtype string }
	groups := map[key]*endpoint.Endpoint{}
	var endpoints []*endpoint.Endpoint
	for _, record := range records {
		rr := record.RR()
		name := strings.TrimSuffix(libdns.AbsoluteName(rr.Name, ensureTrailingDot(zone)), ".")
		k := key{name, rr.Type}
		if ep, ok := groups[k]; ok {
			ep.Targets = append(ep.Targets, rr.Data)
			continue
		}
		ep := endpoint.NewEndpointWithTTL(name, rr.Type, endpoint.TTL(rr.TTL/time.Second), rr.Data)
		groups[k] = ep
		endpoints = append(endpoints, ep)
	}
	return endpoints
}

// ApplyChanges applies create/update/delete changes to the libdns module.
func (p *Provider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	setByZone := map[string][]libdns.Record{}
	deleteByZone := map[string][]libdns.Record{}

	// Create and UpdateNew describe the desired final state of an RRset, which
	// maps onto libdns SetRecords semantics.
	for _, ep := range append(append([]*endpoint.Endpoint{}, changes.Create...), changes.UpdateNew...) {
		zone, ok := p.zoneFor(ep.DNSName)
		if !ok {
			log.Warnf("libdns: no managed zone for %q, skipping", ep.DNSName)
			continue
		}
		setByZone[zone] = append(setByZone[zone], toRecords(zone, ep)...)
	}
	for _, ep := range changes.Delete {
		zone, ok := p.zoneFor(ep.DNSName)
		if !ok {
			log.Warnf("libdns: no managed zone for %q, skipping", ep.DNSName)
			continue
		}
		deleteByZone[zone] = append(deleteByZone[zone], toRecords(zone, ep)...)
	}

	if p.dryRun {
		for zone, recs := range deleteByZone {
			log.Infof("libdns: would delete %d record(s) in zone %q", len(recs), zone)
		}
		for zone, recs := range setByZone {
			log.Infof("libdns: would set %d record(s) in zone %q", len(recs), zone)
		}
		return nil
	}

	for zone, recs := range deleteByZone {
		if _, err := p.client.DeleteRecords(ctx, ensureTrailingDot(zone), recs); err != nil {
			return provider.NewSoftError(fmt.Errorf("deleting records in zone %q: %w", zone, err))
		}
	}
	for zone, recs := range setByZone {
		if err := p.setRecords(ctx, zone, recs); err != nil {
			return provider.NewSoftError(fmt.Errorf("setting records in zone %q: %w", zone, err))
		}
	}
	return nil
}

// setRecords writes an RRset. It prefers RecordSetter; modules lacking it fall
// back to clearing the whole RRset then appending the desired records. The
// delete must target the RRset by (name, type) — not the new records — otherwise
// targets dropped in an update would survive.
func (p *Provider) setRecords(ctx context.Context, zone string, records []libdns.Record) error {
	absZone := ensureTrailingDot(zone)
	if setter, ok := p.client.(libdns.RecordSetter); ok {
		_, err := setter.SetRecords(ctx, absZone, records)
		return err
	}
	if _, err := p.client.DeleteRecords(ctx, absZone, rrsetStubs(records)); err != nil {
		return err
	}
	_, err := p.client.AppendRecords(ctx, absZone, records)
	return err
}

// rrsetStubs reduces records to one (name, type) stub per RRset. libdns
// DeleteRecords treats an empty Data field as a wildcard, so each stub deletes
// every record sharing that name and type.
func rrsetStubs(records []libdns.Record) []libdns.Record {
	seen := map[string]bool{}
	var stubs []libdns.Record
	for _, r := range records {
		rr := r.RR()
		k := rr.Name + "\x00" + rr.Type
		if seen[k] {
			continue
		}
		seen[k] = true
		stubs = append(stubs, libdns.RR{Name: rr.Name, Type: rr.Type})
	}
	return stubs
}

// AdjustEndpoints clears SetIdentifier, which flat libdns backends cannot
// represent. Keeping it would make the plan key on a non-empty SetIdentifier
// that the backend returns empty on read-back, causing perpetual create/delete
// churn. Provider-native routing (weighted/latency/geo) lives in ProviderSpecific
// and is only ever set by routing-aware sources, so nothing else needs stripping.
func (p *Provider) AdjustEndpoints(endpoints []*endpoint.Endpoint) ([]*endpoint.Endpoint, error) {
	for _, ep := range endpoints {
		if ep.SetIdentifier != "" {
			log.Warnf("libdns: ignoring unsupported set-identifier %q on %q", ep.SetIdentifier, ep.DNSName)
			ep.SetIdentifier = ""
		}
	}
	return endpoints, nil
}

// GetDomainFilter returns the configured domain filter.
func (p *Provider) GetDomainFilter() endpoint.DomainFilterInterface {
	return p.domainFilter
}

// zoneFor returns the managed zone with the longest suffix match for an FQDN.
func (p *Provider) zoneFor(dnsName string) (string, bool) {
	name := strings.TrimSuffix(dnsName, ".")
	var best string
	for _, zone := range p.zones {
		if name == zone || strings.HasSuffix(name, "."+zone) {
			if len(zone) > len(best) {
				best = zone
			}
		}
	}
	return best, best != ""
}

// toRecords converts an endpoint to one libdns record per target.
func toRecords(zone string, ep *endpoint.Endpoint) []libdns.Record {
	name := libdns.RelativeName(ep.DNSName, ensureTrailingDot(zone))
	var ttl time.Duration
	if ep.RecordTTL.IsConfigured() {
		ttl = time.Duration(ep.RecordTTL) * time.Second
	}
	records := make([]libdns.Record, 0, len(ep.Targets))
	for _, target := range ep.Targets {
		records = append(records, libdns.RR{
			Name: name,
			Type: ep.RecordType,
			TTL:  ttl,
			Data: target,
		})
	}
	return records
}

// moduleNames returns the registered libdns module names, sorted.
func moduleNames() []string {
	names := make([]string, 0, len(modules))
	for name := range modules {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func ensureTrailingDot(zone string) string {
	return strings.TrimSuffix(zone, ".") + "."
}
