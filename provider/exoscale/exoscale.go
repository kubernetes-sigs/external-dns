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

package exoscale

import (
	"context"
	"fmt"
	"strings"
	"time"

	v3 "github.com/exoscale/egoscale/v3"
	"github.com/exoscale/egoscale/v3/credentials"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
	"sigs.k8s.io/external-dns/provider/blueprint"
)

// EgoscaleClientI for replaceable implementation
type EgoscaleClientI interface {
	ListDNSDomains(context.Context) ([]v3.DNSDomain, error)
	ListDNSDomainRecords(context.Context, v3.UUID) ([]v3.DNSDomainRecord, error)
	CreateDNSDomainRecord(context.Context, v3.UUID, v3.CreateDNSDomainRecordRequest) error
	DeleteDNSDomainRecord(context.Context, v3.UUID, v3.UUID) error
	UpdateDNSDomainRecord(context.Context, v3.UUID, v3.UUID, v3.UpdateDNSDomainRecordRequest) error
}

// exoscaleClient wraps the real v3 client to satisfy EgoscaleClientI.
type exoscaleClient struct {
	c *v3.Client
}

func (w *exoscaleClient) ListDNSDomains(ctx context.Context) ([]v3.DNSDomain, error) {
	resp, err := w.c.ListDNSDomains(ctx)
	if err != nil {
		return nil, err
	}
	return resp.DNSDomains, nil
}

func (w *exoscaleClient) ListDNSDomainRecords(ctx context.Context, domainID v3.UUID) ([]v3.DNSDomainRecord, error) {
	resp, err := w.c.ListDNSDomainRecords(ctx, domainID)
	if err != nil {
		return nil, err
	}
	return resp.DNSDomainRecords, nil
}

func (w *exoscaleClient) CreateDNSDomainRecord(ctx context.Context, domainID v3.UUID, req v3.CreateDNSDomainRecordRequest) error {
	op, err := w.c.CreateDNSDomainRecord(ctx, domainID, req)
	if err != nil {
		return err
	}
	_, err = w.c.Wait(ctx, op, v3.OperationStateSuccess)
	return err
}

func (w *exoscaleClient) DeleteDNSDomainRecord(ctx context.Context, domainID v3.UUID, recordID v3.UUID) error {
	op, err := w.c.DeleteDNSDomainRecord(ctx, domainID, recordID)
	if err != nil {
		return err
	}
	_, err = w.c.Wait(ctx, op, v3.OperationStateSuccess)
	return err
}

func (w *exoscaleClient) UpdateDNSDomainRecord(ctx context.Context, domainID v3.UUID, recordID v3.UUID, req v3.UpdateDNSDomainRecordRequest) error {
	op, err := w.c.UpdateDNSDomainRecord(ctx, domainID, recordID, req)
	if err != nil {
		return err
	}
	_, err = w.c.Wait(ctx, op, v3.OperationStateSuccess)
	return err
}

// ExoscaleProvider initialized as dns provider with no records
type ExoscaleProvider struct {
	provider.BaseProvider
	domain         *endpoint.DomainFilter
	client         EgoscaleClientI
	filter         *zoneFilter
	zoneCache      *blueprint.ZoneCache[map[string]string]
	OnApplyChanges func(changes *plan.Changes)
	dryRun         bool
}

// ExoscaleOption for Provider options
type ExoscaleOption func(*ExoscaleProvider)

// New creates an Exoscale provider from the given configuration.
func New(_ context.Context, cfg *externaldns.Config, domainFilter *endpoint.DomainFilter) (provider.Provider, error) {
	return newProvider(
		cfg.ExoscaleAPIEnvironment,
		cfg.ExoscaleAPIZone,
		cfg.ExoscaleAPIKey,
		cfg.ExoscaleAPISecret,
		cfg.DryRun,
		cfg.ExoscaleZoneCacheDuration,
		ExoscaleWithDomain(domainFilter),
		ExoscaleWithLogging(),
	)
}

// newProvider returns ExoscaleProvider DNS provider interface implementation
func newProvider(env, zone, key, secret string, dryRun bool, zoneCacheDuration time.Duration, opts ...ExoscaleOption) (*ExoscaleProvider, error) {
	creds := credentials.NewStaticCredentials(key, secret)

	// Build endpoint from env and zone, e.g. env="api", zone="ch-gva-2"
	// yields "https://api-ch-gva-2.exoscale.com/v2"
	endpoint := v3.Endpoint(fmt.Sprintf("https://%s-%s.exoscale.com/v2", env, zone))

	c, err := v3.NewClient(creds, v3.ClientOptWithEndpoint(endpoint))
	if err != nil {
		return nil, err
	}

	return NewExoscaleProviderWithClient(&exoscaleClient{c: c}, dryRun, zoneCacheDuration, opts...), nil
}

// NewExoscaleProviderWithClient returns ExoscaleProvider DNS provider interface implementation (Client provided)
func NewExoscaleProviderWithClient(client EgoscaleClientI, dryRun bool, zoneCacheDuration time.Duration, opts ...ExoscaleOption) *ExoscaleProvider {
	ep := &ExoscaleProvider{
		filter:         &zoneFilter{},
		OnApplyChanges: func(_ *plan.Changes) {},
		domain:         endpoint.NewDomainFilter([]string{""}),
		client:         client,
		zoneCache:      blueprint.NewZoneCache[map[string]string](zoneCacheDuration),
		dryRun:         dryRun,
	}
	for _, opt := range opts {
		opt(ep)
	}
	return ep
}

func (ep *ExoscaleProvider) getZones(ctx context.Context) (map[string]string, error) {
	if !ep.zoneCache.Expired() {
		return ep.zoneCache.Get(), nil
	}

	domains, err := ep.client.ListDNSDomains(ctx)
	if err != nil {
		return nil, err
	}

	zones := map[string]string{}
	for _, domain := range domains {
		zones[string(domain.ID)] = domain.UnicodeName
	}

	ep.zoneCache.Reset(zones)
	return zones, nil
}

// GetDomainFilter returns the domain filter built from the zones the provider actually manages.
func (ep *ExoscaleProvider) GetDomainFilter() endpoint.DomainFilterInterface {
	zones, err := ep.getZones(context.Background())
	if err != nil {
		log.Errorf("Failed to list zones for domain filter: %v", err)
		return endpoint.NewDomainFilter(nil)
	}

	names := make([]string, 0, len(zones)*2)
	for _, name := range zones {
		names = append(names, name, "."+name)
	}
	return endpoint.NewDomainFilter(names)
}

// ApplyChanges simply modifies DNS via exoscale API
func (ep *ExoscaleProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	ep.OnApplyChanges(changes)

	if ep.dryRun {
		log.Infof("Will NOT delete these records: %+v", changes.Delete)
		log.Infof("Will NOT create these records: %+v", changes.Create)
		log.Infof("Will NOT update these records: %+v", merge(changes.UpdateOld, changes.UpdateNew))
		return nil
	}

	zones, err := ep.getZones(ctx)
	if err != nil {
		return err
	}

	for _, epoint := range changes.Create {
		if !ep.domain.Match(epoint.DNSName) {
			continue
		}

		zoneID, name := ep.filter.EndpointZoneID(epoint, zones)
		if zoneID == "" {
			continue
		}

		req := v3.CreateDNSDomainRecordRequest{
			Name:    name,
			Type:    v3.CreateDNSDomainRecordRequestType(epoint.RecordType),
			Content: epoint.Targets[0],
		}
		if epoint.RecordTTL != 0 {
			req.Ttl = int64(epoint.RecordTTL)
		}

		if err := ep.client.CreateDNSDomainRecord(ctx, v3.UUID(zoneID), req); err != nil {
			return err
		}
	}

	for _, epoint := range changes.UpdateNew {
		if !ep.domain.Match(epoint.DNSName) {
			continue
		}

		zoneID, name := ep.filter.EndpointZoneID(epoint, zones)
		if zoneID == "" {
			continue
		}

		records, err := ep.client.ListDNSDomainRecords(ctx, v3.UUID(zoneID))
		if err != nil {
			return err
		}

		for _, record := range records {
			if record.Name != name {
				continue
			}
			if string(record.Type) != epoint.RecordType {
				continue
			}

			req := v3.UpdateDNSDomainRecordRequest{
				Content: epoint.Targets[0],
			}
			if epoint.RecordTTL != 0 {
				req.Ttl = int64(epoint.RecordTTL)
			}

			if err := ep.client.UpdateDNSDomainRecord(ctx, v3.UUID(zoneID), record.ID, req); err != nil {
				return err
			}

			break
		}
	}

	for _, epoint := range changes.UpdateOld {
		// Since Exoscale "Patches", we've ignored UpdateOld
		// We leave this logging here for information
		log.Debugf("UPDATE-OLD (ignored) for epoint: %+v", epoint)
	}

	for _, epoint := range changes.Delete {
		if !ep.domain.Match(epoint.DNSName) {
			continue
		}

		zoneID, name := ep.filter.EndpointZoneID(epoint, zones)
		if zoneID == "" {
			continue
		}

		records, err := ep.client.ListDNSDomainRecords(ctx, v3.UUID(zoneID))
		if err != nil {
			return err
		}

		for _, record := range records {
			if record.Name != name {
				continue
			}
			if string(record.Type) != epoint.RecordType {
				continue
			}

			if err := ep.client.DeleteDNSDomainRecord(ctx, v3.UUID(zoneID), record.ID); err != nil {
				return err
			}

			break
		}
	}

	return nil
}

// Records returns the list of endpoints
func (ep *ExoscaleProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := ep.getZones(ctx)
	if err != nil {
		return nil, err
	}

	endpoints := make([]*endpoint.Endpoint, 0)

	for zoneID, zoneName := range zones {
		records, err := ep.client.ListDNSDomainRecords(ctx, v3.UUID(zoneID))
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			if string(record.Type) != endpoint.RecordTypeA && string(record.Type) != endpoint.RecordTypeCNAME && string(record.Type) != endpoint.RecordTypeTXT {
				continue
			}

			fqdn := zoneName
			if record.Name != "" {
				fqdn = record.Name + "." + zoneName
			}

			e := endpoint.NewEndpointWithTTL(fqdn, string(record.Type), endpoint.TTL(record.Ttl), record.Content)
			endpoints = append(endpoints, e)
		}
	}

	log.Infof("Called Records() with %d items", len(endpoints))
	return endpoints, nil
}

// ExoscaleWithDomain modifies the domain on which dns zones are filtered
func ExoscaleWithDomain(domainFilter *endpoint.DomainFilter) ExoscaleOption {
	return func(p *ExoscaleProvider) {
		p.domain = domainFilter
	}
}

// ExoscaleWithLogging injects logging when ApplyChanges is called
func ExoscaleWithLogging() ExoscaleOption {
	return func(p *ExoscaleProvider) {
		p.OnApplyChanges = func(changes *plan.Changes) {
			for _, v := range changes.Create {
				log.Infof("CREATE: %v", v)
			}
			for _, v := range changes.UpdateOld {
				log.Infof("UPDATE (old): %v", v)
			}
			for _, v := range changes.UpdateNew {
				log.Infof("UPDATE (new): %v", v)
			}
			for _, v := range changes.Delete {
				log.Infof("DELETE: %v", v)
			}
		}
	}
}

type zoneFilter struct {
	domain string
}

// Zones filters map[zoneID]zoneName for names having f.domain as suffix
func (f *zoneFilter) Zones(zones map[string]string) map[string]string {
	result := map[string]string{}
	for zoneID, zoneName := range zones {
		if strings.HasSuffix(zoneName, f.domain) {
			result[zoneID] = zoneName
		}
	}
	return result
}

// EndpointZoneID determines zoneID for endpoint from map[zoneID]zoneName by taking longest suffix zoneName match in endpoint DNSName
// returns empty string if no matches are found
func (f *zoneFilter) EndpointZoneID(endpoint *endpoint.Endpoint, zones map[string]string) (string, string) {
	var matchZoneID, matchZoneName, name string
	for zoneID, zoneName := range zones {
		if strings.HasSuffix(endpoint.DNSName, "."+zoneName) && len(zoneName) > len(matchZoneName) {
			matchZoneName = zoneName
			matchZoneID = zoneID
			name = strings.TrimSuffix(endpoint.DNSName, "."+zoneName)
		} else if endpoint.DNSName == zoneName && len(zoneName) > len(matchZoneName) {
			matchZoneName = zoneName
			matchZoneID = zoneID
			name = ""
		}
	}
	return matchZoneID, name
}

func merge(updateOld, updateNew []*endpoint.Endpoint) []*endpoint.Endpoint {
	findMatch := func(template *endpoint.Endpoint) *endpoint.Endpoint {
		for _, record := range updateNew {
			if template.DNSName == record.DNSName &&
				template.RecordType == record.RecordType {
				return record
			}
		}
		return nil
	}

	var result []*endpoint.Endpoint
	for _, old := range updateOld {
		matchingNew := findMatch(old)
		if matchingNew == nil {
			// no match shouldn't happen
			continue
		}

		if !matchingNew.Targets.Same(old.Targets) {
			// new target: always update, TTL will be overwritten too if necessary
			result = append(result, matchingNew)
			continue
		}

		if matchingNew.RecordTTL != 0 && matchingNew.RecordTTL != old.RecordTTL {
			// same target, but new non-zero TTL set in k8s, must update
			// probably would happen only if there is a bug in the code calling the provider
			result = append(result, matchingNew)
		}
	}

	return result
}
