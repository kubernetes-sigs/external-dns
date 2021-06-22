// +build !appengine,!js,windows

package logrus

import (
	"io"
	"os"
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD

	"golang.org/x/sys/windows"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		handle := windows.Handle(v.Fd())
		var mode uint32
		if err := windows.GetConsoleMode(handle, &mode); err != nil {
			return false
		}
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		if err := windows.SetConsoleMode(handle, mode); err != nil {
			return false
		}
		return true
	}
	return false
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"syscall"
||||||| parent of 5ce8c7613 (update vendored files)
	"syscall"
=======
>>>>>>> 5ce8c7613 (update vendored files)

	"golang.org/x/sys/windows"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		handle := windows.Handle(v.Fd())
		var mode uint32
		if err := windows.GetConsoleMode(handle, &mode); err != nil {
			return false
		}
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		if err := windows.SetConsoleMode(handle, mode); err != nil {
			return false
		}
		return true
	}
<<<<<<< HEAD
	if ret {
		initTerminal(w)
	}
	return ret
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
	if ret {
		initTerminal(w)
	}
	return ret
=======
	return false
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"syscall"
||||||| parent of 6b7ce455e (update vendored files)
	"syscall"
=======
>>>>>>> 6b7ce455e (update vendored files)

	"golang.org/x/sys/windows"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		handle := windows.Handle(v.Fd())
		var mode uint32
		if err := windows.GetConsoleMode(handle, &mode); err != nil {
			return false
		}
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		if err := windows.SetConsoleMode(handle, mode); err != nil {
			return false
		}
		return true
	}
<<<<<<< HEAD
	if ret {
		initTerminal(w)
	}
	return ret
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
	if ret {
		initTerminal(w)
	}
	return ret
=======
	return false
>>>>>>> 6b7ce455e (update vendored files)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"syscall"
||||||| parent of 4d7e5ad26 (update vendored files)
	"syscall"
=======
>>>>>>> 4d7e5ad26 (update vendored files)

	"golang.org/x/sys/windows"
)

func checkIfTerminal(w io.Writer) bool {
	switch v := w.(type) {
	case *os.File:
		handle := windows.Handle(v.Fd())
		var mode uint32
		if err := windows.GetConsoleMode(handle, &mode); err != nil {
			return false
		}
		mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
		if err := windows.SetConsoleMode(handle, mode); err != nil {
			return false
		}
		return true
	}
<<<<<<< HEAD
	if ret {
		initTerminal(w)
	}
	return ret
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4d7e5ad26 (update vendored files)
	if ret {
		initTerminal(w)
	}
	return ret
=======
	return false
>>>>>>> 4d7e5ad26 (update vendored files)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
	"syscall"

	sequences "github.com/konsorten/go-windows-terminal-sequences"
)

func initTerminal(w io.Writer) {
	switch v := w.(type) {
	case *os.File:
		sequences.EnableVirtualTerminalProcessing(syscall.Handle(v.Fd()), true)
	}
}

func checkIfTerminal(w io.Writer) bool {
	var ret bool
	switch v := w.(type) {
	case *os.File:
		var mode uint32
		err := syscall.GetConsoleMode(syscall.Handle(v.Fd()), &mode)
		ret = (err == nil)
	default:
		ret = false
	}
	if ret {
		initTerminal(w)
	}
	return ret
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
}
