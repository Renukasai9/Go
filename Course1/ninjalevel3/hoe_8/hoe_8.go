package main

import (
	"fmt"
)

// program with no switch expression specified
func main() {
	switch {
	case 3 == 0:
		fmt.Println("condition 1")
	case 2 == 2:
		fmt.Println("condition 2")
	}
}
