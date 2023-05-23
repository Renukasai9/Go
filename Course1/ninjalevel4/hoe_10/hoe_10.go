package main

import (
	"fmt"
)

// delete a record from the map
func main() {
	m1 := map[string][]string{"john": []string{"icecream", "chocolate", "cakes"}, "Rani": []string{"biryani", "sweets", "cakes"}, "Raju": []string{"sweets", "icecream", "chocolates"}}
	m1["jack"] = []string{"cakes", "cookies", "chocolates"}
	fmt.Println("before deleting the map is:")
	for i, v := range m1 {
		fmt.Println("record name:", i)
		for j, k := range v {
			fmt.Printf("index : %v and value : %v \n", j, k)
		}
	}
	fmt.Println()
	fmt.Println("after deleting the map is:")
	delete(m1, "jack")
	for i, v := range m1 {
		fmt.Println("record name:", i)
		for j, k := range v {
			fmt.Printf("index : %v and value : %v \n", j, k)
		}
	}
}
