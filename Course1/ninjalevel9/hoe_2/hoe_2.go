package main

import (
	"fmt"
)

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("hii")
}
func saysomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		first: "Renuka sai",
	}
	//will not work with saysomething(p1)
	saysomething(&p1)
	p1.speak()
}
