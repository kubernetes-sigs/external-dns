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

// Package binarylog implementation binary logging as defined in
// https://github.com/grpc/proposal/blob/master/A16-binary-logging.md.
package binarylog

import (
	"fmt"
	"os"

	"google.golang.org/grpc/grpclog"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
	"google.golang.org/grpc/internal/grpcutil"
)

var grpclogLogger = grpclog.Component("binarylog")

// Logger specifies MethodLoggers for method names with a Log call that
// takes a context.
//
// This is used in the 1.0 release of gcp/observability, and thus must not be
// deleted or changed.
type Logger interface {
	GetMethodLogger(methodName string) MethodLogger
}

// binLogger is the global binary logger for the binary. One of this should be
// built at init time from the configuration (environment variable or flags).
//
// It is used to get a MethodLogger for each individual method.
var binLogger Logger

// SetLogger sets the binary logger.
//
// Only call this at init time.
func SetLogger(l Logger) {
	binLogger = l
}

// GetLogger gets the binary logger.
//
// Only call this at init time.
func GetLogger() Logger {
	return binLogger
}

// GetMethodLogger returns the MethodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each MethodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func GetMethodLogger(methodName string) MethodLogger {
	if binLogger == nil {
		return nil
	}
	return binLogger.GetMethodLogger(methodName)
}

func init() {
	const envStr = "GRPC_BINARY_LOG_FILTER"
	configStr := os.Getenv(envStr)
	binLogger = NewLoggerFromConfigString(configStr)
}

// MethodLoggerConfig contains the setting for logging behavior of a method
// logger. Currently, it contains the max length of header and message.
type MethodLoggerConfig struct {
	// Max length of header and message.
	Header, Message uint64
}

// LoggerConfig contains the config for loggers to create method loggers.
type LoggerConfig struct {
	All      *MethodLoggerConfig
	Services map[string]*MethodLoggerConfig
	Methods  map[string]*MethodLoggerConfig

	Blacklist map[string]struct{}
}

type logger struct {
	config LoggerConfig
}

// NewLoggerFromConfig builds a logger with the given LoggerConfig.
func NewLoggerFromConfig(config LoggerConfig) Logger {
	return &logger{config: config}
}

// newEmptyLogger creates an empty logger. The map fields need to be filled in
// using the set* functions.
func newEmptyLogger() *logger {
	return &logger{}
}

// Set method logger for "*".
func (l *logger) setDefaultMethodLogger(ml *MethodLoggerConfig) error {
	if l.config.All != nil {
		return fmt.Errorf("conflicting global rules found")
	}
	l.config.All = ml
	return nil
}

// Set method logger for "service/*".
//
// New MethodLogger with same service overrides the old one.
func (l *logger) setServiceMethodLogger(service string, ml *MethodLoggerConfig) error {
	if _, ok := l.config.Services[service]; ok {
		return fmt.Errorf("conflicting service rules for service %v found", service)
	}
	if l.config.Services == nil {
		l.config.Services = make(map[string]*MethodLoggerConfig)
	}
	l.config.Services[service] = ml
	return nil
}

// Set method logger for "service/method".
//
// New MethodLogger with same method overrides the old one.
func (l *logger) setMethodMethodLogger(method string, ml *MethodLoggerConfig) error {
	if _, ok := l.config.Blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.config.Methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.config.Methods == nil {
		l.config.Methods = make(map[string]*MethodLoggerConfig)
	}
	l.config.Methods[method] = ml
	return nil
}

// Set blacklist method for "-service/method".
func (l *logger) setBlacklist(method string) error {
	if _, ok := l.config.Blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.config.Methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.config.Blacklist == nil {
		l.config.Blacklist = make(map[string]struct{})
	}
	l.config.Blacklist[method] = struct{}{}
	return nil
}

// getMethodLogger returns the MethodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each MethodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func (l *logger) GetMethodLogger(methodName string) MethodLogger {
	s, m, err := grpcutil.ParseMethod(methodName)
	if err != nil {
		grpclogLogger.Infof("binarylogging: failed to parse %q: %v", methodName, err)
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 5ce8c7613 (update vendored files)
=======
	"google.golang.org/grpc/internal/grpcutil"
>>>>>>> 5ce8c7613 (update vendored files)
)

// Logger is the global binary logger. It can be used to get binary logger for
// each method.
type Logger interface {
	getMethodLogger(methodName string) *MethodLogger
}

// binLogger is the global binary logger for the binary. One of this should be
// built at init time from the configuration (environment variable or flags).
//
// It is used to get a methodLogger for each individual method.
var binLogger Logger

var grpclogLogger = grpclog.Component("binarylog")

// SetLogger sets the binarg logger.
//
// Only call this at init time.
func SetLogger(l Logger) {
	binLogger = l
}

// GetMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func GetMethodLogger(methodName string) *MethodLogger {
	if binLogger == nil {
		return nil
	}
	return binLogger.getMethodLogger(methodName)
}

func init() {
	const envStr = "GRPC_BINARY_LOG_FILTER"
	configStr := os.Getenv(envStr)
	binLogger = NewLoggerFromConfigString(configStr)
}

type methodLoggerConfig struct {
	// Max length of header and message.
	hdr, msg uint64
}

type logger struct {
	all      *methodLoggerConfig
	services map[string]*methodLoggerConfig
	methods  map[string]*methodLoggerConfig

	blacklist map[string]struct{}
}

// newEmptyLogger creates an empty logger. The map fields need to be filled in
// using the set* functions.
func newEmptyLogger() *logger {
	return &logger{}
}

// Set method logger for "*".
func (l *logger) setDefaultMethodLogger(ml *methodLoggerConfig) error {
	if l.all != nil {
		return fmt.Errorf("conflicting global rules found")
	}
	l.all = ml
	return nil
}

// Set method logger for "service/*".
//
// New methodLogger with same service overrides the old one.
func (l *logger) setServiceMethodLogger(service string, ml *methodLoggerConfig) error {
	if _, ok := l.services[service]; ok {
		return fmt.Errorf("conflicting service rules for service %v found", service)
	}
	if l.services == nil {
		l.services = make(map[string]*methodLoggerConfig)
	}
	l.services[service] = ml
	return nil
}

// Set method logger for "service/method".
//
// New methodLogger with same method overrides the old one.
func (l *logger) setMethodMethodLogger(method string, ml *methodLoggerConfig) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.methods == nil {
		l.methods = make(map[string]*methodLoggerConfig)
	}
	l.methods[method] = ml
	return nil
}

// Set blacklist method for "-service/method".
func (l *logger) setBlacklist(method string) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.blacklist == nil {
		l.blacklist = make(map[string]struct{})
	}
	l.blacklist[method] = struct{}{}
	return nil
}

// getMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func (l *logger) getMethodLogger(methodName string) *MethodLogger {
	s, m, err := grpcutil.ParseMethod(methodName)
	if err != nil {
<<<<<<< HEAD
		grpclog.Infof("binarylogging: failed to parse %q: %v", methodName, err)
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
		grpclog.Infof("binarylogging: failed to parse %q: %v", methodName, err)
=======
		grpclogLogger.Infof("binarylogging: failed to parse %q: %v", methodName, err)
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 6b7ce455e (update vendored files)
=======
	"google.golang.org/grpc/internal/grpcutil"
>>>>>>> 6b7ce455e (update vendored files)
)

// Logger is the global binary logger. It can be used to get binary logger for
// each method.
type Logger interface {
	getMethodLogger(methodName string) *MethodLogger
}

// binLogger is the global binary logger for the binary. One of this should be
// built at init time from the configuration (environment variable or flags).
//
// It is used to get a methodLogger for each individual method.
var binLogger Logger

var grpclogLogger = grpclog.Component("binarylog")

// SetLogger sets the binarg logger.
//
// Only call this at init time.
func SetLogger(l Logger) {
	binLogger = l
}

// GetMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func GetMethodLogger(methodName string) *MethodLogger {
	if binLogger == nil {
		return nil
	}
	return binLogger.getMethodLogger(methodName)
}

func init() {
	const envStr = "GRPC_BINARY_LOG_FILTER"
	configStr := os.Getenv(envStr)
	binLogger = NewLoggerFromConfigString(configStr)
}

type methodLoggerConfig struct {
	// Max length of header and message.
	hdr, msg uint64
}

type logger struct {
	all      *methodLoggerConfig
	services map[string]*methodLoggerConfig
	methods  map[string]*methodLoggerConfig

	blacklist map[string]struct{}
}

// newEmptyLogger creates an empty logger. The map fields need to be filled in
// using the set* functions.
func newEmptyLogger() *logger {
	return &logger{}
}

// Set method logger for "*".
func (l *logger) setDefaultMethodLogger(ml *methodLoggerConfig) error {
	if l.all != nil {
		return fmt.Errorf("conflicting global rules found")
	}
	l.all = ml
	return nil
}

// Set method logger for "service/*".
//
// New methodLogger with same service overrides the old one.
func (l *logger) setServiceMethodLogger(service string, ml *methodLoggerConfig) error {
	if _, ok := l.services[service]; ok {
		return fmt.Errorf("conflicting service rules for service %v found", service)
	}
	if l.services == nil {
		l.services = make(map[string]*methodLoggerConfig)
	}
	l.services[service] = ml
	return nil
}

// Set method logger for "service/method".
//
// New methodLogger with same method overrides the old one.
func (l *logger) setMethodMethodLogger(method string, ml *methodLoggerConfig) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.methods == nil {
		l.methods = make(map[string]*methodLoggerConfig)
	}
	l.methods[method] = ml
	return nil
}

// Set blacklist method for "-service/method".
func (l *logger) setBlacklist(method string) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.blacklist == nil {
		l.blacklist = make(map[string]struct{})
	}
	l.blacklist[method] = struct{}{}
	return nil
}

// getMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func (l *logger) getMethodLogger(methodName string) *MethodLogger {
	s, m, err := grpcutil.ParseMethod(methodName)
	if err != nil {
<<<<<<< HEAD
		grpclog.Infof("binarylogging: failed to parse %q: %v", methodName, err)
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
		grpclog.Infof("binarylogging: failed to parse %q: %v", methodName, err)
=======
		grpclogLogger.Infof("binarylogging: failed to parse %q: %v", methodName, err)
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
||||||| parent of 4d7e5ad26 (update vendored files)
=======
	"google.golang.org/grpc/internal/grpcutil"
>>>>>>> 4d7e5ad26 (update vendored files)
)

// Logger is the global binary logger. It can be used to get binary logger for
// each method.
type Logger interface {
	getMethodLogger(methodName string) *MethodLogger
}

// binLogger is the global binary logger for the binary. One of this should be
// built at init time from the configuration (environment variable or flags).
//
// It is used to get a methodLogger for each individual method.
var binLogger Logger

var grpclogLogger = grpclog.Component("binarylog")

// SetLogger sets the binarg logger.
//
// Only call this at init time.
func SetLogger(l Logger) {
	binLogger = l
}

// GetMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func GetMethodLogger(methodName string) *MethodLogger {
	if binLogger == nil {
		return nil
	}
	return binLogger.getMethodLogger(methodName)
}

func init() {
	const envStr = "GRPC_BINARY_LOG_FILTER"
	configStr := os.Getenv(envStr)
	binLogger = NewLoggerFromConfigString(configStr)
}

type methodLoggerConfig struct {
	// Max length of header and message.
	hdr, msg uint64
}

type logger struct {
	all      *methodLoggerConfig
	services map[string]*methodLoggerConfig
	methods  map[string]*methodLoggerConfig

	blacklist map[string]struct{}
}

// newEmptyLogger creates an empty logger. The map fields need to be filled in
// using the set* functions.
func newEmptyLogger() *logger {
	return &logger{}
}

// Set method logger for "*".
func (l *logger) setDefaultMethodLogger(ml *methodLoggerConfig) error {
	if l.all != nil {
		return fmt.Errorf("conflicting global rules found")
	}
	l.all = ml
	return nil
}

// Set method logger for "service/*".
//
// New methodLogger with same service overrides the old one.
func (l *logger) setServiceMethodLogger(service string, ml *methodLoggerConfig) error {
	if _, ok := l.services[service]; ok {
		return fmt.Errorf("conflicting service rules for service %v found", service)
	}
	if l.services == nil {
		l.services = make(map[string]*methodLoggerConfig)
	}
	l.services[service] = ml
	return nil
}

// Set method logger for "service/method".
//
// New methodLogger with same method overrides the old one.
func (l *logger) setMethodMethodLogger(method string, ml *methodLoggerConfig) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.methods == nil {
		l.methods = make(map[string]*methodLoggerConfig)
	}
	l.methods[method] = ml
	return nil
}

// Set blacklist method for "-service/method".
func (l *logger) setBlacklist(method string) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.blacklist == nil {
		l.blacklist = make(map[string]struct{})
	}
	l.blacklist[method] = struct{}{}
	return nil
}

// getMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func (l *logger) getMethodLogger(methodName string) *MethodLogger {
	s, m, err := grpcutil.ParseMethod(methodName)
	if err != nil {
<<<<<<< HEAD
		grpclog.Infof("binarylogging: failed to parse %q: %v", methodName, err)
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
		grpclog.Infof("binarylogging: failed to parse %q: %v", methodName, err)
=======
		grpclogLogger.Infof("binarylogging: failed to parse %q: %v", methodName, err)
>>>>>>> 4d7e5ad26 (update vendored files)
		return nil
	}
	if ml, ok := l.config.Methods[s+"/"+m]; ok {
		return NewTruncatingMethodLogger(ml.Header, ml.Message)
	}
	if _, ok := l.config.Blacklist[s+"/"+m]; ok {
		return nil
	}
	if ml, ok := l.config.Services[s]; ok {
		return NewTruncatingMethodLogger(ml.Header, ml.Message)
	}
	if l.config.All == nil {
		return nil
	}
	return NewTruncatingMethodLogger(l.config.All.Header, l.config.All.Message)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
)

// Logger is the global binary logger. It can be used to get binary logger for
// each method.
type Logger interface {
	getMethodLogger(methodName string) *MethodLogger
}

// binLogger is the global binary logger for the binary. One of this should be
// built at init time from the configuration (environment variable or flags).
//
// It is used to get a methodLogger for each individual method.
var binLogger Logger

// SetLogger sets the binarg logger.
//
// Only call this at init time.
func SetLogger(l Logger) {
	binLogger = l
}

// GetMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func GetMethodLogger(methodName string) *MethodLogger {
	if binLogger == nil {
		return nil
	}
	return binLogger.getMethodLogger(methodName)
}

func init() {
	const envStr = "GRPC_BINARY_LOG_FILTER"
	configStr := os.Getenv(envStr)
	binLogger = NewLoggerFromConfigString(configStr)
}

type methodLoggerConfig struct {
	// Max length of header and message.
	hdr, msg uint64
}

type logger struct {
	all      *methodLoggerConfig
	services map[string]*methodLoggerConfig
	methods  map[string]*methodLoggerConfig

	blacklist map[string]struct{}
}

// newEmptyLogger creates an empty logger. The map fields need to be filled in
// using the set* functions.
func newEmptyLogger() *logger {
	return &logger{}
}

// Set method logger for "*".
func (l *logger) setDefaultMethodLogger(ml *methodLoggerConfig) error {
	if l.all != nil {
		return fmt.Errorf("conflicting global rules found")
	}
	l.all = ml
	return nil
}

// Set method logger for "service/*".
//
// New methodLogger with same service overrides the old one.
func (l *logger) setServiceMethodLogger(service string, ml *methodLoggerConfig) error {
	if _, ok := l.services[service]; ok {
		return fmt.Errorf("conflicting service rules for service %v found", service)
	}
	if l.services == nil {
		l.services = make(map[string]*methodLoggerConfig)
	}
	l.services[service] = ml
	return nil
}

// Set method logger for "service/method".
//
// New methodLogger with same method overrides the old one.
func (l *logger) setMethodMethodLogger(method string, ml *methodLoggerConfig) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.methods == nil {
		l.methods = make(map[string]*methodLoggerConfig)
	}
	l.methods[method] = ml
	return nil
}

// Set blacklist method for "-service/method".
func (l *logger) setBlacklist(method string) error {
	if _, ok := l.blacklist[method]; ok {
		return fmt.Errorf("conflicting blacklist rules for method %v found", method)
	}
	if _, ok := l.methods[method]; ok {
		return fmt.Errorf("conflicting method rules for method %v found", method)
	}
	if l.blacklist == nil {
		l.blacklist = make(map[string]struct{})
	}
	l.blacklist[method] = struct{}{}
	return nil
}

// getMethodLogger returns the methodLogger for the given methodName.
//
// methodName should be in the format of "/service/method".
//
// Each methodLogger returned by this method is a new instance. This is to
// generate sequence id within the call.
func (l *logger) getMethodLogger(methodName string) *MethodLogger {
	s, m, err := parseMethodName(methodName)
	if err != nil {
		grpclog.Infof("binarylogging: failed to parse %q: %v", methodName, err)
		return nil
	}
	if ml, ok := l.methods[s+"/"+m]; ok {
		return newMethodLogger(ml.hdr, ml.msg)
	}
	if _, ok := l.blacklist[s+"/"+m]; ok {
		return nil
	}
	if ml, ok := l.services[s]; ok {
		return newMethodLogger(ml.hdr, ml.msg)
	}
	if l.all == nil {
		return nil
	}
	return newMethodLogger(l.all.hdr, l.all.msg)
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
