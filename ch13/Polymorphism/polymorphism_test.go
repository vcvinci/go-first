package Polymorphism

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (p *GoProgrammer) WriteHelloWorld() Code {
	return "go hello world"
}

type JavaProgrammer struct {
}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "java hello world"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolymorphism(t *testing.T) {
	goProgrammer := new(GoProgrammer)
	JavaProgrammer := new(JavaProgrammer)
	writeFirstProgram(goProgrammer)
	writeFirstProgram(JavaProgrammer)
}
