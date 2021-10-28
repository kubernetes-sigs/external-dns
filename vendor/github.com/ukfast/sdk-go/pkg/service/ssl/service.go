package ssl

import (
	"github.com/ukfast/sdk-go/pkg/connection"
)

// SSLService is an interface for managing SSL certificates
type SSLService interface {
	GetCertificates(parameters connection.APIRequestParameters) ([]Certificate, error)
	GetCertificatesPaginated(parameters connection.APIRequestParameters) (*PaginatedCertificate, error)
	GetCertificate(certificateID int) (Certificate, error)
	GetCertificateContent(certificateID int) (CertificateContent, error)
	GetCertificatePrivateKey(certificateID int) (CertificatePrivateKey, error)
	GetReport(domainName string) (Report, error)
	GetRecommendations(domainName string) (Recommendations, error)
	ValidateCertificate(req ValidateRequest) (CertificateValidation, error)
}

// Service implements SSLService for managing
// SSL certificates via the UKFast API
type Service struct {
	connection connection.Connection
}

// NewService returns a new instance of SSLService
func NewService(connection connection.Connection) *Service {
	return &Service{
		connection: connection,
	}
}
