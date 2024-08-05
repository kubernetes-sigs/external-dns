// Copyright (c) 2023, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdStructLazy struct {
	base
	cache          map[reflect.Type]*tdStruct
	expectedFields StructFields
	strict         bool
}

var _ TestDeep = &tdStructLazy{}

func newStructLazy(expectedFields StructFields, strict bool) TestDeep {
	return &tdStructLazy{
		base:           newBase(4),
		cache:          map[reflect.Type]*tdStruct{},
		expectedFields: expectedFields,
		strict:         strict,
	}
}

func (s *tdStructLazy) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	gotType := got.Type()
	tds := s.cache[gotType]
	if tds == nil {
		switch gotType.Kind() {
		case reflect.Struct:
		case reflect.Ptr:
			if gotType.Elem().Kind() == reflect.Struct {
				break
			}
			fallthrough
		default:
			if ctx.BooleanError {
				return ctxerr.BooleanError
			}
			return ctx.CollectError(ctxerr.BadKind(got, "struct OR *struct"))
		}

		tds = anyStruct(s.base, reflect.New(got.Type()).Elem(), s.expectedFields, s.strict)
		tds.location = s.location
		s.cache[gotType] = tds
	}

	return tds.Match(ctx, got)
}

func (s *tdStructLazy) String() string {
	buf := bytes.NewBufferString(s.location.Func)
	buf.WriteString("(<any struct type>{")

	if len(s.expectedFields) > 0 {
		buf.WriteByte('\n')

		fields := make([]string, 0, len(s.expectedFields))
		maxLen := 0
		for name := range s.expectedFields {
			fields = append(fields, name)
			if len(name) > maxLen {
				maxLen = len(name)
			}
		}
		sort.Strings(fields)

		maxLen++
		for _, name := range fields {
			fmt.Fprintf(buf, "  %-*s %s\n", //nolint: errcheck
				maxLen, name+":", util.ToString(s.expectedFields[name]))
		}
	}
	buf.WriteString("})")

	return buf.String()
}
