package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetBillingMetrics retrieves a list of billing metrics
func (s *Service) GetBillingMetrics(parameters connection.APIRequestParameters) ([]BillingMetric, error) {
	var metrics []BillingMetric

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetBillingMetricsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, metric := range response.(*PaginatedBillingMetric).Items {
			metrics = append(metrics, metric)
		}
	}

	return metrics, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetBillingMetricsPaginated retrieves a paginated list of billing metrics
func (s *Service) GetBillingMetricsPaginated(parameters connection.APIRequestParameters) (*PaginatedBillingMetric, error) {
	body, err := s.getBillingMetricsPaginatedResponseBody(parameters)

	return NewPaginatedBillingMetric(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetBillingMetricsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getBillingMetricsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetBillingMetricSliceResponseBody, error) {
	body := &GetBillingMetricSliceResponseBody{}

	response, err := s.connection.Get("/ecloud/v2/billing-metrics", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetBillingMetric retrieves a single billing metrics by id
func (s *Service) GetBillingMetric(metricID string) (BillingMetric, error) {
	body, err := s.getBillingMetricResponseBody(metricID)

	return body.Data, err
}

func (s *Service) getBillingMetricResponseBody(metricID string) (*GetBillingMetricResponseBody, error) {
	body := &GetBillingMetricResponseBody{}

	if metricID == "" {
		return body, fmt.Errorf("invalid metric id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/billing-metrics/%s", metricID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &BillingMetricNotFoundError{ID: metricID}
		}

		return nil
	})
}
