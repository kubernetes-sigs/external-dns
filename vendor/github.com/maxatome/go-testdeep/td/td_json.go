// Copyright (c) 2019, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/dark"
	"github.com/maxatome/go-testdeep/internal/types"
	"github.com/maxatome/go-testdeep/internal/util"
)

var jsonEscaper = strings.NewReplacer(
	"\n", `<testdeep:nl>`,
	`\`, `<testdeep:bs>`,
	`"`, `<testdeep:qq>`,
)

var jsonUnescaper = strings.NewReplacer(
	`<testdeep:nl>`, "\n",
	`<testdeep:bs>`, `\`,
	`<testdeep:qq>`, `"`,
)

const (
	commentStart = `<testdeep:opOn>`
	commentEnd   = `<testdeep:opOff>"` // note final "

	commentStartRepl = `" /* `
	commentEndRepl   = ` */`
)

type tdJSONPlaceholder struct {
	TestDeep
	name string
	num  uint64
}

func (p *tdJSONPlaceholder) MarshalJSON() ([]byte, error) {
	var b bytes.Buffer

	if p.num == 0 {
		fmt.Fprintf(&b, `"$%s`, p.name)

		// Don't add a comment for operator shorcuts (aka $^NotZero)
		if p.name[0] == '^' {
			b.WriteByte('"')
			return b.Bytes(), nil
		}
	} else {
		fmt.Fprintf(&b, `"$%d`, p.num)
	}

	b.WriteString(commentStart)
	b.WriteString(jsonEscaper.Replace(util.ToString(p.TestDeep)))
	b.WriteString(commentEnd)

	return b.Bytes(), nil
}

type tdJSON struct {
	baseOKNil
	expected reflect.Value
}

var _ TestDeep = &tdJSON{}

func unmarshal(expectedJSON interface{}, target interface{}) {
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
				panic(fmt.Sprintf("JSON file %s cannot be read: %s", data, err))
			}
			break
		}
		b = []byte(data)

	case []byte:
		b = data

	case io.Reader:
		b, err = ioutil.ReadAll(data)
		if err != nil {
			panic(fmt.Sprintf("JSON read error: %s", err))
		}

	default:
		panic("usage: JSON(STRING_JSON|STRING_FILENAME|[]byte|io.Reader, ...)")
	}

	err = util.UnmarshalJSON(b, target)
	if err != nil {
		panic("JSON unmarshal error: " + err.Error())
	}
}

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

func getTags(params []interface{}) (byTag map[string]*tdTag) {
	for _, op := range params {
		if tag, ok := op.(*tdTag); ok {
			if byTag[tag.tag] != nil {
				panic(`2 params have the same tag "` + tag.tag + `"`)
			}
			if byTag == nil {
				byTag = map[string]*tdTag{}
			}
			byTag[tag.tag] = tag
		}
	}
	return
}

// scan scans "*v" data structure to find strings containing
// placeholders (like $123 or $name) corresponding to a value or
// TestDeep operator contained in "params" and "byTag".
func scan(v *interface{}, params []interface{}, byTag map[string]*tdTag, path string) {
	if *v == nil {
		return
	}

	switch tv := (*v).(type) {
	case map[string]interface{}:
		for k, v := range tv {
			scan(&v, params, byTag, path+`["`+k+`"]`)
			tv[k] = v
		}
	case []interface{}:
		for i := range tv {
			scan(&tv[i], params, byTag, path+"["+strconv.Itoa(i)+"]")
		}
	case string:
		if strings.HasPrefix(tv, "$") && len(tv) > 1 {
			// Double $$ at start of strings escape a $
			if strings.HasPrefix(tv[1:], "$") {
				*v = tv[1:]
				break
			}

			firstRune, _ := utf8.DecodeRuneInString(tv[1:])
			// Test for $123
			if firstRune >= '0' && firstRune <= '9' {
				np, err := strconv.ParseUint(tv[1:], 10, 64)
				if err != nil {
					panic(fmt.Sprintf(
						`JSON obj%s contains a bad numeric placeholder "%s"`,
						path, tv))
				}
				if np == 0 {
					panic(fmt.Sprintf(
						`JSON obj%s contains invalid numeric placeholder "%s", it should start at "$1"`,
						path, tv))
				}
				if np > uint64(len(params)) {
					panic(fmt.Sprintf(
						`JSON obj%s contains numeric placeholder "%s", but only %d params given`,
						path, tv, len(params)))
				}
				val := params[np-1]
				if op, ok := val.(TestDeep); ok {
					*v = &tdJSONPlaceholder{
						TestDeep: op,
						num:      np,
					}
				} else {
					*v = val
				}
				break
			}

			// Test for operator shortcut
			if firstRune == '^' {
				fn := jsonOpShortcuts[tv[2:]]
				if fn == nil {
					panic(fmt.Sprintf(`JSON obj%s contains a bad operator shortcut "%s"`,
						path, tv))
				}

				*v = &tdJSONPlaceholder{
					TestDeep: fn(),
					name:     tv[1:],
				}
				break
			}

			// Test for $tag
			err := util.CheckTag(tv[1:])
			if err != nil {
				panic(fmt.Sprintf(`JSON obj%s contains a bad placeholder "%s"`,
					path, tv))
			}
			op := byTag[tv[1:]]
			if op == nil {
				panic(fmt.Sprintf(`JSON obj%s contains a unknown placeholder "%s"`,
					path, tv))
			}
			*v = &tdJSONPlaceholder{
				TestDeep: op,
				name:     tv[1:],
			}
		}
	}
}

func gotViaJSON(ctx ctxerr.Context, pGot *reflect.Value) *ctxerr.Error {
	gotIf, ok := dark.GetInterface(*pGot, true)
	if !ok {
		return ctx.CannotCompareError()
	}

	b, err := json.Marshal(gotIf)
	if err != nil {
		if ctx.BooleanError {
			return ctxerr.BooleanError
		}
		return &ctxerr.Error{
			Message: "json.Marshal failed",
			Summary: ctxerr.NewSummary(err.Error()),
		}
	}

	// As Marshal succeeded, Unmarshal in an interface{} cannot fail
	var vgot interface{}
	json.Unmarshal(b, &vgot) //nolint: errcheck

	*pGot = reflect.ValueOf(vgot)
	return nil
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
// Last but not least, simple operators can be directly embedded in
// JSON data without requiring any placeholder but using directly
// $^OperatorName. They are operator shortcuts:
//
//   td.Cmp(t, gotValue, td.JSON(`{"id": $1}`, td.NotZero()))
//
// can be written as:
//
//   td.Cmp(t, gotValue, td.JSON(`{"id": $^NotZero}`))
//
// Unfortunately, only simple operators (in fact those which take no
// parameters) have shortcuts. They follow:
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
	var v interface{}
	unmarshal(expectedJSON, &v)

	scan(&v, params, getTags(params), "")

	return &tdJSON{
		baseOKNil: newBaseOKNil(3),
		expected:  reflect.ValueOf(v),
	}
}

func (j *tdJSON) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
	err := gotViaJSON(ctx, &got)
	if err != nil {
		return ctx.CollectError(err)
	}

	ctx.BeLax = true

	return deepValueEqual(ctx, got, j.expected)
}

func (j *tdJSON) String() string {
	return jsonStringify("JSON", j.expected)
}

func jsonStringify(opName string, v reflect.Value) string {
	if !v.IsValid() {
		return "JSON(null)"
	}

	var b bytes.Buffer

	b.WriteString(opName)
	b.WriteByte('(')

	enc := json.NewEncoder(&b)
	enc.SetEscapeHTML(false)
	enc.SetIndent(strings.Repeat(" ", len(opName)+1), "  ")

	// cannot return an error here
	enc.Encode(v.Interface()) //nolint: errcheck

	b.Truncate(b.Len() - 1)
	b.WriteByte(')')

	str := jsonUnescaper.Replace(b.String())

	b.Reset()
	for {
		beginPos := strings.Index(str, commentStart)
		if beginPos < 0 {
			break
		}
		endPos := strings.Index(str[beginPos+len(commentStart):], commentEnd)
		if endPos < 0 {
			break
		}

		b.WriteString(str[:beginPos])
		b.WriteString(commentStartRepl)

		// Multiline comment, compute indent level
		indent := ""
		if strings.ContainsRune(str[beginPos+len(commentStart):endPos+beginPos+len(commentStart)], '\n') {
			begLine := strings.LastIndexByte(str[:beginPos], '\n') + 1
			indent = strings.Repeat(" ", beginPos-begLine+len(commentStartRepl))
		}

		str = str[beginPos+len(commentStart):]

		// Multiline comment?
		if indent != "" {
			b.WriteString(strings.Replace(str[:endPos], "\n", "\n"+indent, -1))
		} else {
			b.WriteString(str[:endPos])
		}

		b.WriteString(commentEndRepl)

		str = str[endPos+len(commentEnd):]
	}

	if b.Len() == 0 {
		return str
	}

	b.WriteString(str)
	return b.String()
}

func (j *tdJSON) TypeBehind() reflect.Type {
	if j.expected.IsValid() {
		return j.expected.Type()
	}
	return interfaceInterface
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
// data against "expectedJSON". Unlike JSON operator, marshalled data
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
// Last but not least, simple operators can be directly embedded in
// JSON data without requiring any placeholder but using directly
// $^OperatorName. They are operator shortcuts:
//
//   td.Cmp(t, gotValue, td.SubJSONOf(`{"id": $1}`, td.NotZero()))
//
// can be written as:
//
//   td.Cmp(t, gotValue, td.SubJSONOf(`{"id": $^NotZero}`))
//
// Unfortunately, only simple operators (in fact those which take no
// parameters) have shortcuts. They follow:
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
	var v interface{}
	unmarshal(expectedJSON, &v)

	scan(&v, params, getTags(params), "")

	_, ok := v.(map[string]interface{})
	if !ok {
		panic("SubJSONOf only accepts JSON objects {…}")
	}

	m := tdMapJSON{
		tdMap: tdMap{
			tdExpectedType: tdExpectedType{
				base:         newBase(3),
				expectedType: reflect.TypeOf((map[string]interface{})(nil)),
			},
			kind: subMap,
		},
		expected: reflect.ValueOf(v),
	}
	m.populateExpectedEntries(nil, m.expected)

	return &m
}

// summary(SuperJSONOf): compares struct or map against JSON
// representation but with potentially extra entries
// input(SuperJSONOf): map,struct,ptr(ptr on map/struct)

// SuperJSONOf operator allows to compare the JSON representation of
// data against "expectedJSON". Unlike JSON operator, marshalled data
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
// Last but not least, simple operators can be directly embedded in
// JSON data without requiring any placeholder but using directly
// $^OperatorName. They are operator shortcuts:
//
//   td.Cmp(t, gotValue, td.SuperJSONOf(`{"id": $1}`, td.NotZero()))
//
// can be written as:
//
//   td.Cmp(t, gotValue, td.SuperJSONOf(`{"id": $^NotZero}`))
//
// Unfortunately, only simple operators (in fact those which take no
// parameters) have shortcuts. They follow:
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
	var v interface{}
	unmarshal(expectedJSON, &v)

	scan(&v, params, getTags(params), "")

	_, ok := v.(map[string]interface{})
	if !ok {
		panic("SuperJSONOf only accepts JSON objects {…}")
	}

	m := tdMapJSON{
		tdMap: tdMap{
			tdExpectedType: tdExpectedType{
				base:         newBase(3),
				expectedType: reflect.TypeOf((map[string]interface{})(nil)),
			},
			kind: superMap,
		},
		expected: reflect.ValueOf(v),
	}
	m.populateExpectedEntries(nil, m.expected)

	return &m
}

func (m *tdMapJSON) Match(ctx ctxerr.Context, got reflect.Value) *ctxerr.Error {
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
	return jsonStringify(m.GetLocation().Func, m.expected)
}

func (m *tdMapJSON) HandleInvalid() bool {
	return true
}
