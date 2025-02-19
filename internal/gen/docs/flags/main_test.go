/*
Copyright 2017 The Kubernetes Authors.

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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const pathToDocs = "%s/../../../../docs"

func TestComputeFlags(t *testing.T) {
	flags := computeFlags()

	if len(flags) == 0 {
		t.Errorf("Expected non-zero flags, got %d", len(flags))
	}

	for _, flag := range flags {
		if strings.Contains(flag.Name, "help") || strings.Contains(flag.Name, "completion-") {
			t.Errorf("Unexpected flag: %s", flag.Name)
		}
	}
}

func TestGenerateMarkdownTableRenderer(t *testing.T) {
	flags := Flags{
		{Name: "flag1", Description: "description1"},
	}

	got, err := flags.generateMarkdownTable()
	assert.NoError(t, err)

	assert.Contains(t, got, "<!-- THIS FILE MUST NOT BE EDITED BY HAND -->")
	assert.Contains(t, got, "| flag1 | description1 |")
}

func TestFlagsMdExists(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf(pathToDocs, testPath))
	fileName := "flags.md"
	st, err := fs.Stat(fsys, fileName)
	assert.NoError(t, err, "expected file %s to exist", fileName)
	assert.Equal(t, fileName, st.Name())
}

func TestFlagsMdUpToDate(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf(pathToDocs, testPath))
	fileName := "flags.md"
	expected, err := fs.ReadFile(fsys, fileName)
	assert.NoError(t, err, "expected file %s to exist", fileName)

	flags := computeFlags()
	actual, err := flags.generateMarkdownTable()
	assert.NoError(t, err)
	actual = actual + "\n"
	assert.True(t, len(expected) == len(actual), "expected file '%s' to be up to date. execute 'make generate-flags-documentation", fileName)
}

func TestFlagsMdExtraFlagAdded(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf(pathToDocs, testPath))
	filePath := "flags.md"
	expected, err := fs.ReadFile(fsys, filePath)
	assert.NoError(t, err, "expected file %s to exist", filePath)

	flags := computeFlags()
	flags.addFlag("new-flag", "description2")
	actual, err := flags.generateMarkdownTable()

	assert.NoError(t, err)
	assert.NotEqual(t, string(expected), actual)
}
