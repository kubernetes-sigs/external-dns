/*
Copyright 2025 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package flags

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type badSetter struct{}

func (b *badSetter) Set(_ string) error { return errors.New("bad default") }

func TestKingpinBinderParsesAllTypes(t *testing.T) {
	app := kingpin.New("test", "")
	b := NewKingpinBinder(app)

	var (
		s    string
		bval bool
		d    time.Duration
		i    int
		i64  int64
		ss   []string
		e    string
	)

	b.StringVar("s", "string flag", "def", &s)
	b.BoolVar("b", "bool flag", true, &bval)
	b.DurationVar("d", "duration flag", 5*time.Second, &d)
	b.IntVar("i", "int flag", 7, &i)
	b.Int64Var("i64", "int64 flag", 9, &i64)
	b.StringsVar("ss", "strings flag", []string{"x"}, &ss)
	b.EnumVar("e", "enum flag", "a", &e, "a", "b")

	_, err := app.Parse([]string{"--s=abc", "--no-b", "--d=2s", "--i=42", "--i64=64", "--ss=one", "--ss=two", "--e=b"})
	require.NoError(t, err)

	assert.Equal(t, "abc", s)
	assert.False(t, bval)
	assert.Equal(t, 2*time.Second, d)
	assert.Equal(t, 42, i)
	assert.Equal(t, int64(64), i64)
	assert.ElementsMatch(t, []string{"one", "two"}, ss)
	assert.Equal(t, "b", e)
}

func TestKingpinBinderEnumValidation(t *testing.T) {
	app := kingpin.New("test", "")
	b := NewKingpinBinder(app)

	var e string
	b.EnumVar("e", "enum flag", "a", &e, "a", "b")

	_, err := app.Parse([]string{"--e=c"})
	require.Error(t, err)
}

func TestKingpinBinderStringsVarNoDefaultAndBoolDefaultFalse(t *testing.T) {
	app := kingpin.New("test", "")
	b := NewKingpinBinder(app)

	var (
		ss []string
		b2 bool
	)

	b.StringsVar("ss", "strings flag", nil, &ss)
	b.BoolVar("b2", "bool2 flag", false, &b2)

	_, err := app.Parse([]string{})
	require.NoError(t, err)

	assert.Empty(t, ss)
	assert.False(t, b2)
}

func TestCobraRegexValueSetStringType(t *testing.T) {
	var r *regexp.Regexp
	rv := &regexpValue{target: &r}

	require.Equal(t, "regexp", rv.Type())
	// empty when target nil
	assert.Empty(t, rv.String())

	// invalid pattern returns error
	err := rv.Set("(")
	require.Error(t, err)

	// valid pattern sets target
	err = rv.Set("^foo$")
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "^foo$", r.String())
	assert.Equal(t, "^foo$", rv.String())
}

func TestKingpinRegexpVarDefaultAndParse(t *testing.T) {
	app := kingpin.New("test", "")
	b := NewKingpinBinder(app)

	var r *regexp.Regexp
	b.RegexpVar("re", "help", regexp.MustCompile("^a+$"), &r)

	_, err := app.Parse([]string{})
	require.NoError(t, err)
	require.NotNil(t, r)
	assert.Equal(t, "^a+$", r.String())

	// user-provided value should override default
	var r2 *regexp.Regexp
	app2 := kingpin.New("test2", "")
	b2 := NewKingpinBinder(app2)
	b2.RegexpVar("re", "help", nil, &r2)
	_, err = app2.Parse([]string{"--re=^b+$"})
	require.NoError(t, err)
	require.NotNil(t, r2)
	assert.Equal(t, "^b+$", r2.String())
}

func TestKingpinStringsEnumVarWithAndWithoutDefault(t *testing.T) {
	app := kingpin.New("test", "")
	b := NewKingpinBinder(app)

	var vals []string
	b.StringsEnumVar("se", "help", []string{"a", "b"}, &vals, "a", "b", "c")
	_, err := app.Parse([]string{})
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"a", "b"}, vals)

	// without default
	app2 := kingpin.New("test2", "")
	b2 := NewKingpinBinder(app2)
	var vals2 []string
	b2.StringsEnumVar("se", "help", nil, &vals2, "a", "b", "c")
	_, err = app2.Parse([]string{"--se=a", "--se=c"})
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{"a", "c"}, vals2)
}

func TestSetRegexDefaultPanicsOnInvalidDefault(t *testing.T) {
	bs := &badSetter{}
	def := regexp.MustCompile("^")
	require.Panics(t, func() { setRegexpDefault(bs, def, "flag") })
}
