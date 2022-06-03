package ssl

import "github.com/ans-group/sdk-go/pkg/connection"

// ValidateCertificate validates a certificate
func (s *Service) ValidateCertificate(req ValidateRequest) (CertificateValidation, error) {
	body, err := s.validateCertificateResponseBody(req)

	return body.Data, err
}

func (s *Service) validateCertificateResponseBody(req ValidateRequest) (*connection.APIResponseBodyData[CertificateValidation], error) {
	body := &connection.APIResponseBodyData[CertificateValidation]{}

	response, err := s.connection.Post("/ssl/v1/validate", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
