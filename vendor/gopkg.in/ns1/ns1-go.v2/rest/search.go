package rest

import (
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// RecordSearchService handles 'dns/record/search' endpoint.
type RecordSearchService service

// Find takes query parameters and returns matching DNS records.
func (s *RecordSearchService) Search(params string) (*dns.SearchResult, *http.Response, error) {
	path := fmt.Sprintf("dns/record/search?%s", params)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var r dns.SearchResult
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return &r, resp, nil
}

// ZoneSearchService handles 'dns/zone/search' endpoint.
type ZoneSearchService service

// Find takes query parameters and returns matching DNS zones.
func (s *ZoneSearchService) Search(params string) (*dns.SearchResult, *http.Response, error) {
	path := fmt.Sprintf("dns/zone/search?%s", params)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var r dns.SearchResult
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	return &r, resp, nil
}
