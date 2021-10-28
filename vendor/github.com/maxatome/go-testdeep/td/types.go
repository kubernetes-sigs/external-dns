// Copyright (c) 2018-2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"
	"strings"
	"testing"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/location"
	"github.com/maxatome/go-testdeep/internal/types"
)

var (
	testDeeper         = reflect.TypeOf((*TestDeep)(nil)).Elem()
	smuggledGotType    = reflect.TypeOf(SmuggledGot{})
	smuggledGotPtrType = reflect.TypeOf((*SmuggledGot)(nil))
)

// TestingT is the minimal interface used by Cmp to report errors. It
// is commonly implemented by *testing.T and *testing.B.
type TestingT interface {
	Error(args ...interface{})
	Fatal(args ...interface{})
	Helper()
}

// TestingFT is a deprecated alias of testing.TB. Use testing.TB
// directly in new code.
type TestingFT = testing.TB

// TestDeep is the representation of a go-testdeep operator. It is not
// intended to be used directly, but through Cmp* functions.
type TestDeep interface {
	types.TestDeepStringer
	location.GetLocationer
	// Match checks "got" against the operator. It returns nil if it matches.
	Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error
	setLocation(int)
	replaceLocation(location.Location)
	// HandleInvalid returns true if the operator is able to handle
	// untyped nil value. Otherwise the untyped nil value is handled
	// generically.
	HandleInvalid() bool
	// TypeBehind returns the type handled by the operator or nil if it
	// is not known. tdhttp helper uses it to know how to unmarshal HTTP
	// responses bodies before comparing them using the operator.
	TypeBehind() reflect.Type
	// Error returns nil if the operator is operational, the
	// corresponding error otherwise.
	Error() error
}

// base is a base type providing some methods needed by the TestDeep
// interface.
type base struct {
	types.TestDeepStamp
	location location.Location
	err      *ctxerr.Error
}

func pkgFunc(full string) (string, string) {
	// the/package.Foo         → "the/package", "Foo"
	// the/package.(*T).Foo    → "the/package", "(*T).Foo"
	// the/package.glob..func1 → "the/package", "glob..func1"
	sp := strings.LastIndexByte(full, '/')
	if sp < 0 {
		sp = 0 // std package without any '/' in name
	}

	dp := strings.IndexByte(full[sp:], '.')
	if dp < 0 {
		return full, ""
	}
	dp += sp
	return full[:dp], full[dp+1:]
}

// setLocation sets location using the stack trace going "callDepth" levels up.
func (t *base) setLocation(callDepth int) {
	var ok bool
	t.location, ok = location.New(callDepth)
	if !ok {
		t.location.File = "???"
		t.location.Line = 0
		return
	}

	// Here package is github.com/maxatome/go-testdeep, or its vendored
	// counterpart
	var pkg string
	pkg, t.location.Func = pkgFunc(t.location.Func)

	// Try to go one level upper, if we are still in go-testdeep package
	cmpLoc, ok := location.New(callDepth + 1)
	if ok {
		cmpPkg, _ := pkgFunc(cmpLoc.Func)
		if cmpPkg == pkg {
			t.location.File = cmpLoc.File
			t.location.Line = cmpLoc.Line
			t.location.BehindCmp = true
		}
	}
}

// replaceLocation replaces the location by "loc".
func (t *base) replaceLocation(loc location.Location) {
	t.location = loc
}

// GetLocation returns a copy of the location.Location where the TestDeep
// operator has been created.
func (t *base) GetLocation() location.Location {
	return t.location
}

// HandleInvalid tells go-testdeep internals that this operator does
// not handle nil values directly.
func (t base) HandleInvalid() bool {
	return false
}

// TypeBehind returns the type handled by the operator. Only few operators
// knows the type they are handling. If they do not know, nil is
// returned.
func (t base) TypeBehind() reflect.Type {
	return nil
}

// Error returns nil if the operator is operational, the corresponding
// error otherwise.
func (t base) Error() error {
	if t.err == nil {
		return nil
	}
	return t.err
}

// stringError is a convenience method to call in String()
// implementations when the operator is in error.
func (t base) stringError() string {
	return t.GetLocation().Func + "(<ERROR>)"
}

// MarshalJSON implements encoding/json.Marshaler only to returns an
// error, as a TestDeep operator should never be JSON marshaled. So
// it is better to tell the user he/she does a mistake.
func (t base) MarshalJSON() ([]byte, error) {
	return nil, types.OperatorNotJSONMarshallableError(t.location.Func)
}

// newBase returns a new base struct with location.Location set to the
// "callDepth" depth.
func newBase(callDepth int) (b base) {
	b.setLocation(callDepth)
	return
}

// baseOKNil is a base type providing some methods needed by the TestDeep
// interface, for operators handling nil values.
type baseOKNil struct {
	base
}

// HandleInvalid tells go-testdeep internals that this operator
// handles nil values directly.
func (t baseOKNil) HandleInvalid() bool {
	return true
}

// newBaseOKNil returns a new baseOKNil struct with location.Location set to
// the "callDepth" depth.
func newBaseOKNil(callDepth int) (b baseOKNil) {
	b.setLocation(callDepth)
	return
}
