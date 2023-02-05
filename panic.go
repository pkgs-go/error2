package error2

import (
	"errors"
	"fmt"
)

/*
PanicMessage returns a message from the panic object
*/
func MessageOf(e interface{}) string {
	if p, ok := e.(error); ok {
		return p.Error()
	}
	return fmt.Sprint(e)
}

/*
PanicError returns an error from the panic object
*/
func ErrorOf(e interface{}) error {
	if p, ok := e.(error); ok {
		return p
	}
	return errors.New(fmt.Sprint(e))
}
