package main

import (
	"fmt"
)

// Printing years from my birth to until now using for{...}
func main() {
	birth := 2002
	for {
		if birth <= 2023 {
			fmt.Println(birth)
			birth++
		} else {
			break
		}

	}
}
