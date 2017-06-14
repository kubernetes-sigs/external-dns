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
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/cloudflare/cloudflare-go"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

// cloudFlareDNS is the subset of the CloudFlare API that we actually use.  Add methods as required. Signatures must match exactly.
type cloudFlareDNS interface {
	UserDetails() (cloudflare.User, error)
	ZoneIDByName(zoneName string) (string, error)
	ListZones(zoneID ...string) ([]cloudflare.Zone, error)
	DNSRecords(zoneID string, rr cloudflare.DNSRecord) ([]cloudflare.DNSRecord, error)
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

// cloudFlareProvider is an implementation of Provider for CloudFlare DNS.
type cloudFlareProvider struct {
	Client cloudFlareDNS
	// only consider hosted zones managing domains ending in this suffix
	domainFilter string
	DryRun       bool
}

// cloudFlareChange differentiates between ChangActions
type cloudFlareChange struct {
	Action            string
	ResourceRecordSet cloudflare.DNSRecord
}

const (
	// cloudFlareCreate is a ChangeAction enum value
	cloudFlareCreate = "CREATE"
	// cloudFlareDelete is a ChangeAction enum value
	cloudFlareDelete = "DELETE"
	// cloudFlareUpdate is a ChangeAction enum value
	cloudFlareUpdate = "UPDATE"
)

// NewCloudFlareProvider initializes a new CloudFlare DNS based Provider.
func NewCloudFlareProvider(domainFilter string, dryRun bool) (Provider, error) {
	// initialize via API email and API key and returns new API object
	config, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		return nil, err
	}
	provider := &cloudFlareProvider{
		//Client: config,
		Client:       zoneService{config},
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}
	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *cloudFlareProvider) Zones() ([]cloudflare.Zone, error) {
	result := []cloudflare.Zone{}

	zones, err := p.Client.ListZones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		if strings.HasSuffix(zone.Name, p.domainFilter) {
			result = append(result, zone)
		}
	}

	return result, nil
}

// Records returns the list of records.
func (p *cloudFlareProvider) Records() ([]*endpoint.Endpoint, error) {
	zones, err := p.Zones()
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
	for _, zone := range zones {
		records, err := p.Client.DNSRecords(zone.ID, cloudflare.DNSRecord{})
		if err != nil {
			return nil, err
		}

		for _, r := range records {
			switch r.Type {
			case "A", "CNAME", "TXT":
				break
			default:
				continue
			}
			endpoints = append(endpoints, endpoint.NewEndpoint(r.Name, r.Content, r.Type))
		}
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *cloudFlareProvider) ApplyChanges(changes *plan.Changes) error {
	combinedChanges := make([]*cloudFlareChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareDelete, changes.Delete)...)

	return p.submitChanges(combinedChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *cloudFlareProvider) submitChanges(changes []*cloudFlareChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	zones, err := p.Zones()
	if err != nil {
		return err
	}
	// separate into per-zone change sets to be passed to the API.
	changesByZone := cloudflareChangesByZone(zones, changes)

	for zoneID, changes := range changesByZone {
		records, err := p.Client.DNSRecords(zoneID, cloudflare.DNSRecord{})
		if err != nil {
			return fmt.Errorf("could not fetch records from zone, %v", err)
		}
		for _, change := range changes {
			logFields := log.Fields{
				"record": change.ResourceRecordSet.Name,
				"type":   change.ResourceRecordSet.Type,
				"action": change.Action,
				"zone":   zoneID,
			}

			log.WithFields(logFields).Info("Changing record.")

			if p.DryRun {
				continue
			}
			recordID := p.getRecordID(records, zoneID, change.ResourceRecordSet)
			switch change.Action {
			case cloudFlareCreate:
				_, err := p.Client.CreateDNSRecord(zoneID, change.ResourceRecordSet)
				if err != nil {
					log.WithFields(logFields).Errorf("failed to create record: %v", err)
				}
			case cloudFlareDelete:
				err := p.Client.DeleteDNSRecord(zoneID, recordID)
				if err != nil {
					log.WithFields(logFields).Errorf("failed to delete record: %v", err)
				}
			case cloudFlareUpdate:
				err := p.Client.UpdateDNSRecord(zoneID, recordID, change.ResourceRecordSet)
				if err != nil {
					log.WithFields(logFields).Errorf("failed to update record: %v", err)
				}
			}
		}
	}
	return nil
}

// changesByZone separates a multi-zone change into a single change per zone.
func cloudflareChangesByZone(zones []cloudflare.Zone, changeSet []*cloudFlareChange) map[string][]*cloudFlareChange {
	changes := make(map[string][]*cloudFlareChange)

	for _, z := range zones {
		changes[z.ID] = []*cloudFlareChange{}
	}

	for _, c := range changeSet {
		zone := cloudflareSuitableZone(c.ResourceRecordSet.Name, zones)
		if zone == nil {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected ", c.ResourceRecordSet.Name)
			continue
		}
		changes[zone.ID] = append(changes[zone.ID], c)
	}

	return changes
}

// cloudflareSuitableZone returns the most suitable zone for a given hostname
// and a set of zones.
func cloudflareSuitableZone(hostname string, zones []cloudflare.Zone) *cloudflare.Zone {
	var zone cloudflare.Zone
	for _, z := range zones {
		if strings.HasSuffix(hostname, z.Name) {
			if len(z.Name) > len(zone.Name) {
				zone = z
			}
		}
	}
	return &zone
}

func (p *cloudFlareProvider) getRecordID(records []cloudflare.DNSRecord, zoneID string, record cloudflare.DNSRecord) string {
	for _, zoneRecord := range records {
		if zoneRecord.Name == record.Name && zoneRecord.Type == record.Type {
			return zoneRecord.ID
		}
	}
	return ""
}

// newCloudFlareChanges returns a collection of Changes based on the given records and action.
func newCloudFlareChanges(action string, endpoints []*endpoint.Endpoint) []*cloudFlareChange {
	changes := make([]*cloudFlareChange, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newCloudFlareChange(action, endpoint))
	}

	return changes
}

func newCloudFlareChange(action string, endpoint *endpoint.Endpoint) *cloudFlareChange {
	typ := suitableType(endpoint)

	return &cloudFlareChange{
		Action: action,
		ResourceRecordSet: cloudflare.DNSRecord{
			Name: endpoint.DNSName,
			// TTL Value of 1 is 'automatic'
			TTL:     1,
			Proxied: false,
			Type:    typ,
			Content: endpoint.Target,
		},
	}
}
