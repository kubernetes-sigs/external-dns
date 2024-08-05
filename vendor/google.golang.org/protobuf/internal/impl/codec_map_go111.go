// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
<<<<<<< HEAD
//go:build !go1.12
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
//go:build !go1.12
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build !go1.12

package impl

import "reflect"

type mapIter struct {
	v    reflect.Value
	keys []reflect.Value
}

// mapRange provides a less-efficient equivalent to
// the Go 1.12 reflect.Value.MapRange method.
func mapRange(v reflect.Value) *mapIter {
	return &mapIter{v: v}
}

func (i *mapIter) Next() bool {
	if i.keys == nil {
		i.keys = i.v.MapKeys()
	} else {
		i.keys = i.keys[1:]
	}
	return len(i.keys) > 0
}

func (i *mapIter) Key() reflect.Value {
	return i.keys[0]
}

func (i *mapIter) Value() reflect.Value {
	return i.v.MapIndex(i.keys[0])
}
