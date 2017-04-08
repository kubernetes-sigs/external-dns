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
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/cloudflare/cloudflare-go"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// CloudFlareDNSInterface is the subset of the CloudFlare API that we actually use.  Add methods as required. Signatures must match exactly.
type CloudFlareDNSInterface interface {
	UserDetails() (cloudflare.User, error)
	ZoneIDByName(zoneName string) (string, error)
	Zones() ([]cloudflare.Zone, error)
	Zone(zoneID ...string) ([]cloudflare.Zone, error)
	ListZones(zoneID ...string) ([]cloudflare.Zone, error)
	DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error)
	CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error)
	DeleteZone(zoneID string) (cloudflare.ZoneID, error)
	CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error)
	DeleteDNSRecord(zoneID, recordID string) error
	UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error
}

type zoneService struct {
	service *cloudflare.API
}

func (z zoneService) UserDetails() (cloudflare.User, error) {
	return z.service.UserDetails()
}

func (z zoneService) Zones() ([]cloudflare.Zone, error) {
	return z.service.ListZones()
}

func (z zoneService) Zone(zoneID ...string) ([]cloudflare.Zone, error) {
	return z.service.ListZones(zoneID...)
}

func (z zoneService) ListZones(zoneID ...string) ([]cloudflare.Zone, error) {
	return z.service.ListZones(zoneID...)
}

func (z zoneService) ZoneIDByName(zoneName string) (string, error) {
	return z.service.ZoneIDByName(zoneName)
}

func (z zoneService) CreateDNSRecord(zoneID string, rr cloudflare.DNSRecord) (*cloudflare.DNSRecordResponse, error) {
	return z.service.CreateDNSRecord(zoneID, rr)
}

func (z zoneService) DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error) {
	return z.service.DNSRecords(zoneID, rr)
}
func (z zoneService) UpdateDNSRecord(zoneID, recordID string, rr cloudflare.DNSRecord) error {
	return z.service.UpdateDNSRecord(zoneID, recordID, rr)
}
func (z zoneService) DeleteDNSRecord(zoneID, recordID string) error {
	return z.service.DeleteDNSRecord(zoneID, recordID)
}
func (z zoneService) CreateZone(name string, jumpstart bool, org cloudflare.Organization) (cloudflare.Zone, error) {
	return z.service.CreateZone(name, jumpstart, org)
}

func (z zoneService) DeleteZone(zoneID string) (cloudflare.ZoneID, error) {
	return z.service.DeleteZone(zoneID)
}

// CloudFlareProvider is an implementation of Provider for CloudFlare DNS.
type CloudFlareProvider struct {
	Client CloudFlareDNSInterface
	DryRun bool
}

// CloudFlareChange differentiates between ChangActions
type CloudFlareChange struct {
	Action            string
	ResourceRecordSet cloudflare.DNSRecord
}

const (
	// CloudFlareCreate is a ChangeAction enum value
	CloudFlareCreate = "CREATE"
	// CloudFlareDelete is a ChangeAction enum value
	CloudFlareDelete = "DELETE"
	// CloudFlareUpdate is a ChangeAction enum value
	CloudFlareUpdate = "UPDATE"
)

// NewCloudFlareProvider initializes a new CloudFlare DNS based Provider.
func NewCloudFlareProvider(dryRun bool) (Provider, error) {
	// initialize via API email and API key and returns new API object
	config, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		return nil, err
	}
	provider := &CloudFlareProvider{
		//Client: config,
		Client: zoneService{config},
		DryRun: dryRun,
	}
	return provider, nil
}

// Zones returns the list of zones.
func (p *CloudFlareProvider) Zones() ([]cloudflare.Zone, error) {
	zones, err := p.Client.ListZones()
	if err != nil {
		return []cloudflare.Zone{}, nil
	}
	return zones, nil
}

// Zone returns a single zone given a DNS name.
func (p *CloudFlareProvider) Zone(name string) ([]cloudflare.Zone, error) {
	zoneID, err := p.Client.ZoneIDByName(name)
	if err != nil {
		return []cloudflare.Zone{}, err
	}
	zones, err := p.Client.ListZones(zoneID)
	if err != nil {
		return []cloudflare.Zone{}, nil
	}
	return zones, nil
}

// CreateZone creates a hosted zone given a name.
func (p *CloudFlareProvider) CreateZone(name string) (*cloudflare.Zone, error) {
	user, err := p.Client.UserDetails()
	if err != nil {
		return nil, err
	}
	zone, err := p.Client.CreateZone(name, true, cloudflare.Organization{ID: user.ID})
	if err != nil {
		return nil, err
	}

	return &zone, nil
}

// DeleteZone deletes a hosted zone given a name.
func (p *CloudFlareProvider) DeleteZone(name string) (cloudflare.ZoneID, error) {
	zoneID, err := p.Client.ZoneIDByName(name)
	if err != nil {
		return cloudflare.ZoneID{}, err
	}
	zone, err := p.Client.DeleteZone(zoneID)
	if err != nil {
		return cloudflare.ZoneID{}, err
	}
	return zone, nil
}

// Records returns the list of records in a given zone.
func (p *CloudFlareProvider) Records(zone string) ([]*endpoint.Endpoint, error) {
	var record cloudflare.DNSRecord
	zoneID, err := p.Client.ZoneIDByName(zone)
	if err != nil {
		return nil, err
	}
	records, err := p.Client.DNSRecords(zoneID, record)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
	for _, r := range records {
		endpoints = append(endpoints, endpoint.NewEndpoint(r.Name, r.Content))
	}

	return endpoints, nil
}

// CreateRecords creates a given set of DNS records in the given hosted zone.
func (p *CloudFlareProvider) CreateRecords(zone string, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newCloudFlareChanges(CloudFlareCreate, endpoints))
}

// UpdateRecords updates a given set of old records to a new set of records in a given hosted zone.
func (p *CloudFlareProvider) UpdateRecords(zone string, endpoints, _ []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newCloudFlareChanges(CloudFlareUpdate, endpoints))
}

// DeleteRecords deletes a given set of DNS records in a given zone.
func (p *CloudFlareProvider) DeleteRecords(zone string, endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(zone, newCloudFlareChanges(CloudFlareDelete, endpoints))
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *CloudFlareProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	combinedChanges := make([]*CloudFlareChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newCloudFlareChanges(CloudFlareCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(CloudFlareUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(CloudFlareDelete, changes.Delete)...)

	return p.submitChanges(zone, combinedChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *CloudFlareProvider) submitChanges(zone string, changes []*CloudFlareChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	zoneID, err := p.Client.ZoneIDByName(zone)
	if err != nil {
		return err
	}

	if p.DryRun {
		for _, change := range changes {
			log.Infof("Changing records: %s %+v", change.Action, change.ResourceRecordSet)
		}

		return nil
	}
	for _, change := range changes {
		switch {
		case change.Action == CloudFlareCreate:
			p.Client.CreateDNSRecord(zoneID, change.ResourceRecordSet)
		case change.Action == CloudFlareDelete:
			recordID, err := p.getRecordID(zoneID, change.ResourceRecordSet)
			if err != nil {
				return err
			}
			p.Client.DeleteDNSRecord(zoneID, recordID)
		case change.Action == CloudFlareUpdate:
			recordID, err := p.getRecordID(zoneID, change.ResourceRecordSet)
			if err != nil {
				return err
			}
			p.Client.UpdateDNSRecord(zoneID, recordID, change.ResourceRecordSet)
		}
	}
	return nil
}

func (p *CloudFlareProvider) getRecordID(zoneID string, record cloudflare.DNSRecord) (string, error) {
	records := cloudflare.DNSRecord{}
	zoneRecords, err := p.Client.DNSRecords(zoneID, records)
	if err != nil {
		return "", err
	}
	for _, zoneRecord := range zoneRecords {
		if zoneRecord.Name == record.Name {
			return zoneRecord.ID, nil
		}
	}
	return "", fmt.Errorf("No record id found")
}

// newCloudFlareChanges returns a collection of Changes based on the given records and action.
func newCloudFlareChanges(action string, endpoints []*endpoint.Endpoint) []*CloudFlareChange {
	changes := make([]*CloudFlareChange, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newCloudFlareChange(action, endpoint))
	}

	return changes
}

func newCloudFlareChange(action string, endpoint *endpoint.Endpoint) *CloudFlareChange {
	change := &CloudFlareChange{
		Action: action,
		ResourceRecordSet: cloudflare.DNSRecord{
			Name: endpoint.DNSName,
			//TTL Value of 1 is 'automatic'
			TTL: 1,
			//record is receiving the performance and security benefits of Cloudflare
			Proxied: true,
			Type:    suitableType(endpoint.Target),
			Content: endpoint.Target,
		},
	}
	return change
}
