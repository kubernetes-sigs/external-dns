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
	"os"
	"testing"
)

func TestDiscoverSources(t *testing.T) {
	// Change to project root for testing
	// Skip if we can't find the source directory
	_, err := os.Stat("../../../../source")
	if err != nil {
		t.Skip("Skipping test: source directory not found (test should run from project root)")
	}

	err = os.Chdir("../../../..")
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}

	sources := discoverSources()

	// We expect at least the core Kubernetes sources that have annotations
	if len(sources) < 5 {
		t.Errorf("Expected at least 5 sources with annotations, got %d", len(sources))
	}
}

func TestGenerateMarkdownTable(t *testing.T) {
	sources := Sources{
		{Name: "test", Type: "testSource", File: "source/test.go", Category: "Test"},
	}

	content, err := sources.generateMarkdownTable()
	if err != nil {
		t.Errorf("Failed to generate markdown table: %v", err)
	}

	if content == "" {
		t.Error("Expected non-empty content, but got empty string")
	}
}
