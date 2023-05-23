package main

import "fmt"

// callback
func foo(f func(xi []int) int, ii []int) int {
	res := f(ii)
	res++
	return res
}
func main() {
	g := func(xi []int) int {
		if len(xi)%2 == 0 {
			return 10
		} else {
			return 20
		}
	}
	x := foo(g, []int{10, 20, 30, 40, 50})
	fmt.Println(x)
}
