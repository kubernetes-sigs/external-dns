// Copyright (c) 2019-2021, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	ejson "encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/flat"
	"github.com/maxatome/go-testdeep/internal/json"
	"github.com/maxatome/go-testdeep/internal/location"
	"github.com/maxatome/go-testdeep/internal/types"
)

// forbiddenOpsInJSON contains operators forbidden inside JSON,
// SubJSONOf or SuperJSONOf, optionally with an alternative to help
// the user.
var forbiddenOpsInJSON = map[string]string{
	"Array":        "literal []",
	"Cap":          "",
	"Catch":        "",
	"Code":         "",
	"Delay":        "",
	"Isa":          "",
	"JSON":         "literal JSON",
	"Lax":          "",
	"Map":          "literal {}",
	"PPtr":         "",
	"Ptr":          "",
	"SStruct":      "",
	"Shallow":      "",
	"Slice":        "literal []",
	"Smuggle":      "",
	"String":       `literal ""`,
	"SubJSONOf":    "SubMapOf operator",
	"SuperJSONOf":  "SuperMapOf operator",
	"SuperSliceOf": "All and JSONPointer operators",
	"Struct":       "",
	"Tag":          "",
	"TruncTime":    "",
}

// jsonOpShortcuts contains operator that can be used as
// $^OperatorName inside JSON, SubJSONOf or SuperJSONOf.
var jsonOpShortcuts = map[string]func() TestDeep{
	"Empty":    Empty,
	"Ignore":   Ignore,
	"NaN":      NaN,
	"Nil":      Nil,
	"NotEmpty": NotEmpty,
	"NotNaN":   NotNaN,
	"NotNil":   NotNil,
	"NotZero":  NotZero,
	"Zero":     Zero,
}

// tdJSONUnmarshaler handles the JSON unmarshaling of JSON, SubJSONOf
// and SuperJSONOf first parameter.
type tdJSONUnmarshaler struct {
	location.Location // position of the operator
}

// newJSONUnmarshaler returns a new instance of tdJSONUnmarshaler.
func newJSONUnmarshaler(pos location.Location) tdJSONUnmarshaler {
	return tdJSONUnmarshaler{
		Location: pos,
	}
}

// replaceLocation replaces the location of tdOp by the
// JSON/SubJSONOf/SuperJSONOf one then add the position of the
// operator inside the JSON string.
func (u tdJSONUnmarshaler) replaceLocation(tdOp TestDeep, posInJSON json.Position) {
	// The goal, instead of:
	//    [under operator Len at value.go:476]
	// having:
	//    [under operator Len at line 12:7 (pos 123) inside operator JSON at file.go:23]
	//                 so add ^------------------------------------------^

	newPos := u.Location
	newPos.Inside = fmt.Sprintf("%s inside operator %s ", posInJSON, u.Func)
	newPos.Func = tdOp.GetLocation().Func
	tdOp.replaceLocation(newPos)
}

// unmarshal unmarshals "expectedJSON" using placeholder parameters "params".
func (u tdJSONUnmarshaler) unmarshal(expectedJSON interface{}, params []interface{}) (interface{}, *ctxerr.Error) {
	var (
		err error
		b   []byte
	)

	switch data := expectedJSON.(type) {
	case string:
		// Try to load this file (if it seems it can be a filename and not
		// a JSON content)
		if strings.HasSuffix(data, ".json") {
			// It could be a file name, try to read from it
			b, err = ioutil.ReadFile(data)
			if err != nil {
				return nil, ctxerr.OpBad(u.Func, "JSON file %s cannot be read: %s", data, err)
			}
			break
		}
		b = []byte(data)

	case []byte:
		b = data

	case io.Reader:
		b, err = ioutil.ReadAll(data)
		if err != nil {
			return nil, ctxerr.OpBad(u.Func, "JSON read error: %s", err)
		}

	default:
		return nil, ctxerr.OpBadUsage(
			u.Func, "(STRING_JSON|STRING_FILENAME|[]byte|io.Reader, ...)",
			expectedJSON, 1, false)
	}

	params = flat.Interfaces(params...)
	var byTag map[string]interface{}

	for i, p := range params {
		switch op := p.(type) {
		case *tdTag:
			if byTag[op.tag] != nil {
				return nil, ctxerr.OpBad(u.Func, `2 params have the same tag "%s"`, op.tag)
			}
			if byTag == nil {
				byTag = map[string]interface{}{}
			}
			// Don't keep the tag layer
			p = nil
			if op.expectedValue.IsValid() {
				p = op.expectedValue.Interface()
			}
			byTag[op.tag] = newJSONNamedPlaceholder(op.tag, p)

		default:
			params[i] = newJSONNumPlaceholder(uint64(i+1), p)
		}
	}

	final, err := json.Parse(b, json.ParseOpts{
		Placeholders:       params,
		PlaceholdersByName: byTag,
		OpShortcutFn:       u.resolveOpShortcut(),
		OpFn:               u.resolveOp(),
	})

	if err != nil {
		return nil, ctxerr.OpBad(u.Func, "JSON unmarshal error: %s", err)
	}

	return final, nil
}

// resolveOp returns a closure usable as json.ParseOpts.OpFn.
func (u tdJSONUnmarshaler) resolveOp() func(json.Operator, json.Position) (interface{}, error) {
	return func(jop json.Operator, posInJSON json.Position) (interface{}, error) {
		op, exists := allOperators[jop.Name]
		if !exists {
			return nil, fmt.Errorf("unknown operator %s()", jop.Name)
		}

		if hint, exists := forbiddenOpsInJSON[jop.Name]; exists {
			if hint == "" {
				return nil, fmt.Errorf("%s() is not usable in JSON()", jop.Name)
			}
			return nil, fmt.Errorf("%s() is not usable in JSON(), use %s instead",
				jop.Name, hint)
		}

		vfn := reflect.ValueOf(op)
		tfn := vfn.Type()

		// If some parameters contain a placeholder, dereference it
		for i, p := range jop.Params {
			if ph, ok := p.(*tdJSONPlaceholder); ok {
				jop.Params[i] = ph.expectedValue.Interface()
			}
		}

		// Special cases
		var min, max int
		addNilParam := false
		switch jop.Name {
		case "Between":
			min, max = 2, 3
			if len(jop.Params) == 3 {
				switch s, _ := jop.Params[2].(string); s {
				case "[]", "BoundsInIn":
					jop.Params[2] = BoundsInIn
				case "[[", "BoundsInOut":
					jop.Params[2] = BoundsInOut
				case "]]", "BoundsOutIn":
					jop.Params[2] = BoundsOutIn
				case "][", "BoundsOutOut":
					jop.Params[2] = BoundsOutOut
				default:
					return nil, errors.New(`Between() bad 3rd parameter, use "[]", "[[", "]]" or "]["`)
				}
			}
		case "N":
			min, max = 1, 2
		case "Re":
			min, max = 1, 2
		case "SubMapOf":
			min, max, addNilParam = 1, 1, true
		case "SuperMapOf":
			min, max, addNilParam = 1, 1, true
		default:
			min = tfn.NumIn()
			if tfn.IsVariadic() {
				// for All(expected ...interface{}) → min == 1, as All() is a non-sense
				max = -1
			} else {
				max = min
			}
		}
		if len(jop.Params) < min || (max >= 0 && len(jop.Params) > max) {
			switch {
			case max < 0:
				return nil, fmt.Errorf("%s() requires at least one parameter", jop.Name)
			case max == 0:
				return nil, fmt.Errorf("%s() requires no parameters", jop.Name)
			case min == max:
				if min == 1 {
					return nil, fmt.Errorf("%s() requires only one parameter", jop.Name)
				}
				return nil, fmt.Errorf("%s() requires %d parameters", jop.Name, min)
			default:
				return nil, fmt.Errorf("%s() requires %d or %d parameters", jop.Name, min, max)
			}
		}

		var in []reflect.Value
		if len(jop.Params) > 0 {
			in = make([]reflect.Value, len(jop.Params))
			for i, p := range jop.Params {
				in[i] = reflect.ValueOf(p)
			}
			if addNilParam {
				in = append(in, reflect.ValueOf(MapEntries(nil)))
			}

			// If the function is variadic, no need to check each param as all
			// variadic operators have always a ...interface{}
			numCheck := len(in)
			if tfn.IsVariadic() {
				numCheck = tfn.NumIn() - 1
			}
			for i, p := range in[:numCheck] {
				fpt := tfn.In(i)
				if fpt.Kind() != reflect.Interface && p.Type() != fpt {
					return nil, fmt.Errorf(
						"%s() bad #%d parameter type: %s required but %s received",
						jop.Name, i+1,
						fpt, p.Type(),
					)
				}
			}
		}

		tdOp := vfn.Call(in)[0].Interface().(TestDeep)

		// let erroneous operators (tdOp.err != nil) pass

		// replace the location by the JSON/SubJSONOf/SuperJSONOf one
		u.replaceLocation(tdOp, posInJSON)
		return &tdJSONEmbedded{tdOp}, nil
	}
}

// resolveOpShortcut returns a closure usable as json.ParseOpts.OpShortcutFn.
func (u tdJSONUnmarshaler) resolveOpShortcut() func(string, json.Position) (interface{}, bool) {
	return func(opName string, posInJSON json.Position) (interface{}, bool) {
		opFn := jsonOpShortcuts[opName]
		if opFn != nil {
			tdOp := opFn()

			// replace the location by the JSON/SubJSONOf/SuperJSONOf one
			u.replaceLocation(tdOp, posInJSON)

			return newJSONNamedPlaceholder("^"+opName, tdOp), true
		}
		return nil, false
	}
}

// tdJSONPlaceholder is an internal smuggler operator. It represents a
// JSON placeholder in an unmarshaled JSON expected data
// structure. It takes the JSON representation of data and compares it
// to "expectedValue".
//
// It does its best to convert back the JSON pointed data to the type
// of "expectedValue" or to the type behind the "expectedValue".
type tdJSONPlaceholder struct {
	tdSmugglerBase // ignored by tools/gen_funcs.pl
	name           string
	num            uint64
}

func newJSONNamedPlaceholder(name string, expectedValue interface{}) TestDeep {
	p := tdJSONPlaceholder{
		tdSmugglerBase: newSmugglerBase(expectedValue, 1),
		name:           name,
	}

	if !p.isTestDeeper {
		p.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &p
}

func newJSONNumPlaceholder(num uint64, expectedValue interface{}) TestDeep {
	p := tdJSONPlaceholder{
		tdSmugglerBase: newSmugglerBase(expectedValue, 1),
		num:            num,
	}

	if !p.isTestDeeper {
		p.expectedValue = reflect.ValueOf(expectedValue)
	}
	return &p
}

func (p *tdJSONPlaceholder) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	vgot, _ := jsonify(ctx, got) // Cannot fail

	// Here, vgot type is either a bool, float64, string,
	// []interface{}, a map[string]interface{} or simply nil

	return p.jsonValueEqual(ctx, vgot)
}

func (p *tdJSONPlaceholder) String() string {
	// Only called by MarshalJSON(), and so only when p.isTestDeeper
	return p.expectedValue.Interface().(TestDeep).String()
}

func (p *tdJSONPlaceholder) HandleInvalid() bool {
	return true
}

func (p *tdJSONPlaceholder) MarshalJSON() ([]byte, error) {
	if !p.isTestDeeper {
		var expected interface{}
		if p.expectedValue.IsValid() {
			expected = p.expectedValue.Interface()
		}
		return ejson.Marshal(expected)
	}

	var b bytes.Buffer

	start := b.Len()
	if p.num == 0 {
		fmt.Fprintf(&b, `"$%s"`, p.name)

		// Don't add a comment for operator shortcuts (aka $^NotZero)
		if p.name[0] == '^' {
			return b.Bytes(), nil
		}
	} else {
		fmt.Fprintf(&b, `"$%d"`, p.num)
	}

	b.WriteString(` /* `)

	indent := "\n" + strings.Repeat(" ", b.Len()-start)
	b.WriteString(strings.Replace(p.String(), "\n", indent, -1)) //nolint: gocritic

	b.WriteString(` */`)

	return b.Bytes(), nil
}

// tdJSONEmbedded represents a MarshalJSON'able operator.
type tdJSONEmbedded struct {
	TestDeep
}

func (e *tdJSONEmbedded) MarshalJSON() ([]byte, error) {
	return []byte(e.String()), nil
}

type tdJSON struct {
	baseOKNil
	expected reflect.Value
}

var _ TestDeep = &tdJSON{}

func gotViaJSON(ctx ctxerr.Context, pGot *reflect.Value) *ctxerr.Error {
	got, err := jsonify(ctx, *pGot)
	if err != nil {
		return err
	}
	*pGot = reflect.ValueOf(got)
	return nil
}

func jsonify(ctx ctxerr.Context, got reflect.Value) (interface{}, *ctxerr.Error) {
	gotIf, ok := dark.GetInterface(got, true)
	if !ok {
		return nil, ctx.CannotCompareError()
	}

	b, err := ejson.Marshal(gotIf)
	if err != nil {
		if ctx.BooleanError {
			return nil, ctxerr.BooleanError
		}
		return nil, &ctxerr.Error{
			Message: "json.Marshal failed",
			Summary: ctxerr.NewSummary(err.Error()),
		}
	}

	// As Marshal succeeded, Unmarshal in an interface{} cannot fail
	var vgot interface{}
	ejson.Unmarshal(b, &vgot) //nolint: errcheck
	return vgot, nil
}

// summary(JSON): compares against JSON representation
// input(JSON): nil,bool,str,int,float,array,slice,map,struct,ptr

// JSON operator allows to compare the JSON representation of data
// against "expectedJSON". "expectedJSON" can be a:
//
//   - string containing JSON data like `{"fullname":"Bob","age":42}`
//   - string containing a JSON filename, ending with ".json" (its
//     content is ioutil.ReadFile before unmarshaling)
//   - []byte containing JSON data
//   - io.Reader stream containing JSON data (is ioutil.ReadAll before
//     unmarshaling)
//
// "expectedJSON" JSON value can contain placeholders. The "params"
// are for any placeholder parameters in "expectedJSON". "params" can
// contain TestDeep operators as well as raw values. A placeholder can
// be numeric like $2 or named like $name and always references an
// item in "params".
//
// Numeric placeholders reference the n'th "operators" item (starting
// at 1). Named placeholders are used with Tag operator as follows:
//
//   td.Cmp(t, gotValue,
//     td.JSON(`{"fullname": $name, "age": $2, "gender": $3}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43),                  // matches only $2
//       "male"))                             // matches only $3
//
// Note that placeholders can be double-quoted as in:
//
//   td.Cmp(t, gotValue,
//     td.JSON(`{"fullname": "$name", "age": "$2", "gender": "$3"}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43),                  // matches only $2
//       "male"))                             // matches only $3
//
// It makes no difference whatever the underlying type of the replaced
// item is (= double quoting a placeholder matching a number is not a
// problem). It is just a matter of taste, double-quoting placeholders
// can be preferred when the JSON data has to conform to the JSON
// specification, like when used in a ".json" file.
//
// JSON does its best to convert back the JSON corresponding to a
// placeholder to the type of the placeholder or, if the placeholder
// is an operator, to the type behind the operator. Allowing to do
// things like:
//
//   td.Cmp(t, gotValue, td.JSON(`{"foo":$1}`, []int{1, 2, 3, 4}))
//   td.Cmp(t, gotValue,
//     td.JSON(`{"foo":$1}`, []interface{}{1, 2, td.Between(2, 4), 4}))
//   td.Cmp(t, gotValue, td.JSON(`{"foo":$1}`, td.Between(27, 32)))
//
// Of course, it does this conversion only if the expected type can be
// guessed. In the case the conversion cannot occur, data is compared
// as is, in its freshly unmarshaled JSON form (so as bool, float64,
// string, []interface{}, map[string]interface{} or simply nil).
//
// Note "expectedJSON" can be a []byte, JSON filename or io.Reader:
//
//   td.Cmp(t, gotValue, td.JSON("file.json", td.Between(12, 34)))
//   td.Cmp(t, gotValue, td.JSON([]byte(`[1, $1, 3]`), td.Between(12, 34)))
//   td.Cmp(t, gotValue, td.JSON(osFile, td.Between(12, 34)))
//
// A JSON filename ends with ".json".
//
// To avoid a legit "$" string prefix causes a bad placeholder error,
// just double it to escape it. Note it is only needed when the "$" is
// the first character of a string:
//
//   td.Cmp(t, gotValue,
//     td.JSON(`{"fullname": "$name", "details": "$$info", "age": $2}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43)))                 // matches only $2
//
// For the "details" key, the raw value "$info" is expected, no
// placeholders are involved here.
//
// Note that Lax mode is automatically enabled by JSON operator to
// simplify numeric tests.
//
// Comments can be embedded in JSON data:
//
//   td.Cmp(t, gotValue,
//     td.JSON(`
//   {
//     // A guy properties:
//     "fullname": "$name",  // The full name of the guy
//     "details":  "$$info", // Literally "$info", thanks to "$" escape
//     "age":      $2        /* The age of the guy:
//                              - placeholder unquoted, but could be without
//                                any change
//                              - to demonstrate a multi-lines comment */
//   }`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43)))                 // matches only $2
//
// Comments, like in go, have 2 forms. To quote the Go language specification:
//   - line comments start with the character sequence // and stop at the
//     end of the line.
//   - multi-lines comments start with the character sequence /* and stop
//     with the first subsequent character sequence */.
//
// Most operators can be directly embedded in JSON without requiring
// any placeholder.
//
//   td.Cmp(t, gotValue,
//     td.JSON(`
//   {
//     "fullname": HasPrefix("Foo"),
//     "age":      Between(41, 43),
//     "details":  SuperMapOf({
//       "address": NotEmpty(),
//       "car":     Any("Peugeot", "Tesla", "Jeep") // any of these
//     })
//   }`))
//
// Placeholders can be used anywhere, even in operators parameters as in:
//
//   td.Cmp(t, gotValue, td.JSON(`{"fullname": HasPrefix($1)}`, "Zip"))
//
// A few notes about operators embedding:
//   - SubMapOf and SuperMapOf take only one parameter, a JSON object;
//   - the optional 3rd parameter of Between has to be specified as a string
//     and can be: "[]" or "BoundsInIn" (default), "[[" or "BoundsInOut",
//     "]]" or "BoundsOutIn", "][" or "BoundsOutOut";
//   - not all operators are embeddable only the following are;
//   - All, Any, ArrayEach, Bag, Between, Contains, ContainsKey, Empty, Gt,
//     Gte, HasPrefix, HasSuffix, Ignore, JSONPointer, Keys, Len, Lt, Lte,
//     MapEach, N, NaN, Nil, None, Not, NotAny, NotEmpty, NotNaN, NotNil,
//     NotZero, Re, ReAll, Set, SubBagOf, SubMapOf, SubSetOf, SuperBagOf,
//     SuperMapOf, SuperSetOf, Values and Zero.
//
// Operators taking no parameters can also be directly embedded in
// JSON data using $^OperatorName or "$^OperatorName" notation. They
// are named shortcut operators (they predate the above operators embedding
// but they subsist for compatibility):
//
//   td.Cmp(t, gotValue, td.JSON(`{"id": $1}`, td.NotZero()))
//
// can be written as:
//
//   td.Cmp(t, gotValue, td.JSON(`{"id": $^NotZero}`))
//
// or
//
//   td.Cmp(t, gotValue, td.JSON(`{"id": "$^NotZero"}`))
//
// As for placeholders, there is no differences between $^NotZero and
// "$^NotZero".
//
// The allowed shortcut operators follow:
//   - Empty    → $^Empty
//   - Ignore   → $^Ignore
//   - NaN      → $^NaN
//   - Nil      → $^Nil
//   - NotEmpty → $^NotEmpty
//   - NotNaN   → $^NotNaN
//   - NotNil   → $^NotNil
//   - NotZero  → $^NotZero
//   - Zero     → $^Zero
//
// TypeBehind method returns the reflect.Type of the "expectedJSON"
// json.Unmarshal'ed. So it can be bool, string, float64,
// []interface{}, map[string]interface{} or interface{} in case
// "expectedJSON" is "null".
func JSON(expectedJSON interface{}, params ...interface{}) TestDeep {
	j := &tdJSON{
		baseOKNil: newBaseOKNil(3),
	}

	v, err := newJSONUnmarshaler(j.GetLocation()).unmarshal(expectedJSON, params)
	if err != nil {
		j.err = err
	} else {
		j.expected = reflect.ValueOf(v)
	}

	return j
}

func (j *tdJSON) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if j.err != nil {
		return ctx.CollectError(j.err)
	}

	err := gotViaJSON(ctx, &got)
	if err != nil {
		return ctx.CollectError(err)
	}

	ctx.BeLax = true

	return deepValueEqual(ctx, got, j.expected)
}

func (j *tdJSON) String() string {
	if j.err != nil {
		return j.stringError()
	}

	return jsonStringify("JSON", j.expected)
}

func jsonStringify(opName string, v reflect.Value) string {
	if !v.IsValid() {
		return "JSON(null)"
	}

	var b bytes.Buffer

	b.WriteString(opName)
	b.WriteByte('(')
	json.AppendMarshal(&b, v.Interface(), len(opName)+1) //nolint: errcheck
	b.WriteByte(')')

	return b.String()
}

func (j *tdJSON) TypeBehind() reflect.Type {
	if j.err != nil {
		return nil
	}

	if j.expected.IsValid() {
		// In case we have an operator at the root, delegate it the call
		if je, ok := j.expected.Interface().(*tdJSONEmbedded); ok {
			return je.TypeBehind()
		}
		return j.expected.Type()
	}
	return types.Interface
}

type tdMapJSON struct {
	tdMap
	expected reflect.Value
}

var _ TestDeep = &tdMapJSON{}

// summary(SubJSONOf): compares struct or map against JSON
// representation but with potentially some exclusions
// input(SubJSONOf): map,struct,ptr(ptr on map/struct)

// SubJSONOf operator allows to compare the JSON representation of
// data against "expectedJSON". Unlike JSON operator, marshaled data
// must be a JSON object/map (aka {…}). "expectedJSON" can be a:
//
//   - string containing JSON data like `{"fullname":"Bob","age":42}`
//   - string containing a JSON filename, ending with ".json" (its
//     content is ioutil.ReadFile before unmarshaling)
//   - []byte containing JSON data
//   - io.Reader stream containing JSON data (is ioutil.ReadAll before
//     unmarshaling)
//
// JSON data contained in "expectedJSON" must be a JSON object/map
// (aka {…}) too. During a match, each expected entry should match in
// the compared map. But some expected entries can be missing from the
// compared map.
//
//   type MyStruct struct {
//     Name string `json:"name"`
//     Age  int    `json:"age"`
//   }
//   got := MyStruct{
//     Name: "Bob",
//     Age:  42,
//   }
//   td.Cmp(t, got, td.SubJSONOf(`{"name": "Bob", "age": 42, "city": "NY"}`)) // succeeds
//   td.Cmp(t, got, td.SubJSONOf(`{"name": "Bob", "zip": 666}`))              // fails, extra "age"
//
// "expectedJSON" JSON value can contain placeholders. The "params"
// are for any placeholder parameters in "expectedJSON". "params" can
// contain TestDeep operators as well as raw values. A placeholder can
// be numeric like $2 or named like $name and always references an
// item in "params".
//
// Numeric placeholders reference the n'th "operators" item (starting
// at 1). Named placeholders are used with Tag operator as follows:
//
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`{"fullname": $name, "age": $2, "gender": $3}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43),                  // matches only $2
//       "male"))                             // matches only $3
//
// Note that placeholders can be double-quoted as in:
//
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`{"fullname": "$name", "age": "$2", "gender": "$3"}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43),                  // matches only $2
//       "male"))                             // matches only $3
//
// It makes no difference whatever the underlying type of the replaced
// item is (= double quoting a placeholder matching a number is not a
// problem). It is just a matter of taste, double-quoting placeholders
// can be preferred when the JSON data has to conform to the JSON
// specification, like when used in a ".json" file.
//
// SubJSONOf does its best to convert back the JSON corresponding to a
// placeholder to the type of the placeholder or, if the placeholder
// is an operator, to the type behind the operator. Allowing to do
// things like:
//
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`{"foo":$1, "bar": 12}`, []int{1, 2, 3, 4}))
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`{"foo":$1, "bar": 12}`, []interface{}{1, 2, td.Between(2, 4), 4}))
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`{"foo":$1, "bar": 12}`, td.Between(27, 32)))
//
// Of course, it does this conversion only if the expected type can be
// guessed. In the case the conversion cannot occur, data is compared
// as is, in its freshly unmarshaled JSON form (so as bool, float64,
// string, []interface{}, map[string]interface{} or simply nil).
//
// Note "expectedJSON" can be a []byte, JSON filename or io.Reader:
//
//   td.Cmp(t, gotValue, td.SubJSONOf("file.json", td.Between(12, 34)))
//   td.Cmp(t, gotValue, td.SubJSONOf([]byte(`[1, $1, 3]`), td.Between(12, 34)))
//   td.Cmp(t, gotValue, td.SubJSONOf(osFile, td.Between(12, 34)))
//
// A JSON filename ends with ".json".
//
// To avoid a legit "$" string prefix causes a bad placeholder error,
// just double it to escape it. Note it is only needed when the "$" is
// the first character of a string:
//
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`{"fullname": "$name", "details": "$$info", "age": $2}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43)))                 // matches only $2
//
// For the "details" key, the raw value "$info" is expected, no
// placeholders are involved here.
//
// Note that Lax mode is automatically enabled by SubJSONOf operator to
// simplify numeric tests.
//
// Comments can be embedded in JSON data:
//
//   td.Cmp(t, gotValue,
//     SubJSONOf(`
//   {
//     // A guy properties:
//     "fullname": "$name",  // The full name of the guy
//     "details":  "$$info", // Literally "$info", thanks to "$" escape
//     "age":      $2        /* The age of the guy:
//                              - placeholder unquoted, but could be without
//                                any change
//                              - to demonstrate a multi-lines comment */
//   }`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43)))                 // matches only $2
//
// Comments, like in go, have 2 forms. To quote the Go language specification:
//   - line comments start with the character sequence // and stop at the
//     end of the line.
//   - multi-lines comments start with the character sequence /* and stop
//     with the first subsequent character sequence */.
//
// Most operators can be directly embedded in SubJSONOf without requiring
// any placeholder.
//
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`
//   {
//     "fullname": HasPrefix("Foo"),
//     "age":      Between(41, 43),
//     "details":  SuperMapOf({
//       "address": NotEmpty(),
//       "car":     Any("Peugeot", "Tesla", "Jeep") // any of these
//     })
//   }`))
//
// Placeholders can be used anywhere, even in operators parameters as in:
//
//   td.Cmp(t, gotValue,
//     td.SubJSONOf(`{"fullname": HasPrefix($1), "bar": 42}`, "Zip"))
//
// A few notes about operators embedding:
//   - SubMapOf and SuperMapOf take only one parameter, a JSON object;
//   - the optional 3rd parameter of Between has to be specified as a string
//     and can be: "[]" or "BoundsInIn" (default), "[[" or "BoundsInOut",
//     "]]" or "BoundsOutIn", "][" or "BoundsOutOut";
//   - not all operators are embeddable only the following are;
//   - All, Any, ArrayEach, Bag, Between, Contains, ContainsKey, Empty, Gt,
//     Gte, HasPrefix, HasSuffix, Ignore, JSONPointer, Keys, Len, Lt, Lte,
//     MapEach, N, NaN, Nil, None, Not, NotAny, NotEmpty, NotNaN, NotNil,
//     NotZero, Re, ReAll, Set, SubBagOf, SubMapOf, SubSetOf, SuperBagOf,
//     SuperMapOf, SuperSetOf, Values and Zero.
//
// Operators taking no parameters can also be directly embedded in
// JSON data using $^OperatorName or "$^OperatorName" notation. They
// are named shortcut operators (they predate the above operators embedding
// but they subsist for compatibility):
//
//   td.Cmp(t, gotValue, td.SubJSONOf(`{"id": $1, "bar": 42}`, td.NotZero()))
//
// can be written as:
//
//   td.Cmp(t, gotValue, td.SubJSONOf(`{"id": $^NotZero, "bar": 42}`))
//
// or
//
//   td.Cmp(t, gotValue, td.SubJSONOf(`{"id": "$^NotZero", "bar": 42}`))
//
// As for placeholders, there is no differences between $^NotZero and
// "$^NotZero".
//
// The allowed shortcut operators follow:
//   - Empty    → $^Empty
//   - Ignore   → $^Ignore
//   - NaN      → $^NaN
//   - Nil      → $^Nil
//   - NotEmpty → $^NotEmpty
//   - NotNaN   → $^NotNaN
//   - NotNil   → $^NotNil
//   - NotZero  → $^NotZero
//   - Zero     → $^Zero
//
// TypeBehind method returns the map[string]interface{} type.
func SubJSONOf(expectedJSON interface{}, params ...interface{}) TestDeep {
	m := &tdMapJSON{
		tdMap: tdMap{
			tdExpectedType: tdExpectedType{
				base:         newBase(3),
				expectedType: reflect.TypeOf((map[string]interface{})(nil)),
			},
			kind: subMap,
		},
	}

	v, err := newJSONUnmarshaler(m.GetLocation()).unmarshal(expectedJSON, params)
	if err != nil {
		m.err = err
		return m
	}

	_, ok := v.(map[string]interface{})
	if !ok {
		m.err = ctxerr.OpBad("SubJSONOf", "SubJSONOf() only accepts JSON objects {…}")
		return m
	}

	m.expected = reflect.ValueOf(v)

	m.populateExpectedEntries(nil, m.expected)
	return m
}

// summary(SuperJSONOf): compares struct or map against JSON
// representation but with potentially extra entries
// input(SuperJSONOf): map,struct,ptr(ptr on map/struct)

// SuperJSONOf operator allows to compare the JSON representation of
// data against "expectedJSON". Unlike JSON operator, marshaled data
// must be a JSON object/map (aka {…}). "expectedJSON" can be a:
//
//   - string containing JSON data like `{"fullname":"Bob","age":42}`
//   - string containing a JSON filename, ending with ".json" (its
//     content is ioutil.ReadFile before unmarshaling)
//   - []byte containing JSON data
//   - io.Reader stream containing JSON data (is ioutil.ReadAll before
//     unmarshaling)
//
// JSON data contained in "expectedJSON" must be a JSON object/map
// (aka {…}) too. During a match, each expected entry should match in
// the compared map. But some entries in the compared map may not be
// expected.
//
//   type MyStruct struct {
//     Name string `json:"name"`
//     Age  int    `json:"age"`
//     City string `json:"city"`
//   }
//   got := MyStruct{
//     Name: "Bob",
//     Age:  42,
//     City: "TestCity",
//   }
//   td.Cmp(t, got, td.SuperJSONOf(`{"name": "Bob", "age": 42}`))  // succeeds
//   td.Cmp(t, got, td.SuperJSONOf(`{"name": "Bob", "zip": 666}`)) // fails, miss "zip"
//
// "expectedJSON" JSON value can contain placeholders. The "params"
// are for any placeholder parameters in "expectedJSON". "params" can
// contain TestDeep operators as well as raw values. A placeholder can
// be numeric like $2 or named like $name and always references an
// item in "params".
//
// Numeric placeholders reference the n'th "operators" item (starting
// at 1). Named placeholders are used with Tag operator as follows:
//
//   td.Cmp(t, gotValue,
//     SuperJSONOf(`{"fullname": $name, "age": $2, "gender": $3}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43),                  // matches only $2
//       "male"))                             // matches only $3
//
// Note that placeholders can be double-quoted as in:
//
//   td.Cmp(t, gotValue,
//     td.SuperJSONOf(`{"fullname": "$name", "age": "$2", "gender": "$3"}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43),                  // matches only $2
//       "male"))                             // matches only $3
//
// It makes no difference whatever the underlying type of the replaced
// item is (= double quoting a placeholder matching a number is not a
// problem). It is just a matter of taste, double-quoting placeholders
// can be preferred when the JSON data has to conform to the JSON
// specification, like when used in a ".json" file.
//
// SuperJSONOf does its best to convert back the JSON corresponding to a
// placeholder to the type of the placeholder or, if the placeholder
// is an operator, to the type behind the operator. Allowing to do
// things like:
//
//   td.Cmp(t, gotValue,
//     td.SuperJSONOf(`{"foo":$1}`, []int{1, 2, 3, 4}))
//   td.Cmp(t, gotValue,
//     td.SuperJSONOf(`{"foo":$1}`, []interface{}{1, 2, td.Between(2, 4), 4}))
//   td.Cmp(t, gotValue,
//     td.SuperJSONOf(`{"foo":$1}`, td.Between(27, 32)))
//
// Of course, it does this conversion only if the expected type can be
// guessed. In the case the conversion cannot occur, data is compared
// as is, in its freshly unmarshaled JSON form (so as bool, float64,
// string, []interface{}, map[string]interface{} or simply nil).
//
// Note "expectedJSON" can be a []byte, JSON filename or io.Reader:
//
//   td.Cmp(t, gotValue, td.SuperJSONOf("file.json", td.Between(12, 34)))
//   td.Cmp(t, gotValue, td.SuperJSONOf([]byte(`[1, $1, 3]`), td.Between(12, 34)))
//   td.Cmp(t, gotValue, td.SuperJSONOf(osFile, td.Between(12, 34)))
//
// A JSON filename ends with ".json".
//
// To avoid a legit "$" string prefix causes a bad placeholder error,
// just double it to escape it. Note it is only needed when the "$" is
// the first character of a string:
//
//   td.Cmp(t, gotValue,
//     td.SuperJSONOf(`{"fullname": "$name", "details": "$$info", "age": $2}`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43)))                 // matches only $2
//
// For the "details" key, the raw value "$info" is expected, no
// placeholders are involved here.
//
// Note that Lax mode is automatically enabled by SuperJSONOf operator to
// simplify numeric tests.
//
// Comments can be embedded in JSON data:
//
//   td.Cmp(t, gotValue,
//     td.SuperJSONOf(`
//   {
//     // A guy properties:
//     "fullname": "$name",  // The full name of the guy
//     "details":  "$$info", // Literally "$info", thanks to "$" escape
//     "age":      $2        /* The age of the guy:
//                              - placeholder unquoted, but could be without
//                                any change
//                              - to demonstrate a multi-lines comment */
//   }`,
//       td.Tag("name", td.HasPrefix("Foo")), // matches $1 and $name
//       td.Between(41, 43)))                 // matches only $2
//
// Comments, like in go, have 2 forms. To quote the Go language specification:
//   - line comments start with the character sequence // and stop at the
//     end of the line.
//   - multi-lines comments start with the character sequence /* and stop
//     with the first subsequent character sequence */.
//
// Most operators can be directly embedded in SuperJSONOf without requiring
// any placeholder.
//
//   td.Cmp(t, gotValue,
//     td.SuperJSONOf(`
//   {
//     "fullname": HasPrefix("Foo"),
//     "age":      Between(41, 43),
//     "details":  SuperMapOf({
//       "address": NotEmpty(),
//       "car":     Any("Peugeot", "Tesla", "Jeep") // any of these
//     })
//   }`))
//
// Placeholders can be used anywhere, even in operators parameters as in:
//
//   td.Cmp(t, gotValue, td.SuperJSONOf(`{"fullname": HasPrefix($1)}`, "Zip"))
//
// A few notes about operators embedding:
//   - SubMapOf and SuperMapOf take only one parameter, a JSON object;
//   - the optional 3rd parameter of Between has to be specified as a string
//     and can be: "[]" or "BoundsInIn" (default), "[[" or "BoundsInOut",
//     "]]" or "BoundsOutIn", "][" or "BoundsOutOut";
//   - not all operators are embeddable only the following are;
//   - All, Any, ArrayEach, Bag, Between, Contains, ContainsKey, Empty, Gt,
//     Gte, HasPrefix, HasSuffix, Ignore, JSONPointer, Keys, Len, Lt, Lte,
//     MapEach, N, NaN, Nil, None, Not, NotAny, NotEmpty, NotNaN, NotNil,
//     NotZero, Re, ReAll, Set, SubBagOf, SubMapOf, SubSetOf, SuperBagOf,
//     SuperMapOf, SuperSetOf, Values and Zero.
//
// Operators taking no parameters can also be directly embedded in
// JSON data using $^OperatorName or "$^OperatorName" notation. They
// are named shortcut operators (they predate the above operators embedding
// but they subsist for compatibility):
//
//   td.Cmp(t, gotValue, td.SuperJSONOf(`{"id": $1}`, td.NotZero()))
//
// can be written as:
//
//   td.Cmp(t, gotValue, td.SuperJSONOf(`{"id": $^NotZero}`))
//
// or
//
//   td.Cmp(t, gotValue, td.SuperJSONOf(`{"id": "$^NotZero"}`))
//
// As for placeholders, there is no differences between $^NotZero and
// "$^NotZero".
//
// The allowed shortcut operators follow:
//   - Empty    → $^Empty
//   - Ignore   → $^Ignore
//   - NaN      → $^NaN
//   - Nil      → $^Nil
//   - NotEmpty → $^NotEmpty
//   - NotNaN   → $^NotNaN
//   - NotNil   → $^NotNil
//   - NotZero  → $^NotZero
//   - Zero     → $^Zero
//
// TypeBehind method returns the map[string]interface{} type.
func SuperJSONOf(expectedJSON interface{}, params ...interface{}) TestDeep {
	m := &tdMapJSON{
		tdMap: tdMap{
			tdExpectedType: tdExpectedType{
				base:         newBase(3),
				expectedType: reflect.TypeOf((map[string]interface{})(nil)),
			},
			kind: superMap,
		},
	}

	v, err := newJSONUnmarshaler(m.GetLocation()).unmarshal(expectedJSON, params)
	if err != nil {
		m.err = err
		return m
	}

	_, ok := v.(map[string]interface{})
	if !ok {
		m.err = ctxerr.OpBad("SuperJSONOf", "SuperJSONOf() only accepts JSON objects {…}")
		return m
	}

	m.expected = reflect.ValueOf(v)

	m.populateExpectedEntries(nil, m.expected)
	return m
}

func (m *tdMapJSON) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	if m.err != nil {
		return ctx.CollectError(m.err)
	}

	err := gotViaJSON(ctx, &got)
	if err != nil {
		return ctx.CollectError(err)
	}

	// nil case
	if !got.IsValid() {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return ctx.CollectError(&ctxerr.Error{
			Message:  "values differ",
			Got:      types.RawString("null"),
			Expected: types.RawString("non-null"),
		})
	}

	ctx.BeLax = true

	return m.match(ctx, got)
}

func (m *tdMapJSON) String() string {
	if m.err != nil {
		return m.stringError()
	}

	return jsonStringify(m.GetLocation().Func, m.expected)
}

func (m *tdMapJSON) HandleInvalid() bool {
	return true
}
