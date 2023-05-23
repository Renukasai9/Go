package main

import (
	"fmt"
)

// Range over the records and over the data in each record using multidimensional slices
func main() {
	s1 := []string{"James", "Bond", "shaken,not stirred"}
	s2 := []string{"Miss", "Moneypenny", "hello"}
	res := [][]string{s1, s2}
	fmt.Println(res)
	for i, v := range res {
		fmt.Println("record number", i)
		for j, k := range v {
			fmt.Printf("index : %v and value : %v \n", j, k)
		}
	}
}
