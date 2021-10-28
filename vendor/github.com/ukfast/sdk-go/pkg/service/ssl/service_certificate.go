package ssl

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetCertificates retrieves a list of certificates
func (s *Service) GetCertificates(parameters connection.APIRequestParameters) ([]Certificate, error) {
	var certificates []Certificate

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetCertificatesPaginated(p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, certificate := range response.(*PaginatedCertificate).Items {
			certificates = append(certificates, certificate)
		}
	}

	return certificates, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetCertificatesPaginated retrieves a paginated list of certificates
func (s *Service) GetCertificatesPaginated(parameters connection.APIRequestParameters) (*PaginatedCertificate, error) {
	body, err := s.getCertificatesPaginatedResponseBody(parameters)

	return NewPaginatedCertificate(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetCertificatesPaginated(p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getCertificatesPaginatedResponseBody(parameters connection.APIRequestParameters) (*GetCertificateSliceResponseBody, error) {
	body := &GetCertificateSliceResponseBody{}

	response, err := s.connection.Get("/ssl/v1/certificates", parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetCertificate retrieves a single certificate by id
func (s *Service) GetCertificate(certificateID int) (Certificate, error) {
	body, err := s.getCertificateResponseBody(certificateID)

	return body.Data, err
}

func (s *Service) getCertificateResponseBody(certificateID int) (*GetCertificateResponseBody, error) {
	body := &GetCertificateResponseBody{}

	if certificateID < 1 {
		return body, fmt.Errorf("invalid certificate id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ssl/v1/certificates/%d", certificateID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CertificateNotFoundError{ID: certificateID}
		}

		return nil
	})
}

// GetCertificateContent retrieves the content of an SSL certificate
func (s *Service) GetCertificateContent(certificateID int) (CertificateContent, error) {
	body, err := s.getCertificateContentResponseBody(certificateID)

	return body.Data, err
}

func (s *Service) getCertificateContentResponseBody(certificateID int) (*GetCertificateContentResponseBody, error) {
	body := &GetCertificateContentResponseBody{}

	if certificateID < 1 {
		return body, fmt.Errorf("invalid certificate id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ssl/v1/certificates/%d/download", certificateID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CertificateNotFoundError{ID: certificateID}
		}

		return nil
	})
}

// GetCertificatePrivateKey retrieves an SSL certificate private key
func (s *Service) GetCertificatePrivateKey(certificateID int) (CertificatePrivateKey, error) {
	body, err := s.getCertificatePrivateKeyResponseBody(certificateID)

	return body.Data, err
}

func (s *Service) getCertificatePrivateKeyResponseBody(certificateID int) (*GetCertificatePrivateKeyResponseBody, error) {
	body := &GetCertificatePrivateKeyResponseBody{}

	if certificateID < 1 {
		return body, fmt.Errorf("invalid certificate id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/ssl/v1/certificates/%d/private-key", certificateID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CertificateNotFoundError{ID: certificateID}
		}

		return nil
	})
}
