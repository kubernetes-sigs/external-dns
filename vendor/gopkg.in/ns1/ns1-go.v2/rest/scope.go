package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dhcp"
)

// ScopeService handles the 'scope' endpoints.
type ScopeService service

// List returns a list of all scopes.
//
// NS1 API docs: https://ns1.com/api#getlist-scopes
func (s *ScopeService) List() ([]dhcp.Scope, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "dhcp/scope", nil)
	if err != nil {
		return nil, nil, err
	}

	scs := make([]dhcp.Scope, 0)
	resp, err := s.client.Do(req, &scs)
	if err != nil {
		return nil, resp, err
	}

	return scs, resp, nil
}

// Get returns the scope corresponding to the provided scope ID.
//
// NS1 API docs: https://ns1.com/api#getview-scope-details
func (s *ScopeService) Get(scID int) (*dhcp.Scope, *http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/scope/%d", scID)
	req, err := s.client.NewRequest(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, nil, err
	}

	sc := &dhcp.Scope{}
	var resp *http.Response
	resp, err = s.client.Do(req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, nil
}

// Create creates a scope.
// The IDAddress field is required.
//
// NS1 API docs: https://ns1.com/api#putcreate-a-scope
func (s *ScopeService) Create(sc *dhcp.Scope) (*dhcp.Scope, *http.Response, error) {
	switch {
	case sc.IDAddress == nil:
		return nil, nil, errors.New("the IDAddress field is required")
	}

	req, err := s.client.NewRequest(http.MethodPut, "dhcp/scope", sc)
	if err != nil {
		return nil, nil, err
	}

	respSc := new(dhcp.Scope)
	var resp *http.Response
	resp, err = s.client.Do(req, respSc)
	if err != nil {
		return nil, resp, err
	}

	return respSc, resp, nil
}

// Edit updates an existing scope.
// The IDAddress field is required.
//
// NS1 API docs: https://ns1.com/api#postmodify-a-scope
func (s *ScopeService) Edit(sc *dhcp.Scope) (*dhcp.Scope, *http.Response, error) {
	switch {
	case sc.IDAddress == nil:
		return nil, nil, errors.New("the IDAddress field is required")
	}

	reqPath := fmt.Sprintf("dhcp/scope/%d", sc.ID)
	req, err := s.client.NewRequest(http.MethodPost, reqPath, sc)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, nil
}

// Delete removes a scope entirely.
//
// NS1 API docs: https://ns1.com/api#deleteremove-a-scope
func (s *ScopeService) Delete(id int) (*http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/scope/%d", id)
	req, err := s.client.NewRequest(http.MethodDelete, reqPath, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
