package type_test

import (
	"testing"
)

type MyInt int64

// 不支持隐式类型转换 别名也不可以
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	var c MyInt
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// point 不支持运算
	//aPtr += aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

// string init
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
}
