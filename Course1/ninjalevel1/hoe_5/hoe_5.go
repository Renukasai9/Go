package main

import "fmt"

type hotdog int

var y hotdog
var x int

func main() {
	fmt.Println(y)
	fmt.Printf("%T", y)
	y = 20
	fmt.Println()
	fmt.Println(y)
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T", x)
	fmt.Println()
}
