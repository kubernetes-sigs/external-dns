// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package tdutil

import (
	"reflect"
	"sort"
)

// MapSortedKeys returns a slice of all sorted keys of map "m". It
// panics if "m"'s reflect.Kind is not reflect.Map.
func MapSortedKeys(m reflect.Value) []reflect.Value {
	ks := m.MapKeys()
	sort.Sort(SortableValues(ks))
	return ks
}
