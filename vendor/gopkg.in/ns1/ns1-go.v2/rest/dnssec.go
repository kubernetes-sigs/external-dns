package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// DNSSECService handles 'zones/ZONE/dnssec' endpoint.
type DNSSECService service

// Get takes a zone, and returns DNSSEC information.
//
// NS1 API docs: https://ns1.com/api#get-get-dnssec-details-for-a-zone
func (s *DNSSECService) Get(zone string) (*dns.ZoneDNSSEC, *http.Response, error) {
	path := fmt.Sprintf("zones/%s/dnssec", zone)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var d dns.ZoneDNSSEC
	resp, err := s.client.Do(req, &d)

	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "zone not found" {
				return nil, resp, ErrZoneMissing
			}
			if err.(*Error).Message == "DNSSEC is not enabled on the zone" {
				return nil, resp, ErrDNSECNotEnabled
			}
		}
		return nil, resp, err
	}

	return &d, resp, nil
}

var (
	// ErrDNSECNotEnabled if DNSSEC is not enabled for the zone, regardless of
	// account-level DNSSEC permission.
	ErrDNSECNotEnabled = errors.New("DNSSEC is not enabled on the zone")
)
