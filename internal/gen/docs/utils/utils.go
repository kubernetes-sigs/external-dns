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
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func WriteToFile(filename string, content string) error {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		_ = fmt.Errorf("failed to create file: %w", fileErr)
	}
	defer file.Close()
	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		_ = fmt.Errorf("failed to write to file: %s", filename)
	}
	return nil
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
	w := len(header)
	for _, v := range values {
		w = max(w, len(v))
	}
	return w
}
