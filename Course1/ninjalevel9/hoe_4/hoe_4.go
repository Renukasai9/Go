package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	incrementer := 0
	a := 100
	wg.Add(a)
	var n sync.Mutex
	for i := 0; i < a; i++ {
		go func() {
			n.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			n.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", incrementer)
}
