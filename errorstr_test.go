package error2

import (
	"fmt"
	"testing"
)

func f() error {
	e := New("text")
	return With(e, "chained error of {error?}")
}

func Test_errorStr_1(t *testing.T) {
	w := f()
	fmt.Println(MessageOf(w))
	fmt.Println(w.(*errorStr).String())
}
