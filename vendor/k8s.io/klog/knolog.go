package klog

import "flag"

func InitFlags(flagset *flag.FlagSet) {}

type Verbose bool

type Level int32

var verbosity Level

func V(level Level) Verbose                                { return verbosity >= level }
func (v Verbose) Info(args ...interface{})                 {}
func (v Verbose) Infoln(args ...interface{})               {}
func (v Verbose) Infof(format string, args ...interface{}) {}

func Info(args ...interface{})                    {}
func InfoDepth(depth int, args ...interface{})    {}
func Infoln(args ...interface{})                  {}
func Infof(format string, args ...interface{})    {}
func Warning(args ...interface{})                 {}
func WarningDepth(depth int, args ...interface{}) {}
func Warningln(args ...interface{})               {}
func Warningf(format string, args ...interface{}) {}
func Error(args ...interface{})                   {}
func ErrorDepth(depth int, args ...interface{})   {}
func Errorln(args ...interface{})                 {}
func Errorf(format string, args ...interface{})   {}
func Fatal(args ...interface{})                   {}
func FatalDepth(depth int, args ...interface{})   {}
func Fatalln(args ...interface{})                 {}
func Fatalf(format string, args ...interface{})   {}
func Exit(args ...interface{})                    {}
func ExitDepth(depth int, args ...interface{})    {}
func Exitln(args ...interface{})                  {}
func Exitf(format string, args ...interface{})    {}
