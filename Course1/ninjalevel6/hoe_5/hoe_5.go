package main

import (
	"fmt"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}
func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func main() {
	s1 := square{
		side: 10,
	}
	c1 := circle{
		radius: 2.5,
	}
	fmt.Println("Area of the square is:", s1.area())
	fmt.Println("Area of the circle is:", c1.area())
}
