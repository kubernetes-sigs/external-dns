package rest

import (
	"encoding/json"
	"fmt"
	"time"
)

// Error is used to unpack every error returned by the api
type Error struct {
	// Message contains the error from the api as a string
	Message string `json:"error"`
	// StatusCode contains a HTTP status code that the api server responded with
	StatusCode int
}

func (e *Error) Error() string {
	return e.Message
}

// Response will contain a body (which can be empty), status code and the Method.
// This struct will be used to decode a response from the api server.
type Response struct {
	Body       []byte
	StatusCode int
	Method     Method
}

// Time is defined because the transip api server does not return a rfc 3339 time string
// and golang requires this. So we need to do manual time parsing, by defining our own time struct
// encapsulating time.Time.
type Time struct {
	// Time item containing the actual parsed time object
	time.Time
}

// Date is defined because the transip api server returns date strings, not parsed by golang by default.
// So we need to do manual time parsing, by defining our own date struct encapsulating time.Time.
type Date struct {
	// Time item containing the actual parsed time object
	time.Time
}

// UnmarshalJSON parses datetime strings returned by the transip api
func (tt *Time) UnmarshalJSON(input []byte) error {
	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		return err
	}
	// don't parse on empty dates
	if string(input) == `""` {
		return nil
	}
	newTime, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(input), loc)
	if err != nil {
		return err
	}

	tt.Time = newTime
	return nil
}

// UnmarshalJSON parses date strings returned by the transip api
func (td *Date) UnmarshalJSON(input []byte) error {
	loc, err := time.LoadLocation("Europe/Amsterdam")
	if err != nil {
		return err
	}
	// don't parse on empty dates
	if string(input) == `""` {
		return nil
	}
	newTime, err := time.ParseInLocation(`"2006-01-02"`, string(input), loc)
	if err != nil {
		return err
	}

	td.Time = newTime
	return nil
}

// ParseResponse will convert a Response struct to the given interface.
// When the rest response has no body it will return without filling the dest variable.
func (r *Response) ParseResponse(dest interface{}) error {
	// do response error checking
	if !r.Method.StatusCodeOK(r.StatusCode) {
		return r.parseErrorResponse()
	}

	if len(r.Body) == 0 {
		return nil
	}

	return json.Unmarshal(r.Body, &dest)
}

// parseErrorResponse tries to unmarshal the error response body
// so we can return it to the user
func (r *Response) parseErrorResponse() error {
	// there is no response content so we also don't need to parse it
	if len(r.Body) == 0 {
		return &Error{
			Message:    fmt.Sprintf("error response without body from api server status code '%d'", r.StatusCode),
			StatusCode: r.StatusCode,
		}
	}

	var errorResponse Error
	err := json.Unmarshal(r.Body, &errorResponse)
	if err != nil {
		return &Error{
			Message:    fmt.Sprintf("response error could not be decoded '%s'", string(r.Body)),
			StatusCode: r.StatusCode,
		}
	}

	// set the exposed status code so users can check on this
	errorResponse.StatusCode = r.StatusCode

	return &errorResponse
}
