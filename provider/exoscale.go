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

package provider

import (
	"strings"

	"github.com/exoscale/egoscale"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	log "github.com/sirupsen/logrus"
)

// EgoscaleClientI for replaceable implementation
type EgoscaleClientI interface {
	GetRecords(string) ([]egoscale.DNSRecord, error)
	GetDomains() ([]egoscale.DNSDomain, error)
	CreateRecord(string, egoscale.DNSRecord) (*egoscale.DNSRecord, error)
	DeleteRecord(string, int64) error
	UpdateRecord(string, egoscale.UpdateDNSRecord) (*egoscale.DNSRecord, error)
}

// ExoscaleProvider initialized as dns provider with no records
type ExoscaleProvider struct {
	domain         DomainFilter
	client         EgoscaleClientI
	filter         *zoneFilter
	OnApplyChanges func(changes *plan.Changes)
	dryRun         bool
}

// ExoscaleOption for Provider options
type ExoscaleOption func(*ExoscaleProvider)

// NewExoscaleProvider returns ExoscaleProvider DNS provider interface implementation
func NewExoscaleProvider(endpoint, apiKey, apiSecret string, dryRun bool, opts ...ExoscaleOption) *ExoscaleProvider {
	client := egoscale.NewClient(endpoint, apiKey, apiSecret)
	return NewExoscaleProviderWithClient(endpoint, apiKey, apiSecret, client, dryRun, opts...)
}

// NewExoscaleProviderWithClient returns ExoscaleProvider DNS provider interface implementation (Client provided)
func NewExoscaleProviderWithClient(endpoint, apiKey, apiSecret string, client EgoscaleClientI, dryRun bool, opts ...ExoscaleOption) *ExoscaleProvider {
	ep := &ExoscaleProvider{
		filter:         &zoneFilter{},
		OnApplyChanges: func(changes *plan.Changes) {},
		domain:         NewDomainFilter([]string{""}),
		client:         client,
		dryRun:         dryRun,
	}
	for _, opt := range opts {
		opt(ep)
	}
	return ep
}

func (ep *ExoscaleProvider) getZones() (map[int64]string, error) {
	dom, err := ep.client.GetDomains()
	if err != nil {
		return nil, err
	}

	zones := map[int64]string{}
	for _, d := range dom {
		zones[d.ID] = d.Name
	}
	return zones, nil
}

// ApplyChanges simply modifies DNS via exoscale API
func (ep *ExoscaleProvider) ApplyChanges(changes *plan.Changes) error {
	ep.OnApplyChanges(changes)

	if ep.dryRun {
		log.Infof("Will NOT delete these records: %+v", changes.Delete)
		log.Infof("Will NOT create these records: %+v", changes.Create)
		log.Infof("Will NOT update these records: %+v", merge(changes.UpdateOld, changes.UpdateNew))
		return nil
	}

	zones, err := ep.getZones()
	if err != nil {
		return err
	}

	for _, epoint := range changes.Create {
		if ep.domain.Match(epoint.DNSName) {
			if zoneID, name := ep.filter.EndpointZoneID(epoint, zones); zoneID != 0 {
				rec := egoscale.DNSRecord{
					Name:       name,
					RecordType: epoint.RecordType,
					TTL:        int(epoint.RecordTTL),
					Content:    epoint.Targets[0],
				}
				_, err := ep.client.CreateRecord(zones[zoneID], rec)
				if err != nil {
					return err
				}
			}
		}
	}
	for _, epoint := range changes.UpdateNew {
		if ep.domain.Match(epoint.DNSName) {
			if zoneID, name := ep.filter.EndpointZoneID(epoint, zones); zoneID != 0 {
				records, err := ep.client.GetRecords(zones[zoneID])
				if err != nil {
					return err
				}
				for _, r := range records {
					if r.Name == name {
						rec := egoscale.UpdateDNSRecord{
							ID:         r.ID,
							DomainID:   r.DomainID,
							Name:       name,
							RecordType: epoint.RecordType,
							TTL:        int(epoint.RecordTTL),
							Content:    epoint.Targets[0],
							Prio:       r.Prio,
						}
						if _, err := ep.client.UpdateRecord(zones[zoneID], rec); err != nil {
							return err
						}
						break
					}
				}
			}
		}
	}

	for _, epoint := range changes.UpdateOld {
		// Since Exoscale "Patches", we ignore UpdateOld
		// We leave this logging here for information
		log.Debugf("UPDATE-OLD (ignored) for epoint: %+v", epoint)
	}

	for _, epoint := range changes.Delete {
		if ep.domain.Match(epoint.DNSName) {
			if zoneID, name := ep.filter.EndpointZoneID(epoint, zones); zoneID != 0 {
				records, err := ep.client.GetRecords(zones[zoneID])
				if err != nil {
					return err
				}

				for _, r := range records {
					if r.Name == name {
						if err := ep.client.DeleteRecord(zones[zoneID], r.ID); err != nil {
							return err
						}
						break
					}
				}
			}
		}
	}

	return nil
}

// Records returns the list of endpoints
func (ep *ExoscaleProvider) Records() ([]*endpoint.Endpoint, error) {
	endpoints := make([]*endpoint.Endpoint, 0)

	domains, err := ep.client.GetDomains()
	if err != nil {
		return nil, err
	}

	for _, d := range domains {
		record, err := ep.client.GetRecords(d.Name)
		if err != nil {
			return nil, err
		}
		for _, r := range record {
			switch r.RecordType {
			case egoscale.A.String(), egoscale.CNAME.String(), egoscale.TXT.String():
				break
			default:
				continue
			}
			ep := endpoint.NewEndpointWithTTL(r.Name+"."+d.Name, r.RecordType, endpoint.TTL(r.TTL), r.Content)
			endpoints = append(endpoints, ep)
		}
	}

	log.Infof("called Records() with %d items", len(endpoints))
	return endpoints, nil
}

// ExoscaleWithDomain modifies the domain on which dns zones are filtered
func ExoscaleWithDomain(domainFilter DomainFilter) ExoscaleOption {
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
func (f *zoneFilter) Zones(zones map[int64]string) map[int64]string {
	result := map[int64]string{}
	for zoneID, zoneName := range zones {
		if strings.HasSuffix(zoneName, f.domain) {
			result[zoneID] = zoneName
		}
	}
	return result
}

// EndpointZoneID determines zoneID for endpoint from map[zoneID]zoneName by taking longest suffix zoneName match in endpoint DNSName
// returns 0 if no matches are found
func (f *zoneFilter) EndpointZoneID(endpoint *endpoint.Endpoint, zones map[int64]string) (zoneID int64, name string) {
	var matchZoneID int64
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
