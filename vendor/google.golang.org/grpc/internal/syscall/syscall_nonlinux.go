<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
//go:build !linux
// +build !linux

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

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Package syscall provides functionalities that grpc uses to get low-level
// operating system stats/info.
package syscall

import (
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/grpclog"
)

var once sync.Once
var logger = grpclog.Component("core")

func log() {
	once.Do(func() {
		logger.Info("CPU time info is unavailable on non-linux or appengine environment.")
	})
}

// GetCPUTime returns the how much CPU time has passed since the start of this process.
// It always returns 0 under non-linux or appengine environment.
func GetCPUTime() int64 {
	log()
	return 0
}

// Rusage is an empty struct under non-linux or appengine environment.
type Rusage struct{}

// GetRusage is a no-op function under non-linux or appengine environment.
func GetRusage() *Rusage {
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
// Package syscall provides functionalities that grpc uses to get low-level
// operating system stats/info.
>>>>>>> 5ce8c7613 (update vendored files)
package syscall

import (
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/grpclog"
)

var once sync.Once
var logger = grpclog.Component("core")

func log() {
	once.Do(func() {
		logger.Info("CPU time info is unavailable on non-linux or appengine environment.")
	})
}

// GetCPUTime returns the how much CPU time has passed since the start of this process.
// It always returns 0 under non-linux or appengine environment.
func GetCPUTime() int64 {
	log()
	return 0
}

// Rusage is an empty struct under non-linux or appengine environment.
type Rusage struct{}

// GetRusage is a no-op function under non-linux or appengine environment.
<<<<<<< HEAD
func GetRusage() (rusage *Rusage) {
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
func GetRusage() (rusage *Rusage) {
=======
func GetRusage() *Rusage {
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
// Package syscall provides functionalities that grpc uses to get low-level
// operating system stats/info.
>>>>>>> 6b7ce455e (update vendored files)
package syscall

import (
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/grpclog"
)

var once sync.Once
var logger = grpclog.Component("core")

func log() {
	once.Do(func() {
		logger.Info("CPU time info is unavailable on non-linux environments.")
	})
}

// GetCPUTime returns the how much CPU time has passed since the start of this
// process. It always returns 0 under non-linux environments.
func GetCPUTime() int64 {
	log()
	return 0
}

// Rusage is an empty struct under non-linux environments.
type Rusage struct{}

<<<<<<< HEAD
// GetRusage is a no-op function under non-linux or appengine environment.
func GetRusage() (rusage *Rusage) {
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// GetRusage is a no-op function under non-linux or appengine environment.
func GetRusage() (rusage *Rusage) {
=======
// GetRusage is a no-op function under non-linux environments.
func GetRusage() *Rusage {
>>>>>>> 6b7ce455e (update vendored files)
	log()
	return nil
}

// CPUTimeDiff returns the differences of user CPU time and system CPU time used
// between two Rusage structs. It a no-op function for non-linux environments.
func CPUTimeDiff(first *Rusage, latest *Rusage) (float64, float64) {
	log()
	return 0, 0
}

// SetTCPUserTimeout is a no-op function under non-linux environments.
func SetTCPUserTimeout(conn net.Conn, timeout time.Duration) error {
	log()
	return nil
}

// GetTCPUserTimeout is a no-op function under non-linux environments.
// A negative return value indicates the operation is not supported
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !linux appengine
||||||| parent of 4d7e5ad26 (update vendored files)
// +build !linux appengine
=======
//go:build !linux
// +build !linux
>>>>>>> 4d7e5ad26 (update vendored files)

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

// Package syscall provides functionalities that grpc uses to get low-level
// operating system stats/info.
package syscall

import (
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/grpclog"
)

var once sync.Once
var logger = grpclog.Component("core")

func log() {
	once.Do(func() {
		logger.Info("CPU time info is unavailable on non-linux environments.")
	})
}

// GetCPUTime returns the how much CPU time has passed since the start of this
// process. It always returns 0 under non-linux environments.
func GetCPUTime() int64 {
	log()
	return 0
}

// Rusage is an empty struct under non-linux environments.
type Rusage struct{}

// GetRusage is a no-op function under non-linux environments.
func GetRusage() *Rusage {
	log()
	return nil
}

// CPUTimeDiff returns the differences of user CPU time and system CPU time used
// between two Rusage structs. It a no-op function for non-linux environments.
func CPUTimeDiff(first *Rusage, latest *Rusage) (float64, float64) {
	log()
	return 0, 0
}

// SetTCPUserTimeout is a no-op function under non-linux environments.
func SetTCPUserTimeout(conn net.Conn, timeout time.Duration) error {
	log()
	return nil
}

<<<<<<< HEAD
// GetTCPUserTimeout is a no-op function under non-linux or appengine environments
// a negative return value indicates the operation is not supported
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
// GetTCPUserTimeout is a no-op function under non-linux or appengine environments
// a negative return value indicates the operation is not supported
=======
// GetTCPUserTimeout is a no-op function under non-linux environments.
// A negative return value indicates the operation is not supported
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !linux appengine
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// +build !linux appengine
=======
//go:build !linux
// +build !linux
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

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

// Package syscall provides functionalities that grpc uses to get low-level
// operating system stats/info.
package syscall

import (
	"net"
	"sync"
	"time"

	"google.golang.org/grpc/grpclog"
)

var once sync.Once
var logger = grpclog.Component("core")

func log() {
	once.Do(func() {
		logger.Info("CPU time info is unavailable on non-linux environments.")
	})
}

// GetCPUTime returns the how much CPU time has passed since the start of this
// process. It always returns 0 under non-linux environments.
func GetCPUTime() int64 {
	log()
	return 0
}

// Rusage is an empty struct under non-linux environments.
type Rusage struct{}

// GetRusage is a no-op function under non-linux environments.
func GetRusage() *Rusage {
	log()
	return nil
}

// CPUTimeDiff returns the differences of user CPU time and system CPU time used
// between two Rusage structs. It a no-op function for non-linux environments.
func CPUTimeDiff(first *Rusage, latest *Rusage) (float64, float64) {
	log()
	return 0, 0
}

// SetTCPUserTimeout is a no-op function under non-linux environments.
func SetTCPUserTimeout(conn net.Conn, timeout time.Duration) error {
	log()
	return nil
}

<<<<<<< HEAD
// GetTCPUserTimeout is a no-op function under non-linux or appengine environments
// a negative return value indicates the operation is not supported
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
// GetTCPUserTimeout is a no-op function under non-linux or appengine environments
// a negative return value indicates the operation is not supported
=======
// GetTCPUserTimeout is a no-op function under non-linux environments.
// A negative return value indicates the operation is not supported
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func GetTCPUserTimeout(conn net.Conn) (int, error) {
	log()
	return -1, nil
}
