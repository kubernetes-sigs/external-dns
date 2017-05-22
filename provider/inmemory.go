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
	domain         string
	client         *inMemoryClient
	filter         *filter
	OnApplyChanges func(changes *plan.Changes)
	OnRecords      func()
}

// NewInMemoryProvider returns InMemoryProvider DNS provider interface implementation
func NewInMemoryProvider() *InMemoryProvider {
	return &InMemoryProvider{
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

// Zones returns filtered zones as specified by domain
func (im *InMemoryProvider) Zones() map[string]string {
	return im.filter.Zones(im.client.Zones())
}

// Records returns the list of endpoints
func (im *InMemoryProvider) Records() ([]*endpoint.Endpoint, error) {
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
func (im *InMemoryProvider) ApplyChanges(changes *plan.Changes) error {
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
			Create:    convertToInMemoryRecord(perZoneChanges[zoneID].Create),
			UpdateNew: convertToInMemoryRecord(perZoneChanges[zoneID].UpdateNew),
			UpdateOld: convertToInMemoryRecord(perZoneChanges[zoneID].UpdateOld),
			Delete:    convertToInMemoryRecord(perZoneChanges[zoneID].Delete),
		}
		err := im.client.ApplyChanges(zoneID, change)
		if err != nil {
			return err
		}
	}

	return nil
}

func convertToInMemoryRecord(endpoints []*endpoint.Endpoint) []*inMemoryRecord {
	records := []*inMemoryRecord{}
	for _, ep := range endpoints {
		records = append(records, &inMemoryRecord{
			Type:   suitableType(ep),
			Name:   ep.DNSName,
			Target: ep.Target,
		})
	}
	return records
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

// inMemoryRecord - record stored in memory
// Type - type of string
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

func (c *inMemoryClient) ApplyChanges(zoneID string, changes *inMemoryChange) error {
	if err := c.validateChangeBatch(zoneID, changes); err != nil {
		return err
	}
	for _, newEndpoint := range changes.Create {
		if _, ok := c.zones[zoneID][newEndpoint.Name]; !ok {
			c.zones[zoneID][newEndpoint.Name] = make([]*inMemoryRecord, 0)
		}
		c.zones[zoneID][newEndpoint.Name] = append(c.zones[zoneID][newEndpoint.Name], newEndpoint)
	}
	for _, updateEndpoint := range changes.UpdateNew {
		for _, rec := range c.zones[zoneID][updateEndpoint.Name] {
			if rec.Type == updateEndpoint.Type {
				rec.Target = updateEndpoint.Target
				break
			}
		}
	}
	for _, deleteEndpoint := range changes.Delete {
		newSet := make([]*inMemoryRecord, 0)
		for _, rec := range c.zones[zoneID][deleteEndpoint.Name] {
			if rec.Type != deleteEndpoint.Type {
				newSet = append(newSet, rec)
			}
		}
		c.zones[zoneID][deleteEndpoint.Name] = newSet
	}
	return nil
}

func (c *inMemoryClient) updateMesh(mesh map[string]map[string]bool, endpoint *inMemoryRecord) error {
	if _, exists := mesh[endpoint.Name]; exists {
		if mesh[endpoint.Name][endpoint.Type] {
			return ErrInvalidBatchRequest
		}
		mesh[endpoint.Name][endpoint.Type] = true
		return nil
	}
	mesh[endpoint.Name] = map[string]bool{endpoint.Type: true}
	return nil
}

// validateChangeBatch validates that the changes passed to InMemory DNS provider is valid
func (c *inMemoryClient) validateChangeBatch(zone string, changes *inMemoryChange) error {
	curZone, ok := c.zones[zone]
	if !ok {
		return ErrZoneNotFound
	}
	mesh := map[string]map[string]bool{}
	for _, newEndpoint := range changes.Create {
		if c.findByType(newEndpoint.Type, curZone[newEndpoint.Name]) != nil {
			return ErrRecordAlreadyExists
		}
		if err := c.updateMesh(mesh, newEndpoint); err != nil {
			return err
		}
	}
	for _, updateEndpoint := range changes.UpdateNew {
		if c.findByType(updateEndpoint.Type, curZone[updateEndpoint.Name]) == nil {
			return ErrRecordNotFound
		}
		if err := c.updateMesh(mesh, updateEndpoint); err != nil {
			return err
		}
	}
	for _, updateOldEndpoint := range changes.UpdateOld {
		if rec := c.findByType(updateOldEndpoint.Type, curZone[updateOldEndpoint.Name]); rec == nil || rec.Target != updateOldEndpoint.Target {
			return ErrRecordNotFound
		}
	}
	for _, deleteEndpoint := range changes.Delete {
		if rec := c.findByType(deleteEndpoint.Type, curZone[deleteEndpoint.Name]); rec == nil || rec.Target != deleteEndpoint.Target {
			return ErrRecordNotFound
		}
		if err := c.updateMesh(mesh, deleteEndpoint); err != nil {
			return err
		}
	}
	return nil
}

func (c *inMemoryClient) findByType(recordType string, records []*inMemoryRecord) *inMemoryRecord {
	for _, record := range records {
		if record.Type == recordType {
			return record
		}
	}
	return nil
}
