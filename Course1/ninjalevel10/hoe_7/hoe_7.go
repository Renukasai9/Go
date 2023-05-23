package main

import "fmt"

func main() {
	a := make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				a <- i
			}
		}()
	}
	for k := 0; k < 100; k++ {
		fmt.Println(k, <-a)
	}

	fmt.Println("about to exit")
}
