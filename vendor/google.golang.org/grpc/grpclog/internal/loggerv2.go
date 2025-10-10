/*
 *
 * Copyright 2024 gRPC authors.
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

package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

<<<<<<< HEAD:vendor/google.golang.org/grpc/internal/grpclog/grpclog.go
// Logger is the logger used for the non-depth log functions.
var Logger LoggerV2

// DepthLogger is the logger used for the depth log functions.
var DepthLogger DepthLoggerV2

// InfoDepth logs to the INFO log at the specified depth.
<<<<<<< HEAD
func InfoDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.InfoDepth(depth, args...)
	} else {
		Logger.Infoln(args...)
	}
}

// WarningDepth logs to the WARNING log at the specified depth.
func WarningDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.WarningDepth(depth, args...)
	} else {
		Logger.Warningln(args...)
	}
}

// ErrorDepth logs to the ERROR log at the specified depth.
func ErrorDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.ErrorDepth(depth, args...)
	} else {
		Logger.Errorln(args...)
	}
}

// FatalDepth logs to the FATAL log at the specified depth.
func FatalDepth(depth int, args ...interface{}) {
	if DepthLogger != nil {
		DepthLogger.FatalDepth(depth, args...)
	} else {
		Logger.Fatalln(args...)
	}
	os.Exit(1)
}

// LoggerV2 does underlying logging work for grpclog.
// This is a copy of the LoggerV2 defined in the external grpclog package. It
// is defined here to avoid a circular dependency.
type LoggerV2 interface {
	// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Info(args ...interface{})
	// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
	Infoln(args ...interface{})
	// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...interface{})
	// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
	Warning(args ...interface{})
	// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
	Warningln(args ...interface{})
	// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
	Warningf(format string, args ...interface{})
	// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	Error(args ...interface{})
	// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
	Errorln(args ...interface{})
	// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...interface{})
	// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatal(args ...interface{})
	// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalln(args ...interface{})
	// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalf(format string, args ...interface{})
	// V reports whether verbosity level l is at least the requested verbose level.
	V(l int) bool
}

// DepthLoggerV2 logs at a specified call frame. If a LoggerV2 also implements
// DepthLoggerV2, the below functions will be called with the appropriate stack
// depth set for trivial functions the logger may ignore.
// This is a copy of the DepthLoggerV2 defined in the external grpclog package.
// It is defined here to avoid a circular dependency.
//
// # Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
type DepthLoggerV2 interface {
<<<<<<< HEAD
<<<<<<< HEAD
	// InfoDepth logs to INFO log at the specified depth. Arguments are handled in the manner of fmt.Print.
	InfoDepth(depth int, args ...interface{})
	// WarningDepth logs to WARNING log at the specified depth. Arguments are handled in the manner of fmt.Print.
	WarningDepth(depth int, args ...interface{})
	// ErrorDetph logs to ERROR log at the specified depth. Arguments are handled in the manner of fmt.Print.
	ErrorDepth(depth int, args ...interface{})
	// FatalDepth logs to FATAL log at the specified depth. Arguments are handled in the manner of fmt.Print.
||||||| parent of 6b7ce455e (update vendored files)
=======
	// InfoDepth logs to INFO log at the specified depth. Arguments are handled in the manner of fmt.Println.
	InfoDepth(depth int, args ...interface{})
	// WarningDepth logs to WARNING log at the specified depth. Arguments are handled in the manner of fmt.Println.
	WarningDepth(depth int, args ...interface{})
	// ErrorDepth logs to ERROR log at the specified depth. Arguments are handled in the manner of fmt.Println.
	ErrorDepth(depth int, args ...interface{})
	// FatalDepth logs to FATAL log at the specified depth. Arguments are handled in the manner of fmt.Println.
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	// InfoDepth logs to INFO log at the specified depth. Arguments are handled in the manner of fmt.Println.
	InfoDepth(depth int, args ...interface{})
	// WarningDepth logs to WARNING log at the specified depth. Arguments are handled in the manner of fmt.Println.
	WarningDepth(depth int, args ...interface{})
	// ErrorDepth logs to ERROR log at the specified depth. Arguments are handled in the manner of fmt.Println.
	ErrorDepth(depth int, args ...interface{})
	// FatalDepth logs to FATAL log at the specified depth. Arguments are handled in the manner of fmt.Println.
>>>>>>> 4d7e5ad26 (update vendored files)
	FatalDepth(depth int, args ...interface{})
||||||| parent of d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
=======
func InfoDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.InfoDepth(depth, args...)
	} else {
		Logger.Infoln(args...)
	}
}

// WarningDepth logs to the WARNING log at the specified depth.
func WarningDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.WarningDepth(depth, args...)
	} else {
		Logger.Warningln(args...)
	}
}

// ErrorDepth logs to the ERROR log at the specified depth.
func ErrorDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.ErrorDepth(depth, args...)
	} else {
		Logger.Errorln(args...)
	}
}

// FatalDepth logs to the FATAL log at the specified depth.
func FatalDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.FatalDepth(depth, args...)
	} else {
		Logger.Fatalln(args...)
	}
	os.Exit(1)
}

||||||| parent of c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules):vendor/google.golang.org/grpc/internal/grpclog/grpclog.go
// Logger is the logger used for the non-depth log functions.
var Logger LoggerV2

// DepthLogger is the logger used for the depth log functions.
var DepthLogger DepthLoggerV2

// InfoDepth logs to the INFO log at the specified depth.
func InfoDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.InfoDepth(depth, args...)
	} else {
		Logger.Infoln(args...)
	}
}

// WarningDepth logs to the WARNING log at the specified depth.
func WarningDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.WarningDepth(depth, args...)
	} else {
		Logger.Warningln(args...)
	}
}

// ErrorDepth logs to the ERROR log at the specified depth.
func ErrorDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.ErrorDepth(depth, args...)
	} else {
		Logger.Errorln(args...)
	}
}

// FatalDepth logs to the FATAL log at the specified depth.
func FatalDepth(depth int, args ...any) {
	if DepthLogger != nil {
		DepthLogger.FatalDepth(depth, args...)
	} else {
		Logger.Fatalln(args...)
	}
	os.Exit(1)
}

=======
>>>>>>> c5487e6d6 (NE-2142: UPSTREAM: 5739: Bump k8s and controller-runtime modules):vendor/google.golang.org/grpc/grpclog/internal/loggerv2.go
// LoggerV2 does underlying logging work for grpclog.
type LoggerV2 interface {
	// Info logs to INFO log. Arguments are handled in the manner of fmt.Print.
	Info(args ...any)
	// Infoln logs to INFO log. Arguments are handled in the manner of fmt.Println.
	Infoln(args ...any)
	// Infof logs to INFO log. Arguments are handled in the manner of fmt.Printf.
	Infof(format string, args ...any)
	// Warning logs to WARNING log. Arguments are handled in the manner of fmt.Print.
	Warning(args ...any)
	// Warningln logs to WARNING log. Arguments are handled in the manner of fmt.Println.
	Warningln(args ...any)
	// Warningf logs to WARNING log. Arguments are handled in the manner of fmt.Printf.
	Warningf(format string, args ...any)
	// Error logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	Error(args ...any)
	// Errorln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
	Errorln(args ...any)
	// Errorf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	Errorf(format string, args ...any)
	// Fatal logs to ERROR log. Arguments are handled in the manner of fmt.Print.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatal(args ...any)
	// Fatalln logs to ERROR log. Arguments are handled in the manner of fmt.Println.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalln(args ...any)
	// Fatalf logs to ERROR log. Arguments are handled in the manner of fmt.Printf.
	// gRPC ensures that all Fatal logs will exit with os.Exit(1).
	// Implementations may also call os.Exit() with a non-zero exit code.
	Fatalf(format string, args ...any)
	// V reports whether verbosity level l is at least the requested verbose level.
	V(l int) bool
}

// DepthLoggerV2 logs at a specified call frame. If a LoggerV2 also implements
// DepthLoggerV2, the below functions will be called with the appropriate stack
// depth set for trivial functions the logger may ignore.
//
// # Experimental
//
// Notice: This type is EXPERIMENTAL and may be changed or removed in a
// later release.
type DepthLoggerV2 interface {
	LoggerV2
	// InfoDepth logs to INFO log at the specified depth. Arguments are handled in the manner of fmt.Println.
	InfoDepth(depth int, args ...any)
	// WarningDepth logs to WARNING log at the specified depth. Arguments are handled in the manner of fmt.Println.
	WarningDepth(depth int, args ...any)
	// ErrorDepth logs to ERROR log at the specified depth. Arguments are handled in the manner of fmt.Println.
	ErrorDepth(depth int, args ...any)
	// FatalDepth logs to FATAL log at the specified depth. Arguments are handled in the manner of fmt.Println.
	FatalDepth(depth int, args ...any)
>>>>>>> d03b4fbe9 (UPSTREAM: <carry>: update vendored files after rebase to v0.14.2)
}

const (
	// infoLog indicates Info severity.
	infoLog int = iota
	// warningLog indicates Warning severity.
	warningLog
	// errorLog indicates Error severity.
	errorLog
	// fatalLog indicates Fatal severity.
	fatalLog
)

// severityName contains the string representation of each severity.
var severityName = []string{
	infoLog:    "INFO",
	warningLog: "WARNING",
	errorLog:   "ERROR",
	fatalLog:   "FATAL",
}

// loggerT is the default logger used by grpclog.
type loggerT struct {
	m          []*log.Logger
	v          int
	jsonFormat bool
}

func (g *loggerT) output(severity int, s string) {
	sevStr := severityName[severity]
	if !g.jsonFormat {
		g.m[severity].Output(2, fmt.Sprintf("%v: %v", sevStr, s))
		return
	}
	// TODO: we can also include the logging component, but that needs more
	// (API) changes.
	b, _ := json.Marshal(map[string]string{
		"severity": sevStr,
		"message":  s,
	})
	g.m[severity].Output(2, string(b))
}

func (g *loggerT) Info(args ...any) {
	g.output(infoLog, fmt.Sprint(args...))
}

func (g *loggerT) Infoln(args ...any) {
	g.output(infoLog, fmt.Sprintln(args...))
}

func (g *loggerT) Infof(format string, args ...any) {
	g.output(infoLog, fmt.Sprintf(format, args...))
}

func (g *loggerT) Warning(args ...any) {
	g.output(warningLog, fmt.Sprint(args...))
}

func (g *loggerT) Warningln(args ...any) {
	g.output(warningLog, fmt.Sprintln(args...))
}

func (g *loggerT) Warningf(format string, args ...any) {
	g.output(warningLog, fmt.Sprintf(format, args...))
}

func (g *loggerT) Error(args ...any) {
	g.output(errorLog, fmt.Sprint(args...))
}

func (g *loggerT) Errorln(args ...any) {
	g.output(errorLog, fmt.Sprintln(args...))
}

func (g *loggerT) Errorf(format string, args ...any) {
	g.output(errorLog, fmt.Sprintf(format, args...))
}

func (g *loggerT) Fatal(args ...any) {
	g.output(fatalLog, fmt.Sprint(args...))
	os.Exit(1)
}

func (g *loggerT) Fatalln(args ...any) {
	g.output(fatalLog, fmt.Sprintln(args...))
	os.Exit(1)
}

func (g *loggerT) Fatalf(format string, args ...any) {
	g.output(fatalLog, fmt.Sprintf(format, args...))
	os.Exit(1)
}

func (g *loggerT) V(l int) bool {
	return l <= g.v
}

// LoggerV2Config configures the LoggerV2 implementation.
type LoggerV2Config struct {
	// Verbosity sets the verbosity level of the logger.
	Verbosity int
	// FormatJSON controls whether the logger should output logs in JSON format.
	FormatJSON bool
}

// NewLoggerV2 creates a new LoggerV2 instance with the provided configuration.
// The infoW, warningW, and errorW writers are used to write log messages of
// different severity levels.
func NewLoggerV2(infoW, warningW, errorW io.Writer, c LoggerV2Config) LoggerV2 {
	var m []*log.Logger
	flag := log.LstdFlags
	if c.FormatJSON {
		flag = 0
	}
	m = append(m, log.New(infoW, "", flag))
	m = append(m, log.New(io.MultiWriter(infoW, warningW), "", flag))
	ew := io.MultiWriter(infoW, warningW, errorW) // ew will be used for error and fatal.
	m = append(m, log.New(ew, "", flag))
	m = append(m, log.New(ew, "", flag))
	return &loggerT{m: m, v: c.Verbosity, jsonFormat: c.FormatJSON}
}
