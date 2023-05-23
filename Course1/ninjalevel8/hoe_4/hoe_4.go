package main

import (
	"fmt"
	"sort"
)

// using sort
func main() {
	a := []int{4, 2, 5, 6, 1, 0}
	b := []string{"hello", "hii", "bye", "how"}
	fmt.Println("before sorting:", a)
	sort.Ints(a)
	fmt.Println("after sorting:", a)
	fmt.Println("before sorting:", b)
	sort.Strings(b)
	fmt.Println("after sorting:", b)

}
