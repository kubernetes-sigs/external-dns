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

// Package grpclog defines logging for grpc.
//
// In the default logger, severity level can be set by environment variable
// GRPC_GO_LOG_SEVERITY_LEVEL, verbosity level can be set by
// GRPC_GO_LOG_VERBOSITY_LEVEL.
package grpclog

<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
import (
	"os"

	"google.golang.org/grpc/grpclog/internal"
)

func init() {
	SetLoggerV2(newLoggerV2())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return internal.LoggerV2Impl.V(l)
}

// Info logs to the INFO log.
func Info(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
	grpclog.Logger.Infoln(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...interface{}) {
	grpclog.Logger.Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {
	grpclog.Logger.Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...interface{}) {
	grpclog.Logger.Warningln(args...)
}

// Error logs to the ERROR log.
func Error(args ...interface{}) {
	grpclog.Logger.Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	grpclog.Logger.Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
	grpclog.Logger.Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func Fatal(args ...interface{}) {
	grpclog.Logger.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func Fatalf(format string, args ...interface{}) {
	grpclog.Logger.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calle os.Exit()) with exit code 1.
func Fatalln(args ...interface{}) {
	grpclog.Logger.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
func Print(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
func Printf(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
func Println(args ...interface{}) {
	grpclog.Logger.Infoln(args...)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "os"
||||||| parent of 5ce8c7613 (update vendored files)
import "os"
=======
import (
	"os"
>>>>>>> 5ce8c7613 (update vendored files)

	"google.golang.org/grpc/internal/grpclog"
)

func init() {
	SetLoggerV2(newLoggerV2())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return grpclog.Logger.V(l)
}

// Info logs to the INFO log.
func Info(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
	grpclog.Logger.Infoln(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...interface{}) {
	grpclog.Logger.Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {
	grpclog.Logger.Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...interface{}) {
	grpclog.Logger.Warningln(args...)
}

// Error logs to the ERROR log.
func Error(args ...interface{}) {
	grpclog.Logger.Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	grpclog.Logger.Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
	grpclog.Logger.Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func Fatal(args ...interface{}) {
	grpclog.Logger.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func Fatalf(format string, args ...interface{}) {
	grpclog.Logger.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calle os.Exit()) with exit code 1.
func Fatalln(args ...interface{}) {
	grpclog.Logger.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
func Print(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
func Printf(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
func Println(args ...interface{}) {
<<<<<<< HEAD
	logger.Infoln(args...)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	logger.Infoln(args...)
=======
	grpclog.Logger.Infoln(args...)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "os"
||||||| parent of 6b7ce455e (update vendored files)
import "os"
=======
import (
	"os"
>>>>>>> 6b7ce455e (update vendored files)

	"google.golang.org/grpc/internal/grpclog"
)

func init() {
	SetLoggerV2(newLoggerV2())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return grpclog.Logger.V(l)
}

// Info logs to the INFO log.
func Info(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
	grpclog.Logger.Infoln(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...interface{}) {
	grpclog.Logger.Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {
	grpclog.Logger.Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...interface{}) {
	grpclog.Logger.Warningln(args...)
}

// Error logs to the ERROR log.
func Error(args ...interface{}) {
	grpclog.Logger.Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	grpclog.Logger.Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
	grpclog.Logger.Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func Fatal(args ...interface{}) {
	grpclog.Logger.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func Fatalf(format string, args ...interface{}) {
	grpclog.Logger.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calle os.Exit()) with exit code 1.
func Fatalln(args ...interface{}) {
	grpclog.Logger.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
func Print(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
func Printf(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
func Println(args ...interface{}) {
<<<<<<< HEAD
	logger.Infoln(args...)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	logger.Infoln(args...)
=======
	grpclog.Logger.Infoln(args...)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "os"
||||||| parent of 4d7e5ad26 (update vendored files)
import "os"
=======
import (
	"os"
>>>>>>> 4d7e5ad26 (update vendored files)

	"google.golang.org/grpc/internal/grpclog"
)

func init() {
	SetLoggerV2(newLoggerV2())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return grpclog.Logger.V(l)
}

// Info logs to the INFO log.
func Info(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...interface{}) {
	grpclog.Logger.Infoln(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...interface{}) {
	grpclog.Logger.Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...interface{}) {
	grpclog.Logger.Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...interface{}) {
	grpclog.Logger.Warningln(args...)
}

// Error logs to the ERROR log.
func Error(args ...interface{}) {
	grpclog.Logger.Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...interface{}) {
	grpclog.Logger.Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...interface{}) {
	grpclog.Logger.Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func Fatal(args ...interface{}) {
	grpclog.Logger.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func Fatalf(format string, args ...interface{}) {
	grpclog.Logger.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calle os.Exit()) with exit code 1.
func Fatalln(args ...interface{}) {
	grpclog.Logger.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
func Print(args ...interface{}) {
	grpclog.Logger.Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
func Printf(format string, args ...interface{}) {
	grpclog.Logger.Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
func Println(args ...interface{}) {
<<<<<<< HEAD
	logger.Infoln(args...)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	logger.Infoln(args...)
=======
	grpclog.Logger.Infoln(args...)
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
import "os"
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
import "os"
=======
import (
	"os"
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)

	"google.golang.org/grpc/internal/grpclog"
)

func init() {
	SetLoggerV2(newLoggerV2())
}

// V reports whether verbosity level l is at least the requested verbose level.
func V(l int) bool {
	return grpclog.Logger.V(l)
}

// Info logs to the INFO log.
func Info(args ...any) {
	internal.LoggerV2Impl.Info(args...)
}

// Infof logs to the INFO log. Arguments are handled in the manner of fmt.Printf.
func Infof(format string, args ...any) {
	internal.LoggerV2Impl.Infof(format, args...)
}

// Infoln logs to the INFO log. Arguments are handled in the manner of fmt.Println.
func Infoln(args ...any) {
	internal.LoggerV2Impl.Infoln(args...)
}

// Warning logs to the WARNING log.
func Warning(args ...any) {
	internal.LoggerV2Impl.Warning(args...)
}

// Warningf logs to the WARNING log. Arguments are handled in the manner of fmt.Printf.
func Warningf(format string, args ...any) {
	internal.LoggerV2Impl.Warningf(format, args...)
}

// Warningln logs to the WARNING log. Arguments are handled in the manner of fmt.Println.
func Warningln(args ...any) {
	internal.LoggerV2Impl.Warningln(args...)
}

// Error logs to the ERROR log.
func Error(args ...any) {
	internal.LoggerV2Impl.Error(args...)
}

// Errorf logs to the ERROR log. Arguments are handled in the manner of fmt.Printf.
func Errorf(format string, args ...any) {
	internal.LoggerV2Impl.Errorf(format, args...)
}

// Errorln logs to the ERROR log. Arguments are handled in the manner of fmt.Println.
func Errorln(args ...any) {
	internal.LoggerV2Impl.Errorln(args...)
}

// Fatal logs to the FATAL log. Arguments are handled in the manner of fmt.Print.
// It calls os.Exit() with exit code 1.
func Fatal(args ...any) {
	internal.LoggerV2Impl.Fatal(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalf logs to the FATAL log. Arguments are handled in the manner of fmt.Printf.
// It calls os.Exit() with exit code 1.
func Fatalf(format string, args ...any) {
	internal.LoggerV2Impl.Fatalf(format, args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Fatalln logs to the FATAL log. Arguments are handled in the manner of fmt.Println.
// It calls os.Exit() with exit code 1.
func Fatalln(args ...any) {
	internal.LoggerV2Impl.Fatalln(args...)
	// Make sure fatal logs will exit.
	os.Exit(1)
}

// Print prints to the logger. Arguments are handled in the manner of fmt.Print.
//
// Deprecated: use Info.
func Print(args ...any) {
	internal.LoggerV2Impl.Info(args...)
}

// Printf prints to the logger. Arguments are handled in the manner of fmt.Printf.
//
// Deprecated: use Infof.
func Printf(format string, args ...any) {
	internal.LoggerV2Impl.Infof(format, args...)
}

// Println prints to the logger. Arguments are handled in the manner of fmt.Println.
//
// Deprecated: use Infoln.
<<<<<<< HEAD
func Println(args ...interface{}) {
	logger.Infoln(args...)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
func Println(args ...interface{}) {
	logger.Infoln(args...)
=======
func Println(args ...any) {
<<<<<<< HEAD
	grpclog.Logger.Infoln(args...)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
	grpclog.Logger.Infoln(args...)
=======
	internal.LoggerV2Impl.Infoln(args...)
}

// InfoDepth logs to the INFO log at the specified depth.
//
// # Experimental
//
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
// later release.
func InfoDepth(depth int, args ...any) {
	if internal.DepthLoggerV2Impl != nil {
		internal.DepthLoggerV2Impl.InfoDepth(depth, args...)
	} else {
		internal.LoggerV2Impl.Infoln(args...)
	}
}

// WarningDepth logs to the WARNING log at the specified depth.
//
// # Experimental
//
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
// later release.
func WarningDepth(depth int, args ...any) {
	if internal.DepthLoggerV2Impl != nil {
		internal.DepthLoggerV2Impl.WarningDepth(depth, args...)
	} else {
		internal.LoggerV2Impl.Warningln(args...)
	}
}

// ErrorDepth logs to the ERROR log at the specified depth.
//
// # Experimental
//
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
// later release.
func ErrorDepth(depth int, args ...any) {
	if internal.DepthLoggerV2Impl != nil {
		internal.DepthLoggerV2Impl.ErrorDepth(depth, args...)
	} else {
		internal.LoggerV2Impl.Errorln(args...)
	}
}

// FatalDepth logs to the FATAL log at the specified depth.
//
// # Experimental
//
// Notice: This API is EXPERIMENTAL and may be changed or removed in a
// later release.
func FatalDepth(depth int, args ...any) {
	if internal.DepthLoggerV2Impl != nil {
		internal.DepthLoggerV2Impl.FatalDepth(depth, args...)
	} else {
		internal.LoggerV2Impl.Fatalln(args...)
	}
	os.Exit(1)
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules)
}
