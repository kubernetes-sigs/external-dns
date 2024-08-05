package rest

import (
	"errors"
	"fmt"
	"gopkg.in/ns1/ns1-go.v2/rest/model/dhcp"
	"net/http"
)

// OptionDefService handles the 'scope group' endpoints.
type OptionDefService service

// List returns a list of all option definitions.
//
// NS1 API docs: https://ns1.com/api#getlist-dhcp-option-definitions
func (s *OptionDefService) List() ([]dhcp.OptionDef, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "dhcp/optiondef", nil)
	if err != nil {
		return nil, nil, err
	}

	ods := make([]dhcp.OptionDef, 0)
	resp, err := s.client.Do(req, &ods)
	return ods, resp, err
}

// Get returns the option definition corresponding to the provided ID.
//
// NS1 API docs: https://ns1.com/api#getview-dhcp-option-definition
func (s *OptionDefService) Get(odSpace, odKey string) (*dhcp.OptionDef, *http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/optiondef/%s/%s", odSpace, odKey)
	req, err := s.client.NewRequest(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, nil, err
	}

	od := &dhcp.OptionDef{}
	var resp *http.Response
	resp, err = s.client.Do(req, od)
	if err != nil {
		return nil, resp, err
	}

	return od, resp, nil
}

// Create creates or updates an option definition.
// The FriendlyName, Description, Code, Schema.Type fields are required.
//
// NS1 API docs: https://ns1.com/api#putcreate-an-custom-dhcp-option-definition
func (s *OptionDefService) Create(od *dhcp.OptionDef, odSpace, odKey string) (*dhcp.OptionDef, *http.Response, error) {
	switch {
	case od.FriendlyName == "":
		return nil, nil, errors.New("the FriendlyName field is required")
	case od.Description == "":
		return nil, nil, errors.New("the Description field is required")
	case od.Code == 0:
		return nil, nil, errors.New("the Code field is required")
	case od.Schema.Type == "":
		return nil, nil, errors.New("the Schema.Type field is required")
	}

	reqPath := fmt.Sprintf("dhcp/optiondef/%s/%s", odSpace, odKey)
	req, err := s.client.NewRequest(http.MethodPut, reqPath, od)
	if err != nil {
		return nil, nil, err
	}

	respOd := new(dhcp.OptionDef)
	var resp *http.Response
	resp, err = s.client.Do(req, respOd)
	if err != nil {
		return nil, resp, err
	}

	return respOd, resp, nil
}

// Delete removes a option definition entirely.
//
// NS1 API docs: https://ns1.com/api#deletedelete-a-custom-dhcp-option-definition
func (s *OptionDefService) Delete(odSpace, odKey string) (*http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/optiondef/%s/%s", odSpace, odKey)
	req, err := s.client.NewRequest(http.MethodDelete, reqPath, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
