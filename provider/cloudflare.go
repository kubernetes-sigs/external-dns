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
	"encoding/json"
	"fmt"
	"os"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
	"github.com/kubernetes-incubator/external-dns/srv"
)

const (
	// cloudFlareCreate is a ChangeAction enum value
	cloudFlareCreate = "CREATE"
	// cloudFlareDelete is a ChangeAction enum value
	cloudFlareDelete = "DELETE"
	// cloudFlareUpdate is a ChangeAction enum value
	cloudFlareUpdate = "UPDATE"
	// defaultCloudFlareRecordTTL 1 = automatic
	defaultCloudFlareRecordTTL = 1
)

var cloudFlareTypeNotSupported = map[string]bool{
	"LOC": true,
	"MX":  true,
	"NS":  true,
	"SPF": true,
	"TXT": true,
	"SRV": true,
}

// cloudFlareSRVData is attached to a DNSRecord to create an SRV record.
type cloudFlareSRVData struct {
	// Service is the service name e.g. _http
	Service string `json:"service"`
	// Proto is the protocol type e.g. _tcp
	Proto string `json:"proto"`
	// Name is the DNS domain name
	Name string `json:"name"`
	// Priority is the priority to select the record
	Priority int `json:"priority"`
	// Weight is the weighting of the endpoint
	Weight int `json:"weight"`
	// Port is the port to dial e.g. 443
	Port int `json:"port"`
	// Target is the target A or CNAME record
	Target string `json:"target"`
}

// Equal compares two SRV data records and returns true if the target portion
// of the records match
func (d *cloudFlareSRVData) Equal(o *cloudFlareSRVData) bool {
	return d.Priority == o.Priority && d.Weight == o.Weight && d.Port == o.Port && d.Target == o.Target
}

// unmarshalCloudFlareSRVData takes raw input from the API and decodes it into
// the internal datatype.
func unmarshalCloudFlareSRVData(data interface{}) *cloudFlareSRVData {
	tmp, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	d := &cloudFlareSRVData{}
	if err := json.Unmarshal(tmp, d); err != nil {
		return nil
	}

	return d
}

// newCloudFlareSRVEndpoint takes raw API SRV data and returns an endpoint
// used for planning.
func newCloudFlareSRVEndpoint(rr cloudflare.DNSRecord) *endpoint.Endpoint {
	d := unmarshalCloudFlareSRVData(rr.Data)

	in := srv.Record{
		Name: srv.ParseName(rr.Name),
		Target: srv.Target{
			Port:     d.Port,
			Priority: d.Priority,
			Target:   d.Target,
			Weight:   d.Weight,
		},
	}

	return endpoint.NewEndpointWithTTL(in.Name.Format(), endpoint.RecordTypeSRV, endpoint.TTL(rr.TTL), in.Target.Format())
}

func newEndpoint(rr cloudflare.DNSRecord) *endpoint.Endpoint {
	switch rr.Type {
	case endpoint.RecordTypeSRV:
		return newCloudFlareSRVEndpoint(rr)
	}

	return endpoint.NewEndpointWithTTL(rr.Name, rr.Type, endpoint.TTL(rr.TTL), rr.Content)
}

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

// CloudFlareProvider is an implementation of Provider for CloudFlare DNS.
type CloudFlareProvider struct {
	Client cloudFlareDNS
	// only consider hosted zones managing domains ending in this suffix
	domainFilter DomainFilter
	zoneIDFilter ZoneIDFilter
	proxied      bool
	DryRun       bool
}

// cloudFlareChange differentiates between ChangActions
type cloudFlareChange struct {
	Action            string
	ResourceRecordSet cloudflare.DNSRecord
}

// NewCloudFlareProvider initializes a new CloudFlare DNS based Provider.
func NewCloudFlareProvider(domainFilter DomainFilter, zoneIDFilter ZoneIDFilter, proxied bool, dryRun bool) (*CloudFlareProvider, error) {
	// initialize via API email and API key and returns new API object
	config, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudflare provider: %v", err)
	}
	provider := &CloudFlareProvider{
		//Client: config,
		Client:       zoneService{config},
		domainFilter: domainFilter,
		zoneIDFilter: zoneIDFilter,
		proxied:      proxied,
		DryRun:       dryRun,
	}
	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *CloudFlareProvider) Zones() ([]cloudflare.Zone, error) {
	result := []cloudflare.Zone{}

	zones, err := p.Client.ListZones()
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		if !p.domainFilter.Match(zone.Name) {
			continue
		}

		if !p.zoneIDFilter.Match(zone.ID) {
			continue
		}

		result = append(result, zone)
	}

	return result, nil
}

// Records returns the list of records.
func (p *CloudFlareProvider) Records() ([]*endpoint.Endpoint, error) {
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
			if supportedRecordType(r.Type) {
				endpoints = append(endpoints, newEndpoint(r))
			}
		}
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *CloudFlareProvider) ApplyChanges(changes *plan.Changes) error {
	combinedChanges := make([]*cloudFlareChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareCreate, changes.Create, p.proxied)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareUpdate, changes.UpdateNew, p.proxied)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareDelete, changes.Delete, p.proxied)...)

	return p.submitChanges(combinedChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *CloudFlareProvider) submitChanges(changes []*cloudFlareChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	zones, err := p.Zones()
	if err != nil {
		return err
	}
	// separate into per-zone change sets to be passed to the API.
	changesByZone := p.changesByZone(zones, changes)

	for zoneID, changes := range changesByZone {
		records, err := p.Client.DNSRecords(zoneID, cloudflare.DNSRecord{})
		if err != nil {
			return fmt.Errorf("could not fetch records from zone, %v", err)
		}
		for _, change := range changes {
			logFields := log.Fields{
				"record": change.ResourceRecordSet.Name,
				"type":   change.ResourceRecordSet.Type,
				"ttl":    change.ResourceRecordSet.TTL,
				"action": change.Action,
				"zone":   zoneID,
			}

			log.WithFields(logFields).Info("Changing record.")

			if p.DryRun {
				continue
			}

			recordID := p.getRecordID(records, change.ResourceRecordSet)
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
func (p *CloudFlareProvider) changesByZone(zones []cloudflare.Zone, changeSet []*cloudFlareChange) map[string][]*cloudFlareChange {
	changes := make(map[string][]*cloudFlareChange)
	zoneNameIDMapper := zoneIDName{}

	for _, z := range zones {
		zoneNameIDMapper.Add(z.ID, z.Name)
		changes[z.ID] = []*cloudFlareChange{}
	}

	for _, c := range changeSet {
		zoneID, _ := zoneNameIDMapper.FindZone(c.ResourceRecordSet.Name)
		if zoneID == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected ", c.ResourceRecordSet.Name)
			continue
		}
		changes[zoneID] = append(changes[zoneID], c)
	}

	return changes
}

func (p *CloudFlareProvider) getRecordID(records []cloudflare.DNSRecord, record cloudflare.DNSRecord) string {
	for _, zoneRecord := range records {
		if zoneRecord.Name == record.Name && zoneRecord.Type == record.Type {
			// SRV records need to check the full record for a match, including the target data
			if record.Type == endpoint.RecordTypeSRV {
				// The zone record is fresh from the API so decode the SRV data.
				zoneData := unmarshalCloudFlareSRVData(zoneRecord.Data)

				// The record will have been decoded by newCloudFlareChange
				recordData, ok := record.Data.(*cloudFlareSRVData)
				if !ok {
					continue
				}
				if zoneData.Equal(recordData) {
					return zoneRecord.ID
				}
				continue
			}
			// TXT records may have a target associated with them, which helps to disambiguate
			// when we have multiple records with the same name e.g. SRV and MX
			if record.Type == endpoint.RecordTypeTXT {
				recordLabels, err := endpoint.NewLabelsFromString(record.Content)
				if err != nil {
					continue
				}
				if ok, recordTarget := recordLabels.HasNonUniqueRecords(); ok {
					zoneLabels, err := endpoint.NewLabelsFromString(zoneRecord.Content)
					if err != nil {
						continue
					}
					if recordTarget == zoneLabels[endpoint.TargetLabel] {
						return zoneRecord.ID
					}
					continue
				}
			}
			return zoneRecord.ID
		}
	}
	return ""
}

// newCloudFlareChanges returns a collection of Changes based on the given records and action.
func newCloudFlareChanges(action string, endpoints []*endpoint.Endpoint, proxied bool) []*cloudFlareChange {
	changes := make([]*cloudFlareChange, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newCloudFlareChange(action, endpoint, proxied))
	}

	return changes
}

func newCloudFlareChange(action string, ep *endpoint.Endpoint, proxied bool) *cloudFlareChange {
	ttl := defaultCloudFlareRecordTTL
	if proxied && (cloudFlareTypeNotSupported[ep.RecordType] || strings.Contains(ep.DNSName, "*")) {
		proxied = false
	}
	if ep.RecordTTL.IsConfigured() {
		ttl = int(ep.RecordTTL)
	}

	switch ep.RecordType {
	case endpoint.RecordTypeSRV:
		// To maintain multiple SRV records for the same name we encode all
		// the target information in the name so they appear unique.
		in := srv.ParseRecord(ep.DNSName, ep.Targets[0])

		return &cloudFlareChange{
			Action: action,
			ResourceRecordSet: cloudflare.DNSRecord{
				Name: ep.DNSName,
				TTL:  ttl,
				Type: ep.RecordType,
				Data: &cloudFlareSRVData{
					Name:     in.Name.Name,
					Service:  in.Name.Service,
					Proto:    in.Name.Proto,
					Priority: in.Target.Priority,
					Weight:   in.Target.Weight,
					Port:     in.Target.Port,
					Target:   in.Target.Target,
				},
			},
		}
	}

	return &cloudFlareChange{
		Action: action,
		ResourceRecordSet: cloudflare.DNSRecord{
			Name:    ep.DNSName,
			TTL:     ttl,
			Proxied: proxied,
			Type:    ep.RecordType,
			Content: ep.Targets[0],
		},
	}
}
