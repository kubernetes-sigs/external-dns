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
	"errors"
	"strings"

	log "github.com/sirupsen/logrus"

	"sigs.k8s.io/external-dns/endpoint"
	"sigs.k8s.io/external-dns/plan"
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
	// ErrDuplicateRecordFound when record is repeated in create/update/delete
	ErrDuplicateRecordFound = errors.New("invalid batch request")
)

// InMemoryProvider - dns provider only used for testing purposes
// initialized as dns provider with no records
type InMemoryProvider struct {
	domain         endpoint.DomainFilter
	client         *inMemoryClient
	filter         *filter
	OnApplyChanges func(ctx context.Context, changes *plan.Changes)
	OnRecords      func()
}

// InMemoryOption allows to extend in-memory provider
type InMemoryOption func(*InMemoryProvider)

// InMemoryWithLogging injects logging when ApplyChanges is called
func InMemoryWithLogging() InMemoryOption {
	return func(p *InMemoryProvider) {
		p.OnApplyChanges = func(ctx context.Context, changes *plan.Changes) {
			for _, v := range changes.Create {
				log.Infof("CREATE: %v", v)
			}
			for _, v := range changes.UpdateOld {
				log.Infof("UPDATE (old): %v", v)
			}
			for _, v := range changes.UpdateNew {
				log.Infof("UPDATE (new): %v", v)
			}
			for _, v := range changes.Delete {
				log.Infof("DELETE: %v", v)
			}
		}
	}
}

// InMemoryWithDomain modifies the domain on which dns zones are filtered
func InMemoryWithDomain(domainFilter endpoint.DomainFilter) InMemoryOption {
	return func(p *InMemoryProvider) {
		p.domain = domainFilter
	}
}

// InMemoryInitZones pre-seeds the InMemoryProvider with given zones
func InMemoryInitZones(zones []string) InMemoryOption {
	return func(p *InMemoryProvider) {
		for _, z := range zones {
			if err := p.CreateZone(z); err != nil {
				log.Warnf("Unable to initialize zones for inmemory provider")
			}
		}
	}
}

// NewInMemoryProvider returns InMemoryProvider DNS provider interface implementation
func NewInMemoryProvider(opts ...InMemoryOption) *InMemoryProvider {
	im := &InMemoryProvider{
		filter:         &filter{},
		OnApplyChanges: func(ctx context.Context, changes *plan.Changes) {},
		OnRecords:      func() {},
		domain:         endpoint.NewDomainFilter([]string{""}),
		client:         newInMemoryClient(),
	}

	for _, opt := range opts {
		opt(im)
	}

	return im
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
func (im *InMemoryProvider) Records(ctx context.Context) ([]*endpoint.Endpoint, error) {
	defer im.OnRecords()

	endpoints := make([]*endpoint.Endpoint, 0)

	for zoneID := range im.Zones() {
		records, err := im.client.Records(zoneID)
		if err != nil {
			return nil, err
		}

		for _, record := range records {
			ep := endpoint.NewEndpoint(record.Name, record.Type, record.Target).WithSetIdentifier(record.SetIdentifier)
			ep.Labels = record.Labels
			endpoints = append(endpoints, ep)
		}
	}

	return endpoints, nil
}

// ApplyChanges simply modifies records in memory
// error checking occurs before any modifications are made, i.e. batch processing
// create record - record should not exist
// update/delete record - record should exist
// create/update/delete lists should not have overlapping records
func (im *InMemoryProvider) ApplyChanges(ctx context.Context, changes *plan.Changes) error {
	defer im.OnApplyChanges(ctx, changes)

	perZoneChanges := map[string]*plan.Changes{}

	zones := im.Zones()
	for zoneID := range zones {
		perZoneChanges[zoneID] = &plan.Changes{}
	}

	for _, ep := range changes.Create {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		if zoneID == "" {
			continue
		}
		perZoneChanges[zoneID].Create = append(perZoneChanges[zoneID].Create, ep)
	}
	for _, ep := range changes.UpdateNew {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		if zoneID == "" {
			continue
		}
		perZoneChanges[zoneID].UpdateNew = append(perZoneChanges[zoneID].UpdateNew, ep)
	}
	for _, ep := range changes.UpdateOld {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		if zoneID == "" {
			continue
		}
		perZoneChanges[zoneID].UpdateOld = append(perZoneChanges[zoneID].UpdateOld, ep)
	}
	for _, ep := range changes.Delete {
		zoneID := im.filter.EndpointZoneID(ep, zones)
		if zoneID == "" {
			continue
		}
		perZoneChanges[zoneID].Delete = append(perZoneChanges[zoneID].Delete, ep)
	}

	for zoneID := range perZoneChanges {
		change := &inMemoryChange{
			Create:    convertToInMemoryRecord(perZoneChanges[zoneID].Create),
			UpdateNew: convertToInMemoryRecord(perZoneChanges[zoneID].UpdateNew),
			UpdateOld: convertToInMemoryRecord(perZoneChanges[zoneID].UpdateOld),
			Delete:    convertToInMemoryRecord(perZoneChanges[zoneID].Delete),
		}
		err := im.client.ApplyChanges(ctx, zoneID, change)
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
			Type:          ep.RecordType,
			Name:          ep.DNSName,
			Target:        ep.Targets[0],
			SetIdentifier: ep.SetIdentifier,
			Labels:        ep.Labels,
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
// Type - type of record
// Name - DNS name assigned to the record
// Target - target of the record
type inMemoryRecord struct {
	Type          string
	SetIdentifier string
	Name          string
	Target        string
	Labels        endpoint.Labels
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

func (c *inMemoryClient) ApplyChanges(ctx context.Context, zoneID string, changes *inMemoryChange) error {
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

func (c *inMemoryClient) updateMesh(mesh map[string]map[string]map[string]bool, record *inMemoryRecord) error {
	if _, exists := mesh[record.Name]; exists {
		if _, exists := mesh[record.Name][record.Type]; exists {
			if mesh[record.Name][record.Type][record.SetIdentifier] {
				return ErrDuplicateRecordFound
			}
			mesh[record.Name][record.Type][record.SetIdentifier] = true
			return nil
		}
		mesh[record.Name][record.Type] = map[string]bool{record.SetIdentifier: true}
		return nil
	}
	mesh[record.Name] = map[string]map[string]bool{record.Type: {record.SetIdentifier: true}}
	return nil
}

// validateChangeBatch validates that the changes passed to InMemory DNS provider is valid
func (c *inMemoryClient) validateChangeBatch(zone string, changes *inMemoryChange) error {
	curZone, ok := c.zones[zone]
	if !ok {
		return ErrZoneNotFound
	}
	mesh := map[string]map[string]map[string]bool{}
	for _, newEndpoint := range changes.Create {
		if c.findByTypeAndSetIdentifier(newEndpoint.Type, newEndpoint.SetIdentifier, curZone[newEndpoint.Name]) != nil {
			return ErrRecordAlreadyExists
		}
		if err := c.updateMesh(mesh, newEndpoint); err != nil {
			return err
		}
	}
	for _, updateEndpoint := range changes.UpdateNew {
		if c.findByTypeAndSetIdentifier(updateEndpoint.Type, updateEndpoint.SetIdentifier, curZone[updateEndpoint.Name]) == nil {
			return ErrRecordNotFound
		}
		if err := c.updateMesh(mesh, updateEndpoint); err != nil {
			return err
		}
	}
	for _, updateOldEndpoint := range changes.UpdateOld {
		if rec := c.findByTypeAndSetIdentifier(updateOldEndpoint.Type, updateOldEndpoint.SetIdentifier, curZone[updateOldEndpoint.Name]); rec == nil || rec.Target != updateOldEndpoint.Target {
			return ErrRecordNotFound
		}
	}
	for _, deleteEndpoint := range changes.Delete {
		if rec := c.findByTypeAndSetIdentifier(deleteEndpoint.Type, deleteEndpoint.SetIdentifier, curZone[deleteEndpoint.Name]); rec == nil || rec.Target != deleteEndpoint.Target {
			return ErrRecordNotFound
		}
		if err := c.updateMesh(mesh, deleteEndpoint); err != nil {
			return err
		}
	}
	return nil
}

func (c *inMemoryClient) findByTypeAndSetIdentifier(recordType, setIdentifier string, records []*inMemoryRecord) *inMemoryRecord {
	for _, record := range records {
		if record.Type == recordType && record.SetIdentifier == setIdentifier {
			return record
		}
	}
	return nil
}
