// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package hooks

import (
	"errors"
	"fmt"
	"reflect"
	"sync"

	"github.com/maxatome/go-testdeep/internal/types"
)

type properties struct {
	cmp              reflect.Value
	smuggle          reflect.Value
	ignoreUnexported bool
	useEqual         bool
}

// Info gathers all hooks information.
type Info struct {
	sync.Mutex
	props map[reflect.Type]properties
}

// NewInfo returns a new instance of *Info.
func NewInfo() *Info {
	return &Info{
		props: map[reflect.Type]properties{},
	}
}

var ErrBoolean = errors.New("CmpHook(got, expected) failed")

// Copy returns a new instance of *Info with the same hooks as i. As a
// special case, if i is nil, returned instance is non-nil.
func (i *Info) Copy() *Info {
	ni := NewInfo()

	if i == nil {
		return ni
	}

	i.Lock()
	defer i.Unlock()

	if len(i.props) == 0 {
		return ni
	}

	ni.props = make(map[reflect.Type]properties, len(i.props))
	for t, p := range i.props {
		ni.props[t] = p
	}

	return ni
}

// AddCmpHooks records new Cmp hooks using functions contained in "fns".
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
// It returns an error if an item of "fns" is not a function or if its
// signature does not match the expected ones.
func (i *Info) AddCmpHooks(fns []interface{}) error {
	for n, fn := range fns {
		vfn := reflect.ValueOf(fn)

		if vfn.Kind() != reflect.Func {
			return fmt.Errorf("expects a function, not a %s (@%d)", vfn.Kind(), n)
		}

		ft := vfn.Type()
		if !ft.IsVariadic() &&
			ft.NumIn() == 2 &&
			ft.NumOut() == 1 &&
			ft.In(0) == ft.In(1) &&
			ft.In(0).Kind() != reflect.Interface &&
			(ft.Out(0) == types.Bool || ft.Out(0) == types.Error) {
			i.Lock()
			prop := i.props[ft.In(0)]
			prop.cmp = vfn
			i.props[ft.In(0)] = prop
			i.Unlock()
			continue
		}

		return fmt.Errorf("expects: func (T, T) bool|error not %s (@%d)", ft, n)
	}
	return nil
}

// Cmp checks if a Cmp hook exists matching "got" and "expected" types.
//
// If no, it returns (false, nil)
//
// If yes, it calls it and returns (true, nil) if it succeeds,
// (true, <an error>) if it fails. If the hook returns a false bool, the
// error returned is ErrBoolean.
func (i *Info) Cmp(got, expected reflect.Value) (bool, error) {
	if i == nil {
		return false, nil
	}

	tg := got.Type()

	i.Lock()
	prop, ok := i.props[tg]
	i.Unlock()
	if !ok || !prop.cmp.IsValid() {
		return false, nil
	}

	if !expected.Type().AssignableTo(prop.cmp.Type().In(1)) {
		return false, nil
	}

	res := prop.cmp.Call([]reflect.Value{got, expected})[0]
	if res.Kind() == reflect.Bool {
		if res.Bool() {
			return true, nil
		}
		return true, ErrBoolean
	}
	err, _ := res.Interface().(error)
	return true, err
}

// AddSmuggleHooks records new Smuggle hooks using functions contained
// in "fns".
//
// Each function in "fns" has to be a function with the following
// possible signatures:
//   func (A) B
//   func (A) (B, error)
//
// A cannot be an interface. This retriction can be removed in the
// future, if really needed.
//
// B can be an interface.
//
// It returns an error if an item of "fns" is not a function or if its
// signature does not match the expected ones.
func (i *Info) AddSmuggleHooks(fns []interface{}) error {
	for n, fn := range fns {
		vfn := reflect.ValueOf(fn)

		if vfn.Kind() != reflect.Func {
			return fmt.Errorf("expects a function, not a %s (@%d)", vfn.Kind(), n)
		}

		ft := vfn.Type()
		if !ft.IsVariadic() &&
			ft.NumIn() == 1 &&
			ft.In(0).Kind() != reflect.Interface &&
			(ft.NumOut() == 1 || (ft.NumOut() == 2 && ft.Out(1) == types.Error)) &&
			ft.Out(0).Kind() != reflect.Interface {
			i.Lock()
			prop := i.props[ft.In(0)]
			prop.smuggle = vfn
			i.props[ft.In(0)] = prop
			i.Unlock()
			continue
		}

		return fmt.Errorf("expects: func (A) (B[, error]) not %s (@%d)", ft, n)
	}
	return nil
}

// Smuggle checks if a Smuggle hook exists matching "*got" type.
//
// If no, it returns (false, nil)
//
// If yes, it calls it and returns (true, nil) if it succeeds,
// (true, <an error>) if it fails.
func (i *Info) Smuggle(got *reflect.Value) (bool, error) {
	if i == nil {
		return false, nil
	}

	tg := got.Type()

	i.Lock()
	prop, ok := i.props[tg]
	i.Unlock()
	if !ok || !prop.smuggle.IsValid() {
		return false, nil
	}

	res := prop.smuggle.Call([]reflect.Value{*got})
	if len(res) == 2 {
		if err, _ := res[1].Interface().(error); err != nil {
			return true, err
		}
	}

	*got = res[0]
	return true, nil
}

// AddUseEqual records types of values contained in "ts" as using
// Equal method. "ts" can also contain reflect.Type instances.
func (i *Info) AddUseEqual(ts []interface{}) error {
	if len(ts) == 0 {
		return nil
	}
	for n, typ := range ts {
		t, ok := typ.(reflect.Type)
		if !ok {
			t = reflect.TypeOf(typ)
			ts[n] = t
		}

		equal, ok := t.MethodByName("Equal")
		if !ok {
			return fmt.Errorf("expects type %s owns an Equal method (@%d)", t, n)
		}

		ft := equal.Type
		if ft.IsVariadic() ||
			ft.NumIn() != 2 ||
			ft.NumOut() != 1 ||
			!ft.In(0).AssignableTo(ft.In(1)) ||
			ft.Out(0) != types.Bool {
			return fmt.Errorf("expects type %[1]s Equal method signature be Equal(%[1]s) bool (@%[2]d)", t, n)
		}
	}

	i.Lock()
	defer i.Unlock()

	for _, typ := range ts {
		t := typ.(reflect.Type)
		prop := i.props[t]
		prop.useEqual = true
		i.props[t] = prop
	}
	return nil
}

// UseEqual returns true if the type "t" needs to use its Equal method
// to be compared.
func (i *Info) UseEqual(t reflect.Type) bool {
	if i == nil {
		return false
	}

	i.Lock()
	defer i.Unlock()
	return i.props[t].useEqual
}

// AddIgnoreUnexported records types of values contained in "ts" as ignoring
// unexported struct fields. "ts" can also contain reflect.Type instances.
func (i *Info) AddIgnoreUnexported(ts []interface{}) error {
	if len(ts) == 0 {
		return nil
	}
	for n, typ := range ts {
		t, ok := typ.(reflect.Type)
		if !ok {
			t = reflect.TypeOf(typ)
			ts[n] = t
		}

		if t.Kind() != reflect.Struct {
			return fmt.Errorf("expects type %s be a struct, not a %s (@%d)", t, t.Kind(), n)
		}
	}

	i.Lock()
	defer i.Unlock()

	for _, typ := range ts {
		t := typ.(reflect.Type)
		prop := i.props[t]
		prop.ignoreUnexported = true
		i.props[t] = prop
	}
	return nil
}

// IgnoreUnexported returns true if the unexported fields of the type
// "t" have to be ignored.
func (i *Info) IgnoreUnexported(t reflect.Type) bool {
	if i == nil {
		return false
	}

	i.Lock()
	defer i.Unlock()
	return i.props[t].ignoreUnexported
}
