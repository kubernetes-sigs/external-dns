package certificate

import (
	"encoding/json"

	"github.com/go-gandi/go-gandi/config"
	"github.com/go-gandi/go-gandi/internal/client"
)

// New returns an instance of the Certificate API client
func New(config config.Config) *Certificate {
	client := client.New(config.APIKey, config.PersonalAccessToken, config.APIURL, config.SharingID, config.Debug, config.DryRun, config.Timeout)
	client.SetEndpoint("certificate/")
	return &Certificate{client: *client}
}

// ListCertificates requests the list of issued certificates
func (g *Certificate) ListCertificates() (certificates []CertificateType, err error) {
	_, elements, err := g.client.GetCollection("issued-certs", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var certificate CertificateType
		err := json.Unmarshal(element, &certificate)
		if err != nil {
			return nil, err
		}
		certificates = append(certificates, certificate)
	}
	return certificates, nil
}

// GetCertificate request details of an issued certificates
func (g *Certificate) GetCertificate(certificateId string) (certificate CertificateType, err error) {
	_, err = g.client.Get("issued-certs/"+certificateId, nil, &certificate)
	return
}

// GetCertificateData requests certificate data for the specified ID.
func (g *Certificate) GetCertificateData(certificateId string) (data []byte, err error) {
	_, data, err = g.client.GetBytes("issued-certs/"+certificateId+"/crt", nil)
	return
}

// CreateCertificate creates a certificate
func (g *Certificate) CreateCertificate(req CreateCertificateRequest) (response CreateCertificateResponse, err error) {
	_, err = g.client.Post("issued-certs", req, &response)
	return
}

// DeleteCertificate revokes a certificate
func (g *Certificate) DeleteCertificate(certificateId string) (response ErrorResponse, err error) {
	_, err = g.client.Delete("issued-certs/"+certificateId, nil, &response)
	return
}

// ListPackages lists certificate package types
func (g *Certificate) ListPackages() (packages []Package, err error) {
	_, elements, err := g.client.GetCollection("packages", nil)
	if err != nil {
		return nil, err
	}
	for _, element := range elements {
		var package_ Package
		err := json.Unmarshal(element, &package_)
		if err != nil {
			return nil, err
		}
		packages = append(packages, package_)
	}
	return packages, nil
}

// GetIntermediateCertificate requests intermediate certificate for the
// specified type, which can be one of "cert_std", "cert_free", "cert_bus",
// "cert_pro".
func (g *Certificate) GetIntermediateCertificate(typ string) (data []byte, err error) {
	_, data, err = g.client.GetBytes("pem/"+typ, nil)
	return
}
