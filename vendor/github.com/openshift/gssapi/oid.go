// Copyright 2013-2015 Apcera Inc. All rights reserved.

package gssapi

/*
#include <stdlib.h>
#include <string.h>

#include <gssapi/gssapi.h>

const size_t gss_OID_size=sizeof(gss_OID_desc);

void helper_gss_OID_desc_free_elements(gss_OID oid) {
	free(oid->elements);
}

void helper_gss_OID_desc_set_elements(gss_OID oid, OM_uint32 l, void *p) {
	oid->length = l;
	oid->elements = p;
}

void helper_gss_OID_desc_get_elements(gss_OID oid, OM_uint32 *l, char **p) {
	*l = oid->length;
	*p = oid->elements;
}

int
wrap_gss_oid_equal(void *fp, gss_OID oid1, gss_OID oid2)
{
	return ((int(*) (gss_OID, gss_OID)) fp)(oid1, oid2);
}

*/
import "C"

import (
	"bytes"
	"fmt"
	"unsafe"
)

// NewOID initializes a new OID. (Object Identifier)
func (lib *Lib) NewOID() *OID {
	return &OID{Lib: lib}
}

// MakeOIDBytes makes an OID encapsulating a byte slice. Note that it does not
// duplicate the data, but rather it points to it directly.
func (lib *Lib) MakeOIDBytes(data []byte) (*OID, error) {
	oid := lib.NewOID()

	s := C.malloc(C.gss_OID_size) // s for struct
	if s == nil {
		return nil, ErrMallocFailed
	}
	C.memset(s, 0, C.gss_OID_size)

	l := C.size_t(len(data))
	e := C.malloc(l) // c for contents
	if e == nil {
		return nil, ErrMallocFailed
	}
	C.memmove(e, (unsafe.Pointer)(&data[0]), l)

	oid.C_gss_OID = C.gss_OID(s)
	oid.alloc = allocMalloc

	// because of the alignment issues I can't access o.oid's fields from go,
	// so invoking a C function to do the same as:
	// oid.C_gss_OID.length = l
	// oid.C_gss_OID.elements = c
	C.helper_gss_OID_desc_set_elements(oid.C_gss_OID, C.OM_uint32(l), e)

	return oid, nil
}

// MakeOIDString makes an OID from a string.
func (lib *Lib) MakeOIDString(data string) (*OID, error) {
	return lib.MakeOIDBytes([]byte(data))
}

// Release safely frees the contents of an OID if it's allocated with malloc by
// MakeOIDBytes.
func (oid *OID) Release() error {
	if oid == nil || oid.C_gss_OID == nil {
		return nil
	}

	switch oid.alloc {
	case allocMalloc:
		// same as with get and set, use a C helper to free(oid.C_gss_OID.elements)
		C.helper_gss_OID_desc_free_elements(oid.C_gss_OID)
		C.free(unsafe.Pointer(oid.C_gss_OID))
		oid.C_gss_OID = nil
		oid.alloc = allocNone
	}

	return nil
}

// Bytes displays the bytes of an OID.
func (oid OID) Bytes() []byte {
	var l C.OM_uint32
	var p *C.char

	C.helper_gss_OID_desc_get_elements(oid.C_gss_OID, &l, &p)

	return C.GoBytes(unsafe.Pointer(p), C.int(l))
}

// String displays a string representation of an OID.
func (oid *OID) String() string {
	var l C.OM_uint32
	var p *C.char

	C.helper_gss_OID_desc_get_elements(oid.C_gss_OID, &l, &p)

	return fmt.Sprintf(`%x`, C.GoStringN(p, C.int(l)))
}

// Returns a symbolic name for a known OID, or the string. Note that this
// function is intended for debugging and is not at all performant.
func (oid *OID) DebugString() string {
	switch {
	case bytes.Equal(oid.Bytes(), oid.GSS_C_NT_USER_NAME.Bytes()):
		return "GSS_C_NT_USER_NAME"
	case bytes.Equal(oid.Bytes(), oid.GSS_C_NT_MACHINE_UID_NAME.Bytes()):
		return "GSS_C_NT_MACHINE_UID_NAME"
	case bytes.Equal(oid.Bytes(), oid.GSS_C_NT_STRING_UID_NAME.Bytes()):
		return "GSS_C_NT_STRING_UID_NAME"
	case bytes.Equal(oid.Bytes(), oid.GSS_C_NT_HOSTBASED_SERVICE_X.Bytes()):
		return "GSS_C_NT_HOSTBASED_SERVICE_X"
	case bytes.Equal(oid.Bytes(), oid.GSS_C_NT_HOSTBASED_SERVICE.Bytes()):
		return "GSS_C_NT_HOSTBASED_SERVICE"
	case bytes.Equal(oid.Bytes(), oid.GSS_C_NT_ANONYMOUS.Bytes()):
		return "GSS_C_NT_ANONYMOUS"
	case bytes.Equal(oid.Bytes(), oid.GSS_C_NT_EXPORT_NAME.Bytes()):
		return "GSS_C_NT_EXPORT_NAME"
	case bytes.Equal(oid.Bytes(), oid.GSS_KRB5_NT_PRINCIPAL_NAME.Bytes()):
		return "GSS_KRB5_NT_PRINCIPAL_NAME"
	case bytes.Equal(oid.Bytes(), oid.GSS_KRB5_NT_PRINCIPAL.Bytes()):
		return "GSS_KRB5_NT_PRINCIPAL"
	case bytes.Equal(oid.Bytes(), oid.GSS_MECH_KRB5.Bytes()):
		return "GSS_MECH_KRB5"
	case bytes.Equal(oid.Bytes(), oid.GSS_MECH_KRB5_LEGACY.Bytes()):
		return "GSS_MECH_KRB5_LEGACY"
	case bytes.Equal(oid.Bytes(), oid.GSS_MECH_KRB5_OLD.Bytes()):
		return "GSS_MECH_KRB5_OLD"
	case bytes.Equal(oid.Bytes(), oid.GSS_MECH_SPNEGO.Bytes()):
		return "GSS_MECH_SPNEGO"
	case bytes.Equal(oid.Bytes(), oid.GSS_MECH_IAKERB.Bytes()):
		return "GSS_MECH_IAKERB"
	case bytes.Equal(oid.Bytes(), oid.GSS_MECH_NTLMSSP.Bytes()):
		return "GSS_MECH_NTLMSSP"
	}

	return oid.String()
}
