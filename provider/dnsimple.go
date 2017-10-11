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
	"strconv"
	"strings"

	"github.com/dnsimple/dnsimple-go/dnsimple"
	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	log "github.com/sirupsen/logrus"
)

type identityService struct {
	service *dnsimple.IdentityService
}

func (i identityService) Whoami() (*dnsimple.WhoamiResponse, error) {
	return i.service.Whoami()
}

// Returns the account ID given dnsimple credentials
func (p *dnsimpleProvider) GetAccountID(credentials dnsimple.Credentials, client dnsimple.Client) (accountID string, err error) {
	// get DNSimple client accountID
	whoamiResponse, err := client.Identity.Whoami()
	if err != nil {
		return "", err
	}
	return strconv.Itoa(whoamiResponse.Data.Account.ID), nil
}

// dnsimpleZoneServiceInterface is an interface that contains all necessary zone services from dnsimple
type dnsimpleZoneServiceInterface interface {
	ListZones(accountID string, options *dnsimple.ZoneListOptions) (*dnsimple.ZonesResponse, error)
	ListRecords(accountID string, zoneID string, options *dnsimple.ZoneRecordListOptions) (*dnsimple.ZoneRecordsResponse, error)
	CreateRecord(accountID string, zoneID string, recordAttributes dnsimple.ZoneRecord) (*dnsimple.ZoneRecordResponse, error)
	DeleteRecord(accountID string, zoneID string, recordID int) (*dnsimple.ZoneRecordResponse, error)
	UpdateRecord(accountID string, zoneID string, recordID int, recordAttributes dnsimple.ZoneRecord) (*dnsimple.ZoneRecordResponse, error)
}

type dnsimpleZoneService struct {
	service *dnsimple.ZonesService
}

func (z dnsimpleZoneService) ListZones(accountID string, options *dnsimple.ZoneListOptions) (*dnsimple.ZonesResponse, error) {
	return z.service.ListZones(accountID, options)
}

func (z dnsimpleZoneService) ListRecords(accountID string, zoneID string, options *dnsimple.ZoneRecordListOptions) (*dnsimple.ZoneRecordsResponse, error) {
	return z.service.ListRecords(accountID, zoneID, options)
}

func (z dnsimpleZoneService) CreateRecord(accountID string, zoneID string, recordAttributes dnsimple.ZoneRecord) (*dnsimple.ZoneRecordResponse, error) {
	return z.service.CreateRecord(accountID, zoneID, recordAttributes)
}

func (z dnsimpleZoneService) DeleteRecord(accountID string, zoneID string, recordID int) (*dnsimple.ZoneRecordResponse, error) {
	return z.service.DeleteRecord(accountID, zoneID, recordID)
}

func (z dnsimpleZoneService) UpdateRecord(accountID string, zoneID string, recordID int, recordAttributes dnsimple.ZoneRecord) (*dnsimple.ZoneRecordResponse, error) {
	return z.service.UpdateRecord(accountID, zoneID, recordID, recordAttributes)
}

type dnsimpleProvider struct {
	client       dnsimpleZoneServiceInterface
	identity     identityService
	accountID    string
	domainFilter DomainFilter
	dryRun       bool
}

type dnsimpleChange struct {
	Action            string
	ResourceRecordSet dnsimple.ZoneRecord
}

const (
	dnsimpleCreate = "CREATE"
	dnsimpleDelete = "DELETE"
	dnsimpleUpdate = "UPDATE"
)

// NewDnsimpleProvider initializes a new Dnsimple based provider
func NewDnsimpleProvider(domainFilter DomainFilter, dryRun bool) (Provider, error) {
	oauthToken := os.Getenv("DNSIMPLE_OAUTH")
	if len(oauthToken) == 0 {
		return nil, fmt.Errorf("No dnsimple oauth token provided")
	}
	client := dnsimple.NewClient(dnsimple.NewOauthTokenCredentials(oauthToken))
	provider := &dnsimpleProvider{
		client:       dnsimpleZoneService{service: client.Zones},
		identity:     identityService{service: client.Identity},
		domainFilter: domainFilter,
		dryRun:       dryRun,
	}
	whoamiResponse, err := provider.identity.service.Whoami()
	if err != nil {
		return nil, err
	}
	provider.accountID = strconv.Itoa(whoamiResponse.Data.Account.ID)
	return provider, nil
}

// Returns a list of Zones that end with the provider's domainFilter
func (p *dnsimpleProvider) Zones() (map[string]dnsimple.Zone, error) {
	zones := make(map[string]dnsimple.Zone)
	zonesResponse, err := p.client.ListZones(p.accountID, &dnsimple.ZoneListOptions{})
	if err != nil {
		return nil, err
	}
	for _, zone := range zonesResponse.Data {
		if p.domainFilter.Match(zone.Name) {
			zones[strconv.Itoa(zone.ID)] = zone
		}
	}
	return zones, nil
}

// Records retuns a list of endpoints in a given zone
func (p *dnsimpleProvider) Records() (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.Zones()
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		records, err := p.client.ListRecords(p.accountID, zone.Name, &dnsimple.ZoneRecordListOptions{})
		if err != nil {
			return nil, err
		}
		for _, record := range records.Data {
			switch record.Type {
			case "A", "CNAME", "TXT":
				break
			default:
				continue
			}
			endpoints = append(endpoints, endpoint.NewEndpoint(record.Name+"."+record.ZoneID, record.Content, record.Type))
		}
	}
	return endpoints, nil
}

// newDnsimpleChange initializes a new change to dns records
func newDnsimpleChange(action string, e *endpoint.Endpoint) *dnsimpleChange {
	change := &dnsimpleChange{
		Action: action,
		ResourceRecordSet: dnsimple.ZoneRecord{
			Name:    e.DNSName,
			Type:    e.RecordType,
			Content: e.Target,
		},
	}
	return change
}

// newDnsimpleChanges returns a slice of changes based on given action and record
func newDnsimpleChanges(action string, endpoints []*endpoint.Endpoint) []*dnsimpleChange {
	changes := make([]*dnsimpleChange, 0, len(endpoints))
	for _, e := range endpoints {
		changes = append(changes, newDnsimpleChange(action, e))
	}
	return changes
}

// submitChanges takes a zone and a collection of changes and makes all changes from the collection
func (p *dnsimpleProvider) submitChanges(changes []*dnsimpleChange) error {
	if len(changes) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}
	zones, err := p.Zones()
	if err != nil {
		return err
	}
	for _, change := range changes {
		zone := dnsimpleSuitableZone(change.ResourceRecordSet.Name, zones)
		if zone.ID == 0 {
			return fmt.Errorf("No suitable zone name found")
		}
		log.Infof("Changing records: %s %v in zone: %s", change.Action, change.ResourceRecordSet, zone.Name)

		change.ResourceRecordSet.Name = strings.TrimSuffix(change.ResourceRecordSet.Name, "."+zone.Name)
		if !p.dryRun {
			switch change.Action {
			case dnsimpleCreate:
				_, err := p.client.CreateRecord(p.accountID, zone.Name, change.ResourceRecordSet)
				if err != nil {
					return err
				}
			case dnsimpleDelete:
				recordID, err := p.GetRecordID(zone.Name, change.ResourceRecordSet.Name)
				if err != nil {
					return err
				}
				_, err = p.client.DeleteRecord(p.accountID, zone.Name, recordID)
				if err != nil {
					return err
				}
			case dnsimpleUpdate:
				recordID, err := p.GetRecordID(zone.Name, change.ResourceRecordSet.Name)
				if err != nil {
					return err
				}
				_, err = p.client.UpdateRecord(p.accountID, zone.Name, recordID, change.ResourceRecordSet)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Returns the record ID for a given record name and zone
func (p *dnsimpleProvider) GetRecordID(zone string, recordName string) (recordID int, err error) {
	records, err := p.client.ListRecords(p.accountID, zone, &dnsimple.ZoneRecordListOptions{})
	if err != nil {
		return 0, err
	}
	for _, record := range records.Data {
		if record.Name == recordName {
			return record.ID, nil
		}
	}
	return 0, fmt.Errorf("No record id found")
}

// dnsimpleSuitableZone returns the most suitable zone for a given hostname and a set of zones.
func dnsimpleSuitableZone(hostname string, zones map[string]dnsimple.Zone) dnsimple.Zone {
	var zone dnsimple.Zone
	for _, z := range zones {
		if strings.HasSuffix(hostname, z.Name) {
			if zone.ID == 0 || len(z.Name) > len(zone.Name) {
				zone = z
			}
		}
	}
	return zone
}

// CreateRecords creates records for a given slice of endpoints
func (p *dnsimpleProvider) CreateRecords(endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(newDnsimpleChanges(dnsimpleCreate, endpoints))
}

// DeleteRecords deletes records for a given slice of endpoints
func (p *dnsimpleProvider) DeleteRecords(endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(newDnsimpleChanges(dnsimpleDelete, endpoints))
}

// UpdateRecords updates records for a given slice of endpoints
func (p *dnsimpleProvider) UpdateRecords(endpoints []*endpoint.Endpoint) error {
	return p.submitChanges(newDnsimpleChanges(dnsimpleUpdate, endpoints))
}

// ApplyChanges applies a given set of changes
func (p *dnsimpleProvider) ApplyChanges(changes *plan.Changes) error {
	combinedChanges := make([]*dnsimpleChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newDnsimpleChanges(dnsimpleCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newDnsimpleChanges(dnsimpleUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newDnsimpleChanges(dnsimpleDelete, changes.Delete)...)

	return p.submitChanges(combinedChanges)
}
