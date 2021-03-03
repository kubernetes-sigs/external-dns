/*
Copyright 2020 The Kubernetes Authors.

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

package vultr

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/vultr/govultr"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	vultrCreate = "CREATE"
	vultrDelete = "DELETE"
	vultrUpdate = "UPDATE"
	vultrTTL    = 3600
)

// VultrProvider is an implementation of Provider for Vultr DNS.
type VultrProvider struct {
	provider.BaseProvider
	client govultr.Client

	domainFilter endpoint.DomainFilter
	DryRun       bool
}

// VultrChanges differentiates between ChangActions.
type VultrChanges struct {
	Action string

	ResourceRecordSet govultr.DNSRecord
}

// NewVultrProvider initializes a new Vultr BNS based provider
func NewVultrProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*VultrProvider, error) {
	apiKey, ok := os.LookupEnv("VULTR_API_KEY")
	if !ok {
		return nil, fmt.Errorf("no token found")
	}

	client := govultr.NewClient(nil, apiKey)
	client.SetUserAgent(fmt.Sprintf("ExternalDNS/%s", client.UserAgent))

	provider := &VultrProvider{
		client:       *client,
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}

	return provider, nil
}

// Zones returns list of hosted zones
func (p *VultrProvider) Zones(ctx context.Context) ([]govultr.DNSDomain, error) {
	zones, err := p.fetchZones(ctx)
	if err != nil {
		return nil, err
	}

	return zones, nil
}

// Records returns the list of records.
func (p *VultrProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	var endpoints []*endpoint.Endpoint

	for _, zone := range zones {
		records, err := p.fetchRecords(ctx, zone.Domain)
		if err != nil {
			return nil, err
		}

		for _, r := range records {
			if provider.SupportedRecordType(r.Type) {
				name := fmt.Sprintf("%s.%s", r.Name, zone.Domain)

				// root name is identified by the empty string and should be
				// translated to zone name for the endpoint entry.
				if r.Name == "" {
					name = zone.Domain
				}

				endPointTTL := endpoint.NewEndpointWithTTL(name, r.Type, endpoint.TTL(r.TTL), r.Data)
				endpoints = append(endpoints, endPointTTL)
			}
		}
	}
	return endpoints, nil
}

func (p *VultrProvider) fetchRecords(ctx context.Context, domain string) ([]govultr.DNSRecord, error) {
	records, err := p.client.DNSRecord.List(ctx, domain)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (p *VultrProvider) fetchZones(ctx context.Context) ([]govultr.DNSDomain, error) {
	var zones []govultr.DNSDomain

	allZones, err := p.client.DNSDomain.List(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range allZones {
		if p.domainFilter.Match(zone.Domain) {
			zones = append(zones, zone)
		}
	}

	return zones, nil
}

func (p *VultrProvider) submitChanges(ctx context.Context, changes []*VultrChanges) error {
	if len(changes) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}

	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}

	zoneChanges := seperateChangesByZone(zones, changes)

	for zoneName, changes := range zoneChanges {
		for _, change := range changes {
			log.WithFields(log.Fields{
				"record":   change.ResourceRecordSet.Name,
				"type":     change.ResourceRecordSet.Type,
				"ttl":      change.ResourceRecordSet.TTL,
				"priority": change.ResourceRecordSet.Priority,
				"action":   change.Action,
				"zone":     zoneName,
			}).Info("Changing record.")

			switch change.Action {
			case vultrCreate:
				priority := getPriority(change.ResourceRecordSet.Priority)
				err = p.client.DNSRecord.Create(ctx, zoneName, change.ResourceRecordSet.Type, change.ResourceRecordSet.Name, change.ResourceRecordSet.Data, change.ResourceRecordSet.TTL, priority)
				if err != nil {
					return err
				}
			case vultrDelete:
				id, err := p.getRecordID(ctx, zoneName, change.ResourceRecordSet)
				if err != nil {
					return err
				}

				err = p.client.DNSRecord.Delete(ctx, zoneName, strconv.Itoa(id))
				if err != nil {
					return err
				}
			case vultrUpdate:
				id, err := p.getRecordID(ctx, zoneName, change.ResourceRecordSet)
				if err != nil {
					return err
				}

				record := &govultr.DNSRecord{
					RecordID: id,
					Type:     change.ResourceRecordSet.Type,
					Name:     change.ResourceRecordSet.Name,
					Data:     change.ResourceRecordSet.Data,
					TTL:      change.ResourceRecordSet.TTL,
				}

				err = p.client.DNSRecord.Update(ctx, zoneName, record)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *VultrProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	combinedChanges := make([]*VultrChanges, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newVultrChanges(vultrCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newVultrChanges(vultrUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newVultrChanges(vultrDelete, changes.Delete)...)

	return p.submitChanges(ctx, combinedChanges)
}

func newVultrChanges(action string, endpoints []*endpoint.Endpoint) []*VultrChanges {
	changes := make([]*VultrChanges, 0, len(endpoints))
	ttl := vultrTTL
	for _, e := range endpoints {
		if e.RecordTTL.IsConfigured() {
			ttl = int(e.RecordTTL)
		}

		change := &VultrChanges{
			Action: action,
			ResourceRecordSet: govultr.DNSRecord{
				Type: e.RecordType,
				Name: e.DNSName,
				Data: e.Targets[0],
				TTL:  ttl,
			},
		}
		changes = append(changes, change)
	}
	return changes
}

func seperateChangesByZone(zones []govultr.DNSDomain, changes []*VultrChanges) map[string][]*VultrChanges {
	change := make(map[string][]*VultrChanges)
	zoneNameID := provider.ZoneIDName{}

	for _, z := range zones {
		zoneNameID.Add(z.Domain, z.Domain)
		change[z.Domain] = []*VultrChanges{}
	}

	for _, c := range changes {
		zone, _ := zoneNameID.FindZone(c.ResourceRecordSet.Name)
		if zone == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", c.ResourceRecordSet.Name)
			continue
		}
		change[zone] = append(change[zone], c)
	}
	return change
}

func (p *VultrProvider) getRecordID(ctx context.Context, zone string, record govultr.DNSRecord) (recordID int, err error) {
	records, err := p.client.DNSRecord.List(ctx, zone)
	if err != nil {
		return 0, err
	}

	for _, r := range records {
		strippedName := strings.TrimSuffix(record.Name, "."+zone)
		if record.Name == zone {
			strippedName = ""
		}

		if r.Name == strippedName && r.Type == record.Type {
			return r.RecordID, nil
		}
	}

	return 0, fmt.Errorf("no record was found")
}

func getPriority(priority *int) int {
	p := 0
	if priority != nil {
		p = *priority
	}
	return p
}
