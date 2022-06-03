package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetBinds retrieves a list of binds
func (s *Service) GetBinds(parameters connection.APIRequestParameters) ([]Bind, error) {
	return connection.InvokeRequestAll(s.GetBindsPaginated, parameters)
}

// GetBindsPaginated retrieves a paginated list of binds
func (s *Service) GetBindsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Bind], error) {
	body, err := s.getBindsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetBindsPaginated), err
}

func (s *Service) getBindsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Bind], error) {
	body := &connection.APIResponseBodyData[[]Bind]{}

	response, err := s.connection.Get("/loadbalancers/v2/binds", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
