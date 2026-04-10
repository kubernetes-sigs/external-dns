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
<<<<<<< HEAD
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
||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
=======
	// EnableTXTServiceConfig is set if the DNS resolver should perform TXT
	// lookups for service config ("GRPC_ENABLE_TXT_SERVICE_CONFIG" is not
	// "false").
	EnableTXTServiceConfig = boolFromEnv("GRPC_ENABLE_TXT_SERVICE_CONFIG", true)

	// TXTErrIgnore is set if TXT errors should be ignored
	// ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
	TXTErrIgnore = boolFromEnv("GRPC_GO_IGNORE_TXT_ERRORS", true)

	// RingHashCap indicates the maximum ring size which defaults to 4096
	// entries but may be overridden by setting the environment variable
	// "GRPC_RING_HASH_CAP".  This does not override the default bounds
	// checking which NACKs configs specifying ring sizes > 8*1024*1024 (~8M).
	RingHashCap = uint64FromEnv("GRPC_RING_HASH_CAP", 4096, 1, 8*1024*1024)
<<<<<<< HEAD
<<<<<<< HEAD
	// PickFirstLBConfig is set if we should support configuration of the
	// pick_first LB policy, which can be enabled by setting the environment
	// variable "GRPC_EXPERIMENTAL_PICKFIRST_LB_CONFIG" to "true".
	PickFirstLBConfig = boolFromEnv("GRPC_EXPERIMENTAL_PICKFIRST_LB_CONFIG", false)
>>>>>>> 5d0416aaf (UPSTREAM: 3984: CVE-2023-44487 - bump golang.org/x/net v0.17.0)
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	// LeastRequestLB is set if we should support the least_request_experimental
	// LB policy, which can be enabled by setting the environment variable
	// "GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST" to "true".
	LeastRequestLB = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST", false)
	// ALTSMaxConcurrentHandshakes is the maximum number of concurrent ALTS
	// handshakes that can be performed.
	ALTSMaxConcurrentHandshakes = uint64FromEnv("GRPC_ALTS_MAX_CONCURRENT_HANDSHAKES", 100, 1, 100)
=======
	// LeastRequestLB is set if we should support the least_request_experimental
	// LB policy, which can be enabled by setting the environment variable
	// "GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST" to "true".
	LeastRequestLB = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST", false)
||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
	// LeastRequestLB is set if we should support the least_request_experimental
	// LB policy, which can be enabled by setting the environment variable
	// "GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST" to "true".
	LeastRequestLB = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST", false)
=======

>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
	// ALTSMaxConcurrentHandshakes is the maximum number of concurrent ALTS
	// handshakes that can be performed.
	ALTSMaxConcurrentHandshakes = uint64FromEnv("GRPC_ALTS_MAX_CONCURRENT_HANDSHAKES", 100, 1, 100)

	// EnforceALPNEnabled is set if TLS connections to servers with ALPN disabled
	// should be rejected. The HTTP/2 protocol requires ALPN to be enabled, this
	// option is present for backward compatibility. This option may be overridden
	// by setting the environment variable "GRPC_ENFORCE_ALPN_ENABLED" to "true"
	// or "false".
	EnforceALPNEnabled = boolFromEnv("GRPC_ENFORCE_ALPN_ENABLED", true)

	// XDSEndpointHashKeyBackwardCompat controls the parsing of the endpoint hash
	// key from EDS LbEndpoint metadata. Endpoint hash keys can be disabled by
	// setting "GRPC_XDS_ENDPOINT_HASH_KEY_BACKWARD_COMPAT" to "true". A future
	// release will remove this environment variable, enabling the new behavior
	// unconditionally.
	XDSEndpointHashKeyBackwardCompat = boolFromEnv("GRPC_XDS_ENDPOINT_HASH_KEY_BACKWARD_COMPAT", false)

	// RingHashSetRequestHashKey is set if the ring hash balancer can get the
	// request hash header by setting the "requestHashHeader" field, according
	// to gRFC A76. It can be disabled by setting the environment variable
	// "GRPC_EXPERIMENTAL_RING_HASH_SET_REQUEST_HASH_KEY" to "false".
	RingHashSetRequestHashKey = boolFromEnv("GRPC_EXPERIMENTAL_RING_HASH_SET_REQUEST_HASH_KEY", true)

	// ALTSHandshakerKeepaliveParams is set if we should add the
	// KeepaliveParams when dial the ALTS handshaker service.
	ALTSHandshakerKeepaliveParams = boolFromEnv("GRPC_EXPERIMENTAL_ALTS_HANDSHAKER_KEEPALIVE_PARAMS", false)

	// EnableDefaultPortForProxyTarget controls whether the resolver adds a default port 443
	// to a target address that lacks one. This flag only has an effect when all of
	// the following conditions are met:
	//   - A connect proxy is being used.
	//   - Target resolution is disabled.
	//   - The DNS resolver is being used.
	EnableDefaultPortForProxyTarget = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_DEFAULT_PORT_FOR_PROXY_TARGET", true)

	// CaseSensitiveBalancerRegistries is set if the balancer registry should be
	// case-sensitive. This is disabled by default, but can be enabled by setting
	// the env variable "GRPC_GO_EXPERIMENTAL_CASE_SENSITIVE_BALANCER_REGISTRIES"
	// to "true".
<<<<<<< HEAD
	NewPickFirstEnabled = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_NEW_PICK_FIRST", false)
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
||||||| parent of 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
	NewPickFirstEnabled = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_NEW_PICK_FIRST", false)
=======
	//
	// TODO: After 2 releases, we will enable the env var by default.
	CaseSensitiveBalancerRegistries = boolFromEnv("GRPC_GO_EXPERIMENTAL_CASE_SENSITIVE_BALANCER_REGISTRIES", false)

	// XDSAuthorityRewrite indicates whether xDS authority rewriting is enabled.
	// This feature is defined in gRFC A81 and is enabled by setting the
	// environment variable GRPC_EXPERIMENTAL_XDS_AUTHORITY_REWRITE to "true".
	XDSAuthorityRewrite = boolFromEnv("GRPC_EXPERIMENTAL_XDS_AUTHORITY_REWRITE", false)

	// PickFirstWeightedShuffling indicates whether weighted endpoint shuffling
	// is enabled in the pick_first LB policy, as defined in gRFC A113. This
	// feature can be disabled by setting the environment variable
	// GRPC_EXPERIMENTAL_PF_WEIGHTED_SHUFFLING to "false".
	PickFirstWeightedShuffling = boolFromEnv("GRPC_EXPERIMENTAL_PF_WEIGHTED_SHUFFLING", true)

	// XDSRecoverPanicInResourceParsing indicates whether the xdsclient should
	// recover from panics while parsing xDS resources.
	//
	// This feature can be disabled (e.g. for fuzz testing) by setting the
	// environment variable "GRPC_GO_EXPERIMENTAL_XDS_RESOURCE_PANIC_RECOVERY"
	// to "false".
	XDSRecoverPanicInResourceParsing = boolFromEnv("GRPC_GO_EXPERIMENTAL_XDS_RESOURCE_PANIC_RECOVERY", true)

	// DisableStrictPathChecking indicates whether strict path checking is
	// disabled. This feature can be disabled by setting the environment
	// variable GRPC_GO_EXPERIMENTAL_DISABLE_STRICT_PATH_CHECKING to "true".
	//
	// When strict path checking is enabled, gRPC will reject requests with
	// paths that do not conform to the gRPC over HTTP/2 specification found at
	// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-HTTP2.md.
	//
	// When disabled, gRPC will allow paths that do not contain a leading slash.
	// Enabling strict path checking is recommended for security reasons, as it
	// prevents potential path traversal vulnerabilities.
	//
	// A future release will remove this environment variable, enabling strict
	// path checking behavior unconditionally.
	DisableStrictPathChecking = boolFromEnv("GRPC_GO_EXPERIMENTAL_DISABLE_STRICT_PATH_CHECKING", false)

	// EnablePriorityLBChildPolicyCache controls whether the priority balancer
	// should cache child balancers that are removed from the LB policy config,
	// for a period of 15 minutes. This is disabled by default, but can be
	// enabled by setting the env variable
	// GRPC_EXPERIMENTAL_ENABLE_PRIORITY_LB_CHILD_POLICY_CACHE to true.
	EnablePriorityLBChildPolicyCache = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_PRIORITY_LB_CHILD_POLICY_CACHE", false)
>>>>>>> 53ef3ded0 (UPSTREAM: 6362: OCPBUGS-79591: Bump deps to get google.golang.org/grpc v1.80.0)
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
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
	"strconv"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
	"strings"
)

var (
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
	TXTErrIgnore = boolFromEnv("GRPC_GO_IGNORE_TXT_ERRORS", true)
	// AdvertiseCompressors is set if registered compressor should be advertised
	// ("GRPC_GO_ADVERTISE_COMPRESSORS" is not "false").
	AdvertiseCompressors = boolFromEnv("GRPC_GO_ADVERTISE_COMPRESSORS", true)
	// RingHashCap indicates the maximum ring size which defaults to 4096
	// entries but may be overridden by setting the environment variable
	// "GRPC_RING_HASH_CAP".  This does not override the default bounds
	// checking which NACKs configs specifying ring sizes > 8*1024*1024 (~8M).
	RingHashCap = uint64FromEnv("GRPC_RING_HASH_CAP", 4096, 1, 8*1024*1024)
	// LeastRequestLB is set if we should support the least_request_experimental
	// LB policy, which can be enabled by setting the environment variable
	// "GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST" to "true".
	LeastRequestLB = boolFromEnv("GRPC_EXPERIMENTAL_ENABLE_LEAST_REQUEST", false)
	// ALTSMaxConcurrentHandshakes is the maximum number of concurrent ALTS
	// handshakes that can be performed.
	ALTSMaxConcurrentHandshakes = uint64FromEnv("GRPC_ALTS_MAX_CONCURRENT_HANDSHAKES", 100, 1, 100)
)

<<<<<<< HEAD
var (
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
var (
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
)
=======
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
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
