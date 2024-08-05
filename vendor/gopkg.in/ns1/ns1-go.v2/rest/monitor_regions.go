package rest

import (
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/monitor"
)

// MonitorRegionsService handles 'monitoring/regions' endpoint.
type MonitorRegionsService service

// List returns all available monitoring regions.
//
// API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#listMonitoringRegions
func (s *MonitorRegionsService) List() ([]*monitor.Region, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "monitoring/regions", nil)
	if err != nil {
		return nil, nil, err
	}

	regions := []*monitor.Region{}

	resp, err := s.client.Do(req, &regions)
	if err != nil {
		return nil, resp, err
	}

	return regions, resp, nil
}
