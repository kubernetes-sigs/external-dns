package rest

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/ns1/ns1-go.v2/rest/model/redirect"
)

// RedirectService handles 'redirect' endpoint.
type RedirectService service

// List returns the configured redirects.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectService) List() ([]*redirect.Configuration, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "redirect", nil)
	if err != nil {
		return nil, nil, err
	}

	cfgList := redirect.ConfigurationList{}
	var resp *http.Response
	if s.client.FollowPagination {
		resp, err = s.client.DoWithPagination(req, &cfgList, s.nextCfgs)
	} else {
		resp, err = s.client.Do(req, &cfgList)
	}
	if err != nil {
		return nil, resp, err
	}

	return cfgList.Results, resp, nil
}

// Get takes a redirect config id and returns a single config.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectService) Get(cfgId string) (*redirect.Configuration, *http.Response, error) {
	path := fmt.Sprintf("redirect/%s", cfgId)

	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var cfg redirect.Configuration
	var resp *http.Response
	resp, err = s.client.Do(req, &cfg)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if strings.HasSuffix(err.Message, " not found") {
				return nil, resp, ErrRedirectNotFound
			}
		}
		return nil, resp, err
	}

	return &cfg, resp, nil
}

// Create takes a *Configuration and creates a new redirect.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectService) Create(cfg *redirect.Configuration) (*redirect.Configuration, *http.Response, error) {
	if cfg == nil {
		return nil, nil, ErrRedirectNil
	}

	req, err := s.client.NewRequest("PUT", "redirect", &cfg)
	if err != nil {
		return nil, nil, err
	}

	// Update redirect fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &cfg)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if err.Message == "configuration already exists" {
				return nil, resp, ErrRedirectExists
			}
		}
		return nil, resp, err
	}

	return cfg, resp, nil
}

// Update takes a *Configuration and modifies basic details of a redirect.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectService) Update(cfg *redirect.Configuration) (*redirect.Configuration, *http.Response, error) {
	if cfg == nil || cfg.ID == nil {
		return nil, nil, ErrRedirectNil
	}

	path := fmt.Sprintf("redirect/%s", *cfg.ID)

	req, err := s.client.NewRequest("POST", path, &cfg)
	if err != nil {
		return nil, nil, err
	}

	// Update redirect fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &cfg)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if strings.HasSuffix(err.Message, " not found") {
				return nil, resp, ErrRedirectNotFound
			}
		}
		return nil, resp, err
	}

	return cfg, resp, nil
}

// Delete takes a configuration id and destroys the associated redirect configuration.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/Getting+Started
// Feature docs: https://www.ibm.com/docs/en/ns1-connect?topic=url-redirects
func (s *RedirectService) Delete(cfgId string) (*http.Response, error) {
	path := fmt.Sprintf("redirect/%s", cfgId)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		switch err := err.(type) {
		case *Error:
			if strings.HasSuffix(err.Message, " not found") {
				return resp, ErrRedirectNotFound
			}
		}
		return resp, err
	}

	return resp, nil
}

// nextCfgs is a pagination helper than gets and appends another list of redirect configs
// to the passed list.
func (s *RedirectService) nextCfgs(v *interface{}, uri string) (*http.Response, error) {
	tmpCfgList := redirect.ConfigurationList{}
	resp, err := s.client.getURI(&tmpCfgList, uri)
	if err != nil {
		return resp, err
	}
	cfgList, ok := (*v).(*redirect.ConfigurationList)
	if !ok {
		return nil, fmt.Errorf(
			"incorrect value for v, expected value of type *redirect.ConfigurationList, got: %T", v,
		)
	}
	cfgList.Total = tmpCfgList.Total
	cfgList.Count += tmpCfgList.Count
	cfgList.Results = append(cfgList.Results, tmpCfgList.Results...)
	return resp, nil
}

var (
	ErrRedirectNil = errors.New("parameter missing")
	// ErrRedirectExists bundles PUT create error.
	ErrRedirectExists = errors.New("redirect configuration id already exists")
	// ErrRedirectNotFound bundles GET/POST/DELETE error.
	ErrRedirectNotFound = errors.New("redirect configuration id not found")
)
