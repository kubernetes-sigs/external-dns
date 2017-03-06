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

package dnsprovider

import (
	"errors"

	"github.com/kubernetes-incubator/external-dns/endpoint"
	"github.com/kubernetes-incubator/external-dns/plan"
)

var (
	defaultType = "A"

	// ErrZoneNotFound error returned when specified zone does not exists
	ErrZoneNotFound = errors.New("specified zone not found")
	// ErrRecordAlreadyExists when create request is sent but record already exists
	ErrRecordAlreadyExists = errors.New("record already exists")
	// ErrRecordNotFound when update/delete request is sent but record not found
	ErrRecordNotFound = errors.New("record not found")
	// ErrInvalidBatchRequest when record is repeated in create/update/delete
	ErrInvalidBatchRequest = errors.New("record should only be specified in one list")
)

type zone map[string][]*InMemoryRecord

// InMemory - dns provider only used for testing purposes
// initialized as dns provider with no records
type InMemory struct {
	zones map[string]zone
}

// NewInMemory returns InMemory DNS provider interface
func NewInMemory() *InMemory {
	return &InMemory{
		zones: map[string]zone{},
	}
}

// InMemoryRecord - record stored in memory
// has additional fields:
// Type - type of string (TODO: Type should probably be part of endpoint struct)
// Payload - string - additional information stored
type InMemoryRecord struct {
	Type    string
	Payload string
	endpoint.Endpoint
}

// Records returns the list of endpoints
func (im *InMemory) Records(zone string) ([]endpoint.Endpoint, error) {
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
func (im *InMemory) ApplyChanges(zone string, changes *plan.Changes) error {
	if err := im.validateChangeBatch(zone, changes); err != nil {
		return err
	}

	for _, newEndpoint := range changes.Create {
		im.zones[zone][newEndpoint.DNSName] = append(im.zones[zone][newEndpoint.DNSName], &InMemoryRecord{
			Type:     defaultType,
			Endpoint: newEndpoint,
		})
	}
	for _, updateEndpoint := range changes.UpdateNew {
		recordToUpdate := im.findByType(defaultType, im.zones[zone][updateEndpoint.DNSName])
		recordToUpdate.Target = updateEndpoint.Target
	}
	for _, deleteEndpoint := range changes.Delete {
		newRecordSet := make([]*InMemoryRecord, 0)
		for _, record := range im.zones[zone][deleteEndpoint.DNSName] {
			if record.Type != defaultType {
				newRecordSet = append(newRecordSet, record)
			}
		}
		im.zones[zone][deleteEndpoint.DNSName] = newRecordSet
	}
	return nil
}

// validateChangeBatch validates that the changes passed to InMemory DNS provider is valid
func (im *InMemory) validateChangeBatch(zone string, changes *plan.Changes) error {
	existing := im.zones[zone]
	mesh := map[string]bool{}
	for _, newEndpoint := range changes.Create {
		if im.findByType(defaultType, existing[newEndpoint.DNSName]) != nil {
			return ErrRecordAlreadyExists
		}
		if _, exists := mesh[newEndpoint.DNSName]; exists {
			return ErrInvalidBatchRequest
		}
		mesh[newEndpoint.DNSName] = true
	}
	for _, updateEndpoint := range changes.UpdateNew {
		if im.findByType(defaultType, existing[updateEndpoint.DNSName]) == nil {
			return ErrRecordNotFound
		}
		if _, exists := mesh[updateEndpoint.DNSName]; exists {
			return ErrInvalidBatchRequest
		}
		mesh[updateEndpoint.DNSName] = true
	}
	for _, deleteEndpoint := range changes.Delete {
		if im.findByType(defaultType, existing[deleteEndpoint.DNSName]) == nil {
			return ErrRecordNotFound
		}
		if _, exists := mesh[deleteEndpoint.DNSName]; exists {
			return ErrInvalidBatchRequest
		}
		mesh[deleteEndpoint.DNSName] = true
	}
	return nil
}

func (im *InMemory) findByType(recordType string, records []*InMemoryRecord) *InMemoryRecord {
	for _, record := range records {
		if record.Type == recordType {
			return record
		}
	}
	return nil
}

func (im *InMemory) endpoints(zone string) []endpoint.Endpoint {
	endpoints := make([]endpoint.Endpoint, 0)
	if zoneRecords, exists := im.zones[zone]; exists {
		for _, recordsPerName := range zoneRecords {
			for _, record := range recordsPerName {
				endpoints = append(endpoints, record.Endpoint)
			}
		}
	}
	return endpoints
}
