package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVIPs retrieves a list of VIPs
func (s *Service) GetVIPs(parameters connection.APIRequestParameters) ([]VIP, error) {
	var vips []VIP

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVIPsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, vip := range response.(*PaginatedVIP).Items {
			vips = append(vips, vip)
		}
	}

	return vips, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVIPsPaginated retrieves a paginated list of VIPs
func (s *Service) GetVIPsPaginated(parameters connection.APIRequestParameters) (*PaginatedVIP, error) {
	body, err := s.getVIPsPaginatedResponseBody(parameters)

	return NewPaginatedVIP(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVIPsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVIPsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVIPSliceResponseBody, error) {
	body := &GetVIPSliceResponseBody{}

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

func (s *Service) getVIPResponseBody(vipID int) (*GetVIPResponseBody, error) {
	body := &GetVIPResponseBody{}

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
