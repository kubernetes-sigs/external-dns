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
	"strings"
)

const (
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	prefix          = "GRPC_GO_"
	retryStr        = prefix + "RETRY"
	txtErrIgnoreStr = prefix + "IGNORE_TXT_ERRORS"
)

var (
	// Retry is set if retry is explicitly enabled via "GRPC_GO_RETRY=on".
	Retry = strings.EqualFold(os.Getenv(retryStr), "on")
	// TXTErrIgnore is set if TXT errors should be ignored ("GRPC_GO_IGNORE_TXT_ERRORS" is not "false").
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
)
