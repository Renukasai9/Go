package main

import (
	"fmt"
)

// returning values from functions and printing them
func foo() int {
	return 1
}

func bar() (int, string) {
	return 19, "renuka"
}
func main() {
	a1 := foo()
	b1, c1 := bar()
	fmt.Println(a1, b1, c1)
}
