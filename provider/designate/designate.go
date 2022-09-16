/*
Copyright 2022 The Kubernetes Authors.

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

package designate

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gophercloud/gophercloud/openstack/dns/v2/recordsets"
	"github.com/gophercloud/gophercloud/openstack/dns/v2/zones"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// designate provider type
type designateProvider struct {
	provider.BaseProvider
	client designateClientInterface

	// only consider hosted zones managing domains ending in this suffix
	domainFilter endpoint.DomainFilter
	dryRun       bool

	// cache Timeout
	cacheTimeout time.Duration
	cacheRefresh time.Time

	// cache zone answers
	zoneMu    sync.Mutex
	zoneCache provider.ZoneIDName

	// cache recordsets
	rsMu    sync.Mutex
	rsCache map[string]*recordsets.RecordSet
}

var (
	_ provider.Provider = &designateProvider{}
)

// NewDesignateProvider is a factory function for OpenStack designate providers
func NewDesignateProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*designateProvider, error) {
	client, err := newDesignateClient()
	if err != nil {
		return nil, err
	}
	return &designateProvider{
		client:       client,
		domainFilter: domainFilter,
		dryRun:       dryRun,
		cacheTimeout: 5 * time.Second,
	}, nil
}

// converts domain names to FQDN
func canonicalizeDomainNames(domains []string) []string {
	var cDomains []string
	for _, d := range domains {
		cDomains = append(cDomains, provider.EnsureTrailingDot(d))
	}
	return cDomains
}

// returns ZoneID -> ZoneName mapping for zones that are managed by the Designate and match domain filter
func (p *designateProvider) getZones() (provider.ZoneIDName, error) {
	if p.zoneCache != nil && time.Since(p.cacheRefresh) < p.cacheTimeout {
		log.Debug("Returning cached zones")
		p.zoneMu.Lock()
		defer p.zoneMu.Unlock()
		return p.zoneCache, nil
	}
	log.Debug("Calculating zones")
	result := provider.ZoneIDName{}

	err := p.client.ForEachZone(
		func(zone *zones.Zone) error {
			if zone.Type != "" && strings.ToUpper(zone.Type) != "PRIMARY" || zone.Status != "ACTIVE" {
				return nil
			}

			zoneName := provider.EnsureTrailingDot(zone.Name)
			if !p.domainFilter.Match(zoneName) {
				return nil
			}
			result[zone.ID] = zoneName
			return nil
		},
	)

	p.zoneMu.Lock()
	p.zoneCache = result
	p.cacheRefresh = time.Now()
	p.zoneMu.Unlock()

	return result, err
}

// Records returns the list of records.
func (p *designateProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var result []*endpoint.Endpoint

	managedZones, err := p.getZones()
	if err != nil {
		return nil, err
	}

	recordsByKey, err := p.getRecordSets(ctx, managedZones)
	if err != nil {
		return nil, err
	}
	for _, rs := range recordsByKey {
		ep := endpoint.NewEndpoint(rs.Name, rs.Type, rs.Records...)
		result = append(result, ep)
	}
	return result, err
}

func (p *designateProvider) getRecordSets(ctx context.Context, zones provider.ZoneIDName) (map[string]*recordsets.RecordSet, error) {
	if p.rsCache != nil && time.Since(p.cacheRefresh) < p.cacheTimeout {
		log.Debug("Returning cached recordSets")
		p.rsMu.Lock()
		defer p.rsMu.Unlock()
		return p.rsCache, nil
	}

	log.Debug("Calculating recordSets")
	recordSetsByZone := make(map[string]*recordsets.RecordSet)
	for zoneID := range zones {
		err := p.client.ForEachRecordSet(zoneID,
			func(rSet *recordsets.RecordSet) error {
				rs := *rSet
				if rs.Type != endpoint.RecordTypeA && rs.Type != endpoint.RecordTypeTXT && rs.Type != endpoint.RecordTypeCNAME {
					log.WithFields(log.Fields{
						"dnsName": rs.Name,
						"type":    rs.Type,
						"id":      rs.ID,
						"zone":    rs.ZoneID,
					}).Debug("Skipping")
					return nil
				}
				key := fmt.Sprintf("%s/%s", rs.Name, rs.Type)
				if dup, ok := recordSetsByZone[key]; ok && dup != nil {
					log.WithFields(log.Fields{
						"key":             key,
						"dnsName":         rs.Name,
						"type":            rs.Type,
						"id":              rs.ID,
						"zone":            rs.ZoneID,
						"duplicateID":     dup.ID,
						"duplicateZoneID": dup.ZoneID,
					}).Warn("Detected duplicate")
				}
				recordSetsByZone[key] = &rs
				return nil
			},
		)
		if err != nil {
			return nil, err
		}
	}

	p.rsMu.Lock()
	p.rsCache = recordSetsByZone
	p.cacheRefresh = time.Now()
	p.rsMu.Unlock()

	return recordSetsByZone, nil
}

// temporary structure to hold recordset parameters so that we could aggregate endpoints into recordsets
type recordSet struct {
	dnsName     string
	recordType  string
	zoneID      string
	recordSetID string
	targets     []string
}

// adds endpoint into recordset aggregation, loading original values from endpoint labels first
func addEndpoint(ep *endpoint.Endpoint, existingRecordSets map[string]*recordsets.RecordSet, recordSets map[string]*recordSet, delete bool) {
	key := fmt.Sprintf("%s/%s", provider.EnsureTrailingDot(ep.DNSName), ep.RecordType)
	rs := recordSets[key]

	if rs == nil {
		rs = &recordSet{
			dnsName:    provider.EnsureTrailingDot(ep.DNSName),
			recordType: ep.RecordType,
		}
	}

	if existingRs := existingRecordSets[key]; existingRs != nil {
		if rs.zoneID == "" {
			rs.zoneID = existingRs.ZoneID
		}
		if rs.recordSetID == "" {
			rs.recordSetID = existingRs.ID
		}
	}

	if !delete {
		targets := ep.Targets
		if ep.RecordType == endpoint.RecordTypeCNAME {
			targets = canonicalizeDomainNames(targets)
		}
		rs.targets = targets
	}

	recordSets[key] = rs
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *designateProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	managedZones, err := p.getZones()
	if err != nil {
		return err
	}
	existingRecordSets, err := p.getRecordSets(ctx, managedZones)
	if err != nil {
		return err
	}
	recordSets := map[string]*recordSet{}
	for _, ep := range changes.Create {
		addEndpoint(ep, existingRecordSets, recordSets, false)
	}
	for _, ep := range changes.UpdateNew {
		addEndpoint(ep, existingRecordSets, recordSets, false)
	}
	for _, ep := range changes.Delete {
		addEndpoint(ep, existingRecordSets, recordSets, true)
	}
	for _, rs := range recordSets {
		if upsertErr := p.upsertRecordSet(rs, managedZones); upsertErr != nil {
			log.WithFields(log.Fields{
				"dnsName":    rs.dnsName,
				"recordType": rs.recordType,
				"content":    strings.Join(rs.targets, ","),
				"zoneID":     rs.zoneID,
			}).Error(upsertErr)
			if err == nil {
				// We only capture the first error and return it
				err = upsertErr
			}
		}
	}
	return err
}

// apply recordset changes by inserting/updating/deleting recordsets
func (p *designateProvider) upsertRecordSet(rs *recordSet, managedZones provider.ZoneIDName) error {
	if rs.zoneID == "" {
		rs.zoneID, _ = managedZones.FindZone(provider.EnsureTrailingDot(rs.dnsName))
		if rs.zoneID == "" {
			log.WithFields(log.Fields{
				"dnsName": rs.dnsName,
			}).Warn("Skipping record because no hosted zone matching record DNS Name was detected")
			return nil
		}
		log.WithFields(log.Fields{
			"dnsName": rs.dnsName,
			"zoneID":  rs.zoneID,
		}).Debug("Fetched zoneID by FindZone")
	}
	if rs.recordSetID == "" && rs.targets == nil {
		return nil
	}
	if rs.recordSetID == "" {
		opts := recordsets.CreateOpts{
			Name:    rs.dnsName,
			Type:    rs.recordType,
			Records: rs.targets,
		}
		log.WithFields(log.Fields{
			"dnsName":    rs.dnsName,
			"recordType": rs.recordType,
			"content":    strings.Join(rs.targets, ","),
			"zoneID":     rs.zoneID,
		}).Info("Creating records")
		if p.dryRun {
			return nil
		}
		_, err := p.client.CreateRecordSet(rs.zoneID, opts)
		return err
	} else if len(rs.targets) == 0 {
		log.WithFields(log.Fields{
			"dnsName":    rs.dnsName,
			"recordType": rs.recordType,
			"zoneID":     rs.zoneID,
			"recordID":   rs.recordSetID,
		}).Info("Deleting records")
		if p.dryRun {
			return nil
		}
		return p.client.DeleteRecordSet(rs.zoneID, rs.recordSetID)
	} else {
		ttl := 0
		opts := recordsets.UpdateOpts{
			Records: rs.targets,
			TTL:     &ttl,
		}
		log.WithFields(log.Fields{
			"dnsName":    rs.dnsName,
			"recordType": rs.recordType,
			"zoneID":     rs.zoneID,
			"recordID":   rs.recordSetID,
			"content":    strings.Join(rs.targets, ","),
		}).Infof("Updating records")
		if p.dryRun {
			return nil
		}
		return p.client.UpdateRecordSet(rs.zoneID, rs.recordSetID, opts)
	}
}
