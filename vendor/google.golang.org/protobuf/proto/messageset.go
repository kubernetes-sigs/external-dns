// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/internal/flags"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"google.golang.org/protobuf/internal/order"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func (o MarshalOptions) sizeMessageSet(m protoreflect.Message) (size int) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += messageset.SizeField(fd.Number())
		size += protowire.SizeTag(messageset.FieldMessage)
		size += protowire.SizeBytes(o.size(v.Message()))
		return true
	})
	size += messageset.SizeUnknown(m.GetUnknown())
	return size
}

func (o MarshalOptions) marshalMessageSet(b []byte, m protoreflect.Message) ([]byte, error) {
	if !flags.ProtoLegacy {
		return b, errors.New("no support for message_set_wire_format")
	}
	fieldOrder := order.AnyFieldOrder
	if o.Deterministic {
		fieldOrder = order.NumberFieldOrder
	}
	var err error
	order.RangeFields(m, fieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		b, err = o.marshalMessageSetField(b, fd, v)
		return err == nil
	})
	if err != nil {
		return b, err
	}
	return messageset.AppendUnknown(b, m.GetUnknown())
}

func (o MarshalOptions) marshalMessageSetField(b []byte, fd protoreflect.FieldDescriptor, value protoreflect.Value) ([]byte, error) {
	b = messageset.AppendFieldStart(b, fd.Number())
	b = protowire.AppendTag(b, messageset.FieldMessage, protowire.BytesType)
	b = protowire.AppendVarint(b, uint64(o.Size(value.Message().Interface())))
	b, err := o.marshalMessage(b, value.Message())
	if err != nil {
		return b, err
	}
	b = messageset.AppendFieldEnd(b)
	return b, nil
}

func (o UnmarshalOptions) unmarshalMessageSet(b []byte, m protoreflect.Message) error {
	if !flags.ProtoLegacy {
		return errors.New("no support for message_set_wire_format")
	}
	return messageset.Unmarshal(b, false, func(num protowire.Number, v []byte) error {
		err := o.unmarshalMessageSetField(m, num, v)
		if err == errUnknown {
			unknown := m.GetUnknown()
			unknown = protowire.AppendTag(unknown, num, protowire.BytesType)
			unknown = protowire.AppendBytes(unknown, v)
			m.SetUnknown(unknown)
			return nil
		}
		return err
	})
}

func (o UnmarshalOptions) unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte) error {
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	"google.golang.org/protobuf/internal/order"
>>>>>>> 5ce8c7613 (update vendored files)
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func (o MarshalOptions) sizeMessageSet(m protoreflect.Message) (size int) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += messageset.SizeField(fd.Number())
		size += protowire.SizeTag(messageset.FieldMessage)
		size += protowire.SizeBytes(o.size(v.Message()))
		return true
	})
	size += messageset.SizeUnknown(m.GetUnknown())
	return size
}

func (o MarshalOptions) marshalMessageSet(b []byte, m protoreflect.Message) ([]byte, error) {
	if !flags.ProtoLegacy {
		return b, errors.New("no support for message_set_wire_format")
	}
	fieldOrder := order.AnyFieldOrder
	if o.Deterministic {
		fieldOrder = order.NumberFieldOrder
	}
	var err error
	order.RangeFields(m, fieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		b, err = o.marshalMessageSetField(b, fd, v)
		return err == nil
	})
	if err != nil {
		return b, err
	}
	return messageset.AppendUnknown(b, m.GetUnknown())
}

func (o MarshalOptions) marshalMessageSetField(b []byte, fd protoreflect.FieldDescriptor, value protoreflect.Value) ([]byte, error) {
	b = messageset.AppendFieldStart(b, fd.Number())
	b = protowire.AppendTag(b, messageset.FieldMessage, protowire.BytesType)
	b = protowire.AppendVarint(b, uint64(o.Size(value.Message().Interface())))
	b, err := o.marshalMessage(b, value.Message())
	if err != nil {
		return b, err
	}
	b = messageset.AppendFieldEnd(b)
	return b, nil
}

func (o UnmarshalOptions) unmarshalMessageSet(b []byte, m protoreflect.Message) error {
	if !flags.ProtoLegacy {
		return errors.New("no support for message_set_wire_format")
	}
	return messageset.Unmarshal(b, false, func(num protowire.Number, v []byte) error {
		err := o.unmarshalMessageSetField(m, num, v)
		if err == errUnknown {
			unknown := m.GetUnknown()
			unknown = protowire.AppendTag(unknown, num, protowire.BytesType)
			unknown = protowire.AppendBytes(unknown, v)
			m.SetUnknown(unknown)
			return nil
		}
		return err
	})
}

<<<<<<< HEAD
func unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte, o UnmarshalOptions) error {
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
func unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte, o UnmarshalOptions) error {
=======
func (o UnmarshalOptions) unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte) error {
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	"google.golang.org/protobuf/internal/order"
>>>>>>> 6b7ce455e (update vendored files)
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func (o MarshalOptions) sizeMessageSet(m protoreflect.Message) (size int) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += messageset.SizeField(fd.Number())
		size += protowire.SizeTag(messageset.FieldMessage)
		size += protowire.SizeBytes(o.size(v.Message()))
		return true
	})
	size += messageset.SizeUnknown(m.GetUnknown())
	return size
}

func (o MarshalOptions) marshalMessageSet(b []byte, m protoreflect.Message) ([]byte, error) {
	if !flags.ProtoLegacy {
		return b, errors.New("no support for message_set_wire_format")
	}
	fieldOrder := order.AnyFieldOrder
	if o.Deterministic {
		fieldOrder = order.NumberFieldOrder
	}
	var err error
	order.RangeFields(m, fieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		b, err = o.marshalMessageSetField(b, fd, v)
		return err == nil
	})
	if err != nil {
		return b, err
	}
	return messageset.AppendUnknown(b, m.GetUnknown())
}

func (o MarshalOptions) marshalMessageSetField(b []byte, fd protoreflect.FieldDescriptor, value protoreflect.Value) ([]byte, error) {
	b = messageset.AppendFieldStart(b, fd.Number())
	b = protowire.AppendTag(b, messageset.FieldMessage, protowire.BytesType)
	b = protowire.AppendVarint(b, uint64(o.Size(value.Message().Interface())))
	b, err := o.marshalMessage(b, value.Message())
	if err != nil {
		return b, err
	}
	b = messageset.AppendFieldEnd(b)
	return b, nil
}

func (o UnmarshalOptions) unmarshalMessageSet(b []byte, m protoreflect.Message) error {
	if !flags.ProtoLegacy {
		return errors.New("no support for message_set_wire_format")
	}
	return messageset.Unmarshal(b, false, func(num protowire.Number, v []byte) error {
		err := o.unmarshalMessageSetField(m, num, v)
		if err == errUnknown {
			unknown := m.GetUnknown()
			unknown = protowire.AppendTag(unknown, num, protowire.BytesType)
			unknown = protowire.AppendBytes(unknown, v)
			m.SetUnknown(unknown)
			return nil
		}
		return err
	})
}

<<<<<<< HEAD
func unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte, o UnmarshalOptions) error {
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
func unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte, o UnmarshalOptions) error {
=======
func (o UnmarshalOptions) unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte) error {
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"google.golang.org/protobuf/internal/order"
>>>>>>> 4d7e5ad26 (update vendored files)
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func (o MarshalOptions) sizeMessageSet(m protoreflect.Message) (size int) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += messageset.SizeField(fd.Number())
		size += protowire.SizeTag(messageset.FieldMessage)
		size += protowire.SizeBytes(o.size(v.Message()))
		return true
	})
	size += messageset.SizeUnknown(m.GetUnknown())
	return size
}

func (o MarshalOptions) marshalMessageSet(b []byte, m protoreflect.Message) ([]byte, error) {
	if !flags.ProtoLegacy {
		return b, errors.New("no support for message_set_wire_format")
	}
	fieldOrder := order.AnyFieldOrder
	if o.Deterministic {
		fieldOrder = order.NumberFieldOrder
	}
	var err error
	order.RangeFields(m, fieldOrder, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		b, err = o.marshalMessageSetField(b, fd, v)
		return err == nil
	})
	if err != nil {
		return b, err
	}
	return messageset.AppendUnknown(b, m.GetUnknown())
}

func (o MarshalOptions) marshalMessageSetField(b []byte, fd protoreflect.FieldDescriptor, value protoreflect.Value) ([]byte, error) {
	b = messageset.AppendFieldStart(b, fd.Number())
	b = protowire.AppendTag(b, messageset.FieldMessage, protowire.BytesType)
	b = protowire.AppendVarint(b, uint64(o.Size(value.Message().Interface())))
	b, err := o.marshalMessage(b, value.Message())
	if err != nil {
		return b, err
	}
	b = messageset.AppendFieldEnd(b)
	return b, nil
}

func (o UnmarshalOptions) unmarshalMessageSet(b []byte, m protoreflect.Message) error {
	if !flags.ProtoLegacy {
		return errors.New("no support for message_set_wire_format")
	}
	return messageset.Unmarshal(b, false, func(num protowire.Number, v []byte) error {
		err := o.unmarshalMessageSetField(m, num, v)
		if err == errUnknown {
			unknown := m.GetUnknown()
			unknown = protowire.AppendTag(unknown, num, protowire.BytesType)
			unknown = protowire.AppendBytes(unknown, v)
			m.SetUnknown(unknown)
			return nil
		}
		return err
	})
}

<<<<<<< HEAD
func unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte, o UnmarshalOptions) error {
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
func unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte, o UnmarshalOptions) error {
=======
func (o UnmarshalOptions) unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte) error {
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func sizeMessageSet(m protoreflect.Message) (size int) {
	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		size += messageset.SizeField(fd.Number())
		size += protowire.SizeTag(messageset.FieldMessage)
		size += protowire.SizeBytes(sizeMessage(v.Message()))
		return true
	})
	size += messageset.SizeUnknown(m.GetUnknown())
	return size
}

func marshalMessageSet(b []byte, m protoreflect.Message, o MarshalOptions) ([]byte, error) {
	if !flags.ProtoLegacy {
		return b, errors.New("no support for message_set_wire_format")
	}
	var err error
	o.rangeFields(m, func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		b, err = marshalMessageSetField(b, fd, v, o)
		return err == nil
	})
	if err != nil {
		return b, err
	}
	return messageset.AppendUnknown(b, m.GetUnknown())
}

func marshalMessageSetField(b []byte, fd protoreflect.FieldDescriptor, value protoreflect.Value, o MarshalOptions) ([]byte, error) {
	b = messageset.AppendFieldStart(b, fd.Number())
	b = protowire.AppendTag(b, messageset.FieldMessage, protowire.BytesType)
	b = protowire.AppendVarint(b, uint64(o.Size(value.Message().Interface())))
	b, err := o.marshalMessage(b, value.Message())
	if err != nil {
		return b, err
	}
	b = messageset.AppendFieldEnd(b)
	return b, nil
}

func unmarshalMessageSet(b []byte, m protoreflect.Message, o UnmarshalOptions) error {
	if !flags.ProtoLegacy {
		return errors.New("no support for message_set_wire_format")
	}
	return messageset.Unmarshal(b, false, func(num protowire.Number, v []byte) error {
		err := unmarshalMessageSetField(m, num, v, o)
		if err == errUnknown {
			unknown := m.GetUnknown()
			unknown = protowire.AppendTag(unknown, num, protowire.BytesType)
			unknown = protowire.AppendBytes(unknown, v)
			m.SetUnknown(unknown)
			return nil
		}
		return err
	})
}

func unmarshalMessageSetField(m protoreflect.Message, num protowire.Number, v []byte, o UnmarshalOptions) error {
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	md := m.Descriptor()
	if !md.ExtensionRanges().Has(num) {
		return errUnknown
	}
	xt, err := o.Resolver.FindExtensionByNumber(md.FullName(), num)
	if err == protoregistry.NotFound {
		return errUnknown
	}
	if err != nil {
		return errors.New("%v: unable to resolve extension %v: %v", md.FullName(), num, err)
	}
	xd := xt.TypeDescriptor()
	if err := o.unmarshalMessage(v, m.Mutable(xd).Message()); err != nil {
		return err
	}
	return nil
}
