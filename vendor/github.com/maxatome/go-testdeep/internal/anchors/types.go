// Copyright (c) 2020, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package anchors

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

var intType = reflect.TypeOf(42)

type anchorableType struct {
	typ     reflect.Type
	builder reflect.Value
}

// AnchorableTypes contains all non-native types that can be
// anchorable. See AddAnchorableStructType to add a new type to it.
var AnchorableTypes []anchorableType

func init() {
	AddAnchorableStructType(func(nextAnchor int) time.Time { //nolint: errcheck
		return time.Unix(math.MaxInt64-1000424443-int64(nextAnchor), 42)
	})
}

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
// is declared as:
//
//   AddAnchorableStructType(func (nextAnchor int) time.Time {
//     return time.Unix(math.MaxInt64-1000424443-int64(nextAnchor), 42)
//   })
//
// Just as a note, the 1000424443 constant allows to avoid to flirt
// with the math.MaxInt64 extreme limit and so avoid possible
// collision with real world values.
//
// It returns an error if the provided "fn" is not a function or if it
// has not the expected signature (see above).
func AddAnchorableStructType(fn interface{}) error {
	vfn := reflect.ValueOf(fn)

	if vfn.Kind() == reflect.Func {
		fnType := vfn.Type()

		if !fnType.IsVariadic() &&
			fnType.NumIn() == 1 && fnType.NumOut() == 1 &&
			fnType.In(0) == intType &&
			fnType.Out(0).Kind() == reflect.Struct {
			typ := fnType.Out(0)
			if !typ.Comparable() {
				return fmt.Errorf(
					"type %s is not comparable, it cannot be anchorable", typ)
			}

			for i, at := range AnchorableTypes {
				if at.typ == typ {
					AnchorableTypes[i].builder = vfn
					return nil
				}
			}

			AnchorableTypes = append(AnchorableTypes, anchorableType{
				typ:     typ,
				builder: vfn,
			})
			return nil
		}
	}

	return errors.New("usage: AddAnchorableStructType(func (nextAnchor int) STRUCT_TYPE)")
}
