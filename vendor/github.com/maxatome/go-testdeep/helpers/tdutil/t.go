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
)

// T can be used in tests, to test [testing.T] behavior as it overrides
// [testing.T.Run] method.
type T struct {
	testing.T
	name string
}

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
type tFailedNow struct{}

// NewT returns a new [*T] instance. name is the string returned by
// method Name.
func NewT(name string) *T {
	return &T{name: name}
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
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) // nolint: govet
}

// FailNow simulates the original (*testing.T).FailNow using
// panic. CatchFailNow should be used to properly intercept it.
func (t *T) FailNow() {
	t.Fail()
	panic(tFailedNow{})
}

// Fatal simulates the original (*testing.T).Fatal.
func (t *T) Fatal(args ...interface{}) {
	t.Helper()
	t.Error(args...)
	t.FailNow()
}

// Fatal simulates the original (*testing.T).Fatalf.
func (t *T) Fatalf(format string, args ...interface{}) {
	t.Helper()
	t.Errorf(format, args...)
	t.FailNow()
}

// CatchFailNow returns true if a FailNow, Fatal or Fatalf call
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
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
type tFailedNow struct{}

>>>>>>> 5ce8c7613 (update vendored files)
// NewT returns a new *T instance. "name" is the string returned by
// method Name.
func NewT(name string) *T {
	return &T{name: name}
}

// Run is a simplified version of testing.T.Run() method, without edge
// cases.
func (t *T) Run(name string, f func(*testing.T)) bool {
	t.CatchFailNow(func() { f(&t.T) })
	return !t.Failed()
}

// Name returns the name of the running test (in fact the one set by NewT).
func (t *T) Name() string {
	return t.name
}

// LogBuf is an ugly hack allowing to access internal testing.T log
// buffer. Keep cool, it is only used for internal unit tests.
func (t *T) LogBuf() string {
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) // nolint: govet
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}

// FailNow simulates the original (*testing.T).FailNow using
// panic. CatchFailNow should be used to properly intercept it.
func (t *T) FailNow() {
	t.Fail()
	panic(tFailedNow{})
}

// Fatal simulates the original (*testing.T).Fatal.
func (t *T) Fatal(args ...interface{}) {
	t.Helper()
	t.Error(args...)
	t.FailNow()
}

// Fatal simulates the original (*testing.T).Fatalf.
func (t *T) Fatalf(format string, args ...interface{}) {
	t.Helper()
	t.Errorf(format, args...)
	t.FailNow()
}

// CatchFailNow returns true if a FailNow, Fatal or Fatalf call
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
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
type tFailedNow struct{}

>>>>>>> 6b7ce455e (update vendored files)
// NewT returns a new *T instance. "name" is the string returned by
// method Name.
func NewT(name string) *T {
	return &T{name: name}
}

// Run is a simplified version of testing.T.Run() method, without edge
// cases.
func (t *T) Run(name string, f func(*testing.T)) bool {
	t.CatchFailNow(func() { f(&t.T) })
	return !t.Failed()
}

// Name returns the name of the running test (in fact the one set by NewT).
func (t *T) Name() string {
	return t.name
}

// LogBuf is an ugly hack allowing to access internal testing.T log
// buffer. Keep cool, it is only used for internal unit tests.
func (t *T) LogBuf() string {
<<<<<<< HEAD
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) // nolint: govet
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) // nolint: govet
=======
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
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
type tFailedNow struct{}

>>>>>>> 4d7e5ad26 (update vendored files)
// NewT returns a new *T instance. "name" is the string returned by
// method Name.
func NewT(name string) *T {
	return &T{name: name}
}

// Run is a simplified version of testing.T.Run() method, without edge
// cases.
func (t *T) Run(name string, f func(*testing.T)) bool {
	t.CatchFailNow(func() { f(&t.T) })
	return !t.Failed()
}

// Name returns the name of the running test (in fact the one set by NewT).
func (t *T) Name() string {
	return t.name
}

// LogBuf is an ugly hack allowing to access internal testing.T log
// buffer. Keep cool, it is only used for internal unit tests.
func (t *T) LogBuf() string {
<<<<<<< HEAD
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) // nolint: govet
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) // nolint: govet
=======
	return string(reflect.ValueOf(t.T).FieldByName("output").Bytes()) //nolint: govet
}

// FailNow simulates the original (*testing.T).FailNow using
// panic. CatchFailNow should be used to properly intercept it.
func (t *T) FailNow() {
	t.Fail()
	panic(tFailedNow{})
}

// Fatal simulates the original (*testing.T).Fatal.
func (t *T) Fatal(args ...interface{}) {
	t.Helper()
	t.Error(args...)
	t.FailNow()
}

// Fatal simulates the original (*testing.T).Fatalf.
func (t *T) Fatalf(format string, args ...interface{}) {
	t.Helper()
	t.Errorf(format, args...)
	t.FailNow()
}

// CatchFailNow returns true if a FailNow, Fatal or Fatalf call
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
>>>>>>> 4d7e5ad26 (update vendored files)
}
