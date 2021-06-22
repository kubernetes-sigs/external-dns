// Copyright 2013-2015 Apcera Inc. All rights reserved.

// +build darwin linux freebsd

package gssapi

/*
#cgo linux LDFLAGS: -ldl
#cgo freebsd pkg-config: heimdal-gssapi

#include <gssapi/gssapi.h>
#include <dlfcn.h>
#include <stdlib.h>

// Name-Types.  These are standardized in the RFCs.  The library requires that
// a given name be usable for resolution, but it's typically a macro, there's
// no guarantee about the name exported from the library.  But since they're
// static, and well-defined, we can just define them ourselves.

// RFC2744-mandated values, mapping from as-near-as-possible to cut&paste
const gss_OID_desc *_GSS_C_NT_USER_NAME           = & (gss_OID_desc) { 10, "\x2a\x86\x48\x86\xf7\x12\x01\x02\x01\x01" };
const gss_OID_desc *_GSS_C_NT_MACHINE_UID_NAME    = & (gss_OID_desc) { 10, "\x2a\x86\x48\x86\xf7\x12\x01\x02\x01\x02" };
const gss_OID_desc *_GSS_C_NT_STRING_UID_NAME     = & (gss_OID_desc) { 10, "\x2a\x86\x48\x86\xf7\x12\x01\x02\x01\x03" };
const gss_OID_desc *_GSS_C_NT_HOSTBASED_SERVICE_X = & (gss_OID_desc) {  6, "\x2b\x06\x01\x05\x06\x02" };
const gss_OID_desc *_GSS_C_NT_HOSTBASED_SERVICE   = & (gss_OID_desc) { 10, "\x2a\x86\x48\x86\xf7\x12\x01\x02\x01\x04" };
const gss_OID_desc *_GSS_C_NT_ANONYMOUS           = & (gss_OID_desc) {  6, "\x2b\x06\x01\x05\x06\x03" };  // original had \01
const gss_OID_desc *_GSS_C_NT_EXPORT_NAME         = & (gss_OID_desc) {  6, "\x2b\x06\x01\x05\x06\x04" };

// from gssapi_krb5.h: This name form shall be represented by the Object
// Identifier {iso(1) member-body(2) United States(840) mit(113554) infosys(1)
// gssapi(2) krb5(2) krb5_name(1)}.  The recommended symbolic name for this
// type is "GSS_KRB5_NT_PRINCIPAL_NAME".
const gss_OID_desc *_GSS_KRB5_NT_PRINCIPAL_NAME   = & (gss_OID_desc) { 10, "\x2a\x86\x48\x86\xf7\x12\x01\x02\x02\x01" };

// { 1 2 840 113554 1 2 2 2 }
const gss_OID_desc *_GSS_KRB5_NT_PRINCIPAL         = & (gss_OID_desc) { 10, "\x2A\x86\x48\x86\xF7\x12\x01\x02\x02\x02" };

// known mech OIDs
const gss_OID_desc *_GSS_MECH_KRB5                 = & (gss_OID_desc) {  9, "\x2A\x86\x48\x86\xF7\x12\x01\x02\x02" };
const gss_OID_desc *_GSS_MECH_KRB5_LEGACY          = & (gss_OID_desc) {  9, "\x2A\x86\x48\x82\xF7\x12\x01\x02\x02" };
const gss_OID_desc *_GSS_MECH_KRB5_OLD             = & (gss_OID_desc) {  5, "\x2B\x05\x01\x05\x02" };
const gss_OID_desc *_GSS_MECH_SPNEGO               = & (gss_OID_desc) {  6, "\x2b\x06\x01\x05\x05\x02" };
const gss_OID_desc *_GSS_MECH_IAKERB               = & (gss_OID_desc) {  6, "\x2b\x06\x01\x05\x02\x05" };
const gss_OID_desc *_GSS_MECH_NTLMSSP              = & (gss_OID_desc) { 10, "\x2b\x06\x01\x04\x01\x82\x37\x02\x02\x0a" };

*/
import "C"

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

// Values for Options.LoadDefault
const (
	MIT = iota
	Heimdal
)

type Severity uint

// Values for Options.Log severity indices
const (
	Emerg = Severity(iota)
	Alert
	Crit
	Err
	Warn
	Notice
	Info
	Debug
	MaxSeverity
)

var severityNames = []string{
	"Emerg",
	"Alert",
	"Crit",
	"Err",
	"Warn",
	"Notice",
	"Info",
	"Debug",
}

// String returns the string name of a log Severity.
func (s Severity) String() string {
	if s >= MaxSeverity {
		return ""
	}
	return severityNames[s]
}

// Printer matches the log package, not fmt
type Printer interface {
	Print(a ...interface{})
}

// Options denote the options used to load a GSSAPI library. If a user supplies
// a LibPath, we use that. Otherwise, based upon the default and the current OS,
// we try to construct the library path.
type Options struct {
	LibPath     string
	Krb5Config  string
	Krb5Ktname  string
	LoadDefault int

	Printers []Printer `json:"-"`
}

// ftable fields will be initialized to the corresponding function pointers from
// the GSSAPI library. They must be of form Fp_function_name (Capital 'F' so
// that we can use reflect.
type ftable struct {
	// buffer.go
	Fp_gss_release_buffer unsafe.Pointer
	Fp_gss_import_name    unsafe.Pointer

	// context.go
	Fp_gss_init_sec_context      unsafe.Pointer
	Fp_gss_accept_sec_context    unsafe.Pointer
	Fp_gss_delete_sec_context    unsafe.Pointer
	Fp_gss_process_context_token unsafe.Pointer
	Fp_gss_context_time          unsafe.Pointer
	Fp_gss_inquire_context       unsafe.Pointer
	Fp_gss_wrap_size_limit       unsafe.Pointer
	Fp_gss_export_sec_context    unsafe.Pointer
	Fp_gss_import_sec_context    unsafe.Pointer

	// credential.go
	Fp_gss_acquire_cred         unsafe.Pointer
	Fp_gss_add_cred             unsafe.Pointer
	Fp_gss_inquire_cred         unsafe.Pointer
	Fp_gss_inquire_cred_by_mech unsafe.Pointer
	Fp_gss_release_cred         unsafe.Pointer

	// message.go
	Fp_gss_get_mic    unsafe.Pointer
	Fp_gss_verify_mic unsafe.Pointer
	Fp_gss_wrap       unsafe.Pointer
	Fp_gss_unwrap     unsafe.Pointer

	// misc.go
	Fp_gss_indicate_mechs unsafe.Pointer

	// name.go
	Fp_gss_canonicalize_name      unsafe.Pointer
	Fp_gss_compare_name           unsafe.Pointer
	Fp_gss_display_name           unsafe.Pointer
	Fp_gss_duplicate_name         unsafe.Pointer
	Fp_gss_export_name            unsafe.Pointer
	Fp_gss_inquire_mechs_for_name unsafe.Pointer
	Fp_gss_inquire_names_for_mech unsafe.Pointer
	Fp_gss_release_name           unsafe.Pointer

	// oid_set.go
	Fp_gss_create_empty_oid_set unsafe.Pointer
	Fp_gss_add_oid_set_member   unsafe.Pointer
	Fp_gss_release_oid_set      unsafe.Pointer
	Fp_gss_test_oid_set_member  unsafe.Pointer

	// status.go
	Fp_gss_display_status unsafe.Pointer

	// krb5_keytab.go -- where does this come from?
	// Fp_gsskrb5_register_acceptor_identity unsafe.Pointer
}

// constants are a number of constant initialized in initConstants.
type constants struct {
	GSS_C_NO_BUFFER     *Buffer
	GSS_C_NO_OID        *OID
	GSS_C_NO_OID_SET    *OIDSet
	GSS_C_NO_CONTEXT    *CtxId
	GSS_C_NO_CREDENTIAL *CredId

	// when adding new OID constants also need to update OID.DebugString
	GSS_C_NT_USER_NAME           *OID
	GSS_C_NT_MACHINE_UID_NAME    *OID
	GSS_C_NT_STRING_UID_NAME     *OID
	GSS_C_NT_HOSTBASED_SERVICE_X *OID
	GSS_C_NT_HOSTBASED_SERVICE   *OID
	GSS_C_NT_ANONYMOUS           *OID
	GSS_C_NT_EXPORT_NAME         *OID
	GSS_KRB5_NT_PRINCIPAL_NAME   *OID
	GSS_KRB5_NT_PRINCIPAL        *OID
	GSS_MECH_KRB5                *OID
	GSS_MECH_KRB5_LEGACY         *OID
	GSS_MECH_KRB5_OLD            *OID
	GSS_MECH_SPNEGO              *OID
	GSS_MECH_IAKERB              *OID
	GSS_MECH_NTLMSSP             *OID

	GSS_C_NO_CHANNEL_BINDINGS ChannelBindings // implicitly initialized as nil
}

// Lib encapsulates both the GSSAPI and the library dlopen()'d for it. The
// handle represents the dynamically-linked gssapi library handle.
type Lib struct {
	LastStatus *Error

	// Should contain a gssapi.Printer for each severity level to be
	// logged, up to gssapi.MaxSeverity items
	Printers []Printer

	handle unsafe.Pointer

	ftable
	constants
}

const (
	fpPrefix = "Fp_"
)

// Path returns the chosen gssapi library path that we're looking for.
func (o *Options) Path() string {
	switch {
	case o.LibPath != "":
		return o.LibPath

	case o.LoadDefault == MIT:
		return appendOSExt("libgssapi_krb5")

	case o.LoadDefault == Heimdal:
		return appendOSExt("libgssapi")
	}
	return ""
}

// Load attempts to load a dynamically-linked gssapi library from the path
// specified by the supplied Options.
func Load(o *Options) (*Lib, error) {
	if o == nil {
		o = &Options{}
	}

	// We get the error in a separate call, so we need to lock OS thread
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	lib := &Lib{
		Printers: o.Printers,
	}

	if o.Krb5Config != "" {
		err := os.Setenv("KRB5_CONFIG", o.Krb5Config)
		if err != nil {
			return nil, err
		}
	}

	if o.Krb5Ktname != "" {
		err := os.Setenv("KRB5_KTNAME", o.Krb5Ktname)
		if err != nil {
			return nil, err
		}
	}

	path := o.Path()
	lib.Debug(fmt.Sprintf("Loading %q", path))
	lib_cs := C.CString(path)
	defer C.free(unsafe.Pointer(lib_cs))

	// we don't use RTLD_FIRST, it might be the case that the GSSAPI lib
	// delegates symbols to other libs it links against (eg, Kerberos)
	lib.handle = C.dlopen(lib_cs, C.RTLD_NOW|C.RTLD_LOCAL)
	if lib.handle == nil {
		return nil, fmt.Errorf("%s", C.GoString(C.dlerror()))
	}

	err := lib.populateFunctions()
	if err != nil {
		lib.Unload()
		return nil, err
	}

	lib.initConstants()

	return lib, nil
}

// Unload closes the handle to the dynamically-linked gssapi library.
func (lib *Lib) Unload() error {
	if lib == nil || lib.handle == nil {
		return nil
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	i := C.dlclose(lib.handle)
	if i == -1 {
		return fmt.Errorf("%s", C.GoString(C.dlerror()))
	}

	lib.handle = nil
	return nil
}

func appendOSExt(path string) string {
	ext := ".so"
	if runtime.GOOS == "darwin" {
		ext = ".dylib"
	}
	if !strings.HasSuffix(path, ext) {
		path += ext
	}
	return path
}

// populateFunctions ranges over the library's ftable, initializing each
// function inside. Assumes that the caller executes runtime.LockOSThread.
func (lib *Lib) populateFunctions() error {
	libT := reflect.TypeOf(lib.ftable)
	functionsV := reflect.ValueOf(lib).Elem().FieldByName("ftable")

	n := libT.NumField()
	for i := 0; i < n; i++ {
		// Get the field name, and make sure it's an Fp_.
		f := libT.FieldByIndex([]int{i})

		if !strings.HasPrefix(f.Name, fpPrefix) {
			return fmt.Errorf(
				"Unexpected: field %q does not start with %q",
				f.Name, fpPrefix)
		}

		// Resolve the symbol.
		cfname := C.CString(f.Name[len(fpPrefix):])
		v := C.dlsym(lib.handle, cfname)
		C.free(unsafe.Pointer(cfname))
		if v == nil {
			return fmt.Errorf("%s", C.GoString(C.dlerror()))
		}

		// Save the value into the struct
		functionsV.FieldByIndex([]int{i}).SetPointer(v)
	}

	return nil
}

// initConstants sets the initial values of a library's set of 'constants'.
func (lib *Lib) initConstants() {
	lib.GSS_C_NO_BUFFER = &Buffer{
		Lib: lib,
		// C_gss_buffer_t: C.GSS_C_NO_BUFFER, already nil
		// alloc: allocNone, already 0
	}
	lib.GSS_C_NO_OID = lib.NewOID()
	lib.GSS_C_NO_OID_SET = lib.NewOIDSet()
	lib.GSS_C_NO_CONTEXT = lib.NewCtxId()
	lib.GSS_C_NO_CREDENTIAL = lib.NewCredId()

	lib.GSS_C_NT_USER_NAME = &OID{Lib: lib, C_gss_OID: C._GSS_C_NT_USER_NAME}
	lib.GSS_C_NT_MACHINE_UID_NAME = &OID{Lib: lib, C_gss_OID: C._GSS_C_NT_MACHINE_UID_NAME}
	lib.GSS_C_NT_STRING_UID_NAME = &OID{Lib: lib, C_gss_OID: C._GSS_C_NT_MACHINE_UID_NAME}
	lib.GSS_C_NT_HOSTBASED_SERVICE_X = &OID{Lib: lib, C_gss_OID: C._GSS_C_NT_HOSTBASED_SERVICE_X}
	lib.GSS_C_NT_HOSTBASED_SERVICE = &OID{Lib: lib, C_gss_OID: C._GSS_C_NT_HOSTBASED_SERVICE}
	lib.GSS_C_NT_ANONYMOUS = &OID{Lib: lib, C_gss_OID: C._GSS_C_NT_ANONYMOUS}
	lib.GSS_C_NT_EXPORT_NAME = &OID{Lib: lib, C_gss_OID: C._GSS_C_NT_EXPORT_NAME}

	lib.GSS_KRB5_NT_PRINCIPAL_NAME = &OID{Lib: lib, C_gss_OID: C._GSS_KRB5_NT_PRINCIPAL_NAME}
	lib.GSS_KRB5_NT_PRINCIPAL = &OID{Lib: lib, C_gss_OID: C._GSS_KRB5_NT_PRINCIPAL}

	lib.GSS_MECH_KRB5 = &OID{Lib: lib, C_gss_OID: C._GSS_MECH_KRB5}
	lib.GSS_MECH_KRB5_LEGACY = &OID{Lib: lib, C_gss_OID: C._GSS_MECH_KRB5_LEGACY}
	lib.GSS_MECH_KRB5_OLD = &OID{Lib: lib, C_gss_OID: C._GSS_MECH_KRB5_OLD}
	lib.GSS_MECH_SPNEGO = &OID{Lib: lib, C_gss_OID: C._GSS_MECH_SPNEGO}
	lib.GSS_MECH_IAKERB = &OID{Lib: lib, C_gss_OID: C._GSS_MECH_IAKERB}
	lib.GSS_MECH_NTLMSSP = &OID{Lib: lib, C_gss_OID: C._GSS_MECH_NTLMSSP}
}

// Print outputs a log line to the specified severity.
func (lib *Lib) Print(level Severity, a ...interface{}) {
	if lib == nil || lib.Printers == nil || level >= Severity(len(lib.Printers)) {
		return
	}
	lib.Printers[level].Print(a...)
}

func (lib *Lib) Emerg(a ...interface{})  { lib.Print(Emerg, a...) }
func (lib *Lib) Alert(a ...interface{})  { lib.Print(Alert, a...) }
func (lib *Lib) Crit(a ...interface{})   { lib.Print(Crit, a...) }
func (lib *Lib) Err(a ...interface{})    { lib.Print(Err, a...) }
func (lib *Lib) Warn(a ...interface{})   { lib.Print(Warn, a...) }
func (lib *Lib) Notice(a ...interface{}) { lib.Print(Notice, a...) }
func (lib *Lib) Info(a ...interface{})   { lib.Print(Info, a...) }
func (lib *Lib) Debug(a ...interface{})  { lib.Print(Debug, a...) }
