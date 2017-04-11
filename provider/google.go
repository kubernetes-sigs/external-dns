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
	"net"
	"strings"

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
	// TODO
	domain string
	// A client for managing resource record sets
	resourceRecordSetsClient resourceRecordSetsClientInterface
	// A client for managing hosted zones
	managedZonesClient managedZonesServiceInterface
	// A client for managing change sets
	changesClient changesServiceInterface
}

// NewGoogleProvider initializes a new Google CloudDNS based Provider.
func NewGoogleProvider(project string, domain string, dryRun bool) (Provider, error) {
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
		domain:  domain,
		dryRun:  dryRun,
		resourceRecordSetsClient: resourceRecordSetsService{dnsClient.ResourceRecordSets},
		managedZonesClient:       managedZonesService{dnsClient.ManagedZones},
		changesClient:            changesService{dnsClient.Changes},
	}

	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *googleProvider) Zones() (map[string]*dns.ManagedZone, error) {
	zones := make(map[string]*dns.ManagedZone)

	f := func(resp *dns.ManagedZonesListResponse) error {
		for _, zone := range resp.ManagedZones {
			if strings.HasSuffix(zone.DnsName, p.domain) {
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
func (p *googleProvider) Records(_ string) (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.Zones()
	if err != nil {
		return nil, err
	}

	f := func(resp *dns.ResourceRecordSetsListResponse) error {
		for _, r := range resp.Rrsets {
			// TODO(linki, ownership): Remove once ownership system is in place.
			// See: https://github.com/kubernetes-incubator/external-dns/pull/122/files/74e2c3d3e237411e619aefc5aab694742001cdec#r109863370
			switch r.Type {
			case "A", "CNAME":
				break
			default:
				continue
			}

			for _, rr := range r.Rrdatas {
				// each page is processed sequentially, no need for a mutex here.
				endpoints = append(endpoints, endpoint.NewEndpoint(r.Name, rr))
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
func (p *googleProvider) CreateRecords(endpoints []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(endpoints)...)

	return p.submitChange(change)
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *googleProvider) UpdateRecords(records, oldRecords []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(records)...)
	change.Deletions = append(change.Deletions, newRecords(oldRecords)...)

	return p.submitChange(change)
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *googleProvider) DeleteRecords(endpoints []*endpoint.Endpoint) error {
	change := &dns.Change{}

	change.Deletions = append(change.Deletions, newRecords(endpoints)...)

	return p.submitChange(change)
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *googleProvider) ApplyChanges(_ string, changes *plan.Changes) error {
	change := &dns.Change{}

	change.Additions = append(change.Additions, newRecords(changes.Create)...)

	change.Additions = append(change.Additions, newRecords(changes.UpdateNew)...)
	change.Deletions = append(change.Deletions, newRecords(changes.UpdateOld)...)

	change.Deletions = append(change.Deletions, newRecords(changes.Delete)...)

	return p.submitChange(change)
}

// submitChange takes a zone and a Change and sends it to Google.
func (p *googleProvider) submitChange(change *dns.Change) error {
	if p.dryRun {
		for _, del := range change.Deletions {
			log.Infof("Del records: %s %s %s", del.Name, del.Type, del.Rrdatas)
		}
		for _, add := range change.Additions {
			log.Infof("Add records: %s %s %s", add.Name, add.Type, add.Rrdatas)
		}

		return nil
	}

	if len(change.Additions) == 0 && len(change.Deletions) == 0 {
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
			if !isNotFound(err) {
				return err
			}
		}
	}

	return nil
}

// separateChange separates a multi-zone change into a single change per zone.
func separateChange(zones map[string]*dns.ManagedZone, change *dns.Change) map[string]*dns.Change {
	changes := make(map[string]*dns.Change)

	for _, z := range zones {
		changes[z.Name] = &dns.Change{
			Additions: []*dns.ResourceRecordSet{},
			Deletions: []*dns.ResourceRecordSet{},
		}

		for _, a := range change.Additions {
			if strings.HasSuffix(ensureTrailingDot(a.Name), z.DnsName) {
				changes[z.Name].Additions = append(changes[z.Name].Additions, a)
			}
		}

		for _, d := range change.Deletions {
			if strings.HasSuffix(ensureTrailingDot(d.Name), z.DnsName) {
				changes[z.Name].Deletions = append(changes[z.Name].Deletions, d)
			}
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
func newRecord(endpoint *endpoint.Endpoint) *dns.ResourceRecordSet {
	return &dns.ResourceRecordSet{
		Name:    ensureTrailingDot(endpoint.DNSName),
		Rrdatas: []string{ensureTrailingDot(endpoint.Target)},
		Ttl:     300,
		Type:    suitableType(endpoint.Target),
	}
}

// ensureTrailingDot ensures that the hostname receives a trailing dot if it hasn't already.
func ensureTrailingDot(hostname string) string {
	if net.ParseIP(hostname) != nil {
		return hostname
	}

	return strings.TrimSuffix(hostname, ".") + "."
}

// isNotFound returns true if a given error is due to a resource not being found.
func isNotFound(err error) bool {
	return strings.Contains(err.Error(), "notFound")
}
