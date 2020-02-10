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
	"strings"

	"github.com/digitalocean/godo"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
)

const (
	// DigitalOceanCreate is a ChangeAction enum value
	DigitalOceanCreate = "CREATE"
	// DigitalOceanDelete is a ChangeAction enum value
	DigitalOceanDelete = "DELETE"
	// DigitalOceanUpdate is a ChangeAction enum value
	DigitalOceanUpdate = "UPDATE"

	// digitalOceanRecordTTL is the default TTL value
	digitalOceanRecordTTL = 300
)

// DigitalOceanProvider is an implementation of Provider for Digital Ocean's DNS.
type DigitalOceanProvider struct {
	Client godo.DomainsService
	// only consider hosted zones managing domains ending in this suffix
	domainFilter endpoint.DomainFilter
	DryRun       bool
}

// DigitalOceanChange differentiates between ChangActions
type DigitalOceanChange struct {
	Action            string
	ResourceRecordSet godo.DomainRecord
}

// NewDigitalOceanProvider initializes a new DigitalOcean DNS based Provider.
func NewDigitalOceanProvider(ctx context.Context, domainFilter endpoint.DomainFilter, dryRun bool) (*DigitalOceanProvider, error) {
	token, ok := os.LookupEnv("DO_TOKEN")
	if !ok {
		return nil, fmt.Errorf("No token found")
	}
	oauthClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	}))
	client := godo.NewClient(oauthClient)

	provider := &DigitalOceanProvider{
		Client:       client.Domains,
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}
	return provider, nil
}

// Zones returns the list of hosted zones.
func (p *DigitalOceanProvider) Zones(ctx context.Context) ([]godo.Domain, error) {
	result := []godo.Domain{}

	zones, err := p.fetchZones(ctx)
	if err != nil {
		return nil, err
	}

	for _, zone := range zones {
		if p.domainFilter.Match(zone.Name) {
			result = append(result, zone)
		}
	}

	return result, nil
}

// Records returns the list of records in a given zone.
func (p *DigitalOceanProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}
	endpoints := []*endpoint.Endpoint{}
	for _, zone := range zones {
		records, err := p.fetchRecords(ctx, zone.Name)
		if err != nil {
			return nil, err
		}

		for _, r := range records {
			if supportedRecordType(r.Type) {
				name := r.Name + "." + zone.Name

				// root name is identified by @ and should be
				// translated to zone name for the endpoint entry.
				if r.Name == "@" {
					name = zone.Name
				}

				endpoints = append(endpoints, endpoint.NewEndpoint(name, r.Type, r.Data))
			}
		}
	}

	return endpoints, nil
}

func (p *DigitalOceanProvider) fetchRecords(ctx context.Context, zoneName string) ([]godo.DomainRecord, error) {
	allRecords := []godo.DomainRecord{}
	listOptions := &godo.ListOptions{}
	for {
		records, resp, err := p.Client.Records(ctx, zoneName, listOptions)
		if err != nil {
			return nil, err
		}
		allRecords = append(allRecords, records...)

		if resp == nil || resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		listOptions.Page = page + 1
	}

	return allRecords, nil
}

func (p *DigitalOceanProvider) fetchZones(ctx context.Context) ([]godo.Domain, error) {
	allZones := []godo.Domain{}
	listOptions := &godo.ListOptions{}
	for {
		zones, resp, err := p.Client.List(ctx, listOptions)
		if err != nil {
			return nil, err
		}
		allZones = append(allZones, zones...)

		if resp == nil || resp.Links == nil || resp.Links.IsLastPage() {
			break
		}

		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}

		listOptions.Page = page + 1
	}

	return allZones, nil
}

// submitChanges takes a zone and a collection of Changes and sends them as a single transaction.
func (p *DigitalOceanProvider) submitChanges(ctx context.Context, changes []*DigitalOceanChange) error {
	// return early if there is nothing to change
	if len(changes) == 0 {
		return nil
	}

	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}

	// separate into per-zone change sets to be passed to the API.
	changesByZone := digitalOceanChangesByZone(zones, changes)
	for zoneName, changes := range changesByZone {
		records, err := p.fetchRecords(ctx, zoneName)
		if err != nil {
			log.Errorf("Failed to list records in the zone: %s", zoneName)
			continue
		}
		for _, change := range changes {
			logFields := log.Fields{
				"record": change.ResourceRecordSet.Name,
				"type":   change.ResourceRecordSet.Type,
				"ttl":    change.ResourceRecordSet.TTL,
				"action": change.Action,
				"zone":   zoneName,
			}

			log.WithFields(logFields).Info("Changing record.")

			if p.DryRun {
				continue
			}

			change.ResourceRecordSet.Name = strings.TrimSuffix(change.ResourceRecordSet.Name, "."+zoneName)

			// record at the root should be defined as @ instead of
			// the full domain name
			if change.ResourceRecordSet.Name == zoneName {
				change.ResourceRecordSet.Name = "@"
			}

			// for some reason the DO API requires the '.' at the end of "data" in case of CNAME request
			// Example: {"type":"CNAME","name":"hello","data":"www.example.com."}
			if change.ResourceRecordSet.Type == endpoint.RecordTypeCNAME {
				change.ResourceRecordSet.Data += "."
			}

			switch change.Action {
			case DigitalOceanCreate:
				_, _, err = p.Client.CreateRecord(ctx, zoneName,
					&godo.DomainRecordEditRequest{
						Data: change.ResourceRecordSet.Data,
						Name: change.ResourceRecordSet.Name,
						Type: change.ResourceRecordSet.Type,
						TTL:  change.ResourceRecordSet.TTL,
					})
				if err != nil {
					return err
				}
			case DigitalOceanDelete:
				recordID := p.getRecordID(records, change.ResourceRecordSet)
				_, err = p.Client.DeleteRecord(ctx, zoneName, recordID)
				if err != nil {
					return err
				}
			case DigitalOceanUpdate:
				recordID := p.getRecordID(records, change.ResourceRecordSet)
				_, _, err = p.Client.EditRecord(ctx, zoneName, recordID,
					&godo.DomainRecordEditRequest{
						Data: change.ResourceRecordSet.Data,
						Name: change.ResourceRecordSet.Name,
						Type: change.ResourceRecordSet.Type,
						TTL:  change.ResourceRecordSet.TTL,
					})
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *DigitalOceanProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	combinedChanges := make([]*DigitalOceanChange, 0, len(changes.Create)+len(changes.UpdateNew)+len(changes.Delete))

	combinedChanges = append(combinedChanges, newDigitalOceanChanges(DigitalOceanCreate, changes.Create)...)
	combinedChanges = append(combinedChanges, newDigitalOceanChanges(DigitalOceanUpdate, changes.UpdateNew)...)
	combinedChanges = append(combinedChanges, newDigitalOceanChanges(DigitalOceanDelete, changes.Delete)...)

	return p.submitChanges(ctx, combinedChanges)
}

// newDigitalOceanChanges returns a collection of Changes based on the given records and action.
func newDigitalOceanChanges(action string, endpoints []*endpoint.Endpoint) []*DigitalOceanChange {
	changes := make([]*DigitalOceanChange, 0, len(endpoints))

	for _, endpoint := range endpoints {
		changes = append(changes, newDigitalOceanChange(action, endpoint))
	}

	return changes
}

func newDigitalOceanChange(action string, endpoint *endpoint.Endpoint) *DigitalOceanChange {
	// no annotation results in a TTL of 0, default to 300 for consistency with other providers
	var ttl = digitalOceanRecordTTL
	if endpoint.RecordTTL.IsConfigured() {
		ttl = int(endpoint.RecordTTL)
	}

	change := &DigitalOceanChange{
		Action: action,
		ResourceRecordSet: godo.DomainRecord{
			Name: endpoint.DNSName,
			Type: endpoint.RecordType,
			Data: endpoint.Targets[0],
			TTL:  ttl,
		},
	}
	return change
}

// getRecordID returns the ID from a record.
// the ID is mandatory to update and delete records
func (p *DigitalOceanProvider) getRecordID(records []godo.DomainRecord, record godo.DomainRecord) int {
	for _, zoneRecord := range records {
		if zoneRecord.Name == record.Name && zoneRecord.Type == record.Type {
			return zoneRecord.ID
		}
	}
	return 0
}

// digitalOceanchangesByZone separates a multi-zone change into a single change per zone.
func digitalOceanChangesByZone(zones []godo.Domain, changeSet []*DigitalOceanChange) map[string][]*DigitalOceanChange {
	changes := make(map[string][]*DigitalOceanChange)
	zoneNameIDMapper := zoneIDName{}
	for _, z := range zones {
		zoneNameIDMapper.Add(z.Name, z.Name)
		changes[z.Name] = []*DigitalOceanChange{}
	}

	for _, c := range changeSet {
		zone, _ := zoneNameIDMapper.FindZone(c.ResourceRecordSet.Name)
		if zone == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", c.ResourceRecordSet.Name)
			continue
		}
		changes[zone] = append(changes[zone], c)
	}

	return changes
}
