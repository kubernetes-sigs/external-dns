package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/alerting"
)

// AlertsService handles 'alerting/v1/alerts' endpoint.
type AlertsService service

// The base for the alerting api relative to /v1
// client.NewRequest will call ResolveReference and remove /v1/../
const alertingRelativeBase = "../alerting/v1"

type alertListResponse struct {
	Limit        *int64            `json:"limit,omitempty"`
	Next         *string           `json:"next,omitempty"`
	Results      []*alerting.Alert `json:"results"`
	TotalResults *int64            `json:"total_results,omitempty"`
}

// List returns all configured alerts.
//
// NS1 API docs: https://ns1.com/api/#alerts-get
func (s *AlertsService) List() ([]*alerting.Alert, *http.Response, error) {
	path := fmt.Sprintf("%s/%s", alertingRelativeBase, "alerts")
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	alertListResp := alertListResponse{}
	var resp *http.Response
	if s.client.FollowPagination {
		resp, err = s.client.DoWithPagination(req, &alertListResp, s.nextAlerts)
	} else {
		resp, err = s.client.Do(req, &alertListResp)
	}
	if err != nil {
		return nil, resp, err
	}
	alerts := alertListResp.Results
	return alerts, resp, nil
}

// nextAlerts is a pagination helper than gets and appends another list of alerts
// to the passed alerts.
func (s *AlertsService) nextAlerts(v *interface{}, uri string) (*http.Response, error) {
	nextAlerts := &alertListResponse{}
	resp, err := s.client.getURI(&nextAlerts, uri)
	if err != nil {
		return resp, err
	}
	alertListResp, ok := (*v).(*alertListResponse)
	if !ok {
		return nil, fmt.Errorf(
			"incorrect value for v, expected value of type *[]*alerting.Alert, got: %T", v,
		)
	}
	alertListResp.Results = append(alertListResp.Results, nextAlerts.Results...)
	return resp, nil
}

// Get returns the details of a specific alert.
//
// NS1 API docs: https://ns1.com/api/#alert-alertid-get
func (s *AlertsService) Get(alertID string) (*alerting.Alert, *http.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", alertingRelativeBase, "alerts", alertID)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var alert alerting.Alert
	resp, err := s.client.Do(req, &alert)
	if err != nil {
		switch err.(type) {
		case *Error:
			if resourceMissingMatch(err.(*Error).Message) {
				return nil, resp, ErrAlertMissing
			}
		}
		return nil, resp, err
	}

	return &alert, resp, nil
}

// Create takes a *alerting.Alert and creates a new alert.
//
// NS1 API docs: https://ns1.com/api/#alert-post
func (s *AlertsService) Create(alert *alerting.Alert) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s", alertingRelativeBase, "alerts")
	req, err := s.client.NewRequest("POST", path, &alert)
	if err != nil {
		return nil, err
	}

	// Update the alerts fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &alert)
	if err != nil {
		switch err.(type) {
		case *Error:
			if err.(*Error).Resp.StatusCode == http.StatusConflict {
				return resp, ErrAlertExists
			}
		}
		return resp, err
	}

	return resp, nil
}

// Update updates the fields specified in the passed alert object.
//
// NS1 API docs: https://ns1.com/api/#alert-alertid-patch
func (s *AlertsService) Update(alert *alerting.Alert) (*http.Response, error) {
	alertID := ""
	if alert != nil && alert.ID != nil {
		alertID = *alert.ID
	}
	path := fmt.Sprintf("%s/%s/%s", alertingRelativeBase, "alerts", alertID)

	req, err := s.client.NewRequest("PATCH", path, &alert)
	if err != nil {
		return nil, err
	}

	// Update the alerts fields with data from api(ensure consistent)
	resp, err := s.client.Do(req, &alert)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Replace replaces the values in an alert with the values in the passed object.
//
// NS1 API docs: https://ns1.com/api/#alert-alertid-put
func (s *AlertsService) Replace(alert *alerting.Alert) (*http.Response, error) {
	alertID := ""
	if alert != nil && alert.ID != nil {
		alertID = *alert.ID
	}
	path := fmt.Sprintf("%s/%s/%s", alertingRelativeBase, "alerts", alertID)

	req, err := s.client.NewRequest("PUT", path, &alert)
	if err != nil {
		return nil, err
	}

	// Update the alerts fields with data from api (ensure consistent)
	resp, err := s.client.Do(req, &alert)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Delete deletes an existing alert.
//
// NS1 API docs: https://ns1.com/api/#alert-alertid-delete
func (s *AlertsService) Delete(alertID string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s/%s", alertingRelativeBase, "alerts", alertID)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Test an existing alert, triggers notifications for the given alert id.
//
// NS1 API docs: https://ns1.com/api/#alert-alertid-test
func (s *AlertsService) Test(alertID string) (*http.Response, error) {
	path := fmt.Sprintf("%s/%s/%s/test", alertingRelativeBase, "alerts", alertID)
	req, err := s.client.NewRequest("POST", path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

var (
	// ErrAlertExists bundles POST create error.
	ErrAlertExists = errors.New("alert already exists")

	// ErrAlertMissing bundles GET/PUT/PATCH/DELETE error.
	ErrAlertMissing = errors.New("alert does not exist")
)
