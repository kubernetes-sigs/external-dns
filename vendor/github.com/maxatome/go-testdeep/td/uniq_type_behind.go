// Copyright (c) 2018-2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"
)

// uniqTypeBehindSlice can return a non-nil reflect.Type if all items
// known non-interface types are equal, or if only interface types
// are found (mostly issued from Isa()) and they are equal.
func uniqTypeBehindSlice(items []reflect.Value) reflect.Type {
	var (
		lastIfType, lastType, curType reflect.Type
		severalIfTypes                bool
	)

	for _, item := range items {
		if !item.IsValid() {
			return nil // no need to go further
		}

		if item.Type().Implements(testDeeper) {
			curType = item.Interface().(TestDeep).TypeBehind()

			// Ignore unknown TypeBehind
			if curType == nil {
				continue
			}

			// Ignore interfaces & interface pointers too (see Isa), but
			// keep them in mind in case we encounter always the same
			// interface pointer
			if curType.Kind() == reflect.Interface ||
				(curType.Kind() == reflect.Ptr &&
					curType.Elem().Kind() == reflect.Interface) {
				if lastIfType == nil {
					lastIfType = curType
				} else if lastIfType != curType {
					severalIfTypes = true
				}
				continue
			}
		} else {
			curType = item.Type()
		}

		if lastType != curType {
			if lastType != nil {
				return nil
			}
			lastType = curType
		}
	}

	// Only one type found
	if lastType != nil {
		return lastType
	}

	// Only one interface type found
	if lastIfType != nil && !severalIfTypes {
		return lastIfType
	}
	return nil
}
