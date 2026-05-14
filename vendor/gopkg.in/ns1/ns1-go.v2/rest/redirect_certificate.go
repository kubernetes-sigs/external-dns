package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/ns1/ns1-go.v2/rest/model/redirect"
)

// RedirectCertificateService handles 'redirect/certificates' endpoint.
type RedirectCertificateService service

// List returns the existing redirect certificates.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectCertificateService) List() ([]*redirect.Certificate, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "redirect/certificates", nil)
	if err != nil {
		return nil, nil, err
	}

	certList := redirect.CertificateList{}
	var resp *http.Response
	if s.client.FollowPagination {
		resp, err = s.client.DoWithPagination(req, &certList, s.nextCerts)
	} else {
		resp, err = s.client.Do(req, &certList)
	}
	if err != nil {
		return nil, resp, err
	}

	return certList.Results, resp, nil
}

// Get takes a redirect config id and returns a single config.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectCertificateService) Get(certId string) (*redirect.Certificate, *http.Response, error) {
	path := fmt.Sprintf("redirect/certificates/%s", certId)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var cert redirect.Certificate
	var resp *http.Response
	resp, err = s.client.Do(req, &cert)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if strings.HasSuffix(err.Message, " not found") {
				return nil, resp, ErrRedirectCertificateNotFound
			}
		}
		return nil, resp, err
	}

	return &cert, resp, nil
}

// Create takes a *Certificate and creates a new redirect.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectCertificateService) Create(domain string) (*redirect.Certificate, *http.Response, error) {

	req, err := s.client.NewRequest("PUT", "redirect/certificates", redirect.NewCertificate(domain))
	if err != nil {
		return nil, nil, err
	}

	// Update redirect fields with data from api(ensure consistent)
	var cert redirect.Certificate
	resp, err := s.client.Do(req, &cert)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if err.Message == "certificate already exists" {
				return nil, resp, ErrRedirectCertificateExists
			}
		}
		return nil, resp, err
	}

	return &cert, resp, nil
}

// Update takes a certificate id and requests it to be renewed.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectCertificateService) Update(certId string) (*http.Response, error) {

	path := fmt.Sprintf("redirect/certificates/%s", certId)

	req, err := s.client.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	// Update redirect fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if strings.HasSuffix(err.Message, " not found") {
				return resp, ErrRedirectCertificateNotFound
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete takes a certificate id and requests it to be revoked.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectCertificateService) Delete(certId string) (*http.Response, error) {
	path := fmt.Sprintf("redirect/certificates/%s", certId)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if strings.HasSuffix(err.Message, " not found") {
				return resp, ErrRedirectCertificateNotFound
			}
		}
		return resp, err
	}

	return resp, nil
}

// nextCerts is a pagination helper than gets and appends another list of redirect configs
// to the passed list.
func (s *RedirectCertificateService) nextCerts(v *interface{}, uri string) (*http.Response, error) {
	tmpcertList := redirect.CertificateList{}
	resp, err := s.client.getURI(&tmpcertList, uri)
	if err != nil {
		return resp, err
	}

	certList, ok := (*v).(*redirect.CertificateList)
	if !ok {
		return nil, fmt.Errorf(
			"incorrect value for v, expected value of type *redirect.CertificateList, got: %T", v,
		)
	}
	certList.Total = tmpcertList.Total
	certList.Count += tmpcertList.Count
	certList.Results = append(certList.Results, tmpcertList.Results...)
	return resp, nil
}

var (
	ErrRedirectCertificateNil = errors.New("parameter missing")
	// ErrRedirectCertificateExists bundles PUT create error.
	ErrRedirectCertificateExists = errors.New("redirect certificate id already exists")
	// ErrRedirectCertificateNotFound bundles GET/POST/DELETE error.
	ErrRedirectCertificateNotFound = errors.New("redirect certificate id not found")
)
