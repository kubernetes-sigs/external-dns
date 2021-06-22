// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"

	"github.com/maxatome/go-testdeep/internal/anchors"
)

// Anchors are stored globally by TestingFT
var allAnchors = map[string]*anchors.Info{}
var allAnchorsMu sync.Mutex

// AddAnchorableStructType declares a struct type as anchorable. "fn"
// is a function allowing to return a unique and identifiable instance
// of the struct type.
//
// "fn" has to have the following signature:
//
//   func (nextAnchor int) TYPE
//
// TYPE is the struct type to make anchorable and "nextAnchor" is an
// index to allow to differentiate several instances of the same type.
//
// For example, the time.Time type which is anchrorable by default,
// could be declared as:
//
//   AddAnchorableStructType(func (nextAnchor int) time.Time {
//     return time.Unix(int64(math.MaxInt64-1000424443-nextAnchor), 42)
//   })
//
// Just as a note, the 1000424443 constant allows to avoid to flirt
// with the math.MaxInt64 extreme limit and so avoid possible
// collision with real world values.
//
// It panics if the provided "fn" is not a function or if it has not
// the expected signature (see above).
func AddAnchorableStructType(fn interface{}) {
	err := anchors.AddAnchorableStructType(fn)
	if err != nil {
		panic(err.Error())
	}
}

// Anchor returns a typed value allowing to anchor the TestDeep
// operator "operator" in a go classic litteral like a struct, slice,
// array or map value.
//
// If the TypeBehind method of "operator" returns non-nil, "model" can
// be omitted (like with Between operator in the example
// below). Otherwise, "model" should contain only one value
// corresponding to the returning type. It can be:
//   - a go value: returning type is the type of the value, whatever the value is
//   - a reflect.Type
//
// It returns a typed value ready to be embed in a go data structure to
// be compared using T.Cmp or T.CmpLax:
//
//   import (
//     "testing"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestFunc(tt *testing.T) {
//     got := Func()
//
//     t := td.NewT(tt)
//     t.Cmp(got, &MyStruct{
//       Name:    "Bob",
//       Details: &MyDetails{
//         Nick: t.Anchor(td.HasPrefix("Bobby"), "").(string),
//         Age:  t.Anchor(td.Between(40, 50)).(int),
//       },
//     })
//   }
//
// In this example:
//
// HasPrefix operates on several input types (string, fmt.Stringer,
// error, …), so its TypeBehind method returns always nil as it can
// not guess in advance on which type it operates. In this case, we
// must pass "" as "model" parameter in order to tell it to return the
// string type. Note that the .(string) type assertion is then
// mandatory to conform to the strict type checking.
//
// Between, on its side, knows the type on which it operates, as it is
// the same as the one of its parameters. So its TypeBehind method
// returns the right type, and so no need to pass it as "model"
// parameter. Note that the .(int) type assertion is still mandatory
// to conform to the strict type checking.
//
// Without operator anchoring feature, the previous example would have
// been:
//
//   import (
//     "testing"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestFunc(tt *testing.T) {
//     got := Func()
//
//     t := td.NewT(tt)
//     t.Cmp(got, td.Struct(&MyStruct{Name: "Bob"},
//       td.StructFields{
//       "Details": td.Struct(&MyDetails{},
//         td.StructFields{
//           "Nick": td.HasPrefix("Bobby"),
//           "Age":  td.Between(40, 50),
//         }),
//     }))
//   }
//
// using two times the Struct operator to work around the strict type
// checking of golang.
//
// By default, the value returned by Anchor can only be used in the
// next T.Cmp or T.CmpLax call. To make it persistent across calls,
// see SetAnchorsPersist and AnchorsPersistTemporarily methods.
//
// See A method for a shorter synonym of Anchor.
func (t *T) Anchor(operator TestDeep, model ...interface{}) interface{} {
	if operator == nil {
		panic("Cannot anchor a nil TestDeep operator")
	}

	var typ reflect.Type
	if len(model) > 0 {
		if len(model) != 1 {
			panic("usage: Anchor(OPERATOR[, MODEL])")
		}
		var ok bool
		typ, ok = model[0].(reflect.Type)
		if !ok {
			vm := reflect.ValueOf(model[0])
			if !vm.IsValid() {
				panic("Untyped nil value is not valid as model for an anchor")
			}
			typ = vm.Type()
		}

		typeBehind := operator.TypeBehind()
		if typeBehind != nil && typeBehind != typ {
			panic(fmt.Sprintf("Operator %s TypeBehind() returned %s which differs from model type %s. Omit model or ensure its type is %[2]s",
				operator.GetLocation().Func, typeBehind, typ))
		}
	} else {
		typ = operator.TypeBehind()
		if typ == nil {
			panic(fmt.Sprintf("Cannot anchor operator %s as TypeBehind() returned nil. Use model parameter to specify the type to return",
				operator.GetLocation().Func))
		}
	}

	nvm := t.Config.anchors.AddAnchor(typ, reflect.ValueOf(operator))

	return nvm.Interface()
}

// A is a synonym for Anchor.
//
//   import (
//     "testing"
//
//     "github.com/maxatome/go-testdeep/td"
//   )
//
//   func TestFunc(tt *testing.T) {
//     got := Func()
//
//     t := td.NewT(tt)
//     t.Cmp(got, &MyStruct{
//       Name:    "Bob",
//       Details: &MyDetails{
//         Nick: t.A(td.HasPrefix("Bobby"), "").(string),
//         Age:  t.A(td.Between(40, 50)).(int),
//       },
//     })
//   }
func (t *T) A(operator TestDeep, model ...interface{}) interface{} {
	return t.Anchor(operator, model...)
}

func (t *T) resetNonPersistentAnchors() {
	t.Config.anchors.ResetAnchors(false)
}

// ResetAnchors frees all operators anchored with Anchor
// method. Unless operators anchoring persistence has been enabled
// with SetAnchorsPersist, there is no need to call this
// method. Anchored operators are automatically freed after each Cmp,
// CmpDeeply and CmpPanic call (or others methods calling them behind
// the scene).
func (t *T) ResetAnchors() {
	t.Config.anchors.ResetAnchors(true)
}

// AnchorsPersistTemporarily is used by helpers to temporarily enable
// anchors persistence. See tdhttp package for an example of use. It
// returns a function to be deferred, to restore the normal behavior
// (clear anchored operators if persistence was false, do nothing
// otherwise).
//
// Typically used as:
//   defer t.AnchorsPersistTemporarily()()
func (t *T) AnchorsPersistTemporarily() func() {
	// If already persistent, do nothing on defer
	if t.DoAnchorsPersist() {
		return func() {}
	}

	t.SetAnchorsPersist(true)
	return func() {
		t.SetAnchorsPersist(false)
		t.Config.anchors.ResetAnchors(true)
	}
}

// DoAnchorsPersist returns true if anchors persistence is enabled,
// false otherwise.
func (t *T) DoAnchorsPersist() bool {
	return t.Config.anchors.DoAnchorsPersist()
}

// SetAnchorsPersist allows to enable or disable anchors persistence.
func (t *T) SetAnchorsPersist(persist bool) {
	t.Config.anchors.SetAnchorsPersist(persist)
}

func (t *T) initAnchors() {
	if t.Config.anchors != nil {
		return
	}

	name := t.Name()

	allAnchorsMu.Lock()
	defer allAnchorsMu.Unlock()

	t.Config.anchors = allAnchors[name]
	if t.Config.anchors == nil {
		t.Config.anchors = anchors.NewInfo()
		allAnchors[name] = t.Config.anchors

		if name != "" {
			// Do not record a finalizer if no name (should not happen
			// except perhaps in tests)
			finalize := func() {
				allAnchorsMu.Lock()
				defer allAnchorsMu.Unlock()
				delete(allAnchors, name)
			}

			// From go 1.14, use Cleanup() method
			if tc, ok := t.TestingFT.(interface{ Cleanup(func()) }); ok {
				tc.Cleanup(finalize)
			} else {
				runtime.SetFinalizer(t.TestingFT, func(t TestingFT) { finalize() })
			}
		}
	}
}
