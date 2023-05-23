package main

import "fmt"

type hello int

var x hello

func main() {
	x = 23
	fmt.Println(x)
}
