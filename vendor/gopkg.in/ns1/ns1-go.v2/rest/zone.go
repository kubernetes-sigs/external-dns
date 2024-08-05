package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// ZonesService handles 'zones' endpoint.
type ZonesService service

// List returns all active zones and basic zone configuration details for each.
//
// NS1 API docs: https://ns1.com/api/#zones-get
func (s *ZonesService) List() ([]*dns.Zone, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "zones", nil)
	if err != nil {
		return nil, nil, err
	}

	zl := []*dns.Zone{}
	var resp *http.Response
	if s.client.FollowPagination {
		resp, err = s.client.DoWithPagination(req, &zl, s.nextZones)
	} else {
		resp, err = s.client.Do(req, &zl)
	}
	if err != nil {
		return nil, resp, err
	}

	return zl, resp, nil
}

// Get takes a zone name and returns a single active zone and its basic configuration details.
//
//	records Optional Query Parameter, if false records array in payload returns empty
//
// NS1 API docs: https://ns1.com/api/#zones-zone-get
func (s *ZonesService) Get(zone string, records bool) (*dns.Zone, *http.Response, error) {
	path := fmt.Sprintf("zones/%s", zone)
	if !records {
		path = fmt.Sprintf("%s%s", path, "?records=false")
	}

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var z dns.Zone
	var resp *http.Response
	if s.client.FollowPagination {
		resp, err = s.client.DoWithPagination(req, &z, s.nextRecords)
	} else {
		resp, err = s.client.Do(req, &z)
	}
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "zone not found" {
				return nil, resp, ErrZoneMissing
			}
		}
		return nil, resp, err
	}

	return &z, resp, nil
}

// Create takes a *Zone and creates a new DNS zone.
//
// NS1 API docs: https://ns1.com/api/#zones-put
func (s *ZonesService) Create(z *dns.Zone) (*http.Response, error) {
	path := fmt.Sprintf("zones/%s", z.Zone)

	req, err := s.client.NewRequest("PUT", path, &z)
	if err != nil {
		return nil, err
	}

	// Update zones fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &z)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "zone already exists" ||
				err.(*Error).Message == "invalid: FQDN already exists" ||
				err.(*Error).Message == "invalid: FQDN already exists in the view" {
				return resp, ErrZoneExists
			}
		}
		return resp, err
	}

	return resp, nil
}

// Update takes a *Zone and modifies basic details of a DNS zone.
//
// NS1 API docs: https://ns1.com/api/#zones-post
func (s *ZonesService) Update(z *dns.Zone) (*http.Response, error) {
	path := fmt.Sprintf("zones/%s", z.Zone)

	req, err := s.client.NewRequest("POST", path, &z)
	if err != nil {
		return nil, err
	}

	// Update zones fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &z)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "zone not found" {
				return resp, ErrZoneMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete takes a zone and destroys an existing DNS zone and all records in the zone.
//
// NS1 API docs: https://ns1.com/api/#zones-delete
func (s *ZonesService) Delete(zone string) (*http.Response, error) {
	path := fmt.Sprintf("zones/%s", zone)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Message == "zone not found" {
				return resp, ErrZoneMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// nextZones is a pagination helper than gets and appends another list of zones
// to the passed list.
func (s *ZonesService) nextZones(v *interface{}, uri string) (*http.Response, error) {
	tmpZl := []*dns.Zone{}
	resp, err := s.client.getURI(&tmpZl, uri)
	if err != nil {
		return resp, err
	}
	zoneList, ok := (*v).(*[]*dns.Zone)
	if !ok {
		return nil, fmt.Errorf(
			"incorrect value for v, expected value of type *[]*dns.Zone, got: %T", v,
		)
	}
	*zoneList = append(*zoneList, tmpZl...)
	return resp, nil
}

// nextRecords is a pagination helper tha gets and appends another set of
// records to the passed zone.
func (s *ZonesService) nextRecords(v *interface{}, uri string) (*http.Response, error) {
	var tmpZone dns.Zone
	resp, err := s.client.getURI(&tmpZone, uri)
	if err != nil {
		return resp, err
	}
	zone, ok := (*v).(*dns.Zone)
	if !ok {
		return nil, fmt.Errorf(
			"incorrect value for v, expected value of type *dns.Zone, got: %T", v,
		)
	}
	// Aside from Records, the rest of the zone data is identical in the
	// paginated response.
	zone.Records = append(zone.Records, tmpZone.Records...)
	return resp, nil
}

var (
	// ErrZoneExists bundles PUT create error.
	ErrZoneExists = errors.New("zone already exists")
	// ErrZoneMissing bundles GET/POST/DELETE error.
	ErrZoneMissing = errors.New("zone does not exist")
)
