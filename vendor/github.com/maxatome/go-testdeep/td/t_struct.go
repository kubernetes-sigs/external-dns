// Copyright (c) 2018, 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"
	"strings"
	"sync"
	"testing"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/color"
	"github.com/maxatome/go-testdeep/internal/trace"
	"github.com/maxatome/go-testdeep/internal/types"
)

// T is a type that encapsulates [testing.TB] interface (which is
// implemented by [*testing.T] and [*testing.B]) allowing to easily use
// [*testing.T] methods as well as T ones.
type T struct {
	testing.TB
	Config ContextConfig // defaults to DefaultContextConfig
}

var _ testing.TB = T{}

// NewT returns a new [*T] instance. Typically used as:
//
//	import (
//	  "testing"
//
//	  "github.com/maxatome/go-testdeep/td"
//	)
//
//	type Record struct {
//	  Id        uint64
//	  Name      string
//	  Age       int
//	  CreatedAt time.Time
//	}
//
//	func TestCreateRecord(tt *testing.T) {
//	  t := NewT(tt, ContextConfig{
//	    MaxErrors: 3, // in case of failure, will dump up to 3 errors
//	  })
//
//	  before := time.Now()
//	  record, err := CreateRecord()
//
//	  if t.CmpNoError(err) {
//	    t.Log("No error, can now check struct contents")
//
//	    ok := t.Struct(record,
//	      &Record{
//	        Name: "Bob",
//	        Age:  23,
//	      },
//	      td.StructFields{
//	        "Id":        td.NotZero(),
//	        "CreatedAt": td.Between(before, time.Now()),
//	      },
//	      "Newly created record")
//	    if ok {
//	      t.Log(Record created successfully!")
//	    }
//	  }
//	}
//
// config is an optional parameter and, if passed, must be unique. It
// allows to configure how failures will be rendered during the
// lifetime of the returned instance.
//
//	t := NewT(tt)
//	t.Cmp(
//	  Record{Age: 12, Name: "Bob", Id: 12},  // got
//	  Record{Age: 21, Name: "John", Id: 28}) // expected
//
// will produce:
//
//	=== RUN   TestFoobar
//	--- FAIL: TestFoobar (0.00s)
//	        foobar_test.go:88: Failed test
//	                DATA.Id: values differ
//	                             got: (uint64) 12
//	                        expected: (uint64) 28
//	                DATA.Name: values differ
//	                             got: "Bob"
//	                        expected: "John"
//	                DATA.Age: values differ
//	                             got: 12
//	                        expected: 28
//	FAIL
//
// Now with a special configuration:
//
//	t := NewT(tt, ContextConfig{
//	    RootName:  "RECORD", // got data named "RECORD" instead of "DATA"
//	    MaxErrors: 2,        // stops after 2 errors instead of default 10
//	  })
//	t.Cmp(
//	  Record{Age: 12, Name: "Bob", Id: 12},  // got
//	  Record{Age: 21, Name: "John", Id: 28}, // expected
//	)
//
// will produce:
//
//	=== RUN   TestFoobar
//	--- FAIL: TestFoobar (0.00s)
//	        foobar_test.go:96: Failed test
//	                RECORD.Id: values differ
//	                             got: (uint64) 12
//	                        expected: (uint64) 28
//	                RECORD.Name: values differ
//	                             got: "Bob"
//	                        expected: "John"
//	                Too many errors (use TESTDEEP_MAX_ERRORS=-1 to see all)
//	FAIL
//
// See [T.RootName] method to configure RootName in a more specific fashion.
//
// Note that setting MaxErrors to a negative value produces a dump
// with all errors.
//
// If MaxErrors is not set (or set to 0), it is set to
// DefaultContextConfig.MaxErrors which is potentially dependent from
// the TESTDEEP_MAX_ERRORS environment variable (else defaults to 10.)
// See [ContextConfig] documentation for details.
//
// Of course t can already be a [*T], in this special case if config
// is omitted, the Config of the new instance is a copy of the t
// Config, including hooks.
//
// See also other constructors [Assert], [Require] and [AssertRequire].
//
// See also configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax], [T.IgnoreUnexported]
// and [T.TestDeepInGotOK].
func NewT(t testing.TB, config ...ContextConfig) *T {
	var newT T

	const usage = "NewT(testing.TB[, ContextConfig])"
	if t == nil {
		panic(color.BadUsage(usage, nil, 1, false))
	}
	if len(config) > 1 {
		t.Helper()
		t.Fatal(color.TooManyParams(usage))
	}

	// Already a *T, so steal its testing.TB and its Config if needed
	if tdT, ok := t.(*T); ok {
		newT.TB = tdT.TB
		if len(config) == 0 {
			newT.Config = tdT.Config
		} else {
			newT.Config = config[0]
		}
	} else {
		newT.TB = t
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

// Assert returns a new [*T] instance with FailureIsFatal flag set to
// false.
//
//	assert := Assert(t)
//
// is roughly equivalent to:
//
//	assert := NewT(t).FailureIsFatal(false)
//
// See [NewT] documentation for usefulness of config optional parameter.
//
// See also other constructors [Require] and [AssertRequire].
//
// See also configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax], [T.IgnoreUnexported]
// and [T.TestDeepInGotOK].
func Assert(t testing.TB, config ...ContextConfig) *T {
	return NewT(t, config...).FailureIsFatal(false)
}

// Require returns a new [*T] instance with FailureIsFatal flag set to
// true.
//
//	require := Require(t)
//
// is roughly equivalent to:
//
//	require := NewT(t).FailureIsFatal(true)
//
// See [NewT] documentation for usefulness of config optional parameter.
//
// See also other constructors [Assert] and [AssertRequire].
//
// See also configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax], [T.IgnoreUnexported]
// and [T.TestDeepInGotOK].
func Require(t testing.TB, config ...ContextConfig) *T {
	return NewT(t, config...).FailureIsFatal()
}

// AssertRequire returns 2 instances of [*T]. assert with
// FailureIsFatal flag set to false, and require with FailureIsFatal
// flag set to true.
//
//	assert, require := AssertRequire(t)
//
// is roughly equivalent to:
//
//	assert, require := Assert(t), Require(t)
//
// See [NewT] documentation for usefulness of config optional parameter.
//
// See also other constructors [Assert] and [Require].
//
// See also configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax], [T.IgnoreUnexported]
// and [T.TestDeepInGotOK].
func AssertRequire(t testing.TB, config ...ContextConfig) (assert, require *T) {
	assert = Assert(t, config...)
	require = assert.FailureIsFatal()
	return
}

// RootName changes the name of the got data. By default it is
// "DATA". For an HTTP response body, it could be "BODY" for example.
//
// It returns a new instance of [*T] so does not alter the original t
// and is used as follows:
//
//	t.RootName("RECORD").
//	  Struct(record,
//	    &Record{
//	      Name: "Bob",
//	      Age:  23,
//	    },
//	    td.StructFields{
//	      "Id":        td.NotZero(),
//	      "CreatedAt": td.Between(before, time.Now()),
//	    },
//	    "Newly created record")
//
// In case of error for the field Age, the failure message will contain:
//
//	RECORD.Age: values differ
//
// Which is more readable than the generic:
//
//	DATA.Age: values differ
//
// If "" is passed the name is set to "DATA", the default value.
//
// See also other configurators [T.Assert], [T.Require],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax], [T.IgnoreUnexported]
// and [T.TestDeepInGotOK].
func (t *T) RootName(rootName string) *T {
	nt := *t
	if rootName == "" {
		rootName = contextDefaultRootName
	}
	nt.Config.RootName = rootName
	return &nt
}

// FailureIsFatal allows to choose whether t.TB.Fatal() or
// t.TB.Error() will be used to print the next failure reports. When
// enable is true (or missing) testing.Fatal() will be called, else
// testing.Error(). Using [*testing.T] or [*testing.B] instance as
// t.TB value, FailNow() method is called behind the scenes when
// Fatal() is called. See [testing] documentation for details.
//
// It returns a new instance of [*T] so does not alter the original t
// and used as follows:
//
//	// Following t.Cmp() will call Fatal() if failure
//	t = t.FailureIsFatal()
//	t.Cmp(...)
//	t.Cmp(...)
//	// Following t.Cmp() won't call Fatal() if failure
//	t = t.FailureIsFatal(false)
//	t.Cmp(...)
//
// or, if only one call is critic:
//
//	// This Cmp() call will call Fatal() if failure
//	t.FailureIsFatal().Cmp(...)
//	// Following t.Cmp() won't call Fatal() if failure
//	t.Cmp(...)
//	t.Cmp(...)
//
// Note that t.FailureIsFatal() acts as t.FailureIsFatal(true).
//
// See also other configurators [T.Assert], [T.Require], [T.RootName],
// [T.UseEqual], [T.BeLax], [T.IgnoreUnexported] and
// [T.TestDeepInGotOK].
func (t *T) FailureIsFatal(enable ...bool) *T {
	nt := *t
	nt.Config.FailureIsFatal = len(enable) == 0 || enable[0]
	return &nt
}

// Assert returns a new [*T] instance inheriting the t config but with
// FailureIsFatal flag set to false.
//
// It returns a new instance of [*T] so does not alter the original t
//
// It is a shortcut for:
//
//	t.FailureIsFatal(false)
//
// See also other configurators [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax], [T.IgnoreUnexported]
// and [T.TestDeepInGotOK].
func (t *T) Assert() *T {
	return t.FailureIsFatal(false)
}

// Require returns a new [*T] instance inheriting the t config but
// with FailureIsFatal flag set to true.
//
// It returns a new instance of [*T] so does not alter the original t
//
// It is a shortcut for:
//
//	t.FailureIsFatal(true)
//
// See also other configurators [T.Assert], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax], [T.IgnoreUnexported]
// and [T.TestDeepInGotOK].
func (t *T) Require() *T {
	return t.FailureIsFatal(true)
}

// UseEqual tells go-testdeep to delegate the comparison of items
// whose type is one of types to their Equal() method.
//
// The signature this method should be:
//
//	(A) Equal(B) bool
//
// with B assignable to A.
//
// See [time.Time.Equal] as an example of accepted Equal() method.
//
// It always returns a new instance of [*T] so does not alter the
// original t.
//
//	t = t.UseEqual(time.Time{}, net.IP{})
//
// types items can also be [reflect.Type] items. In this case, the
// target type is the one reflected by the [reflect.Type].
//
//	t = t.UseEqual(reflect.TypeOf(time.Time{}), reflect.typeOf(net.IP{}))
//
// As a special case, calling t.UseEqual() or t.UseEqual(true) returns
// an instance using the Equal() method globally, for all types owning
// an Equal() method. Other types fall back to the default comparison
// mechanism. t.UseEqual(false) returns an instance not using Equal()
// method anymore, except for types already recorded using a previous
// UseEqual call.
//
// See also other configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.BeLax], [T.IgnoreUnexported] and
// [T.TestDeepInGotOK].
func (t *T) UseEqual(types ...any) *T {
	// special case: UseEqual()
	if len(types) == 0 {
		nt := *t
		nt.Config.UseEqual = true
		return &nt
	}

	// special cases: UseEqual(true) or UseEqual(false)
	if len(types) == 1 {
		if ignore, ok := types[0].(bool); ok {
			nt := *t
			nt.Config.UseEqual = ignore
			return &nt
		}
	}

	// Enable UseEqual only for types types
	t = t.copyWithHooks()

	err := t.Config.hooks.AddUseEqual(types)
	if err != nil {
		t.Helper()
		t.Fatal(color.Bad("UseEqual " + err.Error()))
	}

	return t
}

// BeLax allows to compare different but convertible types. If set to
// false, got and expected types must be the same. If set to true and
// expected type is convertible to got one, expected is first
// converted to go type before its comparison. See [CmpLax] or
// [T.CmpLax] and [Lax] operator to set this flag without providing a
// specific configuration.
//
// It returns a new instance of [*T] so does not alter the original t.
//
// Note that t.BeLax() acts as t.BeLax(true).
//
// See also other configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.IgnoreUnexported] and
// [T.TestDeepInGotOK].
func (t *T) BeLax(enable ...bool) *T {
	nt := *t
	nt.Config.BeLax = len(enable) == 0 || enable[0]
	return &nt
}

// IgnoreUnexported tells go-testdeep to ignore unexported fields of
// structs whose type is one of types.
//
// It always returns a new instance of [*T] so does not alter the original t.
//
//	t = t.IgnoreUnexported(MyStruct1{}, MyStruct2{})
//
// types items can also be [reflect.Type] items. In this case, the
// target type is the one reflected by the [reflect.Type].
//
//	t = t.IgnoreUnexported(reflect.TypeOf(MyStruct1{}))
//
// As a special case, calling t.IgnoreUnexported() or
// t.IgnoreUnexported(true) returns an instance ignoring unexported
// fields globally, for all struct types. t.IgnoreUnexported(false)
// returns an instance not ignoring unexported fields anymore, except
// for types already recorded using a previous IgnoreUnexported call.
//
// See also other configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax] and
// [T.TestDeepInGotOK].
func (t *T) IgnoreUnexported(types ...any) *T {
	// special case: IgnoreUnexported()
	if len(types) == 0 {
		nt := *t
		nt.Config.IgnoreUnexported = true
		return &nt
	}

	// special cases: IgnoreUnexported(true) or IgnoreUnexported(false)
	if len(types) == 1 {
		if ignore, ok := types[0].(bool); ok {
			nt := *t
			nt.Config.IgnoreUnexported = ignore
			return &nt
		}
	}

	// Enable IgnoreUnexported only for types types
	t = t.copyWithHooks()

	err := t.Config.hooks.AddIgnoreUnexported(types)
	if err != nil {
		t.Helper()
		t.Fatal(color.Bad("IgnoreUnexported " + err.Error()))
	}

	return t
}

// TestDeepInGotOK tells go-testdeep not to panic when a [TestDeep]
// operator is found on got side. By default it is forbidden because
// most of the time it is a mistake to compare (expected, got) instead
// of official (got, expected).
//
// It returns a new instance of [*T] so does not alter the original t.
//
// Note that t.TestDeepInGotOK() acts as t.TestDeepInGotOK(true).
//
// See also other configurators [T.Assert], [T.Require], [T.RootName],
// [T.FailureIsFatal], [T.UseEqual], [T.BeLax] and
// [T.IgnoreUnexported].
func (t *T) TestDeepInGotOK(enable ...bool) *T {
	nt := *t
	nt.Config.TestDeepInGotOK = len(enable) == 0 || enable[0]
	return &nt
}

// Cmp is mostly a shortcut for:
//
//	Cmp(t.TB, got, expected, args...)
//
// with the exception that t.Config is used to configure the test
// [ContextConfig].
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
func (t *T) Cmp(got, expected any, args ...any) bool {
	t.Helper()
	defer t.resetNonPersistentAnchors()
	return cmpDeeply(newContext(t), t.TB, got, expected, args...)
}

// CmpDeeply works the same as [Cmp] and is still available for
// compatibility purpose. Use shorter [Cmp] in new code.
func (t *T) CmpDeeply(got, expected any, args ...any) bool {
	t.Helper()
	defer t.resetNonPersistentAnchors()
	return cmpDeeply(newContext(t), t.TB, got, expected, args...)
}

// True is shortcut for:
//
//	t.Cmp(got, true, args...)
//
// Returns true if the test is OK, false if it fails.
//
//	t.True(IsAvailable(x), "x should be available")
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [T.False].
func (t *T) True(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, true, args...)
}

// False is shortcut for:
//
//	t.Cmp(got, false, args...)
//
// Returns true if the test is OK, false if it fails.
//
//	t.False(IsAvailable(x), "x should not be available")
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [T.True].
func (t *T) False(got any, args ...any) bool {
	t.Helper()
	return t.Cmp(got, false, args...)
}

// CmpError checks that got is non-nil error.
//
//	_, err := MyFunction(1, 2, 3)
//	t.CmpError(err, "MyFunction(1, 2, 3) should return an error")
//
// CmpError and not Error to avoid collision with t.TB.Error method.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [T.CmpNoError].
func (t *T) CmpError(got error, args ...any) bool {
	t.Helper()
	return cmpError(newContext(t), t.TB, got, args...)
}

// CmpNoError checks that got is nil error.
//
//	value, err := MyFunction(1, 2, 3)
//	if t.CmpNoError(err) {
//	  // one can now check value...
//	}
//
// CmpNoError and not NoError to be consistent with [T.CmpError] method.
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [T.CmpError].
func (t *T) CmpNoError(got error, args ...any) bool {
	t.Helper()
	return cmpNoError(newContext(t), t.TB, got, args...)
}

// CmpPanic calls fn and checks a panic() occurred with the
// expectedPanic parameter. It returns true only if both conditions
// are fulfilled.
//
// Note that calling panic(nil) in fn body is always detected as a
// panic. [runtime] package says: before Go 1.21, programs that called
// panic(nil) observed recover returning nil. Starting in Go 1.21,
// programs that call panic(nil) observe recover returning a
// [*runtime.PanicNilError]. Programs can change back to the old
// behavior by setting GODEBUG=panicnil=1.
//
//	t.CmpPanic(func() { panic("I am panicking!") },
//	  "I am panicking!",
//	  "The function should panic with the right string")
//
//	t.CmpPanic(func() { panic("I am panicking!") },
//	  Contains("panicking!"),
//	  "The function should panic with a string containing `panicking!`")
//
//	expected := &runtime.PanicNilError{} // should be nil if GODEBUG=panicnil=1
//	t.CmpPanic(t, func() { panic(nil) }, expected, "Checks for panic(nil)")
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [T.CmpNotPanic].
func (t *T) CmpPanic(fn func(), expected any, args ...any) bool {
	t.Helper()
	defer t.resetNonPersistentAnchors()
	return cmpPanic(newContext(t), t, fn, expected, args...)
}

// CmpNotPanic calls fn and checks no panic() occurred. If a panic()
// occurred false is returned then the panic() parameter and the stack
// trace appear in the test report.
//
// Note that calling panic(nil) in fn body is always detected as a
// panic. [runtime] package says: before Go 1.21, programs that called
// panic(nil) observed recover returning nil. Starting in Go 1.21,
// programs that call panic(nil) observe recover returning a
// [*runtime.PanicNilError]. Programs can change back to the old
// behavior by setting GODEBUG=panicnil=1.
//
//	t.CmpNotPanic(func() {}) // succeeds as function does not panic
//
//	t.CmpNotPanic(func() { panic("I am panicking!") }) // fails
//	t.CmpNotPanic(func() { panic(nil) })               // fails too
//
// args... are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of args is a string and contains a '%' rune then
// [fmt.Fprintf] is used to compose the name, else args are passed to
// [fmt.Fprint]. Do not forget it is the name of the test, not the
// reason of a potential failure.
//
// See also [T.CmpPanic].
func (t *T) CmpNotPanic(fn func(), args ...any) bool {
	t.Helper()
	return cmpNotPanic(newContext(t), t, fn, args...)
}

// Parallel marks this test as runnable in parallel with other
// parallel tests.  If t.TB implements Parallel(), as [*testing.T]
// does, it is usually used to mark top-level tests and/or subtests as
// safe for parallel execution:
//
//	func TestCreateRecord(tt *testing.T) {
//	  t := td.NewT(tt)
//	  t.Parallel()
//
//	  t.Run("no error", func(t *td.T) {
//	    t.Parallel()
//
//	    // ...
//	  })
//
// If t.TB does not implement Parallel(), this method is a no-op.
func (t *T) Parallel() {
	p, ok := t.TB.(interface{ Parallel() })
	if ok {
		p.Parallel()
	}
}

type runtFuncs struct {
	run reflect.Value
	fnt reflect.Type
}

var (
	runtMu sync.Mutex
	runt   = map[reflect.Type]runtFuncs{}
)

func (t *T) getRunFunc() (runtFuncs, bool) {
	ttb := reflect.TypeOf(t.TB)

	runtMu.Lock()
	defer runtMu.Unlock()

	vfuncs, ok := runt[ttb]
	if !ok {
		run, ok := ttb.MethodByName("Run")
		if ok {
			mt := run.Type
			if mt.NumIn() == 3 && mt.NumOut() == 1 && !mt.IsVariadic() &&
				mt.In(1) == types.String && mt.Out(0) == types.Bool {
				fnt := mt.In(2)
				if fnt.Kind() == reflect.Func &&
					fnt.NumIn() == 1 && fnt.NumOut() == 0 &&
					fnt.In(0) == mt.In(0) {
					vfuncs = runtFuncs{
						run: run.Func,
						fnt: fnt,
					}
					runt[ttb] = vfuncs
					ok = true
				}
			}
		}
		if !ok {
			runt[ttb] = vfuncs
		}
	}

	return vfuncs, vfuncs != (runtFuncs{})
}

// Run runs f as a subtest of t called name.
//
// If t.TB implement a method with the following signature:
//
//	(X) Run(string, func(X)) bool
//
// it calls it with a function of its own in which it creates a new
// instance of [*T] on the fly before calling f with it.
//
// So if t.TB is a [*testing.T] or a [*testing.B] (which is in normal
// cases), let's quote the [testing.T.Run] & [testing.B.Run]
// documentation: f is called in a separate goroutine and blocks
// until f returns or calls t.Parallel to become a parallel
// test. Run reports whether f succeeded (or at least did not fail
// before calling t.Parallel). Run may be called simultaneously from
// multiple goroutines, but all such calls must return before the
// outer test function for t returns.
//
// If this Run() method is not found, it simply logs name then
// executes f using a new [*T] instance in the current goroutine. Note
// that it is only done for convenience.
//
// The t param of f inherits the configuration of the self-reference.
//
// See also [T.RunAssertRequire],
// [github.com/maxatome/go-testdeep/helpers/tdsynctest.Run] and
// [github.com/maxatome/go-testdeep/helpers/tdsynctest.RunAssertRequire].
func (t *T) Run(name string, f func(t *T)) bool {
	t.Helper()

	vfuncs, ok := t.getRunFunc()
	if !ok {
		t = NewT(t)
		t.Logf("++++ %s", name)
		f(t)
		return !t.Failed()
	}

	conf := t.Config
	ret := vfuncs.run.Call([]reflect.Value{
		reflect.ValueOf(t.TB),
		reflect.ValueOf(name),
		reflect.MakeFunc(vfuncs.fnt,
			func(args []reflect.Value) (results []reflect.Value) {
				f(NewT(args[0].Interface().(testing.TB), conf))
				return nil
			}),
	})

	return ret[0].Bool()
}

// RunAssertRequire runs f as a subtest of t called name.
//
// If t.TB implement a method with the following signature:
//
//	(X) Run(string, func(X)) bool
//
// it calls it with a function of its own in which it creates two new
// instances of [*T] using [AssertRequire] on the fly before calling f
// with them.
//
// So if t.TB is a [*testing.T] or a [*testing.B] (which is in normal
// cases), let's quote the [testing.T.Run] & [testing.B.Run]
// documentation: f is called in a separate goroutine and blocks
// until f returns or calls t.Parallel to become a parallel
// test. Run reports whether f succeeded (or at least did not fail
// before calling t.Parallel). Run may be called simultaneously from
// multiple goroutines, but all such calls must return before the
// outer test function for t returns.
//
// If this Run() method is not found, it simply logs name then
// executes f using two new instances of [*T] (built with
// [AssertRequire]) in the current goroutine. Note that it is only
// done for convenience.
//
// The assert and require params of f inherit the configuration
// of the self-reference, except that a failure is never fatal using
// assert and always fatal using require.
//
// See also [T.Run],
// [github.com/maxatome/go-testdeep/helpers/tdsynctest.RunAssertRequire]
// and [github.com/maxatome/go-testdeep/helpers/tdsynctest.Run].
func (t *T) RunAssertRequire(name string, f func(assert, require *T)) bool {
	t.Helper()

	vfuncs, ok := t.getRunFunc()
	if !ok {
		assert, require := AssertRequire(t)
		t.Logf("++++ %s", name)
		f(assert, require)
		return !t.Failed()
	}

	conf := t.Config
	ret := vfuncs.run.Call([]reflect.Value{
		reflect.ValueOf(t.TB),
		reflect.ValueOf(name),
		reflect.MakeFunc(vfuncs.fnt,
			func(args []reflect.Value) (results []reflect.Value) {
				f(AssertRequire(args[0].Interface().(testing.TB), conf))
				return nil
			}),
	})

	return ret[0].Bool()
}

// RunT runs f as a subtest of t called name.
//
// Deprecated: RunT has been superseded by [T.Run] method. It is kept
// for compatibility.
func (t *T) RunT(name string, f func(t *T)) bool {
	t.Helper()
	return t.Run(name, f)
}

func getTrace(args ...any) string {
	var b strings.Builder
	tdutil.FbuildTestName(&b, args...)

	if b.Len() == 0 {
		b.WriteString("Stack trace:\n")
	} else if !strings.HasSuffix(b.String(), "\n") {
		b.WriteByte('\n')
	}

	s := stripTrace(trace.Retrieve(1, "testing.tRunner"))
	if len(s) == 0 {
		b.WriteString("\tEmpty stack trace")
		return b.String()
	}

	s.Dump(&b)
	return b.String()
}

// LogTrace uses t.TB.Log() to log a stack trace.
//
// args... are optional and allow to prefix the trace by a
// message. If empty, this message defaults to "Stack trace:\n". If
// this message does not end with a "\n", one is automatically
// added. If len(args) > 1 and the first item of args is a string
// and contains a '%' rune then [fmt.Fprintf] is used to compose the
// name, else args are passed to [fmt.Fprint].
//
// See also [T.ErrorTrace] and [T.FatalTrace].
func (t *T) LogTrace(args ...any) {
	t.Helper()
	t.Log(getTrace(args...))
}

// ErrorTrace uses t.TB.Error() to log a stack trace.
//
// args... are optional and allow to prefix the trace by a
// message. If empty, this message defaults to "Stack trace:\n". If
// this message does not end with a "\n", one is automatically
// added. If len(args) > 1 and the first item of args is a string
// and contains a '%' rune then [fmt.Fprintf] is used to compose the
// name, else args are passed to [fmt.Fprint].
//
// See also [T.LogTrace] and [T.FatalTrace].
func (t *T) ErrorTrace(args ...any) {
	t.Helper()
	t.Error(getTrace(args...))
}

// FatalTrace uses t.TB.Fatal() to log a stack trace.
//
// args... are optional and allow to prefix the trace by a
// message. If empty, this message defaults to "Stack trace:\n". If
// this message does not end with a "\n", one is automatically
// added. If len(args) > 1 and the first item of args is a string
// and contains a '%' rune then [fmt.Fprintf] is used to compose the
// name, else args are passed to [fmt.Fprint].
//
// See also [T.LogTrace] and [T.ErrorTrace].
func (t *T) FatalTrace(args ...any) {
	t.Helper()
	t.Fatal(getTrace(args...))
}
