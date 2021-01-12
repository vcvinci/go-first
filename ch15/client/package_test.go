package client

import (
	"fmt"
	"go-first/ch15/series"
	"testing"
)

func init() {
	fmt.Println("init0")
}
func init() {

	fmt.Println("init1")
}
func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacci(10))
}
