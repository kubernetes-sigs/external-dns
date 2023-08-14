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
	"strings"

	egoscale "github.com/exoscale/egoscale/v2"
	exoapi "github.com/exoscale/egoscale/v2/api"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// EgoscaleClientI for replaceable implementation
type EgoscaleClientI interface {
	ListDNSDomainRecords(context.Context, string, string) ([]egoscale.DNSDomainRecord, error)
	ListDNSDomains(context.Context, string) ([]egoscale.DNSDomain, error)
	GetDNSDomainRecord(context.Context, string, string, string) (*egoscale.DNSDomainRecord, error)
	CreateDNSDomainRecord(context.Context, string, string, *egoscale.DNSDomainRecord) (*egoscale.DNSDomainRecord, error)
	DeleteDNSDomainRecord(context.Context, string, string, *egoscale.DNSDomainRecord) error
	UpdateDNSDomainRecord(context.Context, string, string, *egoscale.DNSDomainRecord) error
}

// ExoscaleProvider initialized as dns provider with no records
type ExoscaleProvider struct {
	provider.BaseProvider
	domain         endpoint.DomainFilter
	client         EgoscaleClientI
	apiEnv         string
	apiZone        string
	filter         *zoneFilter
	OnApplyChanges func(changes *plan.Changes)
	dryRun         bool
}

// ExoscaleOption for Provider options
type ExoscaleOption func(*ExoscaleProvider)

// NewExoscaleProvider returns ExoscaleProvider DNS provider interface implementation
func NewExoscaleProvider(env, zone, key, secret string, dryRun bool, opts ...ExoscaleOption) (*ExoscaleProvider, error) {
	client, err := egoscale.NewClient(
		key,
		secret,
	)
	if err != nil {
		return nil, err
	}

	return NewExoscaleProviderWithClient(client, env, zone, dryRun, opts...), nil
}

// NewExoscaleProviderWithClient returns ExoscaleProvider DNS provider interface implementation (Client provided)
func NewExoscaleProviderWithClient(client EgoscaleClientI, env, zone string, dryRun bool, opts ...ExoscaleOption) *ExoscaleProvider {
	ep := &ExoscaleProvider{
		filter:         &zoneFilter{},
		OnApplyChanges: func(changes *plan.Changes) {},
		domain:         endpoint.NewDomainFilter([]string{""}),
		client:         client,
		apiEnv:         env,
		apiZone:        zone,
		dryRun:         dryRun,
	}
	for _, opt := range opts {
		opt(ep)
	}
	return ep
}

func (ep *ExoscaleProvider) getZones(ctx context.Context) (map[string]string, error) {
	ctx = exoapi.WithEndpoint(ctx, exoapi.NewReqEndpoint(ep.apiEnv, ep.apiZone))
	domains, err := ep.client.ListDNSDomains(ctx, ep.apiZone)
	if err != nil {
		return nil, err
	}

	zones := map[string]string{}
	for _, domain := range domains {
		zones[*domain.ID] = *domain.UnicodeName
	}

	return zones, nil
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

	ctx = exoapi.WithEndpoint(ctx, exoapi.NewReqEndpoint(ep.apiEnv, ep.apiZone))

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

		// API does not accept 0 as default TTL but wants nil pointer instead
		var ttl *int64
		if epoint.RecordTTL != 0 {
			t := int64(epoint.RecordTTL)
			ttl = &t
		}
		record := egoscale.DNSDomainRecord{
			Name:    &name,
			Type:    &epoint.RecordType,
			TTL:     ttl,
			Content: &epoint.Targets[0],
		}
		_, err := ep.client.CreateDNSDomainRecord(ctx, ep.apiZone, zoneID, &record)
		if err != nil {
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

		records, err := ep.client.ListDNSDomainRecords(ctx, ep.apiZone, zoneID)
		if err != nil {
			return err
		}

		for _, r := range records {
			if *r.Name != name {
				continue
			}

			record, err := ep.client.GetDNSDomainRecord(ctx, ep.apiZone, zoneID, *r.ID)
			if err != nil {
				return err
			}

			record.Type = &epoint.RecordType
			record.Content = &epoint.Targets[0]
			if epoint.RecordTTL != 0 {
				ttl := int64(epoint.RecordTTL)
				record.TTL = &ttl
			}

			err = ep.client.UpdateDNSDomainRecord(ctx, ep.apiZone, zoneID, record)
			if err != nil {
				return err
			}

			break
		}
	}

	for _, epoint := range changes.UpdateOld {
		// Since Exoscale "Patches", we ignore UpdateOld
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

		records, err := ep.client.ListDNSDomainRecords(ctx, ep.apiZone, zoneID)
		if err != nil {
			return err
		}

		for _, record := range records {
			if *record.Name != name {
				continue
			}

			err = ep.client.DeleteDNSDomainRecord(ctx, ep.apiZone, zoneID, &egoscale.DNSDomainRecord{ID: record.ID})
			if err != nil {
				return err
			}

			break
		}
	}

	return nil
}

// Records returns the list of endpoints
func (ep *ExoscaleProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	ctx = exoapi.WithEndpoint(ctx, exoapi.NewReqEndpoint(ep.apiEnv, ep.apiZone))
	endpoints := make([]*endpoint.Endpoint, 0)

	domains, err := ep.client.ListDNSDomains(ctx, ep.apiZone)
	if err != nil {
		return nil, err
	}

	for _, domain := range domains {
		records, err := ep.client.ListDNSDomainRecords(ctx, ep.apiZone, *domain.ID)
		if err != nil {
			return nil, err
		}

		for _, r := range records {
			record, err := ep.client.GetDNSDomainRecord(ctx, ep.apiZone, *domain.ID, *r.ID)
			if err != nil {
				return nil, err
			}
			switch *record.Type {
			case "A", "CNAME", "TXT":
				break
			default:
				continue
			}

			e := endpoint.NewEndpointWithTTL((*record.Name)+"."+(*domain.UnicodeName), *record.Type, endpoint.TTL(*r.TTL), *record.Content)
			endpoints = append(endpoints, e)
		}
	}

	log.Infof("called Records() with %d items", len(endpoints))
	return endpoints, nil
}

// ExoscaleWithDomain modifies the domain on which dns zones are filtered
func ExoscaleWithDomain(domainFilter endpoint.DomainFilter) ExoscaleOption {
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
func (f *zoneFilter) EndpointZoneID(endpoint *endpoint.Endpoint, zones map[string]string) (zoneID string, name string) {
	var matchZoneID string
	var matchZoneName string
	for zoneID, zoneName := range zones {
		if strings.HasSuffix(endpoint.DNSName, "."+zoneName) && len(zoneName) > len(matchZoneName) {
			matchZoneName = zoneName
			matchZoneID = zoneID
			name = strings.TrimSuffix(endpoint.DNSName, "."+zoneName)
		}
	}
	return matchZoneID, name
}

func merge(updateOld, updateNew []*endpoint.Endpoint) []*endpoint.Endpoint {
	findMatch := func(template *endpoint.Endpoint) *endpoint.Endpoint {
		for _, new := range updateNew {
			if template.DNSName == new.DNSName &&
				template.RecordType == new.RecordType {
				return new
			}
		}
		return nil
	}

	var result []*endpoint.Endpoint
	for _, old := range updateOld {
		matchingNew := findMatch(old)
		if matchingNew == nil {
			// no match, shouldn't happen
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
