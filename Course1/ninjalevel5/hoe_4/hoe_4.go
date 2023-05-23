package main

import (
	"fmt"
)

// anonymous struct
func main() {
	p1 := struct {
		firstname           string
		lastname            string
		favourite_ice_cream []string
	}{
		firstname:           "renuka sai",
		lastname:            "masina",
		favourite_ice_cream: []string{"vanila", "strawberry"},
	}
	fmt.Println("firstname:", p1.firstname)
	fmt.Println("lastname:", p1.lastname)
	for j, k := range p1.favourite_ice_cream {
		fmt.Printf("index is: %v  value : %v \n", j, k)
	}
}
