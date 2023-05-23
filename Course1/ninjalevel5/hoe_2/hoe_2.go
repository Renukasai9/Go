package main

import (
	"fmt"
)

//creating a map and storing two values of type person and ranging over the values

type person struct {
	firstname           string
	lastname            string
	favourite_ice_cream []string
}

func main() {
	p1 := person{firstname: "renuka sai", lastname: "masina", favourite_ice_cream: []string{"vanila", "strawberry"}}
	p2 := person{firstname: "mickey", lastname: "mouse", favourite_ice_cream: []string{"vanila", "strawberry"}}
	fmt.Println("person 1")
	fmt.Println("firstname:", p1.firstname)
	fmt.Println("lastname:", p1.lastname)
	for j, k := range p1.favourite_ice_cream {
		fmt.Printf("index is: %v  value : %v \n", j, k)
	}
	fmt.Println("person 2")
	fmt.Println("firstname:", p2.firstname)
	fmt.Println("lastname:", p2.lastname)
	for j, k := range p2.favourite_ice_cream {
		fmt.Printf("index is: %v  value : %v \n", j, k)
	}
	fmt.Println()
	m := map[string]person{p1.lastname: p1, p2.lastname: p2}
	for i, v := range m {
		fmt.Println(i)
		fmt.Println("firstname:", v.firstname)
		fmt.Println("lastname:", v.lastname)
		for j, k := range v.favourite_ice_cream {
			fmt.Printf("index : %v and value : %v \n", j, k)
		}
	}

}
