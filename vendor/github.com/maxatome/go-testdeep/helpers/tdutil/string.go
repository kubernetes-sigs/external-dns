// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package tdutil

import (
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

// FormatString formats s to a printable string, trying to enclose it
<<<<<<< HEAD
<<<<<<< HEAD
// in double-quotes or back-quotes and defaulting to using [SpewString].
func FormatString(s string) string {
	var unquotable, unbackquotable bool
	for _, chr := range s {
		if !unicode.IsPrint(chr) {
			if chr != '\n' {
				return SpewString(s)
			}
			unquotable = true
			if unbackquotable {
				break
			}
			continue
		}
		if chr == '"' {
			unquotable = true
			if unbackquotable {
				break
			}
		} else if chr == '`' {
			unbackquotable = true
			if unquotable {
				break
			}
		}
	}
	if unquotable {
		if unbackquotable {
			return SpewString(s)
		}
		return "`" + s + "`"
	}
	return `"` + s + `"`
}

// SpewString uses [spew.Sdump] to format val.
func SpewString(val any) string {
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// in "" or `` and defaulting to using SpewString.
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// in "" or `` and defaulting to using SpewString.
=======
// in double-quotes or back-quotes and defaulting to using [SpewString].
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func FormatString(s string) string {
	var unquotable, unbackquotable bool
	for _, chr := range s {
		if !unicode.IsPrint(chr) {
			if chr != '\n' {
				return SpewString(s)
			}
			unquotable = true
			if unbackquotable {
				break
			}
			continue
		}
		if chr == '"' {
			unquotable = true
			if unbackquotable {
				break
			}
		} else if chr == '`' {
			unbackquotable = true
			if unquotable {
				break
			}
		}
	}
	if unquotable {
		if unbackquotable {
			return SpewString(s)
		}
		return "`" + s + "`"
	}
	return `"` + s + `"`
}

<<<<<<< HEAD
// SpewString uses github.com/davecgh/go-spew/spew.Sdump() to format val.
func SpewString(val interface{}) string {
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// SpewString uses github.com/davecgh/go-spew/spew.Sdump() to format val.
func SpewString(val interface{}) string {
=======
// SpewString uses [spew.Sdump] to format val.
func SpewString(val any) string {
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	return strings.TrimRight(spew.Sdump(val), "\n")
}
