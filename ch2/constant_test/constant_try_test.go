package constant_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
)

const (
	Readble = 1 << iota
	Writable
	Executable
)

func TestConstanTry(t *testing.T) {
	t.Log(Monday, Tuesday)
}

func TestConstanTry1(t *testing.T) {
	a := 7
	t.Log(Readble, Writable, Executable)
	t.Log(a&Readble == Readble, a&Writable == Writable, a&Executable == Executable)
}
