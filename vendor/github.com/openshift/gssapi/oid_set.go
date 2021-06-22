// Copyright 2013-2015 Apcera Inc. All rights reserved.

package gssapi

/*
#include <gssapi/gssapi.h>

OM_uint32
wrap_gss_create_empty_oid_set(void *fp,
	OM_uint32 *minor_status,
	gss_OID_set * set)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		gss_OID_set *)) fp)(
			minor_status,
			set);
}

OM_uint32
wrap_gss_release_oid_set(void *fp,
	OM_uint32 *minor_status,
	gss_OID_set * set)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		gss_OID_set *)) fp)(
			minor_status, set);
}

OM_uint32
wrap_gss_add_oid_set_member(void *fp,
	OM_uint32 *minor_status,
	const gss_OID member_oid,
	gss_OID_set * set)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_OID,
		gss_OID_set *)) fp)(
			minor_status, member_oid, set);
}

OM_uint32
wrap_gss_test_oid_set_member(void *fp,
	OM_uint32 *minor_status,
	const gss_OID member_oid,
	const gss_OID_set set,
	int * present)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_OID,
		const gss_OID_set,
		int *)) fp)(
			minor_status, member_oid, set, present);
}

gss_OID
get_oid_set_member(
	gss_OID_set set,
	int index)
{
	return &(set->elements[index]);
}

*/
import "C"

import (
	"fmt"
	"strings"
)

// NewOIDSet constructs a new empty OID set.
func (lib *Lib) NewOIDSet() *OIDSet {
	return &OIDSet{
		Lib: lib,
		// C_gss_OID_set: (C.gss_OID_set)(unsafe.Pointer(nil)),
	}
}

// MakeOIDSet makes an OIDSet prepopulated with the given OIDs.
func (lib *Lib) MakeOIDSet(oids ...*OID) (s *OIDSet, err error) {
	s = &OIDSet{
		Lib: lib,
	}

	var min C.OM_uint32
	maj := C.wrap_gss_create_empty_oid_set(s.Fp_gss_create_empty_oid_set,
		&min, &s.C_gss_OID_set)
	err = s.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	err = s.Add(oids...)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Release frees all C memory associated with an OIDSet.
func (s *OIDSet) Release() (err error) {
	if s == nil || s.C_gss_OID_set == nil {
		return nil
	}

	var min C.OM_uint32
	maj := C.wrap_gss_release_oid_set(s.Fp_gss_release_oid_set, &min, &s.C_gss_OID_set)
	return s.stashLastStatus(maj, min)
}

// Add adds OIDs to an OIDSet.
func (s *OIDSet) Add(oids ...*OID) (err error) {
	var min C.OM_uint32
	for _, oid := range oids {
		maj := C.wrap_gss_add_oid_set_member(s.Fp_gss_add_oid_set_member,
			&min, oid.C_gss_OID, &s.C_gss_OID_set)
		err = s.stashLastStatus(maj, min)
		if err != nil {
			return err
		}
	}

	return nil
}

// TestOIDSetMember a wrapper to determine if an OIDSet contains an OID.
func (s *OIDSet) TestOIDSetMember(oid *OID) (contains bool, err error) {
	var min C.OM_uint32
	var isPresent C.int

	maj := C.wrap_gss_test_oid_set_member(s.Fp_gss_test_oid_set_member,
		&min, oid.C_gss_OID, s.C_gss_OID_set, &isPresent)
	err = s.stashLastStatus(maj, min)
	if err != nil {
		return false, err
	}

	return isPresent != 0, nil
}

// Contains (gss_test_oid_set_member) checks if an OID is present OIDSet.
func (s *OIDSet) Contains(oid *OID) bool {
	contains, _ := s.TestOIDSetMember(oid)
	return contains
}

// Length returns the number of OIDs in a set.
func (s *OIDSet) Length() int {
	if s == nil {
		return 0
	}
	return int(s.C_gss_OID_set.count)
}

// Get returns a specific OID from the set. The memory will be released when the
// set itself is released.
func (s *OIDSet) Get(index int) (*OID, error) {
	if s == nil || index < 0 || index >= int(s.C_gss_OID_set.count) {
		return nil, fmt.Errorf("index %d out of bounds", index)
	}
	oid := s.NewOID()
	oid.C_gss_OID = C.get_oid_set_member(s.C_gss_OID_set, C.int(index))
	return oid, nil
}

func (s *OIDSet) DebugString() string {
	names := make([]string, 0)
	for i := 0; i < s.Length(); i++ {
		oid, _ := s.Get(i)
		names = append(names, oid.DebugString())
	}

	return "[" + strings.Join(names, ", ") + "]"
}
