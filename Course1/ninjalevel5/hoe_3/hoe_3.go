package main

import (
	"fmt"
)

//creating a embedded structs and ranging over the values

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{vehicle: vehicle{doors: 4, color: "white"}, fourwheel: true}
	s1 := sedan{vehicle: vehicle{doors: 2, color: "black"}, luxury: false}
	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(t1.doors)
	fmt.Println(s1.doors)

}
