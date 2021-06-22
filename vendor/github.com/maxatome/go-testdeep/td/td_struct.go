// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"sort"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/util"
)

type tdStruct struct {
	tdExpectedType
	expectedFields fieldInfoSlice
}

var _ TestDeep = &tdStruct{}

type fieldInfo struct {
	name     string
	expected reflect.Value
	index    []int
}

type fieldInfoSlice []fieldInfo

func (e fieldInfoSlice) Len() int           { return len(e) }
func (e fieldInfoSlice) Less(i, j int) bool { return e[i].name < e[j].name }
func (e fieldInfoSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// StructFields allows to pass struct fields to check in functions
// Struct and SStruct. It is a map whose each key is the expected
// field name and the corresponding value the expected field value
// (which can be a TestDeep operator as well as a zero value.)
type StructFields map[string]interface{}

func newStruct(model interface{}, strict bool) (*tdStruct, reflect.Value) {
	vmodel := reflect.ValueOf(model)

	st := tdStruct{
		tdExpectedType: tdExpectedType{
			base: newBase(5),
		},
	}

	switch vmodel.Kind() {
	case reflect.Ptr:
		if vmodel.Type().Elem().Kind() != reflect.Struct {
			break
		}

		st.isPtr = true

		if vmodel.IsNil() {
			st.expectedType = vmodel.Type().Elem()
			return &st, reflect.Value{}
		}

		vmodel = vmodel.Elem()
		fallthrough

	case reflect.Struct:
		st.expectedType = vmodel.Type()
		return &st, vmodel
	}

	var s string
	if strict {
		s = "S"
	}
	panic("usage: " + s + "Struct(STRUCT|&STRUCT, EXPECTED_FIELDS)")
}

func anyStruct(model interface{}, expectedFields StructFields, strict bool) *tdStruct {
	st, vmodel := newStruct(model, strict)

	st.expectedFields = make([]fieldInfo, 0, len(expectedFields))
	checkedFields := make(map[string]bool, len(expectedFields))

	// Check that all given fields are available in model
	stType := st.expectedType
	var vexpectedValue reflect.Value
	for fieldName, expectedValue := range expectedFields {
		field, found := stType.FieldByName(fieldName)
		if !found {
			panic(fmt.Sprintf("struct %s has no field `%s'", stType, fieldName))
		}

		if expectedValue == nil {
			switch field.Type.Kind() {
			case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
				reflect.Ptr, reflect.Slice:
				vexpectedValue = reflect.Zero(field.Type) // change to a typed nil
			default:
				panic(fmt.Sprintf(
					"expected value of field %s cannot be nil as it is a %s",
					fieldName,
					field.Type))
			}
		} else {
			vexpectedValue = reflect.ValueOf(expectedValue)

			if _, ok := expectedValue.(TestDeep); !ok {
				if !vexpectedValue.Type().AssignableTo(field.Type) {
					panic(fmt.Sprintf(
						"type %s of field expected value %s differs from struct one (%s)",
						vexpectedValue.Type(),
						fieldName,
						field.Type))
				}
			}
		}

		st.expectedFields = append(st.expectedFields, fieldInfo{
			name:     fieldName,
			expected: vexpectedValue,
			index:    field.Index,
		})
		checkedFields[fieldName] = true
	}

	// Get all field names
	allFields := map[string]struct{}{}
	stType.FieldByNameFunc(func(fieldName string) bool {
		allFields[fieldName] = struct{}{}
		return false
	})

	// Check initialized fields in model
	if vmodel.IsValid() {
		for fieldName := range allFields {
			field, _ := stType.FieldByName(fieldName)
			if field.Anonymous {
				continue
			}

			vfield := vmodel.FieldByIndex(field.Index)

			// Try to force access to unexported fields
			fieldIf, ok := dark.GetInterface(vfield, true)
			if !ok {
				// Probably in an environment where "unsafe" package is forbidden… :(
				fmt.Fprintf(os.Stderr, // nolint: errcheck
					"field %s is unexported and cannot be overridden, skip it from model.\n",
					fieldName)
				continue
			}

			// If non-zero field
			if !reflect.DeepEqual(reflect.Zero(field.Type).Interface(), fieldIf) {
				if checkedFields[fieldName] {
					panic(fmt.Sprintf(
						"non zero field %s in model already exists in expectedFields",
						fieldName))
				}

				st.expectedFields = append(st.expectedFields, fieldInfo{
					name:     fieldName,
					expected: vfield,
					index:    field.Index,
				})
				checkedFields[fieldName] = true
			}
		}
	}

	// If strict, fill non explicitly expected fields to zero
	if strict {
		for fieldName := range allFields {
			if checkedFields[fieldName] {
				continue
			}

			field, _ := stType.FieldByName(fieldName)
			if field.Anonymous {
				continue
			}

			st.expectedFields = append(st.expectedFields, fieldInfo{
				name:     fieldName,
				expected: reflect.New(field.Type).Elem(), // zero
				index:    field.Index,
			})
		}
	}

	sort.Sort(st.expectedFields)

	return st
}

// summary(Struct): compares the contents of a struct or a pointer on
// a struct
// input(Struct): struct,ptr(ptr on struct)

// Struct operator compares the contents of a struct or a pointer on a
// struct against the non-zero values of "model" (if any) and the
// values of "expectedFields". See SStruct to compares against zero
// fields without specifying them in "expectedFields".
//
// "model" must be the same type as compared data.
//
// "expectedFields" can be nil, if no zero entries are expected and
// no TestDeep operators are involved.
//
//   td.Cmp(t, td.Struct(
//     Person{
//       Name: "John Doe",
//     },
//     td.StructFields{
//       Age:      td.Between(40, 45),
//       Children: 0,
//     }),
//   )
//
// During a match, all expected fields must be found to
// succeed. Non-expected fields are ignored.
//
// TypeBehind method returns the reflect.Type of "model".
func Struct(model interface{}, expectedFields StructFields) TestDeep {
	return anyStruct(model, expectedFields, false)
}

// summary(SStruct): strictly compares the contents of a struct or a
// pointer on a struct
// input(SStruct): struct,ptr(ptr on struct)

// SStruct operator (aka strict-Struct) compares the contents of a
// struct or a pointer on a struct against values of "model" (if any)
// and the values of "expectedFields". The zero values are compared
// too even if they are omitted from "expectedFields": that is the
// difference with Struct operator.
//
// "model" must be the same type as compared data.
//
// "expectedFields" can be nil, if no TestDeep operators are involved.
//
// To ignore a field, one has to specify it in "expectedFields" and
// use the Ignore operator.
//
//   td.Cmp(t, td.SStruct(
//     Person{
//       Name: "John Doe",
//     },
//     td.StructFields{
//       Age:      td.Between(40, 45),
//       Children: td.Ignore(),
//     }),
//   )
//
// During a match, all expected and zero fields must be found to
// succeed.
//
// TypeBehind method returns the reflect.Type of "model".
func SStruct(model interface{}, expectedFields StructFields) TestDeep {
	return anyStruct(model, expectedFields, true)
}

func (s *tdStruct) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	err = s.checkPtr(ctx, &got, false)
	if err != nil {
		return ctx.CollectError(err)
	}

	err = s.checkType(ctx, got)
	if err != nil {
		return ctx.CollectError(err)
	}

	for _, fieldInfo := range s.expectedFields {
		err = deepValueEqual(ctx.AddField(fieldInfo.name),
			got.FieldByIndex(fieldInfo.index), fieldInfo.expected)
		if err != nil {
			return
		}
	}
	return nil
}

func (s *tdStruct) String() string {
	buf := bytes.NewBufferString(s.location.Func)
	buf.WriteByte('(')

	if s.isPtr {
		buf.WriteByte('*')
	}

	buf.WriteString(s.expectedType.String())

	if len(s.expectedFields) == 0 {
		buf.WriteString("{})")
	} else {
		buf.WriteString("{\n")

		for _, fieldInfo := range s.expectedFields {
			fmt.Fprintf(buf, "  %s: %s\n", // nolint: errcheck
				fieldInfo.name, util.ToString(fieldInfo.expected))
		}

		buf.WriteString("})")
	}

	return buf.String()
}
