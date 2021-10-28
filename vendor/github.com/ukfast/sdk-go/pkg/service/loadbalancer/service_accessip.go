package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetAccessIP retrieves a single access IP by id
func (s *Service) GetAccessIP(accessID int) (AccessIP, error) {
	body, err := s.getAccessIPResponseBody(accessID)

	return body.Data, err
}

func (s *Service) getAccessIPResponseBody(accessID int) (*GetAccessIPResponseBody, error) {
	body := &GetAccessIPResponseBody{}

	if accessID < 1 {
		return body, fmt.Errorf("invalid access id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/access-ips/%d", accessID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AccessIPNotFoundError{ID: accessID}
		}

		return nil
	})
}

// PatchAccessIP patches an access IP
func (s *Service) PatchAccessIP(accessID int, req PatchAccessIPRequest) error {
	_, err := s.patchAccessIPResponseBody(accessID, req)

	return err
}

func (s *Service) patchAccessIPResponseBody(accessID int, req PatchAccessIPRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if accessID < 1 {
		return body, fmt.Errorf("invalid access id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/access-ips/%d", accessID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AccessIPNotFoundError{ID: accessID}
		}

		return nil
	})
}

// DeleteAccessIP deletes an access IP
func (s *Service) DeleteAccessIP(accessID int) error {
	_, err := s.deleteAccessIPResponseBody(accessID)

	return err
}

func (s *Service) deleteAccessIPResponseBody(accessID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if accessID < 1 {
		return body, fmt.Errorf("invalid access id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/access-ips/%d", accessID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &AccessIPNotFoundError{ID: accessID}
		}

		return nil
	})
}
