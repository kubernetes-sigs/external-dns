package ecloud

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetVPNProfileGroups retrieves a list of VPN profile groups
func (s *Service) GetVPNProfileGroups(parameters connection.APIRequestParameters) ([]VPNProfileGroup, error) {
	var profileGroups []VPNProfileGroup

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNProfileGroupsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, group := range response.(*PaginatedVPNProfileGroup).Items {
			profileGroups = append(profileGroups, group)
		}
	}

	return profileGroups, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetVPNProfileGroupsPaginated retrieves a paginated list of VPN profile groups
func (s *Service) GetVPNProfileGroupsPaginated(parameters connection.APIRequestParameters) (*PaginatedVPNProfileGroup, error) {
	body, err := s.getVPNProfileGroupsPaginatedResponseBody(parameters)

	return NewPaginatedVPNProfileGroup(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetVPNProfileGroupsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getVPNProfileGroupsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetVPNProfileGroupSliceResponseBody, error) {
	body := &GetVPNProfileGroupSliceResponseBody{}

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

func (s *Service) getVPNProfileGroupResponseBody(profileGroupID string) (*GetVPNProfileGroupResponseBody, error) {
	body := &GetVPNProfileGroupResponseBody{}

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
