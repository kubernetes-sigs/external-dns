// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
//go:build purego || appengine
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build purego appengine

package strs

import pref "google.golang.org/protobuf/reflect/protoreflect"

func UnsafeString(b []byte) string {
	return string(b)
}

func UnsafeBytes(s string) []byte {
	return []byte(s)
}

type Builder struct{}

func (*Builder) AppendFullName(prefix pref.FullName, name pref.Name) pref.FullName {
	return prefix.Append(name)
}

func (*Builder) MakeString(b []byte) string {
	return string(b)
}
