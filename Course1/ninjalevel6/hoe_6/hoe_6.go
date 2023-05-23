package main

import (
	"fmt"
)

// using anonymous function
func main() {
	func() {
		fmt.Println("hello")
	}()
	func(x int) {
		fmt.Println("number is", x)
	}(23)
}
