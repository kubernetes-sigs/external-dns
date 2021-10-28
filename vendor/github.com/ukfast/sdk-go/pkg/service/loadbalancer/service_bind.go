package loadbalancer

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetBinds retrieves a list of binds
func (s *Service) GetBinds(parameters connection.APIRequestParameters) ([]Bind, error) {
	var binds []Bind

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetBindsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, bind := range response.(*PaginatedBind).Items {
			binds = append(binds, bind)
		}
	}

	return binds, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetBindsPaginated retrieves a paginated list of binds
func (s *Service) GetBindsPaginated(parameters connection.APIRequestParameters) (*PaginatedBind, error) {
	body, err := s.getBindsPaginatedResponseBody(parameters)

	return NewPaginatedBind(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetBindsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getBindsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetBindSliceResponseBody, error) {
	body := &GetBindSliceResponseBody{}

	response, err := s.connection.Get("/loadbalancers/v2/binds", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
