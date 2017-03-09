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

package dnsprovider

import (
	"context"
	"strings"

	log "github.com/Sirupsen/logrus"

	"google.golang.org/api/dns/v1"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// GoogleProvider is an implementation of DNSProvider for Google CloudDNS.
type GoogleProvider struct {
	// The Google project to work in
	Project string
	// Enabled dry-run will print any modifying actions rather than execute them.
	DryRun bool
	// A client for managing resource record sets
	ResourceRecordSetsClient *dns.ResourceRecordSetsService
	// A client for managing hosted zones
	ManagedZonesClient *dns.ManagedZonesService
	// A client for managing change sets
	ChangesClient *dns.ChangesService
}

// Zones returns the list of hosted zones.
func (p *GoogleProvider) Zones() (zones []*dns.ManagedZone, _ error) {
	f := func(resp *dns.ManagedZonesListResponse) error {
		// each page is processed sequentially, no need for a mutex here.
		zones = append(zones, resp.ManagedZones...)
		return nil
	}

	err := p.ManagedZonesClient.List(p.Project).Pages(context.TODO(), f)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

// CreateZone creates a hosted zone given a name.
func (p *GoogleProvider) CreateZone(name, domain string) error {
	zone := &dns.ManagedZone{
		Name:        name,
		DnsName:     domain,
		Description: "Automatically managed zone by kubernetes.io/external-dns",
	}

	_, err := p.ManagedZonesClient.Create(p.Project, zone).Do()
	if err != nil {
		return err
	}

	return nil
}

// DeleteZone deletes a hosted zone given a name.
func (p *GoogleProvider) DeleteZone(name string) error {
	err := p.ManagedZonesClient.Delete(p.Project, name).Do()
	if err != nil {
		if !isNotFound(err) {
			return err
		}
	}

	return nil
}

// Records returns the list of A records in a given hosted zone.
func (p *GoogleProvider) Records(zone string) (endpoints []endpoint.Endpoint, _ error) {
	f := func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {
			if r.Type != "A" {
				continue
			}

			for _, rr := range r.Rrdatas {
				// each page is processed sequentially, no need for a mutex here.
				endpoints = append(endpoints, endpoint.Endpoint{
					DNSName: r.Name,
					Target:  rr,
				})
			}
		}

		return nil
	}

	err := p.ResourceRecordSetsClient.List(p.Project, zone).Pages(context.TODO(), f)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *GoogleProvider) CreateRecords(zone string, endpoints []endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(endpoints)...)

	return p.submitChange(zone, change)
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *GoogleProvider) UpdateRecords(zone string, records, oldRecords []endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(records)...)
	change.Deletions = append(change.Deletions, newRecords(oldRecords)...)

	return p.submitChange(zone, change)
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *GoogleProvider) DeleteRecords(zone string, endpoints []endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Deletions = append(change.Deletions, newRecords(endpoints)...)

	return p.submitChange(zone, change)
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *GoogleProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(changes.Create)...)

	change.Additions = append(change.Additions, newRecords(changes.UpdateNew)...)
	change.Deletions = append(change.Deletions, newRecords(changes.UpdateOld)...)

	change.Deletions = append(change.Deletions, newRecords(changes.Delete)...)

	return p.submitChange(zone, change)
}

// submitChange takes a zone and a Change and sends it to Google.
func (p *GoogleProvider) submitChange(zone string, change *dns.Change) error {
	if p.DryRun {
		for _, add := range change.Additions {
			log.Infof("Add records: %s %s %s", add.Name, add.Type, add.Rrdatas)
		}

		for _, del := range change.Deletions {
			log.Infof("Del records: %s %s %s", del.Name, del.Type, del.Rrdatas)
		}

		return nil
	}

	if len(change.Additions) == 0 && len(change.Deletions) == 0 {
		return nil
	}

	_, err := p.ChangesClient.Create(p.Project, zone, change).Do()
	if err != nil {
		if !isNotFound(err) {
			return err
		}
	}

	return nil
}

// newRecords returns a collection of RecordSets based on the given endpoints.
func newRecords(endpoints []endpoint.Endpoint) []*dns.ResourceRecordSet {
	records := make([]*dns.ResourceRecordSet, len(endpoints))

	for i, endpoint := range endpoints {
		records[i] = newRecord(endpoint)
	}

	return records
}

// newRecord returns a RecordSet based on the given endpoint.
func newRecord(endpoint endpoint.Endpoint) *dns.ResourceRecordSet {
	return &dns.ResourceRecordSet{
		Name:    endpoint.DNSName,
		Rrdatas: []string{endpoint.Target},
		Ttl:     300,
		Type:    "A",
	}
}

// isNotFound returns true if a given error is due to a resource not being found.
func isNotFound(err error) bool {
	return strings.Contains(err.Error(), "notFound")
}
