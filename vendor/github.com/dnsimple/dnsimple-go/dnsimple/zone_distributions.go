package dnsimple

import (
	"context"
	"fmt"
)

// ZoneDistribution is the result of the zone distribution check.
type ZoneDistribution struct {
	Distributed bool `json:"distributed"`
}

// ZoneDistributionResponse represents a response from an API method that returns a ZoneDistribution struct.
type ZoneDistributionResponse struct {
	Response
	Data *ZoneDistribution `json:"data"`
}

// CheckZoneDistribution checks if a zone is fully distributed across DNSimple nodes.
//
// See https://developer.dnsimple.com/v2/zones/#checkZoneDistribution
func (s *ZonesService) CheckZoneDistribution(ctx context.Context, accountID string, zoneName string) (*ZoneDistributionResponse, error) {
	path := versioned(fmt.Sprintf("/%v/zones/%v/distribution", accountID, zoneName))
	zoneDistributionResponse := &ZoneDistributionResponse{}

	resp, err := s.client.get(ctx, path, zoneDistributionResponse)
	if err != nil {
		return nil, err
	}

	zoneDistributionResponse.HTTPResponse = resp
	return zoneDistributionResponse, nil
}

// CheckZoneRecordDistribution checks if a zone is fully distributed across DNSimple nodes.
//
// See https://developer.dnsimple.com/v2/zones/#checkZoneRecordDistribution
func (s *ZonesService) CheckZoneRecordDistribution(ctx context.Context, accountID string, zoneName string, recordID int64) (*ZoneDistributionResponse, error) {
	path := versioned(fmt.Sprintf("/%v/zones/%v/records/%v/distribution", accountID, zoneName, recordID))
	zoneDistributionResponse := &ZoneDistributionResponse{}

	resp, err := s.client.get(ctx, path, zoneDistributionResponse)
	if err != nil {
		return nil, err
	}

	zoneDistributionResponse.HTTPResponse = resp
	return zoneDistributionResponse, nil
}
