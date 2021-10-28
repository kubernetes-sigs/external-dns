// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package packagesdriver fetches type sizes for go/packages and go/analysis.
package packagesdriver

import (
<<<<<<< HEAD
<<<<<<< HEAD
	"context"
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/internal/gocommand"
)

var debug = false

func GetSizesGolist(ctx context.Context, inv gocommand.Invocation, gocmdRunner *gocommand.Runner) (types.Sizes, error) {
	inv.Verb = "list"
	inv.Args = []string{"-f", "{{context.GOARCH}} {{context.Compiler}}", "--", "unsafe"}
	stdout, stderr, friendlyErr, rawErr := gocmdRunner.RunRaw(ctx, inv)
	var goarch, compiler string
	if rawErr != nil {
		if rawErrMsg := rawErr.Error(); strings.Contains(rawErrMsg, "cannot find main module") || strings.Contains(rawErrMsg, "go.mod file not found") {
			// User's running outside of a module. All bets are off. Get GOARCH and guess compiler is gc.
			// TODO(matloob): Is this a problem in practice?
			inv.Verb = "env"
			inv.Args = []string{"GOARCH"}
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"bytes"
||||||| parent of 4d7e5ad26 (update vendored files)
	"bytes"
=======
>>>>>>> 4d7e5ad26 (update vendored files)
	"context"
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/internal/gocommand"
)

var debug = false

func GetSizesGolist(ctx context.Context, inv gocommand.Invocation, gocmdRunner *gocommand.Runner) (types.Sizes, error) {
	inv.Verb = "list"
	inv.Args = []string{"-f", "{{context.GOARCH}} {{context.Compiler}}", "--", "unsafe"}
	stdout, stderr, friendlyErr, rawErr := gocmdRunner.RunRaw(ctx, inv)
	var goarch, compiler string
	if rawErr != nil {
		if rawErrMsg := rawErr.Error(); strings.Contains(rawErrMsg, "cannot find main module") || strings.Contains(rawErrMsg, "go.mod file not found") {
			// User's running outside of a module. All bets are off. Get GOARCH and guess compiler is gc.
			// TODO(matloob): Is this a problem in practice?
<<<<<<< HEAD
			inv := gocommand.Invocation{
				Verb:       "env",
				Args:       []string{"GOARCH"},
				Env:        env,
				WorkingDir: dir,
			}
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
			inv := gocommand.Invocation{
				Verb:       "env",
				Args:       []string{"GOARCH"},
				Env:        env,
				WorkingDir: dir,
			}
=======
			inv.Verb = "env"
			inv.Args = []string{"GOARCH"}
>>>>>>> 4d7e5ad26 (update vendored files)
			envout, enverr := gocmdRunner.Run(ctx, inv)
			if enverr != nil {
				return nil, enverr
			}
			goarch = strings.TrimSpace(envout.String())
			compiler = "gc"
		} else {
			return nil, friendlyErr
		}
	} else {
		fields := strings.Fields(stdout.String())
		if len(fields) < 2 {
			return nil, fmt.Errorf("could not parse GOARCH and Go compiler in format \"<GOARCH> <compiler>\":\nstdout: <<%s>>\nstderr: <<%s>>",
				stdout.String(), stderr.String())
		}
		goarch = fields[0]
		compiler = fields[1]
	}
	return types.SizesFor(compiler, goarch), nil
}
