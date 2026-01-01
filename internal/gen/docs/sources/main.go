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

package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"sort"
	"text/template"

	"sigs.k8s.io/external-dns/internal/gen/docs/utils"
)

var (
	//go:embed "templates/*"
	templates embed.FS
)

// Source represents metadata about a source implementation
type Source struct {
	Name         string // e.g., "service", "ingress", "crd"
	Type         string // e.g., "serviceSource"
	File         string // e.g., "source/service.go"
	Description  string // Description of what this source does
	Category     string // e.g., "Kubernetes", "Gateway", "Service Mesh", "Wrapper"
	Resources    string // Kubernetes resources watched, e.g., "Service", "Ingress"
	Filters      string // Supported filters, e.g., "annotation,label"
	Namespace    string // Namespace support: "all", "single", "multiple"
	FQDNTemplate string // FQDN template support: "true", "false"
}

type Sources []Source

// It generates a markdown file
// with the supported sources and writes it to the 'docs/sources/index.md' file.
// to re-generate `docs/sources/index.md` execute 'go run internal/gen/docs/sources/main.go'
func main() {
	testPath, _ := os.Getwd()
	path := fmt.Sprintf("%s/docs/sources/index.md", testPath)
	fmt.Printf("generate file '%s' with supported sources\n", path)

	sources := discoverSources()
	content, err := sources.generateMarkdownTable()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to generate markdown file '%s': %v\n", path, err)
		os.Exit(1)
	}
	content += "\n"
	_ = utils.WriteToFile(path, content)
}

// discoverSources scans the source directory and discovers all source implementations
// by parsing Go files and extracting +externaldns:source annotations
func discoverSources() Sources {
	// Get the source directory path
	testPath, _ := os.Getwd()
	sourceDir := fmt.Sprintf("%s/source", testPath)

	// Parse all source files for annotations
	sources, err := parseSourceAnnotations(sourceDir)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error parsing source annotations: %v\n", err)
		os.Exit(1)
	}

	if len(sources) == 0 {
		_, _ = fmt.Fprintf(os.Stderr, "Warning: No sources with annotations found\n")
	}

	// Sort sources by category, then by name
	sort.Slice(sources, func(i, j int) bool {
		if sources[i].Category == sources[j].Category {
			return sources[i].Name < sources[j].Name
		}
		return sources[i].Category < sources[j].Category
	})

	return sources
}

func (s *Sources) generateMarkdownTable() (string, error) {
	tmpl := template.New("").Funcs(utils.FuncMap())
	template.Must(tmpl.ParseFS(templates, "templates/*.gotpl"))

	var b bytes.Buffer
	err := tmpl.ExecuteTemplate(&b, "sources.gotpl", s)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
