package main

import (
	"fmt"
)

func main() {
	states := []string{"Alabama", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Idaho", "Illinois"}
	for i := 0; i < len(states); i++ {
		fmt.Printf("index : %v   value:%v \n", i, states[i])
	}
}
