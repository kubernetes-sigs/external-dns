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

	"github.com/linki/instrumented_http"
	log "github.com/sirupsen/logrus"

	dns "google.golang.org/api/dns/v1"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"

	googleapi "google.golang.org/api/googleapi"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

type managedZonesCreateCallInterface interface {
	Do(opts ...googleapi.CallOption) (*dns.ManagedZone, error)
}

type managedZonesListCallInterface interface {
	Pages(ctx context.Context, f func(*dns.ManagedZonesListResponse) error) error
}

type managedZonesServiceInterface interface {
	Create(project string, managedzone *dns.ManagedZone) managedZonesCreateCallInterface
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

func (m managedZonesService) List(project string) managedZonesListCallInterface {
	return m.service.List(project)
}

type changesService struct {
	service *dns.ChangesService
}

func (c changesService) Create(project string, managedZone string, change *dns.Change) changesCreateCallInterface {
	return c.service.Create(project, managedZone, change)
}

// GoogleProvider is an implementation of Provider for Google CloudDNS.
type GoogleProvider struct {
	// The Google project to work in
	project string
	// Enabled dry-run will print any modifying actions rather than execute them.
	dryRun bool
	// only consider hosted zones managing domains ending in this suffix
	domainFilter DomainFilter
	// A client for managing resource record sets
	resourceRecordSetsClient resourceRecordSetsClientInterface
	// A client for managing hosted zones
	managedZonesClient managedZonesServiceInterface
	// A client for managing change sets
	changesClient changesServiceInterface
}

// NewGoogleProvider initializes a new Google CloudDNS based Provider.
func NewGoogleProvider(project string, domainFilter DomainFilter, dryRun bool) (*GoogleProvider, error) {
	gcloud, err := google.DefaultClient(context.TODO(), dns.NdevClouddnsReadwriteScope)
	if err != nil {
		return nil, err
	}

	gcloud = instrumented_http.NewClient(gcloud, &instrumented_http.Callbacks{
		PathProcessor: func(path string) string {
			parts := strings.Split(path, "/")
			return parts[len(parts)-1]
		},
	})

	dnsClient, err := dns.New(gcloud)
	if err != nil {
		return nil, err
	}

	provider := &GoogleProvider{
		project:      project,
		domainFilter: domainFilter,
		dryRun:       dryRun,
		resourceRecordSetsClient: resourceRecordSetsService{dnsClient.ResourceRecordSets},
		managedZonesClient:       managedZonesService{dnsClient.ManagedZones},
		changesClient:            changesService{dnsClient.Changes},
	}

	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *GoogleProvider) Zones() (map[string]*dns.ManagedZone, error) {
	zones := make(map[string]*dns.ManagedZone)

	f := func(resp *dns.ManagedZonesListResponse) error {
		for _, zone := range resp.ManagedZones {
			if p.domainFilter.Match(zone.DnsName) {
				zones[zone.Name] = zone
			}
		}

		return nil
	}

	if err := p.managedZonesClient.List(p.project).Pages(context.TODO(), f); err != nil {
		return nil, err
	}

	return zones, nil
}

// Records returns the list of records in all relevant zones.
func (p *GoogleProvider) Records() (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.Zones()
	if err != nil {
		return nil, err
	}

	f := func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {

			for _, rr := range r.Rrdatas {
				// each page is processed sequentially, no need for a mutex here.
				if supportedRecordType(r.Type) {
					endpoints = append(endpoints, endpoint.NewEndpoint(r.Name, rr, r.Type))
				}
			}
		}

		return nil
	}

	for _, z := range zones {
		if err := p.resourceRecordSetsClient.List(p.project, z.Name).Pages(context.TODO(), f); err != nil {
			return nil, err
		}
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *GoogleProvider) CreateRecords(endpoints []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(endpoints)...)

	return p.submitChange(change)
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *GoogleProvider) UpdateRecords(records, oldRecords []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(records)...)
	change.Deletions = append(change.Deletions, newRecords(oldRecords)...)

	return p.submitChange(change)
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *GoogleProvider) DeleteRecords(endpoints []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Deletions = append(change.Deletions, newRecords(endpoints)...)

	return p.submitChange(change)
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *GoogleProvider) ApplyChanges(changes *plan.Changes) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(changes.Create)...)

	change.Additions = append(change.Additions, newRecords(changes.UpdateNew)...)
	change.Deletions = append(change.Deletions, newRecords(changes.UpdateOld)...)

	change.Deletions = append(change.Deletions, newRecords(changes.Delete)...)

	return p.submitChange(change)
}

// submitChange takes a zone and a Change and sends it to Google.
func (p *GoogleProvider) submitChange(change *dns.Change) error {
	if len(change.Additions) == 0 && len(change.Deletions) == 0 {
		log.Info("All records are already up to date")
		return nil
	}

	for _, del := range change.Deletions {
		log.Infof("Del records: %s %s %s", del.Name, del.Type, del.Rrdatas)
	}
	for _, add := range change.Additions {
		log.Infof("Add records: %s %s %s", add.Name, add.Type, add.Rrdatas)
	}

	if p.dryRun {
		return nil
	}

	zones, err := p.Zones()
	if err != nil {
		return err
	}

	// separate into per-zone change sets to be passed to the API.
	changes := separateChange(zones, change)

	for z, c := range changes {
		if _, err := p.changesClient.Create(p.project, z, c).Do(); err != nil {
			return err
		}
	}

	return nil
}

// separateChange separates a multi-zone change into a single change per zone.
func separateChange(zones map[string]*dns.ManagedZone, change *dns.Change) map[string]*dns.Change {
	changes := make(map[string]*dns.Change)
	zoneNameIDMapper := zoneIDName{}
	for _, z := range zones {
		zoneNameIDMapper[z.Name] = z.DnsName
		changes[z.Name] = &dns.Change{
			Additions: []*dns.ResourceRecordSet{},
			Deletions: []*dns.ResourceRecordSet{},
		}
	}
	for _, a := range change.Additions {
		if zoneName, _ := zoneNameIDMapper.FindZone(ensureTrailingDot(a.Name)); zoneName != "" {
			changes[zoneName].Additions = append(changes[zoneName].Additions, a)
		}
	}

	for _, d := range change.Deletions {
		if zoneName, _ := zoneNameIDMapper.FindZone(ensureTrailingDot(d.Name)); zoneName != "" {
			changes[zoneName].Deletions = append(changes[zoneName].Deletions, d)
		}
	}

	// separating a change could lead to empty sub changes, remove them here.
	for zone, change := range changes {
		if len(change.Additions) == 0 && len(change.Deletions) == 0 {
			delete(changes, zone)
		}
	}

	return changes
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
func newRecord(ep *endpoint.Endpoint) *dns.ResourceRecordSet {
	// TODO(linki): works around appending a trailing dot to TXT records. I think
	// we should go back to storing DNS names with a trailing dot internally. This
	// way we can use it has is here and trim it off if it exists when necessary.
	target := ep.Target
	if ep.RecordType == endpoint.RecordTypeCNAME {
		target = ensureTrailingDot(target)
	}

	return &dns.ResourceRecordSet{
		Name:    ensureTrailingDot(ep.DNSName),
		Rrdatas: []string{target},
		Ttl:     300,
		Type:    ep.RecordType,
	}
}
