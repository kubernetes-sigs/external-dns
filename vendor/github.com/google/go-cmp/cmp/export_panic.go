// Copyright 2017, The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
<<<<<<< HEAD
// license that can be found in the LICENSE file.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// license that can be found in the LICENSE.md file.
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)

// +build purego

package cmp

import "reflect"

const supportExporters = false

func retrieveUnexportedField(reflect.Value, reflect.StructField, bool) reflect.Value {
	panic("no support for forcibly accessing unexported fields")
}
