/*
Copyright 2026 The Kubernetes Authors.

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

package credentials

import (
	"context"
	"maps"
	"os"
)

// Source provides credential values to provider constructors.
type Source interface {
	LookupEnv(key string) (string, bool)
	Getenv(key string) string
}

type sourceKey struct{}

type systemSource struct{}

// SystemSource returns a Source backed by the process environment.
func SystemSource() Source {
	return systemSource{}
}

func (systemSource) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (systemSource) Getenv(key string) string {
	return os.Getenv(key)
}

type mapSource map[string]string

// NewMapSource returns a Source backed by the provided values.
func NewMapSource(values map[string]string) Source {
	copied := make(mapSource, len(values))
	maps.Copy(copied, values)
	return copied
}

func (s mapSource) LookupEnv(key string) (string, bool) {
	value, ok := s[key]
	return value, ok
}

func (s mapSource) Getenv(key string) string {
	return s[key]
}

// NewContext returns a context that carries a provider credential source.
func NewContext(ctx context.Context, source Source) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if source == nil {
		source = SystemSource()
	}
	return context.WithValue(ctx, sourceKey{}, source)
}

// FromContext returns the credential source carried by ctx, or the process
// environment source when no explicit source was attached.
func FromContext(ctx context.Context) Source {
	if ctx == nil {
		return SystemSource()
	}
	source, ok := ctx.Value(sourceKey{}).(Source)
	if !ok || source == nil {
		return SystemSource()
	}
	return source
}
