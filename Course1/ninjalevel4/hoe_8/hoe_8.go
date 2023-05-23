package main

import (
	"fmt"
)

// range over the records and values in a map
func main() {
	m1 := map[string][]string{"john": []string{"icecream", "chocolate", "cakes"}, "Rani": []string{"biryani", "sweets", "cakes"}, "Raju": []string{"sweets", "icecream", "chocolates"}}
	for i, v := range m1 {
		fmt.Println("record name:", i)
		for j, k := range v {
			fmt.Printf("index : %v and value : %v \n", j, k)
		}
	}
}
