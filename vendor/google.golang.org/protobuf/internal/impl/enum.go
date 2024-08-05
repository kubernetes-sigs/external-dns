// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"reflect"

<<<<<<< HEAD
<<<<<<< HEAD
	"google.golang.org/protobuf/reflect/protoreflect"
)

type EnumInfo struct {
	GoReflectType reflect.Type // int32 kind
	Desc          protoreflect.EnumDescriptor
}

func (t *EnumInfo) New(n protoreflect.EnumNumber) protoreflect.Enum {
	return reflect.ValueOf(n).Convert(t.GoReflectType).Interface().(protoreflect.Enum)
}
func (t *EnumInfo) Descriptor() protoreflect.EnumDescriptor { return t.Desc }
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	pref "google.golang.org/protobuf/reflect/protoreflect"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	pref "google.golang.org/protobuf/reflect/protoreflect"
=======
	"google.golang.org/protobuf/reflect/protoreflect"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
)

type EnumInfo struct {
	GoReflectType reflect.Type // int32 kind
	Desc          protoreflect.EnumDescriptor
}

func (t *EnumInfo) New(n protoreflect.EnumNumber) protoreflect.Enum {
	return reflect.ValueOf(n).Convert(t.GoReflectType).Interface().(protoreflect.Enum)
}
<<<<<<< HEAD
func (t *EnumInfo) Descriptor() pref.EnumDescriptor { return t.Desc }
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func (t *EnumInfo) Descriptor() pref.EnumDescriptor { return t.Desc }
=======
func (t *EnumInfo) Descriptor() protoreflect.EnumDescriptor { return t.Desc }
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
