package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetVPNProfileGroups retrieves a list of VPN profile groups
func (s *Service) GetVPNProfileGroups(parameters connection.APIRequestParameters) ([]VPNProfileGroup, error) {
	return connection.InvokeRequestAll(s.GetVPNProfileGroupsPaginated, parameters)
}

// GetVPNProfileGroupsPaginated retrieves a paginated list of VPN profile groups
func (s *Service) GetVPNProfileGroupsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[VPNProfileGroup], error) {
	body, err := s.getVPNProfileGroupsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetVPNProfileGroupsPaginated), err
}

func (s *Service) getVPNProfileGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]VPNProfileGroup], error) {
	body := &connection.APIResponseBodyData[[]VPNProfileGroup]{}

	response, err := s.connection.Get("/ecloud/v2/vpn-profile-groups", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetVPNProfileGroup retrieves a single VPN profile group by id
func (s *Service) GetVPNProfileGroup(profileGroupID string) (VPNProfileGroup, error) {
	body, err := s.getVPNProfileGroupResponseBody(profileGroupID)

	return body.Data, err
}

func (s *Service) getVPNProfileGroupResponseBody(profileGroupID string) (*connection.APIResponseBodyData[VPNProfileGroup], error) {
	body := &connection.APIResponseBodyData[VPNProfileGroup]{}

	if profileGroupID == "" {
		return body, fmt.Errorf("invalid vpn profile group id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/vpn-profile-groups/%s", profileGroupID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &VPNProfileGroupNotFoundError{ID: profileGroupID}
		}

		return nil
	})
}
