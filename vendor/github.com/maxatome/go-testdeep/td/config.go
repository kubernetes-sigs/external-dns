// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package td

import (
	"os"
	"strconv"
	"testing"

	"github.com/maxatome/go-testdeep/internal/anchors"
	"github.com/maxatome/go-testdeep/internal/ctxerr"
	"github.com/maxatome/go-testdeep/internal/hooks"
	"github.com/maxatome/go-testdeep/internal/visited"
)

// ContextConfig allows to configure finely how tests failures are rendered.
//
// See [NewT] function to use it.
type ContextConfig struct {
	// RootName is the string used to represent the root of got data. It
	// defaults to "DATA". For an HTTP response body, it could be "BODY"
	// for example.
	RootName      string
	forkedFromCtx *ctxerr.Context
	// MaxErrors is the maximal number of errors to dump in case of Cmp*
	// failure.
	//
	// It defaults to 10 except if the environment variable
	// TESTDEEP_MAX_ERRORS is set. In this latter case, the
	// TESTDEEP_MAX_ERRORS value is converted to an int and used as is.
	//
	// Setting it to 0 has the same effect as 1: only the first error
	// will be dumped without the "Too many errors" error.
	//
	// Setting it to a negative number means no limit: all errors
	// will be dumped.
	MaxErrors int
	anchors   *anchors.Info
	hooks     *hooks.Info
	// FailureIsFatal allows to Fatal() (instead of Error()) when a test
	// fails. Using *testing.T or *testing.B instance as t.TB value, FailNow()
	// is called behind the scenes when Fatal() is called. See testing
	// documentation for details.
	FailureIsFatal bool
	// UseEqual allows to use the Equal method on got (if it exists) or
	// on any of its component to compare got and expected values.
	//
	// The signature of the Equal method should be:
	//   (A) Equal(B) bool
	// with B assignable to A.
	//
	// See time.Time as an example of accepted Equal() method.
	//
	// See (*T).UseEqual method to only apply this property to some
	// specific types.
	UseEqual bool
	// BeLax allows to compare different but convertible types. If set
	// to false (default), got and expected types must be the same. If
	// set to true and expected type is convertible to got one, expected
	// is first converted to go type before its comparison. See CmpLax
	// function/method and Lax operator to set this flag without
	// providing a specific configuration.
	BeLax bool
	// IgnoreUnexported allows to ignore unexported struct fields. Be
	// careful about structs entirely composed of unexported fields
	// (like time.Time for example). With this flag set to true, they
	// are all equal. In such case it is advised to set UseEqual flag,
	// to use (*T).UseEqual method or to add a Cmp hook using
	// (*T).WithCmpHooks method.
	//
	// See (*T).IgnoreUnexported method to only apply this property to some
	// specific types.
	IgnoreUnexported bool
	// TestDeepInGotOK allows to accept TestDeep operator in got Cmp*
	// parameter. By default it is forbidden and a panic occurs, because
	// most of the time it is a mistake to compare (expected, got)
	// instead of official (got, expected).
	TestDeepInGotOK bool
}

// Equal returns true if both c and o are equal. Only public fields
// are taken into account to check equality.
func (c ContextConfig) Equal(o ContextConfig) bool {
	return c.RootName == o.RootName &&
		c.MaxErrors == o.MaxErrors &&
		c.FailureIsFatal == o.FailureIsFatal &&
		c.UseEqual == o.UseEqual &&
		c.BeLax == o.BeLax &&
		c.IgnoreUnexported == o.IgnoreUnexported &&
		c.TestDeepInGotOK == o.TestDeepInGotOK
}

// OriginalPath returns the current path when the [ContextConfig] has
// been built. It always returns ContextConfig.RootName except if c
// has been built by [Code] operator. See [Code] documentation for an
// example of use.
func (c ContextConfig) OriginalPath() string {
	if c.forkedFromCtx == nil {
		return c.RootName
	}
	return c.forkedFromCtx.Path.String()
}

const (
	contextDefaultRootName = "DATA"
	contextPanicRootName   = "FUNCTION"
	envMaxErrors           = "TESTDEEP_MAX_ERRORS"
)

func getMaxErrorsFromEnv() int {
	env := os.Getenv(envMaxErrors)
	if env != "" {
		n, err := strconv.Atoi(env)
		if err == nil {
			return n
		}
	}
	return 10
}

// DefaultContextConfig is the default configuration used to render
// tests failures. If overridden, new settings will impact all Cmp*
// functions and [*T] methods (if not specifically configured.)
var DefaultContextConfig = ContextConfig{
	RootName:         contextDefaultRootName,
	MaxErrors:        getMaxErrorsFromEnv(),
	FailureIsFatal:   false,
	UseEqual:         false,
	BeLax:            false,
	IgnoreUnexported: false,
	TestDeepInGotOK:  false,
}

func (c *ContextConfig) sanitize() {
	if c.RootName == "" {
		c.RootName = DefaultContextConfig.RootName
	}
	if c.MaxErrors == 0 {
		c.MaxErrors = DefaultContextConfig.MaxErrors
	}
}

// newContext creates a new ctxerr.Context using DefaultContextConfig
// configuration.
func newContext(t TestingT) ctxerr.Context {
	if tt, ok := t.(*T); ok {
		return newContextWithConfig(tt, tt.Config)
	}
	tb, _ := t.(testing.TB)
	return newContextWithConfig(tb, DefaultContextConfig)
}

// newContextWithConfig creates a new ctxerr.Context using a specific
// configuration.
func newContextWithConfig(tb testing.TB, config ContextConfig) (ctx ctxerr.Context) {
	config.sanitize()

	ctx = ctxerr.Context{
		Path:             ctxerr.NewPath(config.RootName),
		Visited:          visited.NewVisited(),
		MaxErrors:        config.MaxErrors,
		Anchors:          config.anchors,
		Hooks:            config.hooks,
		OriginalTB:       tb,
		FailureIsFatal:   config.FailureIsFatal,
		UseEqual:         config.UseEqual,
		BeLax:            config.BeLax,
		IgnoreUnexported: config.IgnoreUnexported,
		TestDeepInGotOK:  config.TestDeepInGotOK,
	}

	ctx.InitErrors()
	return
}

// newBooleanContext creates a new boolean ctxerr.Context.
func newBooleanContext() ctxerr.Context {
	return ctxerr.Context{
		Visited:          visited.NewVisited(),
		BooleanError:     true,
		UseEqual:         DefaultContextConfig.UseEqual,
		BeLax:            DefaultContextConfig.BeLax,
		IgnoreUnexported: DefaultContextConfig.IgnoreUnexported,
		TestDeepInGotOK:  DefaultContextConfig.TestDeepInGotOK,
	}
}
