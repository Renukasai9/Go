package main

import (
	"fmt"
)

// using array to store and print values
func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	for i, v := range arr {
		fmt.Printf("index : %v and value : %v\n", i, v)
	}
	fmt.Printf("type of the array is:%T \n", arr)
}
