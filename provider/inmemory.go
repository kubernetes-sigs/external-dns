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

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

var (
	defaultType = ""

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

type zone map[string][]*InMemoryRecord

// InMemoryProvider - dns provider only used for testing purposes
// initialized as dns provider with no records
type InMemoryProvider struct {
	zones          map[string]zone
	OnApplyChanges func(changes *plan.Changes)
	OnRecords      func()
}

// NewInMemoryProvider returns InMemoryProvider DNS provider interface implementation
func NewInMemoryProvider() *InMemoryProvider {
	return &InMemoryProvider{
		zones:          map[string]zone{},
		OnApplyChanges: func(changes *plan.Changes) {},
		OnRecords:      func() {},
	}
}

// InMemoryRecord - record stored in memory
// has additional fields:
// Type - type of string (TODO: Type should probably be part of endpoint struct)
// Payload - string - additional information stored
type InMemoryRecord struct {
	Type    string
	Payload string
	*endpoint.Endpoint
}

// CreateZone adds new zone if not present
func (im *InMemoryProvider) CreateZone(newZone string) error {
	if _, exist := im.zones[newZone]; exist {
		return ErrZoneAlreadyExists
	}
	im.zones[newZone] = zone{}
	return nil
}

// Records returns the list of endpoints
func (im *InMemoryProvider) Records(zone string) ([]*endpoint.Endpoint, error) {
	defer im.OnRecords()

	if _, exists := im.zones[zone]; !exists {
		return nil, ErrZoneNotFound
	}
	return im.endpoints(zone), nil
}

// ApplyChanges simply modifies records in memory
// error checking occurs before any modifications are made, i.e. batch processing
// create record - record should not exist
// update/delete record - record should exist
// create/update/delete lists should not have overlapping records
func (im *InMemoryProvider) ApplyChanges(zone string, changes *plan.Changes) error {
	defer im.OnApplyChanges(changes)

	if err := im.validateChangeBatch(zone, changes); err != nil {
		return err
	}

	for _, newEndpoint := range changes.Create {
		im.zones[zone][newEndpoint.DNSName] = append(im.zones[zone][newEndpoint.DNSName], &InMemoryRecord{
			Type:     newEndpoint.RecordType,
			Endpoint: newEndpoint,
		})
	}
	for _, updateEndpoint := range changes.UpdateNew {
		recordToUpdate := im.findByType(updateEndpoint.RecordType, im.zones[zone][updateEndpoint.DNSName])
		recordToUpdate.Target = updateEndpoint.Target
	}
	for _, deleteEndpoint := range changes.Delete {
		newRecordSet := make([]*InMemoryRecord, 0)
		for _, record := range im.zones[zone][deleteEndpoint.DNSName] {
			if record.Type != deleteEndpoint.RecordType {
				newRecordSet = append(newRecordSet, record)
			}
		}
		im.zones[zone][deleteEndpoint.DNSName] = newRecordSet
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
		if im.findByType(newEndpoint.RecordType, existing[newEndpoint.DNSName]) != nil {
			return ErrRecordAlreadyExists
		}
		if _, exists := mesh[newEndpoint.DNSName]; exists {
			if mesh[newEndpoint.DNSName][newEndpoint.RecordType] {
				return ErrInvalidBatchRequest
			}
			mesh[newEndpoint.DNSName][newEndpoint.RecordType] = true
			continue
		}
		mesh[newEndpoint.DNSName] = map[string]bool{newEndpoint.RecordType: true}
	}
	for _, updateEndpoint := range changes.UpdateNew {
		if im.findByType(updateEndpoint.RecordType, existing[updateEndpoint.DNSName]) == nil {
			return ErrRecordNotFound
		}
		if _, exists := mesh[updateEndpoint.DNSName]; exists {
			if mesh[updateEndpoint.DNSName][updateEndpoint.RecordType] {
				return ErrInvalidBatchRequest
			}
			mesh[updateEndpoint.DNSName][updateEndpoint.RecordType] = true
			continue
		}
		mesh[updateEndpoint.DNSName] = map[string]bool{updateEndpoint.RecordType: true}
	}
	for _, updateOldEndpoint := range changes.UpdateOld {
		if rec := im.findByType(updateOldEndpoint.RecordType, existing[updateOldEndpoint.DNSName]); rec == nil || rec.Target != updateOldEndpoint.Target {
			return ErrRecordNotFound
		}
	}
	for _, deleteEndpoint := range changes.Delete {
		if rec := im.findByType(deleteEndpoint.RecordType, existing[deleteEndpoint.DNSName]); rec == nil || rec.Target != deleteEndpoint.Target {
			return ErrRecordNotFound
		}
		if _, exists := mesh[deleteEndpoint.DNSName]; exists {
			if mesh[deleteEndpoint.DNSName][deleteEndpoint.RecordType] {
				return ErrInvalidBatchRequest
			}
			mesh[deleteEndpoint.DNSName][deleteEndpoint.RecordType] = true
			continue
		}
		mesh[deleteEndpoint.DNSName] = map[string]bool{deleteEndpoint.RecordType: true}
	}
	return nil
}

func (im *InMemoryProvider) findByType(recordType string, records []*InMemoryRecord) *InMemoryRecord {
	for _, record := range records {
		if record.Type == recordType {
			return record
		}
	}
	return nil
}

func (im *InMemoryProvider) endpoints(zone string) []*endpoint.Endpoint {
	endpoints := make([]*endpoint.Endpoint, 0)
	if zoneRecords, exists := im.zones[zone]; exists {
		for _, recordsPerName := range zoneRecords {
			for _, record := range recordsPerName {
				endpoints = append(endpoints, record.Endpoint)
			}
		}
	}
	return endpoints
}
