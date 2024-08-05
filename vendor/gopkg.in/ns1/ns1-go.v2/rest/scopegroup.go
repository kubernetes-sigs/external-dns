package rest

import (
	"errors"
	"fmt"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dhcp"
	"net/http"
)

// ScopeGroupService handles the 'scope group' endpoints.
type ScopeGroupService service

// List returns a list of all scope groups.
//
// NS1 API docs: https://ns1.com/api#getlist-scope-groups
func (s *ScopeGroupService) List() ([]dhcp.ScopeGroup, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "dhcp/scopegroup", nil)
	if err != nil {
		return nil, nil, err
	}

	sgs := make([]dhcp.ScopeGroup, 0)
	resp, err := s.client.Do(req, &sgs)
	return sgs, resp, err
}

// Get returns the Scope Group corresponding to the provided scope group ID.
//
// NS1 API docs: https://ns1.com/api#getview-scope-group
func (s *ScopeGroupService) Get(sgID int) (*dhcp.ScopeGroup, *http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/scopegroup/%d", sgID)
	req, err := s.client.NewRequest(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, nil, err
	}

	sg := &dhcp.ScopeGroup{}
	var resp *http.Response
	resp, err = s.client.Do(req, sg)
	if err != nil {
		return nil, resp, err
	}

	return sg, resp, nil
}

// Create creates a scope group.
// The Name field is required.
//
// NS1 API docs: https://ns1.com/api#putcreate-a-scope-group
func (s *ScopeGroupService) Create(sg *dhcp.ScopeGroup) (*dhcp.ScopeGroup, *http.Response, error) {
	switch {
	case sg.Name == "":
		return nil, nil, errors.New("the Name field is required")
	}

	req, err := s.client.NewRequest(http.MethodPut, "dhcp/scopegroup", sg)
	if err != nil {
		return nil, nil, err
	}

	respSg := new(dhcp.ScopeGroup)
	var resp *http.Response
	resp, err = s.client.Do(req, respSg)
	if err != nil {
		return nil, resp, err
	}

	return respSg, resp, nil
}

// Edit updates an existing scope group.
// The ID field is required.
//
// NS1 API docs: https://ns1.com/api#postedit-scope-group
func (s *ScopeGroupService) Edit(sg *dhcp.ScopeGroup) (*dhcp.ScopeGroup, *http.Response, error) {
	if sg.ID == nil {
		return nil, nil, errors.New("the ID field is required")
	}

	reqPath := fmt.Sprintf("dhcp/scopegroup/%d", *sg.ID)
	req, err := s.client.NewRequest(http.MethodPost, reqPath, sg)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(req, sg)
	if err != nil {
		return nil, resp, err
	}

	return sg, resp, nil
}

// Delete removes a Scope Group entirely.
//
// NS1 API docs: https://ns1.com/api#deleteremove-scope-group-by-id
func (s *ScopeGroupService) Delete(id int) (*http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/scopegroup/%d", id)
	req, err := s.client.NewRequest(http.MethodDelete, reqPath, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
