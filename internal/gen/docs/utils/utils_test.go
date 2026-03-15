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

package utils

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"testing/fstest"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteToFile(t *testing.T) {
	filename := fmt.Sprintf("%s/testfile", t.TempDir())
	content := "Hello, World!"

	defer os.Remove(filename)

	err := WriteToFile(filename, content)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("expected no error reading file, got %v", err)
	}

	if string(data) != content {
		t.Errorf("expected content %q, got %q", content, string(data))
	}
}

func TestFuncs(t *testing.T) {
	tests := []struct {
		tpl, expect string
		vars        any
	}{
		{
			tpl:    `{{ backtick 3 }}`,
			expect: "```",
			vars:   map[string]any{},
		},
		{
			tpl:    `{{ capitalize .name }}`,
			expect: "Capital",
			vars:   map[string]any{"name": "capital"},
		},
		{
			tpl:    `{{ replace .resources "," "<br/>" }}`,
			expect: "one<br/>two<br/>tree",
			vars:   map[string]any{"resources": "one,two,tree"},
		},
	}

	for _, tt := range tests {
		var b strings.Builder
		err := template.Must(template.New("test").Funcs(FuncMap()).Parse(tt.tpl)).Execute(&b, tt.vars)
		assert.NoError(t, err)
		assert.Equal(t, tt.expect, b.String(), tt.tpl)
	}
}

func TestRenderTemplate(t *testing.T) {
	fsys := fstest.MapFS{
		"templates/test.gotpl": &fstest.MapFile{
			Data: []byte("Hello {{ .Name }}!"),
		},
	}

	result, err := RenderTemplate(fsys, "test.gotpl", struct{ Name string }{Name: "World"})
	require.NoError(t, err)
	assert.Equal(t, "Hello World!", result)
}

func TestRenderTemplateWithFuncMap(t *testing.T) {
	fsys := fstest.MapFS{
		"templates/test.gotpl": &fstest.MapFile{
			Data: []byte("{{ backtick 3 }}go\nfmt.Println({{ capitalize .Lang }})\n{{ backtick 3 }}"),
		},
	}

	result, err := RenderTemplate(fsys, "test.gotpl", struct{ Lang string }{Lang: "go"})
	require.NoError(t, err)
	assert.Contains(t, result, "```go")
	assert.Contains(t, result, "Go")
}
