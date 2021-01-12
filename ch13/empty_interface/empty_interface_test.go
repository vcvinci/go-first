package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	i, ok := p.(int)
	if ok {
		fmt.Println("Integer", i)
		return
	}
	s, ok := p.(string)
	if ok {
		fmt.Println("string", s)
		return
	}
	fmt.Println("Unknow Type")
}

func DoSomething1(p interface{}) {
	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("Unknow Type")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
	DoSomething(false)
}

func TestEmptyInterfaceAssertion1(t *testing.T) {
	DoSomething1(10)
	DoSomething1("10")
	DoSomething1(false)
}
