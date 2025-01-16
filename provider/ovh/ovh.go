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

package ovh

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/miekg/dns"
	"github.com/ovh/go-ovh/ovh"
	"github.com/patrickmn/go-cache"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"

	"go.uber.org/ratelimit"
)

const (
	ovhDefaultTTL = 0
	ovhCreate     = iota
	ovhDelete
)

var (
	// ErrRecordToMutateNotFound when ApplyChange has to update/delete and didn't found the record in the existing zone (Change with no record ID)
	ErrRecordToMutateNotFound = errors.New("record to mutate not found in current zone")
	// ErrNoDryRun No dry run support for the moment
	ErrNoDryRun = errors.New("dry run not supported")
)

// OVHProvider is an implementation of Provider for OVH DNS.
type OVHProvider struct {
	provider.BaseProvider

	client ovhClient

	apiRateLimiter ratelimit.Limiter

	domainFilter endpoint.DomainFilter
	DryRun       bool

	// UseCache controls if the OVHProvider will cache records in memory, and serve them
	// without recontacting the OVHcloud API if the SOA of the domain zone hasn't changed.
	// Note that, when disabling cache, OVHcloud API has rate-limiting that will hit if
	// your refresh rate/number of records is too big, which might cause issue with the
	// provider.
	// Default value: true
	UseCache bool

	cacheInstance *cache.Cache
	dnsClient     dnsClient
}

type ovhClient interface {
	Post(string, interface{}, interface{}) error
	Get(string, interface{}) error
	Delete(string, interface{}) error
}

type dnsClient interface {
	ExchangeContext(ctx context.Context, m *dns.Msg, a string) (*dns.Msg, time.Duration, error)
}

type ovhRecordFields struct {
	FieldType string `json:"fieldType"`
	SubDomain string `json:"subDomain"`
	TTL       int64  `json:"ttl"`
	Target    string `json:"target"`
}

type ovhRecord struct {
	ovhRecordFields
	ID   uint64 `json:"id"`
	Zone string `json:"zone"`
}

type ovhChange struct {
	ovhRecord
	Action int
}

// NewOVHProvider initializes a new OVH DNS based Provider.
func NewOVHProvider(ctx context.Context, domainFilter endpoint.DomainFilter, endpoint string, apiRateLimit int, dryRun bool) (*OVHProvider, error) {
	client, err := ovh.NewEndpointClient(endpoint)
	if err != nil {
		return nil, err
	}

	client.UserAgent = externaldns.Version

	// TODO: Add Dry Run support
	if dryRun {
		return nil, ErrNoDryRun
	}
	return &OVHProvider{
		client:         client,
		domainFilter:   domainFilter,
		apiRateLimiter: ratelimit.New(apiRateLimit),
		DryRun:         dryRun,
		cacheInstance:  cache.New(cache.NoExpiration, cache.NoExpiration),
		dnsClient:      new(dns.Client),
		UseCache:       true,
	}, nil
}

// Records returns the list of records in all relevant zones.
func (p *OVHProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	_, records, err := p.zonesRecords(ctx)
	if err != nil {
		return nil, err
	}
	endpoints := ovhGroupByNameAndType(records)
	log.Infof("OVH: %d endpoints have been found", len(endpoints))
	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *OVHProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) (err error) {
	zones, records, err := p.zonesRecords(ctx)
	if err != nil {
		return provider.NewSoftError(err)
	}

	zonesChangeUniques := map[string]bool{}

	// Always refresh zones even in case of errors.
	defer func() {
		log.Debugf("OVH: %d zones will be refreshed", len(zonesChangeUniques))

		eg, _ := errgroup.WithContext(ctx)
		for zone := range zonesChangeUniques {
			// This is necessary because the loop variable zone is reused in each iteration of the loop,
			// and without this line, the goroutines launched by eg.Go would all reference the same zone variable.
			zone := zone
			eg.Go(func() error { return p.refresh(zone) })
		}

		if e := eg.Wait(); e != nil && err == nil { // return the error only if there is no error during the changes
			err = provider.NewSoftError(e)
		}
	}()

	allChanges := make([]ovhChange, 0, countTargets(changes.Create, changes.UpdateNew, changes.UpdateOld, changes.Delete))
	allChanges = append(allChanges, newOvhChange(ovhCreate, changes.Create, zones, records)...)
	allChanges = append(allChanges, newOvhChange(ovhCreate, changes.UpdateNew, zones, records)...)
	allChanges = append(allChanges, newOvhChange(ovhDelete, changes.UpdateOld, zones, records)...)
	allChanges = append(allChanges, newOvhChange(ovhDelete, changes.Delete, zones, records)...)

	log.Infof("OVH: %d changes will be done", len(allChanges))

	eg, _ := errgroup.WithContext(ctx)
	for _, change := range allChanges {
		change := change
		zonesChangeUniques[change.Zone] = true
		eg.Go(func() error { return p.change(change) })
	}
	if err := eg.Wait(); err != nil {
		return provider.NewSoftError(err)
	}

	return nil
}

func (p *OVHProvider) refresh(zone string) error {
	log.Debugf("OVH: Refresh %s zone", zone)

	// Zone has been altered so we invalidate the cache
	// so that the next run will reload it.
	p.invalidateCache(zone)

	p.apiRateLimiter.Take()
	if err := p.client.Post(fmt.Sprintf("/domain/zone/%s/refresh", zone), nil, nil); err != nil {
		return provider.NewSoftError(err)
	}
	return nil
}

func (p *OVHProvider) change(change ovhChange) error {
	p.apiRateLimiter.Take()

	switch change.Action {
	case ovhCreate:
		log.Debugf("OVH: Add an entry to %s", change.String())
		return p.client.Post(fmt.Sprintf("/domain/zone/%s/record", change.Zone), change.ovhRecordFields, nil)
	case ovhDelete:
		if change.ID == 0 {
			return ErrRecordToMutateNotFound
		}
		log.Debugf("OVH: Delete an entry to %s", change.String())
		return p.client.Delete(fmt.Sprintf("/domain/zone/%s/record/%d", change.Zone, change.ID), nil)
	}
	return nil
}

func (p *OVHProvider) invalidateCache(zone string) {
	p.cacheInstance.Delete(zone + "#soa")
}

func (p *OVHProvider) zonesRecords(ctx context.Context) ([]string, []ovhRecord, error) {
	var allRecords []ovhRecord
	zones, err := p.zones()
	if err != nil {
		return nil, nil, provider.NewSoftError(err)
	}

	chRecords := make(chan []ovhRecord, len(zones))
	eg, ctx := errgroup.WithContext(ctx)
	for _, zone := range zones {
		zone := zone
		eg.Go(func() error { return p.records(&ctx, &zone, chRecords) })
	}
	if err := eg.Wait(); err != nil {
		return nil, nil, provider.NewSoftError(err)
	}
	close(chRecords)
	for records := range chRecords {
		allRecords = append(allRecords, records...)
	}
	return zones, allRecords, nil
}

func (p *OVHProvider) zones() ([]string, error) {
	zones := []string{}
	filteredZones := []string{}

	p.apiRateLimiter.Take()
	if err := p.client.Get("/domain/zone", &zones); err != nil {
		return nil, err
	}

	for _, zoneName := range zones {
		if p.domainFilter.Match(zoneName) {
			filteredZones = append(filteredZones, zoneName)
		}
	}
	log.Infof("OVH: %d zones found", len(filteredZones))
	return filteredZones, nil
}

type ovhSoa struct {
	Server  string `json:"server"`
	Serial  uint32 `json:"serial"`
	records []ovhRecord
}

func (p *OVHProvider) records(ctx *context.Context, zone *string, records chan<- []ovhRecord) error {
	var recordsIds []uint64
	ovhRecords := make([]ovhRecord, len(recordsIds))
	eg, _ := errgroup.WithContext(*ctx)

	if p.UseCache {
		if cachedSoaItf, ok := p.cacheInstance.Get(*zone + "#soa"); ok {
			cachedSoa := cachedSoaItf.(ovhSoa)

			m := new(dns.Msg)
			m.SetQuestion(dns.Fqdn(*zone), dns.TypeSOA)
			in, _, err := p.dnsClient.ExchangeContext(*ctx, m, strings.TrimSuffix(cachedSoa.Server, ".")+":53")
			if err == nil {
				if s, ok := in.Answer[0].(*dns.SOA); ok {
					// do something with t.Txt
					if s.Serial == cachedSoa.Serial {
						records <- cachedSoa.records
						return nil
					}
				}
			}

			p.invalidateCache(*zone)
		}
	}

	log.Debugf("OVH: Getting records for %s", *zone)

	p.apiRateLimiter.Take()
	var soa ovhSoa
	if p.UseCache {
		if err := p.client.Get("/domain/zone/"+*zone+"/soa", &soa); err != nil {
			return err
		}
	}

	if err := p.client.Get(fmt.Sprintf("/domain/zone/%s/record", *zone), &recordsIds); err != nil {
		return err
	}
	chRecords := make(chan ovhRecord, len(recordsIds))
	for _, id := range recordsIds {
		id := id
		eg.Go(func() error { return p.record(zone, id, chRecords) })
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	close(chRecords)
	for record := range chRecords {
		ovhRecords = append(ovhRecords, record)
	}

	if p.UseCache {
		soa.records = ovhRecords
		_ = p.cacheInstance.Add(*zone+"#soa", soa, time.Hour)
	}

	records <- ovhRecords
	return nil
}

func (p *OVHProvider) record(zone *string, id uint64, records chan<- ovhRecord) error {
	record := ovhRecord{}

	log.Debugf("OVH: Getting record %d for %s", id, *zone)

	p.apiRateLimiter.Take()
	if err := p.client.Get(fmt.Sprintf("/domain/zone/%s/record/%d", *zone, id), &record); err != nil {
		return err
	}
	if provider.SupportedRecordType(record.FieldType) {
		log.Debugf("OVH: Record %d for %s is %+v", id, *zone, record)
		records <- record
	}
	return nil
}

func ovhGroupByNameAndType(records []ovhRecord) []*endpoint.Endpoint {
	endpoints := []*endpoint.Endpoint{}

	// group supported records by name and type
	groups := map[string][]ovhRecord{}

	for _, r := range records {
		groupBy := r.Zone + r.SubDomain + r.FieldType
		if _, ok := groups[groupBy]; !ok {
			groups[groupBy] = []ovhRecord{}
		}

		groups[groupBy] = append(groups[groupBy], r)
	}

	// create single endpoint with all the targets for each name/type
	for _, records := range groups {
		targets := []string{}
		for _, record := range records {
			targets = append(targets, record.Target)
		}
		endpoint := endpoint.NewEndpointWithTTL(
			strings.TrimPrefix(records[0].SubDomain+"."+records[0].Zone, "."),
			records[0].FieldType,
			endpoint.TTL(records[0].TTL),
			targets...,
		)
		endpoints = append(endpoints, endpoint)
	}

	return endpoints
}

func newOvhChange(action int, endpoints []*endpoint.Endpoint, zones []string, records []ovhRecord) []ovhChange {
	// Copy the records because we need to mutate the list.
	newRecords := make([]ovhRecord, len(records))
	copy(newRecords, records)

	zoneNameIDMapper := provider.ZoneIDName{}
	ovhChanges := make([]ovhChange, 0, countTargets(endpoints))
	for _, zone := range zones {
		zoneNameIDMapper.Add(zone, zone)
	}

	for _, e := range endpoints {
		zone, _ := zoneNameIDMapper.FindZone(e.DNSName)
		if zone == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", e.DNSName)
			continue
		}
		for _, target := range e.Targets {
			if e.RecordType == endpoint.RecordTypeCNAME {
				target = target + "."
			}
			change := ovhChange{
				Action: action,
				ovhRecord: ovhRecord{
					Zone: zone,
					ovhRecordFields: ovhRecordFields{
						FieldType: e.RecordType,
						SubDomain: strings.TrimSuffix(e.DNSName, "."+zone),
						TTL:       ovhDefaultTTL,
						Target:    target,
					},
				},
			}
			if e.RecordTTL.IsConfigured() {
				change.TTL = int64(e.RecordTTL)
			}

			// The Zone might have multiple records with the same target. In order to avoid applying the action to the
			// same OVH record, we remove a record from the list when a match is found.
			for i := 0; i < len(newRecords); i++ {
				rec := newRecords[i]
				if rec.Zone == change.Zone && rec.SubDomain == change.SubDomain && rec.FieldType == change.FieldType && rec.Target == change.Target {
					change.ID = rec.ID
					// Deleting this record from the list to avoid retargetting it later if a change with a similar target exists.
					newRecords = append(newRecords[:i], newRecords[i+1:]...)
					break
				}
			}
			ovhChanges = append(ovhChanges, change)
		}
	}

	return ovhChanges
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

func (c *ovhChange) String() string {
	if c.ID != 0 {
		return fmt.Sprintf("%s zone (ID : %d) : %s %d IN %s %s", c.Zone, c.ID, c.SubDomain, c.TTL, c.FieldType, c.Target)
	}
	return fmt.Sprintf("%s zone : %s %d IN %s %s", c.Zone, c.SubDomain, c.TTL, c.FieldType, c.Target)
}
