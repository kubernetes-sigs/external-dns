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

package hetzner

import (
	"context"
	"errors"
	"os"
	"strings"

	hclouddns "git.blindage.org/21h/hcloud-dns"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

const (
       hetznerTTL    = 600
)

type RecordKey struct {
	DNSName    string
	RecordType string
}

type RecordMap map[RecordKey][]*hclouddns.HCloudRecord

func (rmap RecordMap) Add(record hclouddns.HCloudRecord, zone hclouddns.HCloudZone) {
	name := record.Name + "." + zone.Name
	if record.Name == "@" {
		name = zone.Name
	}
	key := RecordKey{name, string(record.RecordType)}
	rmap[key] = append(rmap[key], &record)
}

func (rmap RecordMap) Lookup(endpoint *endpoint.Endpoint) []*hclouddns.HCloudRecord {
	return rmap[RecordKey{endpoint.DNSName, endpoint.RecordType}]
}


type HetznerZoneMapper struct {
	provider.ZoneIDName
}

func (p *HetznerProvider) NewHetznerZoneMapper() (*HetznerZoneMapper, error) {
	zones, err := p.Client.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		return nil, err
	}

	hzm := &HetznerZoneMapper{provider.ZoneIDName{}}
	for _, z := range zones.Zones {
		hzm.Add(z.ID, z.Name)
	}
	return hzm, nil
}

func (hzm *HetznerZoneMapper) endpointToHCloudRecords(endpoint *endpoint.Endpoint) ([]*hclouddns.HCloudRecord, error) {
	zoneID, zoneName := hzm.FindZone(endpoint.DNSName)
	if zoneName == "" {
		log.Debugf("Skipping record %s because no hosted zone matching record DNS Name was detected", endpoint.DNSName)
		return nil, nil
	}
	name := strings.TrimSuffix(endpoint.DNSName, "." + zoneName)
	if name == zoneName {
		name = "@"
	}
	common := hclouddns.HCloudRecord{
		Name:       name,
		RecordType: hclouddns.RecordType(endpoint.RecordType),
		ZoneID:     zoneID,
		TTL:        int(endpoint.RecordTTL),
	}
	if common.TTL == 0 {
		common.TTL = hetznerTTL
	}
	result := make([]*hclouddns.HCloudRecord, 0, len(endpoint.Targets))
	for _, target := range endpoint.Targets {
		record := common
		record.Value = target
		if endpoint.RecordType == "CNAME" && !strings.HasSuffix(record.Value, ".") {
			record.Value += "."
		}
		result = append(result, &record)
	}
	return result, nil
}

type HetznerProvider struct {
	provider.BaseProvider
	Client       hclouddns.HCloudClientAdapter
	domainFilter endpoint.DomainFilter
	DryRun       bool
}

func NewHetznerProvider(ctx context.Context, domainFilter endpoint.DomainFilter, dryRun bool) (*HetznerProvider, error) {
	token, ok := os.LookupEnv("HETZNER_TOKEN")
	if !ok {
		return nil, errors.New("no environment variable HETZNER_TOKEN provided")
	}

	client := hclouddns.New(token)

	provider := &HetznerProvider{
		Client:       client,
		domainFilter: domainFilter,
		DryRun:       dryRun,
	}
	return provider, nil
}

func (p *HetznerProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	records, err := p.fetchRecords()
	if err != nil {
		return nil, err
	}

	endpoints := []*endpoint.Endpoint{}
	for key, rrset := range records {
		ttl := 0
		targets := make(endpoint.Targets, 0, len(rrset))
		for _, rr := range rrset {
			targets = append(targets, rr.Value)
			if ttl == 0 || rr.TTL < ttl {
				ttl = rr.TTL
			}
		}
		endpoints = append(endpoints, endpoint.NewEndpointWithTTL(key.DNSName, key.RecordType, endpoint.TTL(ttl), targets...))
	}

	return endpoints, nil
}

// ApplyChanges as to adapt the Plan with 1:N endpoint, target relation to Hetzner's 1:1 record, value relation
// There is some overlap with what Plan should have done. If other providers work in a similar way, consider
// Refactoring this to Plan
func (p *HetznerProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	allRecords, err := p.fetchRecords()
	if err != nil {
		return err
	}

	mapper, err := p.NewHetznerZoneMapper()
	if err != nil {
		return err
	}

	createRecords := []*hclouddns.HCloudRecord{}
	deleteRecords := []*hclouddns.HCloudRecord{}
	updateRecords := []*hclouddns.HCloudRecord{}

	// Create one record per target of each endpoint in Create changes
	for _, endpoint := range changes.Create {
		records, _ := mapper.endpointToHCloudRecords(endpoint)
		if records != nil {
			createRecords = append(createRecords, records...)
		}
	}

	// Delete all records with matching name, type for Delete changes
	for _, endpoint := range changes.Delete {
		records := allRecords.Lookup(endpoint)
		if records != nil {
			deleteRecords = append(deleteRecords, records...)
		}
	}

	// For Update changes, check existing records with matching name, type:
	//   if targets match, update the TTL
	//   change any existing records to match desired but missing target, TTL
	//   delete any other existing records
	//   create new records for still missing targets
	for _, endpoint := range changes.UpdateNew {
		oldRecords := allRecords.Lookup(endpoint)
		newRecords, _ := mapper.endpointToHCloudRecords(endpoint)
		if newRecords != nil {
			reserve := []*hclouddns.HCloudRecord{}
			for _, oldRecord := range oldRecords {
				found := false
				for _, newRecord := range newRecords {
					if oldRecord.Value == newRecord.Value {
						newRecord.ID = oldRecord.ID
						found = true
						if oldRecord.TTL != newRecord.TTL {
							updateRecords = append(updateRecords, newRecord)
						}
						break
					}
				}
				if !found {
					reserve = append(reserve, oldRecord)
				}
			}
			update := 0
			for _, newRecord := range newRecords {
				if newRecord.ID == "" {
					if len(reserve) > update {
						reserve[update].Value = newRecord.Value
						reserve[update].TTL = newRecord.TTL
						update++
					} else {
						createRecords = append(createRecords, newRecord)
					}
				}
			}
			if update > 0 {
				updateRecords = append(updateRecords, reserve[0:update]...)
			}
			if len(reserve) > update {
				deleteRecords = append(deleteRecords, reserve[update:]...)
			}
		}
	}

	if len(createRecords) == 0 && len(deleteRecords) == 0 && len(updateRecords) == 0 {
		log.Infof("All records are already up to date")
		return nil
	}

	// Apply the changes collected above
	for _, record := range createRecords {
		p.logRecord(record, "CREATE")
		if p.DryRun {
			continue
		}
		answer, err := p.Client.CreateRecord(*record)
		if err != nil {
			log.WithFields(log.Fields{
				"Code":         answer.Error.Code,
				"Message":      answer.Error.Message,
				"Record name":  answer.Record.Name,
				"Record type":  answer.Record.RecordType,
				"Record value": answer.Record.Value,
			}).Warning("Create problem")
			return err
		}
	}

	for _, record := range deleteRecords {
		p.logRecord(record, "DELETE")
		if p.DryRun {
			continue
		}
		answer, err := p.Client.DeleteRecord(record.ID)
		if err != nil {
			log.WithFields(log.Fields{
				"Code":         answer.Error.Code,
				"Message":      answer.Error.Message,
			}).Warning("Delete problem")
			return err
		}
	}

	for _, record := range updateRecords {
		p.logRecord(record, "UPDATE")
		if p.DryRun {
			continue
		}
		answer, err := p.Client.UpdateRecord(*record)
		if err != nil {
			log.WithFields(log.Fields{
				"Code":         answer.Error.Code,
				"Message":      answer.Error.Message,
				"Record name":  answer.Record.Name,
				"Record type":  answer.Record.RecordType,
				"Record value": answer.Record.Value,
			}).Warning("Update problem")
			return err
		}
	}

	return nil
}

func (p *HetznerProvider) fetchRecords() (RecordMap, error) {
	zones, err := p.Client.GetZones(hclouddns.HCloudGetZonesParams{})
	if err != nil {
		return nil, err
	}
	result := RecordMap{}
	for _, zone := range zones.Zones {
		records, err := p.Client.GetRecords(hclouddns.HCloudGetRecordsParams{ZoneID: zone.ID})
		if err != nil {
			return nil, err
		}

		for _, r := range records.Records {
			if provider.SupportedRecordType(string(r.RecordType)) {
				result.Add(r, zone)
			}
		}
	}

	return result, nil
}

func (p *HetznerProvider) logRecord(record *hclouddns.HCloudRecord, message string) {
	if p.DryRun {
		message += " (dry run)"
	}
	log.WithFields(log.Fields{
		"id":      record.ID,
		"record":  record.Name,
		"type":    record.RecordType,
		"value":   record.Value,
		"ttl":     record.TTL,
	}).Info(message)
}
