package cloudflare

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

var errMissingAccessSeatUID = errors.New("missing required access seat UID")

// AccessUpdateAccessUserSeatResult represents a Access User Seat.
type AccessUpdateAccessUserSeatResult struct {
	AccessSeat  *bool      `json:"access_seat"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	GatewaySeat *bool      `json:"gateway_seat"`
	SeatUID     string     `json:"seat_uid,omitempty"`
}

// UpdateAccessUserSeatParams represents the update payload for access seats.
type UpdateAccessUserSeatParams struct {
	SeatUID     string `json:"seat_uid,omitempty"`
	AccessSeat  *bool  `json:"access_seat"`
	GatewaySeat *bool  `json:"gateway_seat"`
}

// UpdateAccessUsersSeatsParams represents the update payload for multiple access seats.
type UpdateAccessUsersSeatsParams []struct {
	SeatUID     string `json:"seat_uid,omitempty"`
	AccessSeat  *bool  `json:"access_seat"`
	GatewaySeat *bool  `json:"gateway_seat"`
}

// AccessUserSeatResponse represents the response from the access user seat endpoints.
type UpdateAccessUserSeatResponse struct {
	Response
	Result     []AccessUpdateAccessUserSeatResult `json:"result"`
	ResultInfo `json:"result_info"`
}

// UpdateAccessUserSeat updates a single Access User Seat.
//
// API documentation: https://developers.cloudflare.com/api/operations/zero-trust-seats-update-a-user-seat
func (api *API) UpdateAccessUserSeat(ctx context.Context, rc *ResourceContainer, params UpdateAccessUserSeatParams) ([]AccessUpdateAccessUserSeatResult, error) {
	if rc.Level != AccountRouteLevel {
		return []AccessUpdateAccessUserSeatResult{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	if params.SeatUID == "" {
		return []AccessUpdateAccessUserSeatResult{}, errMissingAccessSeatUID
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/seats",
		rc.Level,
		rc.Identifier,
	)

	// this requests expects an array of params, but this method only accepts a single param
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, []UpdateAccessUserSeatParams{params})
	if err != nil {
		return []AccessUpdateAccessUserSeatResult{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var updateAccessUserSeatResponse UpdateAccessUserSeatResponse
	err = json.Unmarshal(res, &updateAccessUserSeatResponse)
	if err != nil {
		return []AccessUpdateAccessUserSeatResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return updateAccessUserSeatResponse.Result, nil
}

// UpdateAccessUsersSeats updates many Access User Seats.
//
// API documentation: https://developers.cloudflare.com/api/operations/zero-trust-seats-update-a-user-seat
func (api *API) UpdateAccessUsersSeats(ctx context.Context, rc *ResourceContainer, params UpdateAccessUsersSeatsParams) ([]AccessUpdateAccessUserSeatResult, error) {
	if rc.Level != AccountRouteLevel {
		return []AccessUpdateAccessUserSeatResult{}, fmt.Errorf(errInvalidResourceContainerAccess, rc.Level)
	}

	for _, param := range params {
		if param.SeatUID == "" {
			return []AccessUpdateAccessUserSeatResult{}, errMissingAccessSeatUID
		}
	}

	uri := fmt.Sprintf(
		"/%s/%s/access/seats",
		rc.Level,
		rc.Identifier,
	)

	// this requests expects an array of params, but this method only accepts a single param
	res, err := api.makeRequestContext(ctx, http.MethodPatch, uri, params)
	if err != nil {
		return []AccessUpdateAccessUserSeatResult{}, fmt.Errorf("%s: %w", errMakeRequestError, err)
	}

	var updateAccessUserSeatResponse UpdateAccessUserSeatResponse
	err = json.Unmarshal(res, &updateAccessUserSeatResponse)
	if err != nil {
		return []AccessUpdateAccessUserSeatResult{}, fmt.Errorf("%s: %w", errUnmarshalError, err)
	}

	return updateAccessUserSeatResponse.Result, nil
}
