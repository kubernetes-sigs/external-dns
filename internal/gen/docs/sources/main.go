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
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"sigs.k8s.io/external-dns/internal/gen/docs/utils"
)

const (
	annotationPrefix       = "+externaldns:source:"
	annotationName         = annotationPrefix + "name="
	annotationCategory     = annotationPrefix + "category="
	annotationDesc         = annotationPrefix + "description="
	annotationResources    = annotationPrefix + "resources="
	annotationFilters      = annotationPrefix + "filters="
	annotationNamespace    = annotationPrefix + "namespace="
	annotationFQDNTemplate = annotationPrefix + "fqdn-template="
)

var (
	//go:embed "templates/*"
	templates embed.FS
	// Regex to match source type names (must end with "Source")
	sourceTypeRegex = regexp.MustCompile(`^(\w+)Source$`)
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
	content, err := sources.generateMarkdown()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to generate markdown file '%s': %v\n", path, err)
		os.Exit(1)
	}
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

func (s *Sources) generateMarkdown() (string, error) {
	tmpl := template.New("").Funcs(utils.FuncMap())
	template.Must(tmpl.ParseFS(templates, "templates/*.gotpl"))

	var b bytes.Buffer
	err := tmpl.ExecuteTemplate(&b, "sources.gotpl", s)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}

// parseSourceAnnotations parses all Go files in the source directory
// and extracts source metadata from +externaldns:source annotations
func parseSourceAnnotations(sourceDir string) (Sources, error) {
	var sources Sources

	// Walk through the source directory
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and non-Go files
		if info.IsDir() || !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Skip test files
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		// Parse the Go file
		fileSources, err := parseFile(path, sourceDir)
		if err != nil {
			return fmt.Errorf("failed to parse %s: %w", path, err)
		}

		sources = append(sources, fileSources...)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return sources, nil
}

// parseFile parses a single Go file and extracts source annotations
func parseFile(filePath, baseDir string) (Sources, error) {
	var sources Sources

	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	// Get relative path for the File field
	relPath, err := filepath.Rel(baseDir, filePath)
	if err != nil {
		relPath = filePath
	}
	// Normalize to use forward slashes
	relPath = filepath.ToSlash(relPath)

	// Create a map of all comments by their starting position
	cmap := ast.NewCommentMap(fset, node, node.Comments)

	// Inspect the AST for type declarations
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for type declarations that are GenDecl (general declarations)
		genDecl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}

		// Get comments associated with this declaration
		comments := cmap[genDecl]
		if len(comments) == 0 {
			return true
		}

		// Check each spec in the declaration
		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// Check if it's a struct type
			_, ok = typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// Check if the type name matches *Source pattern
			typeName := typeSpec.Name.Name
			if !sourceTypeRegex.MatchString(typeName) {
				continue
			}

			// Combine all comment text
			var commentText strings.Builder
			for _, cg := range comments {
				commentText.WriteString(cg.Text())
			}

			if commentText.Len() == 0 {
				continue
			}

			extractedSources := extractSourcesFromComments(commentText.String(), typeName, relPath)
			sources = append(sources, extractedSources...)
		}

		return true
	})

	return sources, nil
}

// extractSourcesFromComments extracts source metadata from comment text.
// It can extract multiple sources from the same comment block (e.g., for gateway routes).
func extractSourcesFromComments(comments, typeName, filePath string) Sources {
	lines := strings.Split(comments, "\n")
	var sources Sources
	var currentSource *Source

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// When we see a name annotation, start a new source
		if strings.HasPrefix(line, annotationName) {
			// Save previous source if it exists
			if currentSource != nil && currentSource.Name != "" {
				sources = append(sources, *currentSource)
			}

			// Start new source
			currentSource = &Source{
				Type: typeName,
				File: filePath,
				Name: strings.TrimPrefix(line, annotationName),
			}
		} else if currentSource != nil {
			// Add other annotations to the current source
			if strings.HasPrefix(line, annotationCategory) {
				currentSource.Category = strings.TrimPrefix(line, annotationCategory)
			} else if strings.HasPrefix(line, annotationDesc) {
				currentSource.Description = strings.TrimPrefix(line, annotationDesc)
			} else if strings.HasPrefix(line, annotationResources) {
				currentSource.Resources = strings.TrimPrefix(line, annotationResources)
			} else if strings.HasPrefix(line, annotationFilters) {
				currentSource.Filters = strings.TrimPrefix(line, annotationFilters)
			} else if strings.HasPrefix(line, annotationNamespace) {
				currentSource.Namespace = strings.TrimPrefix(line, annotationNamespace)
			} else if strings.HasPrefix(line, annotationFQDNTemplate) {
				currentSource.FQDNTemplate = strings.TrimPrefix(line, annotationFQDNTemplate)
			}
		}
	}

	// Don't forget the last source
	if currentSource != nil && currentSource.Name != "" {
		sources = append(sources, *currentSource)
	}

	// Validate and set defaults for each source
	var validSources Sources
	for _, source := range sources {
		if source.Name == "" {
			fmt.Fprintf(os.Stderr, "Warning: source %s in %s is missing +externaldns:source:name annotation\n", typeName, filePath)
			continue
		}

		// Set defaults for optional fields
		if source.Category == "" {
			source.Category = "Uncategorized"
		}
		if source.Description == "" {
			source.Description = "No description available"
		}

		validSources = append(validSources, source)
	}

	return validSources
}
