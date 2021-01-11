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
	gdDefaultTTL = 600
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
	DryRun       bool
}

type gdRecordField struct {
	Data     string
	Name     string
	Port     int
	Priority int
	Protocol *string
	Service  *string
	TTL      int
	Type     string
	Weight   int
}

type gdRecord struct {
	gdRecordField
	zone *string
}

type gdChange struct {
	gdRecord
	Action int
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

// NewGoDaddyProvider initializes a new OVH DNS based Provider.
func NewGoDaddyProvider(ctx context.Context, domainFilter endpoint.DomainFilter, apiKey, apiSecret string, production, dryRun bool) (*GDProvider, error) {
	client, err := NewClient(production, apiKey, apiSecret)

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
		}
	}

	log.Infof("GoDaddy: %d zones found", len(filteredZones))

	return filteredZones, nil
}

func (p *GDProvider) zonesRecords(ctx context.Context) ([]string, []gdRecord, error) {
	var allRecords []gdRecord
	zones, err := p.zones()

	if err != nil {
		return nil, nil, err
	}

	chRecords := make(chan []gdRecord, len(zones))

	eg, ctx := errgroup.WithContext(ctx)

	for _, zone := range zones {
		zone := zone
		eg.Go(func() error {
			return p.records(&ctx, &zone, chRecords)
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, nil, err
	}

	close(chRecords)

	for records := range chRecords {
		allRecords = append(allRecords, records...)
	}

	return zones, allRecords, nil
}

func (p *GDProvider) records(ctx *context.Context, zone *string, records chan<- []gdRecord) error {
	var recordsIds []gdRecord

	log.Debugf("GoDaddy: Getting records for %s", *zone)

	if err := p.client.Get(fmt.Sprintf("/v1/domains/%s/records", *zone), &recordsIds); err != nil {
		return err
	}

	results := make([]gdRecord, 0, len(recordsIds))

	for _, rec := range recordsIds {
		if provider.SupportedRecordType(rec.Type) {
			log.Debugf("GoDaddy: Record %s for %s is %+v", rec.Name, *zone, rec)

			rec.zone = zone
			results = append(results, rec)
		}
	}

	records <- results

	return nil
}

func (p *GDProvider) groupByNameAndType(records []gdRecord) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}

	// group supported records by name and type
	groups := map[string][]gdRecord{}

	for _, r := range records {
		groupBy := fmt.Sprintf("%s - %s.%s", r.Type, r.Name, *r.zone)

		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []gdRecord{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	// create single endpoint with all the targets for each name/type
	for _, records := range groups {
		targets := []string{}

		for _, record := range records {
			targets = append(targets, record.Data)
		}

		var recordName string

		if records[0].Name == "@" {
			recordName = strings.TrimPrefix(*records[0].zone, ".")
		} else {
			recordName = strings.TrimPrefix(fmt.Sprintf("%s.%s", records[0].Name, *records[0].zone), ".")
		}

		endpoint := endpoint.NewEndpointWithTTL(
			recordName,
			records[0].Type,
			endpoint.TTL(records[0].TTL),
			targets...,
		)

		endpoints = append(endpoints, endpoint)
	}

	return endpoints
}

// Records returns the list of records in all relevant zones.
func (p *GDProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	_, records, err := p.zonesRecords(ctx)

	if err != nil {
		return nil, err
	}

	endpoints := p.groupByNameAndType(records)

	log.Infof("GoDaddy: %d endpoints have been found", len(endpoints))

	return endpoints, nil
}

func (p *GDProvider) change(change gdChange) error {
	switch change.Action {
	case gdCreate:
		log.Debugf("GoDaddy: Add an entry to %s", change.String())
		return p.client.Patch(fmt.Sprintf("/v1/domains/%s/records", *change.zone), []gdRecordField{change.gdRecord.gdRecordField}, nil)
	case gdUpdate:
		log.Debugf("GoDaddy: Update an entry to %s", change.String())
		return p.client.Put(fmt.Sprintf("/v1/domains/%s/records/%s/%s", *change.zone, change.Type, change.Name), []gdRecordField{change.gdRecord.gdRecordField}, nil)
	case gdDelete:
		log.Debugf("GoDaddy: Delete an entry to %s", change.String())
		return p.client.Delete(fmt.Sprintf("/v1/domains/%s/records/%s/%s", *change.zone, change.Type, change.Name), nil)
	}
	return nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *GDProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	zones, records, err := p.zonesRecords(ctx)
	zonesChangeUniques := map[string]bool{}
	if err != nil {
		return err
	}

	allChanges := make([]gdChange, 0, countTargets(changes.Create, changes.UpdateNew, changes.UpdateOld, changes.Delete))

	allChanges = append(allChanges, newGoDaddyChange(gdCreate, changes.Create, zones, records)...)
	allChanges = append(allChanges, newGoDaddyChange(gdCreate, changes.UpdateNew, zones, records)...)

	allChanges = append(allChanges, newGoDaddyChange(gdDelete, changes.UpdateOld, zones, records)...)
	allChanges = append(allChanges, newGoDaddyChange(gdDelete, changes.Delete, zones, records)...)

	log.Infof("GoDaddy: %d changes will be done", len(allChanges))

	eg, _ := errgroup.WithContext(ctx)

	for _, change := range allChanges {
		change := change
		zonesChangeUniques[*change.zone] = true

		eg.Go(func() error { return p.change(change) })
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func countTargets(allEndpoints ...[]*endpoint.Endpoint) int {
	count := 0
	for _, endpoints := range allEndpoints {
		for _, endpoint := range endpoints {
			count += len(endpoint.Targets)
		}
	}
	return count
}

func newGoDaddyChange(action int, endpoints []*endpoint.Endpoint, zones []string, records []gdRecord) []gdChange {
	gdChanges := make([]gdChange, 0, countTargets(endpoints))
	zoneNameIDMapper := provider.ZoneIDName{}

	for _, zone := range zones {
		zoneNameIDMapper.Add(zone, zone)
	}

	for _, e := range endpoints {
		zone, _ := zoneNameIDMapper.FindZone(e.DNSName)

		if zone == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", e.DNSName)
			continue
		}

		dnsName := strings.TrimSuffix(e.DNSName, "."+zone)

		if e.RecordType == endpoint.RecordTypeA && (len(dnsName) == 0) {
			dnsName = "@"
		}

		for _, target := range e.Targets {

			if e.RecordType == endpoint.RecordTypeCNAME {
				target = target + "."
			} else if e.RecordType == endpoint.RecordTypeA && (len(dnsName) == 0 || dnsName == ".") {
				dnsName = "@"
			}

			change := gdChange{
				Action: action,
				gdRecord: gdRecord{
					zone: &zone,
					gdRecordField: gdRecordField{
						Type: e.RecordType,
						Name: dnsName,
						TTL:  gdDefaultTTL,
						Data: target,
					},
				},
			}

			if e.RecordTTL.IsConfigured() {
				change.TTL = int(e.RecordTTL)
			}

			gdChanges = append(gdChanges, change)
		}
	}

	return gdChanges
}

func (c *gdChange) String() string {
	return fmt.Sprintf("%s zone : %s %d IN %s %s", *c.zone, c.Name, c.TTL, c.Type, c.Data)
}
