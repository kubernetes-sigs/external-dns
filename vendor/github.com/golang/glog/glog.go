package glog

import (
	"flag"
	"io/ioutil"
	"strconv"
	"sync"

	"github.com/sirupsen/logrus"
)

var logger logrus.FieldLogger = &logrus.Logger{
	Out:   ioutil.Discard,
	Level: logrus.PanicLevel,
}
var mu = sync.Mutex{}

func SetLogger(l logrus.FieldLogger) {
	mu.Lock()
	logger = l
	mu.Unlock()
}

type Level int32

// Set is part of the flag.Value interface.
func (l *Level) Set(value string) error {
	v, err := strconv.Atoi(value)
	if err != nil {
		return err
	}

	*l = Level(v)
	return nil
}

// String is part of the flag.Value interface.
func (l *Level) String() string {
	return strconv.FormatInt(int64(*l), 10)
}

type Verbose bool

var verbosity Level

// init replicates glog's verbosity level functionality,
// allowing us to show and hide high-verbosity-level
// messages from various kubernetes components.
func init() {
	flag.Var(&verbosity, "v", "log level for V logs")
}

func V(level Level) Verbose {
	return verbosity >= level
}

func (v Verbose) Info(args ...interface{}) {
	if v {
		logger.WithField("func", "Info").Debug(args...)
	}
}

func (v Verbose) Infoln(args ...interface{}) {
	if v {
		logger.WithField("func", "Infoln").Debugln(args...)
	}
}

func (v Verbose) Infof(format string, args ...interface{}) {
	if v {
		logger.WithField("func", "Infof").Debugf(format, args...)
	}
}

func Info(args ...interface{}) {
	logger.WithField("func", "Info").Info(args...)
}

func InfoDepth(depth int, args ...interface{}) {
	logger.WithField("func", "InfoDepth").Info(args...)
}

func Infoln(args ...interface{}) {
	logger.WithField("func", "Infoln").Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	logger.WithField("func", "Infof").Infof(format, args...)
}

func Warning(args ...interface{}) {
	logger.WithField("func", "Warning").Warn(args...)
}

func WarningDepth(depth int, args ...interface{}) {
	logger.WithField("func", "WarningDepth").Warn(args...)
}

func Warningln(args ...interface{}) {
	logger.WithField("func", "Warningln").Warnln(args...)
}

func Warningf(format string, args ...interface{}) {
	logger.WithField("func", "Warningf").Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.WithField("func", "Error").Error(args...)
}

func ErrorDepth(depth int, args ...interface{}) {
	logger.WithField("func", "ErrorDepth").Error(args...)
}

func Errorln(args ...interface{}) {
	logger.WithField("func", "Errorln").Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.WithField("func", "Errorf").Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.WithField("func", "Fatal").Fatal(args...)
}

func FatalDepth(depth int, args ...interface{}) {
	logger.WithField("func", "FatalDepth").Fatal(args...)
}

func Fatalln(args ...interface{}) {
	logger.WithField("func", "Fatalln").Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.WithField("func", "Fatalf").Fatalf(format, args...)
}

func Exit(args ...interface{}) {
	logger.WithField("func", "Exit").Fatal(args...)
}

func ExitDepth(depth int, args ...interface{}) {
	logger.WithField("func", "ExitDepth").Fatal(args...)
}

func Exitln(args ...interface{}) {
	logger.WithField("func", "Exitln").Fatalln(args...)
}

func Exitf(format string, args ...interface{}) {
	logger.WithField("func", "Exitf").Fatalf(format, args...)
}
