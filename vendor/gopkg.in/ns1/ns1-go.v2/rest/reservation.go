package rest

import (
	"errors"
	"fmt"
	"net/http"

	"gopkg.in/ns1/ns1-go.v2/rest/model/dhcp"
)

// ReservationService handles the 'reservation' endpoints.
type ReservationService service

// List returns a list of all reservations.
//
// NS1 API docs: https://ns1.com/api#getlist-reservations
func (s *ReservationService) List() ([]dhcp.Reservation, *http.Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "dhcp/reservation", nil)
	if err != nil {
		return nil, nil, err
	}

	scs := make([]dhcp.Reservation, 0)
	resp, err := s.client.Do(req, &scs)
	if err != nil {
		return nil, resp, err
	}

	return scs, resp, nil
}

// Get returns the reservation corresponding to the provided reservation ID.
//
// NS1 API docs: https://ns1.com/api#getview-a-reservations-details
func (s *ReservationService) Get(scID int) (*dhcp.Reservation, *http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/reservation/%d", scID)
	req, err := s.client.NewRequest(http.MethodGet, reqPath, nil)
	if err != nil {
		return nil, nil, err
	}

	sc := &dhcp.Reservation{}
	var resp *http.Response
	resp, err = s.client.Do(req, sc)
	if err != nil {
		return nil, resp, err
	}

	return sc, resp, nil
}

// Create creates a reservation.
// The Options field is required.
//
// NS1 API docs: https://ns1.com/api#putcreate-a-reservation
func (s *ReservationService) Create(sc *dhcp.Reservation) (*dhcp.Reservation, *http.Response, error) {
	switch {
	case sc.Options == nil:
		return nil, nil, errors.New("the Options field is required")
	}

	req, err := s.client.NewRequest(http.MethodPut, "dhcp/reservation", sc)
	if err != nil {
		return nil, nil, err
	}

	respSc := new(dhcp.Reservation)
	var resp *http.Response
	resp, err = s.client.Do(req, respSc)
	if err != nil {
		return nil, resp, err
	}

	return respSc, resp, nil
}

// Edit updates an existing reservation.
// The ID, Options fields are required.
//
// NS1 API docs: https://ns1.com/api#postmodify-a-reservation
func (s *ReservationService) Edit(sc *dhcp.Reservation) (*dhcp.Reservation, *http.Response, error) {
	switch {
	case sc.ID == nil:
		return nil, nil, errors.New("the ID field is required")
	case sc.Options == nil:
		return nil, nil, errors.New("the Options field is required")
	}

	reqPath := fmt.Sprintf("dhcp/reservation/%d", *sc.ID)
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

// Delete removes a reservation entirely.
//
// NS1 API docs: https://ns1.com/api#deletedelete-a-reservation
func (s *ReservationService) Delete(id int) (*http.Response, error) {
	reqPath := fmt.Sprintf("dhcp/reservation/%d", id)
	req, err := s.client.NewRequest(http.MethodDelete, reqPath, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}
