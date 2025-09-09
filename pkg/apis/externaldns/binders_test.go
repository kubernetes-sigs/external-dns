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
	"testing"
	"time"

	"github.com/alecthomas/kingpin/v2"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
