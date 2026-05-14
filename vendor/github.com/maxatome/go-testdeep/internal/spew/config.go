/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spew

// ConfigState houses the configuration options used by spew to format and
// display values.  There is a global instance, Config, that is used to control
// top-level Sdump functionality.
//
// The zero value for ConfigState provides no indentation.  You would typically
// want to set it to a space or a tab.
type ConfigState struct {
	// Indent specifies the string to use for each indentation level.  The
	// global config instance that all top-level functions use set this to a
	// single space by default.  If you would like more indentation, you might
	// set this to a tab with "\t" or perhaps two spaces with "  ".
	Indent string

	// MaxDepth controls the maximum number of levels to descend into nested
	// data structures.  The default, 0, means there is no limit.
	//
	// NOTE: Circular data structures are properly detected, so it is not
	// necessary to set this value unless you specifically want to limit deeply
	// nested data structures.
	MaxDepth int

	// DisableMethods specifies whether or not error and fmt.Stringer
	// interfaces are invoked for types that implement them.
	DisableMethods bool

	// DisablePointerMethods specifies whether or not to check for and invoke
	// error and fmt.Stringer interfaces on types which only accept a pointer
	// receiver when the current type is not a pointer.
	//
	// NOTE: This might be an unsafe action since calling one of these methods
	// with a pointer receiver could technically mutate the value, however,
	// in practice, types which choose to satisfy an error or fmt.Stringer
	// interface with a pointer receiver should not be mutating their state
	// inside these interface methods.  As a result, this option relies on
	// access to the unsafe package, so it will not have any effect when
	// running in environments without access to the unsafe package such as
	// Google App Engine or with the "safe" build tag specified.
	DisablePointerMethods bool

	// DisablePointerAddresses specifies whether to disable the printing of
	// pointer addresses. This is useful when diffing data structures in tests.
	DisablePointerAddresses bool

	// EnableCapacities specifies whether to enable the printing of capacities
	// for arrays, slices and channels.
	EnableCapacities bool
}

// Config is the active configuration of the top-level functions.
// The configuration can be changed by modifying the contents of spew.Config.
var Config = ConfigState{Indent: " "}
