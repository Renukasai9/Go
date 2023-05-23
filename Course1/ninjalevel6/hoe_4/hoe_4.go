package main

import (
	"fmt"
)

// using receiver in function
type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first)
}
func main() {
	p1 := person{
		first: "renuka sai",
		last:  "masina",
		age:   21,
	}
	p1.speak()
}
