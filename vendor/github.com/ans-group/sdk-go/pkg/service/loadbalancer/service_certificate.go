package loadbalancer

import (
	"github.com/ans-group/sdk-go/pkg/connection"
)

// GetCertificates retrieves a list of certificates
func (s *Service) GetCertificates(parameters connection.APIRequestParameters) ([]Certificate, error) {
	return connection.InvokeRequestAll(s.GetCertificatesPaginated, parameters)
}

// GetCertificatesPaginated retrieves a paginated list of certificates
func (s *Service) GetCertificatesPaginated(parameters connection.APIRequestParameters) (*connection.Paginated[Certificate], error) {
	body, err := s.getCertificatesPaginatedResponseBody(parameters)
	return connection.NewPaginated(body, parameters, s.GetCertificatesPaginated), err
}

func (s *Service) getCertificatesPaginatedResponseBody(parameters connection.APIRequestParameters) (*connection.APIResponseBodyData[[]Certificate], error) {
	return connection.Get[[]Certificate](s.connection, "/loadbalancers/v2/certs", parameters)
}
