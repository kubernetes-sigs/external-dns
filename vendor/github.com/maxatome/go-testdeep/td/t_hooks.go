// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"github.com/maxatome/go-testdeep/internal/color"
)

// WithCmpHooks returns a new *T instance with new Cmp hooks recorded
// using functions passed in "fns".
//
// Each function in "fns" has to be a function with the following
// possible signatures:
//   func (A, A) bool
//   func (A, A) error
// First arg is always "got", and second is always "expected".
//
// A cannot be an interface. This retriction can be removed in the
// future, if really needed.
//
// This function is called as soon as possible each time the type A is
// encountered for "got" while "expected" type is assignable to A.
//
// When it returns a bool, false means A is not equal to B.
//
// When it returns a non-nil error (meaning "got" ≠ "expected"), its
// contents is used to tell the reason of the failure.
//
// Cmp hooks are checked before UseEqual feature.
// Cmp hooks are run just after Smuggle hooks.
//
//   func TestCmpHook(tt *testing.T) {
//     t := td.NewT(tt)
//
//     // Test reflect.Value contents instead of default field/field
//     t = t.WithCmpHooks(func (got, expected reflect.Value) bool {
//       return td.EqDeeply(got.Interface(), expected.Interface())
//     })
//     a, b := 1, 1
//     t.Cmp(reflect.ValueOf(&a), reflect.ValueOf(&b)) // succeeds
//
//     // Test reflect.Type correctly instead of default field/field
//     t = t.WithCmpHooks(func (got, expected reflect.Type) bool {
//       return got == expected
//     })
//
//     // Test time.Time via its Equal() method instead of default
//     // field/field (note it bypasses the UseEqual flag)
//     t = t.WithCmpHooks((time.Time).Equal)
//     date, _ := time.Parse(time.RFC3339, "2020-09-08T22:13:54+02:00")
//     t.Cmp(date, date.UTC()) // succeeds
//
//     // Several hooks can be declared at once
//     t = t.WithCmpHooks(
//       func (got, expected reflect.Value) bool {
//         return td.EqDeeply(got.Interface(), expected.Interface())
//       },
//       func (got, expected reflect.Type) bool {
//         return got == expected
//       },
//       (time.Time).Equal,
//     )
//   }
//
// There is no way to add or remove hooks of an existing *T instance,
// only create a new *T instance with this method or WithSmuggleHooks
// to add some.
//
// WithCmpHooks calls t.Fatal if an item of "fns" is not a function or
// if its signature does not match the expected ones.
func (t *T) WithCmpHooks(fns ...interface{}) *T {
	t = t.copyWithHooks()

	err := t.Config.hooks.AddCmpHooks(fns)
	if err != nil {
		t.Helper()
		t.Fatal(color.Bad("WithCmpHooks " + err.Error()))
	}

	return t
}

// WithSmuggleHooks returns a new *T instance with new Smuggle hooks
// recorded using functions passed in "fns".
//
// Each function in "fns" has to be a function with the following
// possible signatures:
//   func (A) B
//   func (A) (B, error)
//
// A cannot be an interface. This retriction can be removed in the
// future, if really needed.
//
// B cannot be an interface. If you have a use case, we can talk about it.
//
// This function is called as soon as possible each time the type A is
// encountered for "got".
//
// The B value returned replaces the "got" value for subsequent tests.
// Smuggle hooks are NOT run again for this returned value to avoid
// easy infinite loop recursion.
//
// When it returns non-nil error (meaning something wrong happened
// during the conversion of A to B), it raises a global error and its
// contents is used to tell the reason of the failure.
//
// Smuggle hooks are run just before Cmp hooks.
//
//   func TestSmuggleHook(tt *testing.T) {
//     t := td.NewT(tt)
//
//     // Each encountered int is changed to a bool
//     t = t.WithSmuggleHooks(func (got int) bool {
//       return got != 0
//     })
//     t.Cmp(map[string]int{"ok": 1, "no": 0},
//       map[string]bool{"ok", true, "no", false}) // succeeds
//
//     // Each encountered string is converted to int
//     t = t.WithSmuggleHooks(strconv.Atoi)
//     t.Cmp("123", 123) // succeeds
//
//     // Several hooks can be declared at once
//     t = t.WithSmuggleHooks(
//       func (got int) bool { return got != 0 },
//       strconv.Atoi,
//     )
//   }
//
// There is no way to add or remove hooks of an existing *T instance,
// only create a new *T instance with this method or WithCmpHooks to add some.
//
// WithSmuggleHooks calls t.Fatal if an item of "fns" is not a
// function or if its signature does not match the expected ones.
func (t *T) WithSmuggleHooks(fns ...interface{}) *T {
	t = t.copyWithHooks()

	err := t.Config.hooks.AddSmuggleHooks(fns)
	if err != nil {
		t.Helper()
		t.Fatal(color.Bad("WithSmuggleHooks " + err.Error()))
	}

	return t
}

func (t *T) copyWithHooks() *T {
	nt := NewT(t)
	nt.Config.hooks = t.Config.hooks.Copy()
	return nt
}
