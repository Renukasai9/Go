package main

import (
	"fmt"
	"sync"
)

// using waitgroups
func main() {
	var wg sync.WaitGroup
	wg.Add((2))
	go func() {
		fmt.Println("function 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("function 2")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("completed")
}
