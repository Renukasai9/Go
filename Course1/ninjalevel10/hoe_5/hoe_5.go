package main

import "fmt"

func main() {
	a := make(chan int)
	go func() {
		a <- 23
	}()
	v, ok := <-a
	fmt.Println(v, ok)
	close(a)
	v, ok = <-a
	fmt.Println(v, ok)
}
