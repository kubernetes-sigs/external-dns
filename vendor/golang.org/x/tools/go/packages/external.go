// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package packages

// This file defines the protocol that enables an external "driver"
// tool to supply package metadata in place of 'go list'.

import (
	"bytes"
	"encoding/json"
	"fmt"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	exec "golang.org/x/sys/execabs"
	"os"
	"strings"
)

// The Driver Protocol
//
// The driver, given the inputs to a call to Load, returns metadata about the packages specified.
// This allows for different build systems to support go/packages by telling go/packages how the
// packages' source is organized.
// The driver is a binary, either specified by the GOPACKAGESDRIVER environment variable or in
// the path as gopackagesdriver. It's given the inputs to load in its argv. See the package
// documentation in doc.go for the full description of the patterns that need to be supported.
// A driver receives as a JSON-serialized driverRequest struct in standard input and will
// produce a JSON-serialized driverResponse (see definition in packages.go) in its standard output.

// driverRequest is used to provide the portion of Load's Config that is needed by a driver.
type driverRequest struct {
	Mode LoadMode `json:"mode"`
	// Env specifies the environment the underlying build system should be run in.
	Env []string `json:"env"`
	// BuildFlags are flags that should be passed to the underlying build system.
	BuildFlags []string `json:"build_flags"`
	// Tests specifies whether the patterns should also return test packages.
	Tests bool `json:"tests"`
	// Overlay maps file paths (relative to the driver's working directory) to the byte contents
	// of overlay files.
	Overlay map[string][]byte `json:"overlay"`
}

// findExternalDriver returns the file path of a tool that supplies
// the build system package structure, or "" if not found."
// If GOPACKAGESDRIVER is set in the environment findExternalTool returns its
// value, otherwise it searches for a binary named gopackagesdriver on the PATH.
func findExternalDriver(cfg *Config) driver {
	const toolPrefix = "GOPACKAGESDRIVER="
	tool := ""
	for _, env := range cfg.Env {
		if val := strings.TrimPrefix(env, toolPrefix); val != env {
			tool = val
		}
	}
	if tool != "" && tool == "off" {
		return nil
	}
	if tool == "" {
		var err error
		tool, err = exec.LookPath("gopackagesdriver")
		if err != nil {
			return nil
		}
	}
	return func(cfg *Config, words ...string) (*driverResponse, error) {
		req, err := json.Marshal(driverRequest{
			Mode:       cfg.Mode,
			Env:        cfg.Env,
			BuildFlags: cfg.BuildFlags,
			Tests:      cfg.Tests,
			Overlay:    cfg.Overlay,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to encode message to driver tool: %v", err)
		}

		buf := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		cmd := exec.CommandContext(cfg.Context, tool, words...)
		cmd.Dir = cfg.Dir
		cmd.Env = cfg.Env
		cmd.Stdin = bytes.NewReader(req)
		cmd.Stdout = buf
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("%v: %v: %s", tool, err, cmd.Stderr)
		}
		if len(stderr.Bytes()) != 0 && os.Getenv("GOPACKAGESPRINTDRIVERERRORS") != "" {
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd), stderr)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	exec "golang.org/x/sys/execabs"
>>>>>>> 4d7e5ad26 (update vendored files)
	"os"
	"strings"
)

// The Driver Protocol
//
// The driver, given the inputs to a call to Load, returns metadata about the packages specified.
// This allows for different build systems to support go/packages by telling go/packages how the
// packages' source is organized.
// The driver is a binary, either specified by the GOPACKAGESDRIVER environment variable or in
// the path as gopackagesdriver. It's given the inputs to load in its argv. See the package
// documentation in doc.go for the full description of the patterns that need to be supported.
// A driver receives as a JSON-serialized driverRequest struct in standard input and will
// produce a JSON-serialized driverResponse (see definition in packages.go) in its standard output.

// driverRequest is used to provide the portion of Load's Config that is needed by a driver.
type driverRequest struct {
	Mode LoadMode `json:"mode"`
	// Env specifies the environment the underlying build system should be run in.
	Env []string `json:"env"`
	// BuildFlags are flags that should be passed to the underlying build system.
	BuildFlags []string `json:"build_flags"`
	// Tests specifies whether the patterns should also return test packages.
	Tests bool `json:"tests"`
	// Overlay maps file paths (relative to the driver's working directory) to the byte contents
	// of overlay files.
	Overlay map[string][]byte `json:"overlay"`
}

// findExternalDriver returns the file path of a tool that supplies
// the build system package structure, or "" if not found."
// If GOPACKAGESDRIVER is set in the environment findExternalTool returns its
// value, otherwise it searches for a binary named gopackagesdriver on the PATH.
func findExternalDriver(cfg *Config) driver {
	const toolPrefix = "GOPACKAGESDRIVER="
	tool := ""
	for _, env := range cfg.Env {
		if val := strings.TrimPrefix(env, toolPrefix); val != env {
			tool = val
		}
	}
	if tool != "" && tool == "off" {
		return nil
	}
	if tool == "" {
		var err error
		tool, err = exec.LookPath("gopackagesdriver")
		if err != nil {
			return nil
		}
	}
	return func(cfg *Config, words ...string) (*driverResponse, error) {
		req, err := json.Marshal(driverRequest{
			Mode:       cfg.Mode,
			Env:        cfg.Env,
			BuildFlags: cfg.BuildFlags,
			Tests:      cfg.Tests,
			Overlay:    cfg.Overlay,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to encode message to driver tool: %v", err)
		}

		buf := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		cmd := exec.CommandContext(cfg.Context, tool, words...)
		cmd.Dir = cfg.Dir
		cmd.Env = cfg.Env
		cmd.Stdin = bytes.NewReader(req)
		cmd.Stdout = buf
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("%v: %v: %s", tool, err, cmd.Stderr)
		}
		if len(stderr.Bytes()) != 0 && os.Getenv("GOPACKAGESPRINTDRIVERERRORS") != "" {
<<<<<<< HEAD
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd, words...), stderr)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd, words...), stderr)
=======
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd), stderr)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"os"
	"os/exec"
	"strings"
)

// DriverRequest defines the schema of a request for package metadata
// from an external driver program. The JSON-encoded DriverRequest
// message is provided to the driver program's standard input. The
// query patterns are provided as command-line arguments.
//
// See the package documentation for an overview.
type DriverRequest struct {
	Mode LoadMode `json:"mode"`

	// Env specifies the environment the underlying build system should be run in.
	Env []string `json:"env"`

	// BuildFlags are flags that should be passed to the underlying build system.
	BuildFlags []string `json:"build_flags"`

	// Tests specifies whether the patterns should also return test packages.
	Tests bool `json:"tests"`

	// Overlay maps file paths (relative to the driver's working directory) to the byte contents
	// of overlay files.
	Overlay map[string][]byte `json:"overlay"`
}

// DriverResponse defines the schema of a response from an external
// driver program, providing the results of a query for package
// metadata. The driver program must write a JSON-encoded
// DriverResponse message to its standard output.
//
// See the package documentation for an overview.
type DriverResponse struct {
	// NotHandled is returned if the request can't be handled by the current
	// driver. If an external driver returns a response with NotHandled, the
	// rest of the DriverResponse is ignored, and go/packages will fallback
	// to the next driver. If go/packages is extended in the future to support
	// lists of multiple drivers, go/packages will fall back to the next driver.
	NotHandled bool

	// Compiler and Arch are the arguments pass of types.SizesFor
	// to get a types.Sizes to use when type checking.
	Compiler string
	Arch     string

	// Roots is the set of package IDs that make up the root packages.
	// We have to encode this separately because when we encode a single package
	// we cannot know if it is one of the roots as that requires knowledge of the
	// graph it is part of.
	Roots []string `json:",omitempty"`

	// Packages is the full set of packages in the graph.
	// The packages are not connected into a graph.
	// The Imports if populated will be stubs that only have their ID set.
	// Imports will be connected and then type and syntax information added in a
	// later pass (see refine).
	Packages []*Package

	// GoVersion is the minor version number used by the driver
	// (e.g. the go command on the PATH) when selecting .go files.
	// Zero means unknown.
	GoVersion int
}

// driver is the type for functions that query the build system for the
// packages named by the patterns.
type driver func(cfg *Config, patterns ...string) (*DriverResponse, error)

// findExternalDriver returns the file path of a tool that supplies
// the build system package structure, or "" if not found."
// If GOPACKAGESDRIVER is set in the environment findExternalTool returns its
// value, otherwise it searches for a binary named gopackagesdriver on the PATH.
func findExternalDriver(cfg *Config) driver {
	const toolPrefix = "GOPACKAGESDRIVER="
	tool := ""
	for _, env := range cfg.Env {
		if val := strings.TrimPrefix(env, toolPrefix); val != env {
			tool = val
		}
	}
	if tool != "" && tool == "off" {
		return nil
	}
	if tool == "" {
		var err error
		tool, err = exec.LookPath("gopackagesdriver")
		if err != nil {
			return nil
		}
	}
	return func(cfg *Config, words ...string) (*DriverResponse, error) {
		req, err := json.Marshal(DriverRequest{
			Mode:       cfg.Mode,
			Env:        cfg.Env,
			BuildFlags: cfg.BuildFlags,
			Tests:      cfg.Tests,
			Overlay:    cfg.Overlay,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to encode message to driver tool: %v", err)
		}

		buf := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		cmd := exec.CommandContext(cfg.Context, tool, words...)
		cmd.Dir = cfg.Dir
		cmd.Env = cfg.Env
		cmd.Stdin = bytes.NewReader(req)
		cmd.Stdout = buf
		cmd.Stderr = stderr

		if err := cmd.Run(); err != nil {
			return nil, fmt.Errorf("%v: %v: %s", tool, err, cmd.Stderr)
		}
		if len(stderr.Bytes()) != 0 && os.Getenv("GOPACKAGESPRINTDRIVERERRORS") != "" {
<<<<<<< HEAD
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd, words...), stderr)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd, words...), stderr)
=======
			fmt.Fprintf(os.Stderr, "%s stderr: <<%s>>\n", cmdDebugStr(cmd), stderr)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
		}

		var response DriverResponse
		if err := json.Unmarshal(buf.Bytes(), &response); err != nil {
			return nil, err
		}
		return &response, nil
	}
}
