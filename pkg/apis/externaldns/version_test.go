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

package externaldns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBanner(t *testing.T) {
	// Set variables to known values
	Version = "1.0.0"
	goVersion = "go1.17"
	GitCommit = "49a0c57c7"

	want := Banner()
	assert.Contains(t, want, "GoVersion=go1.17")
	assert.Contains(t, want, "GitCommitShort=49a0c57c7")
	assert.Contains(t, want, "UserAgent=ExternalDNS/1.0.0")
}
