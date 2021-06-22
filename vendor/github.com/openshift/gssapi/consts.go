// Copyright 2013-2015 Apcera Inc. All rights reserved.

// A number of constants for C binding of GSSAPI.
//
// Unless otherwise stated, values come from RFC 2744 Appendix A.
//
// See also the GSS_S_* values in status.go, together with some related GSS_C_*
// values.

package gssapi

/*
#include <gssapi/gssapi.h>
*/
import "C"

import (
	"time"
)

// Flag bits for context-level services
const (
	GSS_C_DELEG_FLAG      uint32 = 1
	GSS_C_MUTUAL_FLAG            = 2
	GSS_C_REPLAY_FLAG            = 4
	GSS_C_SEQUENCE_FLAG          = 8
	GSS_C_CONF_FLAG              = 16
	GSS_C_INTEG_FLAG             = 32
	GSS_C_ANON_FLAG              = 64
	GSS_C_PROT_READY_FLAG        = 128
	GSS_C_TRANS_FLAG             = 256
)

// Credential usage options
const (
	GSS_C_BOTH     CredUsage = 0
	GSS_C_INITIATE           = 1
	GSS_C_ACCEPT             = 2
)

// Status code types for gss_display_status
const (
	GSS_C_GSS_CODE  int = 1
	GSS_C_MECH_CODE     = 2
)

// The constant definitions for channel-bindings address families
const (
	GSS_C_AF_UNSPEC    ChannelBindingAddressFamily = 0
	GSS_C_AF_LOCAL                                 = 1
	GSS_C_AF_INET                                  = 2
	GSS_C_AF_IMPLINK                               = 3
	GSS_C_AF_PUP                                   = 4
	GSS_C_AF_CHAOS                                 = 5
	GSS_C_AF_NS                                    = 6
	GSS_C_AF_NBS                                   = 7
	GSS_C_AF_ECMA                                  = 8
	GSS_C_AF_DATAKIT                               = 9
	GSS_C_AF_CCITT                                 = 10
	GSS_C_AF_SNA                                   = 11
	GSS_C_AF_DECnet                                = 12
	GSS_C_AF_DLI                                   = 13
	GSS_C_AF_LAT                                   = 14
	GSS_C_AF_HYLINK                                = 15
	GSS_C_AF_APPLETALK                             = 16
	GSS_C_AF_BSC                                   = 17
	GSS_C_AF_DSS                                   = 18
	GSS_C_AF_OSI                                   = 19
	GSS_C_AF_X25                                   = 21
	GSS_C_AF_INET6                                 = 24
	GSS_C_AF_NULLADDR                              = 255

	// Note: GSS_C_AF_INET6 is not in RFC2744 and not in MIT Kerberos.
	// The value here is from Heimdal.
	// Searching reveals that at IETF-64 the Kitten WG discussed the lack of
	// GSS_C_AF_INET6 and problems with standardising, but I can find no
	// further reference to standardising the value.
	// MIT does not have such a value, there are suggestions that GSS_C_AF_INET
	// is used instead.  If this CB value is actually used, interoperability
	// must be ... "limited".
	//
	// Fiat decision: adopt the Heimdal value.
)

const (
	// Quality Of Protection
	GSS_C_QOP_DEFAULT = 0

	// Infinite Lifetime, defined as 2^32-1
	GSS_C_INDEFINITE = 0xffffffff * time.Second
)
