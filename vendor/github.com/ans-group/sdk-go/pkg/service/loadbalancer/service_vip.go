package loadbalancer

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVIPs retrieves a list of VIPs
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	return connection.InvokeRequestAll(s.GetVIPsPaginated, parameters)
}

// GetVIPsPaginated retrieves a paginated list of VIPs
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VIP], error) {
	body, err := s.getVIPsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVIPsPaginated), err
}

func (s *Service) getVIPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VIP], error) {
	body := &connection.APIResponseBodyData[[]VIP]{}

	response, err := s.connection.Get("/loadbalancers/v2/vips", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVIP retrieves a single VIP by id
func (s *Service) GetVIP(vipID int) (VIP, error) {
	body, err := s.getVIPResponseBody(vipID)

	return body.Data, err
}

func (s *Service) getVIPResponseBody(vipID int) (*connection.APIResponseBodyData[VIP], error) {
	body := &connection.APIResponseBodyData[VIP]{}

	if vipID < 1 {
		return body, fmt.Errorf("invalid vip id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/vips/%d", vipID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VIPNotFoundError{ID: vipID}
		}

		return nil
	})
}
