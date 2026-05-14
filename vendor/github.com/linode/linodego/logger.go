package linodego

import (
	"log"
	"os"
)

//nolint:unused
type httpLogger interface {
	Errorf(format string, v ...any)
	Warnf(format string, v ...any)
	Debugf(format string, v ...any)
}

//nolint:unused
type logger struct {
	l *log.Logger
}

//nolint:unused
func createLogger() *logger {
	l := &logger{l: log.New(os.Stderr, "", log.Ldate|log.Lmicroseconds)}
	return l
}

//nolint:unused
var _ httpLogger = (*logger)(nil)

//nolint:unused
func (l *logger) Errorf(format string, v ...any) {
	l.output("ERROR RESTY "+format, v...)
}

//nolint:unused
func (l *logger) Warnf(format string, v ...any) {
	l.output("WARN RESTY "+format, v...)
}

//nolint:unused
func (l *logger) Debugf(format string, v ...any) {
	l.output("DEBUG RESTY "+format, v...)
}

//nolint:unused
func (l *logger) output(format string, v ...any) { //nolint:goprintffuncname
	if len(v) == 0 {
		l.l.Print(format)
		return
	}

	l.l.Printf(format, v...)
}
