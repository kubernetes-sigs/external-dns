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
	log "github.com/Sirupsen/logrus"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/dns/v1"
	googleapi "google.golang.org/api/googleapi"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

type managedZonesCreateCallInterface interface {
	Do(opts ...googleapi.CallOption) (*dns.ManagedZone, error)
}

type managedZonesDeleteCallInterface interface {
	Do(opts ...googleapi.CallOption) error
}

type managedZonesListCallInterface interface {
	Pages(ctx context.Context, f func(*dns.ManagedZonesListResponse) error) error
}

type managedZonesServiceInterface interface {
	Create(project string, managedzone *dns.ManagedZone) managedZonesCreateCallInterface
	Delete(project string, managedZone string) managedZonesDeleteCallInterface
	List(project string) managedZonesListCallInterface
}

type resourceRecordSetsListCallInterface interface {
	Pages(ctx context.Context, f func(*dns.ResourceRecordSetsListResponse) error) error
}

type resourceRecordSetsClientInterface interface {
	List(project string, managedZone string) resourceRecordSetsListCallInterface
}

type changesCreateCallInterface interface {
	Do(opts ...googleapi.CallOption) (*dns.Change, error)
}

type changesServiceInterface interface {
	Create(project string, managedZone string, change *dns.Change) changesCreateCallInterface
}

type resourceRecordSetsService struct {
	service *dns.ResourceRecordSetsService
}

func (r resourceRecordSetsService) List(project string, managedZone string) resourceRecordSetsListCallInterface {
	return r.service.List(project, managedZone)
}

type managedZonesService struct {
	service *dns.ManagedZonesService
}

func (m managedZonesService) Create(project string, managedzone *dns.ManagedZone) managedZonesCreateCallInterface {
	return m.service.Create(project, managedzone)
}

func (m managedZonesService) Delete(project string, managedZone string) managedZonesDeleteCallInterface {
	return m.service.Delete(project, managedZone)
}

func (m managedZonesService) List(project string) managedZonesListCallInterface {
	return m.service.List(project)
}

type changesService struct {
	service *dns.ChangesService
}

func (c changesService) Create(project string, managedZone string, change *dns.Change) changesCreateCallInterface {
	return c.service.Create(project, managedZone, change)
}

// googleProvider is an implementation of Provider for Google CloudDNS.
type googleProvider struct {
	// The Google project to work in
	project string
	// Enabled dry-run will print any modifying actions rather than execute them.
	dryRun bool
	// A client for managing resource record sets
	resourceRecordSetsClient resourceRecordSetsClientInterface
	// A client for managing hosted zones
	managedZonesClient managedZonesServiceInterface
	// A client for managing change sets
	changesClient changesServiceInterface
}

// NewGoogleProvider initializes a new Google CloudDNS based Provider.
func NewGoogleProvider(project string, dryRun bool) (Provider, error) {
	gcloud, err := google.DefaultClient(context.TODO(), dns.NdevClouddnsReadwriteScope)
	if err != nil {
		return nil, err
	}

	dnsClient, err := dns.New(gcloud)
	if err != nil {
		return nil, err
	}

	provider := &googleProvider{
		project: project,
		dryRun:  dryRun,
		resourceRecordSetsClient: resourceRecordSetsService{dnsClient.ResourceRecordSets},
		managedZonesClient:       managedZonesService{dnsClient.ManagedZones},
		changesClient:            changesService{dnsClient.Changes},
	}

	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *googleProvider) Zones() (zones []*dns.ManagedZone, _ error) {
	f := func(resp *dns.ManagedZonesListResponse) error {
		// each page is processed sequentially, no need for a mutex here.
		zones = append(zones, resp.ManagedZones...)
		return nil
	}

	err := p.managedZonesClient.List(p.project).Pages(context.TODO(), f)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

// CreateZone creates a hosted zone given a name.
func (p *googleProvider) CreateZone(name, domain string) error {
	zone := &dns.ManagedZone{
		Name:        name,
		DnsName:     domain,
		Description: "Automatically managed zone by kubernetes.io/external-dns",
	}

	_, err := p.managedZonesClient.Create(p.project, zone).Do()
	if err != nil {
		return err
	}

	return nil
}

// DeleteZone deletes a hosted zone given a name.
func (p *googleProvider) DeleteZone(name string) error {
	err := p.managedZonesClient.Delete(p.project, name).Do()
	if err != nil {
		return err
	}

	return nil
}

// Records returns the list of A records in a given hosted zone.
func (p *googleProvider) Records(zone string) (endpoints []*endpoint.Endpoint, _ error) {
	f := func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {
			// TODO(linki, ownership): Remove once ownership system is in place.
			// See: https://github.com/kubernetes-incubator/external-dns/pull/122/files/74e2c3d3e237411e619aefc5aab694742001cdec#r109863370
			switch r.Type {
			case "A", "CNAME", "TXT":
				break
			default:
				continue
			}

			for _, rr := range r.Rrdatas {
				// each page is processed sequentially, no need for a mutex here.
				endpoints = append(endpoints, endpoint.NewEndpoint(r.Name, rr, r.Type))
			}
		}

		return nil
	}

	err := p.resourceRecordSetsClient.List(p.project, zone).Pages(context.TODO(), f)
	if err != nil {
		return nil, err
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *googleProvider) CreateRecords(zone string, endpoints []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(endpoints)...)

	return p.submitChange(zone, change)
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *googleProvider) UpdateRecords(zone string, records, oldRecords []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(records)...)
	change.Deletions = append(change.Deletions, newRecords(oldRecords)...)

	return p.submitChange(zone, change)
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *googleProvider) DeleteRecords(zone string, endpoints []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Deletions = append(change.Deletions, newRecords(endpoints)...)

	return p.submitChange(zone, change)
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *googleProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(changes.Create)...)

	change.Additions = append(change.Additions, newRecords(changes.UpdateNew)...)
	change.Deletions = append(change.Deletions, newRecords(changes.UpdateOld)...)

	change.Deletions = append(change.Deletions, newRecords(changes.Delete)...)

	return p.submitChange(zone, change)
}

// submitChange takes a zone and a Change and sends it to Google.
func (p *googleProvider) submitChange(zone string, change *dns.Change) error {

	if len(change.Additions) == 0 && len(change.Deletions) == 0 {
		log.Infoln("All records are already up to date")
		return nil
	}

	for _, del := range change.Deletions {
		log.Infof("Del records: %s %s %s", del.Name, del.Type, del.Rrdatas)
	}
	for _, add := range change.Additions {
		log.Infof("Add records: %s %s %s", add.Name, add.Type, add.Rrdatas)
	}

	if !p.dryRun {
		_, err := p.changesClient.Create(p.project, zone, change).Do()
		if err != nil {
			return err
		}
	}

	return nil
}

// newRecords returns a collection of RecordSets based on the given endpoints.
func newRecords(endpoints []*endpoint.Endpoint) []*dns.ResourceRecordSet {
	records := make([]*dns.ResourceRecordSet, len(endpoints))

	for i, endpoint := range endpoints {
		records[i] = newRecord(endpoint)
	}

	return records
}

// newRecord returns a RecordSet based on the given endpoint.
func newRecord(endpoint *endpoint.Endpoint) *dns.ResourceRecordSet {
	return &dns.ResourceRecordSet{
		Name:    ensureTrailingDot(endpoint.DNSName),
		Rrdatas: []string{ensureTrailingDot(endpoint.Target)},
		Ttl:     300,
		Type:    suitableType(endpoint),
	}
}
