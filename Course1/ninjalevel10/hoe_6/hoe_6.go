package main

import "fmt"

func main() {
	a := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			a <- i
		}
		close(a)
	}()
	for v := range a {
		fmt.Println(v)
	}
	fmt.Println("about to exit")
}
