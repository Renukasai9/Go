package main

import (
	"fmt"
)

func main() {
	a := 20
	fmt.Printf("%d \t %b \t %x\t", a, a, a)
	fmt.Println()
	b := a << 1
	fmt.Printf("%d \t %b \t %x\t", b, b, b)

}
