package main

import (
	"fmt"
)

// using defer keyword in functions
func foo() {
	fmt.Println("function 1")
}

func bar() {
	fmt.Println("function 2")
}
func main() {
	defer foo()
	bar()
}
