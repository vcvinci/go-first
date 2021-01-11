package string_fun_test

import (
	"testing"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))
	//s[1] = '3' // string 是不可变的 byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据
	//s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	s = "中"
	t.Log(len(s)) // 是byte数

	c := []rune(s)

	//t.Log(" rune size :", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

	//var s1 string = "\xE4\xB8\xAD"
	//t.Log(s1)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, i := range s {
		t.Logf("%[1]c %[1]d", i)
	}
}
