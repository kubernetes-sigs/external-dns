// Copyright 2013-2015 Apcera Inc. All rights reserved.

// Wrappers for the main gssapi types, all in one file for consistency.

package gssapi

/*
#include <gssapi/gssapi.h>
*/
import "C"

// Struct types. The structs themselves are allocated in Go and are therefore
// GCed, the contents may comes from C/gssapi calls, and therefore must be
// explicitly released.  Calling the Release method is safe on uninitialized
// objects, and nil pointers.

const (
	allocNone = iota
	allocMalloc
	allocGSSAPI
)

// A Buffer is an underlying C buffer represented in Golang. Must be .Release'd.
type Buffer struct {
	*Lib
	C_gss_buffer_t C.gss_buffer_t

	// indicates if the contents of the buffer must be released with
	// gss_release_buffer (allocGSSAPI) or free-ed (allocMalloc)
	alloc int
}

// A Name represents a binary string labeling a security principal. In the case
// of Kerberos, this could be a name like 'user@EXAMPLE.COM'.
type Name struct {
	*Lib
	C_gss_name_t C.gss_name_t
}

// An OID is the wrapper for gss_OID_desc type. IMPORTANT: In gssapi, OIDs are
// not released explicitly, only as part of an OIDSet. However we malloc the OID
// bytes ourselves, so need to free them. To keep it simple, assume that OIDs
// obtained from gssapi must be Release()-ed. It will be safely ignored on those
// allocated by gssapi
type OID struct {
	*Lib
	C_gss_OID C.gss_OID

	// indicates if the contents of the buffer must be released with
	// gss_release_buffer (allocGSSAPI) or free-ed (allocMalloc)
	alloc int
}

// An OIDSet is a set of OIDs.
type OIDSet struct {
	*Lib
	C_gss_OID_set C.gss_OID_set
}

// A CredId represents information like a cryptographic secret. In Kerberos,
// this likely represents a keytab.
type CredId struct {
	*Lib
	C_gss_cred_id_t C.gss_cred_id_t
}

// A CtxId represents a security context. Contexts maintain the state of one end
// of an authentication protocol.
type CtxId struct {
	*Lib
	C_gss_ctx_id_t C.gss_ctx_id_t
}

// Aliases for the simple types
type CredUsage C.gss_cred_usage_t // C.int
type ChannelBindingAddressFamily uint32
type QOP C.OM_uint32

// A struct pointer technically, but not really used yet, and it's a static,
// non-releaseable struct so an alias will suffice
type ChannelBindings C.gss_channel_bindings_t
