package dnsimple

import (
	"context"
	"fmt"
)

// CertificatesService handles communication with the certificate related
// methods of the DNSimple API.
//
// See https://developer.dnsimple.com/v2/certificates/
type CertificatesService struct {
	client *Client
}

// Certificate represents a Certificate in DNSimple.
type Certificate struct {
	ID                  int64    `json:"id,omitempty"`
	DomainID            int64    `json:"domain_id,omitempty"`
	ContactID           int64    `json:"contact_id,omitempty"`
	CommonName          string   `json:"common_name,omitempty"`
	AlternateNames      []string `json:"alternate_names,omitempty"`
	Years               int      `json:"years,omitempty"`
	State               string   `json:"state,omitempty"`
	AuthorityIdentifier string   `json:"authority_identifier,omitempty"`
	AutoRenew           bool     `json:"auto_renew"`
	CreatedAt           string   `json:"created_at,omitempty"`
	UpdatedAt           string   `json:"updated_at,omitempty"`
	ExpiresOn           string   `json:"expires_on,omitempty"`
	CertificateRequest  string   `json:"csr,omitempty"`
}

// CertificateBundle represents a container for all the PEM-encoded X509 certificate entities,
// such as the private key, the server certificate and the intermediate chain.
type CertificateBundle struct {
	// CertificateRequest       string   `json:"csr,omitempty"`
	PrivateKey               string   `json:"private_key,omitempty"`
	ServerCertificate        string   `json:"server,omitempty"`
	RootCertificate          string   `json:"root,omitempty"`
	IntermediateCertificates []string `json:"chain,omitempty"`
}

// CertificatePurchase represents a Certificate Purchase in DNSimple.
type CertificatePurchase struct {
	ID            int64  `json:"id,omitempty"`
	CertificateID int64  `json:"new_certificate_id,omitempty"`
	State         string `json:"state,omitempty"`
	AutoRenew     bool   `json:"auto_renew,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	UpdatedAt     string `json:"updated_at,omitempty"`
}

// CertificateRenewal represents a Certificate Renewal in DNSimple.
type CertificateRenewal struct {
	ID               int64  `json:"id,omitempty"`
	OldCertificateID int64  `json:"old_certificate_id,omitempty"`
	NewCertificateID int64  `json:"new_certificate_id,omitempty"`
	State            string `json:"state,omitempty"`
	AutoRenew        bool   `json:"auto_renew,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

// LetsencryptCertificateAttributes is a set of attributes to purchase a Let's Encrypt certificate.
type LetsencryptCertificateAttributes struct {
	ContactID      int64    `json:"contact_id,omitempty"`
	Name           string   `json:"name,omitempty"`
	AutoRenew      bool     `json:"auto_renew,omitempty"`
	AlternateNames []string `json:"alternate_names,omitempty"`
}

func certificatePath(accountID, domainIdentifier string, certificateID int64) (path string) {
	path = fmt.Sprintf("%v/certificates", domainPath(accountID, domainIdentifier))
	if certificateID != 0 {
		path += fmt.Sprintf("/%v", certificateID)
	}
	return
}

func letsencryptCertificatePath(accountID, domainIdentifier string, certificateID int64) (path string) {
	path = fmt.Sprintf("%v/certificates/letsencrypt", domainPath(accountID, domainIdentifier))
	if certificateID != 0 {
		path += fmt.Sprintf("/%v", certificateID)
	}
	return
}

// CertificateResponse represents a response from an API method that returns a Certificate struct.
type CertificateResponse struct {
	Response
	Data *Certificate `json:"data"`
}

// CertificateBundleResponse represents a response from an API method that returns a CertificatBundle struct.
type CertificateBundleResponse struct {
	Response
	Data *CertificateBundle `json:"data"`
}

// CertificatesResponse represents a response from an API method that returns a collection of Certificate struct.
type CertificatesResponse struct {
	Response
	Data []Certificate `json:"data"`
}

// CertificatePurchaseResponse represents a response from an API method that returns a CertificatePurchase struct.
type CertificatePurchaseResponse struct {
	Response
	Data *CertificatePurchase `json:"data"`
}

// CertificateRenewalResponse represents a response from an API method that returns a CertificateRenewal struct.
type CertificateRenewalResponse struct {
	Response
	Data *CertificateRenewal `json:"data"`
}

// ListCertificates lists the certificates for a domain in the account.
//
// See https://developer.dnsimple.com/v2/certificates#listCertificates
func (s *CertificatesService) ListCertificates(ctx context.Context, accountID, domainIdentifier string, options *ListOptions) (*CertificatesResponse, error) {
	path := versioned(certificatePath(accountID, domainIdentifier, 0))
	certificatesResponse := &CertificatesResponse{}

	path, err := addURLQueryOptions(path, options)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.get(ctx, path, certificatesResponse)
	if err != nil {
		return certificatesResponse, err
	}

	certificatesResponse.HTTPResponse = resp
	return certificatesResponse, nil
}

// GetCertificate gets the details of a certificate.
//
// See https://developer.dnsimple.com/v2/certificates#getCertificate
func (s *CertificatesService) GetCertificate(ctx context.Context, accountID, domainIdentifier string, certificateID int64) (*CertificateResponse, error) {
	path := versioned(certificatePath(accountID, domainIdentifier, certificateID))
	certificateResponse := &CertificateResponse{}

	resp, err := s.client.get(ctx, path, certificateResponse)
	if err != nil {
		return nil, err
	}

	certificateResponse.HTTPResponse = resp
	return certificateResponse, nil
}

// DownloadCertificate gets the PEM-encoded certificate,
// along with the root certificate and intermediate chain.
//
// See https://developer.dnsimple.com/v2/certificates#downloadCertificate
func (s *CertificatesService) DownloadCertificate(ctx context.Context, accountID, domainIdentifier string, certificateID int64) (*CertificateBundleResponse, error) {
	path := versioned(certificatePath(accountID, domainIdentifier, certificateID) + "/download")
	certificateBundleResponse := &CertificateBundleResponse{}

	resp, err := s.client.get(ctx, path, certificateBundleResponse)
	if err != nil {
		return nil, err
	}

	certificateBundleResponse.HTTPResponse = resp
	return certificateBundleResponse, nil
}

// GetCertificatePrivateKey gets the PEM-encoded certificate private key.
//
// See https://developer.dnsimple.com/v2/certificates#getCertificatePrivateKey
func (s *CertificatesService) GetCertificatePrivateKey(ctx context.Context, accountID, domainIdentifier string, certificateID int64) (*CertificateBundleResponse, error) {
	path := versioned(certificatePath(accountID, domainIdentifier, certificateID) + "/private_key")
	certificateBundleResponse := &CertificateBundleResponse{}

	resp, err := s.client.get(ctx, path, certificateBundleResponse)
	if err != nil {
		return nil, err
	}

	certificateBundleResponse.HTTPResponse = resp
	return certificateBundleResponse, nil
}

// PurchaseLetsencryptCertificate purchases a Let's Encrypt certificate.
//
// See https://developer.dnsimple.com/v2/certificates/#purchaseLetsencryptCertificate
func (s *CertificatesService) PurchaseLetsencryptCertificate(ctx context.Context, accountID, domainIdentifier string, certificateAttributes LetsencryptCertificateAttributes) (*CertificatePurchaseResponse, error) {
	path := versioned(letsencryptCertificatePath(accountID, domainIdentifier, 0))
	certificatePurchaseResponse := &CertificatePurchaseResponse{}

	resp, err := s.client.post(ctx, path, certificateAttributes, certificatePurchaseResponse)
	if err != nil {
		return nil, err
	}

	certificatePurchaseResponse.HTTPResponse = resp
	return certificatePurchaseResponse, nil
}

// IssueLetsencryptCertificate issues a pending Let's Encrypt certificate purchase order.
//
// See https://developer.dnsimple.com/v2/certificates/#issueLetsencryptCertificate
func (s *CertificatesService) IssueLetsencryptCertificate(ctx context.Context, accountID, domainIdentifier string, certificateID int64) (*CertificateResponse, error) {
	path := versioned(letsencryptCertificatePath(accountID, domainIdentifier, certificateID) + "/issue")
	certificateResponse := &CertificateResponse{}

	resp, err := s.client.post(ctx, path, nil, certificateResponse)
	if err != nil {
		return nil, err
	}

	certificateResponse.HTTPResponse = resp
	return certificateResponse, nil
}

// PurchaseLetsencryptCertificateRenewal purchases a Let's Encrypt certificate renewal.
//
// See https://developer.dnsimple.com/v2/certificates/#purchaseRenewalLetsencryptCertificate
func (s *CertificatesService) PurchaseLetsencryptCertificateRenewal(ctx context.Context, accountID, domainIdentifier string, certificateID int64, certificateAttributes LetsencryptCertificateAttributes) (*CertificateRenewalResponse, error) {
	path := versioned(letsencryptCertificatePath(accountID, domainIdentifier, certificateID) + "/renewals")
	certificateRenewalResponse := &CertificateRenewalResponse{}

	resp, err := s.client.post(ctx, path, certificateAttributes, certificateRenewalResponse)
	if err != nil {
		return nil, err
	}

	certificateRenewalResponse.HTTPResponse = resp
	return certificateRenewalResponse, nil
}

// IssueLetsencryptCertificateRenewal issues a pending Let's Encrypt certificate renewal order.
//
// See https://developer.dnsimple.com/v2/certificates/#issueRenewalLetsencryptCertificate
func (s *CertificatesService) IssueLetsencryptCertificateRenewal(ctx context.Context, accountID, domainIdentifier string, certificateID, certificateRenewalID int64) (*CertificateResponse, error) {
	path := versioned(letsencryptCertificatePath(accountID, domainIdentifier, certificateID) + fmt.Sprintf("/renewals/%d/issue", certificateRenewalID))
	certificateResponse := &CertificateResponse{}

	resp, err := s.client.post(ctx, path, nil, certificateResponse)
	if err != nil {
		return nil, err
	}

	certificateResponse.HTTPResponse = resp
	return certificateResponse, nil
}
