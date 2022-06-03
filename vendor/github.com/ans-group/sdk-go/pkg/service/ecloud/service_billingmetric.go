package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBillingMetrics retrieves a list of billing metrics
func (s *Service) GetBillingMetrics(parameters connection.APIRequestParameters) ([]BillingMetric, error) {
	return connection.InvokeRequestAll(s.GetBillingMetricsPaginated, parameters)
}

// GetBillingMetricsPaginated retrieves a paginated list of billing metrics
func (s *Service) GetBillingMetricsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[BillingMetric], error) {
	body, err := s.getBillingMetricsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetBillingMetricsPaginated), err
}

func (s *Service) getBillingMetricsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]BillingMetric], error) {
	body := &connection.APIResponseBodyData[[]BillingMetric]{}

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

func (s *Service) getBillingMetricResponseBody(metricID string) (*connection.APIResponseBodyData[BillingMetric], error) {
	body := &connection.APIResponseBodyData[BillingMetric]{}

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
