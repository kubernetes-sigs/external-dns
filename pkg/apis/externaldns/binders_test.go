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

package externaldns

import (
	"errors"
	"io"
	"regexp"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type badSetter struct{}

func (b *badSetter) Set(s string) error { return errors.New("bad default") }

func TestCobraBinderParsesAllTypes(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	b := NewCobraBinder(cmd)

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

	cmd.SetArgs([]string{"--s=abc", "--b=false", "--d=2s", "--i=42", "--i64=64", "--ss=one", "--ss=two", "--e=b"})
	err := cmd.Execute()
	require.NoError(t, err)

	assert.Equal(t, "abc", s)
	assert.False(t, bval)
	assert.Equal(t, 2*time.Second, d)
	assert.Equal(t, 42, i)
	assert.Equal(t, int64(64), i64)
	assert.ElementsMatch(t, []string{"one", "two"}, ss)
	assert.Equal(t, "b", e)
}

func TestCobraBinderEnumNotValidatedHere(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	b := NewCobraBinder(cmd)

	var e string
	b.EnumVar("e", "enum flag", "a", &e, "a", "b")

	cmd.SetArgs([]string{"--e=c"})
	err := cmd.Execute()
	require.NoError(t, err)
	assert.Equal(t, "c", e)
}

// Cobra requires --<flag>=false
func TestCobraBinderNoBooleanNegationFormUnsupported(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.SetErr(io.Discard)
	cmd.SetOut(io.Discard)
	b := NewCobraBinder(cmd)

	var v bool
	b.BoolVar("v", "bool flag", true, &v)

	cmd.SetArgs([]string{"--no-v"})
	err := cmd.Execute()
	require.Error(t, err)
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

func TestCobraRegexpVarDefaultAndInvalidValue(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	cmd.SetErr(io.Discard)
	cmd.SetOut(io.Discard)
	b := NewCobraBinder(cmd)

	var r *regexp.Regexp
	// Provide a valid default: should set target immediately
	b.RegexpVar("re", "help", regexp.MustCompile("^x+$"), &r)
	require.NotNil(t, r)
	assert.Equal(t, "^x+$", r.String())

	// Executing with an invalid value should produce an error
	cmd2 := &cobra.Command{Use: "test2"}
	cmd2.SetErr(io.Discard)
	cmd2.SetOut(io.Discard)
	b2 := NewCobraBinder(cmd2)
	var r2 *regexp.Regexp
	b2.RegexpVar("re", "help", nil, &r2)
	cmd2.SetArgs([]string{"--re=invalid("})
	err := cmd2.Execute()
	require.Error(t, err)
}

func TestCobraStringMapVarDefaultEmpty(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	b := NewCobraBinder(cmd)

	var m map[string]string
	b.StringMapVar("m", "help", &m)

	cmd.SetArgs([]string{})
	err := cmd.Execute()
	require.NoError(t, err)
	require.NotNil(t, m)
	assert.Empty(t, m)
}

func TestCobraStringsEnumVarWithAndWithoutDefault(t *testing.T) {
	cmd := &cobra.Command{Use: "test"}
	b := NewCobraBinder(cmd)

	var vals []string
	b.StringsEnumVar("se", "help", []string{"x", "y"}, &vals, "x", "y")
	cmd.SetArgs([]string{})
	require.NoError(t, cmd.Execute())
	assert.ElementsMatch(t, []string{"x", "y"}, vals)

	// without default
	cmd2 := &cobra.Command{Use: "test2"}
	b2 := NewCobraBinder(cmd2)
	var vals2 []string
	b2.StringsEnumVar("se", "help", nil, &vals2, "x", "y")
	cmd2.SetArgs([]string{"--se=a", "--se=b"})
	require.NoError(t, cmd2.Execute())
	assert.ElementsMatch(t, []string{"a", "b"}, vals2)
}

func TestSetRegexDefaultPanicsOnInvalidDefault(t *testing.T) {
	bs := &badSetter{}
	def := regexp.MustCompile("^")
	require.Panics(t, func() { setRegexpDefault(bs, def, "flag") })
}
