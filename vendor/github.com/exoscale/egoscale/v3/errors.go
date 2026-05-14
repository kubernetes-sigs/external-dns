package v3

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	ErrBadRequest                    = errors.New(http.StatusText(http.StatusBadRequest))
	ErrUnauthorized                  = errors.New(http.StatusText(http.StatusUnauthorized))
	ErrPaymentRequired               = errors.New(http.StatusText(http.StatusPaymentRequired))
	ErrForbidden                     = errors.New(http.StatusText(http.StatusForbidden))
	ErrNotFound                      = errors.New(http.StatusText(http.StatusNotFound))
	ErrMethodNotAllowed              = errors.New(http.StatusText(http.StatusMethodNotAllowed))
	ErrNotAcceptable                 = errors.New(http.StatusText(http.StatusNotAcceptable))
	ErrProxyAuthRequired             = errors.New(http.StatusText(http.StatusProxyAuthRequired))
	ErrRequestTimeout                = errors.New(http.StatusText(http.StatusRequestTimeout))
	ErrConflict                      = errors.New(http.StatusText(http.StatusConflict))
	ErrGone                          = errors.New(http.StatusText(http.StatusGone))
	ErrLengthRequired                = errors.New(http.StatusText(http.StatusLengthRequired))
	ErrPreconditionFailed            = errors.New(http.StatusText(http.StatusPreconditionFailed))
	ErrRequestEntityTooLarge         = errors.New(http.StatusText(http.StatusRequestEntityTooLarge))
	ErrRequestURITooLong             = errors.New(http.StatusText(http.StatusRequestURITooLong))
	ErrUnsupportedMediaType          = errors.New(http.StatusText(http.StatusUnsupportedMediaType))
	ErrRequestedRangeNotSatisfiable  = errors.New(http.StatusText(http.StatusRequestedRangeNotSatisfiable))
	ErrExpectationFailed             = errors.New(http.StatusText(http.StatusExpectationFailed))
	ErrTeapot                        = errors.New(http.StatusText(http.StatusTeapot))
	ErrMisdirectedRequest            = errors.New(http.StatusText(http.StatusMisdirectedRequest))
	ErrUnprocessableEntity           = errors.New(http.StatusText(http.StatusUnprocessableEntity))
	ErrLocked                        = errors.New(http.StatusText(http.StatusLocked))
	ErrFailedDependency              = errors.New(http.StatusText(http.StatusFailedDependency))
	ErrTooEarly                      = errors.New(http.StatusText(http.StatusTooEarly))
	ErrUpgradeRequired               = errors.New(http.StatusText(http.StatusUpgradeRequired))
	ErrPreconditionRequired          = errors.New(http.StatusText(http.StatusPreconditionRequired))
	ErrTooManyRequests               = errors.New(http.StatusText(http.StatusTooManyRequests))
	ErrRequestHeaderFieldsTooLarge   = errors.New(http.StatusText(http.StatusRequestHeaderFieldsTooLarge))
	ErrUnavailableForLegalReasons    = errors.New(http.StatusText(http.StatusUnavailableForLegalReasons))
	ErrInternalServerError           = errors.New(http.StatusText(http.StatusInternalServerError))
	ErrNotImplemented                = errors.New(http.StatusText(http.StatusNotImplemented))
	ErrBadGateway                    = errors.New(http.StatusText(http.StatusBadGateway))
	ErrServiceUnavailable            = errors.New(http.StatusText(http.StatusServiceUnavailable))
	ErrGatewayTimeout                = errors.New(http.StatusText(http.StatusGatewayTimeout))
	ErrHTTPVersionNotSupported       = errors.New(http.StatusText(http.StatusHTTPVersionNotSupported))
	ErrVariantAlsoNegotiates         = errors.New(http.StatusText(http.StatusVariantAlsoNegotiates))
	ErrInsufficientStorage           = errors.New(http.StatusText(http.StatusInsufficientStorage))
	ErrLoopDetected                  = errors.New(http.StatusText(http.StatusLoopDetected))
	ErrNotExtended                   = errors.New(http.StatusText(http.StatusNotExtended))
	ErrNetworkAuthenticationRequired = errors.New(http.StatusText(http.StatusNetworkAuthenticationRequired))
)

var httpStatusCodeErrors = map[int]error{
	http.StatusBadRequest:                    ErrBadRequest,
	http.StatusUnauthorized:                  ErrUnauthorized,
	http.StatusPaymentRequired:               ErrPaymentRequired,
	http.StatusForbidden:                     ErrForbidden,
	http.StatusNotFound:                      ErrNotFound,
	http.StatusMethodNotAllowed:              ErrMethodNotAllowed,
	http.StatusNotAcceptable:                 ErrNotAcceptable,
	http.StatusProxyAuthRequired:             ErrProxyAuthRequired,
	http.StatusRequestTimeout:                ErrRequestTimeout,
	http.StatusConflict:                      ErrConflict,
	http.StatusGone:                          ErrGone,
	http.StatusLengthRequired:                ErrLengthRequired,
	http.StatusPreconditionFailed:            ErrPreconditionFailed,
	http.StatusRequestEntityTooLarge:         ErrRequestEntityTooLarge,
	http.StatusRequestURITooLong:             ErrRequestURITooLong,
	http.StatusUnsupportedMediaType:          ErrUnsupportedMediaType,
	http.StatusRequestedRangeNotSatisfiable:  ErrRequestedRangeNotSatisfiable,
	http.StatusExpectationFailed:             ErrExpectationFailed,
	http.StatusTeapot:                        ErrTeapot,
	http.StatusMisdirectedRequest:            ErrMisdirectedRequest,
	http.StatusUnprocessableEntity:           ErrUnprocessableEntity,
	http.StatusLocked:                        ErrLocked,
	http.StatusFailedDependency:              ErrFailedDependency,
	http.StatusTooEarly:                      ErrTooEarly,
	http.StatusUpgradeRequired:               ErrUpgradeRequired,
	http.StatusPreconditionRequired:          ErrPreconditionRequired,
	http.StatusTooManyRequests:               ErrTooManyRequests,
	http.StatusRequestHeaderFieldsTooLarge:   ErrRequestHeaderFieldsTooLarge,
	http.StatusUnavailableForLegalReasons:    ErrUnavailableForLegalReasons,
	http.StatusInternalServerError:           ErrInternalServerError,
	http.StatusNotImplemented:                ErrNotImplemented,
	http.StatusBadGateway:                    ErrBadGateway,
	http.StatusServiceUnavailable:            ErrServiceUnavailable,
	http.StatusGatewayTimeout:                ErrGatewayTimeout,
	http.StatusHTTPVersionNotSupported:       ErrHTTPVersionNotSupported,
	http.StatusVariantAlsoNegotiates:         ErrVariantAlsoNegotiates,
	http.StatusInsufficientStorage:           ErrInsufficientStorage,
	http.StatusLoopDetected:                  ErrLoopDetected,
	http.StatusNotExtended:                   ErrNotExtended,
	http.StatusNetworkAuthenticationRequired: ErrNetworkAuthenticationRequired,
}

func handleHTTPErrorResp(resp *http.Response) error {
	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		var res struct {
			Message string `json:"message"`
			Error   string `json:"error"`
		}

		data, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading response body: %s", err)
		}

		if json.Valid(data) {
			if err = json.Unmarshal(data, &res); err != nil {
				return fmt.Errorf("error unmarshaling response: %s", err)
			}
		} else {
			res.Message = string(data)
		}

		message := res.Message
		if message == "" && res.Error != "" {
			message = res.Error
		}

		err, ok := httpStatusCodeErrors[resp.StatusCode]
		if ok {
			return fmt.Errorf("%w: %s", err, message)
		}

		return fmt.Errorf("unmapped HTTP error: status code %d, message: %s", resp.StatusCode, message)
	}

	return nil
}
