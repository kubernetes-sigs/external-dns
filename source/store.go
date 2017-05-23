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

import "fmt"

var store = map[string]Source{}

// Register registers a Source under a given name.
func Register(name string, source Source) {
	store[name] = source
}

// Lookup returns a Source by the given name.
func Lookup(name string) Source {
	return store[name]
}

// LookupMultiple returns multiple Sources given multiple names.
func LookupMultiple(names []string) ([]Source, error) {
	sources := []Source{}

	for _, name := range names {
		source := Lookup(name)
		if source == nil {
			return nil, fmt.Errorf("%s source could not be identified", name)
		}
		sources = append(sources, source)
	}

	return sources, nil
}
