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

import "errors"

func init() {
	// Register all selectable sources under an identifier.
	RegisterFunc("service", NewServiceSource)
	RegisterFunc("ingress", NewIngressSource)
	RegisterFunc("fake", NewFakeSource)
}

// sourceFunc is a constructor function that returns a Source and an error.
type sourceFunc func(cfg *Config) (Source, error)

// store is a global store for known sources.
var store = map[string]sourceFunc{}

// ErrSourceNotFound is returned when a requested source doesn't exist.
var ErrSourceNotFound = errors.New("source not found")

// Register registers a Source under a given name.
func Register(name string, source Source) {
	store[name] = func(_ *Config) (Source, error) { return source, nil }
}

// RegisterFunc registers a Source under a given name via a constructor function.
func RegisterFunc(name string, source sourceFunc) {
	store[name] = source
}

// Clear de-registers all sources from the global store.
func Clear() {
	store = map[string]sourceFunc{}
}

// Lookup returns a Source by the given name.
func Lookup(name string, cfg *Config) (Source, error) {
	sf, ok := store[name]
	if !ok {
		return nil, ErrSourceNotFound
	}

	source, err := sf(cfg)
	if err != nil {
		return nil, err
	}

	return source, nil
}

// LookupMultiple returns multiple Sources given multiple names.
func LookupMultiple(names []string, cfg *Config) ([]Source, error) {
	sources := []Source{}

	for _, name := range names {
		source, err := Lookup(name, cfg)
		if err != nil {
			return nil, ErrSourceNotFound
		}
		sources = append(sources, source)
	}

	return sources, nil
}
