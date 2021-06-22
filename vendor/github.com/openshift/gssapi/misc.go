// Copyright 2013-2015 Apcera Inc. All rights reserved.

package gssapi

/*
#include <gssapi/gssapi.h>
#include <stdlib.h>

OM_uint32
wrap_gss_indicate_mechs(void *fp,
	OM_uint32 *minor_status,
	gss_OID_set * mech_set)
{
	gss_OID_set_desc *ms = NULL;
	OM_uint32 maj;
	maj = ((OM_uint32(*)(
		OM_uint32 *,
		gss_OID_set *))fp) (
			minor_status,
			mech_set);

	return maj;
}

*/
import "C"

// IndicateMechs implements the gss_Indicate_mechs call, according to https://tools.ietf.org/html/rfc2743#page-69.
// This returns an OIDSet of the Mechs supported on the current OS.
func (lib *Lib) IndicateMechs() (*OIDSet, error) {

	mechs := lib.NewOIDSet()

	var min C.OM_uint32
	maj := C.wrap_gss_indicate_mechs(
		lib.Fp_gss_indicate_mechs,
		&min,
		&mechs.C_gss_OID_set)
	err := lib.stashLastStatus(maj, min)
	if err != nil {
		return nil, err
	}

	return mechs, nil
}
