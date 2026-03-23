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

package render

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

func TestComputeColumnWidth(t *testing.T) {
	tests := []struct {
		name   string
		header string
		values []string
		want   int
	}{
		{
			name:   "header wins when all values are shorter",
			header: "Metric Type",
			values: []string{"gauge", "counter"},
			want:   len("Metric Type"),
		},
		{
			name:   "value wins when longer than header",
			header: "Name",
			values: []string{"last_reconcile_timestamp_seconds"},
			want:   len("last_reconcile_timestamp_seconds"),
		},
		{
			name:   "empty values returns header length",
			header: "Subsystem",
			values: []string{},
			want:   len("Subsystem"),
		},
		{
			name:   "empty header and empty values returns zero",
			header: "",
			values: []string{},
			want:   0,
		},
		{
			name:   "empty header defers to longest value",
			header: "",
			values: []string{"short", "much longer value"},
			want:   len("much longer value"),
		},
		{
			name:   "empty string values do not shrink below header",
			header: "Help",
			values: []string{"", ""},
			want:   len("Help"),
		},
		{
			name:   "tie between header and value returns that length",
			header: "exact",
			values: []string{"exact"},
			want:   len("exact"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ComputeColumnWidth(tt.header, tt.values)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapColumn(t *testing.T) {
	type item struct{ val string }
	fn := func(i item) string { return i.val }

	tests := []struct {
		name   string
		header string
		items  []item
		want   int
	}{
		{
			name:   "header wins when all values are shorter",
			header: "Metric Type",
			items:  []item{{"gauge"}, {"counter"}},
			want:   len("Metric Type"),
		},
		{
			name:   "value wins when longer than header",
			header: "Name",
			items:  []item{{"last_reconcile_timestamp_seconds"}},
			want:   len("last_reconcile_timestamp_seconds"),
		},
		{
			name:   "empty items returns header length",
			header: "Subsystem",
			items:  []item{},
			want:   len("Subsystem"),
		},
		{
			name:   "empty header and empty items returns zero",
			header: "",
			items:  []item{},
			want:   0,
		},
		{
			name:   "empty header defers to longest value",
			header: "",
			items:  []item{{"short"}, {"much longer value"}},
			want:   len("much longer value"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapColumn(tt.header, tt.items, fn)
			assert.Equal(t, tt.want, got)
		})
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
