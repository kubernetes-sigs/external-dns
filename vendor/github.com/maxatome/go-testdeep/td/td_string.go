// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdStringBase struct {
	base
	expected string
}

func newStringBase(expected string) tdStringBase {
	return tdStringBase{
		base:     newBase(4),
		expected: expected,
	}
}

func getString(ctx ctxerr.Context, got reflect.Value) (string, *ctxerr.Error) {
	switch got.Kind() {
	case reflect.String:
		return got.String(), nil

	case reflect.Slice:
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
		if got.Type().Elem() == types.Uint8 {
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		if got.Type().Elem() == uint8Type {
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		if got.Type().Elem() == uint8Type {
=======
		if got.Type().Elem() == types.Uint8 {
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		if got.Type().Elem() == uint8Type {
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		if got.Type().Elem() == uint8Type {
=======
		if got.Type().Elem() == types.Uint8 {
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		if got.Type().Elem() == uint8Type {
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		if got.Type().Elem() == uint8Type {
=======
		if got.Type().Elem() == types.Uint8 {
>>>>>>> 4d7e5ad26 (update vendored files)
			return string(got.Bytes()), nil
		}
		fallthrough

	default:
		if got.CanInterface() {
			switch iface := got.Interface().(type) {
			case error:
				return iface.Error(), nil
			case fmt.Stringer:
				return iface.String(), nil
			}
		}
	}

	if ctx.BooleanError {
		return "", ctxerr.BooleanError
	}
	return "", ctx.CollectError(&ctxerr.Error{
		Message: "bad type",
		Got:     types.RawString(got.Type().String()),
		Expected: types.RawString(
			"string (convertible) OR []byte (convertible) OR fmt.Stringer OR error"),
	})
}

type tdString struct {
	tdStringBase
}

var _ TestDeep = &tdString{}

// summary(String): checks a string, []byte, error or fmt.Stringer
// interfaces string contents
// input(String): str,slice([]byte),if(✓ + fmt.Stringer/error)

// String operator allows to compare a string (or convertible), []byte
// (or convertible), error or [fmt.Stringer] interface (error interface
// is tested before [fmt.Stringer]).
//
//	err := errors.New("error!")
//	td.Cmp(t, err, td.String("error!")) // succeeds
//
//	bstr := bytes.NewBufferString("fmt.Stringer!")
//	td.Cmp(t, bstr, td.String("fmt.Stringer!")) // succeeds
//
// See also [Contains], [HasPrefix], [HasSuffix], [Re] and [ReAll].
func String(expected string) TestDeep {
	return &tdString{
		tdStringBase: newStringBase(expected),
	}
}

func (s *tdString) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	str, err := getString(ctx, got)
	if err != nil {
		return err
	}

	if str == s.expected {
		return nil
	}
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "does not match",
		Got:      str,
		Expected: s,
	})
}

func (s *tdString) String() string {
	return util.ToString(s.expected)
}

type tdHasPrefix struct {
	tdStringBase
}

var _ TestDeep = &tdHasPrefix{}

// summary(HasPrefix): checks the prefix of a string, []byte, error or
// fmt.Stringer interfaces
// input(HasPrefix): str,slice([]byte),if(✓ + fmt.Stringer/error)

// HasPrefix operator allows to compare the prefix of a string (or
// convertible), []byte (or convertible), error or [fmt.Stringer]
// interface (error interface is tested before [fmt.Stringer]).
//
//	td.Cmp(t, []byte("foobar"), td.HasPrefix("foo")) // succeeds
//
//	type Foobar string
//	td.Cmp(t, Foobar("foobar"), td.HasPrefix("foo")) // succeeds
//
//	err := errors.New("error!")
//	td.Cmp(t, err, td.HasPrefix("err")) // succeeds
//
//	bstr := bytes.NewBufferString("fmt.Stringer!")
//	td.Cmp(t, bstr, td.HasPrefix("fmt")) // succeeds
//
// See also [Contains], [HasSuffix], [Re], [ReAll] and [String].
func HasPrefix(expected string) TestDeep {
	return &tdHasPrefix{
		tdStringBase: newStringBase(expected),
	}
}

func (s *tdHasPrefix) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	str, err := getString(ctx, got)
	if err != nil {
		return err
	}

	if strings.HasPrefix(str, s.expected) {
		return nil
	}
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "has not prefix",
		Got:      str,
		Expected: s,
	})
}

func (s *tdHasPrefix) String() string {
	return "HasPrefix(" + util.ToString(s.expected) + ")"
}

type tdHasSuffix struct {
	tdStringBase
}

var _ TestDeep = &tdHasSuffix{}

// summary(HasSuffix): checks the suffix of a string, []byte, error or
// fmt.Stringer interfaces
// input(HasSuffix): str,slice([]byte),if(✓ + fmt.Stringer/error)

// HasSuffix operator allows to compare the suffix of a string (or
// convertible), []byte (or convertible), error or [fmt.Stringer]
// interface (error interface is tested before [fmt.Stringer]).
//
//	td.Cmp(t, []byte("foobar"), td.HasSuffix("bar")) // succeeds
//
//	type Foobar string
//	td.Cmp(t, Foobar("foobar"), td.HasSuffix("bar")) // succeeds
//
//	err := errors.New("error!")
//	td.Cmp(t, err, td.HasSuffix("!")) // succeeds
//
//	bstr := bytes.NewBufferString("fmt.Stringer!")
//	td.Cmp(t, bstr, td.HasSuffix("!")) // succeeds
//
// See also [Contains], [HasPrefix], [Re], [ReAll] and [String].
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
		if got.Type().Elem() == uint8Type {
			return string(got.Bytes()), nil
		}
		fallthrough

	default:
		if got.CanInterface() {
			switch iface := got.Interface().(type) {
			case error:
				return iface.Error(), nil
			case fmt.Stringer:
				return iface.String(), nil
			}
		}
	}

	if ctx.BooleanError {
		return "", ctxerr.BooleanError
	}
	return "", ctx.CollectError(&ctxerr.Error{
		Message: "bad type",
		Got:     types.RawString(got.Type().String()),
		Expected: types.RawString(
			"string (convertible) OR []byte (convertible) OR fmt.Stringer OR error"),
	})
}

type tdString struct {
	tdStringBase
}

var _ TestDeep = &tdString{}

// summary(String): checks a string, []byte, error or fmt.Stringer
// interfaces string contents
// input(String): str,slice([]byte),if(✓ + fmt.Stringer/error)

// String operator allows to compare a string (or convertible), []byte
// (or convertible), error or fmt.Stringer interface (error interface
// is tested before fmt.Stringer).
//
//   err := errors.New("error!")
//   td.Cmp(t, err, td.String("error!")) // succeeds
//
//   bstr := bytes.NewBufferString("fmt.Stringer!")
//   td.Cmp(t, bstr, td.String("fmt.Stringer!")) // succeeds
func String(expected string) TestDeep {
	return &tdString{
		tdStringBase: newStringBase(expected),
	}
}

func (s *tdString) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	str, err := getString(ctx, got)
	if err != nil {
		return err
	}

	if str == s.expected {
		return nil
	}
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "does not match",
		Got:      str,
		Expected: s,
	})
}

func (s *tdString) String() string {
	return util.ToString(s.expected)
}

type tdHasPrefix struct {
	tdStringBase
}

var _ TestDeep = &tdHasPrefix{}

// summary(HasPrefix): checks the prefix of a string, []byte, error or
// fmt.Stringer interfaces
// input(HasPrefix): str,slice([]byte),if(✓ + fmt.Stringer/error)

// HasPrefix operator allows to compare the prefix of a string (or
// convertible), []byte (or convertible), error or fmt.Stringer
// interface (error interface is tested before fmt.Stringer).
//
//   td.Cmp(t, []byte("foobar"), td.HasPrefix("foo")) // succeeds
//
//   type Foobar string
//   td.Cmp(t, Foobar("foobar"), td.HasPrefix("foo")) // succeeds
//
//   err := errors.New("error!")
//   td.Cmp(t, err, td.HasPrefix("err")) // succeeds
//
//   bstr := bytes.NewBufferString("fmt.Stringer!")
//   td.Cmp(t, bstr, td.HasPrefix("fmt")) // succeeds
func HasPrefix(expected string) TestDeep {
	return &tdHasPrefix{
		tdStringBase: newStringBase(expected),
	}
}

func (s *tdHasPrefix) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	str, err := getString(ctx, got)
	if err != nil {
		return err
	}

	if strings.HasPrefix(str, s.expected) {
		return nil
	}
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "has not prefix",
		Got:      str,
		Expected: s,
	})
}

func (s *tdHasPrefix) String() string {
	return "HasPrefix(" + util.ToString(s.expected) + ")"
}

type tdHasSuffix struct {
	tdStringBase
}

var _ TestDeep = &tdHasSuffix{}

// summary(HasSuffix): checks the suffix of a string, []byte, error or
// fmt.Stringer interfaces
// input(HasSuffix): str,slice([]byte),if(✓ + fmt.Stringer/error)

// HasSuffix operator allows to compare the suffix of a string (or
// convertible), []byte (or convertible), error or fmt.Stringer
// interface (error interface is tested before fmt.Stringer).
//
//   td.Cmp(t, []byte("foobar"), td.HasSuffix("bar")) // succeeds
//
//   type Foobar string
//   td.Cmp(t, Foobar("foobar"), td.HasSuffix("bar")) // succeeds
//
//   err := errors.New("error!")
//   td.Cmp(t, err, td.HasSuffix("!")) // succeeds
//
//   bstr := bytes.NewBufferString("fmt.Stringer!")
//   td.Cmp(t, bstr, td.HasSuffix("!")) // succeeds
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
func HasSuffix(expected string) TestDeep {
	return &tdHasSuffix{
		tdStringBase: newStringBase(expected),
	}
}

func (s *tdHasSuffix) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	str, err := getString(ctx, got)
	if err != nil {
		return err
	}

	if strings.HasSuffix(str, s.expected) {
		return nil
	}
	if ctx.BooleanError {
		return ctxerr.BooleanError
	}
	return ctx.CollectError(&ctxerr.Error{
		Message:  "has not suffix",
		Got:      str,
		Expected: s,
	})
}

func (s *tdHasSuffix) String() string {
	return "HasSuffix(" + util.ToString(s.expected) + ")"
}
