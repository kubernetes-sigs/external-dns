/*
 *
 * Copyright 2018 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package envconfig contains grpc settings configured by environment variables.
package envconfig

import (
	"os"
<<<<<<< HEAD
	"strconv"
	"strings"
)

<<<<<<< HEAD
const (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	prefix          = "GRPC_GO_"
	retryStr        = prefix + "RETRY"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
)

||||||| parent of 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0)
const (
	prefix          = "GRPC_GO_"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
)

=======
>>>>>>> 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0)
var (
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
<<<<<<< HEAD
	TXTErrIgnore = !strings.EqualFold(os.Getenv(txtErrIgnoreStr), "false")
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	prefix   = "GRPC_GO_"
	retryStr = prefix + "RETRY"
||||||| parent of 5ce8c7613 (update vendored files)
	prefix   = "GRPC_GO_"
	retryStr = prefix + "RETRY"
=======
	prefix          = "GRPC_GO_"
	retryStr        = prefix + "RETRY"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
>>>>>>> 5ce8c7613 (update vendored files)
)

var (
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
<<<<<<< HEAD
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
=======
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
	TXTErrIgnore = !strings.EqualFold(os.Getenv(txtErrIgnoreStr), "false")
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	prefix   = "GRPC_GO_"
	retryStr = prefix + "RETRY"
||||||| parent of 6b7ce455e (update vendored files)
	prefix   = "GRPC_GO_"
	retryStr = prefix + "RETRY"
=======
	prefix          = "GRPC_GO_"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
>>>>>>> 6b7ce455e (update vendored files)
)

var (
<<<<<<< HEAD
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
=======
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
	TXTErrIgnore = !strings.EqualFold(os.Getenv(txtErrIgnoreStr), "false")
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	prefix   = "GRPC_GO_"
	retryStr = prefix + "RETRY"
||||||| parent of 4d7e5ad26 (update vendored files)
	prefix   = "GRPC_GO_"
	retryStr = prefix + "RETRY"
=======
	prefix          = "GRPC_GO_"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
>>>>>>> 4d7e5ad26 (update vendored files)
)

var (
<<<<<<< HEAD
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
=======
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
	TXTErrIgnore = !strings.EqualFold(os.Getenv(txtErrIgnoreStr), "false")
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0)
	TXTErrIgnore = !strings.EqualFold(os.Getenv(txtErrIgnoreStr), "false")
=======
	TXTErrIgnore = boolFromEnv("GRPC_GO_IGNORE_TXT_ERRORS", true)
	// AdvertiseCompressors is set if registered compressor should be advertised
	// ("GRPC_GO_ADVERTISE_COMPRESSORS" is not "false").
	AdvertiseCompressors = boolFromEnv("GRPC_GO_ADVERTISE_COMPRESSORS", true)
	// RingHashCap indicates the maximum ring size which defaults to 4096
	// entries but may be overridden by setting the environment variable
	// "GRPC_RING_HASH_CAP".  This does not override the default bounds
	// checking which NACKs configs specifying ring sizes > 8*1024*1024 (~8M).
	RingHashCap = uint64FromEnv("GRPC_RING_HASH_CAP", 4096, 1, 8*1024*1024)
	// PickFirstLBConfig is set if we should support configuration of the
	// pick_first LB policy, which can be enabled by setting the environment
	// variable "GRPC_EXPERIMENTAL_PICKFIRST_LB_CONFIG" to "true".
	PickFirstLBConfig = boolFromEnv("GRPC_EXPERIMENTAL_PICKFIRST_LB_CONFIG", false)
>>>>>>> 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0)
)

func boolFromEnv(envVar string, def bool) bool {
	if def {
		// The default is true; return true unless the variable is "false".
		return !strings.EqualFold(os.Getenv(envVar), "false")
	}
	// The default is false; return false unless the variable is "true".
	return strings.EqualFold(os.Getenv(envVar), "true")
}

func uint64FromEnv(envVar string, def, min, max uint64) uint64 {
	v, err := strconv.ParseUint(os.Getenv(envVar), 10, 64)
	if err != nil {
		return def
	}
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"strings"
)

const (
	prefix   = "GRPC_GO_"
	retryStr = prefix + "RETRY"
)

var (
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
