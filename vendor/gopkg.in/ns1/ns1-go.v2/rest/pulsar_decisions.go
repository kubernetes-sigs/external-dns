package rest

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gopkg.in/ns1/ns1-go.v2/rest/model/pulsar"
)

// PulsarDecisionsService handles 'pulsar/apps/APPID/jobs/JOBID/decisions' endpoint.
type PulsarDecisionsService service

// addQueryParams adds query parameters from DecisionsQueryParams to the path
func addQueryParams(path string, params *pulsar.DecisionsQueryParams) string {
	values := url.Values{}

	if params.Start != 0 {
		values.Add("start", strconv.FormatInt(params.Start, 10))
	}
	if params.End != 0 {
		values.Add("end", strconv.FormatInt(params.End, 10))
	}
	if params.Period != "" {
		values.Add("period", params.Period)
	}
	if params.Area != "" {
		values.Add("area", params.Area)
	}
	if params.ASN != "" {
		values.Add("asn", params.ASN)
	}
	if params.Job != "" {
		values.Add("job", params.Job)
	}
	if len(params.Jobs) > 0 {
		values.Add("jobs", strings.Join(params.Jobs, ","))
	}
	if params.Record != "" {
		values.Add("record", params.Record)
	}
	if params.Result != "" {
		values.Add("result", params.Result)
	}
	if params.Agg != "" {
		values.Add("agg", params.Agg)
	}
	if params.Geo != "" {
		values.Add("geo", params.Geo)
	}
	if params.ZoneID != "" {
		values.Add("zone_id", params.ZoneID)
	}
	if params.CustomerID != 0 {
		values.Add("customer_id", strconv.FormatInt(params.CustomerID, 10))
	}

	if len(values) > 0 {
		return path + "?" + values.Encode()
	}
	return path
}

// GetDecisions retrieves decisions data with optional filtering parameters.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-analytics-get
func (s *PulsarDecisionsService) GetDecisions(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsResponse, *http.Response, error) {
	path := "pulsar/query/decisions"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var decisions pulsar.DecisionsResponse
	resp, err := s.client.Do(req, &decisions)
	if err != nil {
		return nil, resp, err
	}

	return &decisions, resp, nil
}

// GetDecisionsGraphRegion retrieves regional graph data for decisions.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-graph-region-get
func (s *PulsarDecisionsService) GetDecisionsGraphRegion(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsGraphRegionResponse, *http.Response, error) {
	path := "pulsar/query/decisions/graph/region"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var graphRegion pulsar.DecisionsGraphRegionResponse
	resp, err := s.client.Do(req, &graphRegion)
	if err != nil {
		return nil, resp, err
	}

	return &graphRegion, resp, nil
}

// GetDecisionsGraphTime retrieves time-series graph data for decisions.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-graph-time-get
func (s *PulsarDecisionsService) GetDecisionsGraphTime(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsGraphTimeResponse, *http.Response, error) {
	path := "pulsar/query/decisions/graph/time"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var graphTime pulsar.DecisionsGraphTimeResponse
	resp, err := s.client.Do(req, &graphTime)
	if err != nil {
		return nil, resp, err
	}

	return &graphTime, resp, nil
}

// GetDecisionsArea retrieves area-based decisions data.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-area-get
func (s *PulsarDecisionsService) GetDecisionsArea(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsAreaResponse, *http.Response, error) {
	path := "pulsar/query/decisions/area"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var area pulsar.DecisionsAreaResponse
	resp, err := s.client.Do(req, &area)
	if err != nil {
		return nil, resp, err
	}

	return &area, resp, nil
}

// GetDecisionsASN retrieves ASN-based decisions data.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-asn-get
func (s *PulsarDecisionsService) GetDecisionsASN(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsASNResponse, *http.Response, error) {
	path := "pulsar/query/decisions/asn"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var asn pulsar.DecisionsASNResponse
	resp, err := s.client.Do(req, &asn)
	if err != nil {
		return nil, resp, err
	}

	return &asn, resp, nil
}

// GetDecisionsResultsTime retrieves time-based results data for decisions.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-results-time-get
func (s *PulsarDecisionsService) GetDecisionsResultsTime(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsResultsTimeResponse, *http.Response, error) {
	path := "pulsar/query/decisions/results/time"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var resultsTime pulsar.DecisionsResultsTimeResponse
	resp, err := s.client.Do(req, &resultsTime)
	if err != nil {
		return nil, resp, err
	}

	return &resultsTime, resp, nil
}

// GetDecisionsResultsArea retrieves area-based results data for decisions.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-results-area-get
func (s *PulsarDecisionsService) GetDecisionsResultsArea(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsResultsAreaResponse, *http.Response, error) {
	path := "pulsar/query/decisions/results/area"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var resultsArea pulsar.DecisionsResultsAreaResponse
	resp, err := s.client.Do(req, &resultsArea)
	if err != nil {
		return nil, resp, err
	}

	return &resultsArea, resp, nil
}

// GetFiltersTime retrieves time-based filter data for decisions.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-filters-time-get
func (s *PulsarDecisionsService) GetFiltersTime(params *pulsar.DecisionsQueryParams) (*pulsar.FiltersTimeResponse, *http.Response, error) {
	path := "pulsar/query/decisions/filters/time"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var filtersTime pulsar.FiltersTimeResponse
	resp, err := s.client.Do(req, &filtersTime)
	if err != nil {
		return nil, resp, err
	}

	return &filtersTime, resp, nil
}

// GetDecisionCustomer retrieves customer-specific decision data.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decision-customer-get
func (s *PulsarDecisionsService) GetDecisionCustomer(customerID string, params *pulsar.DecisionsQueryParams) (*pulsar.DecisionCustomerResponse, *http.Response, error) {
	path := fmt.Sprintf("pulsar/query/decision/customer/%s", customerID)
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var customer pulsar.DecisionCustomerResponse
	resp, err := s.client.Do(req, &customer)
	if err != nil {
		return nil, resp, err
	}

	return &customer, resp, nil
}

// GetDecisionCustomerUndetermined retrieves undetermined customer-specific decision data.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decision-customer-undetermined-get
func (s *PulsarDecisionsService) GetDecisionCustomerUndetermined(customerID string, params *pulsar.DecisionsQueryParams) (*pulsar.DecisionCustomerResponse, *http.Response, error) {
	path := fmt.Sprintf("pulsar/query/decision/customer/%s/undetermined", customerID)
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var customer pulsar.DecisionCustomerResponse
	resp, err := s.client.Do(req, &customer)
	if err != nil {
		return nil, resp, err
	}

	return &customer, resp, nil
}

// GetDecisionRecord retrieves record-specific decision data.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decision-record-get
func (s *PulsarDecisionsService) GetDecisionRecord(customerID, domain, recType string, params *pulsar.DecisionsQueryParams) (*pulsar.DecisionCustomerResponse, *http.Response, error) {
	path := fmt.Sprintf("pulsar/query/decision/customer/%s/record/%s/%s", customerID, domain, recType)
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var record pulsar.DecisionCustomerResponse
	resp, err := s.client.Do(req, &record)
	if err != nil {
		return nil, resp, err
	}

	return &record, resp, nil
}

// GetDecisionRecordUndetermined retrieves undetermined record-specific decision data.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decision-record-undetermined-get
func (s *PulsarDecisionsService) GetDecisionRecordUndetermined(customerID, domain, recType string, params *pulsar.DecisionsQueryParams) (*pulsar.DecisionCustomerResponse, *http.Response, error) {
	path := fmt.Sprintf("pulsar/query/decision/customer/%s/record/%s/%s/undetermined", customerID, domain, recType)
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var record pulsar.DecisionCustomerResponse
	resp, err := s.client.Do(req, &record)
	if err != nil {
		return nil, resp, err
	}

	return &record, resp, nil
}

// GetDecisionTotal retrieves total decision count.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decision-total-get
func (s *PulsarDecisionsService) GetDecisionTotal(customerID string, params *pulsar.DecisionsQueryParams) (*pulsar.DecisionTotalResponse, *http.Response, error) {
	path := fmt.Sprintf("pulsar/query/decision/customer/%s/total", customerID)
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var total pulsar.DecisionTotalResponse
	resp, err := s.client.Do(req, &total)
	if err != nil {
		return nil, resp, err
	}

	return &total, resp, nil
}

// GetPulsarDecisionsRecords retrieves records-based decisions data.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-records-get
func (s *PulsarDecisionsService) GetPulsarDecisionsRecords(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsRecordsResponse, *http.Response, error) {
	path := "pulsar/query/decisions/records"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var records pulsar.DecisionsRecordsResponse
	resp, err := s.client.Do(req, &records)
	if err != nil {
		return nil, resp, err
	}

	return &records, resp, nil
}

// GetPulsarDecisionsResultsRecord retrieves record-based results data for decisions.
//
// NS1 API docs: https://ns1.com/api/#pulsar-decisions-results-record-get
func (s *PulsarDecisionsService) GetPulsarDecisionsResultsRecord(params *pulsar.DecisionsQueryParams) (*pulsar.DecisionsResultsRecordResponse, *http.Response, error) {
	path := "pulsar/query/decisions/results/record"
	if params != nil {
		path = addQueryParams(path, params)
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var resultsRecord pulsar.DecisionsResultsRecordResponse
	resp, err := s.client.Do(req, &resultsRecord)
	if err != nil {
		return nil, resp, err
	}

	return &resultsRecord, resp, nil
}
