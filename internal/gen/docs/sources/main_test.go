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
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	pathToDocs = "%s/../../../../docs/sources"
	fileName   = "index.md"
)

func TestIndexMdExists(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf(pathToDocs, testPath))
	st, err := fs.Stat(fsys, fileName)
	assert.NoError(t, err, "expected file %s to exist", fileName)
	assert.Equal(t, fileName, st.Name())
}

func TestIndexMdUpToDate(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf(pathToDocs, testPath))
	expected, err := fs.ReadFile(fsys, fileName)
	assert.NoError(t, err, "expected file %s to exist", fileName)

	// path to sources folder
	ssys := os.DirFS(fmt.Sprintf("%s/../../../../source", testPath))
	sources, err := discoverSources(fmt.Sprintf("%s", ssys))
	require.NoError(t, err, "expected to find sources")
	actual, err := sources.generateMarkdown()
	assert.NoError(t, err)
	assert.Contains(t, string(expected), actual, "expected file 'docs/source/index.md' to be up to date. execute 'make generate-sources-documentation'")
}

func TestDiscoverSources(t *testing.T) {
	testPath, _ := os.Getwd()
	ssys := os.DirFS(fmt.Sprintf("%s/../../../../source", testPath))

	sources, err := discoverSources(fmt.Sprintf("%s", ssys))
	require.NoError(t, err)

	assert.GreaterOrEqual(t, len(sources), 5, "Expected at least 5 sources with annotations")

	// Verify sources are sorted by category, then by name
	for i := range len(sources) - 1 {
		prev, curr := sources[i], sources[i+1]
		if prev.Name > curr.Name {
			t.Errorf("Sources not sorted correctly: %s should come before %s", curr.Name, prev.Name)
		}
	}
}

func TestGenerateMarkdown(t *testing.T) {
	sources := Sources{
		{
			Name:         "test",
			Type:         "testSource",
			File:         "source/test.go",
			Category:     "Test",
			Description:  "Test source",
			Resources:    "TestResource",
			Filters:      "annotation,label",
			Namespace:    "all,single",
			FQDNTemplate: "true",
		},
	}

	content, err := sources.generateMarkdown()
	require.NoError(t, err)
	assert.NotEmpty(t, content)

	assert.Contains(t, content, "# Supported Sources")
	assert.Contains(t, content, "## Available Sources")
}

func TestParseSourceAnnotations(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_source.go")
	content := `package main

// testSource is a test source implementation.
//
// +externaldns:source:name=test-source
// +externaldns:source:category=Testing
// +externaldns:source:description=A test source for unit testing
// +externaldns:source:resources=TestResource
// +externaldns:source:filters=annotation,label
// +externaldns:source:namespace=all,single
// +externaldns:source:fqdn-template=true
// +externaldns:source:provider-specific=true
type testSource struct {
	client string
}
`
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	sources, err := parseSourceAnnotations(tmpDir)
	require.NoError(t, err)
	assert.Len(t, sources, 1)

	source := sources[0]
	assert.Equal(t, "test-source", source.Name)
	assert.Equal(t, "Testing", source.Category)
	assert.Equal(t, "TestResource", source.Resources)
	assert.Equal(t, "annotation,label", source.Filters)
	assert.Equal(t, "all,single", source.Namespace)
	assert.Equal(t, "true", source.FQDNTemplate)
	assert.Equal(t, "false", source.Events)
	assert.Equal(t, "true", source.ProviderSpecific)
}

func TestParseSourceAnnotations_SkipsTestFiles(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a test file that should be skipped
	testFile := filepath.Join(tmpDir, "test_source_test.go")
	content := `package main

// +externaldns:source:name=should-be-skipped
// +externaldns:source:category=Test
// +externaldns:source:description=Should be skipped
type testSource struct {}
`
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	sources, err := parseSourceAnnotations(tmpDir)
	require.NoError(t, err)
	assert.Empty(t, sources)
}

func TestParseFile_MultipleSourcesInOneFile(t *testing.T) {
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "multi.go")
	content := `package main

// firstSource is the first source.
//
// +externaldns:source:name=first
// +externaldns:source:category=Testing
// +externaldns:source:description=First source
type firstSource struct {}

// secondSource is the second source.
//
// +externaldns:source:name=second
// +externaldns:source:category=Testing
// +externaldns:source:description=Second source
// +externaldns:source:events=true
type secondSource struct {}
`
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	sources, err := parseFile(testFile, tmpDir)
	require.NoError(t, err)
	assert.Len(t, sources, 2)
	assert.Equal(t, "first", sources[0].Name)
	assert.Equal(t, "false", sources[0].Events)
	assert.Equal(t, "false", sources[0].ProviderSpecific)
	assert.Equal(t, "second", sources[1].Name)
	assert.Equal(t, "true", sources[1].Events)
	assert.Equal(t, "false", sources[1].ProviderSpecific)
}

func TestParseFile_IgnoresNonSourceTypes(t *testing.T) {
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "nonsource.go")
	content := `package main

// regularStruct is not a source (doesn't end with "Source").
//
// +externaldns:source:name=should-not-parse
// +externaldns:source:category=Test
// +externaldns:source:description=Should not be parsed
type regularStruct struct {}
`
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	sources, err := parseFile(testFile, tmpDir)
	require.NoError(t, err)
	assert.Empty(t, sources)
}

func TestParseSourceAnnotations_ErrorOnInvalidFile(t *testing.T) {
	tmpDir := t.TempDir()

	// Create a file with invalid Go syntax
	testFile := filepath.Join(tmpDir, "invalid.go")
	content := `package main

this is not valid go syntax
`
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	_, err = parseSourceAnnotations(tmpDir)
	require.Error(t, err)
}

func TestParseFile_InvalidGoFile(t *testing.T) {
	tmpDir := t.TempDir()

	testFile := filepath.Join(tmpDir, "invalid.go")
	content := `this is not valid go code`

	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	_, err = parseFile(testFile, tmpDir)
	require.Error(t, err)
}

func TestParseSourceAnnotations_WithSubdirectories(t *testing.T) {
	tmpDir := t.TempDir()
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}

	// Create a test file in subdirectory
	testFile := filepath.Join(subDir, "nested_source.go")
	content := `package main

// nestedSource is in a subdirectory.
//
// +externaldns:source:name=nested
// +externaldns:source:category=Testing
// +externaldns:source:description=Nested source
type nestedSource struct {}
`
	err := os.WriteFile(testFile, []byte(content), 0644)
	require.NoError(t, err)

	sources, err := parseSourceAnnotations(tmpDir)
	require.NoError(t, err)
	assert.Len(t, sources, 1)
	assert.Equal(t, "nested", sources[0].Name)
	assert.Contains(t, sources[0].File, "subdir/nested_source.go")
}

func TestGenerateMarkdown_WithMultipleCategories(t *testing.T) {
	sources := Sources{
		{
			Name:         "service",
			Category:     "Kubernetes Core",
			Description:  "Service source",
			Resources:    "Service",
			Filters:      "annotation,label",
			Namespace:    "all,single",
			FQDNTemplate: "true",
		},
		{
			Name:         "ingress",
			Category:     "Kubernetes Core",
			Description:  "Ingress source",
			Resources:    "Ingress",
			Filters:      "annotation,label",
			Namespace:    "all,single",
			FQDNTemplate: "true",
		},
		{
			Name:         "gateway-httproute",
			Category:     "Gateway API",
			Description:  "HTTP route source",
			Resources:    "HTTPRoute.gateway.networking.k8s.io",
			Filters:      "annotation,label",
			Namespace:    "all,single",
			FQDNTemplate: "false",
		},
	}

	content, err := sources.generateMarkdown()
	require.NoError(t, err)
	assert.Contains(t, content, "service")
	assert.Contains(t, content, "ingress")
	assert.Contains(t, content, "gateway-httproute")
}

func TestExtractSourcesFromComments(t *testing.T) {
	tests := []struct {
		name        string
		comments    string
		typeName    string
		filePath    string
		wantSources int
		wantErr     bool
		validate    func(*testing.T, Source)
	}{
		{
			name: "valid single source",
			comments: `testSource is a test implementation.

+externaldns:source:name=test
+externaldns:source:category=Testing
+externaldns:source:description=A test source
+externaldns:source:resources=TestResource
+externaldns:source:filters=annotation
+externaldns:source:namespace=all
+externaldns:source:fqdn-template=false
`,
			typeName:    "testSource",
			filePath:    "test.go",
			wantSources: 1,
			validate: func(t *testing.T, s Source) {
				assert.Equal(t, "test", s.Name)
				assert.Equal(t, "Testing", s.Category)
				assert.Equal(t, "A test source", s.Description)
			},
		},
		{
			name: "multiple sources in same comment block",
			comments: `gatewaySource handles multiple gateway types.

+externaldns:source:name=http-route
+externaldns:source:category=Gateway
+externaldns:source:description=Handles HTTP routes
+externaldns:source:resources=HTTPRoute

+externaldns:source:name=tcp-route
+externaldns:source:category=Gateway
+externaldns:source:description=Handles TCP routes
+externaldns:source:resources=TCPRoute
`,
			typeName:    "gatewaySource",
			filePath:    "gateway.go",
			wantSources: 2,
			validate: func(t *testing.T, s Source) {
				assert.Contains(t, []string{"http-route", "tcp-route"}, s.Name)
			},
		},
		{
			name: "missing required name annotation",
			comments: `testSource without name.

+externaldns:source:category=Testing
+externaldns:source:description=Missing name
`,
			typeName: "testSource",
			filePath: "test.go",
			wantErr:  true,
		},
		{
			name: "optional annotations can be missing",
			comments: `testSource with minimal annotations.

+externaldns:source:name=minimal
+externaldns:source:category=Testing
+externaldns:source:description=Minimal source
`,
			typeName:    "testSource",
			filePath:    "test.go",
			wantSources: 1,
			validate: func(t *testing.T, s Source) {
				assert.Equal(t, "minimal", s.Name)
				assert.Empty(t, s.Resources)
				assert.Empty(t, s.Filters)
			},
		},
		{
			name: "missing name annotation",
			comments: `testSource with minimal annotations.

+externaldns:source:name=
+externaldns:source:category=Testing
+externaldns:source:description=Minimal source
`,
			typeName: "testSource",
			filePath: "test.go",
			validate: func(t *testing.T, s Source) {
				require.Nil(t, s)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sources, err := extractSourcesFromComments(tt.comments, tt.typeName, tt.filePath)

			if tt.wantErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			assert.Len(t, sources, tt.wantSources)

			if tt.validate != nil {
				for _, source := range sources {
					tt.validate(t, source)
				}
			}

			// Verify all sources have required fields
			for _, source := range sources {
				assert.Equal(t, source.Type, tt.typeName)
				assert.Equal(t, source.File, tt.filePath)
			}
		})
	}
}
