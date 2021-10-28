// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/maxatome/go-testdeep/helpers/tdutil"
	"github.com/maxatome/go-testdeep/internal/color"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/flat"
	"github.com/maxatome/go-testdeep/internal/trace"
)

func init() {
	trace.Init()
	trace.IgnorePackage()
}

// stripTrace removes go-testdeep useless calls in a trace returned by
// trace.Retrieve() to make it clearer for the reader.
func stripTrace(s trace.Stack) trace.Stack {
	if len(s) <= 1 {
		return s
	}

	const (
		tdPkg      = "github.com/maxatome/go-testdeep/td"
		tdhttpPkg  = "github.com/maxatome/go-testdeep/helpers/tdhttp"
		tdsuitePkg = "github.com/maxatome/go-testdeep/helpers/tdsuite"
	)

	// Remove useless possible (*T).Run() or (*T).RunAssertRequire() first call
	if s.Match(-1, tdPkg, "(*T).Run.func1", "(*T).RunAssertRequire.func1") {
		// Remove useless tdhttp (*TestAPI).Run() call
		//
		// ✓ xxx      Subtest.func1()
		// ✗ …/tdhttp (*TestAPI).Run.func1
		// ✗ …/td     (*T).Run.func1()
		if s.Match(-2, tdhttpPkg, "(*TestAPI).Run.func1") {
			return s[:len(s)-2]
		}

		// Remove useless tdsuite calls
		//
		// ✓ xxx       Suite.TestSuite
		// ✗ reflect   Value.call
		// ✗ reflect   Value.Call
		// ✗ …/tdsuite run.func2
		// ✗ …/td      (*T).Run.func1() or (*T).RunAssertRequire.func1()
		//
		// or for PostTest
		// ✓ xxx       Suite.PostTest
		// ✗ …/tdsuite run.func2.1
		// ✗ …/tdsuite run.func2
		// ✗ …/td      (*T).Run.func1() or (*T).RunAssertRequire.func1()
		if s.Match(-2, tdsuitePkg, "run.func*") {
			// PostTest
			if s.Match(-3, tdsuitePkg, "run.func*") &&
				len(s) > 4 &&
				strings.HasSuffix(s[len(s)-4].Func, ".PostTest") {
				return s[:len(s)-3]
			}

			for i := len(s) - 3; i >= 1; i-- {
				if !s.Match(i, "reflect") {
					return s[:i+1]
				}
			}
			return nil
		}

		return s[:len(s)-1]
	}

	// Remove testing.Cleanup() stack
	//
	// ✓ xxx     TestCleanup.func2
	// ✗ testing (*common).Cleanup.func1
	// ✗ testing (*common).runCleanup
	// ✗ testing tRunner.func2
	if s.Match(-1, "testing", "tRunner.func*") &&
		s.Match(-2, "testing", "(*common).runCleanup") &&
		s.Match(-3, "testing", "(*common).Cleanup.func1") {
		return s[:len(s)-3]
	}

	// Remove tdsuite pre-Setup/BetweenTests/Destroy stack
	//
	// ✓ xxx       Suite.Destroy
	// ✗ …/tdsuite run.func1
	// ✗ …/tdsuite run
	// ✗ …/tdsuite Run
	// ✓ xxx       TestSuiteDestroy
	if !s.Match(-1, tdsuitePkg) &&
		s.Match(-2, tdsuitePkg, "Run") {
		for i := len(s) - 3; i >= 0; i-- {
			if !s.Match(i, tdsuitePkg) {
				s[i+1] = s[len(s)-1]
				return s[:i+2]
			}
		}
		return s[:1]
	}

	return s
}

func formatError(t TestingT, isFatal bool, err *ctxerr.Error, args ...interface{}) {
	t.Helper()

	const failedTest = "Failed test"

	args = flat.Interfaces(args...)

	var buf bytes.Buffer
	color.AppendTestNameOn(&buf)
	if len(args) == 0 {
		buf.WriteString(failedTest)
	} else {
		buf.WriteString(failedTest + " '")
		tdutil.FbuildTestName(&buf, args...)
		buf.WriteString("'")
	}
	color.AppendTestNameOff(&buf)
	buf.WriteString("\n")

	err.Append(&buf, "")

	// Stask trace
	if s := stripTrace(trace.Retrieve(0, "testing.tRunner")); len(s) > 1 {
		buf.WriteString("\nThis is how we got here:\n")

		fnMaxLen := 0
		for _, level := range s {
			if len(level.Func) > fnMaxLen {
				fnMaxLen = len(level.Func)
			}
		}
		fnMaxLen += 2

		nl := ""
		for _, level := range s {
			fmt.Fprintf(&buf, "%s\t%-*s %s", nl, fnMaxLen, level.Func+"()", level.FileLine)
			nl = "\n"
		}
	}

	if isFatal {
		t.Fatal(buf.String())
	} else {
		t.Error(buf.String())
	}
}

func cmpDeeply(ctx ctxerr.Context, t TestingT, got, expected interface{},
	args ...interface{}) bool {
	err := deepValueEqualFinal(ctx,
		reflect.ValueOf(got), reflect.ValueOf(expected))
	if err == nil {
		return true
	}

	t.Helper()
	formatError(t, ctx.FailureIsFatal, err, args...)
	return false
}

// Cmp returns true if "got" matches "expected". "expected" can
// be the same type as "got" is, or contains some TestDeep
// operators. If "got" does not match "expected", it returns false and
// the reason of failure is logged with the help of "t" Error()
// method.
//
//   got := "foobar"
//   td.Cmp(t, got, "foobar")            // succeeds
//   td.Cmp(t, got, td.HasPrefix("foo")) // succeeds
//
// "args..." are optional and allow to name the test. This name is
// used in case of failure to qualify the test. If len(args) > 1 and
// the first item of "args" is a string and contains a '%' rune then
// fmt.Fprintf is used to compose the name, else "args" are passed to
// fmt.Fprint. Do not forget it is the name of the test, not the
// reason of a potential failure.
func Cmp(t TestingT, got, expected interface{}, args ...interface{}) bool {
	t.Helper()
	return cmpDeeply(newContext(), t, got, expected, args...)
}

// CmpDeeply works the same as Cmp and is still available for
// compatibility purpose. Use shorter Cmp in new code.
//
//   got := "foobar"
//   td.CmpDeeply(t, got, "foobar")            // succeeds
//   td.CmpDeeply(t, got, td.HasPrefix("foo")) // succeeds
func CmpDeeply(t TestingT, got, expected interface{}, args ...interface{}) bool {
	t.Helper()
	return cmpDeeply(newContext(), t, got, expected, args...)
}
