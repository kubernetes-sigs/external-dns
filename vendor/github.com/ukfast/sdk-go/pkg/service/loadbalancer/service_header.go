package loadbalancer

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetHeaders retrieves a list of headers
func (s *Service) GetHeaders(parameters connection.APIRequestParameters) ([]Header, error) {
	var headers []Header

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHeadersPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, header := range response.(*PaginatedHeader).Items {
			headers = append(headers, header)
		}
	}

	return headers, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetHeadersPaginated retrieves a paginated list of headers
func (s *Service) GetHeadersPaginated(parameters connection.APIRequestParameters) (*PaginatedHeader, error) {
	body, err := s.getHeadersPaginatedResponseBody(parameters)

	return NewPaginatedHeader(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetHeadersPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getHeadersPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetHeaderSliceResponseBody, error) {
	body := &GetHeaderSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/headers", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
