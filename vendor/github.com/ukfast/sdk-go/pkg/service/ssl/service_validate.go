package ssl

// ValidateCertificate validates a certificate
func (s *Service) ValidateCertificate(req ValidateRequest) (CertificateValidation, error) {
	body, err := s.validateCertificateResponseBody(req)

	return body.Data, err
}

func (s *Service) validateCertificateResponseBody(req ValidateRequest) (*GetCertificateValidationResponseBody, error) {
	body := &GetCertificateValidationResponseBody{}

	response, err := s.connection.Post("/ssl/v1/validate", &req)
	if err != nil {
		return body, err
	}

	return body, response.HandleResponse(body, nil)
}
