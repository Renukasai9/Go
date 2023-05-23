package main

import (
	"fmt"
)

// Printing the remainder which is found for each number between 10 to 100 when it is divided by 4
func main() {
	r := 4
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v when divided by %v gives remainder %v \n", i, r, i%4)
	}
}
