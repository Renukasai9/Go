package main

import (
	"fmt"
)

// assigning a function to a variable and calling that variable

func main() {
	a := func() {
		fmt.Println("hello")
	}
	a()
}
