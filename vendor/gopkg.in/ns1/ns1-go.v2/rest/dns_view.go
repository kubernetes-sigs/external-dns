package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// DNSViewService handles 'views/' endpoint.
type DNSViewService service

// List returns all DNS Views
//
// NS1 API docs: https://ns1.com/api#getlist-all-dns-views
func (s *DNSViewService) List() ([]*dns.View, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "views", nil)
	if err != nil {
		return nil, nil, err
	}

	var vl []*dns.View
	resp, err := s.client.Do(req, &vl)
	if err != nil {
		return nil, resp, err
	}

	return vl, resp, nil
}

// Create takes a *dns.DNSView and creates a new DNS View.
//
// The given DNSView must have at least the name
// NS1 API docs: https://ns1.com/api#putcreate-a-dns-view
func (s *DNSViewService) Create(v *dns.View) (*http.Response, error) {
	req, err := s.client.NewRequest("PUT", fmt.Sprintf("/v1/views/%s", v.Name), v)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusConflict {
				return nil, ErrViewExists
			}
		}

		return resp, err
	}

	return resp, nil
}

// Get takes a DNS view name and returns DNSView struct.
//
// NS1 API docs: https://ns1.com/api#getview-dns-view-details
func (s *DNSViewService) Get(viewName string) (*dns.View, *http.Response, error) {
	path := fmt.Sprintf("views/%s", viewName)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var v dns.View
	resp, err := s.client.Do(req, &v)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusNotFound {
				return nil, resp, ErrViewMissing
			}
		}
		return nil, resp, err
	}

	return &v, resp, nil
}

// Update takes a *dns.DNSView and updates the DNS view with same name on NS1.
//
// NS1 API docs: https://ns1.com/api#postedit-a-dns-view
func (s *DNSViewService) Update(v *dns.View) (*http.Response, error) {
	path := fmt.Sprintf("views/%s", v.Name)

	req, err := s.client.NewRequest("POST", path, &v)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &v)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusNotFound {
				return resp, ErrViewMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete takes a DNS view name, and removes an existing DNS view
//
// NS1 API docs: https://ns1.com/api#deletedelete-a-dns-view
func (s *DNSViewService) Delete(viewName string) (*http.Response, error) {
	path := fmt.Sprintf("views/%s", viewName)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusNotFound {
				return resp, ErrViewMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// GetPreferences returns a map[string]int of preferences.
//
// NS1 API docs: https://ns1.com/api#getget-dns-view-preference
func (s *DNSViewService) GetPreferences() (map[string]int, *http.Response, error) {
	path := "config/views/preference"

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	m := make(map[string]int)
	resp, err := s.client.Do(req, &m)
	if err != nil {
		return nil, resp, err
	}

	return m, resp, nil
}

// UpdatePreferences takes a map[string]int and returns a map[string]int of preferences.
//
// NS1 API docs: https://ns1.com/api#postedit-dns-view-preference
func (s *DNSViewService) UpdatePreferences(m map[string]int) (map[string]int, *http.Response, error) {
	path := "config/views/preference"

	req, err := s.client.NewRequest("POST", path, m)
	if err != nil {
		return nil, nil, err
	}

	mapUpdated := make(map[string]int)
	resp, err := s.client.Do(req, &mapUpdated)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusNotFound {
				return nil, resp, ErrViewMissing
			}
		}
		return nil, resp, err
	}

	return mapUpdated, resp, nil
}

var (
	// ErrViewExists bundles CREATE error.
	ErrViewExists = errors.New("DNS view already exists")

	// ErrViewMissing bundles GET error.
	ErrViewMissing = errors.New("DNS view not found")
)
