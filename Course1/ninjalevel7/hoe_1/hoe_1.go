package main

import "fmt"

// usage of ampersand(&,to find address)
func main() {
	a := 23
	fmt.Println("value of a :", a)
	fmt.Println("address of a :", &a)
}
