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
func (p *GoogleProvider) Zones() ([]*dns.ManagedZone, error) {
	zones, err := p.ManagedZonesClient.List(p.Project).Do()
	if err != nil {
		return nil, err
	}

	return zones.ManagedZones, nil
}

// CreateZone creates a hosted zone given a name.
func (p *GoogleProvider) CreateZone(name, domain string) error {
	zone := &dns.ManagedZone{
		Name:        name,
		DnsName:     domain,
		Description: "TODO",
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

// Records returns the list of records in a given hosted zone.
func (p *GoogleProvider) Records(zone string) ([]endpoint.Endpoint, error) {
	records, err := p.ResourceRecordSetsClient.List(p.Project, zone).Do()
	if err != nil {
		return nil, err
	}

	endpoints := []endpoint.Endpoint{}

	for _, r := range records.Rrsets {
		for _, rr := range r.Rrdatas {
			endpoint := endpoint.Endpoint{
				DNSName: r.Name,
				Target:  rr,
			}

			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints, nil
}

// CreateRecord creates a given DNS record in the given hosted zone.
func (p *GoogleProvider) CreateRecord(zone string, record endpoint.Endpoint) error {
	createRecord := &dns.ResourceRecordSet{
		Name:    record.DNSName,
		Rrdatas: []string{record.Target},
		Ttl:     300,
		Type:    "A",
	}

	change := &dns.Change{
		Additions: []*dns.ResourceRecordSet{createRecord},
	}

	if p.DryRun {
		log.Infof("Creating record: %#v", createRecord)
		return nil
	}

	_, err := p.ChangesClient.Create(p.Project, zone, change).Do()
	if err != nil {
		return err
	}

	return nil
}

// UpdateRecord updates a given old record to a new record in a given hosted zone.
func (p *GoogleProvider) UpdateRecord(zone string, newRecord, oldRecord endpoint.Endpoint) error {
	deleteRecord := &dns.ResourceRecordSet{
		Name:    oldRecord.DNSName,
		Rrdatas: []string{oldRecord.Target},
		Ttl:     300,
		Type:    "A",
	}

	createRecord := &dns.ResourceRecordSet{
		Name:    newRecord.DNSName,
		Rrdatas: []string{newRecord.Target},
		Ttl:     300,
		Type:    "A",
	}

	change := &dns.Change{
		Deletions: []*dns.ResourceRecordSet{deleteRecord},
		Additions: []*dns.ResourceRecordSet{createRecord},
	}

	if p.DryRun {
		log.Infof("Update record: %#v %#v", deleteRecord, createRecord)
		return nil
	}

	_, err := p.ChangesClient.Create(p.Project, zone, change).Do()
	if err != nil {
		return err
	}

	return nil
}

// DeleteRecord deletes a given DNS record in a given zone.
func (p *GoogleProvider) DeleteRecord(zone string, record endpoint.Endpoint) error {
	deleteRecord := &dns.ResourceRecordSet{
		Name:    record.DNSName,
		Rrdatas: []string{record.Target},
		Ttl:     300,
		Type:    "A",
	}

	change := &dns.Change{
		Deletions: []*dns.ResourceRecordSet{deleteRecord},
	}

	if p.DryRun {
		log.Infof("Delete record: %#v %#v", deleteRecord)
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

// ApplyChanges applies a given set of changes in a given zone.
func (p *GoogleProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	for _, record := range changes.Create {
		err := p.CreateRecord(zone, record)
		if err != nil {
			return err
		}
	}

	for i := range changes.UpdateNew {
		err := p.UpdateRecord(zone, changes.UpdateNew[i], changes.UpdateOld[i])
		if err != nil {
			return err
		}
	}

	for _, record := range changes.Delete {
		err := p.DeleteRecord(zone, record)
		if err != nil {
			return err
		}
	}

	return nil
}

// isNotFound returns true if a given error is due to a resource not being found.
func isNotFound(err error) bool {
	return strings.Contains(err.Error(), "notFound")
}
