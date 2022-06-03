package ddosx

import (
	"fmt"

	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetSSLs retrieves a list of ssls
func (s *Service) GetSSLs(parameters connection.APIRequestParameters) ([]SSL, error) {
	return connection.InvokeRequestAll(s.GetSSLsPaginated, parameters)
}

// GetSSLsPaginated retrieves a paginated list of ssls
func (s *Service) GetSSLsPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[SSL], error) {
	body, err := s.getSSLsPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetSSLsPaginated), err
}

func (s *Service) getSSLsPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]SSL], error) {
	body := &connection.APIResponseBodyData[[]SSL]{}

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

func (s *Service) getSSLResponseBody(sslID string) (*connection.APIResponseBodyData[SSL], error) {
	body := &connection.APIResponseBodyData[SSL]{}

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

func (s *Service) createSSLResponseBody(req CreateSSLRequest) (*connection.APIResponseBodyData[SSL], error) {
	body := &connection.APIResponseBodyData[SSL]{}

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

func (s *Service) patchSSLResponseBody(sslID string, req PatchSSLRequest) (*connection.APIResponseBodyData[SSL], error) {
	body := &connection.APIResponseBodyData[SSL]{}

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

func (s *Service) getSSLContentResponseBody(sslID string) (*connection.APIResponseBodyData[SSLContent], error) {
	body := &connection.APIResponseBodyData[SSLContent]{}

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

func (s *Service) getSSLPrivateKeyResponseBody(sslID string) (*connection.APIResponseBodyData[SSLPrivateKey], error) {
	body := &connection.APIResponseBodyData[SSLPrivateKey]{}

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
