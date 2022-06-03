package ecloud

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSSHKeyPairs retrieves a list of keypairs
func (s *Service) GetSSHKeyPairs(parameters connection.APIRequestParameters) ([]SSHKeyPair, error) {
	return connection.InvokeRequestAll(s.GetSSHKeyPairsPaginated, parameters)
}

// GetSSHKeyPairsPaginated retrieves a paginated list of keypairs
func (s *Service) GetSSHKeyPairsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SSHKeyPair], error) {
	body, err := s.getSSHKeyPairsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetSSHKeyPairsPaginated), err
}

func (s *Service) getSSHKeyPairsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]SSHKeyPair], error) {
	body := &connection.APIResponseBodyData[[]SSHKeyPair]{}

	response, err := s.connection.Get("/ecloud/v2/ssh-key-pairs", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetSSHKeyPair retrieves a single keypair by id
func (s *Service) GetSSHKeyPair(keypairID string) (SSHKeyPair, error) {
	body, err := s.getSSHKeyPairResponseBody(keypairID)

	return body.Data, err
}

func (s *Service) getSSHKeyPairResponseBody(keypairID string) (*connection.APIResponseBodyData[SSHKeyPair], error) {
	body := &connection.APIResponseBodyData[SSHKeyPair]{}

	if keypairID == "" {
		return body, fmt.Errorf("invalid SSH key pair id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ecloud/v2/ssh-key-pairs/%s", keypairID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSHKeyPairNotFoundError{ID: keypairID}
		}

		return nil
	})
}

// CreateSSHKeyPair creates a new SSHKeyPair
func (s *Service) CreateSSHKeyPair(req CreateSSHKeyPairRequest) (string, error) {
	body, err := s.createSSHKeyPairResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createSSHKeyPairResponseBody(req CreateSSHKeyPairRequest) (*connection.APIResponseBodyData[SSHKeyPair], error) {
	body := &connection.APIResponseBodyData[SSHKeyPair]{}

	response, err := s.connection.Post("/ecloud/v2/ssh-key-pairs", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchSSHKeyPair patches a SSHKeyPair
func (s *Service) PatchSSHKeyPair(keypairID string, req PatchSSHKeyPairRequest) error {
	_, err := s.patchSSHKeyPairResponseBody(keypairID, req)

	return err
}

func (s *Service) patchSSHKeyPairResponseBody(keypairID string, req PatchSSHKeyPairRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if keypairID == "" {
		return body, fmt.Errorf("invalid SSH key pair id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ecloud/v2/ssh-key-pairs/%s", keypairID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSHKeyPairNotFoundError{ID: keypairID}
		}

		return nil
	})
}

// DeleteSSHKeyPair deletes a SSHKeyPair
func (s *Service) DeleteSSHKeyPair(keypairID string) error {
	_, err := s.deleteSSHKeyPairResponseBody(keypairID)

	return err
}

func (s *Service) deleteSSHKeyPairResponseBody(keypairID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if keypairID == "" {
		return body, fmt.Errorf("invalid SSH key pair id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ecloud/v2/ssh-key-pairs/%s", keypairID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSHKeyPairNotFoundError{ID: keypairID}
		}

		return nil
	})
}
