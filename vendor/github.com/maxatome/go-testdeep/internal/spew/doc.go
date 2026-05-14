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

/*
Package spew implements a deep pretty printer for Go data structures to aid in
debugging.

Originally copied from https://github.com/davecgh/go-spew it has been
patched to fit go-testdeep needs as the original repository seems to
have no activity anymore.

A quick overview of the additional features spew provides over the built-in
printing facilities for Go data types are as follows:

  - Pointers are dereferenced and followed
  - Circular data structures are detected and handled properly
  - Custom fmt.Stringer/error interfaces are optionally invoked, including
    on unexported types
  - Custom types which only implement the fmt.Stringer/error interfaces via
    a pointer receiver are optionally invoked when passing non-pointer
    variables
  - Byte arrays and slices are dumped like the hexdump -C command which
    includes offsets, byte values in hex, and ASCII output

spew dumps Go data structures using a style which prints with newlines,
customizable indentation, and additional debug information such as
types and all pointer addresses used to indirect to the final value.

# Quick Start

This section demonstrates how to quickly get started with spew.  See the
sections below for further details on formatting and configuration options.

To dump a variable with full newlines, indentation, type and pointer
information use Sdump:

	str := spew.Sdump(myVar1)

# Configuration Options

Configuration of spew is handled by fields in the ConfigState type.  For
convenience, all of the top-level functions use a global state available
via the spew.Config global.

It is also possible to create a ConfigState instance that provides methods
equivalent to the top-level functions.  This allows concurrent configuration
options.  See the ConfigState documentation for more details.

The following configuration options are available:

  - Indent
    String to use for each indentation level for Sdump function.
    It is a single space by default.  A popular alternative is "\t".

  - MaxDepth
    Maximum number of levels to descend into nested data structures.
    There is no limit by default.

  - DisableMethods
    Disables invocation of error and fmt.Stringer interface methods.
    Method invocation is enabled by default.

  - DisablePointerMethods
    Disables invocation of error and fmt.Stringer interface methods on types
    which only accept pointer receivers from non-pointer variables.
    Pointer method invocation is enabled by default.

  - DisablePointerAddresses
    DisablePointerAddresses specifies whether to disable the printing of
    pointer addresses. This is useful when diffing data structures in tests.

  - EnableCapacities
    EnableCapacities specifies whether to enaable the printing of
    capacities for arrays, slices and channels.

# Sdump Usage

Simply call spew.Sdump with a variable you want to dump as a string:

	str := spew.Sdump(myVar1)

# Sample Sdump Output

	(main.Foo) {
	 unexportedField: (*main.Bar)(0xf84002e210)({
	  flag: (main.Flag) flagTwo,
	  data: (uintptr) <nil>
	 }),
	 ExportedField: (map[interface {}]interface {}) (len=1) {
	  (string) (len=3) "one": (bool) true
	 }
	}

Byte (and uint8) arrays and slices are displayed uniquely like the hexdump -C
command as shown.

	([]uint8) (len=32) {
	 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20  |............... |
	 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30  |!"#$%&'()*+,-./0|
	 00000020  31 32                                             |12|
	}

# Errors

Since it is possible for custom fmt.Stringer/error interfaces to panic,
spew detects them and handles them internally by printing the panic
information inline with the output.  Since spew is intended to
provide deep pretty printing capabilities on structures, it
intentionally does not return any errors.
*/
package spew
