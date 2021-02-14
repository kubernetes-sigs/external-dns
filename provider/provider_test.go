/*
Copyright 2017 The Kubernetes Authors.

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

package provider

import (
	"io/ioutil"
	"os"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	os.Exit(m.Run())
}

func TestEnsureTrailingDot(t *testing.T) {
	for _, tc := range []struct {
		input, expected string
	}{
		{"example.org", "example.org."},
		{"example.org.", "example.org."},
		{"8.8.8.8", "8.8.8.8"},
	} {
		output := EnsureTrailingDot(tc.input)

		if output != tc.expected {
			t.Errorf("expected %s, got %s", tc.expected, output)
		}
	}
}

func TestDifference(t *testing.T) {
	current := []string{"foo", "bar"}
	desired := []string{"bar", "baz"}
	add, remove, leave := Difference(current, desired)
	assert.Equal(t, add, []string{"baz"})
	assert.Equal(t, remove, []string{"foo"})
	assert.Equal(t, leave, []string{"bar"})
}

func TestBaseProviderPropertyEquality(t *testing.T) {
	p := BaseProvider{}
	assert.True(t, p.PropertyValuesEqual("some.property", "", ""), "Both properties not present")
	assert.False(t, p.PropertyValuesEqual("some.property", "", "Foo"), "First property missing")
	assert.False(t, p.PropertyValuesEqual("some.property", "Foo", ""), "Second property missing")
	assert.True(t, p.PropertyValuesEqual("some.property", "Foo", "Foo"), "Properties the same")
	assert.False(t, p.PropertyValuesEqual("some.property", "Foo", "Bar"), "Attributes differ")
}
