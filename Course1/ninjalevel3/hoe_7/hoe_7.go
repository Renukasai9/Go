package main

import (
	"fmt"
)

// program showing if,else if and else statements in action
func main() {
	res := "Renuka"
	if res == "sai" {
		fmt.Println("condition 1")
	} else if res == "Renuka" {
		fmt.Println("condition 2")
	} else {
		fmt.Println("condition 3")
	}
}
