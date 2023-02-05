package error2

// TODO: add backtrace and chain report

import (
	"errors"
	"fmt"
	"strings"

	. "github.com/pkgs-go/fu"
)

func With(e error, f string, a ...interface{}) error {
	f = strings.Replace(f, "{error?}", e.Error(), 1)
	return &errorStr{fmt.Sprintf(f, a...), e, caller(1)}
}

func Errorf(f string, a ...interface{}) error {
	return &errorStr{fmt.Sprintf(f, a...), nil, caller(1)}
}

func New(msg string, skip ...int) error {
	s := Fnz(skip...)
	return &errorStr{msg, nil, caller(s + 1)}
}

func (e *errorStr) Unwrap() error {
	return e.error
}

func (e *errorStr) Is(err error) bool {
	return errors.Is(e.error, err)
}
