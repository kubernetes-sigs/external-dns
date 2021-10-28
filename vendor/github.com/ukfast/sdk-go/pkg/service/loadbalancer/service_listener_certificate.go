package loadbalancer

import (
	"fmt"

	"github.com/ukfast/sdk-go/pkg/connection"
)

// GetListenerCertificates retrieves a list of certificates
func (s *Service) GetListenerCertificates(listenerID int, parameters connection.APIRequestParameters) ([]Certificate, error) {
	var certificates []Certificate

	getFunc := func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenerCertificatesPaginated(listenerID, p)
	}

	responseFunc := func(response connection.Paginated) {
		for _, certificate := range response.(*PaginatedCertificate).Items {
			certificates = append(certificates, certificate)
		}
	}

	return certificates, connection.InvokeRequestAll(getFunc, responseFunc, parameters)
}

// GetListenerCertificatesPaginated retrieves a paginated list of certificates
func (s *Service) GetListenerCertificatesPaginated(listenerID int, parameters connection.APIRequestParameters) (*PaginatedCertificate, error) {
	body, err := s.getListenerCertificatesPaginatedResponseBody(listenerID, parameters)

	return NewPaginatedCertificate(func(p connection.APIRequestParameters) (connection.Paginated, error) {
		return s.GetListenerCertificatesPaginated(listenerID, p)
	}, parameters, body.Metadata.Pagination, body.Data), err
}

func (s *Service) getListenerCertificatesPaginatedResponseBody(listenerID int, parameters connection.APIRequestParameters) (*GetCertificateSliceResponseBody, error) {
	body := &GetCertificateSliceResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs", listenerID), parameters)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}

// GetListenerCertificate retrieves a single certificate by id
func (s *Service) GetListenerCertificate(listenerID int, certificateID int) (Certificate, error) {
	body, err := s.getListenerCertificateResponseBody(listenerID, certificateID)

	return body.Data, err
}

func (s *Service) getListenerCertificateResponseBody(listenerID int, certificateID int) (*GetCertificateResponseBody, error) {
	body := &GetCertificateResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	if certificateID < 1 {
		return body, fmt.Errorf("invalid certificate id")
	}

	response, err := s.connection.Get(fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs/%d", listenerID, certificateID), connection.APIRequestParameters{})
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CertificateNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// CreateListenerCertificate creates an certificate
func (s *Service) CreateListenerCertificate(listenerID int, req CreateCertificateRequest) (int, error) {
	body, err := s.createListenerCertificateResponseBody(listenerID, req)

	return body.Data.ID, err
}

func (s *Service) createListenerCertificateResponseBody(listenerID int, req CreateCertificateRequest) (*GetCertificateResponseBody, error) {
	body := &GetCertificateResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	response, err := s.connection.Post(fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs", listenerID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CertificateNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// PatchListenerCertificate patches an certificate
func (s *Service) PatchListenerCertificate(listenerID int, certificateID int, req PatchCertificateRequest) error {
	_, err := s.patchListenerCertificateResponseBody(listenerID, certificateID, req)

	return err
}

func (s *Service) patchListenerCertificateResponseBody(listenerID int, certificateID int, req PatchCertificateRequest) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	if certificateID < 1 {
		return body, fmt.Errorf("invalid certificate id")
	}

	response, err := s.connection.Patch(fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs/%d", listenerID, certificateID), &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CertificateNotFoundError{ID: listenerID}
		}

		return nil
	})
}

// DeleteListenerCertificate deletes a certificate
func (s *Service) DeleteListenerCertificate(listenerID int, certificateID int) error {
	_, err := s.deleteListenerCertificateResponseBody(listenerID, certificateID)

	return err
}

func (s *Service) deleteListenerCertificateResponseBody(listenerID int, certificateID int) (*connection.APIResponseBody, error) {
	body := &connection.APIResponseBody{}

	if listenerID < 1 {
		return body, fmt.Errorf("invalid listener id")
	}

	if certificateID < 1 {
		return body, fmt.Errorf("invalid certificate id")
	}

	response, err := s.connection.Delete(fmt.Sprintf("/loadbalancers/v2/listeners/%d/certs/%d", listenerID, certificateID), nil)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, func(resp *connection.APIResponse) error {
		if response.StatusCode == 404 {
			return &CertificateNotFoundError{ID: listenerID}
		}

		return nil
	})
}
