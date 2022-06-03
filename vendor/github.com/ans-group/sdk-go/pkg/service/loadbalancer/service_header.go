package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetHeaders retrieves a list of headers
func (s *Service) GetHeaders(parameters connection.APIRequestParameters) ([]Header, error) {
	return connection.InvokeRequestAll(s.GetHeadersPaginated, parameters)
}

// GetHeadersPaginated retrieves a paginated list of headers
func (s *Service) GetHeadersPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Header], error) {
	body, err := s.getHeadersPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetHeadersPaginated), err
}

func (s *Service) getHeadersPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Header], error) {
	body := &connection.APIResponseBodyData[[]Header]{}

	response, err := s.connection.Get("/loadbalancers/v2/headers", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
