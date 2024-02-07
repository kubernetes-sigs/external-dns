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

package dnsimple

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/dnsimple/dnsimple-go/dnsimple"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const dnsimpleRecordTTL = 3600 // Default TTL of 1 hour if not set (DNSimple's default)

type dnsimpleIdentityService struct {
	service *dnsimple.IdentityService
}

func (i dnsimpleIdentityService) Whoami(ctx context.Context) (*dnsimple.WhoamiResponse, error) {
	return i.service.Whoami(ctx)
}

// dnsimpleZoneServiceInterface is an interface that contains all necessary zone services from DNSimple
type dnsimpleZoneServiceInterface interface {
	ListZones(ctx context.Context, accountID string, options *dnsimple.ZoneListOptions) (*dnsimple.ZonesResponse, error)
	ListRecords(ctx context.Context, accountID string, zoneID string, options *dnsimple.ZoneRecordListOptions) (*dnsimple.ZoneRecordsResponse, error)
	CreateRecord(ctx context.Context, accountID string, zoneID string, recordAttributes dnsimple.ZoneRecordAttributes) (*dnsimple.ZoneRecordResponse, error)
	DeleteRecord(ctx context.Context, accountID string, zoneID string, recordID int64) (*dnsimple.ZoneRecordResponse, error)
	UpdateRecord(ctx context.Context, accountID string, zoneID string, recordID int64, recordAttributes dnsimple.ZoneRecordAttributes) (*dnsimple.ZoneRecordResponse, error)
}

type dnsimpleZoneService struct {
	service *dnsimple.ZonesService
}

func (z dnsimpleZoneService) ListZones(ctx context.Context, accountID string, options *dnsimple.ZoneListOptions) (*dnsimple.ZonesResponse, error) {
	return z.service.ListZones(ctx, accountID, options)
}

func (z dnsimpleZoneService) ListRecords(ctx context.Context, accountID string, zoneID string, options *dnsimple.ZoneRecordListOptions) (*dnsimple.ZoneRecordsResponse, error) {
	return z.service.ListRecords(ctx, accountID, zoneID, options)
}

func (z dnsimpleZoneService) CreateRecord(ctx context.Context, accountID string, zoneID string, recordAttributes dnsimple.ZoneRecordAttributes) (*dnsimple.ZoneRecordResponse, error) {
	return z.service.CreateRecord(ctx, accountID, zoneID, recordAttributes)
}

func (z dnsimpleZoneService) DeleteRecord(ctx context.Context, accountID string, zoneID string, recordID int64) (*dnsimple.ZoneRecordResponse, error) {
	return z.service.DeleteRecord(ctx, accountID, zoneID, recordID)
}

func (z dnsimpleZoneService) UpdateRecord(ctx context.Context, accountID string, zoneID string, recordID int64, recordAttributes dnsimple.ZoneRecordAttributes) (*dnsimple.ZoneRecordResponse, error) {
	return z.service.UpdateRecord(ctx, accountID, zoneID, recordID, recordAttributes)
}

type dnsimpleProvider struct {
	provider.BaseProvider
	client       dnsimpleZoneServiceInterface
	identity     dnsimpleIdentityService
	accountID    string
	domainFilter endpoint.DomainFilter
	zoneIDFilter provider.ZoneIDFilter
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
func NewDnsimpleProvider(domainFilter endpoint.DomainFilter, zoneIDFilter provider.ZoneIDFilter, dryRun bool) (provider.Provider, error) {
	oauthToken := os.Getenv("DNSIMPLE_OAUTH")
	if len(oauthToken) == 0 {
		return nil, fmt.Errorf("no dnsimple oauth token provided")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: oauthToken})
	tc := oauth2.NewClient(context.Background(), ts)

	client := dnsimple.NewClient(tc)
	client.SetUserAgent(fmt.Sprintf("Kubernetes ExternalDNS/%s", externaldns.Version))

	provider := &dnsimpleProvider{
		client:       dnsimpleZoneService{service: client.Zones},
		identity:     dnsimpleIdentityService{service: client.Identity},
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		dryRun:       dryRun,
	}

	whoamiResponse, err := provider.identity.Whoami(context.Background())
	if err != nil {
		return nil, err
	}
	provider.accountID = int64ToString(whoamiResponse.Data.Account.ID)
	return provider, nil
}

// GetAccountID returns the account ID given DNSimple credentials.
func (p *dnsimpleProvider) GetAccountID(ctx context.Context) (accountID string, err error) {
	// get DNSimple client accountID
	whoamiResponse, err := p.identity.Whoami(ctx)
	if err != nil {
		return "", err
	}
	return int64ToString(whoamiResponse.Data.Account.ID), nil
}

// Returns a list of filtered Zones
func (p *dnsimpleProvider) Zones(ctx context.Context) (map[string]dnsimple.Zone, error) {
	zones := make(map[string]dnsimple.Zone)
	page := 1
	listOptions := &dnsimple.ZoneListOptions{}
	for {
		listOptions.Page = &page
		zonesResponse, err := p.client.ListZones(ctx, p.accountID, listOptions)
		if err != nil {
			return nil, err
		}
		for _, zone := range zonesResponse.Data {
			if !p.domainFilter.Match(zone.Name) {
				continue
			}

			if !p.zoneIDFilter.Match(int64ToString(zone.ID)) {
				continue
			}

			zones[int64ToString(zone.ID)] = zone
		}

		page++
		if page > zonesResponse.Pagination.TotalPages {
			break
		}
	}
	return zones, nil
}

// Records returns a list of endpoints in a given zone
func (p *dnsimpleProvider) Records(ctx context.Context) (endpoints []*endpoint.Endpoint, _ error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		page := 1
		listOptions := &dnsimple.ZoneRecordListOptions{}
		for {
			listOptions.Page = &page
			records, err := p.client.ListRecords(ctx, p.accountID, zone.Name, listOptions)
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
				// Apex records have an empty string for their name.
				// Consider this when creating the endpoint dnsName
				dnsName := fmt.Sprintf("%s.%s", record.Name, record.ZoneID)
				if record.Name == "" {
					dnsName = record.ZoneID
				}
				endpoints = append(endpoints, endpoint.NewEndpointWithTTL(dnsName, record.Type, endpoint.TTL(record.TTL), record.Content))
			}
			page++
			if page > records.Pagination.TotalPages {
				break
			}
		}
	}
	return endpoints, nil
}

// newDnsimpleChange initializes a new change to dns records
func newDnsimpleChange(action string, e *endpoint.Endpoint) *dnsimpleChange {
	ttl := dnsimpleRecordTTL
	if e.RecordTTL.IsConfigured() {
		ttl = int(e.RecordTTL)
	}

	change := &dnsimpleChange{
		Action: action,
		ResourceRecordSet: dnsimple.ZoneRecord{
			Name:    e.DNSName,
			Type:    e.RecordType,
			Content: e.Targets[0],
			TTL:     ttl,
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
func (p *dnsimpleProvider) submitChanges(ctx context.Context, changes []*dnsimpleChange) error {
	if len(changes) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}
	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}
	for _, change := range changes {
		zone := dnsimpleSuitableZone(change.ResourceRecordSet.Name, zones)
		if zone == nil {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", change.ResourceRecordSet.Name)
			continue
		}

		log.Infof("Changing records: %s %v in zone: %s", change.Action, change.ResourceRecordSet, zone.Name)

		if change.ResourceRecordSet.Name == zone.Name {
			change.ResourceRecordSet.Name = "" // Apex records have an empty name
		} else {
			change.ResourceRecordSet.Name = strings.TrimSuffix(change.ResourceRecordSet.Name, fmt.Sprintf(".%s", zone.Name))
		}

		recordAttributes := dnsimple.ZoneRecordAttributes{
			Name:    &change.ResourceRecordSet.Name,
			Type:    change.ResourceRecordSet.Type,
			Content: change.ResourceRecordSet.Content,
			TTL:     change.ResourceRecordSet.TTL,
		}

		if !p.dryRun {
			switch change.Action {
			case dnsimpleCreate:
				_, err := p.client.CreateRecord(ctx, p.accountID, zone.Name, recordAttributes)
				if err != nil {
					return err
				}
			case dnsimpleDelete:
				recordID, err := p.GetRecordID(ctx, zone.Name, *recordAttributes.Name)
				if err != nil {
					return err
				}
				_, err = p.client.DeleteRecord(ctx, p.accountID, zone.Name, recordID)
				if err != nil {
					return err
				}
			case dnsimpleUpdate:
				recordID, err := p.GetRecordID(ctx, zone.Name, *recordAttributes.Name)
				if err != nil {
					return err
				}
				_, err = p.client.UpdateRecord(ctx, p.accountID, zone.Name, recordID, recordAttributes)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// GetRecordID returns the record ID for a given record name and zone.
func (p *dnsimpleProvider) GetRecordID(ctx context.Context, zone string, recordName string) (recordID int64, err error) {
	page := 1
	listOptions := &dnsimple.ZoneRecordListOptions{Name: &recordName}
	for {
		listOptions.Page = &page
		records, err := p.client.ListRecords(ctx, p.accountID, zone, listOptions)
		if err != nil {
			return 0, err
		}

		for _, record := range records.Data {
			if record.Name == recordName {
				return record.ID, nil
			}
		}

		page++
		if page > records.Pagination.TotalPages {
			break
		}
	}
	return 0, fmt.Errorf("no record id found")
}

// dnsimpleSuitableZone returns the most suitable zone for a given hostname and a set of zones.
func dnsimpleSuitableZone(hostname string, zones map[string]dnsimple.Zone) *dnsimple.Zone {
	var zone *dnsimple.Zone
	for _, z := range zones {
		if strings.HasSuffix(hostname, z.Name) {
			if zone == nil || len(z.Name) > len(zone.Name) {
				newZ := z
				zone = &newZ
			}
		}
	}
	return zone
}

// ApplyChanges applies a given set of changes
func (p *dnsimpleProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	combinedChanges := make([]*dnsimpleChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newDnsimpleChanges(dnsimpleCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newDnsimpleChanges(dnsimpleUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newDnsimpleChanges(dnsimpleDelete, changes.Delete)...)

	return p.submitChanges(ctx, combinedChanges)
}

func int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}
