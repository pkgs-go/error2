package error2

import (
	"fmt"
	"runtime"
)

type frame [3]uintptr

type errorStr struct {
	message string
	error   error
	frame   frame
}

func caller(skip int) (r frame) {
	runtime.Callers(skip+1, r[:])
	return
}

func (f frame) location() (function, file string, line int) {
	frames := runtime.CallersFrames(f[:])
	if _, ok := frames.Next(); !ok {
		return "", "", 0
	}
	fr, ok := frames.Next()
	if !ok {
		return "", "", 0
	}
	return fr.Function, fr.File, fr.Line
}

func (e *errorStr) Error() string {
	return e.message
}

func (e *errorStr) String() string {
	return e.message + " [" + e.frame.String() + "]"
}

type ErrorStr interface {
	error
	String() string
	Location() (function, file string, line int)
}

func (e *errorStr) Location() (function, file string, line int) {
	return e.frame.location()
}

func (f frame) String() (str string) {
	function, file, line := f.location()
	str = function
	if str == "" {
		str = "<unknown func>"
	}
	if file != "" {
		str += fmt.Sprintf(" %s:%d\n", file, line)
	}
	return
}
