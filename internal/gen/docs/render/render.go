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
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// WriteToFile writes the given content to a file, creating or truncating it.
func WriteToFile(filename string, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// RenderTemplate parses and executes a named Go template from the given filesystem.
func RenderTemplate(fsys fs.FS, name string, data any) (string, error) {
	tmpl := template.New("").Funcs(FuncMap())
	template.Must(tmpl.ParseFS(fsys, "templates/*.gotpl"))

	var b bytes.Buffer
	if err := tmpl.ExecuteTemplate(&b, name, data); err != nil {
		return "", err
	}
	return b.String(), nil
}

// FuncMap returns a mapping of all of the functions that Engine has.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"backtick": func(times int) string {
			return strings.Repeat("`", times)
		},
		"capitalize": cases.Title(language.English, cases.Compact).String,
		"replace":    strings.ReplaceAll,
		"lower":      strings.ToLower,
		"bold": func(s string) string {
			return "**" + s + "**"
		},
		// padRight pads s with spaces on the right to the given width.
		"padRight": func(width int, s string) string {
			return fmt.Sprintf("%-*s", width, s)
		},
		// leftSep generates a left-aligned markdown table separator of the given column width.
		"leftSep": func(width int) string {
			return strings.Repeat("-", width+1)
		},
	}
}

// ComputeColumnWidth returns the maximum string length among the header and all values.
func ComputeColumnWidth(header string, values []string) int {
	return MapColumn(header, values, func(s string) string { return s })
}

// MapColumn returns the max width among the header and fn applied to each item.
func MapColumn[T any](header string, items []T, fn func(T) string) int {
	w := len(header)
	for _, item := range items {
		w = max(w, len(fn(item)))
	}
	return w
}
