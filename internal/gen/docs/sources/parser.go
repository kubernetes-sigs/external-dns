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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	// Regex to match source type names (must end with "Source")
	sourceTypeRegex = regexp.MustCompile(`^(\w+)Source$`)
)

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
