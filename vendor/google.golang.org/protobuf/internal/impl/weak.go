// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"

<<<<<<< HEAD
<<<<<<< HEAD
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// weakFields adds methods to the exported WeakFields type for internal use.
//
// The exported type is an alias to an unnamed type, so methods can't be
// defined directly on it.
type weakFields WeakFields

func (w weakFields) get(num protoreflect.FieldNumber) (protoreflect.ProtoMessage, bool) {
	m, ok := w[int32(num)]
	return m, ok
}

func (w *weakFields) set(num protoreflect.FieldNumber, m protoreflect.ProtoMessage) {
	if *w == nil {
		*w = make(weakFields)
	}
	(*w)[int32(num)] = m
}

func (w *weakFields) clear(num protoreflect.FieldNumber) {
	delete(*w, int32(num))
}

func (Export) HasWeak(w WeakFields, num protoreflect.FieldNumber) bool {
	_, ok := w[int32(num)]
	return ok
}

func (Export) ClearWeak(w *WeakFields, num protoreflect.FieldNumber) {
	delete(*w, int32(num))
}

func (Export) GetWeak(w WeakFields, num protoreflect.FieldNumber, name protoreflect.FullName) protoreflect.ProtoMessage {
	if m, ok := w[int32(num)]; ok {
		return m
	}
	mt, _ := protoregistry.GlobalTypes.FindMessageByName(name)
	if mt == nil {
		panic(fmt.Sprintf("message %v for weak field is not linked in", name))
	}
	return mt.Zero().Interface()
}

func (Export) SetWeak(w *WeakFields, num protoreflect.FieldNumber, name protoreflect.FullName, m protoreflect.ProtoMessage) {
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	pref "google.golang.org/protobuf/reflect/protoreflect"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	pref "google.golang.org/protobuf/reflect/protoreflect"
=======
	"google.golang.org/protobuf/reflect/protoreflect"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"google.golang.org/protobuf/reflect/protoregistry"
)

// weakFields adds methods to the exported WeakFields type for internal use.
//
// The exported type is an alias to an unnamed type, so methods can't be
// defined directly on it.
type weakFields WeakFields

func (w weakFields) get(num protoreflect.FieldNumber) (protoreflect.ProtoMessage, bool) {
	m, ok := w[int32(num)]
	return m, ok
}

func (w *weakFields) set(num protoreflect.FieldNumber, m protoreflect.ProtoMessage) {
	if *w == nil {
		*w = make(weakFields)
	}
	(*w)[int32(num)] = m
}

func (w *weakFields) clear(num protoreflect.FieldNumber) {
	delete(*w, int32(num))
}

func (Export) HasWeak(w WeakFields, num protoreflect.FieldNumber) bool {
	_, ok := w[int32(num)]
	return ok
}

func (Export) ClearWeak(w *WeakFields, num protoreflect.FieldNumber) {
	delete(*w, int32(num))
}

func (Export) GetWeak(w WeakFields, num protoreflect.FieldNumber, name protoreflect.FullName) protoreflect.ProtoMessage {
	if m, ok := w[int32(num)]; ok {
		return m
	}
	mt, _ := protoregistry.GlobalTypes.FindMessageByName(name)
	if mt == nil {
		panic(fmt.Sprintf("message %v for weak field is not linked in", name))
	}
	return mt.Zero().Interface()
}

<<<<<<< HEAD
func (Export) SetWeak(w *WeakFields, num pref.FieldNumber, name pref.FullName, m pref.ProtoMessage) {
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func (Export) SetWeak(w *WeakFields, num pref.FieldNumber, name pref.FullName, m pref.ProtoMessage) {
=======
func (Export) SetWeak(w *WeakFields, num protoreflect.FieldNumber, name protoreflect.FullName, m protoreflect.ProtoMessage) {
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	if m != nil {
		mt, _ := protoregistry.GlobalTypes.FindMessageByName(name)
		if mt == nil {
			panic(fmt.Sprintf("message %v for weak field is not linked in", name))
		}
		if mt != m.ProtoReflect().Type() {
			panic(fmt.Sprintf("invalid message type for weak field: got %T, want %T", m, mt.Zero().Interface()))
		}
	}
	if m == nil || !m.ProtoReflect().IsValid() {
		delete(*w, int32(num))
		return
	}
	if *w == nil {
		*w = make(weakFields)
	}
	(*w)[int32(num)] = m
}
