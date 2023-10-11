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

package glesys

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"sigs.k8s.io/external-dns/pkg/apis/externaldns"
	"sigs.k8s.io/external-dns/plan"
	"strings"

	"github.com/glesys/glesys-go/v7"
	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/provider"
)

type glesysChanges struct {
	Creates []*glesys.AddRecordParams
	Updates []*glesys.UpdateRecordParams
	Deletes []*glesys.DNSDomainRecord
}

const (
	// digitalOceanRecordTTL is the default TTL value
	glesysRecordTTL = 3600
)

type GlesysProvider struct {
	provider.BaseProvider
	Client glesys.Client
	// only consider hosted zones managing domains ending in this suffix
	domainFilter endpoint.DomainFilter
	// page size when querying paginated APIs
	apiPageSize int
	DryRun      bool
}

func NewGlesysProvider(ctx context.Context, dryRun bool) (*GlesysProvider, error) {
	project, ok := os.LookupEnv("GLESYS_PROJECT")
	if !ok {
		return nil, fmt.Errorf("GLESYS_PROJECT not found")
	}

	apikey, ok := os.LookupEnv("GLESYS_ACCESS_KEY")
	if !ok {
		return nil, fmt.Errorf("GLESYS_APIKEY not found")
	}

	var client = *glesys.NewClient(project, apikey, "ExternalDNS/"+externaldns.Version)
	p := &GlesysProvider{
		Client: client,
		DryRun: dryRun,
	}

	return p, nil
}

func (p *GlesysProvider) Zones(ctx context.Context) ([]glesys.DNSDomain, error) {
	result, err := p.Client.DNSDomains.List(ctx)
	if err != nil {
		return nil, err
	}
	return *result, nil
}

// Records returns the list of records in a given zone.
func (p *GlesysProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	zones, _ := p.Client.DNSDomains.List(ctx)
	for _, z := range *zones {
		var records, _ = p.Client.DNSDomains.ListRecords(ctx, z.Name)

		for _, r := range *records {
			ep := endpoint.NewEndpointWithTTL(r.Host, r.Type, endpoint.TTL(r.TTL), r.Data)
			endpoints = append(endpoints, ep)
		}
	}
	return endpoints, nil
}

func makeUpdateRecordParams(recordId int, domain string, host string, target string, recordType string, ttl int) glesys.UpdateRecordParams {
	adjustedName := strings.TrimSuffix(host, "."+domain)

	// Record at the root should be defined as @ instead of the full domain name.
	if adjustedName == domain {
		adjustedName = "@"
	}
	return glesys.UpdateRecordParams{
		RecordID: recordId,
		Data:     target,
		Host:     host,
		Type:     recordType,
		TTL:      ttl,
	}
}

func makeAddRecordParams(domain string, host string, target string, recordType string, ttl int) glesys.AddRecordParams {
	target = strings.Replace(target, "\"", "", -1)
	adjustedName := strings.TrimSuffix(host, "."+domain)

	// Record at the root should be defined as @ instead of the full domain name.
	if adjustedName == domain {
		adjustedName = "@"
	}
	return glesys.AddRecordParams{
		DomainName: domain,
		Data:       target,
		Host:       adjustedName,
		Type:       recordType,
		TTL:        ttl,
	}
}

func doesRecordExistInList(records []*glesys.DNSDomainRecord, record glesys.DNSDomainRecord) bool {

	for _, r := range records {
		if r.RecordID == record.RecordID {
			return true
		}
	}
	return false
}

func processDeleteActions(
	recordsByDomain map[string][]glesys.DNSDomainRecord,
	deletesByDomain map[string][]*endpoint.Endpoint,
	changes *glesysChanges) error {
	for domain, endpoints := range deletesByDomain {
		records := recordsByDomain[domain]
		for _, endpoint := range endpoints {

			matchingRecords := getMatchingDomainRecords(records, domain, endpoint)
			for _, record := range matchingRecords {
				//Skip if the record already exists in delete tasks
				if doesRecordExistInList(changes.Deletes, record) {
					continue
				}
				changes.Deletes = append(changes.Deletes, &record)
			}
		}
	}
	return nil
}
func getTTLFromEndpoint(ep *endpoint.Endpoint) int {
	if ep.RecordTTL.IsConfigured() {
		return int(ep.RecordTTL)
	}
	return glesysRecordTTL
}
func processUpdateActions(
	recordsByDomain map[string][]glesys.DNSDomainRecord,
	updatesByDomain map[string][]*endpoint.Endpoint,
	changes *glesysChanges) error {
	// Generate creates and updates based on existing
	for domain, updates := range updatesByDomain {
		if len(updates) == 0 {
			log.WithFields(log.Fields{
				"domain": domain,
			}).Debug("Skipping Zone, no updates found.")
			continue
		}

		records := recordsByDomain[domain]
		log.WithFields(log.Fields{
			"domain":  domain,
			"records": records,
		}).Debug("Records for domain")

		for _, ep := range updates {
			matchingRecords := getMatchingDomainRecords(records, domain, ep)

			log.WithFields(log.Fields{
				"endpoint":        ep,
				"matchingRecords": matchingRecords,
			}).Debug("matching records")

			if len(matchingRecords) == 0 {
				log.WithFields(log.Fields{
					"domain":     domain,
					"dnsName":    ep.DNSName,
					"recordType": ep.RecordType,
				}).Warn("Planning an update but no existing records found.")
			}

			matchingRecordsByTarget := map[string]glesys.DNSDomainRecord{}
			for _, r := range matchingRecords {
				matchingRecordsByTarget[r.Data] = r
			}

			ttl := getTTLFromEndpoint(ep)

			// Generate create and delete actions based on existence of a record for each target.
			for _, target := range ep.Targets {
				target = strings.Replace(target, "\"", "", -1)
				target := strings.TrimSuffix(target, ".")
				if record, ok := matchingRecordsByTarget[target]; ok {
					if target == record.Data {
						delete(matchingRecordsByTarget, target)
						continue
					}
					log.WithFields(log.Fields{
						"domain":     domain,
						"dnsName":    ep.DNSName,
						"recordType": ep.RecordType,
						"target":     target,
					}).Warn("Updating existing target")
					updateRecord := makeUpdateRecordParams(record.RecordID, domain, target, record.Host, ep.RecordType, record.TTL)

					changes.Updates = append(changes.Updates, &updateRecord)

					delete(matchingRecordsByTarget, target)
				} else {
					// Record did not previously exist, create new 'target'
					log.WithFields(log.Fields{
						"domain":     domain,
						"dnsName":    ep.DNSName,
						"recordType": ep.RecordType,
						"target":     target,
					}).Warn("Creating new target")
					addRecord := makeAddRecordParams(domain, ep.DNSName, target, ep.RecordType, ttl)
					changes.Creates = append(changes.Creates, &addRecord)
				}
			}

			// Any remaining records have been removed, delete them
			for _, record := range matchingRecordsByTarget {
				target := strings.Replace(record.Data, "\"", "", -1)
				log.WithFields(log.Fields{
					"domain":     domain,
					"dnsName":    ep.DNSName,
					"recordType": ep.RecordType,
					"target":     target,
				}).Warn("Deleting target")
				record.Data = target
				changes.Deletes = append(changes.Deletes, &glesys.DNSDomainRecord{
					DomainName: record.DomainName,
					Data:       record.Data,
					Host:       record.Host,
					RecordID:   record.RecordID,
					TTL:        record.TTL,
					Type:       record.Type,
				})
			}
		}
	}

	return nil
}

func endpointsByZone(zoneNameIDMapper provider.ZoneIDName, endpoints []*endpoint.Endpoint) map[string][]*endpoint.Endpoint {
	endpointsByZone := make(map[string][]*endpoint.Endpoint)

	for _, ep := range endpoints {
		zoneID, _ := zoneNameIDMapper.FindZone(ep.DNSName)
		if zoneID == "" {
			log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", ep.DNSName)
			continue
		}
		endpointsByZone[zoneID] = append(endpointsByZone[zoneID], ep)
	}

	return endpointsByZone
}
func (p *GlesysProvider) submitChanges(ctx context.Context, changes *glesysChanges) error {

	for _, create := range changes.Creates {
		log.WithFields(log.Fields{
			"Domain": create.DomainName,
			"Host":   create.Host,
			"Data":   create.Data,
			"TTL":    create.TTL,
		}).Info("Creating new record")
		if !p.DryRun {
			p.Client.DNSDomains.AddRecord(ctx, *create)
		}

	}
	for _, updates := range changes.Updates {
		if !p.DryRun {
			p.Client.DNSDomains.UpdateRecord(ctx, *updates)
		}
	}
	for _, deletes := range changes.Deletes {
		if !p.DryRun {
			p.Client.DNSDomains.DeleteRecord(ctx, deletes.RecordID)
		}
	}
	return nil
}

// Records returns the list of records in a given zone.
func (p *GlesysProvider) ApplyChanges(ctx context.Context, planChanges *plan.Changes) error {
	recordsByDomain, zoneNameIDMapper, err := p.getRecordsByDomain(ctx)
	if err != nil {
		return err
	}

	//TODO: Apparently they send deletes in the updatelist
	createsByDomain := endpointsByZone(zoneNameIDMapper, planChanges.Create)
	updatesByDomain := endpointsByZone(zoneNameIDMapper, planChanges.UpdateNew)
	deletesByDomain := endpointsByZone(zoneNameIDMapper, planChanges.Delete)

	var changes glesysChanges
	if err := processUpdateActions(recordsByDomain, createsByDomain, &changes); err != nil {
		return err
	}
	if err := processUpdateActions(recordsByDomain, updatesByDomain, &changes); err != nil {
		return err
	}
	if err := processDeleteActions(recordsByDomain, deletesByDomain, &changes); err != nil {
		return err
	}
	return p.submitChanges(ctx, &changes)

}

func getMatchingDomainRecords(records []glesys.DNSDomainRecord, domain string, ep *endpoint.Endpoint) []glesys.DNSDomainRecord {
	var name string
	if ep.DNSName != domain {
		name = strings.TrimSuffix(ep.DNSName, "."+domain)
	} else {
		name = "@"
	}

	var result []glesys.DNSDomainRecord
	for _, r := range records {
		if r.Host == name && r.Type == ep.RecordType {
			r.Data = strings.TrimSuffix(r.Data, ".")
			result = append(result, r)
		}
	}
	return result
}

func (p *GlesysProvider) ZoneRecords(ctx context.Context) ([]glesys.DNSDomainRecord, error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	var dNSDomainRecords []glesys.DNSDomainRecord
	for _, zone := range zones {
		// For each zone in the zonelist, get all records of an ExternalDNS supported type.
		records, err := p.Client.DNSDomains.ListRecords(ctx, zone.Name)
		if err != nil {
			return nil, err
		}
		for _, r := range *records {
			zoneRecord := glesys.DNSDomainRecord{
				RecordID:   r.RecordID,
				Host:       r.Host,
				Type:       r.Type,
				TTL:        r.TTL,
				DomainName: zone.Name,
				Data:       r.Data,
			}
			dNSDomainRecords = append(dNSDomainRecords, zoneRecord)
		}
	}
	return dNSDomainRecords, nil
}

func (p *GlesysProvider) getRecordsByDomain(ctx context.Context) (map[string][]glesys.DNSDomainRecord, provider.ZoneIDName, error) {
	recordsByDomain := map[string][]glesys.DNSDomainRecord{}

	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, nil, err
	}

	zonesByDomain := make(map[string]glesys.DNSDomain)
	zoneNameIDMapper := provider.ZoneIDName{}
	for _, z := range zones {
		zoneNameIDMapper.Add(z.Name, z.Name)
		zonesByDomain[z.Name] = z
	}

	// Fetch records for each zone
	for _, zone := range zones {

		records, err := p.fetchRecords(ctx, zone.Name)
		if err != nil {
			return nil, nil, err
		}

		recordsByDomain[zone.Name] = append(recordsByDomain[zone.Name], records...)
	}

	return recordsByDomain, zoneNameIDMapper, nil
}

func (p *GlesysProvider) fetchRecords(ctx context.Context, zoneName string) ([]glesys.DNSDomainRecord, error) {
	records, err := p.Client.DNSDomains.ListRecords(ctx, zoneName)
	return *records, err

}

func splitDnsName(Zonename string, DNSName string) string {
	return strings.Trim(strings.Replace(DNSName, Zonename, "", -1), ".")

}
