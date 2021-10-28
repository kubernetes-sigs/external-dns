// Copyright (c) 2018, Maxime Soulé
// All rights reserved.
//
// This source code is licensed under the BSD-style license found in the
// LICENSE file in the root directory of this source tree.

package ctxerr

import (
	"github.com/maxatome/go-testdeep/internal/anchors"
	"github.com/maxatome/go-testdeep/internal/hooks"
	"github.com/maxatome/go-testdeep/internal/location"
	"github.com/maxatome/go-testdeep/internal/visited"
)

// Context is used internally to keep track of the CmpDeeply in-Depth
// traversal.
type Context struct {
	Path        Path
	Visited     visited.Visited
	CurOperator location.GetLocationer
	Depth       int
	// 0 ≤ MaxErrors ≤ 1 stops when first error encoutered (without the
	// "Too many errors" error);
	// MaxErrors > 1 stops when MaxErrors'th error encoutered (with a
	// last "Too many errors" error);
	// < 0 do not stop until comparison ends.
	MaxErrors int
	Errors    *[]*Error
	Anchors   *anchors.Info
	Hooks     *hooks.Info
	// If true, the contents of the returned *Error will not be
	// checked. Can be used to avoid filling Error{} with expensive
	// computations.
	BooleanError bool
	// See ContextConfig.FailureIsFatal for details.
	FailureIsFatal bool
	// See ContextConfig.UseEqual for details.
	UseEqual bool
	// See ContextConfig.BeLax for details.
	BeLax bool
	// See ContextConfig.IgnoreUnexported for details.
	IgnoreUnexported bool
}

// InitErrors initializes Context *Errors slice, if MaxErrors < 0 or
// MaxErrors > 1.
func (c *Context) InitErrors() {
	if c.MaxErrors != 0 && c.MaxErrors != 1 {
		var errors []*Error
		c.Errors = &errors
	}
}

// ResetErrors returns a new Context without any Error set.
func (c Context) ResetErrors() (new Context) {
	new = c
	new.InitErrors()
	return
}

// CollectError collects an error in the context. It returns an error
// if the collector is full, nil otherwise.
//
// In boolean context, ignore the passed error and return the BooleanError.
func (c Context) CollectError(err *Error) *Error {
	if err == nil {
		return nil
	}

	// The boolean error must not be altered!
	if c.BooleanError {
		return BooleanError
	}

	// Error context not initialized yet
	if err.Context.Depth == 0 {
		err.Context = c
	}

	if !err.Location.IsInitialized() && c.CurOperator != nil {
		err.Location = c.CurOperator.GetLocation()
	}

	// Stop when first error encoutered
	if c.Errors == nil {
		return err
	}

	// Else, accumulate...
	*c.Errors = append(*c.Errors, err)
	if c.MaxErrors >= 0 && len(*c.Errors) >= c.MaxErrors {
		*c.Errors = append(*c.Errors, ErrTooManyErrors)
		return c.MergeErrors()
	}
	return nil
}

// MergeErrors merges all collected errors in the first one and
// returns it. It returns nil if no errors have been collected.
func (c Context) MergeErrors() *Error {
	if c.Errors == nil || len(*c.Errors) == 0 {
		return nil
	}

	if len(*c.Errors) > 1 {
		for idx, last := 0, len(*c.Errors)-2; idx <= last; idx++ {
			(*c.Errors)[idx].Next = (*c.Errors)[idx+1]
		}
	}
	return (*c.Errors)[0]
}

// CannotCompareError returns a generic error used when the access of
// unexported fields cannot be overridden.
func (c Context) CannotCompareError() *Error {
	if c.BooleanError {
		return BooleanError
	}
	return &Error{
		Message: "cannot compare",
		Summary: NewSummary("unexported field that cannot be overridden"),
	}
}

// AddCustomLevel creates a new Context from current one plus pathAdd.
func (c Context) AddCustomLevel(pathAdd string) (new Context) {
	new = c
	new.Path = new.Path.AddCustomLevel(pathAdd)
	new.Depth++
	return
}

// AddField creates a new Context from current one plus "." + field.
func (c Context) AddField(field string) (new Context) {
	new = c
	new.Path = new.Path.AddField(field)
	new.Depth++
	return
}

// AddArrayIndex creates a new Context from current one plus an array
// dereference for index-th item.
func (c Context) AddArrayIndex(index int) (new Context) {
	new = c
	new.Path = new.Path.AddArrayIndex(index)
	new.Depth++
	return
}

// AddMapKey creates a new Context from current one plus a map
// dereference for key key.
func (c Context) AddMapKey(key interface{}) (new Context) {
	new = c
	new.Path = new.Path.AddMapKey(key)
	new.Depth++
	return
}

// AddPtr creates a new Context from current one plus a pointer dereference.
func (c Context) AddPtr(num int) (new Context) {
	new = c
	new.Path = new.Path.AddPtr(num)
	new.Depth++
	return
}

// AddFunctionCall creates a new Context from current one inside a
// function call.
func (c Context) AddFunctionCall(fn string) (new Context) {
	new = c
	new.Path = new.Path.AddFunctionCall(fn)
	new.Depth++
	return
}

// ResetPath creates a new Context from current one but reinitializing Path.
func (c Context) ResetPath(newRoot string) (new Context) {
	new = c
	new.Path = NewPath(newRoot)
	new.Depth++
	return
}
