// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

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
	name       string
	expected   reflect.Value
	index      []int
	unexported bool
}

type fieldInfoSlice []fieldInfo

func (e fieldInfoSlice) Len() int           { return len(e) }
func (e fieldInfoSlice) Less(i, j int) bool { return e[i].name < e[j].name }
func (e fieldInfoSlice) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

type fieldMatcher struct {
	name     string
	match    func(string) (bool, error)
	expected any
	order    int
	ok       bool
}

var (
	reMatcherOnce  sync.Once
	reMatcher      *regexp.Regexp
	errNotAMatcher = errors.New("Not a matcher")
)

// parseMatcher parses " [NUM] OP PATTERN " and returns 3 strings
// corresponding to each part or nil if "s" is not a matcher.
func parseMatcher(s string) []string {
	reMatcherOnce.Do(func() {
		reMatcher = regexp.MustCompile(`^(?:(\d+)\s*)?([=!]~?)\s*(.+)`)
	})
	subs := reMatcher.FindStringSubmatch(strings.TrimSpace(s))
	if subs != nil {
		subs = subs[1:]
	}
	return subs
}

// newFieldMatcher checks name matches "[NUM] OP PATTERN" where NUM
// is an optional number used to sort patterns, OP is "=~", "!~", "="
// or "!" and PATTERN is a regexp (when OP is either "=~" or "!~") or
// a shell pattern (when OP is either "=" or "!").
//
// NUM, OP and PATTERN can be separated by spaces (or not).
func newFieldMatcher(name string, expected any) (fieldMatcher, error) {
	subs := parseMatcher(name)
	if subs == nil {
		return fieldMatcher{}, errNotAMatcher
	}

	fm := fieldMatcher{
		name:     name,
		expected: expected,
		ok:       subs[1][0] == '=',
	}

	if subs[0] != "" {
		fm.order, _ = strconv.Atoi(subs[0]) //nolint: errcheck
	}

	// Shell pattern
	if subs[1] == "=" || subs[1] == "!" {
		pattern := subs[2]
		fm.match = func(s string) (bool, error) {
			return path.Match(pattern, s)
		}
		return fm, nil
	}

	// Regexp
	r, err := regexp.Compile(subs[2])
	if err != nil {
		return fieldMatcher{}, fmt.Errorf("bad regexp field %#q: %s", name, err)
	}
	fm.match = func(s string) (bool, error) {
		return r.MatchString(s), nil
	}
	return fm, nil
}

type fieldMatcherSlice []fieldMatcher

func (m fieldMatcherSlice) Len() int { return len(m) }
func (m fieldMatcherSlice) Less(i, j int) bool {
	if m[i].order != m[j].order {
		return m[i].order < m[j].order
	}
	return m[i].name < m[j].name
}
func (m fieldMatcherSlice) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

// StructFields allows to pass struct fields to check in functions
// [Struct] and [SStruct]. It is a map whose each key is the expected
// field name (or a regexp or a shell pattern matching a field name,
// see [Struct] & [SStruct] docs for details) and the corresponding
// value the expected field value (which can be a [TestDeep] operator
// as well as a zero value.)
type StructFields map[string]any

// canonStructField canonicalizes name, a key in a StructFields map,
// so it can be compared with other keys during a mergeStructFields().
//   - "name"                 → "name"
//   - ">  name  "            → ">name"
//   - "  22 =~ [A-Z].*At$  " → "22=~[A-Z].*At$"
func canonStructField(name string) string {
	r, _ := utf8.DecodeRuneInString(name)
	if r == utf8.RuneError || unicode.IsLetter(r) {
		return name // shortcut
	}

	// Overwrite a field
	if strings.HasPrefix(name, ">") {
		nn := strings.TrimSpace(name[1:])
		if 1+len(nn) == len(name) {
			return name // already canonicalized
		}
		return ">" + nn
	}

	// Matcher
	if subs := parseMatcher(name); subs != nil {
		if len(subs[0])+len(subs[1])+len(subs[2]) == len(name) {
			return name // already canonicalized
		}
		return subs[0] + subs[1] + subs[2]
	}

	// Will probably raise an error later as it cannot be a field, not
	// an overwritter and not a matcher
	return name
}

// mergeStructFields merges all sfs items into one StructFields and
// returns it.
func mergeStructFields(sfs ...StructFields) StructFields {
	switch len(sfs) {
	case 0:
		return nil

	case 1:
		return sfs[0]

	default:
		// Do a smart merge so ">  pipo" replaces ">pipo  " for example.
		canon2field := map[string]string{}
		ret := make(StructFields, len(sfs[0]))
		for _, sf := range sfs {
			for field, value := range sf {
				canon := canonStructField(field)
				if prevField, ok := canon2field[canon]; ok {
					delete(ret, prevField)
					delete(canon2field, canon)
				} else {
					delete(ret, canon)
				}
				if canon != field {
					canon2field[canon] = field
				}
				ret[field] = value
			}
		}
		return ret
	}
}

func newStruct(base base, vmodel reflect.Value) (*tdStruct, reflect.Value) {
	st := tdStruct{
		tdExpectedType: tdExpectedType{
			base: base,
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

	st.err = ctxerr.OpBadUsage(st.location.Func,
		"(STRUCT|&STRUCT|nil, EXPECTED_FIELDS)",
		vmodel.Interface(), 1, true)
	return &st, reflect.Value{}
}

// structTypeString returns stringified t. It is the caller
// responsibility to check t is a struct type.
//   - struct{}   → "struct {}"
//   - pkg.MyType → "struct pkg.MyType"
func structTypeString(t reflect.Type) string {
	if t.Name() == "" {
		return t.String()
	}
	return "struct " + t.String()
}

func anyStruct(base base, model reflect.Value, expectedFields StructFields, strict bool) *tdStruct {
	st, vmodel := newStruct(base, model)
	if st.err != nil {
		return st
	}

	st.expectedFields = make([]fieldInfo, 0, len(expectedFields))
	checkedFields := make(map[string]bool, len(expectedFields))
	var matchers fieldMatcherSlice //nolint: prealloc

	// Check that all given fields are available in model
	stType := st.expectedType
	for fieldName, expectedValue := range expectedFields {
		field, found := stType.FieldByName(fieldName)
		if found {
			st.addExpectedValue(field, expectedValue, "")
			checkedFields[fieldName] = false
			continue
		}

		// overwrite model field: ">fieldName", "> fieldName"
		if strings.HasPrefix(fieldName, ">") {
			name := strings.TrimSpace(fieldName[1:])
			field, found = stType.FieldByName(name)
			if !found {
				st.err = ctxerr.OpBad(st.location.Func,
					"%s has no field %q (from %q)", structTypeString(stType), name, fieldName)
				return st
			}
			st.addExpectedValue(
				field, expectedValue,
				fmt.Sprintf(" (from %q)", fieldName))
			checkedFields[name] = true
			continue
		}

		// matcher: "=~At$", "!~At$", "=*At", "!*At"
		matcher, err := newFieldMatcher(fieldName, expectedValue)
		if err != nil {
			if err == errNotAMatcher {
				st.err = ctxerr.OpBad(st.location.Func,
					"%s has no field %q", structTypeString(stType), fieldName)
			} else {
				st.err = ctxerr.OpBad(st.location.Func, err.Error())
			}
			return st
		}
		matchers = append(matchers, matcher)
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
			overwrite, alreadySet := checkedFields[fieldName]
			if overwrite {
				continue
			}

			field, _ := stType.FieldByName(fieldName)
			if field.Anonymous {
				continue
			}

			vfield := vmodel.FieldByIndex(field.Index)

			// Try to force access to unexported fields
			fieldIf, ok := dark.GetInterface(vfield, true)
			if !ok {
				// Probably in an environment where "unsafe" package is forbidden… :(
				fmt.Fprintf(os.Stderr, //nolint: errcheck
					"%s(): field %s is unexported and cannot be overridden, skip it from model.\n",
					st.location.Func,
					fieldName)
				continue
			}

			// If non-zero field
			if !reflect.DeepEqual(reflect.Zero(field.Type).Interface(), fieldIf) {
				if alreadySet {
					st.err = ctxerr.OpBad(st.location.Func,
						"non zero field %s in model already exists in expectedFields",
						fieldName)
					return st
				}

				st.expectedFields = append(st.expectedFields, fieldInfo{
					name:       fieldName,
					expected:   vfield,
					index:      field.Index,
					unexported: field.PkgPath != "",
				})
				checkedFields[fieldName] = true
			}
		}
	}

	// At least one matcher (regexp/shell pattern)
	if matchers != nil {
		sort.Sort(matchers) // always process matchers in the same order
		for _, m := range matchers {
			for fieldName := range allFields {
				if _, ok := checkedFields[fieldName]; ok {
					continue
				}
				field, _ := stType.FieldByName(fieldName)
				if field.Anonymous {
					continue
				}

				ok, err := m.match(fieldName)
				if err != nil {
					st.err = ctxerr.OpBad(st.location.Func,
						"bad shell pattern field %#q: %s", m.name, err)
					return st
				}
				if ok == m.ok {
					st.addExpectedValue(
						field, m.expected,
						fmt.Sprintf(" (from pattern %#q)", m.name))
					checkedFields[fieldName] = true
				}
			}
		}
	}

	// If strict, fill non explicitly expected fields to zero
	if strict {
		for fieldName := range allFields {
			if _, ok := checkedFields[fieldName]; ok {
				continue
			}

			field, _ := stType.FieldByName(fieldName)
			if field.Anonymous {
				continue
			}

			st.expectedFields = append(st.expectedFields, fieldInfo{
				name:       fieldName,
				expected:   reflect.New(field.Type).Elem(), // zero
				index:      field.Index,
				unexported: field.PkgPath != "",
			})
		}
	}

	sort.Sort(st.expectedFields)

	return st
}

func (s *tdStruct) addExpectedValue(field reflect.StructField, expectedValue any, ctxInfo string) {
	var vexpectedValue reflect.Value
	if expectedValue == nil {
		switch field.Type.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice:
			vexpectedValue = reflect.Zero(field.Type) // change to a typed nil
			// default:
			//   Don't raise an error if field cannot be nil as a smuggle hook can
			//   change it at fly during the comparison
		}
	} else {
		vexpectedValue = reflect.ValueOf(expectedValue)
		// Don't check vexpectedValue type against field one as a
		// smuggle hook can change it at fly during the comparison
	}

	s.expectedFields = append(s.expectedFields, fieldInfo{
		name:       field.Name + ctxInfo,
		expected:   vexpectedValue,
		index:      field.Index,
		unexported: field.PkgPath != "",
	})
}

// summary(Struct): compares the contents of a struct or a pointer on
// a struct
// input(Struct): struct,ptr(ptr on struct)

// Struct operator compares the contents of a struct or a pointer on a
// struct against the non-zero values of model (if any) and the
// values of expectedFields. See [SStruct] to compares against zero
// fields without specifying them in expectedFields.
//
// model must be the same type as compared data. If the expected type
// is anonymous or private, model can be nil. In this case it is
// considered lazy and determined each time the operator is involved
// in a match, see below.
//
// expectedFields can be omitted, if no zero entries are expected
// and no [TestDeep] operators are involved. If expectedFields
// contains more than one item, all items are merged before their use,
// from left to right.
//
//	td.Cmp(t, got, td.Struct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    "Children": 4,
//	  },
//	  td.StructFields{
//	    "Age":      td.Between(40, 45),
//	    "Children": 0, // overwrite 4
//	  }),
//	)
//
// It is an error to set a non-zero field in model AND to set the
// same field in expectedFields, as in such cases the Struct
// operator does not know if the user wants to override the non-zero
// model field value or if it is an error. To explicitly override a
// non-zero model in expectedFields, just prefix its name with a
// ">" (followed by some optional spaces), as in:
//
//	td.Cmp(t, got, td.Struct(
//	  Person{
//	    Name:     "John Doe",
//	    Age:      23,
//	    Children: 4,
//	  },
//	  td.StructFields{
//	    "> Age":     td.Between(40, 45),
//	    ">Children": 0, // spaces after ">" are optional
//	  }),
//	)
//
// expectedFields can also contain regexps or shell patterns to
// match multiple fields not explicitly listed in model and in
// expectedFields. Regexps are prefixed by "=~" or "!~" to
// respectively match or don't-match. Shell patterns are prefixed by "="
// or "!" to respectively match or don't-match.
//
//	td.Cmp(t, got, td.Struct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    "=*At":     td.Lte(time.Now()), // matches CreatedAt & UpdatedAt fields using shell pattern
//	    "=~^[a-z]": td.Ignore(),        // explicitly ignore private fields using a regexp
//	  }),
//	)
//
// When several patterns can match a same field, it is advised to tell
// go-testdeep in which order patterns should be tested, as once a
// pattern matches a field, the other patterns are ignored for this
// field. To do so, each pattern can be prefixed by a number, as in:
//
//	td.Cmp(t, got, td.Struct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    "1=*At":     td.Lte(time.Now()),
//	    "2=~^[a-z]": td.NotNil(),
//	  }),
//	)
//
// This way, "*At" shell pattern is always used before "^[a-z]"
// regexp, so if a field "createdAt" exists it is tested against
// time.Now() and never against [NotNil]. A pattern without a
// prefix number is the same as specifying "0" as prefix.
//
// To make it clearer, some spaces can be added, as well as bigger
// numbers used:
//
//	td.Cmp(t, got, td.Struct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    " 900 =  *At":    td.Lte(time.Now()),
//	    "2000 =~ ^[a-z]": td.NotNil(),
//	  }),
//	)
//
// The following example combines all possibilities:
//
//	td.Cmp(t, got, td.Struct(
//	  Person{
//	    NickName: "Joe",
//	  },
//	  td.StructFields{
//	    "Firstname":               td.Any("John", "Johnny"),
//	    "1 =  *[nN]ame":           td.NotEmpty(), // matches LastName, lastname, …
//	    "2 !  [A-Z]*":             td.NotZero(),  // matches all private fields
//	    "3 =~ ^(Crea|Upda)tedAt$": td.Gte(time.Now()),
//	    "4 !~ ^(Dogs|Children)$":  td.Zero(),   // matches all remaining fields except Dogs and Children
//	    "5 =~ .":                  td.NotNil(), // matches all remaining fields (same as "5 = *")
//	  }),
//	)
//
// If the expected type is private to the current package, it cannot
// be passed as model. To overcome this limitation, model can be nil,
// it is then considered as lazy. This way, the model is automatically
// set during each match to the same type (still requiring struct or
// struct pointer) of the compared data. Similarly, testing an
// anonymous struct can be boring as all fields have to be re-declared
// to define model. A nil model avoids that:
//
//	got := struct {
//	  name string
//	  age  int
//	}{"Bob", 42}
//	td.Cmp(t, got, td.Struct(nil, td.StructFields{"age": td.Between(40, 42)}))
//
// During a match, all expected fields must be found to
// succeed. Non-expected fields (and so zero model fields) are
// ignored.
//
// TypeBehind method returns the [reflect.Type] of model.
//
// See also [SStruct].
func Struct(model any, expectedFields ...StructFields) TestDeep {
	ef := mergeStructFields(expectedFields...)
	if model == nil {
		return newStructLazy(ef, false)
	}
	return anyStruct(newBase(3), reflect.ValueOf(model), ef, false)
}

// summary(SStruct): strictly compares the contents of a struct or a
// pointer on a struct
// input(SStruct): struct,ptr(ptr on struct)

// SStruct operator (aka strict-[Struct]) compares the contents of a
// struct or a pointer on a struct against values of model (if any)
// and the values of expectedFields. The zero values are compared
// too even if they are omitted from expectedFields: that is the
// difference with [Struct] operator.
//
// model must be the same type as compared data. If the expected type
// is private or anonymous, model can be nil. In this case it is
// considered lazy and determined each time the operator is involved
// in a match, see below.
//
// expectedFields can be omitted, if no [TestDeep] operators are
// involved. If expectedFields contains more than one item, all
// items are merged before their use, from left to right.
//
// To ignore a field, one has to specify it in expectedFields and
// use the [Ignore] operator.
//
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    "Children": 4,
//	  },
//	  td.StructFields{
//	    "Age":      td.Between(40, 45),
//	    "Children": td.Ignore(), // overwrite 4
//	  }),
//	)
//
// It is an error to set a non-zero field in model AND to set the
// same field in expectedFields, as in such cases the SStruct
// operator does not know if the user wants to override the non-zero
// model field value or if it is an error. To explicitly override a
// non-zero model in expectedFields, just prefix its name with a
// ">" (followed by some optional spaces), as in:
//
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    Name:     "John Doe",
//	    Age:      23,
//	    Children: 4,
//	  },
//	  td.StructFields{
//	    "> Age":     td.Between(40, 45),
//	    ">Children": 0, // spaces after ">" are optional
//	  }),
//	)
//
// expectedFields can also contain regexps or shell patterns to
// match multiple fields not explicitly listed in model and in
// expectedFields. Regexps are prefixed by "=~" or "!~" to
// respectively match or don't-match. Shell patterns are prefixed by "="
// or "!" to respectively match or don't-match.
//
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    "=*At":     td.Lte(time.Now()), // matches CreatedAt & UpdatedAt fields using shell pattern
//	    "=~^[a-z]": td.Ignore(),        // explicitly ignore private fields using a regexp
//	  }),
//	)
//
// When several patterns can match a same field, it is advised to tell
// go-testdeep in which order patterns should be tested, as once a
// pattern matches a field, the other patterns are ignored for this
// field. To do so, each pattern can be prefixed by a number, as in:
//
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    "1=*At":     td.Lte(time.Now()),
//	    "2=~^[a-z]": td.NotNil(),
//	  }),
//	)
//
// This way, "*At" shell pattern is always used before "^[a-z]"
// regexp, so if a field "createdAt" exists it is tested against
// time.Now() and never against [NotNil]. A pattern without a
// prefix number is the same as specifying "0" as prefix.
//
// To make it clearer, some spaces can be added, as well as bigger
// numbers used:
//
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    Name: "John Doe",
//	  },
//	  td.StructFields{
//	    " 900 =  *At":    td.Lte(time.Now()),
//	    "2000 =~ ^[a-z]": td.NotNil(),
//	  }),
//	)
//
// The following example combines all possibilities:
//
//	td.Cmp(t, got, td.SStruct(
//	  Person{
//	    NickName: "Joe",
//	  },
//	  td.StructFields{
//	    "Firstname":               td.Any("John", "Johnny"),
//	    "1 =  *[nN]ame":           td.NotEmpty(), // matches LastName, lastname, …
//	    "2 !  [A-Z]*":             td.NotZero(),  // matches all private fields
//	    "3 =~ ^(Crea|Upda)tedAt$": td.Gte(time.Now()),
//	    "4 !~ ^(Dogs|Children)$":  td.Zero(),   // matches all remaining fields except Dogs and Children
//	    "5 =~ .":                  td.NotNil(), // matches all remaining fields (same as "5 = *")
//	  }),
//	)
//
// If the expected type is private to the current package, it cannot
// be passed as model. To overcome this limitation, model can be nil,
// it is then considered as lazy. This way, the model is automatically
// set during each match to the same type (still requiring struct or
// struct pointer) of the compared data. Similarly, testing an
// anonymous struct can be boring as all fields have to be re-declared
// to define model. A nil model avoids that:
//
//	got := struct {
//	  name string
//	  age  int
//	}{"Bob", 42}
//	td.Cmp(t, got, td.SStruct(nil, td.StructFields{
//	  "name": "Bob",
//	  "age":  td.Between(40, 42),
//	}))
//
// During a match, all expected and zero fields must be found to
// succeed.
//
// TypeBehind method returns the [reflect.Type] of model.
//
// See also [SStruct].
func SStruct(model any, expectedFields ...StructFields) TestDeep {
	ef := mergeStructFields(expectedFields...)
	if model == nil {
		return newStructLazy(ef, false)
	}
	return anyStruct(newBase(3), reflect.ValueOf(model), ef, true)
}

func (s *tdStruct) Match(ctx ctxerr.Context, got reflect.Value) (err *ctxerr.Error) {
	if s.err != nil {
		return ctx.CollectError(s.err)
	}

	err = s.checkPtr(ctx, &got, false)
	if err != nil {
		return ctx.CollectError(err)
	}

	err = s.checkType(ctx, got)
	if err != nil {
		return ctx.CollectError(err)
	}

	ignoreUnexported := ctx.IgnoreUnexported || ctx.Hooks.IgnoreUnexported(got.Type())

	for _, fieldInfo := range s.expectedFields {
		if ignoreUnexported && fieldInfo.unexported {
			continue
		}
		err = deepValueEqual(ctx.AddField(fieldInfo.name),
			got.FieldByIndex(fieldInfo.index), fieldInfo.expected)
		if err != nil {
			return
		}
	}
	return nil
}

func (s *tdStruct) String() string {
	if s.err != nil {
		return s.stringError()
	}

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

		maxLen := 0
		for _, fieldInfo := range s.expectedFields {
			if len(fieldInfo.name) > maxLen {
				maxLen = len(fieldInfo.name)
			}
		}
		maxLen++
		for _, fieldInfo := range s.expectedFields {
			fmt.Fprintf(buf, "  %-*s %s\n", //nolint: errcheck
				maxLen, fieldInfo.name+":", util.ToString(fieldInfo.expected))
		}

		buf.WriteString("})")
	}

	return buf.String()
}
