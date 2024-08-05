package rest

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dataset"
)

var (
	// ErrDatasetNotFound bundles GET/POST/DELETE not found errors.
	ErrDatasetNotFound = errors.New("dataset not found")
)

// DatasetsService handles 'datasets' endpoint.
type DatasetsService service

// List returns the configured datasets.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#listDataset
func (s *DatasetsService) List() ([]*dataset.Dataset, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "datasets", nil)
	if err != nil {
		return nil, nil, err
	}

	dts := make([]*dataset.Dataset, 0)
	resp, err := s.client.Do(req, &dts)

	return dts, resp, err
}

// Get takes a dataset id and returns all its data.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#getDataset
func (s *DatasetsService) Get(dtID string) (*dataset.Dataset, *http.Response, error) {
	path := fmt.Sprintf("datasets/%s", dtID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var dt dataset.Dataset
	resp, err := s.client.Do(req, &dt)
	if err != nil {
		var clientErr *Error
		switch {
		case errors.As(err, &clientErr):
			if strings.HasSuffix(clientErr.Message, " not found") {
				return nil, resp, ErrDatasetNotFound
			}
		}
		return nil, resp, err
	}

	return &dt, resp, nil
}

// Create takes a *Dataset and creates a new dataset.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#createDataset
func (s *DatasetsService) Create(dt *dataset.Dataset) (*dataset.Dataset, *http.Response, error) {
	req, err := s.client.NewRequest("PUT", "datasets", dt)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, dt)
	return dt, resp, err
}

// Delete takes a dataset id and deletes it.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#deleteDataset
func (s *DatasetsService) Delete(dtID string) (*http.Response, error) {
	path := fmt.Sprintf("datasets/%s", dtID)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		var clientErr *Error
		switch {
		case errors.As(err, &clientErr):
			if strings.HasSuffix(clientErr.Message, " not found") {
				return resp, ErrDatasetNotFound
			}
		}
		return resp, err
	}

	return resp, nil
}

// GetReport takes a dataset id and a report id and returns bytes.Buffer which contains the file contents
// Additionally, file name can be grabbed from the 'Content-Disposition' header in the http.Response
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#getDatasetReport
func (s *DatasetsService) GetReport(dtID string, reportID string) (*bytes.Buffer, *http.Response, error) {
	path := fmt.Sprintf("datasets/%s/reports/%s", dtID, reportID)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var buf bytes.Buffer

	resp, err := s.client.Do(req, &buf)
	if err != nil {
		var clientErr *Error
		switch {
		case errors.As(err, &clientErr):
			if strings.HasSuffix(clientErr.Message, " not found") {
				return nil, resp, ErrDatasetNotFound
			}
		}
		return nil, resp, err
	}

	return &buf, resp, nil
}
