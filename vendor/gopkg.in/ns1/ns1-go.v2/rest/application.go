package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/pulsar"
)

// ApplicationsService handles 'pulsar/apps/' endpoint.
type ApplicationsService service

// List returns all pulsar Applications
//
// NS1 API docs: https://ns1.com/api#get-list-pulsar-applications
func (s *ApplicationsService) List() ([]*pulsar.Application, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "pulsar/apps", nil)
	if err != nil {
		return nil, nil, err
	}

	var al []*pulsar.Application
	resp, err := s.client.Do(req, &al)
	if err != nil {
		return nil, resp, err
	}

	return al, resp, nil
}

// Get takes a application id and returns application struct.
//
// NS1 API docs: https://ns1.com/api#get-list-pulsar-applications
func (s *ApplicationsService) Get(id string) (*pulsar.Application, *http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s", id)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var a pulsar.Application
	resp, err := s.client.Do(req, &a)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Resp.StatusCode == 404 {
				return nil, resp, ErrApplicationMissing
			}
		}
		return nil, resp, err
	}

	return &a, resp, nil
}

// Create takes a *pulsar.Application and creates a new Application.
//
// The given application must have at least the name
// NS1 API docs: https://ns1.com/api#put-create-a-pulsar-application
func (s *ApplicationsService) Create(a *pulsar.Application) (*http.Response, error) {
	req, err := s.client.NewRequest("PUT", "pulsar/apps", a)
	if err != nil {
		return nil, err
	}
	resp, err := s.client.Do(req, a)
	return resp, err
}

// Update takes a *pulsar.Application and updates the application with same id on Ns1.
//
// NS1 API docs: https://ns1.com/api#post-modify-an-application
func (s *ApplicationsService) Update(a *pulsar.Application) (*http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s", a.ID)

	req, err := s.client.NewRequest("POST", path, &a)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, &a)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Resp.StatusCode == 404 {
				return resp, ErrApplicationMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete takes a application Id, and removes an existing application
//
// NS1 API docs: https://ns1.com/api#delete-delete-a-pulsar-application
func (s *ApplicationsService) Delete(id string) (*http.Response, error) {
	path := fmt.Sprintf("pulsar/apps/%s", id)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Resp.StatusCode == 404 {
				return resp, ErrApplicationMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

var (
	// ErrApplicationMissing bundles GET/POST/DELETE error.
	ErrApplicationMissing = errors.New("application does not exist")
)
