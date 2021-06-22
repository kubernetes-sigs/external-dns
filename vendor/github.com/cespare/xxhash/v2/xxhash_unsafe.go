<<<<<<< HEAD
//go:build !appengine
// +build !appengine

// This file encapsulates usage of unsafe.
// xxhash_safe.go contains the safe implementations.

package xxhash

import (
	"unsafe"
)

// In the future it's possible that compiler optimizations will make these
// XxxString functions unnecessary by realizing that calls such as
// Sum64([]byte(s)) don't need to copy s. See https://go.dev/issue/2205.
// If that happens, even if we keep these functions they can be replaced with
// the trivial safe code.

// NOTE: The usual way of doing an unsafe string-to-[]byte conversion is:
//
//   var b []byte
//   bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
//   bh.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
//   bh.Len = len(s)
//   bh.Cap = len(s)
//
// Unfortunately, as of Go 1.15.3 the inliner's cost model assigns a high enough
// weight to this sequence of expressions that any function that uses it will
// not be inlined. Instead, the functions below use a different unsafe
// conversion designed to minimize the inliner weight and allow both to be
// inlined. There is also a test (TestInlining) which verifies that these are
// inlined.
//
// See https://github.com/golang/go/issues/42739 for discussion.

// Sum64String computes the 64-bit xxHash digest of s.
// It may be faster than Sum64([]byte(s)) by avoiding a copy.
func Sum64String(s string) uint64 {
	b := *(*[]byte)(unsafe.Pointer(&sliceHeader{s, len(s)}))
	return Sum64(b)
}

// WriteString adds more data to d. It always returns len(s), nil.
// It may be faster than Write([]byte(s)) by avoiding a copy.
func (d *Digest) WriteString(s string) (n int, err error) {
	d.Write(*(*[]byte)(unsafe.Pointer(&sliceHeader{s, len(s)})))
	// d.Write always returns len(s), nil.
	// Ignoring the return output and returning these fixed values buys a
	// savings of 6 in the inliner's cost model.
	return len(s), nil
}

// sliceHeader is similar to reflect.SliceHeader, but it assumes that the layout
// of the first two words is the same as the layout of a string.
type sliceHeader struct {
	s   string
	cap int
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !appengine

// This file encapsulates usage of unsafe.
// xxhash_safe.go contains the safe implementations.

package xxhash

import (
	"reflect"
	"unsafe"
)

// Notes:
//
// See https://groups.google.com/d/msg/golang-nuts/dcjzJy-bSpw/tcZYBzQqAQAJ
// for some discussion about these unsafe conversions.
//
// In the future it's possible that compiler optimizations will make these
// unsafe operations unnecessary: https://golang.org/issue/2205.
//
// Both of these wrapper functions still incur function call overhead since they
// will not be inlined. We could write Go/asm copies of Sum64 and Digest.Write
// for strings to squeeze out a bit more speed. Mid-stack inlining should
// eventually fix this.

// Sum64String computes the 64-bit xxHash digest of s.
// It may be faster than Sum64([]byte(s)) by avoiding a copy.
func Sum64String(s string) uint64 {
	var b []byte
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	bh.Len = len(s)
	bh.Cap = len(s)
	return Sum64(b)
}

// WriteString adds more data to d. It always returns len(s), nil.
// It may be faster than Write([]byte(s)) by avoiding a copy.
func (d *Digest) WriteString(s string) (n int, err error) {
	var b []byte
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	bh.Data = (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	bh.Len = len(s)
	bh.Cap = len(s)
	return d.Write(b)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
