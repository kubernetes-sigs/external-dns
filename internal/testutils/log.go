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
	"bytes"
	"flag"
	"testing"

	log "github.com/sirupsen/logrus"
	"k8s.io/klog/v2"
)

// LogsToBuffer redirects log(s) output to a buffer for testing purposes
//
// Usage: LogsToBuffer(t)
// Example:
//
//	buf := LogsToBuffer(log.DebugLevel, t)
//	... do something that logs ...
//	assert.Contains(t, buf.String(), "expected debug log message")
func LogsToBuffer(level log.Level, t *testing.T) *bytes.Buffer {
	t.Helper()
	buf := new(bytes.Buffer)
	log.SetOutput(buf)
	log.SetLevel(level)
	klog.SetOutput(buf)
	flags := &flag.FlagSet{}
	klog.InitFlags(flags)
	// make sure klog doesn't write to stderr by default in tests
	_ = flags.Set("logtostderr", "false")
	_ = flags.Set("alsologtostderr", "false")
	_ = flags.Set("stderrthreshold", "4")
	return buf
}
