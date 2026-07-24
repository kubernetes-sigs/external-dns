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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapSource(t *testing.T) {
	values := map[string]string{
		"API_TOKEN": "pipeline-a",
	}
	source := NewMapSource(values)

	values["API_TOKEN"] = "mutated"

	value, ok := source.LookupEnv("API_TOKEN")
	assert.True(t, ok)
	assert.Equal(t, "pipeline-a", value)
	assert.Equal(t, "pipeline-a", source.Getenv("API_TOKEN"))

	_, ok = source.LookupEnv("MISSING")
	assert.False(t, ok)
	assert.Empty(t, source.Getenv("MISSING"))
}

func TestContextSource(t *testing.T) {
	ctx := NewContext(t.Context(), NewMapSource(map[string]string{
		"API_TOKEN": "pipeline-a",
	}))

	value, ok := FromContext(ctx).LookupEnv("API_TOKEN")
	assert.True(t, ok)
	assert.Equal(t, "pipeline-a", value)
}

func TestContextDefaultsToSystemSource(t *testing.T) {
	t.Setenv("API_TOKEN", "system")

	value, ok := FromContext(t.Context()).LookupEnv("API_TOKEN")
	assert.True(t, ok)
	assert.Equal(t, "system", value)
}

func TestNilContextDefaultsToSystemSource(t *testing.T) {
	t.Setenv("API_TOKEN", "system")

	value, ok := FromContext(nil).LookupEnv("API_TOKEN")
	assert.True(t, ok)
	assert.Equal(t, "system", value)

	ctx := NewContext(nil, NewMapSource(map[string]string{
		"API_TOKEN": "pipeline-a",
	}))
	value, ok = FromContext(ctx).LookupEnv("API_TOKEN")
	assert.True(t, ok)
	assert.Equal(t, "pipeline-a", value)
}
