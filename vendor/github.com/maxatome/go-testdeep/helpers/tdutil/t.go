// Copyright (c) 2019, 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

// Package tdutil allows to write unit tests for go-testdeep helpers
// and so provides some helpful functions.
//
// It is not intended to be used in tests outside go-testdeep and its
// helpers perimeter.
package tdutil

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/maxatome/go-testdeep/internal/or"
)

var testingT = reflect.TypeOf(testing.T{})

// T can be used in tests, to test [testing.T] behavior as it overrides
// [testing.T.Run] method.
type T struct {
	testing.T
	name string
}

type tFailedNow struct{}

// NewT returns a new [*T] instance. name is the string returned by
// method Name.
func NewT(name string) *T {
	t := &T{name: name}
	t.prepareForLogs()
	return t
}

// Run is a simplified version of [testing.T.Run] method, without edge
// cases.
func (t *T) Run(name string, f func(*testing.T)) bool {
	t.CatchFailNow(func() { f(&t.T) })
	return !t.Failed()
}

// Name returns the name of the running test (in fact the one set by [NewT]).
func (t *T) Name() string {
	return t.name
}

// LogBuf is an ugly hack allowing to access internal [testing.T] log
// buffer. Keep cool, it is only used for internal unit tests.
func (t *T) LogBuf() string {
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) //nolint: govet
}

// FailNow simulates the original [testing.T.FailNow] using
// panic. [T.CatchFailNow] should be used to properly intercept it.
func (t *T) FailNow() {
	t.Fail()
	panic(tFailedNow{})
}

// Fatal simulates the original [testing.T.Fatal].
func (t *T) Fatal(args ...any) {
	t.Helper()
	t.Error(args...)
	t.FailNow()
}

// Fatal simulates the original [testing.T.Fatalf].
func (t *T) Fatalf(format string, args ...any) {
	t.Helper()
	t.Errorf(format, args...)
	t.FailNow()
}

// Starting go1.18, unsafe.Add can be used instead.
func ptrAdd(p unsafe.Pointer, x uintptr) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) + x)
}

// CatchFailNow returns true if a [T.FailNow], [T.Fatal] or [T.Fatalf] call
// occurred during the execution of fn.
func (t *T) CatchFailNow(fn func()) (failNowOccurred bool) {
	defer func() {
		if x := recover(); x != nil {
			_, failNowOccurred = x.(tFailedNow)
			if !failNowOccurred {
				panic(x) // rethrow
			}
		}
	}()

	fn()
	return
}

func (t *T) prepareForLogs() {
	_, ok := testingT.FieldByName("output")
	or.Panic(ok, "testing.T.output field not found!")

	// o is a new go1.25 testing.common field
	if o, ok := testingT.FieldByName("o"); ok {
		ot := o.Type.Elem()

		// With go1.25 t.T.o needs to be set to feed t.T.output
		type outputWriter struct {
			c       *testing.T
			partial []byte //nolint: unused
		}
		owt := reflect.TypeOf(outputWriter{})

		or.Panic(
			ot.NumField() == owt.NumField() &&
				ot.Field(0).Name == owt.Field(0).Name && // ot.c is *testing.common
				ot.Field(1).Type == owt.Field(1).Type,
			"testing.T.outputWriter changed")

		// Use ptrAdd to support go1.16 & go1.17, but starting go1.18 we can use:
		// oAddr := (**outputWriter)(unsafe.Add(unsafe.Pointer(&t.T), o.Offset))
		oAddr := (**outputWriter)(ptrAdd(unsafe.Pointer(&t.T), o.Offset))
		*oAddr = &outputWriter{
			c: &t.T,
		}
	}
}
