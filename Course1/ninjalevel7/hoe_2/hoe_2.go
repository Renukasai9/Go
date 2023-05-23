package main

import "fmt"

//using *(value) and &(address)

type person struct {
	name string
}

func changeMe(p *person) {
	p.name = "rani"
}
func main() {
	p1 := person{
		name: "raju",
	}
	fmt.Println("current value is:", p1)
	changeMe(&p1)
	fmt.Println("after change:", p1)
}
