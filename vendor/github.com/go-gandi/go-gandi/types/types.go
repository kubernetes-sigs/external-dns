package types

import (
	"fmt"
)

// StandardResponse is a standard response
type StandardResponse struct {
	Code    int             `json:"code,omitempty"`
	Message string          `json:"message,omitempty"`
	UUID    string          `json:"uuid,omitempty"`
	Object  string          `json:"object,omitempty"`
	Cause   string          `json:"cause,omitempty"`
	Status  string          `json:"status,omitempty"`
	Errors  []StandardError `json:"errors,omitempty"`
}

// StandardError is embedded in a standard error
type StandardError struct {
	Location    string `json:"location"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RequestError struct {
	Err        error
	StatusCode int
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("StatusCode: %d ; Err: %s ", e.StatusCode, e.Err)
}
