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

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/sirupsen/logrus/hooks/test"
)

// LogsUnderTestWithLogLevel redirects log(s) output to a buffer for testing purposes
//
// Usage: LogsUnderTestWithLogLevel(t)
// Example:
//
//	hook := testutils.LogsUnderTestWithLogLevel(log.DebugLevel, t)
//	... do something that logs ...
//
// testutils.TestHelperLogContains("expected debug log message", hook, t)
func LogsUnderTestWithLogLevel(level log.Level, t *testing.T) *test.Hook {
	t.Helper()
	logger, hook := test.NewNullLogger()

	log.AddHook(hook)
	log.SetOutput(logger.Out)
	log.SetLevel(level)
	return hook
}

// TestHelperLogContains verifies that a specific log message is present
// in the captured log entries. It asserts that the provided message `msg`
// appears in at least one of the log entries recorded by the `hook`.
//
// Parameters:
// - msg: The log message that should be found.
// - hook: The test hook capturing log entries.
// - t: The testing object used for assertions.
//
// Usage:
// This helper is used in tests to ensure that certain log messages are
// logged during the execution of the code under test.
func TestHelperLogContains(msg string, hook *test.Hook, t *testing.T) {
	t.Helper()
	isContains := false
	for _, entry := range hook.AllEntries() {
		if strings.Contains(entry.Message, msg) {
			isContains = true
		}
	}
	assert.True(t, isContains, "Expected log message not found: %s", msg)
}

// TestHelperLogNotContains verifies that a specific log message is not present
// in the captured log entries. It asserts that the provided message `msg`
// does not appear in any of the log entries recorded by the `hook`.
//
// Parameters:
// - msg: The log message that should not be found.
// - hook: The test hook capturing log entries.
// - t: The testing object used for assertions.
//
// Usage:
// This helper is used in tests to ensure that certain log messages are not
// logged during the execution of the code under test.
func TestHelperLogNotContains(msg string, hook *test.Hook, t *testing.T) {
	t.Helper()
	isContains := false
	for _, entry := range hook.AllEntries() {
		if strings.Contains(entry.Message, msg) {
			isContains = true
		}
	}
	assert.False(t, isContains, "Expected log message found when should not: %s", msg)
}

// TestHelperLogContainsWithLogLevel verifies that a specific log message with a given log level
// is present in the captured log entries. It asserts that the provided message `msg`
// appears in at least one of the log entries recorded by the `hook` with the specified log level.
//
// Parameters:
// - msg: The log message that should be found.
// - level: The log level that the message should have.
// - hook: The test hook capturing log entries.
// - t: The testing object used for assertions.
//
// Usage:
// This helper is used in tests to ensure that certain log messages with a specific log level
// are logged during the execution of the code under test.
func TestHelperLogContainsWithLogLevel(msg string, level log.Level, hook *test.Hook, t *testing.T) {
	t.Helper()
	isContains := false
	for _, entry := range hook.AllEntries() {
		if strings.Contains(entry.Message, msg) && entry.Level == level {
			isContains = true
		}
	}
	assert.True(t, isContains, "Expected log message not found: %s with level %s", msg, level)
}
