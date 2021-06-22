// Copyright 2013-2015 Apcera Inc. All rights reserved.

/*
This is a GSSAPI provider for Go, which expects to be initialized with the name
of a dynamically loadable module which can be dlopen'd to get at a C language
binding GSSAPI library.

The GSSAPI concepts are explained in RFC 2743, "Generic Security Service
Application Program Interface Version 2, Update 1".

The API calls for C, together with a number of values for constants, come from
RFC 2744, "Generic Security Service API Version 2 : C-bindings".

Note that the basic GSSAPI bindings for C use the Latin-1 character set.  UTF-8
interfaces are specified in RFC 5178, "Generic Security Service Application
Program Interface (GSS-API) Internationalization and Domain-Based Service Names
and Name Type", in 2008.  Looking in 2013, this API does not appear to be
provided by either MIT or Heimdal.  This API applies solely to hostnames
though, which can also be supplied in ACE encoding, bypassing the issue.

For now, we assume that hostnames and usercodes are all ASCII-ish and pass
UTF-8 into the library.  Patches for more comprehensive support welcome.
*/
package gssapi
