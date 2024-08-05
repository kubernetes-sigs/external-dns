package rest

import (
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/account"
)

// ActivityService handles 'account/activity' endpoint.
type ActivityService service

// List returns all activity in the account. It accepts a variadic number of
// optional URL parameters that can be used to edit the endpoint's behavior.
// Parameters are in the form of a `rest.Param` struct. The full list of valid
// parameters for this endpoint are available in the documentation.
//
// NS1 API docs: https://developer.ibm.com/apis/catalog/ns1--ibm-ns1-connect-api/api/API--ns1--ibm-ns1-connect-api#getActivity
func (s *ActivityService) List(params ...Param) ([]*account.Activity, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "account/activity", nil)
	if err != nil {
		return nil, nil, err
	}

	al := []*account.Activity{}
	resp, err := s.client.Do(req, &al, params...)
	if err != nil {
		return nil, resp, err
	}

	return al, resp, nil
}
