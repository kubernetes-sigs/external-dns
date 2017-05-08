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
	"errors"

	"strings"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

var (
	// ErrZoneAlreadyExists error returned when zone cannot be created when it already exists
	ErrZoneAlreadyExists = errors.New("specified zone already exists")
	// ErrZoneNotFound error returned when specified zone does not exists
	ErrZoneNotFound = errors.New("specified zone not found")
	// ErrRecordAlreadyExists when create request is sent but record already exists
	ErrRecordAlreadyExists = errors.New("record already exists")
	// ErrRecordNotFound when update/delete request is sent but record not found
	ErrRecordNotFound = errors.New("record not found")
	// ErrInvalidBatchRequest when record is repeated in create/update/delete
	ErrInvalidBatchRequest = errors.New("invalid batch request")
)

// InMemoryProvider - dns provider only used for testing purposes
// initialized as dns provider with no records
type InMemoryProvider struct {
	zones          map[string]zone
	domain         string
	client         *inMemoryClient
	filter         *filter
	OnApplyChanges func(changes *plan.Changes)
	OnRecords      func()
}

// NewInMemoryProvider returns InMemoryProvider DNS provider interface implementation
func NewInMemoryProvider() *InMemoryProvider {
	return &InMemoryProvider{
		zones:          map[string]zone{},
		filter:         &filter{},
		OnApplyChanges: func(changes *plan.Changes) {},
		OnRecords:      func() {},
		domain:         "",
		client:         newInMemoryClient(),
	}
}

// CreateZone adds new zone if not present
func (im *InMemoryProvider) CreateZone(newZone string) error {
	return im.client.CreateZone(newZone)
}

func (im *InMemoryProvider) Zones() map[string]string {
	return im.filter.Zones(im.client.Zones())
}

// Records returns the list of endpoints
func (im *InMemoryProvider) Records(_ string) ([]*endpoint.Endpoint, error) {
	defer im.OnRecords()

	endpoints := make([]*endpoint.Endpoint, 0)

	for zoneID := range im.Zones() {
		records, err := im.client.Records(zoneID)
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			endpoints = append(endpoints, endpoint.NewEndpoint(record.Name, record.Target, record.Type))
		}
	}

	return endpoints, nil
}

// ApplyChanges simply modifies records in memory
// error checking occurs before any modifications are made, i.e. batch processing
// create record - record should not exist
// update/delete record - record should exist
// create/update/delete lists should not have overlapping records
func (im *InMemoryProvider) ApplyChanges(_ string, changes *plan.Changes) error {
	defer im.OnApplyChanges(changes)

	perZoneChanges := map[string]*plan.Changes{}

	zones := im.Zones()
	for zoneID := range zones {
		perZoneChanges[zoneID] = &plan.Changes{}
	}

	for _, ep := range changes.Create {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		perZoneChanges[zoneID].Create = append(perZoneChanges[zoneID].Create, ep)
	}
	for _, ep := range changes.UpdateNew {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		perZoneChanges[zoneID].UpdateNew = append(perZoneChanges[zoneID].UpdateNew, ep)
	}
	for _, ep := range changes.UpdateOld {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		perZoneChanges[zoneID].UpdateOld = append(perZoneChanges[zoneID].UpdateOld, ep)
	}
	for _, ep := range changes.Delete {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		perZoneChanges[zoneID].Delete = append(perZoneChanges[zoneID].Delete, ep)
	}

	for zoneID := range perZoneChanges {
		change := &inMemoryChange{
			Create:    im.convertToInMemoryRecord(perZoneChanges[zoneID].Create),
			UpdateNew: im.convertToInMemoryRecord(perZoneChanges[zoneID].UpdateNew),
			UpdateOld: im.convertToInMemoryRecord(perZoneChanges[zoneID].UpdateOld),
			Delete:    im.convertToInMemoryRecord(perZoneChanges[zoneID].Delete),
		}
		err := im.client.ApplyChanges(zoneID, change)
		if err != nil {
			return err
		}
	}

	return nil
}

func (im *InMemoryProvider) convertToInMemoryRecord(endpoints []*endpoint.Endpoint) []*inMemoryRecord {
	records := []*inMemoryRecord{}
	for _, ep := range endpoints {
		records = append(records, &inMemoryRecord{
			Type:   suitableType(ep),
			Name:   ep.DNSName,
			Target: ep.Target,
		})
	}
	return nil
}

// validateChangeBatch validates that the changes passed to InMemory DNS provider is valid
func (im *InMemoryProvider) validateChangeBatch(zone string, changes *plan.Changes) error {
	existing, ok := im.zones[zone]
	if !ok {
		return ErrZoneNotFound
	}
	mesh := map[string]map[string]bool{}
	for _, newEndpoint := range changes.Create {
		if im.findByType(suitableType(newEndpoint), existing[newEndpoint.DNSName]) != nil {
			return ErrRecordAlreadyExists
		}
		if _, exists := mesh[newEndpoint.DNSName]; exists {
			if mesh[newEndpoint.DNSName][suitableType(newEndpoint)] {
				return ErrInvalidBatchRequest
			}
			mesh[newEndpoint.DNSName][suitableType(newEndpoint)] = true
			continue
		}
		mesh[newEndpoint.DNSName] = map[string]bool{suitableType(newEndpoint): true}
	}
	for _, updateEndpoint := range changes.UpdateNew {
		if im.findByType(suitableType(updateEndpoint), existing[updateEndpoint.DNSName]) == nil {
			return ErrRecordNotFound
		}
		if _, exists := mesh[updateEndpoint.DNSName]; exists {
			if mesh[updateEndpoint.DNSName][suitableType(updateEndpoint)] {
				return ErrInvalidBatchRequest
			}
			mesh[updateEndpoint.DNSName][suitableType(updateEndpoint)] = true
			continue
		}
		mesh[updateEndpoint.DNSName] = map[string]bool{suitableType(updateEndpoint): true}
	}
	for _, updateOldEndpoint := range changes.UpdateOld {
		if rec := im.findByType(suitableType(updateOldEndpoint), existing[updateOldEndpoint.DNSName]); rec == nil || rec.Target != updateOldEndpoint.Target {
			return ErrRecordNotFound
		}
	}
	for _, deleteEndpoint := range changes.Delete {
		if rec := im.findByType(suitableType(deleteEndpoint), existing[deleteEndpoint.DNSName]); rec == nil || rec.Target != deleteEndpoint.Target {
			return ErrRecordNotFound
		}
		if _, exists := mesh[deleteEndpoint.DNSName]; exists {
			if mesh[deleteEndpoint.DNSName][suitableType(deleteEndpoint)] {
				return ErrInvalidBatchRequest
			}
			mesh[deleteEndpoint.DNSName][suitableType(deleteEndpoint)] = true
			continue
		}
		mesh[deleteEndpoint.DNSName] = map[string]bool{suitableType(deleteEndpoint): true}
	}
	return nil
}

func (im *InMemoryProvider) findByType(recordType string, records []*inMemoryRecord) *inMemoryRecord {
	for _, record := range records {
		if record.Type == recordType {
			return record
		}
	}
	return nil
}

// inMemoryRecord - record stored in memory
// Type - type of string (TODO: Type should probably be part of endpoint struct)
// Name - DNS name assigned to the record
// Target - target of the record
// Payload - string - additional information stored
type inMemoryRecord struct {
	Type    string
	Payload string
	Name    string
	Target  string
}

type zone map[string][]*inMemoryRecord

type inMemoryChange struct {
	Create    []*inMemoryRecord
	UpdateNew []*inMemoryRecord
	UpdateOld []*inMemoryRecord
	Delete    []*inMemoryRecord
}

type inMemoryClient struct {
	zones map[string]zone
}

func newInMemoryClient() *inMemoryClient {
	return &inMemoryClient{map[string]zone{}}
}

func (c *inMemoryClient) Records(zone string) ([]*inMemoryRecord, error) {
	if _, ok := c.zones[zone]; !ok {
		return nil, ErrZoneNotFound
	}

	records := []*inMemoryRecord{}
	for _, rec := range c.zones[zone] {
		records = append(records, rec...)
	}
	return records, nil
}

func (c *inMemoryClient) Zones() map[string]string {
	zones := map[string]string{}
	for zone := range c.zones {
		zones[zone] = zone
	}
	return zones
}

func (c *inMemoryClient) CreateZone(zone string) error {
	if _, ok := c.zones[zone]; ok {
		return ErrZoneAlreadyExists
	}
	c.zones[zone] = map[string][]*inMemoryRecord{}

	return nil
}

func (c *inMemoryClient) ApplyChanges(zoneID string, change *inMemoryChange) error {
	return nil
}

type filter struct {
	domain string
}

// Zones filters map[zoneID]zoneName for names having f.domain as suffix
func (f *filter) Zones(zones map[string]string) map[string]string {
	result := map[string]string{}
	for zoneID, zoneName := range zones {
		if strings.HasSuffix(zoneName, f.domain) {
			result[zoneID] = zoneName
		}
	}
	return result
}

// EndpointZoneID determines zoneID for endpoint from map[zoneID]zoneName by taking longest suffix zoneName match in endpoint DNSName
// returns empty string if no match found
func (f *filter) EndpointZoneID(endpoint *endpoint.Endpoint, zones map[string]string) (zoneID string) {
	var matchZoneID, matchZoneName string
	for zoneID, zoneName := range zones {
		if strings.HasSuffix(endpoint.DNSName, zoneName) && len(zoneName) > len(matchZoneName) {
			matchZoneName = zoneName
			matchZoneID = zoneID
		}
	}
	return matchZoneID
}
