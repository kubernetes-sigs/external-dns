package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
	fsys := os.DirFS(fmt.Sprintf("%s/../../docs", testPath))
	filePath := "flags.md"
	_, err := fs.Stat(fsys, filePath)
	assert.NoError(t, err, "expected file %s to exist", filePath)
}

func TestFlagsMdUpToDate(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf("%s/../../docs", testPath))
	filePath := "flags.md"
	expected, err := fs.ReadFile(fsys, filePath)
	assert.NoError(t, err, "expected file %s to exist", filePath)

	flags := computeFlags()
	actual, err := flags.generateMarkdownTable()
	assert.NoError(t, err)

	assert.True(t, len(expected) == len(actual), "expected file '%s' to be up to date. execute 'go run internal/gen/main.g", filePath)
}

func TestFlagsMdExtraFlagAdded(t *testing.T) {
	testPath, _ := os.Getwd()
	fsys := os.DirFS(fmt.Sprintf("%s/../../docs", testPath))
	filePath := "flags.md"
	expected, err := fs.ReadFile(fsys, filePath)
	assert.NoError(t, err, "expected file %s to exist", filePath)

	flags := computeFlags()
	flags.AddFlag("new-flag", "description2")
	actual, err := flags.generateMarkdownTable()

	assert.NoError(t, err)
	assert.NotEqual(t, string(expected), actual)
}
