// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Message is the top-level interface that all messages must implement.
// It provides access to a reflective view of a message.
// Any implementation of this interface may be used with all functions in the
// protobuf module that accept a Message, except where otherwise specified.
//
// This is the v2 interface definition for protobuf messages.
// The v1 interface definition is "github.com/golang/protobuf/proto".Message.
//
// To convert a v1 message to a v2 message,
// use "github.com/golang/protobuf/proto".MessageV2.
// To convert a v2 message to a v1 message,
// use "github.com/golang/protobuf/proto".MessageV1.
type Message = protoreflect.ProtoMessage

// Error matches all errors produced by packages in the protobuf module.
//
// That is, errors.Is(err, Error) reports whether an error is produced
// by this module.
var Error error

func init() {
	Error = errors.Error
}
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD

// MessageName returns the full name of m.
// If m is nil, it returns an empty string.
func MessageName(m Message) protoreflect.FullName {
	if m == nil {
		return ""
	}
	return m.ProtoReflect().Descriptor().FullName()
}
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======

// MessageName returns the full name of m.
// If m is nil, it returns an empty string.
func MessageName(m Message) protoreflect.FullName {
	if m == nil {
		return ""
	}
	return m.ProtoReflect().Descriptor().FullName()
}
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
=======

// MessageName returns the full name of m.
// If m is nil, it returns an empty string.
func MessageName(m Message) protoreflect.FullName {
	if m == nil {
		return ""
	}
	return m.ProtoReflect().Descriptor().FullName()
}
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
=======

// MessageName returns the full name of m.
// If m is nil, it returns an empty string.
func MessageName(m Message) protoreflect.FullName {
	if m == nil {
		return ""
	}
	return m.ProtoReflect().Descriptor().FullName()
}
>>>>>>> 4d7e5ad26 (update vendored files)
