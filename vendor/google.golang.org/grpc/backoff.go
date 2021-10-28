/*
 *
 * Copyright 2017 gRPC authors.
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

// See internal/backoff package for the backoff implementation. This file is
// kept for the exported types and API backward compatibility.

package grpc

import (
	"time"

	"google.golang.org/grpc/backoff"
)

// DefaultBackoffConfig uses values specified for backoff in
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//
// Deprecated: use ConnectParams instead. Will be supported throughout 1.x.
var DefaultBackoffConfig = BackoffConfig{
	MaxDelay: 120 * time.Second,
}

// BackoffConfig defines the parameters for the default gRPC backoff strategy.
//
// Deprecated: use ConnectParams instead. Will be supported throughout 1.x.
type BackoffConfig struct {
	// MaxDelay is the upper bound of backoff delay.
	MaxDelay time.Duration
}

// ConnectParams defines the parameters for connecting and retrying. Users are
// encouraged to use this instead of the BackoffConfig type defined above. See
// here for more details:
// https://github.com/grpc/grpc/blob/master/doc/connection-backoff.md.
//
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// This API is EXPERIMENTAL.
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// This API is EXPERIMENTAL.
=======
// Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// This API is EXPERIMENTAL.
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// This API is EXPERIMENTAL.
=======
// Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// This API is EXPERIMENTAL.
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// This API is EXPERIMENTAL.
=======
// Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
>>>>>>> 4d7e5ad26 (update vendored files)
type ConnectParams struct {
	// Backoff specifies the configuration options for connection backoff.
	Backoff backoff.Config
	// MinConnectTimeout is the minimum amount of time we are willing to give a
	// connection to complete.
	MinConnectTimeout time.Duration
}
