// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD:vendor/google.golang.org/protobuf/reflect/protoreflect/value_unsafe_go120.go
<<<<<<< HEAD
<<<<<<< HEAD:vendor/google.golang.org/protobuf/reflect/protoreflect/value_unsafe.go
<<<<<<< HEAD
//go:build !purego && !appengine
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
// +build !purego,!appengine
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2):vendor/google.golang.org/protobuf/reflect/protoreflect/value_unsafe.go
// +build !purego,!appengine
=======
//go:build !purego && !appengine && !go1.21
// +build !purego,!appengine,!go1.21
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2):vendor/google.golang.org/protobuf/reflect/protoreflect/value_unsafe_go120.go
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
//go:build !purego && !appengine && !go1.21
// +build !purego,!appengine,!go1.21
=======
//go:build !go1.21
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)

||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0):vendor/google.golang.org/protobuf/reflect/protoreflect/value_unsafe_go120.go
//go:build !go1.21

=======
>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0):vendor/google.golang.org/protobuf/reflect/protoreflect/value_unsafe.go
package protoreflect

import (
	"unsafe"

	"google.golang.org/protobuf/internal/pragma"
)

type (
	ifaceHeader struct {
		_    [0]any // if interfaces have greater alignment than unsafe.Pointer, this will enforce it.
		Type unsafe.Pointer
		Data unsafe.Pointer
	}
)

var (
	nilType     = typeOf(nil)
	boolType    = typeOf(*new(bool))
	int32Type   = typeOf(*new(int32))
	int64Type   = typeOf(*new(int64))
	uint32Type  = typeOf(*new(uint32))
	uint64Type  = typeOf(*new(uint64))
	float32Type = typeOf(*new(float32))
	float64Type = typeOf(*new(float64))
	stringType  = typeOf(*new(string))
	bytesType   = typeOf(*new([]byte))
	enumType    = typeOf(*new(EnumNumber))
)

// typeOf returns a pointer to the Go type information.
// The pointer is comparable and equal if and only if the types are identical.
func typeOf(t any) unsafe.Pointer {
	return (*ifaceHeader)(unsafe.Pointer(&t)).Type
}

// value is a union where only one type can be represented at a time.
// The struct is 24B large on 64-bit systems and requires the minimum storage
// necessary to represent each possible type.
//
// The Go GC needs to be able to scan variables containing pointers.
// As such, pointers and non-pointers cannot be intermixed.
type value struct {
	pragma.DoNotCompare // 0B

	// typ stores the type of the value as a pointer to the Go type.
	typ unsafe.Pointer // 8B

	// ptr stores the data pointer for a String, Bytes, or interface value.
	ptr unsafe.Pointer // 8B

	// num stores a Bool, Int32, Int64, Uint32, Uint64, Float32, Float64, or
	// Enum value as a raw uint64.
	//
	// It is also used to store the length of a String or Bytes value;
	// the capacity is ignored.
	num uint64 // 8B
}

func valueOfString(v string) Value {
	return Value{typ: stringType, ptr: unsafe.Pointer(unsafe.StringData(v)), num: uint64(len(v))}
}
func valueOfBytes(v []byte) Value {
	return Value{typ: bytesType, ptr: unsafe.Pointer(unsafe.SliceData(v)), num: uint64(len(v))}
}
func valueOfIface(v any) Value {
	p := (*ifaceHeader)(unsafe.Pointer(&v))
	return Value{typ: p.Type, ptr: p.Data}
}

func (v Value) getString() string {
	return unsafe.String((*byte)(v.ptr), v.num)
}
func (v Value) getBytes() []byte {
	return unsafe.Slice((*byte)(v.ptr), v.num)
}
func (v Value) getIface() (x any) {
	*(*ifaceHeader)(unsafe.Pointer(&x)) = ifaceHeader{Type: v.typ, Data: v.ptr}
	return x
}
