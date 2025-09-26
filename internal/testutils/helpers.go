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

package testutils

import (
	"strings"
	"testing"
)

// ToPtr returns a pointer to the given value of any type.
// Example usage:
//
//	foo := 42
//	fooPtr := ToPtr(foo)
//	fmt.Println(*fooPtr) // Output: 42
func ToPtr[T any](v T) *T {
	return &v
}

// IsParallel checks if the current test has been marked to run in parallel (i.e. t.Parallel was called).
// It returns true if the test is marked as parallel, false otherwise.
//
// Note: This function uses a deliberate call to t.Setenv that will panic if t.Parallel was previously used;
// the panic is recovered and inspected to determine if parallel execution was attempted.
func IsParallel(t *testing.T) (paralell bool) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			if msg, ok := r.(string); ok && strings.HasPrefix(msg, "testing:") {
				paralell = true
			}
		}
	}()
	t.Setenv(" ", "") // panic if t.Parallel was called
	return
}
