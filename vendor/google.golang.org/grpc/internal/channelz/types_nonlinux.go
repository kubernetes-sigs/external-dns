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

package channelz

import (
	"sync"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
)

var once sync.Once

// SocketOptionData defines the struct to hold socket option data, and related
// getter function to obtain info from fd.
// Windows OS doesn't support Socket Option
type SocketOptionData struct {
}

// Getsockopt defines the function to get socket options requested by channelz.
// It is to be passed to syscall.RawConn.Control().
// Windows OS doesn't support Socket Option
func (s *SocketOptionData) Getsockopt(fd uintptr) {
	once.Do(func() {
		logger.Warning("Channelz: socket options are not supported on non-linux os and appengine.")
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

	"google.golang.org/grpc/grpclog"
||||||| parent of 5ce8c7613 (update vendored files)

	"google.golang.org/grpc/grpclog"
=======
>>>>>>> 5ce8c7613 (update vendored files)
)

var once sync.Once

// SocketOptionData defines the struct to hold socket option data, and related
// getter function to obtain info from fd.
// Windows OS doesn't support Socket Option
type SocketOptionData struct {
}

// Getsockopt defines the function to get socket options requested by channelz.
// It is to be passed to syscall.RawConn.Control().
// Windows OS doesn't support Socket Option
func (s *SocketOptionData) Getsockopt(fd uintptr) {
	once.Do(func() {
<<<<<<< HEAD
		grpclog.Warningln("Channelz: socket options are not supported on non-linux os and appengine.")
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		grpclog.Warningln("Channelz: socket options are not supported on non-linux os and appengine.")
=======
		logger.Warning("Channelz: socket options are not supported on non-linux os and appengine.")
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======

	"google.golang.org/grpc/grpclog"
||||||| parent of 6b7ce455e (update vendored files)

	"google.golang.org/grpc/grpclog"
=======
>>>>>>> 6b7ce455e (update vendored files)
)

var once sync.Once

// SocketOptionData defines the struct to hold socket option data, and related
// getter function to obtain info from fd.
// Windows OS doesn't support Socket Option
type SocketOptionData struct {
}

// Getsockopt defines the function to get socket options requested by channelz.
// It is to be passed to syscall.RawConn.Control().
// Windows OS doesn't support Socket Option
func (s *SocketOptionData) Getsockopt(fd uintptr) {
	once.Do(func() {
<<<<<<< HEAD
		grpclog.Warningln("Channelz: socket options are not supported on non-linux os and appengine.")
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		grpclog.Warningln("Channelz: socket options are not supported on non-linux os and appengine.")
=======
		logger.Warning("Channelz: socket options are not supported on non-linux environments")
>>>>>>> 6b7ce455e (update vendored files)
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

package channelz

import (
	"sync"
)

var once sync.Once

// SocketOptionData defines the struct to hold socket option data, and related
// getter function to obtain info from fd.
// Windows OS doesn't support Socket Option
type SocketOptionData struct {
}

// Getsockopt defines the function to get socket options requested by channelz.
// It is to be passed to syscall.RawConn.Control().
// Windows OS doesn't support Socket Option
func (s *SocketOptionData) Getsockopt(fd uintptr) {
	once.Do(func() {
<<<<<<< HEAD
		grpclog.Warningln("Channelz: socket options are not supported on non-linux os and appengine.")
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		grpclog.Warningln("Channelz: socket options are not supported on non-linux os and appengine.")
=======
		logger.Warning("Channelz: socket options are not supported on non-linux environments")
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// +build !linux appengine

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

package channelz

import (
	"sync"

	"google.golang.org/grpc/grpclog"
)

var once sync.Once

// SocketOptionData defines the struct to hold socket option data, and related
// getter function to obtain info from fd.
// Windows OS doesn't support Socket Option
type SocketOptionData struct {
}

// Getsockopt defines the function to get socket options requested by channelz.
// It is to be passed to syscall.RawConn.Control().
// Windows OS doesn't support Socket Option
func (s *SocketOptionData) Getsockopt(fd uintptr) {
	once.Do(func() {
		grpclog.Warningln("Channelz: socket options are not supported on non-linux os and appengine.")
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
	})
}
