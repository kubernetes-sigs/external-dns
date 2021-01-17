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

package godaddy

import (
	"context"
	"errors"
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
	gdMinimalTTL = 600
	gdCreate     = iota
	gdUpdate
	gdDelete
)

var (
	// ErrRecordToMutateNotFound when ApplyChange has to update/delete and didn't found the record in the existing zone (Change with no record ID)
	ErrRecordToMutateNotFound = errors.New("record to mutate not found in current zone")
	// ErrNoDryRun No dry run support for the moment
	ErrNoDryRun = errors.New("dry run not supported")
)

type gdClient interface {
	Patch(string, interface{}, interface{}) error
	Post(string, interface{}, interface{}) error
	Put(string, interface{}, interface{}) error
	Get(string, interface{}) error
	Delete(string, interface{}) error
}

// GDProvider declare GoDaddy provider
type GDProvider struct {
	provider.BaseProvider

	domainFilter endpoint.DomainFilter
	client       gdClient
	ttl          int64
	DryRun       bool
}

type gdEndpoint struct {
	endpoint *endpoint.Endpoint
	action   int
}

type gdRecordField struct {
	Data     string  `json:"data"`
	Name     string  `json:"name"`
	TTL      int64   `json:"ttl"`
	Type     string  `json:"type"`
	Port     *int    `json:"port,omitempty"`
	Priority *int    `json:"priority,omitempty"`
	Weight   *int64  `json:"weight,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Service  *string `json:"service,omitempty"`
}

type gdRecords struct {
	records []gdRecordField
	changed bool
	zone    string
}

type gdZone struct {
	CreatedAt           string
	Domain              string
	DomainID            int64
	ExpirationProtected bool
	Expires             string
	ExposeWhois         bool
	HoldRegistrar       bool
	Locked              bool
	NameServers         *[]string
	Privacy             bool
	RenewAuto           bool
	RenewDeadline       string
	Renewable           bool
	Status              string
	TransferProtected   bool
}

type gdZoneIDName map[string]*gdRecords

func (z gdZoneIDName) add(zoneID string, zoneRecord *gdRecords) {
	z[zoneID] = zoneRecord
}

func (z gdZoneIDName) findZoneRecord(hostname string) (suitableZoneID string, suitableZoneRecord *gdRecords) {
	for zoneID, zoneRecord := range z {
		if hostname == zoneRecord.zone || strings.HasSuffix(hostname, "."+zoneRecord.zone) {
			if suitableZoneRecord == nil || len(zoneRecord.zone) > len(suitableZoneRecord.zone) {
				suitableZoneID = zoneID
				suitableZoneRecord = zoneRecord
			}
		}
	}

	return
}

// NewGoDaddyProvider initializes a new GoDaddy DNS based Provider.
func NewGoDaddyProvider(ctx context.Context, domainFilter endpoint.DomainFilter, ttl int64, apiKey, apiSecret string, useOTE, dryRun bool) (*GDProvider, error) {
	client, err := NewClient(useOTE, apiKey, apiSecret)

	if err != nil {
		return nil, err
	}

	// TODO: Add Dry Run support
	if dryRun {
		return nil, ErrNoDryRun
	}

	return &GDProvider{
		client:       client,
		domainFilter: domainFilter,
		ttl:          maxOf(gdMinimalTTL, ttl),
		DryRun:       dryRun,
	}, nil
}

func (p *GDProvider) zones() ([]string, error) {
	zones := []gdZone{}
	filteredZones := []string{}

	if err := p.client.Get("/v1/domains?statuses=ACTIVE", &zones); err != nil {
		return nil, err
	}

	for _, zone := range zones {
		if p.domainFilter.Match(zone.Domain) {
			filteredZones = append(filteredZones, zone.Domain)
			log.Debugf("GoDaddy: %s zone found", zone.Domain)
		}
	}

	log.Infof("GoDaddy: %d zones found", len(filteredZones))

	return filteredZones, nil
}

func (p *GDProvider) zonesRecords(ctx context.Context, all bool) ([]string, []gdRecords, error) {
	var allRecords []gdRecords
	zones, err := p.zones()

	if err != nil {
		return nil, nil, err
	}

	if len(zones) == 0 {
		allRecords = []gdRecords{}
	} else if len(zones) == 1 {
		record, err := p.records(&ctx, zones[0], all)

		if err != nil {
			return nil, nil, err
		}

		allRecords = append(allRecords, *record)
	} else {
		chRecords := make(chan gdRecords, len(zones))

		eg, ctx := errgroup.WithContext(ctx)

		for _, zoneName := range zones {
			zone := zoneName
			eg.Go(func() error {
				record, err := p.records(&ctx, zone, all)

				if err != nil {
					return err
				}

				chRecords <- *record

				return nil
			})
		}

		if err := eg.Wait(); err != nil {
			return nil, nil, err
		}

		close(chRecords)

		for records := range chRecords {
			allRecords = append(allRecords, records)
		}
	}

	return zones, allRecords, nil
}

func (p *GDProvider) records(ctx *context.Context, zone string, all bool) (*gdRecords, error) {
	var recordsIds []gdRecordField

	log.Debugf("GoDaddy: Getting records for %s", zone)

	if err := p.client.Get(fmt.Sprintf("/v1/domains/%s/records", zone), &recordsIds); err != nil {
		return nil, err
	}

	if all {
		return &gdRecords{
			zone:    zone,
			records: recordsIds,
		}, nil
	}

	results := &gdRecords{
		zone:    zone,
		records: make([]gdRecordField, 0, len(recordsIds)),
	}

	for _, rec := range recordsIds {
		if provider.SupportedRecordType(rec.Type) {
			log.Debugf("GoDaddy: Record %s for %s is %+v", rec.Name, zone, rec)

			results.records = append(results.records, rec)
		} else {
			log.Infof("GoDaddy: Discard record %s for %s is %+v", rec.Name, zone, rec)
		}
	}

	return results, nil
}

func (p *GDProvider) groupByNameAndType(zoneRecords []gdRecords) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}

	// group supported records by name and type
	groupsByZone := map[string]map[string][]gdRecordField{}

	for _, zone := range zoneRecords {
		groups := map[string][]gdRecordField{}

		groupsByZone[zone.zone] = groups

		for _, r := range zone.records {
			groupBy := fmt.Sprintf("%s - %s", r.Type, r.Name)

			if _, ok := groups[groupBy]; !ok {
				groups[groupBy] = []gdRecordField{}
			}

			groups[groupBy] = append(groups[groupBy], r)
		}
	}

	// create single endpoint with all the targets for each name/type
	for zoneName, groups := range groupsByZone {
		for _, records := range groups {
			targets := []string{}

			for _, record := range records {
				targets = append(targets, record.Data)
			}

			var recordName string

			if records[0].Name == "@" {
				recordName = strings.TrimPrefix(zoneName, ".")
			} else {
				recordName = strings.TrimPrefix(fmt.Sprintf("%s.%s", records[0].Name, zoneName), ".")
			}

			endpoint := endpoint.NewEndpointWithTTL(
				recordName,
				records[0].Type,
				endpoint.TTL(records[0].TTL),
				targets...,
			)

			endpoints = append(endpoints, endpoint)
		}
	}

	return endpoints
}

// Records returns the list of records in all relevant zones.
func (p *GDProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	_, records, err := p.zonesRecords(ctx, false)

	if err != nil {
		return nil, err
	}

	endpoints := p.groupByNameAndType(records)

	log.Infof("GoDaddy: %d endpoints have been found", len(endpoints))

	return endpoints, nil
}

func (p *GDProvider) flushRecords(patch bool, zoneRecord *gdRecords) error {
	if patch {
		return p.client.Patch(fmt.Sprintf("/v1/domains/%s/records", zoneRecord.zone), zoneRecord.records, nil)
	}

	return p.client.Put(fmt.Sprintf("/v1/domains/%s/records", zoneRecord.zone), zoneRecord.records, nil)
}

func (p *GDProvider) appendChange(action int, endpoints []*endpoint.Endpoint, allChanges []gdEndpoint) []gdEndpoint {
	for _, e := range endpoints {
		allChanges = append(allChanges, gdEndpoint{
			action:   action,
			endpoint: e,
		})
	}

	return allChanges
}

func (p *GDProvider) changeAllRecords(patch bool, endpoints []gdEndpoint, zoneRecords []*gdRecords) error {
	zoneNameIDMapper := gdZoneIDName{}

	for _, zoneRecord := range zoneRecords {
		if patch {
			zoneRecord.changed = false
			zoneRecord.records = nil
		}

		zoneNameIDMapper.add(zoneRecord.zone, zoneRecord)
	}

	for _, e := range endpoints {
		dnsName := e.endpoint.DNSName
		zone, zoneRecord := zoneNameIDMapper.findZoneRecord(dnsName)

		if zone == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", dnsName)
		} else {
			dnsName = strings.TrimSuffix(dnsName, "."+zone)

			if e.endpoint.RecordType == endpoint.RecordTypeA && (len(dnsName) == 0) {
				dnsName = "@"
			}

			for _, target := range e.endpoint.Targets {
				change := gdRecordField{
					Type: e.endpoint.RecordType,
					Name: dnsName,
					TTL:  p.ttl,
					Data: target,
				}

				if e.endpoint.RecordTTL.IsConfigured() {
					change.TTL = maxOf(gdMinimalTTL, int64(e.endpoint.RecordTTL))
				}

				zoneRecord.applyChange(e.action, change)
			}
		}
	}

	return nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *GDProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	if countTargets(changes) == 0 {
		return nil
	}

	_, records, err := p.zonesRecords(ctx, true)

	if err != nil {
		return err
	}

	changedZoneRecords := make([]*gdRecords, len(records))

	for i := range records {
		changedZoneRecords[i] = &records[i]
	}

	allChanges := make([]gdEndpoint, 0, countTargets(changes))

	allChanges = p.appendChange(gdCreate, changes.Create, allChanges)
	allChanges = p.appendChange(gdCreate, changes.UpdateNew, allChanges)
	allChanges = p.appendChange(gdDelete, changes.UpdateOld, allChanges)
	allChanges = p.appendChange(gdDelete, changes.Delete, allChanges)

	log.Infof("GoDaddy: %d changes will be done", len(allChanges))

	patch := len(changes.UpdateOld)+len(changes.Delete) == 0

	if err = p.changeAllRecords(patch, allChanges, changedZoneRecords); err != nil {
		return err
	}

	for _, record := range changedZoneRecords {
		if record.changed {
			if err = p.flushRecords(patch, record); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *gdRecords) addRecord(change gdRecordField) {
	log.Debugf("GoDaddy: Add an entry %s to zone %s", change.String(), p.zone)

	p.records = append(p.records, change)
	p.changed = true
}

func (p *gdRecords) updateRecord(change gdRecordField) {
	log.Debugf("GoDaddy: Update an entry %s to zone %s", change.String(), p.zone)

	for index, record := range p.records {
		if record.Type == change.Type && record.Name == change.Name {
			p.records[index] = change
			p.changed = true
			break
		}
	}
}

// Remove one record from the record list
func (p *gdRecords) deleteRecord(change gdRecordField) {
	log.Debugf("GoDaddy: Delete an entry %s to zone %s", change.String(), p.zone)

	deleteIndex := -1

	for index, record := range p.records {
		if record.Type == change.Type && record.Name == change.Name && record.Data == change.Data {
			deleteIndex = index
			break
		}
	}

	if deleteIndex >= 0 {
		p.records[deleteIndex] = p.records[len(p.records)-1]

		p.records = p.records[:len(p.records)-1]
		p.changed = true
	} else {
		log.Warnf("GoDaddy: record in zone %s not found %s to delete", p.zone, change.String())
	}
}

func (p *gdRecords) applyChange(action int, change gdRecordField) {
	switch action {
	case gdCreate:
		p.addRecord(change)
	case gdUpdate:
		p.updateRecord(change)
	case gdDelete:
		p.deleteRecord(change)
	}
}

func (c gdRecordField) String() string {
	return fmt.Sprintf("%s %d IN %s %s", c.Name, c.TTL, c.Type, c.Data)
}

func countTargets(p *plan.Changes) int {
	changes := [][]*endpoint.Endpoint{p.Create, p.UpdateNew, p.UpdateOld, p.Delete}
	count := 0

	for _, endpoints := range changes {
		for _, endpoint := range endpoints {
			count += len(endpoint.Targets)
		}
	}

	return count
}

func maxOf(vars ...int64) int64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
