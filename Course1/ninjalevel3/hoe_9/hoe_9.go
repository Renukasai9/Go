package main

import (
	"fmt"
)

// program with switch expression specified
func main() {
	favsport := "throw ball"
	switch favsport {
	case "cricket":
		fmt.Println("I love cricket")
	case "chess":
		fmt.Println("I love chess")
	case "throw ball":
		fmt.Println("I love throw ball")

	}
}
