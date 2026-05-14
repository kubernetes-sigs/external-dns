package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// RecordSearchService handles 'dns/record/search' endpoint.
type RecordSearchService service

// Find takes query parameters and returns matching DNS records.
func (s *RecordSearchService) Search(params string) (*dns.RecordSearchResult, *http.Response, error) {
	path := fmt.Sprintf("dns/record/search?%s", params)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var r dns.RecordSearchResult
	resp, err := s.client.Do(req, &r)
	if err != nil {
		return nil, resp, err
	}

	for _, record := range r.Results {
		for _, answer := range record.Answers {
			if err := processSearchAnswer(answer); err != nil {
				fmt.Printf("error processing search answer: %v", err)
				continue
			}
		}
	}

	return &r, resp, nil
}

func processSearchAnswer(answer *dns.SearchAnswer) error {
	if len(answer.Raw) == 0 {
		return nil
	}

	var rawArray []interface{}
	if err := json.Unmarshal(answer.Raw, &rawArray); err != nil {
		return err
	}

	answer.Rdata = make([]string, len(rawArray))
	for i, v := range rawArray {
		switch val := v.(type) {
		case string:
			answer.Rdata[i] = val
		case float64:
			answer.Rdata[i] = fmt.Sprintf("%.0f", val)
		case int:
			answer.Rdata[i] = fmt.Sprintf("%d", val)
		default:
			answer.Rdata[i] = fmt.Sprintf("%v", val)
		}
	}

	return nil
}

// ZoneSearchService handles 'dns/zone/search' endpoint.
type ZoneSearchService service

func (s *ZoneSearchService) Search(params string) (*dns.ZoneSearchResult, *http.Response, error) {
	path := fmt.Sprintf("dns/zone/search?%s", params)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var result dns.ZoneSearchResult
	resp, err := s.client.Do(req, &result)
	if err != nil {
		return nil, resp, err
	}

	return &result, resp, nil
}
