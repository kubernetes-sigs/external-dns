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
	"context"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	cloudflare "github.com/cloudflare/cloudflare-go"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/source"
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

// cloudFlareDNS is the subset of the CloudFlare API that we actually use.  Add methods as required. Signatures must match exactly.
type cloudFlareDNS interface {
	UserDetails() (cloudflare.User, error)
	ZoneIDByName(zoneName string) (string, error)
	ListZones(zoneID ...string) ([]cloudflare.Zone, error)
	ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error)
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

func (z zoneService) ListZonesContext(ctx context.Context, opts ...cloudflare.ReqOption) (cloudflare.ZonesResponse, error) {
	return z.service.ListZonesContext(ctx, opts...)
}

// CloudFlareProvider is an implementation of Provider for CloudFlare DNS.
type CloudFlareProvider struct {
	Client cloudFlareDNS
	// only consider hosted zones managing domains ending in this suffix
	domainFilter      endpoint.DomainFilter
	zoneIDFilter      ZoneIDFilter
	proxiedByDefault  bool
	DryRun            bool
	PaginationOptions cloudflare.PaginationOptions
}

// cloudFlareChange differentiates between ChangActions
type cloudFlareChange struct {
	Action            string
	ResourceRecordSet []cloudflare.DNSRecord
}

// NewCloudFlareProvider initializes a new CloudFlare DNS based Provider.
func NewCloudFlareProvider(domainFilter endpoint.DomainFilter, zoneIDFilter ZoneIDFilter, zonesPerPage int, proxiedByDefault bool, dryRun bool) (*CloudFlareProvider, error) {
	// initialize via chosen auth method and returns new API object
	var (
		config *cloudflare.API
		err    error
	)
	if os.Getenv("CF_API_TOKEN") != "" {
		config, err = cloudflare.NewWithAPIToken(os.Getenv("CF_API_TOKEN"))
	} else {
		config, err = cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to initialize cloudflare provider: %v", err)
	}
	provider := &CloudFlareProvider{
		//Client: config,
		Client:           zoneService{config},
		domainFilter:     domainFilter,
		zoneIDFilter:     zoneIDFilter,
		proxiedByDefault: proxiedByDefault,
		DryRun:           dryRun,
		PaginationOptions: cloudflare.PaginationOptions{
			PerPage: zonesPerPage,
			Page:    1,
		},
	}
	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *CloudFlareProvider) Zones(ctx context.Context) ([]cloudflare.Zone, error) {
	result := []cloudflare.Zone{}
	p.PaginationOptions.Page = 1

	for {
		zonesResponse, err := p.Client.ListZonesContext(ctx, cloudflare.WithPagination(p.PaginationOptions))
		if err != nil {
			return nil, err
		}

		for _, zone := range zonesResponse.Result {
			if !p.domainFilter.Match(zone.Name) {
				continue
			}

			if !p.zoneIDFilter.Match(zone.ID) {
				continue
			}
			result = append(result, zone)
		}
		if p.PaginationOptions.Page == zonesResponse.ResultInfo.TotalPages {
			break
		}
		p.PaginationOptions.Page++
	}
	return result, nil
}

// Records returns the list of records.
func (p *CloudFlareProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
	for _, zone := range zones {
		records, err := p.Client.DNSRecords(zone.ID, cloudflare.DNSRecord{})
		if err != nil {
			return nil, err
		}

		// As CloudFlare does not support "sets" of targets, but instead returns
		// a single entry for each name/type/target, we have to group by name
		// and record to allow the planner to calculate the correct plan. See #992.
		endpoints = append(endpoints, groupByNameAndType(records)...)
	}

	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *CloudFlareProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	proxiedByDefault := p.proxiedByDefault

	combinedChanges := make([]*cloudFlareChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareCreate, changes.Create, proxiedByDefault)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareUpdate, changes.UpdateNew, proxiedByDefault)...)
	combinedChanges = append(combinedChanges, newCloudFlareChanges(cloudFlareDelete, changes.Delete, proxiedByDefault)...)

	return p.submitChanges(ctx, combinedChanges)
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *CloudFlareProvider) submitChanges(ctx context.Context, changes []*cloudFlareChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	zones, err := p.Zones(ctx)
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
				"record":  change.ResourceRecordSet[0].Name,
				"type":    change.ResourceRecordSet[0].Type,
				"ttl":     change.ResourceRecordSet[0].TTL,
				"targets": len(change.ResourceRecordSet),
				"action":  change.Action,
				"zone":    zoneID,
			}

			log.WithFields(logFields).Info("Changing record.")

			if p.DryRun {
				continue
			}

			recordIDs := p.getRecordIDs(records, change.ResourceRecordSet[0])

			// to simplify bookkeeping for multiple records, an update is executed as delete+create
			if change.Action == cloudFlareDelete || change.Action == cloudFlareUpdate {
				for _, recordID := range recordIDs {
					err := p.Client.DeleteDNSRecord(zoneID, recordID)
					if err != nil {
						log.WithFields(logFields).Errorf("failed to delete record: %v", err)
					}
				}
			}

			if change.Action == cloudFlareCreate || change.Action == cloudFlareUpdate {
				for _, record := range change.ResourceRecordSet {
					_, err := p.Client.CreateDNSRecord(zoneID, record)
					if err != nil {
						log.WithFields(logFields).Errorf("failed to create record: %v", err)
					}
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
		zoneID, _ := zoneNameIDMapper.FindZone(c.ResourceRecordSet[0].Name)
		if zoneID == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", c.ResourceRecordSet[0].Name)
			continue
		}
		changes[zoneID] = append(changes[zoneID], c)
	}

	return changes
}

func (p *CloudFlareProvider) getRecordIDs(records []cloudflare.DNSRecord, record cloudflare.DNSRecord) []string {
	recordIDs := make([]string, 0)
	for _, zoneRecord := range records {
		if zoneRecord.Name == record.Name && zoneRecord.Type == record.Type {
			recordIDs = append(recordIDs, zoneRecord.ID)
		}
	}
	sort.Strings(recordIDs)
	return recordIDs
}

// newCloudFlareChanges returns a collection of Changes based on the given records and action.
func newCloudFlareChanges(action string, endpoints []*endpoint.Endpoint, proxiedByDefault bool) []*cloudFlareChange {
	changes := make([]*cloudFlareChange, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newCloudFlareChange(action, endpoint, proxiedByDefault))
	}

	return changes
}

func newCloudFlareChange(action string, endpoint *endpoint.Endpoint, proxiedByDefault bool) *cloudFlareChange {
	ttl := defaultCloudFlareRecordTTL
	proxied := shouldBeProxied(endpoint, proxiedByDefault)

	if endpoint.RecordTTL.IsConfigured() {
		ttl = int(endpoint.RecordTTL)
	}

	resourceRecordSet := make([]cloudflare.DNSRecord, len(endpoint.Targets))

	for i := range endpoint.Targets {
		resourceRecordSet[i] = cloudflare.DNSRecord{
			Name:    endpoint.DNSName,
			TTL:     ttl,
			Proxied: proxied,
			Type:    endpoint.RecordType,
			Content: endpoint.Targets[i],
		}
	}

	return &cloudFlareChange{
		Action:            action,
		ResourceRecordSet: resourceRecordSet,
	}
}

func shouldBeProxied(endpoint *endpoint.Endpoint, proxiedByDefault bool) bool {
	proxied := proxiedByDefault

	for _, v := range endpoint.ProviderSpecific {
		if v.Name == source.CloudflareProxiedKey {
			b, err := strconv.ParseBool(v.Value)
			if err != nil {
				log.Errorf("Failed to parse annotation [%s]: %v", source.CloudflareProxiedKey, err)
			} else {
				proxied = b
			}
			break
		}
	}

	if cloudFlareTypeNotSupported[endpoint.RecordType] || strings.Contains(endpoint.DNSName, "*") {
		proxied = false
	}
	return proxied
}

func groupByNameAndType(records []cloudflare.DNSRecord) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}

	// group supported records by name and type
	groups := map[string][]cloudflare.DNSRecord{}

	for _, r := range records {
		if !supportedRecordType(r.Type) {
			continue
		}

		groupBy := r.Name + r.Type
		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []cloudflare.DNSRecord{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	// create single endpoint with all the targets for each name/type
	for _, records := range groups {
		targets := make([]string, len(records))
		for i, record := range records {
			targets[i] = record.Content
		}
		endpoints = append(endpoints,
			endpoint.NewEndpointWithTTL(
				records[0].Name,
				records[0].Type,
				endpoint.TTL(records[0].TTL),
				targets...).
				WithProviderSpecific(source.CloudflareProxiedKey, strconv.FormatBool(records[0].Proxied)))
	}

	return endpoints
}
