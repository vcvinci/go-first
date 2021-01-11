package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello world")
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
}
