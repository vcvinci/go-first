package operator_test

import "testing"

const (
	Readble = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c:=[...]int{1,2,3,4,5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// 长度不同的数组 比较会compile fail
	//t.Log(a==c)
	t.Log(a == d)

}

// &^ 按位清0
// 1 &^ 0 --- 1
// 1 &^ 1 --- 0
// 0 &^ 1 --- 0
// 0 &^ 0 --- 0
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	a = a &^ Readble
	a = a &^ Executable
	t.Log(a)
	t.Log(a&Readble == Readble, a&Writable == Writable, a&Executable == Executable)
}
