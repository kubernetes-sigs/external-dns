/*
Copyright 2021 The Kubernetes Authors.

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

package safedns

import (
	"context"
	"fmt"
	"os"

	ansClient "github.com/ans-group/sdk-go/pkg/client"
	ansConnection "github.com/ans-group/sdk-go/pkg/connection"
	"github.com/ans-group/sdk-go/pkg/service/safedns"
	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
	"sigs.k8s.io/external-dns/provider"
)

// SafeDNS is an interface that is a subset of the SafeDNS service API that are actually used.
// Signatures must match exactly.
type SafeDNS interface {
	CreateZoneRecord(zoneName string, req safedns.CreateRecordRequest) (int, error)
	DeleteZoneRecord(zoneName string, recordID int) error
	GetZone(zoneName string) (safedns.Zone, error)
	GetZoneRecord(zoneName string, recordID int) (safedns.Record, error)
	GetZoneRecords(zoneName string, parameters ansConnection.APIRequestParameters) ([]safedns.Record, error)
	GetZones(parameters ansConnection.APIRequestParameters) ([]safedns.Zone, error)
	PatchZoneRecord(zoneName string, recordID int, patch safedns.PatchRecordRequest) (int, error)
	UpdateZoneRecord(zoneName string, record safedns.Record) (int, error)
}

// SafeDNSProvider implements the DNS provider spec for UKFast SafeDNS.
type SafeDNSProvider struct {
	provider.BaseProvider
	Client SafeDNS
	// Only consider hosted zones managing domains ending in this suffix
	domainFilter     endpoint.DomainFilter
	DryRun           bool
	APIRequestParams ansConnection.APIRequestParameters
}

// ZoneRecord is a datatype to simplify management of a record in a zone.
type ZoneRecord struct {
	ID      int
	Name    string
	Type    safedns.RecordType
	TTL     safedns.RecordTTL
	Zone    string
	Content string
}

func NewSafeDNSProvider(domainFilter endpoint.DomainFilter, dryRun bool) (*SafeDNSProvider, error) {
	token, ok := os.LookupEnv("SAFEDNS_TOKEN")
	if !ok {
		return nil, fmt.Errorf("no SAFEDNS_TOKEN found in environment")
	}

	ukfAPIConnection := ansConnection.NewAPIKeyCredentialsAPIConnection(token)
	ansClient := ansClient.NewClient(ukfAPIConnection)
	safeDNS := ansClient.SafeDNSService()

	provider := &SafeDNSProvider{
		Client:           safeDNS,
		domainFilter:     domainFilter,
		DryRun:           dryRun,
		APIRequestParams: *ansConnection.NewAPIRequestParameters(),
	}
	return provider, nil
}

// Zones returns the list of hosted zones in the SafeDNS account
func (p *SafeDNSProvider) Zones(ctx context.Context) ([]safedns.Zone, error) {
	var zones []safedns.Zone

	allZones, err := p.Client.GetZones(p.APIRequestParams)
	if err != nil {
		return nil, err
	}

	// Check each found zone to see whether they match the domain filter provided. If they do, append it to the array of
	// zones defined above. If not, continue to the next item in the loop.
	for _, zone := range allZones {
		if p.domainFilter.Match(zone.Name) {
			zones = append(zones, zone)
		} else {
			continue
		}
	}
	return zones, nil
}

func (p *SafeDNSProvider) ZoneRecords(ctx context.Context) ([]ZoneRecord, error) {
	zones, err := p.Zones(ctx)
	if err != nil {
		return nil, err
	}

	var zoneRecords []ZoneRecord
	for _, zone := range zones {
		// For each zone in the zonelist, get all records of an ExternalDNS supported type.
		records, err := p.Client.GetZoneRecords(zone.Name, p.APIRequestParams)
		if err != nil {
			return nil, err
		}
		for _, r := range records {
			zoneRecord := ZoneRecord{
				ID:      r.ID,
				Name:    r.Name,
				Type:    r.Type,
				TTL:     r.TTL,
				Zone:    zone.Name,
				Content: r.Content,
			}
			zoneRecords = append(zoneRecords, zoneRecord)
		}
	}
	return zoneRecords, nil
}

// Records returns a list of Endpoint resources created from all records in supported zones.
func (p *SafeDNSProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	var endpoints []*endpoint.Endpoint
	zoneRecords, err := p.ZoneRecords(ctx)
	if err != nil {
		return nil, err
	}
	for _, r := range zoneRecords {
		if provider.SupportedRecordType(string(r.Type)) {
			endpoints = append(endpoints, endpoint.NewEndpointWithTTL(r.Name, string(r.Type), endpoint.TTL(r.TTL), r.Content))
		}
	}
	return endpoints, nil
}

// ApplyChanges applies a given set of changes in a given zone.
func (p *SafeDNSProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	// Identify the zone name for each record
	zoneNameIDMapper := provider.ZoneIDName{}

	zones, err := p.Zones(ctx)
	if err != nil {
		return err
	}
	for _, zone := range zones {
		zoneNameIDMapper.Add(zone.Name, zone.Name)
	}

	zoneRecords, err := p.ZoneRecords(ctx)
	if err != nil {
		return err
	}

	for _, endpoint := range changes.Create {
		_, ZoneName := zoneNameIDMapper.FindZone(endpoint.DNSName)
		for _, target := range endpoint.Targets {
			request := safedns.CreateRecordRequest{
				Name:    endpoint.DNSName,
				Type:    endpoint.RecordType,
				Content: target,
			}
			log.WithFields(log.Fields{
				"zoneID":     ZoneName,
				"dnsName":    endpoint.DNSName,
				"recordType": endpoint.RecordType,
				"Value":      target,
			}).Info("Creating record")
			_, err := p.Client.CreateZoneRecord(ZoneName, request)
			if err != nil {
				return err
			}
		}
	}
	for _, endpoint := range changes.UpdateNew {
		// Currently iterates over each zoneRecord in ZoneRecords for each Endpoint
		// in UpdateNew; the same will go for Delete. As it's double-iteration,
		// that's O(n^2), which isn't great. No performance issues have been noted
		// thus far.
		var zoneRecord ZoneRecord
		for _, target := range endpoint.Targets {
			for _, zr := range zoneRecords {
				if zr.Name == endpoint.DNSName && zr.Content == target {
					zoneRecord = zr
					break
				}
			}

			newTTL := safedns.RecordTTL(int(endpoint.RecordTTL))
			newRecord := safedns.PatchRecordRequest{
				Name:    endpoint.DNSName,
				Content: target,
				TTL:     &newTTL,
				Type:    endpoint.RecordType,
			}
			log.WithFields(log.Fields{
				"zoneID":     zoneRecord.Zone,
				"dnsName":    newRecord.Name,
				"recordType": newRecord.Type,
				"Value":      newRecord.Content,
				"Priority":   newRecord.Priority,
			}).Info("Patching record")
			_, err = p.Client.PatchZoneRecord(zoneRecord.Zone, zoneRecord.ID, newRecord)
			if err != nil {
				return err
			}
		}
	}
	for _, endpoint := range changes.Delete {
		// As above, currently iterates in O(n^2). May be a good start for optimisations.
		var zoneRecord ZoneRecord
		for _, zr := range zoneRecords {
			if zr.Name == endpoint.DNSName && string(zr.Type) == endpoint.RecordType {
				zoneRecord = zr
				break
			}
		}
		log.WithFields(log.Fields{
			"zoneID":     zoneRecord.Zone,
			"dnsName":    zoneRecord.Name,
			"recordType": zoneRecord.Type,
		}).Info("Deleting record")
		err := p.Client.DeleteZoneRecord(zoneRecord.Zone, zoneRecord.ID)
		if err != nil {
			return err
		}
	}
	return nil
}
