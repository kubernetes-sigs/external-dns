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

// Package version reports the build identity of external-dns.
// Stdlib-only to keep the import lightweight for webhook providers.
package version

import (
	"fmt"
	"runtime"
)

const (
	bannerTemplate = `GitCommitShort=%s, GoVersion=%s, Platform=%s, UserAgent=%s`
)

var (
	// Version is the release version, set via -ldflags -X.
	Version = "unknown"
	// GitCommit is the short commit hash, set via -ldflags -X.
	GitCommit = "unknown"
	// UserAgentProduct prefixes UserAgent; webhook providers may override it.
	UserAgentProduct = "ExternalDNS"
	goVersion        = runtime.Version()
)

// UserAgent returns the HTTP User-Agent, "<UserAgentProduct>/<Version>".
func UserAgent() string {
	return fmt.Sprintf("%s/%s", UserAgentProduct, Version)
}

// Banner returns a one-line summary of the build for startup logs.
func Banner() string {
	platform := fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
	return fmt.Sprintf(
		bannerTemplate,
		GitCommit,
		goVersion,
		platform,
		UserAgent(),
	)
}
