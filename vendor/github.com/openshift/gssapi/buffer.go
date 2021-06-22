// Copyright 2013-2015 Apcera Inc. All rights reserved.

package gssapi

/*
#include <stdlib.h>
#include <string.h>

#include <gssapi/gssapi.h>

const size_t gss_buffer_size=sizeof(gss_buffer_desc);

OM_uint32
wrap_gss_release_buffer(void *fp,
	OM_uint32 *minor_status,
	gss_buffer_t buf)
{
	return ((OM_uint32(*)(
		OM_uint32*,
		gss_buffer_t))fp) (minor_status, buf);
}

OM_uint32
wrap_gss_import_name(void *fp,
	OM_uint32 *minor_status,
	const gss_buffer_t input_name_buffer,
	const gss_OID input_name_type,
	gss_name_t *output_name)
{
	return ((OM_uint32(*)(
		OM_uint32 *,
		const gss_buffer_t,
		const gss_OID,
		gss_name_t *)) fp) (
			minor_status,
			input_name_buffer,
			input_name_type,
			output_name);
}

int
wrap_gss_buffer_equal(
	gss_buffer_t b1,
	gss_buffer_t b2)
{
	return
		b1 != NULL &&
		b2 != NULL &&
		b1->length == b2->length &&
		(memcmp(b1->value,b2->value,b1->length) == 0);
}

*/
import "C"

import (
	"errors"
	"unsafe"
)

// ErrMallocFailed is returned when the malloc call has failed.
var ErrMallocFailed = errors.New("malloc failed, out of memory?")

// MakeBuffer returns a Buffer with an empty malloc-ed gss_buffer_desc in it.
// The return value must be .Release()-ed
func (lib *Lib) MakeBuffer(alloc int) (*Buffer, error) {
	s := C.malloc(C.gss_buffer_size)
	if s == nil {
		return nil, ErrMallocFailed
	}
	C.memset(s, 0, C.gss_buffer_size)

	b := &Buffer{
		Lib:            lib,
		C_gss_buffer_t: C.gss_buffer_t(s),
		alloc:          alloc,
	}
	return b, nil
}

// MakeBufferBytes makes a Buffer encapsulating a byte slice.
func (lib *Lib) MakeBufferBytes(data []byte) (*Buffer, error) {
	if len(data) == 0 {
		return lib.GSS_C_NO_BUFFER, nil
	}

	// have to allocate the memory in C land and copy
	b, err := lib.MakeBuffer(allocMalloc)
	if err != nil {
		return nil, err
	}

	l := C.size_t(len(data))
	c := C.malloc(l)
	if b == nil {
		return nil, ErrMallocFailed
	}
	C.memmove(c, (unsafe.Pointer)(&data[0]), l)

	b.C_gss_buffer_t.length = l
	b.C_gss_buffer_t.value = c
	b.alloc = allocMalloc

	return b, nil
}

// MakeBufferString makes a Buffer encapsulating the contents of a string.
func (lib *Lib) MakeBufferString(content string) (*Buffer, error) {
	return lib.MakeBufferBytes([]byte(content))
}

// Release safely frees the contents of a Buffer.
func (b *Buffer) Release() error {
	if b == nil || b.C_gss_buffer_t == nil {
		return nil
	}

	defer func() {
		C.free(unsafe.Pointer(b.C_gss_buffer_t))
		b.C_gss_buffer_t = nil
		b.alloc = allocNone
	}()

	// free the value as needed
	switch {
	case b.C_gss_buffer_t.value == nil:
		// do nothing

	case b.alloc == allocMalloc:
		C.free(b.C_gss_buffer_t.value)

	case b.alloc == allocGSSAPI:
		var min C.OM_uint32
		maj := C.wrap_gss_release_buffer(b.Fp_gss_release_buffer, &min, b.C_gss_buffer_t)
		err := b.stashLastStatus(maj, min)
		if err != nil {
			return err
		}
	}

	return nil
}

// Length returns the number of bytes in the Buffer.
func (b *Buffer) Length() int {
	if b == nil || b.C_gss_buffer_t == nil || b.C_gss_buffer_t.length == 0 {
		return 0
	}
	return int(b.C_gss_buffer_t.length)
}

// Bytes returns the contents of a Buffer as a byte slice.
func (b *Buffer) Bytes() []byte {
	if b == nil || b.C_gss_buffer_t == nil || b.C_gss_buffer_t.length == 0 {
		return make([]byte, 0)
	}
	return C.GoBytes(b.C_gss_buffer_t.value, C.int(b.C_gss_buffer_t.length))
}

// String returns the contents of a Buffer as a string.
func (b *Buffer) String() string {
	if b == nil || b.C_gss_buffer_t == nil || b.C_gss_buffer_t.length == 0 {
		return ""
	}
	return C.GoStringN((*C.char)(b.C_gss_buffer_t.value), C.int(b.C_gss_buffer_t.length))
}

// Name converts a Buffer representing a name into a Name (internal opaque
// representation) using the specified nametype.
func (b Buffer) Name(nametype *OID) (*Name, error) {
	var min C.OM_uint32
	var result C.gss_name_t

	maj := C.wrap_gss_import_name(b.Fp_gss_import_name, &min,
		b.C_gss_buffer_t, nametype.C_gss_OID, &result)
	err := b.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	n := &Name{
		Lib:          b.Lib,
		C_gss_name_t: result,
	}
	return n, nil
}

// Equal determines if a Buffer receiver is equivalent to the supplied Buffer.
func (b *Buffer) Equal(other *Buffer) bool {
	isEqual := C.wrap_gss_buffer_equal(b.C_gss_buffer_t, other.C_gss_buffer_t)
	return isEqual != 0
}
