package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dns"
)

// TsigService handles 'tsig' endpoint.
type TsigService service

// List returns all tsig keys and basic tsig keys configuration details for each.
//
// NS1 API docs: https://ns1.com/api/#getlist-tsig-keys
func (s *TsigService) List() ([]*dns.TSIGKey, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "tsig", nil)
	if err != nil {
		return nil, nil, err
	}

	tsigKeyList := []*dns.TSIGKey{}
	var resp *http.Response
	resp, err = s.client.Do(req, &tsigKeyList)
	if err != nil {
		return nil, resp, err
	}

	return tsigKeyList, resp, nil
}

// Get takes a TSIG key name and returns a single TSIG key and its basic configuration details.
//
// NS1 API docs: https://ns1.com/api/#getview-tsig-key-details
func (s *TsigService) Get(name string) (*dns.TSIGKey, *http.Response, error) {
	path := fmt.Sprintf("tsig/%s", name)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var tk dns.TSIGKey
	var resp *http.Response
	resp, err = s.client.Do(req, &tk)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusNotFound {
				return nil, resp, ErrTsigKeyMissing
			}
		}
		return nil, resp, err
	}

	return &tk, resp, nil
}

// Create takes a *TSIGkey and creates a new TSIG key.
//
// NS1 API docs: https://ns1.com/api/#putcreate-a-tsig-key
func (s *TsigService) Create(tk *dns.TSIGKey) (*http.Response, error) {
	path := fmt.Sprintf("tsig/%s", tk.Name)

	req, err := s.client.NewRequest("PUT", path, &tk)
	if err != nil {
		return nil, err
	}

	// Update TSIG key fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &tk)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusConflict {
				return resp, ErrTsigKeyExists
			}
		}
		return resp, err
	}

	return resp, nil
}

// Update takes a *TSIGKey and modifies basic details of a TSIG key.
//
// NS1 API docs: https://ns1.com/api/#postmodify-a-tsig-key
func (s *TsigService) Update(tk *dns.TSIGKey) (*http.Response, error) {
	path := fmt.Sprintf("tsig/%s", tk.Name)

	req, err := s.client.NewRequest("POST", path, &tk)
	if err != nil {
		return nil, err
	}

	// Update TSIG key fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &tk)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusNotFound {
				return resp, ErrTsigKeyMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

// Delete takes a TSIG key name and destroys an existing TSIG key.
//
// NS1 API docs: https://ns1.com/api/#deleteremove-a-tsig-key
func (s *TsigService) Delete(name string) (*http.Response, error) {
	path := fmt.Sprintf("tsig/%s", name)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch errType := err.(type) {
		case *Error:
			if errType.Resp.StatusCode == http.StatusNotFound {
				return resp, ErrTsigKeyMissing
			}
		}
		return resp, err
	}

	return resp, nil
}

var (
	// ErrTsigKeyExists bundles PUT create error.
	ErrTsigKeyExists = errors.New("TSIG key already exists")
	// ErrTsigKeyMissing bundles GET/POST/DELETE error.
	ErrTsigKeyMissing = errors.New("TSIG key does not exist")
)
