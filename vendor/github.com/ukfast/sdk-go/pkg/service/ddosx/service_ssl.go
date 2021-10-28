package ddosx

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetSSLs retrieves a list of ssls
func (s *Service) GetSSLs(parameters connection.APIRequestParameters) ([]SSL, error) {
	var ssls []SSL

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSSLsPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, ssl := range response.(*PaginatedSSL).Items {
			ssls = append(ssls, ssl)
		}
	}

	return ssls, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetSSLsPaginated retrieves a paginated list of ssls
func (s *Service) GetSSLsPaginated(parameters connection.APIRequestParameters) (*PaginatedSSL, error) {
	body, err := s.getSSLsPaginatedResponseBody(parameters)

	return NewPaginatedSSL(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetSSLsPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getSSLsPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetSSLSliceResponseBody, error) {
	body := &GetSSLSliceResponseBody{}

	response, err := s.connection.Get("/ddosx/v1/ssls", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetSSL retrieves a single ssl by id
func (s *Service) GetSSL(sslID string) (SSL, error) {
	body, err := s.getSSLResponseBody(sslID)

	return body.Data, err
}

func (s *Service) getSSLResponseBody(sslID string) (*GetSSLResponseBody, error) {
	body := &GetSSLResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSLNotFoundError{ID: sslID}
		}

		return nil
	})
}

// CreateSSL retrieves creates an SSL
func (s *Service) CreateSSL(req CreateSSLRequest) (string, error) {
	body, err := s.createSSLResponseBody(req)

	return body.Data.ID, err
}

func (s *Service) createSSLResponseBody(req CreateSSLRequest) (*GetSSLResponseBody, error) {
	body := &GetSSLResponseBody{}

	response, err := s.connection.Post("/ddosx/v1/ssls", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// PatchSSL retrieves patches an SSL
func (s *Service) PatchSSL(sslID string, req PatchSSLRequest) (string, error) {
	body, err := s.patchSSLResponseBody(sslID, req)

	return body.Data.ID, err
}

func (s *Service) patchSSLResponseBody(sslID string, req PatchSSLRequest) (*GetSSLResponseBody, error) {
	body := &GetSSLResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSLNotFoundError{ID: sslID}
		}

		return nil
	})
}

// DeleteSSL deletes patches an SSL
func (s *Service) DeleteSSL(sslID string) error {
	_, err := s.deleteSSLResponseBody(sslID)

	return err
}

func (s *Service) deleteSSLResponseBody(sslID string) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/ddosx/v1/ssls/%s", sslID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSLNotFoundError{ID: sslID}
		}

		return nil
	})
}

// GetSSLContent retrieves a single ssl by id
func (s *Service) GetSSLContent(sslID string) (SSLContent, error) {
	body, err := s.getSSLContentResponseBody(sslID)

	return body.Data, err
}

func (s *Service) getSSLContentResponseBody(sslID string) (*GetSSLContentResponseBody, error) {
	body := &GetSSLContentResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/ssls/%s/certificates", sslID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSLNotFoundError{ID: sslID}
		}

		return nil
	})
}

// GetSSLPrivateKey retrieves a single ssl by id
func (s *Service) GetSSLPrivateKey(sslID string) (SSLPrivateKey, error) {
	body, err := s.getSSLPrivateKeyResponseBody(sslID)

	return body.Data, err
}

func (s *Service) getSSLPrivateKeyResponseBody(sslID string) (*GetSSLPrivateKeyResponseBody, error) {
	body := &GetSSLPrivateKeyResponseBody{}

	if sslID == "" {
		return body, fmt.Errorf("invalid ssl id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ddosx/v1/ssls/%s/private-key", sslID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &SSLNotFoundError{ID: sslID}
		}

		return nil
	})
}
