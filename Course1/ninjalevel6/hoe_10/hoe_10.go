package main

import "fmt"

// program to demonstrate the closure(when we have enclosed the scope of a variable in some code block)
func main() {
	func() {
		x := 10
		fmt.Println(x)
	}()
	//we cannot access variable x because of its block scope
	//fmt.Println(x)
}
