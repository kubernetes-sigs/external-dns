// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package types

import (
	"encoding/json"
	"strconv"
)

// TestDeepStringer is a TestDeep specific interface for objects which
// know how to stringify themselves.
type TestDeepStringer interface {
	_TestDeep()
	String() string
}

// TestDeepStamp is a useful type providing the _TestDeep() method
// needed to implement TestDeepStringer interface.
type TestDeepStamp struct{}

func (t TestDeepStamp) _TestDeep() {}

// RawString implements TestDeepStringer interface.
type RawString string

func (s RawString) _TestDeep() {}

func (s RawString) String() string {
	return string(s)
}

// RawInt implements TestDeepStringer interface.
type RawInt int

func (i RawInt) _TestDeep() {}

func (i RawInt) String() string {
	return strconv.Itoa(int(i))
}

var _ = []TestDeepStringer{RawString(""), RawInt(0)}

// OperatorNotJSONMarshallableError implements error interface. It
// is returned by (*td.TestDeep).MarshalJSON() to notice the user an
// operator cannot be JSON Marshal'led.
type OperatorNotJSONMarshallableError string

// Error implements error interface.
func (e OperatorNotJSONMarshallableError) Error() string {
	return string(e) + " TestDeep operator cannot be json.Marshal'led"
}

// Operator returns the operator behind this error.
func (e OperatorNotJSONMarshallableError) Operator() string {
	return string(e)
}

// AsOperatorNotJSONMarshallableError checks that err is or contains
// an OperatorNotJSONMarshallableError and if yes, returns it and
// true.
func AsOperatorNotJSONMarshallableError(err error) (OperatorNotJSONMarshallableError, bool) {
	switch err := err.(type) {
	case OperatorNotJSONMarshallableError:
		return err, true

	case *json.MarshalerError:
		if err, ok := err.Err.(OperatorNotJSONMarshallableError); ok {
			return err, true
		}
	}

	return "", false
}
