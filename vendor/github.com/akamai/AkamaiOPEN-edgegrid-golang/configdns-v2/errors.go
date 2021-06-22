package dnsv2

import (
	"fmt"
	client "github.com/akamai/AkamaiOPEN-edgegrid-golang/client-v1"
)

type ConfigDNSError interface {
	error
	Network() bool
	NotFound() bool
	FailedToSave() bool
	ValidationFailed() bool
	ConcurrencyConflict() bool
}

func IsConfigDNSError(e error) bool {
	_, ok := e.(ConfigDNSError)
	return ok
}

type ZoneError struct {
	zoneName         string
	httpErrorMessage string
	apiErrorMessage  string
	err              error
}

func (e *ZoneError) Network() bool {
	if e.httpErrorMessage != "" {
		return true
	}
	return false
}

func (e *ZoneError) NotFound() bool {
	if e.err == nil && e.httpErrorMessage == "" && e.apiErrorMessage == "" {
		return true
	} else if e.err != nil {
		_, ok := e.err.(client.APIError)
		if ok && e.err.(client.APIError).Response.StatusCode == 404 {
			return true
		}
	}
	return false
}

func (e *ZoneError) FailedToSave() bool {
	return false
}

func (e *ZoneError) ValidationFailed() bool {
	if e.apiErrorMessage != "" {
		return true
	}
	return false
}

func (e *ZoneError) ConcurrencyConflict() bool {
	_, ok := e.err.(client.APIError)
	if ok && e.err.(client.APIError).Response.StatusCode == 409 {
		return true
	}
	return false
}

func (e *ZoneError) Error() string {
	if e.Network() {
		return fmt.Sprintf("Zone \"%s\" network error: [%s]", e.zoneName, e.httpErrorMessage)
	}

	if e.NotFound() {
		return fmt.Sprintf("Zone \"%s\" not found.", e.zoneName)
	}

	if e.ConcurrencyConflict() {
		return fmt.Sprintf("Modification Confict: [%s]", e.apiErrorMessage)
	}

	if e.FailedToSave() {
		return fmt.Sprintf("Zone \"%s\" failed to save: [%s]", e.zoneName, e.err.Error())
	}

	if e.ValidationFailed() {
		return fmt.Sprintf("Zone \"%s\" validation failed: [%s]", e.zoneName, e.apiErrorMessage)
	}

	if e.err != nil {
		return e.err.Error()
	}

	return "<nil>"
}

type RecordError struct {
	fieldName        string
	httpErrorMessage string
	apiErrorMessage  string
	err              error
}

func (e *RecordError) Network() bool {
	if e.httpErrorMessage != "" {
		return true
	}
	return false
}

func (e *RecordError) NotFound() bool {
	if e.err == nil && e.httpErrorMessage == "" && e.apiErrorMessage == "" {
		return true
	} else if e.err != nil {
		_, ok := e.err.(client.APIError)
		if ok && e.err.(client.APIError).Response.StatusCode == 404 {
			return true
		}
	}
	return false
}

func (e *RecordError) FailedToSave() bool {
	if e.fieldName == "" {
		return true
	}
	return false
}

func (e *RecordError) ValidationFailed() bool {
	if e.fieldName != "" && e.err == nil {
		return true
	}
	return false
}

func (e *RecordError) ConcurrencyConflict() bool {
	_, ok := e.err.(client.APIError)
	if ok && e.err.(client.APIError).Response.StatusCode == 409 {
		return true
	}
	return false
}

func (e *RecordError) BadRequest() bool {
	_, ok := e.err.(client.APIError)
	if ok && e.err.(client.APIError).Status == 400 {
		return true
	}
	return false
}

func (e *RecordError) Error() string {
	if e.Network() {
		return fmt.Sprintf("Record network error: [%s]", e.httpErrorMessage)
	}

	if e.ConcurrencyConflict() {
		return fmt.Sprintf("Modification Confict: [%s]", e.apiErrorMessage)
	}

	if e.BadRequest() {
		return fmt.Sprintf("Invalid Operation: [%s]", e.apiErrorMessage)
	}

	if e.NotFound() {
		return fmt.Sprintf("Record not found.")
	}

	if e.FailedToSave() {
		return fmt.Sprintf("Record failed to save: [%s]", e.err.Error())
	}

	if e.ValidationFailed() {
		return fmt.Sprintf("Record validation failed for field [%s]", e.fieldName)
	}

	if e.err != nil {
		return fmt.Sprintf("%s", e.err.Error())
	}

	return "<nil>"
}

type TsigError struct {
	keyName          string
	httpErrorMessage string
	apiErrorMessage  string
	err              error
}

func (e *TsigError) Network() bool {
	if e.httpErrorMessage != "" {
		return true
	}
	return false
}

func (e *TsigError) NotFound() bool {
	if e.err == nil && e.httpErrorMessage == "" && e.apiErrorMessage == "" {
		return true
	}
	return false
}

func (e *TsigError) FailedToSave() bool {
	return false
}

func (e *TsigError) ValidationFailed() bool {
	if e.apiErrorMessage != "" {
		return true
	}
	return false
}

func (e *TsigError) Error() string {
	if e.Network() {
		return fmt.Sprintf("Tsig network error: [%s]", e.httpErrorMessage)
	}

	if e.NotFound() {
		return fmt.Sprintf("tsig key not found.")
	}

	if e.FailedToSave() {
		return fmt.Sprintf("tsig key failed to save: [%s]", e.err.Error())
	}

	if e.ValidationFailed() {
		return fmt.Sprintf("tsig key validation failed: [%s]", e.apiErrorMessage)
	}

	return "<nil>"
}
