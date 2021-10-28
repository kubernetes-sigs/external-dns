// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"reflect"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdTag struct {
	tdSmugglerBase
	tag string
}

var _ TestDeep = &tdTag{}

// summary(Tag): names an operator or a value. Only useful as a
// parameter of JSON operator, to name placeholders
// input(Tag): all

// Tag is a smuggler operator. It only allows to name "expectedValue",
// which can be an operator or a value. The data is then compared
// against "expectedValue" as if Tag was never called. It is only
// useful as JSON operator parameter, to name placeholders. See JSON
// operator for more details.
//
//   td.Cmp(t, gotValue,
//     td.JSON(`{"fullname": $name, "age": $age, "gender": $gender}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $name
//       td.Tag("age", td.Between(41, 43)),   // matches $age
//       td.Tag("gender", "male")))           // matches $gender
//
// TypeBehind method is delegated to "expectedValue" one if
// "expectedValue" is a TestDeep operator, otherwise it returns the
// type of "expectedValue" (or nil if it is originally untyped nil).
func Tag(tag string, expectedValue interface{}) TestDeep {
	t := tdTag{
		tdSmugglerBase: newSmugglerBase(expectedValue),
		tag:            tag,
	}

	if err := util.CheckTag(tag); err != nil {
		t.err = ctxerr.OpBad("Tag", err.Error())
		return &t
	}

	if !t.isTestDeeper {
		t.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &t
}

func (t *tdTag) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if t.err != nil {
		return ctx.CollectError(t.err)
	}
	return deepValueEqual(ctx, got, t.expectedValue)
}

func (t *tdTag) HandleInvalid() bool {
	return true // Knows how to handle untyped nil values (aka invalid values)
}

func (t *tdTag) String() string {
	if t.err != nil {
		return t.stringError()
	}
	if t.isTestDeeper {
		return t.expectedValue.Interface().(TestDeep).String()
	}
	return util.ToString(t.expectedValue)
}

func (t *tdTag) TypeBehind() reflect.Type {
	if t.err != nil {
		return nil
	}
	if t.isTestDeeper {
		return t.expectedValue.Interface().(TestDeep).TypeBehind()
	}
	if t.expectedValue.IsValid() {
		return t.expectedValue.Type()
	}
	return nil
}
