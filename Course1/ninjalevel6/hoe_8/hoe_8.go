package main

import (
	"fmt"
)

// calling the returned function
func speak() func() int {
	return func() int {
		return 23
	}
}
func main() {
	a := speak()
	fmt.Println("the number is:", a())
}
