// Copyright 2013-2015 Apcera Inc. All rights reserved.

// GSS status and errors

package gssapi

/*
#include <gssapi/gssapi.h>

OM_uint32
wrap_gss_display_status(void *fp,
	OM_uint32 *minor_status,
	OM_uint32 status_value,
	int status_type,
	const gss_OID mech_type,
	OM_uint32 *message_context,
	gss_buffer_t status_string)
{
	return ((OM_uint32(*)(
		OM_uint32 *,
		OM_uint32,
		int,
		const gss_OID,
		OM_uint32 *,
		gss_buffer_t)
		)fp)(minor_status,
			status_value,
			status_type,
			mech_type,
			message_context,
			status_string);
}

*/
import "C"

import (
	"errors"
	"fmt"
	"strings"
)

// Constant values are specified for C-language bindings in RFC 2744.
/*
"""
   These errors are encoded into the 32-bit GSS status code as follows:

      MSB                                                        LSB
      |------------------------------------------------------------|
      |  Calling Error | Routine Error  |    Supplementary Info    |
      |------------------------------------------------------------|
   Bit 31            24 23            16 15                       0
"""

Note that the first two fields hold integer consts, whereas Supplementary Info
is a bit-field.
*/

const (
	shiftCALLING = 24
	shiftROUTINE = 16
	maskCALLING  = 0xFF000000
	maskROUTINE  = 0x00FF0000
	maskSUPPINFO = 0x0000FFFF
)

// Status values are returned by gssapi calls to indicate the result of a call.
// Declared according to: https://tools.ietf.org/html/rfc2743#page-17
const (
	GSS_S_COMPLETE MajorStatus = 0

	GSS_S_CALL_INACCESSIBLE_READ  MajorStatus = 1 << shiftCALLING
	GSS_S_CALL_INACCESSIBLE_WRITE             = 2 << shiftCALLING
	GSS_S_CALL_BAD_STRUCTURE                  = 3 << shiftCALLING

	GSS_S_BAD_MECH             MajorStatus = 1 << shiftROUTINE
	GSS_S_BAD_NAME                         = 2 << shiftROUTINE
	GSS_S_BAD_NAMETYPE                     = 3 << shiftROUTINE
	GSS_S_BAD_BINDINGS                     = 4 << shiftROUTINE
	GSS_S_BAD_STATUS                       = 5 << shiftROUTINE
	GSS_S_BAD_MIC                          = 6 << shiftROUTINE
	GSS_S_BAD_SIG                          = 6 << shiftROUTINE // duplication deliberate
	GSS_S_NO_CRED                          = 7 << shiftROUTINE
	GSS_S_NO_CONTEXT                       = 8 << shiftROUTINE
	GSS_S_DEFECTIVE_TOKEN                  = 9 << shiftROUTINE
	GSS_S_DEFECTIVE_CREDENTIAL             = 10 << shiftROUTINE
	GSS_S_CREDENTIALS_EXPIRED              = 11 << shiftROUTINE
	GSS_S_CONTEXT_EXPIRED                  = 12 << shiftROUTINE
	GSS_S_FAILURE                          = 13 << shiftROUTINE
	GSS_S_BAD_QOP                          = 14 << shiftROUTINE
	GSS_S_UNAUTHORIZED                     = 15 << shiftROUTINE
	GSS_S_UNAVAILABLE                      = 16 << shiftROUTINE
	GSS_S_DUPLICATE_ELEMENT                = 17 << shiftROUTINE
	GSS_S_NAME_NOT_MN                      = 18 << shiftROUTINE

	field_GSS_S_CONTINUE_NEEDED = 1 << 0
	field_GSS_S_DUPLICATE_TOKEN = 1 << 1
	field_GSS_S_OLD_TOKEN       = 1 << 2
	field_GSS_S_UNSEQ_TOKEN     = 1 << 3
	field_GSS_S_GAP_TOKEN       = 1 << 4
)

// These are GSSAPI-defined:
// TODO: should MajorStatus be defined as C.OM_uint32?
type MajorStatus uint32

// CallingError is equivalent to C GSS_CALLING_ERROR() macro.
func (st MajorStatus) CallingError() MajorStatus {
	return st & maskCALLING
}

// RoutineError is equivalent to C GSS_ROUTINE_ERROR() macro.
func (st MajorStatus) RoutineError() MajorStatus {
	return st & maskROUTINE
}

// SupplementaryInfo is equivalent to C GSS_SUPPLEMENTARY_INFO() macro.
func (st MajorStatus) SupplementaryInfo() MajorStatus {
	return st & maskSUPPINFO
}

// IsError is equivalent to C GSS_ERROR() macro. Not written as 'Error' because
// that's special in Go conventions. (i.e. conforming to error interface)
func (st MajorStatus) IsError() bool {
	return st&(maskCALLING|maskROUTINE) != 0
}

// ContinueNeeded is equivalent to a C bitfield set test against the
// GSS_S_CONTINUE_NEEDED macro.
func (st MajorStatus) ContinueNeeded() bool {
	return st&field_GSS_S_CONTINUE_NEEDED != 0
}

// DuplicateToken is equivalent to a C bitfield set test against the
// GSS_S_DUPLICATE_TOKEN macro.
func (st MajorStatus) DuplicateToken() bool {
	return st&field_GSS_S_DUPLICATE_TOKEN != 0
}

// OldToken is equivalent to a C bitfield set test against the
// GSS_S_OLD_TOKEN macro.
func (st MajorStatus) OldToken() bool {
	return st&field_GSS_S_OLD_TOKEN != 0
}

// UnseqToken is equivalent to a C bitfield set test against the
// GSS_S_UNSEQ_TOKEN macro.
func (st MajorStatus) UnseqToken() bool {
	return st&field_GSS_S_UNSEQ_TOKEN != 0
}

// GapToken is equivalent to a C bitfield set test against the
// GSS_S_GAP_TOKEN macro.
func (st MajorStatus) GapToken() bool {
	return st&field_GSS_S_GAP_TOKEN != 0
}

// Error is designed to serve both as an error, and as a general gssapi status
// container. If Major is GSS_S_FAILURE, then information will be in Minor.
// The GoError method will return a nil if it doesn't represent a real error.
type Error struct {
	// gssapi lib binding, so that we can convert the results of an
	// operation to a string for diagnosis.
	*Lib

	// Specified by gssapi
	Major MajorStatus

	// Mechanism-specific:
	Minor C.OM_uint32
}

// MakeError creates a golang Error object from a gssapi major & minor status.
func (lib *Lib) MakeError(major, minor C.OM_uint32) *Error {
	return &Error{
		Lib:   lib,
		Major: MajorStatus(major),
		Minor: minor,
	}
}

// ErrContinueNeeded may be returned by InitSecContext or AcceptSecContext to
// indicate that another iteration is needed
var ErrContinueNeeded = errors.New("continue needed")

func (lib *Lib) stashLastStatus(major, minor C.OM_uint32) error {
	lib.LastStatus = lib.MakeError(major, minor)
	return lib.LastStatus.GoError()
}

// GoError returns an untyped error interface object.
func (e *Error) GoError() error {
	if e.Major.IsError() {
		return e
	}
	return nil
}

// Error returns a string representation of an Error object.
func (e *Error) Error() string {
	messages := []string{}
	nOther := 0
	context := C.OM_uint32(0)
	inquiry := C.OM_uint32(0)
	code_type := 0
	first := true

	if e.Major.RoutineError() == GSS_S_FAILURE {
		inquiry = e.Minor
		code_type = GSS_C_MECH_CODE
	} else {
		inquiry = C.OM_uint32(e.Major)
		code_type = GSS_C_GSS_CODE
	}

	for first || context != C.OM_uint32(0) {
		first = false
		min := C.OM_uint32(0)

		b, err := e.MakeBuffer(allocGSSAPI)
		if err != nil {
			break
		}

		// TODO: store a mech_type at the lib level?  Or context? For now GSS_C_NO_OID...
		maj := C.wrap_gss_display_status(
			e.Fp_gss_display_status,
			&min,
			inquiry,
			C.int(code_type),
			nil,
			&context,
			b.C_gss_buffer_t)

		err = e.MakeError(maj, min).GoError()
		if err != nil {
			nOther = nOther + 1
		}
		messages = append(messages, b.String())
		b.Release()
	}
	if nOther > 0 {
		messages = append(messages, fmt.Sprintf("additionally, %d conversions failed", nOther))
	}
	messages = append(messages, "")
	return strings.Join(messages, "\n")
}
