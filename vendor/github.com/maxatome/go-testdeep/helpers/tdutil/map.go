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

<<<<<<< HEAD
// MapSortedKeys returns a slice of all sorted keys of map m. It
// panics if m's [reflect.Kind] is not [reflect.Map].
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// MapSortedKeys returns a slice of all sorted keys of map "m". It
// panics if "m"'s reflect.Kind is not reflect.Map.
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
func MapSortedKeys(m reflect.Value) []reflect.Value {
	ks := m.MapKeys()
	sort.Sort(SortableValues(ks))
	return ks
}
