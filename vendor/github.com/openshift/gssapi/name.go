// Copyright 2013-2015 Apcera Inc. All rights reserved.

package gssapi

// Side-note: gss_const_name_t is defined in RFC5587 as a bug-fix over RFC2744,
// since "const gss_name_t foo" says that the foo pointer is const, not the item
// pointed to is const.  Ideally, we'd be able to detect that, or have a macro
// which indicates availability of the 5587 extensions.  Instead, we're stuck with
// the ancient system GSSAPI headers on MacOS not supporting this.
//
// Choosing between "correctness" on the target platform and losing that for others,
// I've chosen to pull in /opt/local/include for MacPorts on MacOS; that should get
// us a functioning type; it's a pointer, at the ABI level the typing doesn't matter,
// so once we compile we're good.  If modern (correct) headers are available in other
// locations, just add them to the search path for the relevant OS below.
//
// Using "MacPorts" on MacOS gives us: -I/opt/local/include
// Using "brew" on MacOS gives us: -I/usr/local/opt/heimdal/include

/*
#cgo darwin CFLAGS: -I/opt/local/include -I/usr/local/opt/heimdal/include
#include <stdio.h>

#include <gssapi/gssapi.h>

OM_uint32
wrap_gss_display_name(void *fp,
	OM_uint32 *minor_status,
	const gss_name_t input_name,
	gss_buffer_t output_name_buffer,
	gss_OID *output_name_type)
{
	return ((OM_uint32(*)(
		OM_uint32 *, const gss_name_t, gss_buffer_t, gss_OID *)
	)fp)(
		minor_status, input_name, output_name_buffer, output_name_type);
}

OM_uint32
wrap_gss_compare_name(void *fp,
	OM_uint32 *minor_status,
	const gss_name_t name1,
	const gss_name_t name2,
	int * name_equal)
{
	return ((OM_uint32(*)(
		OM_uint32 *, const gss_name_t, const gss_name_t, int *)
	)fp)(
		minor_status, name1, name2, name_equal);
}

OM_uint32
wrap_gss_release_name(void *fp,
	OM_uint32 *minor_status,
	gss_name_t *input_name)
{
	return ((OM_uint32(*)(
		OM_uint32 *, gss_name_t *)
	)fp)(
		minor_status, input_name);
}

OM_uint32
wrap_gss_inquire_mechs_for_name(void *fp,
	OM_uint32 *minor_status,
	const gss_name_t input_name,
	gss_OID_set *mech_types)
{
	return ((OM_uint32(*)(
		OM_uint32 *, const gss_name_t, gss_OID_set *)
	)fp)(
		minor_status, input_name, mech_types);
}

OM_uint32
wrap_gss_inquire_names_for_mech(void *fp,
	OM_uint32 *minor_status,
	const gss_OID mechanism,
	gss_OID_set * name_types)
{
	return ((OM_uint32(*)(
		OM_uint32 *, const gss_OID, gss_OID_set *)
	)fp)(
		minor_status, mechanism, name_types);
}

OM_uint32
wrap_gss_canonicalize_name(void *fp,
	OM_uint32 *minor_status,
	gss_const_name_t input_name,
	const gss_OID mech_type,
	gss_name_t *output_name)
{
	return ((OM_uint32(*)(
		OM_uint32 *, gss_const_name_t, const gss_OID, gss_name_t *)
	)fp)(
		minor_status, input_name, mech_type, output_name);
}

OM_uint32
wrap_gss_export_name(void *fp,
	OM_uint32 *minor_status,
	const gss_name_t input_name,
	gss_buffer_t exported_name)
{
	OM_uint32 maj;

	maj = ((OM_uint32(*)(
		OM_uint32 *, const gss_name_t, gss_buffer_t)
	)fp)(
		minor_status, input_name, exported_name);

	return maj;
}

OM_uint32
wrap_gss_duplicate_name(void *fp,
	OM_uint32 *minor_status,
	const gss_name_t src_name,
	gss_name_t *dest_name)
{
	return ((OM_uint32(*)(
		OM_uint32 *, const gss_name_t, gss_name_t *)
	)fp)(
		minor_status, src_name, dest_name);
}

*/
import "C"

// NewName initializes a new principal name.
func (lib *Lib) NewName() *Name {
	return &Name{
		Lib: lib,
	}
}

// GSS_C_NO_NAME is a Name where the value is NULL, used to request special
// behavior in some GSSAPI calls.
func (lib *Lib) GSS_C_NO_NAME() *Name {
	return lib.NewName()
}

// Release frees the memory associated with an internal representation of the
// name.
func (n *Name) Release() error {
	if n == nil || n.C_gss_name_t == nil {
		return nil
	}

	var min C.OM_uint32
	maj := C.wrap_gss_release_name(n.Fp_gss_release_name, &min, &n.C_gss_name_t)
	err := n.stashLastStatus(maj, min)
	if err == nil {
		n.C_gss_name_t = nil
	}
	return err
}

// Equal tests 2 names for semantic equality (refer to the same entity)
func (n Name) Equal(other Name) (equal bool, err error) {
	var min C.OM_uint32
	var isEqual C.int

	maj := C.wrap_gss_compare_name(n.Fp_gss_compare_name, &min,
		n.C_gss_name_t, other.C_gss_name_t, &isEqual)
	err = n.stashLastStatus(maj, min)
	if err != nil {
		return false, err
	}

	return isEqual != 0, nil
}

// Display "allows an application to obtain a textual representation of an
// opaque internal-form name for display purposes"
func (n Name) Display() (name string, oid *OID, err error) {
	var min C.OM_uint32
	b, err := n.MakeBuffer(allocGSSAPI)
	if err != nil {
		return "", nil, err
	}
	defer b.Release()

	oid = n.NewOID()

	maj := C.wrap_gss_display_name(n.Fp_gss_display_name, &min,
		n.C_gss_name_t, b.C_gss_buffer_t, &oid.C_gss_OID)

	err = n.stashLastStatus(maj, min)
	if err != nil {
		oid.Release()
		return "", nil, err
	}

	return b.String(), oid, err
}

// String displays a Go-friendly version of a name. ("" on error)
func (n Name) String() string {
	s, _, _ := n.Display()
	return s
}

// Canonicalize returns a copy of this name, canonicalized for the specified
// mechanism
func (n Name) Canonicalize(mech_type *OID) (canonical *Name, err error) {
	canonical = &Name{
		Lib: n.Lib,
	}

	var min C.OM_uint32
	maj := C.wrap_gss_canonicalize_name(n.Fp_gss_canonicalize_name, &min,
		n.C_gss_name_t, mech_type.C_gss_OID, &canonical.C_gss_name_t)
	err = n.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	return canonical, nil
}

// Duplicate creates a new independent imported name; after this, both the original and
// the duplicate will need to be .Released().
func (n *Name) Duplicate() (duplicate *Name, err error) {
	duplicate = &Name{
		Lib: n.Lib,
	}

	var min C.OM_uint32
	maj := C.wrap_gss_duplicate_name(n.Fp_gss_duplicate_name, &min,
		n.C_gss_name_t, &duplicate.C_gss_name_t)
	err = n.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	return duplicate, nil
}

// Export makes a text (Buffer) version from an internal representation
func (n *Name) Export() (b *Buffer, err error) {
	b, err = n.MakeBuffer(allocGSSAPI)
	if err != nil {
		return nil, err
	}

	var min C.OM_uint32
	maj := C.wrap_gss_export_name(n.Fp_gss_export_name, &min,
		n.C_gss_name_t, b.C_gss_buffer_t)
	err = n.stashLastStatus(maj, min)
	if err != nil {
		b.Release()
		return nil, err
	}

	return b, nil
}

// InquireMechs returns the set of mechanisms supported by the GSS-API
// implementation that may be able to process the specified name
func (n *Name) InquireMechs() (oids *OIDSet, err error) {
	oidset := n.NewOIDSet()
	if err != nil {
		return nil, err
	}

	var min C.OM_uint32
	maj := C.wrap_gss_inquire_mechs_for_name(n.Fp_gss_inquire_mechs_for_name, &min,
		n.C_gss_name_t, &oidset.C_gss_OID_set)
	err = n.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	return oidset, nil
}

// InquireNameForMech returns the set of name types supported by
// the specified mechanism
func (lib *Lib) InquireNamesForMechs(mech *OID) (name_types *OIDSet, err error) {
	oidset := lib.NewOIDSet()
	if err != nil {
		return nil, err
	}

	var min C.OM_uint32
	maj := C.wrap_gss_inquire_names_for_mech(lib.Fp_gss_inquire_mechs_for_name, &min,
		mech.C_gss_OID, &oidset.C_gss_OID_set)
	err = lib.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	return oidset, nil
}
