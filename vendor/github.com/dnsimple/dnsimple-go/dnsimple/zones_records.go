package dnsimple

import (
	"context"
	"fmt"
)

// ZoneRecord represents a zone record in DNSimple.
type ZoneRecord struct {
	ID           int64    `json:"id,omitempty"`
	ZoneID       string   `json:"zone_id,omitempty"`
	ParentID     int64    `json:"parent_id,omitempty"`
	Type         string   `json:"type,omitempty"`
	Name         string   `json:"name"`
	Content      string   `json:"content,omitempty"`
	TTL          int      `json:"ttl,omitempty"`
	Priority     int      `json:"priority,omitempty"`
	SystemRecord bool     `json:"system_record,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	CreatedAt    string   `json:"created_at,omitempty"`
	UpdatedAt    string   `json:"updated_at,omitempty"`
}

// ZoneRecordAttributes represents the attributes you can send to create/update a zone record.
//
// Compared to most other calls in this library, you should not use ZoneRecord as payload for record calls.
// This is because it can lead to side effects due to the inability of go to distinguish between a non-present string
// and an empty string. Name can be both, therefore a specific struct is required.
type ZoneRecordAttributes struct {
	ZoneID   string   `json:"zone_id,omitempty"`
	Type     string   `json:"type,omitempty"`
	Name     *string  `json:"name,omitempty"`
	Content  string   `json:"content,omitempty"`
	TTL      int      `json:"ttl,omitempty"`
	Priority int      `json:"priority,omitempty"`
	Regions  []string `json:"regions,omitempty"`
}

func zoneRecordPath(accountID string, zoneName string, recordID int64) (path string) {
	path = fmt.Sprintf("/%v/zones/%v/records", accountID, zoneName)
	if recordID != 0 {
		path += fmt.Sprintf("/%v", recordID)
	}
	return
}

// ZoneRecordResponse represents a response from an API method that returns a ZoneRecord struct.
type ZoneRecordResponse struct {
	Response
	Data *ZoneRecord `json:"data"`
}

// ZoneRecordsResponse represents a response from an API method that returns a collection of ZoneRecord struct.
type ZoneRecordsResponse struct {
	Response
	Data []ZoneRecord `json:"data"`
}

// ZoneRecordListOptions specifies the optional parameters you can provide
// to customize the ZonesService.ListZoneRecords method.
type ZoneRecordListOptions struct {
	// Select records where the name matches given string.
	Name *string `url:"name,omitempty"`

	// Select records where the name contains given string.
	NameLike *string `url:"name_like,omitempty"`

	// Select records of given type.
	// Eg. TXT, A, NS.
	Type *string `url:"type,omitempty"`

	ListOptions
}

// ListRecords lists the zone records for a zone.
//
// See https://developer.dnsimple.com/v2/zones/records/#listZoneRecords
func (s *ZonesService) ListRecords(ctx context.Context, accountID string, zoneName string, options *ZoneRecordListOptions) (*ZoneRecordsResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneName, 0))
	recordsResponse := &ZoneRecordsResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, recordsResponse)
	if err != nil {
		return nil, err
	}

	recordsResponse.HTTPResponse = resp
	return recordsResponse, nil
}

// CreateRecord creates a zone record.
//
// See https://developer.dnsimple.com/v2/zones/records/#createZoneRecord
func (s *ZonesService) CreateRecord(ctx context.Context, accountID string, zoneName string, recordAttributes ZoneRecordAttributes) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneName, 0))
	recordResponse := &ZoneRecordResponse{}

	resp, err := s.client.post(ctx, path, recordAttributes, recordResponse)
	if err != nil {
		return nil, err
	}

	recordResponse.HTTPResponse = resp
	return recordResponse, nil
}

// GetRecord fetches a zone record.
//
// See https://developer.dnsimple.com/v2/zones/records/#getZoneRecord
func (s *ZonesService) GetRecord(ctx context.Context, accountID string, zoneName string, recordID int64) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneName, recordID))
	recordResponse := &ZoneRecordResponse{}

	resp, err := s.client.get(ctx, path, recordResponse)
	if err != nil {
		return nil, err
	}

	recordResponse.HTTPResponse = resp
	return recordResponse, nil
}

// UpdateRecord updates a zone record.
//
// See https://developer.dnsimple.com/v2/zones/records/#updateZoneRecord
func (s *ZonesService) UpdateRecord(ctx context.Context, accountID string, zoneName string, recordID int64, recordAttributes ZoneRecordAttributes) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneName, recordID))
	recordResponse := &ZoneRecordResponse{}
	resp, err := s.client.patch(ctx, path, recordAttributes, recordResponse)

	if err != nil {
		return nil, err
	}

	recordResponse.HTTPResponse = resp
	return recordResponse, nil
}

// DeleteRecord PERMANENTLY deletes a zone record from the zone.
//
// See https://developer.dnsimple.com/v2/zones/records/#deleteZoneRecord
func (s *ZonesService) DeleteRecord(ctx context.Context, accountID string, zoneName string, recordID int64) (*ZoneRecordResponse, error) {
	path := versioned(zoneRecordPath(accountID, zoneName, recordID))
	recordResponse := &ZoneRecordResponse{}

	resp, err := s.client.delete(ctx, path, nil, nil)
	if err != nil {
		return nil, err
	}

	recordResponse.HTTPResponse = resp
	return recordResponse, nil
}
