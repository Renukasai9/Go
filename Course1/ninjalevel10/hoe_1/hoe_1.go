package main

import "fmt"

func main() {
	a := make(chan int)
	go func() {
		a <- 23
	}()
	fmt.Println(<-a)
}
