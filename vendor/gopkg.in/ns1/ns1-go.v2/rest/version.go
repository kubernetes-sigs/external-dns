package rest

import (
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// VersionsService handles 'zones/ZONE/versions' related endpoints.
type VersionsService service

// List returns all versions for a zone.
//
// NS1 API docs: https://ns1.com/api/#zones-get
func (s *VersionsService) List(zone string) ([]*dns.Version, *http.Response, error) {
	path := fmt.Sprintf("zones/%s/versions", zone)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	vl := []*dns.Version{}
	var resp *http.Response
	resp, err = s.client.Do(req, &vl)

	if err != nil {
		return nil, resp, err
	}

	return vl, resp, nil
}

// Create creates a new version for a zone
//
// NS1 API docs: https://ns1.com/api/#zones-get
func (s *VersionsService) Create(zone string, force bool) (*dns.Version, *http.Response, error) {
	path := fmt.Sprintf("zones/%s/versions?force=%t", zone, force)
	req, err := s.client.NewRequest("PUT", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var v dns.Version
	var resp *http.Response
	resp, err = s.client.Do(req, &v)
	if err != nil {
		return nil, resp, err
	}

	return &v, resp, nil
}

// Delete deletes a zone version
//
// NS1 API docs: https://ns1.com/api/#zones-get
func (s *VersionsService) Delete(zone string, versionID int) (*http.Response, error) {
	path := fmt.Sprintf("zones/%s/versions/%d", zone, versionID)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	var v dns.Version
	var resp *http.Response
	resp, err = s.client.Do(req, &v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Activate activates a zone version
//
// NS1 API docs: https://ns1.com/api/#zones-get
func (s *VersionsService) Activate(zone string, versionID int) (*http.Response, error) {
	path := fmt.Sprintf("/v1/zones/%s/versions/%d/activate", zone, versionID)
	req, err := s.client.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	var v dns.Version
	var resp *http.Response
	resp, err = s.client.Do(req, &v)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
