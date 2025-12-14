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
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
)

type boolPairTestCase struct {
	name            string
	defaultVal      bool
	flags           []string
	expectedResult  bool
	expectedErrText string
}

func TestBooleanPair(t *testing.T) {
	cases := []*boolPairTestCase{
		{"foo", true, []string{}, true, ""},
		{"foo", true, []string{"--foo"}, true, ""},
		{"foo", true, []string{"--no-foo"}, false, ""},
		{"foo", false, []string{"--foo"}, true, ""},
		{"foo", false, []string{}, false, ""},
		{"foo", false, []string{"--no-foo"}, false, ""},
		{"foo", true, []string{"--foo", "--no-foo"}, false, "only one of"},
		{"foo", true, []string{"--no-foo", "--foo"}, false, "only one of"},
		// Disallow confusing "false" value.
		{"foo", true, []string{"--foo=false"}, false, "use --no-foo instead of providing false to --foo"},
		{"foo", true, []string{"--no-foo=false"}, false, "use --foo instead of providing false to --no-foo"},

		// Ensure tests still pass if positive sorts after no- alphabetically.
		{"zoo", true, []string{}, true, ""},
		{"zoo", true, []string{"--zoo"}, true, ""},
		{"zoo", true, []string{"--no-zoo"}, false, ""},
		{"zoo", false, []string{"--zoo"}, true, ""},
		{"zoo", false, []string{}, false, ""},
		{"zoo", false, []string{"--no-zoo"}, false, ""},
		{"zoo", true, []string{"--zoo", "--no-zoo"}, false, "only one of"},
		{"zoo", true, []string{"--no-zoo", "--zoo"}, false, "only one of"},
		// Disallow confusing "false" value.
		{"zoo", true, []string{"--zoo=false"}, false, "use --no-zoo instead of providing false to --zoo"},
		{"zoo", true, []string{"--no-zoo=false"}, false, "use --zoo instead of providing false to --no-zoo"},
	}
	for _, c := range cases {
		f := &pflag.FlagSet{}
		var result bool
		AddNegationToBoolFlags(f, &result, c.name, "", c.defaultVal, "set "+c.name)
		_ = f.Parse(c.flags)
		err := ReconcileBoolFlags(f)
		if c.expectedErrText != "" {
			require.ErrorContains(t, err, c.expectedErrText)
		} else {
			require.Equal(t, result, c.expectedResult)
		}
	}
}
