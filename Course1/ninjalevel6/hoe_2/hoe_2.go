package main

import (
	"fmt"
)

// unfurling a slice
func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func main() {
	x := []int{10, 20, 30, 40, 50}
	res1 := foo(x...)
	fmt.Println(res1)
	y := []int{1, 2, 3, 4, 5}
	res2 := bar(y)
	fmt.Println(res2)
}
