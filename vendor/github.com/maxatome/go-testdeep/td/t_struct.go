// Copyright (c) 2018, 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"testing"
)

// T is a type that encapsulates *testing.T (in fact TestingFT
// interface which is implemented by *testing.T) allowing to easily
// use *testing.T methods as well as T ones.
type T struct {
	TestingFT
	Config ContextConfig // defaults to DefaultContextConfig
}

var _ TestingFT = T{}

// NewT returns a new T instance. Typically used as:
//
//   import (
//     "testing"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   type Record struct {
//     Id        uint64
//     Name      string
//     Age       int
//     CreatedAt time.Time
//   }
//
//   func TestCreateRecord(tt *testing.T) {
//     t := NewT(tt, ContextConfig{
//       MaxErrors: 3, // in case of failure, will dump up to 3 errors
//     })
//
//     before := time.Now()
//     record, err := CreateRecord()
//
//     if t.CmpNoError(err) {
//       t.Log("No error, can now check struct contents")
//
//       ok := t.Struct(record,
//         &Record{
//           Name: "Bob",
//           Age:  23,
//         },
//         td.StructFields{
//           "Id":        td.NotZero(),
//           "CreatedAt": td.Between(before, time.Now()),
//         },
//         "Newly created record")
//       if ok {
//         t.Log(Record created successfully!")
//       }
//     }
//   }
//
// "config" is an optional argument and, if passed, must be unique. It
// allows to configure how failures will be rendered during the
// lifetime of the returned instance.
//
//   t := NewT(tt)
//   t.Cmp(
//     Record{Age: 12, Name: "Bob", Id: 12},  // got
//     Record{Age: 21, Name: "John", Id: 28}) // expected
//
// will produce:
//
//   === RUN   TestFoobar
//   --- FAIL: TestFoobar (0.00s)
//           foobar_test.go:88: Failed test
//                   DATA.Id: values differ
//                                got: (uint64) 12
//                           expected: (uint64) 28
//                   DATA.Name: values differ
//                                got: "Bob"
//                           expected: "John"
//                   DATA.Age: values differ
//                                got: 12
//                           expected: 28
//   FAIL
//
// Now with a special configuration:
//
//   t := NewT(tt, ContextConfig{
//       RootName:  "RECORD", // got data named "RECORD" instead of "DATA"
//       MaxErrors: 2,        // stops after 2 errors instead of default 10
//     })
//   t.Cmp(
//     Record{Age: 12, Name: "Bob", Id: 12},  // got
//     Record{Age: 21, Name: "John", Id: 28}, // expected
//   )
//
// will produce:
//
//   === RUN   TestFoobar
//   --- FAIL: TestFoobar (0.00s)
//           foobar_test.go:96: Failed test
//                   RECORD.Id: values differ
//                                got: (uint64) 12
//                           expected: (uint64) 28
//                   RECORD.Name: values differ
//                                got: "Bob"
//                           expected: "John"
//                   Too many errors (use TESTDEEP_MAX_ERRORS=-1 to see all)
//   FAIL
//
// See RootName method to configure RootName in a more specific fashion.
//
// Note that setting MaxErrors to a negative value produces a dump
// with all errors.
//
// If MaxErrors is not set (or set to 0), it is set to
// DefaultContextConfig.MaxErrors which is potentially dependent from
// the TESTDEEP_MAX_ERRORS environment variable (else defaults to 10.)
// See ContextConfig documentation for details.
//
// Of course "t" can already be a *T, in this special case if "config"
// is omitted, the Config of the new instance is a copy of the "t"
// Config.
func NewT(t TestingFT, config ...ContextConfig) *T {
	var newT T

	if len(config) > 1 || t == nil {
		panic("usage: NewT(TestingFT[, ContextConfig]")
	}

	// Already a *T, so steal its TestingFT and its Config if needed
	if tdT, ok := t.(*T); ok {
		newT.TestingFT = tdT.TestingFT
		if len(config) == 0 {
			newT.Config = tdT.Config
		} else {
			newT.Config = config[0]
		}
	} else {
		newT.TestingFT = t
		if len(config) == 0 {
			newT.Config = DefaultContextConfig
		} else {
			newT.Config = config[0]
		}
	}
	newT.Config.sanitize()

	newT.initAnchors()

	return &newT
}

// Assert return a new T instance with FailureIsFatal flag set to
// false.
//
//   assert := Assert(t)
//
// is roughly equivalent to:
//
//   assert := NewT(t).FailureIsFatal(false)
//
// See NewT documentation for usefulness of "config" optional parameter.
func Assert(t TestingFT, config ...ContextConfig) *T {
	return NewT(t, config...).FailureIsFatal(false)
}

// Require return a new T instance with FailureIsFatal flag set to
// true.
//
//   require := Require(t)
//
// is roughly equivalent to:
//
//   require := NewT(t).FailureIsFatal(true)
//
// See NewT documentation for usefulness of "config" optional parameter.
func Require(t TestingFT, config ...ContextConfig) *T {
	return NewT(t, config...).FailureIsFatal()
}

// AssertRequire returns 2 instances of T. The first one called
// "assert" with FailureIsFatal flag set to false, and the second
// called "require" with FailureIsFatal flag set to true.
//
//   assert, require := AssertRequire(t)
//
// is roughly equivalent to:
//
//   assert, require := Assert(t), Require(t)
//
// See NewT documentation for usefulness of "config" optional parameter.
func AssertRequire(t TestingFT, config ...ContextConfig) (*T, *T) {
	assert := Assert(t, config...)
	return assert, assert.FailureIsFatal()
}

// RootName changes the name of the got data. By default it is
// "DATA". For an HTTP response body, it could be "BODY" for example.
//
// It returns a new instance of *T so does not alter the original t
// and used as follows:
//
//   t.RootName("RECORD").
//     Struct(record,
//       &Record{
//         Name: "Bob",
//         Age:  23,
//       },
//       td.StructFields{
//         "Id":        td.NotZero(),
//         "CreatedAt": td.Between(before, time.Now()),
//       },
//       "Newly created record")
//
// In case of error for the field Age, the failure message will contain:
//
//   RECORD.Age: values differ
//
// Which is more readable than the generic:
//
//   DATA.Age: values differ
//
// If "" is passed the name is set to "DATA", the default value.
func (t *T) RootName(rootName string) *T {
	new := *t
	if rootName == "" {
		rootName = contextDefaultRootName
	}
	new.Config.RootName = rootName
	return &new
}

// FailureIsFatal allows to choose whether t.TestingFT.Fatal() or
// t.TestingFT.Error() will be used to print the next failure
// reports. When "enable" is true (or missing) testing.Fatal() will be
// called, else testing.Error(). Using *testing.T instance as
// t.TestingFT value, FailNow() is called behind the scenes when
// Fatal() is called. See testing documentation for details.
//
// It returns a new instance of *T so does not alter the original t
// and used as follows:
//
//   // Following t.Cmp() will call Fatal() if failure
//   t = t.FailureIsFatal()
//   t.Cmp(...)
//   t.Cmp(...)
//   // Following t.Cmp() won't call Fatal() if failure
//   t = t.FailureIsFatal(false)
//   t.Cmp(...)
//
// or, if only one call is critic:
//
//   // This Cmp() call will call Fatal() if failure
//   t.FailureIsFatal().Cmp(...)
//   // Following t.Cmp() won't call Fatal() if failure
//   t.Cmp(...)
//   t.Cmp(...)
//
// Note that t.FailureIsFatal() acts as t.FailureIsFatal(true).
func (t *T) FailureIsFatal(enable ...bool) *T {
	new := *t
	new.Config.FailureIsFatal = len(enable) == 0 || enable[0]
	return &new
}

// UseEqual allows to use the Equal method on got (if it exists) or
// on any of its component to compare got and expected values.
//
// The signature should be:
//   (A) Equal(B) bool
// with B assignable to A.
//
// See time.Time as an example of accepted Equal() method.
//
// It returns a new instance of *T so does not alter the original t.
//
// Note that t.UseEqual() acts as t.UseEqual(true).
func (t *T) UseEqual(enable ...bool) *T {
	new := *t
	new.Config.UseEqual = len(enable) == 0 || enable[0]
	return &new
}

// BeLax allows to to compare different but convertible types. If
// set to false, got and expected types must be the same. If set to
// true and expected type is convertible to got one, expected is
// first converted to go type before its comparison. See CmpLax
// function/method and Lax operator to set this flag without
// providing a specific configuration.
//
// It returns a new instance of *T so does not alter the original t.
//
// Note that t.BeLax() acts as t.BeLax(true).
func (t *T) BeLax(enable ...bool) *T {
	new := *t
	new.Config.BeLax = len(enable) == 0 || enable[0]
	return &new
}

// Cmp is mostly a shortcut for:
//
//   Cmp(t.TestingFT, got, expected, args...)
//
// with the exception that t.Config is used to configure the test
// Context.
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Cmp(got, expected interface{}, args ...interface{}) bool {
	t.Helper()
	defer t.resetNonPersistentAnchors()
	return cmpDeeply(newContextWithConfig(t.Config),
		t.TestingFT, got, expected, args...)
}

// CmpDeeply works the same as Cmp and is still available for
// compatibility purpose. Use shorter Cmp in new code.
func (t *T) CmpDeeply(got, expected interface{}, args ...interface{}) bool {
	t.Helper()
	defer t.resetNonPersistentAnchors()
	return cmpDeeply(newContextWithConfig(t.Config),
		t.TestingFT, got, expected, args...)
}

// True is shortcut for:
//
//   t.Cmp(got, true, args...)
//
// Returns true if the test is OK, false if it fails.
//
//   t.True(IsAvailable(x), "x should be available")
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) True(got interface{}, args ...interface{}) bool {
	t.Helper()
	return t.Cmp(got, true, args...)
}

// False is shortcut for:
//
//   t.Cmp(got, false, args...)
//
// Returns true if the test is OK, false if it fails.
//
//   t.False(IsAvailable(x), "x should not be available")
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) False(got interface{}, args ...interface{}) bool {
	t.Helper()
	return t.Cmp(got, false, args...)
}

// CmpError checks that "got" is non-nil error.
//
//   _, err := MyFunction(1, 2, 3)
//   t.CmpError(err, "MyFunction(1, 2, 3) should return an error")
//
// CmpError and not Error to avoid collision with t.TestingFT.Error method.
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) CmpError(got error, args ...interface{}) bool {
	t.Helper()
	return cmpError(newContextWithConfig(t.Config), t.TestingFT, got, args...)
}

// CmpNoError checks that "got" is nil error.
//
//   value, err := MyFunction(1, 2, 3)
//   if t.CmpNoError(err) {
//     // one can now check value...
//   }
//
// CmpNoError and not NoError to be consistent with CmpError method.
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) CmpNoError(got error, args ...interface{}) bool {
	t.Helper()
	return cmpNoError(newContextWithConfig(t.Config), t.TestingFT, got, args...)
}

// CmpPanic calls "fn" and checks a panic() occurred with the
// "expectedPanic" parameter. It returns true only if both conditions
// are fulfilled.
//
// Note that calling panic(nil) in "fn" body is detected as a panic
// (in this case "expectedPanic" has to be nil).
//
//   t.CmpPanic(func() { panic("I am panicking!") },
//     "I am panicking!",
//     "The function should panic with the right string")
//
//   t.CmpPanic(func() { panic("I am panicking!") },
//     Contains("panicking!"),
//     "The function should panic with a string containing `panicking!`")
//
//   t.CmpPanic(t, func() { panic(nil) }, nil, "Checks for panic(nil)")
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) CmpPanic(fn func(), expected interface{}, args ...interface{}) bool {
	t.Helper()
	defer t.resetNonPersistentAnchors()
	return cmpPanic(newContextWithConfig(t.Config), t, fn, expected, args...)
}

// CmpNotPanic calls "fn" and checks no panic() occurred. If a panic()
// occurred false is returned then the panic() parameter and the stack
// trace appear in the test report.
//
// Note that calling panic(nil) in "fn" body is detected as a panic.
//
//   t.CmpNotPanic(func() {}) // succeeds as function does not panic
//
//   t.CmpNotPanic(func() { panic("I am panicking!") }) // fails
//   t.CmpNotPanic(func() { panic(nil) })               // fails too
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) CmpNotPanic(fn func(), args ...interface{}) bool {
	t.Helper()
	return cmpNotPanic(newContextWithConfig(t.Config), t, fn, args...)
}

// RunT runs "f" as a subtest of t called "name". It runs "f" in a
// separate goroutine and blocks until "f" returns or calls t.Parallel
// to become a parallel test. RunT reports whether "f" succeeded (or at
// least did not fail before calling t.Parallel).
//
// RunT may be called simultaneously from multiple goroutines, but all
// such calls must return before the outer test function for t
// returns.
//
// Under the hood, RunT delegates all this stuff to testing.Run. That
// is why this documentation is a copy/paste of testing.Run one.
//
// In versions up to v1.0.8, the name of this function was Run. As *T
// now implements TestingFT interface, the original
// (*testing.T).Run(string, func(t *testing.T)) is callable directly
// on *T.
func (t *T) RunT(name string, f func(t *T)) bool {
	t.Helper()
	return t.TestingFT.Run(name, func(tt *testing.T) { f(NewT(tt, t.Config)) })
}
