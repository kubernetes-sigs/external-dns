// Copyright 2013 Apcera Inc. All rights reserved.

package gssapi

/*
#include <gssapi/gssapi.h>

OM_uint32
wrap_gss_acquire_cred(void *fp,
	OM_uint32 * minor_status,
	const gss_name_t desired_name,
	OM_uint32 time_req,
	const gss_OID_set desired_mechs,
	gss_cred_usage_t cred_usage,
	gss_cred_id_t * output_cred_handle,
	gss_OID_set * actual_mechs,
	OM_uint32 * time_rec)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_name_t,
		OM_uint32,
		const gss_OID_set,
		gss_cred_usage_t,
		gss_cred_id_t *,
		gss_OID_set *,
		OM_uint32 *)
	) fp)(
		minor_status,
		desired_name,
		time_req,
		desired_mechs,
		cred_usage,
		output_cred_handle,
		actual_mechs,
		time_rec);
}

OM_uint32
wrap_gss_add_cred(void *fp,
	OM_uint32 * minor_status,
	const gss_cred_id_t input_cred_handle,
	const gss_name_t desired_name,
	const gss_OID desired_mech,
	gss_cred_usage_t cred_usage,
	OM_uint32 initiator_time_req,
	OM_uint32 acceptor_time_req,
	gss_cred_id_t * output_cred_handle,
	gss_OID_set * actual_mechs,
	OM_uint32 * initiator_time_rec,
	OM_uint32 * acceptor_time_rec)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_cred_id_t,
		const gss_name_t,
		const gss_OID,
		gss_cred_usage_t,
		OM_uint32,
		OM_uint32,
		gss_cred_id_t *,
		gss_OID_set *,
		OM_uint32 *,
		OM_uint32 *)
	) fp)(
		minor_status,
		input_cred_handle,
		desired_name,
		desired_mech,
		cred_usage,
		initiator_time_req,
		acceptor_time_req,
		output_cred_handle,
		actual_mechs,
		initiator_time_rec,
		acceptor_time_rec);
}

OM_uint32
wrap_gss_inquire_cred (void *fp,
	OM_uint32           *minor_status,
	const gss_cred_id_t cred_handle,
	gss_name_t          *name,
	OM_uint32           *lifetime,
	gss_cred_usage_t    *cred_usage,
	gss_OID_set         *mechanisms )
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_cred_id_t,
		gss_name_t *,
		OM_uint32 *,
		gss_cred_usage_t *,
		gss_OID_set *)
	) fp)(
		minor_status,
		cred_handle,
		name,
		lifetime,
		cred_usage,
		mechanisms);
}

OM_uint32
wrap_gss_inquire_cred_by_mech (void *fp,
	OM_uint32           *minor_status,
	const gss_cred_id_t cred_handle,
	const gss_OID       mech_type,
	gss_name_t          *name,
	OM_uint32           *initiator_lifetime,
	OM_uint32           *acceptor_lifetime,
	gss_cred_usage_t    *cred_usage )
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		const gss_cred_id_t,
		const gss_OID,
		gss_name_t *,
		OM_uint32 *,
		OM_uint32 *,
		gss_cred_usage_t *)
	) fp)(
		minor_status,
		cred_handle,
		mech_type,
		name,
		initiator_lifetime,
		acceptor_lifetime,
		cred_usage);
}

OM_uint32
wrap_gss_release_cred(void *fp,
	OM_uint32 * minor_status,
	gss_cred_id_t * cred_handle)
{
	return ((OM_uint32(*) (
		OM_uint32 *,
		gss_cred_id_t *)
	) fp)(
		minor_status,
		cred_handle);
}

*/
import "C"

import (
	"time"
)

// NewCredId instantiates a new credential.
func (lib *Lib) NewCredId() *CredId {
	return &CredId{
		Lib: lib,
	}
}

// AcquireCred implements gss_acquire_cred API, as per
// https://tools.ietf.org/html/rfc2743#page-31. outputCredHandle, actualMechs
// must be .Release()-ed by the caller
func (lib *Lib) AcquireCred(desiredName *Name, timeReq time.Duration,
	desiredMechs *OIDSet, credUsage CredUsage) (outputCredHandle *CredId,
	actualMechs *OIDSet, timeRec time.Duration, err error) {

	min := C.OM_uint32(0)
	actualMechs = lib.NewOIDSet()
	outputCredHandle = lib.NewCredId()
	timerec := C.OM_uint32(0)

	maj := C.wrap_gss_acquire_cred(lib.Fp_gss_acquire_cred,
		&min,
		desiredName.C_gss_name_t,
		C.OM_uint32(timeReq.Seconds()),
		desiredMechs.C_gss_OID_set,
		C.gss_cred_usage_t(credUsage),
		&outputCredHandle.C_gss_cred_id_t,
		&actualMechs.C_gss_OID_set,
		&timerec)

	err = lib.stashLastStatus(maj, min)
	if err != nil {
		return nil, nil, 0, err
	}

	return outputCredHandle, actualMechs, time.Duration(timerec) * time.Second, nil
}

// AddCred implements gss_add_cred API, as per
// https://tools.ietf.org/html/rfc2743#page-36. outputCredHandle, actualMechs
// must be .Release()-ed by the caller
func (lib *Lib) AddCred(inputCredHandle *CredId,
	desiredName *Name, desiredMech *OID, credUsage CredUsage,
	initiatorTimeReq time.Duration, acceptorTimeReq time.Duration) (
	outputCredHandle *CredId, actualMechs *OIDSet,
	initiatorTimeRec time.Duration, acceptorTimeRec time.Duration,
	err error) {

	min := C.OM_uint32(0)
	actualMechs = lib.NewOIDSet()
	outputCredHandle = lib.NewCredId()
	initSeconds := C.OM_uint32(0)
	acceptSeconds := C.OM_uint32(0)

	maj := C.wrap_gss_add_cred(lib.Fp_gss_add_cred,
		&min,
		inputCredHandle.C_gss_cred_id_t,
		desiredName.C_gss_name_t,
		desiredMech.C_gss_OID,
		C.gss_cred_usage_t(credUsage),
		C.OM_uint32(initiatorTimeReq.Seconds()),
		C.OM_uint32(acceptorTimeReq.Seconds()),
		&outputCredHandle.C_gss_cred_id_t,
		&actualMechs.C_gss_OID_set,
		&initSeconds,
		&acceptSeconds)

	err = lib.stashLastStatus(maj, min)
	if err != nil {
		return nil, nil, 0, 0, err
	}

	return outputCredHandle,
		actualMechs,
		time.Duration(initSeconds) * time.Second,
		time.Duration(acceptSeconds) * time.Second,
		nil
}

// InquireCred implements gss_inquire_cred API, as per
// https://tools.ietf.org/html/rfc2743#page-34. name and mechanisms must be
// .Release()-ed by the caller
func (lib *Lib) InquireCred(credHandle *CredId) (
	name *Name, lifetime time.Duration, credUsage CredUsage, mechanisms *OIDSet,
	err error) {

	min := C.OM_uint32(0)
	name = lib.NewName()
	life := C.OM_uint32(0)
	credUsage = CredUsage(0)
	mechanisms = lib.NewOIDSet()

	maj := C.wrap_gss_inquire_cred(lib.Fp_gss_inquire_cred,
		&min,
		credHandle.C_gss_cred_id_t,
		&name.C_gss_name_t,
		&life,
		(*C.gss_cred_usage_t)(&credUsage),
		&mechanisms.C_gss_OID_set)
	err = lib.stashLastStatus(maj, min)
	if err != nil {
		return nil, 0, 0, nil, err
	}

	return name,
		time.Duration(life) * time.Second,
		credUsage,
		mechanisms,
		nil
}

// InquireCredByMech implements gss_inquire_cred_by_mech API, as per
// https://tools.ietf.org/html/rfc2743#page-39. name must be .Release()-ed by
// the caller
func (lib *Lib) InquireCredByMech(credHandle *CredId, mechType *OID) (
	name *Name, initiatorLifetime time.Duration, acceptorLifetime time.Duration,
	credUsage CredUsage, err error) {

	min := C.OM_uint32(0)
	name = lib.NewName()
	ilife := C.OM_uint32(0)
	alife := C.OM_uint32(0)
	credUsage = CredUsage(0)

	maj := C.wrap_gss_inquire_cred_by_mech(lib.Fp_gss_inquire_cred_by_mech,
		&min,
		credHandle.C_gss_cred_id_t,
		mechType.C_gss_OID,
		&name.C_gss_name_t,
		&ilife,
		&alife,
		(*C.gss_cred_usage_t)(&credUsage))
	err = lib.stashLastStatus(maj, min)
	if err != nil {
		return nil, 0, 0, 0, err
	}

	return name,
		time.Duration(ilife) * time.Second,
		time.Duration(alife) * time.Second,
		credUsage,
		nil
}

// Release frees a credential.
func (c *CredId) Release() error {
	if c == nil || c.C_gss_cred_id_t == nil {
		return nil
	}

	min := C.OM_uint32(0)
	maj := C.wrap_gss_release_cred(c.Fp_gss_release_cred,
		&min,
		&c.C_gss_cred_id_t)

	return c.stashLastStatus(maj, min)
}

//TODO: Test for AddCred with existing cred
