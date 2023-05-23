package main

import (
	"fmt"
)

// using slice to store and print values
func main() {
	arr := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i, v := range arr {
		fmt.Printf("index : %v and value : %v\n", i, v)
	}
	fmt.Printf("type of the array is:%T \n", arr)
}
