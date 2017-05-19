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

package source

import "testing"

func TestStore(t *testing.T) {
	t.Run("RegisterAndLookup", testRegisterAndLookup)
	t.Run("RegisterAndLookupFunc", testRegisterAndLookupFunc)
	t.Run("LookupMultiple", testLookupMultiple)
}

// testRegisterAndLookup tests that a Source can be registered and looked up by name.
func testRegisterAndLookup(t *testing.T) {
	for _, tc := range []struct {
		title   string
		sources map[string]Source
	}{
		{
			"registered source is found by name",
			map[string]Source{
				"foo": NewMockSource(nil),
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			// Clear already registered sources.
			Clear()

			// Register the source objects under test.
			for k, v := range tc.sources {
				Register(k, v)
			}

			// Validate that we can lookup the registered sources by name.
			for k, v := range tc.sources {
				src, err := Lookup(k, nil)
				if err != nil {
					t.Fatal(err)
				}

				if src != v {
					t.Errorf("expected %#v, got %#v", v, src)
				}
			}
		})
	}
}

// testRegisterAndLookupFunc tests that a Source can be registered and looked up
// by name via constructor functions.
func testRegisterAndLookupFunc(t *testing.T) {
	for _, tc := range []struct {
		title   string
		sources map[string]Source
	}{
		{
			"registered source is found by name via constructor func",
			map[string]Source{
				"foo": NewMockSource(nil),
			},
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			// Clear already registered sources.
			Clear()

			// Register the source objects under test via contructor functions.
			for k, v := range tc.sources {
				RegisterFunc(k, func(_ *Config) (Source, error) { return v, nil })
			}

			// Validate that we can lookup the registered sources by name.
			for k, v := range tc.sources {
				src, err := Lookup(k, nil)
				if err != nil {
					t.Fatal(err)
				}

				if src != v {
					t.Errorf("expected %#v, got %#v", v, src)
				}
			}
		})
	}
}

// testLookupMultiple tests that Sources can be looked up by providing multiple names.
func testLookupMultiple(t *testing.T) {
	for _, tc := range []struct {
		title       string
		sources     map[string]Source
		names       []string
		expectError bool
	}{
		{
			"multiple registered sources are found by names",
			map[string]Source{
				"foo": NewMockSource(nil),
				"bar": NewMockSource(nil),
			},
			[]string{"foo", "bar"},
			false,
		},
		{
			"multiple registered sources, one source not registered",
			map[string]Source{
				"foo": NewMockSource(nil),
			},
			[]string{"foo", "bar"},
			true,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			// Clear already registered sources.
			Clear()

			// Register the source objects under test.
			for k, v := range tc.sources {
				Register(k, v)
			}

			// Lookup multiple sources by names.
			lookup, err := LookupMultiple(tc.names, nil)
			if !tc.expectError && err != nil {
				t.Fatal(err)
			}

			if tc.expectError {
				if err == nil {
					t.Fatal("look up should fail if source not registered")
				}
				t.Skip()
			}

			// Validate that we can lookup the registered sources by name.
			for i, name := range tc.names {
				if lookup[i] != tc.sources[name] {
					t.Errorf("expected %#v, got %#v", tc.sources[name], lookup[i])
				}
			}
		})
	}
}
