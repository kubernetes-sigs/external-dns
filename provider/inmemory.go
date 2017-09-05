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

	log "github.com/Sirupsen/logrus"

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
	// ErrDuplicateRecordFound when record is repeated in create/update/delete
	ErrDuplicateRecordFound = errors.New("invalid batch request")
)

// InMemoryProvider - dns provider only used for testing purposes
// initialized as dns provider with no records
type InMemoryProvider struct {
	domain         DomainFilter
	client         *inMemoryClient
	filter         *filter
	OnApplyChanges func(changes *plan.Changes)
	OnRecords      func()
}

// InMemoryOption allows to extend in-memory provider
type InMemoryOption func(*InMemoryProvider)

// InMemoryZone stores endpoints for each zone
type InMemoryZone struct {
	ZoneID    string
	Endpoints []*endpoint.Endpoint
}

// InMemoryWithLogging injects logging when ApplyChanges is called
func InMemoryWithLogging() InMemoryOption {
	return func(p *InMemoryProvider) {
		p.OnApplyChanges = func(changes *plan.Changes) {
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
func InMemoryWithDomain(domainFilter DomainFilter) InMemoryOption {
	return func(p *InMemoryProvider) {
		p.domain = domainFilter
	}
}

// NewInMemoryProvider returns InMemoryProvider DNS provider interface implementation
func NewInMemoryProvider(opts ...InMemoryOption) *InMemoryProvider {
	im := &InMemoryProvider{
		filter:         &filter{},
		OnApplyChanges: func(changes *plan.Changes) {},
		OnRecords:      func() {},
		domain:         NewDomainFilter([]string{""}),
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
func (im *InMemoryProvider) Zones() []InMemoryZone {
	return im.filter.Zones(im.client.Zones())
}

// Records returns the list of endpoints
func (im *InMemoryProvider) Records() ([]*endpoint.Endpoint, error) {
	defer im.OnRecords()

	endpoints := make([]*endpoint.Endpoint, 0)

	for _, zone := range im.Zones() {
		records, err := im.client.Records(zone.ZoneID)
		if err != nil {
			return nil, err
		}

		endpoints = append(endpoints, records...)
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
	for _, zone := range zones {
		perZoneChanges[zone.ZoneID] = &plan.Changes{}
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
			Create:    perZoneChanges[zoneID].Create,
			UpdateNew: perZoneChanges[zoneID].UpdateNew,
			UpdateOld: perZoneChanges[zoneID].UpdateOld,
			Delete:    perZoneChanges[zoneID].Delete,
		}
		err := im.client.ApplyChanges(zoneID, change)
		if err != nil {
			return err
		}
	}

	return nil
}

type filter struct {
	domain string
}

// Zones filters map[zoneID]zoneName for names having f.domain as suffix
func (f *filter) Zones(zones []InMemoryZone) []InMemoryZone {
	result := []InMemoryZone{}
	for _, zone := range zones {
		if strings.HasSuffix(zone.ZoneID, f.domain) {
			result = append(result, InMemoryZone{ZoneID: zone.ZoneID, Endpoints: zone.Endpoints})
		}
	}
	return result
}

// EndpointZoneID determines zoneID for an endpoint from []InMemoryZone by taking the highest level domain which has an
// endpoint.DNSName that matches the suffix of the zone
// returns empty string if no match found
func (f *filter) EndpointZoneID(endpoint *endpoint.Endpoint, zones []InMemoryZone) string {
	var matchZoneName string
	var domainLevel int
	for _, zone := range zones {
		domainLevel = len(strings.Split(endpoint.DNSName, ".")[1:])
		if strings.HasSuffix(endpoint.DNSName, zone.ZoneID) &&
			domainLevel >= len(strings.Split(matchZoneName, ".")) {
			matchZoneName = zone.ZoneID
		}
	}
	return matchZoneName
}

type inMemoryChange struct {
	Create    []*endpoint.Endpoint
	UpdateNew []*endpoint.Endpoint
	UpdateOld []*endpoint.Endpoint
	Delete    []*endpoint.Endpoint
}

type inMemoryClient struct {
	zones []InMemoryZone
}

func newInMemoryClient() *inMemoryClient {
	return &inMemoryClient{zones: []InMemoryZone{}}
}

func (c *inMemoryClient) Records(zone string) ([]*endpoint.Endpoint, error) {
	zoneExists := false
	records := []*endpoint.Endpoint{}
	for _, z := range c.zones {
		if z.ZoneID == zone {
			zoneExists = true
			records = append(records, z.Endpoints...)
		}
	}
	if !zoneExists {
		return nil, ErrZoneNotFound
	}

	return records, nil
}

func (c *inMemoryClient) Zones() []InMemoryZone {
	return c.zones
}

func (c *inMemoryClient) CreateZone(zone string) error {
	for _, z := range c.zones {
		if z.ZoneID == zone {
			return ErrZoneAlreadyExists
		}
	}
	c.zones = append(c.zones, InMemoryZone{ZoneID: zone})

	return nil
}

func (c *inMemoryClient) ApplyChanges(zoneID string, changes *inMemoryChange) error {
	if err := c.validateChangeBatch(zoneID, changes); err != nil {
		return err
	}
	for _, newEndpoint := range changes.Create {
		for i, z := range c.zones {
			if z.ZoneID == zoneID {
				c.zones[i].Endpoints = append(c.zones[i].Endpoints, newEndpoint)
			}
		}
	}
	for _, updateEndpoint := range changes.UpdateNew {
		for i, z := range c.zones {
			if z.ZoneID == zoneID {
				for i2, rec := range z.Endpoints {
					if c.recordsAreTheSame(rec, updateEndpoint) {
						c.zones[i].Endpoints[i2] = updateEndpoint
					}
				}
			}
		}
	}
	for _, deleteEndpoint := range changes.Delete {
		newSet := make([]*endpoint.Endpoint, 0)
		for _, z := range c.zones {
			if z.ZoneID == zoneID {
				for _, rec := range z.Endpoints {
					if !c.recordsAreTheSame(rec, deleteEndpoint) {
						newSet = append(newSet, rec)
					}
				}
			}
		}
		for i, z := range c.zones {
			if z.ZoneID == zoneID {
				c.zones[i].Endpoints = newSet
			}
		}

	}
	return nil
}

func (c *inMemoryClient) updateMesh(mesh *InMemoryZone, record *endpoint.Endpoint) error {
	if rec := c.findRecord(record, mesh.Endpoints); rec != nil {
		return ErrDuplicateRecordFound
	}
	mesh.Endpoints = append(mesh.Endpoints, record)
	return nil
}

// validateChangeBatch validates that the changes passed to InMemory DNS provider is valid
func (c *inMemoryClient) validateChangeBatch(zone string, changes *inMemoryChange) error {
	zoneExists := false
	var curZone InMemoryZone
	for _, z := range c.zones {
		if z.ZoneID == zone {
			curZone = z
			zoneExists = true
		}
	}
	if !zoneExists {
		return ErrZoneNotFound
	}

	mesh := &InMemoryZone{}
	for _, newEndpoint := range changes.Create {
		if c.findRecord(newEndpoint, curZone.Endpoints) != nil {
			return ErrRecordAlreadyExists
		}
		if err := c.updateMesh(mesh, newEndpoint); err != nil {
			return err
		}
	}
	for _, updateEndpoint := range changes.UpdateNew {
		if c.findRecord(updateEndpoint, curZone.Endpoints) == nil {
			return ErrRecordNotFound
		}
		if err := c.updateMesh(mesh, updateEndpoint); err != nil {
			return err
		}
	}
	for _, updateOldEndpoint := range changes.UpdateOld {
		if rec := c.findRecord(updateOldEndpoint, curZone.Endpoints); rec == nil || rec.Target != updateOldEndpoint.Target {
			return ErrRecordNotFound
		}
	}
	for _, deleteEndpoint := range changes.Delete {
		if rec := c.findRecord(deleteEndpoint, curZone.Endpoints); rec == nil || rec.Target != deleteEndpoint.Target {
			return ErrRecordNotFound
		}
		if err := c.updateMesh(mesh, deleteEndpoint); err != nil {
			return err
		}
	}
	return nil
}

func (c *inMemoryClient) findRecord(record *endpoint.Endpoint, records []*endpoint.Endpoint) *endpoint.Endpoint {
	for _, r := range records {
		if c.recordsAreTheSame(r, record) {
			return r
		}
	}
	return nil
}

func (c *inMemoryClient) recordsAreTheSame(recordOne *endpoint.Endpoint, recordTwo *endpoint.Endpoint) bool {
	if recordOne != nil && recordOne.DNSName == recordTwo.DNSName && recordOne.RecordType == recordTwo.RecordType {
		if recordOne.Policy.HasAWSRoute53Policy() {
			if recordOne.Policy.AWSRoute53.SetIdentifier == recordTwo.Policy.AWSRoute53.SetIdentifier {
				return true
			}
		} else {
			return true
		}
	}
	return false
}
